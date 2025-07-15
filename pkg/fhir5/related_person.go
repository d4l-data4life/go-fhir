// Package fhir5 contains FHIR R5 resource definitions
package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// RelatedPerson represents information about a person that is involved in the care for a patient
type RelatedPerson struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "RelatedPerson"

	// Whether this related person record is in active use
	Active *bool `json:"active,omitempty"`

	// Address where the related person can be contacted or visited
	Address []Address `json:"address,omitempty"`

	// The date on which the related person was born
	BirthDate *string `json:"birthDate,omitempty"`

	// Administrative Gender - the gender that the person is considered to have for administration and record keeping purposes
	Gender *string `json:"gender,omitempty"`

	// Identifier for a person within a particular scope
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The name of the related person
	Name []HumanName `json:"name,omitempty"`

	// The patient this person is related to
	Patient common.Reference `json:"patient"`

	// The nature of the relationship between a patient and the related person
	Relationship []common.CodeableConcept `json:"relationship,omitempty"`

	// A contact detail for the person
	Telecom []ContactPoint `json:"telecom,omitempty"`

	// Image of the person
	Photo []Attachment `json:"photo,omitempty"`

	// Period of time that this relationship is considered valid
	Period *common.Period `json:"period,omitempty"`

	// A language which may be used to communicate with the related person about the patient's health
	Communication []interface{} `json:"communication,omitempty"`
}
