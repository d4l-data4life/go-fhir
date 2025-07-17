package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// NutritionProductStatus represents the status of a nutrition product
type NutritionProductStatus string

const (
	NutritionProductStatusActive         NutritionProductStatus = "active"
	NutritionProductStatusInactive       NutritionProductStatus = "inactive"
	NutritionProductStatusEnteredInError NutritionProductStatus = "entered-in-error"
)

// NutritionProductNutrient represents nutrient information
type NutritionProductNutrient struct {
	common.BackboneElement

	// The amount of nutrient expressed in one or more units
	Amount []Ratio `json:"amount,omitempty"`

	// The (relevant) nutrients in the product
	Item *CodeableReference `json:"item,omitempty"`
}

// NutritionProductIngredient represents ingredients contained in this product
type NutritionProductIngredient struct {
	common.BackboneElement

	// The amount of ingredient that is in the product
	Amount []Ratio `json:"amount,omitempty"`

	// The ingredient contained in the product
	Item CodeableReference `json:"item"`
}

// NutritionProductCharacteristic represents specifies descriptive properties of the nutrition product
type NutritionProductCharacteristic struct {
	common.BackboneElement

	// A code specifying which characteristic of the product is being described
	Type common.CodeableConcept `json:"type"`

	// The description of the characteristic
	ValueCodeableConcept *common.CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueString          *string                 `json:"valueString,omitempty"`
	ValueQuantity        *common.Quantity        `json:"valueQuantity,omitempty"`
	ValueBase64Binary    *string                 `json:"valueBase64Binary,omitempty"`
	ValueAttachment      *Attachment             `json:"valueAttachment,omitempty"`
	ValueBoolean         *bool                   `json:"valueBoolean,omitempty"`
}

// NutritionProductInstance represents conveys instance-level information about this product item
type NutritionProductInstance struct {
	common.BackboneElement

	// Biological source event identifier
	BiologicalSourceEvent *common.Identifier `json:"biologicalSourceEvent,omitempty"`

	// The time after which the product is no longer expected to be in proper condition
	Expiry *string `json:"expiry,omitempty"`

	// The identifier for the physical instance
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The identification of the batch or lot of the product
	LotNumber *string `json:"lotNumber,omitempty"`

	// The name for the specific product
	Name *string `json:"name,omitempty"`

	// The amount of items or instances that the resource considers
	Quantity *common.Quantity `json:"quantity,omitempty"`

	// The time after which the product is no longer expected to be in proper condition
	UseBy *string `json:"useBy,omitempty"`
}

// NutritionProduct represents a product used for nutritional purposes
type NutritionProduct struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "NutritionProduct"

	// Nutrition products can have different classifications
	Category []common.CodeableConcept `json:"category,omitempty"`

	// Specifies descriptive properties of the nutrition product
	Characteristic []NutritionProductCharacteristic `json:"characteristic,omitempty"`

	// The code assigned to the product
	Code *common.CodeableConcept `json:"code,omitempty"`

	// Ingredients contained in this product
	Ingredient []NutritionProductIngredient `json:"ingredient,omitempty"`

	// Conveys instance-level information about this product item
	Instance []NutritionProductInstance `json:"instance,omitempty"`

	// Allergens that are known or suspected to be a part of this nutrition product
	KnownAllergen []CodeableReference `json:"knownAllergen,omitempty"`

	// The organisation responsible for the product
	Manufacturer []common.Reference `json:"manufacturer,omitempty"`

	// Comments made about the product
	Note []Annotation `json:"note,omitempty"`

	// Nutrient information
	Nutrient []NutritionProductNutrient `json:"nutrient,omitempty"`

	// active | inactive | entered-in-error
	Status NutritionProductStatus `json:"status"`
}
