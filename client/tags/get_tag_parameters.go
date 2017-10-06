// Code generated by go-swagger; DO NOT EDIT.

package tags

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

// NewGetTagParams creates a new GetTagParams object
// with the default values initialized.
func NewGetTagParams() *GetTagParams {
	var ()
	return &GetTagParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTagParamsWithTimeout creates a new GetTagParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTagParamsWithTimeout(timeout time.Duration) *GetTagParams {
	var ()
	return &GetTagParams{

		timeout: timeout,
	}
}

// NewGetTagParamsWithContext creates a new GetTagParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTagParamsWithContext(ctx context.Context) *GetTagParams {
	var ()
	return &GetTagParams{

		Context: ctx,
	}
}

// NewGetTagParamsWithHTTPClient creates a new GetTagParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTagParamsWithHTTPClient(client *http.Client) *GetTagParams {
	var ()
	return &GetTagParams{
		HTTPClient: client,
	}
}

/*GetTagParams contains all the parameters to send to the API endpoint
for the get tag operation typically these are written to a http.Request
*/
type GetTagParams struct {

	/*TagName
	  The tag identifier

	*/
	TagName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get tag params
func (o *GetTagParams) WithTimeout(timeout time.Duration) *GetTagParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get tag params
func (o *GetTagParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get tag params
func (o *GetTagParams) WithContext(ctx context.Context) *GetTagParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get tag params
func (o *GetTagParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get tag params
func (o *GetTagParams) WithHTTPClient(client *http.Client) *GetTagParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get tag params
func (o *GetTagParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTagName adds the tagName to the get tag params
func (o *GetTagParams) WithTagName(tagName string) *GetTagParams {
	o.SetTagName(tagName)
	return o
}

// SetTagName adds the tagName to the get tag params
func (o *GetTagParams) SetTagName(tagName string) {
	o.TagName = tagName
}

// WriteToRequest writes these params to a swagger request
func (o *GetTagParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param tagName
	if err := r.SetPathParam("tagName", o.TagName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
