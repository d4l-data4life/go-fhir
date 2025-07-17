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
	Offsets *string `json:"offSets,omitempty"`

	// The base quantity that a measured value of zero represents
	Origin Quantity `json:"origin"`

	// The upper limit of detection of the measured points
	UpperLimit *float64 `json:"upperLimit,omitempty"`
}

// ContactDetail represents contact information for a person or organization
type ContactDetail struct {
	Element

	// The name of an individual to contact
	Name *string `json:"name,omitempty"`

	// The contact details for the individual (if a name was provided) or the organization
	Telecom []ContactPoint `json:"telecom,omitempty"`
}

// ContactPoint represents details for all kinds of technology-mediated contact points
type ContactPoint struct {
	Element

	// Telecommunications form for contact point - what communications system is required to make use of this contact
	System *ContactPointSystem `json:"system,omitempty"`

	// The actual contact point details, in a form that is meaningful to the designated communication system
	Value *string `json:"value,omitempty"`

	// Identifies the purpose for the contact point
	Use *ContactPointUse `json:"use,omitempty"`

	// Specifies a preferred order in which to use a set of contacts
	Rank *int `json:"rank,omitempty"`

	// Time period when the contact point was/is in use
	Period *Period `json:"period,omitempty"`
}

// ContactPointSystem represents the telecommunications form for contact point
type ContactPointSystem string

const (
	ContactPointSystemPhone ContactPointSystem = "phone"
	ContactPointSystemFax   ContactPointSystem = "fax"
	ContactPointSystemEmail ContactPointSystem = "email"
	ContactPointSystemPager ContactPointSystem = "pager"
	ContactPointSystemURL   ContactPointSystem = "url"
	ContactPointSystemSMS   ContactPointSystem = "sms"
	ContactPointSystemOther ContactPointSystem = "other"
)

// ContactPointUse represents the purpose for the contact point
type ContactPointUse string

const (
	ContactPointUseHome   ContactPointUse = "home"
	ContactPointUseWork   ContactPointUse = "work"
	ContactPointUseTemp   ContactPointUse = "temp"
	ContactPointUseOld    ContactPointUse = "old"
	ContactPointUseMobile ContactPointUse = "mobile"
)

// UsageContext represents the context in which a resource is intended to be used
type UsageContext struct {
	Element

	// A code that identifies the type of context being specified by this usage context
	Code Coding `json:"code"`

	// A value that defines the context specified in this context of use
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueQuantity        *Quantity        `json:"valueQuantity,omitempty"`
	ValueRange           *Range           `json:"valueRange,omitempty"`
	ValueReference       *Reference       `json:"valueReference,omitempty"`
}

// Range represents a set of ordered Quantities defined by a low and high limit
type Range struct {
	Element

	// The low limit. The boundary is inclusive
	Low *Quantity `json:"low,omitempty"`

	// The high limit. The boundary is inclusive
	High *Quantity `json:"high,omitempty"`
}

// ParameterDefinition represents the parameters to the module
type ParameterDefinition struct {
	Element

	// The name of the parameter used to allow access to the value of the parameter in evaluation contexts
	Name *string `json:"name,omitempty"`

	// The type of the parameter
	Type *string `json:"type,omitempty"`

	// Whether the parameter is input or output for the module
	Use *ParameterUse `json:"use,omitempty"`

	// The minimum number of times this parameter SHALL appear in the request or response
	Min *int `json:"min,omitempty"`

	// The maximum number of times this element is permitted to appear in the request or response
	Max *string `json:"max,omitempty"`

	// A brief discussion of what the parameter is for and how it is used by the module
	Documentation *string `json:"documentation,omitempty"`

	// If specified, this indicates a profile that the input data must conform to, or that the output data will conform to
	Profile *string `json:"profile,omitempty"`
}

// ParameterUse represents whether the parameter is input or output for the module
type ParameterUse string

const (
	ParameterUseIn  ParameterUse = "in"
	ParameterUseOut ParameterUse = "out"
)

// Contributor represents a contributor to the content of a knowledge asset
type Contributor struct {
	Element

	// The type of contributor
	Type ContributorType `json:"type"`

	// The name of the individual or organization responsible for some aspect of the content being contributed
	Name string `json:"name"`

	// Contact details to assist a user in finding and communicating with the contributor
	Contact []ContactDetail `json:"contact,omitempty"`
}

