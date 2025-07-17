package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// SupplyRequest represents a record of a request for a medication, substance or device used in the healthcare setting
type SupplyRequest struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "SupplyRequest"

	// When the request was made
	AuthoredOn *string `json:"authoredOn,omitempty"`

	// Category of supply, e.g. central, non-stock, etc. This is used to support work flows associated with the supply process
	Category *common.CodeableConcept `json:"category,omitempty"`

	// Where the supply is expected to come from
	DeliverFrom *common.Reference `json:"deliverFrom,omitempty"`

	// Where the supply is destined to go
	DeliverTo *common.Reference `json:"deliverTo,omitempty"`

	// The identifier.type element is used to distinguish between the identifiers assigned by the requester/placer and the performer/filler
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Note that there's a difference between a prescription - an instruction to take a medication, along with a (sometimes) implicit supply, and an explicit request to supply, with no explicit instructions
	ItemCodeableConcept *common.CodeableConcept `json:"itemCodeableConcept,omitempty"`

	// Note that there's a difference between a prescription - an instruction to take a medication, along with a (sometimes) implicit supply, and an explicit request to supply, with no explicit instructions
	ItemReference *common.Reference `json:"itemReference,omitempty"`

	// When the request should be fulfilled
	OccurrenceDateTime *string `json:"occurrenceDateTime,omitempty"`

	// When the request should be fulfilled
	OccurrencePeriod *common.Period `json:"occurrencePeriod,omitempty"`

	// When the request should be fulfilled
	OccurrenceTiming *common.Timing `json:"occurrenceTiming,omitempty"`

	// Specific parameters for the ordered item. For example, the size of the indicated item
	Parameter []SupplyRequestParameter `json:"parameter,omitempty"`

	// Indicates how quickly this SupplyRequest should be addressed with respect to other requests
	Priority *string `json:"priority,omitempty"` // "routine" | "urgent" | "asap" | "stat"

	// The amount that is being ordered of the indicated item
	Quantity common.Quantity `json:"quantity"`

	// The reason why the supply item was requested
	ReasonCode []common.CodeableConcept `json:"reasonCode,omitempty"`

	// The reason why the supply item was requested
	ReasonReference []common.Reference `json:"reasonReference,omitempty"`

	// The device, practitioner, etc. who initiated the request
	Requester *common.Reference `json:"requester,omitempty"`

	// Status of the supply request
	Status *string `json:"status,omitempty"` // "draft" | "active" | "suspended" | "cancelled" | "completed" | "entered-in-error" | "unknown"

	// Who is intended to fulfill the request
	Supplier []common.Reference `json:"supplier,omitempty"`
}

// SupplyRequestParameter represents specific parameters for the ordered item. For example, the size of the indicated item
type SupplyRequestParameter struct {
	common.BackboneElement

	// A code or string that identifies the device detail being asserted
	Code *common.CodeableConcept `json:"code,omitempty"`

	// Range means device should have a value that falls somewhere within the specified range
	ValueCodeableConcept *common.CodeableConcept `json:"valueCodeableConcept,omitempty"`

	// Range means device should have a value that falls somewhere within the specified range
	ValueQuantity *common.Quantity `json:"valueQuantity,omitempty"`

	// Range means device should have a value that falls somewhere within the specified range
	ValueRange *common.Range `json:"valueRange,omitempty"`

	// Range means device should have a value that falls somewhere within the specified range
	ValueBoolean *bool `json:"valueBoolean,omitempty"`
}
