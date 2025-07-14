// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// CapabilityStatementKind represents the kind of capability statement
type CapabilityStatementKind string

const (
	CapabilityStatementKindInstance     CapabilityStatementKind = "instance"
	CapabilityStatementKindCapability   CapabilityStatementKind = "capability"
	CapabilityStatementKindRequirements CapabilityStatementKind = "requirements"
)

// CapabilityStatementStatus represents the status of a capability statement
type CapabilityStatementStatus string

const (
	CapabilityStatementStatusDraft   CapabilityStatementStatus = "draft"
	CapabilityStatementStatusActive  CapabilityStatementStatus = "active"
	CapabilityStatementStatusRetired CapabilityStatementStatus = "retired"
	CapabilityStatementStatusUnknown CapabilityStatementStatus = "unknown"
)

// CapabilityStatementRestMode represents the mode of a REST interface
type CapabilityStatementRestMode string

const (
	CapabilityStatementRestModeClient CapabilityStatementRestMode = "client"
	CapabilityStatementRestModeServer CapabilityStatementRestMode = "server"
)

// CapabilityStatementRestResourceInteractionCode represents the code of a resource interaction
type CapabilityStatementRestResourceInteractionCode string

const (
	CapabilityStatementRestResourceInteractionCodeRead            CapabilityStatementRestResourceInteractionCode = "read"
	CapabilityStatementRestResourceInteractionCodeVread           CapabilityStatementRestResourceInteractionCode = "vread"
	CapabilityStatementRestResourceInteractionCodeUpdate          CapabilityStatementRestResourceInteractionCode = "update"
	CapabilityStatementRestResourceInteractionCodePatch           CapabilityStatementRestResourceInteractionCode = "patch"
	CapabilityStatementRestResourceInteractionCodeDelete          CapabilityStatementRestResourceInteractionCode = "delete"
	CapabilityStatementRestResourceInteractionCodeHistoryInstance CapabilityStatementRestResourceInteractionCode = "history-instance"
	CapabilityStatementRestResourceInteractionCodeHistoryType     CapabilityStatementRestResourceInteractionCode = "history-type"
	CapabilityStatementRestResourceInteractionCodeCreate          CapabilityStatementRestResourceInteractionCode = "create"
	CapabilityStatementRestResourceInteractionCodeSearchType      CapabilityStatementRestResourceInteractionCode = "search-type"
)

// CapabilityStatementRestInteractionCode represents the code of a system interaction
type CapabilityStatementRestInteractionCode string

const (
	CapabilityStatementRestInteractionCodeTransaction   CapabilityStatementRestInteractionCode = "transaction"
	CapabilityStatementRestInteractionCodeBatch         CapabilityStatementRestInteractionCode = "batch"
	CapabilityStatementRestInteractionCodeSearchSystem  CapabilityStatementRestInteractionCode = "search-system"
	CapabilityStatementRestInteractionCodeHistorySystem CapabilityStatementRestInteractionCode = "history-system"
)

// CapabilityStatementRestResourceConditionalDelete represents conditional delete support
type CapabilityStatementRestResourceConditionalDelete string

const (
	CapabilityStatementRestResourceConditionalDeleteNotSupported CapabilityStatementRestResourceConditionalDelete = "not-supported"
	CapabilityStatementRestResourceConditionalDeleteSingle       CapabilityStatementRestResourceConditionalDelete = "single"
	CapabilityStatementRestResourceConditionalDeleteMultiple     CapabilityStatementRestResourceConditionalDelete = "multiple"
)

// CapabilityStatementRestResourceConditionalRead represents conditional read support
type CapabilityStatementRestResourceConditionalRead string

const (
	CapabilityStatementRestResourceConditionalReadNotSupported  CapabilityStatementRestResourceConditionalRead = "not-supported"
	CapabilityStatementRestResourceConditionalReadModifiedSince CapabilityStatementRestResourceConditionalRead = "modified-since"
	CapabilityStatementRestResourceConditionalReadNotMatch      CapabilityStatementRestResourceConditionalRead = "not-match"
	CapabilityStatementRestResourceConditionalReadFullSupport   CapabilityStatementRestResourceConditionalRead = "full-support"
)

// CapabilityStatementRestResourceReferencePolicy represents reference policy
type CapabilityStatementRestResourceReferencePolicy string

const (
	CapabilityStatementRestResourceReferencePolicyLiteral  CapabilityStatementRestResourceReferencePolicy = "literal"
	CapabilityStatementRestResourceReferencePolicyLogical  CapabilityStatementRestResourceReferencePolicy = "logical"
	CapabilityStatementRestResourceReferencePolicyResolves CapabilityStatementRestResourceReferencePolicy = "resolves"
	CapabilityStatementRestResourceReferencePolicyEnforced CapabilityStatementRestResourceReferencePolicy = "enforced"
	CapabilityStatementRestResourceReferencePolicyLocal    CapabilityStatementRestResourceReferencePolicy = "local"
)

