package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"io"
	"io/ioutil"
	"net/http"
	"repo.nefrosovet.ru/maximus-platform/recognition/api/models"
	"repo.nefrosovet.ru/maximus-platform/recognition/api/restapi/operations/photo"
	"repo.nefrosovet.ru/maximus-platform/recognition/api/restapi/operations/recognize"
	"repo.nefrosovet.ru/maximus-platform/recognition/logger"
	"repo.nefrosovet.ru/maximus-platform/recognition/services"
	"repo.nefrosovet.ru/maximus-platform/recognition/services/AWS"
	errs "github.com/go-openapi/errors"
)

type PhotoController struct {
	s services.CloudServicer

	l       logger.APIEntrier
	version string
}

func NewPhotoController(s services.CloudServicer, l logger.APIEntrier, version string) *PhotoController {
	return &PhotoController{s: s, l: l, version: version}
}

func (phc *PhotoController) Post(params photo.CreateParams) (m middleware.Responder) {
	var (
		file io.ReadCloser
		err  error
	)
	if params.PersonID == "" {
		payload := new(photo.CreateBadRequestBody)
		payload.Version = &phc.version
		payload.Message = &PayloadValidationErrorMessage
		payload.Errors = &photo.CreateBadRequestBodyAO1Errors{
			Validation: errs.Required("personID", "body"),
		}
		payload.Data = []interface{}{}
		return photo.NewCreateBadRequest().WithPayload(payload)
	}

	file = params.File
	defer file.Close()

	var img []byte
	img, err = ioutil.ReadAll(file)
	if err != nil {
		phc.l.Photo().Debug(err)
		payload := new(photo.CreateInternalServerErrorBody)
		payload.Version = &phc.version
		payload.Message = InternalServerErrorMessage
		payload.Errors = err.Error()
		payload.Data = []interface{}{}
		return photo.NewCreateInternalServerError().WithPayload(payload)
	}

	var (
		format string
		ok     bool
	)
	format, ok = AWS.ValidateImageType(img)
	if !ok {
		phc.l.Photo().Debug(AWS.ErrUnsupportedImage)
		payload := new(recognize.RecognizeBadRequestBody)
		payload.Message = &PayloadValidationErrorMessage
		payload.Version = &phc.version
		payload.Errors = &recognize.RecognizeBadRequestBodyAO1Errors{
			Validation: AWS.ErrUnsupportedImage.Error(),
		}
		payload.Data = []interface{}{}
		return recognize.NewRecognizeBadRequest().WithPayload(payload)
	}

	var imgID *strfmt.UUID
	imgID, err = phc.s.Set(params.PersonID, format, img)
	if err != nil {
		phc.l.Photo().Debug(err)
		payload := new(photo.CreateInternalServerErrorBody)
		payload.Message = InternalServerErrorMessage
		payload.Version = &phc.version
		payload.Errors = err.Error()
		payload.Data = []interface{}{}
		return photo.NewCreateInternalServerError().WithPayload(payload)
	}

	phc.l.Photo().Info(imgID.String(), params.PersonID, logger.APIUPLOADED)

	var payload = new(photo.CreateOKBody)
	payload = &photo.CreateOKBody{
		SuccessData: models.SuccessData{},
		Data: []*photo.DataItems0{
			{
				models.PhotoObject{
					ID:         imgID,
					URL:        buildURL(params.HTTPRequest, imgID.String()),
					ExtService: phc.s.ServiceName(),
					PersonID:   strfmt.UUID(params.PersonID),
				},
			},
		},
	}

	payload.Version = &phc.version
	payload.Message = &PayloadSuccessMessage

	return photo.NewCreateOK().WithPayload(payload)
}

