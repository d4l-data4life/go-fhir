package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// Provenance represents a record of activity performed on a resource
type Provenance struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Provenance"

	// An activity is something that occurs over a period of time and acts upon or with entities
	Activity *common.CodeableConcept `json:"activity,omitempty"`

	// Several agents may be associated (i.e. has some responsibility for an activity) with an activity and vice-versa
	Agent []ProvenanceAgent `json:"agent"`

	// An entity used in this activity
	Entity []ProvenanceEntity `json:"entity,omitempty"`

	// Where the activity occurred, if relevant
	Location *common.Reference `json:"location,omitempty"`

	// The period can be a little arbitrary; where possible, the time should correspond to human assessment of the activity time
	OccurredPeriod *common.Period `json:"occurredPeriod,omitempty"`

	// The period can be a little arbitrary; where possible, the time should correspond to human assessment of the activity time
	OccurredDateTime *string `json:"occurredDateTime,omitempty"`

	// For example: Where an OAuth token authorizes, the unique identifier from the OAuth token is placed into the policy element
	Policy []string `json:"policy,omitempty"`

	// The reason that the activity was taking place
	Reason []common.CodeableConcept `json:"reason,omitempty"`

	// This can be a little different from the time stamp on the resource if there is a delay between recording the event and updating the provenance and target resource
	Recorded string `json:"recorded"`

	// A digital signature on the target Reference(s)
	Signature []common.Signature `json:"signature,omitempty"`

	// Target references are usually version specific, but might not be, if a version has not been assigned
	Target []common.Reference `json:"target"`
}

// ProvenanceAgent represents several agents may be associated with an activity
type ProvenanceAgent struct {
	common.BackboneElement

	// The function of the agent with respect to the activity
	Function *common.CodeableConcept `json:"function,omitempty"`

	// The individual, device or organization that participated in the event
	Who common.Reference `json:"who"`

	// The individual, device, or organization for whom the change was made
	OnBehalfOf *common.Reference `json:"onBehalfOf,omitempty"`
}

// ProvenanceEntity represents an entity used in this activity
type ProvenanceEntity struct {
	common.BackboneElement

	// How the entity was used during the activity
	Role ProvenanceEntityRole `json:"role"`

	// Identity of the Entity used. May be a logical or physical uri and maybe absolute or relative
	What common.Reference `json:"what"`

	// Identity of the Entity used. May be a logical or physical uri and maybe absolute or relative
	Agent []ProvenanceAgent `json:"agent,omitempty"`
}

// ProvenanceEntityRole represents how the entity was used during the activity
type ProvenanceEntityRole string

const (
	ProvenanceEntityRoleDerivation ProvenanceEntityRole = "derivation"
	ProvenanceEntityRoleRevision   ProvenanceEntityRole = "revision"
	ProvenanceEntityRoleQuotation  ProvenanceEntityRole = "quotation"
	ProvenanceEntityRoleSource     ProvenanceEntityRole = "source"
	ProvenanceEntityRoleRemoval    ProvenanceEntityRole = "removal"
)
