// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// Evidence represents the Evidence resource
type Evidence struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Evidence"

	// The 'date' element may be more recent than the approval date
	ApprovalDate *string `json:"approvalDate,omitempty"`

	// Declarative description of the Evidence
	Assertion *string `json:"assertion,omitempty"`

	// An individual or organization primarily involved in the creation and maintenance of the content
	Author []ContactDetail `json:"author,omitempty"`

	// Assessment of certainty, confidence in the estimates, or quality of the evidence
	Certainty []EvidenceCertainty `json:"certainty,omitempty"`

	// Citation Resource or display of suggested citation for this evidence
	CiteAsReference *common.Reference `json:"citeAsReference,omitempty"`
	CiteAsMarkdown  *string           `json:"citeAsMarkdown,omitempty"`

	// May be a web site, an email address, a telephone number, etc
	Contact []ContactDetail `json:"contact,omitempty"`

	// Use and/or publishing restrictions
	Copyright *string `json:"copyright,omitempty"`

	// Copyright holder and year(s)
	CopyrightLabel *string `json:"copyrightLabel,omitempty"`

	// Date last changed
	Date *string `json:"date,omitempty"`

	// Natural language description of the evidence
	Description *string `json:"description,omitempty"`

	// An individual or organization primarily responsible for internal coherence of the content
	Editor []ContactDetail `json:"editor,omitempty"`

	// The effective period for an evidence determines when the content is applicable for usage
	EffectivePeriod *common.Period `json:"effectivePeriod,omitempty"`

	// See guidance around (not) making local changes to elements
	Endorser []ContactDetail `json:"endorser,omitempty"`

	// For testing purposes, not real usage
	Experimental *bool `json:"experimental,omitempty"`

	// Additional identifier for the evidence
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// If specified, this date follows the original approval date
	LastReviewDate *string `json:"lastReviewDate,omitempty"`

	// Name for this evidence (computer friendly)
	Name *string `json:"name,omitempty"`

	// Used for footnotes or explanatory notes
	Note []Annotation `json:"note,omitempty"`

	// Name of the publisher (organization or individual)
	Publisher *string `json:"publisher,omitempty"`

	// Why this evidence is defined
	Purpose *string `json:"purpose,omitempty"`

	// Link or citation to artifact associated with the summary
	RelatedArtifact []RelatedArtifact `json:"relatedArtifact,omitempty"`

	// See guidance around (not) making local changes to elements
	Reviewer []ContactDetail `json:"reviewer,omitempty"`

	// Values and parameters for a single statistic
	Statistic []EvidenceStatistic `json:"statistic,omitempty"`

	// The status of this evidence
	Status EventDefinitionStatus `json:"status"`

	// The design of the study that produced this evidence
	StudyDesign []common.CodeableConcept `json:"studyDesign,omitempty"`

	// The method to combine studies
	SynthesisType *common.CodeableConcept `json:"synthesisType,omitempty"`

	// Name for this evidence (human friendly)
	Title *string `json:"title,omitempty"`

	// Canonical identifier for this evidence
	URL *string `json:"url,omitempty"`

	// The context that the content is intended to support
	UseContext []UsageContext `json:"useContext,omitempty"`

	// Evidence variable such as population, exposure, or outcome
	VariableDefinition []EvidenceVariableDefinition `json:"variableDefinition"`

	// Business version of the evidence
	Version *string `json:"version,omitempty"`

	// If set as a string, this is a FHIRPath expression
	VersionAlgorithmString *string        `json:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *common.Coding `json:"versionAlgorithmCoding,omitempty"`
}

// (Add related types: EvidenceCertainty, EvidenceStatistic, etc. as needed)
