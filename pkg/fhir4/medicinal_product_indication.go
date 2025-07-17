package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// MedicinalProductIndication represents indication for the Medicinal Product
type MedicinalProductIndication struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "MedicinalProductIndication"

	// Comorbidity (concurrent condition) or co-infection as part of the indication
	Comorbidity []common.CodeableConcept `json:"comorbidity,omitempty"`

	// The status of the disease or symptom for which the indication applies
	DiseaseStatus *common.CodeableConcept `json:"diseaseStatus,omitempty"`

	// The disease, symptom or procedure that is the indication for treatment
	DiseaseSymptomProcedure *common.CodeableConcept `json:"diseaseSymptomProcedure,omitempty"`

	// Timing or duration information as part of the indication
	Duration *common.Quantity `json:"duration,omitempty"`

	// The intended effect, aim or strategy to be achieved by the indication
	IntendedEffect *common.CodeableConcept `json:"intendedEffect,omitempty"`

	// Information about the use of the medicinal product in relation to other therapies described as part of the indication
	OtherTherapy []MedicinalProductIndicationOtherTherapy `json:"otherTherapy,omitempty"`

	// The population group to which this applies
	Population []MedicinalProductIndicationPopulation `json:"population,omitempty"`

	// The medication for which this is an indication
	Subject []common.Reference `json:"subject,omitempty"`

	// Describe the undesirable effects of the medicinal product
	UndesirableEffect []common.Reference `json:"undesirableEffect,omitempty"`
}

// MedicinalProductIndicationOtherTherapy represents information about the use of the medicinal product in relation to other therapies
type MedicinalProductIndicationOtherTherapy struct {
	common.BackboneElement

	// Reference to a specific medication (active substance, medicinal product or class of products) as part of an indication or contraindication
	MedicationCodeableConcept *common.CodeableConcept `json:"medicationCodeableConcept,omitempty"`

	// Reference to a specific medication (active substance, medicinal product or class of products) as part of an indication or contraindication
	MedicationReference *common.Reference `json:"medicationReference,omitempty"`

	// The type of relationship between the medicinal product indication or contraindication and another therapy
	TherapyRelationshipType common.CodeableConcept `json:"therapyRelationshipType"`
}

// MedicinalProductIndicationPopulation represents a population of people with some set of grouping criteria
type MedicinalProductIndicationPopulation struct {
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
