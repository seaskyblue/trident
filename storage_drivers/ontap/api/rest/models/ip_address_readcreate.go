// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// IPAddressReadcreate IPv4 or IPv6 address
// Example: 10.10.10.7
//
// swagger:model ip_address_readcreate
type IPAddressReadcreate string

// Validate validates this ip address readcreate
func (m IPAddressReadcreate) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ip address readcreate based on context it is used
func (m IPAddressReadcreate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}