package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/spiegela/gorackhd/models"
)

// GetBootstrapReader is a Reader for the GetBootstrap structure.
type GetBootstrapReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBootstrapReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetBootstrapOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetBootstrapDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetBootstrapOK creates a GetBootstrapOK with default headers values
func NewGetBootstrapOK() *GetBootstrapOK {
	return &GetBootstrapOK{}
}

/*GetBootstrapOK handles this case with default header values.

Successfully retrieved bootstrap.js
*/
type GetBootstrapOK struct {
	Payload *models.VersionsResponse
}

func (o *GetBootstrapOK) Error() string {
	return fmt.Sprintf("[GET /tasks/bootstrap.js][%d] getBootstrapOK  %+v", 200, o.Payload)
}

func (o *GetBootstrapOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VersionsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBootstrapDefault creates a GetBootstrapDefault with default headers values
func NewGetBootstrapDefault(code int) *GetBootstrapDefault {
	return &GetBootstrapDefault{
		_statusCode: code,
	}
}

/*GetBootstrapDefault handles this case with default header values.

Error
*/
type GetBootstrapDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get bootstrap default response
func (o *GetBootstrapDefault) Code() int {
	return o._statusCode
}

func (o *GetBootstrapDefault) Error() string {
	return fmt.Sprintf("[GET /tasks/bootstrap.js][%d] getBootstrap default  %+v", o._statusCode, o.Payload)
}

func (o *GetBootstrapDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