// ContributorType represents the type of contributor
type ContributorType string

const (
	ContributorTypeAuthor   ContributorType = "author"
	ContributorTypeEditor   ContributorType = "editor"
	ContributorTypeReviewer ContributorType = "reviewer"
	ContributorTypeEndorser ContributorType = "endorser"
)

// CodeableReference represents a reference to a resource (by instance) or a concept (by class)
type CodeableReference struct {
	Element

	// Reference to a concept (by class)
	Concept *CodeableConcept `json:"concept,omitempty"`

	// Reference to a resource instance
	Reference *Reference `json:"reference,omitempty"`
}

// Address represents an address expressed using postal conventions
type Address struct {
	Element

	// The purpose of this address
	Use *AddressUse `json:"use,omitempty"`

	// Distinguishes between physical addresses (those you can visit) and mailing addresses (e.g. PO Boxes and care-of addresses)
	Type *AddressType `json:"type,omitempty"`

	// A full text representation of the address
	Text *string `json:"text,omitempty"`

	// This component contains the house number, apartment number, street name, street direction, P.O. Box number, delivery hints, and similar address information
	Line []string `json:"line,omitempty"`

	// The name of the city, town, suburb, village or other community or delivery center
	City *string `json:"city,omitempty"`

	// Sub-unit of a country with limited sovereignty in a federally organized country
	District *string `json:"district,omitempty"`

	// Sub-unit of a country with limited sovereignty in a federally organized country
	State *string `json:"state,omitempty"`

	// A postal code designating a region defined by the postal service
	PostalCode *string `json:"postalCode,omitempty"`

	// Country - a nation as commonly understood or generally accepted
	Country *string `json:"country,omitempty"`

	// Time period when address was/is in use
	Period *Period `json:"period,omitempty"`
}

// AddressUse represents the purpose of an address
type AddressUse string

const (
	AddressUseHome    AddressUse = "home"
	AddressUseWork    AddressUse = "work"
	AddressUseTemp    AddressUse = "temp"
	AddressUseOld     AddressUse = "old"
	AddressUseBilling AddressUse = "billing"
)

// AddressType represents the type of address
type AddressType string

const (
	AddressTypePostal   AddressType = "postal"
	AddressTypePhysical AddressType = "physical"
	AddressTypeBoth     AddressType = "both"
)

// Annotation represents a text note which also contains information about who made the statement and when
type Annotation struct {
	Element

	// The individual responsible for making the annotation
	AuthorReference *Reference `json:"authorReference,omitempty"`
	AuthorString    *string    `json:"authorString,omitempty"`

	// The text of the annotation in markdown format
	Text string `json:"text"`

	// Indicates when this particular annotation was made
	Time *time.Time `json:"time,omitempty"`
}

// Attachment represents content in a base64format
type Attachment struct {
	Element

	// Identifies the type of the data in the attachment and allows a method to be chosen to interpret or render the data
	ContentType *string `json:"contentType,omitempty"`

	// The human language of the content
	Language *string `json:"language,omitempty"`

	// The actual data of the attachment - a sequence of bytes, base64 encoded
	Data *string `json:"data,omitempty"`

	// A location where the data can be accessed
	URL *string `json:"url,omitempty"`

	// The number of bytes of data that make up this attachment
	Size *int64 `json:"size,omitempty"`

	// The calculated hash of the data using SHA-1
	Hash *string `json:"hash,omitempty"`

	// A label or set of text to display in place of the data
	Title *string `json:"title,omitempty"`

	// The date that the attachment was first created
	Creation *time.Time `json:"creation,omitempty"`
}

// Count represents a measured amount
type Count struct {
	Element

	// The value of the measured amount
	Value *float64 `json:"value,omitempty"`

	// How the value should be understood and represented
	Comparator *QuantityComparator `json:"comparator,omitempty"`

	// A human-readable form of the value
	Unit *string `json:"unit,omitempty"`

	// The identification of the system that provides the coded form of the unit
	System *string `json:"system,omitempty"`

	// A computer processable form of the unit in some unit representation system
	Code *string `json:"code,omitempty"`
}

