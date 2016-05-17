package tags

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPatchNodesIdentifierTagsParams creates a new PatchNodesIdentifierTagsParams object
// with the default values initialized.
func NewPatchNodesIdentifierTagsParams() *PatchNodesIdentifierTagsParams {
	var ()
	return &PatchNodesIdentifierTagsParams{}
}

/*PatchNodesIdentifierTagsParams contains all the parameters to send to the API endpoint
for the patch nodes identifier tags operation typically these are written to a http.Request
*/
type PatchNodesIdentifierTagsParams struct {

	/*Body
	  tags to apply


	*/
	Body interface{}
	/*Identifier
	  Mac addresses and unique aliases to identify the node by


	*/
	Identifier string
}

// WithBody adds the body to the patch nodes identifier tags params
func (o *PatchNodesIdentifierTagsParams) WithBody(Body interface{}) *PatchNodesIdentifierTagsParams {
	o.Body = Body
	return o
}

// WithIdentifier adds the identifier to the patch nodes identifier tags params
func (o *PatchNodesIdentifierTagsParams) WithIdentifier(Identifier string) *PatchNodesIdentifierTagsParams {
	o.Identifier = Identifier
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PatchNodesIdentifierTagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param identifier
	if err := r.SetPathParam("identifier", o.Identifier); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
