package skus

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/codedellemc/gorackhd/models"
)

// GetSkusIdentifierReader is a Reader for the GetSkusIdentifier structure.
type GetSkusIdentifierReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetSkusIdentifierReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSkusIdentifierOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetSkusIdentifierNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetSkusIdentifierDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetSkusIdentifierOK creates a GetSkusIdentifierOK with default headers values
func NewGetSkusIdentifierOK() *GetSkusIdentifierOK {
	return &GetSkusIdentifierOK{}
}

/*GetSkusIdentifierOK handles this case with default header values.

return sku

*/
type GetSkusIdentifierOK struct {
	Payload *models.Sku
}

func (o *GetSkusIdentifierOK) Error() string {
	return fmt.Sprintf("[GET /skus/{identifier}][%d] getSkusIdentifierOK  %+v", 200, o.Payload)
}

func (o *GetSkusIdentifierOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Sku)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSkusIdentifierNotFound creates a GetSkusIdentifierNotFound with default headers values
func NewGetSkusIdentifierNotFound() *GetSkusIdentifierNotFound {
	return &GetSkusIdentifierNotFound{}
}

/*GetSkusIdentifierNotFound handles this case with default header values.

There is no sku with identifier.

*/
type GetSkusIdentifierNotFound struct {
	Payload *models.Error
}

func (o *GetSkusIdentifierNotFound) Error() string {
	return fmt.Sprintf("[GET /skus/{identifier}][%d] getSkusIdentifierNotFound  %+v", 404, o.Payload)
}

func (o *GetSkusIdentifierNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSkusIdentifierDefault creates a GetSkusIdentifierDefault with default headers values
func NewGetSkusIdentifierDefault(code int) *GetSkusIdentifierDefault {
	return &GetSkusIdentifierDefault{
		_statusCode: code,
	}
}

/*GetSkusIdentifierDefault handles this case with default header values.

Unexpected error
*/
type GetSkusIdentifierDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get skus identifier default response
func (o *GetSkusIdentifierDefault) Code() int {
	return o._statusCode
}

func (o *GetSkusIdentifierDefault) Error() string {
	return fmt.Sprintf("[GET /skus/{identifier}][%d] GetSkusIdentifier default  %+v", o._statusCode, o.Payload)
}

func (o *GetSkusIdentifierDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
