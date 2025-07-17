package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// MedicationDispense represents indicates that a medication product is to be or has been dispensed for a named person/patient
type MedicationDispense struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "MedicationDispense"

	// Maps to basedOn in Event logical model
	AuthorizingPrescription []common.Reference `json:"authorizingPrescription,omitempty"`

	// The category can be used to include where the medication is expected to be consumed or other types of dispenses
	Category *common.CodeableConcept `json:"category,omitempty"`

	// The encounter or episode of care that establishes the context for this event
	Context *common.Reference `json:"context,omitempty"`

	// The amount of medication expressed as a timing amount
	DaysSupply *common.Quantity `json:"daysSupply,omitempty"`

	// Identification of the facility/location where the medication was shipped to, as part of the dispense event
	Destination *common.Reference `json:"destination,omitempty"`

	// This element can include a detected issue that has been identified either by a decision support system or by a clinician
	DetectedIssue []common.Reference `json:"detectedIssue,omitempty"`

	// When the dose or rate is intended to change over the entire administration period
	DosageInstruction []common.Dosage `json:"dosageInstruction,omitempty"`

	// This might not include provenances for all versions of the request
	EventHistory []common.Reference `json:"eventHistory,omitempty"`

	// This is a business identifier, not a resource identifier
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The principal physical location where the dispense was performed
	Location *common.Reference `json:"location,omitempty"`

	// If only a code is specified, then it needs to be a code for a specific product
	MedicationCodeableConcept *common.CodeableConcept `json:"medicationCodeableConcept,omitempty"`

	// If only a code is specified, then it needs to be a code for a specific product
	MedicationReference *common.Reference `json:"medicationReference,omitempty"`

	// Extra information about the dispense that could not be conveyed in the other attributes
	Note []common.Annotation `json:"note,omitempty"`

	// The procedure that trigger the dispense
	PartOf []common.Reference `json:"partOf,omitempty"`

	// Indicates who or what performed the event
	Performer []MedicationDispensePerformer `json:"performer,omitempty"`

	// The amount of medication that has been dispensed. Includes unit of measure
	Quantity *common.Quantity `json:"quantity,omitempty"`

	// Identifies the person who picked up the medication
	Receiver []common.Reference `json:"receiver,omitempty"`

	// This element is labeled as a modifier because the status contains codes that mark the resource as not currently valid
	Status MedicationDispenseStatus `json:"status"`

	// Indicates the reason why a dispense was not performed
	StatusReasonCodeableConcept *common.CodeableConcept `json:"statusReasonCodeableConcept,omitempty"`

	// Indicates the reason why a dispense was not performed
	StatusReasonReference *common.Reference `json:"statusReasonReference,omitempty"`

	// SubstanceAdministration->subject->Patient
	Subject *common.Reference `json:"subject,omitempty"`

	// Indicates whether or not substitution was made as part of the dispense
	Substitution *MedicationDispenseSubstitution `json:"substitution,omitempty"`

	// Additional information that supports the medication being dispensed
	SupportingInformation []common.Reference `json:"supportingInformation,omitempty"`

	// Indicates the type of dispensing event that is performed
	Type *common.CodeableConcept `json:"type,omitempty"`

	// The time the dispensed product was provided to the patient or their representative
	WhenHandedOver *string `json:"whenHandedOver,omitempty"`

	// The time when the dispensed product was packaged and reviewed
	WhenPrepared *string `json:"whenPrepared,omitempty"`
}

// MedicationDispenseStatus represents the status of the medication dispense
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

// MedicationDispensePerformer represents indicates who or what performed the event
type MedicationDispensePerformer struct {
	common.BackboneElement

	// The device, practitioner, etc. who performed the action
	Actor common.Reference `json:"actor"`

	// Distinguishes the type of performer in the dispense
	Function *common.CodeableConcept `json:"function,omitempty"`
}

// MedicationDispenseSubstitution represents indicates whether or not substitution was made as part of the dispense
type MedicationDispenseSubstitution struct {
	common.BackboneElement

	// Indicates the reason for the substitution (or lack of substitution) from what was prescribed
	Reason []common.CodeableConcept `json:"reason,omitempty"`

	// The person or organization that has primary responsibility for the substitution
	ResponsibleParty []common.Reference `json:"responsibleParty,omitempty"`

	// A code signifying whether a different drug was dispensed from what was prescribed
	Type *common.CodeableConcept `json:"type,omitempty"`

	// True if the dispenser dispensed a different drug or product from what was prescribed
	WasSubstituted bool `json:"wasSubstituted"`
}
