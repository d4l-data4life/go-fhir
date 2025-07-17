package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// DeviceRequestParameter represents specific parameters for the ordered item
type DeviceRequestParameter struct {
	common.BackboneElement

	// A code or string that identifies the device detail being asserted
	Code *common.CodeableConcept `json:"code,omitempty"`

	// Range means device should have a value that falls somewhere within the specified range
	ValueCodeableConcept *common.CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueQuantity        *common.Quantity        `json:"valueQuantity,omitempty"`
	ValueRange           *Range                  `json:"valueRange,omitempty"`
	ValueBoolean         *bool                   `json:"valueBoolean,omitempty"`
}

// DeviceRequest represents a request for a patient to employ a medical device
type DeviceRequest struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "DeviceRequest"

	// This status is to indicate whether the request is a PRN or may be given as needed
	AsNeeded *bool `json:"asNeeded,omitempty"`

	// The reason for using the device
	AsNeededFor *common.CodeableConcept `json:"asNeededFor,omitempty"`

	// When the request transitioned to being actionable
	AuthoredOn *string `json:"authoredOn,omitempty"`

	// Plan/proposal/order fulfilled by this request
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// The details of the device to be used
	Code CodeableReference `json:"code"`

	// If do not perform is not specified, the request is a positive request
	DoNotPerform *bool `json:"doNotPerform,omitempty"`

	// An encounter that provides additional context in which this request is made
	Encounter *common.Reference `json:"encounter,omitempty"`

	// A shared identifier common to multiple independent Request instances
	GroupIdentifier *common.Identifier `json:"groupIdentifier,omitempty"`

	// Identifiers assigned to this order by the orderer or by the receiver
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Note: This is a business identifier, not a resource identifier
	InstantiatesCanonical []string `json:"instantiatesCanonical,omitempty"`

	// This might be an HTML page, PDF, etc. or could just be a non-resolvable URI identifier
	InstantiatesUri []string `json:"instantiatesUri,omitempty"`

	// Insurance plans, coverage extensions, pre-authorizations and/or pre-determinations
	Insurance []common.Reference `json:"insurance,omitempty"`

	// Whether the request is a proposal, plan, an original order or a reflex order
	Intent DeviceRequestIntent `json:"intent"`

	// Details about this request that were not represented at all or sufficiently in one of the attributes
	Note []Annotation `json:"note,omitempty"`

	// The timing schedule for the use of the device
	OccurrenceDateTime *string        `json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod   *common.Period `json:"occurrencePeriod,omitempty"`
	OccurrenceTiming   *Timing        `json:"occurrenceTiming,omitempty"`

	// Specific parameters for the ordered item
	Parameter []DeviceRequestParameter `json:"parameter,omitempty"`

	// The desired individual or entity to provide the device to the subject of the request
	Performer *CodeableReference `json:"performer,omitempty"`

	// Indicates how quickly the request should be addressed with respect to other requests
	Priority *DeviceRequestPriority `json:"priority,omitempty"`

	// The number of devices to be provided
	Quantity *int `json:"quantity,omitempty"`

	// Reason or justification for the use of this device
	Reason []CodeableReference `json:"reason,omitempty"`

	// This might not include provenances for all versions of the request
	RelevantHistory []common.Reference `json:"relevantHistory,omitempty"`

	// The request takes the place of the referenced completed or terminated request(s)
	Replaces []common.Reference `json:"replaces,omitempty"`

	// The individual or entity who initiated the request and has responsibility for its activation
	Requester *common.Reference `json:"requester,omitempty"`

	// This element is labeled as a modifier because the status contains the codes revoked and entered-in-error
	Status *DeviceRequestStatus `json:"status,omitempty"`

	// The patient who will use the device
	Subject common.Reference `json:"subject"`

	// Additional clinical information about the patient that may influence the request fulfilment
	SupportingInfo []common.Reference `json:"supportingInfo,omitempty"`
}

// DeviceRequestIntent represents the intent of the device request
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

// DeviceRequestPriority represents the priority of the device request
type DeviceRequestPriority string

const (
	DeviceRequestPriorityRoutine DeviceRequestPriority = "routine"
	DeviceRequestPriorityUrgent  DeviceRequestPriority = "urgent"
	DeviceRequestPriorityASAP    DeviceRequestPriority = "asap"
	DeviceRequestPriorityStat    DeviceRequestPriority = "stat"
)

// DeviceRequestStatus represents the status of the device request
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
