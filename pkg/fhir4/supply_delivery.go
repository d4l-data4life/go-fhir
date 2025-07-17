package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// SupplyDelivery represents record of delivery of what is supplied
type SupplyDelivery struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "SupplyDelivery"

	// A plan, proposal or order that is fulfilled in whole or in part by this event
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// Identification of the facility/location where the Supply was shipped to, as part of the dispense event
	Destination *common.Reference `json:"destination,omitempty"`

	// This identifier is typically assigned by the dispenser, and may be used to reference the delivery when exchanging information about it with other systems
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// [The list of types may be constrained as appropriate for the type of event]
	OccurrenceDateTime *string `json:"occurrenceDateTime,omitempty"`

	// [The list of types may be constrained as appropriate for the type of event]
	OccurrencePeriod *common.Period `json:"occurrencePeriod,omitempty"`

	// [The list of types may be constrained as appropriate for the type of event]
	OccurrenceTiming *common.Timing `json:"occurrenceTiming,omitempty"`

	// Not to be used to link an event to an Encounter - use Event.context for that
	PartOf []common.Reference `json:"partOf,omitempty"`

	// A link to a resource representing the person whom the delivered item is for
	Patient *common.Reference `json:"patient,omitempty"`

	// Identifies the person who picked up the Supply
	Receiver []common.Reference `json:"receiver,omitempty"`

	// This element is labeled as a modifier because the status contains codes that mark the resource as not currently valid
	Status *string `json:"status,omitempty"` // "in-progress" | "completed" | "abandoned" | "entered-in-error"

	// The item that is being delivered or has been supplied
	SuppliedItem *SupplyDeliverySuppliedItem `json:"suppliedItem,omitempty"`

	// The individual responsible for dispensing the medication, supplier or device
	Supplier *common.Reference `json:"supplier,omitempty"`

	// Indicates the type of dispensing event that is performed. Examples include: Trial Fill, Completion of Trial, Partial Fill, Emergency Fill, Samples, etc
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// SupplyDeliverySuppliedItem represents the item that is being delivered or has been supplied
type SupplyDeliverySuppliedItem struct {
	common.BackboneElement

	// Identifies the medication, substance or device being dispensed. This is either a link to a resource representing the details of the item or a code that identifies the item from a known list
	ItemCodeableConcept *common.CodeableConcept `json:"itemCodeableConcept,omitempty"`

	// Identifies the medication, substance or device being dispensed. This is either a link to a resource representing the details of the item or a code that identifies the item from a known list
	ItemReference *common.Reference `json:"itemReference,omitempty"`

	// The amount of supply that has been dispensed. Includes unit of measure
	Quantity *common.Quantity `json:"quantity,omitempty"`
}
