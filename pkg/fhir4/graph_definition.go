package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// GraphDefinition represents a formal computable definition of a graph of resources
type GraphDefinition struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "GraphDefinition"

	// May be a web site, an email address, a telephone number, etc
	Contact []common.ContactDetail `json:"contact,omitempty"`

	// Note that this is not the same as the resource last-modified-date
	Date *string `json:"date,omitempty"`

	// This description can be used to capture details such as why the graph definition was built
	Description *string `json:"description,omitempty"`

	// Allows filtering of graph definitions that are appropriate for use versus not
	Experimental *bool `json:"experimental,omitempty"`

	// It may be possible for the graph definition to be used in jurisdictions other than those for which it was originally designed or intended
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`

	// Links this graph makes rules about
	Link []GraphDefinitionLink `json:"link,omitempty"`

	// The name is not expected to be globally unique
	Name string `json:"name"`

	// The code does not include the '$' prefix that is always included in the URL when the operation is invoked
	Profile *string `json:"profile,omitempty"`

	// Usually an organization but may be an individual
	Publisher *string `json:"publisher,omitempty"`

	// This element does not describe the usage of the graph definition
	Purpose *string `json:"purpose,omitempty"`

	// The type of FHIR resource at which instances of this graph start
	Start string `json:"start"`

	// Allows filtering of graph definitions that are appropriate for use versus not
	Status GraphDefinitionStatus `json:"status"`

	// Can be a urn:uuid: or a urn:oid: but real http: addresses are preferred
	Url *string `json:"url,omitempty"`

	// When multiple useContexts are specified, there is no expectation that all or any of the contexts apply
	UseContext []common.UsageContext `json:"useContext,omitempty"`

	// There may be different graph definition instances that have the same identifier but different versions
	Version *string `json:"version,omitempty"`
}

// GraphDefinitionStatus represents the status of the graph definition
type GraphDefinitionStatus string

const (
	GraphDefinitionStatusDraft   GraphDefinitionStatus = "draft"
	GraphDefinitionStatusActive  GraphDefinitionStatus = "active"
	GraphDefinitionStatusRetired GraphDefinitionStatus = "retired"
	GraphDefinitionStatusUnknown GraphDefinitionStatus = "unknown"
)

// GraphDefinitionLink represents links this graph makes rules about
type GraphDefinitionLink struct {
	common.BackboneElement

	// Information about why this link is of interest in this graph definition
	Description *string `json:"description,omitempty"`

	// Maximum occurrences for this link
	Max *string `json:"max,omitempty"`

	// Minimum occurrences for this link
	Min *int `json:"min,omitempty"`

	// The path expression cannot contain a resolve() function
	Path *string `json:"path,omitempty"`

	// Which slice (if profiled)
	SliceName *string `json:"sliceName,omitempty"`

	// Potential target for the link
	Target []GraphDefinitionLinkTarget `json:"target,omitempty"`
}

// GraphDefinitionLinkTarget represents potential target for the link
type GraphDefinitionLinkTarget struct {
	common.BackboneElement

	// Compartment Consistency Rules
	Compartment []GraphDefinitionLinkTargetCompartment `json:"compartment,omitempty"`

	// Additional links from target resource
	Link []GraphDefinitionLink `json:"link,omitempty"`

	// At least one of the parameters must have the value {ref} which identifies the focus resource
	Params *string `json:"params,omitempty"`

	// Profile for the target resource
	Profile *string `json:"profile,omitempty"`

	// Type of resource this link refers to
	Type string `json:"type"`
}

// GraphDefinitionLinkTargetCompartment represents compartment consistency rules
type GraphDefinitionLinkTargetCompartment struct {
	common.BackboneElement

	// Identifies the compartment
	Code GraphDefinitionLinkTargetCompartmentCode `json:"code"`

	// Documentation for FHIRPath expression
	Description *string `json:"description,omitempty"`

	// Custom rule, as a FHIRPath expression
	Expression *string `json:"expression,omitempty"`

	// identical | matching | different | no-rule | custom
	Rule GraphDefinitionLinkTargetCompartmentRule `json:"rule"`

	// All conditional rules are evaluated; if they are true, then the rules are evaluated
	Use GraphDefinitionLinkTargetCompartmentUse `json:"use"`
}

// GraphDefinitionLinkTargetCompartmentCode represents the compartment code
type GraphDefinitionLinkTargetCompartmentCode string

const (
	GraphDefinitionLinkTargetCompartmentCodePatient       GraphDefinitionLinkTargetCompartmentCode = "Patient"
	GraphDefinitionLinkTargetCompartmentCodeEncounter     GraphDefinitionLinkTargetCompartmentCode = "Encounter"
	GraphDefinitionLinkTargetCompartmentCodeRelatedPerson GraphDefinitionLinkTargetCompartmentCode = "RelatedPerson"
	GraphDefinitionLinkTargetCompartmentCodePractitioner  GraphDefinitionLinkTargetCompartmentCode = "Practitioner"
	GraphDefinitionLinkTargetCompartmentCodeDevice        GraphDefinitionLinkTargetCompartmentCode = "Device"
)

// GraphDefinitionLinkTargetCompartmentRule represents the compartment rule
type GraphDefinitionLinkTargetCompartmentRule string

const (
	GraphDefinitionLinkTargetCompartmentRuleIdentical GraphDefinitionLinkTargetCompartmentRule = "identical"
	GraphDefinitionLinkTargetCompartmentRuleMatching  GraphDefinitionLinkTargetCompartmentRule = "matching"
	GraphDefinitionLinkTargetCompartmentRuleDifferent GraphDefinitionLinkTargetCompartmentRule = "different"
	GraphDefinitionLinkTargetCompartmentRuleCustom    GraphDefinitionLinkTargetCompartmentRule = "custom"
)

// GraphDefinitionLinkTargetCompartmentUse represents the compartment use
type GraphDefinitionLinkTargetCompartmentUse string

const (
	GraphDefinitionLinkTargetCompartmentUseCondition   GraphDefinitionLinkTargetCompartmentUse = "condition"
	GraphDefinitionLinkTargetCompartmentUseRequirement GraphDefinitionLinkTargetCompartmentUse = "requirement"
)
