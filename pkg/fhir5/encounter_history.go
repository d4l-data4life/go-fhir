// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// EncounterHistoryStatus represents the status of an encounter history (reuses EncounterStatus)
type EncounterHistoryStatus = EncounterStatus

// EncounterHistoryLocation represents location information in encounter history
type EncounterHistoryLocation struct {
	common.BackboneElement

	// Physical form of the location
	Form *common.CodeableConcept `json:"form,omitempty"`

	// The location where the encounter takes place
	Location common.Reference `json:"location"`
}

// EncounterHistory represents a record of significant events/milestones during an Encounter
type EncounterHistory struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "EncounterHistory"

	// The encounter whose history is being described
	Encounter common.Reference `json:"encounter"`

	// The status of the encounter at the period
	Status EncounterHistoryStatus `json:"status"`

	// The period covered by this history
	Period common.Period `json:"period"`

	// The locations involved in this period
	Location []EncounterHistoryLocation `json:"location,omitempty"`
}
