package obms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/spiegela/gorackhd/models"
)

// NewObmsPutParams creates a new ObmsPutParams object
// with the default values initialized.
func NewObmsPutParams() *ObmsPutParams {
	var ()
	return &ObmsPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewObmsPutParamsWithTimeout creates a new ObmsPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewObmsPutParamsWithTimeout(timeout time.Duration) *ObmsPutParams {
	var ()
	return &ObmsPutParams{

		timeout: timeout,
	}
}

/*ObmsPutParams contains all the parameters to send to the API endpoint
for the obms put operation typically these are written to a http.Request
*/
type ObmsPutParams struct {

	/*Body
	  The OBM settings information to create

	*/
	Body *models.IPMIObmServiceObm

	timeout time.Duration
}

// WithBody adds the body to the obms put params
func (o *ObmsPutParams) WithBody(body *models.IPMIObmServiceObm) *ObmsPutParams {
	o.Body = body
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *ObmsPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Body == nil {
		o.Body = new(models.IPMIObmServiceObm)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}