// Distance represents a length - a value with a unit that is a physical distance
type Distance struct {
	Element

	// The value of the measured amount
	Value *float64 `json:"value,omitempty"`

	// How the value should be understood and represented
	Comparator *QuantityComparator `json:"comparator,omitempty"`

	// A human-readable form of the value
	Unit *string `json:"unit,omitempty"`

	// The identification of the system that provides the coded form of the unit
	System *string `json:"system,omitempty"`

	// A computer processable form of the unit in some unit representation system
	Code *string `json:"code,omitempty"`
}

// HumanName represents a human's name with the ability to identify parts and usage
type HumanName struct {
	Element

	// Identifies the purpose for this name
	Use *HumanNameUse `json:"use,omitempty"`

	// The full name formatted for display
	Text *string `json:"text,omitempty"`

	// The part of a name that links to the genealogy
	Family *string `json:"family,omitempty"`

	// Given name
	Given []string `json:"given,omitempty"`

	// Part that comes before the name
	Prefix []string `json:"prefix,omitempty"`

	// Part that comes after the name
	Suffix []string `json:"suffix,omitempty"`

	// Time period when name was/is in use
	Period *Period `json:"period,omitempty"`
}

// HumanNameUse represents the purpose of a human name
type HumanNameUse string

const (
	HumanNameUseUsual     HumanNameUse = "usual"
	HumanNameUseOfficial  HumanNameUse = "official"
	HumanNameUseTemp      HumanNameUse = "temp"
	HumanNameUseNickname  HumanNameUse = "nickname"
	HumanNameUseAnonymous HumanNameUse = "anonymous"
	HumanNameUseOld       HumanNameUse = "old"
	HumanNameUseMaiden    HumanNameUse = "maiden"
)

// Money represents an amount of economic utility in some recognized currency
type Money struct {
	Element

	// Numerical value (with implicit precision)
	Value *float64 `json:"value,omitempty"`

	// ISO 4217 Currency Code
	Currency *string `json:"currency,omitempty"`
}

// Ratio represents a relationship of two Quantity values expressed as a numerator and a denominator
type Ratio struct {
	Element

	// The value of the numerator
	Numerator *Quantity `json:"numerator,omitempty"`

	// The value of the denominator
	Denominator *Quantity `json:"denominator,omitempty"`
}

// Signature represents a digital signature along with supporting context
type Signature struct {
	Element

	// An indication of the reason that the entity signed this document
	Type []Coding `json:"type,omitempty"`

	// When the digital signature was signed
	When *time.Time `json:"when,omitempty"`

	// A reference to an application-usable description of the identity that signed
	Who *Reference `json:"who,omitempty"`

	// A reference to an application-usable description of the identity that is represented by the signature
	OnBehalfOf *Reference `json:"onBehalfOf,omitempty"`

	// A mime type that indicates the technical format of the target resources signed by the signature
	TargetFormat *string `json:"targetFormat,omitempty"`

	// A mime type that indicates the technical format of the signature
	SigFormat *string `json:"sigFormat,omitempty"`

	// The base64 encoding of the Signature content
	Data *string `json:"data,omitempty"`
}

// Timing represents a timing schedule
type Timing struct {
	Element

	// Identifies what events are expected to occur
	Event []time.Time `json:"event,omitempty"`

	// A set of rules that describe when the event should occur
	Repeat *TimingRepeat `json:"repeat,omitempty"`

	// A code for the timing schedule
	Code *CodeableConcept `json:"code,omitempty"`
}

