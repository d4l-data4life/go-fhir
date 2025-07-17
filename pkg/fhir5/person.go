package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// PersonGender represents the gender of a person
type PersonGender string

const (
	PersonGenderMale    PersonGender = "male"
	PersonGenderFemale  PersonGender = "female"
	PersonGenderOther   PersonGender = "other"
	PersonGenderUnknown PersonGender = "unknown"
)

// PersonLinkAssurance represents the level of assurance that this link is associated with the target resource
type PersonLinkAssurance string

const (
	PersonLinkAssuranceLevel1 PersonLinkAssurance = "level1"
	PersonLinkAssuranceLevel2 PersonLinkAssurance = "level2"
	PersonLinkAssuranceLevel3 PersonLinkAssurance = "level3"
	PersonLinkAssuranceLevel4 PersonLinkAssurance = "level4"
)

// PersonLink represents a link to a resource that concerns the same actual person
type PersonLink struct {
	common.BackboneElement

	// Level of assurance that this link is associated with the target resource
	Assurance *PersonLinkAssurance `json:"assurance,omitempty"`

	// The resource to which this actual person is associated
	Target common.Reference `json:"target"`
}

// Person represents a person
type Person struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Person"

	// Whether this person's record is in active use
	Active *bool `json:"active,omitempty"`

	// One or more addresses for the person
	Address []Address `json:"address,omitempty"`

	// The date on which the person was born
	BirthDate *string `json:"birthDate,omitempty"`

	// Administrative Gender - the gender that the person is considered to have for administration and record keeping purposes
	Gender *PersonGender `json:"gender,omitempty"`

	// Identifier for a person within a particular scope
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// A link to a resource that concerns the same actual person
	Link []PersonLink `json:"link,omitempty"`

	// The organization that is the custodian of the person record
	ManagingOrganization *common.Reference `json:"managingOrganization,omitempty"`

	// A name associated with the person
	Name []HumanName `json:"name,omitempty"`

	// An image of the person
	Photo *Attachment `json:"photo,omitempty"`

	// A contact detail for the person
	Telecom []ContactPoint `json:"telecom,omitempty"`
}
