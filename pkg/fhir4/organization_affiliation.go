package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// OrganizationAffiliation represents a formal or informal recognized grouping of people or organizations
type OrganizationAffiliation struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "OrganizationAffiliation"

	// If this value is false, you may refer to the period to see when the role was in active use
	Active *bool `json:"active,omitempty"`

	// Definition of the role the participatingOrganization plays in the association
	Code []common.CodeableConcept `json:"code,omitempty"`

	// Technical endpoints providing access to services operated for this role
	Endpoint []common.Reference `json:"endpoint,omitempty"`

	// Healthcare services provided through the role
	HealthcareService []common.Reference `json:"healthcareService,omitempty"`

	// Business identifiers that are specific to this role
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The location(s) at which the role occurs
	Location []common.Reference `json:"location,omitempty"`

	// Health insurance provider network in which the participatingOrganization provides the role's services
	Network []common.Reference `json:"network,omitempty"`

	// Organization where the role is available (primary organization/has members)
	Organization *common.Reference `json:"organization,omitempty"`

	// The Participating Organization provides/performs the role(s) defined by the code to the Primary Organization
	ParticipatingOrganization *common.Reference `json:"participatingOrganization,omitempty"`

	// The period during which the participatingOrganization is affiliated with the primary organization
	Period *common.Period `json:"period,omitempty"`

	// Specific specialty of the participatingOrganization in the context of the role
	Specialty []common.CodeableConcept `json:"specialty,omitempty"`

	// Contact details at the participatingOrganization relevant to this Affiliation
	Telecom []common.ContactPoint `json:"telecom,omitempty"`
}
