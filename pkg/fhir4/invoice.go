package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// Invoice represents invoice containing collected ChargeItems from an Account with calculated individual and total price
type Invoice struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Invoice"

	// Systems posting the ChargeItems might not always be able to determine, which accounts the Items need to be places into
	Account *common.Reference `json:"account,omitempty"`

	// Derived Profiles may choose to add invariants requiring this field to be populated
	CancelledReason *string `json:"cancelledReason,omitempty"`

	// The list of types may be constrained as appropriate for the type of charge item
	Date *string `json:"date,omitempty"`

	// Identifier of this Invoice, often used for reference in correspondence about this invoice
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Practitioners and Devices can be associated with multiple organizations
	Issuer *common.Reference `json:"issuer,omitempty"`

	// Each line item represents one charge for goods and services rendered
	LineItem []InvoiceLineItem `json:"lineItem,omitempty"`

	// Comments made about the invoice by the issuer, subject, or other participants
	Note []common.Annotation `json:"note,omitempty"`

	// Indicates who or what performed or participated in the charged service
	Participant []InvoiceParticipant `json:"participant,omitempty"`

	// Derived Profiles may chose to add invariants requiring this field to be populated
	PaymentTerms *string `json:"paymentTerms,omitempty"`

	// The individual or Organization responsible for balancing of this invoice
	Recipient *common.Reference `json:"recipient,omitempty"`

	// The current state of the Invoice
	Status InvoiceStatus `json:"status"`

	// The individual or set of individuals receiving the goods and services billed in this invoice
	Subject *common.Reference `json:"subject,omitempty"`

	// There is no reason to carry the price in the instance of a ChargeItem unless circumstances require a manual override
	TotalGross *common.Money `json:"totalGross,omitempty"`

	// There is no reason to carry the price in the instance of a ChargeItem unless circumstances require a manual override
	TotalNet *common.Money `json:"totalNet,omitempty"`

	// The total amount for the Invoice may be calculated as the sum of the line items with surcharges/deductions
	TotalPriceComponent []InvoiceLineItemPriceComponent `json:"totalPriceComponent,omitempty"`

	// Type of Invoice depending on domain, realm an usage (e.g. internal/external, dental, preliminary)
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// InvoiceStatus represents the status of the invoice
type InvoiceStatus string

const (
	InvoiceStatusDraft          InvoiceStatus = "draft"
	InvoiceStatusIssued         InvoiceStatus = "issued"
	InvoiceStatusBalanced       InvoiceStatus = "balanced"
	InvoiceStatusCancelled      InvoiceStatus = "cancelled"
	InvoiceStatusEnteredInError InvoiceStatus = "entered-in-error"
)

// InvoiceParticipant represents indicates who or what performed or participated in the charged service
type InvoiceParticipant struct {
	common.BackboneElement

	// The device, practitioner, etc. who performed or participated in the service
	Actor common.Reference `json:"actor"`

	// Describes the type of involvement (e.g. transcriptionist, creator etc.)
	Role *common.CodeableConcept `json:"role,omitempty"`
}

// InvoiceLineItem represents each line item represents one charge for goods and services rendered
type InvoiceLineItem struct {
	common.BackboneElement

	// The ChargeItem contains information such as the billing code, date, amount etc
	ChargeItemReference *common.Reference `json:"chargeItemReference,omitempty"`

	// The ChargeItem contains information such as the billing code, date, amount etc
	ChargeItemCodeableConcept *common.CodeableConcept `json:"chargeItemCodeableConcept,omitempty"`

	// The price for a ChargeItem may be calculated as a base price with surcharges/deductions
	PriceComponent []InvoiceLineItemPriceComponent `json:"priceComponent,omitempty"`

	// Sequence in which the items appear on the invoice
	Sequence *int `json:"sequence,omitempty"`
}

// InvoiceLineItemPriceComponent represents the price for a ChargeItem may be calculated as a base price with surcharges/deductions
type InvoiceLineItemPriceComponent struct {
	common.BackboneElement

	// There is no reason to carry the price in the instance of a ChargeItem unless circumstances require a manual override
	Amount *common.Money `json:"amount,omitempty"`

	// A code that identifies the component. Codes may be used to differentiate between kinds of taxes, surcharges, discounts etc
	Code *common.CodeableConcept `json:"code,omitempty"`

	// There is no reason to carry the price in the instance of a ChargeItem unless circumstances require a manual override
	Factor *float64 `json:"factor,omitempty"`

	// This code identifies the type of the component
	Type InvoiceLineItemPriceComponentType `json:"type"`
}

// InvoiceLineItemPriceComponentType represents the type of price component
type InvoiceLineItemPriceComponentType string

const (
	InvoiceLineItemPriceComponentTypeBase          InvoiceLineItemPriceComponentType = "base"
	InvoiceLineItemPriceComponentTypeSurcharge     InvoiceLineItemPriceComponentType = "surcharge"
	InvoiceLineItemPriceComponentTypeDeduction     InvoiceLineItemPriceComponentType = "deduction"
	InvoiceLineItemPriceComponentTypeDiscount      InvoiceLineItemPriceComponentType = "discount"
	InvoiceLineItemPriceComponentTypeTax           InvoiceLineItemPriceComponentType = "tax"
	InvoiceLineItemPriceComponentTypeInformational InvoiceLineItemPriceComponentType = "informational"
)
