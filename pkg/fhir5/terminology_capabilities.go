package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// TerminologyCapabilitiesClosure represents closure operation capabilities
type TerminologyCapabilitiesClosure struct {
	common.BackboneElement

	// If cross-system closure is supported
	Translation *bool `json:"translation,omitempty"`
}

// TerminologyCapabilitiesCodeSystem represents code system capabilities
type TerminologyCapabilitiesCodeSystem struct {
	common.BackboneElement

	// URI for the Code System
	Uri *string `json:"uri,omitempty"`

	// Version of Code System supported
	Version []TerminologyCapabilitiesCodeSystemVersion `json:"version,omitempty"`

	// True if subsumption is supported
	Subsumption *bool `json:"subsumption,omitempty"`
}

// TerminologyCapabilitiesCodeSystemVersion represents version-specific capabilities
type TerminologyCapabilitiesCodeSystemVersion struct {
	common.BackboneElement

	// Version identifier for this version
	Code *string `json:"code,omitempty"`

	// If this is the default version for this code system
	IsDefault *bool `json:"isDefault,omitempty"`

	// If compositional grammar is supported
	Compositional *bool `json:"compositional,omitempty"`

	// Language Displays supported
	Language []string `json:"language,omitempty"`

	// Filter Properties supported
	Filter []TerminologyCapabilitiesCodeSystemVersionFilter `json:"filter,omitempty"`

	// Properties supported for $lookup
	Property []string `json:"property,omitempty"`
}

// TerminologyCapabilitiesCodeSystemVersionFilter represents filter capabilities
type TerminologyCapabilitiesCodeSystemVersionFilter struct {
	common.BackboneElement

	// Code of the property supported
	Code string `json:"code"`

	// Operations supported for the property
	Op []string `json:"op"`
}

// TerminologyCapabilitiesExpansion represents expansion operation capabilities
type TerminologyCapabilitiesExpansion struct {
	common.BackboneElement

	// Whether the server can return nested value sets
	Hierarchical *bool `json:"hierarchical,omitempty"`

	// Whether the server supports paging on expansion
	Paging *bool `json:"paging,omitempty"`

	// Allow request for incomplete expansions
	Incomplete *bool `json:"incomplete,omitempty"`

	// Supported expansion parameter
	Parameter []TerminologyCapabilitiesExpansionParameter `json:"parameter,omitempty"`

	// Documentation about text searching
	TextFilter *string `json:"textFilter,omitempty"`
}

// TerminologyCapabilitiesExpansionParameter represents expansion parameter capabilities
type TerminologyCapabilitiesExpansionParameter struct {
	common.BackboneElement

	// Expansion Parameter name
	Name string `json:"name"`

	// Description of support for parameter
	Documentation *string `json:"documentation,omitempty"`
}

// TerminologyCapabilitiesImplementation represents implementation details
type TerminologyCapabilitiesImplementation struct {
	common.BackboneElement

	// Describes this specific instance
	Description string `json:"description"`

	// Base URL for the implementation
	Url *string `json:"url,omitempty"`
}

// TerminologyCapabilitiesSoftware represents software details
type TerminologyCapabilitiesSoftware struct {
	common.BackboneElement

	// A name the software is known by
	Name string `json:"name"`

	// Version covered by this statement
	Version *string `json:"version,omitempty"`

	// Date this version was released
	ReleaseDate *string `json:"releaseDate,omitempty"`
}

// TerminologyCapabilitiesTranslation represents translation operation capabilities
type TerminologyCapabilitiesTranslation struct {
	common.BackboneElement

	// Whether the client must identify the map
	NeedsMap bool `json:"needsMap"`
}

// TerminologyCapabilitiesValidateCode represents validate-code operation capabilities
type TerminologyCapabilitiesValidateCode struct {
	common.BackboneElement

	// Whether translations are validated
	Translations bool `json:"translations"`
}

