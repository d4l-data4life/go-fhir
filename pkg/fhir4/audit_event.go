package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// AuditEventAction represents the type of action performed during the event
type AuditEventAction string

const (
	AuditEventActionC AuditEventAction = "C" // Create
	AuditEventActionR AuditEventAction = "R" // Read
	AuditEventActionU AuditEventAction = "U" // Update
	AuditEventActionD AuditEventAction = "D" // Delete
	AuditEventActionE AuditEventAction = "E" // Execute
)

// AuditEventOutcome represents the outcome of the event
type AuditEventOutcome string

const (
	AuditEventOutcome0  AuditEventOutcome = "0"  // Success
	AuditEventOutcome4  AuditEventOutcome = "4"  // Minor failure
	AuditEventOutcome8  AuditEventOutcome = "8"  // Serious failure
	AuditEventOutcome12 AuditEventOutcome = "12" // Major failure
)

// AuditEventAgentNetworkType represents the type of network access point
type AuditEventAgentNetworkType string

const (
	AuditEventAgentNetworkType1 AuditEventAgentNetworkType = "1" // Machine Name
	AuditEventAgentNetworkType2 AuditEventAgentNetworkType = "2" // IP Address
	AuditEventAgentNetworkType3 AuditEventAgentNetworkType = "3" // Telephone Number
	AuditEventAgentNetworkType4 AuditEventAgentNetworkType = "4" // Email address
	AuditEventAgentNetworkType5 AuditEventAgentNetworkType = "5" // URI
)

// AuditEventAgent represents an agent involved in the event
type AuditEventAgent struct {
	common.BackboneElement

	// Specification of the participation type the user plays when performing the event
	Type *common.CodeableConcept `json:"type,omitempty"`

	// The security role that the user was acting under, that come from local codes defined by the access control security system (e.g. RBAC, ABAC) used in the local context
	Role []common.CodeableConcept `json:"role,omitempty"`

	// Direct reference to a resource that identifies the agent
	Who *common.Reference `json:"who,omitempty"`

	// Alternative agent Identifier. For a human, this should be the user ID
	AltID *string `json:"altId,omitempty"`

	// Human-meaningful name for the agent
	Name *string `json:"name,omitempty"`

	// Indicator that the user is or is not the requestor, or initiator, for the event being audited
	Requestor bool `json:"requestor"`

	// Where the event occurred
	Location *common.Reference `json:"location,omitempty"`

	// The policy or plan that authorized the activity being recorded
	Policy []string `json:"policy,omitempty"`

	// Type of media involved
	Media *common.Coding `json:"media,omitempty"`

	// Logical network location for application activity, if the activity has a network location
	Network *AuditEventAgentNetwork `json:"network,omitempty"`

	// The reason (purpose of use) that was used during the event being recorded
	PurposeOfUse []common.CodeableConcept `json:"purposeOfUse,omitempty"`
}

// AuditEventAgentNetwork represents logical network location for application activity
type AuditEventAgentNetwork struct {
	common.BackboneElement

	// An identifier for the network access point of the user device for the audit event
	Address *string `json:"address,omitempty"`

	// An identifier for the type of network access point that originated the audit event
	Type *AuditEventAgentNetworkType `json:"type,omitempty"`
}

// AuditEventSource represents the source of the event
type AuditEventSource struct {
	common.BackboneElement

	// Logical source location within the healthcare enterprise network
	Site *string `json:"site,omitempty"`

	// Identifier of the source where the event originated
	Observer common.Reference `json:"observer"`

	// Code specifying the type of source where event originated
	Type []common.Coding `json:"type,omitempty"`
}

// AuditEventEntity represents an entity involved in the event
type AuditEventEntity struct {
	common.BackboneElement

	// Identifies a specific instance of the entity
	What *common.Reference `json:"what,omitempty"`

	// Code representing the role the entity played in the event being audited
	Type *common.Coding `json:"type,omitempty"`

	// Code representing the role the entity played in the event being audited
	Role *common.Coding `json:"role,omitempty"`

	// Security labels on the entity
	SecurityLabel []common.Coding `json:"securityLabel,omitempty"`

	// A value that describes the entity
	Name *string `json:"name,omitempty"`

	// Text that describes the entity in more detail
	Description *string `json:"description,omitempty"`

	// The query parameters for a query-type entities
	Query *string `json:"query,omitempty"`

	// Tagged value pairs for conveying additional information about the entity
	Detail []AuditEventEntityDetail `json:"detail,omitempty"`
}

// AuditEventEntityDetail represents additional information about the entity
type AuditEventEntityDetail struct {
	common.BackboneElement

	// The type of extra detail provided in the value
	Type string `json:"type"`

	// The value of the extra detail
	ValueString       *string `json:"valueString,omitempty"`
	ValueBase64Binary *string `json:"valueBase64Binary,omitempty"`
}

// AuditEvent represents a record of an event made for purposes of maintaining a security log
type AuditEvent struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "AuditEvent"

	// Identifier for a family of the event
	Type common.Coding `json:"type"`

	// Identifier for the category of event
	Subtype []common.Coding `json:"subtype,omitempty"`

	// Indicator for type of action performed during the event that generated the audit
	Action *AuditEventAction `json:"action,omitempty"`

	// The period during which the recorded activity occurred
	Period *common.Period `json:"period,omitempty"`

	// The time when the event was recorded
	Recorded *string `json:"recorded"`

	// Indicates whether the event succeeded or failed
	Outcome *AuditEventOutcome `json:"outcome,omitempty"`

	// A free text description of the outcome of the event
	OutcomeDesc *string `json:"outcomeDesc,omitempty"`

	// The purposeOfUse (reason) that was used during the event being recorded
	PurposeOfEvent []common.CodeableConcept `json:"purposeOfEvent,omitempty"`

	// An actor taking an active role in the event or activity that is logged
	Agent []AuditEventAgent `json:"agent"`

	// The system that is reporting the event
	Source AuditEventSource `json:"source"`

	// Specific instances of data or objects that have been accessed
	Entity []AuditEventEntity `json:"entity,omitempty"`
}
