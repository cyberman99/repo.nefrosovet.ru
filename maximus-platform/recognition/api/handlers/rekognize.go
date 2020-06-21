package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"io"
	"io/ioutil"
	"repo.nefrosovet.ru/maximus-platform/recognition/api/models"
	"repo.nefrosovet.ru/maximus-platform/recognition/api/restapi/operations/photo"
	"repo.nefrosovet.ru/maximus-platform/recognition/api/restapi/operations/recognize"
	"repo.nefrosovet.ru/maximus-platform/recognition/logger"
	"repo.nefrosovet.ru/maximus-platform/recognition/services"
	"repo.nefrosovet.ru/maximus-platform/recognition/services/AWS"
	"strings"
)

type RekognitionController struct {
	s services.CloudServicer

	l       logger.APIEntrier
	version string
}

func NewRekognitionController(s services.CloudServicer, l logger.APIEntrier, version string) *RekognitionController {
	return &RekognitionController{s: s, l: l, version: version}
}

func (rc *RekognitionController) Rekognize(params recognize.RecognizeParams) middleware.Responder {
	var (
		file io.ReadCloser
		err  error
	)

	file = params.File
	defer file.Close()

	var img []byte
	img, err = ioutil.ReadAll(file)
	if err != nil {
		rc.l.Rekognition().Debug(err)
		payload := new(recognize.RecognizeInternalServerErrorBody)
		payload.Message = InternalServerErrorMessage
		payload.Version = &rc.version
		payload.Errors = &recognize.RecognizeBadRequestBodyAO1Errors{
			Validation: err.Error(),
		}
		return recognize.NewRecognizeInternalServerError().WithPayload(payload)
	}

	var ok bool
	_, ok = AWS.ValidateImageType(img)
	if !ok {
		rc.l.Rekognition().Debug(AWS.ErrUnsupportedImage)
		payload := new(recognize.RecognizeBadRequestBody)
		payload.Message = &PayloadValidationErrorMessage
		payload.Version = &rc.version
		payload.Errors = &recognize.RecognizeBadRequestBodyAO1Errors{
			Validation: AWS.ErrUnsupportedImage.Error(),
		}
		payload.Data = []interface{}{}
		return recognize.NewRecognizeBadRequest().WithPayload(payload)
	}

	var matches AWS.MatchedFaces
	matches, err = rc.s.Rekognize(img)
	if err == AWS.ErrUnrekognized || err == AWS.ErrObjectNotFound {
		rc.l.Rekognition().Debug(err)
		payload := new(recognize.RecognizeNotFoundBody)
		payload.Message = NotFoundMessage
		payload.Version = &rc.version
		payload.Errors = []interface{}{err.Error()}
		payload.Data = []interface{}{}
		return recognize.NewRecognizeNotFound().WithPayload(payload)
	}
	if err != nil {
		rc.l.Rekognition().Debug(err)
		if !strings.Contains(err.Error(), AWS.ErrPartiallyNotCompared.Error()) {
			payload := new(photo.CreateInternalServerErrorBody)
			payload.Message = InternalServerErrorMessage
			payload.Version = &rc.version
			payload.Errors = err.Error()
			payload.Data = []interface{}{}
			return photo.NewCreateInternalServerError().WithPayload(payload)
		}
	}

	var items = make([]*recognize.DataItems0, len(matches))
	for _, m := range matches {

		rc.l.Rekognition().RekInfo(m.Key.String(), logger.APIREKOGNIZED)

		items = append(
			items,
			&recognize.DataItems0{
				PhotoObject: models.PhotoObject{
					ID:         m.Key,
					URL:        buildURL(params.HTTPRequest, m.Key.String()),
					ExtService: rc.s.ServiceName(),
					PersonID:   "",
				},
				RecognizeObject: models.RecognizeObject{
					Similarity: m.Similarity,
				},
			},
		)
	}

	var payload = &recognize.RecognizeOKBody{
		SuccessData: models.SuccessData{},
		Data:        items,
	}
	payload.Version = &rc.version
	payload.Message = &PayloadSuccessMessage

	return recognize.NewRecognizeOK().WithPayload(payload)
}
