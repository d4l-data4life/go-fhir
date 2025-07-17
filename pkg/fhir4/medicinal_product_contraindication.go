package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// MedicinalProductContraindication represents the clinical particulars - indications, contraindications etc. of a medicinal product
type MedicinalProductContraindication struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "MedicinalProductContraindication"

	// A comorbidity (concurrent condition) or coinfection
	Comorbidity []common.CodeableConcept `json:"comorbidity,omitempty"`

	// The disease, symptom or procedure for the contraindication
	Disease *common.CodeableConcept `json:"disease,omitempty"`

	// The status of the disease or symptom for the contraindication
	DiseaseStatus *common.CodeableConcept `json:"diseaseStatus,omitempty"`

	// Information about the use of the medicinal product in relation to other therapies described as part of the indication
	OtherTherapy []MedicinalProductContraindicationOtherTherapy `json:"otherTherapy,omitempty"`

	// The population group to which this applies
	Population []MedicinalProductContraindicationPopulation `json:"population,omitempty"`

	// The medication for which this is an indication
	Subject []common.Reference `json:"subject,omitempty"`

	// Information about the use of the medicinal product in relation to other therapies as part of the indication
	TherapeuticIndication []common.Reference `json:"therapeuticIndication,omitempty"`
}

// MedicinalProductContraindicationOtherTherapy represents information about the use of the medicinal product in relation to other therapies
type MedicinalProductContraindicationOtherTherapy struct {
	common.BackboneElement

	// Reference to a specific medication (active substance, medicinal product or class of products) as part of an indication or contraindication
	MedicationCodeableConcept *common.CodeableConcept `json:"medicationCodeableConcept,omitempty"`

	// Reference to a specific medication (active substance, medicinal product or class of products) as part of an indication or contraindication
	MedicationReference *common.Reference `json:"medicationReference,omitempty"`

	// The type of relationship between the medicinal product indication or contraindication and another therapy
	TherapyRelationshipType common.CodeableConcept `json:"therapyRelationshipType"`
}

// MedicinalProductContraindicationPopulation represents a population of people with some set of grouping criteria
type MedicinalProductContraindicationPopulation struct {
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
