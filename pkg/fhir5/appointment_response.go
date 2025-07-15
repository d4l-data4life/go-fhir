// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// AppointmentResponse represents a reply to an appointment request for a patient and/or practitioner(s)
type AppointmentResponse struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "AppointmentResponse"

	// A Person, Location, HealthcareService, or Device that is participating in the appointment
	Actor *common.Reference `json:"actor,omitempty"`

	// Appointment that this response is replying to
	Appointment common.Reference `json:"appointment"`

	// This comment is particularly important when the responder is declining, tentatively accepting or requesting another time
	Comment *string `json:"comment,omitempty"`

	// This may be either the same as the appointment request to confirm the details of the appointment, or alternately a new time to request a re-negotiation of the end time
	End *string `json:"end,omitempty"`

	// Business identifier for this appointment response
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The original date within a recurring request
	OccurrenceDate *string `json:"occurrenceDate,omitempty"`

	// This element is labeled as a modifier because the status contains the code entered-in-error
	ParticipantStatus AppointmentParticipantStatus `json:"participantStatus"`

	// The role of the participant can be used to declare what the actor will be doing in the scope of the referenced appointment
	ParticipantType []common.CodeableConcept `json:"participantType,omitempty"`

	// Indicates that the response is proposing a different time that was initially requested
	ProposedNewTime *bool `json:"proposedNewTime,omitempty"`

	// The recurrence ID (sequence number) of the specific appointment when responding to a recurring request
	RecurrenceId *int `json:"recurrenceId,omitempty"`

	// When a recurring appointment is requested, the participant may choose to respond to each individual occurrence
	Recurring *bool `json:"recurring,omitempty"`

	// This may be either the same as the appointment request to confirm the details of the appointment, or alternately a new time to request a re-negotiation of the start time
	Start *string `json:"start,omitempty"`
}
