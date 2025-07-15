// Package fhir5 contains FHIR R5 resource definitions
package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// Slot represents a slot of time on a schedule that may be available for booking appointments
type Slot struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Slot"

	// A slot may be allow multiple appointment types, but when booked, would only be used for one of the given appointment types
	AppointmentType []common.CodeableConcept `json:"appointmentType,omitempty"`

	// Comments on the slot to describe any extended information
	Comment *string `json:"comment,omitempty"`

	// Date/Time that the slot is to conclude
	End string `json:"end"`

	// External Ids for this item
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// This slot has already been overbooked, appointments are unlikely to be accepted for this time
	Overbooked *bool `json:"overbooked,omitempty"`

	// The schedule resource that this slot defines an interval of status information
	Schedule common.Reference `json:"schedule"`

	// A broad categorization of the service that is to be performed during this appointment
	ServiceCategory []common.CodeableConcept `json:"serviceCategory,omitempty"`

	// The type of appointments that can be booked into this slot
	ServiceType []CodeableReference `json:"serviceType,omitempty"`

	// The specialty of a practitioner that would be required to perform the service requested in this appointment
	Specialty []common.CodeableConcept `json:"specialty,omitempty"`

	// Date/Time that the slot is to begin
	Start string `json:"start"`

	// busy | free | busy-unavailable | busy-tentative | entered-in-error
	Status string `json:"status"`
}
