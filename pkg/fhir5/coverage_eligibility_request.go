// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// CoverageEligibilityRequestEvent represents an event for the request
type CoverageEligibilityRequestEvent struct {
	common.BackboneElement

	Type         common.CodeableConcept `json:"type"`
	WhenDateTime *string                `json:"whenDateTime,omitempty"`
	WhenPeriod   *common.Period         `json:"whenPeriod,omitempty"`
}

// CoverageEligibilityRequestSupportingInfo represents supporting information
type CoverageEligibilityRequestSupportingInfo struct {
	common.BackboneElement

	AppliesToAll *bool            `json:"appliesToAll,omitempty"`
	Information  common.Reference `json:"information"`
	Sequence     int              `json:"sequence"`
}

// CoverageEligibilityRequestInsurance represents insurance information
type CoverageEligibilityRequestInsurance struct {
	common.BackboneElement

	BusinessArrangement *string          `json:"businessArrangement,omitempty"`
	Coverage            common.Reference `json:"coverage"`
	Focal               *bool            `json:"focal,omitempty"`
}

// CoverageEligibilityRequestItemDiagnosis represents diagnosis for an item
type CoverageEligibilityRequestItemDiagnosis struct {
	common.BackboneElement

	DiagnosisCodeableConcept *common.CodeableConcept `json:"diagnosisCodeableConcept,omitempty"`
	DiagnosisReference       *common.Reference       `json:"diagnosisReference,omitempty"`
}

// CoverageEligibilityRequestItem represents an item in the request
type CoverageEligibilityRequestItem struct {
	common.BackboneElement

	Category               *common.CodeableConcept                   `json:"category,omitempty"`
	Detail                 []common.Reference                        `json:"detail,omitempty"`
	Diagnosis              []CoverageEligibilityRequestItemDiagnosis `json:"diagnosis,omitempty"`
	Facility               *common.Reference                         `json:"facility,omitempty"`
	Modifier               []common.CodeableConcept                  `json:"modifier,omitempty"`
	ProductOrService       *common.CodeableConcept                   `json:"productOrService,omitempty"`
	Provider               *common.Reference                         `json:"provider,omitempty"`
	Quantity               *common.Quantity                          `json:"quantity,omitempty"`
	SupportingInfoSequence []int                                     `json:"supportingInfoSequence,omitempty"`
	UnitPrice              *Money                                    `json:"unitPrice,omitempty"`
}

// CoverageEligibilityRequest represents a request for coverage eligibility
type CoverageEligibilityRequest struct {
	DomainResource

	ResourceType string `json:"resourceType"` // Always "CoverageEligibilityRequest"

	Created        string                                     `json:"created"`
	Enterer        *common.Reference                          `json:"enterer,omitempty"`
	Event          []CoverageEligibilityRequestEvent          `json:"event,omitempty"`
	Facility       *common.Reference                          `json:"facility,omitempty"`
	Identifier     []common.Identifier                        `json:"identifier,omitempty"`
	Insurance      []CoverageEligibilityRequestInsurance      `json:"insurance,omitempty"`
	Insurer        common.Reference                           `json:"insurer"`
	Item           []CoverageEligibilityRequestItem           `json:"item,omitempty"`
	Patient        common.Reference                           `json:"patient"`
	Priority       *common.CodeableConcept                    `json:"priority,omitempty"`
	Provider       *common.Reference                          `json:"provider,omitempty"`
	Purpose        []CoverageEligibilityRequestPurpose        `json:"purpose"`
	ServicedDate   *string                                    `json:"servicedDate,omitempty"`
	ServicedPeriod *common.Period                             `json:"servicedPeriod,omitempty"`
	Status         CoverageEligibilityRequestStatus           `json:"status"`
	SupportingInfo []CoverageEligibilityRequestSupportingInfo `json:"supportingInfo,omitempty"`
}

// CoverageEligibilityRequestPurpose represents the purpose of the request
type CoverageEligibilityRequestPurpose string

const (
	CoverageEligibilityRequestPurposeAuthRequirements CoverageEligibilityRequestPurpose = "auth-requirements"
	CoverageEligibilityRequestPurposeBenefits         CoverageEligibilityRequestPurpose = "benefits"
	CoverageEligibilityRequestPurposeDiscovery        CoverageEligibilityRequestPurpose = "discovery"
	CoverageEligibilityRequestPurposeValidation       CoverageEligibilityRequestPurpose = "validation"
)

// CoverageEligibilityRequestStatus represents the status of the request
type CoverageEligibilityRequestStatus string

const (
	CoverageEligibilityRequestStatusActive         CoverageEligibilityRequestStatus = "active"
	CoverageEligibilityRequestStatusCancelled      CoverageEligibilityRequestStatus = "cancelled"
	CoverageEligibilityRequestStatusDraft          CoverageEligibilityRequestStatus = "draft"
	CoverageEligibilityRequestStatusEnteredInError CoverageEligibilityRequestStatus = "entered-in-error"
)
