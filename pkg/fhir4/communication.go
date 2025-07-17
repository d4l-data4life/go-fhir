// Package fhir4 contains FHIR R4 (version 4.0.1) resource definitions
package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// Communication represents an instance of information being transmitted
type Communication struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Communication"

	// Unique identifier
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Instantiates FHIR protocol or definition
	InstantiatesCanonical []string `json:"instantiatesCanonical,omitempty"`

	// Instantiates external protocol or definition
	InstantiatesUri []string `json:"instantiatesUri,omitempty"`

	// Request fulfilled by this communication
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// Part of this action
	PartOf []common.Reference `json:"partOf,omitempty"`

	// Reply to
	InResponseTo []common.Reference `json:"inResponseTo,omitempty"`

	// preparation | in-progress | not-done | on-hold | stopped | completed | entered-in-error | unknown
	Status CommunicationStatus `json:"status"`

	// Reason for current status
	StatusReason *common.CodeableConcept `json:"statusReason,omitempty"`

	// Message category
	Category []common.CodeableConcept `json:"category,omitempty"`

	// routine | urgent | asap | stat
	Priority *CommunicationPriority `json:"priority,omitempty"`

	// A channel of communication
	Medium []common.CodeableConcept `json:"medium,omitempty"`

	// Focus of message
	Subject *common.Reference `json:"subject,omitempty"`

	// Description of the purpose/content
	Topic *common.CodeableConcept `json:"topic,omitempty"`

	// Resources that pertain to this communication
	About []common.Reference `json:"about,omitempty"`

	// Encounter created as part of
	Encounter *common.Reference `json:"encounter,omitempty"`

	// When sent
	Sent *string `json:"sent,omitempty"`

	// When received
	Received *string `json:"received,omitempty"`

	// Message recipient
	Recipient []common.Reference `json:"recipient,omitempty"`

	// Message sender
	Sender *common.Reference `json:"sender,omitempty"`

	// Indication for message
	ReasonCode []common.CodeableConcept `json:"reasonCode,omitempty"`

	// Why was communication done?
	ReasonReference []common.Reference `json:"reasonReference,omitempty"`

	// Message payload
	Payload []CommunicationPayload `json:"payload,omitempty"`

	// Comments made about the communication
	Note []Annotation `json:"note,omitempty"`
}

// CommunicationStatus represents the status of the communication
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

// CommunicationPriority represents the priority of the communication
type CommunicationPriority string

const (
	CommunicationPriorityRoutine CommunicationPriority = "routine"
	CommunicationPriorityUrgent  CommunicationPriority = "urgent"
	CommunicationPriorityASAP    CommunicationPriority = "asap"
	CommunicationPriorityStat    CommunicationPriority = "stat"
)

// CommunicationPayload represents a message part content
type CommunicationPayload struct {
	common.BackboneElement

	// Message part content
	ContentString     *string           `json:"contentString,omitempty"`
	ContentAttachment *Attachment       `json:"contentAttachment,omitempty"`
	ContentReference  *common.Reference `json:"contentReference,omitempty"`
}
