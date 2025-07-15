// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// ExplanationOfBenefitStatus represents the status of an explanation of benefit
type ExplanationOfBenefitStatus string

const (
	ExplanationOfBenefitStatusActive         ExplanationOfBenefitStatus = "active"
	ExplanationOfBenefitStatusCancelled      ExplanationOfBenefitStatus = "cancelled"
	ExplanationOfBenefitStatusDraft          ExplanationOfBenefitStatus = "draft"
	ExplanationOfBenefitStatusEnteredInError ExplanationOfBenefitStatus = "entered-in-error"
)

// ExplanationOfBenefitOutcome represents the outcome of the claim processing
type ExplanationOfBenefitOutcome string

const (
	ExplanationOfBenefitOutcomeQueued   ExplanationOfBenefitOutcome = "queued"
	ExplanationOfBenefitOutcomeComplete ExplanationOfBenefitOutcome = "complete"
	ExplanationOfBenefitOutcomeError    ExplanationOfBenefitOutcome = "error"
	ExplanationOfBenefitOutcomePartial  ExplanationOfBenefitOutcome = "partial"
)

// ExplanationOfBenefitUse represents the use of the explanation of benefit
type ExplanationOfBenefitUse string

const (
	ExplanationOfBenefitUseClaim            ExplanationOfBenefitUse = "claim"
	ExplanationOfBenefitUsePreauthorization ExplanationOfBenefitUse = "preauthorization"
	ExplanationOfBenefitUsePredetermination ExplanationOfBenefitUse = "predetermination"
)

// ExplanationOfBenefitEvent represents events that occurred during the claim processing
type ExplanationOfBenefitEvent struct {
	common.BackboneElement

	// A coded event such as when a service is expected or a card printed
	Type common.CodeableConcept `json:"type"`

	// A date or period in the past or future indicating when the event occurred
	WhenDateTime *string        `json:"whenDateTime,omitempty"`
	WhenPeriod   *common.Period `json:"whenPeriod,omitempty"`
}

// ExplanationOfBenefitSupportingInfo represents supporting information for the claim
type ExplanationOfBenefitSupportingInfo struct {
	common.BackboneElement

	// This may contain a category for the local bill type codes
	Category common.CodeableConcept `json:"category"`

	// This may contain the local bill type codes
	Code *common.CodeableConcept `json:"code,omitempty"`

	// For example: the reason for the additional stay, or why a tooth is missing
	Reason *common.Coding `json:"reason,omitempty"`

	// A number to uniquely identify supporting information entries
	Sequence int `json:"sequence"`

	// The date when or period to which this information refers
	TimingDate   *string        `json:"timingDate,omitempty"`
	TimingPeriod *common.Period `json:"timingPeriod,omitempty"`

	// Could be used to provide references to other resources, document
	ValueBoolean    *bool              `json:"valueBoolean,omitempty"`
	ValueString     *string            `json:"valueString,omitempty"`
	ValueQuantity   *common.Quantity   `json:"valueQuantity,omitempty"`
	ValueAttachment *Attachment        `json:"valueAttachment,omitempty"`
	ValueReference  *common.Reference  `json:"valueReference,omitempty"`
	ValueIdentifier *common.Identifier `json:"valueIdentifier,omitempty"`
}

// ExplanationOfBenefitDiagnosis represents diagnoses relevant to the claim
type ExplanationOfBenefitDiagnosis struct {
	common.BackboneElement

	// The nature of illness or problem in a coded form or as a reference to an external defined Condition
	DiagnosisCodeableConcept *common.CodeableConcept `json:"diagnosisCodeableConcept,omitempty"`
	DiagnosisReference       *common.Reference       `json:"diagnosisReference,omitempty"`

	// Indication of whether the diagnosis was present on admission to a facility
	OnAdmission *common.CodeableConcept `json:"onAdmission,omitempty"`

	// Diagnosis are presented in list order to their expected importance
	Sequence int `json:"sequence"`

	// For example: admitting, primary, secondary, discharge
	Type []common.CodeableConcept `json:"type,omitempty"`
}

