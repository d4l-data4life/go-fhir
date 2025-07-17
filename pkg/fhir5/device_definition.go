package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// DeviceDefinitionUdiDeviceIdentifierMarketDistribution represents market distribution information
type DeviceDefinitionUdiDeviceIdentifierMarketDistribution struct {
	common.BackboneElement

	// Begin and end dates for the commercial distribution of the device
	MarketPeriod common.Period `json:"marketPeriod"`

	// National state or territory to which the marketDistribution refers
	SubJurisdiction string `json:"subJurisdiction"`
}

// DeviceDefinitionUdiDeviceIdentifier represents UDI device identifier information
type DeviceDefinitionUdiDeviceIdentifier struct {
	common.BackboneElement

	// The identifier that is to be associated with every Device
	DeviceIdentifier string `json:"deviceIdentifier"`

	// The organization that assigns the identifier algorithm
	Issuer string `json:"issuer"`

	// The jurisdiction to which the deviceIdentifier applies
	Jurisdiction string `json:"jurisdiction"`

	// Indicates where and when the device is available on the market
	MarketDistribution []DeviceDefinitionUdiDeviceIdentifierMarketDistribution `json:"marketDistribution,omitempty"`
}

// DeviceDefinitionRegulatoryIdentifier represents regulatory identifier information
type DeviceDefinitionRegulatoryIdentifier struct {
	common.BackboneElement

	// The identifier itself
	DeviceIdentifier string `json:"deviceIdentifier"`

	// The organization that issued this identifier
	Issuer string `json:"issuer"`

	// The jurisdiction to which the deviceIdentifier applies
	Jurisdiction string `json:"jurisdiction"`

	// The type of identifier itself
	Type DeviceDefinitionRegulatoryIdentifierType `json:"type"`
}

// DeviceDefinitionRegulatoryIdentifierType represents the type of regulatory identifier
type DeviceDefinitionRegulatoryIdentifierType string

const (
	DeviceDefinitionRegulatoryIdentifierTypeBasic   DeviceDefinitionRegulatoryIdentifierType = "basic"
	DeviceDefinitionRegulatoryIdentifierTypeMaster  DeviceDefinitionRegulatoryIdentifierType = "master"
	DeviceDefinitionRegulatoryIdentifierTypeLicense DeviceDefinitionRegulatoryIdentifierType = "license"
)

// DeviceDefinitionDeviceName represents device name information
type DeviceDefinitionDeviceName struct {
	common.BackboneElement

	// A human-friendly name that is used to refer to the device
	Name string `json:"name"`

	// The type of deviceName
	Type DeviceDefinitionDeviceNameType `json:"type"`
}

// DeviceDefinitionDeviceNameType represents the type of device name
type DeviceDefinitionDeviceNameType string

const (
	DeviceDefinitionDeviceNameTypeRegisteredName      DeviceDefinitionDeviceNameType = "registered-name"
	DeviceDefinitionDeviceNameTypeUserFriendlyName    DeviceDefinitionDeviceNameType = "user-friendly-name"
	DeviceDefinitionDeviceNameTypePatientReportedName DeviceDefinitionDeviceNameType = "patient-reported-name"
)

// DeviceDefinitionClassification represents device classification information
type DeviceDefinitionClassification struct {
	common.BackboneElement

	// Further information qualifying this classification of the device model
	Justification []RelatedArtifact `json:"justification,omitempty"`

	// A classification or risk class of the device model
	Type common.CodeableConcept `json:"type"`
}

// DeviceDefinitionConformsTo represents conformance information
type DeviceDefinitionConformsTo struct {
	common.BackboneElement

	// Describes the type of the standard, specification, or formal guidance
	Category *common.CodeableConcept `json:"category,omitempty"`

	// Standard, regulation, certification, or guidance website, document, or other publication
	Source []RelatedArtifact `json:"source,omitempty"`

	// Code that identifies the specific standard, specification, protocol, formal guidance, regulation, legislation, or certification scheme
	Specification common.CodeableConcept `json:"specification"`

	// Identifies the specific form or variant of the standard, specification, or formal guidance
	Version []string `json:"version,omitempty"`
}

// DeviceDefinitionHasPart represents device component information
type DeviceDefinitionHasPart struct {
	common.BackboneElement

	// Number of instances of the component device in the current device
	Count *int `json:"count,omitempty"`

	// Reference to the device that is part of the current device
	Reference common.Reference `json:"reference"`
}

// DeviceDefinitionPackagingDistributor represents packaging distributor information
type DeviceDefinitionPackagingDistributor struct {
	common.BackboneElement

	// Distributor's human-readable name
	Name *string `json:"name,omitempty"`

	// Distributor as an Organization resource
	OrganizationReference []common.Reference `json:"organizationReference,omitempty"`
}

