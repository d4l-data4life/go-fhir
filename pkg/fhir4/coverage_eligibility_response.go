package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// CoverageEligibilityResponse represents the CoverageEligibilityResponse resource
type CoverageEligibilityResponse struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "CoverageEligibilityResponse"

	// The date this resource was created
	Created string `json:"created"`

	// A human readable description of the status of the adjudication
	Disposition *string `json:"disposition,omitempty"`

	// Errors encountered during the processing of the request
	Error []CoverageEligibilityResponseError `json:"error,omitempty"`

	// May be needed to identify specific jurisdictional forms
	Form *common.CodeableConcept `json:"form,omitempty"`

	// A unique identifier assigned to this coverage eligibility request
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// All insurance coverages for the patient which may be applicable for reimbursement
	Insurance []CoverageEligibilityResponseInsurance `json:"insurance,omitempty"`

	// The Insurer who issued the coverage in question and is the author of the response
	Insurer common.Reference `json:"insurer"`

	// The resource may be used to indicate that: the request has been held (queued) for processing
	Outcome CoverageEligibilityResponseOutcome `json:"outcome"`

	// The party who is the beneficiary of the supplied coverage and for whom eligibility is sought
	Patient common.Reference `json:"patient"`

	// A reference from the Insurer to which these services pertain to be used on further communication
	PreAuthRef *string `json:"preAuthRef,omitempty"`

	// Code to specify whether requesting: prior authorization requirements for some service categories or billing codes
	Purpose []CoverageEligibilityResponsePurpose `json:"purpose"`

	// Reference to the original request resource
	Request common.Reference `json:"request"`

	// Typically this field would be 1..1 where this party is responsible for the claim
	Requestor *common.Reference `json:"requestor,omitempty"`

	// The date or dates when the enclosed suite of services were performed or completed
	ServicedDate *string `json:"servicedDate,omitempty"`

	// The date or dates when the enclosed suite of services were performed or completed
	ServicedPeriod *common.Period `json:"servicedPeriod,omitempty"`

	// This element is labeled as a modifier because the status contains codes that mark the resource as not currently valid
	Status CoverageEligibilityResponseStatus `json:"status"`
}

// CoverageEligibilityResponseOutcome represents the outcome of the response
type CoverageEligibilityResponseOutcome string

const (
	CoverageEligibilityResponseOutcomeQueued   CoverageEligibilityResponseOutcome = "queued"
	CoverageEligibilityResponseOutcomeComplete CoverageEligibilityResponseOutcome = "complete"
	CoverageEligibilityResponseOutcomeError    CoverageEligibilityResponseOutcome = "error"
	CoverageEligibilityResponseOutcomePartial  CoverageEligibilityResponseOutcome = "partial"
)

// CoverageEligibilityResponsePurpose represents the purpose of the response
type CoverageEligibilityResponsePurpose string

const (
	CoverageEligibilityResponsePurposeAuthRequirements CoverageEligibilityResponsePurpose = "auth-requirements"
	CoverageEligibilityResponsePurposeBenefits         CoverageEligibilityResponsePurpose = "benefits"
	CoverageEligibilityResponsePurposeDiscovery        CoverageEligibilityResponsePurpose = "discovery"
	CoverageEligibilityResponsePurposeValidation       CoverageEligibilityResponsePurpose = "validation"
)

// CoverageEligibilityResponseStatus represents the status of the response
type CoverageEligibilityResponseStatus string

const (
	CoverageEligibilityResponseStatusActive         CoverageEligibilityResponseStatus = "active"
	CoverageEligibilityResponseStatusCancelled      CoverageEligibilityResponseStatus = "cancelled"
	CoverageEligibilityResponseStatusDraft          CoverageEligibilityResponseStatus = "draft"
	CoverageEligibilityResponseStatusEnteredInError CoverageEligibilityResponseStatus = "entered-in-error"
)

