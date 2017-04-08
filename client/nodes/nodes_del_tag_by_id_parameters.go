package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewNodesDelTagByIDParams creates a new NodesDelTagByIDParams object
// with the default values initialized.
func NewNodesDelTagByIDParams() *NodesDelTagByIDParams {
	var ()
	return &NodesDelTagByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewNodesDelTagByIDParamsWithTimeout creates a new NodesDelTagByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewNodesDelTagByIDParamsWithTimeout(timeout time.Duration) *NodesDelTagByIDParams {
	var ()
	return &NodesDelTagByIDParams{

		timeout: timeout,
	}
}

/*NodesDelTagByIDParams contains all the parameters to send to the API endpoint
for the nodes del tag by Id operation typically these are written to a http.Request
*/
type NodesDelTagByIDParams struct {

	/*Identifier
	  The node identifier

	*/
	Identifier string
	/*TagName
	  The name of the tag

	*/
	TagName string

	timeout time.Duration
}

// WithIdentifier adds the identifier to the nodes del tag by Id params
func (o *NodesDelTagByIDParams) WithIdentifier(identifier string) *NodesDelTagByIDParams {
	o.Identifier = identifier
	return o
}

// WithTagName adds the tagName to the nodes del tag by Id params
func (o *NodesDelTagByIDParams) WithTagName(tagName string) *NodesDelTagByIDParams {
	o.TagName = tagName
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *NodesDelTagByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param identifier
	if err := r.SetPathParam("identifier", o.Identifier); err != nil {
		return err
	}

	// path param tagName
	if err := r.SetPathParam("tagName", o.TagName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}