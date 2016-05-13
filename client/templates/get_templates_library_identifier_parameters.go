package templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetTemplatesLibraryIdentifierParams creates a new GetTemplatesLibraryIdentifierParams object
// with the default values initialized.
func NewGetTemplatesLibraryIdentifierParams() *GetTemplatesLibraryIdentifierParams {
	var ()
	return &GetTemplatesLibraryIdentifierParams{}
}

/*GetTemplatesLibraryIdentifierParams contains all the parameters to send to the API endpoint
for the get templates library identifier operation typically these are written to a http.Request
*/
type GetTemplatesLibraryIdentifierParams struct {

	/*Identifier
	  template identifier


	*/
	Identifier string
}

// WithIdentifier adds the identifier to the get templates library identifier params
func (o *GetTemplatesLibraryIdentifierParams) WithIdentifier(Identifier string) *GetTemplatesLibraryIdentifierParams {
	o.Identifier = Identifier
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetTemplatesLibraryIdentifierParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
