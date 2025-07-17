package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// OperationDefinition represents a formal computable definition of an operation
type OperationDefinition struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "OperationDefinition"

	// What http methods can be used for the operation depends on the .affectsState value
	AffectsState *bool `json:"affectsState,omitempty"`

	// A constrained profile can make optional parameters required or not used and clarify documentation
	Base *string `json:"base,omitempty"`

	// The name used to invoke the operation
	Code string `json:"code"`

	// Additional information about how to use this operation or named query
	Comment *string `json:"comment,omitempty"`

	// May be a web site, an email address, a telephone number, etc
	Contact []common.ContactDetail `json:"contact,omitempty"`

	// Note that this is not the same as the resource last-modified-date
	Date *string `json:"date,omitempty"`

	// This description can be used to capture details such as why the operation definition was built
	Description *string `json:"description,omitempty"`

	// Allows filtering of operation definitions that are appropriate for use versus not
	Experimental *bool `json:"experimental,omitempty"`

	// If present the profile shall not conflict with what is specified in the parameters
	InputProfile *string `json:"inputProfile,omitempty"`

	// Indicates whether this operation can be invoked on a particular instance of one of the given types
	Instance bool `json:"instance"`

	// It may be possible for the operation definition to be used in jurisdictions other than those for which it was originally designed
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`

	// Named queries are invoked differently, and have different capabilities
	Kind OperationDefinitionKind `json:"kind"`

	// The name is not expected to be globally unique
	Name string `json:"name"`

	// If present the profile shall not conflict with what is specified in the parameters
	OutputProfile *string `json:"outputProfile,omitempty"`

	// The combinations are suggestions as to which sets of parameters to use together
	Overload []OperationDefinitionOverload `json:"overload,omitempty"`

	// Query Definitions only have one output parameter, named "result"
	Parameter []OperationDefinitionParameter `json:"parameter,omitempty"`

	// Usually an organization but may be an individual
	Publisher *string `json:"publisher,omitempty"`

	// This element does not describe the usage of the operation definition
	Purpose *string `json:"purpose,omitempty"`

	// If the type is an abstract resource ("Resource" or "DomainResource") then the operation can be invoked on any concrete specialization
	Resource []string `json:"resource,omitempty"`

	// Allows filtering of operation definitions that are appropriate for use versus not
	Status OperationDefinitionStatus `json:"status"`

	// Indicates whether this operation or named query can be invoked at the system level
	System bool `json:"system"`

	// This name does not need to be machine-processing friendly and may contain punctuation, white-space, etc
	Title *string `json:"title,omitempty"`

	// Indicates whether this operation or named query can be invoked at the resource type level
	Type bool `json:"type"`

	// Can be a urn:uuid: or a urn:oid: but real http: addresses are preferred
	Url *string `json:"url,omitempty"`

	// When multiple useContexts are specified, there is no expectation that all or any of the contexts apply
	UseContext []common.UsageContext `json:"useContext,omitempty"`

	// There may be different operation definition instances that have the same identifier but different versions
	Version *string `json:"version,omitempty"`
}

// OperationDefinitionKind represents named queries are invoked differently
type OperationDefinitionKind string

const (
	OperationDefinitionKindOperation OperationDefinitionKind = "operation"
	OperationDefinitionKindQuery     OperationDefinitionKind = "query"
)

// OperationDefinitionStatus represents the status of the operation definition
type OperationDefinitionStatus string

const (
	OperationDefinitionStatusDraft   OperationDefinitionStatus = "draft"
	OperationDefinitionStatusActive  OperationDefinitionStatus = "active"
	OperationDefinitionStatusRetired OperationDefinitionStatus = "retired"
	OperationDefinitionStatusUnknown OperationDefinitionStatus = "unknown"
)

// OperationDefinitionParameter represents query definitions only have one output parameter
type OperationDefinitionParameter struct {
	common.BackboneElement

	// Binds to a value set if this parameter is coded (code, Coding, CodeableConcept)
	Binding *OperationDefinitionParameterBinding `json:"binding,omitempty"`

	// Describes the meaning or use of this parameter
	Documentation *string `json:"documentation,omitempty"`

	// The maximum number of times this element is permitted to appear in the request or response
	Max string `json:"max"`

	// The minimum number of times this parameter SHALL appear in the request or response
	Min int `json:"min"`

	// This name must be a token (start with a letter in a..z, and only contain letters, numerals, and underscore
	Name string `json:"name"`

	// Query Definitions only have one output parameter, named "result"
	Part []OperationDefinitionParameter `json:"part,omitempty"`

	// Resolution applies if the referenced parameter exists
	ReferencedFrom []OperationDefinitionParameterReferencedFrom `json:"referencedFrom,omitempty"`

	// How the parameter is understood as a search parameter
	SearchType *OperationDefinitionParameterSearchType `json:"searchType,omitempty"`

	// Often, these profiles are the base definitions from the spec
	TargetProfile []string `json:"targetProfile,omitempty"`

	// if there is no stated parameter, then the parameter is a multi-part parameter
	Type *string `json:"type,omitempty"`

	// If a parameter name is used for both an input and an output parameter, the parameter should be defined twice
	Use OperationDefinitionParameterUse `json:"use"`
}

// OperationDefinitionParameterUse represents if a parameter name is used for both an input and an output parameter
type OperationDefinitionParameterUse string

const (
	OperationDefinitionParameterUseIn  OperationDefinitionParameterUse = "in"
	OperationDefinitionParameterUseOut OperationDefinitionParameterUse = "out"
)

// OperationDefinitionParameterSearchType represents how the parameter is understood as a search parameter
type OperationDefinitionParameterSearchType string

const (
	OperationDefinitionParameterSearchTypeNumber    OperationDefinitionParameterSearchType = "number"
	OperationDefinitionParameterSearchTypeDate      OperationDefinitionParameterSearchType = "date"
	OperationDefinitionParameterSearchTypeString    OperationDefinitionParameterSearchType = "string"
	OperationDefinitionParameterSearchTypeToken     OperationDefinitionParameterSearchType = "token"
	OperationDefinitionParameterSearchTypeReference OperationDefinitionParameterSearchType = "reference"
	OperationDefinitionParameterSearchTypeComposite OperationDefinitionParameterSearchType = "composite"
	OperationDefinitionParameterSearchTypeQuantity  OperationDefinitionParameterSearchType = "quantity"
	OperationDefinitionParameterSearchTypeUri       OperationDefinitionParameterSearchType = "uri"
	OperationDefinitionParameterSearchTypeSpecial   OperationDefinitionParameterSearchType = "special"
)

// OperationDefinitionParameterBinding represents binds to a value set if this parameter is coded
type OperationDefinitionParameterBinding struct {
	common.BackboneElement

	// For further discussion, see Using Terminologies
	Strength OperationDefinitionParameterBindingStrength `json:"strength"`

	// For value sets with a referenceResource, the display can contain the value set description
	ValueSet string `json:"valueSet"`
}

// OperationDefinitionParameterBindingStrength represents for further discussion, see Using Terminologies
type OperationDefinitionParameterBindingStrength string

const (
	OperationDefinitionParameterBindingStrengthRequired   OperationDefinitionParameterBindingStrength = "required"
	OperationDefinitionParameterBindingStrengthExtensible OperationDefinitionParameterBindingStrength = "extensible"
	OperationDefinitionParameterBindingStrengthPreferred  OperationDefinitionParameterBindingStrength = "preferred"
	OperationDefinitionParameterBindingStrengthExample    OperationDefinitionParameterBindingStrength = "example"
)

// OperationDefinitionParameterReferencedFrom represents resolution applies if the referenced parameter exists
type OperationDefinitionParameterReferencedFrom struct {
	common.BackboneElement

	// The name of the parameter or dot-separated path of parameter names pointing to the resource parameter
	Source string `json:"source"`

	// The id of the element in the referencing resource that is expected to resolve to this resource
	SourceId *string `json:"sourceId,omitempty"`
}

// OperationDefinitionOverload represents the combinations are suggestions as to which sets of parameters to use together
type OperationDefinitionOverload struct {
	common.BackboneElement

	// Comments to go on overload
	Comment *string `json:"comment,omitempty"`

	// Name of parameter to include in overload
	ParameterName []string `json:"parameterName,omitempty"`
}
