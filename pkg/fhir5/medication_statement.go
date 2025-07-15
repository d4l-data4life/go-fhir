package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// MedicationStatementStatus represents the status of a medication statement
type MedicationStatementStatus string

const (
	MedicationStatementStatusRecorded       MedicationStatementStatus = "recorded"
	MedicationStatementStatusEnteredInError MedicationStatementStatus = "entered-in-error"
	MedicationStatementStatusDraft          MedicationStatementStatus = "draft"
)

// MedicationStatementAdherence represents adherence information for a medication statement
type MedicationStatementAdherence struct {
	common.BackboneElement

	// Type of the adherence for the medication
	Code common.CodeableConcept `json:"code"`

	// This is generally only used for "exception" statuses such as "entered-in-error"
	Reason *common.CodeableConcept `json:"reason,omitempty"`
}

// MedicationStatement represents a record of a medication that is being consumed by a patient
type MedicationStatement struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "MedicationStatement"

	// This element can be used to indicate whether a patient is following a course of treatment as instructed
	Adherence *MedicationStatementAdherence `json:"adherence,omitempty"`

	// Type of medication statement
	Category []common.CodeableConcept `json:"category,omitempty"`

	// The date when the Medication Statement was asserted by the information source
	DateAsserted *string `json:"dateAsserted,omitempty"`

	// Likely references would be to MedicationRequest, MedicationDispense, Claim, Observation or QuestionnaireAnswers
	DerivedFrom []common.Reference `json:"derivedFrom,omitempty"`

	// The dates included in the dosage on a Medication Statement reflect the dates for a given dose
	Dosage []Dosage `json:"dosage,omitempty"`

	// This attribute reflects the period over which the patient consumed the medication
	EffectiveDateTime *string        `json:"effectiveDateTime,omitempty"`
	EffectivePeriod   *common.Period `json:"effectivePeriod,omitempty"`
	EffectiveTiming   *Timing        `json:"effectiveTiming,omitempty"`

	// The encounter that establishes the context for this MedicationStatement
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Business identifier for this medication statement
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The person or organization that provided the information about the taking of this medication
	InformationSource []common.Reference `json:"informationSource,omitempty"`

	// If only a code is specified, then it needs to be a code for a specific product
	Medication CodeableReference `json:"medication"`

	// Provides extra information about the Medication Statement
	Note []Annotation `json:"note,omitempty"`

	// A larger event of which this particular MedicationStatement is a component or step
	PartOf []common.Reference `json:"partOf,omitempty"`

	// This could be a diagnosis code
	Reason []CodeableReference `json:"reason,omitempty"`

	// Link to information that is relevant to a medication statement
	RelatedClinicalInformation []common.Reference `json:"relatedClinicalInformation,omitempty"`

	// The full representation of the dose of the medication included in all dosage instructions
	RenderedDosageInstruction *string `json:"renderedDosageInstruction,omitempty"`

	// This status concerns just the recording of the medication statement
	Status MedicationStatementStatus `json:"status"`

	// The person, animal or group who is/was taking the medication
	Subject common.Reference `json:"subject"`
}
