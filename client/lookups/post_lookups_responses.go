package lookups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/codedellemc/gorackhd/models"
)

// PostLookupsReader is a Reader for the PostLookups structure.
type PostLookupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *PostLookupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostLookupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostLookupsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewPostLookupsOK creates a PostLookupsOK with default headers values
func NewPostLookupsOK() *PostLookupsOK {
	return &PostLookupsOK{}
}

/*PostLookupsOK handles this case with default header values.

waterline response
*/
type PostLookupsOK struct {
	Payload PostLookupsOKBodyBody
}

func (o *PostLookupsOK) Error() string {
	return fmt.Sprintf("[POST /lookups][%d] postLookupsOK  %+v", 200, o.Payload)
}

func (o *PostLookupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLookupsDefault creates a PostLookupsDefault with default headers values
func NewPostLookupsDefault(code int) *PostLookupsDefault {
	return &PostLookupsDefault{
		_statusCode: code,
	}
}

/*PostLookupsDefault handles this case with default header values.

Unexpected error
*/
type PostLookupsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post lookups default response
func (o *PostLookupsDefault) Code() int {
	return o._statusCode
}

func (o *PostLookupsDefault) Error() string {
	return fmt.Sprintf("[POST /lookups][%d] PostLookups default  %+v", o._statusCode, o.Payload)
}

func (o *PostLookupsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostLookupsOKBodyBody post lookups o k body body

swagger:model PostLookupsOKBodyBody
*/
type PostLookupsOKBodyBody interface{}