// DeviceDefinitionPackaging represents packaging information
type DeviceDefinitionPackaging struct {
	common.BackboneElement

	// The number of items contained in the package (devices or sub-packages)
	Count *int `json:"count,omitempty"`

	// An organization that distributes the packaged device
	Distributor []DeviceDefinitionPackagingDistributor `json:"distributor,omitempty"`

	// The business identifier of the packaged medication
	Identifier *common.Identifier `json:"identifier,omitempty"`

	// Allows packages within packages
	Packaging []DeviceDefinitionPackaging `json:"packaging,omitempty"`

	// A code that defines the specific type of packaging
	Type *common.CodeableConcept `json:"type,omitempty"`

	// Unique Device Identifier (UDI) Barcode string on the packaging
	UdiDeviceIdentifier []DeviceDefinitionUdiDeviceIdentifier `json:"udiDeviceIdentifier,omitempty"`
}

// DeviceDefinitionVersion represents device version information
type DeviceDefinitionVersion struct {
	common.BackboneElement

	// Note that the module of the device would not need to be tracked as a separate device
	Component *common.Identifier `json:"component,omitempty"`

	// The type of the device version, e.g. manufacturer, approved, internal
	Type *common.CodeableConcept `json:"type,omitempty"`

	// The version of the device or software
	Value string `json:"value"`
}

// DeviceDefinitionProperty represents device property information
type DeviceDefinitionProperty struct {
	common.BackboneElement

	// Code that specifies the property being represented
	Type common.CodeableConcept `json:"type"`

	// Property value as a quantity
	ValueQuantity []common.Quantity `json:"valueQuantity,omitempty"`

	// Property value as a code, e.g. NTP4 (synced to NTP)
	ValueCode []common.CodeableConcept `json:"valueCode,omitempty"`
}

// DeviceDefinitionLink represents device link information
type DeviceDefinitionLink struct {
	common.BackboneElement

	// The type indicates the relationship of the related device to the device instance
	Relation common.CodeableConcept `json:"relation"`

	// A reference to the linked device
	RelatedDevice CodeableReference `json:"relatedDevice"`
}

// DeviceDefinitionMaterial represents device material information
type DeviceDefinitionMaterial struct {
	common.BackboneElement

	// Indicates an alternative material of the device
	Substance *common.CodeableConcept `json:"substance,omitempty"`

	// Indicates if the material is an alternative to the base material
	Alternate *bool `json:"alternate,omitempty"`

	// Whether the substance is a known or suspected allergen
	AllergenicIndicator *bool `json:"allergenicIndicator,omitempty"`
}

// DeviceDefinitionGuideline represents device guideline information
type DeviceDefinitionGuideline struct {
	common.BackboneElement

	// The circumstances that form the setting for using the device
	UseContext []UsageContext `json:"useContext,omitempty"`

	// Indicates what sort of device this is - if it is a sterile, disposable one, etc
	Indication []common.CodeableConcept `json:"indication,omitempty"`

	// A specific situation when a device should not be used because it may cause harm
	Contraindication []common.CodeableConcept `json:"contraindication,omitempty"`

	// Specific hazard alert information that a user needs to know before using the device
	Warning []common.CodeableConcept `json:"warning,omitempty"`

	// A description of the general purpose or medical use of the device or function of the device
	IntendedUse *string `json:"intendedUse,omitempty"`
}

// DeviceDefinitionCorrectiveAction represents corrective action information
type DeviceDefinitionCorrectiveAction struct {
	common.BackboneElement

	// The type of corrective action
	Type common.CodeableConcept `json:"type"`

	// The date and time when the corrective action was initiated
	Period *common.Period `json:"period,omitempty"`

	// The device that has been corrected
	Device common.Reference `json:"device"`

	// The scope of the corrective action - whether the action targeted all instances of this device
	Scope *DeviceDefinitionCorrectiveActionScope `json:"scope,omitempty"`
}

// DeviceDefinitionCorrectiveActionScope represents the scope of corrective action
type DeviceDefinitionCorrectiveActionScope string

const (
	DeviceDefinitionCorrectiveActionScopeModel         DeviceDefinitionCorrectiveActionScope = "model"
	DeviceDefinitionCorrectiveActionScopeLotNumbers    DeviceDefinitionCorrectiveActionScope = "lot-numbers"
	DeviceDefinitionCorrectiveActionScopeSerialNumbers DeviceDefinitionCorrectiveActionScope = "serial-numbers"
	DeviceDefinitionCorrectiveActionScopeDevices       DeviceDefinitionCorrectiveActionScope = "devices"
)

// DeviceDefinitionChargeItem represents charge item information
type DeviceDefinitionChargeItem struct {
	common.BackboneElement

	// The type of charge item
	ChargeItemCode common.CodeableConcept `json:"chargeItemCode"`

	// The context to which this charge item applies
	UseContext []UsageContext `json:"useContext,omitempty"`
}

