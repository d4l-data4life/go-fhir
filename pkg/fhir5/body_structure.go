// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark represents distance from landmark
type BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark struct {
	common.BackboneElement

	// An instrument, tool, analyzer, etc. used in the measurement
	Device []CodeableReference `json:"device,omitempty"`

	// The measured distance (e.g., in cm) from a body landmark
	Value []common.Quantity `json:"value,omitempty"`
}

// BodyStructureIncludedStructureBodyLandmarkOrientation represents body landmark orientation
type BodyStructureIncludedStructureBodyLandmarkOrientation struct {
	common.BackboneElement

	// An description of the direction away from a landmark something is located based on a radial clock dial
	ClockFacePosition []common.CodeableConcept `json:"clockFacePosition,omitempty"`

	// The distance in centimeters a certain observation is made from a body landmark
	DistanceFromLandmark []BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark `json:"distanceFromLandmark,omitempty"`

	// A description of a landmark on the body used as a reference to locate something else
	LandmarkDescription []common.CodeableConcept `json:"landmarkDescription,omitempty"`

	// The surface area a body location is in relation to a landmark
	SurfaceOrientation []common.CodeableConcept `json:"surfaceOrientation,omitempty"`
}

// BodyStructureIncludedStructure represents the anatomical location(s) or region(s)
type BodyStructureIncludedStructure struct {
	common.BackboneElement

	// Body locations in relation to a specific body landmark (tatoo, scar, other body structure)
	BodyLandmarkOrientation []BodyStructureIncludedStructureBodyLandmarkOrientation `json:"bodyLandmarkOrientation,omitempty"`

	// Code that represents the included structure laterality
	Laterality *common.CodeableConcept `json:"laterality,omitempty"`

	// Code that represents the included structure qualifier
	Qualifier []common.CodeableConcept `json:"qualifier,omitempty"`

	// XY or XYZ-coordinate orientation for structure
	SpatialReference []common.Reference `json:"spatialReference,omitempty"`

	// Code that represents the included structure
	Structure common.CodeableConcept `json:"structure"`
}

// BodyStructure represents a body structure
type BodyStructure struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "BodyStructure"

	// This element is labeled as a modifier because it may be used to mark that the resource was created in error
	Active *bool `json:"active,omitempty"`

	// This description could include any visual markings used to orientate the viewer
	Description *string `json:"description,omitempty"`

	// The anatomical location(s) or region(s) not occupied or represented by the specimen, lesion, or body structure
	ExcludedStructure []BodyStructureIncludedStructure `json:"excludedStructure,omitempty"`

	// Identifier for this instance of the anatomical structure
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Image or images used to identify a location
	Image []Attachment `json:"image,omitempty"`

	// The anatomical location(s) or region(s) of the specimen, lesion, or body structure
	IncludedStructure []BodyStructureIncludedStructure `json:"includedStructure"`

	// The minimum cardinality of 0 supports the use case of specifying a location without defining a morphology
	Morphology *common.CodeableConcept `json:"morphology,omitempty"`

	// The person to which the body site belongs
	Patient common.Reference `json:"patient"`
}
