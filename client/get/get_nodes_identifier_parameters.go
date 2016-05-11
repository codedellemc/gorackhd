package get

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetNodesIdentifierParams creates a new GetNodesIdentifierParams object
// with the default values initialized.
func NewGetNodesIdentifierParams() *GetNodesIdentifierParams {
	var ()
	return &GetNodesIdentifierParams{}
}

/*GetNodesIdentifierParams contains all the parameters to send to the API endpoint
for the get nodes identifier operation typically these are written to a http.Request
*/
type GetNodesIdentifierParams struct {

	/*Identifier
	  Mac addresses and unique aliases to identify the node by, |
	expect a string or an array of strings.


	*/
	Identifier string
}

// WithIdentifier adds the identifier to the get nodes identifier params
func (o *GetNodesIdentifierParams) WithIdentifier(Identifier string) *GetNodesIdentifierParams {
	o.Identifier = Identifier
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetNodesIdentifierParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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