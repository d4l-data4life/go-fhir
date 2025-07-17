package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// ImmunizationEvaluation represents the evaluation of a vaccine administration event
type ImmunizationEvaluation struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "ImmunizationEvaluation"

	// Indicates the authority who published the protocol (e.g. ACIP)
	Authority *common.Reference `json:"authority,omitempty"`

	// The date the evaluation of the vaccine administration event was performed
	Date *string `json:"date,omitempty"`

	// Additional information about the evaluation
	Description *string `json:"description,omitempty"`

	// The use of an integer is preferred if known
	DoseNumberPositiveInt *int `json:"doseNumberPositiveInt,omitempty"`

	// The use of an integer is preferred if known
	DoseNumberString *string `json:"doseNumberString,omitempty"`

	// Indicates if the dose is valid or not valid with respect to the published recommendations
	DoseStatus common.CodeableConcept `json:"doseStatus"`

	// Provides an explanation as to why the vaccine administration event is valid or not relative to the published recommendations
	DoseStatusReason []common.CodeableConcept `json:"doseStatusReason,omitempty"`

	// A unique identifier assigned to this immunization evaluation record
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The vaccine administration event being evaluated
	ImmunizationEvent common.Reference `json:"immunizationEvent"`

	// The individual for whom the evaluation is being done
	Patient common.Reference `json:"patient"`

	// One possible path to achieve presumed immunity against a disease - within the context of an authority
	Series *string `json:"series,omitempty"`

	// The use of an integer is preferred if known
	SeriesDosesPositiveInt *int `json:"seriesDosesPositiveInt,omitempty"`

	// The use of an integer is preferred if known
	SeriesDosesString *string `json:"seriesDosesString,omitempty"`

	// Indicates the current status of the evaluation of the vaccination administration event
	Status ImmunizationEvaluationStatus `json:"status"`

	// The vaccine preventable disease the dose is being evaluated against
	TargetDisease common.CodeableConcept `json:"targetDisease"`
}

// ImmunizationEvaluationStatus represents the status of the immunization evaluation
type ImmunizationEvaluationStatus string

const (
	ImmunizationEvaluationStatusCompleted      ImmunizationEvaluationStatus = "completed"
	ImmunizationEvaluationStatusEnteredInError ImmunizationEvaluationStatus = "entered-in-error"
)
