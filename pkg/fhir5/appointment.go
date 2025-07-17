// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// AppointmentStatus represents the status of an appointment
type AppointmentStatus string

const (
	AppointmentStatusProposed       AppointmentStatus = "proposed"
	AppointmentStatusPending        AppointmentStatus = "pending"
	AppointmentStatusBooked         AppointmentStatus = "booked"
	AppointmentStatusArrived        AppointmentStatus = "arrived"
	AppointmentStatusFulfilled      AppointmentStatus = "fulfilled"
	AppointmentStatusCancelled      AppointmentStatus = "cancelled"
	AppointmentStatusNoShow         AppointmentStatus = "no-show"
	AppointmentStatusEnteredInError AppointmentStatus = "entered-in-error"
	AppointmentStatusCheckedIn      AppointmentStatus = "checked-in"
	AppointmentStatusWaitlist       AppointmentStatus = "waitlist"
)

// AppointmentParticipantStatus represents the status of an appointment participant
type AppointmentParticipantStatus string

const (
	AppointmentParticipantStatusAccepted    AppointmentParticipantStatus = "accepted"
	AppointmentParticipantStatusDeclined    AppointmentParticipantStatus = "declined"
	AppointmentParticipantStatusTentative   AppointmentParticipantStatus = "tentative"
	AppointmentParticipantStatusNeedsAction AppointmentParticipantStatus = "needs-action"
)

// AppointmentRecurrenceTemplateWeeklyTemplate represents weekly recurrence template
type AppointmentRecurrenceTemplateWeeklyTemplate struct {
	common.BackboneElement

	// Days of the week when the appointment recurs
	Monday    *bool `json:"monday,omitempty"`
	Tuesday   *bool `json:"tuesday,omitempty"`
	Wednesday *bool `json:"wednesday,omitempty"`
	Thursday  *bool `json:"thursday,omitempty"`
	Friday    *bool `json:"friday,omitempty"`
	Saturday  *bool `json:"saturday,omitempty"`
	Sunday    *bool `json:"sunday,omitempty"`

	// The number of weeks between recurrences
	WeekInterval *int `json:"weekInterval,omitempty"`
}

// AppointmentRecurrenceTemplateMonthlyTemplate represents monthly recurrence template
type AppointmentRecurrenceTemplateMonthlyTemplate struct {
	common.BackboneElement

	// The day of the month when the appointment recurs
	DayOfMonth *int `json:"dayOfMonth,omitempty"`

	// Indicates which week within the month the appointment should recur
	NthWeekOfMonth *int `json:"nthWeekOfMonth,omitempty"`

	// Indicates which day of the week the appointment should recur
	DayOfWeek *common.Coding `json:"dayOfWeek,omitempty"`

	// The number of months between recurrences
	MonthInterval *int `json:"monthInterval,omitempty"`
}

// AppointmentRecurrenceTemplateYearlyTemplate represents yearly recurrence template
type AppointmentRecurrenceTemplateYearlyTemplate struct {
	common.BackboneElement

	// The number of years between recurrences
	YearInterval *int `json:"yearInterval,omitempty"`
}

// AppointmentRecurrenceTemplate represents recurrence template for appointments
type AppointmentRecurrenceTemplate struct {
	common.BackboneElement

	// The timezone of the recurring appointment occurrences
	Timezone *common.CodeableConcept `json:"timezone,omitempty"`

	// How often the appointment series should recur
	RecurrenceType common.CodeableConcept `json:"recurrenceType"`

	// The date when the recurrence should end
	LastOccurrenceDate *string `json:"lastOccurrenceDate,omitempty"`

	// How many appointment occurrences are planned for the series
	OccurrenceCount *int `json:"occurrenceCount,omitempty"`

	// Recurring dates that should be excluded from the series
	ExcludingDate []string `json:"excludingDate,omitempty"`

	// Identifiers of occurrences that should be excluded from the series
	ExcludingRecurrenceId []int `json:"excludingRecurrenceId,omitempty"`

	// Template for weekly recurring appointments
	WeeklyTemplate *AppointmentRecurrenceTemplateWeeklyTemplate `json:"weeklyTemplate,omitempty"`

	// Template for monthly recurring appointments
	MonthlyTemplate *AppointmentRecurrenceTemplateMonthlyTemplate `json:"monthlyTemplate,omitempty"`

	// Template for yearly recurring appointments
	YearlyTemplate *AppointmentRecurrenceTemplateYearlyTemplate `json:"yearlyTemplate,omitempty"`
}

