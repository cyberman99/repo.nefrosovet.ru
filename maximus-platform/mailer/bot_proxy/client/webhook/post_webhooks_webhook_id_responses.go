// Code generated by go-swagger; DO NOT EDIT.

package webhook

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "repo.nefrosovet.ru/maximus-platform/mailer/bot_proxy/models"
)

// PostWebhooksWebhookIDReader is a Reader for the PostWebhooksWebhookID structure.
type PostWebhooksWebhookIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostWebhooksWebhookIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostWebhooksWebhookIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewPostWebhooksWebhookIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewPostWebhooksWebhookIDMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostWebhooksWebhookIDOK creates a PostWebhooksWebhookIDOK with default headers values
func NewPostWebhooksWebhookIDOK() *PostWebhooksWebhookIDOK {
	return &PostWebhooksWebhookIDOK{}
}

/*PostWebhooksWebhookIDOK handles this case with default header values.

Коллекция каналов
*/
type PostWebhooksWebhookIDOK struct {
	Payload *models.PostWebhooksWebhookIDOKBody
}

func (o *PostWebhooksWebhookIDOK) Error() string {
	return fmt.Sprintf("[POST /webhooks/{webhookID}][%d] postWebhooksWebhookIdOK  %+v", 200, o.Payload)
}

func (o *PostWebhooksWebhookIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostWebhooksWebhookIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWebhooksWebhookIDNotFound creates a PostWebhooksWebhookIDNotFound with default headers values
func NewPostWebhooksWebhookIDNotFound() *PostWebhooksWebhookIDNotFound {
	return &PostWebhooksWebhookIDNotFound{}
}

/*PostWebhooksWebhookIDNotFound handles this case with default header values.

Not found
*/
type PostWebhooksWebhookIDNotFound struct {
	Payload *models.PostWebhooksWebhookIDNotFoundBody
}

func (o *PostWebhooksWebhookIDNotFound) Error() string {
	return fmt.Sprintf("[POST /webhooks/{webhookID}][%d] postWebhooksWebhookIdNotFound  %+v", 404, o.Payload)
}

func (o *PostWebhooksWebhookIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostWebhooksWebhookIDNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWebhooksWebhookIDMethodNotAllowed creates a PostWebhooksWebhookIDMethodNotAllowed with default headers values
func NewPostWebhooksWebhookIDMethodNotAllowed() *PostWebhooksWebhookIDMethodNotAllowed {
	return &PostWebhooksWebhookIDMethodNotAllowed{}
}

/*PostWebhooksWebhookIDMethodNotAllowed handles this case with default header values.

Invalid Method
*/
type PostWebhooksWebhookIDMethodNotAllowed struct {
	Payload *models.PostWebhooksWebhookIDMethodNotAllowedBody
}

func (o *PostWebhooksWebhookIDMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /webhooks/{webhookID}][%d] postWebhooksWebhookIdMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *PostWebhooksWebhookIDMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostWebhooksWebhookIDMethodNotAllowedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
