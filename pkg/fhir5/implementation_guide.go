package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// ImplementationGuideStatus represents the status of an implementation guide
type ImplementationGuideStatus string

const (
	ImplementationGuideStatusDraft   ImplementationGuideStatus = "draft"
	ImplementationGuideStatusActive  ImplementationGuideStatus = "active"
	ImplementationGuideStatusRetired ImplementationGuideStatus = "retired"
	ImplementationGuideStatusUnknown ImplementationGuideStatus = "unknown"
)

// ImplementationGuideDefinitionPageGeneration represents how the page is generated
type ImplementationGuideDefinitionPageGeneration string

const (
	ImplementationGuideDefinitionPageGenerationHTML      ImplementationGuideDefinitionPageGeneration = "html"
	ImplementationGuideDefinitionPageGenerationMarkdown  ImplementationGuideDefinitionPageGeneration = "markdown"
	ImplementationGuideDefinitionPageGenerationXML       ImplementationGuideDefinitionPageGeneration = "xml"
	ImplementationGuideDefinitionPageGenerationGenerated ImplementationGuideDefinitionPageGeneration = "generated"
)

// ImplementationGuideDependsOn represents another implementation guide that this implementation depends on
type ImplementationGuideDependsOn struct {
	common.BackboneElement

	// The NPM package name for the Implementation Guide that this IG depends on
	PackageId *string `json:"packageId,omitempty"`

	// This doesn't need to enumerate every resource used, but should give some sense of why the dependency exists
	Reason *string `json:"reason,omitempty"`

	// Usually, A canonical reference to the implementation guide is the same as the master location
	Uri string `json:"uri"`

	// This follows the syntax of the NPM packaging version field
	Version *string `json:"version,omitempty"`
}

// ImplementationGuideGlobal represents global profiles for resources
type ImplementationGuideGlobal struct {
	common.BackboneElement

	// A reference to the profile that all instances must conform to
	Profile string `json:"profile"`

	// The type must match that of the profile that is referred to
	Type string `json:"type"`
}

// ImplementationGuideDefinitionGrouping represents groupings of content
type ImplementationGuideDefinitionGrouping struct {
	common.BackboneElement

	// Human readable text describing the package
	Description *string `json:"description,omitempty"`

	// The human-readable title to display for the package of resources
	Name string `json:"name"`
}

// ImplementationGuideDefinitionResource represents a resource that is part of the implementation guide
type ImplementationGuideDefinitionResource struct {
	common.BackboneElement

	// This is mostly used with examples to explain why it is present
	Description *string `json:"description,omitempty"`

	// The resource SHALL be valid against all the versions it is specified to apply to
	FhirVersion []string `json:"fhirVersion,omitempty"`

	// This must correspond to a group.id element within this implementation guide
	GroupingId *string `json:"groupingId,omitempty"`

	// If true, indicates the resource is an example instance
	IsExample *bool `json:"isExample,omitempty"`

	// A human assigned name for the resource
	Name *string `json:"name,omitempty"`

	// Examples: StructureDefinition -> Any, ValueSet -> expansion, etc.
	Profile []string `json:"profile,omitempty"`

	// Usually this is a relative URL that locates the resource within the implementation guide
	Reference common.Reference `json:"reference"`
}

// ImplementationGuideDefinitionPage represents pages in the implementation guide
type ImplementationGuideDefinitionPage struct {
	common.BackboneElement

	// A code that indicates how the page is generated
	Generation ImplementationGuideDefinitionPageGeneration `json:"generation"`

	// This SHALL be a local reference, expressed with respect to the root of the IG output folder
	Name string `json:"name"`

	// Sub-pages (recursive structure)
	Page []ImplementationGuideDefinitionPage `json:"page,omitempty"`

	// If absent and the page isn't a generated page, this may be inferred from the page name
	SourceUrl      *string `json:"sourceUrl,omitempty"`
	SourceString   *string `json:"sourceString,omitempty"`
	SourceMarkdown *string `json:"sourceMarkdown,omitempty"`

	// A short title used to represent this page in navigational structures
	Title string `json:"title"`
}

// ImplementationGuideDefinitionParameter represents parameters for the implementation guide
type ImplementationGuideDefinitionParameter struct {
	common.BackboneElement

	// A tool-specific code that defines the parameter
	Code common.Coding `json:"code"`

	// Value for named type
	Value string `json:"value"`
}

// ImplementationGuideDefinitionTemplate represents a template for building resources
type ImplementationGuideDefinitionTemplate struct {
	common.BackboneElement

	// Type of template specified
	Code string `json:"code"`

	// The scope in which the template applies
	Scope *string `json:"scope,omitempty"`

	// The source location for the template
	Source string `json:"source"`
}

