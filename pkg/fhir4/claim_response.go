package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// ClaimResponse represents the ClaimResponse resource
type ClaimResponse struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "ClaimResponse"

	// The first-tier service adjudications for payor added product or service lines
	AddItem []ClaimResponseAddItem `json:"addItem,omitempty"`

	// The adjudication results which are presented at the header level rather than at the line-item or add-item levels
	Adjudication []ClaimResponseItemAdjudication `json:"adjudication,omitempty"`

	// For example: professional reports, documents, images, clinical resources, or accident reports
	CommunicationRequest []common.Reference `json:"communicationRequest,omitempty"`

	// The date this resource was created
	Created string `json:"created"`

	// A human readable description of the status of the adjudication
	Disposition *string `json:"disposition,omitempty"`

	// If the request contains errors then an error element should be provided
	Error []ClaimResponseError `json:"error,omitempty"`

	// Needed to permit insurers to include the actual form
	Form *common.Attachment `json:"form,omitempty"`

	// May be needed to identify specific jurisdictional forms
	FormCode *common.CodeableConcept `json:"formCode,omitempty"`

	// Fund would be release by a future claim quoting the preAuthRef of this response
	FundsReserve *common.CodeableConcept `json:"fundsReserve,omitempty"`

	// A unique identifier assigned to this claim response
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// All insurance coverages for the patient which may be applicable for reimbursement
	Insurance []ClaimResponseInsurance `json:"insurance,omitempty"`

	// The party responsible for authorization, adjudication and reimbursement
	Insurer common.Reference `json:"insurer"`

	// A claim line. Either a simple (a product or service) or a 'group' of details
	Item []ClaimResponseItem `json:"item,omitempty"`

	// The resource may be used to indicate that: the request has been held (queued) for processing
	Outcome ClaimResponseOutcome `json:"outcome"`

	// The party to whom the professional services and/or products have been supplied
	Patient common.Reference `json:"patient"`

	// Type of Party to be reimbursed: subscriber, provider, other
	PayeeType *common.CodeableConcept `json:"payeeType,omitempty"`

	// Payment details for the adjudication of the claim
	Payment *ClaimResponsePayment `json:"payment,omitempty"`

	// The time frame during which this authorization is effective
	PreAuthPeriod *common.Period `json:"preAuthPeriod,omitempty"`

	// This value is only present on preauthorization adjudications
	PreAuthRef *string `json:"preAuthRef,omitempty"`

	// A note that describes or explains adjudication results in a human readable form
	ProcessNote []ClaimResponseProcessNote `json:"processNote,omitempty"`

	// The provider-required urgency of processing the request
	Request *common.Reference `json:"request,omitempty"`

	// The provider which is responsible for the claim, predetermination or preauthorization
	Requestor *common.Reference `json:"requestor,omitempty"`

	// Totals for amounts submitted, co-pays, benefits payable etc
	Total []ClaimResponseTotal `json:"total,omitempty"`

	// A code to indicate the type of claim
	Type common.CodeableConcept `json:"type"`

	// A code, used only on a response to a preauthorization, to indicate whether the benefits payable have been reserved and for whom
	Use ClaimUse `json:"use"`
}

// ClaimUse represents the use of a claim
type ClaimUse string

const (
	ClaimUseClaim   ClaimUse = "claim"
	ClaimUsePreauth ClaimUse = "preauthorization"
	ClaimUsePredet  ClaimUse = "predetermination"
)

// ClaimResponseOutcome represents the outcome of the claim response
type ClaimResponseOutcome string

const (
	ClaimResponseOutcomeQueued   ClaimResponseOutcome = "queued"
	ClaimResponseOutcomeComplete ClaimResponseOutcome = "complete"
	ClaimResponseOutcomeError    ClaimResponseOutcome = "error"
	ClaimResponseOutcomePartial  ClaimResponseOutcome = "partial"
)

// ClaimResponseItemAdjudication represents the adjudication results
type ClaimResponseItemAdjudication struct {
	common.BackboneElement

	// For example: amount submitted, eligible amount, co-payment, and benefit payable
	Amount *common.Money `json:"amount,omitempty"`

	// For example codes indicating: Co-Pay, deductible, eligible, benefit, tax, etc
	Category common.CodeableConcept `json:"category"`

	// For example may indicate that the funds for this benefit type have been exhausted
	Reason *common.CodeableConcept `json:"reason,omitempty"`

	// For example: eligible percentage or co-payment percentage
	Value *float64 `json:"value,omitempty"`
}

// ClaimResponseItemDetailSubDetail represents a sub-detail adjudication
type ClaimResponseItemDetailSubDetail struct {
	common.BackboneElement

	// The adjudication results
	Adjudication []ClaimResponseItemAdjudication `json:"adjudication,omitempty"`

	// The numbers associated with notes below which apply to the adjudication of this item
	NoteNumber []int `json:"noteNumber,omitempty"`

	// A number to uniquely reference the claim sub-detail entry
	SubDetailSequence int `json:"subDetailSequence"`
}

