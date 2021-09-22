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

// NvmeSubsystemController A Non-Volatile Memory Express (NVMe) subsystem controller represents a connection between a host and a storage solution.<br/>
// An NVMe subsystem controller is identified by the NVMe subsystem UUID and the controller ID.
//
//
// swagger:model nvme_subsystem_controller
type NvmeSubsystemController struct {

	// links
	Links *NvmeSubsystemControllerLinks `json:"_links,omitempty"`

	// admin queue
	AdminQueue *NvmeSubsystemControllerAdminQueue `json:"admin_queue,omitempty"`

	// host
	Host *NvmeSubsystemControllerHost `json:"host,omitempty"`

	// The identifier of the subsystem controller. This field consists of 4 zero-filled hexadecimal digits followed by an 'h'.
	//
	// Example: 0040h
	// Read Only: true
	ID string `json:"id,omitempty"`

	// interface
	Interface *NvmeSubsystemControllerInterface `json:"interface,omitempty"`

	// io queue
	IoQueue *NvmeSubsystemControllerIoQueue `json:"io_queue,omitempty"`

	// node
	Node *NvmeSubsystemControllerNode `json:"node,omitempty"`

	// subsystem
	Subsystem *NvmeSubsystemControllerSubsystem `json:"subsystem,omitempty"`

	// svm
	Svm *NvmeSubsystemControllerSvm `json:"svm,omitempty"`
}

// Validate validates this nvme subsystem controller
func (m *NvmeSubsystemController) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAdminQueue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInterface(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIoQueue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubsystem(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSvm(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NvmeSubsystemController) validateLinks(formats strfmt.Registry) error {
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

func (m *NvmeSubsystemController) validateAdminQueue(formats strfmt.Registry) error {
	if swag.IsZero(m.AdminQueue) { // not required
		return nil
	}

	if m.AdminQueue != nil {
		if err := m.AdminQueue.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("admin_queue")
			}
			return err
		}
	}

	return nil
}

func (m *NvmeSubsystemController) validateHost(formats strfmt.Registry) error {
	if swag.IsZero(m.Host) { // not required
		return nil
	}

	if m.Host != nil {
		if err := m.Host.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("host")
			}
			return err
		}
	}

	return nil
}

func (m *NvmeSubsystemController) validateInterface(formats strfmt.Registry) error {
	if swag.IsZero(m.Interface) { // not required
		return nil
	}

	if m.Interface != nil {
		if err := m.Interface.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("interface")
			}
			return err
		}
	}

	return nil
}

func (m *NvmeSubsystemController) validateIoQueue(formats strfmt.Registry) error {
	if swag.IsZero(m.IoQueue) { // not required
		return nil
	}

	if m.IoQueue != nil {
		if err := m.IoQueue.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("io_queue")
			}
			return err
		}
	}

	return nil
}

func (m *NvmeSubsystemController) validateNode(formats strfmt.Registry) error {
	if swag.IsZero(m.Node) { // not required
		return nil
	}

	if m.Node != nil {
		if err := m.Node.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node")
			}
			return err
		}
	}

	return nil
}

func (m *NvmeSubsystemController) validateSubsystem(formats strfmt.Registry) error {
	if swag.IsZero(m.Subsystem) { // not required
		return nil
	}

	if m.Subsystem != nil {
		if err := m.Subsystem.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subsystem")
			}
			return err
		}
	}

	return nil
}

func (m *NvmeSubsystemController) validateSvm(formats strfmt.Registry) error {
	if swag.IsZero(m.Svm) { // not required
		return nil
	}

	if m.Svm != nil {
		if err := m.Svm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this nvme subsystem controller based on the context it is used
func (m *NvmeSubsystemController) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAdminQueue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHost(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInterface(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIoQueue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSubsystem(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSvm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NvmeSubsystemController) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *NvmeSubsystemController) contextValidateAdminQueue(ctx context.Context, formats strfmt.Registry) error {

	if m.AdminQueue != nil {
		if err := m.AdminQueue.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("admin_queue")
			}
			return err
		}
	}

	return nil
}

func (m *NvmeSubsystemController) contextValidateHost(ctx context.Context, formats strfmt.Registry) error {

	if m.Host != nil {
		if err := m.Host.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("host")
			}
			return err
		}
	}

	return nil
}

