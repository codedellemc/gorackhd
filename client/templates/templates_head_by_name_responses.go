package templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/spiegela/gorackhd/models"
)

// TemplatesHeadByNameReader is a Reader for the TemplatesHeadByName structure.
type TemplatesHeadByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TemplatesHeadByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewTemplatesHeadByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewTemplatesHeadByNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewTemplatesHeadByNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTemplatesHeadByNameOK creates a TemplatesHeadByNameOK with default headers values
func NewTemplatesHeadByNameOK() *TemplatesHeadByNameOK {
	return &TemplatesHeadByNameOK{}
}

/*TemplatesHeadByNameOK handles this case with default header values.

Successfuly retrieved the specified template headers
*/
type TemplatesHeadByNameOK struct {
	Payload TemplatesHeadByNameOKBody
}

func (o *TemplatesHeadByNameOK) Error() string {
	return fmt.Sprintf("[HEAD /templates/{name}][%d] templatesHeadByNameOK  %+v", 200, o.Payload)
}

func (o *TemplatesHeadByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTemplatesHeadByNameNotFound creates a TemplatesHeadByNameNotFound with default headers values
func NewTemplatesHeadByNameNotFound() *TemplatesHeadByNameNotFound {
	return &TemplatesHeadByNameNotFound{}
}

/*TemplatesHeadByNameNotFound handles this case with default header values.

The template with specified identifier was not found
*/
type TemplatesHeadByNameNotFound struct {
	Payload *models.Error
}

func (o *TemplatesHeadByNameNotFound) Error() string {
	return fmt.Sprintf("[HEAD /templates/{name}][%d] templatesHeadByNameNotFound  %+v", 404, o.Payload)
}

func (o *TemplatesHeadByNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTemplatesHeadByNameDefault creates a TemplatesHeadByNameDefault with default headers values
func NewTemplatesHeadByNameDefault(code int) *TemplatesHeadByNameDefault {
	return &TemplatesHeadByNameDefault{
		_statusCode: code,
	}
}

/*TemplatesHeadByNameDefault handles this case with default header values.

Unexpected error
*/
type TemplatesHeadByNameDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the templates head by name default response
func (o *TemplatesHeadByNameDefault) Code() int {
	return o._statusCode
}

func (o *TemplatesHeadByNameDefault) Error() string {
	return fmt.Sprintf("[HEAD /templates/{name}][%d] templatesHeadByName default  %+v", o._statusCode, o.Payload)
}

func (o *TemplatesHeadByNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*TemplatesHeadByNameOKBody templates head by name o k body
swagger:model TemplatesHeadByNameOKBody
*/
type TemplatesHeadByNameOKBody interface{}
