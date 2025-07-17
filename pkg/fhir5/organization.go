// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// OrganizationQualification represents official certifications and accreditations
type OrganizationQualification struct {
	common.BackboneElement

	// Coded representation of the qualification
	Code common.CodeableConcept `json:"code"`

	// An identifier for this qualification for the organization
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Organization that regulates and issues the qualification
	Issuer *common.Reference `json:"issuer,omitempty"`

	// Period during which the qualification is valid
	Period *common.Period `json:"period,omitempty"`
}

// Organization represents a formally or informally recognized grouping of people or organizations (R5)
type Organization struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Organization"

	// Whether the organization record is still in active use
	Active *bool `json:"active,omitempty"`

	// A list of alternate names that the organization is known as, or was known as in the past
	Alias []string `json:"alias,omitempty"`

	// Official contact details for the organization
	Contact []ExtendedContactDetail `json:"contact,omitempty"`

	// Additional details about the organization that could be displayed as further information
	Description *string `json:"description,omitempty"`

	// Technical endpoints providing access to services operated for the organization
	Endpoint []common.Reference `json:"endpoint,omitempty"`

	// Identifier for the organization that is used to identify the organization across multiple disparate systems
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Name used for the organization
	Name *string `json:"name,omitempty"`

	// The organization of which this organization forms a part
	PartOf *common.Reference `json:"partOf,omitempty"`

	// Official certifications, accreditations, training, designations and licenses
	Qualification []OrganizationQualification `json:"qualification,omitempty"`

	// Kind of organization
	Type []common.CodeableConcept `json:"type,omitempty"`
}

// ExtendedContactDetail represents extended contact information (R5)
type ExtendedContactDetail struct {
	common.Element

	// The type of contact
	Purpose *common.CodeableConcept `json:"purpose,omitempty"`

	// Name of an individual to contact
	Name []HumanName `json:"name,omitempty"`

	// Contact details (telephone, email, etc.) for a contact
	Telecom []ContactPoint `json:"telecom,omitempty"`

	// Address for the contact
	Address *Address `json:"address,omitempty"`

	// The organization that is associated with the contact
	Organization *common.Reference `json:"organization,omitempty"`

	// Period that this contact was valid for usage
	Period *common.Period `json:"period,omitempty"`
}
