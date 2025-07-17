package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

type MedicationRequest struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "MedicationRequest"

	// Identifiers assigned to this order
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// active | on-hold | cancelled | completed | entered-in-error | stopped | draft | unknown
	Status MedicationRequestStatus `json:"status"`

	// Reason for current status
	StatusReason *common.CodeableConcept `json:"statusReason,omitempty"`

	// proposal | plan | order | original-order | reflex-order | filler-order | instance-order | option
	Intent MedicationRequestIntent `json:"intent"`

	// Type of medication usage
	Category []common.CodeableConcept `json:"category,omitempty"`

	// routine | urgent | asap | stat
	Priority *MedicationRequestPriority `json:"priority,omitempty"`

	// True if request is prohibiting action
	DoNotPerform *bool `json:"doNotPerform,omitempty"`

	// Reported rather than primary record
	ReportedBoolean   *bool             `json:"reportedBoolean,omitempty"`
	ReportedReference *common.Reference `json:"reportedReference,omitempty"`

	// Medication to be taken
	MedicationCodeableConcept *common.CodeableConcept `json:"medicationCodeableConcept,omitempty"`
	MedicationReference       *common.Reference       `json:"medicationReference,omitempty"`

	// Who or group medication request is for
	Subject common.Reference `json:"subject"`

	// Encounter created as part of encounter/admission/stay
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Information to support ordering of the medication
	SupportingInformation []common.Reference `json:"supportingInformation,omitempty"`

	// When request was initially authored
	AuthoredOn *string `json:"authoredOn,omitempty"`

	// Who/What requested the Request
	Requester *common.Reference `json:"requester,omitempty"`

	// Person who entered the order on behalf of another
	Recorder *common.Reference `json:"recorder,omitempty"`

	// Reason or indication for ordering or not ordering the medication
	ReasonCode      []common.CodeableConcept `json:"reasonCode,omitempty"`
	ReasonReference []common.Reference       `json:"reasonReference,omitempty"`

	// Instantiates FHIR protocol or definition
	InstantiatesCanonical []string `json:"instantiatesCanonical,omitempty"`
	InstantiatesUri       []string `json:"instantiatesUri,omitempty"`

	// What request fulfills
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// Composite Request ID
	GroupIdentifier *common.Identifier `json:"groupIdentifier,omitempty"`

	// Overall pattern of medication administration
	CourseOfTherapyType *common.CodeableConcept `json:"courseOfTherapyType,omitempty"`

	// Associated insurance coverage
	Insurance []common.Reference `json:"insurance,omitempty"`

	// Information about the prescription
	Note []Annotation `json:"note,omitempty"`

	// How the medication should be taken
	DosageInstruction []Dosage `json:"dosageInstruction,omitempty"`

	// Medication supply authorization
	DispenseRequest *MedicationRequestDispenseRequest `json:"dispenseRequest,omitempty"`

	// Any restrictions on medication substitution
	Substitution *MedicationRequestSubstitution `json:"substitution,omitempty"`

	// An order/prescription that is being replaced
	PriorPrescription *common.Reference `json:"priorPrescription,omitempty"`

	// Clinical Issue with action
	DetectedIssue []common.Reference `json:"detectedIssue,omitempty"`

	// A list of events of interest in the lifecycle
	EventHistory []common.Reference `json:"eventHistory,omitempty"`
}

type MedicationRequestStatus string

const (
	MedicationRequestStatusActive         MedicationRequestStatus = "active"
	MedicationRequestStatusOnHold         MedicationRequestStatus = "on-hold"
	MedicationRequestStatusCancelled      MedicationRequestStatus = "cancelled"
	MedicationRequestStatusCompleted      MedicationRequestStatus = "completed"
	MedicationRequestStatusEnteredInError MedicationRequestStatus = "entered-in-error"
	MedicationRequestStatusStopped        MedicationRequestStatus = "stopped"
	MedicationRequestStatusDraft          MedicationRequestStatus = "draft"
	MedicationRequestStatusUnknown        MedicationRequestStatus = "unknown"
)

type MedicationRequestIntent string

const (
	MedicationRequestIntentProposal      MedicationRequestIntent = "proposal"
	MedicationRequestIntentPlan          MedicationRequestIntent = "plan"
	MedicationRequestIntentOrder         MedicationRequestIntent = "order"
	MedicationRequestIntentOriginalOrder MedicationRequestIntent = "original-order"
	MedicationRequestIntentReflexOrder   MedicationRequestIntent = "reflex-order"
	MedicationRequestIntentFillerOrder   MedicationRequestIntent = "filler-order"
	MedicationRequestIntentInstanceOrder MedicationRequestIntent = "instance-order"
	MedicationRequestIntentOption        MedicationRequestIntent = "option"
)

type MedicationRequestPriority string

const (
	MedicationRequestPriorityRoutine MedicationRequestPriority = "routine"
	MedicationRequestPriorityUrgent  MedicationRequestPriority = "urgent"
	MedicationRequestPriorityAsap    MedicationRequestPriority = "asap"
	MedicationRequestPriorityStat    MedicationRequestPriority = "stat"
)

type MedicationRequestDispenseRequest struct {
	common.BackboneElement

	// First fill details
	InitialFill *MedicationRequestDispenseRequestInitialFill `json:"initialFill,omitempty"`

	// Minimum period of time between dispenses
	DispenseInterval *Duration `json:"dispenseInterval,omitempty"`

	// Time period supply is authorized for
	ValidityPeriod *common.Period `json:"validityPeriod,omitempty"`

	// Number of refills authorized
	NumberOfRepeatsAllowed *int `json:"numberOfRepeatsAllowed,omitempty"`

	// Amount of medication to supply per dispense
	Quantity *common.Quantity `json:"quantity,omitempty"`

	// Number of days supply per dispense
	ExpectedSupplyDuration *Duration `json:"expectedSupplyDuration,omitempty"`

	// Intended dispenser
	Performer *common.Reference `json:"performer,omitempty"`
}

type MedicationRequestDispenseRequestInitialFill struct {
	common.BackboneElement

	// First fill quantity
	Quantity *common.Quantity `json:"quantity,omitempty"`

	// First fill duration
	Duration *Duration `json:"duration,omitempty"`
}

type MedicationRequestSubstitution struct {
	common.BackboneElement

	// Whether substitution is allowed or not
	AllowedBoolean         *bool                   `json:"allowedBoolean,omitempty"`
	AllowedCodeableConcept *common.CodeableConcept `json:"allowedCodeableConcept,omitempty"`

	// Why should (not) substitution be made
	Reason *common.CodeableConcept `json:"reason,omitempty"`
}
