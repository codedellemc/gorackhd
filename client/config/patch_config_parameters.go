package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPatchConfigParams creates a new PatchConfigParams object
// with the default values initialized.
func NewPatchConfigParams() *PatchConfigParams {
	var ()
	return &PatchConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchConfigParamsWithTimeout creates a new PatchConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchConfigParamsWithTimeout(timeout time.Duration) *PatchConfigParams {
	var ()
	return &PatchConfigParams{

		timeout: timeout,
	}
}

/*PatchConfigParams contains all the parameters to send to the API endpoint
for the patch config operation typically these are written to a http.Request
*/
type PatchConfigParams struct {

	/*Body*/
	Body interface{}

	timeout time.Duration
}

// WithBody adds the body to the patch config params
func (o *PatchConfigParams) WithBody(body interface{}) *PatchConfigParams {
	o.Body = body
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PatchConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
