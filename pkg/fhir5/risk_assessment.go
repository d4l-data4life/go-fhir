// Package fhir5 contains FHIR R5 resource definitions
package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// RiskAssessmentPrediction represents a prediction in a risk assessment
type RiskAssessmentPrediction struct {
	common.BackboneElement

	// One of the potential outcomes for the patient (e.g. remission, death, a particular condition)
	Outcome *common.CodeableConcept `json:"outcome,omitempty"`

	// If range is used, it represents the lower and upper bounds of certainty; e.g. 40-60%
	ProbabilityDecimal *float64 `json:"probabilityDecimal,omitempty"`
	ProbabilityRange   *Range   `json:"probabilityRange,omitempty"`

	// Indicates how likely the outcome is (in the specified timeframe), expressed as a qualitative value
	QualitativeRisk *common.CodeableConcept `json:"qualitativeRisk,omitempty"`

	// Additional information explaining the basis for the prediction
	Rationale *string `json:"rationale,omitempty"`

	// Indicates the risk for this particular subject divided by the risk of the population in general
	RelativeRisk *float64 `json:"relativeRisk,omitempty"`

	// If not specified, the risk applies "over the subject's lifespan"
	WhenPeriod *common.Period `json:"whenPeriod,omitempty"`
	WhenRange  *Range         `json:"whenRange,omitempty"`
}

// RiskAssessment represents an assessment of the likely outcome(s) for a patient or other subject
type RiskAssessment struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "RiskAssessment"

	// A reference to the request that is fulfilled by this risk assessment
	BasedOn *common.Reference `json:"basedOn,omitempty"`

	// Indicates the source data considered as part of the assessment
	Basis []common.Reference `json:"basis,omitempty"`

	// The type of the risk assessment performed
	Code *common.CodeableConcept `json:"code,omitempty"`

	// For assessments or prognosis specific to a particular condition, indicates the condition being assessed
	Condition *common.Reference `json:"condition,omitempty"`

	// The encounter where the assessment was performed
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Business identifier assigned to the risk assessment
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The algorithm, process or mechanism used to evaluate the risk
	Method *common.CodeableConcept `json:"method,omitempty"`

	// A description of the steps that might be taken to reduce the identified risk(s)
	Mitigation *string `json:"mitigation,omitempty"`

	// Additional comments about the risk assessment
	Note []Annotation `json:"note,omitempty"`

	// The date (and possibly time) the risk assessment was performed
	OccurrenceDateTime *string        `json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod   *common.Period `json:"occurrencePeriod,omitempty"`

	// A reference to a resource that this risk assessment is part of, such as a Procedure
	Parent *common.Reference `json:"parent,omitempty"`

	// The provider, patient, related person, or software application that performed the assessment
	Performer *common.Reference `json:"performer,omitempty"`

	// Multiple repetitions can be used to identify the same type of outcome in different timeframes
	Prediction []RiskAssessmentPrediction `json:"prediction,omitempty"`

	// The reason the risk assessment was performed
	Reason []CodeableReference `json:"reason,omitempty"`

	// The status of the RiskAssessment, using the same statuses as an Observation
	Status string `json:"status"`

	// The patient or group the risk assessment applies to
	Subject common.Reference `json:"subject"`
}
