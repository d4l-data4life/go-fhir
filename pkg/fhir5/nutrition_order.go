// Package fhir5 contains FHIR R5 resource definitions
package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// NutritionOrderIntent represents intent of nutrition order
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

// NutritionOrderStatus represents status of nutrition order
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

// NutritionOrderPriority represents priority of nutrition order
type NutritionOrderPriority string

const (
	NutritionOrderPriorityRoutine NutritionOrderPriority = "routine"
	NutritionOrderPriorityUrgent  NutritionOrderPriority = "urgent"
	NutritionOrderPriorityASAP    NutritionOrderPriority = "asap"
	NutritionOrderPriorityStat    NutritionOrderPriority = "stat"
)

// NutritionOrder represents a request for a diet, oral nutritional supplement, or enteral or parenteral formula
type NutritionOrder struct {
	DomainResource
	ResourceType           string                   `json:"resourceType"` // Always "NutritionOrder"
	AllergyIntolerance     []common.Reference       `json:"allergyIntolerance,omitempty"`
	BasedOn                []common.Reference       `json:"basedOn,omitempty"`
	DateTime               string                   `json:"dateTime"`
	Encounter              *common.Reference        `json:"encounter,omitempty"`
	EnteralFormula         *interface{}             `json:"enteralFormula,omitempty"`
	ExcludeFoodModifier    []common.CodeableConcept `json:"excludeFoodModifier,omitempty"`
	FoodPreferenceModifier []common.CodeableConcept `json:"foodPreferenceModifier,omitempty"`
	GroupIdentifier        *common.Identifier       `json:"groupIdentifier,omitempty"`
	Identifier             []common.Identifier      `json:"identifier,omitempty"`
	Instantiates           []string                 `json:"instantiates,omitempty"`
	InstantiatesCanonical  []string                 `json:"instantiatesCanonical,omitempty"`
	InstantiatesUri        []string                 `json:"instantiatesUri,omitempty"`
	Intent                 NutritionOrderIntent     `json:"intent"`
	Note                   []Annotation             `json:"note,omitempty"`
	OralDiet               *interface{}             `json:"oralDiet,omitempty"`
	Orderer                *common.Reference        `json:"orderer,omitempty"`
	OutsideFoodAllowed     *bool                    `json:"outsideFoodAllowed,omitempty"`
	Performer              []CodeableReference      `json:"performer,omitempty"`
	Priority               *NutritionOrderPriority  `json:"priority,omitempty"`
	Status                 NutritionOrderStatus     `json:"status"`
	Subject                common.Reference         `json:"subject"`
	Supplement             []interface{}            `json:"supplement,omitempty"`
	SupportingInformation  []common.Reference       `json:"supportingInformation,omitempty"`
}