// ClaimResponseItemDetail represents a claim detail
type ClaimResponseItemDetail struct {
	common.BackboneElement

	// The adjudication results
	Adjudication []ClaimResponseItemAdjudication `json:"adjudication"`

	// A number to uniquely reference the claim detail entry
	DetailSequence int `json:"detailSequence"`

	// The numbers associated with notes below which apply to the adjudication of this item
	NoteNumber []int `json:"noteNumber,omitempty"`

	// A sub-detail adjudication of a simple product or service
	SubDetail []ClaimResponseItemDetailSubDetail `json:"subDetail,omitempty"`
}

// ClaimResponseItem represents a claim line
type ClaimResponseItem struct {
	common.BackboneElement

	// If this item is a group then the values here are a summary of the adjudication of the detail items
	Adjudication []ClaimResponseItemAdjudication `json:"adjudication"`

	// A claim detail. Either a simple (a product or service) or a 'group' of sub-details
	Detail []ClaimResponseItemDetail `json:"detail,omitempty"`

	// A number to uniquely reference the claim item entries
	ItemSequence int `json:"itemSequence"`

	// The numbers associated with notes below which apply to the adjudication of this item
	NoteNumber []int `json:"noteNumber,omitempty"`
}

// ClaimResponseAddItemDetailSubDetail represents the third-tier service adjudications
type ClaimResponseAddItemDetailSubDetail struct {
	common.BackboneElement

	// The adjudication results
	Adjudication []ClaimResponseItemAdjudication `json:"adjudication"`

	// To show a 10% senior's discount, the value entered is: 0.90 (1.00 - 0.10)
	Factor *float64 `json:"factor,omitempty"`

	// For example in Oral whether the treatment is cosmetic or associated with TMJ
	Modifier []common.CodeableConcept `json:"modifier,omitempty"`

	// For example, the formula: quantity * unitPrice * factor = net
	Net *common.Money `json:"net,omitempty"`

	// The numbers associated with notes below which apply to the adjudication of this item
	NoteNumber []int `json:"noteNumber,omitempty"`

	// If this is an actual service or product line, i.e. not a Group
	ProductOrService common.CodeableConcept `json:"productOrService"`

	// The number of repetitions of a service or product
	Quantity *common.Quantity `json:"quantity,omitempty"`

	// If the item is not a group then this is the fee for the product or service
	UnitPrice *common.Money `json:"unitPrice,omitempty"`
}

// ClaimResponseAddItemDetail represents the second-tier service adjudications
type ClaimResponseAddItemDetail struct {
	common.BackboneElement

	// The adjudication results
	Adjudication []ClaimResponseItemAdjudication `json:"adjudication"`

	// To show a 10% senior's discount, the value entered is: 0.90 (1.00 - 0.10)
	Factor *float64 `json:"factor,omitempty"`

	// For example in Oral whether the treatment is cosmetic or associated with TMJ
	Modifier []common.CodeableConcept `json:"modifier,omitempty"`

	// For example, the formula: quantity * unitPrice * factor = net
	Net *common.Money `json:"net,omitempty"`

	// The numbers associated with notes below which apply to the adjudication of this item
	NoteNumber []int `json:"noteNumber,omitempty"`

	// If this is an actual service or product line, i.e. not a Group
	ProductOrService common.CodeableConcept `json:"productOrService"`

	// The number of repetitions of a service or product
	Quantity *common.Quantity `json:"quantity,omitempty"`

	// The third-tier service adjudications for payor added services
	SubDetail []ClaimResponseAddItemDetailSubDetail `json:"subDetail,omitempty"`

	// If the item is not a group then this is the fee for the product or service
	UnitPrice *common.Money `json:"unitPrice,omitempty"`
}

