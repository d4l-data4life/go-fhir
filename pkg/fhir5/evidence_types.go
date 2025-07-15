// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// EvidenceVariableDefinition represents evidence variable definition
type EvidenceVariableDefinition struct {
	common.BackboneElement

	// A text description or summary of the variable
	Description *string `json:"description,omitempty"`

	// Indication of quality of match between intended variable to actual variable
	DirectnessMatch *common.CodeableConcept `json:"directnessMatch,omitempty"`

	// Definition of the intended variable related to the Evidence
	Intended *common.Reference `json:"intended,omitempty"`

	// Footnotes and/or explanatory notes
	Note []Annotation `json:"note,omitempty"`

	// Definition of the actual variable related to the statistic(s)
	Observed *common.Reference `json:"observed,omitempty"`

	// population | subpopulation | exposure | referenceExposure | measuredVariable | confounder
	VariableRole common.CodeableConcept `json:"variableRole"`
}

// EvidenceStatisticSampleSize represents sample size information
type EvidenceStatisticSampleSize struct {
	common.BackboneElement

	// Human-readable summary of population sample size
	Description *string `json:"description,omitempty"`

	// Number of participants with known results for measured variables
	KnownDataCount *int `json:"knownDataCount,omitempty"`

	// Footnote or explanatory note about the sample size
	Note []Annotation `json:"note,omitempty"`

	// Number of participants in the population
	NumberOfParticipants *int `json:"numberOfParticipants,omitempty"`

	// Number of studies
	NumberOfStudies *int `json:"numberOfStudies,omitempty"`
}

