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

// InfrastructureAccountUpdate infrastructure account update
// swagger:model InfrastructureAccountUpdate
type InfrastructureAccountUpdate struct {

	// Cost center of the Owner/infrastructure account
	CostCenter string `json:"cost_center,omitempty"`

	// Level of criticality as defined by tech controlling. 1 is non
	// critical, 2 is standard production, 3 is PCI
	//
	CriticalityLevel int32 `json:"criticality_level,omitempty"`

	// Environment. possible values are "production" or "test".
	//
	Environment string `json:"environment,omitempty"`

	// The external identifier of the account (i.e. AWS account ID)
	ExternalID string `json:"external_id,omitempty"`

	// Globally unique ID of the infrastructure account.
	ID string `json:"id,omitempty"`

	// Lifecycle Status is used to describe the current status of the account.
	// Enum: [requested creating ready decommissioned]
	LifecycleStatus string `json:"lifecycle_status,omitempty"`

	// Name of the infrastructure account
	Name string `json:"name,omitempty"`

	// Owner of the infrastructure account (references an object in the organization service)
	Owner string `json:"owner,omitempty"`

	// Type of the infrastructure account. Possible types are "aws", "gcp",
	// "dc".
	//
	Type string `json:"type,omitempty"`
}

// Validate validates this infrastructure account update
func (m *InfrastructureAccountUpdate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLifecycleStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var infrastructureAccountUpdateTypeLifecycleStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["requested","creating","ready","decommissioned"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		infrastructureAccountUpdateTypeLifecycleStatusPropEnum = append(infrastructureAccountUpdateTypeLifecycleStatusPropEnum, v)
	}
}

const (

	// InfrastructureAccountUpdateLifecycleStatusRequested captures enum value "requested"
	InfrastructureAccountUpdateLifecycleStatusRequested string = "requested"

	// InfrastructureAccountUpdateLifecycleStatusCreating captures enum value "creating"
	InfrastructureAccountUpdateLifecycleStatusCreating string = "creating"

	// InfrastructureAccountUpdateLifecycleStatusReady captures enum value "ready"
	InfrastructureAccountUpdateLifecycleStatusReady string = "ready"

	// InfrastructureAccountUpdateLifecycleStatusDecommissioned captures enum value "decommissioned"
	InfrastructureAccountUpdateLifecycleStatusDecommissioned string = "decommissioned"
)

// prop value enum
func (m *InfrastructureAccountUpdate) validateLifecycleStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, infrastructureAccountUpdateTypeLifecycleStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *InfrastructureAccountUpdate) validateLifecycleStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.LifecycleStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateLifecycleStatusEnum("lifecycle_status", "body", m.LifecycleStatus); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InfrastructureAccountUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InfrastructureAccountUpdate) UnmarshalBinary(b []byte) error {
	var res InfrastructureAccountUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
