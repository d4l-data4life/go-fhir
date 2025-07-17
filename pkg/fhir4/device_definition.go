package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// DeviceDefinition represents the DeviceDefinition resource
type DeviceDefinition struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "DeviceDefinition"

	// Device capabilities
	Capability []DeviceDefinitionCapability `json:"capability,omitempty"`

	// Used for troubleshooting etc
	Contact []common.ContactPoint `json:"contact,omitempty"`

	// A name given to the device to identify it
	DeviceName []DeviceDefinitionDeviceName `json:"deviceName,omitempty"`

	// Unique instance identifiers assigned to a device by the software, manufacturers, other organizations or owners
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Language code for the human-readable text strings produced by the device (all supported)
	LanguageCode []common.CodeableConcept `json:"languageCode,omitempty"`

	// A name of the manufacturer
	ManufacturerString *string `json:"manufacturerString,omitempty"`

	// A name of the manufacturer
	ManufacturerReference *common.Reference `json:"manufacturerReference,omitempty"`

	// A substance used to create the material(s) of which the device is made
	Material []DeviceDefinitionMaterial `json:"material,omitempty"`

	// The model number for the device
	ModelNumber *string `json:"modelNumber,omitempty"`

	// Descriptive information, usage information or implantation information that is not captured in an existing element
	Note []common.Annotation `json:"note,omitempty"`

	// Access to on-line information about the device
	OnlineInformation *string `json:"onlineInformation,omitempty"`

	// An organization that is responsible for the provision and ongoing maintenance of the device
	Owner *common.Reference `json:"owner,omitempty"`

	// The parent device it can be part of
	ParentDevice *common.Reference `json:"parentDevice,omitempty"`

	// Dimensions, color etc
	PhysicalCharacteristics *ProdCharacteristic `json:"physicalCharacteristics,omitempty"`

	// The actual configuration settings of a device as it actually operates, e.g., regulation status, time properties
	Property []DeviceDefinitionProperty `json:"property,omitempty"`

	// The quantity of the device present in the packaging (e.g. the number of devices present in a pack, or the number of devices in the same package of the medicinal product)
	Quantity *common.Quantity `json:"quantity,omitempty"`

	// Safety characteristics of the device
	Safety []common.CodeableConcept `json:"safety,omitempty"`

	// Shelf Life and storage information
	ShelfLifeStorage []ProductShelfLife `json:"shelfLifeStorage,omitempty"`

	// The capabilities supported on a device, the standards to which the device conforms for a particular purpose, and used for the communication
	Specialization []DeviceDefinitionSpecialization `json:"specialization,omitempty"`

	// What kind of device or device system this is
	Type *common.CodeableConcept `json:"type,omitempty"`

	// Unique device identifier (UDI) assigned to device label or package
	UdiDeviceIdentifier []DeviceDefinitionUdiDeviceIdentifier `json:"udiDeviceIdentifier,omitempty"`

	// If the device is running a FHIR server, the network address should be the Base URL from which a conformance statement may be retrieved
	Url *string `json:"url,omitempty"`

	// The available versions of the device, e.g., software versions
	Version []string `json:"version,omitempty"`
}

// DeviceDefinitionUdiDeviceIdentifier represents UDI device identifier
type DeviceDefinitionUdiDeviceIdentifier struct {
	common.BackboneElement

	// The identifier that is to be associated with every Device that references this DeviceDefinition
	DeviceIdentifier string `json:"deviceIdentifier"`

	// The organization that assigns the identifier algorithm
	Issuer string `json:"issuer"`

	// The jurisdiction to which the deviceIdentifier applies
	Jurisdiction string `json:"jurisdiction"`
}

// DeviceDefinitionDeviceName represents a name given to the device to identify it
type DeviceDefinitionDeviceName struct {
	common.BackboneElement

	// The name of the device
	Name string `json:"name"`

	// The type of deviceName
	Type DeviceDefinitionDeviceNameType `json:"type"`
}

