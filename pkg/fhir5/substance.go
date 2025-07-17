// Package fhir5 contains FHIR R5 resource definitions
package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// SubstanceIngredient represents an ingredient in a substance
type SubstanceIngredient struct {
	common.BackboneElement

	// Indication of whether this ingredient affects the therapeutic action of the drug
	IsActive *bool `json:"isActive,omitempty"`

	// The actual ingredient - either a substance (simple ingredient) or another medication
	ItemCodeableConcept *common.CodeableConcept `json:"itemCodeableConcept,omitempty"`
	ItemReference       *common.Reference       `json:"itemReference,omitempty"`

	// Specifies how many (or how much) of the items there are in this Substance
	Quantity *common.Quantity `json:"quantity,omitempty"`
}

// Substance represents a homogeneous material with a definite composition
type Substance struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Substance"

	// The level of granularity is defined by the category concepts in the value set
	Category []common.CodeableConcept `json:"category,omitempty"`

	// This could be a reference to an externally defined code
	Code common.CodeableReference `json:"code"`

	// A description of the substance - its appearance, handling requirements, and other usage notes
	Description *string `json:"description,omitempty"`

	// When the substance is no longer valid to use
	Expiry *string `json:"expiry,omitempty"`

	// Unique identifier for the substance
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// A substance can be composed of other substances
	Ingredient []SubstanceIngredient `json:"ingredient,omitempty"`

	// A boolean to indicate if this an instance of a substance or a kind of one (a definition)
	Instance bool `json:"instance"`

	// The amount of the substance
	Quantity *common.Quantity `json:"quantity,omitempty"`

	// A code to indicate if the substance is actively used
	Status *SubstanceStatus `json:"status,omitempty"`
}

// SubstanceStatus represents the status of a substance
type SubstanceStatus string

const (
	SubstanceStatusActive         SubstanceStatus = "active"
	SubstanceStatusInactive       SubstanceStatus = "inactive"
	SubstanceStatusEnteredInError SubstanceStatus = "entered-in-error"
)
