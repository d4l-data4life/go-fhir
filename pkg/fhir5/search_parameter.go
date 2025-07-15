// Package fhir5 contains FHIR R5 resource definitions
package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// SearchParameterComponent represents a component of a composite search parameter
type SearchParameterComponent struct {
	common.BackboneElement

	// The definition of the search parameter that describes this part
	Definition string `json:"definition"`

	// This expression overrides the expression in the definition and extracts the index values from the outcome of the composite expression
	Expression string `json:"expression"`
}

// SearchParameter represents a search parameter that defines a named search item that can be used to search/filter on a resource
type SearchParameter struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "SearchParameter"

	// A search parameter must always apply to at least one resource type
	Base []string `json:"base"`

	// Notes about chaining support
	Chain []string `json:"chain,omitempty"`

	// For maximum compatibility, use only lowercase ASCII characters
	Code string `json:"code"`

	// Used to define the parts of a composite search parameter
	Component []SearchParameterComponent `json:"component,omitempty"`

	// The type of value a search parameter refers to, and how the content is interpreted
	Comparator []string `json:"comparator,omitempty"`

	// A compound search parameter that can be used to define a complex search parameter
	Composite *string `json:"composite,omitempty"`

	// Contact details to assist a user in finding and communicating with the publisher
	Contact []ContactDetail `json:"contact,omitempty"`

	// The date when the search parameter was last reviewed
	Date *string `json:"date,omitempty"`

	// A FHIRPath expression that returns a set of elements for the search parameter
	DerivedFrom *string `json:"derivedFrom,omitempty"`

	// This description can be used to capture details
	Description *string `json:"description,omitempty"`

	// Allows filtering of search parameters that are appropriate for use versus not
	Experimental *bool `json:"experimental,omitempty"`

	// A FHIRPath expression that returns a set of elements for the search parameter
	Expression *string `json:"expression,omitempty"`

	// The type of value a search parameter refers to, and how the content is interpreted
	Modifier []string `json:"modifier,omitempty"`

	// A natural language name identifying the search parameter
	Name string `json:"name"`

	// The type of value a search parameter refers to, and how the content is interpreted
	MultipleOr *bool `json:"multipleOr,omitempty"`

	// The type of value a search parameter refers to, and how the content is interpreted
	MultipleAnd *bool `json:"multipleAnd,omitempty"`

	// The type of value a search parameter refers to, and how the content is interpreted
	Operator []string `json:"operator,omitempty"`

	// Usually an organization but may be an individual
	Publisher *string `json:"publisher,omitempty"`

	// This element does not describe the usage of the search parameter
	Purpose *string `json:"purpose,omitempty"`

	// The type of value a search parameter refers to, and how the content is interpreted
	Status string `json:"status"`

	// The type of value a search parameter refers to, and how the content is interpreted
	Target []string `json:"target,omitempty"`

	// The type of value a search parameter refers to, and how the content is interpreted
	Type string `json:"type"`

	// Can be a urn:uuid: or a urn:oid: but real http: addresses are preferred
	Url string `json:"url"`

	// A detailed description of how the search parameter is used from a clinical perspective
	Usage *string `json:"usage,omitempty"`

	// When multiple useContexts are specified, there is no expectation that all or any of the contexts apply
	UseContext []interface{} `json:"useContext,omitempty"`

	// There may be different search parameter instances that have the same identifier but different versions
	Version *string `json:"version,omitempty"`

	// The type of value a search parameter refers to, and how the content is interpreted
	Xpath *string `json:"xpath,omitempty"`

	// The type of value a search parameter refers to, and how the content is interpreted
	XpathUsage *string `json:"xpathUsage,omitempty"`
}
