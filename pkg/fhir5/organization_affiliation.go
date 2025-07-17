// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// OrganizationAffiliation represents an affiliation between organizations
type OrganizationAffiliation struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "OrganizationAffiliation"

	// If this value is false, you may refer to the period to see when the role was in active use
	Active *bool `json:"active,omitempty"`

	// Definition of the role the participatingOrganization plays in the association
	Code []common.CodeableConcept `json:"code,omitempty"`

	// The contact details of communication devices available at the participatingOrganization
	Contact []ExtendedContactDetail `json:"contact,omitempty"`

	// Technical endpoints providing access to services operated for this role
	Endpoint []common.Reference `json:"endpoint,omitempty"`

	// Healthcare services provided through the role
	HealthcareService []common.Reference `json:"healthcareService,omitempty"`

	// Business identifiers that are specific to this role
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The location(s) at which the role occurs
	Location []common.Reference `json:"location,omitempty"`

	// e.g. Commonly used for Health Insurance provider networks
	Network []common.Reference `json:"network,omitempty"`

	// For example, a Spotless Cleaning Services is a supplier to General Hospital
	Organization *common.Reference `json:"organization,omitempty"`

	// See comments for OrganizationAffiliation.organization above
	ParticipatingOrganization *common.Reference `json:"participatingOrganization,omitempty"`

	// The period during which the participatingOrganization is affiliated
	Period *common.Period `json:"period,omitempty"`

	// Specific specialty of the participatingOrganization in the context of the role
	Specialty []common.CodeableConcept `json:"specialty,omitempty"`
}
