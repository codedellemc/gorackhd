package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPostNodesIdentifierObmIdentifyParams creates a new PostNodesIdentifierObmIdentifyParams object
// with the default values initialized.
func NewPostNodesIdentifierObmIdentifyParams() *PostNodesIdentifierObmIdentifyParams {
	var ()
	return &PostNodesIdentifierObmIdentifyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostNodesIdentifierObmIdentifyParamsWithTimeout creates a new PostNodesIdentifierObmIdentifyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostNodesIdentifierObmIdentifyParamsWithTimeout(timeout time.Duration) *PostNodesIdentifierObmIdentifyParams {
	var ()
	return &PostNodesIdentifierObmIdentifyParams{

		timeout: timeout,
	}
}

/*PostNodesIdentifierObmIdentifyParams contains all the parameters to send to the API endpoint
for the post nodes identifier obm identify operation typically these are written to a http.Request
*/
type PostNodesIdentifierObmIdentifyParams struct {

	/*Body
	  obm settings to apply.


	*/
	Body *bool
	/*Identifier
	  Mac addresses and unique aliases to identify the node by, |
	expect a string or an array of strings.


	*/
	Identifier string

	timeout time.Duration
}

// WithBody adds the body to the post nodes identifier obm identify params
func (o *PostNodesIdentifierObmIdentifyParams) WithBody(body *bool) *PostNodesIdentifierObmIdentifyParams {
	o.Body = body
	return o
}

// WithIdentifier adds the identifier to the post nodes identifier obm identify params
func (o *PostNodesIdentifierObmIdentifyParams) WithIdentifier(identifier string) *PostNodesIdentifierObmIdentifyParams {
	o.Identifier = identifier
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PostNodesIdentifierObmIdentifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param identifier
	if err := r.SetPathParam("identifier", o.Identifier); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
