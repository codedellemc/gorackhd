// Code generated by go-swagger; DO NOT EDIT.

package lookups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/spiegela/gorackhd/models"
)

// NewLookupsPatchByIDParams creates a new LookupsPatchByIDParams object
// with the default values initialized.
func NewLookupsPatchByIDParams() *LookupsPatchByIDParams {
	var ()
	return &LookupsPatchByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLookupsPatchByIDParamsWithTimeout creates a new LookupsPatchByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLookupsPatchByIDParamsWithTimeout(timeout time.Duration) *LookupsPatchByIDParams {
	var ()
	return &LookupsPatchByIDParams{

		timeout: timeout,
	}
}

// NewLookupsPatchByIDParamsWithContext creates a new LookupsPatchByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewLookupsPatchByIDParamsWithContext(ctx context.Context) *LookupsPatchByIDParams {
	var ()
	return &LookupsPatchByIDParams{

		Context: ctx,
	}
}

// NewLookupsPatchByIDParamsWithHTTPClient creates a new LookupsPatchByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLookupsPatchByIDParamsWithHTTPClient(client *http.Client) *LookupsPatchByIDParams {
	var ()
	return &LookupsPatchByIDParams{
		HTTPClient: client,
	}
}

/*LookupsPatchByIDParams contains all the parameters to send to the API endpoint
for the lookups patch by Id operation typically these are written to a http.Request
*/
type LookupsPatchByIDParams struct {

	/*Body
	  The lookup properties to be modified

	*/
	Body *models.LookupBase
	/*ID
	  The id of the lookup to patch

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the lookups patch by Id params
func (o *LookupsPatchByIDParams) WithTimeout(timeout time.Duration) *LookupsPatchByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the lookups patch by Id params
func (o *LookupsPatchByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the lookups patch by Id params
func (o *LookupsPatchByIDParams) WithContext(ctx context.Context) *LookupsPatchByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the lookups patch by Id params
func (o *LookupsPatchByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the lookups patch by Id params
func (o *LookupsPatchByIDParams) WithHTTPClient(client *http.Client) *LookupsPatchByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the lookups patch by Id params
func (o *LookupsPatchByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the lookups patch by Id params
func (o *LookupsPatchByIDParams) WithBody(body *models.LookupBase) *LookupsPatchByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the lookups patch by Id params
func (o *LookupsPatchByIDParams) SetBody(body *models.LookupBase) {
	o.Body = body
}

// WithID adds the id to the lookups patch by Id params
func (o *LookupsPatchByIDParams) WithID(id string) *LookupsPatchByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the lookups patch by Id params
func (o *LookupsPatchByIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *LookupsPatchByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models.LookupBase)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
