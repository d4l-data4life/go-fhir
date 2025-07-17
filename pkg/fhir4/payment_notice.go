package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// PaymentNotice represents the status of the payment for goods and services rendered
type PaymentNotice struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "PaymentNotice"

	// The amount sent to the payee
	Amount common.Money `json:"amount"`

	// The date when this resource was created
	Created string `json:"created"`

	// A unique identifier assigned to this payment notice
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The party who will receive or has received payment that is the subject of this notification
	Payee *common.Reference `json:"payee,omitempty"`

	// A reference to the payment which is the subject of this notice
	Payment common.Reference `json:"payment"`

	// The date when the above payment action occurred
	PaymentDate *string `json:"paymentDate,omitempty"`

	// Typically paid: payment sent, cleared: payment received
	PaymentStatus *common.CodeableConcept `json:"paymentStatus,omitempty"`

	// The practitioner who is responsible for the services rendered to the patient
	Provider *common.Reference `json:"provider,omitempty"`

	// The party who is notified of the payment status
	Recipient common.Reference `json:"recipient"`

	// Reference of resource for which payment is being made
	Request *common.Reference `json:"request,omitempty"`

	// Reference of response to resource for which payment is being made
	Response *common.Reference `json:"response,omitempty"`

	// This element is labeled as a modifier because the status contains codes that mark the resource as not currently valid
	Status PaymentNoticeStatus `json:"status"`
}

// PaymentNoticeStatus represents this element is labeled as a modifier because the status contains codes that mark the resource as not currently valid
type PaymentNoticeStatus string

const (
	PaymentNoticeStatusActive         PaymentNoticeStatus = "active"
	PaymentNoticeStatusCancelled      PaymentNoticeStatus = "cancelled"
	PaymentNoticeStatusDraft          PaymentNoticeStatus = "draft"
	PaymentNoticeStatusEnteredInError PaymentNoticeStatus = "entered-in-error"
)
