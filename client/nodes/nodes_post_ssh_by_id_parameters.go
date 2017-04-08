package nodes

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

// NewNodesPostSSHByIDParams creates a new NodesPostSSHByIDParams object
// with the default values initialized.
func NewNodesPostSSHByIDParams() *NodesPostSSHByIDParams {
	var ()
	return &NodesPostSSHByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewNodesPostSSHByIDParamsWithTimeout creates a new NodesPostSSHByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewNodesPostSSHByIDParamsWithTimeout(timeout time.Duration) *NodesPostSSHByIDParams {
	var ()
	return &NodesPostSSHByIDParams{

		timeout: timeout,
	}
}

/*NodesPostSSHByIDParams contains all the parameters to send to the API endpoint
for the nodes post Ssh by Id operation typically these are written to a http.Request
*/
type NodesPostSSHByIDParams struct {

	/*Body
	  The ssh properties to create

	*/
	Body *models.SSHIbmServiceIbm
	/*Identifier
	  The node identifier

	*/
	Identifier string

	timeout time.Duration
}

// WithBody adds the body to the nodes post Ssh by Id params
func (o *NodesPostSSHByIDParams) WithBody(body *models.SSHIbmServiceIbm) *NodesPostSSHByIDParams {
	o.Body = body
	return o
}

// WithIdentifier adds the identifier to the nodes post Ssh by Id params
func (o *NodesPostSSHByIDParams) WithIdentifier(identifier string) *NodesPostSSHByIDParams {
	o.Identifier = identifier
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *NodesPostSSHByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Body == nil {
		o.Body = new(models.SSHIbmServiceIbm)
	}

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