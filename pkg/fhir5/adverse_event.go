// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// AdverseEventActuality represents whether the event actually happened or was potential
type AdverseEventActuality string

const (
	AdverseEventActualityActual    AdverseEventActuality = "actual"
	AdverseEventActualityPotential AdverseEventActuality = "potential"
)

// AdverseEventStatus represents the status of the adverse event
type AdverseEventStatus string

const (
	AdverseEventStatusInProgress     AdverseEventStatus = "in-progress"
	AdverseEventStatusCompleted      AdverseEventStatus = "completed"
	AdverseEventStatusEnteredInError AdverseEventStatus = "entered-in-error"
	AdverseEventStatusUnknown        AdverseEventStatus = "unknown"
)

// AdverseEventParticipant represents who participated in the adverse event and how they were involved
type AdverseEventParticipant struct {
	common.BackboneElement

	// The physician prescribing a drug, a nurse administering the drug, a device that administered the drug, a witness to the event, or an informant of clinical history
	Actor common.Reference `json:"actor"`

	// Distinguishes the type of involvement of the actor in the adverse event, such as contributor or informant
	Function *common.CodeableConcept `json:"function,omitempty"`
}

// AdverseEventSuspectEntityCausality represents information on the possible cause of the event
type AdverseEventSuspectEntityCausality struct {
	common.BackboneElement

	// The method of evaluating the relatedness of the suspected entity to the event
	AssessmentMethod *common.CodeableConcept `json:"assessmentMethod,omitempty"`

	// The author of the information on the possible cause of the event
	Author *common.Reference `json:"author,omitempty"`

	// The result of the assessment regarding the relatedness of the suspected entity to the event
	EntityRelatedness *common.CodeableConcept `json:"entityRelatedness,omitempty"`
}

// AdverseEventSuspectEntity represents the entity that is suspected to have caused the adverse event
type AdverseEventSuspectEntity struct {
	common.BackboneElement

	// Information on the possible cause of the event
	Causality *AdverseEventSuspectEntityCausality `json:"causality,omitempty"`

	// Identifies the actual instance of what caused the adverse event
	InstanceCodeableConcept *common.CodeableConcept `json:"instanceCodeableConcept,omitempty"`
	InstanceReference       *common.Reference       `json:"instanceReference,omitempty"`
}

// AdverseEventContributingFactor represents the contributing factors suspected to have increased the probability or severity of the adverse event
type AdverseEventContributingFactor struct {
	common.BackboneElement

	// The item that is suspected to have increased the probability or severity of the adverse event
	ItemReference       *common.Reference       `json:"itemReference,omitempty"`
	ItemCodeableConcept *common.CodeableConcept `json:"itemCodeableConcept,omitempty"`
}

// AdverseEventPreventiveAction represents preventive actions that contributed to avoiding the adverse event
type AdverseEventPreventiveAction struct {
	common.BackboneElement

	// The action that contributed to avoiding the adverse event
	ItemReference       *common.Reference       `json:"itemReference,omitempty"`
	ItemCodeableConcept *common.CodeableConcept `json:"itemCodeableConcept,omitempty"`
}

// AdverseEventMitigatingAction represents the ameliorating action taken after the adverse event occurred in order to reduce the extent of harm
type AdverseEventMitigatingAction struct {
	common.BackboneElement

	// The ameliorating action taken after the adverse event occurred in order to reduce the extent of harm
	ItemReference       *common.Reference       `json:"itemReference,omitempty"`
	ItemCodeableConcept *common.CodeableConcept `json:"itemCodeableConcept,omitempty"`
}

// AdverseEventSupportingInfo represents supporting information relevant to the event
type AdverseEventSupportingInfo struct {
	common.BackboneElement

	// Relevant past history for the subject
	ItemReference       *common.Reference       `json:"itemReference,omitempty"`
	ItemCodeableConcept *common.CodeableConcept `json:"itemCodeableConcept,omitempty"`
}

// AdverseEvent represents an event that may be related to unintended effects on a patient or research participant
type AdverseEvent struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "AdverseEvent"

	// Whether the event actually happened or was potential
	Actuality AdverseEventActuality `json:"actuality"`

	// The overall type of event, intended for search and filtering purposes
	Category []common.CodeableConcept `json:"category,omitempty"`

	// Specific event that occurred or that was averted
	Code *common.CodeableConcept `json:"code,omitempty"`

	// The contributing factors suspected to have increased the probability or severity of the adverse event
	ContributingFactor []AdverseEventContributingFactor `json:"contributingFactor,omitempty"`

	// Estimated or actual date the AdverseEvent began, in the opinion of the reporter
	Detected *string `json:"detected,omitempty"`

	// The encounter the event occurred within
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Considered likely or probable or anticipated in the research study
	ExpectedInResearchStudy *bool `json:"expectedInResearchStudy,omitempty"`

	// Business identifier for the adverse event
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The information about where the adverse event occurred
	Location *common.Reference `json:"location,omitempty"`

	// The ameliorating action taken after the adverse event occurred in order to reduce the extent of harm
	MitigatingAction []AdverseEventMitigatingAction `json:"mitigatingAction,omitempty"`

	// Comments made about the adverse event by the performer, subject or other participants
	Note []Annotation `json:"note,omitempty"`

	// The date (and perhaps time) when the adverse event occurred
	OccurrenceDateTime *string        `json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod   *common.Period `json:"occurrencePeriod,omitempty"`
	OccurrenceTiming   *Timing        `json:"occurrenceTiming,omitempty"`

	// Describes the type of outcome from the adverse event
	Outcome []common.CodeableConcept `json:"outcome,omitempty"`

	// Indicates who or what participated in the adverse event and how they were involved
	Participant []AdverseEventParticipant `json:"participant,omitempty"`

	// Preventive actions that contributed to avoiding the adverse event
	PreventiveAction []AdverseEventPreventiveAction `json:"preventiveAction,omitempty"`

	// The recordedDate represents the date when this particular AdverseEvent record was created in the system
	RecordedDate *string `json:"recordedDate,omitempty"`

	// Information on who recorded the adverse event
	Recorder *common.Reference `json:"recorder,omitempty"`

	// Information about the condition that occurred as a result of the adverse event
	ResultingEffect []common.Reference `json:"resultingEffect,omitempty"`

	// Assessment of if the event was of clinical importance
	Seriousness *common.CodeableConcept `json:"seriousness,omitempty"`

	// in-progress | completed | entered-in-error | unknown
	Status AdverseEventStatus `json:"status"`

	// The research study that the subject is enrolled in
	Study []common.Reference `json:"study,omitempty"`

	// Subject impacted by event
	Subject common.Reference `json:"subject"`

	// Supporting information relevant to the event
	SupportingInfo []AdverseEventSupportingInfo `json:"supportingInfo,omitempty"`

	// Describes the entity that is suspected to have caused the adverse event
	SuspectEntity []AdverseEventSuspectEntity `json:"suspectEntity,omitempty"`
}
