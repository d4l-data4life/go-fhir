package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// Measure represents the Measure resource provides the definition of a quality measure
type Measure struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Measure"

	// The 'date' element may be more recent than the approval date because of minor changes or editorial corrections
	ApprovalDate *string `json:"approvalDate,omitempty"`

	// An individual or organization primarily involved in the creation and maintenance of the content
	Author []common.ContactDetail `json:"author,omitempty"`

	// Provides a summary of relevant clinical guidelines or other clinical recommendations supporting the measure
	ClinicalRecommendationStatement *string `json:"clinicalRecommendationStatement,omitempty"`

	// If this is a composite measure, the scoring method used to combine the component measures
	CompositeScoring *common.CodeableConcept `json:"compositeScoring,omitempty"`

	// May be a web site, an email address, a telephone number, etc
	Contact []common.ContactDetail `json:"contact,omitempty"`

	// A copyright statement relating to the measure and/or its contents
	Copyright *string `json:"copyright,omitempty"`

	// Note that this is not the same as the resource last-modified-date
	Date *string `json:"date,omitempty"`

	// Provides a description of an individual term used within the measure
	Definition []string `json:"definition,omitempty"`

	// This description can be used to capture details such as why the measure was built
	Description *string `json:"description,omitempty"`

	// Notices and disclaimers regarding the use of the measure or related to intellectual property
	Disclaimer *string `json:"disclaimer,omitempty"`

	// An individual or organization primarily responsible for internal coherence of the content
	Editor []common.ContactDetail `json:"editor,omitempty"`

	// The effective period for a measure determines when the content is applicable for usage
	EffectivePeriod *common.Period `json:"effectivePeriod,omitempty"`

	// An individual or organization responsible for officially endorsing the content for use in some setting
	Endorser []common.ContactDetail `json:"endorser,omitempty"`

	// Allows filtering of measures that are appropriate for use versus not
	Experimental *bool `json:"experimental,omitempty"`

	// A group of population criteria for the measure
	Group []MeasureGroup `json:"group,omitempty"`

	// Additional guidance for the measure including how it can be used in a clinical context
	Guidance *string `json:"guidance,omitempty"`

	// Typically, this is used for identifiers that can go in an HL7 V3 II (instance identifier) data type
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Information on whether an increase or decrease in score is the preferred result
	ImprovementNotation *common.CodeableConcept `json:"improvementNotation,omitempty"`

	// It may be possible for the measure to be used in jurisdictions other than those for which it was originally designed
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`

	// If specified, this date follows the original approval date
	LastReviewDate *string `json:"lastReviewDate,omitempty"`

	// A reference to a Library resource containing the formal logic used by the measure
	Library []string `json:"library,omitempty"`

	// The name is not expected to be globally unique
	Name *string `json:"name,omitempty"`

	// Usually an organization but may be an individual
	Publisher *string `json:"publisher,omitempty"`

	// This element does not describe the usage of the measure
	Purpose *string `json:"purpose,omitempty"`

	// The measure rate for an organization or clinician is based upon the entities' aggregate data
	RateAggregation *string `json:"rateAggregation,omitempty"`

	// Provides a succinct statement of the need for the measure
	Rationale *string `json:"rationale,omitempty"`

	// Each related artifact is either an attachment, or a reference to another resource, but not both
	RelatedArtifact []common.RelatedArtifact `json:"relatedArtifact,omitempty"`

	// An individual or organization primarily responsible for review of some aspect of the content
	Reviewer []common.ContactDetail `json:"reviewer,omitempty"`

	// Describes the method of adjusting for clinical severity and conditions present at the start of care
	RiskAdjustment *string `json:"riskAdjustment,omitempty"`

	// Indicates how the calculation is performed for the measure
	Scoring *common.CodeableConcept `json:"scoring,omitempty"`

	// Allows filtering of measures that are appropriate for use versus not
	Status MeasureStatus `json:"status"`

	// The subject of the measure is critical in interpreting the criteria definitions
	SubjectCodeableConcept *common.CodeableConcept `json:"subjectCodeableConcept,omitempty"`

	// The subject of the measure is critical in interpreting the criteria definitions
	SubjectReference *common.Reference `json:"subjectReference,omitempty"`

	// An explanatory or alternate title for the measure giving additional information about its content
	Subtitle *string `json:"subtitle,omitempty"`

	// Note that supplemental data are reported as observations for each patient
	SupplementalData []MeasureSupplementalData `json:"supplementalData,omitempty"`

	// This name does not need to be machine-processing friendly and may contain punctuation, white-space, etc
	Title *string `json:"title,omitempty"`

	// Descriptive topics related to the content of the measure
	Topic []common.CodeableConcept `json:"topic,omitempty"`

	// Indicates whether the measure is used to examine a process, an outcome over time, a patient-reported outcome, or a structure measure
	Type []common.CodeableConcept `json:"type,omitempty"`

	// Can be a urn:uuid: or a urn:oid: but real http: addresses are preferred
	Url *string `json:"url,omitempty"`

	// A detailed description, from a clinical perspective, of how the measure is used
	Usage *string `json:"usage,omitempty"`

	// When multiple useContexts are specified, there is no expectation that all or any of the contexts apply
	UseContext []common.UsageContext `json:"useContext,omitempty"`

	// There may be different measure instances that have the same identifier but different versions
	Version *string `json:"version,omitempty"`
}

// MeasureStatus represents the status of the measure
type MeasureStatus string

const (
	MeasureStatusDraft   MeasureStatus = "draft"
	MeasureStatusActive  MeasureStatus = "active"
	MeasureStatusRetired MeasureStatus = "retired"
	MeasureStatusUnknown MeasureStatus = "unknown"
)

// MeasureGroup represents a group of population criteria for the measure
type MeasureGroup struct {
	common.BackboneElement

	// A unique identifier for the group. This identifier will be used to report data for the group in the measure report
	Code *common.CodeableConcept `json:"code,omitempty"`

	// The human readable description of this population group
	Description *string `json:"description,omitempty"`

	// The populations that make up the population group, one for each type of population appropriate for the measure
	Population []MeasureGroupPopulation `json:"population,omitempty"`

	// The stratifier criteria for the measure report, specified as either the name of a valid CQL expression defined within a referenced library, or a valid FHIR Resource Path
	Stratifier []MeasureGroupStratifier `json:"stratifier,omitempty"`
}

// MeasureGroupPopulation represents a population criteria for the measure
type MeasureGroupPopulation struct {
	common.BackboneElement

	// The type of population criteria
	Code *common.CodeableConcept `json:"code,omitempty"`

	// In the case of a continuous-variable or ratio measure, this may be the name of a function that calculates the value
	Criteria common.Expression `json:"criteria"`

	// The human readable description of this population criteria
	Description *string `json:"description,omitempty"`
}

// MeasureGroupStratifier represents the stratifier criteria for the measure report
type MeasureGroupStratifier struct {
	common.BackboneElement

	// An expression that specifies the criteria for the stratifier
	Criteria *common.Expression `json:"criteria,omitempty"`

	// The human readable description of this stratifier
	Description *string `json:"description,omitempty"`

	// The identifier for the stratifier used to coordinate the reported data back to this stratifier
	Identifier *common.Identifier `json:"identifier,omitempty"`

	// The path to an element that defines the stratifier, specified as a valid FHIR resource path
	Path *string `json:"path,omitempty"`
}

// MeasureSupplementalData represents the supplemental data criteria for the measure report
type MeasureSupplementalData struct {
	common.BackboneElement

	// The criteria for the supplemental data
	Criteria common.Expression `json:"criteria"`

	// The human readable description of this supplemental data
	Description *string `json:"description,omitempty"`

	// An identifier for the supplemental data
	Identifier *common.Identifier `json:"identifier,omitempty"`

	// Indicates a meaning for the supplemental data
	Usage []common.CodeableConcept `json:"usage,omitempty"`
}
