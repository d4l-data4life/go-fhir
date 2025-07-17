package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// EvidenceVariable represents the EvidenceVariable resource
type EvidenceVariable struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "EvidenceVariable"

	// The 'date' element may be more recent than the approval date because of minor changes or editorial corrections
	ApprovalDate *string `json:"approvalDate,omitempty"`

	// An individual or organization primarily involved in the creation and maintenance of the content
	Author []common.ContactDetail `json:"author,omitempty"`

	// Characteristics can be defined flexibly to accommodate different use cases for membership criteria
	Characteristic []EvidenceVariableCharacteristic `json:"characteristic"`

	// May be a web site, an email address, a telephone number, etc
	Contact []common.ContactDetail `json:"contact,omitempty"`

	// A copyright statement relating to the evidence variable and/or its contents
	Copyright *string `json:"copyright,omitempty"`

	// Note that this is not the same as the resource last-modified-date
	Date *string `json:"date,omitempty"`

	// This description can be used to capture details such as why the evidence variable was built
	Description *string `json:"description,omitempty"`

	// An individual or organization primarily responsible for internal coherence of the content
	Editor []common.ContactDetail `json:"editor,omitempty"`

	// The effective period for a evidence variable determines when the content is applicable for usage
	EffectivePeriod *common.Period `json:"effectivePeriod,omitempty"`

	// An individual or organization responsible for officially endorsing the content for use in some setting
	Endorser []common.ContactDetail `json:"endorser,omitempty"`

	// Typically, this is used for identifiers that can go in an HL7 V3 II (instance identifier) data type
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// It may be possible for the evidence variable to be used in jurisdictions other than those for which it was originally designed or intended
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`

	// If specified, this date follows the original approval date
	LastReviewDate *string `json:"lastReviewDate,omitempty"`

	// The name is not expected to be globally unique
	Name *string `json:"name,omitempty"`

	// A human-readable string to clarify or explain concepts about the resource
	Note []common.Annotation `json:"note,omitempty"`

	// Usually an organization but may be an individual
	Publisher *string `json:"publisher,omitempty"`

	// Each related artifact is either an attachment, or a reference to another resource, but not both
	RelatedArtifact []common.RelatedArtifact `json:"relatedArtifact,omitempty"`

	// An individual or organization primarily responsible for review of some aspect of the content
	Reviewer []common.ContactDetail `json:"reviewer,omitempty"`

	// The short title provides an alternate title for use in informal descriptive contexts
	ShortTitle *string `json:"shortTitle,omitempty"`

	// Allows filtering of evidence variables that are appropriate for use versus not
	Status EvidenceVariableStatus `json:"status"`

	// An explanatory or alternate title for the EvidenceVariable giving additional information about its content
	Subtitle *string `json:"subtitle,omitempty"`

	// This name does not need to be machine-processing friendly and may contain punctuation, white-space, etc
	Title *string `json:"title,omitempty"`

	// Descriptive topics related to the content of the EvidenceVariable
	Topic []common.CodeableConcept `json:"topic,omitempty"`

	// The type of evidence element, a population, an exposure, or an outcome
	Type *EvidenceVariableType `json:"type,omitempty"`

	// Can be a urn:uuid: or a urn:oid: but real http: addresses are preferred
	Url *string `json:"url,omitempty"`

	// When multiple useContexts are specified, there is no expectation that all or any of the contexts apply
	UseContext []common.UsageContext `json:"useContext,omitempty"`

	// There may be different evidence variable instances that have the same identifier but different versions
	Version *string `json:"version,omitempty"`
}

// EvidenceVariableStatus represents the status of the evidence variable
type EvidenceVariableStatus string

const (
	EvidenceVariableStatusDraft   EvidenceVariableStatus = "draft"
	EvidenceVariableStatusActive  EvidenceVariableStatus = "active"
	EvidenceVariableStatusRetired EvidenceVariableStatus = "retired"
	EvidenceVariableStatusUnknown EvidenceVariableStatus = "unknown"
)

// EvidenceVariableType represents the type of evidence element
type EvidenceVariableType string

const (
	EvidenceVariableTypeDichotomous EvidenceVariableType = "dichotomous"
	EvidenceVariableTypeContinuous  EvidenceVariableType = "continuous"
	EvidenceVariableTypeDescriptive EvidenceVariableType = "descriptive"
)

// EvidenceVariableCharacteristic represents characteristics that can be defined flexibly
type EvidenceVariableCharacteristic struct {
	common.BackboneElement

	// Define members of the evidence element using Codes
	DefinitionReference *common.Reference `json:"definitionReference,omitempty"`

	// Define members of the evidence element using Codes
	DefinitionCanonical *string `json:"definitionCanonical,omitempty"`

	// Define members of the evidence element using Codes
	DefinitionCodeableConcept *common.CodeableConcept `json:"definitionCodeableConcept,omitempty"`

	// Define members of the evidence element using Expressions
	DefinitionExpression *common.Expression `json:"definitionExpression,omitempty"`

	// Define members of the evidence element using DataRequirements
	DefinitionDataRequirement *common.DataRequirement `json:"definitionDataRequirement,omitempty"`

	// Define members of the evidence element using TriggerDefinition
	DefinitionTriggerDefinition *common.TriggerDefinition `json:"definitionTriggerDefinition,omitempty"`

	// A short, natural language description of the characteristic
	Description *string `json:"description,omitempty"`

	// When true, members with this characteristic are excluded from the element
	Exclude *bool `json:"exclude,omitempty"`

	// Indicates how elements are aggregated within the study effective period
	GroupMeasure *EvidenceVariableCharacteristicGroupMeasure `json:"groupMeasure,omitempty"`

	// Indicates what effective period the study covers
	ParticipantEffectiveDateTime *string `json:"participantEffectiveDateTime,omitempty"`

	// Indicates what effective period the study covers
	ParticipantEffectivePeriod *common.Period `json:"participantEffectivePeriod,omitempty"`

	// Indicates what effective period the study covers
	ParticipantEffectiveDuration *common.Duration `json:"participantEffectiveDuration,omitempty"`

	// Indicates what effective period the study covers
	ParticipantEffectiveTiming *common.Timing `json:"participantEffectiveTiming,omitempty"`

	// Indicates duration from the participant's study entry
	TimeFromStart *common.Duration `json:"timeFromStart,omitempty"`

	// Use UsageContext to define the members of the population
	UsageContext []common.UsageContext `json:"usageContext,omitempty"`
}

// EvidenceVariableCharacteristicGroupMeasure represents how elements are aggregated
type EvidenceVariableCharacteristicGroupMeasure string

const (
	EvidenceVariableCharacteristicGroupMeasureMean           EvidenceVariableCharacteristicGroupMeasure = "mean"
	EvidenceVariableCharacteristicGroupMeasureMedian         EvidenceVariableCharacteristicGroupMeasure = "median"
	EvidenceVariableCharacteristicGroupMeasureMeanOfMean     EvidenceVariableCharacteristicGroupMeasure = "mean-of-mean"
	EvidenceVariableCharacteristicGroupMeasureMeanOfMedian   EvidenceVariableCharacteristicGroupMeasure = "mean-of-median"
	EvidenceVariableCharacteristicGroupMeasureMedianOfMean   EvidenceVariableCharacteristicGroupMeasure = "median-of-mean"
	EvidenceVariableCharacteristicGroupMeasureMedianOfMedian EvidenceVariableCharacteristicGroupMeasure = "median-of-median"
)
