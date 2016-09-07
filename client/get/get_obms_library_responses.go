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

// GetObmsLibraryReader is a Reader for the GetObmsLibrary structure.
type GetObmsLibraryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetObmsLibraryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetObmsLibraryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetObmsLibraryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetObmsLibraryOK creates a GetObmsLibraryOK with default headers values
func NewGetObmsLibraryOK() *GetObmsLibraryOK {
	return &GetObmsLibraryOK{}
}

/*GetObmsLibraryOK handles this case with default header values.

get list of possible OBM services

*/
type GetObmsLibraryOK struct {
	Payload []interface{}
}

func (o *GetObmsLibraryOK) Error() string {
	return fmt.Sprintf("[GET /obms/library][%d] getObmsLibraryOK  %+v", 200, o.Payload)
}

func (o *GetObmsLibraryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetObmsLibraryDefault creates a GetObmsLibraryDefault with default headers values
func NewGetObmsLibraryDefault(code int) *GetObmsLibraryDefault {
	return &GetObmsLibraryDefault{
		_statusCode: code,
	}
}

/*GetObmsLibraryDefault handles this case with default header values.

Unexpected error
*/
type GetObmsLibraryDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get obms library default response
func (o *GetObmsLibraryDefault) Code() int {
	return o._statusCode
}

func (o *GetObmsLibraryDefault) Error() string {
	return fmt.Sprintf("[GET /obms/library][%d] GetObmsLibrary default  %+v", o._statusCode, o.Payload)
}

func (o *GetObmsLibraryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
