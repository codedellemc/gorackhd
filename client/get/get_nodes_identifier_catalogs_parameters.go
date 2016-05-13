package get

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetNodesIdentifierCatalogsParams creates a new GetNodesIdentifierCatalogsParams object
// with the default values initialized.
func NewGetNodesIdentifierCatalogsParams() *GetNodesIdentifierCatalogsParams {
	var ()
	return &GetNodesIdentifierCatalogsParams{}
}

/*GetNodesIdentifierCatalogsParams contains all the parameters to send to the API endpoint
for the get nodes identifier catalogs operation typically these are written to a http.Request
*/
type GetNodesIdentifierCatalogsParams struct {

	/*Identifier
	  Mac addresses and unique aliases to identify the node by, |
	expect a string or an array of strings.


	*/
	Identifier string
}

// WithIdentifier adds the identifier to the get nodes identifier catalogs params
func (o *GetNodesIdentifierCatalogsParams) WithIdentifier(Identifier string) *GetNodesIdentifierCatalogsParams {
	o.Identifier = Identifier
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetNodesIdentifierCatalogsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