// CapabilityStatementRestResourceVersioning represents versioning support
type CapabilityStatementRestResourceVersioning string

const (
	CapabilityStatementRestResourceVersioningNoVersion       CapabilityStatementRestResourceVersioning = "no-version"
	CapabilityStatementRestResourceVersioningVersioned       CapabilityStatementRestResourceVersioning = "versioned"
	CapabilityStatementRestResourceVersioningVersionedUpdate CapabilityStatementRestResourceVersioning = "versioned-update"
)

// CapabilityStatementRestResourceSearchParamType represents search parameter type
type CapabilityStatementRestResourceSearchParamType string

const (
	CapabilityStatementRestResourceSearchParamTypeNumber    CapabilityStatementRestResourceSearchParamType = "number"
	CapabilityStatementRestResourceSearchParamTypeDate      CapabilityStatementRestResourceSearchParamType = "date"
	CapabilityStatementRestResourceSearchParamTypeString    CapabilityStatementRestResourceSearchParamType = "string"
	CapabilityStatementRestResourceSearchParamTypeToken     CapabilityStatementRestResourceSearchParamType = "token"
	CapabilityStatementRestResourceSearchParamTypeReference CapabilityStatementRestResourceSearchParamType = "reference"
	CapabilityStatementRestResourceSearchParamTypeComposite CapabilityStatementRestResourceSearchParamType = "composite"
	CapabilityStatementRestResourceSearchParamTypeQuantity  CapabilityStatementRestResourceSearchParamType = "quantity"
	CapabilityStatementRestResourceSearchParamTypeUri       CapabilityStatementRestResourceSearchParamType = "uri"
	CapabilityStatementRestResourceSearchParamTypeSpecial   CapabilityStatementRestResourceSearchParamType = "special"
)

// CapabilityStatementMessagingMode represents messaging mode
type CapabilityStatementMessagingMode string

const (
	CapabilityStatementMessagingModeSender   CapabilityStatementMessagingMode = "sender"
	CapabilityStatementMessagingModeReceiver CapabilityStatementMessagingMode = "receiver"
)

// CapabilityStatementDocumentMode represents document mode
type CapabilityStatementDocumentMode string

const (
	CapabilityStatementDocumentModeProducer CapabilityStatementDocumentMode = "producer"
	CapabilityStatementDocumentModeConsumer CapabilityStatementDocumentMode = "consumer"
)

// CapabilityStatementSoftware represents software covered by this capability statement
type CapabilityStatementSoftware struct {
	common.BackboneElement

	// Name the software is known by
	Name string `json:"name"`

	// Date this version of the software was released
	ReleaseDate *string `json:"releaseDate,omitempty"`

	// If possible, a version should be specified, as statements are likely to be different for different versions of software
	Version *string `json:"version,omitempty"`
}

// CapabilityStatementImplementation represents a specific implementation instance
type CapabilityStatementImplementation struct {
	common.BackboneElement

	// The organization responsible for the management of the instance and oversight of the data on the server
	Custodian *common.Reference `json:"custodian,omitempty"`

	// Information about the specific installation that this capability statement relates to
	Description string `json:"description"`

	// An absolute base URL for the implementation
	URL *string `json:"url,omitempty"`
}

// CapabilityStatementRestSecurity represents security information
type CapabilityStatementRestSecurity struct {
	common.BackboneElement

	// The easiest CORS headers to add are Access-Control-Allow-Origin: * & Access-Control-Request-Method: GET, POST, PUT, DELETE
	CORS *bool `json:"cors,omitempty"`

	// General description of how security works
	Description *string `json:"description,omitempty"`

	// Types of security services that are supported/required by the system
	Service []common.CodeableConcept `json:"service,omitempty"`
}

// CapabilityStatementRestResourceInteraction represents a restful operation supported by the solution
type CapabilityStatementRestResourceInteraction struct {
	common.BackboneElement

	// Coded identifier of the operation, supported by the system resource
	Code CapabilityStatementRestResourceInteractionCode `json:"code"`

	// Guidance specific to the implementation of this operation
	Documentation *string `json:"documentation,omitempty"`
}

// CapabilityStatementRestResourceSearchParam represents a search parameter supported by the system
type CapabilityStatementRestResourceSearchParam struct {
	common.BackboneElement

	// This SHOULD be present, and matches refers to a SearchParameter by its canonical URL
	Definition *string `json:"definition,omitempty"`

	// This allows documentation of any distinct behaviors about how the search parameter is used
	Documentation *string `json:"documentation,omitempty"`

	// Parameter names cannot overlap with standard parameter names
	Name string `json:"name"`

	// While this can be looked up from the definition, it is included here as a convenience
	Type CapabilityStatementRestResourceSearchParamType `json:"type"`
}

