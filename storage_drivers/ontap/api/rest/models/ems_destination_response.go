// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// EmsDestinationResponse ems destination response
//
// swagger:model ems_destination_response
type EmsDestinationResponse struct {

	// links
	Links *EmsDestinationResponseLinks `json:"_links,omitempty"`

	// Number of records
	// Example: 3
	NumRecords int64 `json:"num_records,omitempty"`

	// records
	Records []*EmsDestinationResponseRecordsItems0 `json:"records,omitempty"`
}

// Validate validates this ems destination response
func (m *EmsDestinationResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsDestinationResponse) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *EmsDestinationResponse) validateRecords(formats strfmt.Registry) error {
	if swag.IsZero(m.Records) { // not required
		return nil
	}

	for i := 0; i < len(m.Records); i++ {
		if swag.IsZero(m.Records[i]) { // not required
			continue
		}

		if m.Records[i] != nil {
			if err := m.Records[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this ems destination response based on the context it is used
func (m *EmsDestinationResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsDestinationResponse) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *EmsDestinationResponse) contextValidateRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Records); i++ {

		if m.Records[i] != nil {
			if err := m.Records[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *EmsDestinationResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmsDestinationResponse) UnmarshalBinary(b []byte) error {
	var res EmsDestinationResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EmsDestinationResponseLinks ems destination response links
//
// swagger:model EmsDestinationResponseLinks
type EmsDestinationResponseLinks struct {

	// next
	Next *Href `json:"next,omitempty"`

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this ems destination response links
func (m *EmsDestinationResponseLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsDestinationResponseLinks) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(m.Next) { // not required
		return nil
	}

	if m.Next != nil {
		if err := m.Next.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "next")
			}
			return err
		}
	}

	return nil
}

func (m *EmsDestinationResponseLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ems destination response links based on the context it is used
func (m *EmsDestinationResponseLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNext(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsDestinationResponseLinks) contextValidateNext(ctx context.Context, formats strfmt.Registry) error {

	if m.Next != nil {
		if err := m.Next.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "next")
			}
			return err
		}
	}

	return nil
}

func (m *EmsDestinationResponseLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EmsDestinationResponseLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmsDestinationResponseLinks) UnmarshalBinary(b []byte) error {
	var res EmsDestinationResponseLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EmsDestinationResponseRecordsItems0 ems destination response records items0
//
// swagger:model EmsDestinationResponseRecordsItems0
type EmsDestinationResponseRecordsItems0 struct {

	// links
	Links *EmsDestinationResponseRecordsItems0Links `json:"_links,omitempty"`

	// certificate
	Certificate *EmsDestinationResponseRecordsItems0Certificate `json:"certificate,omitempty"`

	// Event destination
	// Example: administrator@mycompany.com
	Destination string `json:"destination,omitempty"`

	// filters
	Filters []*EmsDestinationResponseRecordsItems0FiltersItems0 `json:"filters,omitempty"`

	// Destination name.  Valid in POST.
	// Example: Admin_Email
	Name string `json:"name,omitempty"`

	// Type of destination. Valid in POST.
	// Example: email
	// Enum: [snmp email syslog rest_api]
	Type string `json:"type,omitempty"`
}

// Validate validates this ems destination response records items0
func (m *EmsDestinationResponseRecordsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCertificate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsDestinationResponseRecordsItems0) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *EmsDestinationResponseRecordsItems0) validateCertificate(formats strfmt.Registry) error {
	if swag.IsZero(m.Certificate) { // not required
		return nil
	}

	if m.Certificate != nil {
		if err := m.Certificate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("certificate")
			}
			return err
		}
	}

	return nil
}

func (m *EmsDestinationResponseRecordsItems0) validateFilters(formats strfmt.Registry) error {
	if swag.IsZero(m.Filters) { // not required
		return nil
	}

	for i := 0; i < len(m.Filters); i++ {
		if swag.IsZero(m.Filters[i]) { // not required
			continue
		}

		if m.Filters[i] != nil {
			if err := m.Filters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("filters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var emsDestinationResponseRecordsItems0TypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["snmp","email","syslog","rest_api"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		emsDestinationResponseRecordsItems0TypeTypePropEnum = append(emsDestinationResponseRecordsItems0TypeTypePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// EmsDestinationResponseRecordsItems0
	// EmsDestinationResponseRecordsItems0
	// type
	// Type
	// snmp
	// END DEBUGGING
	// EmsDestinationResponseRecordsItems0TypeSnmp captures enum value "snmp"
	EmsDestinationResponseRecordsItems0TypeSnmp string = "snmp"

	// BEGIN DEBUGGING
	// EmsDestinationResponseRecordsItems0
	// EmsDestinationResponseRecordsItems0
	// type
	// Type
	// email
	// END DEBUGGING
	// EmsDestinationResponseRecordsItems0TypeEmail captures enum value "email"
	EmsDestinationResponseRecordsItems0TypeEmail string = "email"

	// BEGIN DEBUGGING
	// EmsDestinationResponseRecordsItems0
	// EmsDestinationResponseRecordsItems0
	// type
	// Type
	// syslog
	// END DEBUGGING
	// EmsDestinationResponseRecordsItems0TypeSyslog captures enum value "syslog"
	EmsDestinationResponseRecordsItems0TypeSyslog string = "syslog"

	// BEGIN DEBUGGING
	// EmsDestinationResponseRecordsItems0
	// EmsDestinationResponseRecordsItems0
	// type
	// Type
	// rest_api
	// END DEBUGGING
	// EmsDestinationResponseRecordsItems0TypeRestAPI captures enum value "rest_api"
	EmsDestinationResponseRecordsItems0TypeRestAPI string = "rest_api"
)

// prop value enum
func (m *EmsDestinationResponseRecordsItems0) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, emsDestinationResponseRecordsItems0TypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *EmsDestinationResponseRecordsItems0) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this ems destination response records items0 based on the context it is used
func (m *EmsDestinationResponseRecordsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCertificate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFilters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsDestinationResponseRecordsItems0) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *EmsDestinationResponseRecordsItems0) contextValidateCertificate(ctx context.Context, formats strfmt.Registry) error {

	if m.Certificate != nil {
		if err := m.Certificate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("certificate")
			}
			return err
		}
	}

	return nil
}

