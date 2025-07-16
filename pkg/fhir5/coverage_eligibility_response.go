// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// CoverageEligibilityResponseEvent represents an event for the response
type CoverageEligibilityResponseEvent struct {
	common.BackboneElement

	Type         common.CodeableConcept `json:"type"`
	WhenDateTime *string                `json:"whenDateTime,omitempty"`
	WhenPeriod   *common.Period         `json:"whenPeriod,omitempty"`
}

// CoverageEligibilityResponseError represents errors in the response
type CoverageEligibilityResponseError struct {
	common.BackboneElement

	Code       common.CodeableConcept `json:"code"`
	Expression []string               `json:"expression,omitempty"`
}

// CoverageEligibilityResponseInsuranceItemBenefit represents benefits for an item
type CoverageEligibilityResponseInsuranceItemBenefit struct {
	common.BackboneElement

	AllowedUnsignedInt *int                   `json:"allowedUnsignedInt,omitempty"`
	AllowedString      *string                `json:"allowedString,omitempty"`
	AllowedMoney       *Money                 `json:"allowedMoney,omitempty"`
	Type               common.CodeableConcept `json:"type"`
	UsedUnsignedInt    *int                   `json:"usedUnsignedInt,omitempty"`
	UsedString         *string                `json:"usedString,omitempty"`
	UsedMoney          *Money                 `json:"usedMoney,omitempty"`
}

// CoverageEligibilityResponseInsuranceItem represents an item in the response
type CoverageEligibilityResponseInsuranceItem struct {
	common.BackboneElement

	AuthorizationRequired   *bool                                             `json:"authorizationRequired,omitempty"`
	AuthorizationSupporting []common.CodeableConcept                          `json:"authorizationSupporting,omitempty"`
	AuthorizationUrl        *string                                           `json:"authorizationUrl,omitempty"`
	Benefit                 []CoverageEligibilityResponseInsuranceItemBenefit `json:"benefit,omitempty"`
	Category                *common.CodeableConcept                           `json:"category,omitempty"`
	Description             *string                                           `json:"description,omitempty"`
	Excluded                *bool                                             `json:"excluded,omitempty"`
	Modifier                []common.CodeableConcept                          `json:"modifier,omitempty"`
	Name                    *string                                           `json:"name,omitempty"`
	Network                 *common.CodeableConcept                           `json:"network,omitempty"`
	ProductOrService        *common.CodeableConcept                           `json:"productOrService,omitempty"`
	Provider                *common.Reference                                 `json:"provider,omitempty"`
	Term                    *common.CodeableConcept                           `json:"term,omitempty"`
	Unit                    *common.CodeableConcept                           `json:"unit,omitempty"`
}

// CoverageEligibilityResponseInsurance represents insurance information
type CoverageEligibilityResponseInsurance struct {
	common.BackboneElement

	BenefitPeriod *common.Period                             `json:"benefitPeriod,omitempty"`
	Coverage      common.Reference                           `json:"coverage"`
	Inforce       *bool                                      `json:"inforce,omitempty"`
	Item          []CoverageEligibilityResponseInsuranceItem `json:"item,omitempty"`
}

// CoverageEligibilityResponse represents a response for coverage eligibility
type CoverageEligibilityResponse struct {
	DomainResource

	ResourceType string `json:"resourceType"` // Always "CoverageEligibilityResponse"

	Created        string                                 `json:"created"`
	Disposition    *string                                `json:"disposition,omitempty"`
	Error          []CoverageEligibilityResponseError     `json:"error,omitempty"`
	Event          []CoverageEligibilityResponseEvent     `json:"event,omitempty"`
	Form           *common.CodeableConcept                `json:"form,omitempty"`
	Identifier     []common.Identifier                    `json:"identifier,omitempty"`
	Insurance      []CoverageEligibilityResponseInsurance `json:"insurance,omitempty"`
	Insurer        common.Reference                       `json:"insurer"`
	Outcome        CoverageEligibilityResponseOutcome     `json:"outcome"`
	Patient        common.Reference                       `json:"patient"`
	PreAuthRef     *string                                `json:"preAuthRef,omitempty"`
	Purpose        []CoverageEligibilityResponsePurpose   `json:"purpose"`
	Request        common.Reference                       `json:"request"`
	Requestor      *common.Reference                      `json:"requestor,omitempty"`
	ServicedDate   *string                                `json:"servicedDate,omitempty"`
	ServicedPeriod *common.Period                         `json:"servicedPeriod,omitempty"`
	Status         CoverageEligibilityResponseStatus      `json:"status"`
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
