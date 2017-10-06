// Code generated by go-swagger; DO NOT EDIT.

package skus

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

// NewSkusGetParams creates a new SkusGetParams object
// with the default values initialized.
func NewSkusGetParams() *SkusGetParams {
	var ()
	return &SkusGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSkusGetParamsWithTimeout creates a new SkusGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSkusGetParamsWithTimeout(timeout time.Duration) *SkusGetParams {
	var ()
	return &SkusGetParams{

		timeout: timeout,
	}
}

// NewSkusGetParamsWithContext creates a new SkusGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewSkusGetParamsWithContext(ctx context.Context) *SkusGetParams {
	var ()
	return &SkusGetParams{

		Context: ctx,
	}
}

// NewSkusGetParamsWithHTTPClient creates a new SkusGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSkusGetParamsWithHTTPClient(client *http.Client) *SkusGetParams {
	var ()
	return &SkusGetParams{
		HTTPClient: client,
	}
}

/*SkusGetParams contains all the parameters to send to the API endpoint
for the skus get operation typically these are written to a http.Request
*/
type SkusGetParams struct {

	/*Query
	  SKU properties to search

	*/
	Query *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the skus get params
func (o *SkusGetParams) WithTimeout(timeout time.Duration) *SkusGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the skus get params
func (o *SkusGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the skus get params
func (o *SkusGetParams) WithContext(ctx context.Context) *SkusGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the skus get params
func (o *SkusGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the skus get params
func (o *SkusGetParams) WithHTTPClient(client *http.Client) *SkusGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the skus get params
func (o *SkusGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithQuery adds the query to the skus get params
func (o *SkusGetParams) WithQuery(query *string) *SkusGetParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the skus get params
func (o *SkusGetParams) SetQuery(query *string) {
	o.Query = query
}

// WriteToRequest writes these params to a swagger request
func (o *SkusGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Query != nil {

		// query param query
		var qrQuery string
		if o.Query != nil {
			qrQuery = *o.Query
		}
		qQuery := qrQuery
		if qQuery != "" {
			if err := r.SetQueryParam("query", qQuery); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
