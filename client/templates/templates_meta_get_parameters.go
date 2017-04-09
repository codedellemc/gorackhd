package templates

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

// NewTemplatesMetaGetParams creates a new TemplatesMetaGetParams object
// with the default values initialized.
func NewTemplatesMetaGetParams() *TemplatesMetaGetParams {
	var ()
	return &TemplatesMetaGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTemplatesMetaGetParamsWithTimeout creates a new TemplatesMetaGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTemplatesMetaGetParamsWithTimeout(timeout time.Duration) *TemplatesMetaGetParams {
	var ()
	return &TemplatesMetaGetParams{

		timeout: timeout,
	}
}

// NewTemplatesMetaGetParamsWithContext creates a new TemplatesMetaGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewTemplatesMetaGetParamsWithContext(ctx context.Context) *TemplatesMetaGetParams {
	var ()
	return &TemplatesMetaGetParams{

		Context: ctx,
	}
}

// NewTemplatesMetaGetParamsWithHTTPClient creates a new TemplatesMetaGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTemplatesMetaGetParamsWithHTTPClient(client *http.Client) *TemplatesMetaGetParams {
	var ()
	return &TemplatesMetaGetParams{
		HTTPClient: client,
	}
}

/*TemplatesMetaGetParams contains all the parameters to send to the API endpoint
for the templates meta get operation typically these are written to a http.Request
*/
type TemplatesMetaGetParams struct {

	/*Sort
	  Query string specifying properties to sort with

	*/
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the templates meta get params
func (o *TemplatesMetaGetParams) WithTimeout(timeout time.Duration) *TemplatesMetaGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the templates meta get params
func (o *TemplatesMetaGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the templates meta get params
func (o *TemplatesMetaGetParams) WithContext(ctx context.Context) *TemplatesMetaGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the templates meta get params
func (o *TemplatesMetaGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the templates meta get params
func (o *TemplatesMetaGetParams) WithHTTPClient(client *http.Client) *TemplatesMetaGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the templates meta get params
func (o *TemplatesMetaGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSort adds the sort to the templates meta get params
func (o *TemplatesMetaGetParams) WithSort(sort *string) *TemplatesMetaGetParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the templates meta get params
func (o *TemplatesMetaGetParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *TemplatesMetaGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Sort != nil {

		// query param sort
		var qrSort string
		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {
			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
