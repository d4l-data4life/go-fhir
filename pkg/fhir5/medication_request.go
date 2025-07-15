// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// MedicationRequestDispenseRequestInitialFill represents initial fill details
type MedicationRequestDispenseRequestInitialFill struct {
	common.BackboneElement

	// The length of time that the first dispense is expected to last
	Duration *Duration `json:"duration,omitempty"`

	// The amount or quantity to provide as part of the first dispense
	Quantity *common.Quantity `json:"quantity,omitempty"`
}

// MedicationRequestDispenseRequest represents dispense details
type MedicationRequestDispenseRequest struct {
	common.BackboneElement

	// The minimum period of time that must occur between dispenses of the medication
	DispenseInterval *Duration `json:"dispenseInterval,omitempty"`

	// Indicates the intended performing Organization that will dispense the medication
	Dispenser *common.Reference `json:"dispenser,omitempty"`

	// Provides additional information to the dispenser
	DispenserInstruction []Annotation `json:"dispenserInstruction,omitempty"`

	// Provides information about the type of adherence packaging to be supplied
	DoseAdministrationAid *common.CodeableConcept `json:"doseAdministrationAid,omitempty"`

	// Expected supply duration
	ExpectedSupplyDuration *Duration `json:"expectedSupplyDuration,omitempty"`

	// Initial fill details
	InitialFill *MedicationRequestDispenseRequestInitialFill `json:"initialFill,omitempty"`

	// If displaying "number of authorized fills", add 1 to this number
	NumberOfRepeatsAllowed *int `json:"numberOfRepeatsAllowed,omitempty"`

	// The amount that is to be dispensed for one fill
	Quantity *common.Quantity `json:"quantity,omitempty"`

	// Prescription validity period
	ValidityPeriod *common.Period `json:"validityPeriod,omitempty"`
}

// MedicationRequestSubstitution represents substitution permissions
type MedicationRequestSubstitution struct {
	common.BackboneElement

	// Whether substitution is allowed
	AllowedBoolean         *bool                   `json:"allowedBoolean,omitempty"`
	AllowedCodeableConcept *common.CodeableConcept `json:"allowedCodeableConcept,omitempty"`

	// Indicates the reason for the substitution
	Reason *common.CodeableConcept `json:"reason,omitempty"`
}

// MedicationRequest represents an order or request for medication (R5)
type MedicationRequest struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "MedicationRequest"

	// The date when the prescription was initially written or authored on
	AuthoredOn *string `json:"authoredOn,omitempty"`

	// A plan or request that is fulfilled in whole or in part by this medication request
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// The category can be used to include where the medication is expected to be consumed
	Category []common.CodeableConcept `json:"category,omitempty"`

	// Course of therapy type
	CourseOfTherapyType *common.CodeableConcept `json:"courseOfTherapyType,omitempty"`

	// The intended type of device that is to be used for the administration
	Device []CodeableReference `json:"device,omitempty"`

	// Specific details for the dispense or medication supply part
	DispenseRequest *MedicationRequestDispenseRequest `json:"dispenseRequest,omitempty"`

	// If do not perform is not specified, the request is a positive request
	DoNotPerform *bool `json:"doNotPerform,omitempty"`

	// How to take the medication
	DosageInstruction []Dosage `json:"dosageInstruction,omitempty"`

	// The period over which the medication is to be taken (R5 feature)
	EffectiveDosePeriod *common.Period `json:"effectiveDosePeriod,omitempty"`

	// Encounter during which the request was created
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Event history references
	EventHistory []common.Reference `json:"eventHistory,omitempty"`

	// Shared identifier for multiple requests
	GroupIdentifier *common.Identifier `json:"groupIdentifier,omitempty"`

	// Business identifier
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Who provided the information about this request (R5 feature)
	InformationSource []common.Reference `json:"informationSource,omitempty"`

	// Associated insurance coverage
	Insurance []common.Reference `json:"insurance,omitempty"`

	// proposal | plan | order | original-order | reflex-order | filler-order | instance-order | option
	Intent MedicationRequestIntent `json:"intent"`

	// Medication to be taken (R5 uses CodeableReference)
	Medication CodeableReference `json:"medication"`

	// Extra information about the prescription
	Note []Annotation `json:"note,omitempty"`

	// The specified desired performer of the medication treatment (R5 allows multiple)
	Performer []common.Reference `json:"performer,omitempty"`

	// Performer type requirements
	PerformerType *common.CodeableConcept `json:"performerType,omitempty"`

	// routine | urgent | asap | stat
	Priority *MedicationRequestPriority `json:"priority,omitempty"`

	// Reference to an order/prescription that is being replaced
	PriorPrescription *common.Reference `json:"priorPrescription,omitempty"`

	// The reason for the order (R5 uses CodeableReference)
	Reason []CodeableReference `json:"reason,omitempty"`

	// The person who entered the order on behalf of another individual
	Recorder *common.Reference `json:"recorder,omitempty"`

	// Full representation of the dose of the medication (R5 feature)
	RenderedDosageInstruction *string `json:"renderedDosageInstruction,omitempty"`

	// If not populated, then assume that this is the original record
	Reported *bool `json:"reported,omitempty"`

	// The individual, organization, or device that initiated the request
	Requester *common.Reference `json:"requester,omitempty"`

	// active | on-hold | ended | stopped | completed | cancelled | entered-in-error | draft | unknown
	Status MedicationRequestStatus `json:"status"`

	// The date when the status was changed (R5 feature)
	StatusChanged *string `json:"statusChanged,omitempty"`

	// Reason for current status
	StatusReason *common.CodeableConcept `json:"statusReason,omitempty"`

	// The subject on a medication request is mandatory
	Subject common.Reference `json:"subject"`

	// Whether or not substitution can or should be part of the dispense
	Substitution *MedicationRequestSubstitution `json:"substitution,omitempty"`

	// Additional information that supports the ordering of the medication
	SupportingInformation []common.Reference `json:"supportingInformation,omitempty"`
}

// MedicationRequestIntent represents the intent of the medication request
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

// MedicationRequestStatus represents the status of the medication request
type MedicationRequestStatus string

const (
	MedicationRequestStatusActive         MedicationRequestStatus = "active"
	MedicationRequestStatusOnHold         MedicationRequestStatus = "on-hold"
	MedicationRequestStatusEnded          MedicationRequestStatus = "ended"
	MedicationRequestStatusStopped        MedicationRequestStatus = "stopped"
	MedicationRequestStatusCompleted      MedicationRequestStatus = "completed"
	MedicationRequestStatusCancelled      MedicationRequestStatus = "cancelled"
	MedicationRequestStatusEnteredInError MedicationRequestStatus = "entered-in-error"
	MedicationRequestStatusDraft          MedicationRequestStatus = "draft"
	MedicationRequestStatusUnknown        MedicationRequestStatus = "unknown"
)

// MedicationRequestPriority represents the priority of the medication request
type MedicationRequestPriority string

const (
	MedicationRequestPriorityRoutine MedicationRequestPriority = "routine"
	MedicationRequestPriorityUrgent  MedicationRequestPriority = "urgent"
	MedicationRequestPriorityASAP    MedicationRequestPriority = "asap"
	MedicationRequestPriorityStat    MedicationRequestPriority = "stat"
)
