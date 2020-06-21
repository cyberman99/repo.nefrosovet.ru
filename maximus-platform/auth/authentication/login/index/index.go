package index

import (
	"errors"

	"repo.nefrosovet.ru/maximus-platform/auth/storage"

	"github.com/go-openapi/strfmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	indexClient "repo.nefrosovet.ru/maximus-platform/auth/index/client"
	"repo.nefrosovet.ru/maximus-platform/auth/index/client/auth"
	"repo.nefrosovet.ru/maximus-platform/auth/index/client/employees"
	"repo.nefrosovet.ru/maximus-platform/auth/index/client/patients"
	"repo.nefrosovet.ru/maximus-platform/auth/index/client/search"
	st "repo.nefrosovet.ru/maximus-platform/auth/storage/storage_accessor"
)

// UserParams - params for employee or patient sync
type UserParams struct {
	Username        string `json:"username,omitempty"`
	Email           string `json:"email,omitempty"`
	FirstName       string `json:"firstName,omitempty"`
	LastName        string `json:"lastName,omitempty"`
	Patronymic      string `json:"patronymic,omitempty"`
	Mobile          string `json:"mobile,omitempty"`
	SmartCardNumber string `json:"smartCardNumber,omitempty"`
}

type Credentials struct {
	UserType string

	Login    string
	Password string

	SmartCardNumber string
}

type Result struct {
	EntityID string

	Error error
}

func Auth(credentials *Credentials) *Result {
	if credentials.SmartCardNumber != "" {
		return authBySmartCard(credentials)
	}

	client := getIndexClient()
	var indexResponse *auth.PostAuthOK
	var err error
	var userType string

	for _, userType = range []string{storage.RoleDefaultEmployee, storage.RoleDefaultPatient} {
		params := auth.NewPostAuthParams().WithBody(auth.PostAuthBody{})
		params.Body.Login = credentials.Login
		params.Body.Password = credentials.Password
		params.Body.Type = userType

		indexResponse, err = client.Auth.PostAuth(params)
		// success
		if err == nil {
			break
		}

		switch err := err.(type) {
		case *auth.PostAuthUnauthorized:
			logrus.WithFields(logrus.Fields{
				"context":  "CORE",
				"resource": "index",
				"function": "AuthByIndex",
				"type":     userType,
				"login":    credentials.Login,
				"status":   "FAILED",
				"error":    err.Error(),
			}).Debug("login on index failed")
		default:
			logrus.WithFields(logrus.Fields{
				"context":  "CORE",
				"resource": "index",
				"addr":     viper.GetString("index.http.host") + "/" + viper.GetString("index.http.path"),
				"status":   "FAILED",
			}).Error("index request failed")
			logrus.Debug(err)
		}

	}

	if err != nil {
		return &Result{
			Error: err,
		}
	}

	id := indexResponse.Payload.Data[0].GUID
	if id == "" {
		return &Result{
			Error: errors.New("no user GUID returned from index"),
		}
	}

	us := st.GetStorage().UserStorage
	users, err := us.Get(storage.GetUser{
		ID: &id,
	})
	if err != nil {
		return &Result{
			Error: err,
		}
	}

	if len(users) == 0 {
		in := storage.StoreUser{
			User: storage.User{
				ID: id,
				Roles: map[string]bool{
					userType: true,
				},
			},
		}

		if user, err := us.Store(in); err != nil {
			return &Result{
				Error: err,
			}
		} else {
			return &Result{
				EntityID: user.ID,
			}
		}
	}

	return &Result{
		EntityID: users[0].ID,
	}
}

