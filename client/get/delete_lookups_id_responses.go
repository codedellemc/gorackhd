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

// DeleteLookupsIDReader is a Reader for the DeleteLookupsID structure.
type DeleteLookupsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *DeleteLookupsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteLookupsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteLookupsIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewDeleteLookupsIDOK creates a DeleteLookupsIDOK with default headers values
func NewDeleteLookupsIDOK() *DeleteLookupsIDOK {
	return &DeleteLookupsIDOK{}
}

/*DeleteLookupsIDOK handles this case with default header values.

array of all
*/
type DeleteLookupsIDOK struct {
	Payload []interface{}
}

func (o *DeleteLookupsIDOK) Error() string {
	return fmt.Sprintf("[DELETE /lookups/{id}][%d] deleteLookupsIdOK  %+v", 200, o.Payload)
}

func (o *DeleteLookupsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLookupsIDDefault creates a DeleteLookupsIDDefault with default headers values
func NewDeleteLookupsIDDefault(code int) *DeleteLookupsIDDefault {
	return &DeleteLookupsIDDefault{
		_statusCode: code,
	}
}

/*DeleteLookupsIDDefault handles this case with default header values.

Unexpected error
*/
type DeleteLookupsIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete lookups ID default response
func (o *DeleteLookupsIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteLookupsIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /lookups/{id}][%d] DeleteLookupsID default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteLookupsIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
