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

// LookupsGetReader is a Reader for the LookupsGet structure.
type LookupsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LookupsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewLookupsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewLookupsGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLookupsGetOK creates a LookupsGetOK with default headers values
func NewLookupsGetOK() *LookupsGetOK {
	return &LookupsGetOK{}
}

/*LookupsGetOK handles this case with default header values.

Successfully retrieved the list of lookups
*/
type LookupsGetOK struct {
	Payload models.LookupsGetOKBody
}

func (o *LookupsGetOK) Error() string {
	return fmt.Sprintf("[GET /lookups][%d] lookupsGetOK  %+v", 200, o.Payload)
}

func (o *LookupsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLookupsGetDefault creates a LookupsGetDefault with default headers values
func NewLookupsGetDefault(code int) *LookupsGetDefault {
	return &LookupsGetDefault{
		_statusCode: code,
	}
}

/*LookupsGetDefault handles this case with default header values.

Unexpected error
*/
type LookupsGetDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the lookups get default response
func (o *LookupsGetDefault) Code() int {
	return o._statusCode
}

func (o *LookupsGetDefault) Error() string {
	return fmt.Sprintf("[GET /lookups][%d] lookupsGet default  %+v", o._statusCode, o.Payload)
}

func (o *LookupsGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
