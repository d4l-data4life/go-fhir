package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// PaymentReconciliationOutcome represents the outcome of a request for a reconciliation
type PaymentReconciliationOutcome string

const (
	PaymentReconciliationOutcomeQueued    PaymentReconciliationOutcome = "queued"
	PaymentReconciliationOutcomeComplete  PaymentReconciliationOutcome = "complete"
	PaymentReconciliationOutcomeError     PaymentReconciliationOutcome = "error"
	PaymentReconciliationOutcomePartial   PaymentReconciliationOutcome = "partial"
	PaymentReconciliationOutcomeCancelled PaymentReconciliationOutcome = "cancelled"
)

// PaymentReconciliationStatus represents the status of a payment reconciliation
type PaymentReconciliationStatus string

const (
	PaymentReconciliationStatusActive         PaymentReconciliationStatus = "active"
	PaymentReconciliationStatusCancelled      PaymentReconciliationStatus = "cancelled"
	PaymentReconciliationStatusDraft          PaymentReconciliationStatus = "draft"
	PaymentReconciliationStatusEnteredInError PaymentReconciliationStatus = "entered-in-error"
)

// PaymentReconciliationProcessNoteType represents the business purpose of the note text
type PaymentReconciliationProcessNoteType string

const (
	PaymentReconciliationProcessNoteTypeDisplay PaymentReconciliationProcessNoteType = "display"
	PaymentReconciliationProcessNoteTypePrint   PaymentReconciliationProcessNoteType = "print"
	PaymentReconciliationProcessNoteTypePrintop PaymentReconciliationProcessNoteType = "printop"
)

// PaymentReconciliationProcessNote represents the practitioner who is responsible for the services rendered to the patient
type PaymentReconciliationProcessNote struct {
	common.BackboneElement

	// The explanation or description associated with the processing
	Text *string `json:"text,omitempty"`

	// The business purpose of the note text
	Type *PaymentReconciliationProcessNoteType `json:"type,omitempty"`
}

// PaymentReconciliation represents a payment reconciliation
type PaymentReconciliation struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "PaymentReconciliation"

	// The date when the resource was created
	Created *string `json:"created,omitempty"`

	// A human readable description of the status of the request for the reconciliation
	Disposition *string `json:"disposition,omitempty"`

	// The form code or additional code identifying the form, such as a document definition
	FormCode *common.CodeableConcept `json:"formCode,omitempty"`

	// A unique identifier assigned to this payment reconciliation
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The Insurer who produced this adjudicated response
	Issuer *common.Reference `json:"issuer,omitempty"`

	// The outcome of a request for a reconciliation
	Outcome *PaymentReconciliationOutcome `json:"outcome,omitempty"`

	// The period of time for which payments have been gathered into this bulk payment for settlement
	Period *common.Period `json:"period,omitempty"`

	// The practitioner who is responsible for the services rendered to the patient
	ProcessNote []PaymentReconciliationProcessNote `json:"processNote,omitempty"`

	// Original request resource reference
	Request *common.Reference `json:"request,omitempty"`

	// The organization which is responsible for the services rendered to the patient
	Requestor *common.Reference `json:"requestor,omitempty"`

	// The status of the resource instance
	Status PaymentReconciliationStatus `json:"status"`

	// Total payment amount as indicated on the financial instrument
	TotalAmount Money `json:"totalAmount"`

	// The date when the above payment action occurred
	TransactionDate *string `json:"transactionDate,omitempty"`
}
