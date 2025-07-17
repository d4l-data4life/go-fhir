package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// ValueSetComposeIncludeConceptDesignation represents concepts have both a display and an array of designation
type ValueSetComposeIncludeConceptDesignation struct {
	common.BackboneElement

	// This was added rather than increasing the cardinality of .use to 0..* in order to maintain backward compatibility
	AdditionalUse []common.Coding `json:"additionalUse,omitempty"`

	// In the absence of a language, the resource language applies
	Language *string `json:"language,omitempty"`

	// If no use is provided, the designation can be assumed to be suitable for general display to a human user
	Use *common.Coding `json:"use,omitempty"`

	// The text value for this designation
	Value string `json:"value"`
}

// ValueSetComposeIncludeConcept represents the list of concepts is considered ordered
type ValueSetComposeIncludeConcept struct {
	common.BackboneElement

	// Expressions are allowed if defined by the underlying code system
	Code string `json:"code"`

	// Concepts have both a display and an array of designation
	Designation []ValueSetComposeIncludeConceptDesignation `json:"designation,omitempty"`

	// The value set resource allows for an alternative display to be specified for when this concept is used in this particular value set
	Display *string `json:"display,omitempty"`
}

// ValueSetComposeIncludeFilter represents selecting codes by specifying filters based on properties
type ValueSetComposeIncludeFilter struct {
	common.BackboneElement

	// In case filter.property represents a property of the system, the operation applies to the selected property
	Op ValueSetComposeIncludeFilterOp `json:"op"`

	// A code that identifies a property or a filter defined in the code system
	Property string `json:"property"`

	// The match value may be either a code defined by the system, or a string value, which is a regex match on the literal string of the property value
	Value string `json:"value"`
}

// ValueSetComposeIncludeFilterOp represents the operation to perform
type ValueSetComposeIncludeFilterOp string

const (
	ValueSetComposeIncludeFilterOpEquals         ValueSetComposeIncludeFilterOp = "="
	ValueSetComposeIncludeFilterOpIsA            ValueSetComposeIncludeFilterOp = "is-a"
	ValueSetComposeIncludeFilterOpDescendentOf   ValueSetComposeIncludeFilterOp = "descendent-of"
	ValueSetComposeIncludeFilterOpIsNotA         ValueSetComposeIncludeFilterOp = "is-not-a"
	ValueSetComposeIncludeFilterOpRegex          ValueSetComposeIncludeFilterOp = "regex"
	ValueSetComposeIncludeFilterOpIn             ValueSetComposeIncludeFilterOp = "in"
	ValueSetComposeIncludeFilterOpNotIn          ValueSetComposeIncludeFilterOp = "not-in"
	ValueSetComposeIncludeFilterOpGeneralizes    ValueSetComposeIncludeFilterOp = "generalizes"
	ValueSetComposeIncludeFilterOpChildOf        ValueSetComposeIncludeFilterOp = "child-of"
	ValueSetComposeIncludeFilterOpDescendentLeaf ValueSetComposeIncludeFilterOp = "descendent-leaf"
	ValueSetComposeIncludeFilterOpExists         ValueSetComposeIncludeFilterOp = "exists"
)

// ValueSetComposeInclude represents include one or more codes from a code system or other value set(s)
type ValueSetComposeInclude struct {
	common.BackboneElement

	// The list of concepts is considered ordered, though the order might not have any particular significance
	Concept []ValueSetComposeIncludeConcept `json:"concept,omitempty"`

	// Selecting codes by specifying filters based on properties is only possible where the underlying code system defines appropriate properties
	Filter []ValueSetComposeIncludeFilter `json:"filter,omitempty"`

	// The version of the code system that the codes are selected from
	Version *string `json:"version,omitempty"`

	// The system value is required if the concept is not defined by the value set
	System *string `json:"system,omitempty"`

	// The value set resource allows for an alternative display to be specified for when this concept is used in this particular value set
	ValueSet []string `json:"valueSet,omitempty"`
}

// ValueSetCompose represents a set of criteria that define the contents of the value set by including or excluding codes
type ValueSetCompose struct {
	common.BackboneElement

	// The Locked Date is the effective date that is used to determine the version of all referenced Code Systems and Value Set Definitions included in the compose that are not already tied to a specific version
	LockedDate *string `json:"lockedDate,omitempty"`

	// Whether inactive codes - codes that are not approved for current use - are in the value set
	Inactive *bool `json:"inactive,omitempty"`

	// Include one or more codes from a code system or other value set(s)
	Include []ValueSetComposeInclude `json:"include"`

	// Exclude one or more codes from the value set based on code system filters and/or other value sets
	Exclude []ValueSetComposeInclude `json:"exclude,omitempty"`

	// Property to filter on
	Property []string `json:"property,omitempty"`
}

