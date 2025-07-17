package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// PractitionerRole represents a specific set of roles/labels/specialties that a practitioner may have at an organization
type PractitionerRole struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "PractitionerRole"

	// If this value is false, you may refer to the period to see when the role was in active use
	Active *bool `json:"active,omitempty"`

	// A description of site availability exceptions, e.g. public holiday availability
	AvailabilityExceptions *string `json:"availabilityExceptions,omitempty"`

	// More detailed availability information may be provided in associated Schedule/Slot resources
	AvailableTime []PractitionerRoleAvailableTime `json:"availableTime,omitempty"`

	// A person may have more than one role
	Code []common.CodeableConcept `json:"code,omitempty"`

	// Technical endpoints providing access to services operated for the practitioner with this role
	Endpoint []common.Reference `json:"endpoint,omitempty"`

	// The list of healthcare services that this worker provides for this role's Organization/Location(s)
	HealthcareService []common.Reference `json:"healthcareService,omitempty"`

	// Business Identifiers that are specific to a role/location
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The location(s) at which this practitioner provides care
	Location []common.Reference `json:"location,omitempty"`

	// The practitioner is not available or performing this role during this period of time due to the provided reason
	NotAvailable []PractitionerRoleNotAvailable `json:"notAvailable,omitempty"`

	// The organization where the Practitioner performs the roles associated
	Organization *common.Reference `json:"organization,omitempty"`

	// The period during which the person is authorized to act as a practitioner in these role(s) for the organization
	Period *common.Period `json:"period,omitempty"`

	// Practitioner that is able to provide the defined services for the organization
	Practitioner *common.Reference `json:"practitioner,omitempty"`

	// Specific specialty of the practitioner
	Specialty []common.CodeableConcept `json:"specialty,omitempty"`

	// Contact details that are specific to the role/location/service
	Telecom []common.ContactPoint `json:"telecom,omitempty"`
}

// PractitionerRoleAvailableTime represents more detailed availability information may be provided in associated Schedule/Slot resources
type PractitionerRoleAvailableTime struct {
	common.BackboneElement

	// Indicates which days of the week are available between the start and end Times
	AvailableEndTime *string `json:"availableEndTime,omitempty"`

	// Indicates which days of the week are available between the start and end Times
	AvailableStartTime *string `json:"availableStartTime,omitempty"`

	// Indicates which days of the week are available between the start and end Times
	DaysOfWeek []PractitionerRoleAvailableTimeDaysOfWeek `json:"daysOfWeek,omitempty"`
}

// PractitionerRoleAvailableTimeDaysOfWeek represents indicates which days of the week are available between the start and end Times
type PractitionerRoleAvailableTimeDaysOfWeek string

const (
	PractitionerRoleAvailableTimeDaysOfWeekMon PractitionerRoleAvailableTimeDaysOfWeek = "mon"
	PractitionerRoleAvailableTimeDaysOfWeekTue PractitionerRoleAvailableTimeDaysOfWeek = "tue"
	PractitionerRoleAvailableTimeDaysOfWeekWed PractitionerRoleAvailableTimeDaysOfWeek = "wed"
	PractitionerRoleAvailableTimeDaysOfWeekThu PractitionerRoleAvailableTimeDaysOfWeek = "thu"
	PractitionerRoleAvailableTimeDaysOfWeekFri PractitionerRoleAvailableTimeDaysOfWeek = "fri"
	PractitionerRoleAvailableTimeDaysOfWeekSat PractitionerRoleAvailableTimeDaysOfWeek = "sat"
	PractitionerRoleAvailableTimeDaysOfWeekSun PractitionerRoleAvailableTimeDaysOfWeek = "sun"
)

// PractitionerRoleNotAvailable represents the practitioner is not available or performing this role during this period of time due to the provided reason
type PractitionerRoleNotAvailable struct {
	common.BackboneElement

	// The reason that can be presented to the user as to why this time is not available
	Description string `json:"description"`

	// Service is not available (seasonally or for a public holiday) from this date
	During *common.Period `json:"during,omitempty"`
}
