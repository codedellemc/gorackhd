package profiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/codedellemc/gorackhd/models"
)

// PutProfilesLibraryIdentifierReader is a Reader for the PutProfilesLibraryIdentifier structure.
type PutProfilesLibraryIdentifierReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *PutProfilesLibraryIdentifierReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutProfilesLibraryIdentifierOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewPutProfilesLibraryIdentifierInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPutProfilesLibraryIdentifierDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewPutProfilesLibraryIdentifierOK creates a PutProfilesLibraryIdentifierOK with default headers values
func NewPutProfilesLibraryIdentifierOK() *PutProfilesLibraryIdentifierOK {
	return &PutProfilesLibraryIdentifierOK{}
}

/*PutProfilesLibraryIdentifierOK handles this case with default header values.

profile to put

*/
type PutProfilesLibraryIdentifierOK struct {
	Payload PutProfilesLibraryIdentifierOKBodyBody
}

func (o *PutProfilesLibraryIdentifierOK) Error() string {
	return fmt.Sprintf("[PUT /profiles/library/{identifier}][%d] putProfilesLibraryIdentifierOK  %+v", 200, o.Payload)
}

func (o *PutProfilesLibraryIdentifierOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutProfilesLibraryIdentifierInternalServerError creates a PutProfilesLibraryIdentifierInternalServerError with default headers values
func NewPutProfilesLibraryIdentifierInternalServerError() *PutProfilesLibraryIdentifierInternalServerError {
	return &PutProfilesLibraryIdentifierInternalServerError{}
}

/*PutProfilesLibraryIdentifierInternalServerError handles this case with default header values.

Upload failed.

*/
type PutProfilesLibraryIdentifierInternalServerError struct {
	Payload *models.Error
}

func (o *PutProfilesLibraryIdentifierInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /profiles/library/{identifier}][%d] putProfilesLibraryIdentifierInternalServerError  %+v", 500, o.Payload)
}

func (o *PutProfilesLibraryIdentifierInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutProfilesLibraryIdentifierDefault creates a PutProfilesLibraryIdentifierDefault with default headers values
func NewPutProfilesLibraryIdentifierDefault(code int) *PutProfilesLibraryIdentifierDefault {
	return &PutProfilesLibraryIdentifierDefault{
		_statusCode: code,
	}
}

/*PutProfilesLibraryIdentifierDefault handles this case with default header values.

Unexpected error
*/
type PutProfilesLibraryIdentifierDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put profiles library identifier default response
func (o *PutProfilesLibraryIdentifierDefault) Code() int {
	return o._statusCode
}

func (o *PutProfilesLibraryIdentifierDefault) Error() string {
	return fmt.Sprintf("[PUT /profiles/library/{identifier}][%d] PutProfilesLibraryIdentifier default  %+v", o._statusCode, o.Payload)
}

func (o *PutProfilesLibraryIdentifierDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PutProfilesLibraryIdentifierOKBodyBody put profiles library identifier o k body body

swagger:model PutProfilesLibraryIdentifierOKBodyBody
*/
type PutProfilesLibraryIdentifierOKBodyBody interface{}
