// Code generated by go-swagger; DO NOT EDIT.

package files

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

// NewFilesMetadataGetParams creates a new FilesMetadataGetParams object
// with the default values initialized.
func NewFilesMetadataGetParams() *FilesMetadataGetParams {
	var ()
	return &FilesMetadataGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFilesMetadataGetParamsWithTimeout creates a new FilesMetadataGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFilesMetadataGetParamsWithTimeout(timeout time.Duration) *FilesMetadataGetParams {
	var ()
	return &FilesMetadataGetParams{

		timeout: timeout,
	}
}

// NewFilesMetadataGetParamsWithContext creates a new FilesMetadataGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewFilesMetadataGetParamsWithContext(ctx context.Context) *FilesMetadataGetParams {
	var ()
	return &FilesMetadataGetParams{

		Context: ctx,
	}
}

// NewFilesMetadataGetParamsWithHTTPClient creates a new FilesMetadataGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFilesMetadataGetParamsWithHTTPClient(client *http.Client) *FilesMetadataGetParams {
	var ()
	return &FilesMetadataGetParams{
		HTTPClient: client,
	}
}

/*FilesMetadataGetParams contains all the parameters to send to the API endpoint
for the files metadata get operation typically these are written to a http.Request
*/
type FilesMetadataGetParams struct {

	/*Filename
	  The name of a file as provided when you originally stored it

	*/
	Filename string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the files metadata get params
func (o *FilesMetadataGetParams) WithTimeout(timeout time.Duration) *FilesMetadataGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the files metadata get params
func (o *FilesMetadataGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the files metadata get params
func (o *FilesMetadataGetParams) WithContext(ctx context.Context) *FilesMetadataGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the files metadata get params
func (o *FilesMetadataGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the files metadata get params
func (o *FilesMetadataGetParams) WithHTTPClient(client *http.Client) *FilesMetadataGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the files metadata get params
func (o *FilesMetadataGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilename adds the filename to the files metadata get params
func (o *FilesMetadataGetParams) WithFilename(filename string) *FilesMetadataGetParams {
	o.SetFilename(filename)
	return o
}

// SetFilename adds the filename to the files metadata get params
func (o *FilesMetadataGetParams) SetFilename(filename string) {
	o.Filename = filename
}

// WriteToRequest writes these params to a swagger request
func (o *FilesMetadataGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param filename
	if err := r.SetPathParam("filename", o.Filename); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
