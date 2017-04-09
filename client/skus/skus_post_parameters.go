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

	"github.com/spiegela/gorackhd/models"
)

// NewSkusPostParams creates a new SkusPostParams object
// with the default values initialized.
func NewSkusPostParams() *SkusPostParams {
	var ()
	return &SkusPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSkusPostParamsWithTimeout creates a new SkusPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSkusPostParamsWithTimeout(timeout time.Duration) *SkusPostParams {
	var ()
	return &SkusPostParams{

		timeout: timeout,
	}
}

// NewSkusPostParamsWithContext creates a new SkusPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewSkusPostParamsWithContext(ctx context.Context) *SkusPostParams {
	var ()
	return &SkusPostParams{

		Context: ctx,
	}
}

// NewSkusPostParamsWithHTTPClient creates a new SkusPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSkusPostParamsWithHTTPClient(client *http.Client) *SkusPostParams {
	var ()
	return &SkusPostParams{
		HTTPClient: client,
	}
}

/*SkusPostParams contains all the parameters to send to the API endpoint
for the skus post operation typically these are written to a http.Request
*/
type SkusPostParams struct {

	/*Body
	  The properties used to define the new SKU

	*/
	Body *models.Skus20SkusUpsert

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the skus post params
func (o *SkusPostParams) WithTimeout(timeout time.Duration) *SkusPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the skus post params
func (o *SkusPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the skus post params
func (o *SkusPostParams) WithContext(ctx context.Context) *SkusPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the skus post params
func (o *SkusPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the skus post params
func (o *SkusPostParams) WithHTTPClient(client *http.Client) *SkusPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the skus post params
func (o *SkusPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the skus post params
func (o *SkusPostParams) WithBody(body *models.Skus20SkusUpsert) *SkusPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the skus post params
func (o *SkusPostParams) SetBody(body *models.Skus20SkusUpsert) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SkusPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models.Skus20SkusUpsert)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