// CapabilityStatementRestResourceOperation represents an operation supported by the system
type CapabilityStatementRestResourceOperation struct {
	common.BackboneElement

	// This can be used to build an HTML form to invoke the operation, for instance
	Definition string `json:"definition"`

	// Documentation that describes anything special about the operation behavior
	Documentation *string `json:"documentation,omitempty"`

	// The name here SHOULD be the same as the OperationDefinition.code in the referenced OperationDefinition
	Name string `json:"name"`
}

// CapabilityStatementRestResource represents a resource type supported by the system
type CapabilityStatementRestResource struct {
	common.BackboneElement

	// Conditional Create is mainly appropriate for interface engine scripts converting from other formats
	ConditionalCreate *bool `json:"conditionalCreate,omitempty"`

	// Conditional Delete is mainly appropriate for interface engine scripts converting from other formats
	ConditionalDelete *CapabilityStatementRestResourceConditionalDelete `json:"conditionalDelete,omitempty"`

	// Conditional Patch is mainly appropriate for interface engine scripts converting from other formats
	ConditionalPatch *bool `json:"conditionalPatch,omitempty"`

	// Conditional Read is mainly appropriate for interface engine scripts converting from other formats
	ConditionalRead *CapabilityStatementRestResourceConditionalRead `json:"conditionalRead,omitempty"`

	// Conditional Update is mainly appropriate for interface engine scripts converting from other formats
	ConditionalUpdate *bool `json:"conditionalUpdate,omitempty"`

	// Additional information about the resource type used by the system
	Documentation *string `json:"documentation,omitempty"`

	// In general, a Resource will only appear in a CapabilityStatement if the server actually has some capabilities
	Interaction []CapabilityStatementRestResourceInteraction `json:"interaction,omitempty"`

	// Operations linked from CapabilityStatement.rest.resource.operation
	Operation []CapabilityStatementRestResourceOperation `json:"operation,omitempty"`

	// All other profiles for this type that are listed in `.rest.resource.supportedProfile` must conform to this profile
	Profile *string `json:"profile,omitempty"`

	// It is useful to support the vRead operation for current operations, even if past versions aren't available
	ReadHistory *bool `json:"readHistory,omitempty"`

	// A set of flags that defines how references are supported
	ReferencePolicy []CapabilityStatementRestResourceReferencePolicy `json:"referencePolicy,omitempty"`

	// Documenting `_include` support helps set conformance expectations for the desired system
	SearchInclude []string `json:"searchInclude,omitempty"`

	// The search parameters should include the control search parameters such as _sort, _count, etc.
	SearchParam []CapabilityStatementRestResourceSearchParam `json:"searchParam,omitempty"`

	// See `CapabilityStatement.rest.resource.searchInclude` comments
	SearchRevInclude []string `json:"searchRevInclude,omitempty"`

	// Supported profiles must conform to the resource profile in the `.rest.resource.profile` element
	SupportedProfile []string `json:"supportedProfile,omitempty"`

	// A type of resource exposed via the restful interface
	Type string `json:"type"`

	// Allowing the clients to create new identities on the server
	UpdateCreate *bool `json:"updateCreate,omitempty"`

	// If a server supports versionIds correctly, it SHOULD support vread too
	Versioning *CapabilityStatementRestResourceVersioning `json:"versioning,omitempty"`
}

// CapabilityStatementRestInteraction represents a system interaction
type CapabilityStatementRestInteraction struct {
	common.BackboneElement

	// A coded identifier of the operation, supported by the system
	Code CapabilityStatementRestInteractionCode `json:"code"`

	// Guidance specific to the implementation of this operation
	Documentation *string `json:"documentation,omitempty"`
}

// CapabilityStatementRest represents a REST interface
type CapabilityStatementRest struct {
	common.BackboneElement

	// At present, the only defined compartments are at CompartmentDefinition
	Compartment []string `json:"compartment,omitempty"`

	// Information about the system's restful capabilities that apply across all applications
	Documentation *string `json:"documentation,omitempty"`

	// A specification of restful operations supported by the system
	Interaction []CapabilityStatementRestInteraction `json:"interaction,omitempty"`

	// Identifies whether this portion of the statement is describing the ability to initiate or receive restful operations
	Mode CapabilityStatementRestMode `json:"mode"`

	// CapabilityStatement.rest.operation is for operations invoked at the system level
	Operation []CapabilityStatementRestResourceOperation `json:"operation,omitempty"`

	// Max of one repetition per resource type
	Resource []CapabilityStatementRestResource `json:"resource,omitempty"`

	// Information about security implementation from an interface perspective
	Security *CapabilityStatementRestSecurity `json:"security,omitempty"`
}

