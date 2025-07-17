package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// BodyStructure represents a specific and identified anatomical location
type BodyStructure struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "BodyStructure"

	// Identifier for this instance of the anatomical structure
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Whether this record is in active use
	Active *bool `json:"active,omitempty"`

	// The kind of structure being represented by the resource at BodyStructure.location
	Morphology *common.CodeableConcept `json:"morphology,omitempty"`

	// The anatomical location or region of the specimen, lesion, or body structure
	Location *common.CodeableConcept `json:"location,omitempty"`

	// Qualifier to refine the anatomical location
	LocationQualifier []common.CodeableConcept `json:"locationQualifier,omitempty"`

	// A summary, characterization or explanation of the body structure
	Description *string `json:"description,omitempty"`

	// Image or images used to identify a location
	Image []common.Attachment `json:"image,omitempty"`

	// The person to which the body site belongs
	Patient common.Reference `json:"patient"`
}
