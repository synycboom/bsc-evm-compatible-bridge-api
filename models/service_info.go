// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ServiceInfo service info
//
// swagger:model ServiceInfo
type ServiceInfo struct {

	// bsc erc 1155 swap agent
	BscErc1155SwapAgent string `json:"bsc_erc_1155_swap_agent"`

	// bsc erc 721 swap agent
	BscErc721SwapAgent string `json:"bsc_erc_721_swap_agent"`

	// eth erc 1155 swap agent
	EthErc1155SwapAgent string `json:"eth_erc_1155_swap_agent"`

	// eth erc 721 swap agent
	EthErc721SwapAgent string `json:"eth_erc_721_swap_agent"`
}

// Validate validates this service info
func (m *ServiceInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this service info based on context it is used
func (m *ServiceInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceInfo) UnmarshalBinary(b []byte) error {
	var res ServiceInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
