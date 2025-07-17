package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// MedicinalProductIngredient represents an ingredient of a manufactured item or pharmaceutical product
type MedicinalProductIngredient struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "MedicinalProductIngredient"

	// If the ingredient is a known or suspected allergen
	AllergenicIndicator *bool `json:"allergenicIndicator,omitempty"`

	// The identifier(s) of this Ingredient that are assigned by business processes
	Identifier *common.Identifier `json:"identifier,omitempty"`

	// Manufacturer of this Ingredient
	Manufacturer []common.Reference `json:"manufacturer,omitempty"`

	// Ingredient role e.g. Active ingredient, excipient
	Role common.CodeableConcept `json:"role"`

	// A specified substance that comprises this ingredient
	SpecifiedSubstance []MedicinalProductIngredientSpecifiedSubstance `json:"specifiedSubstance,omitempty"`

	// The ingredient substance
	Substance *MedicinalProductIngredientSubstance `json:"substance,omitempty"`
}

// MedicinalProductIngredientSpecifiedSubstance represents a specified substance that comprises this ingredient
type MedicinalProductIngredientSpecifiedSubstance struct {
	common.BackboneElement

	// The specified substance
	Code common.CodeableConcept `json:"code"`

	// The group of specified substance, e.g. group 1 to 4
	Group common.CodeableConcept `json:"group"`

	// Confidentiality level of the specified substance as the ingredient
	Confidentiality *common.CodeableConcept `json:"confidentiality,omitempty"`

	// Strength expressed in terms of a reference substance
	Strength []MedicinalProductIngredientSpecifiedSubstanceStrength `json:"strength,omitempty"`
}

// MedicinalProductIngredientSpecifiedSubstanceStrength represents strength expressed in terms of a reference substance
type MedicinalProductIngredientSpecifiedSubstanceStrength struct {
	common.BackboneElement

	// The quantity of substance in the unit of presentation, or in the volume (or mass) of the single pharmaceutical product or manufactured item
	Presentation *common.Ratio `json:"presentation,omitempty"`

	// A lower limit for the quantity of substance in the unit of presentation
	PresentationLowLimit *common.Ratio `json:"presentationLowLimit,omitempty"`

	// The strength per unitary volume (or mass)
	Concentration *common.Ratio `json:"concentration,omitempty"`

	// A lower limit for the strength per unitary volume (or mass)
	ConcentrationLowLimit *common.Ratio `json:"concentrationLowLimit,omitempty"`

	// For when strength is measured at a particular point or distance
	MeasurementPoint *string `json:"measurementPoint,omitempty"`

	// The country or countries for which the strength range applies
	Country []common.CodeableConcept `json:"country,omitempty"`

	// Strength expressed in terms of a reference substance
	ReferenceStrength []MedicinalProductIngredientSpecifiedSubstanceStrengthReferenceStrength `json:"referenceStrength,omitempty"`
}

// MedicinalProductIngredientSpecifiedSubstanceStrengthReferenceStrength represents strength expressed in terms of a reference substance
type MedicinalProductIngredientSpecifiedSubstanceStrengthReferenceStrength struct {
	common.BackboneElement

	// Relevant reference substance
	Substance *common.CodeableConcept `json:"substance,omitempty"`

	// Strength expressed in terms of a reference substance
	Strength common.Ratio `json:"strength"`

	// Strength expressed in terms of a reference substance
	StrengthLowLimit *common.Ratio `json:"strengthLowLimit,omitempty"`

	// For when strength is measured at a particular point or distance
	MeasurementPoint *string `json:"measurementPoint,omitempty"`

	// The country or countries for which the strength range applies
	Country []common.CodeableConcept `json:"country,omitempty"`
}

// MedicinalProductIngredientSubstance represents the ingredient substance
type MedicinalProductIngredientSubstance struct {
	common.BackboneElement

	// The ingredient substance
	Code common.CodeableConcept `json:"code"`

	// Quantity of the substance or specified substance present in the manufactured item or pharmaceutical product
	Strength []MedicinalProductIngredientSpecifiedSubstanceStrength `json:"strength,omitempty"`
}
