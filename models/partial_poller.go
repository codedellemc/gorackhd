// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PartialPoller A poller for periodic collection of telemetry data
// swagger:model PartialPoller

type PartialPoller struct {

	// Poller configuration object
	Config interface{} `json:"config,omitempty"`

	// Asserted if poller is paused
	Paused bool `json:"paused,omitempty"`

	// Interval at which poller will run
	PollInterval float64 `json:"pollInterval,omitempty"`

	// Type of poller
	Type string `json:"type,omitempty"`
}

/* polymorph PartialPoller config false */

/* polymorph PartialPoller paused false */

/* polymorph PartialPoller pollInterval false */

/* polymorph PartialPoller type false */

// Validate validates this partial poller
func (m *PartialPoller) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var partialPollerTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ipmi","snmp","redfish","wsman"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		partialPollerTypeTypePropEnum = append(partialPollerTypeTypePropEnum, v)
	}
}

const (
	// PartialPollerTypeIPMI captures enum value "ipmi"
	PartialPollerTypeIPMI string = "ipmi"
	// PartialPollerTypeSnmp captures enum value "snmp"
	PartialPollerTypeSnmp string = "snmp"
	// PartialPollerTypeRedfish captures enum value "redfish"
	PartialPollerTypeRedfish string = "redfish"
	// PartialPollerTypeWsman captures enum value "wsman"
	PartialPollerTypeWsman string = "wsman"
)

// prop value enum
func (m *PartialPoller) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, partialPollerTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PartialPoller) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PartialPoller) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PartialPoller) UnmarshalBinary(b []byte) error {
	var res PartialPoller
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
