package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// ExplanationOfBenefit represents the ExplanationOfBenefit resource
type ExplanationOfBenefit struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "ExplanationOfBenefit"

	// Details of a accident which resulted in injuries which required the products and services listed in the claim
	Accident *ExplanationOfBenefitAccident `json:"accident,omitempty"`

	// The first-tier service adjudications for payor added product or service lines
	AddItem []ExplanationOfBenefitAddItem `json:"addItem,omitempty"`

	// The adjudication results which are presented at the header level rather than at the line-item or add-item levels
	Adjudication []ExplanationOfBenefitItemAdjudication `json:"adjudication,omitempty"`

	// Balance by Benefit Category
	BenefitBalance []ExplanationOfBenefitBenefitBalance `json:"benefitBalance,omitempty"`

	// Not applicable when use=claim
	BenefitPeriod *common.Period `json:"benefitPeriod,omitempty"`

	// Typically this would be today or in the past for a claim, and today or in the future for preauthorizations and prodeterminations
	BillablePeriod *common.Period `json:"billablePeriod,omitempty"`

	// The members of the team who provided the products and services
	CareTeam []ExplanationOfBenefitCareTeam `json:"careTeam,omitempty"`

	// The business identifier for the instance of the adjudication request: claim predetermination or preauthorization
	Claim *common.Reference `json:"claim,omitempty"`

	// The business identifier for the instance of the adjudication response: claim, predetermination or preauthorization response
	ClaimResponse *common.Reference `json:"claimResponse,omitempty"`

	// This field is independent of the date of creation of the resource as it may reflect the creation date of a source document prior to digitization
	Created string `json:"created"`

	// Information about diagnoses relevant to the claim items
	Diagnosis []ExplanationOfBenefitDiagnosis `json:"diagnosis,omitempty"`

	// A human readable description of the status of the adjudication
	Disposition *string `json:"disposition,omitempty"`

	// Individual who created the claim, predetermination or preauthorization
	Enterer *common.Reference `json:"enterer,omitempty"`

	// Facility where the services were provided
	Facility *common.Reference `json:"facility,omitempty"`

	// Needed to permit insurers to include the actual form
	Form *common.Attachment `json:"form,omitempty"`

	// May be needed to identify specific jurisdictional forms
	FormCode *common.CodeableConcept `json:"formCode,omitempty"`

	// Fund would be release by a future claim quoting the preAuthRef of this response
	FundsReserve *common.CodeableConcept `json:"fundsReserve,omitempty"`

	// This field is only used for preauthorizations
	FundsReserveRequested *common.CodeableConcept `json:"fundsReserveRequested,omitempty"`

	// A unique identifier assigned to this explanation of benefit
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// All insurance coverages for the patient which may be applicable for reimbursement
	Insurance []ExplanationOfBenefitInsurance `json:"insurance"`

	// The party responsible for authorization, adjudication and reimbursement
	Insurer common.Reference `json:"insurer"`

	// A claim line. Either a simple (a product or service) or a 'group' of details which can also be a simple items or groups of sub-details
	Item []ExplanationOfBenefitItem `json:"item,omitempty"`

	// For example, a physician may prescribe a medication which the pharmacy determines is contraindicated
	OriginalPrescription *common.Reference `json:"originalPrescription,omitempty"`

	// The resource may be used to indicate that: the request has been held (queued) for processing
	Outcome ExplanationOfBenefitOutcome `json:"outcome"`

	// The party to whom the professional services and/or products have been supplied or are being considered
	Patient common.Reference `json:"patient"`

	// Often providers agree to receive the benefits payable to reduce the near-term costs to the patient
	Payee *ExplanationOfBenefitPayee `json:"payee,omitempty"`

	// Payment details for the adjudication of the claim
	Payment *ExplanationOfBenefitPayment `json:"payment,omitempty"`

	// This value is only present on preauthorization adjudications
	PreAuthRef []string `json:"preAuthRef,omitempty"`

	// This value is only present on preauthorization adjudications
	PreAuthRefPeriod []common.Period `json:"preAuthRefPeriod,omitempty"`

	// This indicates the relative order of a series of EOBs related to different coverages for the same suite of services
	Precedence *int `json:"precedence,omitempty"`

	// Prescription to support the dispensing of pharmacy, device or vision products
	Prescription *common.Reference `json:"prescription,omitempty"`

	// If a claim processor is unable to complete the processing as per the priority then they should generate and error and not process the request
	Priority *common.CodeableConcept `json:"priority,omitempty"`

	// Procedures performed on the patient relevant to the billing items with the claim
	Procedure []ExplanationOfBenefitProcedure `json:"procedure,omitempty"`

	// A note that describes or explains adjudication results in a human readable form
	ProcessNote []ExplanationOfBenefitProcessNote `json:"processNote,omitempty"`

	// Typically this field would be 1..1 where this party is responsible for the claim but not necessarily professionally responsible
	Provider common.Reference `json:"provider"`

	// The referral resource which lists the date, practitioner, reason and other supporting information
	Referral *common.Reference `json:"referral,omitempty"`

	// For example, for the original treatment and follow-up exams
	Related []ExplanationOfBenefitRelated `json:"related,omitempty"`

	// This element is labeled as a modifier because the status contains codes that mark the resource as not currently valid
	Status ExplanationOfBenefitStatus `json:"status"`

	// This may contain the local bill type codes such as the US UB-04 bill type code
	SubType *common.CodeableConcept `json:"subType,omitempty"`

	// Often there are multiple jurisdiction specific valuesets which are required
	SupportingInfo []ExplanationOfBenefitSupportingInfo `json:"supportingInfo,omitempty"`

	// Totals for amounts submitted, co-pays, benefits payable etc
	Total []ExplanationOfBenefitTotal `json:"total,omitempty"`

	// The majority of jurisdictions use: oral, pharmacy, vision, professional and institutional, or variants on those terms
	Type common.CodeableConcept `json:"type"`

	// A code to indicate whether the nature of the request is: to request adjudication of products and services previously rendered
	Use ExplanationOfBenefitUse `json:"use"`
}