func authBySmartCard(credentials *Credentials) *Result {
	var err error
	var id string

	switch credentials.UserType {
	case storage.RoleDefaultPatient:
		id, err = SearchPatient(UserParams{SmartCardNumber: credentials.SmartCardNumber})
	case storage.RoleDefaultEmployee:
		id, err = SearchEmployee(UserParams{SmartCardNumber: credentials.SmartCardNumber})
	case "":
		id, err = SearchEmployee(UserParams{SmartCardNumber: credentials.SmartCardNumber})
		if id == "" {
			id, err = SearchPatient(UserParams{SmartCardNumber: credentials.SmartCardNumber})
		}
	}

	if id == "" {
		return &Result{
			Error: errors.New("no user GUID returned from index"),
		}
	}

	// Already in local DB.
	us := st.GetStorage().UserStorage
	users, err := us.Get(storage.GetUser{
		ID: &id,
	})
	if err != nil {
		return &Result{
			EntityID: id,
			Error:    err,
		}
	}

	if len(users) == 0 {
		in := storage.StoreUser{
			User: storage.User{
				ID:    id,
				Roles: map[string]bool{},
			},
		}

		switch credentials.UserType {
		case storage.RoleDefaultPatient:
			in.Roles[storage.RoleDefaultPatient] = true
		case storage.RoleDefaultEmployee:
			in.Roles[storage.RoleDefaultEmployee] = true
		}

		if user, err := us.Store(in); err != nil {
			return &Result{
				EntityID: id,
				Error:    err,
			}
		} else {
			return &Result{
				EntityID: user.ID,
			}
		}
	}

	return &Result{
		EntityID: users[0].ID,
	}
}

// Search searches index user by GUID
func Search(id string) (*search.GetSearchOK, error) {

	client := getIndexClient()
	params := search.NewGetSearchParams()
	params.GUID = id

	indexResponse, err := client.Search.GetSearch(params)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"context":  "CORE",
			"resource": "index",
			"function": "Search",
			"addr":     viper.GetString("index.http.host") + "/" + viper.GetString("index.http.path"),
			"status":   "FAILED",
		}).Error("index request failed")
		logrus.WithError(err).Debug("index response")

		return nil, errors.New(ErrorTypeInternal)
	}

	return indexResponse, err
}

// CreateOrSearchEmployee creates employee
func CreateOrSearchEmployee(employeeParams UserParams) (string, error) {

	client := getIndexClient()
	params := employees.NewPostEmployeesParams().WithBody(employees.PostEmployeesBody{})
	params.Body.Username = employeeParams.Username
	params.Body.Email = employeeParams.Email
	params.Body.FirstName = employeeParams.FirstName
	params.Body.LastName = employeeParams.LastName
	params.Body.Patronymic = employeeParams.Patronymic
	params.Body.Mobile = employeeParams.Mobile

	indexResponse, err := client.Employees.PostEmployees(params)

	if err != nil {
		switch e := err.(type) {
		// Validation error: try to parse error and search user.
		case *employees.PostEmployeesBadRequest:
			logrus.WithFields(logrus.Fields{
				"context":  "CORE",
				"resource": "index",
				"function": "CreateOrSearchEmployee",
				"status":   "FAILED",
				"errors":   e.Payload.Errors,
			}).Debug("can't create new employee")

			if e.Payload.Errors != nil {
				errorMap := e.Payload.Errors.(map[string]interface{})

				if errorMap["validation"] != nil {
					validation := errorMap["validation"].(map[string]interface{})
					searchEmployeeParams := UserParams{}
					var needSearch bool

					if validation["email"] != nil && validation["email"].(string) == ErrorTypeUnique {
						searchEmployeeParams.Email = employeeParams.Email
						needSearch = true
					}
					if validation["username"] != nil && validation["username"].(string) == ErrorTypeUnique {
						searchEmployeeParams.Username = employeeParams.Username
						needSearch = true
					}
					if validation["mobile"] != nil && validation["mobile"].(string) == ErrorTypeUnique {
						searchEmployeeParams.Mobile = employeeParams.Mobile
						needSearch = true
					}

					if needSearch {
						return SearchEmployee(searchEmployeeParams)
					}
				}
			}

			return "", errors.New(ErrorTypeInternal)
		default:
			logrus.WithFields(logrus.Fields{
				"context":  "CORE",
				"resource": "index",
				"function": "CreateOrSearchEmployee",
				"addr":     viper.GetString("index.http.host") + "/" + viper.GetString("index.http.path"),
				"status":   "FAILED",
			}).Error("index request failed")
			logrus.Debug(err)

			return "", errors.New(ErrorTypeInternal)
		}
	}

	return indexResponse.Payload.Data[0].GUID, nil
}

