package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// AdverseEventActuality represents whether the event actually happened, or just had the potential to
type AdverseEventActuality string

const (
	AdverseEventActualityActual    AdverseEventActuality = "actual"
	AdverseEventActualityPotential AdverseEventActuality = "potential"
)

// AdverseEventSuspectEntityCausality represents information on the possible cause of the event
type AdverseEventSuspectEntityCausality struct {
	common.BackboneElement

	// Assessment of if the entity caused the event
	Assessment *common.CodeableConcept `json:"assessment,omitempty"`

	// AdverseEvent.suspectEntity.causalityAuthor
	Author *common.Reference `json:"author,omitempty"`

	// ProbabilityScale | Bayesian | Checklist
	Method *common.CodeableConcept `json:"method,omitempty"`

	// AdverseEvent.suspectEntity.causalityProductRelatedness
	ProductRelatedness *string `json:"productRelatedness,omitempty"`
}

// AdverseEventSuspectEntity represents the entity that is suspected to have caused the adverse event
type AdverseEventSuspectEntity struct {
	common.BackboneElement

	// Information on the possible cause of the event
	Causality []AdverseEventSuspectEntityCausality `json:"causality,omitempty"`

	// Identifies the actual instance of what caused the adverse event. May be a substance, medication, medication administration, medication statement or a device
	Instance common.Reference `json:"instance"`
}

// AdverseEvent represents actual or potential/avoided event causing unintended physical injury
type AdverseEvent struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "AdverseEvent"

	// Whether the event actually happened, or just had the potential to. Note that this is independent of whether anyone was affected or harmed or how severely
	Actuality AdverseEventActuality `json:"actuality"`

	// The overall type of event, intended for search and filtering purposes
	Category []common.CodeableConcept `json:"category,omitempty"`

	// Parties that may or should contribute or have contributed information to the adverse event, which can consist of one or more activities
	Contributor []common.Reference `json:"contributor,omitempty"`

	// The date (and perhaps time) when the adverse event occurred
	Date *string `json:"date,omitempty"`

	// Estimated or actual date the AdverseEvent began, in the opinion of the reporter
	Detected *string `json:"detected,omitempty"`

	// This will typically be the encounter the event occurred within, but some activities may be initiated prior to or after the official completion of an encounter but still be tied to the context of the encounter
	Encounter *common.Reference `json:"encounter,omitempty"`

	// This element defines the specific type of event that occurred or that was prevented from occurring
	Event *common.CodeableConcept `json:"event,omitempty"`

	// This is a business identifier, not a resource identifier
	Identifier *common.Identifier `json:"identifier,omitempty"`

	// The information about where the adverse event occurred
	Location *common.Reference `json:"location,omitempty"`

	// Describes the type of outcome from the adverse event
	Outcome *common.CodeableConcept `json:"outcome,omitempty"`

	// The recordedDate represents the date when this particular AdverseEvent record was created in the system, not the date of the most recent update
	RecordedDate *string `json:"recordedDate,omitempty"`

	// Information on who recorded the adverse event. May be the patient or a practitioner
	Recorder *common.Reference `json:"recorder,omitempty"`

	// AdverseEvent.referenceDocument
	ReferenceDocument []common.Reference `json:"referenceDocument,omitempty"`

	// Includes information about the reaction that occurred as a result of exposure to a substance (for example, a drug or a chemical)
	ResultingCondition []common.Reference `json:"resultingCondition,omitempty"`

	// Assessment whether this event was of real importance
	Seriousness *common.CodeableConcept `json:"seriousness,omitempty"`

	// Describes the severity of the adverse event, in relation to the subject. Contrast to AdverseEvent.seriousness - a severe rash might not be serious, but a mild heart problem is
	Severity *common.CodeableConcept `json:"severity,omitempty"`

	// AdverseEvent.study
	Study []common.Reference `json:"study,omitempty"`

	// If AdverseEvent.resultingCondition differs among members of the group, then use Patient as the subject
	Subject common.Reference `json:"subject"`

	// AdverseEvent.subjectMedicalHistory
	SubjectMedicalHistory []common.Reference `json:"subjectMedicalHistory,omitempty"`

	// Describes the entity that is suspected to have caused the adverse event
	SuspectEntity []AdverseEventSuspectEntity `json:"suspectEntity,omitempty"`
}
