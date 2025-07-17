// Package fhir4 contains FHIR R4 (version 4.0.1) resource definitions
package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// Condition represents detailed information about conditions, problems, or diagnoses
type Condition struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Condition"

	// Identifier for this condition
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The clinical status of the condition
	ClinicalStatus *common.CodeableConcept `json:"clinicalStatus,omitempty"`

	// The verification status to support the clinical status of the condition
	VerificationStatus *common.CodeableConcept `json:"verificationStatus,omitempty"`

	// A category assigned to the condition
	Category []common.CodeableConcept `json:"category,omitempty"`

	// A subjective assessment of the severity of the condition
	Severity *common.CodeableConcept `json:"severity,omitempty"`

	// Identification of the condition, problem or diagnosis
	Code *common.CodeableConcept `json:"code,omitempty"`

	// Anatomical location, if relevant
	BodySite []common.CodeableConcept `json:"bodySite,omitempty"`

	// Who has the condition
	Subject common.Reference `json:"subject"`

	// Encounter created as part of
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Estimated or actual date, date-time, or age when condition began
	OnsetDateTime *string        `json:"onsetDateTime,omitempty"`
	OnsetAge      *Age           `json:"onsetAge,omitempty"`
	OnsetPeriod   *common.Period `json:"onsetPeriod,omitempty"`
	OnsetRange    *Range         `json:"onsetRange,omitempty"`
	OnsetString   *string        `json:"onsetString,omitempty"`

	// When in resolution/remission
	AbatementDateTime *string        `json:"abatementDateTime,omitempty"`
	AbatementAge      *Age           `json:"abatementAge,omitempty"`
	AbatementPeriod   *common.Period `json:"abatementPeriod,omitempty"`
	AbatementRange    *Range         `json:"abatementRange,omitempty"`
	AbatementString   *string        `json:"abatementString,omitempty"`
	AbatementBoolean  *bool          `json:"abatementBoolean,omitempty"`

	// Date record was first recorded
	RecordedDate *string `json:"recordedDate,omitempty"`

	// Who recorded the record and takes responsibility for its content
	Recorder *common.Reference `json:"recorder,omitempty"`

	// Supporting evidence
	Evidence []ConditionEvidence `json:"evidence,omitempty"`

	// Additional information about the Condition
	Note []Annotation `json:"note,omitempty"`

	// Stage/grade, usually assessed formally
	Stage []ConditionStage `json:"stage,omitempty"`
}

// ConditionEvidence represents supporting evidence for the condition
type ConditionEvidence struct {
	common.BackboneElement

	// Manifestation/symptom
	Code []common.CodeableConcept `json:"code,omitempty"`

	// Supporting information found elsewhere
	Detail []common.Reference `json:"detail,omitempty"`
}

// ConditionStage represents the stage or grade of the condition
type ConditionStage struct {
	common.BackboneElement

	// Simple summary (disease specific)
	Summary *common.CodeableConcept `json:"summary,omitempty"`

	// Formal record of assessment
	Assessment []common.Reference `json:"assessment,omitempty"`

	// The type of staging
	Type *common.CodeableConcept `json:"type,omitempty"`
}