// SearchEmployee searches employee
func SearchEmployee(searchParams UserParams) (string, error) {

	client := getIndexClient()
	params := search.NewGetSearchEmployeesParams()
	if searchParams.Username != "" {
		params.Username = &searchParams.Username
	}
	if searchParams.Email != "" {
		params.Email = &searchParams.Email
	}
	if searchParams.Mobile != "" {
		params.Mobile = &searchParams.Mobile
	}
	if searchParams.SmartCardNumber != "" {
		params.SmartCardNumber = &searchParams.SmartCardNumber
	}

	indexResponse, err := client.Search.GetSearchEmployees(params)

	if err != nil || len(indexResponse.Payload.Data) == 0 {
		logrus.WithFields(logrus.Fields{
			"context":  "CORE",
			"resource": "index",
			"function": "SearchEmployee",
			"addr":     viper.GetString("index.http.host") + "/" + viper.GetString("index.http.path"),
			"status":   "FAILED",
			"error":    err,
		}).Debug("index request failed")

		return "", errors.New(ErrorTypeCore)
	}

	return indexResponse.Payload.Data[0].GUID, nil
}

// CreateOrSearchPatient creates patient
func CreateOrSearchPatient(patientParams UserParams) (string, error) {

	client := getIndexClient()
	params := patients.NewPostPatientsParams().WithBody(patients.PostPatientsBody{})
	params.Body.Username = patientParams.Username
	params.Body.Email = patientParams.Email
	params.Body.FirstName = patientParams.FirstName
	params.Body.LastName = patientParams.LastName
	params.Body.Patronymic = patientParams.Patronymic
	params.Body.Mobile = patientParams.Mobile

	indexResponse, err := client.Patients.PostPatients(params)

	if err != nil {
		switch e := err.(type) {
		// Validation error: try to parse error and search user.
		case *patients.PostPatientsBadRequest:
			logrus.WithFields(logrus.Fields{
				"context":  "CORE",
				"resource": "index",
				"function": "CreateOrSearchEmployee",
				"status":   "FAILED",
				"errors":   e.Payload.Errors,
			}).Debug("can't create new employee")

			if e.Payload.Errors != nil {
				errorMap := e.Payload.Errors.(map[string]interface{})

				if errorMap["validation"] != nil {
					validation := errorMap["validation"].(map[string]interface{})
					searchPatientParams := UserParams{}
					var needSearch bool

					if validation["email"] != nil && validation["email"].(string) == ErrorTypeUnique {
						searchPatientParams.Email = patientParams.Email
						needSearch = true
					}
					if validation["username"] != nil && validation["username"].(string) == ErrorTypeUnique {
						searchPatientParams.Username = patientParams.Username
						needSearch = true
					}
					if validation["mobile"] != nil && validation["mobile"].(string) == ErrorTypeUnique {
						searchPatientParams.Mobile = patientParams.Mobile
						needSearch = true
					}

					if needSearch {
						return SearchPatient(searchPatientParams)
					}
				}
			}

			return "", errors.New(ErrorTypeInternal)
		default:
			logrus.WithFields(logrus.Fields{
				"context":  "CORE",
				"resource": "index",
				"addr":     viper.GetString("index.http.host") + "/" + viper.GetString("index.http.path"),
				"status":   "FAILED",
			}).Error("index request failed")
			logrus.Debug(err)

			return "", errors.New(ErrorTypeInternal)
		}
	}

	return indexResponse.Payload.Data[0].GUID, nil
}

// SearchPatient searches patient
func SearchPatient(searchParams UserParams) (string, error) {

	client := getIndexClient()
	params := search.NewGetSearchPatientsParams()
	if searchParams.Username != "" {
		params.Username = &searchParams.Username
	}
	if searchParams.Email != "" {
		params.Email = &searchParams.Email
	}
	if searchParams.Mobile != "" {
		params.Mobile = &searchParams.Mobile
	}
	if searchParams.SmartCardNumber != "" {
		params.SmartCardNumber = &searchParams.SmartCardNumber
	}

	indexResponse, err := client.Search.GetSearchPatients(params)

	if err != nil || len(indexResponse.Payload.Data) == 0 {
		logrus.WithFields(logrus.Fields{
			"context":  "CORE",
			"resource": "index",
			"function": "SearchPatient",
			"addr":     viper.GetString("index.http.host") + "/" + viper.GetString("index.http.path"),
			"status":   "FAILED",
			"error":    err,
		}).Debug("index request failed")

		return "", errors.New(ErrorTypeCore)
	}

	return indexResponse.Payload.Data[0].GUID, nil
}

