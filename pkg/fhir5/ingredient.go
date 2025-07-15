package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// IngredientStatus represents the status of an ingredient
type IngredientStatus string

const (
	IngredientStatusDraft   IngredientStatus = "draft"
	IngredientStatusActive  IngredientStatus = "active"
	IngredientStatusRetired IngredientStatus = "retired"
	IngredientStatusUnknown IngredientStatus = "unknown"
)

// IngredientManufacturer represents the organization(s) that manufacture this ingredient
type IngredientManufacturer struct {
	common.BackboneElement

	// The organization that manufactures this ingredient
	Manufacturer string `json:"manufacturer"`

	// The way in which this manufacturer is associated with the ingredient
	Role *common.CodeableConcept `json:"role,omitempty"`
}

// IngredientSubstanceStrengthReferenceStrength represents strength expressed in terms of a reference substance
type IngredientSubstanceStrengthReferenceStrength struct {
	common.BackboneElement

	// The country or countries for which the strength range applies
	Country []common.CodeableConcept `json:"country,omitempty"`

	// For when strength is measured at a particular point or distance
	MeasurementPoint *string `json:"measurementPoint,omitempty"`

	// Strength expressed in terms of a reference substance
	StrengthRatio    *Ratio           `json:"strengthRatio,omitempty"`
	StrengthQuantity *common.Quantity `json:"strengthQuantity,omitempty"`

	// Relevant reference substance
	Substance CodeableReference `json:"substance"`
}

// IngredientSubstanceStrength represents the quantity of substance in the unit of presentation
type IngredientSubstanceStrength struct {
	common.BackboneElement

	// A code that indicates if the strength is, for example, based on the ingredient substance as stated or on the substance base
	Basis *common.CodeableConcept `json:"basis,omitempty"`

	// The strength per unitary volume (or mass)
	ConcentrationRatio           *Ratio                  `json:"concentrationRatio,omitempty"`
	ConcentrationCodeableConcept *common.CodeableConcept `json:"concentrationCodeableConcept,omitempty"`
	ConcentrationQuantity        *common.Quantity        `json:"concentrationQuantity,omitempty"`

	// The country or countries for which the strength range applies
	Country []common.CodeableConcept `json:"country,omitempty"`

	// For when strength is measured at a particular point or distance
	MeasurementPoint *string `json:"measurementPoint,omitempty"`

	// The quantity of substance in the unit of presentation
	PresentationRatio           *Ratio                  `json:"presentationRatio,omitempty"`
	PresentationCodeableConcept *common.CodeableConcept `json:"presentationCodeableConcept,omitempty"`
	PresentationQuantity        *common.Quantity        `json:"presentationQuantity,omitempty"`

	// Strength expressed in terms of a reference substance
	ReferenceStrength []IngredientSubstanceStrengthReferenceStrength `json:"referenceStrength,omitempty"`

	// A textual represention of either the whole of the concentration strength or a part of it
	TextConcentration *string `json:"textConcentration,omitempty"`

	// A textual represention of either the whole of the presentation strength or a part of it
	TextPresentation *string `json:"textPresentation,omitempty"`
}

// IngredientSubstance represents the substance that comprises this ingredient
type IngredientSubstance struct {
	common.BackboneElement

	// A code or full resource that represents the ingredient's substance
	Code CodeableReference `json:"code"`

	// The quantity of substance in the unit of presentation
	Strength []IngredientSubstanceStrength `json:"strength,omitempty"`
}

// Ingredient represents an ingredient of a manufactured item or pharmaceutical product
type Ingredient struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Ingredient"

	// If the ingredient is a known or suspected allergen
	AllergenicIndicator *bool `json:"allergenicIndicator,omitempty"`

	// A place for providing any notes that are relevant to the component
	Comment *string `json:"comment,omitempty"`

	// The product which this ingredient is a constituent part of
	For []common.Reference `json:"for,omitempty"`

	// A classification of the ingredient identifying its precise purpose(s) in the drug product
	Function []common.CodeableConcept `json:"function,omitempty"`

	// A classification of the ingredient according to where in the physical item it tends to be used
	Group *common.CodeableConcept `json:"group,omitempty"`

	// The identifier(s) of this Ingredient that are assigned by business processes
	Identifier *common.Identifier `json:"identifier,omitempty"`

	// The organization(s) that manufacture this ingredient
	Manufacturer []IngredientManufacturer `json:"manufacturer,omitempty"`

	// A classification of the ingredient identifying its purpose within the product
	Role common.CodeableConcept `json:"role"`

	// Allows filtering of ingredient that are appropriate for use versus not
	Status IngredientStatus `json:"status"`

	// The substance that comprises this ingredient
	Substance IngredientSubstance `json:"substance"`
}
