// Package fhir4 contains FHIR R4 (version 4.0.1) resource definitions
package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// Organization represents a formally or informally recognized grouping of people or organizations
type Organization struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Organization"

	// Whether the organization's record is still in active use
	Active *bool `json:"active,omitempty"`

	// Identifies this organization across multiple systems
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Name used for the organization
	Name *string `json:"name,omitempty"`

	// A list of alternate names that the organization is known as, or was known as in the past
	Alias []string `json:"alias,omitempty"`

	// Kind of organization
	Type []common.CodeableConcept `json:"type,omitempty"`

	// A contact detail for the organization
	Telecom []ContactPoint `json:"telecom,omitempty"`

	// An address for the organization
	Address []Address `json:"address,omitempty"`

	// The organization of which this organization forms a part
	PartOf *common.Reference `json:"partOf,omitempty"`

	// Contact for the organization for a certain purpose
	Contact []OrganizationContact `json:"contact,omitempty"`

	// Technical endpoints providing access to services operated for the organization
	Endpoint []common.Reference `json:"endpoint,omitempty"`
}

// OrganizationContact represents a contact for the organization
type OrganizationContact struct {
	common.BackboneElement

	// The type of contact
	Purpose *common.CodeableConcept `json:"purpose,omitempty"`

	// A name associated with the contact
	Name *HumanName `json:"name,omitempty"`

	// Contact details (telephone, email, etc.) for a contact
	Telecom []ContactPoint `json:"telecom,omitempty"`

	// Visiting or postal addresses for the contact
	Address *Address `json:"address,omitempty"`
}
