// Code generated by go-swagger; DO NOT EDIT.

package skus

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/spiegela/gorackhd/models"
)

// SkusPutReader is a Reader for the SkusPut structure.
type SkusPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SkusPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewSkusPutCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewSkusPutInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewSkusPutDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSkusPutCreated creates a SkusPutCreated with default headers values
func NewSkusPutCreated() *SkusPutCreated {
	return &SkusPutCreated{}
}

/*SkusPutCreated handles this case with default header values.

Successfully created or updated SKU definition
*/
type SkusPutCreated struct {
	Payload SkusPutCreatedBody
}

func (o *SkusPutCreated) Error() string {
	return fmt.Sprintf("[PUT /skus][%d] skusPutCreated  %+v", 201, o.Payload)
}

func (o *SkusPutCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSkusPutInternalServerError creates a SkusPutInternalServerError with default headers values
func NewSkusPutInternalServerError() *SkusPutInternalServerError {
	return &SkusPutInternalServerError{}
}

/*SkusPutInternalServerError handles this case with default header values.

Update failed
*/
type SkusPutInternalServerError struct {
	Payload *models.Error
}

func (o *SkusPutInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /skus][%d] skusPutInternalServerError  %+v", 500, o.Payload)
}

func (o *SkusPutInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSkusPutDefault creates a SkusPutDefault with default headers values
func NewSkusPutDefault(code int) *SkusPutDefault {
	return &SkusPutDefault{
		_statusCode: code,
	}
}

/*SkusPutDefault handles this case with default header values.

Unexpected error
*/
type SkusPutDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the skus put default response
func (o *SkusPutDefault) Code() int {
	return o._statusCode
}

func (o *SkusPutDefault) Error() string {
	return fmt.Sprintf("[PUT /skus][%d] skusPut default  %+v", o._statusCode, o.Payload)
}

func (o *SkusPutDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*SkusPutCreatedBody skus put created body
swagger:model SkusPutCreatedBody
*/

type SkusPutCreatedBody interface{}
