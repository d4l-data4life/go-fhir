package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// SearchParameter represents a search parameter that defines a named search item that can be used to search/filter on a resource
type SearchParameter struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "SearchParameter"

	// A search parameter must always apply to at least one resource type
	Base []string `json:"base"`

	// Systems are not required to list all the chain names they support
	Chain []string `json:"chain,omitempty"`

	// For maximum compatibility, use only lowercase ASCII characters
	Code string `json:"code"`

	// If no comparators are listed, clients should not expect servers to support any comparators
	Comparator []SearchParameterComparator `json:"comparator,omitempty"`

	// Used to define the parts of a composite search parameter
	Component []SearchParameterComponent `json:"component,omitempty"`

	// May be a web site, an email address, a telephone number, etc
	Contact []common.ContactDetail `json:"contact,omitempty"`

	// Note that this is not the same as the resource last-modified-date
	Date *string `json:"date,omitempty"`

	// The intent of this is that a server can designate that it provides support for a search parameter defined in the specification itself
	DerivedFrom *string `json:"derivedFrom,omitempty"`

	// This description can be used to capture details such as why the search parameter was built
	Description string `json:"description"`

	// Allows filtering of search parameters that are appropriate for use versus not
	Experimental *bool `json:"experimental,omitempty"`

	// Note that the elements returned by the expression are sometimes complex elements where logic is required to determine quite how to handle them
	Expression *string `json:"expression,omitempty"`

	// It may be possible for the search parameter to be used in jurisdictions other than those for which it was originally designed
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`

	// A modifier supported for the search parameter
	Modifier []SearchParameterModifier `json:"modifier,omitempty"`

	// Whether multiple parameters are allowed - e.g. more than one parameter with the same name
	MultipleAnd *bool `json:"multipleAnd,omitempty"`

	// Whether multiple values are allowed for each time the parameter exists
	MultipleOr *bool `json:"multipleOr,omitempty"`

	// The name is not expected to be globally unique
	Name string `json:"name"`

	// Usually an organization but may be an individual
	Publisher *string `json:"publisher,omitempty"`

	// This element does not describe the usage of the search parameter
	Purpose *string `json:"purpose,omitempty"`

	// Allows filtering of search parameters that are appropriate for use versus not
	Status SearchParameterStatus `json:"status"`

	// Types of resource (if a resource is referenced)
	Target []string `json:"target,omitempty"`

	// The type of value that a search parameter may contain, and how the content is interpreted
	Type SearchParameterType `json:"type"`

	// Can be a urn:uuid: or a urn:oid: but real http: addresses are preferred
	Url string `json:"url"`

	// When multiple useContexts are specified, there is no expectation that all or any of the contexts apply
	UseContext []common.UsageContext `json:"useContext,omitempty"`

	// There may be different search parameter instances that have the same identifier but different versions
	Version *string `json:"version,omitempty"`

	// Note that the elements returned by the XPath are sometimes complex elements where logic is required to determine quite how to handle them
	Xpath *string `json:"xpath,omitempty"`

	// How the search parameter relates to the set of elements returned by evaluating the xpath query
	XpathUsage *SearchParameterXpathUsage `json:"xpathUsage,omitempty"`
}

// SearchParameterStatus represents allows filtering of search parameters that are appropriate for use versus not
type SearchParameterStatus string

const (
	SearchParameterStatusDraft   SearchParameterStatus = "draft"
	SearchParameterStatusActive  SearchParameterStatus = "active"
	SearchParameterStatusRetired SearchParameterStatus = "retired"
	SearchParameterStatusUnknown SearchParameterStatus = "unknown"
)

// SearchParameterType represents the type of value that a search parameter may contain, and how the content is interpreted
type SearchParameterType string

const (
	SearchParameterTypeNumber    SearchParameterType = "number"
	SearchParameterTypeDate      SearchParameterType = "date"
	SearchParameterTypeString    SearchParameterType = "string"
	SearchParameterTypeToken     SearchParameterType = "token"
	SearchParameterTypeReference SearchParameterType = "reference"
	SearchParameterTypeComposite SearchParameterType = "composite"
	SearchParameterTypeQuantity  SearchParameterType = "quantity"
	SearchParameterTypeUri       SearchParameterType = "uri"
	SearchParameterTypeSpecial   SearchParameterType = "special"
)

// SearchParameterComparator represents if no comparators are listed, clients should not expect servers to support any comparators
type SearchParameterComparator string

const (
	SearchParameterComparatorEq SearchParameterComparator = "eq"
	SearchParameterComparatorNe SearchParameterComparator = "ne"
	SearchParameterComparatorGt SearchParameterComparator = "gt"
	SearchParameterComparatorLt SearchParameterComparator = "lt"
	SearchParameterComparatorGe SearchParameterComparator = "ge"
	SearchParameterComparatorLe SearchParameterComparator = "le"
	SearchParameterComparatorSa SearchParameterComparator = "sa"
	SearchParameterComparatorEb SearchParameterComparator = "eb"
	SearchParameterComparatorAp SearchParameterComparator = "ap"
)

// SearchParameterModifier represents a modifier supported for the search parameter
type SearchParameterModifier string

const (
	SearchParameterModifierMissing    SearchParameterModifier = "missing"
	SearchParameterModifierExact      SearchParameterModifier = "exact"
	SearchParameterModifierContains   SearchParameterModifier = "contains"
	SearchParameterModifierNot        SearchParameterModifier = "not"
	SearchParameterModifierText       SearchParameterModifier = "text"
	SearchParameterModifierIn         SearchParameterModifier = "in"
	SearchParameterModifierNotIn      SearchParameterModifier = "not-in"
	SearchParameterModifierBelow      SearchParameterModifier = "below"
	SearchParameterModifierAbove      SearchParameterModifier = "above"
	SearchParameterModifierType       SearchParameterModifier = "type"
	SearchParameterModifierIdentifier SearchParameterModifier = "identifier"
	SearchParameterModifierOfType     SearchParameterModifier = "ofType"
)

// SearchParameterXpathUsage represents how the search parameter relates to the set of elements returned by evaluating the xpath query
type SearchParameterXpathUsage string

const (
	SearchParameterXpathUsageNormal   SearchParameterXpathUsage = "normal"
	SearchParameterXpathUsagePhonetic SearchParameterXpathUsage = "phonetic"
	SearchParameterXpathUsageNearby   SearchParameterXpathUsage = "nearby"
	SearchParameterXpathUsageDistance SearchParameterXpathUsage = "distance"
	SearchParameterXpathUsageOther    SearchParameterXpathUsage = "other"
)

// SearchParameterComponent represents used to define the parts of a composite search parameter
type SearchParameterComponent struct {
	common.BackboneElement

	// The definition of the search parameter that describes this part
	Definition string `json:"definition"`

	// This expression overrides the expression in the definition and extracts the index values from the outcome of the composite expression
	Expression string `json:"expression"`
}