// TimingRepeat represents a timing schedule
type TimingRepeat struct {
	Element

	// Length/Range of lengths, or (Start and/or end) limits
	BoundsDuration *Duration `json:"boundsDuration,omitempty"`
	BoundsRange    *Range    `json:"boundsRange,omitempty"`
	BoundsPeriod   *Period   `json:"boundsPeriod,omitempty"`

	// Number of times to repeat
	Count *int `json:"count,omitempty"`

	// Maximum number of times to repeat
	CountMax *int `json:"countMax,omitempty"`

	// How long when it happens
	Duration *float64 `json:"duration,omitempty"`

	// How long when it happens (Max)
	DurationMax *float64 `json:"durationMax,omitempty"`

	// s | min | h | d | wk | mo | a - unit of time (UCUM)
	DurationUnit *string `json:"durationUnit,omitempty"`

	// Event occurs frequency times per period
	Frequency *int `json:"frequency,omitempty"`

	// Event occurs up to frequencyMax times per period
	FrequencyMax *int `json:"frequencyMax,omitempty"`

	// Event occurs frequency times per period
	Period *float64 `json:"period,omitempty"`

	// Upper limit of period (3-4 hours)
	PeriodMax *float64 `json:"periodMax,omitempty"`

	// s | min | h | d | wk | mo | a - unit of time (UCUM)
	PeriodUnit *string `json:"periodUnit,omitempty"`

	// mon | tue | wed | thu | fri | sat | sun
	DayOfWeek []string `json:"dayOfWeek,omitempty"`

	// Time of day for action
	TimeOfDay []string `json:"timeOfDay,omitempty"`

	// Code for time period of occurrence
	When []string `json:"when,omitempty"`

	// Minutes from event (before or after)
	Offset *int `json:"offset,omitempty"`
}

// Duration represents a length of time
type Duration struct {
	Element

	// The value of the measured amount
	Value *float64 `json:"value,omitempty"`

	// How the value should be understood and represented
	Comparator *QuantityComparator `json:"comparator,omitempty"`

	// A human-readable form of the value
	Unit *string `json:"unit,omitempty"`

	// The identification of the system that provides the coded form of the unit
	System *string `json:"system,omitempty"`

	// A computer processable form of the unit in some unit representation system
	Code *string `json:"code,omitempty"`
}

// DataRequirement represents a required data item for evaluation
type DataRequirement struct {
	Element

	// The type of the required data
	Type string `json:"type"`

	// The profile of the required data
	Profile []string `json:"profile,omitempty"`

	// The subject of a data requirement
	SubjectCodeableConcept *CodeableConcept `json:"subjectCodeableConcept,omitempty"`
	SubjectReference       *Reference       `json:"subjectReference,omitempty"`

	// Indicates that specific elements of the type are referenced by the knowledge module
	MustSupport []string `json:"mustSupport,omitempty"`

	// Code filters specify additional constraints on the data
	CodeFilter []DataRequirementCodeFilter `json:"codeFilter,omitempty"`

	// Date filters specify additional constraints on the data
	DateFilter []DataRequirementDateFilter `json:"dateFilter,omitempty"`

	// Value filters specify additional constraints on the data
	ValueFilter []DataRequirementValueFilter `json:"valueFilter,omitempty"`

	// Specifies the order of the results to be returned
	Sort []DataRequirementSort `json:"sort,omitempty"`

	// The maximum number of results to return
	Limit *int `json:"limit,omitempty"`
}

// DataRequirementCodeFilter represents a code filter for data requirements
type DataRequirementCodeFilter struct {
	Element

	// A code-valued attribute to filter on
	Path *string `json:"path,omitempty"`

	// A token that defines the search parameter
	SearchParam *string `json:"searchParam,omitempty"`

	// Valueset for the filter
	ValueSet *string `json:"valueSet,omitempty"`

	// What code is expected
	Code []Coding `json:"code,omitempty"`
}

// DataRequirementDateFilter represents a date filter for data requirements
type DataRequirementDateFilter struct {
	Element

	// A date-valued attribute to filter on
	Path *string `json:"path,omitempty"`

	// A date parameter that refers to a search parameter type on this type
	SearchParam *string `json:"searchParam,omitempty"`

	// The value of the filter, as a Period, DateTime, or Duration value
	ValueDateTime *time.Time `json:"valueDateTime,omitempty"`
	ValuePeriod   *Period    `json:"valuePeriod,omitempty"`
	ValueDuration *Duration  `json:"valueDuration,omitempty"`
}

// DataRequirementValueFilter represents a value filter for data requirements
type DataRequirementValueFilter struct {
	Element

	// The attribute of the sort
	Path *string `json:"path,omitempty"`

	// The search parameter for the filter
	SearchParam *string `json:"searchParam,omitempty"`

	// The comparator to be used to determine whether the value is acceptable
	Comparator *string `json:"comparator,omitempty"`

	// The value of the filter
	Value *string `json:"value,omitempty"`
}

