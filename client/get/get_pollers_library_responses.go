package get

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/emccode/gorackhd/models"
)

// GetPollersLibraryReader is a Reader for the GetPollersLibrary structure.
type GetPollersLibraryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetPollersLibraryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPollersLibraryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetPollersLibraryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetPollersLibraryOK creates a GetPollersLibraryOK with default headers values
func NewGetPollersLibraryOK() *GetPollersLibraryOK {
	return &GetPollersLibraryOK{}
}

/*GetPollersLibraryOK handles this case with default header values.

list of all pollers

*/
type GetPollersLibraryOK struct {
	Payload GetPollersLibraryOKBodyBody
}

func (o *GetPollersLibraryOK) Error() string {
	return fmt.Sprintf("[GET /pollers/library][%d] getPollersLibraryOK  %+v", 200, o.Payload)
}

func (o *GetPollersLibraryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPollersLibraryDefault creates a GetPollersLibraryDefault with default headers values
func NewGetPollersLibraryDefault(code int) *GetPollersLibraryDefault {
	return &GetPollersLibraryDefault{
		_statusCode: code,
	}
}

/*GetPollersLibraryDefault handles this case with default header values.

Unexpected error
*/
type GetPollersLibraryDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get pollers library default response
func (o *GetPollersLibraryDefault) Code() int {
	return o._statusCode
}

func (o *GetPollersLibraryDefault) Error() string {
	return fmt.Sprintf("[GET /pollers/library][%d] GetPollersLibrary default  %+v", o._statusCode, o.Payload)
}

func (o *GetPollersLibraryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetPollersLibraryOKBodyBody get pollers library o k body body

swagger:model GetPollersLibraryOKBodyBody
*/
type GetPollersLibraryOKBodyBody interface{}
