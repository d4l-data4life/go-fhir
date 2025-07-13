// Package common contains base types and structures shared across FHIR versions
package common

import "time"

// Element is the base definition for all elements contained inside a resource.
// All elements, whether defined as a data type (like the ones in this specification)
// or as part of a resource definition, have this base content.
type Element struct {
	// Unique id for the element within a resource (for internal references)
	ID *string `json:"id,omitempty"`

	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
}

// BackboneElement is the base definition for all elements that are defined inside a resource
// but not those in a data type.
type BackboneElement struct {
	Element

	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
}

// Extension represents additional content defined by implementations.
type Extension struct {
	Element

	// identifies the meaning of the extension
	URL string `json:"url"`

	// Value can be one of many types - using interface{} for flexibility
	Value interface{} `json:"value,omitempty"`
}

// Reference is a reference from one resource to another.
type Reference struct {
	Element

	// Literal reference, Relative, internal or absolute URL
	Reference *string `json:"reference,omitempty"`

	// Type the reference refers to (e.g. "Patient")
	Type *string `json:"type,omitempty"`

	// Logical reference, when literal reference is not known
	Identifier *Identifier `json:"identifier,omitempty"`

	// Text alternative for the resource
	Display *string `json:"display,omitempty"`
}

// Identifier represents an identifier intended for computation.
type Identifier struct {
	Element

	// The purpose of this identifier
	Use *IdentifierUse `json:"use,omitempty"`

	// Description of identifier
	Type *CodeableConcept `json:"type,omitempty"`

	// The namespace for the identifier value
	System *string `json:"system,omitempty"`

	// The value that is unique
	Value *string `json:"value,omitempty"`

	// Time period when id is/was valid for use
	Period *Period `json:"period,omitempty"`

	// Organization that issued id (may be just text)
	Assigner *Reference `json:"assigner,omitempty"`
}

// IdentifierUse represents the purpose of an identifier
type IdentifierUse string

const (
	IdentifierUseUsual     IdentifierUse = "usual"
	IdentifierUseOfficial  IdentifierUse = "official"
	IdentifierUseTemp      IdentifierUse = "temp"
	IdentifierUseSecondary IdentifierUse = "secondary"
	IdentifierUseOld       IdentifierUse = "old"
)

// Period represents a time period defined by a start and end date/time.
type Period struct {
	Element

	// Start time with inclusive boundary
	Start *time.Time `json:"start,omitempty"`

	// End time with inclusive boundary, if not ongoing
	End *time.Time `json:"end,omitempty"`
}

// CodeableConcept represents a concept that may be defined by a formal reference
// to a terminology or ontology or may be provided by text.
type CodeableConcept struct {
	Element

	// Code defined by a terminology system
	Coding []Coding `json:"coding,omitempty"`

	// Plain text representation of the concept
	Text *string `json:"text,omitempty"`
}

// Coding represents a reference to a code defined by a terminology system.
type Coding struct {
	Element

	// Identity of the terminology system
	System *string `json:"system,omitempty"`

	// Version of the system - if relevant
	Version *string `json:"version,omitempty"`

	// Symbol in syntax defined by the system
	Code *string `json:"code,omitempty"`

	// Representation defined by the system
	Display *string `json:"display,omitempty"`

	// If this coding was chosen directly by the user
	UserSelected *bool `json:"userSelected,omitempty"`
}

// Quantity represents a measured amount (or an amount that can potentially be measured).
type Quantity struct {
	Element

	// Numerical value (with implicit precision)
	Value *float64 `json:"value,omitempty"`

	// < | <= | >= | > - how to understand the value
	Comparator *QuantityComparator `json:"comparator,omitempty"`

	// Unit representation
	Unit *string `json:"unit,omitempty"`

	// System that defines coded unit form
	System *string `json:"system,omitempty"`

	// Coded form of the unit
	Code *string `json:"code,omitempty"`
}

// QuantityComparator represents how to understand the quantity value
type QuantityComparator string

const (
	QuantityComparatorLess         QuantityComparator = "<"
	QuantityComparatorLessOrEqual  QuantityComparator = "<="
	QuantityComparatorGreater      QuantityComparator = ">"
	QuantityComparatorGreaterEqual QuantityComparator = ">="
)

// SampledData represents a series of measurements taken by a device
type SampledData struct {
	Element

	// The ConceptMap cannot define meanings for the codes 'E', 'U', or 'L'
	CodeMap *string `json:"codeMap,omitempty"`

	// The data may be missing if it is omitted for summarization purposes
	Data *string `json:"data,omitempty"`

	// If there is more than one dimension, the code for the type of data will define the meaning of the dimensions
	Dimensions int `json:"dimensions"`

	// A correction factor that is applied to the sampled data points before they are added to the origin
	Factor *float64 `json:"factor,omitempty"`

	// This is usually a whole number
	Interval *int `json:"interval,omitempty"`

	// The measurement unit in which the sample interval is expressed
	IntervalUnit string `json:"intervalUnit"`

	// The lower limit of detection of the measured points
	LowerLimit *float64 `json:"lowerLimit,omitempty"`

	// If offsets is present, the number of data points must be equal to the number of offsets multiplied by the dimensions
	Offsets *string `json:"offsets,omitempty"`

	// The base quantity that a measured value of zero represents
	Origin Quantity `json:"origin"`

	// The upper limit of detection of the measured points
	UpperLimit *float64 `json:"upperLimit,omitempty"`
}
