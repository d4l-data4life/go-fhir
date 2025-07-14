// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// ServiceRequestStatus represents the status of a service request
type ServiceRequestStatus string

const (
	ServiceRequestStatusDraft          ServiceRequestStatus = "draft"
	ServiceRequestStatusActive         ServiceRequestStatus = "active"
	ServiceRequestStatusOnHold         ServiceRequestStatus = "on-hold"
	ServiceRequestStatusRevoked        ServiceRequestStatus = "revoked"
	ServiceRequestStatusCompleted      ServiceRequestStatus = "completed"
	ServiceRequestStatusEnteredInError ServiceRequestStatus = "entered-in-error"
	ServiceRequestStatusUnknown        ServiceRequestStatus = "unknown"
)

// ServiceRequestPriority represents the priority of a service request
type ServiceRequestPriority string

const (
	ServiceRequestPriorityRoutine ServiceRequestPriority = "routine"
	ServiceRequestPriorityUrgent  ServiceRequestPriority = "urgent"
	ServiceRequestPriorityASAP    ServiceRequestPriority = "asap"
	ServiceRequestPriorityStat    ServiceRequestPriority = "stat"
)

// ServiceRequestIntent represents the intent of a service request
type ServiceRequestIntent string

const (
	ServiceRequestIntentProposal      ServiceRequestIntent = "proposal"
	ServiceRequestIntentPlan          ServiceRequestIntent = "plan"
	ServiceRequestIntentDirective     ServiceRequestIntent = "directive"
	ServiceRequestIntentOrder         ServiceRequestIntent = "order"
	ServiceRequestIntentOriginalOrder ServiceRequestIntent = "original-order"
	ServiceRequestIntentReflexOrder   ServiceRequestIntent = "reflex-order"
	ServiceRequestIntentFillerOrder   ServiceRequestIntent = "filler-order"
	ServiceRequestIntentInstanceOrder ServiceRequestIntent = "instance-order"
	ServiceRequestIntentOption        ServiceRequestIntent = "option"
)

// ServiceRequest represents a record of a request for service such as diagnostic investigations, treatments, or operations
type ServiceRequest struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "ServiceRequest"

	// If a CodeableConcept is present, it indicates the pre-condition for performing the service
	AsNeededBoolean         *bool                   `json:"asNeededBoolean,omitempty"`
	AsNeededCodeableConcept *common.CodeableConcept `json:"asNeededCodeableConcept,omitempty"`

	// When the request transitioned to being actionable
	AuthoredOn *string `json:"authoredOn,omitempty"`

	// Plan/proposal/order fulfilled by this request
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// Anatomical location where the procedure should be performed
	BodySite []common.CodeableConcept `json:"bodySite,omitempty"`

	// Anatomic location where the procedure should be performed (R5 feature)
	BodyStructure *common.Reference `json:"bodyStructure,omitempty"`

	// A code that classifies the procedure for searching, sorting and display purposes
	Category []common.CodeableConcept `json:"category,omitempty"`

	// What is being requested/ordered
	Code *common.CodeableConcept `json:"code,omitempty"`

	// If do not perform is not specified, the request is a positive request
	DoNotPerform *bool `json:"doNotPerform,omitempty"`

	// An encounter that provides additional context in which this request is made
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Business identifier for this service request
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The URL pointing to a FHIR-defined protocol, guideline, order set or other definition
	InstantiatesCanonical []string `json:"instantiatesCanonical,omitempty"`

	// The URL pointing to an externally maintained protocol, guideline, order set
	InstantiatesUri []string `json:"instantiatesUri,omitempty"`

	// Additional clinical information
	Instructions []Instruction `json:"instructions,omitempty"`

	// Insurance plans, coverage extensions, pre-authorizations and/or pre-determinations
	Insurance []common.Reference `json:"insurance,omitempty"`

	// proposal | plan | directive | order | original-order | reflex-order | filler-order | instance-order | option
	Intent ServiceRequestIntent `json:"intent"`

	// Where procedure is going to be done
	Location []common.Reference `json:"location,omitempty"`

	// The subject on a service request is mandatory
	Subject common.Reference `json:"subject"`

	// Additional information that supports the ordering of the service
	Note []Annotation `json:"note,omitempty"`

	// When the request should be fulfilled
	OccurrenceDateTime *string        `json:"occurrenceDateTime,omitempty"`
	OccurrenceAge      *Age           `json:"occurrenceAge,omitempty"`
	OccurrencePeriod   *common.Period `json:"occurrencePeriod,omitempty"`
	OccurrenceRange    *Range         `json:"occurrenceRange,omitempty"`
	OccurrenceString   *string        `json:"occurrenceString,omitempty"`
	OccurrenceTiming   *Timing        `json:"occurrenceTiming,omitempty"`

	// The individual who initiated the request and has responsibility for its activation
	OrderDetail []common.CodeableConcept `json:"orderDetail,omitempty"`

	// The desired performer for doing the requested service
	Performer []common.Reference `json:"performer,omitempty"`

	// Performer type requirements
	PerformerType *common.CodeableConcept `json:"performerType,omitempty"`

	// routine | urgent | asap | stat
	Priority *ServiceRequestPriority `json:"priority,omitempty"`

	// An explanation or justification for why this service is being requested
	Reason []CodeableReference `json:"reason,omitempty"`

	// The request takes the place of referenced completed or terminated request(s)
	Replaces []common.Reference `json:"replaces,omitempty"`

	// The individual who initiated the request and has responsibility for its activation
	Requester *common.Reference `json:"requester,omitempty"`

	// draft | active | on-hold | revoked | completed | entered-in-error | unknown
	Status ServiceRequestStatus `json:"status"`

	// Reason for current status
	StatusReason *common.CodeableConcept `json:"statusReason,omitempty"`

	// Additional information that supports the ordering of the service
	SupportingInfo []common.Reference `json:"supportingInfo,omitempty"`
}