// DataRequirementSort represents a sort specification for data requirements
type DataRequirementSort struct {
	Element

	// The attribute of the sort
	Path string `json:"path"`

	// The direction of the sort
	Direction DataRequirementSortDirection `json:"direction"`
}

// DataRequirementSortDirection represents the direction of sorting
type DataRequirementSortDirection string

const (
	DataRequirementSortDirectionAscending  DataRequirementSortDirection = "ascending"
	DataRequirementSortDirectionDescending DataRequirementSortDirection = "descending"
)

// Expression represents an expression that can be used to generate a value
type Expression struct {
	Element

	// A brief, natural language description of the condition that effectively communicates the intended semantics
	Description *string `json:"description,omitempty"`

	// A short name assigned to the expression to allow for multiple reuse of the expression in the context where it is defined
	Name *string `json:"name,omitempty"`

	// The media type of the language for the expression
	Language *string `json:"language,omitempty"`

	// An expression in the specified language
	Expression *string `json:"expression,omitempty"`

	// A URI that defines where the expression is found
	Reference *string `json:"reference,omitempty"`
}

// RelatedArtifact represents related artifacts such as additional documentation, justification, or bibliographic references
type RelatedArtifact struct {
	Element

	// The type of relationship to the related artifact
	Type RelatedArtifactType `json:"type"`

	// A short label that can be used to reference the citation from elsewhere in the containing artifact
	Label *string `json:"label,omitempty"`

	// A brief description of the document or knowledge resource being referenced
	Display *string `json:"display,omitempty"`

	// A bibliographic citation for the related artifact
	Citation *string `json:"citation,omitempty"`

	// The document being referenced
	URL *string `json:"url,omitempty"`

	// The related artifact, such as a library, value set, profile, or other knowledge resource
	Document *Attachment `json:"document,omitempty"`

	// The related resource
	Resource *string `json:"resource,omitempty"`
}

// RelatedArtifactType represents the type of relationship to the related artifact
type RelatedArtifactType string

const (
	RelatedArtifactTypeDocumentation RelatedArtifactType = "documentation"
	RelatedArtifactTypeJustification RelatedArtifactType = "justification"
	RelatedArtifactTypeCitation      RelatedArtifactType = "citation"
	RelatedArtifactTypePredecessor   RelatedArtifactType = "predecessor"
	RelatedArtifactTypeSuccessor     RelatedArtifactType = "successor"
	RelatedArtifactTypeDerivedFrom   RelatedArtifactType = "derived-from"
	RelatedArtifactTypeDependsOn     RelatedArtifactType = "depends-on"
	RelatedArtifactTypeComposedOf    RelatedArtifactType = "composed-of"
)

// TriggerDefinition represents a description of a triggering event
type TriggerDefinition struct {
	Element

	// The type of triggering event
	Type TriggerDefinitionType `json:"type"`

	// A formal name for the event
	Name *string `json:"name,omitempty"`

	// The timing of the event (if this is a periodic trigger)
	TimingTiming    *Timing    `json:"timingTiming,omitempty"`
	TimingReference *Reference `json:"timingReference,omitempty"`
	TimingDate      *time.Time `json:"timingDate,omitempty"`
	TimingDateTime  *time.Time `json:"timingDateTime,omitempty"`

	// The triggering data of the event (if this is a data trigger)
	Data []DataRequirement `json:"data,omitempty"`

	// A boolean-valued expression that is evaluated in the context of the container of the trigger definition
	Condition *Expression `json:"condition,omitempty"`
}

// TriggerDefinitionType represents the type of triggering event
type TriggerDefinitionType string

const (
	TriggerDefinitionTypeNamedEvent      TriggerDefinitionType = "named-event"
	TriggerDefinitionTypePeriodic        TriggerDefinitionType = "periodic"
	TriggerDefinitionTypeDataChanged     TriggerDefinitionType = "data-changed"
	TriggerDefinitionTypeDataAdded       TriggerDefinitionType = "data-added"
	TriggerDefinitionTypeDataModified    TriggerDefinitionType = "data-modified"
	TriggerDefinitionTypeDataRemoved     TriggerDefinitionType = "data-removed"
	TriggerDefinitionTypeDataAccessed    TriggerDefinitionType = "data-accessed"
	TriggerDefinitionTypeDataAccessEnded TriggerDefinitionType = "data-access-ended"
)

