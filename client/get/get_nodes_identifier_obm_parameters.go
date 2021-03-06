package get

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetNodesIdentifierObmParams creates a new GetNodesIdentifierObmParams object
// with the default values initialized.
func NewGetNodesIdentifierObmParams() *GetNodesIdentifierObmParams {
	var ()
	return &GetNodesIdentifierObmParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNodesIdentifierObmParamsWithTimeout creates a new GetNodesIdentifierObmParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNodesIdentifierObmParamsWithTimeout(timeout time.Duration) *GetNodesIdentifierObmParams {
	var ()
	return &GetNodesIdentifierObmParams{

		timeout: timeout,
	}
}

/*GetNodesIdentifierObmParams contains all the parameters to send to the API endpoint
for the get nodes identifier obm operation typically these are written to a http.Request
*/
type GetNodesIdentifierObmParams struct {

	/*Identifier
	  Mac addresses and unique aliases to identify the node by, |
	expect a string or an array of strings.


	*/
	Identifier string

	timeout time.Duration
}

// WithIdentifier adds the identifier to the get nodes identifier obm params
func (o *GetNodesIdentifierObmParams) WithIdentifier(identifier string) *GetNodesIdentifierObmParams {
	o.Identifier = identifier
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetNodesIdentifierObmParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param identifier
	if err := r.SetPathParam("identifier", o.Identifier); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
