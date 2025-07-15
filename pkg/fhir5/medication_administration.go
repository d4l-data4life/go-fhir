package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// MedicationAdministrationStatus represents the status of a medication administration
type MedicationAdministrationStatus string

const (
	MedicationAdministrationStatusInProgress     MedicationAdministrationStatus = "in-progress"
	MedicationAdministrationStatusNotDone        MedicationAdministrationStatus = "not-done"
	MedicationAdministrationStatusOnHold         MedicationAdministrationStatus = "on-hold"
	MedicationAdministrationStatusCompleted      MedicationAdministrationStatus = "completed"
	MedicationAdministrationStatusEnteredInError MedicationAdministrationStatus = "entered-in-error"
	MedicationAdministrationStatusStopped        MedicationAdministrationStatus = "stopped"
	MedicationAdministrationStatusUnknown        MedicationAdministrationStatus = "unknown"
)

// MedicationAdministrationPerformer represents the performer of the medication administration
type MedicationAdministrationPerformer struct {
	common.BackboneElement

	// Indicates who or what performed the medication administration
	Actor common.Reference `json:"actor"`

	// Distinguishes the type of involvement of the actor in the medication administration
	Function *common.CodeableConcept `json:"function,omitempty"`
}

// MedicationAdministrationDosage represents dosage information for medication administration
type MedicationAdministrationDosage struct {
	common.BackboneElement

	// Amount of medication per dose
	Dose *common.Quantity `json:"dose,omitempty"`

	// How drug should enter body
	Method *common.CodeableConcept `json:"method,omitempty"`

	// Amount of medication per unit of time
	RateRatio    *Ratio           `json:"rateRatio,omitempty"`
	RateQuantity *common.Quantity `json:"rateQuantity,omitempty"`

	// Body site to administer to
	Route *common.CodeableConcept `json:"route,omitempty"`

	// Body site to administer to
	Site *common.CodeableConcept `json:"site,omitempty"`

	// Free text dosage instructions
	Text *string `json:"text,omitempty"`
}

// MedicationAdministration represents the event of a patient consuming or otherwise being administered a medication
type MedicationAdministration struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "MedicationAdministration"

	// A plan that is fulfilled in whole or in part by this MedicationAdministration
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// The type of medication administration
	Category []common.CodeableConcept `json:"category,omitempty"`

	// The device that is to be used for the administration of the medication
	Device []CodeableReference `json:"device,omitempty"`

	// Describes the medication dosage information details
	Dosage *MedicationAdministrationDosage `json:"dosage,omitempty"`

	// The visit, admission, or other contact between patient and health care provider
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Event history references
	EventHistory []common.Reference `json:"eventHistory,omitempty"`

	// Business identifier for this medication administration
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// An indication that the full dose was not administered
	IsSubPotent *bool `json:"isSubPotent,omitempty"`

	// If only a code is specified, then it needs to be a code for a specific product
	Medication CodeableReference `json:"medication"`

	// Extra information about the medication administration
	Note []Annotation `json:"note,omitempty"`

	// A specific date/time or interval of time during which the administration took place
	OccurrenceDateTime *string        `json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod   *common.Period `json:"occurrencePeriod,omitempty"`
	OccurrenceTiming   *Timing        `json:"occurrenceTiming,omitempty"`

	// A larger event of which this particular event is a component or step
	PartOf []common.Reference `json:"partOf,omitempty"`

	// The performer of the medication treatment
	Performer []MedicationAdministrationPerformer `json:"performer,omitempty"`

	// A code, Condition or observation that supports why the medication was administered
	Reason []CodeableReference `json:"reason,omitempty"`

	// The date the occurrence of the MedicationAdministration was first captured in the record
	Recorded *string `json:"recorded,omitempty"`

	// The original request that this medication administration is based on
	Request *common.Reference `json:"request,omitempty"`

	// Will generally be set to show that the administration has been completed
	Status MedicationAdministrationStatus `json:"status"`

	// A code indicating why the administration was not performed
	StatusReason []common.CodeableConcept `json:"statusReason,omitempty"`

	// The person or animal or group receiving the medication
	Subject common.Reference `json:"subject"`

	// The reason or reasons why the full dose was not administered
	SubPotentReason []common.CodeableConcept `json:"subPotentReason,omitempty"`

	// Additional information that supports the administration of the medication
	SupportingInformation []common.Reference `json:"supportingInformation,omitempty"`
}
