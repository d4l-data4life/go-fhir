// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// BiologicallyDerivedProductDispensePerformer represents who or what performed an action
type BiologicallyDerivedProductDispensePerformer struct {
	common.BackboneElement

	// Identifies the person responsible for the action
	Actor common.Reference `json:"actor"`

	// Identifies the function of the performer during the dispense
	Function *common.CodeableConcept `json:"function,omitempty"`
}

// BiologicallyDerivedProductDispense represents a biologically derived product dispense
type BiologicallyDerivedProductDispense struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "BiologicallyDerivedProductDispense"

	// The order or request that the dispense is fulfilling
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// Link to a resource identifying the physical location that the product was dispatched to
	Destination *common.Reference `json:"destination,omitempty"`

	// Unique instance identifiers assigned to a biologically derived product dispense
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The physical location where the dispense was performed
	Location *common.Reference `json:"location,omitempty"`

	// Indicates the type of matching associated with the dispense
	MatchStatus *common.CodeableConcept `json:"matchStatus,omitempty"`

	// Additional notes
	Note []Annotation `json:"note,omitempty"`

	// Indicates the relationship between the donor of the biologically derived product and the intended recipient
	OriginRelationshipType *common.CodeableConcept `json:"originRelationshipType,omitempty"`

	// A larger event of which this particular event is a component
	PartOf []common.Reference `json:"partOf,omitempty"`

	// A link to a resource representing the patient that the product is dispensed for
	Patient common.Reference `json:"patient"`

	// Indicates who or what performed an action
	Performer []BiologicallyDerivedProductDispensePerformer `json:"performer,omitempty"`

	// When the product was selected/ matched
	PreparedDate *string `json:"preparedDate,omitempty"`

	// A link to a resource identifying the biologically derived product that is being dispensed
	Product common.Reference `json:"product"`

	// The amount of product in the dispense
	Quantity *common.Quantity `json:"quantity,omitempty"`

	// A code specifying the state of the dispense event
	Status BiologicallyDerivedProductDispenseStatus `json:"status"`

	// Specific instructions for use
	UsageInstruction *string `json:"usageInstruction,omitempty"`

	// When the product was dispatched for clinical use
	WhenHandedOver *string `json:"whenHandedOver,omitempty"`
}

// BiologicallyDerivedProductDispenseStatus represents the status of a biologically derived product dispense
type BiologicallyDerivedProductDispenseStatus string

const (
	BiologicallyDerivedProductDispenseStatusPreparation    BiologicallyDerivedProductDispenseStatus = "preparation"
	BiologicallyDerivedProductDispenseStatusInProgress     BiologicallyDerivedProductDispenseStatus = "in-progress"
	BiologicallyDerivedProductDispenseStatusAllocated      BiologicallyDerivedProductDispenseStatus = "allocated"
	BiologicallyDerivedProductDispenseStatusIssued         BiologicallyDerivedProductDispenseStatus = "issued"
	BiologicallyDerivedProductDispenseStatusUnfulfilled    BiologicallyDerivedProductDispenseStatus = "unfulfilled"
	BiologicallyDerivedProductDispenseStatusReturned       BiologicallyDerivedProductDispenseStatus = "returned"
	BiologicallyDerivedProductDispenseStatusEnteredInError BiologicallyDerivedProductDispenseStatus = "entered-in-error"
	BiologicallyDerivedProductDispenseStatusUnknown        BiologicallyDerivedProductDispenseStatus = "unknown"
)
