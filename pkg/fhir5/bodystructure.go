// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark represents distance from landmark
type BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark struct {
	common.BackboneElement

	// Distance from the landmark
	Distance *common.Quantity `json:"distance,omitempty"`

	// Landmark relative to which the distance is measured
	Landmark *common.CodeableConcept `json:"landmark,omitempty"`
}

// BodyStructureIncludedStructureBodyLandmarkOrientation represents body landmark orientation
type BodyStructureIncludedStructureBodyLandmarkOrientation struct {
	common.BackboneElement

	// Distance from the landmark
	DistanceFromLandmark []BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark `json:"distanceFromLandmark,omitempty"`

	// Landmark relative to which the distance is measured
	Landmark *common.CodeableConcept `json:"landmark,omitempty"`
}

// BodyStructureIncludedStructure represents included structure
type BodyStructureIncludedStructure struct {
	common.BackboneElement

	// Body landmark orientation
	BodyLandmarkOrientation []BodyStructureIncludedStructureBodyLandmarkOrientation `json:"bodyLandmarkOrientation,omitempty"`

	// Code that represents the included structure
	Code *common.CodeableConcept `json:"code,omitempty"`

	// Code that represents the included structure laterality
	Laterality *common.CodeableConcept `json:"laterality,omitempty"`

	// Code that represents the included structure qualifier
	Qualifier []common.CodeableConcept `json:"qualifier,omitempty"`
}

// BodyStructure represents a specific and identified anatomical structure
type BodyStructure struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "BodyStructure"

	// Whether this record is in active use
	Active *bool `json:"active,omitempty"`

	// Kind of Structure
	Morphology *common.CodeableConcept `json:"morphology,omitempty"`

	// Body site
	Location *common.CodeableConcept `json:"location,omitempty"`

	// Body site modifier
	LocationQualifier []common.CodeableConcept `json:"locationQualifier,omitempty"`

	// Text description
	Description *string `json:"description,omitempty"`

	// Attached images
	Image []Attachment `json:"image,omitempty"`

	// Who this is about
	Patient common.Reference `json:"patient"`

	// Included anatomic location(s)
	IncludedStructure []BodyStructureIncludedStructure `json:"includedStructure,omitempty"`

	// Excluded anatomic location(s)
	ExcludedStructure []BodyStructureIncludedStructure `json:"excludedStructure,omitempty"`
}
