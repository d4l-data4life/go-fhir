package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// DeviceRequest represents the DeviceRequest resource
type DeviceRequest struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "DeviceRequest"

	// When the request transitioned to being actionable
	AuthoredOn *string `json:"authoredOn,omitempty"`

	// Plan/proposal/order fulfilled by this request
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// The details of the device to be used
	CodeReference *common.Reference `json:"codeReference,omitempty"`

	// The details of the device to be used
	CodeCodeableConcept *common.CodeableConcept `json:"codeCodeableConcept,omitempty"`

	// An encounter that provides additional context in which this request is made
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Composite request this is part of
	GroupIdentifier *common.Identifier `json:"groupIdentifier,omitempty"`

	// Identifiers assigned to this order by the orderer or by the receiver
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Note: This is a business identifier, not a resource identifier
	InstantiatesCanonical []string `json:"instantiatesCanonical,omitempty"`

	// This might be an HTML page, PDF, etc. or could just be a non-resolvable URI identifier
	InstantiatesUri []string `json:"instantiatesUri,omitempty"`

	// Insurance plans, coverage extensions, pre-authorizations and/or pre-determinations that may be required for delivering the requested service
	Insurance []common.Reference `json:"insurance,omitempty"`

	// Whether the request is a proposal, plan, an original order or a reflex order
	Intent DeviceRequestIntent `json:"intent"`

	// Details about this request that were not represented at all or sufficiently in one of the attributes provided in a class
	Note []common.Annotation `json:"note,omitempty"`

	// The timing schedule for the use of the device
	OccurrenceDateTime *string `json:"occurrenceDateTime,omitempty"`

	// The timing schedule for the use of the device
	OccurrencePeriod *common.Period `json:"occurrencePeriod,omitempty"`

	// The timing schedule for the use of the device
	OccurrenceTiming *common.Timing `json:"occurrenceTiming,omitempty"`

	// Specific parameters for the ordered item
	Parameter []DeviceRequestParameter `json:"parameter,omitempty"`

	// The desired performer for doing the diagnostic testing
	Performer *common.Reference `json:"performer,omitempty"`

	// Desired type of performer for doing the diagnostic testing
	PerformerType *common.CodeableConcept `json:"performerType,omitempty"`

	// Indicates how quickly the request should be addressed with respect to other requests
	Priority *DeviceRequestPriority `json:"priority,omitempty"`

	// The request takes the place of the referenced completed or terminated request(s)
	PriorRequest []common.Reference `json:"priorRequest,omitempty"`

	// Reason or justification for the use of this device
	ReasonCode []common.CodeableConcept `json:"reasonCode,omitempty"`

	// Reason or justification for the use of this device
	ReasonReference []common.Reference `json:"reasonReference,omitempty"`

	// This might not include provenances for all versions of the request
	RelevantHistory []common.Reference `json:"relevantHistory,omitempty"`

	// The individual who initiated the request and has responsibility for its activation
	Requester *common.Reference `json:"requester,omitempty"`

	// This element is labeled as a modifier because the status contains the codes cancelled and entered-in-error that mark the request as not currently valid
	Status *DeviceRequestStatus `json:"status,omitempty"`

	// The patient who will use the device
	Subject common.Reference `json:"subject"`

	// Additional clinical information about the patient that may influence the request fulfilment
	SupportingInfo []common.Reference `json:"supportingInfo,omitempty"`
}

// DeviceRequestIntent represents whether the request is a proposal, plan, an original order or a reflex order
type DeviceRequestIntent string

const (
	DeviceRequestIntentProposal      DeviceRequestIntent = "proposal"
	DeviceRequestIntentPlan          DeviceRequestIntent = "plan"
	DeviceRequestIntentDirective     DeviceRequestIntent = "directive"
	DeviceRequestIntentOrder         DeviceRequestIntent = "order"
	DeviceRequestIntentOriginalOrder DeviceRequestIntent = "original-order"
	DeviceRequestIntentReflexOrder   DeviceRequestIntent = "reflex-order"
	DeviceRequestIntentFillerOrder   DeviceRequestIntent = "filler-order"
	DeviceRequestIntentInstanceOrder DeviceRequestIntent = "instance-order"
	DeviceRequestIntentOption        DeviceRequestIntent = "option"
)

// DeviceRequestPriority represents how quickly the request should be addressed
type DeviceRequestPriority string

const (
	DeviceRequestPriorityRoutine DeviceRequestPriority = "routine"
	DeviceRequestPriorityUrgent  DeviceRequestPriority = "urgent"
	DeviceRequestPriorityAsap    DeviceRequestPriority = "asap"
	DeviceRequestPriorityStat    DeviceRequestPriority = "stat"
)

// DeviceRequestStatus represents the status of the request
type DeviceRequestStatus string

const (
	DeviceRequestStatusDraft          DeviceRequestStatus = "draft"
	DeviceRequestStatusActive         DeviceRequestStatus = "active"
	DeviceRequestStatusOnHold         DeviceRequestStatus = "on-hold"
	DeviceRequestStatusRevoked        DeviceRequestStatus = "revoked"
	DeviceRequestStatusCompleted      DeviceRequestStatus = "completed"
	DeviceRequestStatusEnteredInError DeviceRequestStatus = "entered-in-error"
	DeviceRequestStatusUnknown        DeviceRequestStatus = "unknown"
)

// DeviceRequestParameter represents specific parameters for the ordered item
type DeviceRequestParameter struct {
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
