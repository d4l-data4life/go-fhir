package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// Procedure represents an action that is or was performed on or for a patient
type Procedure struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Procedure"

	// External Identifiers for this procedure
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// A code specifying the state of the procedure
	Status ProcedureStatus `json:"status"`

	// Reason for current status
	StatusReason *common.CodeableConcept `json:"statusReason,omitempty"`

	// Classification of the procedure
	Category *common.CodeableConcept `json:"category,omitempty"`

	// Identification of the procedure
	Code *common.CodeableConcept `json:"code,omitempty"`

	// Who the procedure was performed on
	Subject common.Reference `json:"subject"`

	// Encounter created as part of
	Encounter *common.Reference `json:"encounter,omitempty"`

	// When the procedure was performed
	PerformedDateTime *string        `json:"performedDateTime,omitempty"`
	PerformedPeriod   *common.Period `json:"performedPeriod,omitempty"`
	PerformedString   *string        `json:"performedString,omitempty"`
	PerformedAge      *Age           `json:"performedAge,omitempty"`
	PerformedRange    *Range         `json:"performedRange,omitempty"`

	// Who recorded the record and takes responsibility for its content
	Recorder *common.Reference `json:"recorder,omitempty"`

	// Person who asserts this procedure
	Asserter *common.Reference `json:"asserter,omitempty"`

	// The people who performed the procedure
	Performer []ProcedurePerformer `json:"performer,omitempty"`

	// Where the procedure happened
	Location *common.Reference `json:"location,omitempty"`

	// Coded reason procedure performed
	ReasonCode []common.CodeableConcept `json:"reasonCode,omitempty"`

	// The justification that the procedure was performed
	ReasonReference []common.Reference `json:"reasonReference,omitempty"`

	// Target body sites
	BodySite []common.CodeableConcept `json:"bodySite,omitempty"`

	// The result of procedure
	Outcome *common.CodeableConcept `json:"outcome,omitempty"`

	// Any report resulting from the procedure
	Report []common.Reference `json:"report,omitempty"`

	// Complication following the procedure
	Complication []common.CodeableConcept `json:"complication,omitempty"`

	// A condition that is a result of the procedure
	ComplicationDetail []common.Reference `json:"complicationDetail,omitempty"`

	// Instructions for follow up
	FollowUp []common.CodeableConcept `json:"followUp,omitempty"`

	// Additional information about the procedure
	Note []Annotation `json:"note,omitempty"`

	// Manipulated, implanted, or removed device
	FocalDevice []ProcedureFocalDevice `json:"focalDevice,omitempty"`

	// Items used during procedure
	UsedReference []common.Reference `json:"usedReference,omitempty"`

	// Coded items used during the procedure
	UsedCode []common.CodeableConcept `json:"usedCode,omitempty"`

	// A request for this procedure
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// Part of referenced event
	PartOf []common.Reference `json:"partOf,omitempty"`
}

// ProcedureStatus represents the status of the procedure
type ProcedureStatus string

const (
	ProcedureStatusPreparation    ProcedureStatus = "preparation"
	ProcedureStatusInProgress     ProcedureStatus = "in-progress"
	ProcedureStatusNotDone        ProcedureStatus = "not-done"
	ProcedureStatusOnHold         ProcedureStatus = "on-hold"
	ProcedureStatusStopped        ProcedureStatus = "stopped"
	ProcedureStatusCompleted      ProcedureStatus = "completed"
	ProcedureStatusEnteredInError ProcedureStatus = "entered-in-error"
	ProcedureStatusUnknown        ProcedureStatus = "unknown"
)

// ProcedurePerformer represents a performer of the procedure
type ProcedurePerformer struct {
	common.BackboneElement

	// Type of performance
	Function *common.CodeableConcept `json:"function,omitempty"`

	// The reference to the practitioner
	Actor common.Reference `json:"actor"`

	// Organization the device or practitioner was acting for
	OnBehalfOf *common.Reference `json:"onBehalfOf,omitempty"`
}

// ProcedureFocalDevice represents a device that was manipulated during the procedure
type ProcedureFocalDevice struct {
	common.BackboneElement

	// Kind of change to device
	Action *common.CodeableConcept `json:"action,omitempty"`

	// Device that was changed
	Manipulated common.Reference `json:"manipulated"`
}