// Dosage represents how medication should be taken
type Dosage struct {
	Element

	// Free text dosage instructions
	Text *string `json:"text,omitempty"`

	// Patient or consumer oriented instructions
	PatientInstruction *string `json:"patientInstruction,omitempty"`

	// Timing for the dosage
	Timing *Timing `json:"timing,omitempty"`

	// Body site to administer to
	Site *CodeableConcept `json:"site,omitempty"`

	// How drug should enter body
	Route *CodeableConcept `json:"route,omitempty"`

	// Technique for administering medication
	Method *CodeableConcept `json:"method,omitempty"`

	// Amount of medication per dose
	DoseAndRate []DosageDoseAndRate `json:"doseAndRate,omitempty"`
}

// DosageDoseAndRate represents dosage dose and rate information
type DosageDoseAndRate struct {
	Element

	// The kind of dose or rate specified
	Type *CodeableConcept `json:"type,omitempty"`

	// Amount of medication per dose
	DoseRange    *Range    `json:"doseRange,omitempty"`
	DoseQuantity *Quantity `json:"doseQuantity,omitempty"`

	// Amount of medication per unit of time
	RateRatio    *Ratio    `json:"rateRatio,omitempty"`
	RateRange    *Range    `json:"rateRange,omitempty"`
	RateQuantity *Quantity `json:"rateQuantity,omitempty"`
}