// AppointmentParticipant represents a participant in an appointment
type AppointmentParticipant struct {
	common.BackboneElement

	// Person, Location, HealthcareService, PractitionerRole, RelatedPerson or Device
	Actor *common.Reference `json:"actor,omitempty"`

	// Whether this participant is required to be present at the meeting
	Required *bool `json:"required,omitempty"`

	// Participation status of the participant
	Status AppointmentParticipantStatus `json:"status"`

	// Role of participant in the appointment
	Type []common.CodeableConcept `json:"type,omitempty"`

	// Participation period of the participant
	Period *common.Period `json:"period,omitempty"`
}

// Appointment represents a booking of a healthcare event among patient(s), practitioner(s), related person(s) and/or device(s)
type Appointment struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Appointment"

	// The style of appointment or patient that has been booked
	AppointmentType *common.CodeableConcept `json:"appointmentType,omitempty"`

	// The request this appointment is allocated to assess
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// The coded reason for the appointment being cancelled
	CancellationReason *common.CodeableConcept `json:"cancellationReason,omitempty"`

	// Concepts representing classification of patient encounter
	Class []common.CodeableConcept `json:"class,omitempty"`

	// The date that this appointment was initially created
	Created *string `json:"created,omitempty"`

	// Shown on a subject line in a meeting request, or appointment list
	Description *string `json:"description,omitempty"`

	// When appointment is to conclude
	End *string `json:"end,omitempty"`

	// External identifier for the appointment
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Can be less than start/end
	MinutesDuration *int `json:"minutesDuration,omitempty"`

	// Additional information to support the appointment
	Note []Annotation `json:"note,omitempty"`

	// The originating appointment in a recurring set of appointments
	OriginatingAppointment *common.Reference `json:"originatingAppointment,omitempty"`

	// List of participants involved in the appointment
	Participant []AppointmentParticipant `json:"participant"`

	// Detailed information and instructions for the patient
	PatientInstruction []CodeableReference `json:"patientInstruction,omitempty"`

	// Used to make informed decisions if needing to re-prioritize
	Priority *common.CodeableConcept `json:"priority,omitempty"`

	// Coded reason this appointment is scheduled
	Reason []CodeableReference `json:"reason,omitempty"`

	// Details of the recurrence pattern/template for recurring appointments
	RecurrenceTemplate []AppointmentRecurrenceTemplate `json:"recurrenceTemplate,omitempty"`

	// Appointment replaced by this Appointment
	Replaces []common.Reference `json:"replaces,omitempty"`

	// Potential date/time interval(s) requested to allocate the appointment within
	RequestedPeriod []common.Period `json:"requestedPeriod,omitempty"`

	// A broad categorization of the service that is to be performed during this appointment
	ServiceCategory []common.CodeableConcept `json:"serviceCategory,omitempty"`

	// The specific service that is to be performed during this appointment
	ServiceType []CodeableReference `json:"serviceType,omitempty"`

	// The slots that this appointment is filling
	Slot []common.Reference `json:"slot,omitempty"`

	// The specialty of a practitioner that would be required to perform the service requested
	Specialty []common.CodeableConcept `json:"specialty,omitempty"`

	// When appointment is to take place
	Start *string `json:"start,omitempty"`

	// The overall status of the Appointment
	Status AppointmentStatus `json:"status"`

	// The patient or group associated with the appointment
	Subject *common.Reference `json:"subject,omitempty"`

	// Additional information to support the appointment
	SupportingInformation []common.Reference `json:"supportingInformation,omitempty"`

	// Connection details of a virtual service
	VirtualService []VirtualServiceDetail `json:"virtualService,omitempty"`
}
