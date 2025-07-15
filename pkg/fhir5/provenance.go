// Package fhir5 contains FHIR R5 resource definitions
package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// ProvenanceEntity represents an entity used in a provenance activity
type ProvenanceEntity struct {
	common.BackboneElement

	// How the entity was used during the activity
	Role string `json:"role"`

	// Identity of the entity used
	What common.Reference `json:"what"`

	// Entity is attributed to this agent
	Agent []interface{} `json:"agent,omitempty"`
}

// Provenance represents provenance of a resource
type Provenance struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Provenance"

	// An activity is something that occurs over a period of time and acts upon or with entities
	Activity *common.CodeableConcept `json:"activity,omitempty"`

	// Several agents may be associated with an activity and vice-versa
	Agent []interface{} `json:"agent"`

	// The authorization that was used during the event being recorded
	Authorization []CodeableReference `json:"authorization,omitempty"`

	// Allows tracing of authorization for the events and tracking whether proposals/recommendations were acted upon
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// This will typically be the encounter the event occurred
	Encounter *common.Reference `json:"encounter,omitempty"`

	// An entity used in this activity
	Entity []ProvenanceEntity `json:"entity,omitempty"`

	// Where the activity occurred, if relevant
	Location *common.Reference `json:"location,omitempty"`

	// The period can be a little arbitrary; where possible, the time should correspond to human assessment of the activity time
	OccurredPeriod   *common.Period `json:"occurredPeriod,omitempty"`
	OccurredDateTime *string        `json:"occurredDateTime,omitempty"`

	// The patient element is available to enable deterministic tracking of activities that involve the patient as the subject
	Patient *common.Reference `json:"patient,omitempty"`

	// For example: Where an OAuth token authorizes, the unique identifier from the OAuth token is placed into the policy element
	Policy []string `json:"policy,omitempty"`

	// This can be a little different from the time stamp on the resource if there is a delay between recording the event
	Recorded *string `json:"recorded,omitempty"`

	// A digital signature on the target Reference(s)
	Signature []Signature `json:"signature,omitempty"`

	// Target references are usually version specific, but might not be
	Target []common.Reference `json:"target"`
}