// Meta represents metadata about a resource
type Meta struct {
	Element

	// Version specific identifier
	VersionID *string `json:"versionId,omitempty"`

	// When the resource version last changed
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`

	// Identifies where the resource comes from
	Source *string `json:"source,omitempty"`

	// Profiles this resource claims to conform to
	Profile []string `json:"profile,omitempty"`

	// Security Labels applied to this resource
	Security []Coding `json:"security,omitempty"`

	// Tags applied to this resource
	Tag []Coding `json:"tag,omitempty"`
}

// Age represents a duration of time during which an organism (or a process) has existed
type Age Quantity

// ElementDefinition represents the formal definition of an element in a resource or data type
type ElementDefinition struct {
	BackboneElement

	// The path of the element (see the ElementDefinition.path property)
	Path string `json:"path"`

	// How many times this element is permitted to appear in the instance
	Min *int `json:"min,omitempty"`

	// The maximum number of times this element is permitted to appear in the instance
	Max *string `json:"max,omitempty"`

	// The data type or resource that the value of this element is permitted to be
	Type []ElementDefinitionType `json:"type,omitempty"`

	// A concise description of what this element means (used for alternate purpose, etc.)
	Short *string `json:"short,omitempty"`

	// Provides a complete explanation of the meaning of the data element for human readability
	Definition *string `json:"definition,omitempty"`

	// Explanatory notes and implementation guidance about the data element
	Comment *string `json:"comment,omitempty"`

	// The name of this element definition slice, when slicing is working
	SliceName *string `json:"sliceName,omitempty"`

	// Whether this element definition is constraining a slicer
	SliceIsConstraining *bool `json:"sliceIsConstraining,omitempty"`

	// A sample value for this element demonstrating the type of information that would typically be captured
	Example []ElementDefinitionExample `json:"example,omitempty"`

	// Formal constraints such as co-occurrence and other constraints that can be computationally evaluated
	Constraint []ElementDefinitionConstraint `json:"constraint,omitempty"`

	// Binds to a value set if this element is coded (code, Coding, CodeableConcept, Quantity), or the data types (string, uri)
	Binding *ElementDefinitionBinding `json:"binding,omitempty"`

	// Identifies a concept from an external specification that roughly corresponds to this element
	Mapping []ElementDefinitionMapping `json:"mapping,omitempty"`
}

// ElementDefinitionType represents the data type or resource that the value of this element is permitted to be
type ElementDefinitionType struct {
	Element

	// URL of the type
	Code string `json:"code"`

	// Profile (StructureDefinition) to apply to this element
	Profile []string `json:"profile,omitempty"`

	// Profile (StructureDefinition) to apply to the target resource (for a reference)
	TargetProfile []string `json:"targetProfile,omitempty"`

	// If the type is a reference to another resource, how the resource is or can be aggregated
	Aggregation []ElementDefinitionTypeAggregation `json:"aggregation,omitempty"`

	// Whether this reference needs to be version specific or version independent, or whether either can be used
	Versioning *ElementDefinitionTypeVersioning `json:"versioning,omitempty"`
}

// ElementDefinitionTypeAggregation represents if the type is a reference to another resource, how the resource is or can be aggregated
type ElementDefinitionTypeAggregation string

const (
	ElementDefinitionTypeAggregationContained  ElementDefinitionTypeAggregation = "contained"
	ElementDefinitionTypeAggregationReferenced ElementDefinitionTypeAggregation = "referenced"
	ElementDefinitionTypeAggregationBundled    ElementDefinitionTypeAggregation = "bundled"
)

// ElementDefinitionTypeVersioning represents whether this reference needs to be version specific or version independent
type ElementDefinitionTypeVersioning string

const (
	ElementDefinitionTypeVersioningEither      ElementDefinitionTypeVersioning = "either"
	ElementDefinitionTypeVersioningIndependent ElementDefinitionTypeVersioning = "independent"
	ElementDefinitionTypeVersioningSpecific    ElementDefinitionTypeVersioning = "specific"
)

// ElementDefinitionExample represents a sample value for this element demonstrating the type of information that would typically be captured
type ElementDefinitionExample struct {
	Element

	// Describes the purpose of this example
	Label string `json:"label"`

	// The actual value for the element (must be one of the allowed types)
	Value interface{} `json:"value"`
}

// ElementDefinitionConstraint represents formal constraints such as co-occurrence and other constraints that can be computationally evaluated
type ElementDefinitionConstraint struct {
	Element

	// Allows identification of which elements have their cardinalities impacted by the constraint
	Key string `json:"key"`

	// Description of why this constraint is necessary or appropriate
	Requirements *string `json:"requirements,omitempty"`

	// Identifies the impact constraint violation has on the conformance of the instance
	Severity ElementDefinitionConstraintSeverity `json:"severity"`

	// Text that can be used to describe the constraint in messages identifying that the constraint has been violated
	Human string `json:"human"`

	// A FHIRPath expression of constraint that can be executed to see if this constraint is met
	Expression *string `json:"expression,omitempty"`

	// An XPath expression of constraint that can be executed to see if this constraint is met
	Xpath *string `json:"xpath,omitempty"`
}

// ElementDefinitionConstraintSeverity represents identifies the impact constraint violation has on the conformance of the instance
type ElementDefinitionConstraintSeverity string

const (
	ElementDefinitionConstraintSeverityError   ElementDefinitionConstraintSeverity = "error"
	ElementDefinitionConstraintSeverityWarning ElementDefinitionConstraintSeverity = "warning"
)

// ElementDefinitionBinding represents binds to a value set if this element is coded
type ElementDefinitionBinding struct {
	Element

	// Indicates the degree of conformance expectations associated with this binding
	Strength ElementDefinitionBindingStrength `json:"strength"`

	// Describes the intended use of this binding
	Description *string `json:"description,omitempty"`

	// Refers to the value set that identifies the set of codes the binding refers to
	ValueSet *string `json:"valueSet,omitempty"`
}

// ElementDefinitionBindingStrength represents indicates the degree of conformance expectations associated with this binding
type ElementDefinitionBindingStrength string

const (
	ElementDefinitionBindingStrengthRequired   ElementDefinitionBindingStrength = "required"
	ElementDefinitionBindingStrengthExtensible ElementDefinitionBindingStrength = "extensible"
	ElementDefinitionBindingStrengthPreferred  ElementDefinitionBindingStrength = "preferred"
	ElementDefinitionBindingStrengthExample    ElementDefinitionBindingStrength = "example"
)

// ElementDefinitionMapping represents identifies a concept from an external specification that roughly corresponds to this element
type ElementDefinitionMapping struct {
	Element

	// An internal reference to the definition of a mapping
	Identity string `json:"identity"`

	// Identifies the computable language in which mapping.map is expressed
	Language *string `json:"language,omitempty"`

	// Expresses what part of the target specification corresponds to this element
	Map string `json:"map"`

	// Comments that provide information about the mapping or its use
	Comment *string `json:"comment,omitempty"`
}
