package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// HealthcareService represents the details of a healthcare service available at a location
type HealthcareService struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "HealthcareService"

	// This element is labeled as a modifier because it may be used to mark that the resource was created in error
	Active *bool `json:"active,omitempty"`

	// Indicates whether or not a prospective consumer will require an appointment for a particular service at a site
	AppointmentRequired *bool `json:"appointmentRequired,omitempty"`

	// A description of site availability exceptions, e.g. public holiday availability
	AvailabilityExceptions *string `json:"availabilityExceptions,omitempty"`

	// More detailed availability information may be provided in associated Schedule/Slot resources
	AvailableTime []HealthcareServiceAvailableTime `json:"availableTime,omitempty"`

	// Selecting a Service Category then determines the list of relevant service types that can be selected in the primary service type
	Category []common.CodeableConcept `json:"category,omitempty"`

	// These could be such things as is wheelchair accessible
	Characteristic []common.CodeableConcept `json:"characteristic,omitempty"`

	// Would expect that a user would not see this information on a search results, and it would only be available when viewing the complete details of the service
	Comment *string `json:"comment,omitempty"`

	// When using this property it indicates that the service is available with this language
	Communication []common.CodeableConcept `json:"communication,omitempty"`

	// The locations referenced by the coverage area can include both specific locations, including areas, and also conceptual domains too
	CoverageArea []common.Reference `json:"coverageArea,omitempty"`

	// Does this service have specific eligibility requirements that need to be met in order to use the service?
	Eligibility []HealthcareServiceEligibility `json:"eligibility,omitempty"`

	// Technical endpoints providing access to services operated for the specific healthcare services defined at this resource
	Endpoint []common.Reference `json:"endpoint,omitempty"`

	// Extra details about the service that can't be placed in the other fields
	ExtraDetails *string `json:"extraDetails,omitempty"`

	// External identifiers for this item
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The location(s) where this healthcare service may be provided
	Location []common.Reference `json:"location,omitempty"`

	// Further description of the service as it would be presented to a consumer while searching
	Name *string `json:"name,omitempty"`

	// The HealthcareService is not available during this period of time due to the provided reason
	NotAvailable []HealthcareServiceNotAvailable `json:"notAvailable,omitempty"`

	// If there is a photo/symbol associated with this HealthcareService, it may be included here to facilitate quick identification of the service in a list
	Photo *common.Attachment `json:"photo,omitempty"`

	// Programs are often defined externally to an Organization, commonly by governments
	Program []common.CodeableConcept `json:"program,omitempty"`

	// This property is recommended to be the same as the Location's managingOrganization
	ProvidedBy *common.Reference `json:"providedBy,omitempty"`

	// Ways that the service accepts referrals, if this is not provided then it is implied that no referral is required
	ReferralMethod []common.CodeableConcept `json:"referralMethod,omitempty"`

	// The provision means being commissioned by, contractually obliged or financially sourced
	ServiceProvisionCode []common.CodeableConcept `json:"serviceProvisionCode,omitempty"`

	// Collection of specialties handled by the service site
	Specialty []common.CodeableConcept `json:"specialty,omitempty"`

	// If this is empty, then refer to the location's contacts
	Telecom []common.ContactPoint `json:"telecom,omitempty"`

	// The specific type of service that may be delivered or performed
	Type []common.CodeableConcept `json:"type,omitempty"`
}

// HealthcareServiceEligibility represents does this service have specific eligibility requirements
type HealthcareServiceEligibility struct {
	common.BackboneElement

	// Coded value for the eligibility
	Code *common.CodeableConcept `json:"code,omitempty"`

	// The description of service eligibility should, in general, not exceed one or two paragraphs
	Comment *string `json:"comment,omitempty"`
}

// HealthcareServiceAvailableTime represents more detailed availability information
type HealthcareServiceAvailableTime struct {
	common.BackboneElement

	// Is this always available? (hence times are irrelevant) e.g. 24 hour service
	AllDay *bool `json:"allDay,omitempty"`

	// The time zone is expected to be for where this HealthcareService is provided at
	AvailableEndTime *string `json:"availableEndTime,omitempty"`

	// The time zone is expected to be for where this HealthcareService is provided at
	AvailableStartTime *string `json:"availableStartTime,omitempty"`

	// Indicates which days of the week are available between the start and end Times
	DaysOfWeek []HealthcareServiceAvailableTimeDaysOfWeek `json:"daysOfWeek,omitempty"`
}

// HealthcareServiceAvailableTimeDaysOfWeek represents the days of the week
type HealthcareServiceAvailableTimeDaysOfWeek string

const (
	HealthcareServiceAvailableTimeDaysOfWeekMon HealthcareServiceAvailableTimeDaysOfWeek = "mon"
	HealthcareServiceAvailableTimeDaysOfWeekTue HealthcareServiceAvailableTimeDaysOfWeek = "tue"
	HealthcareServiceAvailableTimeDaysOfWeekWed HealthcareServiceAvailableTimeDaysOfWeek = "wed"
	HealthcareServiceAvailableTimeDaysOfWeekThu HealthcareServiceAvailableTimeDaysOfWeek = "thu"
	HealthcareServiceAvailableTimeDaysOfWeekFri HealthcareServiceAvailableTimeDaysOfWeek = "fri"
	HealthcareServiceAvailableTimeDaysOfWeekSat HealthcareServiceAvailableTimeDaysOfWeek = "sat"
	HealthcareServiceAvailableTimeDaysOfWeekSun HealthcareServiceAvailableTimeDaysOfWeek = "sun"
)

// HealthcareServiceNotAvailable represents the HealthcareService is not available during this period of time
type HealthcareServiceNotAvailable struct {
	common.BackboneElement

	// The reason that can be presented to the user as to why this time is not available
	Description string `json:"description"`

	// Service is not available (seasonally or for a public holiday) from this date
	During *common.Period `json:"during,omitempty"`
}
