// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// RetentionMetadata the tag retention metadata
// swagger:model RetentionMetadata
type RetentionMetadata struct {

	// supported scope selectors
	ScopeSelectors []*RetentionSelectorMetadata `json:"scope_selectors"`

	// supported tag selectors
	TagSelectors []*RetentionSelectorMetadata `json:"tag_selectors"`

	// templates
	Templates []*RetentionRuleMetadata `json:"templates"`
}

// Validate validates this retention metadata
func (m *RetentionMetadata) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateScopeSelectors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTagSelectors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTemplates(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RetentionMetadata) validateScopeSelectors(formats strfmt.Registry) error {

	if swag.IsZero(m.ScopeSelectors) { // not required
		return nil
	}

	for i := 0; i < len(m.ScopeSelectors); i++ {
		if swag.IsZero(m.ScopeSelectors[i]) { // not required
			continue
		}

		if m.ScopeSelectors[i] != nil {
			if err := m.ScopeSelectors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("scope_selectors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RetentionMetadata) validateTagSelectors(formats strfmt.Registry) error {

	if swag.IsZero(m.TagSelectors) { // not required
		return nil
	}

	for i := 0; i < len(m.TagSelectors); i++ {
		if swag.IsZero(m.TagSelectors[i]) { // not required
			continue
		}

		if m.TagSelectors[i] != nil {
			if err := m.TagSelectors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tag_selectors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RetentionMetadata) validateTemplates(formats strfmt.Registry) error {

	if swag.IsZero(m.Templates) { // not required
		return nil
	}

	for i := 0; i < len(m.Templates); i++ {
		if swag.IsZero(m.Templates[i]) { // not required
			continue
		}

		if m.Templates[i] != nil {
			if err := m.Templates[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("templates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RetentionMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RetentionMetadata) UnmarshalBinary(b []byte) error {
	var res RetentionMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