// ExplanationOfBenefitProcedure represents procedures performed on the patient
type ExplanationOfBenefitProcedure struct {
	common.BackboneElement

	// Date and optionally time the procedure was performed
	Date *string `json:"date,omitempty"`

	// Specific clinical procedure
	ProcedureCodeableConcept *common.CodeableConcept `json:"procedureCodeableConcept,omitempty"`
	ProcedureReference       *common.Reference       `json:"procedureReference,omitempty"`

	// A number to uniquely identify procedure entries
	Sequence int `json:"sequence"`

	// Category of Procedure
	Type []common.CodeableConcept `json:"type,omitempty"`

	// Unique device identifier
	Udi []common.Reference `json:"udi,omitempty"`
}

// ExplanationOfBenefitInsurance represents patient insurance information
type ExplanationOfBenefitInsurance struct {
	common.BackboneElement

	// Insurance information
	Coverage common.Reference `json:"coverage"`

	// Coverage to be used for adjudication
	Focal bool `json:"focal"`

	// Prior authorization reference number
	PreAuthRef []string `json:"preAuthRef,omitempty"`
}

// ExplanationOfBenefitAccident represents details of an accident
type ExplanationOfBenefitAccident struct {
	common.BackboneElement

	// When the incident occurred
	Date *string `json:"date,omitempty"`

	// Where the event occurred
	LocationAddress   *Address          `json:"locationAddress,omitempty"`
	LocationReference *common.Reference `json:"locationReference,omitempty"`

	// The nature of the accident
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// ExplanationOfBenefitCareTeam represents the members of the team who provided the products and services
type ExplanationOfBenefitCareTeam struct {
	common.BackboneElement

	// Practitioner or organization
	Provider common.Reference `json:"provider"`

	// Practitioner credential or specialization
	Qualification *common.CodeableConcept `json:"qualification,omitempty"`

	// Indicator of the lead practitioner
	Responsible *bool `json:"responsible,omitempty"`

	// Function within the team
	Role *common.CodeableConcept `json:"role,omitempty"`

	// Order of care team
	Sequence int `json:"sequence"`
}

// ExplanationOfBenefitPayee represents the recipient of the payment
type ExplanationOfBenefitPayee struct {
	common.BackboneElement

	// Recipient reference
	Party *common.Reference `json:"party,omitempty"`

	// Category of recipient
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// ExplanationOfBenefitPayment represents payment details for the adjudication
type ExplanationOfBenefitPayment struct {
	common.BackboneElement

	// Partial or complete payment adjustment
	Adjustment *Money `json:"adjustment,omitempty"`

	// Explanation for the adjustment
	AdjustmentReason *common.CodeableConcept `json:"adjustmentReason,omitempty"`

	// Payable amount after adjustment
	Amount *Money `json:"amount,omitempty"`

	// Expected date of payment
	Date *string `json:"date,omitempty"`

	// Business identifier for the payment
	Identifier *common.Identifier `json:"identifier,omitempty"`

	// Payment instrument
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// ExplanationOfBenefitProcessNote represents a note that describes or explains adjudication results
type ExplanationOfBenefitProcessNote struct {
	common.BackboneElement

	// Language of the text
	Language *common.CodeableConcept `json:"language,omitempty"`

	// A number to uniquely identify a note entry
	Number *int `json:"number,omitempty"`

	// Note explanatory text
	Text *string `json:"text,omitempty"`

	// The business purpose of the note text
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// ExplanationOfBenefitRelated represents other claims which are related to this claim
type ExplanationOfBenefitRelated struct {
	common.BackboneElement

	// File or case reference
	Claim *common.Reference `json:"claim,omitempty"`

	// Identifier of the related claim
	Reference *common.Identifier `json:"reference,omitempty"`

	// How the reference claim is related
	Relationship *common.CodeableConcept `json:"relationship,omitempty"`
}

// ExplanationOfBenefitTotal represents totals for amounts submitted, co-pays, benefits payable etc
type ExplanationOfBenefitTotal struct {
	common.BackboneElement

	// Financial total for the category
	Amount Money `json:"amount"`

	// Type of adjudication information
	Category common.CodeableConcept `json:"category"`
}

// ExplanationOfBenefitBenefitBalance represents balance by benefit category
type ExplanationOfBenefitBenefitBalance struct {
	common.BackboneElement

	// Examples include Medical Care, Periodontics, Renal Dialysis, Vision Coverage
	Category common.CodeableConcept `json:"category"`

	// For example, 'DENT2 covers 100% of basic, 50% of major but excludes Ortho, Implants and Cosmetic services'
	Description *string `json:"description,omitempty"`

	// True if the indicated class of service is excluded from the plan
	Excluded *bool `json:"excluded,omitempty"`

	// For example: MED01, or DENT2
	Name *string `json:"name,omitempty"`

	// Is a flag to indicate whether the benefits refer to in-network providers
	Network *common.CodeableConcept `json:"network,omitempty"`

	// The term or period of the values such as 'maximum lifetime benefit'
	Term *common.CodeableConcept `json:"term,omitempty"`

	// Indicates if the benefits apply to an individual or to the family
	Unit *common.CodeableConcept `json:"unit,omitempty"`
}

// ExplanationOfBenefitItem represents a claim line
type ExplanationOfBenefitItem struct {
	common.BackboneElement

	// Care team members related to this service or product
	CareTeamSequence []int `json:"careTeamSequence,omitempty"`

	// Examples include Medical Care, Periodontics, Renal Dialysis, Vision Coverage
	Category *common.CodeableConcept `json:"category,omitempty"`

	// Diagnoses applicable for this service or product
	DiagnosisSequence []int `json:"diagnosisSequence,omitempty"`

	// Healthcare encounters related to this claim
	Encounter []common.Reference `json:"encounter,omitempty"`

	// To show a 10% senior's discount, the value entered is: 0.90 (1.00 - 0.10)
	Factor *float64 `json:"factor,omitempty"`

	// Exceptions, special conditions and supporting information
	InformationSequence []int `json:"informationSequence,omitempty"`

	// Where the product or service was provided
	LocationCodeableConcept *common.CodeableConcept `json:"locationCodeableConcept,omitempty"`
	LocationAddress         *Address                `json:"locationAddress,omitempty"`
	LocationReference       *common.Reference       `json:"locationReference,omitempty"`

	// For example, in Oral whether the treatment is cosmetic or associated with TMJ
	Modifier []common.CodeableConcept `json:"modifier,omitempty"`

	// When the value is a group code then this item collects a set of related claim details
	ProductOrService *common.CodeableConcept `json:"productOrService,omitempty"`

	// Item typification or modifiers codes
	ProgramCode []common.CodeableConcept `json:"programCode,omitempty"`

	// Count of products or services
	Quantity *common.Quantity `json:"quantity,omitempty"`

	// Applicable exception and supporting information
	Revenue *common.CodeableConcept `json:"revenue,omitempty"`

	// A number to uniquely identify item entries
	Sequence int `json:"sequence"`

	// Date or dates of service or product delivery
	ServicedDate   *string        `json:"servicedDate,omitempty"`
	ServicedPeriod *common.Period `json:"servicedPeriod,omitempty"`

	// Anatomical sub-location
	SubSite []common.CodeableConcept `json:"subSite,omitempty"`

	// Unique device identifier
	Udi []common.Reference `json:"udi,omitempty"`

	// Fee, charge or cost per item
	UnitPrice *Money `json:"unitPrice,omitempty"`
}

// ExplanationOfBenefit represents the explanation of benefit resource
type ExplanationOfBenefit struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "ExplanationOfBenefit"

	// Details of a accident which resulted in injuries
	Accident *ExplanationOfBenefitAccident `json:"accident,omitempty"`

	// Balance by Benefit Category
	BenefitBalance []ExplanationOfBenefitBenefitBalance `json:"benefitBalance,omitempty"`

	// Not applicable when use=claim
	BenefitPeriod *common.Period `json:"benefitPeriod,omitempty"`

	// Typically this would be today or in the past for a claim
	BillablePeriod *common.Period `json:"billablePeriod,omitempty"`

	// The members of the team who provided the products and services
	CareTeam []ExplanationOfBenefitCareTeam `json:"careTeam,omitempty"`

	// The business identifier for the instance of the adjudication request
	Claim *common.Reference `json:"claim,omitempty"`

	// The business identifier for the instance of the adjudication response
	ClaimResponse *common.Reference `json:"claimResponse,omitempty"`

	// The date this resource was created
	Created string `json:"created"`

	// The element is used to indicate the current state of the adjudication
	Decision *common.CodeableConcept `json:"decision,omitempty"`

	// Information about diagnoses relevant to the claim items
	Diagnosis []ExplanationOfBenefitDiagnosis `json:"diagnosis,omitempty"`

	// For example DRG (Diagnosis Related Group) or a bundled billing code
	DiagnosisRelatedGroup *common.CodeableConcept `json:"diagnosisRelatedGroup,omitempty"`

	// A human readable description of the status of the adjudication
	Disposition *string `json:"disposition,omitempty"`

	// Healthcare encounters related to this claim
	Encounter []common.Reference `json:"encounter,omitempty"`

	// Individual who created the claim, predetermination or preauthorization
	Enterer *common.Reference `json:"enterer,omitempty"`

	// Information code for an event with a corresponding date or period
	Event []ExplanationOfBenefitEvent `json:"event,omitempty"`

	// Facility where the services were provided
	Facility *common.Reference `json:"facility,omitempty"`

	// Needed to permit insurers to include the actual form
	Form *Attachment `json:"form,omitempty"`

	// May be needed to identify specific jurisdictional forms
	FormCode *common.CodeableConcept `json:"formCode,omitempty"`

	// Fund would be release by a future claim
	FundsReserve *common.CodeableConcept `json:"fundsReserve,omitempty"`

	// This field is only used for preauthorizations
	FundsReserveRequested *common.CodeableConcept `json:"fundsReserveRequested,omitempty"`

	// A unique identifier assigned to this explanation of benefit
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// All insurance coverages for the patient
	Insurance []ExplanationOfBenefitInsurance `json:"insurance,omitempty"`

	// The party responsible for authorization, adjudication and reimbursement
	Insurer *common.Reference `json:"insurer,omitempty"`

	// A claim line
	Item []ExplanationOfBenefitItem `json:"item,omitempty"`

	// For example, a physician may prescribe a medication
	OriginalPrescription *common.Reference `json:"originalPrescription,omitempty"`

	// The resource may be used to indicate that the Claim has been received
	Outcome ExplanationOfBenefitOutcome `json:"outcome"`

	// The party to whom the professional services have been supplied
	Patient common.Reference `json:"patient"`

	// The amount paid by the patient
	PatientPaid *Money `json:"patientPaid,omitempty"`

	// Often providers agree to receive the benefits payable
	Payee *ExplanationOfBenefitPayee `json:"payee,omitempty"`

	// Payment details for the adjudication of the claim
	Payment *ExplanationOfBenefitPayment `json:"payment,omitempty"`

	// This value is only present on preauthorization adjudications
	PreAuthRef []string `json:"preAuthRef,omitempty"`

	// This value is only present on preauthorization adjudications
	PreAuthRefPeriod []common.Period `json:"preAuthRefPeriod,omitempty"`

	// This indicates the relative order of a series of EOBs
	Precedence *int `json:"precedence,omitempty"`

	// Prescription is the document/authorization given to the claim author
	Prescription *common.Reference `json:"prescription,omitempty"`

	// If a claim processor is unable to complete the processing
	Priority *common.CodeableConcept `json:"priority,omitempty"`

	// Procedures performed on the patient relevant to the billing items
	Procedure []ExplanationOfBenefitProcedure `json:"procedure,omitempty"`

	// A note that describes or explains adjudication results
	ProcessNote []ExplanationOfBenefitProcessNote `json:"processNote,omitempty"`

	// Typically this field would be 1..1 where this party is accountable
	Provider *common.Reference `json:"provider,omitempty"`

	// The referral resource which lists the date, practitioner, reason
	Referral *common.Reference `json:"referral,omitempty"`

	// For example, for the original treatment and follow-up exams
	Related []ExplanationOfBenefitRelated `json:"related,omitempty"`

	// This element is labeled as a modifier because the status contains codes
	Status ExplanationOfBenefitStatus `json:"status"`

	// This may contain the local bill type codes such as the US UB-04 bill type code
	SubType *common.CodeableConcept `json:"subType,omitempty"`

	// Often there are multiple jurisdiction specific valuesets which are required
	SupportingInfo []ExplanationOfBenefitSupportingInfo `json:"supportingInfo,omitempty"`

	// Totals for amounts submitted, co-pays, benefits payable etc
	Total []ExplanationOfBenefitTotal `json:"total,omitempty"`

	// Trace number for tracking purposes
	TraceNumber []common.Identifier `json:"traceNumber,omitempty"`

	// The majority of jurisdictions use: oral, pharmacy, vision, professional
	Type common.CodeableConcept `json:"type"`

	// A code to indicate whether the nature of the request is
	Use ExplanationOfBenefitUse `json:"use"`
}
