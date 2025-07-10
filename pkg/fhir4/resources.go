// Package fhir4 contains FHIR R4 (version 4.0.1) resource definitions
package fhir4

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// Resource is the base definition for all resources
type Resource struct {
	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"`

	// Logical id of this artifact
	ID *string `json:"id,omitempty"`

	// Metadata about the resource
	Meta *Meta `json:"meta,omitempty"`

	// A set of rules under which this content was created
	ImplicitRules *string `json:"implicitRules,omitempty"`

	// Language of the resource content
	Language *string `json:"language,omitempty"`
}

// DomainResource is a resource that includes narrative, extensions, and contained resources
type DomainResource struct {
	Resource

	// Contained, inline Resources
	Contained []interface{} `json:"contained,omitempty"`

	// Additional content defined by implementations
	Extension []common.Extension `json:"extension,omitempty"`

	// Extensions that cannot be ignored
	ModifierExtension []common.Extension `json:"modifierExtension,omitempty"`

	// Text summary of the resource, for human interpretation
	Text *Narrative `json:"text,omitempty"`
}

// Meta represents metadata about a resource
type Meta struct {
	common.Element

	// Version specific identifier
	VersionID *string `json:"versionId,omitempty"`

	// When the resource version last changed
	LastUpdated *string `json:"lastUpdated,omitempty"`

	// Identifies where the resource comes from
	Source *string `json:"source,omitempty"`

	// Profiles this resource claims to conform to
	Profile []string `json:"profile,omitempty"`

	// Security Labels applied to this resource
	Security []common.Coding `json:"security,omitempty"`

	// Tags applied to this resource
	Tag []common.Coding `json:"tag,omitempty"`
}

// Narrative represents human-readable summary of the resource
type Narrative struct {
	common.Element

	// generated | extensions | additional | empty
	Status NarrativeStatus `json:"status"`

	// Limited xhtml content
	Div string `json:"div"`
}

// NarrativeStatus represents the status of a narrative
type NarrativeStatus string

const (
	NarrativeStatusGenerated  NarrativeStatus = "generated"
	NarrativeStatusExtensions NarrativeStatus = "extensions"
	NarrativeStatusAdditional NarrativeStatus = "additional"
	NarrativeStatusEmpty      NarrativeStatus = "empty"
)

// Patient represents information about an individual or animal receiving care
type Patient struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Patient"

	// Whether this patient record is in active use
	Active *bool `json:"active,omitempty"`

	// An identifier for this patient
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// A name associated with the patient
	Name []HumanName `json:"name,omitempty"`

	// A contact detail for the individual
	Telecom []ContactPoint `json:"telecom,omitempty"`

	// male | female | other | unknown
	Gender *AdministrativeGender `json:"gender,omitempty"`

	// The date of birth for the individual
	BirthDate *string `json:"birthDate,omitempty"`

	// Indicates if the individual is deceased or not
	DeceasedBoolean  *bool   `json:"deceasedBoolean,omitempty"`
	DeceasedDateTime *string `json:"deceasedDateTime,omitempty"`

	// An address for the individual
	Address []Address `json:"address,omitempty"`

	// Marital (civil) status of a patient
	MaritalStatus *common.CodeableConcept `json:"maritalStatus,omitempty"`

	// Whether patient is part of a multiple birth
	MultipleBirthBoolean *bool `json:"multipleBirthBoolean,omitempty"`
	MultipleBirthInteger *int  `json:"multipleBirthInteger,omitempty"`

	// Image of the patient
	Photo []Attachment `json:"photo,omitempty"`

	// A contact party (e.g. guardian, partner, friend) for the patient
	Contact []PatientContact `json:"contact,omitempty"`

	// A language which may be used to communicate with the patient
	Communication []PatientCommunication `json:"communication,omitempty"`

	// Patient's nominated primary care provider
	GeneralPractitioner []common.Reference `json:"generalPractitioner,omitempty"`

	// Organization that is the custodian of the patient record
	ManagingOrganization *common.Reference `json:"managingOrganization,omitempty"`

	// Link to another patient resource that concerns the same actual person
	Link []PatientLink `json:"link,omitempty"`
}