// ValueSetExpansionParameter represents expansion parameter
type ValueSetExpansionParameter struct {
	common.BackboneElement

	// The name of the parameter
	Name string `json:"name"`

	// The value of the parameter
	ValueString   *string  `json:"valueString,omitempty"`
	ValueBoolean  *bool    `json:"valueBoolean,omitempty"`
	ValueInteger  *int     `json:"valueInteger,omitempty"`
	ValueDecimal  *float64 `json:"valueDecimal,omitempty"`
	ValueUri      *string  `json:"valueUri,omitempty"`
	ValueCode     *string  `json:"valueCode,omitempty"`
	ValueDateTime *string  `json:"valueDateTime,omitempty"`
}

// ValueSetExpansionContains represents codes in the value set
type ValueSetExpansionContains struct {
	common.BackboneElement

	// An absolute URI that is used to identify this code in this value set
	System *string `json:"system,omitempty"`

	// If true, this entry is included in the expansion for navigational purposes, and the user cannot select the code directly as a proper value
	Abstract *bool `json:"abstract,omitempty"`

	// If the concept is inactive in the code system that defines it
	Inactive *bool `json:"inactive,omitempty"`

	// The version of the code system from this code was taken
	Version *string `json:"version,omitempty"`

	// The code for this item in the expansion hierarchy
	Code *string `json:"code,omitempty"`

	// The recommended display for this item in the expansion
	Display *string `json:"display,omitempty"`

	// Additional representations for this item
	Designation []ValueSetComposeIncludeConceptDesignation `json:"designation,omitempty"`

	// Other codes and entries contained under this entry in the hierarchy
	Contains []ValueSetExpansionContains `json:"contains,omitempty"`

	// The property value for this concept
	Property []ValueSetExpansionContainsProperty `json:"property,omitempty"`
}

// ValueSetExpansionContainsProperty represents the property value for this concept
type ValueSetExpansionContainsProperty struct {
	common.BackboneElement

	// A code that is a reference to ValueSet.expansion.property.code
	Code string `json:"code"`

	// The value of this property
	ValueCode     *string        `json:"valueCode,omitempty"`
	ValueCoding   *common.Coding `json:"valueCoding,omitempty"`
	ValueString   *string        `json:"valueString,omitempty"`
	ValueInteger  *int           `json:"valueInteger,omitempty"`
	ValueBoolean  *bool          `json:"valueBoolean,omitempty"`
	ValueDateTime *string        `json:"valueDateTime,omitempty"`
	ValueDecimal  *float64       `json:"valueDecimal,omitempty"`
}

// ValueSetExpansion represents expansion is performed to produce a collection of codes
type ValueSetExpansion struct {
	common.BackboneElement

	// An identifier that uniquely identifies this expansion of the valueset
	Identifier *string `json:"identifier,omitempty"`

	// The time at which the expansion was produced by the expanding system
	Timestamp string `json:"timestamp"`

	// The total number of concepts in the expansion
	Total *int `json:"total,omitempty"`

	// If paging is being used, the offset at which this resource starts
	Offset *int `json:"offset,omitempty"`

	// A parameter that controlled the expansion process
	Parameter []ValueSetExpansionParameter `json:"parameter,omitempty"`

	// The codes that are contained in the value set expansion
	Contains []ValueSetExpansionContains `json:"contains,omitempty"`
}

// ValueSetScope represents description of the semantic space the Value Set Expansion is intended to cover
type ValueSetScope struct {
	common.BackboneElement

	// Criteria for inclusion of concepts in the value set expansion
	InclusionCriteria *string `json:"inclusionCriteria,omitempty"`

	// Criteria for exclusion of concepts from the value set expansion
	ExclusionCriteria *string `json:"exclusionCriteria,omitempty"`
}

// ValueSetStatus represents the status of a value set
type ValueSetStatus string

const (
	ValueSetStatusDraft   ValueSetStatus = "draft"
	ValueSetStatusActive  ValueSetStatus = "active"
	ValueSetStatusRetired ValueSetStatus = "retired"
	ValueSetStatusUnknown ValueSetStatus = "unknown"
)

