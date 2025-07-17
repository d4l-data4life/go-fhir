package fhir4

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

// AppointmentParticipantStatus represents the participation status of a participant
type AppointmentParticipantStatus string

const (
	AppointmentParticipantStatusAccepted    AppointmentParticipantStatus = "accepted"
	AppointmentParticipantStatusDeclined    AppointmentParticipantStatus = "declined"
	AppointmentParticipantStatusTentative   AppointmentParticipantStatus = "tentative"
	AppointmentParticipantStatusNeedsAction AppointmentParticipantStatus = "needs-action"
)

// AppointmentParticipant represents participants involved in the appointment
type AppointmentParticipant struct {
	common.BackboneElement

	// Role of participant in the appointment
	Type []common.CodeableConcept `json:"type,omitempty"`

	// A Person, Location/HealthcareService or Device that is participating in the appointment
	Actor *common.Reference `json:"actor,omitempty"`

	// Whether this participant is required to be present at the meeting
	Required *string `json:"required,omitempty"`

	// Participation status of the actor
	Status AppointmentParticipantStatus `json:"status"`

	// Period of participation during the appointment
	Period *common.Period `json:"period,omitempty"`
}

// Appointment represents a booking of a healthcare event among patient(s), practitioner(s), related person(s) and/or device(s) for a specific date/time
type Appointment struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Appointment"

	// This records identifiers associated with this appointment concern that are defined by business processes and/or used to refer to it when a direct URL reference to the resource itself is not appropriate
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The overall status of the Appointment. Each of the participants has their own participation status which indicates their involvement in the process, however this status indicates the shared status
	Status AppointmentStatus `json:"status"`

	// A broad categorization of the service that is to be performed during this appointment
	ServiceCategory []common.CodeableConcept `json:"serviceCategory,omitempty"`

	// The specific service that is to be performed during this appointment
	ServiceType []common.CodeableConcept `json:"serviceType,omitempty"`

	// The specialty of a practitioner that would be required to perform the service requested in this appointment
	Specialty []common.CodeableConcept `json:"specialty,omitempty"`

	// The style of appointment or patient that has been booked in the slot (not service type)
	AppointmentType *common.CodeableConcept `json:"appointmentType,omitempty"`

	// The coded reason that this appointment is being scheduled
	ReasonCode []common.CodeableConcept `json:"reasonCode,omitempty"`

	// Reason the appointment has been scheduled, as specified using information from another resource
	ReasonReference []common.Reference `json:"reasonReference,omitempty"`

	// The priority of the appointment. Can be used to make informed decisions if needing to re-prioritize appointments
	Priority *int `json:"priority,omitempty"`

	// The brief description of the appointment as would be shown on a subject line in a meeting request, or appointment list
	Description *string `json:"description,omitempty"`

	// Additional information to support the appointment
	SupportingInformation []common.Reference `json:"supportingInformation,omitempty"`

	// Date/Time that the appointment is to take place
	Start *string `json:"start,omitempty"`

	// Date/Time that the appointment is to conclude
	End *string `json:"end,omitempty"`

	// Number of minutes that the appointment is to take
	MinutesDuration *int `json:"minutesDuration,omitempty"`

	// The slots that this appointment is filling
	Slot []common.Reference `json:"slot,omitempty"`

	// The date that this appointment was initially created
	Created *string `json:"created,omitempty"`

	// Additional comments about the appointment
	Comment *string `json:"comment,omitempty"`

	// While appointments can be made for a specific patient, they are often made for a group of patients
	PatientInstruction *string `json:"patientInstruction,omitempty"`

	// The service request this appointment is allocated to assess
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// List of participants involved in the appointment
	Participant []AppointmentParticipant `json:"participant"`

	// A set of date ranges (potentially including times) that the appointment is preferred to be scheduled within
	RequestedPeriod []common.Period `json:"requestedPeriod,omitempty"`
}