// HumanName represents a human name
type HumanName struct {
	common.Element

	// usual | official | temp | nickname | anonymous | old | maiden
	Use *NameUse `json:"use,omitempty"`

	// Text representation of the full name
	Text *string `json:"text,omitempty"`

	// Family name (often called 'Surname')
	Family *string `json:"family,omitempty"`

	// Given names (not always 'first'). Includes middle names
	Given []string `json:"given,omitempty"`

	// Parts that come before the name
	Prefix []string `json:"prefix,omitempty"`

	// Parts that come after the name
	Suffix []string `json:"suffix,omitempty"`

	// Time period when name was/is in use
	Period *common.Period `json:"period,omitempty"`
}

// NameUse represents the use of a human name
type NameUse string

const (
	NameUseUsual     NameUse = "usual"
	NameUseOfficial  NameUse = "official"
	NameUseTemp      NameUse = "temp"
	NameUseNickname  NameUse = "nickname"
	NameUseAnonymous NameUse = "anonymous"
	NameUseOld       NameUse = "old"
	NameUseMaiden    NameUse = "maiden"
)

// AdministrativeGender represents administrative gender codes
type AdministrativeGender string

const (
	AdministrativeGenderMale    AdministrativeGender = "male"
	AdministrativeGenderFemale  AdministrativeGender = "female"
	AdministrativeGenderOther   AdministrativeGender = "other"
	AdministrativeGenderUnknown AdministrativeGender = "unknown"
)

// PatientContact represents a contact party for the patient
type PatientContact struct {
	common.BackboneElement

	// The kind of relationship
	Relationship []common.CodeableConcept `json:"relationship,omitempty"`

	// A name associated with the contact person
	Name *HumanName `json:"name,omitempty"`

	// A contact detail for the person
	Telecom []ContactPoint `json:"telecom,omitempty"`

	// Address for the contact person
	Address *Address `json:"address,omitempty"`

	// male | female | other | unknown
	Gender *AdministrativeGender `json:"gender,omitempty"`

	// Organization that is associated with the contact
	Organization *common.Reference `json:"organization,omitempty"`

	// The period during which this contact person or organization is valid
	Period *common.Period `json:"period,omitempty"`
}

// PatientCommunication represents a language which may be used to communicate with the patient
type PatientCommunication struct {
	common.BackboneElement

	// The language which can be used to communicate with the patient
	Language common.CodeableConcept `json:"language"`

	// Language preference indicator
	Preferred *bool `json:"preferred,omitempty"`
}

// PatientLink represents a link to another patient resource that concerns the same actual person
type PatientLink struct {
	common.BackboneElement

	// The other patient resource that the link refers to
	Other common.Reference `json:"other"`

	// replaced-by | replaces | refer | seealso
	Type PatientLinkType `json:"type"`
}

// PatientLinkType represents the type of link between this patient resource and another patient resource
type PatientLinkType string

const (
	PatientLinkTypeReplacedBy PatientLinkType = "replaced-by"
	PatientLinkTypeReplaces   PatientLinkType = "replaces"
	PatientLinkTypeRefer      PatientLinkType = "refer"
	PatientLinkTypeSeeAlso    PatientLinkType = "seealso"
)

