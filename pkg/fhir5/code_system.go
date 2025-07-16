// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// CodeSystemFilter represents a filter that can be used to select a subset of concepts
type CodeSystemFilter struct {
	common.BackboneElement

	// The code that identifies this filter when it is used as a filter in ValueSet
	Code string `json:"code"`

	// A description of how or why the filter is used
	Description *string `json:"description,omitempty"`

	// A list of operators that can be used with the filter
	Operator []string `json:"operator"`

	// A description of what the value for the filter should be
	Value string `json:"value"`
}

// CodeSystemProperty represents a property that can be used to extend the code system
type CodeSystemProperty struct {
	common.BackboneElement

	// A code that is used to identify the property
	Code string `json:"code"`

	// A description of the property- why it is defined, and how its value might be used
	Description *string `json:"description,omitempty"`

	// The type of the property value
	Type CodeSystemPropertyType `json:"type"`

	// Reference to the formal meaning of the property
	URI *string `json:"uri,omitempty"`
}

// CodeSystemPropertyType represents the type of a code system property
type CodeSystemPropertyType string

const (
	CodeSystemPropertyTypeCode     CodeSystemPropertyType = "code"
	CodeSystemPropertyTypeCoding   CodeSystemPropertyType = "Coding"
	CodeSystemPropertyTypeString   CodeSystemPropertyType = "string"
	CodeSystemPropertyTypeInteger  CodeSystemPropertyType = "integer"
	CodeSystemPropertyTypeBoolean  CodeSystemPropertyType = "boolean"
	CodeSystemPropertyTypeDateTime CodeSystemPropertyType = "dateTime"
	CodeSystemPropertyTypeDecimal  CodeSystemPropertyType = "decimal"
)

// CodeSystemConceptDesignation represents a designation for a concept
type CodeSystemConceptDesignation struct {
	common.BackboneElement

	// This was added rather than increasing the cardinality of .use to 0..*
	AdditionalUse []common.Coding `json:"additionalUse,omitempty"`

	// In the absence of a language, the resource language applies
	Language *string `json:"language,omitempty"`

	// If no use is provided, the designation can be assumed to be suitable for general display
	Use *common.Coding `json:"use,omitempty"`

	// The text value for this designation
	Value string `json:"value"`
}

// CodeSystemConceptProperty represents a property value for a concept
type CodeSystemConceptProperty struct {
	common.BackboneElement

	// A code that is a reference to CodeSystem.property.code
	Code string `json:"code"`

	// The value of this property
	ValueCode *string `json:"valueCode,omitempty"`

	// The value of this property
	ValueCoding *common.Coding `json:"valueCoding,omitempty"`

	// The value of this property
	ValueString *string `json:"valueString,omitempty"`

	// The value of this property
	ValueInteger *int `json:"valueInteger,omitempty"`

	// The value of this property
	ValueBoolean *bool `json:"valueBoolean,omitempty"`

	// The value of this property
	ValueDateTime *string `json:"valueDateTime,omitempty"`

	// The value of this property
	ValueDecimal *float64 `json:"valueDecimal,omitempty"`
}

// CodeSystemConcept represents a concept in a code system
type CodeSystemConcept struct {
	common.BackboneElement

	// A code - a text symbol - that uniquely identifies the concept within the code system
	Code string `json:"code"`

	// Defines children of a concept to produce a hierarchy of concepts
	Concept []CodeSystemConcept `json:"concept,omitempty"`

	// The formal definition of the concept
	Definition *string `json:"definition,omitempty"`

	// Concepts have both a display and an array of designation
	Designation []CodeSystemConceptDesignation `json:"designation,omitempty"`

	// A human readable string that is the recommended default way to present this concept to a user
	Display *string `json:"display,omitempty"`

	// A property value for this concept
	Property []CodeSystemConceptProperty `json:"property,omitempty"`
}

