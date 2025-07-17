// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// PractitionerCommunication represents communication details for a practitioner in R5
type PractitionerCommunication struct {
	common.BackboneElement

	// The language which can be used to communicate with the practitioner
	Language common.CodeableConcept `json:"language"`

	// Language preference indicator
	Preferred *bool `json:"preferred,omitempty"`
}

// PractitionerQualification represents qualifications obtained by training and certification
type PractitionerQualification struct {
	common.BackboneElement

	// Coded representation of the qualification
	Code common.CodeableConcept `json:"code"`

	// An identifier for this qualification for the practitioner
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Organization that regulates and issues the qualification
	Issuer *common.Reference `json:"issuer,omitempty"`

	// Period during which the qualification is valid
	Period *common.Period `json:"period,omitempty"`
}

// Practitioner represents a person who is directly or indirectly involved in the provisioning of healthcare or related services
type Practitioner struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Practitioner"

	// Whether this practitioner's record is in active use
	Active *bool `json:"active,omitempty"`

	// Addresses where the practitioner can be found or visited or to which mail can be delivered
	Address []Address `json:"address,omitempty"`

	// The date of birth for the practitioner
	BirthDate *string `json:"birthDate,omitempty"`

	// A language the practitioner can use in patient communication (R5 uses structured communication)
	Communication []PractitionerCommunication `json:"communication,omitempty"`

	// Indicates if the practitioner is deceased (R5 feature)
	DeceasedBoolean  *bool   `json:"deceasedBoolean,omitempty"`
	DeceasedDateTime *string `json:"deceasedDateTime,omitempty"`

	// Administrative Gender - the gender that the person is considered to have for administration and record keeping purposes
	Gender *PractitionerGender `json:"gender,omitempty"`

	// An identifier that applies to this person in this role
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The name of the practitioner
	Name []HumanName `json:"name,omitempty"`

	// Image of the person
	Photo []Attachment `json:"photo,omitempty"`

	// Qualifications obtained by training and certification
	Qualification []PractitionerQualification `json:"qualification,omitempty"`

	// A contact detail for the practitioner
	Telecom []ContactPoint `json:"telecom,omitempty"`
}

// PractitionerGender represents the gender of a practitioner
type PractitionerGender string

const (
	PractitionerGenderMale    PractitionerGender = "male"
	PractitionerGenderFemale  PractitionerGender = "female"
	PractitionerGenderOther   PractitionerGender = "other"
	PractitionerGenderUnknown PractitionerGender = "unknown"
)
