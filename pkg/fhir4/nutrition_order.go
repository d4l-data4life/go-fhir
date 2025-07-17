package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// NutritionOrder represents a request to supply a diet, formula feeding (enteral) or oral nutritional supplement
type NutritionOrder struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "NutritionOrder"

	// Information on a patient's food allergies and intolerances
	AllergyIntolerance []common.Reference `json:"allergyIntolerance,omitempty"`

	// The date and time that this nutrition order was requested
	DateTime string `json:"dateTime"`

	// An encounter that provides additional information about the healthcare context
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Feeding provided through the gastrointestinal tract via a tube, catheter, or stoma
	EnteralFormula *NutritionOrderEnteralFormula `json:"enteralFormula,omitempty"`

	// Information on a patient's food allergies, intolerances and preferences
	ExcludeFoodModifier []common.CodeableConcept `json:"excludeFoodModifier,omitempty"`

	// Information on a patient's food preferences
	FoodPreferenceModifier []common.CodeableConcept `json:"foodPreferenceModifier,omitempty"`

	// The Identifier.type element can be to indicate filler vs. placer if needed
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The URL pointing to a protocol, guideline, orderset or other definition
	Instantiates []string `json:"instantiates,omitempty"`

	// Note: This is a business identifier, not a resource identifier
	InstantiatesCanonical []string `json:"instantiatesCanonical,omitempty"`

	// This might be an HTML page, PDF, etc. or could just be a non-resolvable URI identifier
	InstantiatesUri []string `json:"instantiatesUri,omitempty"`

	// When resources map to this element, they are free to define as many codes as necessary
	Intent NutritionOrderIntent `json:"intent"`

	// This element SHALL NOT be used to supply free text instructions for the diet
	Note []common.Annotation `json:"note,omitempty"`

	// Diet given orally in contrast to enteral (tube) feeding
	OralDiet *NutritionOrderOralDiet `json:"oralDiet,omitempty"`

	// The practitioner that holds legal responsibility for ordering the diet
	Orderer *common.Reference `json:"orderer,omitempty"`

	// The person (patient) who needs the nutrition order
	Patient common.Reference `json:"patient"`

	// Typically the system placing the order sets the status to "requested"
	Status NutritionOrderStatus `json:"status"`

	// Oral nutritional products given in order to add further nutritional value
	Supplement []NutritionOrderSupplement `json:"supplement,omitempty"`
}

// NutritionOrderIntent represents the intent of the nutrition order
type NutritionOrderIntent string

const (
	NutritionOrderIntentProposal      NutritionOrderIntent = "proposal"
	NutritionOrderIntentPlan          NutritionOrderIntent = "plan"
	NutritionOrderIntentDirective     NutritionOrderIntent = "directive"
	NutritionOrderIntentOrder         NutritionOrderIntent = "order"
	NutritionOrderIntentOriginalOrder NutritionOrderIntent = "original-order"
	NutritionOrderIntentReflexOrder   NutritionOrderIntent = "reflex-order"
	NutritionOrderIntentFillerOrder   NutritionOrderIntent = "filler-order"
	NutritionOrderIntentInstanceOrder NutritionOrderIntent = "instance-order"
	NutritionOrderIntentOption        NutritionOrderIntent = "option"
)

// NutritionOrderStatus represents the status of the nutrition order
type NutritionOrderStatus string

const (
	NutritionOrderStatusDraft          NutritionOrderStatus = "draft"
	NutritionOrderStatusActive         NutritionOrderStatus = "active"
	NutritionOrderStatusOnHold         NutritionOrderStatus = "on-hold"
	NutritionOrderStatusRevoked        NutritionOrderStatus = "revoked"
	NutritionOrderStatusCompleted      NutritionOrderStatus = "completed"
	NutritionOrderStatusEnteredInError NutritionOrderStatus = "entered-in-error"
	NutritionOrderStatusUnknown        NutritionOrderStatus = "unknown"
)

