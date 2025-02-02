// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MediatorResponse mediator response
//
// swagger:model mediator_response
type MediatorResponse struct {

	// links
	Links *MediatorResponseLinks `json:"_links,omitempty"`

	// num records
	NumRecords int64 `json:"num_records,omitempty"`

	// records
	Records []*MediatorResponseRecordsItems0 `json:"records,omitempty"`
}

// Validate validates this mediator response
func (m *MediatorResponse) Validate(formats strfmt.Registry) error {
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

func (m *MediatorResponse) validateLinks(formats strfmt.Registry) error {
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

func (m *MediatorResponse) validateRecords(formats strfmt.Registry) error {
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

// ContextValidate validate this mediator response based on the context it is used
func (m *MediatorResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *MediatorResponse) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *MediatorResponse) contextValidateRecords(ctx context.Context, formats strfmt.Registry) error {

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
func (m *MediatorResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MediatorResponse) UnmarshalBinary(b []byte) error {
	var res MediatorResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MediatorResponseLinks mediator response links
//
// swagger:model MediatorResponseLinks
type MediatorResponseLinks struct {

	// next
	Next *Href `json:"next,omitempty"`

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this mediator response links
func (m *MediatorResponseLinks) Validate(formats strfmt.Registry) error {
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

func (m *MediatorResponseLinks) validateNext(formats strfmt.Registry) error {
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

func (m *MediatorResponseLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this mediator response links based on the context it is used
func (m *MediatorResponseLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *MediatorResponseLinks) contextValidateNext(ctx context.Context, formats strfmt.Registry) error {

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

func (m *MediatorResponseLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *MediatorResponseLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MediatorResponseLinks) UnmarshalBinary(b []byte) error {
	var res MediatorResponseLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MediatorResponseRecordsItems0 mediator response records items0
//
// swagger:model MediatorResponseRecordsItems0
type MediatorResponseRecordsItems0 struct {

	// CA certificate for ONTAP Mediator. This is optional if the certificate is already installed.
	CaCertificate string `json:"ca_certificate,omitempty"`

	// dr group
	DrGroup *MediatorResponseRecordsItems0DrGroup `json:"dr_group,omitempty"`

	// The IP address of the mediator.
	// Example: 10.10.10.7
	IPAddress string `json:"ip_address,omitempty"`

	// The password used to connect to the REST server on the mediator.
	// Example: mypassword
	// Format: password
	Password strfmt.Password `json:"password,omitempty"`

	// peer cluster
	PeerCluster *MediatorResponseRecordsItems0PeerCluster `json:"peer_cluster,omitempty"`

	// The REST server's port number on the mediator.
	// Example: 31784
	Port *int64 `json:"port,omitempty"`

	// Indicates the connectivity status of the mediator.
	// Example: true
	// Read Only: true
	Reachable *bool `json:"reachable,omitempty"`

	// The username used to connect to the REST server on the mediator.
	// Example: myusername
	User string `json:"user,omitempty"`

	// The unique identifier for the mediator service.
	// Read Only: true
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this mediator response records items0
func (m *MediatorResponseRecordsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDrGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePeerCluster(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MediatorResponseRecordsItems0) validateDrGroup(formats strfmt.Registry) error {
	if swag.IsZero(m.DrGroup) { // not required
		return nil
	}

	if m.DrGroup != nil {
		if err := m.DrGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dr_group")
			}
			return err
		}
	}

	return nil
}

func (m *MediatorResponseRecordsItems0) validatePassword(formats strfmt.Registry) error {
	if swag.IsZero(m.Password) { // not required
		return nil
	}

	if err := validate.FormatOf("password", "body", "password", m.Password.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MediatorResponseRecordsItems0) validatePeerCluster(formats strfmt.Registry) error {
	if swag.IsZero(m.PeerCluster) { // not required
		return nil
	}

	if m.PeerCluster != nil {
		if err := m.PeerCluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("peer_cluster")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this mediator response records items0 based on the context it is used
func (m *MediatorResponseRecordsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDrGroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePeerCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReachable(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUUID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MediatorResponseRecordsItems0) contextValidateDrGroup(ctx context.Context, formats strfmt.Registry) error {

	if m.DrGroup != nil {
		if err := m.DrGroup.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dr_group")
			}
			return err
		}
	}

	return nil
}

func (m *MediatorResponseRecordsItems0) contextValidatePeerCluster(ctx context.Context, formats strfmt.Registry) error {

	if m.PeerCluster != nil {
		if err := m.PeerCluster.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("peer_cluster")
			}
			return err
		}
	}

	return nil
}

func (m *MediatorResponseRecordsItems0) contextValidateReachable(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "reachable", "body", m.Reachable); err != nil {
		return err
	}

	return nil
}

func (m *MediatorResponseRecordsItems0) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "uuid", "body", string(m.UUID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MediatorResponseRecordsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MediatorResponseRecordsItems0) UnmarshalBinary(b []byte) error {
	var res MediatorResponseRecordsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MediatorResponseRecordsItems0DrGroup DR group reference.
//
// swagger:model MediatorResponseRecordsItems0DrGroup
type MediatorResponseRecordsItems0DrGroup struct {

	// DR Group ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`
}

// Validate validates this mediator response records items0 dr group
func (m *MediatorResponseRecordsItems0DrGroup) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this mediator response records items0 dr group based on the context it is used
func (m *MediatorResponseRecordsItems0DrGroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MediatorResponseRecordsItems0DrGroup) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dr_group"+"."+"id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MediatorResponseRecordsItems0DrGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MediatorResponseRecordsItems0DrGroup) UnmarshalBinary(b []byte) error {
	var res MediatorResponseRecordsItems0DrGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MediatorResponseRecordsItems0PeerCluster The peer cluster that the mediator service is used for.
//
// swagger:model MediatorResponseRecordsItems0PeerCluster
type MediatorResponseRecordsItems0PeerCluster struct {

	// links
	Links *MediatorResponseRecordsItems0PeerClusterLinks `json:"_links,omitempty"`

	// name
	// Example: cluster2
	Name string `json:"name,omitempty"`

	// uuid
	// Example: ebe27c49-1adf-4496-8335-ab862aebebf2
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this mediator response records items0 peer cluster
func (m *MediatorResponseRecordsItems0PeerCluster) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MediatorResponseRecordsItems0PeerCluster) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("peer_cluster" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this mediator response records items0 peer cluster based on the context it is used
func (m *MediatorResponseRecordsItems0PeerCluster) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MediatorResponseRecordsItems0PeerCluster) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("peer_cluster" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MediatorResponseRecordsItems0PeerCluster) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MediatorResponseRecordsItems0PeerCluster) UnmarshalBinary(b []byte) error {
	var res MediatorResponseRecordsItems0PeerCluster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MediatorResponseRecordsItems0PeerClusterLinks mediator response records items0 peer cluster links
//
// swagger:model MediatorResponseRecordsItems0PeerClusterLinks
type MediatorResponseRecordsItems0PeerClusterLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this mediator response records items0 peer cluster links
func (m *MediatorResponseRecordsItems0PeerClusterLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MediatorResponseRecordsItems0PeerClusterLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("peer_cluster" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this mediator response records items0 peer cluster links based on the context it is used
func (m *MediatorResponseRecordsItems0PeerClusterLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MediatorResponseRecordsItems0PeerClusterLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("peer_cluster" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MediatorResponseRecordsItems0PeerClusterLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MediatorResponseRecordsItems0PeerClusterLinks) UnmarshalBinary(b []byte) error {
	var res MediatorResponseRecordsItems0PeerClusterLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
