// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// ArtifactAssessmentContent represents a component comment, classifier, or rating of the artifact
type ArtifactAssessmentContent struct {
	common.BackboneElement

	// Indicates who or what authored the content
	Author *common.Reference `json:"author,omitempty"`

	// Represents a rating, classifier, or assessment of the artifact
	Classifier []common.CodeableConcept `json:"classifier,omitempty"`

	// If the informationType is container, the components of the content
	Component []ArtifactAssessmentContent `json:"component,omitempty"`

	// Acceptable to publicly share the comment, classifier or rating
	FreeToShare *bool `json:"freeToShare,omitempty"`

	// The type of information this component of the content represents
	InformationType *string `json:"informationType,omitempty"`

	// The target element is used to point the comment to aspect of the artifact
	Path []string `json:"path,omitempty"`

	// A quantitative rating of the artifact
	Quantity *common.Quantity `json:"quantity,omitempty"`

	// Additional related artifacts that provide supporting documentation
	RelatedArtifact []interface{} `json:"relatedArtifact,omitempty"`

	// A brief summary of the content of this component
	Summary *string `json:"summary,omitempty"`

	// Indicates what type of content this component represents
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// ArtifactAssessment represents comments, classifiers or ratings about a Resource
// and supports attribution and rights management metadata for the added content.
type ArtifactAssessment struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "ArtifactAssessment"

	// The 'date' element may be more recent than the approval date because of minor changes or editorial corrections
	ApprovalDate *string `json:"approvalDate,omitempty"`

	// A reference to a resource, canonical resource, or non-FHIR resource which the comment or assessment is about
	ArtifactReference *common.Reference `json:"artifactReference,omitempty"`

	// A reference to a resource, canonical resource, or non-FHIR resource which the comment or assessment is about
	ArtifactCanonical *string `json:"artifactCanonical,omitempty"`

	// A reference to a resource, canonical resource, or non-FHIR resource which the comment or assessment is about
	ArtifactUri *string `json:"artifactUri,omitempty"`

	// Display of or reference to the bibliographic citation of the comment, classifier, or rating
	CiteAsReference *common.Reference `json:"citeAsReference,omitempty"`

	// Display of or reference to the bibliographic citation of the comment, classifier, or rating
	CiteAsMarkdown *string `json:"citeAsMarkdown,omitempty"`

	// A component comment, classifier, or rating of the artifact
	Content []ArtifactAssessmentContent `json:"content,omitempty"`

	// A copyright statement relating to the artifact assessment and/or its contents
	Copyright *string `json:"copyright,omitempty"`
}