// NutritionOrderOralDiet represents diet given orally in contrast to enteral (tube) feeding
type NutritionOrderOralDiet struct {
	common.BackboneElement

	// The required consistency of liquids or fluids served to the patient
	FluidConsistencyType []common.CodeableConcept `json:"fluidConsistencyType,omitempty"`

	// Free text dosage instructions can be used for cases where the instructions are too complex to code
	Instruction *string `json:"instruction,omitempty"`

	// Class that defines the quantity and type of nutrient modifications
	Nutrient []NutritionOrderOralDietNutrient `json:"nutrient,omitempty"`

	// The time period and frequency at which the diet should be given
	Schedule []common.Timing `json:"schedule,omitempty"`

	// Class that describes any texture modifications required for the patient
	Texture []NutritionOrderOralDietTexture `json:"texture,omitempty"`

	// The kind of diet or dietary restriction such as fiber restricted diet or diabetic diet
	Type []common.CodeableConcept `json:"type,omitempty"`
}

// NutritionOrderOralDietNutrient represents class that defines the quantity and type of nutrient modifications
type NutritionOrderOralDietNutrient struct {
	common.BackboneElement

	// The quantity of the specified nutrient to include in diet
	Amount *common.Quantity `json:"amount,omitempty"`

	// The nutrient that is being modified such as carbohydrate or sodium
	Modifier *common.CodeableConcept `json:"modifier,omitempty"`
}

// NutritionOrderOralDietTexture represents class that describes any texture modifications
type NutritionOrderOralDietTexture struct {
	common.BackboneElement

	// Coupled with the texture.modifier; could be (All Foods)
	FoodType *common.CodeableConcept `json:"foodType,omitempty"`

	// Coupled with the foodType (Meat)
	Modifier *common.CodeableConcept `json:"modifier,omitempty"`
}

// NutritionOrderSupplement represents oral nutritional products given in order to add further nutritional value
type NutritionOrderSupplement struct {
	common.BackboneElement

	// Free text dosage instructions can be used for cases where the instructions are too complex to code
	Instruction *string `json:"instruction,omitempty"`

	// The product or brand name of the nutritional supplement
	ProductName *string `json:"productName,omitempty"`

	// The amount of the nutritional supplement to be given
	Quantity *common.Quantity `json:"quantity,omitempty"`

	// The time period and frequency at which the supplement(s) should be given
	Schedule []common.Timing `json:"schedule,omitempty"`

	// The kind of nutritional supplement product required
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// NutritionOrderEnteralFormula represents feeding provided through the gastrointestinal tract
type NutritionOrderEnteralFormula struct {
	common.BackboneElement

	// The product or brand name of the type of modular component to be added to the formula
	AdditiveProductName *string `json:"additiveProductName,omitempty"`

	// Indicates the type of modular component such as protein, carbohydrate, fat or fiber
	AdditiveType *common.CodeableConcept `json:"additiveType,omitempty"`

	// See implementation notes below for further discussion on how to order continuous vs bolus enteral feeding
	Administration []NutritionOrderEnteralFormulaAdministration `json:"administration,omitempty"`

	// Free text dosage instructions can be used for cases where the instructions are too complex to code
	AdministrationInstruction *string `json:"administrationInstruction,omitempty"`

	// The product or brand name of the enteral or infant formula product
	BaseFormulaProductName *string `json:"baseFormulaProductName,omitempty"`

	// The type of enteral or infant formula such as an adult standard formula with fiber
	BaseFormulaType *common.CodeableConcept `json:"baseFormulaType,omitempty"`

	// The amount of energy (calories) that the formula should provide per specified volume
	CaloricDensity *common.Quantity `json:"caloricDensity,omitempty"`

	// The maximum total quantity of formula that may be administered to a subject over the period of time
	MaxVolumeToDeliver *common.Quantity `json:"maxVolumeToDeliver,omitempty"`

	// The route or physiological path of administration into the patient's gastrointestinal tract
	RouteofAdministration *common.CodeableConcept `json:"routeofAdministration,omitempty"`
}

// NutritionOrderEnteralFormulaAdministration represents see implementation notes below for further discussion
type NutritionOrderEnteralFormulaAdministration struct {
	common.BackboneElement

	// The volume of formula to provide to the patient per the specified administration schedule
	Quantity *common.Quantity `json:"quantity,omitempty"`

	// Ratio is used when the quantity value in the denominator is not "1"
	RateQuantity *common.Quantity `json:"rateQuantity,omitempty"`

	// Ratio is used when the quantity value in the denominator is not "1"
	RateRatio *common.Ratio `json:"rateRatio,omitempty"`

	// The time period and frequency at which the enteral formula should be delivered to the patient
	Schedule *common.Timing `json:"schedule,omitempty"`
}
