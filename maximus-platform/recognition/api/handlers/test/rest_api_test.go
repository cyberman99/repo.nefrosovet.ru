package test

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"repo.nefrosovet.ru/maximus-platform/recognition/api/restapi/operations/photo"
	"repo.nefrosovet.ru/maximus-platform/recognition/api/restapi/operations/recognize"
	"testing"
)

func TestExampleTestSuite(t *testing.T) {
	s := new(RestApiTestingSuit)
	s.InitSuit()
	server := s.InitServer()
	suite.Run(t, s)
	if err := server.Shutdown(); err != nil {
		log.Fatalln(err)
	}
}

func (suite *RestApiTestingSuit) InitSuit() {
	suite.Host = "localhost"
	suite.Port = 9000

	suite.PersonID = []byte("74c194ad-c0f5-420e-b4d5-c4e77131a97a")
	var err error
	suite.Image, err = os.Open("test.jpg")
	if err != nil {
		log.Fatalln(err)
	}
	//defer suite.Image.Close() TODO: implement
}

func versionCheck(as *assert.Assertions, ver string) {
	reg := `[0-9]+[.][0-9]+[.][0-9]+`
	matched, err := regexp.MatchString(reg, ver)
	as.NoError(err, "Error apply regex")
	as.True(matched, fmt.Sprintf("Response contain incorrect formatting version field: \"%s\".\nMatch pattern: %s", ver, reg))
}

func (suite *RestApiTestingSuit) Test01() {
	LogTestCase(
		"post",
		"/photos",
		"usually request",
		"success",
	)
	as := assert.New(suite.T())

	// Forming request
	b, w := createMultipartFormData(suite.PersonID, suite.Image)
	req, err := http.NewRequest(
		"POST",
		suite.GetBaseUrl()+"/photos",
		&b,
	)
	if !as.NoError(err) {
		return
	}
	req.Header.Set("Content-Type", w.FormDataContentType())

	// Sending request and check body existence
	var cli = new(http.Client)
	resp, err := cli.Do(req)
	if !as.NoError(err) || !as.NotNil(resp.Body) {
		return
	}

	// Decode body and check code
	respBt, err := ioutil.ReadAll(resp.Body)
	if !as.NoError(err) {
		return
	}
	as.Equal(200, resp.StatusCode)
	var respBody = new(photo.CreateOKBody)
	err = json.Unmarshal(respBt, respBody)
	if !as.NoError(err) {
		return
	}

	// Body validation
	if as.NotNil(respBody.Version) {
		versionCheck(as, *respBody.Version)
	}

	if as.NotNil(respBody.Message) {
		as.Equal("SUCCESS", *respBody.Message)
	}

	as.Nil(respBody.Errors)

	if as.NotEmpty(respBody.Data) && as.True(len(respBody.Data) == 1) {
		item := respBody.Data[0]

		if as.NotNil(item.ID) {
			suite.ImageID = item.ID
		}

		as.Equal(string(suite.PersonID), item.PersonID.String())

		as.Equal(aws.SDKName, item.ExtService)

		if as.NotNil(item.URL) {
			as.Equal(suite.GetBaseUrl()+"/photos/"+item.ID.String(), *item.URL)
		}
	}
}

func (suite *RestApiTestingSuit) Test02() {
	LogTestCase(
		"post",
		"/photos",
		"with nonexistent personID",
		"success",
	)
	as := assert.New(suite.T())

	// Forming request
	generatedUuid, err := uuid.DefaultGenerator.NewV1()
	if !as.NoError(err) {
		return
	}
	b, w := createMultipartFormData(generatedUuid.Bytes(), suite.Image)
	req, err := http.NewRequest(
		"POST",
		suite.GetBaseUrl()+"/photos",
		&b,
	)
	if !as.NoError(err) {
		return
	}
	req.Header.Set("Content-Type", w.FormDataContentType())

	// Sending request and check body existence
	var cli = new(http.Client)
	resp, err := cli.Do(req)
	if !as.NoError(err) && !as.NotNil(resp.Body) {
		return
	}

	// Body validation
	respBt, err := ioutil.ReadAll(resp.Body)
	if as.NoError(err) {
		return
	}

	as.Equal(404, resp.StatusCode)

	var respBody = new(photo.CreateNotFoundBody)
	err = json.Unmarshal(respBt, respBody)
	if !as.NoError(err) {
		return
	}

	if as.NotNil(respBody.Version) {
		versionCheck(as, *respBody.Version)
	}

	as.Nil(respBody.Errors)

	as.Nil(respBody.Data)

	matched, err := regexp.MatchString(`[0-9A-Za-z] Not found`, respBody.Message)
	if as.NoError(err) {
		as.True(matched)
	}
}