// DeviceDefinition represents the characteristics, operational status and capabilities of a medical-related component of a medical device
type DeviceDefinition struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "DeviceDefinition"

	// Billing code or reference associated with the device
	ChargeItem []DeviceDefinitionChargeItem `json:"chargeItem,omitempty"`

	// In this element various classifications can be used, such as GMDN, EMDN, SNOMED CT, risk classes, national product codes
	Classification []DeviceDefinitionClassification `json:"classification,omitempty"`

	// Identifies the standards, specifications, or formal guidances for the capabilities supported by the device
	ConformsTo []DeviceDefinitionConformsTo `json:"conformsTo,omitempty"`

	// Used for troubleshooting etc
	Contact []ContactPoint `json:"contact,omitempty"`

	// Tracking of latest field safety corrective action
	CorrectiveAction *DeviceDefinitionCorrectiveAction `json:"correctiveAction,omitempty"`

	// Additional information to describe the device
	Description *string `json:"description,omitempty"`

	// The name or names of the device as given by the manufacturer
	DeviceName []DeviceDefinitionDeviceName `json:"deviceName,omitempty"`

	// For more structured data, a ClinicalUseDefinition that points to the DeviceDefinition can be used
	Guideline *DeviceDefinitionGuideline `json:"guideline,omitempty"`

	// A device that is part (for example a component) of the present device
	HasPart []DeviceDefinitionHasPart `json:"hasPart,omitempty"`

	// Unique instance identifiers assigned to a device by the software, manufacturers, other organizations or owners
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Language code for the human-readable text strings produced by the device (all supported)
	LanguageCode []common.CodeableConcept `json:"languageCode,omitempty"`

	// An associated device, attached to, used with, communicating with or linking a previous or new device model to the focal device
	Link []DeviceDefinitionLink `json:"link,omitempty"`

	// A name of the manufacturer or legal representative e.g. labeler
	Manufacturer *common.Reference `json:"manufacturer,omitempty"`

	// A substance used to create the material(s) of which the device is made
	Material []DeviceDefinitionMaterial `json:"material,omitempty"`

	// The model number for the device for example as defined by the manufacturer or labeler, or other agency
	ModelNumber *string `json:"modelNumber,omitempty"`

	// Descriptive information, usage information or implantation information that is not captured in an existing element
	Note []Annotation `json:"note,omitempty"`

	// An organization that is responsible for the provision and ongoing maintenance of the device
	Owner *common.Reference `json:"owner,omitempty"`

	// Information about the packaging of the device, i.e. how the device is packaged
	Packaging []DeviceDefinitionPackaging `json:"packaging,omitempty"`

	// Alphanumeric Maximum 20
	PartNumber *string `json:"partNumber,omitempty"`

	// Indicates the production identifier(s) that are expected to appear in the UDI carrier on the device label
	ProductionIdentifierInUDI []DeviceDefinitionProductionIdentifierInUDI `json:"productionIdentifierInUDI,omitempty"`

	// Dynamic or current properties, such as settings, of an individual device are described using a Device instance-specific DeviceMetric
	Property []DeviceDefinitionProperty `json:"property,omitempty"`

	// This should not be used for regulatory authorization numbers which are to be captured elsewhere
	RegulatoryIdentifier []DeviceDefinitionRegulatoryIdentifier `json:"regulatoryIdentifier,omitempty"`

	// Safety characteristics of the device
	Safety []common.CodeableConcept `json:"safety,omitempty"`

	// Shelf Life and storage information
	ShelfLifeStorage []ProductShelfLife `json:"shelfLifeStorage,omitempty"`

	// Unique device identifier (UDI) assigned to device label or package
	UdiDeviceIdentifier []DeviceDefinitionUdiDeviceIdentifier `json:"udiDeviceIdentifier,omitempty"`

	// The version of the device or software
	Version []DeviceDefinitionVersion `json:"version,omitempty"`
}

// DeviceDefinitionProductionIdentifierInUDI represents production identifier types
type DeviceDefinitionProductionIdentifierInUDI string

const (
	DeviceDefinitionProductionIdentifierInUDILotNumber        DeviceDefinitionProductionIdentifierInUDI = "lot-number"
	DeviceDefinitionProductionIdentifierInUDIManufacturedDate DeviceDefinitionProductionIdentifierInUDI = "manufactured-date"
	DeviceDefinitionProductionIdentifierInUDISerialNumber     DeviceDefinitionProductionIdentifierInUDI = "serial-number"
	DeviceDefinitionProductionIdentifierInUDIExpirationDate   DeviceDefinitionProductionIdentifierInUDI = "expiration-date"
	DeviceDefinitionProductionIdentifierInUDIBiologicalSource DeviceDefinitionProductionIdentifierInUDI = "biological-source"
	DeviceDefinitionProductionIdentifierInUDISoftwareVersion  DeviceDefinitionProductionIdentifierInUDI = "software-version"
)

// ProductShelfLife represents shelf life and storage information
type ProductShelfLife struct {
	common.BackboneElement

	// The shelf life time period can be specified using a numerical value for the period of time and its unit of time measurement
	Period Duration `json:"period"`

	// Special precautions for storage, if any, can be specified using an appropriate controlled vocabulary
	SpecialPrecautionsForStorage []common.CodeableConcept `json:"specialPrecautionsForStorage,omitempty"`

	// The shelf life type of the product
	Type common.CodeableConcept `json:"type"`
}