// CodeSystem represents a code system
type CodeSystem struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "CodeSystem"

	// The 'date' element may be more recent than the approval date because of minor changes
	ApprovalDate *string `json:"approvalDate,omitempty"`

	// An individual or organization primarily involved in the creation and maintenance of the CodeSystem
	Author []ContactDetail `json:"author,omitempty"`

	// If this value is missing, then it is not specified whether a code system is case sensitive or not
	CaseSensitive *bool `json:"caseSensitive,omitempty"`

	// Note that the code system resource does not define what the compositional grammar is
	Compositional *bool `json:"compositional,omitempty"`

	// If this is empty, it means that the code system resource does not represent the content of the code system
	Concept []CodeSystemConcept `json:"concept,omitempty"`

	// May be a web site, an email address, a telephone number, etc
	Contact []ContactDetail `json:"contact,omitempty"`

	// The extent of the content of the code system
	Content CodeSystemContent `json:"content"`

	// Sometimes, the copyright differs between the code system and the codes that are included
	Copyright *string `json:"copyright,omitempty"`

	// The (c) symbol should NOT be included in this string
	CopyrightLabel *string `json:"copyrightLabel,omitempty"`

	// The count of concepts defined in this resource cannot be more than this value
	Count *int `json:"count,omitempty"`

	// A free text natural language description of the code system from a consumer's perspective
	Description *string `json:"description,omitempty"`

	// A Boolean value to indicate that this code system is authored for testing purposes
	Experimental *bool `json:"experimental,omitempty"`

	// A filter that can be used in a value set compose statement when selecting concepts using a filter
	Filter []CodeSystemFilter `json:"filter,omitempty"`

	// The hierarchyMeaning element describes how the concepts in the code system are organized
	HierarchyMeaning *CodeSystemHierarchyMeaning `json:"hierarchyMeaning,omitempty"`

	// A formal identifier that is used to identify this code system when it is represented in other formats
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// A legal or geographic region in which the code system is intended to be used
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`

	// The name is not expected to be globally unique
	Name *string `json:"name,omitempty"`

	// A property defines an additional slot through which additional information can be provided
	Property []CodeSystemProperty `json:"property,omitempty"`

	// Usually an organization but may be an individual
	Publisher *string `json:"publisher,omitempty"`

	// This element does not describe the usage of the code system. Instead, it provides traceability
	Purpose *string `json:"purpose,omitempty"`

	// The status of this code system
	Status PublicationStatus `json:"status"`

	// A short, descriptive, user-friendly title for the code system
	Title *string `json:"title,omitempty"`

	// An absolute URI that is used to identify this code system when it is referenced in a specification
	URL *string `json:"url,omitempty"`

	// The content was developed with a focus and intent of supporting the contexts that are listed
	UseContext []UsageContext `json:"useContext,omitempty"`

	// The identifier that is used to identify this version of the code system when it is referenced
	Version *string `json:"version,omitempty"`

	// The version is used for this rather than the version of the code system
	VersionNeeded *bool `json:"versionNeeded,omitempty"`
}

// CodeSystemContent represents the extent of the content of the code system
type CodeSystemContent string

const (
	CodeSystemContentNotPresent CodeSystemContent = "not-present"
	CodeSystemContentExample    CodeSystemContent = "example"
	CodeSystemContentFragment   CodeSystemContent = "fragment"
	CodeSystemContentComplete   CodeSystemContent = "complete"
	CodeSystemContentSupplement CodeSystemContent = "supplement"
)

// CodeSystemHierarchyMeaning represents the hierarchy meaning of a code system
type CodeSystemHierarchyMeaning string

const (
	CodeSystemHierarchyMeaningGroupedBy    CodeSystemHierarchyMeaning = "grouped-by"
	CodeSystemHierarchyMeaningIsA          CodeSystemHierarchyMeaning = "is-a"
	CodeSystemHierarchyMeaningPartOf       CodeSystemHierarchyMeaning = "part-of"
	CodeSystemHierarchyMeaningClassifiedBy CodeSystemHierarchyMeaning = "classified-by"
)

// PublicationStatus represents the publication status of a resource
type PublicationStatus string

const (
	PublicationStatusDraft   PublicationStatus = "draft"
	PublicationStatusActive  PublicationStatus = "active"
	PublicationStatusRetired PublicationStatus = "retired"
	PublicationStatusUnknown PublicationStatus = "unknown"
)