// PatchEmployee patches employee
func PatchEmployee(guid string, employeeParams UserParams) error {

	if guid == "" {
		return errors.New("no guid given")
	}
	client := getIndexClient()
	params := employees.NewPatchEmployeesEmployeeGUIDParams().WithBody(employees.PatchEmployeesEmployeeGUIDBody{})
	params.EmployeeGUID = guid
	if employeeParams.Username != "" {
		params.Body.Username = employeeParams.Username
	}
	if employeeParams.Email != "" {
		params.Body.Email = employeeParams.Email
	}
	if employeeParams.FirstName != "" {
		params.Body.FirstName = employeeParams.FirstName
	}
	if employeeParams.LastName != "" {
		params.Body.LastName = employeeParams.LastName
	}
	if employeeParams.Patronymic != "" {
		params.Body.Patronymic = employeeParams.Patronymic
	}
	if employeeParams.Mobile != "" {
		params.Body.Mobile = employeeParams.Mobile
	}

	_, err := client.Employees.PatchEmployeesEmployeeGUID(params)

	if err != nil {
		switch e := err.(type) {
		case *employees.PatchEmployeesEmployeeGUIDBadRequest:
			if e.Payload.Errors != nil {
				logrus.WithFields(logrus.Fields{
					"context":  "CORE",
					"resource": "index",
					"function": "PatchEmployee",
					"errors":   e.Payload.Errors,
					"status":   "FAILED",
				}).Error("can't patch employee")
			}
			return errors.New(ErrorTypeUnique)
		default:
			logrus.WithFields(logrus.Fields{
				"context":  "CORE",
				"resource": "index",
				"addr":     viper.GetString("index.http.host") + "/" + viper.GetString("index.http.path"),
				"status":   "FAILED",
			}).Error("index request failed")
			logrus.Debug(err)

			return errors.New(ErrorTypeInternal)
		}
	}

	return nil
}

// PatchPatient patches patient
func PatchPatient(guid string, employeeParams UserParams) error {

	if guid == "" {
		return errors.New("no guid given")
	}
	client := getIndexClient()
	params := patients.NewPatchPatientsPatientGUIDParams().WithBody(patients.PatchPatientsPatientGUIDBody{})
	params.PatientGUID = guid
	if employeeParams.Username != "" {
		params.Body.Username = employeeParams.Username
	}
	if employeeParams.Email != "" {
		params.Body.Email = employeeParams.Email
	}
	if employeeParams.FirstName != "" {
		params.Body.FirstName = employeeParams.FirstName
	}
	if employeeParams.LastName != "" {
		params.Body.LastName = employeeParams.LastName
	}
	if employeeParams.Patronymic != "" {
		params.Body.Patronymic = employeeParams.Patronymic
	}
	if employeeParams.Mobile != "" {
		params.Body.Mobile = employeeParams.Mobile
	}

	_, err := client.Patients.PatchPatientsPatientGUID(params)

	if err != nil {
		switch e := err.(type) {
		case *employees.PatchEmployeesEmployeeGUIDBadRequest:
			if e.Payload.Errors != nil {
				logrus.WithFields(logrus.Fields{
					"context":  "CORE",
					"resource": "index",
					"function": "PatchPatient",
					"errors":   e.Payload.Errors,
					"status":   "FAILED",
				}).Error("can't patch patient")
			}
			return errors.New(ErrorTypeUnique)
		default:
			logrus.WithFields(logrus.Fields{
				"context":  "CORE",
				"resource": "index",
				"addr":     viper.GetString("index.http.host") + "/" + viper.GetString("index.http.path"),
				"status":   "FAILED",
			}).Error("index request failed")
			logrus.Debug(err)

			return errors.New(ErrorTypeInternal)
		}
	}

	return nil
}

func getIndexClient() *indexClient.Index {
	return indexClient.NewHTTPClientWithConfig(strfmt.Default, &indexClient.TransportConfig{
		Host:     viper.GetString("index.http.host"),
		BasePath: viper.GetString("index.http.path"),
		Schemes:  []string{"http"},
	})
}
