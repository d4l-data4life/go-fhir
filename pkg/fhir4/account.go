package fhir4

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

// AccountCoverage represents coverage information for an account
type AccountCoverage struct {
	common.BackboneElement

	// The party(s) that contribute to payment (or part of) of the charges applied to this account
	Coverage common.Reference `json:"coverage"`

	// The priority of the coverage in the context of this account
	Priority *int `json:"priority,omitempty"`
}

// AccountGuarantor represents guarantor information for an account
type AccountGuarantor struct {
	common.BackboneElement

	// The entity who is responsible
	Party common.Reference `json:"party"`

	// A guarantor may be placed on credit hold or otherwise have their role temporarily suspended
	OnHold *bool `json:"onHold,omitempty"`

	// The timeframe during which the guarantor accepts responsibility for the account
	Period *common.Period `json:"period,omitempty"`
}

// Account represents a financial tool for tracking value accrued for a particular purpose
type Account struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Account"

	// Typically, this may be some form of insurance, internal charges, or self-pay
	Coverage []AccountCoverage `json:"coverage,omitempty"`

	// Provides additional information about what the account tracks and how it is used
	Description *string `json:"description,omitempty"`

	// The parties responsible for balancing the account if other payment options fall short
	Guarantor []AccountGuarantor `json:"guarantor,omitempty"`

	// Unique identifier used to reference the account
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Name used for the account when displaying it to humans in reports, etc
	Name *string `json:"name,omitempty"`

	// Indicates the service area, hospital, department, etc. with responsibility for managing the Account
	Owner *common.Reference `json:"owner,omitempty"`

	// Reference to a parent Account
	PartOf *common.Reference `json:"partOf,omitempty"`

	// It is possible for transactions to be posted outside the service period, as long as the service was provided within the defined service period
	ServicePeriod *common.Period `json:"servicePeriod,omitempty"`

	// This element is labeled as a modifier because the status contains the codes inactive and entered-in-error that mark the Account as not currently valid
	Status AccountStatus `json:"status"`

	// Accounts can be applied to non-patients for tracking other non-patient related activities, such as group services (patients not tracked, and costs charged to another body), or might not be allocated
	Subject []common.Reference `json:"subject,omitempty"`

	// Categorizes the account for reporting and searching purposes
	Type *common.CodeableConcept `json:"type,omitempty"`
}
