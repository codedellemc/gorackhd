package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TagRule tag rule
// swagger:model tag_rule
type TagRule struct {

	// contains
	Contains string `json:"contains,omitempty"`

	// equals
	Equals string `json:"equals,omitempty"`

	// greater than
	GreaterThan float64 `json:"greaterThan,omitempty"`

	// in
	In []string `json:"in"`

	// less than
	LessThan float64 `json:"lessThan,omitempty"`

	// max
	Max float64 `json:"max,omitempty"`

	// min
	Min float64 `json:"min,omitempty"`

	// not contains
	NotContains string `json:"notContains,omitempty"`

	// not in
	NotIn []string `json:"notIn"`

	// not regex
	NotRegex string `json:"notRegex,omitempty"`

	// path
	// Required: true
	Path *string `json:"path"`

	// regex
	Regex string `json:"regex,omitempty"`
}

// Validate validates this tag rule
func (m *TagRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIn(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNotIn(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePath(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TagRule) validateIn(formats strfmt.Registry) error {

	if swag.IsZero(m.In) { // not required
		return nil
	}

	return nil
}

func (m *TagRule) validateNotIn(formats strfmt.Registry) error {

	if swag.IsZero(m.NotIn) { // not required
		return nil
	}

	return nil
}

func (m *TagRule) validatePath(formats strfmt.Registry) error {

	if err := validate.Required("path", "body", m.Path); err != nil {
		return err
	}

	return nil
}
