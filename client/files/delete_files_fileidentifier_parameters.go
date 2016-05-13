package files

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteFilesFileidentifierParams creates a new DeleteFilesFileidentifierParams object
// with the default values initialized.
func NewDeleteFilesFileidentifierParams() *DeleteFilesFileidentifierParams {
	var ()
	return &DeleteFilesFileidentifierParams{}
}

/*DeleteFilesFileidentifierParams contains all the parameters to send to the API endpoint
for the delete files fileidentifier operation typically these are written to a http.Request
*/
type DeleteFilesFileidentifierParams struct {

	/*Fileidentifier
	  filename identifier of the file you wish to delete

	*/
	Fileidentifier string
}

// WithFileidentifier adds the fileidentifier to the delete files fileidentifier params
func (o *DeleteFilesFileidentifierParams) WithFileidentifier(Fileidentifier string) *DeleteFilesFileidentifierParams {
	o.Fileidentifier = Fileidentifier
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteFilesFileidentifierParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	var res []error

	// path param fileidentifier
	if err := r.SetPathParam("fileidentifier", o.Fileidentifier); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