func (m *NvmeSubsystemController) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *NvmeSubsystemController) contextValidateInterface(ctx context.Context, formats strfmt.Registry) error {

	if m.Interface != nil {
		if err := m.Interface.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("interface")
			}
			return err
		}
	}

	return nil
}

func (m *NvmeSubsystemController) contextValidateIoQueue(ctx context.Context, formats strfmt.Registry) error {

	if m.IoQueue != nil {
		if err := m.IoQueue.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("io_queue")
			}
			return err
		}
	}

	return nil
}

func (m *NvmeSubsystemController) contextValidateNode(ctx context.Context, formats strfmt.Registry) error {

	if m.Node != nil {
		if err := m.Node.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node")
			}
			return err
		}
	}

	return nil
}

func (m *NvmeSubsystemController) contextValidateSubsystem(ctx context.Context, formats strfmt.Registry) error {

	if m.Subsystem != nil {
		if err := m.Subsystem.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subsystem")
			}
			return err
		}
	}

	return nil
}

func (m *NvmeSubsystemController) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

	if m.Svm != nil {
		if err := m.Svm.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NvmeSubsystemController) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NvmeSubsystemController) UnmarshalBinary(b []byte) error {
	var res NvmeSubsystemController
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NvmeSubsystemControllerAdminQueue nvme subsystem controller admin queue
//
// swagger:model NvmeSubsystemControllerAdminQueue
type NvmeSubsystemControllerAdminQueue struct {

	// The depth of the admin queue for the controller.
	//
	// Read Only: true
	Depth int64 `json:"depth,omitempty"`
}

// Validate validates this nvme subsystem controller admin queue
func (m *NvmeSubsystemControllerAdminQueue) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this nvme subsystem controller admin queue based on the context it is used
func (m *NvmeSubsystemControllerAdminQueue) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDepth(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NvmeSubsystemControllerAdminQueue) contextValidateDepth(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "admin_queue"+"."+"depth", "body", int64(m.Depth)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NvmeSubsystemControllerAdminQueue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NvmeSubsystemControllerAdminQueue) UnmarshalBinary(b []byte) error {
	var res NvmeSubsystemControllerAdminQueue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NvmeSubsystemControllerHost Properties of the connected host.
//
//
// swagger:model NvmeSubsystemControllerHost
type NvmeSubsystemControllerHost struct {

	// The host identifier registered with the controller.
	//
	// Example: b8546ca6097349e5b1558dc154fc073b
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The NVMe qualified name of the host.
	//
	// Example: nqn.2014-08.org.nvmexpress:uuid:c2846cb1-89d2-4020-a3b0-71ce907b4eef
	// Read Only: true
	// Max Length: 223
	// Min Length: 1
	Nqn string `json:"nqn,omitempty"`

	// The transport address of the host.
	//
	// Example: nn-0x20000090fae00806:pn-0x10000090fae00806
	// Read Only: true
	TransportAddress string `json:"transport_address,omitempty"`
}

// Validate validates this nvme subsystem controller host
func (m *NvmeSubsystemControllerHost) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNqn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NvmeSubsystemControllerHost) validateNqn(formats strfmt.Registry) error {
	if swag.IsZero(m.Nqn) { // not required
		return nil
	}

	if err := validate.MinLength("host"+"."+"nqn", "body", m.Nqn, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("host"+"."+"nqn", "body", m.Nqn, 223); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this nvme subsystem controller host based on the context it is used
func (m *NvmeSubsystemControllerHost) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNqn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTransportAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NvmeSubsystemControllerHost) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "host"+"."+"id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *NvmeSubsystemControllerHost) contextValidateNqn(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "host"+"."+"nqn", "body", string(m.Nqn)); err != nil {
		return err
	}

	return nil
}

