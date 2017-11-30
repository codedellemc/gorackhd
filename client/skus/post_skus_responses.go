package skus

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/codedellemc/gorackhd/models"
)

// PostSkusReader is a Reader for the PostSkus structure.
type PostSkusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *PostSkusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostSkusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewPostSkusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPostSkusDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewPostSkusOK creates a PostSkusOK with default headers values
func NewPostSkusOK() *PostSkusOK {
	return &PostSkusOK{}
}

/*PostSkusOK handles this case with default header values.

sku to create

*/
type PostSkusOK struct {
	Payload PostSkusOKBodyBody
}

func (o *PostSkusOK) Error() string {
	return fmt.Sprintf("[POST /skus][%d] postSkusOK  %+v", 200, o.Payload)
}

func (o *PostSkusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSkusInternalServerError creates a PostSkusInternalServerError with default headers values
func NewPostSkusInternalServerError() *PostSkusInternalServerError {
	return &PostSkusInternalServerError{}
}

/*PostSkusInternalServerError handles this case with default header values.

Upload failed.

*/
type PostSkusInternalServerError struct {
	Payload *models.Error
}

func (o *PostSkusInternalServerError) Error() string {
	return fmt.Sprintf("[POST /skus][%d] postSkusInternalServerError  %+v", 500, o.Payload)
}

func (o *PostSkusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSkusDefault creates a PostSkusDefault with default headers values
func NewPostSkusDefault(code int) *PostSkusDefault {
	return &PostSkusDefault{
		_statusCode: code,
	}
}

/*PostSkusDefault handles this case with default header values.

Unexpected error
*/
type PostSkusDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post skus default response
func (o *PostSkusDefault) Code() int {
	return o._statusCode
}

func (o *PostSkusDefault) Error() string {
	return fmt.Sprintf("[POST /skus][%d] PostSkus default  %+v", o._statusCode, o.Payload)
}

func (o *PostSkusDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostSkusOKBodyBody post skus o k body body

swagger:model PostSkusOKBodyBody
*/
type PostSkusOKBodyBody interface{}