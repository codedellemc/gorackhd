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

// GetWorkflowsReader is a Reader for the GetWorkflows structure.
type GetWorkflowsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetWorkflowsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetWorkflowsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetWorkflowsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetWorkflowsOK creates a GetWorkflowsOK with default headers values
func NewGetWorkflowsOK() *GetWorkflowsOK {
	return &GetWorkflowsOK{}
}

/*GetWorkflowsOK handles this case with default header values.

Fetch workflows

*/
type GetWorkflowsOK struct {
	Payload GetWorkflowsOKBodyBody
}

func (o *GetWorkflowsOK) Error() string {
	return fmt.Sprintf("[GET /workflows][%d] getWorkflowsOK  %+v", 200, o.Payload)
}

func (o *GetWorkflowsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkflowsDefault creates a GetWorkflowsDefault with default headers values
func NewGetWorkflowsDefault(code int) *GetWorkflowsDefault {
	return &GetWorkflowsDefault{
		_statusCode: code,
	}
}

/*GetWorkflowsDefault handles this case with default header values.

Unexpected error
*/
type GetWorkflowsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get workflows default response
func (o *GetWorkflowsDefault) Code() int {
	return o._statusCode
}

func (o *GetWorkflowsDefault) Error() string {
	return fmt.Sprintf("[GET /workflows][%d] GetWorkflows default  %+v", o._statusCode, o.Payload)
}

func (o *GetWorkflowsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetWorkflowsOKBodyBody get workflows o k body body

swagger:model GetWorkflowsOKBodyBody
*/
type GetWorkflowsOKBodyBody interface{}
