// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// PingRegistry ping registry
// swagger:model PingRegistry
type PingRegistry struct {

	// The registry access key.
	AccessKey string `json:"access_key,omitempty"`

	// The registry access secret.
	AccessSecret string `json:"access_secret,omitempty"`

	// Credential type of the registry, e.g. 'basic'.
	CredentialType string `json:"credential_type,omitempty"`

	// The ID of the registry
	ID int64 `json:"id,omitempty"`

	// Whether or not the certificate will be verified when Harbor tries to access the server.
	Insecure bool `json:"insecure,omitempty"`

	// Type of the registry, e.g. 'harbor'.
	Type string `json:"type,omitempty"`

	// The registry address URL string.
	URL string `json:"url,omitempty"`
}

// Validate validates this ping registry
func (m *PingRegistry) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PingRegistry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PingRegistry) UnmarshalBinary(b []byte) error {
	var res PingRegistry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