func (phc *PhotoController) Get(params photo.ViewParams) middleware.Responder {
	var (
		personID string
		err      error
	)

	personID, err = phc.s.Get(params.PhotoID.String())
	if err == AWS.ErrObjectNotFound {
		phc.l.Photo().Debug(AWS.ErrObjectNotFound, params.PhotoID.String())
		payload := new(photo.ViewNotFoundBody)
		payload.Message = NotFoundMessage
		payload.Version = &phc.version
		payload.Errors = []interface{}{err.Error()}
		payload.Data = []interface{}{}
		return photo.NewViewNotFound().WithPayload(payload)
	}
	if err != nil {
		phc.l.Photo().Debug(err)
		payload := new(photo.ViewInternalServerErrorBody)
		payload.Message = InternalServerErrorMessage
		payload.Version = &phc.version
		payload.Errors = err.Error()
		payload.Data = []interface{}{}
		return photo.NewViewInternalServerError().WithPayload(payload)
	}

	var payload = &photo.ViewOKBody{
		SuccessData: models.SuccessData{},
		Data: []*photo.DataItems0{
			{
				models.PhotoObject{
					ID:         &params.PhotoID,
					URL:        buildURL(params.HTTPRequest, ""),
					ExtService: phc.s.ServiceName(),
					PersonID:   strfmt.UUID(personID),
				},
			},
		},
	}
	payload.Version = &phc.version
	payload.Message = &PayloadSuccessMessage

	return photo.NewViewOK().WithPayload(payload)

}

func (phc *PhotoController) Delete(params photo.DeleteParams) middleware.Responder {
	var err error
	err = phc.s.Delete(params.PhotoID.String())
	if err == AWS.ErrObjectNotFound {
		phc.l.Photo().Debug(AWS.ErrObjectNotFound, params.PhotoID.String())
		payload := new(photo.DeleteNotFoundBody)
		payload.Message = NotFoundMessage
		payload.Version = &phc.version
		payload.Errors = []interface{}{err.Error()}
		payload.Data = []interface{}{}
		return photo.NewDeleteNotFound().WithPayload(payload)
	}
	if err != nil {
		phc.l.Photo().Debug(err)
		payload := new(photo.DeleteInternalServerErrorBody)
		payload.Message = InternalServerErrorMessage
		payload.Version = &phc.version
		payload.Errors = err.Error()
		payload.Data = []interface{}{}
		return photo.NewDeleteInternalServerError().WithPayload(payload)
	}

	phc.l.Photo().Info(params.PhotoID.String(), "", logger.APIDELETED)

	var payload = &photo.DeleteOKBody{
		SuccessData: models.SuccessData{},
	}
	payload.Version = &phc.version
	payload.Message = &PayloadSuccessMessage

	return photo.NewDeleteOK().WithPayload(payload)
}

func (phc *PhotoController) List(params photo.CollectionParams) middleware.Responder {
	var offset int64

	if params.Offset != nil {
		offset = *params.Offset
	}

	var (
		keys []*strfmt.UUID
		err  error
	)
	keys, err = phc.s.List(params.Limit)
	if err == AWS.ErrObjectNotFound || err == AWS.ErrBucketDoesntExist {
		phc.l.Photo().Debug(err)
		payload := new(photo.CollectionNotFoundBody)
		payload.Message = NotFoundMessage
		payload.Version = &phc.version
		payload.Errors = []interface{}{err.Error()}
		payload.Data = []interface{}{}
		return photo.NewCollectionNotFound().WithPayload(payload)
	}
	if err != nil {
		phc.l.Photo().Debug(err)
		payload := new(photo.CollectionInternalServerErrorBody)
		payload.Message = InternalServerErrorMessage
		payload.Version = &phc.version
		payload.Errors = err.Error()
		payload.Data = []interface{}{}
		return photo.NewCollectionInternalServerError().WithPayload(payload)
	}

	var items = make([]*photo.DataItems0, (len(keys) - int(offset)))
	for i := int(offset); i < len(keys); i++ {
		if keys[i] == nil {
			continue
		}
		items = append(
			items,
			&photo.DataItems0{
				PhotoObject: models.PhotoObject{
					ID:         keys[i],
					URL:        buildURL(params.HTTPRequest, keys[i].String()),
					ExtService: phc.s.ServiceName(),
					PersonID:   "",
				},
			},
		)
	}

	var payload = &photo.CollectionOKBody{
		SuccessData: models.SuccessData{},
		Data:        items,
	}
	payload.Version = &phc.version
	payload.Message = &PayloadSuccessMessage

	return photo.NewCollectionOK().WithPayload(payload)
}

func buildURL(req *http.Request, id string) *string {
	url := "http://" + req.Host + req.RequestURI + "/" + id
	return &url
}