// Observation represents measurements and simple assertions made about a patient
type Observation struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Observation"

	// Business Identifier for observation
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Fulfills plan, proposal or order
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// Part of referenced event
	PartOf []common.Reference `json:"partOf,omitempty"`

	// registered | preliminary | final | amended | corrected | cancelled | entered-in-error | unknown
	Status ObservationStatus `json:"status"`

	// Classification of type of observation
	Category []common.CodeableConcept `json:"category,omitempty"`

	// Type of observation (code / type)
	Code common.CodeableConcept `json:"code"`

	// Who and/or what the observation is about
	Subject *common.Reference `json:"subject,omitempty"`

	// What the observation is about, when it is not about the subject of record
	Focus []common.Reference `json:"focus,omitempty"`

	// Healthcare event during which this observation is made
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Clinically relevant time/time-period for observation
	EffectiveDateTime *string        `json:"effectiveDateTime,omitempty"`
	EffectivePeriod   *common.Period `json:"effectivePeriod,omitempty"`
	EffectiveTiming   *Timing        `json:"effectiveTiming,omitempty"`
	EffectiveInstant  *string        `json:"effectiveInstant,omitempty"`

	// Date/Time this version was made available
	Issued *string `json:"issued,omitempty"`

	// Who is responsible for the observation
	Performer []common.Reference `json:"performer,omitempty"`

	// Actual result
	ValueQuantity        *common.Quantity        `json:"valueQuantity,omitempty"`
	ValueCodeableConcept *common.CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueString          *string                 `json:"valueString,omitempty"`
	ValueBoolean         *bool                   `json:"valueBoolean,omitempty"`
	ValueInteger         *int                    `json:"valueInteger,omitempty"`
	ValueRange           *Range                  `json:"valueRange,omitempty"`
	ValueRatio           *Ratio                  `json:"valueRatio,omitempty"`
	ValueSampledData     *SampledData            `json:"valueSampledData,omitempty"`
	ValueTime            *string                 `json:"valueTime,omitempty"`
	ValueDateTime        *string                 `json:"valueDateTime,omitempty"`
	ValuePeriod          *common.Period          `json:"valuePeriod,omitempty"`

	// Why the result is missing
	DataAbsentReason *common.CodeableConcept `json:"dataAbsentReason,omitempty"`

	// High, low, normal, etc.
	Interpretation []common.CodeableConcept `json:"interpretation,omitempty"`

	// Comments about the observation
	Note []Annotation `json:"note,omitempty"`

	// Observed body part
	BodySite *common.CodeableConcept `json:"bodySite,omitempty"`

	// How it was done
	Method *common.CodeableConcept `json:"method,omitempty"`

	// Specimen used for this observation
	Specimen *common.Reference `json:"specimen,omitempty"`

	// (Measurement) Device
	Device *common.Reference `json:"device,omitempty"`

	// Provides guide for interpretation
	ReferenceRange []ObservationReferenceRange `json:"referenceRange,omitempty"`

	// Related resource that belongs to the Observation group
	HasMember []common.Reference `json:"hasMember,omitempty"`

	// Related measurements the observation is made from
	DerivedFrom []common.Reference `json:"derivedFrom,omitempty"`

	// Component results
	Component []ObservationComponent `json:"component,omitempty"`
}

// ObservationStatus represents the status of the observation
type ObservationStatus string

const (
	ObservationStatusRegistered     ObservationStatus = "registered"
	ObservationStatusPreliminary    ObservationStatus = "preliminary"
	ObservationStatusFinal          ObservationStatus = "final"
	ObservationStatusAmended        ObservationStatus = "amended"
	ObservationStatusCorrected      ObservationStatus = "corrected"
	ObservationStatusCancelled      ObservationStatus = "cancelled"
	ObservationStatusEnteredInError ObservationStatus = "entered-in-error"
	ObservationStatusUnknown        ObservationStatus = "unknown"
)

// ObservationReferenceRange provides guide for interpretation of component result
type ObservationReferenceRange struct {
	common.BackboneElement

	// Low Range, if relevant
	Low *common.Quantity `json:"low,omitempty"`

	// High Range, if relevant
	High *common.Quantity `json:"high,omitempty"`

	// Reference range qualifier
	Type *common.CodeableConcept `json:"type,omitempty"`

	// Reference range population
	AppliesTo []common.CodeableConcept `json:"appliesTo,omitempty"`

	// Applicable age range, if relevant
	Age *Range `json:"age,omitempty"`

	// Text based reference range in an observation
	Text *string `json:"text,omitempty"`
}

// ObservationComponent represents component results
type ObservationComponent struct {
	common.BackboneElement

	// Type of component observation (code / type)
	Code common.CodeableConcept `json:"code"`

	// Actual component result
	ValueQuantity        *common.Quantity        `json:"valueQuantity,omitempty"`
	ValueCodeableConcept *common.CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueString          *string                 `json:"valueString,omitempty"`
	ValueBoolean         *bool                   `json:"valueBoolean,omitempty"`
	ValueInteger         *int                    `json:"valueInteger,omitempty"`
	ValueRange           *Range                  `json:"valueRange,omitempty"`
	ValueRatio           *Ratio                  `json:"valueRatio,omitempty"`
	ValueSampledData     *SampledData            `json:"valueSampledData,omitempty"`
	ValueTime            *string                 `json:"valueTime,omitempty"`
	ValueDateTime        *string                 `json:"valueDateTime,omitempty"`
	ValuePeriod          *common.Period          `json:"valuePeriod,omitempty"`

	// Why the component result is missing
	DataAbsentReason *common.CodeableConcept `json:"dataAbsentReason,omitempty"`

	// High, low, normal, etc.
	Interpretation []common.CodeableConcept `json:"interpretation,omitempty"`

	// Provides guide for interpretation of component result
	ReferenceRange []ObservationReferenceRange `json:"referenceRange,omitempty"`
}

