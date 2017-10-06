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

// NewSkusIDGetNodesParams creates a new SkusIDGetNodesParams object
// with the default values initialized.
func NewSkusIDGetNodesParams() *SkusIDGetNodesParams {
	var ()
	return &SkusIDGetNodesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSkusIDGetNodesParamsWithTimeout creates a new SkusIDGetNodesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSkusIDGetNodesParamsWithTimeout(timeout time.Duration) *SkusIDGetNodesParams {
	var ()
	return &SkusIDGetNodesParams{

		timeout: timeout,
	}
}

// NewSkusIDGetNodesParamsWithContext creates a new SkusIDGetNodesParams object
// with the default values initialized, and the ability to set a context for a request
func NewSkusIDGetNodesParamsWithContext(ctx context.Context) *SkusIDGetNodesParams {
	var ()
	return &SkusIDGetNodesParams{

		Context: ctx,
	}
}

// NewSkusIDGetNodesParamsWithHTTPClient creates a new SkusIDGetNodesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSkusIDGetNodesParamsWithHTTPClient(client *http.Client) *SkusIDGetNodesParams {
	var ()
	return &SkusIDGetNodesParams{
		HTTPClient: client,
	}
}

/*SkusIDGetNodesParams contains all the parameters to send to the API endpoint
for the skus Id get nodes operation typically these are written to a http.Request
*/
type SkusIDGetNodesParams struct {

	/*Identifier
	  The SKU identifier

	*/
	Identifier string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the skus Id get nodes params
func (o *SkusIDGetNodesParams) WithTimeout(timeout time.Duration) *SkusIDGetNodesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the skus Id get nodes params
func (o *SkusIDGetNodesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the skus Id get nodes params
func (o *SkusIDGetNodesParams) WithContext(ctx context.Context) *SkusIDGetNodesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the skus Id get nodes params
func (o *SkusIDGetNodesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the skus Id get nodes params
func (o *SkusIDGetNodesParams) WithHTTPClient(client *http.Client) *SkusIDGetNodesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the skus Id get nodes params
func (o *SkusIDGetNodesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIdentifier adds the identifier to the skus Id get nodes params
func (o *SkusIDGetNodesParams) WithIdentifier(identifier string) *SkusIDGetNodesParams {
	o.SetIdentifier(identifier)
	return o
}

// SetIdentifier adds the identifier to the skus Id get nodes params
func (o *SkusIDGetNodesParams) SetIdentifier(identifier string) {
	o.Identifier = identifier
}

// WriteToRequest writes these params to a swagger request
func (o *SkusIDGetNodesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
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
