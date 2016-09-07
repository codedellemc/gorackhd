package delete

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/codedellemc/gorackhd/models"
)

// DeleteDhcpLeaseMacReader is a Reader for the DeleteDhcpLeaseMac structure.
type DeleteDhcpLeaseMacReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *DeleteDhcpLeaseMacReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteDhcpLeaseMacOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteDhcpLeaseMacDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewDeleteDhcpLeaseMacOK creates a DeleteDhcpLeaseMacOK with default headers values
func NewDeleteDhcpLeaseMacOK() *DeleteDhcpLeaseMacOK {
	return &DeleteDhcpLeaseMacOK{}
}

/*DeleteDhcpLeaseMacOK handles this case with default header values.

A single lease
*/
type DeleteDhcpLeaseMacOK struct {
	Payload []*models.Lease
}

func (o *DeleteDhcpLeaseMacOK) Error() string {
	return fmt.Sprintf("[DELETE /dhcp/lease/{mac}][%d] deleteDhcpLeaseMacOK  %+v", 200, o.Payload)
}

func (o *DeleteDhcpLeaseMacOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDhcpLeaseMacDefault creates a DeleteDhcpLeaseMacDefault with default headers values
func NewDeleteDhcpLeaseMacDefault(code int) *DeleteDhcpLeaseMacDefault {
	return &DeleteDhcpLeaseMacDefault{
		_statusCode: code,
	}
}

/*DeleteDhcpLeaseMacDefault handles this case with default header values.

NotFound error
*/
type DeleteDhcpLeaseMacDefault struct {
	_statusCode int
}

// Code gets the status code for the delete dhcp lease mac default response
func (o *DeleteDhcpLeaseMacDefault) Code() int {
	return o._statusCode
}

func (o *DeleteDhcpLeaseMacDefault) Error() string {
	return fmt.Sprintf("[DELETE /dhcp/lease/{mac}][%d] DeleteDhcpLeaseMac default ", o._statusCode)
}

func (o *DeleteDhcpLeaseMacDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
