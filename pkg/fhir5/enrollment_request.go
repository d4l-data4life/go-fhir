// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// EnrollmentRequestStatus represents the status of an enrollment request
type EnrollmentRequestStatus string

const (
	EnrollmentRequestStatusActive         EnrollmentRequestStatus = "active"
	EnrollmentRequestStatusCancelled      EnrollmentRequestStatus = "cancelled"
	EnrollmentRequestStatusDraft          EnrollmentRequestStatus = "draft"
	EnrollmentRequestStatusEnteredInError EnrollmentRequestStatus = "entered-in-error"
)

// EnrollmentRequest provides insurance enrollment details to an insurer
type EnrollmentRequest struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "EnrollmentRequest"

	// Patient Resource
	Candidate *common.Reference `json:"candidate,omitempty"`

	// Reference to the program or plan identification, underwriter or payor
	Coverage *common.Reference `json:"coverage,omitempty"`

	// The date when this resource was created
	Created *string `json:"created,omitempty"`

	// The Response business identifier
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The Insurer who is target of the request
	Insurer *common.Reference `json:"insurer,omitempty"`

	// The practitioner who is responsible for the services rendered to the patient
	Provider *common.Reference `json:"provider,omitempty"`

	// This element is labeled as a modifier because the status contains codes that mark the request as not currently valid
	Status *EnrollmentRequestStatus `json:"status,omitempty"`
}