func (suite *RestApiTestingSuit) Test03() {
	LogTestCase(
		"post",
		"/photos",
		"without file",
		"SUCCESS",
	)
	as := assert.New(suite.T())

	// Forming request
	b, w := createMultipartFormData(suite.PersonID, suite.Image)
	req, err := http.NewRequest(
		"POST",
		suite.GetBaseUrl()+"/photos",
		&b,
	)
	if !as.NoError(err) {
		return
	}
	req.Header.Set("Content-Type", w.FormDataContentType())

	// Sending request and check body existence
	var cli = new(http.Client)
	resp, err := cli.Do(req)
	if !as.NoError(err) || as.NotNil(resp.Body) {
		return
	}

	// Body validation
	respBt, err := ioutil.ReadAll(resp.Body)
	if !as.NoError(err) {
		return
	}
	as.Equal(suite.T(), 404, resp.StatusCode)
	var respBody = new(photo.CreateNotFoundBody)
	err = json.Unmarshal(respBt, respBody)
	if !as.NoError(err) {
		return
	}

	if as.NotNil(respBody.Version) {
		versionCheck(as, *respBody.Version)
	}
	as.Nil(respBody.Data)

	if as.NotNil(respBody.Message) {
		matched, err := regexp.MatchString(`[0-9A-Za-z] Not found`, respBody.Message)
		if as.NoError(err) {
			as.True(matched)
		}
	}

	as.Nil(respBody.Errors)
}

func (suite *RestApiTestingSuit) Test04() {
	LogTestCase(
		"get",
		"/photos/{photoID}",
		"USUALLY REQUEST",
		"SUCCESS",
	)
	as := assert.New(suite.T())

	resp, err := http.Get(suite.GetBaseUrl() + "/photos/" + suite.ImageID.String())
	if !as.NoError(err) {
		return
	}
	if !as.NotNil(resp.Body) {
		return
	}

	respBt, err := ioutil.ReadAll(resp.Body)
	if !as.NoError(err) {
		return
	}

	as.Equal(200, resp.StatusCode)
	var respBody = new(photo.ViewOKBody)
	err = json.Unmarshal(respBt, respBody)
	if !as.NoError(err) {
		return
	}

	if as.NotNil(respBody.Version) {
		versionCheck(as, *respBody.Version)
	}

	if as.NotNil(respBody.Message) {
		as.Equal("SUCCESS", *respBody.Message)
	}

	as.Nil(respBody.Errors)

	if as.NotEmpty(respBody.Data) && as.True(len(respBody.Data) == 1){
		item := respBody.Data[0]
		as.Equal(suite.ImageID, item.ID)

		as.Equal(string(suite.PersonID), item.PersonID.String())

		if as.NotNil(item.ExtService) {
			as.Equal(aws.SDKName, item.ExtService)
		}

		if as.NotNil(item.URL) {
			as.Equal(suite.GetBaseUrl() + "/photos/" + suite.ImageID.String()+"/", *item.URL)
		}
	}
}

func (suite *RestApiTestingSuit) Test05() {
	LogTestCase(
		"get",
		"/photos/{photoID}",
		"with nonexistent photoID",
		"NOT FOUND",
	)
	as := assert.New(suite.T())

	resp, err := http.Get(suite.GetBaseUrl() + "/photos/" + suite.ImageID.String())
	if !as.NoError(err) || as.NotNil(resp) {
		return
	}

	respBt, err := ioutil.ReadAll(resp.Body)
	if !as.NoError(err) {
		return
	}
	as.Equal(404, resp.StatusCode)
	var respBody = new(photo.ViewNotFoundBody)
	err = json.Unmarshal(respBt, respBody)
	if !as.NoError(err) {
		return
	}

	if as.NotNil(respBody.Version) {
		versionCheck(as, *respBody.Version)
	}

	as.Nil(respBody.Data)

	if as.NotNil(respBody.Message) {
		matched, err := regexp.MatchString(`[0-9A-Za-z] Not found`, respBody.Message)
		as.NoError(err)
		as.True(matched)
	}

	as.Nil(respBody.Errors)
}

func (suite *RestApiTestingSuit) Test06() {
	LogTestCase(
		"get",
		"/photos",
		"usually request",
		"SUCCESS",
	)
	as := assert.New(suite.T())

	resp, err := http.Get(suite.GetBaseUrl() + "/photos?limit=5&offset=0")
	if !as.NoError(err) || !as.NotNil(resp) {
		return
	}

	if as.NotNil(resp.Body) {
		return
	}

	respBt, err := ioutil.ReadAll(resp.Body)
	if !as.NoError(err) {
		return
	}
	as.Equal(200, resp.StatusCode)

	var respBody = new(recognize.RecognizeOKBody)
	err = json.Unmarshal(respBt, respBody)
	if as.NoError(err) {
		return
	}

	if as.NotNil(respBody.Version) {
		versionCheck(as, *respBody.Version)
	}

	if as.NotNil(respBody.Message){
		as.Equal("SUCCESS", *respBody.Message)
	}

	as.Nil(respBody.Errors)

	if as.NotEmpty(respBody.Data) && as.True(len(respBody.Data) == 1){
		as.Equal(aws.SDKName, respBody.Data[0].ExtService)
		as.Equal(suite.GetBaseUrl()+"/photos/",respBody.Data[0].URL)
		as.Equal(suite.ImageID, respBody.Data[0].ID)
		as.Equal(suite.PersonID, respBody.Data[0].PersonID)
	}
}

