package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// Schedule represents a container for slots of time that may be available for booking appointments
type Schedule struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Schedule"

	// This element is labeled as a modifier because it may be used to mark that the resource was created in error
	Active *bool `json:"active,omitempty"`

	// The capacity to support multiple referenced resource types should be used in cases where the specific resources themselves cannot be scheduled without the other
	Actor []common.Reference `json:"actor"`

	// Comments on the availability to describe any extended information
	Comment *string `json:"comment,omitempty"`

	// External Ids for this item
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The period of time that the slots that reference this Schedule resource cover
	PlanningHorizon *common.Period `json:"planningHorizon,omitempty"`

	// A broad categorization of the service that is to be performed during this appointment
	ServiceCategory []common.CodeableConcept `json:"serviceCategory,omitempty"`

	// The specific service that is to be performed during this appointment
	ServiceType []common.CodeableConcept `json:"serviceType,omitempty"`

	// The specialty of a practitioner that would be required to perform the service requested in this appointment
	Specialty []common.CodeableConcept `json:"specialty,omitempty"`
}
