package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// Coverage represents the Coverage resource
type Coverage struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Coverage"

	// The party who benefits from the insurance coverage; the patient when products and/or services are provided
	Beneficiary common.Reference `json:"beneficiary"`

	// For example may be used to identify a class of coverage or employer group, Policy, Plan
	Class []CoverageClass `json:"class,omitempty"`

	// The policy(s) which constitute this insurance coverage
	Contract []common.Reference `json:"contract,omitempty"`

	// For example by knowing the patient visit co-pay, the provider can collect the amount prior to undertaking treatment
	CostToBeneficiary []CoverageCostToBeneficiary `json:"costToBeneficiary,omitempty"`

	// Periodically the member number is constructed from the subscriberId and the dependant number
	Dependent *string `json:"dependent,omitempty"`

	// The main (and possibly only) identifier for the coverage - often referred to as a Member Id, Certificate number, Personal Health Number or Case ID
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The insurer-specific identifier for the insurer-defined network of providers to which the beneficiary may seek treatment
	Network *string `json:"network,omitempty"`

	// The order of applicability of this coverage relative to other coverages which are currently in force
	Order *int `json:"order,omitempty"`

	// May provide multiple identifiers such as insurance company identifier or business identifier (BIN number)
	Payor []common.Reference `json:"payor"`

	// Time period during which the coverage is in force
	Period *common.Period `json:"period,omitempty"`

	// For example: may be an individual, corporation or the subscriber's employer
	PolicyHolder *common.Reference `json:"policyHolder,omitempty"`

	// Typically, an individual uses policies which are theirs (relationship='self') before policies owned by others
	Relationship *common.CodeableConcept `json:"relationship,omitempty"`

	// This element is labeled as a modifier because the status contains the code entered-in-error that marks the coverage as not currently valid
	Status CoverageStatus `json:"status"`

	// Typically, automotive and worker's compensation policies would be flagged with 'subrogation=true' to enable healthcare payors to collect against accident claims
	Subrogation *bool `json:"subrogation,omitempty"`

	// May be self or a parent in the case of dependants
	Subscriber *common.Reference `json:"subscriber,omitempty"`

	// The insurer assigned ID for the Subscriber
	SubscriberId *string `json:"subscriberId,omitempty"`

	// The type of coverage: social program, medical plan, accident coverage (workers compensation, auto), group health or payment by an individual or organization
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// CoverageStatus represents the status of the coverage
type CoverageStatus string

const (
	CoverageStatusActive         CoverageStatus = "active"
	CoverageStatusCancelled      CoverageStatus = "cancelled"
	CoverageStatusDraft          CoverageStatus = "draft"
	CoverageStatusEnteredInError CoverageStatus = "entered-in-error"
)

// CoverageClass represents a class of coverage
type CoverageClass struct {
	common.BackboneElement

	// A short description for the class
	Name *string `json:"name,omitempty"`

	// The type of classification for which an insurer-specific class label or number and optional name is provided
	Type common.CodeableConcept `json:"type"`

	// For example, the Group or Plan number
	Value string `json:"value"`
}

// CoverageCostToBeneficiaryException represents exceptions or reductions to patient costs
type CoverageCostToBeneficiaryException struct {
	common.BackboneElement

	// The timeframe during when the exception is in force
	Period *common.Period `json:"period,omitempty"`

	// The code for the specific exception
	Type common.CodeableConcept `json:"type"`
}

// CoverageCostToBeneficiary represents cost to beneficiary
type CoverageCostToBeneficiary struct {
	common.BackboneElement

	// A suite of codes indicating exceptions or reductions to patient costs and their effective periods
	Exception []CoverageCostToBeneficiaryException `json:"exception,omitempty"`

	// For example visit, specialist visits, emergency, inpatient care, etc
	Type *common.CodeableConcept `json:"type,omitempty"`

	// Amount may be expressed as a percentage of the service/product cost or a fixed amount of currency
	ValueQuantity *common.Quantity `json:"valueQuantity,omitempty"`

	// Amount may be expressed as a percentage of the service/product cost or a fixed amount of currency
	ValueMoney *common.Money `json:"valueMoney,omitempty"`
}