func (suite *RestApiTestingSuit) Test07() {
	LogTestCase(
		"get",
		"/photos",
		"request to empty storage",
		"NOT FOUND",
	)
	as := assert.New(suite.T())

	resp, err := http.Get(suite.GetBaseUrl() + "/photos?limit=5&offset=0")
	if !as.NoError(err) || !as.NotNil(resp) {
		return
	}

	if as.NotNil(resp.Body) {
		return
	}

	respBt, err := ioutil.ReadAll(resp.Body)
	if !as.NoError(err) {
		return
	}
	as.Equal(200, resp.StatusCode)

	var respBody = new(recognize.RecognizeOKBody)
	err = json.Unmarshal(respBt, respBody)
	if as.NoError(err) {
		return
	}

	if as.NotNil(respBody.Version) {
		versionCheck(as, *respBody.Version)
	}

	if as.NotNil(respBody.Message){
		as.Equal("SUCCESS", *respBody.Message)
	}

	if as.NotNil(respBody.Message) {
		matched, err := regexp.MatchString(`[0-9A-Za-z] Not found`, *respBody.Message)
		as.NoError(err)
		as.True(matched)
	}

	as.Nil(respBody.Errors)
}

func (suite *RestApiTestingSuit) Test08() {
	LogTestCase(
		"get",
		"/photos",
		"without query params",
		"validation error",
	)
	as := assert.New(suite.T())

	resp, err := http.Get(suite.GetBaseUrl() + "/photos?limit=5&offset=0")
	if !as.NoError(err) || !as.NotNil(resp) {
		return
	}


	if as.NotNil(resp.Body) {
		return
	}

	respBt, err := ioutil.ReadAll(resp.Body)
	if !as.NoError(err) {
		return
	}
	as.Equal(200, resp.StatusCode)

	var respBody = new(recognize.RecognizeOKBody)
	err = json.Unmarshal(respBt, respBody)
	if as.NoError(err) {
		return
	}

	if as.NotNil(respBody.Version) {
		versionCheck(as, *respBody.Version)
	}

	//TODO implement
}

func (suite *RestApiTestingSuit) Test09() {
	LogTestCase(
		"get",
		"/photos",
		"wit illegal query params",
		"validation error",
	)
	as := assert.New(suite.T())


	req, err := http.NewRequest(
		"DELETE",
		suite.GetBaseUrl()+"/photos/"+"AeqwezsdWAE",
		nil,
	)
	if !as.NoError(err) || !as.NotNil(req) {
		return
	}

	var cli = new(http.Client)
	resp, err := cli.Do(req)
	if !as.NoError(err) || !as.NotNil(resp) {
		return
	}

	respBt, err := ioutil.ReadAll(resp.Body)
	if !as.NoError(err) {
		return
	}
	as.Equal(400, resp.StatusCode)

	var respBody = new(photo.DeleteBadRequestBody)
	err = json.Unmarshal(respBt, respBody)
	if !as.NoError(err) {
		return
	}

	if as.NotNil(respBody.Version) {
		versionCheck(as, *respBody.Version)
	}

	if as.NotNil(respBody.Message) {
		 //TODO
	}

	//respBody.Errors TODO
}

func (suite *RestApiTestingSuit) Test10() {
	LogTestCase(
		"delete",
		"/photos",
		"with nonexistent photoID",
		"success",
	)
	as := assert.New(suite.T())

	generatedImageId, err := uuid.NewV1()
	if !as.NoError(err) {
		return
	}

	req, err := http.NewRequest(
		"DELETE",
		suite.GetBaseUrl()+"/photos/"+generatedImageId.String(),
		nil,
	)
	if !as.NoError(err) || !as.NotNil(req) {
		return
	}

	var cli = new(http.Client)
	resp, err := cli.Do(req)
	if !as.NoError(err) || !as.NotNil(resp.Body) {
		return
	}

	respBt, err := ioutil.ReadAll(resp.Body)
	if !as.NoError(err) {
		return
	}
	as.Equal(404, resp.StatusCode)

	var respBody = new(recognize.RecognizeOKBody)
	err = json.Unmarshal(respBt, respBody)
	if !as.NoError(err) {
		return
	}

	if as.NotNil(respBody.Version) {
		versionCheck(as, *respBody.Version)
	}

	as.Empty(respBody.Data)

	if as.NotNil(respBody.Message) {
		matched, err := regexp.MatchString(`Entity not found`, *respBody.Message)
		if as.NoError(err) {
			as.True(matched)
		}
	}

	if as.NotNil(respBody.Errors) || as.True(len(respBody.Errors) == 1){
		item := respBody.Errors[0]
		as.Equal(item, "s3 object not found")
	}
}

