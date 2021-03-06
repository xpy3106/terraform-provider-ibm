// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Volume Volume
// swagger:model volume
type Volume struct {

	// If set to true, this volume will be automatically deleted if the only server it is attached to is deleted
	AutoDelete *bool `json:"auto_delete,omitempty"`

	// The capacity of the volume in gigabytes
	// Maximum: 64000
	// Minimum: 10
	Capacity int64 `json:"capacity,omitempty"`

	// The date and time that the volume was created
	CreatedAt string `json:"created_at,omitempty"`

	// The CRN for this volume
	Crn string `json:"crn,omitempty"`

	// This field defines the type of encryption used on the volume.
	// Enum: [provider_managed user_managed]
	Encryption string `json:"encryption,omitempty"`

	// encryption key
	EncryptionKey *VolumeEncryptionKey `json:"encryption_key,omitempty"`

	// The URL for this volume
	// Pattern: ^http(s)?:\/\/([^\/?#]*)([^?#]*)(\?([^#]*))?(#(.*))?$
	Href string `json:"href,omitempty"`

	// The unique identifier for this volume
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// The bandwidth for the volume
	// Enum: [1000 10000 100000]
	Iops int64 `json:"iops,omitempty"`

	// The user-defined name for this volume
	// Pattern: ^[A-Za-z][-A-Za-z0-9_]*$
	Name string `json:"name,omitempty"`

	// profile
	Profile *ResourceReference `json:"profile,omitempty"`

	// resource group
	ResourceGroup *ResourceReference `json:"resource_group,omitempty"`

	// The status of the volume
	// Enum: [pending available pending_deletion failed]
	Status string `json:"status,omitempty"`

	// volume attachments
	VolumeAttachments []*VolumeVolumeAttachmentsItems `json:"volume_attachments"`

	// zone
	Zone *VolumeZone `json:"zone,omitempty"`
}

// Validate validates this volume
func (m *Volume) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCapacity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEncryption(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEncryptionKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHref(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIops(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProfile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolumeAttachments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateZone(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Volume) validateCapacity(formats strfmt.Registry) error {

	if swag.IsZero(m.Capacity) { // not required
		return nil
	}

	if err := validate.MinimumInt("capacity", "body", int64(m.Capacity), 10, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("capacity", "body", int64(m.Capacity), 64000, false); err != nil {
		return err
	}

	return nil
}

var volumeTypeEncryptionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["provider_managed","user_managed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		volumeTypeEncryptionPropEnum = append(volumeTypeEncryptionPropEnum, v)
	}
}

const (

	// VolumeEncryptionProviderManaged captures enum value "provider_managed"
	VolumeEncryptionProviderManaged string = "provider_managed"

	// VolumeEncryptionUserManaged captures enum value "user_managed"
	VolumeEncryptionUserManaged string = "user_managed"
)

// prop value enum
func (m *Volume) validateEncryptionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, volumeTypeEncryptionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Volume) validateEncryption(formats strfmt.Registry) error {

	if swag.IsZero(m.Encryption) { // not required
		return nil
	}

	// value enum
	if err := m.validateEncryptionEnum("encryption", "body", m.Encryption); err != nil {
		return err
	}

	return nil
}

func (m *Volume) validateEncryptionKey(formats strfmt.Registry) error {

	if swag.IsZero(m.EncryptionKey) { // not required
		return nil
	}

	if m.EncryptionKey != nil {
		if err := m.EncryptionKey.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("encryption_key")
			}
			return err
		}
	}

	return nil
}

func (m *Volume) validateHref(formats strfmt.Registry) error {

	if swag.IsZero(m.Href) { // not required
		return nil
	}

	if err := validate.Pattern("href", "body", string(m.Href), `^http(s)?:\/\/([^\/?#]*)([^?#]*)(\?([^#]*))?(#(.*))?$`); err != nil {
		return err
	}

	return nil
}

func (m *Volume) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

var volumeTypeIopsPropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[1000,10000,100000]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		volumeTypeIopsPropEnum = append(volumeTypeIopsPropEnum, v)
	}
}

// prop value enum
func (m *Volume) validateIopsEnum(path, location string, value int64) error {
	if err := validate.Enum(path, location, value, volumeTypeIopsPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Volume) validateIops(formats strfmt.Registry) error {

	if swag.IsZero(m.Iops) { // not required
		return nil
	}

	// value enum
	if err := m.validateIopsEnum("iops", "body", m.Iops); err != nil {
		return err
	}

	return nil
}

func (m *Volume) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.Pattern("name", "body", string(m.Name), `^[A-Za-z][-A-Za-z0-9_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *Volume) validateProfile(formats strfmt.Registry) error {

	if swag.IsZero(m.Profile) { // not required
		return nil
	}

	if m.Profile != nil {
		if err := m.Profile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("profile")
			}
			return err
		}
	}

	return nil
}

func (m *Volume) validateResourceGroup(formats strfmt.Registry) error {

	if swag.IsZero(m.ResourceGroup) { // not required
		return nil
	}

	if m.ResourceGroup != nil {
		if err := m.ResourceGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resource_group")
			}
			return err
		}
	}

	return nil
}

var volumeTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["pending","available","pending_deletion","failed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		volumeTypeStatusPropEnum = append(volumeTypeStatusPropEnum, v)
	}
}

const (

	// VolumeStatusPending captures enum value "pending"
	VolumeStatusPending string = "pending"

	// VolumeStatusAvailable captures enum value "available"
	VolumeStatusAvailable string = "available"

	// VolumeStatusPendingDeletion captures enum value "pending_deletion"
	VolumeStatusPendingDeletion string = "pending_deletion"

	// VolumeStatusFailed captures enum value "failed"
	VolumeStatusFailed string = "failed"
)

// prop value enum
func (m *Volume) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, volumeTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Volume) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *Volume) validateVolumeAttachments(formats strfmt.Registry) error {

	if swag.IsZero(m.VolumeAttachments) { // not required
		return nil
	}

	for i := 0; i < len(m.VolumeAttachments); i++ {
		if swag.IsZero(m.VolumeAttachments[i]) { // not required
			continue
		}

		if m.VolumeAttachments[i] != nil {
			if err := m.VolumeAttachments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("volume_attachments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Volume) validateZone(formats strfmt.Registry) error {

	if swag.IsZero(m.Zone) { // not required
		return nil
	}

	if m.Zone != nil {
		if err := m.Zone.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("zone")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Volume) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Volume) UnmarshalBinary(b []byte) error {
	var res Volume
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
