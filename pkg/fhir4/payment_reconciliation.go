package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// PaymentReconciliation represents the details including amount of a payment and allocates the payment items being paid
type PaymentReconciliation struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "PaymentReconciliation"

	// The date when the resource was created
	Created string `json:"created"`

	// Distribution of the payment amount for a previously acknowledged payable
	Detail []PaymentReconciliationDetail `json:"detail,omitempty"`

	// A human readable description of the status of the request for the reconciliation
	Disposition *string `json:"disposition,omitempty"`

	// May be needed to identify specific jurisdictional forms
	FormCode *common.CodeableConcept `json:"formCode,omitempty"`

	// A unique identifier assigned to this payment reconciliation
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The resource may be used to indicate that: the request has been held (queued) for processing
	Outcome *PaymentReconciliationOutcome `json:"outcome,omitempty"`

	// Total payment amount as indicated on the financial instrument
	PaymentAmount common.Money `json:"paymentAmount"`

	// The date of payment as indicated on the financial instrument
	PaymentDate string `json:"paymentDate"`

	// For example: EFT number or check number
	PaymentIdentifier *common.Identifier `json:"paymentIdentifier,omitempty"`

	// This party is also responsible for the reconciliation
	PaymentIssuer *common.Reference `json:"paymentIssuer,omitempty"`

	// The period of time for which payments have been gathered into this bulk payment for settlement
	Period *common.Period `json:"period,omitempty"`

	// A note that describes or explains the processing in a human readable form
	ProcessNote []PaymentReconciliationProcessNote `json:"processNote,omitempty"`

	// Original request resource reference
	Request *common.Reference `json:"request,omitempty"`

	// The practitioner who is responsible for the services rendered to the patient
	Requestor *common.Reference `json:"requestor,omitempty"`

	// This element is labeled as a modifier because the status contains codes that mark the resource as not currently valid
	Status PaymentReconciliationStatus `json:"status"`
}

// PaymentReconciliationStatus represents this element is labeled as a modifier because the status contains codes that mark the resource as not currently valid
type PaymentReconciliationStatus string

const (
	PaymentReconciliationStatusActive         PaymentReconciliationStatus = "active"
	PaymentReconciliationStatusCancelled      PaymentReconciliationStatus = "cancelled"
	PaymentReconciliationStatusDraft          PaymentReconciliationStatus = "draft"
	PaymentReconciliationStatusEnteredInError PaymentReconciliationStatus = "entered-in-error"
)

// PaymentReconciliationOutcome represents the resource may be used to indicate that: the request has been held (queued) for processing
type PaymentReconciliationOutcome string

const (
	PaymentReconciliationOutcomeQueued   PaymentReconciliationOutcome = "queued"
	PaymentReconciliationOutcomeComplete PaymentReconciliationOutcome = "complete"
	PaymentReconciliationOutcomeError    PaymentReconciliationOutcome = "error"
	PaymentReconciliationOutcomePartial  PaymentReconciliationOutcome = "partial"
)

// PaymentReconciliationDetail represents distribution of the payment amount for a previously acknowledged payable
type PaymentReconciliationDetail struct {
	common.BackboneElement

	// The monetary amount allocated from the total payment to the payable
	Amount *common.Money `json:"amount,omitempty"`

	// The date from the response resource containing a commitment to pay
	Date *string `json:"date,omitempty"`

	// Unique identifier for the current payment item for the referenced payable
	Identifier *common.Identifier `json:"identifier,omitempty"`

	// The party which is receiving the payment
	Payee *common.Reference `json:"payee,omitempty"`

	// Unique identifier for the prior payment item for the referenced payable
	Predecessor *common.Identifier `json:"predecessor,omitempty"`

	// A resource, such as a Claim, the evaluation of which could lead to payment
	Request *common.Reference `json:"request,omitempty"`

	// A resource, such as a ClaimResponse, which contains a commitment to payment
	Response *common.Reference `json:"response,omitempty"`

	// A reference to the individual who is responsible for inquiries regarding the response and its payment
	Responsible *common.Reference `json:"responsible,omitempty"`

	// The party which submitted the claim or financial transaction
	Submitter *common.Reference `json:"submitter,omitempty"`

	// For example: payment, adjustment, funds advance, etc
	Type common.CodeableConcept `json:"type"`
}

// PaymentReconciliationProcessNote represents a note that describes or explains the processing in a human readable form
type PaymentReconciliationProcessNote struct {
	common.BackboneElement

	// The explanation or description associated with the processing
	Text *string `json:"text,omitempty"`

	// The business purpose of the note text
	Type *PaymentReconciliationProcessNoteType `json:"type,omitempty"`
}

// PaymentReconciliationProcessNoteType represents the business purpose of the note text
type PaymentReconciliationProcessNoteType string

const (
	PaymentReconciliationProcessNoteTypeDisplay   PaymentReconciliationProcessNoteType = "display"
	PaymentReconciliationProcessNoteTypePrint     PaymentReconciliationProcessNoteType = "print"
	PaymentReconciliationProcessNoteTypePrintoper PaymentReconciliationProcessNoteType = "printoper"
)