// ExplanationOfBenefitOutcome represents the outcome of the explanation of benefit
type ExplanationOfBenefitOutcome string

const (
	ExplanationOfBenefitOutcomeQueued   ExplanationOfBenefitOutcome = "queued"
	ExplanationOfBenefitOutcomeComplete ExplanationOfBenefitOutcome = "complete"
	ExplanationOfBenefitOutcomeError    ExplanationOfBenefitOutcome = "error"
	ExplanationOfBenefitOutcomePartial  ExplanationOfBenefitOutcome = "partial"
)

// ExplanationOfBenefitStatus represents the status of the explanation of benefit
type ExplanationOfBenefitStatus string

const (
	ExplanationOfBenefitStatusActive         ExplanationOfBenefitStatus = "active"
	ExplanationOfBenefitStatusCancelled      ExplanationOfBenefitStatus = "cancelled"
	ExplanationOfBenefitStatusDraft          ExplanationOfBenefitStatus = "draft"
	ExplanationOfBenefitStatusEnteredInError ExplanationOfBenefitStatus = "entered-in-error"
)

// ExplanationOfBenefitUse represents the use of the explanation of benefit
type ExplanationOfBenefitUse string

const (
	ExplanationOfBenefitUseClaim            ExplanationOfBenefitUse = "claim"
	ExplanationOfBenefitUsePreauthorization ExplanationOfBenefitUse = "preauthorization"
	ExplanationOfBenefitUsePredetermination ExplanationOfBenefitUse = "predetermination"
)

// ExplanationOfBenefitRelated represents related claims
type ExplanationOfBenefitRelated struct {
	common.BackboneElement

	// Reference to a related claim
	Claim *common.Reference `json:"claim,omitempty"`

	// For example, Property/Casualty insurer claim number or Workers Compensation case number
	Reference *common.Identifier `json:"reference,omitempty"`

	// For example, prior claim or umbrella
	Relationship *common.CodeableConcept `json:"relationship,omitempty"`
}

