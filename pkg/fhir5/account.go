// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// AccountStatus represents the status of an account
type AccountStatus string

const (
	AccountStatusActive         AccountStatus = "active"
	AccountStatusInactive       AccountStatus = "inactive"
	AccountStatusEnteredInError AccountStatus = "entered-in-error"
	AccountStatusOnHold         AccountStatus = "on-hold"
	AccountStatusUnknown        AccountStatus = "unknown"
)

// AccountCoverage represents coverage for an account
type AccountCoverage struct {
	common.BackboneElement

	// The party(s) that contribute to payment (or part of) of the charges applied to this account
	Coverage common.Reference `json:"coverage"`

	// The priority of the coverage in the context of this account
	Priority *int `json:"priority,omitempty"`
}

// AccountGuarantor represents parties responsible for balancing the account
type AccountGuarantor struct {
	common.BackboneElement

	// A guarantor may be placed on credit hold or otherwise have their role temporarily suspended
	OnHold *bool `json:"onHold,omitempty"`

	// The entity who is responsible
	Party common.Reference `json:"party"`

	// The timeframe during which the guarantor accepts responsibility for the account
	Period *common.Period `json:"period,omitempty"`
}

// AccountDiagnosis represents diagnoses relevant to the account
type AccountDiagnosis struct {
	common.BackboneElement

	// The diagnosis relevant to the account
	Condition CodeableReference `json:"condition"`

	// Date of the diagnosis
	DateOfDiagnosis *string `json:"dateOfDiagnosis,omitempty"`

	// Was the diagnosis present on admission
	OnAdmission *bool `json:"onAdmission,omitempty"`

	// Package code
	PackageCode []common.CodeableConcept `json:"packageCode,omitempty"`

	// Ranking of the diagnosis
	Sequence *int `json:"sequence,omitempty"`

	// Type that this diagnosis has relevant to the account
	Type []common.CodeableConcept `json:"type,omitempty"`
}

// AccountProcedure represents procedures relevant to the account
type AccountProcedure struct {
	common.BackboneElement

	// The procedure relevant to the account
	Code CodeableReference `json:"code"`

	// Date of the procedure
	DateOfService *string `json:"dateOfService,omitempty"`

	// Devices associated with the procedure
	Device []common.Reference `json:"device,omitempty"`

	// Package code
	PackageCode []common.CodeableConcept `json:"packageCode,omitempty"`

	// Ranking of the procedure
	Sequence *int `json:"sequence,omitempty"`

	// How this procedure value should be used in charging the account
	Type []common.CodeableConcept `json:"type,omitempty"`
}

// AccountRelatedAccount represents related accounts
type AccountRelatedAccount struct {
	common.BackboneElement

	// Reference to an associated Account
	Account common.Reference `json:"account"`

	// Relationship of the associated Account
	Relationship *common.CodeableConcept `json:"relationship,omitempty"`
}

// Money represents an amount of money
type Money struct {
	common.Element

	// Numerical value (with implicit precision)
	Value *float64 `json:"value,omitempty"`

	// ISO 4217 Currency Code
	Currency *string `json:"currency,omitempty"`
}

// AccountBalance represents balance information for an account
type AccountBalance struct {
	common.BackboneElement

	// Who is expected to pay this part of the balance
	Aggregate *common.CodeableConcept `json:"aggregate,omitempty"`

	// The actual balance value calculated
	Amount Money `json:"amount"`

	// The amount is only an estimated value
	Estimate *bool `json:"estimate,omitempty"`

	// The term of the account balances
	Term *common.CodeableConcept `json:"term,omitempty"`
}

// Account represents a financial tool for tracking value accrued for a particular purpose
type Account struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Account"

	// The calculated account balances
	Balance []AccountBalance `json:"balance,omitempty"`

	// The BillingStatus tracks the lifecycle of the account
	BillingStatus *common.CodeableConcept `json:"billingStatus,omitempty"`

	// Time the balances were calculated
	CalculatedAt *string `json:"calculatedAt,omitempty"`

	// The coverage accounts to be used for billing
	Coverage []AccountCoverage `json:"coverage,omitempty"`

	// The default currency for the account
	Currency *common.CodeableConcept `json:"currency,omitempty"`

	// Explanation of purpose/use
	Description *string `json:"description,omitempty"`

	// The set of diagnoses that are relevant for billing
	Diagnosis []AccountDiagnosis `json:"diagnosis,omitempty"`

	// The parties responsible for balancing the account
	Guarantor []AccountGuarantor `json:"guarantor,omitempty"`

	// Unique identifier used to reference the account
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Human-readable label
	Name *string `json:"name,omitempty"`

	// Entity managing the Account
	Owner *common.Reference `json:"owner,omitempty"`

	// The set of procedures that are relevant for billing
	Procedure []AccountProcedure `json:"procedure,omitempty"`

	// Other associated accounts related to this account
	RelatedAccount []AccountRelatedAccount `json:"relatedAccount,omitempty"`

	// Transaction window
	ServicePeriod *common.Period `json:"servicePeriod,omitempty"`

	// active | inactive | entered-in-error | on-hold | unknown
	Status AccountStatus `json:"status"`

	// The entity that caused the expenses
	Subject []common.Reference `json:"subject,omitempty"`

	// E.g. patient, expense, depreciation
	Type *common.CodeableConcept `json:"type,omitempty"`
}
