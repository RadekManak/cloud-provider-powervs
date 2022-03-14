// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StockImage Stock image detail
//
// swagger:model StockImage
type StockImage struct {

	// Image ID
	ID string `json:"id,omitempty"`

	// Storage pool for a stock image
	StoragePool string `json:"storagePool,omitempty"`

	// Storage type for a stock image
	StorageType string `json:"storageType,omitempty"`
}

// Validate validates this stock image
func (m *StockImage) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this stock image based on context it is used
func (m *StockImage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StockImage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StockImage) UnmarshalBinary(b []byte) error {
	var res StockImage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