// ClaimResponseAddItem represents the first-tier service adjudications
type ClaimResponseAddItem struct {
	common.BackboneElement

	// The adjudication results
	Adjudication []ClaimResponseItemAdjudication `json:"adjudication"`

	// For example: Providing a tooth code allows an insurer to identify a provider performing a filling
	BodySite *common.CodeableConcept `json:"bodySite,omitempty"`

	// The second-tier service adjudications for payor added services
	Detail []ClaimResponseAddItemDetail `json:"detail,omitempty"`

	// The sequence number of the details within the claim item which this line is intended to replace
	DetailSequence []int `json:"detailSequence,omitempty"`

	// To show a 10% senior's discount, the value entered is: 0.90 (1.00 - 0.10)
	Factor *float64 `json:"factor,omitempty"`

	// Claim items which this service line is intended to replace
	ItemSequence []int `json:"itemSequence,omitempty"`

	// Where the product or service was provided
	LocationCodeableConcept *common.CodeableConcept `json:"locationCodeableConcept,omitempty"`

	// Where the product or service was provided
	LocationAddress *common.Address `json:"locationAddress,omitempty"`

	// Where the product or service was provided
	LocationReference *common.Reference `json:"locationReference,omitempty"`

	// For example in Oral whether the treatment is cosmetic or associated with TMJ
	Modifier []common.CodeableConcept `json:"modifier,omitempty"`

	// For example, the formula: quantity * unitPrice * factor = net
	Net *common.Money `json:"net,omitempty"`

	// The numbers associated with notes below which apply to the adjudication of this item
	NoteNumber []int `json:"noteNumber,omitempty"`

	// If this is an actual service or product line, i.e. not a Group
	ProductOrService common.CodeableConcept `json:"productOrService"`

	// For example: Neonatal program, child dental program or drug users recovery program
	ProgramCode []common.CodeableConcept `json:"programCode,omitempty"`

	// The providers who are authorized for the services rendered to the patient
	Provider []common.Reference `json:"provider,omitempty"`

	// The number of repetitions of a service or product
	Quantity *common.Quantity `json:"quantity,omitempty"`

	// The date or dates when the service or product was supplied, performed or completed
	ServicedDate *string `json:"servicedDate,omitempty"`

	// The date or dates when the service or product was supplied, performed or completed
	ServicedPeriod *common.Period `json:"servicedPeriod,omitempty"`

	// The sequence number of the sub-details within the details within the claim item which this line is intended to replace
	SubdetailSequence []int `json:"subdetailSequence,omitempty"`

	// A region or surface of the bodySite, e.g. limb region or tooth surface(s)
	SubSite []common.CodeableConcept `json:"subSite,omitempty"`

	// If the item is not a group then this is the fee for the product or service
	UnitPrice *common.Money `json:"unitPrice,omitempty"`
}

// ClaimResponseTotal represents totals for amounts submitted, co-pays, benefits payable etc
type ClaimResponseTotal struct {
	common.BackboneElement

	// Monetary total amount associated with the category
	Amount common.Money `json:"amount"`

	// For example codes indicating: Co-Pay, deductible, eligible, benefit, tax, etc
	Category common.CodeableConcept `json:"category"`
}

// ClaimResponsePayment represents payment details for the adjudication of the claim
type ClaimResponsePayment struct {
	common.BackboneElement

	// Insurers will deduct amounts owing from the provider (adjustment)
	Adjustment *common.Money `json:"adjustment,omitempty"`

	// Reason for the payment adjustment
	AdjustmentReason *common.CodeableConcept `json:"adjustmentReason,omitempty"`

	// Benefits payable less any payment adjustment
	Amount common.Money `json:"amount"`

	// Estimated date the payment will be issued or the actual issue date of payment
	Date *string `json:"date,omitempty"`

	// For example: EFT number or check number
	Identifier *common.Identifier `json:"identifier,omitempty"`

	// Whether this represents partial or complete payment of the benefits payable
	Type common.CodeableConcept `json:"type"`
}

// ClaimResponseProcessNote represents a note that describes or explains adjudication results
type ClaimResponseProcessNote struct {
	common.BackboneElement

	// Only required if the language is different from the resource language
	Language *common.CodeableConcept `json:"language,omitempty"`

	// A number to uniquely identify a note entry
	Number *int `json:"number,omitempty"`

	// The explanation or description associated with the processing
	Text string `json:"text"`

	// The business purpose of the note text
	Type *ClaimResponseProcessNoteType `json:"type,omitempty"`
}

// ClaimResponseProcessNoteType represents the business purpose of the note text
type ClaimResponseProcessNoteType string

const (
	ClaimResponseProcessNoteTypeDisplay   ClaimResponseProcessNoteType = "display"
	ClaimResponseProcessNoteTypePrint     ClaimResponseProcessNoteType = "print"
	ClaimResponseProcessNoteTypePrintoper ClaimResponseProcessNoteType = "printoper"
)

// ClaimResponseInsurance represents insurance coverages for the patient
type ClaimResponseInsurance struct {
	common.BackboneElement

	// A business agreement number established between the provider and the insurer
	BusinessArrangement *string `json:"businessArrangement,omitempty"`

	// Must not be specified when 'focal=true' for this insurance
	ClaimResponse *common.Reference `json:"claimResponse,omitempty"`

	// Reference to the insurance card level information contained in the Coverage resource
	Coverage common.Reference `json:"coverage"`

	// A patient may (will) have multiple insurance policies which provide reimbursement
	Focal bool `json:"focal"`

	// A number to uniquely identify insurance entries and provide a sequence of coverages
	Sequence int `json:"sequence"`
}

// ClaimResponseError represents an error in the claim response
type ClaimResponseError struct {
	common.BackboneElement

	// An error code, from a specified code system, which details why the claim could not be adjudicated
	Code common.CodeableConcept `json:"code"`

	// The sequence number of the detail within the line item submitted which contains the error
	DetailSequence *int `json:"detailSequence,omitempty"`

	// The sequence number of the line item submitted which contains the error
	ItemSequence *int `json:"itemSequence,omitempty"`

	// The sequence number of the sub-detail within the detail within the line item submitted which contains the error
	SubDetailSequence *int `json:"subDetailSequence,omitempty"`
}
