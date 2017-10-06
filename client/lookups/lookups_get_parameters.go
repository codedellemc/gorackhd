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
)

// NewLookupsGetParams creates a new LookupsGetParams object
// with the default values initialized.
func NewLookupsGetParams() *LookupsGetParams {
	var ()
	return &LookupsGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLookupsGetParamsWithTimeout creates a new LookupsGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLookupsGetParamsWithTimeout(timeout time.Duration) *LookupsGetParams {
	var ()
	return &LookupsGetParams{

		timeout: timeout,
	}
}

// NewLookupsGetParamsWithContext creates a new LookupsGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewLookupsGetParamsWithContext(ctx context.Context) *LookupsGetParams {
	var ()
	return &LookupsGetParams{

		Context: ctx,
	}
}

// NewLookupsGetParamsWithHTTPClient creates a new LookupsGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLookupsGetParamsWithHTTPClient(client *http.Client) *LookupsGetParams {
	var ()
	return &LookupsGetParams{
		HTTPClient: client,
	}
}

/*LookupsGetParams contains all the parameters to send to the API endpoint
for the lookups get operation typically these are written to a http.Request
*/
type LookupsGetParams struct {

	/*Q
	  Query string specifying properties to search for

	*/
	Q *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the lookups get params
func (o *LookupsGetParams) WithTimeout(timeout time.Duration) *LookupsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the lookups get params
func (o *LookupsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the lookups get params
func (o *LookupsGetParams) WithContext(ctx context.Context) *LookupsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the lookups get params
func (o *LookupsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the lookups get params
func (o *LookupsGetParams) WithHTTPClient(client *http.Client) *LookupsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the lookups get params
func (o *LookupsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithQ adds the q to the lookups get params
func (o *LookupsGetParams) WithQ(q *string) *LookupsGetParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the lookups get params
func (o *LookupsGetParams) SetQ(q *string) {
	o.Q = q
}

// WriteToRequest writes these params to a swagger request
func (o *LookupsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Q != nil {

		// query param q
		var qrQ string
		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {
			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
