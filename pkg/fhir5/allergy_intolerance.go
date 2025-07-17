// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// AllergyIntoleranceCategory represents the category of an allergy/intolerance
type AllergyIntoleranceCategory string

const (
	AllergyIntoleranceCategoryFood        AllergyIntoleranceCategory = "food"
	AllergyIntoleranceCategoryMedication  AllergyIntoleranceCategory = "medication"
	AllergyIntoleranceCategoryEnvironment AllergyIntoleranceCategory = "environment"
	AllergyIntoleranceCategoryBiologic    AllergyIntoleranceCategory = "biologic"
)

// AllergyIntoleranceCriticality represents the criticality of an allergy/intolerance
type AllergyIntoleranceCriticality string

const (
	AllergyIntoleranceCriticalityLow            AllergyIntoleranceCriticality = "low"
	AllergyIntoleranceCriticalityHigh           AllergyIntoleranceCriticality = "high"
	AllergyIntoleranceCriticalityUnableToAssess AllergyIntoleranceCriticality = "unable-to-assess"
)

// AllergyIntoleranceType represents the type of an allergy/intolerance
type AllergyIntoleranceType string

const (
	AllergyIntoleranceTypeAllergy     AllergyIntoleranceType = "allergy"
	AllergyIntoleranceTypeIntolerance AllergyIntoleranceType = "intolerance"
)

// AllergyIntoleranceReactionSeverity represents the severity of an allergy/intolerance reaction
type AllergyIntoleranceReactionSeverity string

const (
	AllergyIntoleranceReactionSeverityMild     AllergyIntoleranceReactionSeverity = "mild"
	AllergyIntoleranceReactionSeverityModerate AllergyIntoleranceReactionSeverity = "moderate"
	AllergyIntoleranceReactionSeveritySevere   AllergyIntoleranceReactionSeverity = "severe"
)

// AllergyIntoleranceParticipant represents participants in the allergy/intolerance activities
type AllergyIntoleranceParticipant struct {
	common.BackboneElement

	// Indicates who or what participated in the activities related to the allergy or intolerance
	Actor common.Reference `json:"actor"`

	// Distinguishes the type of involvement of the actor in the activities related to the allergy or intolerance
	Function *common.CodeableConcept `json:"function,omitempty"`
}

// AllergyIntoleranceReaction represents adverse reaction event details
type AllergyIntoleranceReaction struct {
	common.BackboneElement

	// Use the description to provide any details of a particular event of the occurred reaction
	Description *string `json:"description,omitempty"`

	// Coding of the route of exposure with a terminology should be used wherever possible
	ExposureRoute *common.CodeableConcept `json:"exposureRoute,omitempty"`

	// Manifestation can be expressed as a single word, phrase or brief description
	Manifestation []CodeableReference `json:"manifestation"`

	// Use this field to record information indirectly related to a particular event
	Note []Annotation `json:"note,omitempty"`

	// Record of the date and/or time of the onset of the Reaction
	Onset *string `json:"onset,omitempty"`

	// It is acknowledged that this assessment is very subjective
	Severity *AllergyIntoleranceReactionSeverity `json:"severity,omitempty"`

	// Identification of the specific substance considered to be responsible for the Adverse Reaction event
	Substance *common.CodeableConcept `json:"substance,omitempty"`
}

// AllergyIntolerance represents a risk of harmful or undesirable physiological response
type AllergyIntolerance struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "AllergyIntolerance"

	// This data element has been included because it is currently being captured in some clinical systems
	Category []AllergyIntoleranceCategory `json:"category,omitempty"`

	// AllergyIntolerance.clinicalStatus should be present if verificationStatus is not entered-in-error
	ClinicalStatus *common.CodeableConcept `json:"clinicalStatus,omitempty"`

	// It is strongly recommended that this element be populated using a terminology
	Code *common.CodeableConcept `json:"code,omitempty"`

	// The default criticality value for any propensity to an adverse reaction should be 'Low Risk'
	Criticality *AllergyIntoleranceCriticality `json:"criticality,omitempty"`

	// The encounter when the allergy or intolerance was asserted
	Encounter *common.Reference `json:"encounter,omitempty"`

	// This is a business identifier, not a resource identifier
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// This date may be replicated by one of the Onset of Reaction dates
	LastOccurrence *string `json:"lastOccurrence,omitempty"`

	// For example: including reason for flagging a seriousness of 'High Risk'
	Note []Annotation `json:"note,omitempty"`

	// Age is generally used when the patient reports an age at which the AllergyIntolerance was noted
	OnsetDateTime *string        `json:"onsetDateTime,omitempty"`
	OnsetAge      *Age           `json:"onsetAge,omitempty"`
	OnsetPeriod   *common.Period `json:"onsetPeriod,omitempty"`
	OnsetRange    *Range         `json:"onsetRange,omitempty"`
	OnsetString   *string        `json:"onsetString,omitempty"`

	// Indicates who or what participated in the activities related to the allergy or intolerance and how they were involved
	Participant []AllergyIntoleranceParticipant `json:"participant,omitempty"`

	// The patient who has the allergy or intolerance
	Patient common.Reference `json:"patient"`

	// Details about each adverse reaction event linked to exposure to the identified substance
	Reaction []AllergyIntoleranceReaction `json:"reaction,omitempty"`

	// When onset date is unknown, recordedDate can be used to establish if the allergy or intolerance was present
	RecordedDate *string `json:"recordedDate,omitempty"`

	// Individual who recorded the record and takes responsibility for its content
	Recorder *common.Reference `json:"recorder,omitempty"`

	// Allergic (typically immune-mediated) reactions have been traditionally regarded as an indicator
	Type *AllergyIntoleranceType `json:"type,omitempty"`

	// The verification status to support the clinical status of the condition
	VerificationStatus *common.CodeableConcept `json:"verificationStatus,omitempty"`
}
