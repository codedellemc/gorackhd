// Code generated by go-swagger; DO NOT EDIT.

package workflows

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/spiegela/gorackhd/models"
)

// WorkflowsGetAllTasksReader is a Reader for the WorkflowsGetAllTasks structure.
type WorkflowsGetAllTasksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WorkflowsGetAllTasksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewWorkflowsGetAllTasksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewWorkflowsGetAllTasksDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWorkflowsGetAllTasksOK creates a WorkflowsGetAllTasksOK with default headers values
func NewWorkflowsGetAllTasksOK() *WorkflowsGetAllTasksOK {
	return &WorkflowsGetAllTasksOK{}
}

/*WorkflowsGetAllTasksOK handles this case with default header values.

Successfully retrieved workflow tasks
*/
type WorkflowsGetAllTasksOK struct {
	Payload WorkflowsGetAllTasksOKBody
}

func (o *WorkflowsGetAllTasksOK) Error() string {
	return fmt.Sprintf("[GET /workflows/tasks][%d] workflowsGetAllTasksOK  %+v", 200, o.Payload)
}

func (o *WorkflowsGetAllTasksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWorkflowsGetAllTasksDefault creates a WorkflowsGetAllTasksDefault with default headers values
func NewWorkflowsGetAllTasksDefault(code int) *WorkflowsGetAllTasksDefault {
	return &WorkflowsGetAllTasksDefault{
		_statusCode: code,
	}
}

/*WorkflowsGetAllTasksDefault handles this case with default header values.

Unexpected error
*/
type WorkflowsGetAllTasksDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the workflows get all tasks default response
func (o *WorkflowsGetAllTasksDefault) Code() int {
	return o._statusCode
}

func (o *WorkflowsGetAllTasksDefault) Error() string {
	return fmt.Sprintf("[GET /workflows/tasks][%d] workflowsGetAllTasks default  %+v", o._statusCode, o.Payload)
}

func (o *WorkflowsGetAllTasksDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*WorkflowsGetAllTasksOKBody workflows get all tasks o k body
swagger:model WorkflowsGetAllTasksOKBody
*/

type WorkflowsGetAllTasksOKBody interface{}
