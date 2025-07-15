// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// GraphDefinitionStatus represents the status of a graph definition
type GraphDefinitionStatus string

const (
	GraphDefinitionStatusDraft   GraphDefinitionStatus = "draft"
	GraphDefinitionStatusActive  GraphDefinitionStatus = "active"
	GraphDefinitionStatusRetired GraphDefinitionStatus = "retired"
	GraphDefinitionStatusUnknown GraphDefinitionStatus = "unknown"
)

// GraphDefinitionNode represents a potential target for the link
type GraphDefinitionNode struct {
	common.BackboneElement

	// Information about why this node is of interest in this graph definition
	Description *string `json:"description,omitempty"`

	// Internal ID of node - target for link references
	NodeId string `json:"nodeId"`

	// Profile for the target resource
	Profile *string `json:"profile,omitempty"`

	// Type of resource this link refers to
	Type string `json:"type"`
}

// GraphDefinitionLinkCompartmentRule represents the compartment rule
type GraphDefinitionLinkCompartmentRule string

const (
	GraphDefinitionLinkCompartmentRuleIdentical GraphDefinitionLinkCompartmentRule = "identical"
	GraphDefinitionLinkCompartmentRuleMatching  GraphDefinitionLinkCompartmentRule = "matching"
	GraphDefinitionLinkCompartmentRuleDifferent GraphDefinitionLinkCompartmentRule = "different"
	GraphDefinitionLinkCompartmentRuleNoRule    GraphDefinitionLinkCompartmentRule = "no-rule"
)

// GraphDefinitionLinkCompartment represents compartment consistency rules
type GraphDefinitionLinkCompartment struct {
	common.BackboneElement

	// Identifies the compartment
	Code common.CodeableConcept `json:"code"`

	// Describes how the compartment rule is used
	Description *string `json:"description,omitempty"`

	// Identifies the compartment
	Expression *string `json:"expression,omitempty"`

	// identical | matching | different | no-rule
	Rule GraphDefinitionLinkCompartmentRule `json:"rule"`

	// Custom rule, as a FHIRPath expression
	Use common.CodeableConcept `json:"use"`
}

// GraphDefinitionLink represents links this graph makes rules about
type GraphDefinitionLink struct {
	common.BackboneElement

	// Compartment consistency rules
	Compartment []GraphDefinitionLinkCompartment `json:"compartment,omitempty"`

	// Information about why this link is of interest in this graph definition
	Description *string `json:"description,omitempty"`

	// Maximum occurrences for this link
	Max *string `json:"max,omitempty"`

	// Minimum occurrences for this link
	Min *int `json:"min,omitempty"`

	// At least one of the parameters must have the value {ref} which identifies the focus resource
	Params *string `json:"params,omitempty"`

	// The path expression cannot contain a resolve() function
	Path *string `json:"path,omitempty"`

	// Which slice (if profiled)
	SliceName *string `json:"sliceName,omitempty"`

	// The source node for this link
	SourceId string `json:"sourceId"`

	// The target node for this link
	TargetId string `json:"targetId"`
}

// GraphDefinition represents a formal computable definition of a graph of resources
type GraphDefinition struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "GraphDefinition"

	// May be a web site, an email address, a telephone number, etc.
	Contact []ContactDetail `json:"contact,omitempty"`

	// Copyright statement
	Copyright *string `json:"copyright,omitempty"`

	// Copyright label
	CopyrightLabel *string `json:"copyrightLabel,omitempty"`

	// The date is often not tracked until the resource is published
	Date *string `json:"date,omitempty"`

	// This description can be used to capture details such as comments about misuse
	Description *string `json:"description,omitempty"`

	// Allows filtering of graph definitions that are appropriate for use versus not
	Experimental *bool `json:"experimental,omitempty"`

	// A formal identifier that is used to identify this GraphDefinition
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// It may be possible for the graph definition to be used in jurisdictions
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`

	// Links this graph makes rules about
	Link []GraphDefinitionLink `json:"link,omitempty"`

	// The name is not expected to be globally unique
	Name string `json:"name"`

	// Potential target for the link
	Node []GraphDefinitionNode `json:"node,omitempty"`

	// Usually an organization but may be an individual
	Publisher *string `json:"publisher,omitempty"`

	// This element does not describe the usage of the graph definition
	Purpose *string `json:"purpose,omitempty"`

	// The Node at which instances of this graph start
	Start *string `json:"start,omitempty"`

	// Allows filtering of graph definitions that are appropriate for use versus not
	Status GraphDefinitionStatus `json:"status"`

	// This name does not need to be machine-processing friendly
	Title *string `json:"title,omitempty"`

	// Can be a urn:uuid: or a urn:oid: but real http: addresses are preferred
	Url *string `json:"url,omitempty"`

	// When multiple useContexts are specified, there is no expectation that all or any of the contexts apply
	UseContext []UsageContext `json:"useContext,omitempty"`

	// There may be different graph definition instances that have the same identifier
	Version *string `json:"version,omitempty"`

	// Version algorithm
	VersionAlgorithmString *string `json:"versionAlgorithmString,omitempty"`
}
