package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// EnrollmentResponse represents enrollment and plan details from the processing of an EnrollmentRequest
type EnrollmentResponse struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "EnrollmentResponse"

	// The date when the enclosed suite of services were performed or completed
	Created *string `json:"created,omitempty"`

	// A description of the status of the adjudication
	Disposition *string `json:"disposition,omitempty"`

	// The Response business identifier
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The Insurer who produced this adjudicated response
	Organization *common.Reference `json:"organization,omitempty"`

	// Processing status: error, complete
	Outcome *EnrollmentResponseOutcome `json:"outcome,omitempty"`

	// Original request resource reference
	Request *common.Reference `json:"request,omitempty"`

	// The practitioner who is responsible for the services rendered to the patient
	RequestProvider *common.Reference `json:"requestProvider,omitempty"`

	// This element is labeled as a modifier because the status contains codes that mark the response as not currently valid
	Status *EnrollmentResponseStatus `json:"status,omitempty"`
}

// EnrollmentResponseOutcome represents the processing status
type EnrollmentResponseOutcome string

const (
	EnrollmentResponseOutcomeQueued   EnrollmentResponseOutcome = "queued"
	EnrollmentResponseOutcomeComplete EnrollmentResponseOutcome = "complete"
	EnrollmentResponseOutcomeError    EnrollmentResponseOutcome = "error"
	EnrollmentResponseOutcomePartial  EnrollmentResponseOutcome = "partial"
)

// EnrollmentResponseStatus represents the status of the enrollment response
type EnrollmentResponseStatus string

const (
	EnrollmentResponseStatusActive         EnrollmentResponseStatus = "active"
	EnrollmentResponseStatusCancelled      EnrollmentResponseStatus = "cancelled"
	EnrollmentResponseStatusDraft          EnrollmentResponseStatus = "draft"
	EnrollmentResponseStatusEnteredInError EnrollmentResponseStatus = "entered-in-error"
)
