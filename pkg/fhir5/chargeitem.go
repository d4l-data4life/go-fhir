// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// ChargeItemStatus represents the status of a charge item
type ChargeItemStatus string

const (
	ChargeItemStatusPlanned        ChargeItemStatus = "planned"
	ChargeItemStatusBillable       ChargeItemStatus = "billable"
	ChargeItemStatusNotBillable    ChargeItemStatus = "not-billable"
	ChargeItemStatusAborted        ChargeItemStatus = "aborted"
	ChargeItemStatusBilled         ChargeItemStatus = "billed"
	ChargeItemStatusEnteredInError ChargeItemStatus = "entered-in-error"
	ChargeItemStatusUnknown        ChargeItemStatus = "unknown"
)

// MonetaryComponentType represents the type of a monetary component
type MonetaryComponentType string

const (
	MonetaryComponentTypeBase          MonetaryComponentType = "base"
	MonetaryComponentTypeSurcharge     MonetaryComponentType = "surcharge"
	MonetaryComponentTypeDeduction     MonetaryComponentType = "deduction"
	MonetaryComponentTypeDiscount      MonetaryComponentType = "discount"
	MonetaryComponentTypeTax           MonetaryComponentType = "tax"
	MonetaryComponentTypeInformational MonetaryComponentType = "informational"
)

// MonetaryComponent represents a monetary value component
type MonetaryComponent struct {
	DataType

	// Explicit value amount to be used
	Amount *Money `json:"amount,omitempty"`

	// Codes may be used to differentiate between kinds of taxes, surcharges, discounts etc.
	Code *common.CodeableConcept `json:"code,omitempty"`

	// Factor used for calculating this component
	Factor *float64 `json:"factor,omitempty"`

	// base | surcharge | deduction | discount | tax | informational
	Type MonetaryComponentType `json:"type"`
}

// ChargeItemPerformer represents who performed or participated in the charged service
type ChargeItemPerformer struct {
	common.BackboneElement

	// The device, practitioner, etc. who performed or participated in the service
	Actor common.Reference `json:"actor"`

	// Describes the type of performance or participation(e.g. primary surgeon, anesthesiologist, etc.)
	Function *common.CodeableConcept `json:"function,omitempty"`
}

// ChargeItem represents the provision of healthcare provider products for a certain patient
type ChargeItem struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "ChargeItem"

	// Systems posting the ChargeItems might not always be able to determine, which accounts the Items need to be placed into
	Account []common.Reference `json:"account,omitempty"`

	// Only used if not implicit in code found in Condition.code
	Bodysite []common.CodeableConcept `json:"bodysite,omitempty"`

	// A code that identifies the charge, like a billing code
	Code common.CodeableConcept `json:"code"`

	// The costCenter could either be given as a reference to an Organization(Role) resource
	CostCenter *common.Reference `json:"costCenter,omitempty"`

	// References the source of pricing information, rules of application for the code this ChargeItem uses
	DefinitionCanonical []string `json:"definitionCanonical,omitempty"`

	// References the (external) source of pricing information, rules of application for the code this ChargeItem uses
	DefinitionUri []string `json:"definitionUri,omitempty"`

	// This ChargeItem may be recorded during planning, execution or after the actual encounter takes place
	Encounter *common.Reference `json:"encounter,omitempty"`

	// The actual date when the service associated with the charge has been rendered is captured in occurrence[x]
	EnteredDate *string `json:"enteredDate,omitempty"`

	// The enterer is also the person considered responsible for factor/price overrides if applicable
	Enterer *common.Reference `json:"enterer,omitempty"`

	// Identifiers assigned to this event performer or other systems
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Comments made about the event by the performer, subject or other participants
	Note []Annotation `json:"note,omitempty"`

	// The list of types may be constrained as appropriate for the type of charge item
	OccurrenceDateTime *string        `json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod   *common.Period `json:"occurrencePeriod,omitempty"`
	OccurrenceTiming   *Timing        `json:"occurrenceTiming,omitempty"`

	// If the list price or the rule-based factor associated with the code is overridden
	OverrideReason *common.CodeableConcept `json:"overrideReason,omitempty"`

	// ChargeItems can be grouped to larger ChargeItems covering the whole set
	PartOf []common.Reference `json:"partOf,omitempty"`

	// Indicates who or what performed or participated in the charged service
	Performer []ChargeItemPerformer `json:"performer,omitempty"`

	// Practitioners and Devices can be associated with multiple organizations
	PerformingOrganization *common.Reference `json:"performingOrganization,omitempty"`

	// Identifies the device, food, drug or other product being charged either by type code or reference to an instance
	Product []CodeableReference `json:"product,omitempty"`

	// In many cases this may just be a value, if the underlying units are implicit in the definition of the charge item code
	Quantity *common.Quantity `json:"quantity,omitempty"`

	// If the application of the charge item requires a reason to be given, it can be captured here
	Reason []common.CodeableConcept `json:"reason,omitempty"`

	// The rendered Service might not be associated with a Request
	RequestingOrganization *common.Reference `json:"requestingOrganization,omitempty"`

	// Indicated the rendered service that caused this charge
	Service []CodeableReference `json:"service,omitempty"`

	// Unknown does not represent "other" - one of the defined statuses must apply
	Status ChargeItemStatus `json:"status"`

	// The individual or set of individuals the action is being or was performed on
	Subject common.Reference `json:"subject"`

	// Further information supporting this charge
	SupportingInformation []common.Reference `json:"supportingInformation,omitempty"`

	// Often, the total price may be calculated and recorded on the Invoice
	TotalPriceComponent *MonetaryComponent `json:"totalPriceComponent,omitempty"`

	// This could be communicated in ChargeItemDefinition
	UnitPriceComponent *MonetaryComponent `json:"unitPriceComponent,omitempty"`
}
