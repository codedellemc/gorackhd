package post

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPostTagsParams creates a new PostTagsParams object
// with the default values initialized.
func NewPostTagsParams() *PostTagsParams {
	var ()
	return &PostTagsParams{}
}

/*PostTagsParams contains all the parameters to send to the API endpoint
for the post tags operation typically these are written to a http.Request
*/
type PostTagsParams struct {

	/*Body*/
	Body interface{}
}

// WithBody adds the body to the post tags params
func (o *PostTagsParams) WithBody(Body interface{}) *PostTagsParams {
	o.Body = Body
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PostTagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
