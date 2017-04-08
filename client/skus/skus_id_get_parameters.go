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

// NewSkusIDGetParams creates a new SkusIDGetParams object
// with the default values initialized.
func NewSkusIDGetParams() *SkusIDGetParams {
	var ()
	return &SkusIDGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSkusIDGetParamsWithTimeout creates a new SkusIDGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSkusIDGetParamsWithTimeout(timeout time.Duration) *SkusIDGetParams {
	var ()
	return &SkusIDGetParams{

		timeout: timeout,
	}
}

/*SkusIDGetParams contains all the parameters to send to the API endpoint
for the skus Id get operation typically these are written to a http.Request
*/
type SkusIDGetParams struct {

	/*Identifier
	  The SKU identifier

	*/
	Identifier string

	timeout time.Duration
}

// WithIdentifier adds the identifier to the skus Id get params
func (o *SkusIDGetParams) WithIdentifier(identifier string) *SkusIDGetParams {
	o.Identifier = identifier
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *SkusIDGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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