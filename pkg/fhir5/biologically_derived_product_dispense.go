// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// BiologicallyDerivedProductDispensePerformer represents performer for BiologicallyDerivedProductDispense
type BiologicallyDerivedProductDispensePerformer struct {
	common.BackboneElement

	// The actor performing the dispense
	Actor *common.Reference `json:"actor,omitempty"`

	// The function of the actor in the dispense
	Function *common.CodeableConcept `json:"function,omitempty"`
}

// BiologicallyDerivedProductDispense represents a dispense event for a biologically derived product
type BiologicallyDerivedProductDispense struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "BiologicallyDerivedProductDispense"

	// The biologically derived product that is dispensed
	Product *common.Reference `json:"product,omitempty"`

	// The patient to whom the product is dispensed
	Subject *common.Reference `json:"subject,omitempty"`

	// The encounter or episode of care that establishes the context for this dispense
	Encounter *common.Reference `json:"encounter,omitempty"`

	// The amount dispensed
	Quantity *common.Quantity `json:"quantity,omitempty"`

	// When the product was dispensed
	DispensedDateTime *string `json:"dispensedDateTime,omitempty"`

	// When the product was dispensed
	DispensedPeriod *common.Period `json:"dispensedPeriod,omitempty"`

	// The individual responsible for dispensing the product
	Performer []BiologicallyDerivedProductDispensePerformer `json:"performer,omitempty"`

	// The location where the dispense occurred
	Location *common.Reference `json:"location,omitempty"`

	// Whether the dispense was or was not performed
	NotDone *bool `json:"notDone,omitempty"`

	// Why a dispense was not performed
	NotDoneReason *common.CodeableConcept `json:"notDoneReason,omitempty"`

	// Additional notes
	Note []Annotation `json:"note,omitempty"`

	// How the product should be used
	UsageInstruction *string `json:"usageInstruction,omitempty"`

	// prepared | dispensed | in-transit | delivered
	Status *string `json:"status,omitempty"`
}
