package obms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetObmsLibraryIdentifierParams creates a new GetObmsLibraryIdentifierParams object
// with the default values initialized.
func NewGetObmsLibraryIdentifierParams() *GetObmsLibraryIdentifierParams {
	var ()
	return &GetObmsLibraryIdentifierParams{}
}

/*GetObmsLibraryIdentifierParams contains all the parameters to send to the API endpoint
for the get obms library identifier operation typically these are written to a http.Request
*/
type GetObmsLibraryIdentifierParams struct {

	/*Identifier
	  The obm service name.


	*/
	Identifier string
}

// WithIdentifier adds the identifier to the get obms library identifier params
func (o *GetObmsLibraryIdentifierParams) WithIdentifier(Identifier string) *GetObmsLibraryIdentifierParams {
	o.Identifier = Identifier
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetObmsLibraryIdentifierParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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