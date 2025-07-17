// Package fhir4 contains FHIR R4 (version 4.0.1) resource definitions
package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// AllergyIntolerance represents risk of harmful or undesirable, physiological response which is unique to an individual and associated with exposure to a substance
type AllergyIntolerance struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "AllergyIntolerance"

	// External ids for this item
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// active | inactive | resolved
	ClinicalStatus *common.CodeableConcept `json:"clinicalStatus,omitempty"`

	// unconfirmed | confirmed | refuted | entered-in-error
	VerificationStatus *common.CodeableConcept `json:"verificationStatus,omitempty"`

	// allergy | intolerance - Underlying mechanism (if known)
	Type *AllergyIntoleranceType `json:"type,omitempty"`

	// food | medication | environment | biologic
	Category []AllergyIntoleranceCategory `json:"category,omitempty"`

	// low | high | unable-to-assess
	Criticality *AllergyIntoleranceCriticality `json:"criticality,omitempty"`

	// Code that identifies the allergy or intolerance
	Code *common.CodeableConcept `json:"code,omitempty"`

	// Who the sensitivity is for
	Patient common.Reference `json:"patient"`

	// Encounter when the allergy or intolerance was asserted
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Date first version of the resource instance was recorded
	OnsetDateTime *string        `json:"onsetDateTime,omitempty"`
	OnsetAge      *Age           `json:"onsetAge,omitempty"`
	OnsetPeriod   *common.Period `json:"onsetPeriod,omitempty"`
	OnsetRange    *Range         `json:"onsetRange,omitempty"`
	OnsetString   *string        `json:"onsetString,omitempty"`

	// Date(/time) of last known occurrence of a reaction
	RecordedDate *string `json:"recordedDate,omitempty"`

	// Who recorded the sensitivity
	Recorder *common.Reference `json:"recorder,omitempty"`

	// Source of the information about the allergy
	Asserter *common.Reference `json:"asserter,omitempty"`

	// Date(/time) of last known occurrence of a reaction
	LastOccurrence *string `json:"lastOccurrence,omitempty"`

	// Additional text not captured in other fields
	Note []Annotation `json:"note,omitempty"`

	// Adverse Reaction Events linked to exposure to substance
	Reaction []AllergyIntoleranceReaction `json:"reaction,omitempty"`
}

// AllergyIntoleranceType represents the type of allergy intolerance
type AllergyIntoleranceType string

const (
	AllergyIntoleranceTypeAllergy     AllergyIntoleranceType = "allergy"
	AllergyIntoleranceTypeIntolerance AllergyIntoleranceType = "intolerance"
)

// AllergyIntoleranceCategory represents the category of allergy intolerance
type AllergyIntoleranceCategory string

const (
	AllergyIntoleranceCategoryFood        AllergyIntoleranceCategory = "food"
	AllergyIntoleranceCategoryMedication  AllergyIntoleranceCategory = "medication"
	AllergyIntoleranceCategoryEnvironment AllergyIntoleranceCategory = "environment"
	AllergyIntoleranceCategoryBiologic    AllergyIntoleranceCategory = "biologic"
)

// AllergyIntoleranceCriticality represents the criticality of allergy intolerance
type AllergyIntoleranceCriticality string

const (
	AllergyIntoleranceCriticalityLow            AllergyIntoleranceCriticality = "low"
	AllergyIntoleranceCriticalityHigh           AllergyIntoleranceCriticality = "high"
	AllergyIntoleranceCriticalityUnableToAssess AllergyIntoleranceCriticality = "unable-to-assess"
)

// AllergyIntoleranceReaction represents a reaction event linked to exposure to substance
type AllergyIntoleranceReaction struct {
	common.BackboneElement

	// Specific substance or pharmaceutical product considered to be responsible
	Substance *common.CodeableConcept `json:"substance,omitempty"`

	// Clinical symptoms/signs associated with the Event
	Manifestation []common.CodeableConcept `json:"manifestation"`

	// Description of the event as a whole
	Description *string `json:"description,omitempty"`

	// Date(/time) when manifestations showed
	Onset *string `json:"onset,omitempty"`

	// mild | moderate | severe (of event as a whole)
	Severity *AllergyIntoleranceReactionSeverity `json:"severity,omitempty"`

	// How the subject was exposed to the substance
	ExposureRoute *common.CodeableConcept `json:"exposureRoute,omitempty"`

	// Text about event not captured in other fields
	Note []Annotation `json:"note,omitempty"`
}

// AllergyIntoleranceReactionSeverity represents the severity of a reaction event
type AllergyIntoleranceReactionSeverity string

const (
	AllergyIntoleranceReactionSeverityMild     AllergyIntoleranceReactionSeverity = "mild"
	AllergyIntoleranceReactionSeverityModerate AllergyIntoleranceReactionSeverity = "moderate"
	AllergyIntoleranceReactionSeveritySevere   AllergyIntoleranceReactionSeverity = "severe"
)
