package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// TerminologyCapabilities represents a statement of system capabilities
type TerminologyCapabilities struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "TerminologyCapabilities"

	// Whether the $closure operation is supported
	Closure *TerminologyCapabilitiesClosure `json:"closure,omitempty"`

	// See notes on the ValueSet resource
	CodeSearch *string `json:"codeSearch,omitempty"` // "explicit" | "all"

	// The code system - identified by its system URL - may also be declared explicitly as a Code System Resource at /CodeSystem, but it might not be
	CodeSystem []TerminologyCapabilitiesCodeSystem `json:"codeSystem,omitempty"`

	// May be a web site, an email address, a telephone number, etc
	Contact []common.ContactDetail `json:"contact,omitempty"`

	// A copyright statement relating to the terminology capabilities and/or its contents
	Copyright *string `json:"copyright,omitempty"`

	// Note that this is not the same as the resource last-modified-date, since the resource may be a secondary representation of the terminology capabilities
	Date string `json:"date"`

	// This description can be used to capture details such as why the terminology capabilities was built
	Description *string `json:"description,omitempty"`

	// Information about the ValueSet/$expand operation
	Expansion *TerminologyCapabilitiesExpansion `json:"expansion,omitempty"`

	// Allows filtering of terminology capabilitiess that are appropriate for use versus not
	Experimental *bool `json:"experimental,omitempty"`

	// Identifies a specific implementation instance that is described by the terminology capability statement
	Implementation *TerminologyCapabilitiesImplementation `json:"implementation,omitempty"`

	// It may be possible for the terminology capabilities to be used in jurisdictions other than those for which it was originally designed or intended
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`

	// The way that this statement is intended to be used, to describe an actual running instance of software, a particular product or a class of implementation
	Kind string `json:"kind"` // "instance" | "capability" | "requirements"

	// Whether the server supports lockedDate
	LockedDate *bool `json:"lockedDate,omitempty"`

	// The name is not expected to be globally unique. The name should be a simple alphanumeric type name to ensure that it is machine-processing friendly
	Name *string `json:"name,omitempty"`

	// Usually an organization but may be an individual. The publisher (or steward) of the terminology capabilities is the organization or individual primarily responsible for the maintenance and upkeep of the terminology capabilities
	Publisher *string `json:"publisher,omitempty"`

	// This element does not describe the usage of the terminology capabilities. Instead, it provides traceability of 'why' the resource is either needed or 'why' it is defined as it is
	Purpose *string `json:"purpose,omitempty"`

	// Software that is covered by this terminology capability statement
	Software *TerminologyCapabilitiesSoftware `json:"software,omitempty"`

	// Allows filtering of terminology capabilitiess that are appropriate for use versus not
	Status string `json:"status"` // "draft" | "active" | "retired" | "unknown"

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
}

// TerminologyCapabilitiesSoftware represents software that is covered by this terminology capability statement
type TerminologyCapabilitiesSoftware struct {
	common.BackboneElement

	// Name the software is known by
	Name string `json:"name"`

	// If possible, a version should be specified, as statements are likely to be different for different versions of software
	Version *string `json:"version,omitempty"`
}

// TerminologyCapabilitiesImplementation represents identifies a specific implementation instance that is described by the terminology capability statement
type TerminologyCapabilitiesImplementation struct {
	common.BackboneElement

	// Information about the specific installation that this terminology capability statement relates to
	Description string `json:"description"`

	// An absolute base URL for the implementation
	Url *string `json:"url,omitempty"`
}

// TerminologyCapabilitiesCodeSystemVersionFilter represents filter Properties supported
type TerminologyCapabilitiesCodeSystemVersionFilter struct {
	common.BackboneElement

	// Code of the property supported
	Code string `json:"code"`

	// Operations supported for the property
	Op []string `json:"op"`
}

// TerminologyCapabilitiesCodeSystemVersion represents language translations might not be available for all codes
type TerminologyCapabilitiesCodeSystemVersion struct {
	common.BackboneElement

	// For version-less code systems, there should be a single version with no identifier
	Code *string `json:"code,omitempty"`

	// If the compositional grammar defined by the code system is supported
	Compositional *bool `json:"compositional,omitempty"`

	// Filter Properties supported
	Filter []TerminologyCapabilitiesCodeSystemVersionFilter `json:"filter,omitempty"`

	// If this is the default version for this code system
	IsDefault *bool `json:"isDefault,omitempty"`

	// Language Displays supported
	Language []string `json:"language,omitempty"`

	// Properties supported for $lookup
	Property []string `json:"property,omitempty"`
}

// TerminologyCapabilitiesCodeSystem represents the code system - identified by its system URL
type TerminologyCapabilitiesCodeSystem struct {
	common.BackboneElement

	// URI for the Code System
	Uri *string `json:"uri,omitempty"`

	// For the code system, a list of versions that are supported by the server
	Version []TerminologyCapabilitiesCodeSystemVersion `json:"version,omitempty"`

	// True if subsumption is supported for this version of the code system
	Subsumption *bool `json:"subsumption,omitempty"`
}

// TerminologyCapabilitiesExpansion represents information about the ValueSet/$expand operation
type TerminologyCapabilitiesExpansion struct {
	common.BackboneElement

	// Whether the server can return nested value sets
	Hierarchical *bool `json:"hierarchical,omitempty"`

	// Whether the server supports paging on expansion
	Paging *bool `json:"paging,omitempty"`

	// Allow request for incomplete expansions?
	Incomplete *bool `json:"incomplete,omitempty"`

	// Supported expansion parameter
	Parameter []TerminologyCapabilitiesExpansionParameter `json:"parameter,omitempty"`

	// Documentation about text searching works
	TextFilter *string `json:"textFilter,omitempty"`
}

// TerminologyCapabilitiesExpansionParameter represents supported expansion parameter
type TerminologyCapabilitiesExpansionParameter struct {
	common.BackboneElement

	// Expansion Parameter name
	Name string `json:"name"`

	// Description of support for parameter
	Documentation *string `json:"documentation,omitempty"`
}

// TerminologyCapabilitiesTranslation represents information about the ConceptMap/$translate operation
type TerminologyCapabilitiesTranslation struct {
	common.BackboneElement

	// Whether the client must identify the map
	NeedsMap bool `json:"needsMap"`
}

// TerminologyCapabilitiesValidateCode represents information about the ValueSet/$validate-code operation
type TerminologyCapabilitiesValidateCode struct {
	common.BackboneElement

	// Whether translations are validated
	Translations bool `json:"translations"`
}

// TerminologyCapabilitiesClosure represents whether the $closure operation is supported
type TerminologyCapabilitiesClosure struct {
	common.BackboneElement

	// If cross-system closure is supported
	Translation *bool `json:"translation,omitempty"`
}
