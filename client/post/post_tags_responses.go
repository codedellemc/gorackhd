package post

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/codedellemc/gorackhd/models"
)

// PostTagsReader is a Reader for the PostTags structure.
type PostTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *PostTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostTagsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 409:
		result := NewPostTagsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPostTagsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewPostTagsCreated creates a PostTagsCreated with default headers values
func NewPostTagsCreated() *PostTagsCreated {
	return &PostTagsCreated{}
}

/*PostTagsCreated handles this case with default header values.

tag accepted successfully
*/
type PostTagsCreated struct {
}

func (o *PostTagsCreated) Error() string {
	return fmt.Sprintf("[POST /tags][%d] postTagsCreated ", 201)
}

func (o *PostTagsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostTagsConflict creates a PostTagsConflict with default headers values
func NewPostTagsConflict() *PostTagsConflict {
	return &PostTagsConflict{}
}

/*PostTagsConflict handles this case with default header values.

tag already exists
*/
type PostTagsConflict struct {
	Payload *models.Error
}

func (o *PostTagsConflict) Error() string {
	return fmt.Sprintf("[POST /tags][%d] postTagsConflict  %+v", 409, o.Payload)
}

func (o *PostTagsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTagsDefault creates a PostTagsDefault with default headers values
func NewPostTagsDefault(code int) *PostTagsDefault {
	return &PostTagsDefault{
		_statusCode: code,
	}
}

/*PostTagsDefault handles this case with default header values.

Unexpected error
*/
type PostTagsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post tags default response
func (o *PostTagsDefault) Code() int {
	return o._statusCode
}

func (o *PostTagsDefault) Error() string {
	return fmt.Sprintf("[POST /tags][%d] PostTags default  %+v", o._statusCode, o.Payload)
}

func (o *PostTagsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
