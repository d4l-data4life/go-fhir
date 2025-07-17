// Package fhir5 contains FHIR R5 resource definitions
package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// Provenance represents who, what, and when for a set of resources (FHIR R5)
type Provenance struct {
	DomainResource

	ResourceType     string                  `json:"resourceType"` // Always "Provenance"
	Target           []common.Reference      `json:"target"`
	OccurredPeriod   *common.Period          `json:"occurredPeriod,omitempty"`
	OccurredDateTime *string                 `json:"occurredDateTime,omitempty"`
	Recorded         string                  `json:"recorded"`
	Policy           []string                `json:"policy,omitempty"`
	Location         *common.Reference       `json:"location,omitempty"`
	Authorization    []CodeableReference     `json:"authorization,omitempty"`
	Activity         *common.CodeableConcept `json:"activity,omitempty"`
	BasedOn          []common.Reference      `json:"basedOn,omitempty"`
	Patient          *common.Reference       `json:"patient,omitempty"`
	Encounter        *common.Reference       `json:"encounter,omitempty"`
	Agent            []ProvenanceAgent       `json:"agent"`
	Entity           []ProvenanceEntity      `json:"entity,omitempty"`
	Signature        []Signature             `json:"signature,omitempty"`
}

type ProvenanceAgent struct {
	common.BackboneElement
	Type       *common.CodeableConcept  `json:"type,omitempty"`
	Role       []common.CodeableConcept `json:"role,omitempty"`
	Who        common.Reference         `json:"who"`
	OnBehalfOf *common.Reference        `json:"onBehalfOf,omitempty"`
}

type ProvenanceEntity struct {
	common.BackboneElement
	Role  string            `json:"role"`
	What  common.Reference  `json:"what"`
	Agent []ProvenanceAgent `json:"agent,omitempty"`
}
