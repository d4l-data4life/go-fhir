// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// ChargeItemDefinitionApplicability represents applicability conditions for a charge item definition
type ChargeItemDefinitionApplicability struct {
	common.BackboneElement

	// Please note that FHIRPath Expressions can only be evaluated in the scope of the current ChargeItem resource
	Condition *Expression `json:"condition,omitempty"`

	// The effective period for a charge item definition determines when the content is applicable for usage
	EffectivePeriod *common.Period `json:"effectivePeriod,omitempty"`

	// Reference to / quotation of the external source of the group of properties
	RelatedArtifact *RelatedArtifact `json:"relatedArtifact,omitempty"`
}

// ChargeItemDefinitionPropertyGroup represents a group of properties which are applicable under the same conditions
type ChargeItemDefinitionPropertyGroup struct {
	common.BackboneElement

	// The applicability conditions can be used to ascertain whether a billing item is allowed in a specific context
	Applicability []ChargeItemDefinitionApplicability `json:"applicability,omitempty"`

	// The price for a ChargeItem may be calculated as a base price with surcharges/deductions
	PriceComponent []MonetaryComponent `json:"priceComponent,omitempty"`
}

// ChargeItemDefinition represents a charge item definition
type ChargeItemDefinition struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "ChargeItemDefinition"

	// The applicability conditions can be used to ascertain whether a billing item is allowed in a specific context
	Applicability []ChargeItemDefinitionApplicability `json:"applicability,omitempty"`

	// The 'date' element may be more recent than the approval date
	ApprovalDate *string `json:"approvalDate,omitempty"`

	// The defined billing details in this resource pertain to the given billing code
	Code *common.CodeableConcept `json:"code,omitempty"`

	// May be a web site, an email address, a telephone number, etc.
	Contact []ContactDetail `json:"contact,omitempty"`

	// A copyright statement relating to the charge item definition
	Copyright *string `json:"copyright,omitempty"`

	// The (c) symbol should NOT be included in this string
	CopyrightLabel *string `json:"copyrightLabel,omitempty"`

	// The date is often not tracked until the resource is published
	Date *string `json:"date,omitempty"`

	// The URL pointing to an externally-defined charge item definition
	DerivedFromUri []string `json:"derivedFromUri,omitempty"`

	// This description can be used to capture details such as comments about misuse
	Description *string `json:"description,omitempty"`

	// Allows filtering of charge item definitions that are appropriate for use versus not
	Experimental *bool `json:"experimental,omitempty"`

	// Typically, this is used for identifiers that can go in an HL7 V3 II
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// In case of highly customized, individually produced or fitted devices/substances
	Instance []common.Reference `json:"instance,omitempty"`

	// It may be possible for the charge item definition to be used in jurisdictions other than those for which it was originally designed
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`

	// If specified, this date follows the original approval date
	LastReviewDate *string `json:"lastReviewDate,omitempty"`

	// The name is not expected to be globally unique
	Name *string `json:"name,omitempty"`

	// A larger definition of which this particular definition is a component or step
	PartOf []string `json:"partOf,omitempty"`

	// Group of properties which are applicable under the same conditions
	PropertyGroup []ChargeItemDefinitionPropertyGroup `json:"propertyGroup,omitempty"`

	// Usually an organization but may be an individual
	Publisher *string `json:"publisher,omitempty"`

	// This element does not describe the usage of the charge item definition
	Purpose *string `json:"purpose,omitempty"`

	// As new versions of a protocol or guideline are defined, allows identification of what versions are replaced
	Replaces []string `json:"replaces,omitempty"`

	// Allows filtering of charge item definitions that are appropriate for use versus not
	Status ChargeItemDefinitionStatus `json:"status"`

	// This name does not need to be machine-processing friendly
	Title *string `json:"title,omitempty"`

	// Can be a urn:uuid: or a urn:oid: but real http: addresses are preferred
	Url *string `json:"url,omitempty"`

	// When multiple useContexts are specified, there is no expectation that all or any of the contexts apply
	UseContext []UsageContext `json:"useContext,omitempty"`

	// There may be different charge item definition instances that have the same identifier but different versions
	Version *string `json:"version,omitempty"`

	// If set as a string, this is a FHIRPath expression
	VersionAlgorithmString *string `json:"versionAlgorithmString,omitempty"`

	// If set as a string, this is a FHIRPath expression
	VersionAlgorithmCoding *common.Coding `json:"versionAlgorithmCoding,omitempty"`
}

// ChargeItemDefinitionStatus represents the status of a charge item definition
type ChargeItemDefinitionStatus string

const (
	ChargeItemDefinitionStatusDraft   ChargeItemDefinitionStatus = "draft"
	ChargeItemDefinitionStatusActive  ChargeItemDefinitionStatus = "active"
	ChargeItemDefinitionStatusRetired ChargeItemDefinitionStatus = "retired"
	ChargeItemDefinitionStatusUnknown ChargeItemDefinitionStatus = "unknown"
)
