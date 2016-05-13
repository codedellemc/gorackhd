package dhcp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/emccode/gorackhd/models"
)

// DeleteNodesMacaddressDhcpWhitelistReader is a Reader for the DeleteNodesMacaddressDhcpWhitelist structure.
type DeleteNodesMacaddressDhcpWhitelistReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *DeleteNodesMacaddressDhcpWhitelistReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteNodesMacaddressDhcpWhitelistNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewDeleteNodesMacaddressDhcpWhitelistNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteNodesMacaddressDhcpWhitelistDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewDeleteNodesMacaddressDhcpWhitelistNoContent creates a DeleteNodesMacaddressDhcpWhitelistNoContent with default headers values
func NewDeleteNodesMacaddressDhcpWhitelistNoContent() *DeleteNodesMacaddressDhcpWhitelistNoContent {
	return &DeleteNodesMacaddressDhcpWhitelistNoContent{}
}

/*DeleteNodesMacaddressDhcpWhitelistNoContent handles this case with default header values.

delete completed successfully

*/
type DeleteNodesMacaddressDhcpWhitelistNoContent struct {
}

func (o *DeleteNodesMacaddressDhcpWhitelistNoContent) Error() string {
	return fmt.Sprintf("[DELETE /nodes/{macaddress}/dhcp/whitelist][%d] deleteNodesMacaddressDhcpWhitelistNoContent ", 204)
}

func (o *DeleteNodesMacaddressDhcpWhitelistNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteNodesMacaddressDhcpWhitelistNotFound creates a DeleteNodesMacaddressDhcpWhitelistNotFound with default headers values
func NewDeleteNodesMacaddressDhcpWhitelistNotFound() *DeleteNodesMacaddressDhcpWhitelistNotFound {
	return &DeleteNodesMacaddressDhcpWhitelistNotFound{}
}

/*DeleteNodesMacaddressDhcpWhitelistNotFound handles this case with default header values.

The node with the identifier was not found

*/
type DeleteNodesMacaddressDhcpWhitelistNotFound struct {
	Payload *models.Error
}

func (o *DeleteNodesMacaddressDhcpWhitelistNotFound) Error() string {
	return fmt.Sprintf("[DELETE /nodes/{macaddress}/dhcp/whitelist][%d] deleteNodesMacaddressDhcpWhitelistNotFound  %+v", 404, o.Payload)
}

func (o *DeleteNodesMacaddressDhcpWhitelistNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNodesMacaddressDhcpWhitelistDefault creates a DeleteNodesMacaddressDhcpWhitelistDefault with default headers values
func NewDeleteNodesMacaddressDhcpWhitelistDefault(code int) *DeleteNodesMacaddressDhcpWhitelistDefault {
	return &DeleteNodesMacaddressDhcpWhitelistDefault{
		_statusCode: code,
	}
}

/*DeleteNodesMacaddressDhcpWhitelistDefault handles this case with default header values.

Unexpected error
*/
type DeleteNodesMacaddressDhcpWhitelistDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete nodes macaddress dhcp whitelist default response
func (o *DeleteNodesMacaddressDhcpWhitelistDefault) Code() int {
	return o._statusCode
}

func (o *DeleteNodesMacaddressDhcpWhitelistDefault) Error() string {
	return fmt.Sprintf("[DELETE /nodes/{macaddress}/dhcp/whitelist][%d] DeleteNodesMacaddressDhcpWhitelist default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteNodesMacaddressDhcpWhitelistDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