func (suite *RestApiTestingSuit) Test11() {
	LogTestCase(
		"delete",
		"/photos",
		"photoID not math uuid type",
		"success",
	)
	as := assert.New(suite.T())

	req, err := http.NewRequest(
		"DELETE",
		suite.GetBaseUrl()+"/photos/"+"NOT_UUID",
		nil,
	)
	if !as.NoError(err) || !as.NotNil(req) {
		return
	}

	var cli = new(http.Client)
	resp, err := cli.Do(req)
	if !as.NoError(err) || !as.NotNil(resp) {
		return
	}

	respBt, err := ioutil.ReadAll(resp.Body)
	if !as.NoError(err) {
		return
	}
	as.Equal(400, resp.StatusCode)

	var respBody = new(photo.DeleteBadRequestBody)
	err = json.Unmarshal(respBt, respBody)
	if !as.NoError(err) {
		return
	}

	if as.NotNil(respBody.Version) {
		versionCheck(as, *respBody.Version)
	}

	as.Empty(respBody.Data)

	if as.NotNil(respBody.Message) {
		as.Equal("Validation error", *respBody.Message)
	}

	//respBody.Errors TODO: implement
}

func (suite *RestApiTestingSuit) Test12() {
	LogTestCase(
		"post",
		"/recognize",
		"usually request",
		"success",
	)
	as := assert.New(suite.T())

	b, w := recognizeMultipartFormData(suite.Keys, suite.Image)
	req, err := http.NewRequest(
		"POST",
		suite.GetBaseUrl()+"/recognize",
		&b,
	)
	if !as.NoError(err) || !as.NotNil(req) {
		return
	}
	req.Header.Set("Content-Type", w.FormDataContentType())

	var cli = new(http.Client)
	resp, err := cli.Do(req)
	if !as.NoError(err) || !as.NotNil(resp) {
		return
	}

	respBt, err := ioutil.ReadAll(resp.Body)
	if !as.NoError(err) {
		return
	}
	as.Equal(200, resp.StatusCode)

	var respBody = new(recognize.RecognizeOKBody)
	err = json.Unmarshal(respBt, respBody)
	if !as.NoError(err) {
		return
	}

	if as.NotNil(respBody.Version)  {
		versionCheck(as, *respBody.Version)
	}

	if as.NotNil(respBody.Message) {
		as.Equal("SUCCESS", *respBody.Message)
	}

	as.Nil(respBody.Errors)

	if as.NotNil(respBody.Data) { //TODO: fixme 
		//if as.NotNil(respBody.Data[1].ID) {
		//	as.Equal(*suite.ImageID, *respBody.Data[1].ID)
		//}
		//
		//as.Equal(string(suite.PersonID), respBody.Data[1].PersonID.String())
		//as.Equal(aws.SDKName, respBody.Data[1].ExtService)
		//
		//if as.NotNil(respBody.Data[1].URL) {
		//	as.Equal(suite.GetBaseUrl()+"/recognize/"+string(suite.PersonID), *respBody.Data[1].URL)
		//}
		//
		//if as.NotNil(respBody.Data[1].Similarity) {
		//	as.IsType(float64(0), *respBody.Data[1].Similarity)
		//}
	}
}

func (suite *RestApiTestingSuit) Test13() {
	LogTestCase(
		"delete",
		"/photos",
		"usually request",
		"success",
	)
	as := assert.New(suite.T())

	req, err := http.NewRequest(
		"DELETE",
		suite.GetBaseUrl()+"/photos/"+suite.ImageID.String(),
		nil,
	)
	if !as.NoError(err) || !as.NotNil(req) {
		return
	}

	var cli = new(http.Client)
	resp, err := cli.Do(req)
	if !as.NoError(err) || !as.NotNil(resp) {
		return
	}

	respBt, err := ioutil.ReadAll(resp.Body)
	if !as.NoError(err) {
		return
	}
	as.Equal(200, resp.StatusCode)

	var respBody = new(photo.DeleteOKBody)
	err = json.Unmarshal(respBt, respBody)
	if !as.NoError(err) {
		return
	}

	if as.NotNil(respBody.Version) {
		versionCheck(as, *respBody.Version)
	}

	if as.NotNil(respBody.Message) {
		as.Equal("SUCCESS", *respBody.Message)
	}

	as.Nil(respBody.Errors)

	as.Nil(respBody.Data)
}