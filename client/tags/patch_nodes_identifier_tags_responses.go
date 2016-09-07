package tags

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/codedellemc/gorackhd/models"
)

// PatchNodesIdentifierTagsReader is a Reader for the PatchNodesIdentifierTags structure.
type PatchNodesIdentifierTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *PatchNodesIdentifierTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPatchNodesIdentifierTagsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewPatchNodesIdentifierTagsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPatchNodesIdentifierTagsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewPatchNodesIdentifierTagsOK creates a PatchNodesIdentifierTagsOK with default headers values
func NewPatchNodesIdentifierTagsOK() *PatchNodesIdentifierTagsOK {
	return &PatchNodesIdentifierTagsOK{}
}

/*PatchNodesIdentifierTagsOK handles this case with default header values.

patch succeeded
*/
type PatchNodesIdentifierTagsOK struct {
	Payload PatchNodesIdentifierTagsOKBodyBody
}

func (o *PatchNodesIdentifierTagsOK) Error() string {
	return fmt.Sprintf("[PATCH /nodes/{identifier}/tags][%d] patchNodesIdentifierTagsOK  %+v", 200, o.Payload)
}

func (o *PatchNodesIdentifierTagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchNodesIdentifierTagsNotFound creates a PatchNodesIdentifierTagsNotFound with default headers values
func NewPatchNodesIdentifierTagsNotFound() *PatchNodesIdentifierTagsNotFound {
	return &PatchNodesIdentifierTagsNotFound{}
}

/*PatchNodesIdentifierTagsNotFound handles this case with default header values.

Not found
*/
type PatchNodesIdentifierTagsNotFound struct {
	Payload *models.Error
}

func (o *PatchNodesIdentifierTagsNotFound) Error() string {
	return fmt.Sprintf("[PATCH /nodes/{identifier}/tags][%d] patchNodesIdentifierTagsNotFound  %+v", 404, o.Payload)
}

func (o *PatchNodesIdentifierTagsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchNodesIdentifierTagsDefault creates a PatchNodesIdentifierTagsDefault with default headers values
func NewPatchNodesIdentifierTagsDefault(code int) *PatchNodesIdentifierTagsDefault {
	return &PatchNodesIdentifierTagsDefault{
		_statusCode: code,
	}
}

/*PatchNodesIdentifierTagsDefault handles this case with default header values.

Unexpected error
*/
type PatchNodesIdentifierTagsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the patch nodes identifier tags default response
func (o *PatchNodesIdentifierTagsDefault) Code() int {
	return o._statusCode
}

func (o *PatchNodesIdentifierTagsDefault) Error() string {
	return fmt.Sprintf("[PATCH /nodes/{identifier}/tags][%d] PatchNodesIdentifierTags default  %+v", o._statusCode, o.Payload)
}

func (o *PatchNodesIdentifierTagsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PatchNodesIdentifierTagsOKBodyBody patch nodes identifier tags o k body body

swagger:model PatchNodesIdentifierTagsOKBodyBody
*/
type PatchNodesIdentifierTagsOKBodyBody interface{}
