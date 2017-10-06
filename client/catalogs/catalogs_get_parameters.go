// Code generated by go-swagger; DO NOT EDIT.

package catalogs

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

// NewCatalogsGetParams creates a new CatalogsGetParams object
// with the default values initialized.
func NewCatalogsGetParams() *CatalogsGetParams {
	var ()
	return &CatalogsGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCatalogsGetParamsWithTimeout creates a new CatalogsGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCatalogsGetParamsWithTimeout(timeout time.Duration) *CatalogsGetParams {
	var ()
	return &CatalogsGetParams{

		timeout: timeout,
	}
}

// NewCatalogsGetParamsWithContext creates a new CatalogsGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewCatalogsGetParamsWithContext(ctx context.Context) *CatalogsGetParams {
	var ()
	return &CatalogsGetParams{

		Context: ctx,
	}
}

// NewCatalogsGetParamsWithHTTPClient creates a new CatalogsGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCatalogsGetParamsWithHTTPClient(client *http.Client) *CatalogsGetParams {
	var ()
	return &CatalogsGetParams{
		HTTPClient: client,
	}
}

/*CatalogsGetParams contains all the parameters to send to the API endpoint
for the catalogs get operation typically these are written to a http.Request
*/
type CatalogsGetParams struct {

	/*Query
	  A query string containing catalog properties to search

	*/
	Query *string
	/*Sort
	  Query string specifying properties to sort with

	*/
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the catalogs get params
func (o *CatalogsGetParams) WithTimeout(timeout time.Duration) *CatalogsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the catalogs get params
func (o *CatalogsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the catalogs get params
func (o *CatalogsGetParams) WithContext(ctx context.Context) *CatalogsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the catalogs get params
func (o *CatalogsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the catalogs get params
func (o *CatalogsGetParams) WithHTTPClient(client *http.Client) *CatalogsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the catalogs get params
func (o *CatalogsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithQuery adds the query to the catalogs get params
func (o *CatalogsGetParams) WithQuery(query *string) *CatalogsGetParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the catalogs get params
func (o *CatalogsGetParams) SetQuery(query *string) {
	o.Query = query
}

// WithSort adds the sort to the catalogs get params
func (o *CatalogsGetParams) WithSort(sort *string) *CatalogsGetParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the catalogs get params
func (o *CatalogsGetParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *CatalogsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