// ExplanationOfBenefitPayee represents the payee information
type ExplanationOfBenefitPayee struct {
	common.BackboneElement

	// Not required if the payee is 'subscriber' or 'provider'
	Party *common.Reference `json:"party,omitempty"`

	// Type of Party to be reimbursed: Subscriber, provider, other
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// ExplanationOfBenefitCareTeam represents the members of the team who provided the products and services
type ExplanationOfBenefitCareTeam struct {
	common.BackboneElement

	// Member of the team who provided the product or service
	Provider common.Reference `json:"provider"`

	// The qualification of the practitioner which is applicable for this service
	Qualification *common.CodeableConcept `json:"qualification,omitempty"`

	// Responsible might not be required when there is only a single provider listed
	Responsible *bool `json:"responsible,omitempty"`

	// Role might not be required when there is only a single provider listed
	Role *common.CodeableConcept `json:"role,omitempty"`

	// A number to uniquely identify care team entries
	Sequence int `json:"sequence"`
}

// ExplanationOfBenefitSupportingInfo represents supporting information
type ExplanationOfBenefitSupportingInfo struct {
	common.BackboneElement

	// This may contain a category for the local bill type codes
	Category common.CodeableConcept `json:"category"`

	// This may contain the local bill type codes such as the US UB-04 bill type code
	Code *common.CodeableConcept `json:"code,omitempty"`

	// For example: the reason for the additional stay, or why a tooth is missing
	Reason *common.Coding `json:"reason,omitempty"`

	// A number to uniquely identify supporting information entries
	Sequence int `json:"sequence"`

	// The date when or period to which this information refers
	TimingDate *string `json:"timingDate,omitempty"`

	// The date when or period to which this information refers
	TimingPeriod *common.Period `json:"timingPeriod,omitempty"`

	// Could be used to provide references to other resources, document
	ValueBoolean *bool `json:"valueBoolean,omitempty"`

	// Could be used to provide references to other resources, document
	ValueString *string `json:"valueString,omitempty"`

	// Could be used to provide references to other resources, document
	ValueQuantity *common.Quantity `json:"valueQuantity,omitempty"`

	// Could be used to provide references to other resources, document
	ValueAttachment *common.Attachment `json:"valueAttachment,omitempty"`

	// Could be used to provide references to other resources, document
	ValueReference *common.Reference `json:"valueReference,omitempty"`
}

// ExplanationOfBenefitDiagnosis represents information about diagnoses relevant to the claim items
type ExplanationOfBenefitDiagnosis struct {
	common.BackboneElement

	// The nature of illness or problem in a coded form or as a reference to an external defined Condition
	DiagnosisCodeableConcept *common.CodeableConcept `json:"diagnosisCodeableConcept,omitempty"`

	// The nature of illness or problem in a coded form or as a reference to an external defined Condition
	DiagnosisReference *common.Reference `json:"diagnosisReference,omitempty"`

	// For example, admitting, primary, secondary, discharge
	PackageCode *common.CodeableConcept `json:"packageCode,omitempty"`

	// A number to uniquely identify diagnosis entries
	Sequence int `json:"sequence"`

	// The type of diagnosis (for example, admitting, primary, secondary, discharge)
	Type []common.CodeableConcept `json:"type,omitempty"`
}

// ExplanationOfBenefitProcedure represents procedures performed on the patient relevant to the billing items
type ExplanationOfBenefitProcedure struct {
	common.BackboneElement

	// Date and optionally time the procedure was performed
	Date *string `json:"date,omitempty"`

	// The code or reference to a Procedure resource which identifies the clinical intervention performed
	ProcedureCodeableConcept *common.CodeableConcept `json:"procedureCodeableConcept,omitempty"`

	// The code or reference to a Procedure resource which identifies the clinical intervention performed
	ProcedureReference *common.Reference `json:"procedureReference,omitempty"`

	// Sequence of procedures which serves to order and provide a link
	Sequence int `json:"sequence"`

	// When the condition was observed or the relative ranking
	Type []common.CodeableConcept `json:"type,omitempty"`

	// Unique Device Identifiers associated with this line item or detail item
	Udi []common.Reference `json:"udi,omitempty"`
}

// ExplanationOfBenefitInsurance represents insurance coverage information
type ExplanationOfBenefitInsurance struct {
	common.BackboneElement

	// Reference to the insurance card level information contained in the Coverage resource
	Coverage common.Reference `json:"coverage"`

	// A flag to indicate that this Coverage is to be used for adjudication of this claim when set to true
	Focal *bool `json:"focal,omitempty"`

	// Reference to the original request resource
	PreAuthRef []string `json:"preAuthRef,omitempty"`

	// Reference to the original request resource
	PreAuthRefPeriod []common.Period `json:"preAuthRefPeriod,omitempty"`

	// Sequence of the referenced insurance card level information contained in the Coverage resource
	Sequence int `json:"sequence"`
}

// ExplanationOfBenefitAccident represents details of an accident
type ExplanationOfBenefitAccident struct {
	common.BackboneElement

	// Date of an accident event related to the products and services contained in the claim
	Date *string `json:"date,omitempty"`

	// The physical location of the accident event
	LocationAddress *common.Address `json:"locationAddress,omitempty"`

	// The physical location of the accident event
	LocationReference *common.Reference `json:"locationReference,omitempty"`

	// The type or context of the accident event for the purposes of selection of potential insurance coverages and determination of coordination between insurers
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// ExplanationOfBenefitItemAdjudication represents adjudication results
type ExplanationOfBenefitItemAdjudication struct {
	common.BackboneElement

	// Monetary amount associated with the category
	Amount *common.Money `json:"amount,omitempty"`

	// A code to indicate the information type of this adjudication record
	Category common.CodeableConcept `json:"category"`

	// A code supporting the understanding of the adjudication result and explaining variance from expected amount
	Reason *common.CodeableConcept `json:"reason,omitempty"`

	// A non-monetary value associated with the category
	Value *float64 `json:"value,omitempty"`
}

// ExplanationOfBenefitItemDetailSubDetail represents sub-details of an item detail
type ExplanationOfBenefitItemDetailSubDetail struct {
	common.BackboneElement

	// The adjudication results
	Adjudication []ExplanationOfBenefitItemAdjudication `json:"adjudication,omitempty"`

	// The type of revenue or cost center providing the product and/or service
	Category *common.CodeableConcept `json:"category,omitempty"`

	// A real world service or product
	Modifier []common.CodeableConcept `json:"modifier,omitempty"`

	// The number of repetitions of a service or product
	Quantity *common.Quantity `json:"quantity,omitempty"`

	// The fee charged for the professional service or product
	Revenue *common.CodeableConcept `json:"revenue,omitempty"`

	// A number to uniquely identify sub-detail entries
	Sequence int `json:"sequence"`

	// A code to identify the general type of benefits under which products and services are provided
	Service *common.CodeableConcept `json:"service,omitempty"`

	// Unique Device Identifiers associated with this line item or detail item
	Udi []common.Reference `json:"udi,omitempty"`

	// The charge amount for the service or product
	UnitPrice *common.Money `json:"unitPrice,omitempty"`
}

// ExplanationOfBenefitItemDetail represents details of an item
type ExplanationOfBenefitItemDetail struct {
	common.BackboneElement

	// The adjudication results
	Adjudication []ExplanationOfBenefitItemAdjudication `json:"adjudication,omitempty"`

	// The type of revenue or cost center providing the product and/or service
	Category *common.CodeableConcept `json:"category,omitempty"`

	// A real world service or product
	Modifier []common.CodeableConcept `json:"modifier,omitempty"`

	// The number of repetitions of a service or product
	Quantity *common.Quantity `json:"quantity,omitempty"`

	// The fee charged for the professional service or product
	Revenue *common.CodeableConcept `json:"revenue,omitempty"`

	// A number to uniquely identify detail entries
	Sequence int `json:"sequence"`

	// A code to identify the general type of benefits under which products and services are provided
	Service *common.CodeableConcept `json:"service,omitempty"`

	// A claim detail line. Either a simple (a product or service) or a 'group' of sub-details which are simple items
	SubDetail []ExplanationOfBenefitItemDetailSubDetail `json:"subDetail,omitempty"`

	// Unique Device Identifiers associated with this line item or detail item
	Udi []common.Reference `json:"udi,omitempty"`

	// The charge amount for the service or product
	UnitPrice *common.Money `json:"unitPrice,omitempty"`
}

// ExplanationOfBenefitItem represents a claim line
type ExplanationOfBenefitItem struct {
	common.BackboneElement

	// The adjudication results
	Adjudication []ExplanationOfBenefitItemAdjudication `json:"adjudication,omitempty"`

	// Physical location where the service is performed or applies
	BodySite []common.CodeableConcept `json:"bodySite,omitempty"`

	// CareTeam members related to this service or product
	CareTeamSequence []int `json:"careTeamSequence,omitempty"`

	// The type of revenue or cost center providing the product and/or service
	Category *common.CodeableConcept `json:"category,omitempty"`

	// A claim detail line. Either a simple (a product or service) or a 'group' of sub-details which are simple items
	Detail []ExplanationOfBenefitItemDetail `json:"detail,omitempty"`

	// Diagnosis applicable for this service or product
	DiagnosisSequence []int `json:"diagnosisSequence,omitempty"`

	// Encounters related to this billed item
	Encounter []common.Reference `json:"encounter,omitempty"`

	// A real world service or product
	Modifier []common.CodeableConcept `json:"modifier,omitempty"`

	// The number of repetitions of a service or product
	Quantity *common.Quantity `json:"quantity,omitempty"`

	// The fee charged for the professional service or product
	Revenue *common.CodeableConcept `json:"revenue,omitempty"`

	// A number to uniquely identify item entries
	Sequence int `json:"sequence"`

	// A code to identify the general type of benefits under which products and services are provided
	Service *common.CodeableConcept `json:"service,omitempty"`

	// The date or dates when the service or product was supplied, performed or completed
	ServicedDate *string `json:"servicedDate,omitempty"`

	// The date or dates when the service or product was supplied, performed or completed
	ServicedPeriod *common.Period `json:"servicedPeriod,omitempty"`

	// Unique Device Identifiers associated with this line item or detail item
	Udi []common.Reference `json:"udi,omitempty"`

	// The charge amount for the service or product
	UnitPrice *common.Money `json:"unitPrice,omitempty"`
}

// ExplanationOfBenefitAddItemDetailSubDetail represents sub-details of an add item detail
type ExplanationOfBenefitAddItemDetailSubDetail struct {
	common.BackboneElement

	// The adjudication results
	Adjudication []ExplanationOfBenefitItemAdjudication `json:"adjudication,omitempty"`

	// The type of revenue or cost center providing the product and/or service
	Category *common.CodeableConcept `json:"category,omitempty"`

	// A real world service or product
	Modifier []common.CodeableConcept `json:"modifier,omitempty"`

	// The number of repetitions of a service or product
	Quantity *common.Quantity `json:"quantity,omitempty"`

	// The fee charged for the professional service or product
	Revenue *common.CodeableConcept `json:"revenue,omitempty"`

	// A code to identify the general type of benefits under which products and services are provided
	Service *common.CodeableConcept `json:"service,omitempty"`

	// Unique Device Identifiers associated with this line item or detail item
	Udi []common.Reference `json:"udi,omitempty"`

	// The charge amount for the service or product
	UnitPrice *common.Money `json:"unitPrice,omitempty"`
}

// ExplanationOfBenefitAddItemDetail represents details of an add item
type ExplanationOfBenefitAddItemDetail struct {
	common.BackboneElement

	// The adjudication results
	Adjudication []ExplanationOfBenefitItemAdjudication `json:"adjudication,omitempty"`

	// The type of revenue or cost center providing the product and/or service
	Category *common.CodeableConcept `json:"category,omitempty"`

	// A real world service or product
	Modifier []common.CodeableConcept `json:"modifier,omitempty"`

	// The number of repetitions of a service or product
	Quantity *common.Quantity `json:"quantity,omitempty"`

	// The fee charged for the professional service or product
	Revenue *common.CodeableConcept `json:"revenue,omitempty"`

	// A code to identify the general type of benefits under which products and services are provided
	Service *common.CodeableConcept `json:"service,omitempty"`

	// A claim detail line. Either a simple (a product or service) or a 'group' of sub-details which are simple items
	SubDetail []ExplanationOfBenefitAddItemDetailSubDetail `json:"subDetail,omitempty"`

	// Unique Device Identifiers associated with this line item or detail item
	Udi []common.Reference `json:"udi,omitempty"`

	// The charge amount for the service or product
	UnitPrice *common.Money `json:"unitPrice,omitempty"`
}

// ExplanationOfBenefitAddItem represents the first-tier service adjudications for payor added product or service lines
type ExplanationOfBenefitAddItem struct {
	common.BackboneElement

	// The adjudication results
	Adjudication []ExplanationOfBenefitItemAdjudication `json:"adjudication,omitempty"`

	// Physical location where the service is performed or applies
	BodySite []common.CodeableConcept `json:"bodySite,omitempty"`

	// CareTeam members related to this service or product
	CareTeamSequence []int `json:"careTeamSequence,omitempty"`

	// The type of revenue or cost center providing the product and/or service
	Category *common.CodeableConcept `json:"category,omitempty"`

	// A claim detail line. Either a simple (a product or service) or a 'group' of sub-details which are simple items
	Detail []ExplanationOfBenefitAddItemDetail `json:"detail,omitempty"`

	// Diagnosis applicable for this service or product
	DiagnosisSequence []int `json:"diagnosisSequence,omitempty"`

	// Encounters related to this billed item
	Encounter []common.Reference `json:"encounter,omitempty"`

	// A real world service or product
	Modifier []common.CodeableConcept `json:"modifier,omitempty"`

	// The number of repetitions of a service or product
	Quantity *common.Quantity `json:"quantity,omitempty"`

	// The fee charged for the professional service or product
	Revenue *common.CodeableConcept `json:"revenue,omitempty"`

	// A code to identify the general type of benefits under which products and services are provided
	Service *common.CodeableConcept `json:"service,omitempty"`

	// The date or dates when the service or product was supplied, performed or completed
	ServicedDate *string `json:"servicedDate,omitempty"`

	// The date or dates when the service or product was supplied, performed or completed
	ServicedPeriod *common.Period `json:"servicedPeriod,omitempty"`

	// Unique Device Identifiers associated with this line item or detail item
	Udi []common.Reference `json:"udi,omitempty"`

	// The charge amount for the service or product
	UnitPrice *common.Money `json:"unitPrice,omitempty"`
}

// ExplanationOfBenefitTotal represents totals for amounts submitted, co-pays, benefits payable etc
type ExplanationOfBenefitTotal struct {
	common.BackboneElement

	// Monetary amount associated with the category
	Amount common.Money `json:"amount"`

	// A code to indicate the information type of this adjudication record
	Category common.CodeableConcept `json:"category"`
}

// ExplanationOfBenefitPayment represents payment details for the adjudication of the claim
type ExplanationOfBenefitPayment struct {
	common.BackboneElement

	// Total amount of all adjustments to this payment
	Adjustment *common.Money `json:"adjustment,omitempty"`

	// Reason for the payment adjustment
	AdjustmentReason *common.CodeableConcept `json:"adjustmentReason,omitempty"`

	// Estimated date, based on the payment date
	Date *string `json:"date,omitempty"`

	// Benefits payable less any payment adjustment
	Amount common.Money `json:"amount"`

	// For example: EOB number, preauthorization number
	Identifier *common.Identifier `json:"identifier,omitempty"`

	// Whether this represents partial or complete payment of the claim
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// ExplanationOfBenefitProcessNote represents a note that describes or explains adjudication results
type ExplanationOfBenefitProcessNote struct {
	common.BackboneElement

	// The explanation or description associated with the processing
	Text *string `json:"text,omitempty"`

	// The business purpose of the note text
	Type *ExplanationOfBenefitProcessNoteType `json:"type,omitempty"`

	// A code to define the language used in the text of the note
	Language *common.CodeableConcept `json:"language,omitempty"`
}

// ExplanationOfBenefitProcessNoteType represents the business purpose of the note text
type ExplanationOfBenefitProcessNoteType string

const (
	ExplanationOfBenefitProcessNoteTypeDisplay   ExplanationOfBenefitProcessNoteType = "display"
	ExplanationOfBenefitProcessNoteTypePrint     ExplanationOfBenefitProcessNoteType = "print"
	ExplanationOfBenefitProcessNoteTypePrintoper ExplanationOfBenefitProcessNoteType = "printoper"
)

// ExplanationOfBenefitBenefitBalanceFinancial represents financial information for benefit balance
type ExplanationOfBenefitBenefitBalanceFinancial struct {
	common.BackboneElement

	// Benefits allowed
	AllowedMoney *common.Money `json:"allowedMoney,omitempty"`

	// Benefits allowed
	AllowedUnsignedInt *int `json:"allowedUnsignedInt,omitempty"`

	// Benefits allowed
	AllowedString *string `json:"allowedString,omitempty"`

	// Benefits used
	UsedMoney *common.Money `json:"usedMoney,omitempty"`

	// Benefits used
	UsedUnsignedInt *int `json:"usedUnsignedInt,omitempty"`

	// Type of benefit (primary care; speciality care; inpatient; outpatient)
	Type common.CodeableConcept `json:"type"`
}

// ExplanationOfBenefitBenefitBalance represents balance by Benefit Category
type ExplanationOfBenefitBenefitBalance struct {
	common.BackboneElement

	// Benefits Used to date
	Financial []ExplanationOfBenefitBenefitBalanceFinancial `json:"financial,omitempty"`

	// A code to indicate the Information Class of this adjudication record
	Category common.CodeableConcept `json:"category"`

	// True if the indicated class of service is excluded from the plan, missing or False indicates the product or service is included in the coverage
	Excluded *bool `json:"excluded,omitempty"`

	// A short name or tag for the benefit
	Name *string `json:"name,omitempty"`

	// A richer description of the benefit or services covered
	Description *string `json:"description,omitempty"`

	// Is a flag to indicate whether the benefits refer to in-network providers or out-of-network providers
	Network *common.CodeableConcept `json:"network,omitempty"`

	// Indicates if the benefits apply to the indicated network location
	Unit *common.CodeableConcept `json:"unit,omitempty"`

	// The term or period of the values such as 'maximum lifetime benefit' or 'maximum annual visits'
	Term *common.CodeableConcept `json:"term,omitempty"`
}
