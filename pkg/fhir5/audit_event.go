// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

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

// AuditEventEntityDetail represents tagged value pairs for conveying additional information about the entity
type AuditEventEntityDetail struct {
	common.BackboneElement

	// The type of extra detail provided in the value
	Type string `json:"type"`

	// The value can be string when known to be a string, else base64 encoding should be used
	ValueString *string `json:"valueString,omitempty"`

	// The value can be base64 encoded for binary content
	ValueBase64Binary *string `json:"valueBase64Binary,omitempty"`
}

// AuditEventEntity represents an entity involved in the audit event
type AuditEventEntity struct {
	common.BackboneElement

	// A usecase where one AuditEvent.entity.agent is used where the Entity that was used in the creation/updating of a target resource
	Agent []AuditEventAgent `json:"agent,omitempty"`

	// Tagged value pairs for conveying additional information about the entity
	Detail []AuditEventEntityDetail `json:"detail,omitempty"`

	// The meaning and secondary-encoding of the content of base64 encoded blob
	Query *string `json:"query,omitempty"`

	// Text that describes the entity in more detail
	Description *string `json:"description,omitempty"`

	// This can be used to provide an audit trail for data, over time, as it passes through the system
	Lifecycle *common.Coding `json:"lifecycle,omitempty"`

	// This field may be used in a query/report to identify audit events for a specific person
	Name *string `json:"name,omitempty"`

	// Code representing the role the entity played in the event being audited
	Role *common.Coding `json:"role,omitempty"`

	// Copied from entity meta security tags
	SecurityLabel []common.Coding `json:"securityLabel,omitempty"`

	// This value is distinct from the user's role or any user relationship to the entity
	Type *common.Coding `json:"type,omitempty"`

	// Identifies a specific instance of the entity. The reference should be version specific
	What *common.Reference `json:"what,omitempty"`
}

// AuditEventAgent represents an agent involved in the audit event
type AuditEventAgent struct {
	common.BackboneElement

	// Alternative agent Identifier. For a human, this should be a user identifier text string from authentication system
	AltId *string `json:"altId,omitempty"`

	// Where the event occurred
	Location *common.Reference `json:"location,omitempty"`

	// Type of media involved. Used when the event is about exporting/importing onto media
	Media *common.Coding `json:"media,omitempty"`

	// Human-meaningful name for the agent
	Name *string `json:"name,omitempty"`

	// Logical network location for application activity, if the activity has a network location
	Network interface{} `json:"network,omitempty"`

	// For example: Where an OAuth token authorizes, the unique identifier from the OAuth token is placed into the policy element
	Policy []string `json:"policy,omitempty"`

	// Use AuditEvent.agent.purposeOfUse when you know that is specific to the agent
	PurposeOfUse []common.CodeableConcept `json:"purposeOfUse,omitempty"`

	// There can only be one initiator. If the initiator is not clear, then do not choose any one agent as the initiator
	Requestor bool `json:"requestor"`
}

// AuditEvent represents a record of an event relevant for purposes such as operations, privacy, security, maintenance, and performance analysis.
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

	// The time or period can be a little arbitrary; where possible, the time should correspond to human assessment of the activity time
	OccurredPeriod *common.Period `json:"occurredPeriod,omitempty"`

	// The time or period can be a little arbitrary; where possible, the time should correspond to human assessment of the activity time
	OccurredDateTime *string `json:"occurredDateTime,omitempty"`

	// In some cases a "success" may be partial
	Outcome *string `json:"outcome,omitempty"`

	// The patient element is available to enable deterministic tracking of activities that involve the patient as the subject
	Patient *common.Reference `json:"patient,omitempty"`

	// In a distributed system, some sort of common time base is a good implementation tactic
	Recorded string `json:"recorded"`

	// ATNA will map this to the SYSLOG PRI element
	Severity *string `json:"severity,omitempty"`

	// Events are reported by the actor that detected them
	Source AuditEventSource `json:"source"`
}
