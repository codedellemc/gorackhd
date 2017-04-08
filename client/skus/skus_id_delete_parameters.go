package skus

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewSkusIDDeleteParams creates a new SkusIDDeleteParams object
// with the default values initialized.
func NewSkusIDDeleteParams() *SkusIDDeleteParams {
	var ()
	return &SkusIDDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSkusIDDeleteParamsWithTimeout creates a new SkusIDDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSkusIDDeleteParamsWithTimeout(timeout time.Duration) *SkusIDDeleteParams {
	var ()
	return &SkusIDDeleteParams{

		timeout: timeout,
	}
}

/*SkusIDDeleteParams contains all the parameters to send to the API endpoint
for the skus Id delete operation typically these are written to a http.Request
*/
type SkusIDDeleteParams struct {

	/*Identifier
	  The SKU identifier

	*/
	Identifier string

	timeout time.Duration
}

// WithIdentifier adds the identifier to the skus Id delete params
func (o *SkusIDDeleteParams) WithIdentifier(identifier string) *SkusIDDeleteParams {
	o.Identifier = identifier
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *SkusIDDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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