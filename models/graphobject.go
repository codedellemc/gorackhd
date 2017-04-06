package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

/*Graphobject graphobject

swagger:model graphobject
*/
type Graphobject struct {

	/* context
	 */
	Context interface{} `json:"context,omitempty"`

	/* definition
	 */
	Definition interface{} `json:"definition,omitempty"`

	/* id
	 */
	ID string `json:"id,omitempty"`

	/* instanceid
	 */
	Instanceid string `json:"instanceid,omitempty"`

	/* node
	 */
	Node *Node `json:"node,omitempty"`

	/* tasks
	 */
	Tasks interface{} `json:"tasks,omitempty"`
}

// Validate validates this graphobject
func (m *Graphobject) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Graphobject) validateNode(formats strfmt.Registry) error {

	if swag.IsZero(m.Node) { // not required
		return nil
	}

	if m.Node != nil {

		if err := m.Node.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}