// EvidenceStatisticAttributeEstimate represents attribute estimate information
type EvidenceStatisticAttributeEstimate struct {
	common.BackboneElement

	// A nested attribute estimate
	AttributeEstimate []EvidenceStatisticAttributeEstimate `json:"attributeEstimate,omitempty"`

	// Human-readable summary of the estimate
	Description *string `json:"description,omitempty"`

	// Use 95 for a 95% confidence interval
	Level *float64 `json:"level,omitempty"`

	// Footnote or explanatory note about the estimate
	Note []Annotation `json:"note,omitempty"`

	// Often the p value
	Quantity *common.Quantity `json:"quantity,omitempty"`

	// Lower bound of confidence interval
	Range *Range `json:"range,omitempty"`

	// The type of attribute estimate, e.g., confidence interval or p value
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// EvidenceStatisticModelCharacteristicVariable represents model characteristic variable
type EvidenceStatisticModelCharacteristicVariable struct {
	common.BackboneElement

	// How the variable is classified for use in adjusted analysis
	Handling *EvidenceVariableHandling `json:"handling,omitempty"`

	// Description for grouping of ordinal or polychotomous variables
	ValueCategory []common.CodeableConcept `json:"valueCategory,omitempty"`

	// Discrete value for grouping of ordinal or polychotomous variables
	ValueQuantity []common.Quantity `json:"valueQuantity,omitempty"`

	// Range of values for grouping of ordinal or polychotomous variables
	ValueRange []Range `json:"valueRange,omitempty"`

	// Description of the variable
	VariableDefinition common.Reference `json:"variableDefinition"`
}

// EvidenceVariableHandling represents how variables are handled in statistical analysis
type EvidenceVariableHandling string

const (
	EvidenceVariableHandlingContinuous    EvidenceVariableHandling = "continuous"
	EvidenceVariableHandlingDichotomous   EvidenceVariableHandling = "dichotomous"
	EvidenceVariableHandlingOrdinal       EvidenceVariableHandling = "ordinal"
	EvidenceVariableHandlingPolychotomous EvidenceVariableHandling = "polychotomous"
)

// EvidenceStatisticModelCharacteristic represents model characteristic
type EvidenceStatisticModelCharacteristic struct {
	common.BackboneElement

	// An attribute of the statistic used as a model characteristic
	AttributeEstimate []EvidenceStatisticAttributeEstimate `json:"attributeEstimate,omitempty"`

	// Description of a component of the method to generate the statistic
	Code common.CodeableConcept `json:"code"`

	// Further specification of the quantified value of the component
	Value *common.Quantity `json:"value,omitempty"`

	// A variable adjusted for in the adjusted analysis
	Variable []EvidenceStatisticModelCharacteristicVariable `json:"variable,omitempty"`
}

// EvidenceStatistic represents values and parameters for a single statistic
type EvidenceStatistic struct {
	common.BackboneElement

	// A statistical attribute of the statistic such as a measure of heterogeneity
	AttributeEstimate []EvidenceStatisticAttributeEstimate `json:"attributeEstimate,omitempty"`

	// Simple strings can be used for descriptive purposes
	Category *common.CodeableConcept `json:"category,omitempty"`

	// A description of the content value of the statistic
	Description *string `json:"description,omitempty"`

	// A component of the method to generate the statistic
	ModelCharacteristic []EvidenceStatisticModelCharacteristic `json:"modelCharacteristic,omitempty"`

	// Footnotes and/or explanatory notes
	Note []Annotation `json:"note,omitempty"`

	// Number of participants with events
	NumberAffected *int `json:"numberAffected,omitempty"`

	// Number of events
	NumberOfEvents *int `json:"numberOfEvents,omitempty"`

	// Statistic value
	Quantity *common.Quantity `json:"quantity,omitempty"`

	// Number of samples in the statistic
	SampleSize *EvidenceStatisticSampleSize `json:"sampleSize,omitempty"`

	// Type of statistic, e.g., relative risk
	StatisticType *common.CodeableConcept `json:"statisticType,omitempty"`
}

// EvidenceCertainty represents certainty assessment
type EvidenceCertainty struct {
	common.BackboneElement

	// Textual description of certainty
	Description *string `json:"description,omitempty"`

	// Footnotes and/or explanatory notes
	Note []Annotation `json:"note,omitempty"`

	// Individual or group who did the rating
	Rater *string `json:"rater,omitempty"`

	// Assessment or judgement of the aspect
	Rating *common.CodeableConcept `json:"rating,omitempty"`

	// A domain or subdomain of certainty
	Subcomponent []EvidenceCertainty `json:"subcomponent,omitempty"`

	// Aspect of certainty being rated
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// EvidenceVariableCharacteristicDefinitionByCombination represents definition by combination
type EvidenceVariableCharacteristicDefinitionByCombination struct {
	common.BackboneElement

	// Provide a code for the characteristic
	Code common.CodeableConcept `json:"code"`

	// Describes the characteristic
	Characteristic []EvidenceVariableCharacteristic `json:"characteristic"`
}

// EvidenceVariableCharacteristicDefinitionByTypeAndValue represents definition by type and value
type EvidenceVariableCharacteristicDefinitionByTypeAndValue struct {
	common.BackboneElement

	// Defines the characteristic when paired with characteristic.definitionByTypeAndValue.value[x]
	Type common.CodeableConcept `json:"type"`

	// Defines the characteristic when paired with characteristic.definitionByTypeAndValue.type
	ValueCodeableConcept *common.CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueBoolean         *bool                   `json:"valueBoolean,omitempty"`
	ValueQuantity        *common.Quantity        `json:"valueQuantity,omitempty"`
	ValueRange           *Range                  `json:"valueRange,omitempty"`
	ValueReference       *common.Reference       `json:"valueReference,omitempty"`
	ValueId              *string                 `json:"valueId,omitempty"`
}

// EvidenceVariableCharacteristicTimeFromEvent represents time from event
type EvidenceVariableCharacteristicTimeFromEvent struct {
	common.BackboneElement

	// Human readable description
	Description *string `json:"description,omitempty"`

	// The event used as a base point (reference point) in time
	EventCodeableConcept *common.CodeableConcept `json:"eventCodeableConcept,omitempty"`
	EventReference       *common.Reference       `json:"eventReference,omitempty"`
	EventDateTime        *string                 `json:"eventDateTime,omitempty"`
	EventId              *string                 `json:"eventId,omitempty"`

	// Used to express the observation at a defined number of time units after the defined event
	Quantity *common.Quantity `json:"quantity,omitempty"`
	Range    *Range           `json:"range,omitempty"`
}

// EvidenceVariableCharacteristic represents characteristics of the evidence variable
type EvidenceVariableCharacteristic struct {
	common.BackboneElement

	// Defines the characteristic as a combination of two or more characteristics
	DefinitionByCombination *EvidenceVariableCharacteristicDefinitionByCombination `json:"definitionByCombination,omitempty"`

	// Defines the characteristic using both a type and value[x] elements
	DefinitionByTypeAndValue *EvidenceVariableCharacteristicDefinitionByTypeAndValue `json:"definitionByTypeAndValue,omitempty"`

	// Defines the characteristic using Canonical
	DefinitionCanonical *string `json:"definitionCanonical,omitempty"`

	// Defines the characteristic using CodeableConcept
	DefinitionCodeableConcept *common.CodeableConcept `json:"definitionCodeableConcept,omitempty"`

	// Defines the characteristic using Expression
	DefinitionExpression *Expression `json:"definitionExpression,omitempty"`

	// Defines the characteristic using id
	DefinitionId *string `json:"definitionId,omitempty"`

	// Defines the characteristic using Reference
	DefinitionReference *common.Reference `json:"definitionReference,omitempty"`

	// A short, natural language description of the characteristic
	Description *string `json:"description,omitempty"`

	// Device used for determining characteristic
	Device *common.Reference `json:"device,omitempty"`

	// When true, members with this characteristic are excluded from the element
	Exclude *bool `json:"exclude,omitempty"`

	// The grouping characteristic
	GroupMeasure *common.CodeableConcept `json:"groupMeasure,omitempty"`

	// Defines the characteristic using instances
	InstancesQuantity *common.Quantity `json:"instancesQuantity,omitempty"`
	InstancesRange    *Range           `json:"instancesRange,omitempty"`

	// The intended method of use for the characteristic
	Method []common.CodeableConcept `json:"method,omitempty"`

	// Length of time in which the characteristic is met
	DurationQuantity *common.Quantity `json:"durationQuantity,omitempty"`
	DurationRange    *Range           `json:"durationRange,omitempty"`

	// Timing for the characteristic
	TimeFromEvent []EvidenceVariableCharacteristicTimeFromEvent `json:"timeFromEvent,omitempty"`
}

// EvidenceVariableCategory represents a grouping for ordinal or polychotomous variables
type EvidenceVariableCategory struct {
	common.BackboneElement

	// Description of the grouping
	Name *string `json:"name,omitempty"`

	// Definition of the grouping
	ValueCodeableConcept *common.CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueQuantity        *common.Quantity        `json:"valueQuantity,omitempty"`
	ValueRange           *Range                  `json:"valueRange,omitempty"`
}