// DeviceDefinitionDeviceNameType represents the type of device name
type DeviceDefinitionDeviceNameType string

const (
	DeviceDefinitionDeviceNameTypeUdiLabelName        DeviceDefinitionDeviceNameType = "udi-label-name"
	DeviceDefinitionDeviceNameTypeUserFriendlyName    DeviceDefinitionDeviceNameType = "user-friendly-name"
	DeviceDefinitionDeviceNameTypePatientReportedName DeviceDefinitionDeviceNameType = "patient-reported-name"
	DeviceDefinitionDeviceNameTypeManufacturerName    DeviceDefinitionDeviceNameType = "manufacturer-name"
	DeviceDefinitionDeviceNameTypeModelName           DeviceDefinitionDeviceNameType = "model-name"
	DeviceDefinitionDeviceNameTypeOther               DeviceDefinitionDeviceNameType = "other"
)

// DeviceDefinitionSpecialization represents the capabilities supported on a device
type DeviceDefinitionSpecialization struct {
	common.BackboneElement

	// The standard that is used to operate and communicate
	SystemType string `json:"systemType"`

	// The version of the standard that is used to operate and communicate
	Version *string `json:"version,omitempty"`
}

// DeviceDefinitionCapability represents device capabilities
type DeviceDefinitionCapability struct {
	common.BackboneElement

	// Description of capability
	Description []common.CodeableConcept `json:"description,omitempty"`

	// Type of capability
	Type common.CodeableConcept `json:"type"`
}

// DeviceDefinitionProperty represents the actual configuration settings of a device
type DeviceDefinitionProperty struct {
	common.BackboneElement

	// Code that specifies the property DeviceDefinitionPropertyCode (Extensible)
	Type common.CodeableConcept `json:"type"`

	// Property value as a code, e.g., NTP4 (synced to NTP)
	ValueCode []common.CodeableConcept `json:"valueCode,omitempty"`

	// Property value as a quantity
	ValueQuantity []common.Quantity `json:"valueQuantity,omitempty"`
}

// DeviceDefinitionMaterial represents a substance used to create the material(s) of which the device is made
type DeviceDefinitionMaterial struct {
	common.BackboneElement

	// Whether the substance is a known or suspected allergen
	AllergenicIndicator *bool `json:"allergenicIndicator,omitempty"`

	// Indicates an alternative material of the device
	Alternate *bool `json:"alternate,omitempty"`

	// The substance
	Substance common.CodeableConcept `json:"substance"`
}

// ProdCharacteristic represents product characteristics
type ProdCharacteristic struct {
	common.BackboneElement

	// A code for the characteristic
	Code *common.CodeableConcept `json:"code,omitempty"`

	// The value of the characteristic
	ValueCodeableConcept *common.CodeableConcept `json:"valueCodeableConcept,omitempty"`

	// The value of the characteristic
	ValueQuantity *common.Quantity `json:"valueQuantity,omitempty"`

	// The value of the characteristic
	ValueString *string `json:"valueString,omitempty"`

	// The value of the characteristic
	ValueBoolean *bool `json:"valueBoolean,omitempty"`

	// The value of the characteristic
	ValueAttachment *common.Attachment `json:"valueAttachment,omitempty"`
}

// ProductShelfLife represents shelf life and storage information
type ProductShelfLife struct {
	common.BackboneElement

	// This describes the shelf life, taking into account various scenarios such as shelf life of the packaged Medicinal Product itself, shelf life after transformation where necessary and shelf life after the first opening of a bottle, etc
	Period *common.Quantity `json:"period,omitempty"`

	// Special precautions for storage, if any, can be specified using an appropriate controlled vocabulary The controlled term and the controlled term identifier shall be specified
	SpecialPrecautionsForStorage []common.CodeableConcept `json:"specialPrecautionsForStorage,omitempty"`

	// Shelf life time period can be specified using a controlled vocabulary The controlled term and the controlled term identifier shall be specified
	Type *common.CodeableConcept `json:"type,omitempty"`
}
