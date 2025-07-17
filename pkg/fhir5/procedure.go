// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// ProcedureStatus represents the status of a procedure
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

// ProcedurePerformer represents who performed the procedure
type ProcedurePerformer struct {
	common.BackboneElement

	// Indicates who or what performed the procedure
	Actor common.Reference `json:"actor"`

	// Distinguishes the type of involvement of the performer in the procedure
	Function *common.CodeableConcept `json:"function,omitempty"`

	// Organization they were acting on behalf of when performing the action
	OnBehalfOf *common.Reference `json:"onBehalfOf,omitempty"`

	// Time period during which the performer performed the procedure
	Period *common.Period `json:"period,omitempty"`
}

// ProcedureFocalDevice represents a device that was manipulated during the procedure
type ProcedureFocalDevice struct {
	common.BackboneElement

	// The kind of change that happened to the device during the procedure
	Action *common.CodeableConcept `json:"action,omitempty"`

	// The device that was manipulated (changed) during the procedure
	Manipulated common.Reference `json:"manipulated"`
}

// Procedure represents an action that is or was performed on or for a patient
type Procedure struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Procedure"

	// A reference to a resource that contains details of the request for this procedure
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// Specific anatomical location where the procedure was performed
	BodySite []common.CodeableConcept `json:"bodySite,omitempty"`

	// A code that classifies the procedure for searching, sorting and display purposes
	Category []common.CodeableConcept `json:"category,omitempty"`

	// The specific procedure that is performed
	Code *common.CodeableConcept `json:"code,omitempty"`

	// Complications that occurred during the procedure
	Complication []CodeableReference `json:"complication,omitempty"`

	// The encounter during which the procedure was performed
	Encounter *common.Reference `json:"encounter,omitempty"`

	// A device that is implanted, removed or otherwise manipulated during the procedure
	FocalDevice []ProcedureFocalDevice `json:"focalDevice,omitempty"`

	// Who is the target of the procedure when it is not the subject of record only
	Focus *common.Reference `json:"focus,omitempty"`

	// If the procedure required specific follow up
	FollowUp []common.CodeableConcept `json:"followUp,omitempty"`

	// Business identifier for this procedure
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The URL pointing to a FHIR-defined protocol, guideline, order set or other definition
	InstantiatesCanonical []string `json:"instantiatesCanonical,omitempty"`

	// The URL pointing to an externally maintained protocol, guideline, order set
	InstantiatesUri []string `json:"instantiatesUri,omitempty"`

	// Where the procedure happened
	Location *common.Reference `json:"location,omitempty"`

	// Any other notes and comments about the procedure
	Note []Annotation `json:"note,omitempty"`

	// When the procedure occurred
	OccurrenceDateTime *string        `json:"occurrenceDateTime,omitempty"`
	OccurrenceAge      *common.Age    `json:"occurrenceAge,omitempty"`
	OccurrencePeriod   *common.Period `json:"occurrencePeriod,omitempty"`
	OccurrenceRange    *Range         `json:"occurrenceRange,omitempty"`
	OccurrenceString   *string        `json:"occurrenceString,omitempty"`
	OccurrenceTiming   *Timing        `json:"occurrenceTiming,omitempty"`

	// The outcome of the procedure
	Outcome *common.CodeableConcept `json:"outcome,omitempty"`

	// Part of referenced resource
	PartOf []common.Reference `json:"partOf,omitempty"`

	// Who performed the procedure and how they were involved
	Performer []ProcedurePerformer `json:"performer,omitempty"`

	// The reason why the procedure was performed
	Reason []CodeableReference `json:"reason,omitempty"`

	// The date the occurrence of the procedure was first captured in the record
	Recorded *string `json:"recorded,omitempty"`

	// Individual who recorded the record and takes responsibility for its content
	Recorder *common.Reference `json:"recorder,omitempty"`

	// There could potentially be multiple reports
	Report []common.Reference `json:"report,omitempty"`

	// Indicates if this record was captured as a secondary 'reported' record
	ReportedBoolean   *bool             `json:"reportedBoolean,omitempty"`
	ReportedReference *common.Reference `json:"reportedReference,omitempty"`

	// The status of the procedure
	Status ProcedureStatus `json:"status"`

	// Captures the reason for the current state of the procedure
	StatusReason *common.CodeableConcept `json:"statusReason,omitempty"`

	// On whom or on what the procedure was performed
	Subject common.Reference `json:"subject"`

	// Other resources from the patient record that may be relevant to the procedure
	SupportingInfo []common.Reference `json:"supportingInfo,omitempty"`

	// For devices actually implanted or removed, use Procedure.focalDevice.manipulated
	Used []CodeableReference `json:"used,omitempty"`
}