func (m *NvmeSubsystemControllerHost) contextValidateTransportAddress(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "host"+"."+"transport_address", "body", string(m.TransportAddress)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NvmeSubsystemControllerHost) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NvmeSubsystemControllerHost) UnmarshalBinary(b []byte) error {
	var res NvmeSubsystemControllerHost
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NvmeSubsystemControllerInterface The logical interface through which the host is connected.
//
//
// swagger:model NvmeSubsystemControllerInterface
type NvmeSubsystemControllerInterface struct {

	// The name of the logical interface.
	//
	// Example: lif1
	// Read Only: true
	Name string `json:"name,omitempty"`

	// The transport address of the logical interface.
	//
	// Example: nn-0x200400a0989a1c8d:pn-0x200500a0989a1c8d
	// Read Only: true
	TransportAddress string `json:"transport_address,omitempty"`

	// The unique identifier of the logical interface.
	//
	// Example: fa1c5941-2593-11e9-94c4-00a0989a1c8e
	// Read Only: true
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this nvme subsystem controller interface
func (m *NvmeSubsystemControllerInterface) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this nvme subsystem controller interface based on the context it is used
func (m *NvmeSubsystemControllerInterface) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTransportAddress(ctx, formats); err != nil {
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

func (m *NvmeSubsystemControllerInterface) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "interface"+"."+"name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

func (m *NvmeSubsystemControllerInterface) contextValidateTransportAddress(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "interface"+"."+"transport_address", "body", string(m.TransportAddress)); err != nil {
		return err
	}

	return nil
}

func (m *NvmeSubsystemControllerInterface) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "interface"+"."+"uuid", "body", string(m.UUID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NvmeSubsystemControllerInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NvmeSubsystemControllerInterface) UnmarshalBinary(b []byte) error {
	var res NvmeSubsystemControllerInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NvmeSubsystemControllerIoQueue Properties of the I/O queues available to the controller.
//
//
// swagger:model NvmeSubsystemControllerIoQueue
type NvmeSubsystemControllerIoQueue struct {

	// The number of I/O queues available to the controller.
	//
	// Read Only: true
	Count int64 `json:"count,omitempty"`

	// The depths of the I/O queues.
	//
	Depth []int64 `json:"depth,omitempty"`
}

// Validate validates this nvme subsystem controller io queue
func (m *NvmeSubsystemControllerIoQueue) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this nvme subsystem controller io queue based on the context it is used
func (m *NvmeSubsystemControllerIoQueue) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDepth(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NvmeSubsystemControllerIoQueue) contextValidateCount(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "io_queue"+"."+"count", "body", int64(m.Count)); err != nil {
		return err
	}

	return nil
}

func (m *NvmeSubsystemControllerIoQueue) contextValidateDepth(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Depth); i++ {

		if err := validate.ReadOnly(ctx, "io_queue"+"."+"depth"+"."+strconv.Itoa(i), "body", int64(m.Depth[i])); err != nil {
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NvmeSubsystemControllerIoQueue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NvmeSubsystemControllerIoQueue) UnmarshalBinary(b []byte) error {
	var res NvmeSubsystemControllerIoQueue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NvmeSubsystemControllerLinks nvme subsystem controller links
//
// swagger:model NvmeSubsystemControllerLinks
type NvmeSubsystemControllerLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this nvme subsystem controller links
func (m *NvmeSubsystemControllerLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NvmeSubsystemControllerLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this nvme subsystem controller links based on the context it is used
func (m *NvmeSubsystemControllerLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NvmeSubsystemControllerLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *NvmeSubsystemControllerLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NvmeSubsystemControllerLinks) UnmarshalBinary(b []byte) error {
	var res NvmeSubsystemControllerLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NvmeSubsystemControllerNode nvme subsystem controller node
//
// swagger:model NvmeSubsystemControllerNode
type NvmeSubsystemControllerNode struct {

	// links
	Links *NvmeSubsystemControllerNodeLinks `json:"_links,omitempty"`

	// name
	// Example: node1
	Name string `json:"name,omitempty"`

	// uuid
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this nvme subsystem controller node
func (m *NvmeSubsystemControllerNode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NvmeSubsystemControllerNode) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this nvme subsystem controller node based on the context it is used
func (m *NvmeSubsystemControllerNode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NvmeSubsystemControllerNode) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NvmeSubsystemControllerNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NvmeSubsystemControllerNode) UnmarshalBinary(b []byte) error {
	var res NvmeSubsystemControllerNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NvmeSubsystemControllerNodeLinks nvme subsystem controller node links
//
// swagger:model NvmeSubsystemControllerNodeLinks
type NvmeSubsystemControllerNodeLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this nvme subsystem controller node links
func (m *NvmeSubsystemControllerNodeLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NvmeSubsystemControllerNodeLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this nvme subsystem controller node links based on the context it is used
func (m *NvmeSubsystemControllerNodeLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NvmeSubsystemControllerNodeLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NvmeSubsystemControllerNodeLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NvmeSubsystemControllerNodeLinks) UnmarshalBinary(b []byte) error {
	var res NvmeSubsystemControllerNodeLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NvmeSubsystemControllerSubsystem nvme subsystem controller subsystem
//
// swagger:model NvmeSubsystemControllerSubsystem
type NvmeSubsystemControllerSubsystem struct {

	// links
	Links *NvmeSubsystemControllerSubsystemLinks `json:"_links,omitempty"`

	// The name of the NVMe subsystem.
	//
	// Max Length: 96
	// Min Length: 1
	Name string `json:"name,omitempty"`

	// The unique identifier of the NVMe subsystem.
	//
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this nvme subsystem controller subsystem
func (m *NvmeSubsystemControllerSubsystem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NvmeSubsystemControllerSubsystem) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subsystem" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (m *NvmeSubsystemControllerSubsystem) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MinLength("subsystem"+"."+"name", "body", m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("subsystem"+"."+"name", "body", m.Name, 96); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this nvme subsystem controller subsystem based on the context it is used
func (m *NvmeSubsystemControllerSubsystem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NvmeSubsystemControllerSubsystem) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subsystem" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NvmeSubsystemControllerSubsystem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NvmeSubsystemControllerSubsystem) UnmarshalBinary(b []byte) error {
	var res NvmeSubsystemControllerSubsystem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NvmeSubsystemControllerSubsystemLinks nvme subsystem controller subsystem links
//
// swagger:model NvmeSubsystemControllerSubsystemLinks
type NvmeSubsystemControllerSubsystemLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this nvme subsystem controller subsystem links
func (m *NvmeSubsystemControllerSubsystemLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NvmeSubsystemControllerSubsystemLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subsystem" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this nvme subsystem controller subsystem links based on the context it is used
func (m *NvmeSubsystemControllerSubsystemLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NvmeSubsystemControllerSubsystemLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subsystem" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NvmeSubsystemControllerSubsystemLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NvmeSubsystemControllerSubsystemLinks) UnmarshalBinary(b []byte) error {
	var res NvmeSubsystemControllerSubsystemLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NvmeSubsystemControllerSvm nvme subsystem controller svm
//
// swagger:model NvmeSubsystemControllerSvm
type NvmeSubsystemControllerSvm struct {

	// links
	Links *NvmeSubsystemControllerSvmLinks `json:"_links,omitempty"`

	// The name of the SVM.
	//
	// Example: svm1
	Name string `json:"name,omitempty"`

	// The unique identifier of the SVM.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this nvme subsystem controller svm
func (m *NvmeSubsystemControllerSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NvmeSubsystemControllerSvm) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this nvme subsystem controller svm based on the context it is used
func (m *NvmeSubsystemControllerSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NvmeSubsystemControllerSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NvmeSubsystemControllerSvm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NvmeSubsystemControllerSvm) UnmarshalBinary(b []byte) error {
	var res NvmeSubsystemControllerSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NvmeSubsystemControllerSvmLinks nvme subsystem controller svm links
//
// swagger:model NvmeSubsystemControllerSvmLinks
type NvmeSubsystemControllerSvmLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this nvme subsystem controller svm links
func (m *NvmeSubsystemControllerSvmLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NvmeSubsystemControllerSvmLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this nvme subsystem controller svm links based on the context it is used
func (m *NvmeSubsystemControllerSvmLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NvmeSubsystemControllerSvmLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NvmeSubsystemControllerSvmLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NvmeSubsystemControllerSvmLinks) UnmarshalBinary(b []byte) error {
	var res NvmeSubsystemControllerSvmLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}