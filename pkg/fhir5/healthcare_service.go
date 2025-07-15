package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// HealthcareServiceEligibility represents specific eligibility requirements for using the service
type HealthcareServiceEligibility struct {
	common.BackboneElement

	// Coded value for the eligibility
	Code *common.CodeableConcept `json:"code,omitempty"`

	// The description of service eligibility
	Comment *string `json:"comment,omitempty"`
}

// HealthcareService represents the details of a healthcare service available at a location
type HealthcareService struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "HealthcareService"

	// This element is labeled as a modifier because it may be used to mark that the resource was created in error
	Active *bool `json:"active,omitempty"`

	// Indicates whether or not a prospective consumer will require an appointment for a particular service
	AppointmentRequired *bool `json:"appointmentRequired,omitempty"`

	// More detailed availability information may be provided in associated Schedule/Slot resources
	Availability []Availability `json:"availability,omitempty"`

	// Selecting a Service Category then determines the list of relevant service types
	Category []common.CodeableConcept `json:"category,omitempty"`

	// These could be such things as is wheelchair accessible
	Characteristic []common.CodeableConcept `json:"characteristic,omitempty"`

	// Would expect that a user would not see this information on a search results
	Comment *string `json:"comment,omitempty"`

	// When using this property it indicates that the service is available with this language
	Communication []common.CodeableConcept `json:"communication,omitempty"`

	// Official contacts for the HealthcareService itself for specific purposes
	Contact []ExtendedContactDetail `json:"contact,omitempty"`

	// The locations referenced by the coverage area can include both specific locations and areas
	CoverageArea []common.Reference `json:"coverageArea,omitempty"`

	// Does this service have specific eligibility requirements that need to be met in order to use the service?
	Eligibility []HealthcareServiceEligibility `json:"eligibility,omitempty"`

	// Technical endpoints providing access to services operated for the specific healthcare services
	Endpoint []common.Reference `json:"endpoint,omitempty"`

	// Extra details about the service that can't be placed in the other fields
	ExtraDetails *string `json:"extraDetails,omitempty"`

	// External identifiers for this item
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The location(s) where this healthcare service may be provided
	Location []common.Reference `json:"location,omitempty"`

	// Further description of the service as it would be presented to a consumer while searching
	Name *string `json:"name,omitempty"`

	// For example, if there is a generic Radiology service that offers CT Scans, MRIs, etc
	OfferedIn []common.Reference `json:"offeredIn,omitempty"`

	// If there is a photo/symbol associated with this HealthcareService
	Photo *Attachment `json:"photo,omitempty"`

	// Programs are often defined externally to an Organization
	Program []common.CodeableConcept `json:"program,omitempty"`

	// This property is recommended to be the same as the Location's managingOrganization
	ProvidedBy *common.Reference `json:"providedBy,omitempty"`

	// Ways that the service accepts referrals
	ReferralMethod []common.CodeableConcept `json:"referralMethod,omitempty"`

	// The provision means being commissioned by, contractually obliged or financially sourced
	ServiceProvisionCode []common.CodeableConcept `json:"serviceProvisionCode,omitempty"`

	// Collection of specialties handled by the Healthcare service
	Specialty []common.CodeableConcept `json:"specialty,omitempty"`

	// The specific type of service that may be delivered or performed
	Type []common.CodeableConcept `json:"type,omitempty"`
}