// ValueSet represents a ValueSet resource instance specifies a set of codes drawn from one or more code systems
type ValueSet struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "ValueSet"

	// The 'date' element may be more recent than the approval date because of minor changes or editorial corrections
	ApprovalDate *string `json:"approvalDate,omitempty"`

	// An individiual or organization primarily involved in the creation and maintenance of the ValueSet
	Author []common.ContactDetail `json:"author,omitempty"`

	// A set of criteria that define the contents of the value set by including or excluding codes selected from the specified code system(s) that the value set draws from
	Compose *ValueSetCompose `json:"compose,omitempty"`

	// May be a web site, an email address, a telephone number, etc
	Contact []common.ContactDetail `json:"contact,omitempty"`

	// Frequently, the copyright differs between the value set and the codes that are included
	Copyright *string `json:"copyright,omitempty"`

	// The (c) symbol should NOT be included in this string
	CopyrightLabel *string `json:"copyrightLabel,omitempty"`

	// Note that this is not the same as the meta.lastUpdated which is specific to an instance of a value set resource on a server
	Date *string `json:"date,omitempty"`

	// Description SHOULD contain instructions for clinical or administrative use and interpretation and information about misuse
	Description *string `json:"description,omitempty"`

	// An individual or organization primarily responsible for internal coherence of the ValueSet
	Editor []common.ContactDetail `json:"editor,omitempty"`

	// The effective period for a ValueSet determines when the content is applicable for usage and is independent of publication and review dates
	EffectivePeriod *common.Period `json:"effectivePeriod,omitempty"`

	// See guidance around (not) making local changes to elements
	Endorser []common.ContactDetail `json:"endorser,omitempty"`

	// Expansion is performed to produce a collection of codes that are ready to use for data entry or validation
	Expansion *ValueSetExpansion `json:"expansion,omitempty"`

	// Allows filtering of value sets that are appropriate for use versus not
	Experimental *bool `json:"experimental,omitempty"`

	// Typically, this is used for identifiers that can go in an HL7 V3 II (instance identifier) data type, and can then identify this value set outside of FHIR, where it is not possible to use the logical URI
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Normally immutability is set to 'false', which is the default assumption if it is not populated
	Immutable *bool `json:"immutable,omitempty"`

	// It may be possible for the value set to be used in jurisdictions other than those for which it was originally designed or intended
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`

	// If specified, this date follows the original approval date
	LastReviewDate *string `json:"lastReviewDate,omitempty"`

	// The name is not expected to be globally unique
	Name *string `json:"name,omitempty"`

	// Usually an organization but may be an individual
	Publisher *string `json:"publisher,omitempty"`

	// This element does not describe the usage of the value set
	Purpose *string `json:"purpose,omitempty"`

	// Each related artifact is either an attachment, or a reference to another resource, but not both
	RelatedArtifact []common.RelatedArtifact `json:"relatedArtifact,omitempty"`

	// See guidance around (not) making local changes to elements
	Reviewer []common.ContactDetail `json:"reviewer,omitempty"`

	// Description of the semantic space the Value Set Expansion is intended to cover and should further clarify the text in ValueSet.description
	Scope *ValueSetScope `json:"scope,omitempty"`

	// Allows filtering of value sets that are appropriate for use versus not
	Status ValueSetStatus `json:"status"`

	// This name does not need to be machine-processing friendly and may contain punctuation, white-space, etc
	Title *string `json:"title,omitempty"`

	// DEPRECATION NOTE: For consistency, implementations are encouraged to migrate to using the new 'topic' code in the useContext element
	Topic []common.CodeableConcept `json:"topic,omitempty"`

	// Can be a urn:uuid: or a urn:oid: but real http: addresses are preferred
	Url *string `json:"url,omitempty"`

	// When multiple useContexts are specified, there is no expectation that all or any of the contexts apply
	UseContext []common.UsageContext `json:"useContext,omitempty"`

	// There may be different value set instances that have the same identifier but different versions
	Version *string `json:"version,omitempty"`

	// If set as a string, this is a FHIRPath expression that has two additional context variables passed in - %version1 and %version2 and will return a negative number if version1 is newer, a positive number if version2 and a 0 if the version ordering can't be successfully be determined
	VersionAlgorithmString *string `json:"versionAlgorithmString,omitempty"`

	// If set as a string, this is a FHIRPath expression that has two additional context variables passed in - %version1 and %version2 and will return a negative number if version1 is newer, a positive number if version2 and a 0 if the version ordering can't be successfully be determined
	VersionAlgorithmCoding *common.Coding `json:"versionAlgorithmCoding,omitempty"`
}
