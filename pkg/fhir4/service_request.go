package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// ServiceRequest represents a record of a request for service such as diagnostic investigations, treatments, or operations to be performed
type ServiceRequest struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "ServiceRequest"

	// If a CodeableConcept is present, it indicates the pre-condition for performing the service
	AsNeededBoolean *bool `json:"asNeededBoolean,omitempty"`

	// If a CodeableConcept is present, it indicates the pre-condition for performing the service
	AsNeededCodeableConcept *common.CodeableConcept `json:"asNeededCodeableConcept,omitempty"`

	// When the request transitioned to being actionable
	AuthoredOn *string `json:"authoredOn,omitempty"`

	// Plan/proposal/order fulfilled by this request
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// Only used if not implicit in the code found in ServiceRequest.code
	BodySite []common.CodeableConcept `json:"bodySite,omitempty"`

	// There may be multiple axis of categorization depending on the context or use case for retrieving or displaying the resource
	Category []common.CodeableConcept `json:"category,omitempty"`

	// Many laboratory and radiology procedure codes embed the specimen/organ system in the test order name
	Code *common.CodeableConcept `json:"code,omitempty"`

	// In general, only the code and timeframe will be present, though occasional additional qualifiers such as body site or even performer could be included
	DoNotPerform *bool `json:"doNotPerform,omitempty"`

	// An encounter that provides additional information about the healthcare context in which this request is made
	Encounter *common.Reference `json:"encounter,omitempty"`

	// The identifier.type element is used to distinguish between the identifiers assigned by the orderer
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Note: This is a business identifier, not a resource identifier
	InstantiatesCanonical []string `json:"instantiatesCanonical,omitempty"`

	// This might be an HTML page, PDF, etc. or could just be a non-resolvable URI identifier
	InstantiatesUri []string `json:"instantiatesUri,omitempty"`

	// Insurance plans, coverage extensions, pre-authorizations and/or pre-determinations that may be needed for delivering the requested service
	Insurance []common.Reference `json:"insurance,omitempty"`

	// This element is labeled as a modifier because the intent alters when and how the resource is actually applicable
	Intent ServiceRequestIntent `json:"intent"`

	// The preferred location(s) where the procedure should actually happen in coded or free text form
	LocationCode []common.CodeableConcept `json:"locationCode,omitempty"`

	// A reference to the the preferred location(s) where the procedure should actually happen
	LocationReference []common.Reference `json:"locationReference,omitempty"`

	// Any other notes and comments made about the service request
	Note []common.Annotation `json:"note,omitempty"`

	// The date/time at which the requested service should occur
	OccurrenceDateTime *string `json:"occurrenceDateTime,omitempty"`

	// The date/time at which the requested service should occur
	OccurrencePeriod *common.Period `json:"occurrencePeriod,omitempty"`

	// The date/time at which the requested service should occur
	OccurrenceTiming *common.Timing `json:"occurrenceTiming,omitempty"`

	// For information from the medical record intended to support the delivery of the requested services
	OrderDetail []common.CodeableConcept `json:"orderDetail,omitempty"`

	// Instructions in terms that are understood by the patient or consumer
	PatientInstruction *string `json:"patientInstruction,omitempty"`

	// If multiple performers are present, it is interpreted as a list of *alternative* performers without any preference regardless of order
	Performer []common.Reference `json:"performer,omitempty"`

	// This is a role, not a participation type
	PerformerType *common.CodeableConcept `json:"performerType,omitempty"`

	// Indicates how quickly the ServiceRequest should be addressed with respect to other requests
	Priority *ServiceRequestPriority `json:"priority,omitempty"`

	// An amount of service being requested which can be a quantity
	QuantityQuantity *common.Quantity `json:"quantityQuantity,omitempty"`

	// An amount of service being requested which can be a ratio
	QuantityRatio *common.Ratio `json:"quantityRatio,omitempty"`

	// An amount of service being requested which can be a range
	QuantityRange *common.Range `json:"quantityRange,omitempty"`

	// This element represents why the referral is being made and may be used to decide how the service will be performed
	ReasonCode []common.CodeableConcept `json:"reasonCode,omitempty"`

	// This element represents why the referral is being made and may be used to decide how the service will be performed
	ReasonReference []common.Reference `json:"reasonReference,omitempty"`

	// This might not include provenances for all versions of the request – only those deemed "relevant" or important
	RelevantHistory []common.Reference `json:"relevantHistory,omitempty"`

	// The request takes the place of the referenced completed or terminated request(s)
	Replaces []common.Reference `json:"replaces,omitempty"`

	// This not the dispatcher, but rather who is the authorizer
	Requester *common.Reference `json:"requester,omitempty"`

	// Requests are linked either by a "basedOn" relationship or by having a common requisition
	Requisition *common.Identifier `json:"requisition,omitempty"`

	// Many diagnostic procedures need a specimen, but the request itself is not actually about the specimen
	Specimen []common.Reference `json:"specimen,omitempty"`

	// The status is generally fully in the control of the requester
	Status ServiceRequestStatus `json:"status"`

	// On whom or what the service is to be performed
	Subject common.Reference `json:"subject"`

	// To represent information about how the services are to be delivered use the `instructions` element
	SupportingInfo []common.Reference `json:"supportingInfo,omitempty"`
}

// ServiceRequestIntent represents this element is labeled as a modifier because the intent alters when and how the resource is actually applicable
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

// ServiceRequestStatus represents the status is generally fully in the control of the requester
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

// ServiceRequestPriority represents indicates how quickly the ServiceRequest should be addressed with respect to other requests
type ServiceRequestPriority string

const (
	ServiceRequestPriorityRoutine ServiceRequestPriority = "routine"
	ServiceRequestPriorityUrgent  ServiceRequestPriority = "urgent"
	ServiceRequestPriorityAsap    ServiceRequestPriority = "asap"
	ServiceRequestPriorityStat    ServiceRequestPriority = "stat"
)
