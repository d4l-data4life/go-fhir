package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// ImplementationGuide represents a set of rules about how FHIR resources are used
type ImplementationGuide struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "ImplementationGuide"

	// May be a web site, an email address, a telephone number, etc
	Contact []common.ContactDetail `json:"contact,omitempty"`

	// A copyright statement relating to the implementation guide and/or its contents
	Copyright *string `json:"copyright,omitempty"`

	// Note that this is not the same as the resource last-modified-date
	Date *string `json:"date,omitempty"`

	// Principally, this consists of information about source resource and file locations, and build parameters and templates
	Definition *ImplementationGuideDefinition `json:"definition,omitempty"`

	// Another implementation guide that this implementation depends on
	DependsOn []ImplementationGuideDependsOn `json:"dependsOn,omitempty"`

	// This description can be used to capture details such as why the implementation guide was built
	Description *string `json:"description,omitempty"`

	// Allows filtering of implementation guides that are appropriate for use versus not
	Experimental *bool `json:"experimental,omitempty"`

	// Most implementation guides target a single version
	FhirVersion []string `json:"fhirVersion"`

	// See Default Profiles for a discussion of which resources are 'covered' by an implementation guide
	Global []ImplementationGuideGlobal `json:"global,omitempty"`

	// It may be possible for the implementation guide to be used in jurisdictions other than those for which it was originally designed or intended
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`

	// The license that applies to this Implementation Guide, using an SPDX license code, or 'not-open-source'
	License *string `json:"license,omitempty"`

	// Information about an assembled implementation guide, created by the publication tooling
	Manifest *ImplementationGuideManifest `json:"manifest,omitempty"`

	// The name is not expected to be globally unique
	Name string `json:"name"`

	// Many (if not all) IG publishing tools will require that this element be present
	PackageId string `json:"packageId"`

	// Usually an organization but may be an individual
	Publisher *string `json:"publisher,omitempty"`

	// Allows filtering of implementation guides that are appropriate for use versus not
	Status ImplementationGuideStatus `json:"status"`

	// This name does not need to be machine-processing friendly and may contain punctuation, white-space, etc
	Title *string `json:"title,omitempty"`

	// Can be a urn:uuid: or a urn:oid: but real http: addresses are preferred
	Url string `json:"url"`

	// When multiple useContexts are specified, there is no expectation that all or any of the contexts apply
	UseContext []common.UsageContext `json:"useContext,omitempty"`

	// There may be different implementation guide instances that have the same identifier but different versions
	Version *string `json:"version,omitempty"`
}

// ImplementationGuideStatus represents the status of the implementation guide
type ImplementationGuideStatus string

const (
	ImplementationGuideStatusDraft   ImplementationGuideStatus = "draft"
	ImplementationGuideStatusActive  ImplementationGuideStatus = "active"
	ImplementationGuideStatusRetired ImplementationGuideStatus = "retired"
	ImplementationGuideStatusUnknown ImplementationGuideStatus = "unknown"
)

// ImplementationGuideDependsOn represents another implementation guide that this implementation depends on
type ImplementationGuideDependsOn struct {
	common.BackboneElement

	// The NPM package name for the Implementation Guide that this IG depends on
	PackageId *string `json:"packageId,omitempty"`

	// Usually, A canonical reference to the implementation guide is the same as the master location at which the implementation guide is published
	Uri string `json:"uri"`

	// This follows the syntax of the NPM packaging version field
	Version *string `json:"version,omitempty"`
}

// ImplementationGuideGlobal represents see Default Profiles for a discussion of which resources are 'covered' by an implementation guide
type ImplementationGuideGlobal struct {
	common.BackboneElement

	// The profile allows the implementation guide to specify the types and structure of profiles that are allowed to be referenced in this implementation guide when using inline resources
	Profile string `json:"profile"`

	// The type of resource this definition applies to
	Type string `json:"type"`
}

// ImplementationGuideDefinition represents principally, this consists of information about source resource and file locations, and build parameters and templates
type ImplementationGuideDefinition struct {
	common.BackboneElement

	// Grouping used to present related resources in the IG
	Grouping []ImplementationGuideDefinitionGrouping `json:"grouping,omitempty"`

	// A page/Section in the implementation guide
	Page *ImplementationGuideDefinitionPage `json:"page,omitempty"`

	// Defines how IG is built by tools
	Parameter []ImplementationGuideDefinitionParameter `json:"parameter,omitempty"`

	// A resource that is part of the implementation guide
	Resource []ImplementationGuideDefinitionResource `json:"resource,omitempty"`

	// A template for building resources
	Template []ImplementationGuideDefinitionTemplate `json:"template,omitempty"`
}

// ImplementationGuideDefinitionGrouping represents grouping used to present related resources in the IG
type ImplementationGuideDefinitionGrouping struct {
	common.BackboneElement

	// Human readable text describing the package
	Description *string `json:"description,omitempty"`

	// The human-readable title to display for the package of resources when rendering the implementation guide
	Name string `json:"name"`
}

// ImplementationGuideDefinitionResource represents a resource that is part of the implementation guide
type ImplementationGuideDefinitionResource struct {
	common.BackboneElement

	// A description of the reason that a resource has been included in the implementation guide
	Description *string `json:"description,omitempty"`

	// Examples, if present, represent resources that conform to this resource, and cannot be represented by inline resources
	ExampleBoolean *bool `json:"exampleBoolean,omitempty"`

	// Examples, if present, represent resources that conform to this resource, and cannot be represented by inline resources
	ExampleCanonical *string `json:"exampleCanonical,omitempty"`

	// A human assigned name for the resource. All resources SHOULD have a name, but the name may be extracted from the resource (e.g. ValueSet.name)
	Name *string `json:"name,omitempty"`

	// Reference of the resource
	Reference common.Reference `json:"reference"`
}

// ImplementationGuideDefinitionPage represents a page/Section in the implementation guide
type ImplementationGuideDefinitionPage struct {
	common.BackboneElement

	// The format of the page
	Generation ImplementationGuideDefinitionPageGeneration `json:"generation"`

	// The name of the page to be generated
	Name string `json:"name"`

	// Nested Pages/Sections under this page
	Page []ImplementationGuideDefinitionPage `json:"page,omitempty"`

	// The source address for the page
	SourceUrl *string `json:"sourceUrl,omitempty"`

	// The title of the page to be generated
	Title *string `json:"title,omitempty"`
}

// ImplementationGuideDefinitionPageGeneration represents the format of the page
type ImplementationGuideDefinitionPageGeneration string

const (
	ImplementationGuideDefinitionPageGenerationHtml      ImplementationGuideDefinitionPageGeneration = "html"
	ImplementationGuideDefinitionPageGenerationMarkdown  ImplementationGuideDefinitionPageGeneration = "markdown"
	ImplementationGuideDefinitionPageGenerationXml       ImplementationGuideDefinitionPageGeneration = "xml"
	ImplementationGuideDefinitionPageGenerationGenerated ImplementationGuideDefinitionPageGeneration = "generated"
)

// ImplementationGuideDefinitionParameter represents defines how IG is built by tools
type ImplementationGuideDefinitionParameter struct {
	common.BackboneElement

	// apply | path-resource | path-pages | path-tx-cache | expansion-parameter | rule-broken-links | generate-xml | generate-json | generate-turtle | html-template
	Code string `json:"code"`

	// Value for named type, or id, path, or url for a resource
	Value string `json:"value"`
}

// ImplementationGuideDefinitionTemplate represents a template for building resources
type ImplementationGuideDefinitionTemplate struct {
	common.BackboneElement

	// The code specifies the type of template
	Code string `json:"code"`

	// The source location for the template
	Source string `json:"source"`

	// The scope in which the template applies
	Scope *string `json:"scope,omitempty"`
}

// ImplementationGuideManifest represents information about an assembled implementation guide, created by the publication tooling
type ImplementationGuideManifest struct {
	common.BackboneElement

	// A pointer to other implementation guides that this implementation depends on
	Image []string `json:"image,omitempty"`

	// A pointer to other implementation guides that this implementation depends on
	Other []string `json:"other,omitempty"`

	// Information about a page within the IG
	Page []ImplementationGuideManifestPage `json:"page,omitempty"`

	// A pointer to other implementation guides that this implementation depends on
	Rendering *string `json:"rendering,omitempty"`

	// A resource that is part of the implementation guide
	Resource []ImplementationGuideManifestResource `json:"resource"`
}

// ImplementationGuideManifestResource represents a resource that is part of the implementation guide
type ImplementationGuideManifestResource struct {
	common.BackboneElement

	// If true or a reference, indicates the resource is an example instance
	ExampleBoolean *bool `json:"exampleBoolean,omitempty"`

	// If true or a reference, indicates the resource is an example instance
	ExampleCanonical *string `json:"exampleCanonical,omitempty"`

	// Relative path for page in IG
	Reference common.Reference `json:"reference"`

	// Relative path for page in IG
	RelativePath *string `json:"relativePath,omitempty"`
}

// ImplementationGuideManifestPage represents information about a page within the IG
type ImplementationGuideManifestPage struct {
	common.BackboneElement

	// Relative path for page in IG
	Anchor []string `json:"anchor,omitempty"`

	// Relative path for page in IG
	Name string `json:"name"`

	// Relative path for page in IG
	Title *string `json:"title,omitempty"`
}
