package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// FamilyMemberHistory represents significant health conditions for a person related to the patient
type FamilyMemberHistory struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "FamilyMemberHistory"

	// Use estimatedAge to indicate whether the age is actual or not
	AgeAge *common.Age `json:"ageAge,omitempty"`

	// Use estimatedAge to indicate whether the age is actual or not
	AgeRange *common.Range `json:"ageRange,omitempty"`

	// Use estimatedAge to indicate whether the age is actual or not
	AgeString *string `json:"ageString,omitempty"`

	// The actual or approximate date of birth of the relative
	BornPeriod *common.Period `json:"bornPeriod,omitempty"`

	// The actual or approximate date of birth of the relative
	BornDate *string `json:"bornDate,omitempty"`

	// The actual or approximate date of birth of the relative
	BornString *string `json:"bornString,omitempty"`

	// The significant Conditions (or condition) that the family member had
	Condition []FamilyMemberHistoryCondition `json:"condition,omitempty"`

	// Describes why the family member's history is not available
	DataAbsentReason *common.CodeableConcept `json:"dataAbsentReason,omitempty"`

	// This should be captured even if the same as the date on the List aggregating the full family history
	Date *string `json:"date,omitempty"`

	// Deceased flag or the actual or approximate age of the relative at the time of death
	DeceasedBoolean *bool `json:"deceasedBoolean,omitempty"`

	// Deceased flag or the actual or approximate age of the relative at the time of death
	DeceasedAge *common.Age `json:"deceasedAge,omitempty"`

	// Deceased flag or the actual or approximate age of the relative at the time of death
	DeceasedRange *common.Range `json:"deceasedRange,omitempty"`

	// Deceased flag or the actual or approximate age of the relative at the time of death
	DeceasedDate *string `json:"deceasedDate,omitempty"`

	// Deceased flag or the actual or approximate age of the relative at the time of death
	DeceasedString *string `json:"deceasedString,omitempty"`

	// This element is labeled as a modifier because the fact that age is estimated can/should change the results
	EstimatedAge *bool `json:"estimatedAge,omitempty"`

	// This is a business identifier, not a resource identifier
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The URL pointing to a FHIR-defined protocol, guideline, orderset or other definition
	InstantiatesCanonical []string `json:"instantiatesCanonical,omitempty"`

	// This might be an HTML page, PDF, etc. or could just be a non-resolvable URI identifier
	InstantiatesUri []string `json:"instantiatesUri,omitempty"`

	// This will either be a name or a description; e.g. "Aunt Susan", "my cousin with the red hair"
	Name *string `json:"name,omitempty"`

	// This property allows a non condition-specific note to the made about the related person
	Note []common.Annotation `json:"note,omitempty"`

	// The person who this history concerns
	Patient common.Reference `json:"patient"`

	// Textual reasons can be captured using reasonCode.text
	ReasonCode []common.CodeableConcept `json:"reasonCode,omitempty"`

	// Indicates a Condition, Observation, AllergyIntolerance, or QuestionnaireResponse that justifies this family member history event
	ReasonReference []common.Reference `json:"reasonReference,omitempty"`

	// The type of relationship this person has to the patient (father, mother, brother etc.)
	Relationship common.CodeableConcept `json:"relationship"`

	// This element should ideally reflect whether the individual is genetically male or female
	Sex *common.CodeableConcept `json:"sex,omitempty"`

	// This element is labeled as a modifier because the status contains codes that mark the resource as not currently valid
	Status FamilyMemberHistoryStatus `json:"status"`
}

// FamilyMemberHistoryStatus represents the status of the family member history
type FamilyMemberHistoryStatus string

const (
	FamilyMemberHistoryStatusPartial        FamilyMemberHistoryStatus = "partial"
	FamilyMemberHistoryStatusCompleted      FamilyMemberHistoryStatus = "completed"
	FamilyMemberHistoryStatusEnteredInError FamilyMemberHistoryStatus = "entered-in-error"
	FamilyMemberHistoryStatusHealthUnknown  FamilyMemberHistoryStatus = "health-unknown"
)

// FamilyMemberHistoryCondition represents the significant conditions that the family member had
type FamilyMemberHistoryCondition struct {
	common.BackboneElement

	// The actual condition specified. Could be a coded condition (like MI or Diabetes) or a less specific string
	Code common.CodeableConcept `json:"code"`

	// This condition contributed to the cause of death of the related person
	ContributedToDeath *bool `json:"contributedToDeath,omitempty"`

	// An area where general notes can be placed about this specific condition
	Note []common.Annotation `json:"note,omitempty"`

	// Either the age of onset, range of approximate age or descriptive string can be recorded
	OnsetAge *common.Age `json:"onsetAge,omitempty"`

	// Either the age of onset, range of approximate age or descriptive string can be recorded
	OnsetRange *common.Range `json:"onsetRange,omitempty"`

	// Either the age of onset, range of approximate age or descriptive string can be recorded
	OnsetPeriod *common.Period `json:"onsetPeriod,omitempty"`

	// Either the age of onset, range of approximate age or descriptive string can be recorded
	OnsetString *string `json:"onsetString,omitempty"`

	// Indicates what happened following the condition
	Outcome *common.CodeableConcept `json:"outcome,omitempty"`
}
