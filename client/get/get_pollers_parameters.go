package get

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetPollersParams creates a new GetPollersParams object
// with the default values initialized.
func NewGetPollersParams() *GetPollersParams {

	return &GetPollersParams{}
}

/*GetPollersParams contains all the parameters to send to the API endpoint
for the get pollers operation typically these are written to a http.Request
*/
type GetPollersParams struct {
}

// WriteToRequest writes these params to a swagger request
func (o *GetPollersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