// Timing represents a timing schedule that specifies an event that may occur multiple times
type Timing struct {
	common.BackboneElement

	// When the event occurs
	Event []string `json:"event,omitempty"`

	// When the event is to occur
	Repeat *TimingRepeat `json:"repeat,omitempty"`

	// BID | TID | QID | AM | PM | QD | QOD | +
	Code *common.CodeableConcept `json:"code,omitempty"`
}

// TimingRepeat represents when the event is to occur
type TimingRepeat struct {
	common.Element

	// Length/Range of lengths, or (Start and/or end) limits
	BoundsDuration *Duration      `json:"boundsDuration,omitempty"`
	BoundsRange    *Range         `json:"boundsRange,omitempty"`
	BoundsPeriod   *common.Period `json:"boundsPeriod,omitempty"`

	// Number of times to repeat
	Count    *int `json:"count,omitempty"`
	CountMax *int `json:"countMax,omitempty"`

	// How long when it happens
	Duration     *float64     `json:"duration,omitempty"`
	DurationMax  *float64     `json:"durationMax,omitempty"`
	DurationUnit *UnitsOfTime `json:"durationUnit,omitempty"`

	// Event occurs frequency times per period
	Frequency    *int `json:"frequency,omitempty"`
	FrequencyMax *int `json:"frequencyMax,omitempty"`

	// Event occurs frequency times per period
	Period     *float64     `json:"period,omitempty"`
	PeriodMax  *float64     `json:"periodMax,omitempty"`
	PeriodUnit *UnitsOfTime `json:"periodUnit,omitempty"`

	// If one or more days of week is provided, then the action happens only on the specified day(s)
	DayOfWeek []DaysOfWeek `json:"dayOfWeek,omitempty"`

	// Time of day for action
	TimeOfDay []string `json:"timeOfDay,omitempty"`

	// Code for time period of occurrence
	When []EventTiming `json:"when,omitempty"`

	// Minutes from event (before or after)
	Offset *int `json:"offset,omitempty"`
}

// UnitsOfTime represents units of time
type UnitsOfTime string

const (
	UnitsOfTimeS   UnitsOfTime = "s"   // seconds
	UnitsOfTimeMin UnitsOfTime = "min" // minutes
	UnitsOfTimeH   UnitsOfTime = "h"   // hours
	UnitsOfTimeD   UnitsOfTime = "d"   // days
	UnitsOfTimeWk  UnitsOfTime = "wk"  // weeks
	UnitsOfTimeMo  UnitsOfTime = "mo"  // months
	UnitsOfTimeA   UnitsOfTime = "a"   // years
)

// DaysOfWeek represents days of the week
type DaysOfWeek string

const (
	DaysOfWeekMon DaysOfWeek = "mon"
	DaysOfWeekTue DaysOfWeek = "tue"
	DaysOfWeekWed DaysOfWeek = "wed"
	DaysOfWeekThu DaysOfWeek = "thu"
	DaysOfWeekFri DaysOfWeek = "fri"
	DaysOfWeekSat DaysOfWeek = "sat"
	DaysOfWeekSun DaysOfWeek = "sun"
)

// EventTiming represents real world events for timing schedules
type EventTiming string

const (
	EventTimingMORN  EventTiming = "MORN"  // Morning
	EventTimingAFT   EventTiming = "AFT"   // Afternoon
	EventTimingEVE   EventTiming = "EVE"   // Evening
	EventTimingNIGHT EventTiming = "NIGHT" // Night
	EventTimingPHS   EventTiming = "PHS"   // After Sleep
)

// SampledData represents a series of measurements taken by a device
type SampledData struct {
	common.Element

	// Zero value and units
	Origin common.Quantity `json:"origin"`

	// Number of milliseconds between samples
	Period float64 `json:"period"`

	// Multiply data by this before adding to origin
	Factor *float64 `json:"factor,omitempty"`

	// Lower limit of detection
	LowerLimit *float64 `json:"lowerLimit,omitempty"`

	// Upper limit of detection
	UpperLimit *float64 `json:"upperLimit,omitempty"`

	// Number of sample points at each time point
	Dimensions int `json:"dimensions"`

	// Decimal values with spaces, or "E" | "U" | "L"
	Data *string `json:"data,omitempty"`
}
