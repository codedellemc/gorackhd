// Code generated by go-swagger; DO NOT EDIT.

package nodes

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

// NewNodesMasterDelTagByIDParams creates a new NodesMasterDelTagByIDParams object
// with the default values initialized.
func NewNodesMasterDelTagByIDParams() *NodesMasterDelTagByIDParams {
	var ()
	return &NodesMasterDelTagByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewNodesMasterDelTagByIDParamsWithTimeout creates a new NodesMasterDelTagByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewNodesMasterDelTagByIDParamsWithTimeout(timeout time.Duration) *NodesMasterDelTagByIDParams {
	var ()
	return &NodesMasterDelTagByIDParams{

		timeout: timeout,
	}
}

// NewNodesMasterDelTagByIDParamsWithContext creates a new NodesMasterDelTagByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewNodesMasterDelTagByIDParamsWithContext(ctx context.Context) *NodesMasterDelTagByIDParams {
	var ()
	return &NodesMasterDelTagByIDParams{

		Context: ctx,
	}
}

// NewNodesMasterDelTagByIDParamsWithHTTPClient creates a new NodesMasterDelTagByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewNodesMasterDelTagByIDParamsWithHTTPClient(client *http.Client) *NodesMasterDelTagByIDParams {
	var ()
	return &NodesMasterDelTagByIDParams{
		HTTPClient: client,
	}
}

/*NodesMasterDelTagByIDParams contains all the parameters to send to the API endpoint
for the nodes master del tag by Id operation typically these are written to a http.Request
*/
type NodesMasterDelTagByIDParams struct {

	/*TagName
	  The tag identifier

	*/
	TagName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the nodes master del tag by Id params
func (o *NodesMasterDelTagByIDParams) WithTimeout(timeout time.Duration) *NodesMasterDelTagByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the nodes master del tag by Id params
func (o *NodesMasterDelTagByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the nodes master del tag by Id params
func (o *NodesMasterDelTagByIDParams) WithContext(ctx context.Context) *NodesMasterDelTagByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the nodes master del tag by Id params
func (o *NodesMasterDelTagByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the nodes master del tag by Id params
func (o *NodesMasterDelTagByIDParams) WithHTTPClient(client *http.Client) *NodesMasterDelTagByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the nodes master del tag by Id params
func (o *NodesMasterDelTagByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTagName adds the tagName to the nodes master del tag by Id params
func (o *NodesMasterDelTagByIDParams) WithTagName(tagName string) *NodesMasterDelTagByIDParams {
	o.SetTagName(tagName)
	return o
}

// SetTagName adds the tagName to the nodes master del tag by Id params
func (o *NodesMasterDelTagByIDParams) SetTagName(tagName string) {
	o.TagName = tagName
}

// WriteToRequest writes these params to a swagger request
func (o *NodesMasterDelTagByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
