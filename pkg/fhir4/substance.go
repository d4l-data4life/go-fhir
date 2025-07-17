package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// Substance represents a homogeneous material with a definite composition
type Substance struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Substance"

	// The level of granularity is defined by the category concepts in the value set
	Category []common.CodeableConcept `json:"category,omitempty"`

	// This could be a reference to an externally defined code
	Code common.CodeableConcept `json:"code"`

	// A description of the substance - its appearance, handling requirements, and other usage notes
	Description *string `json:"description,omitempty"`

	// This identifier is associated with the kind of substance in contrast to the Substance.instance.identifier
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// A substance can be composed of other substances
	Ingredient []SubstanceIngredient `json:"ingredient,omitempty"`

	// Substance may be used to describe a kind of substance, or a specific package/container of the substance
	Instance []SubstanceInstance `json:"instance,omitempty"`

	// A code to indicate if the substance is actively used
	Status *SubstanceStatus `json:"status,omitempty"`
}

// SubstanceStatus represents a code to indicate if the substance is actively used
type SubstanceStatus string

const (
	SubstanceStatusActive         SubstanceStatus = "active"
	SubstanceStatusInactive       SubstanceStatus = "inactive"
	SubstanceStatusEnteredInError SubstanceStatus = "entered-in-error"
)

// SubstanceInstance represents substance may be used to describe a kind of substance, or a specific package/container of the substance
type SubstanceInstance struct {
	common.BackboneElement

	// When the substance is no longer valid to use
	Expiry *string `json:"expiry,omitempty"`

	// Identifier associated with the package/container (usually a label affixed directly)
	Identifier *common.Identifier `json:"identifier,omitempty"`

	// The amount of the substance
	Quantity *common.Quantity `json:"quantity,omitempty"`
}

// SubstanceIngredient represents a substance can be composed of other substances
type SubstanceIngredient struct {
	common.BackboneElement

	// The amount of the ingredient in the substance - a concentration ratio
	Quantity *common.Ratio `json:"quantity,omitempty"`

	// Another substance that is a component of this substance
	SubstanceCodeableConcept *common.CodeableConcept `json:"substanceCodeableConcept,omitempty"`

	// Another substance that is a component of this substance
	SubstanceReference *common.Reference `json:"substanceReference,omitempty"`
}