// CapabilityStatementMessaging represents messaging capabilities
type CapabilityStatementMessaging struct {
	common.BackboneElement

	// Documentation about the system's messaging capabilities for this endpoint not otherwise documented by the capability statement
	Documentation *string `json:"documentation,omitempty"`

	// An endpoint (network accessible address) to which messages and/or replies are to be sent
	Endpoint []interface{} `json:"endpoint,omitempty"`

	// Identifies which messaging standard is supported
	SupportedMessage []interface{} `json:"supportedMessage,omitempty"`
}

// CapabilityStatementDocument represents document capabilities
type CapabilityStatementDocument struct {
	common.BackboneElement

	// Mode of this document declaration - whether an application is a producer or consumer
	Mode CapabilityStatementDocumentMode `json:"mode"`

	// A profile on the DocumentReference for documents that conform to this document definition
	Profile string `json:"profile"`
}

// CapabilityStatement represents a capability statement documenting the capabilities of a FHIR Server or Client
type CapabilityStatement struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "CapabilityStatement"

	// In general, if a server gets a request with an Accept-Language that it doesn't support, it should still return the resource
	AcceptLanguage []string `json:"acceptLanguage,omitempty"`

	// May be a web site, an email address, a telephone number, etc.
	Contact []ContactDetail `json:"contact,omitempty"`

	// Use and/or publishing restrictions
	Copyright *string `json:"copyright,omitempty"`

	// Copyright holder and year(s)
	CopyrightLabel *string `json:"copyrightLabel,omitempty"`

	// The date is often not tracked until the resource is published, but may be present on draft content
	Date string `json:"date"`

	// This description can be used to capture details such as comments about misuse, instructions for clinical use and interpretation
	Description *string `json:"description,omitempty"`

	// A document definition
	Document []CapabilityStatementDocument `json:"document,omitempty"`

	// Allows filtering of capability statements that are appropriate for use versus not
	Experimental *bool `json:"experimental,omitempty"`

	// Servers may implement multiple versions
	FhirVersion string `json:"fhirVersion"`

	// "xml", "json" and "ttl" are allowed, which describe the simple encodings described in the specification
	Format []string `json:"format"`

	// A formal identifier that is used to identify this CapabilityStatement when it is represented in other formats
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Identifies a specific implementation instance that is described by the capability statement
	Implementation *CapabilityStatementImplementation `json:"implementation,omitempty"`

	// Note: this is primarily only relevant in terms of ImplementationGuides that don't define specific CapabilityStatements
	ImplementationGuide []string `json:"implementationGuide,omitempty"`

	// the contents of any directly or indirectly imported CapabilityStatements SHALL NOT overlap
	Imports []string `json:"imports,omitempty"`

	// HL7 defines the following Services: Terminology Service
	Instantiates []string `json:"instantiates,omitempty"`

	// It may be possible for the capability statement to be used in jurisdictions other than those for which it was originally designed
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`

	// The way that this statement is intended to be used, to describe an actual running instance of software
	Kind CapabilityStatementKind `json:"kind"`

	// Multiple repetitions allow the documentation of multiple endpoints per solution
	Messaging []CapabilityStatementMessaging `json:"messaging,omitempty"`

	// The name is not expected to be globally unique
	Name *string `json:"name,omitempty"`

	// At present, the patch mime types application/json-patch+json and application/xml-patch+xml are legal
	PatchFormat []string `json:"patchFormat,omitempty"`

	// Usually an organization but may be an individual
	Publisher *string `json:"publisher,omitempty"`

	// This element does not describe the usage of the capability statement
	Purpose *string `json:"purpose,omitempty"`

	// Multiple repetitions allow definition of both client and/or server behaviors
	Rest []CapabilityStatementRest `json:"rest,omitempty"`

	// Software that is covered by this capability statement
	Software *CapabilityStatementSoftware `json:"software,omitempty"`

	// Allows filtering of capability statements that are appropriate for use versus not
	Status CapabilityStatementStatus `json:"status"`

	// This name does not need to be machine-processing friendly and may contain punctuation, white-space, etc.
	Title *string `json:"title,omitempty"`

	// Can be a urn:uuid: or a urn:oid: but real http: addresses are preferred
	URL *string `json:"url,omitempty"`

	// When multiple useContexts are specified, there is no expectation that all or any of the contexts apply
	UseContext []UsageContext `json:"useContext,omitempty"`

	// There may be different capability statement instances that have the same identifier but different versions
	Version *string `json:"version,omitempty"`

	// If set as a string, this is a FHIRPath expression that has two additional context variables passed in
	VersionAlgorithmString *string        `json:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *common.Coding `json:"versionAlgorithmCoding,omitempty"`
}
