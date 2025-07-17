package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// MedicationAdministration represents describes the event of a patient consuming or otherwise being administered a medication
type MedicationAdministration struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "MedicationAdministration"

	// Indicates where the medication is expected to be consumed or administered
	Category *common.CodeableConcept `json:"category,omitempty"`

	// The visit, admission, or other contact between patient and health care provider during which the medication administration was performed
	Context *common.Reference `json:"context,omitempty"`

	// The device used in administering the medication to the patient
	Device []common.Reference `json:"device,omitempty"`

	// Describes the medication dosage information details e.g. dose, rate, site, route, etc
	Dosage *MedicationAdministrationDosage `json:"dosage,omitempty"`

	// A specific date/time or interval of time during which the administration took place
	EffectiveDateTime *string `json:"effectiveDateTime,omitempty"`

	// A specific date/time or interval of time during which the administration took place
	EffectivePeriod *common.Period `json:"effectivePeriod,omitempty"`

	// This might not include provenances for all versions of the request
	EventHistory []common.Reference `json:"eventHistory,omitempty"`

	// This is a business identifier, not a resource identifier
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// A protocol, guideline, orderset, or other definition that was adhered to in whole or in part by this event
	Instantiates []string `json:"instantiates,omitempty"`

	// If only a code is specified, then it needs to be a code for a specific product
	MedicationCodeableConcept *common.CodeableConcept `json:"medicationCodeableConcept,omitempty"`

	// If only a code is specified, then it needs to be a code for a specific product
	MedicationReference *common.Reference `json:"medicationReference,omitempty"`

	// Extra information about the medication administration that is not conveyed by the other attributes
	Note []common.Annotation `json:"note,omitempty"`

	// A larger event of which this particular event is a component or step
	PartOf []common.Reference `json:"partOf,omitempty"`

	// Indicates who or what performed the medication administration and how they were involved
	Performer []MedicationAdministrationPerformer `json:"performer,omitempty"`

	// A code indicating why the medication was given
	ReasonCode []common.CodeableConcept `json:"reasonCode,omitempty"`

	// This is a reference to a condition that is the reason for the medication request
	ReasonReference []common.Reference `json:"reasonReference,omitempty"`

	// This is a reference to the MedicationRequest where the intent is either order or instance-order
	Request *common.Reference `json:"request,omitempty"`

	// This element is labeled as a modifier because the status contains codes that mark the resource as not currently valid
	Status MedicationAdministrationStatus `json:"status"`

	// A code indicating why the administration was not performed
	StatusReason []common.CodeableConcept `json:"statusReason,omitempty"`

	// The person or animal or group receiving the medication
	Subject common.Reference `json:"subject"`

	// Additional information (for example, patient height and weight) that supports the administration of the medication
	SupportingInformation []common.Reference `json:"supportingInformation,omitempty"`
}

// MedicationAdministrationStatus represents the status of the medication administration
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

// MedicationAdministrationPerformer represents indicates who or what performed the medication administration and how they were involved
type MedicationAdministrationPerformer struct {
	common.BackboneElement

	// Indicates who or what performed the medication administration
	Actor common.Reference `json:"actor"`

	// Distinguishes the type of involvement of the performer in the medication administration
	Function *common.CodeableConcept `json:"function,omitempty"`
}

// MedicationAdministrationDosage represents describes the medication dosage information details
type MedicationAdministrationDosage struct {
	common.BackboneElement

	// The amount of the medication given at one administration event
	Dose *common.Quantity `json:"dose,omitempty"`

	// The method of drug delivery (e.g. oral, intravenous, intramuscular, etc.)
	Method *common.CodeableConcept `json:"method,omitempty"`

	// The rate of medication delivery (e.g. 100ml per hour)
	RateQuantity *common.Quantity `json:"rateQuantity,omitempty"`

	// The rate of medication delivery (e.g. 100ml per hour)
	RateRatio *common.Ratio `json:"rateRatio,omitempty"`

	// The path by which the drug product is taken into or makes contact with the body
	Route *common.CodeableConcept `json:"route,omitempty"`

	// Body site to administer to
	Site *common.CodeableConcept `json:"site,omitempty"`

	// Free text dosage instructions e.g. SIG
	Text *string `json:"text,omitempty"`
}
