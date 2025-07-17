package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// AppointmentResponseParticipantStatus represents the participation status of a participant
type AppointmentResponseParticipantStatus string

const (
	AppointmentResponseParticipantStatusAccepted    AppointmentResponseParticipantStatus = "accepted"
	AppointmentResponseParticipantStatusDeclined    AppointmentResponseParticipantStatus = "declined"
	AppointmentResponseParticipantStatusTentative   AppointmentResponseParticipantStatus = "tentative"
	AppointmentResponseParticipantStatusNeedsAction AppointmentResponseParticipantStatus = "needs-action"
)

// AppointmentResponse represents a reply to an appointment request for a patient and/or practitioner(s)
type AppointmentResponse struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "AppointmentResponse"

	// This records identifiers associated with this appointment response concern that are defined by business processes and/or used to refer to it when a direct URL reference to the resource itself is not appropriate
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Appointment that this response is replying to
	Appointment common.Reference `json:"appointment"`

	// Date/Time that the appointment is to take place, or requested new start time
	Start *string `json:"start,omitempty"`

	// This may be either the same as the appointment request to confirm the details of the appointment, or alternately a new time to request a re-negotiation of the end time
	End *string `json:"end,omitempty"`

	// Role of participant in the appointment
	ParticipantType []common.CodeableConcept `json:"participantType,omitempty"`

	// A Person, Location/HealthcareService or Device that is participating in the appointment
	Actor *common.Reference `json:"actor,omitempty"`

	// Participation status of the participant
	ParticipantStatus AppointmentResponseParticipantStatus `json:"participantStatus"`

	// Additional comments about the appointment
	Comment *string `json:"comment,omitempty"`
}
