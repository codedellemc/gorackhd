package files

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/codedellemc/gorackhd/models"
)

// DeleteFilesFileidentifierReader is a Reader for the DeleteFilesFileidentifier structure.
type DeleteFilesFileidentifierReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *DeleteFilesFileidentifierReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteFilesFileidentifierNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewDeleteFilesFileidentifierNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeleteFilesFileidentifierInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteFilesFileidentifierDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewDeleteFilesFileidentifierNoContent creates a DeleteFilesFileidentifierNoContent with default headers values
func NewDeleteFilesFileidentifierNoContent() *DeleteFilesFileidentifierNoContent {
	return &DeleteFilesFileidentifierNoContent{}
}

/*DeleteFilesFileidentifierNoContent handles this case with default header values.

File successfully deleted.
*/
type DeleteFilesFileidentifierNoContent struct {
}

func (o *DeleteFilesFileidentifierNoContent) Error() string {
	return fmt.Sprintf("[DELETE /files/{fileidentifier}][%d] deleteFilesFileidentifierNoContent ", 204)
}

func (o *DeleteFilesFileidentifierNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteFilesFileidentifierNotFound creates a DeleteFilesFileidentifierNotFound with default headers values
func NewDeleteFilesFileidentifierNotFound() *DeleteFilesFileidentifierNotFound {
	return &DeleteFilesFileidentifierNotFound{}
}

/*DeleteFilesFileidentifierNotFound handles this case with default header values.

File not found.
*/
type DeleteFilesFileidentifierNotFound struct {
	Payload *models.Error
}

func (o *DeleteFilesFileidentifierNotFound) Error() string {
	return fmt.Sprintf("[DELETE /files/{fileidentifier}][%d] deleteFilesFileidentifierNotFound  %+v", 404, o.Payload)
}

func (o *DeleteFilesFileidentifierNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFilesFileidentifierInternalServerError creates a DeleteFilesFileidentifierInternalServerError with default headers values
func NewDeleteFilesFileidentifierInternalServerError() *DeleteFilesFileidentifierInternalServerError {
	return &DeleteFilesFileidentifierInternalServerError{}
}

/*DeleteFilesFileidentifierInternalServerError handles this case with default header values.

Error deleting file from the database.
*/
type DeleteFilesFileidentifierInternalServerError struct {
	Payload *models.Error
}

func (o *DeleteFilesFileidentifierInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /files/{fileidentifier}][%d] deleteFilesFileidentifierInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteFilesFileidentifierInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFilesFileidentifierDefault creates a DeleteFilesFileidentifierDefault with default headers values
func NewDeleteFilesFileidentifierDefault(code int) *DeleteFilesFileidentifierDefault {
	return &DeleteFilesFileidentifierDefault{
		_statusCode: code,
	}
}

/*DeleteFilesFileidentifierDefault handles this case with default header values.

Unexpected error
*/
type DeleteFilesFileidentifierDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete files fileidentifier default response
func (o *DeleteFilesFileidentifierDefault) Code() int {
	return o._statusCode
}

func (o *DeleteFilesFileidentifierDefault) Error() string {
	return fmt.Sprintf("[DELETE /files/{fileidentifier}][%d] DeleteFilesFileidentifier default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteFilesFileidentifierDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
