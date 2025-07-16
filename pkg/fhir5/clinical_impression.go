// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// ClinicalImpressionFinding represents specific findings or diagnoses
type ClinicalImpressionFinding struct {
	common.BackboneElement

	// Which investigations support finding or diagnosis
	Basis *string `json:"basis,omitempty"`

	// Specific text, code or reference for finding or diagnosis
	Item *CodeableReference `json:"item,omitempty"`
}

// ClinicalImpression represents a record of a clinical assessment
type ClinicalImpression struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "ClinicalImpression"

	// Change in the status/pattern of a subject's condition since previously assessed
	ChangePattern *common.CodeableConcept `json:"changePattern,omitempty"`

	// Indicates when the documentation of the assessment was complete
	Date *string `json:"date,omitempty"`

	// A summary of the context and/or cause of the assessment
	Description *string `json:"description,omitempty"`

	// This SHOULD be accurate to at least the minute
	EffectiveDateTime *string `json:"effectiveDateTime,omitempty"`

	// This SHOULD be accurate to at least the minute
	EffectivePeriod *common.Period `json:"effectivePeriod,omitempty"`

	// This will typically be the encounter the event occurred within
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Specific findings or diagnoses that were considered likely or relevant to ongoing treatment
	Finding []ClinicalImpressionFinding `json:"finding,omitempty"`

	// This is a business identifier, not a resource identifier
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Don't use this element for content that should more properly appear as one of the specific elements
	Note []Annotation `json:"note,omitempty"`

	// The clinician performing the assessment
	Performer *common.Reference `json:"performer,omitempty"`

	// It is always likely that multiple previous assessments exist for a patient
	Previous *common.Reference `json:"previous,omitempty"`

	// e.g. The patient is a pregnant, has congestive heart failure, has an Adenocarcinoma
	Problem []common.Reference `json:"problem,omitempty"`

	// Estimate of likely outcome
	PrognosisCodeableConcept []common.CodeableConcept `json:"prognosisCodeableConcept,omitempty"`

	// RiskAssessment expressing likely outcome
	PrognosisReference []common.Reference `json:"prognosisReference,omitempty"`

	// Reference to a specific published clinical protocol that was followed during this assessment
	Protocol []string `json:"protocol,omitempty"`

	// This element is labeled as a modifier because the status contains the code entered-in-error
	Status ClinicalImpressionStatus `json:"status"`

	// This is generally only used for "exception" statuses such as "not-done", "on-hold" or "stopped"
	StatusReason *common.CodeableConcept `json:"statusReason,omitempty"`

	// The patient or group of individuals assessed as part of this record
	Subject common.Reference `json:"subject"`

	// A text summary of the investigations and the diagnosis
	Summary *string `json:"summary,omitempty"`

	// Information supporting the clinical impression, which can contain investigation results
	SupportingInfo []common.Reference `json:"supportingInfo,omitempty"`
}

// ClinicalImpressionStatus represents the status of a clinical impression
type ClinicalImpressionStatus string

const (
	ClinicalImpressionStatusPreparation    ClinicalImpressionStatus = "preparation"
	ClinicalImpressionStatusInProgress     ClinicalImpressionStatus = "in-progress"
	ClinicalImpressionStatusNotDone        ClinicalImpressionStatus = "not-done"
	ClinicalImpressionStatusOnHold         ClinicalImpressionStatus = "on-hold"
	ClinicalImpressionStatusStopped        ClinicalImpressionStatus = "stopped"
	ClinicalImpressionStatusCompleted      ClinicalImpressionStatus = "completed"
	ClinicalImpressionStatusEnteredInError ClinicalImpressionStatus = "entered-in-error"
	ClinicalImpressionStatusUnknown        ClinicalImpressionStatus = "unknown"
)