// ImplementationGuideDefinition represents the definition of the implementation guide
type ImplementationGuideDefinition struct {
	common.BackboneElement

	// Groupings are arbitrary sub-divisions of content
	Grouping []ImplementationGuideDefinitionGrouping `json:"grouping,omitempty"`

	// Pages automatically become sections if they have sub-pages
	Page *ImplementationGuideDefinitionPage `json:"page,omitempty"`

	// Parameters for the implementation guide
	Parameter []ImplementationGuideDefinitionParameter `json:"parameter,omitempty"`

	// A resource that is part of the implementation guide
	Resource []ImplementationGuideDefinitionResource `json:"resource,omitempty"`

	// A template for building resources
	Template []ImplementationGuideDefinitionTemplate `json:"template,omitempty"`
}

// ImplementationGuideManifestResource represents a resource that is part of the implementation guide
type ImplementationGuideManifestResource struct {
	common.BackboneElement

	// If true, indicates the resource is an example instance
	IsExample *bool `json:"isExample,omitempty"`

	// Examples: StructureDefinition -> Any, ValueSet -> expansion, etc.
	Profile []string `json:"profile,omitempty"`

	// Usually this is a relative URL that locates the resource within the implementation guide
	Reference common.Reference `json:"reference"`

	// Appending 'rendering' + "/" + this should resolve to the resource page
	RelativePath *string `json:"relativePath,omitempty"`
}

// ImplementationGuideManifestPage represents information about a page within the IG
type ImplementationGuideManifestPage struct {
	common.BackboneElement

	// Appending 'rendering' + "/" + page.name + "#" + page.anchor should resolve to the anchor
	Anchor []string `json:"anchor,omitempty"`

	// Appending 'rendering' + "/" + this should resolve to the page
	Name string `json:"name"`

	// Label for the page intended for human display
	Title *string `json:"title,omitempty"`
}

// ImplementationGuideManifest represents information about an assembled implementation guide
type ImplementationGuideManifest struct {
	common.BackboneElement

	// Indicates a relative path to an image that exists within the IG
	Image []string `json:"image,omitempty"`

	// Indicates the relative path of an additional non-page, non-image file that is part of the IG
	Other []string `json:"other,omitempty"`

	// Information about a page within the IG
	Page []ImplementationGuideManifestPage `json:"page,omitempty"`

	// A pointer to official web page, PDF or other rendering of the implementation guide
	Rendering *string `json:"rendering,omitempty"`

	// A resource that is part of the implementation guide
	Resource []ImplementationGuideManifestResource `json:"resource"`
}

// ImplementationGuide represents a set of rules of how a particular interoperability or standards problem is solved
type ImplementationGuide struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "ImplementationGuide"

	// May be a web site, an email address, a telephone number, etc.
	Contact []ContactDetail `json:"contact,omitempty"`

	// Use and/or publishing restrictions
	Copyright *string `json:"copyright,omitempty"`

	// Copyright holder and year(s)
	CopyrightLabel *string `json:"copyrightLabel,omitempty"`

	// Date last changed
	Date *string `json:"date,omitempty"`

	// Principally, this consists of information about source resource and file locations
	Definition *ImplementationGuideDefinition `json:"definition,omitempty"`

	// Another implementation guide that this implementation depends on
	DependsOn []ImplementationGuideDependsOn `json:"dependsOn,omitempty"`

	// Natural language description of the implementation guide
	Description *string `json:"description,omitempty"`

	// For testing purposes, not real usage
	Experimental *bool `json:"experimental,omitempty"`

	// FHIR Version(s) this Implementation Guide targets
	FhirVersion []string `json:"fhirVersion"`

	// See Default Profiles for a discussion of which resources are 'covered' by an implementation guide
	Global []ImplementationGuideGlobal `json:"global,omitempty"`

	// Additional identifier for the implementation guide
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Intended jurisdiction for implementation guide (if applicable)
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`

	// The license that applies to this Implementation Guide, using an SPDX license code
	License *string `json:"license,omitempty"`

	// Information about an assembled implementation guide, created by the publication tooling
	Manifest *ImplementationGuideManifest `json:"manifest,omitempty"`

	// Name for this implementation guide (computer friendly)
	Name string `json:"name"`

	// Many (if not all) IG publishing tools will require that this element be present
	PackageId string `json:"packageId"`

	// Name of the publisher (organization or individual)
	Publisher *string `json:"publisher,omitempty"`

	// Why this implementation guide is defined
	Purpose *string `json:"purpose,omitempty"`

	// draft | active | retired | unknown
	Status ImplementationGuideStatus `json:"status"`

	// Name for this implementation guide (human friendly)
	Title *string `json:"title,omitempty"`

	// Canonical identifier for this implementation guide, represented as a URI
	Url string `json:"url"`

	// The context that the content is intended to support
	UseContext []UsageContext `json:"useContext,omitempty"`

	// Business version of the implementation guide
	Version *string `json:"version,omitempty"`

	// How to compare versions
	VersionAlgorithmString *string        `json:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *common.Coding `json:"versionAlgorithmCoding,omitempty"`
}
