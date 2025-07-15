// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// OperationDefinitionKind represents the kind of operation
type OperationDefinitionKind string

const (
	OperationDefinitionKindOperation OperationDefinitionKind = "operation"
	OperationDefinitionKindQuery     OperationDefinitionKind = "query"
)

// OperationDefinitionStatus represents the status of an operation definition
type OperationDefinitionStatus string

const (
	OperationDefinitionStatusDraft   OperationDefinitionStatus = "draft"
	OperationDefinitionStatusActive  OperationDefinitionStatus = "active"
	OperationDefinitionStatusRetired OperationDefinitionStatus = "retired"
	OperationDefinitionStatusUnknown OperationDefinitionStatus = "unknown"
)

// OperationDefinitionParameter represents a parameter for an operation
type OperationDefinitionParameter struct {
	common.BackboneElement

	// In previous versions of FHIR, there was an extension for this
	AllowedType []string `json:"allowedType,omitempty"`

	// Binds to a value set if this parameter is coded
	Binding *interface{} `json:"binding,omitempty"`

	// Describes the meaning or use of this parameter
	Documentation *string `json:"documentation,omitempty"`

	// The maximum number of times this element is permitted to appear
	Max string `json:"max"`

	// The minimum number of times this parameter SHALL appear
	Min int `json:"min"`

	// This name must be a token
	Name string `json:"name"`

	// Query Definitions only have one output parameter, named "result"
	Part []OperationDefinitionParameter `json:"part,omitempty"`

	// Resolution applies if the referenced parameter exists
	ReferencedFrom []interface{} `json:"referencedFrom,omitempty"`

	// If present, indicates that the parameter applies when the operation is being invoked
	Scope []string `json:"scope,omitempty"`

	// How the parameter is understood as a search parameter
	SearchType *string `json:"searchType,omitempty"`

	// The type for this parameter
	Type *string `json:"type,omitempty"`

	// Whether this is an input or an output parameter
	Use string `json:"use"`
}

// OperationDefinition represents a formal computable definition of an operation
type OperationDefinition struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "OperationDefinition"

	// What http methods can be used for the operation
	AffectsState *bool `json:"affectsState,omitempty"`

	// A constrained profile can make optional parameters required or not used
	Base *string `json:"base,omitempty"`

	// For maximum compatibility, use only lowercase ASCII characters
	Code string `json:"code"`

	// Additional information about how to use this operation or named query
	Comment *string `json:"comment,omitempty"`

	// May be a web site, an email address, a telephone number, etc.
	Contact []ContactDetail `json:"contact,omitempty"`

	// Note that this is not the same as the resource last-modified-date
	Date *string `json:"date,omitempty"`

	// This description can be used to capture details
	Description *string `json:"description,omitempty"`

	// Allows filtering of operation definition that are appropriate for use vs. not
	Experimental *bool `json:"experimental,omitempty"`

	// Operations that are idempotent may be invoked by performing an HTTP GET
	Idempotent *bool `json:"idempotent,omitempty"`

	// Indicates whether this operation can be invoked on a particular instance
	Instance bool `json:"instance"`

	// It may be possible for the operation definition to be used in jurisdictions
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`

	// Named queries are invoked differently, and have different capabilities
	Kind OperationDefinitionKind `json:"kind"`

	// The name is not expected to be globally unique
	Name string `json:"name"`

	// The combinations are suggestions as to which sets of parameters to use together
	Overload []interface{} `json:"overload,omitempty"`

	// Query Definitions only have one output parameter, named "result"
	Parameter []OperationDefinitionParameter `json:"parameter,omitempty"`

	// Usually an organization but may be an individual
	Publisher *string `json:"publisher,omitempty"`

	// This element does not describe the usage of the operation definition
	Purpose *string `json:"purpose,omitempty"`

	// If the type is an abstract resource then the operation can be invoked
	Resource []string `json:"resource,omitempty"`

	// draft | active | retired | unknown
	Status OperationDefinitionStatus `json:"status"`

	// Indicates whether this operation or named query can be invoked at the system level
	System bool `json:"system"`

	// This name does not need to be machine-processing friendly
	Title *string `json:"title,omitempty"`

	// Indicates whether this operation can be invoked at the resource type level
	Type bool `json:"type"`

	// Can be a urn:uuid: or a urn:oid: but real http: addresses are preferred
	Url *string `json:"url,omitempty"`

	// When multiple useContexts are specified
	UseContext []UsageContext `json:"useContext,omitempty"`

	// There may be different operation definitions that have the same url
	Version *string `json:"version,omitempty"`
}
