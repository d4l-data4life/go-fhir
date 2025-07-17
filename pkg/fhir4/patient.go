// Package fhir4 contains FHIR R4 (version 4.0.1) resource definitions
package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// Patient represents information about an individual or animal receiving care
type Patient struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Patient"

	// Whether this patient record is in active use
	Active *bool `json:"active,omitempty"`

	// An identifier for this patient
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// A name associated with the patient
	Name []HumanName `json:"name,omitempty"`

	// A contact detail for the individual
	Telecom []ContactPoint `json:"telecom,omitempty"`

	// male | female | other | unknown
	Gender *AdministrativeGender `json:"gender,omitempty"`

	// The date of birth for the individual
	BirthDate *string `json:"birthDate,omitempty"`

	// Indicates if the individual is deceased or not
	DeceasedBoolean  *bool   `json:"deceasedBoolean,omitempty"`
	DeceasedDateTime *string `json:"deceasedDateTime,omitempty"`

	// An address for the individual
	Address []Address `json:"address,omitempty"`

	// Marital (civil) status of a patient
	MaritalStatus *common.CodeableConcept `json:"maritalStatus,omitempty"`

	// Whether patient is part of a multiple birth
	MultipleBirthBoolean *bool `json:"multipleBirthBoolean,omitempty"`
	MultipleBirthInteger *int  `json:"multipleBirthInteger,omitempty"`

	// Image of the patient
	Photo []Attachment `json:"photo,omitempty"`

	// A contact party (e.g. guardian, partner, friend) for the patient
	Contact []PatientContact `json:"contact,omitempty"`

	// A language which may be used to communicate with the patient
	Communication []PatientCommunication `json:"communication,omitempty"`

	// Patient's nominated primary care provider
	GeneralPractitioner []common.Reference `json:"generalPractitioner,omitempty"`

	// Organization that is the custodian of the patient record
	ManagingOrganization *common.Reference `json:"managingOrganization,omitempty"`

	// Link to another patient resource that concerns the same actual person
	Link []PatientLink `json:"link,omitempty"`
}

// HumanName represents a human name
type HumanName struct {
	common.Element

	// usual | official | temp | nickname | anonymous | old | maiden
	Use *NameUse `json:"use,omitempty"`

	// Text representation of the full name
	Text *string `json:"text,omitempty"`

	// Family name (often called 'Surname')
	Family *string `json:"family,omitempty"`

	// Given names (not always 'first'). Includes middle names
	Given []string `json:"given,omitempty"`

	// Parts that come before the name
	Prefix []string `json:"prefix,omitempty"`

	// Parts that come after the name
	Suffix []string `json:"suffix,omitempty"`

	// Time period when name was/is in use
	Period *common.Period `json:"period,omitempty"`
}

// NameUse represents the use of a human name
type NameUse string

const (
	NameUseUsual     NameUse = "usual"
	NameUseOfficial  NameUse = "official"
	NameUseTemp      NameUse = "temp"
	NameUseNickname  NameUse = "nickname"
	NameUseAnonymous NameUse = "anonymous"
	NameUseOld       NameUse = "old"
	NameUseMaiden    NameUse = "maiden"
)

// AdministrativeGender represents administrative gender codes
type AdministrativeGender string

const (
	AdministrativeGenderMale    AdministrativeGender = "male"
	AdministrativeGenderFemale  AdministrativeGender = "female"
	AdministrativeGenderOther   AdministrativeGender = "other"
	AdministrativeGenderUnknown AdministrativeGender = "unknown"
)

// PatientContact represents a contact party for the patient
type PatientContact struct {
	common.BackboneElement

	// The kind of relationship
	Relationship []common.CodeableConcept `json:"relationship,omitempty"`

	// A name associated with the contact person
	Name *HumanName `json:"name,omitempty"`

	// A contact detail for the person
	Telecom []ContactPoint `json:"telecom,omitempty"`

	// Address for the contact person
	Address *Address `json:"address,omitempty"`

	// male | female | other | unknown
	Gender *AdministrativeGender `json:"gender,omitempty"`

	// Organization that is associated with the contact
	Organization *common.Reference `json:"organization,omitempty"`

	// The period during which this contact person or organization is valid
	Period *common.Period `json:"period,omitempty"`
}

// PatientCommunication represents a language which may be used to communicate with the patient
type PatientCommunication struct {
	common.BackboneElement

	// The language which can be used to communicate with the patient
	Language common.CodeableConcept `json:"language"`

	// Language preference indicator
	Preferred *bool `json:"preferred,omitempty"`
}

// PatientLink represents a link to another patient resource that concerns the same actual person
type PatientLink struct {
	common.BackboneElement

	// The other patient resource that the link refers to
	Other common.Reference `json:"other"`

	// replaced-by | replaces | refer | seealso
	Type PatientLinkType `json:"type"`
}

// PatientLinkType represents the type of link between this patient resource and another patient resource
type PatientLinkType string

const (
	PatientLinkTypeReplacedBy PatientLinkType = "replaced-by"
	PatientLinkTypeReplaces   PatientLinkType = "replaces"
	PatientLinkTypeRefer      PatientLinkType = "refer"
	PatientLinkTypeSeeAlso    PatientLinkType = "seealso"
)