// CoverageEligibilityResponseInsuranceItemBenefit represents benefits used to date
type CoverageEligibilityResponseInsuranceItemBenefit struct {
	common.BackboneElement

	// The quantity of the benefit which is permitted under the coverage
	AllowedUnsignedInt *int `json:"allowedUnsignedInt,omitempty"`

	// The quantity of the benefit which is permitted under the coverage
	AllowedString *string `json:"allowedString,omitempty"`

	// The quantity of the benefit which is permitted under the coverage
	AllowedMoney *common.Money `json:"allowedMoney,omitempty"`

	// For example: deductible, visits, benefit amount
	Type common.CodeableConcept `json:"type"`

	// The quantity of the benefit which have been consumed to date
	UsedUnsignedInt *int `json:"usedUnsignedInt,omitempty"`

	// The quantity of the benefit which have been consumed to date
	UsedString *string `json:"usedString,omitempty"`

	// The quantity of the benefit which have been consumed to date
	UsedMoney *common.Money `json:"usedMoney,omitempty"`
}

// CoverageEligibilityResponseInsuranceItem represents benefits and optionally current balances
type CoverageEligibilityResponseInsuranceItem struct {
	common.BackboneElement

	// A boolean flag indicating whether a preauthorization is required prior to actual service delivery
	AuthorizationRequired *bool `json:"authorizationRequired,omitempty"`

	// Codes or comments regarding information or actions associated with the preauthorization
	AuthorizationSupporting []common.CodeableConcept `json:"authorizationSupporting,omitempty"`

	// A web location for obtaining requirements or descriptive information regarding the preauthorization
	AuthorizationUrl *string `json:"authorizationUrl,omitempty"`

	// Benefits used to date
	Benefit []CoverageEligibilityResponseInsuranceItemBenefit `json:"benefit,omitempty"`

	// Examples include Medical Care, Periodontics, Renal Dialysis, Vision Coverage
	Category *common.CodeableConcept `json:"category,omitempty"`

	// For example 'DENT2 covers 100% of basic, 50% of major but excludes Ortho, Implants and Cosmetic services'
	Description *string `json:"description,omitempty"`

	// True if the indicated class of service is excluded from the plan
	Excluded *bool `json:"excluded,omitempty"`

	// For example in Oral whether the treatment is cosmetic or associated with TMJ
	Modifier []common.CodeableConcept `json:"modifier,omitempty"`

	// For example: MED01, or DENT2
	Name *string `json:"name,omitempty"`

	// Is a flag to indicate whether the benefits refer to in-network providers or out-of-network providers
	Network *common.CodeableConcept `json:"network,omitempty"`

	// Code to indicate the Professional Service or Product supplied
	ProductOrService *common.CodeableConcept `json:"productOrService,omitempty"`

	// The practitioner who is eligible for the provision of the product or service
	Provider *common.Reference `json:"provider,omitempty"`

	// The term or period of the values such as 'maximum lifetime benefit' or 'maximum annual visits'
	Term *common.CodeableConcept `json:"term,omitempty"`

	// Indicates if the benefits apply to an individual or to the family
	Unit *common.CodeableConcept `json:"unit,omitempty"`
}

// CoverageEligibilityResponseInsurance represents insurance coverages for the patient
type CoverageEligibilityResponseInsurance struct {
	common.BackboneElement

	// The term of the benefits documented in this response
	BenefitPeriod *common.Period `json:"benefitPeriod,omitempty"`

	// Reference to the insurance card level information contained in the Coverage resource
	Coverage common.Reference `json:"coverage"`

	// Flag indicating if the coverage provided is inforce currently
	Inforce *bool `json:"inforce,omitempty"`

	// Benefits and optionally current balances, and authorization details by category or service
	Item []CoverageEligibilityResponseInsuranceItem `json:"item,omitempty"`
}

// CoverageEligibilityResponseError represents errors encountered during the processing of the request
type CoverageEligibilityResponseError struct {
	common.BackboneElement

	// An error code, from a specified code system, which details why the eligibility check could not be performed
	Code common.CodeableConcept `json:"code"`
}
