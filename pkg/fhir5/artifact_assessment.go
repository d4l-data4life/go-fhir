// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// ArtifactAssessmentContent represents content of an artifact assessment
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
	InformationType *ArtifactAssessmentInformationType `json:"informationType,omitempty"`

	// The target element is used to point the comment to aspect of the artifact
	Path []string `json:"path,omitempty"`

	// A quantitative rating of the artifact
	Quantity *common.Quantity `json:"quantity,omitempty"`

	// Additional related artifacts that provide supporting documentation
	RelatedArtifact []RelatedArtifact `json:"relatedArtifact,omitempty"`

	// A brief summary of the content of this component
	Summary *string `json:"summary,omitempty"`

	// Indicates what type of content this component represents
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// ArtifactAssessmentInformationType represents the type of information in artifact assessment
type ArtifactAssessmentInformationType string

const (
	ArtifactAssessmentInformationTypeComment       ArtifactAssessmentInformationType = "comment"
	ArtifactAssessmentInformationTypeClassifier    ArtifactAssessmentInformationType = "classifier"
	ArtifactAssessmentInformationTypeRating        ArtifactAssessmentInformationType = "rating"
	ArtifactAssessmentInformationTypeContainer     ArtifactAssessmentInformationType = "container"
	ArtifactAssessmentInformationTypeResponse      ArtifactAssessmentInformationType = "response"
	ArtifactAssessmentInformationTypeChangeRequest ArtifactAssessmentInformationType = "change-request"
)

// ArtifactAssessment represents an assessment of an artifact
type ArtifactAssessment struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "ArtifactAssessment"

	// The 'date' element may be more recent than the approval date
	ApprovalDate *string `json:"approvalDate,omitempty"`

	// A reference to a resource which the comment or assessment is about
	ArtifactReference *common.Reference `json:"artifactReference,omitempty"`

	// A reference to a canonical resource which the comment or assessment is about
	ArtifactCanonical *string `json:"artifactCanonical,omitempty"`

	// A reference to a URI which the comment or assessment is about
	ArtifactUri *string `json:"artifactUri,omitempty"`

	// Display of or reference to the bibliographic citation
	CiteAsReference *common.Reference `json:"citeAsReference,omitempty"`

	// Display of or reference to the bibliographic citation
	CiteAsMarkdown *string `json:"citeAsMarkdown,omitempty"`

	// A component comment, classifier, or rating of the artifact
	Content []ArtifactAssessmentContent `json:"content,omitempty"`

	// A copyright statement relating to the artifact assessment
	Copyright *string `json:"copyright,omitempty"`

	// The date is often not tracked until the resource is published
	Date *string `json:"date,omitempty"`

	// Indicates the disposition of the responsible party to the comment
	Disposition *ArtifactAssessmentDisposition `json:"disposition,omitempty"`

	// Typically, this is used for identifiers that can go in an HL7 V3 II
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// If specified, this date follows the original approval date
	LastReviewDate *string `json:"lastReviewDate,omitempty"`

	// A short title for the assessment for use in displaying and selecting
	Title *string `json:"title,omitempty"`

	// Indicates the workflow status of the comment or change request
	WorkflowStatus *ArtifactAssessmentWorkflowStatus `json:"workflowStatus,omitempty"`
}

// ArtifactAssessmentDisposition represents the disposition of an artifact assessment
type ArtifactAssessmentDisposition string

const (
	ArtifactAssessmentDispositionUnresolved                    ArtifactAssessmentDisposition = "unresolved"
	ArtifactAssessmentDispositionNotPersuasive                 ArtifactAssessmentDisposition = "not-persuasive"
	ArtifactAssessmentDispositionPersuasive                    ArtifactAssessmentDisposition = "persuasive"
	ArtifactAssessmentDispositionPersuasiveWithModification    ArtifactAssessmentDisposition = "persuasive-with-modification"
	ArtifactAssessmentDispositionNotPersuasiveWithModification ArtifactAssessmentDisposition = "not-persuasive-with-modification"
)

// ArtifactAssessmentWorkflowStatus represents the workflow status of an artifact assessment
type ArtifactAssessmentWorkflowStatus string

const (
	ArtifactAssessmentWorkflowStatusSubmitted              ArtifactAssessmentWorkflowStatus = "submitted"
	ArtifactAssessmentWorkflowStatusTriaged                ArtifactAssessmentWorkflowStatus = "triaged"
	ArtifactAssessmentWorkflowStatusWaitingForInput        ArtifactAssessmentWorkflowStatus = "waiting-for-input"
	ArtifactAssessmentWorkflowStatusResolvedNoChange       ArtifactAssessmentWorkflowStatus = "resolved-no-change"
	ArtifactAssessmentWorkflowStatusResolvedChangeRequired ArtifactAssessmentWorkflowStatus = "resolved-change-required"
	ArtifactAssessmentWorkflowStatusDeferred               ArtifactAssessmentWorkflowStatus = "deferred"
	ArtifactAssessmentWorkflowStatusDuplicate              ArtifactAssessmentWorkflowStatus = "duplicate"
	ArtifactAssessmentWorkflowStatusApplied                ArtifactAssessmentWorkflowStatus = "applied"
	ArtifactAssessmentWorkflowStatusPublished              ArtifactAssessmentWorkflowStatus = "published"
	ArtifactAssessmentWorkflowStatusEnteredInError         ArtifactAssessmentWorkflowStatus = "entered-in-error"
)