func (m *EmsDestinationResponseRecordsItems0) contextValidateFilters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Filters); i++ {

		if m.Filters[i] != nil {
			if err := m.Filters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("filters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *EmsDestinationResponseRecordsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmsDestinationResponseRecordsItems0) UnmarshalBinary(b []byte) error {
	var res EmsDestinationResponseRecordsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EmsDestinationResponseRecordsItems0Certificate Certificate information is valid for the "rest_api" type.
//
// swagger:model EmsDestinationResponseRecordsItems0Certificate
type EmsDestinationResponseRecordsItems0Certificate struct {

	// Client certificate issuing CA
	// Example: VeriSign
	// Max Length: 256
	// Min Length: 1
	Ca string `json:"ca,omitempty"`

	// Client certificate serial number
	// Example: 1234567890
	// Max Length: 40
	// Min Length: 1
	SerialNumber string `json:"serial_number,omitempty"`
}

// Validate validates this ems destination response records items0 certificate
func (m *EmsDestinationResponseRecordsItems0Certificate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCa(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSerialNumber(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsDestinationResponseRecordsItems0Certificate) validateCa(formats strfmt.Registry) error {
	if swag.IsZero(m.Ca) { // not required
		return nil
	}

	if err := validate.MinLength("certificate"+"."+"ca", "body", m.Ca, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("certificate"+"."+"ca", "body", m.Ca, 256); err != nil {
		return err
	}

	return nil
}

func (m *EmsDestinationResponseRecordsItems0Certificate) validateSerialNumber(formats strfmt.Registry) error {
	if swag.IsZero(m.SerialNumber) { // not required
		return nil
	}

	if err := validate.MinLength("certificate"+"."+"serial_number", "body", m.SerialNumber, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("certificate"+"."+"serial_number", "body", m.SerialNumber, 40); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this ems destination response records items0 certificate based on context it is used
func (m *EmsDestinationResponseRecordsItems0Certificate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EmsDestinationResponseRecordsItems0Certificate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmsDestinationResponseRecordsItems0Certificate) UnmarshalBinary(b []byte) error {
	var res EmsDestinationResponseRecordsItems0Certificate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EmsDestinationResponseRecordsItems0FiltersItems0 ems destination response records items0 filters items0
//
// swagger:model EmsDestinationResponseRecordsItems0FiltersItems0
type EmsDestinationResponseRecordsItems0FiltersItems0 struct {

	// links
	Links *EmsDestinationResponseRecordsItems0FiltersItems0Links `json:"_links,omitempty"`

	// name
	// Example: important-events
	Name string `json:"name,omitempty"`
}

// Validate validates this ems destination response records items0 filters items0
func (m *EmsDestinationResponseRecordsItems0FiltersItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsDestinationResponseRecordsItems0FiltersItems0) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ems destination response records items0 filters items0 based on the context it is used
func (m *EmsDestinationResponseRecordsItems0FiltersItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsDestinationResponseRecordsItems0FiltersItems0) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EmsDestinationResponseRecordsItems0FiltersItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmsDestinationResponseRecordsItems0FiltersItems0) UnmarshalBinary(b []byte) error {
	var res EmsDestinationResponseRecordsItems0FiltersItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EmsDestinationResponseRecordsItems0FiltersItems0Links ems destination response records items0 filters items0 links
//
// swagger:model EmsDestinationResponseRecordsItems0FiltersItems0Links
type EmsDestinationResponseRecordsItems0FiltersItems0Links struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this ems destination response records items0 filters items0 links
func (m *EmsDestinationResponseRecordsItems0FiltersItems0Links) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsDestinationResponseRecordsItems0FiltersItems0Links) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ems destination response records items0 filters items0 links based on the context it is used
func (m *EmsDestinationResponseRecordsItems0FiltersItems0Links) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsDestinationResponseRecordsItems0FiltersItems0Links) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EmsDestinationResponseRecordsItems0FiltersItems0Links) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmsDestinationResponseRecordsItems0FiltersItems0Links) UnmarshalBinary(b []byte) error {
	var res EmsDestinationResponseRecordsItems0FiltersItems0Links
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EmsDestinationResponseRecordsItems0Links ems destination response records items0 links
//
// swagger:model EmsDestinationResponseRecordsItems0Links
type EmsDestinationResponseRecordsItems0Links struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this ems destination response records items0 links
func (m *EmsDestinationResponseRecordsItems0Links) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsDestinationResponseRecordsItems0Links) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ems destination response records items0 links based on the context it is used
func (m *EmsDestinationResponseRecordsItems0Links) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsDestinationResponseRecordsItems0Links) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EmsDestinationResponseRecordsItems0Links) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmsDestinationResponseRecordsItems0Links) UnmarshalBinary(b []byte) error {
	var res EmsDestinationResponseRecordsItems0Links
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
