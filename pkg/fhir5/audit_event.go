// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// AuditEventOutcome represents the outcome of an audit event
type AuditEventOutcome struct {
	common.BackboneElement

	// The outcome code
	Code common.Coding `json:"code"`

	// A human readable description of the error issue
	Detail []common.CodeableConcept `json:"detail,omitempty"`
}

// AuditEventAgent represents an agent involved in an audit event
type AuditEventAgent struct {
	common.BackboneElement

	// Use AuditEvent.agent.authorization when you know that is specific to the agent
	Authorization []common.CodeableConcept `json:"authorization,omitempty"`

	// Where the agent location is known, the agent location when the event occurred
	Location *common.Reference `json:"location,omitempty"`

	// When remote network endpoint is known, another agent representing the remote agent
	NetworkReference *common.Reference `json:"networkReference,omitempty"`

	// When remote network endpoint is known, another agent representing the remote agent
	NetworkUri *string `json:"networkUri,omitempty"`

	// When remote network endpoint is known, another agent representing the remote agent
	NetworkString *string `json:"networkString,omitempty"`

	// For example: Where an OAuth token authorizes, the unique identifier from the OAuth token
	Policy []string `json:"policy,omitempty"`

	// There can only be one initiator. If the initiator is not clear, then do not choose any one agent as the initiator
	Requestor *bool `json:"requestor,omitempty"`

	// For example: Chief-of-Radiology, Nurse, Physician, Medical-Student, etc.
	Role []common.CodeableConcept `json:"role,omitempty"`

	// For example: assembler, author, prescriber, signer, investigator, etc.
	Type *common.CodeableConcept `json:"type,omitempty"`

	// Where a User ID is available it will go into who.identifier
	Who common.Reference `json:"who"`
}

// AuditEventSource represents the source of an audit event
type AuditEventSource struct {
	common.BackboneElement

	// Identifier of the source where the event was detected
	Observer common.Reference `json:"observer"`

	// Logical source location within the healthcare enterprise network
	Site *common.Reference `json:"site,omitempty"`

	// Code specifying the type of source where event originated
	Type []common.CodeableConcept `json:"type,omitempty"`
}

// AuditEventEntity represents an entity involved in an audit event
type AuditEventEntity struct {
	common.BackboneElement

	// A usecase where one AuditEvent.agent.participant is a human user and the other is a process on behalf of that human user
	Agent []AuditEventAgent `json:"agent,omitempty"`

	// Tagged value pairs for conveying additional information about the entity
	Detail []AuditEventEntityDetail `json:"detail,omitempty"`

	// The meaning and processing of the data element/entity
	Description *string `json:"description,omitempty"`

	// The type of the object that was involved in the audit event
	Lifecycle *common.Coding `json:"lifecycle,omitempty"`

	// A name of the entity in the audit event
	Name *string `json:"name,omitempty"`

	// Text that describes the entity in more detail
	Query *string `json:"query,omitempty"`

	// The security labels on the entity
	SecurityLabel []common.Coding `json:"securityLabel,omitempty"`

	// The type of the object that was involved in the audit event
	Type *common.Coding `json:"type,omitempty"`

	// What the entity is
	What common.Reference `json:"what"`
}

// AuditEventEntityDetail represents details about an entity in an audit event
type AuditEventEntityDetail struct {
	common.BackboneElement

	// The type of extra detail provided in the value
	Type string `json:"type"`

	// The value of the extra detail
	ValueString       *string `json:"valueString,omitempty"`
	ValueBase64Binary *string `json:"valueBase64Binary,omitempty"`
}

// AuditEvent represents an audit event
type AuditEvent struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "AuditEvent"

	// Indicator for type of action performed during the event that generated the audit
	Action *string `json:"action,omitempty"`

	// Several agents may be associated with an event or activity
	Agent []AuditEventAgent `json:"agent"`

	// Use AuditEvent.agent.authorization when you know that it is specific to the agent
	Authorization []common.CodeableConcept `json:"authorization,omitempty"`

	// Allows tracing of authorization for the events and tracking whether proposals/recommendations were acted upon
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// Classification of the type of event
	Category []common.CodeableConcept `json:"category,omitempty"`

	// Describes what happened. The most specific code for the event
	Code common.CodeableConcept `json:"code"`

	// This will typically be the encounter the event occurred
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Required unless the values for event identification, agent identification, and audit source identification are sufficient
	Entity []AuditEventEntity `json:"entity,omitempty"`

	// The time or period can be a little arbitrary
	OccurredPeriod *common.Period `json:"occurredPeriod,omitempty"`

	// The time or period can be a little arbitrary
	OccurredDateTime *string `json:"occurredDateTime,omitempty"`

	// In some cases a "success" may be partial
	Outcome *AuditEventOutcome `json:"outcome,omitempty"`

	// The patient element is available to enable deterministic tracking of activities
	Patient *common.Reference `json:"patient,omitempty"`

	// In a distributed system, some sort of common time base is a good implementation tactic
	Recorded string `json:"recorded"`

	// ATNA will map this to the SYSLOG PRI element
	Severity *AuditEventSeverity `json:"severity,omitempty"`

	// Events are reported by the actor that detected them
	Source AuditEventSource `json:"source"`
}

// AuditEventSeverity represents the severity of an audit event
type AuditEventSeverity string

const (
	AuditEventSeverityEmergency     AuditEventSeverity = "emergency"
	AuditEventSeverityAlert         AuditEventSeverity = "alert"
	AuditEventSeverityCritical      AuditEventSeverity = "critical"
	AuditEventSeverityError         AuditEventSeverity = "error"
	AuditEventSeverityWarning       AuditEventSeverity = "warning"
	AuditEventSeverityNotice        AuditEventSeverity = "notice"
	AuditEventSeverityInformational AuditEventSeverity = "informational"
	AuditEventSeverityDebug         AuditEventSeverity = "debug"
)
