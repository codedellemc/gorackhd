package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteNodesIdentifierTagsTagnameParams creates a new DeleteNodesIdentifierTagsTagnameParams object
// with the default values initialized.
func NewDeleteNodesIdentifierTagsTagnameParams() *DeleteNodesIdentifierTagsTagnameParams {
	var ()
	return &DeleteNodesIdentifierTagsTagnameParams{}
}

/*DeleteNodesIdentifierTagsTagnameParams contains all the parameters to send to the API endpoint
for the delete nodes identifier tags tagname operation typically these are written to a http.Request
*/
type DeleteNodesIdentifierTagsTagnameParams struct {

	/*Identifier
	  Mac addresses and unique aliases to identify the node by


	*/
	Identifier string
	/*Tagname
	  tag to remove from node


	*/
	Tagname string
}

// WithIdentifier adds the identifier to the delete nodes identifier tags tagname params
func (o *DeleteNodesIdentifierTagsTagnameParams) WithIdentifier(Identifier string) *DeleteNodesIdentifierTagsTagnameParams {
	o.Identifier = Identifier
	return o
}

// WithTagname adds the tagname to the delete nodes identifier tags tagname params
func (o *DeleteNodesIdentifierTagsTagnameParams) WithTagname(Tagname string) *DeleteNodesIdentifierTagsTagnameParams {
	o.Tagname = Tagname
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteNodesIdentifierTagsTagnameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	var res []error

	// path param identifier
	if err := r.SetPathParam("identifier", o.Identifier); err != nil {
		return err
	}

	// path param tagname
	if err := r.SetPathParam("tagname", o.Tagname); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
