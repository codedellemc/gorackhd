package get

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/codedellemc/gorackhd/models"
)

// PatchLookupsIDReader is a Reader for the PatchLookupsID structure.
type PatchLookupsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *PatchLookupsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPatchLookupsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPatchLookupsIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewPatchLookupsIDOK creates a PatchLookupsIDOK with default headers values
func NewPatchLookupsIDOK() *PatchLookupsIDOK {
	return &PatchLookupsIDOK{}
}

/*PatchLookupsIDOK handles this case with default header values.

array of all
*/
type PatchLookupsIDOK struct {
	Payload []interface{}
}

func (o *PatchLookupsIDOK) Error() string {
	return fmt.Sprintf("[PATCH /lookups/{id}][%d] patchLookupsIdOK  %+v", 200, o.Payload)
}

func (o *PatchLookupsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLookupsIDDefault creates a PatchLookupsIDDefault with default headers values
func NewPatchLookupsIDDefault(code int) *PatchLookupsIDDefault {
	return &PatchLookupsIDDefault{
		_statusCode: code,
	}
}

/*PatchLookupsIDDefault handles this case with default header values.

Unexpected error
*/
type PatchLookupsIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the patch lookups ID default response
func (o *PatchLookupsIDDefault) Code() int {
	return o._statusCode
}

func (o *PatchLookupsIDDefault) Error() string {
	return fmt.Sprintf("[PATCH /lookups/{id}][%d] PatchLookupsID default  %+v", o._statusCode, o.Payload)
}

func (o *PatchLookupsIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
