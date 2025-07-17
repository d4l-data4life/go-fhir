package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// InsurancePlan represents details of a health insurance product/plan provided by an organization
type InsurancePlan struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "InsurancePlan"

	// An organization which administer other services such as underwriting, customer service and/or claims processing
	AdministeredBy *common.Reference `json:"administeredBy,omitempty"`

	// There are no dates associated with the alias/historic names
	Alias []string `json:"alias,omitempty"`

	// Where multiple contacts for the same purpose are provided there is a standard extension
	Contact []InsurancePlanContact `json:"contact,omitempty"`

	// Details about the coverage offered by the insurance product
	Coverage []InsurancePlanCoverage `json:"coverage,omitempty"`

	// The geographic region in which a health insurance product's benefits apply
	CoverageArea []common.Reference `json:"coverageArea,omitempty"`

	// The technical endpoints providing access to services operated for the health insurance product
	Endpoint []common.Reference `json:"endpoint,omitempty"`

	// Business identifiers assigned to this health insurance product
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// If the name of the product/plan changes, consider putting the old name in the alias column
	Name *string `json:"name,omitempty"`

	// Networks are represented as a hierarchy of organization resources
	Network []common.Reference `json:"network,omitempty"`

	// The entity that is providing the health insurance product and underwriting the risk
	OwnedBy *common.Reference `json:"ownedBy,omitempty"`

	// The period of time that the health insurance product is available
	Period *common.Period `json:"period,omitempty"`

	// Details about an insurance plan
	Plan []InsurancePlanPlan `json:"plan,omitempty"`

	// The current state of the health insurance product
	Status *InsurancePlanStatus `json:"status,omitempty"`

	// The kind of health insurance product
	Type []common.CodeableConcept `json:"type,omitempty"`
}

// InsurancePlanStatus represents the status of the insurance plan
type InsurancePlanStatus string

const (
	InsurancePlanStatusDraft   InsurancePlanStatus = "draft"
	InsurancePlanStatusActive  InsurancePlanStatus = "active"
	InsurancePlanStatusRetired InsurancePlanStatus = "retired"
	InsurancePlanStatusUnknown InsurancePlanStatus = "unknown"
)

// InsurancePlanContact represents where multiple contacts for the same purpose are provided
type InsurancePlanContact struct {
	common.BackboneElement

	// Visiting or postal addresses for the contact
	Address *common.Address `json:"address,omitempty"`

	// A name associated with the contact
	Name *common.HumanName `json:"name,omitempty"`

	// Indicates a purpose for which the contact can be reached
	Purpose *common.CodeableConcept `json:"purpose,omitempty"`

	// A contact detail (e.g. a telephone number or an email address) by which the party may be contacted
	Telecom []common.ContactPoint `json:"telecom,omitempty"`
}

// InsurancePlanCoverage represents details about the coverage offered by the insurance product
type InsurancePlanCoverage struct {
	common.BackboneElement

	// Specific benefits under this type of coverage
	Benefit []InsurancePlanCoverageBenefit `json:"benefit"`

	// Networks are represented as a hierarchy of organization resources
	Network []common.Reference `json:"network,omitempty"`

	// Type of coverage (Medical; Dental; Mental Health; Substance Abuse; Vision; Drug; Short Term; Long Term Care; Hospice; Home Health)
	Type common.CodeableConcept `json:"type"`
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

// InsurancePlanCoverageBenefitLimit represents the specific limits on the benefit
type InsurancePlanCoverageBenefitLimit struct {
	common.BackboneElement

	// Use CodeableConcept.text element if the data is free (uncoded) text
	Code *common.CodeableConcept `json:"code,omitempty"`

	// May also be called "eligible expense," "payment allowance," or "negotiated rate"
	Value *common.Quantity `json:"value,omitempty"`
}

// InsurancePlanPlan represents details about an insurance plan
type InsurancePlanPlan struct {
	common.BackboneElement

	// Overall costs associated with the plan
	GeneralCost []InsurancePlanPlanGeneralCost `json:"generalCost,omitempty"`

	// Business identifiers assigned to this health insurance plan
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The geographic region in which a health insurance plan's benefits apply
	CoverageArea []common.Reference `json:"coverageArea,omitempty"`

	// Networks are represented as a hierarchy of organization resources
	Network []common.Reference `json:"network,omitempty"`

	// Overall costs associated with the plan
	SpecificCost []InsurancePlanPlanSpecificCost `json:"specificCost,omitempty"`

	// Type of plan. For example, "Platinum" or "High Deductable"
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// InsurancePlanPlanGeneralCost represents overall costs associated with the plan
type InsurancePlanPlanGeneralCost struct {
	common.BackboneElement

	// Additional information about the general costs associated with this plan
	Comment *string `json:"comment,omitempty"`

	// Value of the cost
	Cost *common.Money `json:"cost,omitempty"`

	// Number of participants enrolled in the plan
	GroupSize *int `json:"groupSize,omitempty"`

	// Type of cost
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// InsurancePlanPlanSpecificCost represents list of the costs associated with a specific benefit
type InsurancePlanPlanSpecificCost struct {
	common.BackboneElement

	// List of the costs associated with a specific benefit
	Benefit []InsurancePlanPlanSpecificCostBenefit `json:"benefit,omitempty"`

	// General category of benefit (Medical; Dental; Vision; Drug; Mental Health; Substance Abuse)
	Category common.CodeableConcept `json:"category"`
}

// InsurancePlanPlanSpecificCostBenefit represents list of the costs associated with a specific benefit
type InsurancePlanPlanSpecificCostBenefit struct {
	common.BackboneElement

	// List of the costs associated with a specific benefit
	Cost []InsurancePlanPlanSpecificCostBenefitCost `json:"cost,omitempty"`

	// Type of specific benefit (preventive; primary care office visit; speciality office visit; hospitalization; emergency room; urgent care)
	Type common.CodeableConcept `json:"type"`
}

// InsurancePlanPlanSpecificCostBenefitCost represents list of the costs associated with a specific benefit
type InsurancePlanPlanSpecificCostBenefitCost struct {
	common.BackboneElement

	// Whether the cost applies to in-network or out-of-network providers (in-network; out-of-network; other)
	Applicability *common.CodeableConcept `json:"applicability,omitempty"`

	// Additional information about the cost, such as information about funding sources (e.g. HSA, HRA, FSA, RRA)
	Qualifiers []common.CodeableConcept `json:"qualifiers,omitempty"`

	// Type of cost (copay; individual; family)
	Type common.CodeableConcept `json:"type"`

	// The actual cost value
	Value *common.Quantity `json:"value,omitempty"`
}
