package catalogs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/codedellemc/gorackhd/models"
)

// GetNodesIdentifierCatalogsReader is a Reader for the GetNodesIdentifierCatalogs structure.
type GetNodesIdentifierCatalogsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetNodesIdentifierCatalogsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetNodesIdentifierCatalogsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetNodesIdentifierCatalogsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetNodesIdentifierCatalogsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetNodesIdentifierCatalogsOK creates a GetNodesIdentifierCatalogsOK with default headers values
func NewGetNodesIdentifierCatalogsOK() *GetNodesIdentifierCatalogsOK {
	return &GetNodesIdentifierCatalogsOK{}
}

/*GetNodesIdentifierCatalogsOK handles this case with default header values.

all catalogs of specified node, empty object if none exist.

*/
type GetNodesIdentifierCatalogsOK struct {
	Payload []interface{}
}

func (o *GetNodesIdentifierCatalogsOK) Error() string {
	return fmt.Sprintf("[GET /nodes/{identifier}/catalogs][%d] getNodesIdentifierCatalogsOK  %+v", 200, o.Payload)
}

func (o *GetNodesIdentifierCatalogsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNodesIdentifierCatalogsNotFound creates a GetNodesIdentifierCatalogsNotFound with default headers values
func NewGetNodesIdentifierCatalogsNotFound() *GetNodesIdentifierCatalogsNotFound {
	return &GetNodesIdentifierCatalogsNotFound{}
}

/*GetNodesIdentifierCatalogsNotFound handles this case with default header values.

The node with the identifier was not found

*/
type GetNodesIdentifierCatalogsNotFound struct {
	Payload *models.Error
}

func (o *GetNodesIdentifierCatalogsNotFound) Error() string {
	return fmt.Sprintf("[GET /nodes/{identifier}/catalogs][%d] getNodesIdentifierCatalogsNotFound  %+v", 404, o.Payload)
}

func (o *GetNodesIdentifierCatalogsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNodesIdentifierCatalogsDefault creates a GetNodesIdentifierCatalogsDefault with default headers values
func NewGetNodesIdentifierCatalogsDefault(code int) *GetNodesIdentifierCatalogsDefault {
	return &GetNodesIdentifierCatalogsDefault{
		_statusCode: code,
	}
}

/*GetNodesIdentifierCatalogsDefault handles this case with default header values.

Unexpected error
*/
type GetNodesIdentifierCatalogsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get nodes identifier catalogs default response
func (o *GetNodesIdentifierCatalogsDefault) Code() int {
	return o._statusCode
}

func (o *GetNodesIdentifierCatalogsDefault) Error() string {
	return fmt.Sprintf("[GET /nodes/{identifier}/catalogs][%d] GetNodesIdentifierCatalogs default  %+v", o._statusCode, o.Payload)
}

func (o *GetNodesIdentifierCatalogsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
