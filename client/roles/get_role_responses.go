// Code generated by go-swagger; DO NOT EDIT.

package roles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/spiegela/gorackhd/models"
)

// GetRoleReader is a Reader for the GetRole structure.
type GetRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetRoleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetRoleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetRoleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetRoleOK creates a GetRoleOK with default headers values
func NewGetRoleOK() *GetRoleOK {
	return &GetRoleOK{}
}

/*GetRoleOK handles this case with default header values.

Successfully retrieved the specified role
*/
type GetRoleOK struct {
}

func (o *GetRoleOK) Error() string {
	return fmt.Sprintf("[GET /roles/{name}][%d] getRoleOK ", 200)
}

func (o *GetRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRoleUnauthorized creates a GetRoleUnauthorized with default headers values
func NewGetRoleUnauthorized() *GetRoleUnauthorized {
	return &GetRoleUnauthorized{}
}

/*GetRoleUnauthorized handles this case with default header values.

Unauthorized
*/
type GetRoleUnauthorized struct {
	Payload GetRoleUnauthorizedBody
}

func (o *GetRoleUnauthorized) Error() string {
	return fmt.Sprintf("[GET /roles/{name}][%d] getRoleUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRoleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoleForbidden creates a GetRoleForbidden with default headers values
func NewGetRoleForbidden() *GetRoleForbidden {
	return &GetRoleForbidden{}
}

/*GetRoleForbidden handles this case with default header values.

Forbidden
*/
type GetRoleForbidden struct {
	Payload GetRoleForbiddenBody
}

func (o *GetRoleForbidden) Error() string {
	return fmt.Sprintf("[GET /roles/{name}][%d] getRoleForbidden  %+v", 403, o.Payload)
}

func (o *GetRoleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoleDefault creates a GetRoleDefault with default headers values
func NewGetRoleDefault(code int) *GetRoleDefault {
	return &GetRoleDefault{
		_statusCode: code,
	}
}

/*GetRoleDefault handles this case with default header values.

Unexpected error
*/
type GetRoleDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get role default response
func (o *GetRoleDefault) Code() int {
	return o._statusCode
}

func (o *GetRoleDefault) Error() string {
	return fmt.Sprintf("[GET /roles/{name}][%d] getRole default  %+v", o._statusCode, o.Payload)
}

func (o *GetRoleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetRoleForbiddenBody get role forbidden body
swagger:model GetRoleForbiddenBody
*/

type GetRoleForbiddenBody interface{}

/*GetRoleUnauthorizedBody get role unauthorized body
swagger:model GetRoleUnauthorizedBody
*/

type GetRoleUnauthorizedBody interface{}
