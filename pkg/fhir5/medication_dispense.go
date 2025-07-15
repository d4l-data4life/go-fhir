package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// MedicationDispenseStatus represents the status of a medication dispense
type MedicationDispenseStatus string

const (
	MedicationDispenseStatusPreparation    MedicationDispenseStatus = "preparation"
	MedicationDispenseStatusInProgress     MedicationDispenseStatus = "in-progress"
	MedicationDispenseStatusCancelled      MedicationDispenseStatus = "cancelled"
	MedicationDispenseStatusOnHold         MedicationDispenseStatus = "on-hold"
	MedicationDispenseStatusCompleted      MedicationDispenseStatus = "completed"
	MedicationDispenseStatusEnteredInError MedicationDispenseStatus = "entered-in-error"
	MedicationDispenseStatusStopped        MedicationDispenseStatus = "stopped"
	MedicationDispenseStatusDeclined       MedicationDispenseStatus = "declined"
	MedicationDispenseStatusUnknown        MedicationDispenseStatus = "unknown"
)

// MedicationDispensePerformer represents the performer of the medication dispense
type MedicationDispensePerformer struct {
	common.BackboneElement

	// Indicates who or what performed the event
	Actor common.Reference `json:"actor"`

	// Distinguishes the type of performer in the dispense
	Function *common.CodeableConcept `json:"function,omitempty"`
}

// MedicationDispenseSubstitution represents whether a substitution was made as part of the dispense
type MedicationDispenseSubstitution struct {
	common.BackboneElement

	// A code signifying whether a different drug was dispensed from what was prescribed
	Type *common.CodeableConcept `json:"type,omitempty"`

	// Indicates the reason for the substitution
	Reason []common.CodeableConcept `json:"reason,omitempty"`

	// The person or organization that has primary responsibility for the substitution
	ResponsibleParty *common.Reference `json:"responsibleParty,omitempty"`

	// True if the dispenser dispensed a different drug or product from what was prescribed
	WasSubstituted *bool `json:"wasSubstituted,omitempty"`
}

// MedicationDispense represents the provision of a medication to a patient
type MedicationDispense struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "MedicationDispense"

	// Indicates the medication order that is being dispensed against
	AuthorizingPrescription []common.Reference `json:"authorizingPrescription,omitempty"`

	// A plan that is fulfilled in whole or in part by this MedicationDispense
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// The category can be used to include where the medication is expected to be consumed
	Category []common.CodeableConcept `json:"category,omitempty"`

	// The amount of medication expressed as a timing amount
	DaysSupply *common.Quantity `json:"daysSupply,omitempty"`

	// Identification of the facility/location where the medication was/will be shipped to
	Destination *common.Reference `json:"destination,omitempty"`

	// Indicates an actual or potential clinical issue with or between one or more active or proposed clinical actions
	DetectedIssue []common.Reference `json:"detectedIssue,omitempty"`

	// How the medication is to be used by the patient
	DosageInstruction []Dosage `json:"dosageInstruction,omitempty"`

	// The encounter that establishes the context for this event
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Event history references
	EventHistory []common.Reference `json:"eventHistory,omitempty"`

	// Business identifier for this medication dispense
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Where the medication was/will be dispensed to
	Location *common.Reference `json:"location,omitempty"`

	// What medication was supplied
	Medication CodeableReference `json:"medication"`

	// Information about the dispense
	Note []Annotation `json:"note,omitempty"`

	// A larger event of which this particular event is a component or step
	PartOf []common.Reference `json:"partOf,omitempty"`

	// Who performed the dispense and what they did
	Performer []MedicationDispensePerformer `json:"performer,omitempty"`

	// The amount of medication that has been dispensed
	Quantity *common.Quantity `json:"quantity,omitempty"`

	// Who collected the medication or where the medication was delivered
	Receiver []common.Reference `json:"receiver,omitempty"`

	// The date the dispensed product was recorded
	Recorded *string `json:"recorded,omitempty"`

	// The status of the dispense
	Status MedicationDispenseStatus `json:"status"`

	// Indicates the reason why a dispense was not performed
	StatusReason *CodeableReference `json:"statusReason,omitempty"`

	// A link to a resource representing the person or the group to whom the medication will be given
	Subject common.Reference `json:"subject"`

	// Whether a substitution was made as part of the dispense
	Substitution *MedicationDispenseSubstitution `json:"substitution,omitempty"`

	// Additional information that supports the medication being dispensed
	SupportingInformation []common.Reference `json:"supportingInformation,omitempty"`

	// The type of medication dispense
	Type *common.CodeableConcept `json:"type,omitempty"`

	// The time when the dispensed product was packaged and reviewed
	WhenHandedOver *string `json:"whenHandedOver,omitempty"`

	// The time the dispensed product was prepared
	WhenPrepared *string `json:"whenPrepared,omitempty"`
}
