// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

type DeviceConformsTo struct {
	common.BackboneElement
	Category      *common.CodeableConcept `json:"category,omitempty"`
	Specification common.CodeableConcept  `json:"specification"`
	Version       *string                 `json:"version,omitempty"`
}

type DeviceName struct {
	common.BackboneElement
	Display *bool          `json:"display,omitempty"`
	Type    DeviceNameType `json:"type"`
	Value   string         `json:"value"`
}

type DeviceVersion struct {
	common.BackboneElement
	Component   *common.Identifier      `json:"component,omitempty"`
	InstallDate *string                 `json:"installDate,omitempty"`
	Type        *common.CodeableConcept `json:"type,omitempty"`
	Value       string                  `json:"value"`
}

type DeviceUdiCarrier struct {
	common.BackboneElement
	CarrierAIDC      *string                    `json:"carrierAIDC,omitempty"`
	CarrierHRF       *string                    `json:"carrierHRF,omitempty"`
	DeviceIdentifier string                     `json:"deviceIdentifier"`
	EntryType        *DeviceUdiCarrierEntryType `json:"entryType,omitempty"`
	Issuer           string                     `json:"issuer"`
	Jurisdiction     *string                    `json:"jurisdiction,omitempty"`
}

type DeviceProperty struct {
	common.BackboneElement
	Type                 common.CodeableConcept  `json:"type"`
	ValueQuantity        *common.Quantity        `json:"valueQuantity,omitempty"`
	ValueCodeableConcept *common.CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueString          *string                 `json:"valueString,omitempty"`
	ValueBoolean         *bool                   `json:"valueBoolean,omitempty"`
	ValueInteger         *int                    `json:"valueInteger,omitempty"`
}

type Device struct {
	DomainResource

	ResourceType          string                   `json:"resourceType"` // Always "Device"
	AvailabilityStatus    *common.CodeableConcept  `json:"availabilityStatus,omitempty"`
	BiologicalSourceEvent *common.Identifier       `json:"biologicalSourceEvent,omitempty"`
	Category              []common.CodeableConcept `json:"category,omitempty"`
	ConformsTo            []DeviceConformsTo       `json:"conformsTo,omitempty"`
	Contact               []ContactPoint           `json:"contact,omitempty"`
	Cycle                 *Count                   `json:"cycle,omitempty"`
	Definition            *CodeableReference       `json:"definition,omitempty"`
	DisplayName           *string                  `json:"displayName,omitempty"`
	Duration              *Duration                `json:"duration,omitempty"`
	Endpoint              []common.Reference       `json:"endpoint,omitempty"`
	ExpirationDate        *string                  `json:"expirationDate,omitempty"`
	Gateway               []CodeableReference      `json:"gateway,omitempty"`
	Identifier            []common.Identifier      `json:"identifier,omitempty"`
	Location              *common.Reference        `json:"location,omitempty"`
	LotNumber             *string                  `json:"lotNumber,omitempty"`
	ManufactureDate       *string                  `json:"manufactureDate,omitempty"`
	Manufacturer          *string                  `json:"manufacturer,omitempty"`
	Mode                  *common.CodeableConcept  `json:"mode,omitempty"`
	ModelNumber           *string                  `json:"modelNumber,omitempty"`
	Name                  []DeviceName             `json:"name,omitempty"`
	Note                  []Annotation             `json:"note,omitempty"`
	Owner                 *common.Reference        `json:"owner,omitempty"`
	Parent                *common.Reference        `json:"parent,omitempty"`
	PartNumber            *string                  `json:"partNumber,omitempty"`
	Property              []DeviceProperty         `json:"property,omitempty"`
	Safety                []common.CodeableConcept `json:"safety,omitempty"`
	SerialNumber          *string                  `json:"serialNumber,omitempty"`
	Status                *DeviceStatus            `json:"status,omitempty"`
	Type                  []common.CodeableConcept `json:"type,omitempty"`
	UdiCarrier            []DeviceUdiCarrier       `json:"udiCarrier,omitempty"`
	Url                   *string                  `json:"url,omitempty"`
	Version               []DeviceVersion          `json:"version,omitempty"`
}

type DeviceNameType string

const (
	DeviceNameTypeRegisteredName      DeviceNameType = "registered-name"
	DeviceNameTypeUserFriendlyName    DeviceNameType = "user-friendly-name"
	DeviceNameTypePatientReportedName DeviceNameType = "patient-reported-name"
)

type DeviceUdiCarrierEntryType string

const (
	DeviceUdiCarrierEntryTypeBarcode                DeviceUdiCarrierEntryType = "barcode"
	DeviceUdiCarrierEntryTypeRfid                   DeviceUdiCarrierEntryType = "rfid"
	DeviceUdiCarrierEntryTypeManual                 DeviceUdiCarrierEntryType = "manual"
	DeviceUdiCarrierEntryTypeCard                   DeviceUdiCarrierEntryType = "card"
	DeviceUdiCarrierEntryTypeSelfReported           DeviceUdiCarrierEntryType = "self-reported"
	DeviceUdiCarrierEntryTypeElectronicTransmission DeviceUdiCarrierEntryType = "electronic-transmission"
	DeviceUdiCarrierEntryTypeUnknown                DeviceUdiCarrierEntryType = "unknown"
)

type DeviceStatus string

const (
	DeviceStatusActive         DeviceStatus = "active"
	DeviceStatusInactive       DeviceStatus = "inactive"
	DeviceStatusEnteredInError DeviceStatus = "entered-in-error"
)
