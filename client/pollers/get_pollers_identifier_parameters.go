package pollers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetPollersIdentifierParams creates a new GetPollersIdentifierParams object
// with the default values initialized.
func NewGetPollersIdentifierParams() *GetPollersIdentifierParams {
	var ()
	return &GetPollersIdentifierParams{}
}

/*GetPollersIdentifierParams contains all the parameters to send to the API endpoint
for the get pollers identifier operation typically these are written to a http.Request
*/
type GetPollersIdentifierParams struct {

	/*Identifier
	  poller identifier


	*/
	Identifier string
}

// WithIdentifier adds the identifier to the get pollers identifier params
func (o *GetPollersIdentifierParams) WithIdentifier(Identifier string) *GetPollersIdentifierParams {
	o.Identifier = Identifier
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetPollersIdentifierParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
