package get

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetPollersIdentifierDataParams creates a new GetPollersIdentifierDataParams object
// with the default values initialized.
func NewGetPollersIdentifierDataParams() *GetPollersIdentifierDataParams {
	var ()
	return &GetPollersIdentifierDataParams{}
}

/*GetPollersIdentifierDataParams contains all the parameters to send to the API endpoint
for the get pollers identifier data operation typically these are written to a http.Request
*/
type GetPollersIdentifierDataParams struct {

	/*Identifier
	  identifier (ip address or NodeId) for the data from a poller


	*/
	Identifier string
}

// WithIdentifier adds the identifier to the get pollers identifier data params
func (o *GetPollersIdentifierDataParams) WithIdentifier(Identifier string) *GetPollersIdentifierDataParams {
	o.Identifier = Identifier
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetPollersIdentifierDataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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