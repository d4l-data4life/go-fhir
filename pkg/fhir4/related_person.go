package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// RelatedPerson represents information about a person that is involved in the care for a patient
type RelatedPerson struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "RelatedPerson"

	// This element is labeled as a modifier because it may be used to mark that the resource was created in error
	Active *bool `json:"active,omitempty"`

	// Address where the related person can be contacted or visited
	Address []common.Address `json:"address,omitempty"`

	// The date on which the related person was born
	BirthDate *string `json:"birthDate,omitempty"`

	// If no language is specified, this *implies* that the default local language is spoken
	Communication []RelatedPersonCommunication `json:"communication,omitempty"`

	// Administrative Gender - the gender that the person is considered to have for administration and record keeping purposes
	Gender *AdministrativeGender `json:"gender,omitempty"`

	// Identifier for a person within a particular scope
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// A name associated with the person
	Name []common.HumanName `json:"name,omitempty"`

	// The patient this person is related to
	Patient common.Reference `json:"patient"`

	// The period of time during which this relationship is or was active
	Period *common.Period `json:"period,omitempty"`

	// Image of the person
	Photo []common.Attachment `json:"photo,omitempty"`

	// The nature of the relationship between a patient and the related person
	Relationship []common.CodeableConcept `json:"relationship,omitempty"`

	// Person may have multiple ways to be contacted with different uses or applicable periods
	Telecom []common.ContactPoint `json:"telecom,omitempty"`
}

// RelatedPersonCommunication represents if no language is specified, this *implies* that the default local language is spoken
type RelatedPersonCommunication struct {
	common.BackboneElement

	// The structure aa-BB with this exact casing is one the most widely used notations for locale
	Language common.CodeableConcept `json:"language"`

	// This language is specifically identified for communicating healthcare information
	Preferred *bool `json:"preferred,omitempty"`
}
