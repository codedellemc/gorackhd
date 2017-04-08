package ibms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewIbmsDeleteByIDParams creates a new IbmsDeleteByIDParams object
// with the default values initialized.
func NewIbmsDeleteByIDParams() *IbmsDeleteByIDParams {
	var ()
	return &IbmsDeleteByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIbmsDeleteByIDParamsWithTimeout creates a new IbmsDeleteByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIbmsDeleteByIDParamsWithTimeout(timeout time.Duration) *IbmsDeleteByIDParams {
	var ()
	return &IbmsDeleteByIDParams{

		timeout: timeout,
	}
}

/*IbmsDeleteByIDParams contains all the parameters to send to the API endpoint
for the ibms delete by Id operation typically these are written to a http.Request
*/
type IbmsDeleteByIDParams struct {

	/*Identifier
	  The IBMS service identifier

	*/
	Identifier string

	timeout time.Duration
}

// WithIdentifier adds the identifier to the ibms delete by Id params
func (o *IbmsDeleteByIDParams) WithIdentifier(identifier string) *IbmsDeleteByIDParams {
	o.Identifier = identifier
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *IbmsDeleteByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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