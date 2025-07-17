// Package fhir4 contains FHIR R4 (version 4.0.1) resource definitions
package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// Practitioner represents a person who is directly or indirectly involved in the provisioning of healthcare
type Practitioner struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Practitioner"

	// Whether this practitioner's record is in active use
	Active *bool `json:"active,omitempty"`

	// An identifier that applies to this person
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The name(s) associated with the practitioner
	Name []HumanName `json:"name,omitempty"`

	// A contact detail for the practitioner (that apply to all roles)
	Telecom []ContactPoint `json:"telecom,omitempty"`

	// Address(es) of the practitioner that are not role specific (typically home address)
	Address []Address `json:"address,omitempty"`

	// male | female | other | unknown
	Gender *AdministrativeGender `json:"gender,omitempty"`

	// The date of birth for the practitioner
	BirthDate *string `json:"birthDate,omitempty"`

	// Image of the person
	Photo []Attachment `json:"photo,omitempty"`

	// Certification, licenses, or training pertaining to the provision of care
	Qualification []PractitionerQualification `json:"qualification,omitempty"`

	// A language the practitioner can use in patient communication
	Communication []common.CodeableConcept `json:"communication,omitempty"`
}

// PractitionerQualification represents a qualification of the practitioner
type PractitionerQualification struct {
	common.BackboneElement

	// An identifier for this qualification for the practitioner
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Coded representation of the qualification
	Code common.CodeableConcept `json:"code"`

	// Period during which the qualification is valid
	Period *common.Period `json:"period,omitempty"`

	// Organization that regulates and issues the qualification
	Issuer *common.Reference `json:"issuer,omitempty"`
}
