// Package fhir5 contains FHIR R5 resource definitions
package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// SupplyRequestParameter represents specific parameters for the ordered item
type SupplyRequestParameter struct {
	common.BackboneElement

	// A code or string that identifies the device detail being asserted
	Code *common.CodeableConcept `json:"code,omitempty"`

	// Range means device should have a value that falls somewhere within the specified range
	ValueCodeableConcept *common.CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueQuantity        *common.Quantity        `json:"valueQuantity,omitempty"`
	ValueRange           *Range                  `json:"valueRange,omitempty"`
	ValueBoolean         *bool                   `json:"valueBoolean,omitempty"`
}

// SupplyRequest represents a record of a request to deliver a medication, substance or device
type SupplyRequest struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "SupplyRequest"

	// When the request was made
	AuthoredOn *string `json:"authoredOn,omitempty"`

	// Plan/proposal/order fulfilled by this request
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// Category of supply, e.g. central, non-stock, etc
	Category *common.CodeableConcept `json:"category,omitempty"`

	// The patient to whom the supply will be given or for whom they will be used
	DeliverFor *common.Reference `json:"deliverFor,omitempty"`

	// Where the supply is expected to come from
	DeliverFrom *common.Reference `json:"deliverFrom,omitempty"`

	// Where the supply is destined to go
	DeliverTo *common.Reference `json:"deliverTo,omitempty"`

	// The identifier.type element is used to distinguish between the identifiers assigned by the requester/placer and the performer/filler
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Note that there's a difference between a prescription - an instruction to take a medication, along with a (sometimes) implicit supply, and an explicit request to supply, with no explicit instructions
	Item CodeableReference `json:"item"`

	// When the request should be fulfilled
	OccurrenceDateTime *string        `json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod   *common.Period `json:"occurrencePeriod,omitempty"`
	OccurrenceTiming   *Timing        `json:"occurrenceTiming,omitempty"`

	// Specific parameters for the ordered item
	Parameter []SupplyRequestParameter `json:"parameter,omitempty"`

	// Indicates how quickly this SupplyRequest should be addressed with respect to other requests
	Priority *SupplyRequestPriority `json:"priority,omitempty"`

	// The amount that is being ordered of the indicated item
	Quantity common.Quantity `json:"quantity"`

	// The reason why the supply item was requested
	Reason []CodeableReference `json:"reason,omitempty"`

	// The device, practitioner, etc. who initiated the request
	Requester *common.Reference `json:"requester,omitempty"`

	// Status of the supply request
	Status *SupplyRequestStatus `json:"status,omitempty"`

	// Who is intended to fulfill the request
	Supplier []common.Reference `json:"supplier,omitempty"`
}

// SupplyRequestPriority represents the priority of a supply request
type SupplyRequestPriority string

const (
	SupplyRequestPriorityRoutine SupplyRequestPriority = "routine"
	SupplyRequestPriorityUrgent  SupplyRequestPriority = "urgent"
	SupplyRequestPriorityASAP    SupplyRequestPriority = "asap"
	SupplyRequestPriorityStat    SupplyRequestPriority = "stat"
)

// SupplyRequestStatus represents the status of a supply request
type SupplyRequestStatus string

const (
	SupplyRequestStatusDraft          SupplyRequestStatus = "draft"
	SupplyRequestStatusActive         SupplyRequestStatus = "active"
	SupplyRequestStatusSuspended      SupplyRequestStatus = "suspended"
	SupplyRequestStatusCancelled      SupplyRequestStatus = "cancelled"
	SupplyRequestStatusCompleted      SupplyRequestStatus = "completed"
	SupplyRequestStatusEnteredInError SupplyRequestStatus = "entered-in-error"
	SupplyRequestStatusUnknown        SupplyRequestStatus = "unknown"
)
