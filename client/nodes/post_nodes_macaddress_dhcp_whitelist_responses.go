package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/emccode/gorackhd/models"
)

// PostNodesMacaddressDhcpWhitelistReader is a Reader for the PostNodesMacaddressDhcpWhitelist structure.
type PostNodesMacaddressDhcpWhitelistReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *PostNodesMacaddressDhcpWhitelistReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostNodesMacaddressDhcpWhitelistCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewPostNodesMacaddressDhcpWhitelistNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPostNodesMacaddressDhcpWhitelistDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewPostNodesMacaddressDhcpWhitelistCreated creates a PostNodesMacaddressDhcpWhitelistCreated with default headers values
func NewPostNodesMacaddressDhcpWhitelistCreated() *PostNodesMacaddressDhcpWhitelistCreated {
	return &PostNodesMacaddressDhcpWhitelistCreated{}
}

/*PostNodesMacaddressDhcpWhitelistCreated handles this case with default header values.

the add was successful and it returns the whitelist

*/
type PostNodesMacaddressDhcpWhitelistCreated struct {
	Payload PostNodesMacaddressDhcpWhitelistCreatedBodyBody
}

func (o *PostNodesMacaddressDhcpWhitelistCreated) Error() string {
	return fmt.Sprintf("[POST /nodes/{macaddress}/dhcp/whitelist][%d] postNodesMacaddressDhcpWhitelistCreated  %+v", 201, o.Payload)
}

func (o *PostNodesMacaddressDhcpWhitelistCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostNodesMacaddressDhcpWhitelistNotFound creates a PostNodesMacaddressDhcpWhitelistNotFound with default headers values
func NewPostNodesMacaddressDhcpWhitelistNotFound() *PostNodesMacaddressDhcpWhitelistNotFound {
	return &PostNodesMacaddressDhcpWhitelistNotFound{}
}

/*PostNodesMacaddressDhcpWhitelistNotFound handles this case with default header values.

The node with the identifier was not found

*/
type PostNodesMacaddressDhcpWhitelistNotFound struct {
	Payload *models.Error
}

func (o *PostNodesMacaddressDhcpWhitelistNotFound) Error() string {
	return fmt.Sprintf("[POST /nodes/{macaddress}/dhcp/whitelist][%d] postNodesMacaddressDhcpWhitelistNotFound  %+v", 404, o.Payload)
}

func (o *PostNodesMacaddressDhcpWhitelistNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostNodesMacaddressDhcpWhitelistDefault creates a PostNodesMacaddressDhcpWhitelistDefault with default headers values
func NewPostNodesMacaddressDhcpWhitelistDefault(code int) *PostNodesMacaddressDhcpWhitelistDefault {
	return &PostNodesMacaddressDhcpWhitelistDefault{
		_statusCode: code,
	}
}

/*PostNodesMacaddressDhcpWhitelistDefault handles this case with default header values.

Unexpected error
*/
type PostNodesMacaddressDhcpWhitelistDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post nodes macaddress dhcp whitelist default response
func (o *PostNodesMacaddressDhcpWhitelistDefault) Code() int {
	return o._statusCode
}

func (o *PostNodesMacaddressDhcpWhitelistDefault) Error() string {
	return fmt.Sprintf("[POST /nodes/{macaddress}/dhcp/whitelist][%d] PostNodesMacaddressDhcpWhitelist default  %+v", o._statusCode, o.Payload)
}

func (o *PostNodesMacaddressDhcpWhitelistDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostNodesMacaddressDhcpWhitelistCreatedBodyBody post nodes macaddress dhcp whitelist created body body

swagger:model PostNodesMacaddressDhcpWhitelistCreatedBodyBody
*/
type PostNodesMacaddressDhcpWhitelistCreatedBodyBody interface{}
