// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// FamilyMemberHistoryStatus represents the status of a family member history record
type FamilyMemberHistoryStatus string

const (
	FamilyMemberHistoryStatusPartial        FamilyMemberHistoryStatus = "partial"
	FamilyMemberHistoryStatusCompleted      FamilyMemberHistoryStatus = "completed"
	FamilyMemberHistoryStatusEnteredInError FamilyMemberHistoryStatus = "entered-in-error"
	FamilyMemberHistoryStatusHealthUnknown  FamilyMemberHistoryStatus = "health-unknown"
)

// FamilyMemberHistoryParticipant represents who or what participated in the activities related to the family member history
type FamilyMemberHistoryParticipant struct {
	common.BackboneElement

	// Indicates who or what participated in the activities related to the family member history
	Actor common.Reference `json:"actor"`

	// Distinguishes the type of involvement of the actor in the activities related to the family member history
	Function *common.CodeableConcept `json:"function,omitempty"`
}

// FamilyMemberHistoryCondition represents the significant conditions that the family member had
type FamilyMemberHistoryCondition struct {
	common.BackboneElement

	// The actual condition specified
	Code common.CodeableConcept `json:"code"`

	// This condition contributed to the cause of death of the related person
	ContributedToDeath *bool `json:"contributedToDeath,omitempty"`

	// An area where general notes can be placed about this specific condition
	Note []Annotation `json:"note,omitempty"`

	// Either the age of onset, range of approximate age or descriptive string can be recorded
	OnsetAge    *Age           `json:"onsetAge,omitempty"`
	OnsetRange  *Range         `json:"onsetRange,omitempty"`
	OnsetPeriod *common.Period `json:"onsetPeriod,omitempty"`
	OnsetString *string        `json:"onsetString,omitempty"`

	// Indicates what happened following the condition
	Outcome *common.CodeableConcept `json:"outcome,omitempty"`
}

// FamilyMemberHistoryProcedure represents the significant procedures that the family member had
type FamilyMemberHistoryProcedure struct {
	common.BackboneElement

	// The actual procedure specified
	Code common.CodeableConcept `json:"code"`

	// This procedure contributed to the cause of death of the related person
	ContributedToDeath *bool `json:"contributedToDeath,omitempty"`

	// An area where general notes can be placed about this specific procedure
	Note []Annotation `json:"note,omitempty"`

	// Indicates what happened following the procedure
	Outcome *common.CodeableConcept `json:"outcome,omitempty"`

	// Estimated or actual date, date-time, period, or age when the procedure was performed
	PerformedAge    *Age           `json:"performedAge,omitempty"`
	PerformedRange  *Range         `json:"performedRange,omitempty"`
	PerformedPeriod *common.Period `json:"performedPeriod,omitempty"`
	PerformedString *string        `json:"performedString,omitempty"`
}

// FamilyMemberHistory represents significant health conditions for a person related to the patient
type FamilyMemberHistory struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "FamilyMemberHistory"

	// Use estimatedAge to indicate whether the age is actual or not
	AgeAge    *Age    `json:"ageAge,omitempty"`
	AgeRange  *Range  `json:"ageRange,omitempty"`
	AgeString *string `json:"ageString,omitempty"`

	// The actual or approximate date of birth of the relative
	BornPeriod *common.Period `json:"bornPeriod,omitempty"`
	BornDate   *string        `json:"bornDate,omitempty"`
	BornString *string        `json:"bornString,omitempty"`

	// The significant conditions that the family member had
	Condition []FamilyMemberHistoryCondition `json:"condition,omitempty"`

	// Describes why the family member's history is not available
	DataAbsentReason *common.CodeableConcept `json:"dataAbsentReason,omitempty"`

	// This should be captured even if the same as the date on the List aggregating the full family history
	Date *string `json:"date,omitempty"`

	// Deceased flag or the actual or approximate age of the relative at the time of death
	DeceasedBoolean *bool   `json:"deceasedBoolean,omitempty"`
	DeceasedAge     *Age    `json:"deceasedAge,omitempty"`
	DeceasedRange   *Range  `json:"deceasedRange,omitempty"`
	DeceasedDate    *string `json:"deceasedDate,omitempty"`
	DeceasedString  *string `json:"deceasedString,omitempty"`

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

	// This property allows a non condition-specific note to be made about the related person
	Note []Annotation `json:"note,omitempty"`

	// Indicates who or what participated in the activities related to the family member history
	Participant []FamilyMemberHistoryParticipant `json:"participant,omitempty"`

	// This is not the family member
	Patient common.Reference `json:"patient"`

	// The significant procedures that the family member had
	Procedure []FamilyMemberHistoryProcedure `json:"procedure,omitempty"`

	// Textual reasons can be captured using reasonCode.text
	Reason []CodeableReference `json:"reason,omitempty"`

	// The type of relationship this person has to the patient
	Relationship common.CodeableConcept `json:"relationship"`

	// This element should ideally reflect whether the individual is genetically male or female
	Sex *common.CodeableConcept `json:"sex,omitempty"`

	// This element is labeled as a modifier because the status contains codes that mark the resource as not currently valid
	Status FamilyMemberHistoryStatus `json:"status"`
}
