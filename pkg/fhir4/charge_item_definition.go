package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// ChargeItemDefinition represents the ChargeItemDefinition resource
type ChargeItemDefinition struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "ChargeItemDefinition"

	// The applicability conditions can be used to ascertain whether a billing item is allowed in a specific context
	Applicability []ChargeItemDefinitionApplicability `json:"applicability,omitempty"`

	// The 'date' element may be more recent than the approval date because of minor changes or editorial corrections
	ApprovalDate *string `json:"approvalDate,omitempty"`

	// The defined billing details in this resource pertain to the given billing code
	Code *common.CodeableConcept `json:"code,omitempty"`

	// May be a web site, an email address, a telephone number, etc.
	Contact []common.ContactDetail `json:"contact,omitempty"`

	// A copyright statement relating to the charge item definition and/or its contents
	Copyright *string `json:"copyright,omitempty"`

	// Note that this is not the same as the resource last-modified-date
	Date *string `json:"date,omitempty"`

	// The URL pointing to an externally-defined charge item definition
	DerivedFromUri []string `json:"derivedFromUri,omitempty"`

	// This description can be used to capture details such as why the charge item definition was built
	Description *string `json:"description,omitempty"`

	// The effective period for a charge item definition determines when the content is applicable for usage
	EffectivePeriod *common.Period `json:"effectivePeriod,omitempty"`

	// Allows filtering of charge item definitions that are appropriate for use versus not
	Experimental *bool `json:"experimental,omitempty"`

	// Typically, this is used for identifiers that can go in an HL7 V3 II data type
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// In case of highly customized, individually produced or fitted devices/substances
	Instance []common.Reference `json:"instance,omitempty"`

	// It may be possible for the charge item definition to be used in jurisdictions other than those for which it was originally designed
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`

	// If specified, this date follows the original approval date
	LastReviewDate *string `json:"lastReviewDate,omitempty"`

	// A larger definition of which this particular definition is a component or step
	PartOf []string `json:"partOf,omitempty"`

	// Group of properties which are applicable under the same conditions
	PropertyGroup []ChargeItemDefinitionPropertyGroup `json:"propertyGroup,omitempty"`

	// Usually an organization but may be an individual
	Publisher *string `json:"publisher,omitempty"`

	// As new versions of a protocol or guideline are defined, allows identification of what versions are replaced
	Replaces []string `json:"replaces,omitempty"`

	// Allows filtering of charge item definitions that are appropriate for use versus not
	Status ChargeItemDefinitionStatus `json:"status"`

	// This name does not need to be machine-processing friendly and may contain punctuation, white-space, etc.
	Title *string `json:"title,omitempty"`

	// Can be a urn:uuid: or a urn:oid: but real http: addresses are preferred
	Url string `json:"url"`

	// When multiple useContexts are specified, there is no expectation that all or any of the contexts apply
	UseContext []common.UsageContext `json:"useContext,omitempty"`

	// There may be different charge item definition instances that have the same identifier but different versions
	Version *string `json:"version,omitempty"`
}

// ChargeItemDefinitionStatus represents the status of the charge item definition
type ChargeItemDefinitionStatus string

const (
	ChargeItemDefinitionStatusDraft   ChargeItemDefinitionStatus = "draft"
	ChargeItemDefinitionStatusActive  ChargeItemDefinitionStatus = "active"
	ChargeItemDefinitionStatusRetired ChargeItemDefinitionStatus = "retired"
	ChargeItemDefinitionStatusUnknown ChargeItemDefinitionStatus = "unknown"
)

// ChargeItemDefinitionApplicability represents the applicability conditions
type ChargeItemDefinitionApplicability struct {
	common.BackboneElement

	// A brief, natural language description of the condition that effectively communicates the intended semantics
	Description *string `json:"description,omitempty"`

	// Please note that FHIRPath Expressions can only be evaluated in the scope of the current ChargeItem resource
	Expression *string `json:"expression,omitempty"`

	// The media type of the language for the expression, e.g. "text/cql" for Clinical Query Language expressions
	Language *string `json:"language,omitempty"`
}

// ChargeItemDefinitionPropertyGroupPriceComponent represents the price component
type ChargeItemDefinitionPropertyGroupPriceComponent struct {
	common.BackboneElement

	// The amount calculated for this component
	Amount *common.Money `json:"amount,omitempty"`

	// A code that identifies the component
	Code *common.CodeableConcept `json:"code,omitempty"`

	// The factor that has been applied on the base price for calculating this component
	Factor *float64 `json:"factor,omitempty"`

	// This code identifies the type of the component
	Type ChargeItemDefinitionPropertyGroupPriceComponentType `json:"type"`
}

// ChargeItemDefinitionPropertyGroupPriceComponentType represents the type of price component
type ChargeItemDefinitionPropertyGroupPriceComponentType string

const (
	ChargeItemDefinitionPropertyGroupPriceComponentTypeBase          ChargeItemDefinitionPropertyGroupPriceComponentType = "base"
	ChargeItemDefinitionPropertyGroupPriceComponentTypeSurcharge     ChargeItemDefinitionPropertyGroupPriceComponentType = "surcharge"
	ChargeItemDefinitionPropertyGroupPriceComponentTypeDeduction     ChargeItemDefinitionPropertyGroupPriceComponentType = "deduction"
	ChargeItemDefinitionPropertyGroupPriceComponentTypeDiscount      ChargeItemDefinitionPropertyGroupPriceComponentType = "discount"
	ChargeItemDefinitionPropertyGroupPriceComponentTypeTax           ChargeItemDefinitionPropertyGroupPriceComponentType = "tax"
	ChargeItemDefinitionPropertyGroupPriceComponentTypeInformational ChargeItemDefinitionPropertyGroupPriceComponentType = "informational"
)

// ChargeItemDefinitionPropertyGroup represents a group of properties
type ChargeItemDefinitionPropertyGroup struct {
	common.BackboneElement

	// The applicability conditions can be used to ascertain whether a billing item is allowed in a specific context
	Applicability []ChargeItemDefinitionApplicability `json:"applicability,omitempty"`

	// The price for a ChargeItem may be calculated as a base price with surcharges/deductions
	PriceComponent []ChargeItemDefinitionPropertyGroupPriceComponent `json:"priceComponent,omitempty"`
}
