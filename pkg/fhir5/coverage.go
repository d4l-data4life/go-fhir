// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// CoveragePaymentBy represents a paying party for coverage
type CoveragePaymentBy struct {
	common.BackboneElement

	Party          common.Reference `json:"party"`
	Responsibility *string          `json:"responsibility,omitempty"`
}

// CoverageClass represents a class of coverage
type CoverageClass struct {
	common.BackboneElement

	Name  *string                `json:"name,omitempty"`
	Type  common.CodeableConcept `json:"type"`
	Value common.Identifier      `json:"value"`
}

// CoverageCostToBeneficiaryException represents exceptions to beneficiary costs
type CoverageCostToBeneficiaryException struct {
	common.BackboneElement

	Period *common.Period         `json:"period,omitempty"`
	Type   common.CodeableConcept `json:"type"`
}

// CoverageCostToBeneficiary represents costs to the beneficiary
type CoverageCostToBeneficiary struct {
	common.BackboneElement

	Category      *common.CodeableConcept              `json:"category,omitempty"`
	Exception     []CoverageCostToBeneficiaryException `json:"exception,omitempty"`
	Network       *common.CodeableConcept              `json:"network,omitempty"`
	Term          *common.CodeableConcept              `json:"term,omitempty"`
	Type          *common.CodeableConcept              `json:"type,omitempty"`
	Unit          *common.CodeableConcept              `json:"unit,omitempty"`
	ValueQuantity *common.Quantity                     `json:"valueQuantity,omitempty"`
	ValueMoney    *Money                               `json:"valueMoney,omitempty"`
}

// Coverage represents insurance coverage
type Coverage struct {
	DomainResource

	ResourceType string `json:"resourceType"` // Always "Coverage"

	Beneficiary       common.Reference            `json:"beneficiary"`
	Class             []CoverageClass             `json:"class,omitempty"`
	Contract          []common.Reference          `json:"contract,omitempty"`
	CostToBeneficiary []CoverageCostToBeneficiary `json:"costToBeneficiary,omitempty"`
	Dependent         *string                     `json:"dependent,omitempty"`
	Identifier        []common.Identifier         `json:"identifier,omitempty"`
	InsurancePlan     *common.Reference           `json:"insurancePlan,omitempty"`
	Insurer           *common.Reference           `json:"insurer,omitempty"`
	Kind              CoverageKind                `json:"kind"`
	Network           *string                     `json:"network,omitempty"`
	Order             *int                        `json:"order,omitempty"`
	PaymentBy         []CoveragePaymentBy         `json:"paymentBy,omitempty"`
	Period            *common.Period              `json:"period,omitempty"`
	PolicyHolder      *common.Reference           `json:"policyHolder,omitempty"`
	Relationship      *common.CodeableConcept     `json:"relationship,omitempty"`
	Status            CoverageStatus              `json:"status"`
	Subrogation       *bool                       `json:"subrogation,omitempty"`
	Subscriber        *common.Reference           `json:"subscriber,omitempty"`
	SubscriberId      []common.Identifier         `json:"subscriberId,omitempty"`
	Type              *common.CodeableConcept     `json:"type,omitempty"`
}

// CoverageKind represents the kind of coverage
type CoverageKind string

const (
	CoverageKindInsurance CoverageKind = "insurance"
	CoverageKindSelfPay   CoverageKind = "self-pay"
	CoverageKindOther     CoverageKind = "other"
)

// CoverageStatus represents the status of coverage
type CoverageStatus string

const (
	CoverageStatusActive         CoverageStatus = "active"
	CoverageStatusCancelled      CoverageStatus = "cancelled"
	CoverageStatusDraft          CoverageStatus = "draft"
	CoverageStatusEnteredInError CoverageStatus = "entered-in-error"
)
