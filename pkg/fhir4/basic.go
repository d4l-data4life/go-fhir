package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// Basic represents a resource for capturing information that doesn't fit into the other resource types
type Basic struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Basic"

	// Identifier for this resource
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Identifies the 'type' of resource - equivalent to the resource name for other resources
	Code common.CodeableConcept `json:"code"`

	// Identifies the patient, practitioner, device or any other resource that is the "focus" of this resource
	Subject *common.Reference `json:"subject,omitempty"`

	// Identifies when the resource was first created
	Created *string `json:"created,omitempty"`

	// Indicates who was responsible for creating the resource
	Author *common.Reference `json:"author,omitempty"`
}
