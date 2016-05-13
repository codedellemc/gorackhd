package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/emccode/gorackhd/models"
)

// PatchConfigReader is a Reader for the PatchConfig structure.
type PatchConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *PatchConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPatchConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPatchConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewPatchConfigOK creates a PatchConfigOK with default headers values
func NewPatchConfigOK() *PatchConfigOK {
	return &PatchConfigOK{}
}

/*PatchConfigOK handles this case with default header values.

An array of configuration objects
*/
type PatchConfigOK struct {
	Payload PatchConfigOKBodyBody
}

func (o *PatchConfigOK) Error() string {
	return fmt.Sprintf("[PATCH /config][%d] patchConfigOK  %+v", 200, o.Payload)
}

func (o *PatchConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConfigDefault creates a PatchConfigDefault with default headers values
func NewPatchConfigDefault(code int) *PatchConfigDefault {
	return &PatchConfigDefault{
		_statusCode: code,
	}
}

/*PatchConfigDefault handles this case with default header values.

Unexpected error
*/
type PatchConfigDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the patch config default response
func (o *PatchConfigDefault) Code() int {
	return o._statusCode
}

func (o *PatchConfigDefault) Error() string {
	return fmt.Sprintf("[PATCH /config][%d] PatchConfig default  %+v", o._statusCode, o.Payload)
}

func (o *PatchConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PatchConfigOKBodyBody patch config o k body body

swagger:model PatchConfigOKBodyBody
*/
type PatchConfigOKBodyBody interface{}
