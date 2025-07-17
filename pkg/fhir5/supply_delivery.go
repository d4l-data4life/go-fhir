// Package fhir5 contains FHIR R5 resource definitions
package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// SupplyDeliverySuppliedItem represents the item that is being delivered or has been supplied
type SupplyDeliverySuppliedItem struct {
	common.BackboneElement

	// Identifies the medication, substance, device or biologically derived product being supplied
	ItemCodeableConcept *common.CodeableConcept `json:"itemCodeableConcept,omitempty"`
	ItemReference       *common.Reference       `json:"itemReference,omitempty"`

	// The amount of the item that has been supplied
	Quantity *common.Quantity `json:"quantity,omitempty"`
}

// SupplyDelivery represents record of delivery of what is supplied
type SupplyDelivery struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "SupplyDelivery"

	// A plan, proposal or order that is fulfilled in whole or in part by this event
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// Identification of the facility/location where the delivery was shipped to
	Destination *common.Reference `json:"destination,omitempty"`

	// This identifier is typically assigned by the supplier
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// When the delivery occurred
	OccurrenceDateTime *string        `json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod   *common.Period `json:"occurrencePeriod,omitempty"`
	OccurrenceTiming   *Timing        `json:"occurrenceTiming,omitempty"`

	// Not to be used to link an event to an Encounter - use Event.context for that
	PartOf []common.Reference `json:"partOf,omitempty"`

	// A link to a resource representing the person whom the delivered item is for
	Patient *common.Reference `json:"patient,omitempty"`

	// Identifies the individual or organization that received the delivery
	Receiver []common.Reference `json:"receiver,omitempty"`

	// This element is labeled as a modifier because the status contains codes that mark the resource as not currently valid
	Status *SupplyDeliveryStatus `json:"status,omitempty"`

	// The item that is being delivered or has been supplied
	SuppliedItem []SupplyDeliverySuppliedItem `json:"suppliedItem,omitempty"`

	// The individual or organization responsible for supplying the delivery
	Supplier *common.Reference `json:"supplier,omitempty"`

	// Indicates the type of supply being provided
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// SupplyDeliveryStatus represents the status of a supply delivery
type SupplyDeliveryStatus string

const (
	SupplyDeliveryStatusInProgress     SupplyDeliveryStatus = "in-progress"
	SupplyDeliveryStatusCompleted      SupplyDeliveryStatus = "completed"
	SupplyDeliveryStatusAbandoned      SupplyDeliveryStatus = "abandoned"
	SupplyDeliveryStatusEnteredInError SupplyDeliveryStatus = "entered-in-error"
)
