// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// PatientGender represents the gender of a patient
type PatientGender string

const (
	PatientGenderMale    PatientGender = "male"
	PatientGenderFemale  PatientGender = "female"
	PatientGenderOther   PatientGender = "other"
	PatientGenderUnknown PatientGender = "unknown"
)

// PatientLinkType represents the type of link between patient resources
type PatientLinkType string

const (
	PatientLinkTypeReplacedBy PatientLinkType = "replaced-by"
	PatientLinkTypeReplaces   PatientLinkType = "replaces"
	PatientLinkTypeRefer      PatientLinkType = "refer"
	PatientLinkTypeSeeAlso    PatientLinkType = "seealso"
)

// PatientContact represents contact information for a patient
type PatientContact struct {
	common.BackboneElement

	// The nature of the relationship between the patient and the contact person
	Relationship []common.CodeableConcept `json:"relationship,omitempty"`

	// A name associated with the contact person
	Name *HumanName `json:"name,omitempty"`

	// A contact detail for the person
	Telecom []ContactPoint `json:"telecom,omitempty"`

	// Address for the contact person
	Address *Address `json:"address,omitempty"`

	// Administrative Gender
	Gender *PatientGender `json:"gender,omitempty"`

	// Organization on behalf of which the contact is acting or for which the contact is working
	Organization *common.Reference `json:"organization,omitempty"`

	// The period during which this contact person or organization is valid to be contacted relating to this patient
	Period *common.Period `json:"period,omitempty"`
}

// PatientCommunication represents patient communication preferences
type PatientCommunication struct {
	common.BackboneElement

	// The language which can be used to communicate with the patient about his or her health
	Language common.CodeableConcept `json:"language"`

	// Indicates whether or not the patient prefers this language
	Preferred *bool `json:"preferred,omitempty"`
}

// PatientLink represents a link to another patient resource
type PatientLink struct {
	common.BackboneElement

	// The other patient resource that the link refers to
	Other common.Reference `json:"other"`

	// The type of link between this patient resource and another patient resource
	Type PatientLinkType `json:"type"`
}

// Patient represents a person receiving care or other health-related services
type Patient struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Patient"

	// A list of names associated with the patient
	Name []HumanName `json:"name,omitempty"`

	// A list of identifiers for the patient
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// A list of telecoms for the patient
	Telecom []ContactPoint `json:"telecom,omitempty"`

	// The gender of the patient
	Gender *PatientGender `json:"gender,omitempty"`

	// The date of birth for the patient
	BirthDate *string `json:"birthDate,omitempty"`

	// Indicates if the patient record is in active use
	Active *bool `json:"active,omitempty"`

	// A list of addresses for the patient
	Address []Address `json:"address,omitempty"`

	// Marital (civil) status of a patient
	MaritalStatus *common.CodeableConcept `json:"maritalStatus,omitempty"`

	// Multiple birth indicator and order
	MultipleBirthBoolean *bool `json:"multipleBirthBoolean,omitempty"`
	MultipleBirthInteger *int  `json:"multipleBirthInteger,omitempty"`

	// A list of photos of the patient
	Photo []Attachment `json:"photo,omitempty"`

	// A list of contacts for the patient
	Contact []PatientContact `json:"contact,omitempty"`

	// A list of communication preferences for the patient
	Communication []PatientCommunication `json:"communication,omitempty"`

	// A list of general practitioner(s) for the patient
	GeneralPractitioner []common.Reference `json:"generalPractitioner,omitempty"`

	// The organization that is the custodian of the patient record
	ManagingOrganization *common.Reference `json:"managingOrganization,omitempty"`

	// A list of links to other patient resources
	Link []PatientLink `json:"link,omitempty"`
}
