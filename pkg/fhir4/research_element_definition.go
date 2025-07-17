package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// ResearchElementDefinition represents the ResearchElementDefinition resource describes a "PICO" element that knowledge is about
type ResearchElementDefinition struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "ResearchElementDefinition"

	// The 'date' element may be more recent than the approval date because of minor changes or editorial corrections
	ApprovalDate *string `json:"approvalDate,omitempty"`

	// An individiual or organization primarily involved in the creation and maintenance of the content
	Author []common.ContactDetail `json:"author,omitempty"`

	// Characteristics can be defined flexibly to accommodate different use cases for membership criteria
	Characteristic []ResearchElementDefinitionCharacteristic `json:"characteristic"`

	// A human-readable string to clarify or explain concepts about the resource
	Comment []string `json:"comment,omitempty"`

	// May be a web site, an email address, a telephone number, etc
	Contact []common.ContactDetail `json:"contact,omitempty"`

	// A copyright statement relating to the research element definition and/or its contents
	Copyright *string `json:"copyright,omitempty"`

	// Note that this is not the same as the resource last-modified-date
	Date *string `json:"date,omitempty"`

	// This description can be used to capture details such as why the research element definition was built
	Description *string `json:"description,omitempty"`

	// An individual or organization primarily responsible for internal coherence of the content
	Editor []common.ContactDetail `json:"editor,omitempty"`

	// The effective period for a research element definition determines when the content is applicable for usage
	EffectivePeriod *common.Period `json:"effectivePeriod,omitempty"`

	// An individual or organization responsible for officially endorsing the content for use in some setting
	Endorser []common.ContactDetail `json:"endorser,omitempty"`

	// Allows filtering of research element definitions that are appropriate for use versus not
	Experimental *bool `json:"experimental,omitempty"`

	// Typically, this is used for identifiers that can go in an HL7 V3 II (instance identifier) data type
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// It may be possible for the research element definition to be used in jurisdictions other than those for which it was originally designed
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`

	// If specified, this date follows the original approval date
	LastReviewDate *string `json:"lastReviewDate,omitempty"`

	// A reference to a Library resource containing the formal logic used by the ResearchElementDefinition
	Library []string `json:"library,omitempty"`

	// The name is not expected to be globally unique
	Name *string `json:"name,omitempty"`

	// Usually an organization but may be an individual
	Publisher *string `json:"publisher,omitempty"`

	// This element does not describe the usage of the research element definition
	Purpose *string `json:"purpose,omitempty"`

	// Each related artifact is either an attachment, or a reference to another resource, but not both
	RelatedArtifact []common.RelatedArtifact `json:"relatedArtifact,omitempty"`

	// An individual or organization primarily responsible for review of some aspect of the content
	Reviewer []common.ContactDetail `json:"reviewer,omitempty"`

	// The short title provides an alternate title for use in informal descriptive contexts
	ShortTitle *string `json:"shortTitle,omitempty"`

	// Allows filtering of research element definitions that are appropriate for use versus not
	Status ResearchElementDefinitionStatus `json:"status"`

	// The subject of the ResearchElementDefinition is critical in interpreting the criteria definitions
	SubjectCodeableConcept *common.CodeableConcept `json:"subjectCodeableConcept,omitempty"`

	// The subject of the ResearchElementDefinition is critical in interpreting the criteria definitions
	SubjectReference *common.Reference `json:"subjectReference,omitempty"`

	// An explanatory or alternate title for the ResearchElementDefinition giving additional information about its content
	Subtitle *string `json:"subtitle,omitempty"`

	// This name does not need to be machine-processing friendly and may contain punctuation, white-space, etc
	Title *string `json:"title,omitempty"`

	// Descriptive topics related to the content of the ResearchElementDefinition
	Topic []common.CodeableConcept `json:"topic,omitempty"`

	// The type of research element, a population, an exposure, or an outcome
	Type ResearchElementDefinitionType `json:"type"`

	// Can be a urn:uuid: or a urn:oid: but real http: addresses are preferred
	Url *string `json:"url,omitempty"`

	// A detailed description, from a clinical perspective, of how the ResearchElementDefinition is used
	Usage *string `json:"usage,omitempty"`

	// When multiple useContexts are specified, there is no expectation that all or any of the contexts apply
	UseContext []common.UsageContext `json:"useContext,omitempty"`

	// The type of the outcome (e.g. Dichotomous, Continuous, or Descriptive)
	VariableType *ResearchElementDefinitionVariableType `json:"variableType,omitempty"`

	// There may be different research element definition instances that have the same identifier but different versions
	Version *string `json:"version,omitempty"`
}

// ResearchElementDefinitionStatus represents allows filtering of research element definitions that are appropriate for use versus not
type ResearchElementDefinitionStatus string

const (
	ResearchElementDefinitionStatusDraft   ResearchElementDefinitionStatus = "draft"
	ResearchElementDefinitionStatusActive  ResearchElementDefinitionStatus = "active"
	ResearchElementDefinitionStatusRetired ResearchElementDefinitionStatus = "retired"
	ResearchElementDefinitionStatusUnknown ResearchElementDefinitionStatus = "unknown"
)

// ResearchElementDefinitionType represents the type of research element, a population, an exposure, or an outcome
type ResearchElementDefinitionType string

const (
	ResearchElementDefinitionTypePopulation ResearchElementDefinitionType = "population"
	ResearchElementDefinitionTypeExposure   ResearchElementDefinitionType = "exposure"
	ResearchElementDefinitionTypeOutcome    ResearchElementDefinitionType = "outcome"
)

// ResearchElementDefinitionVariableType represents the type of the outcome (e.g. Dichotomous, Continuous, or Descriptive)
type ResearchElementDefinitionVariableType string

const (
	ResearchElementDefinitionVariableTypeDichotomous ResearchElementDefinitionVariableType = "dichotomous"
	ResearchElementDefinitionVariableTypeContinuous  ResearchElementDefinitionVariableType = "continuous"
	ResearchElementDefinitionVariableTypeDescriptive ResearchElementDefinitionVariableType = "descriptive"
)

// ResearchElementDefinitionCharacteristic represents characteristics can be defined flexibly to accommodate different use cases for membership criteria
type ResearchElementDefinitionCharacteristic struct {
	common.BackboneElement

	// Define members of the research element using Codes (such as condition, medication, or observation)
	DefinitionCodeableConcept *common.CodeableConcept `json:"definitionCodeableConcept,omitempty"`

	// Define members of the research element using Codes (such as condition, medication, or observation)
	DefinitionCanonical *string `json:"definitionCanonical,omitempty"`

	// Define members of the research element using Codes (such as condition, medication, or observation)
	DefinitionExpression *common.Expression `json:"definitionExpression,omitempty"`

	// Define members of the research element using Codes (such as condition, medication, or observation)
	DefinitionDataRequirement *common.DataRequirement `json:"definitionDataRequirement,omitempty"`

	// When true, members with this characteristic are excluded from the element
	Exclude *bool `json:"exclude,omitempty"`

	// Indicates what effective period the study covers
	ParticipantEffectiveDateTime *string `json:"participantEffectiveDateTime,omitempty"`

	// Indicates what effective period the study covers
	ParticipantEffectivePeriod *common.Period `json:"participantEffectivePeriod,omitempty"`

	// Indicates what effective period the study covers
	ParticipantEffectiveDuration *common.Duration `json:"participantEffectiveDuration,omitempty"`

	// Indicates what effective period the study covers
	ParticipantEffectiveTiming *common.Timing `json:"participantEffectiveTiming,omitempty"`

	// A narrative description of the time period the study covers
	ParticipantEffectiveDescription *string `json:"participantEffectiveDescription,omitempty"`

	// Indicates how elements are aggregated within the study effective period
	ParticipantEffectiveGroupMeasure *ResearchElementDefinitionCharacteristicGroupMeasure `json:"participantEffectiveGroupMeasure,omitempty"`

	// Indicates duration from the participant's study entry
	ParticipantEffectiveTimeFromStart *common.Duration `json:"participantEffectiveTimeFromStart,omitempty"`

	// Indicates what effective period the study covers
	StudyEffectiveDateTime *string `json:"studyEffectiveDateTime,omitempty"`

	// Indicates what effective period the study covers
	StudyEffectivePeriod *common.Period `json:"studyEffectivePeriod,omitempty"`

	// Indicates what effective period the study covers
	StudyEffectiveDuration *common.Duration `json:"studyEffectiveDuration,omitempty"`

	// Indicates what effective period the study covers
	StudyEffectiveTiming *common.Timing `json:"studyEffectiveTiming,omitempty"`

	// A narrative description of the time period the study covers
	StudyEffectiveDescription *string `json:"studyEffectiveDescription,omitempty"`

	// Indicates how elements are aggregated within the study effective period
	StudyEffectiveGroupMeasure *ResearchElementDefinitionCharacteristicGroupMeasure `json:"studyEffectiveGroupMeasure,omitempty"`

	// Indicates duration from the study initiation
	StudyEffectiveTimeFromStart *common.Duration `json:"studyEffectiveTimeFromStart,omitempty"`

	// Specifies the UCUM unit for the outcome
	UnitOfMeasure *common.CodeableConcept `json:"unitOfMeasure,omitempty"`

	// Use UsageContext to define the members of the population, such as Age Ranges, Genders, Settings
	UsageContext []common.UsageContext `json:"usageContext,omitempty"`
}

// ResearchElementDefinitionCharacteristicGroupMeasure represents indicates how elements are aggregated within the study effective period
type ResearchElementDefinitionCharacteristicGroupMeasure string

const (
	ResearchElementDefinitionCharacteristicGroupMeasureMean           ResearchElementDefinitionCharacteristicGroupMeasure = "mean"
	ResearchElementDefinitionCharacteristicGroupMeasureMedian         ResearchElementDefinitionCharacteristicGroupMeasure = "median"
	ResearchElementDefinitionCharacteristicGroupMeasureMeanOfMean     ResearchElementDefinitionCharacteristicGroupMeasure = "mean-of-mean"
	ResearchElementDefinitionCharacteristicGroupMeasureMeanOfMedian   ResearchElementDefinitionCharacteristicGroupMeasure = "mean-of-median"
	ResearchElementDefinitionCharacteristicGroupMeasureMedianOfMean   ResearchElementDefinitionCharacteristicGroupMeasure = "median-of-mean"
	ResearchElementDefinitionCharacteristicGroupMeasureMedianOfMedian ResearchElementDefinitionCharacteristicGroupMeasure = "median-of-median"
)
