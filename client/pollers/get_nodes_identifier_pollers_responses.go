package pollers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/codedellemc/gorackhd/models"
)

// GetNodesIdentifierPollersReader is a Reader for the GetNodesIdentifierPollers structure.
type GetNodesIdentifierPollersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetNodesIdentifierPollersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetNodesIdentifierPollersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetNodesIdentifierPollersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetNodesIdentifierPollersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetNodesIdentifierPollersOK creates a GetNodesIdentifierPollersOK with default headers values
func NewGetNodesIdentifierPollersOK() *GetNodesIdentifierPollersOK {
	return &GetNodesIdentifierPollersOK{}
}

/*GetNodesIdentifierPollersOK handles this case with default header values.

all pollers of specified node, empty object if none exist.

*/
type GetNodesIdentifierPollersOK struct {
	Payload []interface{}
}

func (o *GetNodesIdentifierPollersOK) Error() string {
	return fmt.Sprintf("[GET /nodes/{identifier}/pollers][%d] getNodesIdentifierPollersOK  %+v", 200, o.Payload)
}

func (o *GetNodesIdentifierPollersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNodesIdentifierPollersNotFound creates a GetNodesIdentifierPollersNotFound with default headers values
func NewGetNodesIdentifierPollersNotFound() *GetNodesIdentifierPollersNotFound {
	return &GetNodesIdentifierPollersNotFound{}
}

/*GetNodesIdentifierPollersNotFound handles this case with default header values.

The node with the identifier was not found

*/
type GetNodesIdentifierPollersNotFound struct {
	Payload *models.Error
}

func (o *GetNodesIdentifierPollersNotFound) Error() string {
	return fmt.Sprintf("[GET /nodes/{identifier}/pollers][%d] getNodesIdentifierPollersNotFound  %+v", 404, o.Payload)
}

func (o *GetNodesIdentifierPollersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNodesIdentifierPollersDefault creates a GetNodesIdentifierPollersDefault with default headers values
func NewGetNodesIdentifierPollersDefault(code int) *GetNodesIdentifierPollersDefault {
	return &GetNodesIdentifierPollersDefault{
		_statusCode: code,
	}
}

/*GetNodesIdentifierPollersDefault handles this case with default header values.

Unexpected error
*/
type GetNodesIdentifierPollersDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get nodes identifier pollers default response
func (o *GetNodesIdentifierPollersDefault) Code() int {
	return o._statusCode
}

func (o *GetNodesIdentifierPollersDefault) Error() string {
	return fmt.Sprintf("[GET /nodes/{identifier}/pollers][%d] GetNodesIdentifierPollers default  %+v", o._statusCode, o.Payload)
}

func (o *GetNodesIdentifierPollersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
