// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// InsurancePlanStatus represents the status of an insurance plan
type InsurancePlanStatus string

const (
	InsurancePlanStatusDraft   InsurancePlanStatus = "draft"
	InsurancePlanStatusActive  InsurancePlanStatus = "active"
	InsurancePlanStatusRetired InsurancePlanStatus = "retired"
	InsurancePlanStatusUnknown InsurancePlanStatus = "unknown"
)

// InsurancePlanCoverageBenefitLimit represents the specific limits on the benefit
type InsurancePlanCoverageBenefitLimit struct {
	common.BackboneElement

	// The specific limit on the benefit
	Code *common.CodeableConcept `json:"code,omitempty"`

	// May also be called "eligible expense," "payment allowance," or "negotiated rate"
	Value *common.Quantity `json:"value,omitempty"`
}

// InsurancePlanCoverageBenefit represents specific benefits under this type of coverage
type InsurancePlanCoverageBenefit struct {
	common.BackboneElement

	// The specific limits on the benefit
	Limit []InsurancePlanCoverageBenefitLimit `json:"limit,omitempty"`

	// The referral requirements to have access/coverage for this benefit
	Requirement *string `json:"requirement,omitempty"`

	// Type of benefit (primary care; speciality care; inpatient; outpatient)
	Type common.CodeableConcept `json:"type"`
}

// InsurancePlanCoverage represents details about the coverage offered by the insurance product
type InsurancePlanCoverage struct {
	common.BackboneElement

	// Specific benefits under this type of coverage
	Benefit []InsurancePlanCoverageBenefit `json:"benefit,omitempty"`

	// Networks and tiers
	Network []common.Reference `json:"network,omitempty"`

	// Type of coverage (Medical; Dental; Mental Health; Substance Abuse; Vision; Drug; Short Term; Long Term Care; Hospice; Home Health)
	Type common.CodeableConcept `json:"type"`
}

// InsurancePlanPlanGeneralCost represents general costs associated with the plan
type InsurancePlanPlanGeneralCost struct {
	common.BackboneElement

	// Additional information about the general costs
	Comment *string `json:"comment,omitempty"`

	// Value of the cost
	Cost *Money `json:"cost,omitempty"`

	// Number of enrollees
	GroupSize *int `json:"groupSize,omitempty"`

	// Type of cost (copay; individual-deductible; family-deductible; out-of-pocket-maximum)
	Type common.CodeableConcept `json:"type"`
}

// InsurancePlanPlanSpecificCostBenefitCost represents costs associated with the benefit
type InsurancePlanPlanSpecificCostBenefitCost struct {
	common.BackboneElement

	// Additional information about the cost
	Applicability *common.CodeableConcept `json:"applicability,omitempty"`

	// Qualifiers that apply to the cost
	Qualifiers []common.CodeableConcept `json:"qualifiers,omitempty"`

	// Type of cost (copay; individual-deductible; family-deductible; out-of-pocket-maximum)
	Type common.CodeableConcept `json:"type"`

	// The actual cost value
	Value *common.Quantity `json:"value,omitempty"`
}

// InsurancePlanPlanSpecificCostBenefit represents benefits and their costs for a specific coverage area
type InsurancePlanPlanSpecificCostBenefit struct {
	common.BackboneElement

	// List of the costs associated with this benefit
	Cost []InsurancePlanPlanSpecificCostBenefitCost `json:"cost,omitempty"`

	// Type of the benefit
	Type common.CodeableConcept `json:"type"`
}

// InsurancePlanPlanSpecificCost represents costs for a specific coverage area
type InsurancePlanPlanSpecificCost struct {
	common.BackboneElement

	// Benefits and their costs for a specific coverage area
	Benefit []InsurancePlanPlanSpecificCostBenefit `json:"benefit,omitempty"`

	// General category of benefit
	Category common.CodeableConcept `json:"category"`
}

// InsurancePlanPlan represents details about an insurance plan
type InsurancePlanPlan struct {
	common.BackboneElement

	// The geographic region in which a health insurance plan's benefits apply
	CoverageArea []common.Reference `json:"coverageArea,omitempty"`

	// Overall costs associated with the plan
	GeneralCost []InsurancePlanPlanGeneralCost `json:"generalCost,omitempty"`

	// Business identifier for the plan
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Networks and tiers
	Network []common.Reference `json:"network,omitempty"`

	// Costs for specific coverage areas
	SpecificCost []InsurancePlanPlanSpecificCost `json:"specificCost,omitempty"`

	// Type of plan (Bronze; Silver; Gold; Platinum)
	Type *common.CodeableConcept `json:"type,omitempty"`
}
