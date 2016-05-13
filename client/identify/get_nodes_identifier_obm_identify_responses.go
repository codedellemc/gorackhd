package identify

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/emccode/gorackhd/models"
)

// GetNodesIdentifierObmIdentifyReader is a Reader for the GetNodesIdentifierObmIdentify structure.
type GetNodesIdentifierObmIdentifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetNodesIdentifierObmIdentifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetNodesIdentifierObmIdentifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetNodesIdentifierObmIdentifyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetNodesIdentifierObmIdentifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetNodesIdentifierObmIdentifyOK creates a GetNodesIdentifierObmIdentifyOK with default headers values
func NewGetNodesIdentifierObmIdentifyOK() *GetNodesIdentifierObmIdentifyOK {
	return &GetNodesIdentifierObmIdentifyOK{}
}

/*GetNodesIdentifierObmIdentifyOK handles this case with default header values.

obm identity light settings
*/
type GetNodesIdentifierObmIdentifyOK struct {
	Payload GetNodesIdentifierObmIdentifyOKBodyBody
}

func (o *GetNodesIdentifierObmIdentifyOK) Error() string {
	return fmt.Sprintf("[GET /nodes/{identifier}/obm/identify][%d] getNodesIdentifierObmIdentifyOK  %+v", 200, o.Payload)
}

func (o *GetNodesIdentifierObmIdentifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNodesIdentifierObmIdentifyNotFound creates a GetNodesIdentifierObmIdentifyNotFound with default headers values
func NewGetNodesIdentifierObmIdentifyNotFound() *GetNodesIdentifierObmIdentifyNotFound {
	return &GetNodesIdentifierObmIdentifyNotFound{}
}

/*GetNodesIdentifierObmIdentifyNotFound handles this case with default header values.

The node with the identifier was not found or has no obm settings.

*/
type GetNodesIdentifierObmIdentifyNotFound struct {
	Payload *models.Error
}

func (o *GetNodesIdentifierObmIdentifyNotFound) Error() string {
	return fmt.Sprintf("[GET /nodes/{identifier}/obm/identify][%d] getNodesIdentifierObmIdentifyNotFound  %+v", 404, o.Payload)
}

func (o *GetNodesIdentifierObmIdentifyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNodesIdentifierObmIdentifyDefault creates a GetNodesIdentifierObmIdentifyDefault with default headers values
func NewGetNodesIdentifierObmIdentifyDefault(code int) *GetNodesIdentifierObmIdentifyDefault {
	return &GetNodesIdentifierObmIdentifyDefault{
		_statusCode: code,
	}
}

/*GetNodesIdentifierObmIdentifyDefault handles this case with default header values.

Unexpected error
*/
type GetNodesIdentifierObmIdentifyDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get nodes identifier obm identify default response
func (o *GetNodesIdentifierObmIdentifyDefault) Code() int {
	return o._statusCode
}

func (o *GetNodesIdentifierObmIdentifyDefault) Error() string {
	return fmt.Sprintf("[GET /nodes/{identifier}/obm/identify][%d] GetNodesIdentifierObmIdentify default  %+v", o._statusCode, o.Payload)
}

func (o *GetNodesIdentifierObmIdentifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetNodesIdentifierObmIdentifyOKBodyBody get nodes identifier obm identify o k body body

swagger:model GetNodesIdentifierObmIdentifyOKBodyBody
*/
type GetNodesIdentifierObmIdentifyOKBodyBody interface{}
