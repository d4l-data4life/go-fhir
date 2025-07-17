package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// DeviceDispensePerformer represents who or what performed the event
type DeviceDispensePerformer struct {
	common.BackboneElement

	// The device, practitioner, etc. who performed the action
	Actor common.Reference `json:"actor"`

	// Distinguishes the type of performer in the dispense
	Function *common.CodeableConcept `json:"function,omitempty"`
}

// DeviceDispense represents a record of dispensation of a device
type DeviceDispense struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "DeviceDispense"

	// The order or request that this dispense is fulfilling
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// The category can be used to include where the device is expected to be consumed or other types of dispenses
	Category []common.CodeableConcept `json:"category,omitempty"`

	// Identification of the facility/location where the device was /should be shipped to, as part of the dispense process
	Destination *common.Reference `json:"destination,omitempty"`

	// Identifies the device being dispensed
	Device CodeableReference `json:"device"`

	// The encounter that establishes the context for this event
	Encounter *common.Reference `json:"encounter,omitempty"`

	// This might not include provenances for all versions of the request
	EventHistory []common.Reference `json:"eventHistory,omitempty"`

	// This is a business identifier, not a resource identifier
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The principal physical location where the dispense was performed
	Location *common.Reference `json:"location,omitempty"`

	// Extra information about the dispense that could not be conveyed in the other attributes
	Note []Annotation `json:"note,omitempty"`

	// The bigger event that this dispense is a part of
	PartOf []common.Reference `json:"partOf,omitempty"`

	// Indicates who or what performed the event
	Performer []DeviceDispensePerformer `json:"performer,omitempty"`

	// The time when the dispensed product was packaged and reviewed
	PreparedDate *string `json:"preparedDate,omitempty"`

	// The number of devices that have been dispensed
	Quantity *common.Quantity `json:"quantity,omitempty"`

	// Identifies the person who picked up the device or the person or location where the device was delivered
	Receiver *common.Reference `json:"receiver,omitempty"`

	// This element is labeled as a modifier because the status contains codes that mark the resource as not currently valid
	Status DeviceDispenseStatus `json:"status"`

	// The subject for whom the device was dispensed
	Subject *common.Reference `json:"subject,omitempty"`

	// Indicates the type of dispense event
	Type *common.CodeableConcept `json:"type,omitempty"`

	// The time the dispense was signed
	WhenHandedOver *string `json:"whenHandedOver,omitempty"`

	// The time when dispensed product was packaged and reviewed
	WhenPrepared *string `json:"whenPrepared,omitempty"`
}

// DeviceDispenseStatus represents the status of the device dispense
type DeviceDispenseStatus string

const (
	DeviceDispenseStatusPreparation    DeviceDispenseStatus = "preparation"
	DeviceDispenseStatusInProgress     DeviceDispenseStatus = "in-progress"
	DeviceDispenseStatusCancelled      DeviceDispenseStatus = "cancelled"
	DeviceDispenseStatusOnHold         DeviceDispenseStatus = "on-hold"
	DeviceDispenseStatusCompleted      DeviceDispenseStatus = "completed"
	DeviceDispenseStatusEnteredInError DeviceDispenseStatus = "entered-in-error"
	DeviceDispenseStatusStopped        DeviceDispenseStatus = "stopped"
	DeviceDispenseStatusDeclined       DeviceDispenseStatus = "declined"
	DeviceDispenseStatusUnknown        DeviceDispenseStatus = "unknown"
)
