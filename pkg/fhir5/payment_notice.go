package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// PaymentNoticeStatus represents the status of a payment notice
type PaymentNoticeStatus string

const (
	PaymentNoticeStatusActive         PaymentNoticeStatus = "active"
	PaymentNoticeStatusCancelled      PaymentNoticeStatus = "cancelled"
	PaymentNoticeStatusDraft          PaymentNoticeStatus = "draft"
	PaymentNoticeStatusEnteredInError PaymentNoticeStatus = "entered-in-error"
)

// PaymentNotice represents a notice of a payment
type PaymentNotice struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "PaymentNotice"

	// The amount sent to the payee
	Amount Money `json:"amount"`

	// The date when this resource was created
	Created *string `json:"created,omitempty"`

	// The party who will receive or has received payment that is the subject of this notification
	Payee *common.Reference `json:"payee,omitempty"`

	// The party who notifies the payer about the payment being made
	Payer *common.Reference `json:"payer,omitempty"`

	// A note that describes or explains the processing in a human readable form
	PaymentDate *string `json:"paymentDate,omitempty"`

	// The date when the above payment action occurred
	PaymentStatus *common.CodeableConcept `json:"paymentStatus,omitempty"`

	// The practitioner who is responsible for the services rendered to the patient
	Provider *common.Reference `json:"provider,omitempty"`

	// A reference to the payment which is the subject of this notice
	Request *common.Reference `json:"request,omitempty"`

	// A reference to the payment which is the subject of this notice
	Response *common.Reference `json:"response,omitempty"`

	// The status of the resource instance
	Status PaymentNoticeStatus `json:"status"`

	// The target of the payment notice
	Target *common.Reference `json:"target,omitempty"`
}
