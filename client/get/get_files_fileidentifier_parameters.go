package get

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetFilesFileidentifierParams creates a new GetFilesFileidentifierParams object
// with the default values initialized.
func NewGetFilesFileidentifierParams() *GetFilesFileidentifierParams {
	var ()
	return &GetFilesFileidentifierParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFilesFileidentifierParamsWithTimeout creates a new GetFilesFileidentifierParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFilesFileidentifierParamsWithTimeout(timeout time.Duration) *GetFilesFileidentifierParams {
	var ()
	return &GetFilesFileidentifierParams{

		timeout: timeout,
	}
}

/*GetFilesFileidentifierParams contains all the parameters to send to the API endpoint
for the get files fileidentifier operation typically these are written to a http.Request
*/
type GetFilesFileidentifierParams struct {

	/*Fileidentifier
	  uuid of a file as provided when you originally stored it.

	*/
	Fileidentifier string

	timeout time.Duration
}

// WithFileidentifier adds the fileidentifier to the get files fileidentifier params
func (o *GetFilesFileidentifierParams) WithFileidentifier(fileidentifier string) *GetFilesFileidentifierParams {
	o.Fileidentifier = fileidentifier
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetFilesFileidentifierParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param fileidentifier
	if err := r.SetPathParam("fileidentifier", o.Fileidentifier); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