// TerminologyCapabilities represents a set of capabilities of a FHIR Terminology Server
type TerminologyCapabilities struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "TerminologyCapabilities"

	// Whether the $closure operation is supported
	Closure *TerminologyCapabilitiesClosure `json:"closure,omitempty"`

	// See notes on the ValueSet resource
	CodeSearch *TerminologyCapabilitiesCodeSearch `json:"codeSearch,omitempty"`

	// The code system - identified by its system URL - may also be declared explicitly as a Code System Resource at /CodeSystem, but it might not be
	CodeSystem []TerminologyCapabilitiesCodeSystem `json:"codeSystem,omitempty"`

	// May be a web site, an email address, a telephone number, etc
	Contact []common.ContactDetail `json:"contact,omitempty"`

	// A copyright statement relating to the terminology capabilities and/or its contents
	Copyright *string `json:"copyright,omitempty"`

	// The (c) symbol should NOT be included in this string
	CopyrightLabel *string `json:"copyrightLabel,omitempty"`

	// The date is often not tracked until the resource is published, but may be present on draft content
	Date string `json:"date"`

	// This description can be used to capture details such as comments about misuse, instructions for clinical use and interpretation, literature references, examples from the paper world, etc
	Description *string `json:"description,omitempty"`

	// Information about the ValueSet/$expand operation
	Expansion *TerminologyCapabilitiesExpansion `json:"expansion,omitempty"`

	// Allows filtering of terminology capabilitiess that are appropriate for use versus not
	Experimental *bool `json:"experimental,omitempty"`

	// Typically, this is used for identifiers that can go in an HL7 V3 II (instance identifier) data type, and can then identify this terminology capabilities outside of FHIR, where it is not possible to use the logical URI
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Identifies a specific implementation instance that is described by the terminology capability statement - i.e. a particular installation, rather than the capabilities of a software program
	Implementation *TerminologyCapabilitiesImplementation `json:"implementation,omitempty"`

	// It may be possible for the terminology capabilities to be used in jurisdictions other than those for which it was originally designed or intended
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`

	// The way that this statement is intended to be used, to describe an actual running instance of software, a particular product (kind, not instance of software) or a class of implementation (e.g. a desired purchase)
	Kind TerminologyCapabilitiesKind `json:"kind"`

	// Whether the server supports lockedDate
	LockedDate *bool `json:"lockedDate,omitempty"`

	// The name is not expected to be globally unique
	Name *string `json:"name,omitempty"`

	// Usually an organization but may be an individual
	Publisher *string `json:"publisher,omitempty"`

	// This element does not describe the usage of the terminology capabilities
	Purpose *string `json:"purpose,omitempty"`

	// Software that is covered by this terminology capability statement
	Software *TerminologyCapabilitiesSoftware `json:"software,omitempty"`

	// Allows filtering of terminology capabilitiess that are appropriate for use versus not
	Status TerminologyCapabilitiesStatus `json:"status"`

	// This name does not need to be machine-processing friendly and may contain punctuation, white-space, etc
	Title *string `json:"title,omitempty"`

	// Information about the ConceptMap/$translate operation
	Translation *TerminologyCapabilitiesTranslation `json:"translation,omitempty"`

	// Can be a urn:uuid: or a urn:oid: but real http: addresses are preferred
	Url *string `json:"url,omitempty"`

	// When multiple useContexts are specified, there is no expectation that all or any of the contexts apply
	UseContext []common.UsageContext `json:"useContext,omitempty"`

	// Information about the ValueSet/$validate-code operation
	ValidateCode *TerminologyCapabilitiesValidateCode `json:"validateCode,omitempty"`

	// There may be different terminology capabilities instances that have the same identifier but different versions
	Version *string `json:"version,omitempty"`

	// If set as a string, this is a FHIRPath expression that has two additional context variables passed in - %version1 and %version2 and will return a negative number if version1 is newer, a positive number if version2 and a 0 if the version ordering can't be successfully be determined
	VersionAlgorithmString *string `json:"versionAlgorithmString,omitempty"`

	// If set as a string, this is a FHIRPath expression that has two additional context variables passed in - %version1 and %version2 and will return a negative number if version1 is newer, a positive number if version2 and a 0 if the version ordering can't be successfully be determined
	VersionAlgorithmCoding *common.Coding `json:"versionAlgorithmCoding,omitempty"`
}

// TerminologyCapabilitiesCodeSearch represents code search capabilities
type TerminologyCapabilitiesCodeSearch string

const (
	TerminologyCapabilitiesCodeSearchInCompose            TerminologyCapabilitiesCodeSearch = "in-compose"
	TerminologyCapabilitiesCodeSearchInExpansion          TerminologyCapabilitiesCodeSearch = "in-expansion"
	TerminologyCapabilitiesCodeSearchInComposeOrExpansion TerminologyCapabilitiesCodeSearch = "in-compose-or-expansion"
)

// TerminologyCapabilitiesKind represents the kind of terminology capabilities
type TerminologyCapabilitiesKind string

const (
	TerminologyCapabilitiesKindInstance     TerminologyCapabilitiesKind = "instance"
	TerminologyCapabilitiesKindCapability   TerminologyCapabilitiesKind = "capability"
	TerminologyCapabilitiesKindRequirements TerminologyCapabilitiesKind = "requirements"
)

// TerminologyCapabilitiesStatus represents the status of terminology capabilities
type TerminologyCapabilitiesStatus string

const (
	TerminologyCapabilitiesStatusDraft   TerminologyCapabilitiesStatus = "draft"
	TerminologyCapabilitiesStatusActive  TerminologyCapabilitiesStatus = "active"
	TerminologyCapabilitiesStatusRetired TerminologyCapabilitiesStatus = "retired"
	TerminologyCapabilitiesStatusUnknown TerminologyCapabilitiesStatus = "unknown"
)
