// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// CommunicationPayload represents text, attachment(s), or resource(s) that was communicated
type CommunicationPayload struct {
	common.BackboneElement

	// The content can be codified or textual
	ContentAttachment *Attachment `json:"contentAttachment,omitempty"`

	// The content can be codified or textual
	ContentReference *common.Reference `json:"contentReference,omitempty"`

	// The content can be codified or textual
	ContentCodeableConcept *common.CodeableConcept `json:"contentCodeableConcept,omitempty"`
}

// Communication represents a clinical or business level record of information being transmitted
type Communication struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Communication"

	// Don't use Communication.about element when a more specific element exists
	About []common.Reference `json:"about,omitempty"`

	// This must point to some sort of a 'Request' resource
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// There may be multiple axes of categorization and one communication may serve multiple purposes
	Category []common.CodeableConcept `json:"category,omitempty"`

	// This will typically be the encounter the event occurred within
	Encounter *common.Reference `json:"encounter,omitempty"`

	// This is a business identifier, not a resource identifier
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Prior communication that this communication is in response to
	InResponseTo []common.Reference `json:"inResponseTo,omitempty"`

	// The URL pointing to a FHIR-defined protocol, guideline, orderset or other definition
	InstantiatesCanonical []string `json:"instantiatesCanonical,omitempty"`

	// This might be an HTML page, PDF, etc. or could just be a non-resolvable URI identifier
	InstantiatesUri []string `json:"instantiatesUri,omitempty"`

	// A channel that was used for this communication (e.g. email, fax)
	Medium []common.CodeableConcept `json:"medium,omitempty"`

	// Additional notes or commentary about the communication by the sender, receiver or other interested parties
	Note []Annotation `json:"note,omitempty"`

	// The event the communication was part of
	PartOf []common.Reference `json:"partOf,omitempty"`

	// Text, attachment(s), or resource(s) that was communicated to the recipient
	Payload []CommunicationPayload `json:"payload,omitempty"`

	// Used to prioritize workflow when the communication is planned or in progress
	Priority *CommunicationPriority `json:"priority,omitempty"`

	// Textual reasons can be captured using reason.concept.text
	Reason []CodeableReference `json:"reason,omitempty"`

	// The time when this communication arrived at the destination
	Received *string `json:"received,omitempty"`

	// If receipts need to be tracked by an individual, a separate resource instance will need to be created
	Recipient []common.Reference `json:"recipient,omitempty"`

	// The entity which is the source of the communication
	Sender *common.Reference `json:"sender,omitempty"`

	// The time when this communication was sent
	Sent *string `json:"sent,omitempty"`

	// This element is labeled as a modifier because the status contains the codes aborted and entered-in-error
	Status CommunicationStatus `json:"status"`

	// This is generally only used for "exception" statuses such as "not-done", "suspended" or "aborted"
	StatusReason *common.CodeableConcept `json:"statusReason,omitempty"`

	// The patient or group that was the focus of this communication
	Subject *common.Reference `json:"subject,omitempty"`

	// Communication.topic.text can be used without any codings
	Topic *common.CodeableConcept `json:"topic,omitempty"`
}

// CommunicationPriority represents the priority of a communication
type CommunicationPriority string

const (
	CommunicationPriorityRoutine CommunicationPriority = "routine"
	CommunicationPriorityUrgent  CommunicationPriority = "urgent"
	CommunicationPriorityASAP    CommunicationPriority = "asap"
	CommunicationPriorityStat    CommunicationPriority = "stat"
)

// CommunicationStatus represents the status of a communication
type CommunicationStatus string

const (
	CommunicationStatusPreparation    CommunicationStatus = "preparation"
	CommunicationStatusInProgress     CommunicationStatus = "in-progress"
	CommunicationStatusNotDone        CommunicationStatus = "not-done"
	CommunicationStatusOnHold         CommunicationStatus = "on-hold"
	CommunicationStatusStopped        CommunicationStatus = "stopped"
	CommunicationStatusCompleted      CommunicationStatus = "completed"
	CommunicationStatusEnteredInError CommunicationStatus = "entered-in-error"
	CommunicationStatusUnknown        CommunicationStatus = "unknown"
)
