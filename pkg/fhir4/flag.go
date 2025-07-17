package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// Flag represents prospective warnings of potential issues when providing care to the patient
type Flag struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Flag"

	// The person, organization or device that created the flag
	Author *common.Reference `json:"author,omitempty"`

	// The value set will often need to be adjusted based on local business rules and usage context
	Category []common.CodeableConcept `json:"category,omitempty"`

	// If non-coded, use CodeableConcept.text. This element should always be included in the narrative
	Code common.CodeableConcept `json:"code"`

	// If both Flag.encounter and Flag.period are valued, then Flag.period.start shall not be before Encounter.period.start
	Encounter *common.Reference `json:"encounter,omitempty"`

	// This is a business identifier, not a resource identifier
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The period of time from the activation of the flag to inactivation of the flag
	Period *common.Period `json:"period,omitempty"`

	// This element is labeled as a modifier because the status contains codes that mark the resource as not currently valid
	Status FlagStatus `json:"status"`

	// The patient, location, group, organization, or practitioner etc. this is about record this flag is associated with
	Subject common.Reference `json:"subject"`
}

// FlagStatus represents the status of the flag
type FlagStatus string

const (
	FlagStatusActive         FlagStatus = "active"
	FlagStatusInactive       FlagStatus = "inactive"
	FlagStatusEnteredInError FlagStatus = "entered-in-error"
)
