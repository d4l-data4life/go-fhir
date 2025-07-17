// Package fhir5 contains FHIR R5 resource definitions
package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// NutritionIntakeStatus represents status of nutrition intake
type NutritionIntakeStatus string

const (
	NutritionIntakeStatusActive         NutritionIntakeStatus = "active"
	NutritionIntakeStatusCompleted      NutritionIntakeStatus = "completed"
	NutritionIntakeStatusEnteredInError NutritionIntakeStatus = "entered-in-error"
	NutritionIntakeStatusIntended       NutritionIntakeStatus = "intended"
	NutritionIntakeStatusOnHold         NutritionIntakeStatus = "on-hold"
	NutritionIntakeStatusStopped        NutritionIntakeStatus = "stopped"
	NutritionIntakeStatusUnknown        NutritionIntakeStatus = "unknown"
)

// NutritionIntakeConsumedItem represents an item that was consumed
type NutritionIntakeConsumedItem struct {
	common.BackboneElement
	Amount            *common.Quantity        `json:"amount,omitempty"`
	NotConsumed       *bool                   `json:"notConsumed,omitempty"`
	NotConsumedReason *common.CodeableConcept `json:"notConsumedReason,omitempty"`
	NutritionProduct  CodeableReference       `json:"nutritionProduct"`
	Rate              *common.Quantity        `json:"rate,omitempty"`
	Schedule          *Timing                 `json:"schedule,omitempty"`
	Type              common.CodeableConcept  `json:"type"`
}

// NutritionIntakeIngredientLabel represents ingredient label information
type NutritionIntakeIngredientLabel struct {
	common.BackboneElement
	Amount   common.Quantity   `json:"amount"`
	Nutrient CodeableReference `json:"nutrient"`
}

// NutritionIntakePerformer represents who performed the nutrition intake
type NutritionIntakePerformer struct {
	common.BackboneElement
	Actor    common.Reference        `json:"actor"`
	Function *common.CodeableConcept `json:"function,omitempty"`
}

// NutritionIntake represents a record of food or fluid being taken by a patient
type NutritionIntake struct {
	DomainResource
	ResourceType          string                           `json:"resourceType"` // Always "NutritionIntake"
	BasedOn               []common.Reference               `json:"basedOn,omitempty"`
	Code                  *common.CodeableConcept          `json:"code,omitempty"`
	ConsumedItem          []NutritionIntakeConsumedItem    `json:"consumedItem"`
	DerivedFrom           []common.Reference               `json:"derivedFrom,omitempty"`
	Encounter             *common.Reference                `json:"encounter,omitempty"`
	Identifier            []common.Identifier              `json:"identifier,omitempty"`
	IngredientLabel       []NutritionIntakeIngredientLabel `json:"ingredientLabel,omitempty"`
	InstantiatesCanonical []string                         `json:"instantiatesCanonical,omitempty"`
	InstantiatesUri       []string                         `json:"instantiatesUri,omitempty"`
	Location              *common.Reference                `json:"location,omitempty"`
	Note                  []Annotation                     `json:"note,omitempty"`
	OccurrenceDateTime    *string                          `json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod      *common.Period                   `json:"occurrencePeriod,omitempty"`
	PartOf                []common.Reference               `json:"partOf,omitempty"`
	Performer             []NutritionIntakePerformer       `json:"performer,omitempty"`
	Reason                []CodeableReference              `json:"reason,omitempty"`
	Recorded              *string                          `json:"recorded,omitempty"`
	ReportedBoolean       *bool                            `json:"reportedBoolean,omitempty"`
	ReportedReference     *common.Reference                `json:"reportedReference,omitempty"`
	Status                NutritionIntakeStatus            `json:"status"`
	StatusReason          []common.CodeableConcept         `json:"statusReason,omitempty"`
	Subject               common.Reference                 `json:"subject"`
}
