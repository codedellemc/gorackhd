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

// NewFilesMd5GetParams creates a new FilesMd5GetParams object
// with the default values initialized.
func NewFilesMd5GetParams() *FilesMd5GetParams {
	var ()
	return &FilesMd5GetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFilesMd5GetParamsWithTimeout creates a new FilesMd5GetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFilesMd5GetParamsWithTimeout(timeout time.Duration) *FilesMd5GetParams {
	var ()
	return &FilesMd5GetParams{

		timeout: timeout,
	}
}

// NewFilesMd5GetParamsWithContext creates a new FilesMd5GetParams object
// with the default values initialized, and the ability to set a context for a request
func NewFilesMd5GetParamsWithContext(ctx context.Context) *FilesMd5GetParams {
	var ()
	return &FilesMd5GetParams{

		Context: ctx,
	}
}

// NewFilesMd5GetParamsWithHTTPClient creates a new FilesMd5GetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFilesMd5GetParamsWithHTTPClient(client *http.Client) *FilesMd5GetParams {
	var ()
	return &FilesMd5GetParams{
		HTTPClient: client,
	}
}

/*FilesMd5GetParams contains all the parameters to send to the API endpoint
for the files md5 get operation typically these are written to a http.Request
*/
type FilesMd5GetParams struct {

	/*Filename
	  File name of a file as provided when you originally stored it

	*/
	Filename string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the files md5 get params
func (o *FilesMd5GetParams) WithTimeout(timeout time.Duration) *FilesMd5GetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the files md5 get params
func (o *FilesMd5GetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the files md5 get params
func (o *FilesMd5GetParams) WithContext(ctx context.Context) *FilesMd5GetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the files md5 get params
func (o *FilesMd5GetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the files md5 get params
func (o *FilesMd5GetParams) WithHTTPClient(client *http.Client) *FilesMd5GetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the files md5 get params
func (o *FilesMd5GetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilename adds the filename to the files md5 get params
func (o *FilesMd5GetParams) WithFilename(filename string) *FilesMd5GetParams {
	o.SetFilename(filename)
	return o
}

// SetFilename adds the filename to the files md5 get params
func (o *FilesMd5GetParams) SetFilename(filename string) {
	o.Filename = filename
}

// WriteToRequest writes these params to a swagger request
func (o *FilesMd5GetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
