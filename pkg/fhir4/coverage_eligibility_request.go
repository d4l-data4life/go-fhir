package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// CoverageEligibilityRequest represents the CoverageEligibilityRequest resource
type CoverageEligibilityRequest struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "CoverageEligibilityRequest"

	// The date when this resource was created
	Created string `json:"created"`

	// Person who created the request
	Enterer *common.Reference `json:"enterer,omitempty"`

	// Facility where the services are intended to be provided
	Facility *common.Reference `json:"facility,omitempty"`

	// A unique identifier assigned to this coverage eligibility request
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// All insurance coverages for the patient which may be applicable for reimbursement
	Insurance []CoverageEligibilityRequestInsurance `json:"insurance,omitempty"`

	// The Insurer who issued the coverage in question and is the recipient of the request
	Insurer common.Reference `json:"insurer"`

	// Service categories or billable services for which benefit details and/or an authorization prior to service delivery may be required by the payor
	Item []CoverageEligibilityRequestItem `json:"item,omitempty"`

	// The patient
	Patient common.Reference `json:"patient"`

	// When the requestor expects the processor to complete processing
	Priority *common.CodeableConcept `json:"priority,omitempty"`

	// Typically this field would be 1..1 where this party is responsible for the eligibility request
	Provider *common.Reference `json:"provider,omitempty"`

	// Code to specify whether requesting: prior authorization requirements for some service categories or billing codes
	Purpose []CoverageEligibilityRequestPurpose `json:"purpose"`

	// The date or dates when the enclosed suite of services were performed or completed
	ServicedDate *string `json:"servicedDate,omitempty"`

	// The date or dates when the enclosed suite of services were performed or completed
	ServicedPeriod *common.Period `json:"servicedPeriod,omitempty"`

	// This element is labeled as a modifier because the status contains codes that mark the resource as not currently valid
	Status CoverageEligibilityRequestStatus `json:"status"`

	// Often there are multiple jurisdiction specific valuesets which are required
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

// CoverageEligibilityRequestSupportingInfo represents supporting information
type CoverageEligibilityRequestSupportingInfo struct {
	common.BackboneElement

	// The supporting materials are applicable for all detail items, product/service categories and specific billing codes
	AppliesToAll *bool `json:"appliesToAll,omitempty"`

	// Could be used to provide references to other resources, document
	Information common.Reference `json:"information"`

	// A number to uniquely identify supporting information entries
	Sequence int `json:"sequence"`
}

// CoverageEligibilityRequestInsurance represents insurance coverages for the patient
type CoverageEligibilityRequestInsurance struct {
	common.BackboneElement

	// A business agreement number established between the provider and the insurer for special business processing purposes
	BusinessArrangement *string `json:"businessArrangement,omitempty"`

	// Reference to the insurance card level information contained in the Coverage resource
	Coverage common.Reference `json:"coverage"`

	// A patient may (will) have multiple insurance policies which provide reimbursement for healthcare services and products
	Focal *bool `json:"focal,omitempty"`
}

// CoverageEligibilityRequestItemDiagnosis represents patient diagnosis for which care is sought
type CoverageEligibilityRequestItemDiagnosis struct {
	common.BackboneElement

	// The nature of illness or problem in a coded form or as a reference to an external defined Condition
	DiagnosisCodeableConcept *common.CodeableConcept `json:"diagnosisCodeableConcept,omitempty"`

	// The nature of illness or problem in a coded form or as a reference to an external defined Condition
	DiagnosisReference *common.Reference `json:"diagnosisReference,omitempty"`
}

// CoverageEligibilityRequestItem represents service categories or billable services
type CoverageEligibilityRequestItem struct {
	common.BackboneElement

	// Examples include Medical Care, Periodontics, Renal Dialysis, Vision Coverage
	Category *common.CodeableConcept `json:"category,omitempty"`

	// The plan/proposal/order describing the proposed service in detail
	Detail []common.Reference `json:"detail,omitempty"`

	// Patient diagnosis for which care is sought
	Diagnosis []CoverageEligibilityRequestItemDiagnosis `json:"diagnosis,omitempty"`

	// Facility where the services will be provided
	Facility *common.Reference `json:"facility,omitempty"`

	// For example in Oral whether the treatment is cosmetic or associated with TMJ
	Modifier []common.CodeableConcept `json:"modifier,omitempty"`

	// Code to indicate the Professional Service or Product supplied
	ProductOrService *common.CodeableConcept `json:"productOrService,omitempty"`

	// The practitioner who is responsible for the product or service to be rendered to the patient
	Provider *common.Reference `json:"provider,omitempty"`

	// The number of repetitions of a service or product
	Quantity *common.Quantity `json:"quantity,omitempty"`

	// Exceptions, special conditions and supporting information applicable for this service or product line
	SupportingInfoSequence []int `json:"supportingInfoSequence,omitempty"`

	// The amount charged to the patient by the provider for a single unit
	UnitPrice *common.Money `json:"unitPrice,omitempty"`
}
