package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// MedicinalProductUndesirableEffect represents describe the undesirable effects of the medicinal product
type MedicinalProductUndesirableEffect struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "MedicinalProductUndesirableEffect"

	// Classification of the effect
	Classification *common.CodeableConcept `json:"classification,omitempty"`

	// The frequency of occurrence of the effect
	FrequencyOfOccurrence *common.CodeableConcept `json:"frequencyOfOccurrence,omitempty"`

	// The population group to which this applies
	Population []MedicinalProductUndesirableEffectPopulation `json:"population,omitempty"`

	// The medication for which this is an indication
	Subject []common.Reference `json:"subject,omitempty"`

	// The symptom, condition or undesirable effect
	SymptomConditionEffect *common.CodeableConcept `json:"symptomConditionEffect,omitempty"`
}

// MedicinalProductUndesirableEffectPopulation represents a population of people with some set of grouping criteria
type MedicinalProductUndesirableEffectPopulation struct {
	common.BackboneElement

	// The age of the specific population
	AgeRange *common.Range `json:"ageRange,omitempty"`

	// The age of the specific population
	AgeCodeableConcept *common.CodeableConcept `json:"ageCodeableConcept,omitempty"`

	// The gender of the specific population
	Gender *common.CodeableConcept `json:"gender,omitempty"`

	// The existing physiological conditions of the specific population to which this applies
	PhysiologicalCondition *common.CodeableConcept `json:"physiologicalCondition,omitempty"`

	// Race of the specific population
	Race *common.CodeableConcept `json:"race,omitempty"`
}
