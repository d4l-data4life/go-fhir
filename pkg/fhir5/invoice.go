package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// InvoiceStatus represents the current state of the Invoice
type InvoiceStatus string

const (
	InvoiceStatusDraft          InvoiceStatus = "draft"
	InvoiceStatusIssued         InvoiceStatus = "issued"
	InvoiceStatusBalanced       InvoiceStatus = "balanced"
	InvoiceStatusCancelled      InvoiceStatus = "cancelled"
	InvoiceStatusEnteredInError InvoiceStatus = "entered-in-error"
)

// InvoiceParticipant represents participants in the invoice
type InvoiceParticipant struct {
	common.BackboneElement

	// The device, practitioner, etc. who performed or participated in the service
	Actor common.Reference `json:"actor"`

	// Describes the type of involvement (e.g. transcriptionist, creator etc.)
	Role *common.CodeableConcept `json:"role,omitempty"`
}

// InvoiceLineItem represents each line item representing one charge for goods and services rendered
type InvoiceLineItem struct {
	common.BackboneElement

	// The ChargeItem contains information such as the billing code, date, amount etc.
	ChargeItemReference       *common.Reference       `json:"chargeItemReference,omitempty"`
	ChargeItemCodeableConcept *common.CodeableConcept `json:"chargeItemCodeableConcept,omitempty"`

	// The price for a ChargeItem may be calculated as a base price with surcharges/deductions
	PriceComponent []MonetaryComponent `json:"priceComponent,omitempty"`

	// Sequence in which the items appear on the invoice
	Sequence *int `json:"sequence,omitempty"`

	// Date/time(s) range when this service was delivered or completed
	ServicedDate   *string        `json:"servicedDate,omitempty"`
	ServicedPeriod *common.Period `json:"servicedPeriod,omitempty"`
}

// Invoice represents an invoice for goods and services
type Invoice struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Invoice"

	// Account that is supposed to be balanced with this Invoice
	Account *common.Reference `json:"account,omitempty"`

	// Reason for cancellation of this Invoice
	CancelledReason *string `json:"cancelledReason,omitempty"`

	// When this Invoice was posted
	Creation *string `json:"creation,omitempty"`

	// Deprecated by the element below
	Date *string `json:"date,omitempty"`

	// Identifier of this Invoice
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The issuer of the Invoice
	Issuer *common.Reference `json:"issuer,omitempty"`

	// Each line item represents one charge for goods and services rendered
	LineItem []InvoiceLineItem `json:"lineItem,omitempty"`

	// Comments made about the invoice by the issuer, subject, or other participants
	Note []Annotation `json:"note,omitempty"`

	// Indicates who or what performed or participated in the charged service
	Participant []InvoiceParticipant `json:"participant,omitempty"`

	// Payment details
	PaymentTerms *string `json:"paymentTerms,omitempty"`

	// Date/time(s) range of services included in this invoice
	PeriodDate   *string        `json:"periodDate,omitempty"`
	PeriodPeriod *common.Period `json:"periodPeriod,omitempty"`

	// The individual or Organization responsible for balancing of this invoice
	Recipient *common.Reference `json:"recipient,omitempty"`

	// The current state of the Invoice
	Status InvoiceStatus `json:"status"`

	// The individual or set of individuals receiving the goods and services billed in this invoice
	Subject *common.Reference `json:"subject,omitempty"`

	// Invoice Totals
	TotalGross          *Money              `json:"totalGross,omitempty"`
	TotalNet            *Money              `json:"totalNet,omitempty"`
	TotalPriceComponent []MonetaryComponent `json:"totalPriceComponent,omitempty"`

	// Type of Invoice depending on domain, realm an usage
	Type *common.CodeableConcept `json:"type,omitempty"`
}
