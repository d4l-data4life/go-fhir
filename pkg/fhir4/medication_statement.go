package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// MedicationStatement represents a record of a medication that is being consumed by a patient
type MedicationStatement struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "MedicationStatement"

	// A plan, proposal or order that is fulfilled in whole or in part by this event
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// Indicates where the medication is expected to be consumed or administered
	Category *common.CodeableConcept `json:"category,omitempty"`

	// The encounter or episode of care that establishes the context for this MedicationStatement
	Context *common.Reference `json:"context,omitempty"`

	// The date when the medication statement was asserted by the information source
	DateAsserted *string `json:"dateAsserted,omitempty"`

	// Likely references would be to MedicationRequest, MedicationDispense, Claim, Observation or QuestionnaireAnswers
	DerivedFrom []common.Reference `json:"derivedFrom,omitempty"`

	// The dates included in the dosage on a Medication Statement reflect the dates for a given dose
	Dosage []common.Dosage `json:"dosage,omitempty"`

	// This attribute reflects the period over which the patient consumed the medication
	EffectiveDateTime *string `json:"effectiveDateTime,omitempty"`

	// This attribute reflects the period over which the patient consumed the medication
	EffectivePeriod *common.Period `json:"effectivePeriod,omitempty"`

	// This is a business identifier, not a resource identifier
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The person or organization that provided the information about the taking of this medication
	InformationSource *common.Reference `json:"informationSource,omitempty"`

	// If only a code is specified, then it needs to be a code for a specific product
	MedicationCodeableConcept *common.CodeableConcept `json:"medicationCodeableConcept,omitempty"`

	// If only a code is specified, then it needs to be a code for a specific product
	MedicationReference *common.Reference `json:"medicationReference,omitempty"`

	// Provides extra information about the medication statement that is not conveyed by the other attributes
	Note []common.Annotation `json:"note,omitempty"`

	// A larger event of which this particular event is a component or step
	PartOf []common.Reference `json:"partOf,omitempty"`

	// This could be a diagnosis code
	ReasonCode []common.CodeableConcept `json:"reasonCode,omitempty"`

	// This is a reference to a condition that is the reason why the medication is being/was taken
	ReasonReference []common.Reference `json:"reasonReference,omitempty"`

	// MedicationStatement is a statement at a point in time
	Status MedicationStatementStatus `json:"status"`

	// This is generally only used for "exception" statuses such as "not-taken", "on-hold", "cancelled" or "entered-in-error"
	StatusReason []common.CodeableConcept `json:"statusReason,omitempty"`

	// The person, animal or group who is/was taking the medication
	Subject common.Reference `json:"subject"`
}

// MedicationStatementStatus represents the status of the medication statement
type MedicationStatementStatus string

const (
	MedicationStatementStatusActive         MedicationStatementStatus = "active"
	MedicationStatementStatusCompleted      MedicationStatementStatus = "completed"
	MedicationStatementStatusEnteredInError MedicationStatementStatus = "entered-in-error"
	MedicationStatementStatusIntended       MedicationStatementStatus = "intended"
	MedicationStatementStatusStopped        MedicationStatementStatus = "stopped"
	MedicationStatementStatusOnHold         MedicationStatementStatus = "on-hold"
	MedicationStatementStatusUnknown        MedicationStatementStatus = "unknown"
	MedicationStatementStatusNotTaken       MedicationStatementStatus = "not-taken"
)
