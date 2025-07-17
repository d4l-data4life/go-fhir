package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// Device represents a medical or non-medical device used in healthcare
type Device struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Device"

	// Instance identifier
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Reference to the definition for the device
	Definition *common.Reference `json:"definition,omitempty"`

	// Unique Device Identifier (UDI) Barcode string
	UdiCarrier []DeviceUdiCarrier `json:"udiCarrier,omitempty"`

	// active | inactive | entered-in-error | unknown
	Status *DeviceStatus `json:"status,omitempty"`

	// Reason for the device's status
	StatusReason []common.CodeableConcept `json:"statusReason,omitempty"`

	// For example, this applies to devices in the United States regulated under Code of Federal Regulation 21CFR§1271.290(c)
	DistinctIdentifier *string `json:"distinctIdentifier,omitempty"`

	// Name of device manufacturer
	Manufacturer *string `json:"manufacturer,omitempty"`

	// Date when the device was made
	ManufactureDate *string `json:"manufactureDate,omitempty"`

	// Date and time of expiry of this device (if applicable)
	ExpirationDate *string `json:"expirationDate,omitempty"`

	// Lot number assigned by the manufacturer
	LotNumber *string `json:"lotNumber,omitempty"`

	// The serial number assigned by the organization when the device was manufactured
	SerialNumber *string `json:"serialNumber,omitempty"`

	// The names of the device as given by the manufacturer
	DeviceName []DeviceDeviceName `json:"deviceName,omitempty"`

	// The model number for the device
	ModelNumber *string `json:"modelNumber,omitempty"`

	// The part number of the device
	PartNumber *string `json:"partNumber,omitempty"`

	// The kind or type of device
	Type *common.CodeableConcept `json:"type,omitempty"`

	// Device capabilities, specializations, or standards supported
	Specialization []DeviceSpecialization `json:"specialization,omitempty"`

	// The actual configuration settings of a device
	Property []DeviceProperty `json:"property,omitempty"`

	// Patient to whom Device is affixed
	Patient *common.Reference `json:"patient,omitempty"`

	// Organization responsible for device
	Owner *common.Reference `json:"owner,omitempty"`

	// Details for human/organization for support
	Contact []common.ContactPoint `json:"contact,omitempty"`

	// Where the device is found
	Location *common.Reference `json:"location,omitempty"`

	// Network address to contact device
	Url *string `json:"url,omitempty"`
}

// DeviceStatus represents the status of the device
type DeviceStatus string

const (
	DeviceStatusActive         DeviceStatus = "active"
	DeviceStatusInactive       DeviceStatus = "inactive"
	DeviceStatusEnteredInError DeviceStatus = "entered-in-error"
	DeviceStatusUnknown        DeviceStatus = "unknown"
)

// DeviceUdiCarrier represents UDI may identify an unique instance of a device
type DeviceUdiCarrier struct {
	common.BackboneElement

	// The AIDC form of UDIs should be scanned or otherwise used for the identification of the device
	CarrierAIDC *string `json:"carrierAIDC,omitempty"`

	// If separate barcodes for DI and PI are present, concatenate the string with DI first
	CarrierHRF *string `json:"carrierHRF,omitempty"`

	// The device identifier (DI) is a mandatory, fixed portion of a UDI
	DeviceIdentifier *string `json:"deviceIdentifier,omitempty"`

	// A coded entry to indicate how the data was entered
	EntryType *DeviceUdiCarrierEntryType `json:"entryType,omitempty"`

	// Organization that is charged with issuing UDIs for devices
	Issuer *string `json:"issuer,omitempty"`

	// The identity of the authoritative source for UDI generation within a jurisdiction
	Jurisdiction *string `json:"jurisdiction,omitempty"`
}

// DeviceUdiCarrierEntryType represents a coded entry to indicate how the data was entered
type DeviceUdiCarrierEntryType string

const (
	DeviceUdiCarrierEntryTypeBarcode      DeviceUdiCarrierEntryType = "barcode"
	DeviceUdiCarrierEntryTypeRfid         DeviceUdiCarrierEntryType = "rfid"
	DeviceUdiCarrierEntryTypeManual       DeviceUdiCarrierEntryType = "manual"
	DeviceUdiCarrierEntryTypeCard         DeviceUdiCarrierEntryType = "card"
	DeviceUdiCarrierEntryTypeSelfReported DeviceUdiCarrierEntryType = "self-reported"
	DeviceUdiCarrierEntryTypeUnknown      DeviceUdiCarrierEntryType = "unknown"
)

// DeviceDeviceName represents the manufacturer's name of the device
type DeviceDeviceName struct {
	common.BackboneElement

	// The name of the device
	Name string `json:"name"`

	// The type of deviceName
	Type DeviceDeviceNameType `json:"type"`
}

// DeviceDeviceNameType represents the type of device name
type DeviceDeviceNameType string

const (
	DeviceDeviceNameTypeUdiLabelName        DeviceDeviceNameType = "udi-label-name"
	DeviceDeviceNameTypeUserFriendlyName    DeviceDeviceNameType = "user-friendly-name"
	DeviceDeviceNameTypePatientReportedName DeviceDeviceNameType = "patient-reported-name"
	DeviceDeviceNameTypeManufacturerName    DeviceDeviceNameType = "manufacturer-name"
	DeviceDeviceNameTypeModelName           DeviceDeviceNameType = "model-name"
	DeviceDeviceNameTypeOther               DeviceDeviceNameType = "other"
)

// DeviceSpecialization represents the capabilities supported on a device
type DeviceSpecialization struct {
	common.BackboneElement

	// The standard that is used to operate and communicate
	SystemType common.CodeableConcept `json:"systemType"`

	// The version of the standard that is used to operate and communicate
	Version *string `json:"version,omitempty"`
}

// DeviceProperty represents the actual configuration settings of a device
type DeviceProperty struct {
	common.BackboneElement

	// Code that specifies the property DeviceDefinitionPropertyCode (Extensible)
	Type common.CodeableConcept `json:"type"`

	// Property value as a code, e.g., NTP4 (synced to NTP)
	ValueCode []common.CodeableConcept `json:"valueCode,omitempty"`

	// Property value as a quantity
	ValueQuantity []common.Quantity `json:"valueQuantity,omitempty"`
}
