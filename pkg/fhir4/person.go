package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// Person represents demographics and other administrative information about an individual or animal
type Person struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Person"

	// Whether this person's record is in active use
	Active *bool `json:"active,omitempty"`

	// Person may have multiple addresses with different uses or applicable periods
	Address []common.Address `json:"address,omitempty"`

	// At least an estimated year should be provided as a guess if the real DOB is unknown
	BirthDate *string `json:"birthDate,omitempty"`

	// The gender might not match the biological sex as determined by genetics
	Gender *AdministrativeGender `json:"gender,omitempty"`

	// Identifier for a person within a particular scope
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Link to a resource that concerns the same actual person
	Link []PersonLink `json:"link,omitempty"`

	// The organization that is the custodian of the person record
	ManagingOrganization *common.Reference `json:"managingOrganization,omitempty"`

	// Person may have multiple names with different uses or applicable periods
	Name []common.HumanName `json:"name,omitempty"`

	// An image that can be displayed as a thumbnail of the person to enhance the identification of the individual
	Photo *common.Attachment `json:"photo,omitempty"`

	// Person may have multiple ways to be contacted with different uses or applicable periods
	Telecom []common.ContactPoint `json:"telecom,omitempty"`
}

// PersonLink represents link to a resource that concerns the same actual person
type PersonLink struct {
	common.BackboneElement

	// Level of assurance that this link is associated with the target resource
	Assurance *PersonLinkAssurance `json:"assurance,omitempty"`

	// The resource to which this actual person is associated
	Target common.Reference `json:"target"`
}

// PersonLinkAssurance represents level of assurance that this link is associated with the target resource
type PersonLinkAssurance string

const (
	PersonLinkAssuranceLevel1 PersonLinkAssurance = "level1"
	PersonLinkAssuranceLevel2 PersonLinkAssurance = "level2"
	PersonLinkAssuranceLevel3 PersonLinkAssurance = "level3"
	PersonLinkAssuranceLevel4 PersonLinkAssurance = "level4"
)
