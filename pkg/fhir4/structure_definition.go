package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// StructureDefinition represents a definition of a FHIR structure
type StructureDefinition struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "StructureDefinition"

	// Abstract Resources cannot be instantiated - a concrete sub-type must be used
	Abstract bool `json:"abstract"`

	// If differential constraints are specified in this structure, they are applied to the base in a "differential" fashion
	BaseDefinition *string `json:"baseDefinition,omitempty"`

	// May be a web site, an email address, a telephone number, etc
	Contact []common.ContactDetail `json:"contact,omitempty"`

	// Identifies the types of resource or data type elements to which the extension can be applied
	Context []StructureDefinitionContext `json:"context,omitempty"`

	// The rules are only evaluated when the extension is present
	ContextInvariant []string `json:"contextInvariant,omitempty"`

	// A copyright statement relating to the structure definition and/or its contents
	Copyright *string `json:"copyright,omitempty"`

	// Note that this is not the same as the resource last-modified-date
	Date *string `json:"date,omitempty"`

	// If the definition is a specialization, then it adds new elements in the differential
	Derivation *StructureDefinitionDerivation `json:"derivation,omitempty"`

	// This description can be used to capture details such as why the structure definition was built
	Description *string `json:"description,omitempty"`

	// A differential view is expressed relative to the base StructureDefinition
	Differential *StructureDefinitionDifferential `json:"differential,omitempty"`

	// Allows filtering of structure definitions that are appropriate for use versus not
	Experimental *bool `json:"experimental,omitempty"`

	// A StructureDefinition does not need to specify the target it applies to
	FhirVersion *string `json:"fhirVersion,omitempty"`

	// Typically, this is used for identifiers that can go in an HL7 V3 II (instance identifier) data type
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// It may be possible for the structure definition to be used in jurisdictions other than those for which it was originally designed
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`

	// A set of key words or terms from external terminologies that may be used to assist with indexing and searching
	Keyword []common.Coding `json:"keyword,omitempty"`

	// Defines the kind of structure that this definition is describing
	Kind StructureDefinitionKind `json:"kind"`

	// An external specification that the content is mapped to
	Mapping []StructureDefinitionMapping `json:"mapping,omitempty"`

	// The name is not expected to be globally unique
	Name string `json:"name"`

	// Usually an organization but may be an individual
	Publisher *string `json:"publisher,omitempty"`

	// This element does not describe the usage of the structure definition
	Purpose *string `json:"purpose,omitempty"`

	// A snapshot view is expressed in a standalone form that can be used and interpreted without considering the base StructureDefinition
	Snapshot *StructureDefinitionSnapshot `json:"snapshot,omitempty"`

	// Allows filtering of structure definitions that are appropriate for use versus not
	Status StructureDefinitionStatus `json:"status"`

	// This name does not need to be machine-processing friendly and may contain punctuation, white-space, etc
	Title *string `json:"title,omitempty"`

	// The type must match the elements defined in the differential and the snapshot
	Type string `json:"type"`

	// Can be a urn:uuid: or a urn:oid: but real http: addresses are preferred
	Url string `json:"url"`

	// When multiple useContexts are specified, there is no expectation that all or any of the contexts apply
	UseContext []common.UsageContext `json:"useContext,omitempty"`

	// There may be different structure definition instances that have the same identifier but different versions
	Version *string `json:"version,omitempty"`
}

// StructureDefinitionStatus represents allows filtering of structure definitions that are appropriate for use versus not
type StructureDefinitionStatus string

const (
	StructureDefinitionStatusDraft   StructureDefinitionStatus = "draft"
	StructureDefinitionStatusActive  StructureDefinitionStatus = "active"
	StructureDefinitionStatusRetired StructureDefinitionStatus = "retired"
	StructureDefinitionStatusUnknown StructureDefinitionStatus = "unknown"
)

// StructureDefinitionKind represents defines the kind of structure that this definition is describing
type StructureDefinitionKind string

const (
	StructureDefinitionKindPrimitiveType StructureDefinitionKind = "primitive-type"
	StructureDefinitionKindComplexType   StructureDefinitionKind = "complex-type"
	StructureDefinitionKindResource      StructureDefinitionKind = "resource"
	StructureDefinitionKindLogical       StructureDefinitionKind = "logical"
)

// StructureDefinitionDerivation represents if the definition is a specialization, then it adds new elements in the differential
type StructureDefinitionDerivation string

const (
	StructureDefinitionDerivationSpecialization StructureDefinitionDerivation = "specialization"
	StructureDefinitionDerivationConstraint     StructureDefinitionDerivation = "constraint"
)

// StructureDefinitionMapping represents an external specification that the content is mapped to
type StructureDefinitionMapping struct {
	common.BackboneElement

	// Comments about this mapping, including version notes, issues, scope limitations, and other important notes for usage
	Comment *string `json:"comment,omitempty"`

	// The specification is described once, with general comments, and then specific mappings are made that reference this declaration
	Identity string `json:"identity"`

	// A name for the specification that is being mapped to
	Name *string `json:"name,omitempty"`

	// A formal identity for the specification being mapped to helps with identifying maps consistently
	Uri *string `json:"uri,omitempty"`
}

// StructureDefinitionContext represents identifies the types of resource or data type elements to which the extension can be applied
type StructureDefinitionContext struct {
	common.BackboneElement

	// An expression that defines where an extension can be used in resources
	Expression string `json:"expression"`

	// Defines how to interpret the expression that defines what the context of the extension is
	Type StructureDefinitionContextType `json:"type"`
}

// StructureDefinitionContextType represents defines how to interpret the expression that defines what the context of the extension is
type StructureDefinitionContextType string

const (
	StructureDefinitionContextTypeFhirpath  StructureDefinitionContextType = "fhirpath"
	StructureDefinitionContextTypeElement   StructureDefinitionContextType = "element"
	StructureDefinitionContextTypeExtension StructureDefinitionContextType = "extension"
)

// StructureDefinitionSnapshot represents a snapshot view is expressed in a standalone form that can be used and interpreted without considering the base StructureDefinition
type StructureDefinitionSnapshot struct {
	common.BackboneElement

	// Captures constraints on each element within the resource
	Element []common.ElementDefinition `json:"element"`
}

// StructureDefinitionDifferential represents a differential view is expressed relative to the base StructureDefinition
type StructureDefinitionDifferential struct {
	common.BackboneElement

	// Captures constraints on each element within the resource
	Element []common.ElementDefinition `json:"element"`
}
