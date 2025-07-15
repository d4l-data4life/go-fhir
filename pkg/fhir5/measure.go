package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// MeasureStatus represents the status of a measure
type MeasureStatus string

const (
	MeasureStatusDraft   MeasureStatus = "draft"
	MeasureStatusActive  MeasureStatus = "active"
	MeasureStatusRetired MeasureStatus = "retired"
	MeasureStatusUnknown MeasureStatus = "unknown"
)

// MeasureTerm represents a term used within the measure
type MeasureTerm struct {
	common.BackboneElement

	// A codeable representation of the defined term
	Code *common.CodeableConcept `json:"code,omitempty"`

	// Provides a definition for the term as used within the measure
	Definition *string `json:"definition,omitempty"`
}

// MeasureGroupPopulation represents a population in a measure group
type MeasureGroupPopulation struct {
	common.BackboneElement

	// The code for the population criteria
	Code *common.CodeableConcept `json:"code,omitempty"`

	// The human readable description of this population criteria
	Description *string `json:"description,omitempty"`

	// An expression that specifies the criteria for this population
	Criteria interface{} `json:"criteria"`

	// A unique identifier for the population criteria
	LinkId *string `json:"linkId,omitempty"`
}

// MeasureGroupStratifierComponent represents a component of a stratifier
type MeasureGroupStratifierComponent struct {
	common.BackboneElement

	// The code for the stratifier component
	Code *common.CodeableConcept `json:"code,omitempty"`

	// The expression that is used to define the stratifier criteria
	Criteria interface{} `json:"criteria"`

	// The human readable description of this stratifier component
	Description *string `json:"description,omitempty"`

	// A unique identifier for the stratifier component
	LinkId *string `json:"linkId,omitempty"`
}

// MeasureGroupStratifier represents a stratifier for a measure group
type MeasureGroupStratifier struct {
	common.BackboneElement

	// The code for this stratifier
	Code *common.CodeableConcept `json:"code,omitempty"`

	// Indicates a meaning for the stratifier component
	Component []MeasureGroupStratifierComponent `json:"component,omitempty"`

	// The expression that is used to define the stratifier criteria
	Criteria *interface{} `json:"criteria,omitempty"`

	// The human readable description of this stratifier
	Description *string `json:"description,omitempty"`

	// A unique identifier for the stratifier
	LinkId *string `json:"linkId,omitempty"`
}

// MeasureGroup represents a group of populations for a measure
type MeasureGroup struct {
	common.BackboneElement

	// Meaning of the group
	Code *common.CodeableConcept `json:"code,omitempty"`

	// The human readable description of this population group
	Description *string `json:"description,omitempty"`

	// A unique identifier for the group
	LinkId *string `json:"linkId,omitempty"`

	// A population criteria for the measure
	Population []MeasureGroupPopulation `json:"population,omitempty"`

	// The stratifier criteria for the measure report
	Stratifier []MeasureGroupStratifier `json:"stratifier,omitempty"`
}

// MeasureSupplementalData represents supplemental data for a measure
type MeasureSupplementalData struct {
	common.BackboneElement

	// The code for the supplemental data element
	Code *common.CodeableConcept `json:"code,omitempty"`

	// The expression that is used to define the supplemental data
	Criteria interface{} `json:"criteria"`

	// The human readable description of this supplemental data
	Description *string `json:"description,omitempty"`

	// A unique identifier for the supplemental data element
	LinkId *string `json:"linkId,omitempty"`

	// Indicates how the calculation is performed for the measure
	Usage []common.CodeableConcept `json:"usage,omitempty"`
}

// Measure represents the definition of a quality measure
type Measure struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Measure"

	// The 'date' element may be more recent than the approval date
	ApprovalDate *string `json:"approvalDate,omitempty"`

	// An individual or organization primarily involved in the creation and maintenance of the content
	Author []ContactDetail `json:"author,omitempty"`

	// For a subject-based measure, the population basis is simply boolean
	Basis *string `json:"basis,omitempty"`

	// Provides a summary of relevant clinical guidelines or other clinical recommendations
	ClinicalRecommendationStatement *string `json:"clinicalRecommendationStatement,omitempty"`

	// If this is a composite measure, the scoring method used to combine the component measures
	CompositeScoring *common.CodeableConcept `json:"compositeScoring,omitempty"`

	// May be a web site, an email address, a telephone number, etc.
	Contact []ContactDetail `json:"contact,omitempty"`

	// A copyright statement relating to the measure and/or its contents
	Copyright *string `json:"copyright,omitempty"`

	// Copyright holder and year(s)
	CopyrightLabel *string `json:"copyrightLabel,omitempty"`

	// Note that this is not the same as the resource last-modified-date
	Date *string `json:"date,omitempty"`

	// Natural language description of the measure
	Description *string `json:"description,omitempty"`

	// Provides a summary of key clinical recommendation
	Disclaimer *string `json:"disclaimer,omitempty"`

	// An individual or organization primarily responsible for internal coherence of the content
	Editor []ContactDetail `json:"editor,omitempty"`

	// When the measure is expected to be used
	EffectivePeriod *common.Period `json:"effectivePeriod,omitempty"`

	// Who endorsed the content
	Endorser []ContactDetail `json:"endorser,omitempty"`

	// For testing purposes, not real usage
	Experimental *bool `json:"experimental,omitempty"`

	// Additional guidance for the measure including how it can be used in a clinical context
	Guidance *string `json:"guidance,omitempty"`

	// Population criteria group
	Group []MeasureGroup `json:"group,omitempty"`

	// Additional identifier for the measure
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Information about whether an improvement in the measure is noted by an increase or decrease
	ImprovementNotation *common.CodeableConcept `json:"improvementNotation,omitempty"`

	// Intended jurisdiction for measure (if applicable)
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`

	// When the measure was last reviewed
	LastReviewDate *string `json:"lastReviewDate,omitempty"`

	// Logic used by the measure
	Library []string `json:"library,omitempty"`

	// Name for this measure (computer friendly)
	Name *string `json:"name,omitempty"`

	// Natural language description of the purpose of the measure
	Purpose *string `json:"purpose,omitempty"`

	// Additional documentation, citations, etc.
	RelatedArtifact []interface{} `json:"relatedArtifact,omitempty"`

	// Who reviewed the content
	Reviewer []ContactDetail `json:"reviewer,omitempty"`

	// draft | active | retired | unknown
	Status MeasureStatus `json:"status"`

	// The subject for which the measure is intended
	SubjectCodeableConcept *common.CodeableConcept `json:"subjectCodeableConcept,omitempty"`
	SubjectReference       *common.Reference       `json:"subjectReference,omitempty"`

	// Subordinate title of the measure
	Subtitle *string `json:"subtitle,omitempty"`

	// Name for this measure (human friendly)
	Title *string `json:"title,omitempty"`

	// The category of the measure, such as process, outcome, structure, patient-reported outcome, etc.
	Topic []common.CodeableConcept `json:"topic,omitempty"`

	// Identifies the type of measure such as a process measure, outcome measure, etc.
	Type []common.CodeableConcept `json:"type,omitempty"`

	// Canonical identifier for this measure, represented as a URI
	Url *string `json:"url,omitempty"`

	// Describes the clinical usage of the measure
	Usage *string `json:"usage,omitempty"`

	// The context that the content is intended to support
	UseContext []interface{} `json:"useContext,omitempty"`

	// Business version of the measure
	Version *string `json:"version,omitempty"`

	// How to compare versions
	VersionAlgorithmString *string        `json:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *common.Coding `json:"versionAlgorithmCoding,omitempty"`
}
