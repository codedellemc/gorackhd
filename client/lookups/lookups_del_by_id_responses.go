// Code generated by go-swagger; DO NOT EDIT.

package lookups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/spiegela/gorackhd/models"
)

// LookupsDelByIDReader is a Reader for the LookupsDelByID structure.
type LookupsDelByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LookupsDelByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewLookupsDelByIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewLookupsDelByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewLookupsDelByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLookupsDelByIDNoContent creates a LookupsDelByIDNoContent with default headers values
func NewLookupsDelByIDNoContent() *LookupsDelByIDNoContent {
	return &LookupsDelByIDNoContent{}
}

/*LookupsDelByIDNoContent handles this case with default header values.

Successfully deleted lookup
*/
type LookupsDelByIDNoContent struct {
}

func (o *LookupsDelByIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /lookups/{id}][%d] lookupsDelByIdNoContent ", 204)
}

func (o *LookupsDelByIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLookupsDelByIDNotFound creates a LookupsDelByIDNotFound with default headers values
func NewLookupsDelByIDNotFound() *LookupsDelByIDNotFound {
	return &LookupsDelByIDNotFound{}
}

/*LookupsDelByIDNotFound handles this case with default header values.

The specified lookup was not found
*/
type LookupsDelByIDNotFound struct {
	Payload *models.Error
}

func (o *LookupsDelByIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /lookups/{id}][%d] lookupsDelByIdNotFound  %+v", 404, o.Payload)
}

func (o *LookupsDelByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLookupsDelByIDDefault creates a LookupsDelByIDDefault with default headers values
func NewLookupsDelByIDDefault(code int) *LookupsDelByIDDefault {
	return &LookupsDelByIDDefault{
		_statusCode: code,
	}
}

/*LookupsDelByIDDefault handles this case with default header values.

Unexpected error
*/
type LookupsDelByIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the lookups del by Id default response
func (o *LookupsDelByIDDefault) Code() int {
	return o._statusCode
}

func (o *LookupsDelByIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /lookups/{id}][%d] lookupsDelById default  %+v", o._statusCode, o.Payload)
}

func (o *LookupsDelByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
