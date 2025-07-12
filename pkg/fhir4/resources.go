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

// Bundle represents a container for a collection of resources
type Bundle struct {
	Resource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Bundle"

	// Identifies the purpose of this bundle
	Type BundleType `json:"type"`

	// Optional timestamp when this bundle was assembled
	Timestamp *string `json:"timestamp,omitempty"`

	// If a set of search matches, this is the total number of matches for the search
	Total *int `json:"total,omitempty"`

	// A series of links that provide context to this bundle
	Link []BundleLink `json:"link,omitempty"`

	// An entry in a bundle resource
	Entry []BundleEntry `json:"entry,omitempty"`

	// Digital Signature
	Signature *Signature `json:"signature,omitempty"`

	// Persistent identity generally only matters for batches of type Document, Message, and Collection
	Identifier *common.Identifier `json:"identifier,omitempty"`
}

// BundleType represents the purpose of this bundle
type BundleType string

const (
	BundleTypeDocument            BundleType = "document"
	BundleTypeMessage             BundleType = "message"
	BundleTypeTransaction         BundleType = "transaction"
	BundleTypeTransactionResponse BundleType = "transaction-response"
	BundleTypeBatch               BundleType = "batch"
	BundleTypeBatchResponse       BundleType = "batch-response"
	BundleTypeHistory             BundleType = "history"
	BundleTypeSearchSet           BundleType = "searchset"
	BundleTypeCollection          BundleType = "collection"
)

// BundleLink represents a link related to this Bundle
type BundleLink struct {
	common.BackboneElement

	// A name which details the functional use for this link
	Relation string `json:"relation"`

	// The reference details for the link
	URL string `json:"url"`
}

// BundleEntry represents an entry in a bundle resource
type BundleEntry struct {
	common.BackboneElement

	// URI for resource (Absolute URL or relative URL)
	FullURL *string `json:"fullUrl,omitempty"`

	// A series of links that provide context to this entry
	Link []BundleLink `json:"link,omitempty"`

	// The Resource for the entry
	Resource interface{} `json:"resource,omitempty"`

	// Information about the search process that lead to the creation of this entry
	Search *BundleEntrySearch `json:"search,omitempty"`

	// Additional information about how this entry should be processed as part of a transaction or batch
	Request *BundleEntryRequest `json:"request,omitempty"`

	// Indicates the results of processing the corresponding 'request' entry
	Response *BundleEntryResponse `json:"response,omitempty"`
}

// BundleEntrySearch represents information about the search process that lead to the creation of this entry
type BundleEntrySearch struct {
	common.BackboneElement

	// search | match | include | outcome - why this entry is in the result set
	Mode *BundleEntrySearchMode `json:"mode,omitempty"`

	// Search ranking (between 0..1)
	Score *float64 `json:"score,omitempty"`
}

// BundleEntrySearchMode represents why this entry is in the result set
type BundleEntrySearchMode string

const (
	BundleEntrySearchModeMatch   BundleEntrySearchMode = "match"
	BundleEntrySearchModeInclude BundleEntrySearchMode = "include"
	BundleEntrySearchModeOutcome BundleEntrySearchMode = "outcome"
)

// BundleEntryRequest represents additional information about how this entry should be processed as part of a transaction or batch
type BundleEntryRequest struct {
	common.BackboneElement

	// In a transaction or batch, this is the HTTP action to be executed for this entry
	Method BundleEntryRequestMethod `json:"method"`

	// URL for HTTP equivalent of this entry
	URL string `json:"url"`

	// For managing cache currency
	IfNoneMatch *string `json:"ifNoneMatch,omitempty"`

	// For managing cache currency
	IfModifiedSince *string `json:"ifModifiedSince,omitempty"`

	// For managing update contention
	IfMatch *string `json:"ifMatch,omitempty"`

	// For conditional creates
	IfNoneExist *string `json:"ifNoneExist,omitempty"`
}

// BundleEntryRequestMethod represents the HTTP verb used for this entry
type BundleEntryRequestMethod string

const (
	BundleEntryRequestMethodGET    BundleEntryRequestMethod = "GET"
	BundleEntryRequestMethodHEAD   BundleEntryRequestMethod = "HEAD"
	BundleEntryRequestMethodPOST   BundleEntryRequestMethod = "POST"
	BundleEntryRequestMethodPUT    BundleEntryRequestMethod = "PUT"
	BundleEntryRequestMethodDELETE BundleEntryRequestMethod = "DELETE"
	BundleEntryRequestMethodPATCH  BundleEntryRequestMethod = "PATCH"
)

// BundleEntryResponse represents the results of processing the corresponding 'request' entry
type BundleEntryResponse struct {
	common.BackboneElement

	// Status response code (text optional)
	Status string `json:"status"`

	// The location header created by processing this operation
	Location *string `json:"location,omitempty"`

	// The etag for the resource (if relevant)
	Etag *string `json:"etag,omitempty"`

	// Server's date time modified
	LastModified *string `json:"lastModified,omitempty"`

	// OperationOutcome with hints and warnings (for batch/transaction)
	Outcome interface{} `json:"outcome,omitempty"`
}

// Signature represents a signature along with supporting context
type Signature struct {
	common.Element

	// Indication of the reason the entity signed the object(s)
	Type []common.Coding `json:"type"`

	// When the signature was created
	When string `json:"when"`

	// Who signed
	Who common.Reference `json:"who"`

	// The party represented
	OnBehalfOf *common.Reference `json:"onBehalfOf,omitempty"`

	// The technical format of the signed resources
	TargetFormat *string `json:"targetFormat,omitempty"`

	// The technical format of the signature
	SigFormat *string `json:"sigFormat,omitempty"`

	// The base64 encoding of the Signature content
	Data *string `json:"data,omitempty"`
}

// Organization represents a formally or informally recognized grouping of people or organizations
type Organization struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Organization"

	// Whether the organization's record is still in active use
	Active *bool `json:"active,omitempty"`

	// Identifies this organization across multiple systems
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Name used for the organization
	Name *string `json:"name,omitempty"`

	// A list of alternate names that the organization is known as, or was known as in the past
	Alias []string `json:"alias,omitempty"`

	// Kind of organization
	Type []common.CodeableConcept `json:"type,omitempty"`

	// A contact detail for the organization
	Telecom []ContactPoint `json:"telecom,omitempty"`

	// An address for the organization
	Address []Address `json:"address,omitempty"`

	// The organization of which this organization forms a part
	PartOf *common.Reference `json:"partOf,omitempty"`

	// Contact for the organization for a certain purpose
	Contact []OrganizationContact `json:"contact,omitempty"`

	// Technical endpoints providing access to services operated for the organization
	Endpoint []common.Reference `json:"endpoint,omitempty"`
}

// OrganizationContact represents contact for the organization for a certain purpose
type OrganizationContact struct {
	common.BackboneElement

	// The type of contact
	Purpose *common.CodeableConcept `json:"purpose,omitempty"`

	// A name associated with the contact
	Name *HumanName `json:"name,omitempty"`

	// Contact details (telephone, email, etc.) for a contact
	Telecom []ContactPoint `json:"telecom,omitempty"`

	// Visiting or postal addresses for the contact
	Address *Address `json:"address,omitempty"`
}

// Practitioner represents a person who is directly or indirectly involved in the provisioning of healthcare
type Practitioner struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Practitioner"

	// Whether this practitioner's record is in active use
	Active *bool `json:"active,omitempty"`

	// An identifier that applies to this person
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The name(s) associated with the practitioner
	Name []HumanName `json:"name,omitempty"`

	// A contact detail for the practitioner (that apply to all roles)
	Telecom []ContactPoint `json:"telecom,omitempty"`

	// Address(es) of the practitioner that are not role specific (typically home address)
	Address []Address `json:"address,omitempty"`

	// male | female | other | unknown
	Gender *AdministrativeGender `json:"gender,omitempty"`

	// The date of birth for the practitioner
	BirthDate *string `json:"birthDate,omitempty"`

	// Image of the person
	Photo []Attachment `json:"photo,omitempty"`

	// Certification, licenses, or training pertaining to the provision of care
	Qualification []PractitionerQualification `json:"qualification,omitempty"`

	// A language the practitioner can use in patient communication
	Communication []common.CodeableConcept `json:"communication,omitempty"`
}

// PractitionerQualification represents the official certifications, training, and licenses
type PractitionerQualification struct {
	common.BackboneElement

	// An identifier for this qualification for the practitioner
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Coded representation of the qualification
	Code common.CodeableConcept `json:"code"`

	// Period during which the qualification is valid
	Period *common.Period `json:"period,omitempty"`

	// Organization that regulates and issues the qualification
	Issuer *common.Reference `json:"issuer,omitempty"`
}

// Encounter represents an interaction between a patient and healthcare provider(s)
type Encounter struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Encounter"

	// planned | arrived | triaged | in-progress | onleave | finished | cancelled | entered-in-error | unknown
	Status EncounterStatus `json:"status"`

	// Identifier(s) by which this encounter is known
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Classification of patient encounter
	Class common.Coding `json:"class"`

	// Specific type of encounter
	Type []common.CodeableConcept `json:"type,omitempty"`

	// Indicates the urgency of the encounter
	Priority *common.CodeableConcept `json:"priority,omitempty"`

	// The patient or group present at the encounter
	Subject *common.Reference `json:"subject,omitempty"`

	// Episode(s) of care that this encounter should be recorded against
	EpisodeOfCare []common.Reference `json:"episodeOfCare,omitempty"`

	// The ServiceRequest that initiated this encounter
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// List of participants involved in the encounter
	Participant []EncounterParticipant `json:"participant,omitempty"`

	// The appointment that scheduled this encounter
	Appointment []common.Reference `json:"appointment,omitempty"`

	// The start and end time of the encounter
	Period *common.Period `json:"period,omitempty"`

	// Quantity of time the encounter lasted
	Length *Duration `json:"length,omitempty"`

	// Coded reason the encounter takes place
	ReasonCode []common.CodeableConcept `json:"reasonCode,omitempty"`

	// Reason the encounter takes place (reference)
	ReasonReference []common.Reference `json:"reasonReference,omitempty"`

	// The list of diagnosis relevant to this encounter
	Diagnosis []EncounterDiagnosis `json:"diagnosis,omitempty"`

	// The set of accounts that may be used for billing for this Encounter
	Account []common.Reference `json:"account,omitempty"`

	// Details about the admission to a healthcare service
	Hospitalization *EncounterHospitalization `json:"hospitalization,omitempty"`

	// List of locations where the patient has been during this encounter
	Location []EncounterLocation `json:"location,omitempty"`

	// The organization that is primarily responsible for this Encounter's services
	ServiceProvider *common.Reference `json:"serviceProvider,omitempty"`

	// Broad categorization of the service that is to be provided
	ServiceType *common.CodeableConcept `json:"serviceType,omitempty"`

	// Another Encounter this encounter is part of
	PartOf *common.Reference `json:"partOf,omitempty"`

	// The status history permits the encounter resource to contain the status history
	StatusHistory []EncounterStatusHistory `json:"statusHistory,omitempty"`

	// The class history permits the tracking of the encounters transitions
	ClassHistory []EncounterClassHistory `json:"classHistory,omitempty"`
}

// EncounterStatus represents the current state of the encounter
type EncounterStatus string

const (
	EncounterStatusPlanned         EncounterStatus = "planned"
	EncounterStatusArrived         EncounterStatus = "arrived"
	EncounterStatusTriaged         EncounterStatus = "triaged"
	EncounterStatusInProgress      EncounterStatus = "in-progress"
	EncounterStatusOnLeave         EncounterStatus = "onleave"
	EncounterStatusFinished        EncounterStatus = "finished"
	EncounterStatusCancelled       EncounterStatus = "cancelled"
	EncounterStatusEnteredInError  EncounterStatus = "entered-in-error"
	EncounterStatusUnknown         EncounterStatus = "unknown"
)

// EncounterParticipant represents the list of people responsible for providing the service
type EncounterParticipant struct {
	common.BackboneElement

	// Role of participant in encounter
	Type []common.CodeableConcept `json:"type,omitempty"`

	// Period of time during the encounter that the participant participated
	Period *common.Period `json:"period,omitempty"`

	// Persons involved in the encounter other than the patient
	Individual *common.Reference `json:"individual,omitempty"`
}

// EncounterDiagnosis represents the list of diagnosis relevant to this encounter
type EncounterDiagnosis struct {
	common.BackboneElement

	// The diagnosis or procedure relevant to the encounter
	Condition common.Reference `json:"condition"`

	// Role that this diagnosis has within the encounter (e.g. admission, billing, discharge)
	Use *common.CodeableConcept `json:"use,omitempty"`

	// Ranking of the diagnosis (for each role type)
	Rank *int `json:"rank,omitempty"`
}

// EncounterHospitalization represents details about the admission to a healthcare service
type EncounterHospitalization struct {
	common.BackboneElement

	// Pre-admission identifier
	PreAdmissionIdentifier *common.Identifier `json:"preAdmissionIdentifier,omitempty"`

	// The location/organization from which the patient came before admission
	Origin *common.Reference `json:"origin,omitempty"`

	// From where patient was admitted (physician referral, transfer)
	AdmitSource *common.CodeableConcept `json:"admitSource,omitempty"`

	// The type of hospital re-admission that has occurred
	ReAdmission *common.CodeableConcept `json:"reAdmission,omitempty"`

	// Diet preferences reported by the patient
	DietPreference []common.CodeableConcept `json:"dietPreference,omitempty"`

	// Special courtesies (VIP, board member)
	SpecialCourtesy []common.CodeableConcept `json:"specialCourtesy,omitempty"`

	// Any special requests that have been made for this hospitalization encounter
	SpecialArrangement []common.CodeableConcept `json:"specialArrangement,omitempty"`

	// Location/organization to which the patient is discharged
	Destination *common.Reference `json:"destination,omitempty"`

	// Category or kind of location after discharge
	DischargeDisposition *common.CodeableConcept `json:"dischargeDisposition,omitempty"`
}

// EncounterLocation represents list of locations where the patient has been during this encounter
type EncounterLocation struct {
	common.BackboneElement

	// Location the encounter takes place
	Location common.Reference `json:"location"`

	// planned | active | reserved | completed
	Status *EncounterLocationStatus `json:"status,omitempty"`

	// The physical type of the location (usually the level in the location hierarchy)
	PhysicalType *common.CodeableConcept `json:"physicalType,omitempty"`

	// Time period during which the patient was present at the location
	Period *common.Period `json:"period,omitempty"`
}

// EncounterLocationStatus represents the status of the location
type EncounterLocationStatus string

const (
	EncounterLocationStatusPlanned   EncounterLocationStatus = "planned"
	EncounterLocationStatusActive    EncounterLocationStatus = "active"
	EncounterLocationStatusReserved  EncounterLocationStatus = "reserved"
	EncounterLocationStatusCompleted EncounterLocationStatus = "completed"
)

// EncounterStatusHistory represents the status history permits the encounter resource to contain the status history
type EncounterStatusHistory struct {
	common.BackboneElement

	// The status the encounter was in during this period
	Status EncounterStatus `json:"status"`

	// The time that the episode was in the specified status
	Period common.Period `json:"period"`
}

// EncounterClassHistory represents the class history permits the tracking of the encounters transitions
type EncounterClassHistory struct {
	common.BackboneElement

	// inpatient | outpatient | ambulatory | emergency +
	Class common.Coding `json:"class"`

	// The time that the episode was in the specified class
	Period common.Period `json:"period"`
}

// Condition represents a clinical condition, problem, diagnosis, or other event, situation, issue, or clinical concept
type Condition struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Condition"

	// Identifier for this condition
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The clinical status of the condition
	ClinicalStatus *common.CodeableConcept `json:"clinicalStatus,omitempty"`

	// The verification status to support the clinical status of the condition
	VerificationStatus *common.CodeableConcept `json:"verificationStatus,omitempty"`

	// A category assigned to the condition
	Category []common.CodeableConcept `json:"category,omitempty"`

	// A subjective assessment of the severity of the condition
	Severity *common.CodeableConcept `json:"severity,omitempty"`

	// Identification of the condition, problem or diagnosis
	Code *common.CodeableConcept `json:"code,omitempty"`

	// Anatomical location, if relevant
	BodySite []common.CodeableConcept `json:"bodySite,omitempty"`

	// Who has the condition
	Subject common.Reference `json:"subject"`

	// Encounter created as part of
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Estimated or actual date, date-time, or age when condition began
	OnsetDateTime    *string         `json:"onsetDateTime,omitempty"`
	OnsetAge         *Age            `json:"onsetAge,omitempty"`
	OnsetPeriod      *common.Period  `json:"onsetPeriod,omitempty"`
	OnsetRange       *Range          `json:"onsetRange,omitempty"`
	OnsetString      *string         `json:"onsetString,omitempty"`

	// When in resolution/remission
	AbatementDateTime *string         `json:"abatementDateTime,omitempty"`
	AbatementAge      *Age            `json:"abatementAge,omitempty"`
	AbatementPeriod   *common.Period  `json:"abatementPeriod,omitempty"`
	AbatementRange    *Range          `json:"abatementRange,omitempty"`
	AbatementString   *string         `json:"abatementString,omitempty"`
	AbatementBoolean  *bool           `json:"abatementBoolean,omitempty"`

	// Date record was first recorded
	RecordedDate *string `json:"recordedDate,omitempty"`

	// Who recorded the record and takes responsibility for its content
	Recorder *common.Reference `json:"recorder,omitempty"`

	// Supporting evidence
	Evidence []ConditionEvidence `json:"evidence,omitempty"`

	// Additional information about the Condition
	Note []Annotation `json:"note,omitempty"`

	// Stage/grade, usually assessed formally
	Stage []ConditionStage `json:"stage,omitempty"`
}

// ConditionEvidence represents supporting evidence / manifestations that are the basis of the condition
type ConditionEvidence struct {
	common.BackboneElement

	// Manifestation/symptom
	Code []common.CodeableConcept `json:"code,omitempty"`

	// Supporting information found elsewhere
	Detail []common.Reference `json:"detail,omitempty"`
}

// ConditionStage represents stage/grade, usually assessed formally
type ConditionStage struct {
	common.BackboneElement

	// Simple summary (disease specific)
	Summary *common.CodeableConcept `json:"summary,omitempty"`

	// Formal record of assessment
	Assessment []common.Reference `json:"assessment,omitempty"`

	// Kind of staging
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// Procedure represents an action that is or was performed on or for a patient
type Procedure struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Procedure"

	// External Identifiers for this procedure
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// A code specifying the state of the procedure
	Status ProcedureStatus `json:"status"`

	// Reason for current status
	StatusReason *common.CodeableConcept `json:"statusReason,omitempty"`

	// Classification of the procedure
	Category *common.CodeableConcept `json:"category,omitempty"`

	// Identification of the procedure
	Code *common.CodeableConcept `json:"code,omitempty"`

	// Who the procedure was performed on
	Subject common.Reference `json:"subject"`

	// Encounter created as part of
	Encounter *common.Reference `json:"encounter,omitempty"`

	// When the procedure was performed
	PerformedDateTime *string        `json:"performedDateTime,omitempty"`
	PerformedPeriod   *common.Period `json:"performedPeriod,omitempty"`
	PerformedString   *string        `json:"performedString,omitempty"`
	PerformedAge      *Age           `json:"performedAge,omitempty"`
	PerformedRange    *Range         `json:"performedRange,omitempty"`

	// Who recorded the record and takes responsibility for its content
	Recorder *common.Reference `json:"recorder,omitempty"`

	// Person who asserts this procedure
	Asserter *common.Reference `json:"asserter,omitempty"`

	// The people who performed the procedure
	Performer []ProcedurePerformer `json:"performer,omitempty"`

	// Where the procedure happened
	Location *common.Reference `json:"location,omitempty"`

	// Coded reason procedure performed
	ReasonCode []common.CodeableConcept `json:"reasonCode,omitempty"`

	// The justification that the procedure was performed
	ReasonReference []common.Reference `json:"reasonReference,omitempty"`

	// Target body sites
	BodySite []common.CodeableConcept `json:"bodySite,omitempty"`

	// The result of procedure
	Outcome *common.CodeableConcept `json:"outcome,omitempty"`

	// Any report resulting from the procedure
	Report []common.Reference `json:"report,omitempty"`

	// Complication following the procedure
	Complication []common.CodeableConcept `json:"complication,omitempty"`

	// A condition that is a result of the procedure
	ComplicationDetail []common.Reference `json:"complicationDetail,omitempty"`

	// Instructions for follow up
	FollowUp []common.CodeableConcept `json:"followUp,omitempty"`

	// Additional information about the procedure
	Note []Annotation `json:"note,omitempty"`

	// Manipulated, implanted, or removed device
	FocalDevice []ProcedureFocalDevice `json:"focalDevice,omitempty"`

	// Items used during procedure
	UsedReference []common.Reference `json:"usedReference,omitempty"`

	// Coded items used during the procedure
	UsedCode []common.CodeableConcept `json:"usedCode,omitempty"`

	// A request for this procedure
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// Part of referenced event
	PartOf []common.Reference `json:"partOf,omitempty"`
}

// ProcedureStatus represents the status of the procedure
type ProcedureStatus string

const (
	ProcedureStatusPreparation     ProcedureStatus = "preparation"
	ProcedureStatusInProgress      ProcedureStatus = "in-progress"
	ProcedureStatusNotDone         ProcedureStatus = "not-done"
	ProcedureStatusOnHold          ProcedureStatus = "on-hold"
	ProcedureStatusStopped         ProcedureStatus = "stopped"
	ProcedureStatusCompleted       ProcedureStatus = "completed"
	ProcedureStatusEnteredInError  ProcedureStatus = "entered-in-error"
	ProcedureStatusUnknown         ProcedureStatus = "unknown"
)

// ProcedurePerformer represents the people who performed the procedure
type ProcedurePerformer struct {
	common.BackboneElement

	// Type of performance
	Function *common.CodeableConcept `json:"function,omitempty"`

	// The reference to the practitioner
	Actor common.Reference `json:"actor"`

	// Organization the device or practitioner was acting for
	OnBehalfOf *common.Reference `json:"onBehalfOf,omitempty"`
}

// ProcedureFocalDevice represents a device that is implanted, removed or otherwise manipulated
type ProcedureFocalDevice struct {
	common.BackboneElement

	// Kind of change to device
	Action *common.CodeableConcept `json:"action,omitempty"`

	// Device that was changed
	Manipulated common.Reference `json:"manipulated"`
}

// Location represents details and position information for a physical place
type Location struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Location"

	// Unique code or number identifying the location to its users
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// active | suspended | inactive
	Status *LocationStatus `json:"status,omitempty"`

	// Name of the location as used by humans
	Name *string `json:"name,omitempty"`

	// A list of alternate names that the location is known as, or was known as in the past
	Alias []string `json:"alias,omitempty"`

	// Additional details about the location that could be displayed as further information
	Description *string `json:"description,omitempty"`

	// instance | kind - Indicates whether a resource instance represents a specific location or a class of locations
	Mode *LocationMode `json:"mode,omitempty"`

	// Indicates the type of function performed at the location
	Type []common.CodeableConcept `json:"type,omitempty"`

	// Contact details of the location
	Telecom []ContactPoint `json:"telecom,omitempty"`

	// Physical location
	Address *Address `json:"address,omitempty"`

	// Physical form of the location
	PhysicalType *common.CodeableConcept `json:"physicalType,omitempty"`

	// The absolute geographic location
	Position *LocationPosition `json:"position,omitempty"`

	// Organization responsible for provisioning and upkeep
	ManagingOrganization *common.Reference `json:"managingOrganization,omitempty"`

	// Another Location of which this Location is physically a part of
	PartOf *common.Reference `json:"partOf,omitempty"`

	// What days/times during a week is this location usually open
	HoursOfOperation []LocationHoursOfOperation `json:"hoursOfOperation,omitempty"`

	// Description of availability exceptions
	AvailabilityExceptions *string `json:"availabilityExceptions,omitempty"`

	// Technical endpoints providing access to services operated for the location
	Endpoint []common.Reference `json:"endpoint,omitempty"`

	// The operational status covers operation values most relevant to beds
	OperationalStatus *common.Coding `json:"operationalStatus,omitempty"`
}

// LocationStatus represents the status of the location
type LocationStatus string

const (
	LocationStatusActive    LocationStatus = "active"
	LocationStatusSuspended LocationStatus = "suspended"
	LocationStatusInactive  LocationStatus = "inactive"
)

// LocationMode represents whether the location is a specific instance or class of locations
type LocationMode string

const (
	LocationModeInstance LocationMode = "instance"
	LocationModeKind     LocationMode = "kind"
)

// LocationPosition represents the absolute geographic location
type LocationPosition struct {
	common.BackboneElement

	// Longitude with WGS84 datum
	Longitude float64 `json:"longitude"`

	// Latitude with WGS84 datum
	Latitude float64 `json:"latitude"`

	// Altitude with WGS84 datum
	Altitude *float64 `json:"altitude,omitempty"`
}

// LocationHoursOfOperation represents what days/times during a week is this location usually open
type LocationHoursOfOperation struct {
	common.BackboneElement

	// Indicates which days of the week are available between the start and end Times
	DaysOfWeek []DaysOfWeek `json:"daysOfWeek,omitempty"`

	// The Location is open all day
	AllDay *bool `json:"allDay,omitempty"`

	// Time that the Location opens
	OpeningTime *string `json:"openingTime,omitempty"`

	// Time that the Location closes
	ClosingTime *string `json:"closingTime,omitempty"`
}

// DiagnosticReport represents the findings and interpretation of diagnostic tests
type DiagnosticReport struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "DiagnosticReport"

	// Business identifier for report
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Details concerning a service requested
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// registered | partial | preliminary | final | amended | corrected | appended | cancelled | entered-in-error | unknown
	Status DiagnosticReportStatus `json:"status"`

	// Service category
	Category []common.CodeableConcept `json:"category,omitempty"`

	// Name/Code for this diagnostic report
	Code common.CodeableConcept `json:"code"`

	// The subject of the report - usually, but not always, this is a patient
	Subject *common.Reference `json:"subject,omitempty"`

	// Health care event when test ordered
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Clinically relevant time/time-period for report
	EffectiveDateTime *string        `json:"effectiveDateTime,omitempty"`
	EffectivePeriod   *common.Period `json:"effectivePeriod,omitempty"`

	// DateTime this version was made
	Issued *string `json:"issued,omitempty"`

	// Responsible Diagnostic Service
	Performer []common.Reference `json:"performer,omitempty"`

	// Primary result interpreter
	ResultsInterpreter []common.Reference `json:"resultsInterpreter,omitempty"`

	// Specimens this report is based on
	Specimen []common.Reference `json:"specimen,omitempty"`

	// Observations
	Result []common.Reference `json:"result,omitempty"`

	// Key images associated with this report
	Media []DiagnosticReportMedia `json:"media,omitempty"`

	// Clinical conclusion (interpretation) of test results
	Conclusion *string `json:"conclusion,omitempty"`

	// Codes for the clinical conclusion of test results
	ConclusionCode []common.CodeableConcept `json:"conclusionCode,omitempty"`

	// Entire report as attachment
	PresentedForm []Attachment `json:"presentedForm,omitempty"`
}

// DiagnosticReportStatus represents the status of the diagnostic report
type DiagnosticReportStatus string

const (
	DiagnosticReportStatusRegistered      DiagnosticReportStatus = "registered"
	DiagnosticReportStatusPartial         DiagnosticReportStatus = "partial"
	DiagnosticReportStatusPreliminary     DiagnosticReportStatus = "preliminary"
	DiagnosticReportStatusFinal           DiagnosticReportStatus = "final"
	DiagnosticReportStatusAmended         DiagnosticReportStatus = "amended"
	DiagnosticReportStatusCorrected       DiagnosticReportStatus = "corrected"
	DiagnosticReportStatusAppended        DiagnosticReportStatus = "appended"
	DiagnosticReportStatusCancelled       DiagnosticReportStatus = "cancelled"
	DiagnosticReportStatusEnteredInError  DiagnosticReportStatus = "entered-in-error"
	DiagnosticReportStatusUnknown         DiagnosticReportStatus = "unknown"
)

// DiagnosticReportMedia represents key images associated with this report
type DiagnosticReportMedia struct {
	common.BackboneElement

	// Comment about the image (e.g. explanation)
	Comment *string `json:"comment,omitempty"`

	// Reference to the image source
	Link common.Reference `json:"link"`
}

// Medication represents a definition of a medication
type Medication struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Medication"

	// Business identifier for this medication
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// A code that specifies this medication
	Code *common.CodeableConcept `json:"code,omitempty"`

	// active | inactive | entered-in-error
	Status *MedicationStatus `json:"status,omitempty"`

	// Manufacturer of the medication
	Manufacturer *common.Reference `json:"manufacturer,omitempty"`

	// powder | tablets | capsule +
	Form *common.CodeableConcept `json:"form,omitempty"`

	// Amount of drug in package
	Amount *Ratio `json:"amount,omitempty"`

	// Active or inactive ingredient
	Ingredient []MedicationIngredient `json:"ingredient,omitempty"`

	// Details about packaged medications
	Batch *MedicationBatch `json:"batch,omitempty"`
}

// MedicationStatus represents the status of the medication
type MedicationStatus string

const (
	MedicationStatusActive         MedicationStatus = "active"
	MedicationStatusInactive       MedicationStatus = "inactive"
	MedicationStatusEnteredInError MedicationStatus = "entered-in-error"
)

// MedicationIngredient represents active or inactive ingredient
type MedicationIngredient struct {
	common.BackboneElement

	// The actual ingredient - either a substance (simple ingredient) or another medication
	ItemCodeableConcept *common.CodeableConcept `json:"itemCodeableConcept,omitempty"`
	ItemReference       *common.Reference       `json:"itemReference,omitempty"`

	// Indication of whether this ingredient affects the therapeutic action of the drug
	IsActive *bool `json:"isActive,omitempty"`

	// Quantity of ingredient present
	Strength *Ratio `json:"strength,omitempty"`
}

// MedicationBatch represents details about packaged medications
type MedicationBatch struct {
	common.BackboneElement

	// Identifier assigned to batch
	LotNumber *string `json:"lotNumber,omitempty"`

	// When batch will expire
	ExpirationDate *string `json:"expirationDate,omitempty"`
}

// MedicationRequest represents an order or request for both supply of the medication and the instructions for administration
type MedicationRequest struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "MedicationRequest"

	// Identifiers assigned to this order
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// active | on-hold | cancelled | completed | entered-in-error | stopped | draft | unknown
	Status MedicationRequestStatus `json:"status"`

	// Reason for current status
	StatusReason *common.CodeableConcept `json:"statusReason,omitempty"`

	// proposal | plan | order | original-order | reflex-order | filler-order | instance-order | option
	Intent MedicationRequestIntent `json:"intent"`

	// Type of medication usage
	Category []common.CodeableConcept `json:"category,omitempty"`

	// routine | urgent | asap | stat
	Priority *MedicationRequestPriority `json:"priority,omitempty"`

	// True if request is prohibiting action
	DoNotPerform *bool `json:"doNotPerform,omitempty"`

	// Reported rather than primary record
	ReportedBoolean   *bool             `json:"reportedBoolean,omitempty"`
	ReportedReference *common.Reference `json:"reportedReference,omitempty"`

	// Medication to be taken
	MedicationCodeableConcept *common.CodeableConcept `json:"medicationCodeableConcept,omitempty"`
	MedicationReference       *common.Reference       `json:"medicationReference,omitempty"`

	// Who or group medication request is for
	Subject common.Reference `json:"subject"`

	// Encounter created as part of encounter/admission/stay
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Information to support ordering of the medication
	SupportingInformation []common.Reference `json:"supportingInformation,omitempty"`

	// When request was initially authored
	AuthoredOn *string `json:"authoredOn,omitempty"`

	// Who/What requested the Request
	Requester *common.Reference `json:"requester,omitempty"`

	// Person who entered the order on behalf of another
	Recorder *common.Reference `json:"recorder,omitempty"`

	// Reason or indication for ordering or not ordering the medication
	ReasonCode      []common.CodeableConcept `json:"reasonCode,omitempty"`
	ReasonReference []common.Reference       `json:"reasonReference,omitempty"`

	// Instantiates FHIR protocol or definition
	InstantiatesCanonical []string `json:"instantiatesCanonical,omitempty"`
	InstantiatesUri       []string `json:"instantiatesUri,omitempty"`

	// What request fulfills
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// Composite Request ID
	GroupIdentifier *common.Identifier `json:"groupIdentifier,omitempty"`

	// Overall pattern of medication administration
	CourseOfTherapyType *common.CodeableConcept `json:"courseOfTherapyType,omitempty"`

	// Associated insurance coverage
	Insurance []common.Reference `json:"insurance,omitempty"`

	// Information about the prescription
	Note []Annotation `json:"note,omitempty"`

	// How the medication should be taken
	DosageInstruction []Dosage `json:"dosageInstruction,omitempty"`

	// Medication supply authorization
	DispenseRequest *MedicationRequestDispenseRequest `json:"dispenseRequest,omitempty"`

	// Any restrictions on medication substitution
	Substitution *MedicationRequestSubstitution `json:"substitution,omitempty"`

	// An order/prescription that is being replaced
	PriorPrescription *common.Reference `json:"priorPrescription,omitempty"`

	// Clinical Issue with action
	DetectedIssue []common.Reference `json:"detectedIssue,omitempty"`

	// A list of events of interest in the lifecycle
	EventHistory []common.Reference `json:"eventHistory,omitempty"`
}

// MedicationRequestStatus represents the status of the medication request
type MedicationRequestStatus string

const (
	MedicationRequestStatusActive         MedicationRequestStatus = "active"
	MedicationRequestStatusOnHold         MedicationRequestStatus = "on-hold"
	MedicationRequestStatusCancelled      MedicationRequestStatus = "cancelled"
	MedicationRequestStatusCompleted      MedicationRequestStatus = "completed"
	MedicationRequestStatusEnteredInError MedicationRequestStatus = "entered-in-error"
	MedicationRequestStatusStopped        MedicationRequestStatus = "stopped"
	MedicationRequestStatusDraft          MedicationRequestStatus = "draft"
	MedicationRequestStatusUnknown        MedicationRequestStatus = "unknown"
)

// MedicationRequestIntent represents the intent of the medication request
type MedicationRequestIntent string

const (
	MedicationRequestIntentProposal       MedicationRequestIntent = "proposal"
	MedicationRequestIntentPlan           MedicationRequestIntent = "plan"
	MedicationRequestIntentOrder          MedicationRequestIntent = "order"
	MedicationRequestIntentOriginalOrder  MedicationRequestIntent = "original-order"
	MedicationRequestIntentReflexOrder    MedicationRequestIntent = "reflex-order"
	MedicationRequestIntentFillerOrder    MedicationRequestIntent = "filler-order"
	MedicationRequestIntentInstanceOrder  MedicationRequestIntent = "instance-order"
	MedicationRequestIntentOption         MedicationRequestIntent = "option"
)

// MedicationRequestPriority represents the priority of the medication request
type MedicationRequestPriority string

const (
	MedicationRequestPriorityRoutine MedicationRequestPriority = "routine"
	MedicationRequestPriorityUrgent  MedicationRequestPriority = "urgent"
	MedicationRequestPriorityAsap    MedicationRequestPriority = "asap"
	MedicationRequestPriorityStat    MedicationRequestPriority = "stat"
)

// MedicationRequestDispenseRequest represents medication supply authorization
type MedicationRequestDispenseRequest struct {
	common.BackboneElement

	// First fill details
	InitialFill *MedicationRequestDispenseRequestInitialFill `json:"initialFill,omitempty"`

	// Minimum period of time between dispenses
	DispenseInterval *Duration `json:"dispenseInterval,omitempty"`

	// Time period supply is authorized for
	ValidityPeriod *common.Period `json:"validityPeriod,omitempty"`

	// Number of refills authorized
	NumberOfRepeatsAllowed *int `json:"numberOfRepeatsAllowed,omitempty"`

	// Amount of medication to supply per dispense
	Quantity *common.Quantity `json:"quantity,omitempty"`

	// Number of days supply per dispense
	ExpectedSupplyDuration *Duration `json:"expectedSupplyDuration,omitempty"`

	// Intended dispenser
	Performer *common.Reference `json:"performer,omitempty"`
}

// MedicationRequestDispenseRequestInitialFill represents first fill details
type MedicationRequestDispenseRequestInitialFill struct {
	common.BackboneElement

	// First fill quantity
	Quantity *common.Quantity `json:"quantity,omitempty"`

	// First fill duration
	Duration *Duration `json:"duration,omitempty"`
}

// MedicationRequestSubstitution represents any restrictions on medication substitution
type MedicationRequestSubstitution struct {
	common.BackboneElement

	// Whether substitution is allowed or not
	AllowedBoolean        *bool                     `json:"allowedBoolean,omitempty"`
	AllowedCodeableConcept *common.CodeableConcept  `json:"allowedCodeableConcept,omitempty"`

	// Why should (not) substitution be made
	Reason *common.CodeableConcept `json:"reason,omitempty"`
}

// QuestionnaireResponse represents a structured set of questions and their answers
type QuestionnaireResponse struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "QuestionnaireResponse"

	// Business identifier assigned to a particular completed questionnaire
	Identifier *common.Identifier `json:"identifier,omitempty"`

	// Plan/proposal/order fulfilled by this QuestionnaireResponse
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// Composition this QuestionnaireResponse is part of
	PartOf []common.Reference `json:"partOf,omitempty"`

	// Form being answered
	Questionnaire *string `json:"questionnaire,omitempty"`

	// in-progress | completed | amended | entered-in-error | stopped
	Status QuestionnaireResponseStatus `json:"status"`

	// The subject of the questionnaire response
	Subject *common.Reference `json:"subject,omitempty"`

	// Encounter created as part of
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Date the answers were gathered
	Authored *string `json:"authored,omitempty"`

	// Person who received and recorded the answers
	Author *common.Reference `json:"author,omitempty"`

	// The person who answered the questions
	Source *common.Reference `json:"source,omitempty"`

	// Groups and questions
	Item []QuestionnaireResponseItem `json:"item,omitempty"`
}

// QuestionnaireResponseStatus represents the lifecycle status of the questionnaire response
type QuestionnaireResponseStatus string

const (
	QuestionnaireResponseStatusInProgress      QuestionnaireResponseStatus = "in-progress"
	QuestionnaireResponseStatusCompleted       QuestionnaireResponseStatus = "completed"
	QuestionnaireResponseStatusAmended         QuestionnaireResponseStatus = "amended"
	QuestionnaireResponseStatusEnteredInError  QuestionnaireResponseStatus = "entered-in-error"
	QuestionnaireResponseStatusStopped         QuestionnaireResponseStatus = "stopped"
)

// QuestionnaireResponseItem represents groups and questions
type QuestionnaireResponseItem struct {
	common.BackboneElement

	// Pointer to specific item from Questionnaire
	LinkId string `json:"linkId"`

	// ElementDefinition - details for the item
	Definition *string `json:"definition,omitempty"`

	// Name for group or question text
	Text *string `json:"text,omitempty"`

	// The response(s) to the question
	Answer []QuestionnaireResponseItemAnswer `json:"answer,omitempty"`

	// Nested questionnaire response items
	Item []QuestionnaireResponseItem `json:"item,omitempty"`
}

// QuestionnaireResponseItemAnswer represents the response(s) to the question
type QuestionnaireResponseItemAnswer struct {
	common.BackboneElement

	// Single-valued answer to the question
	ValueBoolean     *bool                   `json:"valueBoolean,omitempty"`
	ValueDecimal     *float64                `json:"valueDecimal,omitempty"`
	ValueInteger     *int                    `json:"valueInteger,omitempty"`
	ValueDate        *string                 `json:"valueDate,omitempty"`
	ValueDateTime    *string                 `json:"valueDateTime,omitempty"`
	ValueTime        *string                 `json:"valueTime,omitempty"`
	ValueString      *string                 `json:"valueString,omitempty"`
	ValueUri         *string                 `json:"valueUri,omitempty"`
	ValueAttachment  *Attachment             `json:"valueAttachment,omitempty"`
	ValueCoding      *common.Coding          `json:"valueCoding,omitempty"`
	ValueQuantity    *common.Quantity        `json:"valueQuantity,omitempty"`
	ValueReference   *common.Reference       `json:"valueReference,omitempty"`

	// Nested items inside this answer
	Item []QuestionnaireResponseItem `json:"item,omitempty"`
}

// ResearchSubject represents a physical entity which is the primary unit of operational and/or administrative interest in a study
type ResearchSubject struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "ResearchSubject"

	// Business identifier for research subject in a study
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// candidate | eligible | follow-up | ineligible | not-registered | off-study | on-study | on-study-intervention | on-study-observation | pending-on-study | potential-candidate | screening | withdrawn
	Status ResearchSubjectStatus `json:"status"`

	// Dates the subject began and ended their participation in the study
	Period *common.Period `json:"period,omitempty"`

	// Study subject is part of
	Study common.Reference `json:"study"`

	// Who or what is part of study
	Individual common.Reference `json:"individual"`

	// What path should be followed
	AssignedArm *string `json:"assignedArm,omitempty"`

	// What path was followed
	ActualArm *string `json:"actualArm,omitempty"`

	// Agreement to participate in study
	Consent *common.Reference `json:"consent,omitempty"`
}

// ResearchSubjectStatus represents the lifecycle status of the research subject
type ResearchSubjectStatus string

const (
	ResearchSubjectStatusCandidate          ResearchSubjectStatus = "candidate"
	ResearchSubjectStatusEligible           ResearchSubjectStatus = "eligible"
	ResearchSubjectStatusFollowUp           ResearchSubjectStatus = "follow-up"
	ResearchSubjectStatusIneligible         ResearchSubjectStatus = "ineligible"
	ResearchSubjectStatusNotRegistered      ResearchSubjectStatus = "not-registered"
	ResearchSubjectStatusOffStudy           ResearchSubjectStatus = "off-study"
	ResearchSubjectStatusOnStudy            ResearchSubjectStatus = "on-study"
	ResearchSubjectStatusOnStudyIntervention ResearchSubjectStatus = "on-study-intervention"
	ResearchSubjectStatusOnStudyObservation ResearchSubjectStatus = "on-study-observation"
	ResearchSubjectStatusPendingOnStudy     ResearchSubjectStatus = "pending-on-study"
	ResearchSubjectStatusPotentialCandidate ResearchSubjectStatus = "potential-candidate"
	ResearchSubjectStatusScreening          ResearchSubjectStatus = "screening"
	ResearchSubjectStatusWithdrawn          ResearchSubjectStatus = "withdrawn"
)

// Consent represents a record of a healthcare consumer's choices
type Consent struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Consent"

	// Identifier for this copy of the consent
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// draft | proposed | active | rejected | inactive | entered-in-error
	Status ConsentStatus `json:"status"`

	// Which of the four areas this resource covers (extensible)
	Scope common.CodeableConcept `json:"scope"`

	// Classification of the consent statement - for indexing/retrieval
	Category []common.CodeableConcept `json:"category"`

	// Who the consent applies to
	Patient *common.Reference `json:"patient,omitempty"`

	// When consent was agreed to
	DateTime *string `json:"dateTime,omitempty"`

	// Who is agreeing to the policy and rules
	Performer []common.Reference `json:"performer,omitempty"`

	// Custodian of the consent
	Organization []common.Reference `json:"organization,omitempty"`

	// Source from which this consent is taken
	SourceAttachment *Attachment       `json:"sourceAttachment,omitempty"`
	SourceReference  *common.Reference `json:"sourceReference,omitempty"`

	// Policies covered by this consent
	Policy []ConsentPolicy `json:"policy,omitempty"`

	// Regulation that this consents to
	PolicyRule *common.CodeableConcept `json:"policyRule,omitempty"`

	// Consent provision(s)
	Provision *ConsentProvision `json:"provision,omitempty"`

	// Contract verification
	Verification []ConsentVerification `json:"verification,omitempty"`
}

// ConsentStatus represents the status of the consent
type ConsentStatus string

const (
	ConsentStatusDraft           ConsentStatus = "draft"
	ConsentStatusProposed        ConsentStatus = "proposed"
	ConsentStatusActive          ConsentStatus = "active"
	ConsentStatusRejected        ConsentStatus = "rejected"
	ConsentStatusInactive        ConsentStatus = "inactive"
	ConsentStatusEnteredInError  ConsentStatus = "entered-in-error"
)

// ConsentPolicy represents policies covered by this consent
type ConsentPolicy struct {
	common.BackboneElement

	// Entity or Organization having regulatory jurisdiction
	Authority *string `json:"authority,omitempty"`

	// Specific policy covered by this consent
	Uri *string `json:"uri,omitempty"`
}

// ConsentProvision represents constraints to the base Consent.policyRule
type ConsentProvision struct {
	common.BackboneElement

	// deny | permit
	Type *ConsentProvisionType `json:"type,omitempty"`

	// Timeframe for this rule
	Period *common.Period `json:"period,omitempty"`

	// Who|what controlled by this rule
	Actor []ConsentProvisionActor `json:"actor,omitempty"`

	// Actions controlled by this rule
	Action []common.CodeableConcept `json:"action,omitempty"`

	// Security Labels that define affected resources
	SecurityLabel []common.Coding `json:"securityLabel,omitempty"`

	// Context of activities covered by this rule
	Purpose []common.Coding `json:"purpose,omitempty"`

	// e.g. Resource Type, Profile, CDA, etc.
	Class []common.Coding `json:"class,omitempty"`

	// e.g. LOINC or SNOMED CT code, etc. in the content
	Code []common.CodeableConcept `json:"code,omitempty"`

	// Timeframe for data controlled by this rule
	DataPeriod *common.Period `json:"dataPeriod,omitempty"`

	// Data controlled by this rule
	Data []ConsentProvisionData `json:"data,omitempty"`

	// Nested Exception Rules
	Provision []ConsentProvision `json:"provision,omitempty"`
}

// ConsentProvisionType represents deny | permit
type ConsentProvisionType string

const (
	ConsentProvisionTypeDeny   ConsentProvisionType = "deny"
	ConsentProvisionTypePermit ConsentProvisionType = "permit"
)

// ConsentProvisionActor represents who|what controlled by this rule
type ConsentProvisionActor struct {
	common.BackboneElement

	// How the actor is involved
	Role common.CodeableConcept `json:"role"`

	// Resource for the actor (or group, by role)
	Reference *common.Reference `json:"reference,omitempty"`
}

// ConsentProvisionData represents data controlled by this rule
type ConsentProvisionData struct {
	common.BackboneElement

	// instance | related | dependents | authoredby
	Meaning ConsentProvisionDataMeaning `json:"meaning"`

	// The actual data reference
	Reference common.Reference `json:"reference"`
}

// ConsentProvisionDataMeaning represents how the resource reference is interpreted
type ConsentProvisionDataMeaning string

const (
	ConsentProvisionDataMeaningInstance   ConsentProvisionDataMeaning = "instance"
	ConsentProvisionDataMeaningRelated    ConsentProvisionDataMeaning = "related"
	ConsentProvisionDataMeaningDependents ConsentProvisionDataMeaning = "dependents"
	ConsentProvisionDataMeaningAuthoredby ConsentProvisionDataMeaning = "authoredby"
)

// ConsentVerification represents contract verification
type ConsentVerification struct {
	common.BackboneElement

	// Has been verified
	Verified bool `json:"verified"`

	// Person who verified
	VerifiedWith *common.Reference `json:"verifiedWith,omitempty"`

	// When consent verified
	VerificationDate *string `json:"verificationDate,omitempty"`
}

// Device represents a medical or non-medical device used in healthcare
type Device struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Device"

	// Instance identifier
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Reference to the definition for the device
	Definition *common.Reference `json:"definition,omitempty"`

	// Unique Device Identifier (UDI) Barcode string
	UdiCarrier []DeviceUdiCarrier `json:"udiCarrier,omitempty"`

	// active | inactive | entered-in-error | unknown
	Status *DeviceStatus `json:"status,omitempty"`

	// Reason for the device's status
	StatusReason []common.CodeableConcept `json:"statusReason,omitempty"`

	// For example, this applies to devices in the United States regulated under Code of Federal Regulation 21CFR§1271.290(c)
	DistinctIdentifier *string `json:"distinctIdentifier,omitempty"`

	// Name of device manufacturer
	Manufacturer *string `json:"manufacturer,omitempty"`

	// Date when the device was made
	ManufactureDate *string `json:"manufactureDate,omitempty"`

	// Date and time of expiry of this device (if applicable)
	ExpirationDate *string `json:"expirationDate,omitempty"`

	// Lot number assigned by the manufacturer
	LotNumber *string `json:"lotNumber,omitempty"`

	// The serial number assigned by the organization when the device was manufactured
	SerialNumber *string `json:"serialNumber,omitempty"`

	// The names of the device as given by the manufacturer
	DeviceName []DeviceDeviceName `json:"deviceName,omitempty"`

	// The model number for the device
	ModelNumber *string `json:"modelNumber,omitempty"`

	// The part number of the device
	PartNumber *string `json:"partNumber,omitempty"`

	// The kind or type of device
	Type *common.CodeableConcept `json:"type,omitempty"`

	// Device capabilities, specializations, or standards supported
	Specialization []DeviceSpecialization `json:"specialization,omitempty"`

	// The actual configuration settings of a device
	Property []DeviceProperty `json:"property,omitempty"`

	// Patient to whom Device is affixed
	Patient *common.Reference `json:"patient,omitempty"`

	// Organization responsible for device
	Owner *common.Reference `json:"owner,omitempty"`

	// Details for human/organization for support
	Contact []ContactPoint `json:"contact,omitempty"`

	// Where the device is found
	Location *common.Reference `json:"location,omitempty"`

	// Network address to contact device
	Url *string `json:"url,omitempty"`

	// Device notes and comments
	Note []Annotation `json:"note,omitempty"`

	// Safety Characteristics of Device
	Safety []common.CodeableConcept `json:"safety,omitempty"`

	// The parent device
	Parent *common.Reference `json:"parent,omitempty"`
}

// DeviceStatus represents the availability status of the device
type DeviceStatus string

const (
	DeviceStatusActive         DeviceStatus = "active"
	DeviceStatusInactive       DeviceStatus = "inactive"
	DeviceStatusEnteredInError DeviceStatus = "entered-in-error"
	DeviceStatusUnknown        DeviceStatus = "unknown"
)

// DeviceUdiCarrier represents Unique Device Identifier
type DeviceUdiCarrier struct {
	common.BackboneElement

	// Mandatory fixed portion of UDI
	DeviceIdentifier *string `json:"deviceIdentifier,omitempty"`

	// UDI Issuing Organization
	Issuer *string `json:"issuer,omitempty"`

	// Regional UDI authority
	Jurisdiction *string `json:"jurisdiction,omitempty"`

	// UDI Machine Readable Barcode String
	CarrierAIDC *string `json:"carrierAIDC,omitempty"`

	// UDI Human Readable Barcode String
	CarrierHRF *string `json:"carrierHRF,omitempty"`

	// barcode | rfid | manual +
	EntryType *UDIEntryType `json:"entryType,omitempty"`
}

// UDIEntryType represents how the UDI was entered
type UDIEntryType string

const (
	UDIEntryTypeBarcode UDIEntryType = "barcode"
	UDIEntryTypeRFID    UDIEntryType = "rfid"
	UDIEntryTypeManual  UDIEntryType = "manual"
	UDIEntryTypeCard    UDIEntryType = "card"
	UDIEntryTypeSelfReported UDIEntryType = "self-reported"
	UDIEntryTypeUnknown UDIEntryType = "unknown"
)

// DeviceDeviceName represents name of device as given by manufacturer
type DeviceDeviceName struct {
	common.BackboneElement

	// The name of the device
	Name string `json:"name"`

	// udi-label-name | user-friendly-name | patient-reported-name | manufacturer-name | model-name | other
	Type DeviceNameType `json:"type"`
}

// DeviceNameType represents the type of device name
type DeviceNameType string

const (
	DeviceNameTypeUdiLabelName        DeviceNameType = "udi-label-name"
	DeviceNameTypeUserFriendlyName    DeviceNameType = "user-friendly-name"
	DeviceNameTypePatientReportedName DeviceNameType = "patient-reported-name"
	DeviceNameTypeManufacturerName    DeviceNameType = "manufacturer-name"
	DeviceNameTypeModelName           DeviceNameType = "model-name"
	DeviceNameTypeOther               DeviceNameType = "other"
)

// DeviceSpecialization represents device capabilities
type DeviceSpecialization struct {
	common.BackboneElement

	// The standard that is used to operate and communicate
	SystemType common.CodeableConcept `json:"systemType"`

	// The version of the standard that is used
	Version *string `json:"version,omitempty"`
}

// DeviceProperty represents device configuration settings
type DeviceProperty struct {
	common.BackboneElement

	// Code that specifies the property DeviceDefinitionPropetyCode (Extensible)
	Type common.CodeableConcept `json:"type"`

	// Property value as a quantity
	ValueQuantity []common.Quantity `json:"valueQuantity,omitempty"`

	// Property value as a code, e.g., NTP4 (synced to NTP)
	ValueCode []common.CodeableConcept `json:"valueCode,omitempty"`
}

// Composition represents a set of healthcare-related information assembled into a single logical package
type Composition struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Composition"

	// Version-independent identifier for the Composition
	Identifier *common.Identifier `json:"identifier,omitempty"`

	// preliminary | final | amended | entered-in-error
	Status CompositionStatus `json:"status"`

	// Kind of composition (LOINC if possible)
	Type common.CodeableConcept `json:"type"`

	// Categorization of Composition
	Category []common.CodeableConcept `json:"category,omitempty"`

	// Who and/or what the composition is about
	Subject *common.Reference `json:"subject,omitempty"`

	// Context of the Composition
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Composition editing time
	Date string `json:"date"`

	// Who and/or what authored the composition
	Author []common.Reference `json:"author"`

	// Human Readable name/title
	Title string `json:"title"`

	// As defined by affinity domain
	Confidentiality *string `json:"confidentiality,omitempty"`

	// Attests to accuracy of composition
	Attester []CompositionAttester `json:"attester,omitempty"`

	// Organization which maintains the composition
	Custodian *common.Reference `json:"custodian,omitempty"`

	// Relationships to other compositions/documents
	RelatesTo []CompositionRelatesTo `json:"relatesTo,omitempty"`

	// The clinical service(s) being documented
	Event []CompositionEvent `json:"event,omitempty"`

	// Composition is broken into sections
	Section []CompositionSection `json:"section,omitempty"`
}

// CompositionStatus represents the workflow/clinical status of the composition
type CompositionStatus string

const (
	CompositionStatusPreliminary    CompositionStatus = "preliminary"
	CompositionStatusFinal          CompositionStatus = "final"
	CompositionStatusAmended        CompositionStatus = "amended"
	CompositionStatusEnteredInError CompositionStatus = "entered-in-error"
)

// CompositionAttester represents a participant who has attested to the accuracy of the composition
type CompositionAttester struct {
	common.BackboneElement

	// personal | professional | legal | official
	Mode CompositionAttesterMode `json:"mode"`

	// When the composition was attested
	Time *string `json:"time,omitempty"`

	// Who attested the composition
	Party *common.Reference `json:"party,omitempty"`
}

// CompositionAttesterMode represents the type of attestation the authenticator offers
type CompositionAttesterMode string

const (
	CompositionAttesterModePersonal     CompositionAttesterMode = "personal"
	CompositionAttesterModeProfessional CompositionAttesterMode = "professional"
	CompositionAttesterModeLegal        CompositionAttesterMode = "legal"
	CompositionAttesterModeOfficial     CompositionAttesterMode = "official"
)

// CompositionRelatesTo represents relationships to other compositions/documents
type CompositionRelatesTo struct {
	common.BackboneElement

	// replaces | transforms | signs | appends
	Code CompositionRelatesToCode `json:"code"`

	// Target of the relationship
	TargetIdentifier *common.Identifier `json:"targetIdentifier,omitempty"`
	TargetReference  *common.Reference  `json:"targetReference,omitempty"`
}

// CompositionRelatesToCode represents the type of relationship that this composition has with another composition or document
type CompositionRelatesToCode string

const (
	CompositionRelatesToCodeReplaces   CompositionRelatesToCode = "replaces"
	CompositionRelatesToCodeTransforms CompositionRelatesToCode = "transforms"
	CompositionRelatesToCodeSigns      CompositionRelatesToCode = "signs"
	CompositionRelatesToCodeAppends    CompositionRelatesToCode = "appends"
)

// CompositionEvent represents the clinical service being documented
type CompositionEvent struct {
	common.BackboneElement

	// Code(s) that apply to the event being documented
	Code []common.CodeableConcept `json:"code,omitempty"`

	// The period covered by the documentation
	Period *common.Period `json:"period,omitempty"`

	// The event(s) being documented
	Detail []common.Reference `json:"detail,omitempty"`
}

// CompositionSection represents the root of the sections that make up the composition
type CompositionSection struct {
	common.BackboneElement

	// Label for section (e.g. for ToC)
	Title *string `json:"title,omitempty"`

	// Classification of section (recommended)
	Code *common.CodeableConcept `json:"code,omitempty"`

	// Who and/or what authored the section
	Author []common.Reference `json:"author,omitempty"`

	// Who/what the section is about, when it is not about the subject of composition
	Focus *common.Reference `json:"focus,omitempty"`

	// Text summary of the section, for human interpretation
	Text *Narrative `json:"text,omitempty"`

	// working | snapshot | changes
	Mode *SectionMode `json:"mode,omitempty"`

	// Order of section entries
	OrderedBy *common.CodeableConcept `json:"orderedBy,omitempty"`

	// A reference to data that supports this section
	Entry []common.Reference `json:"entry,omitempty"`

	// Why the section is empty
	EmptyReason *common.CodeableConcept `json:"emptyReason,omitempty"`

	// Nested Section
	Section []CompositionSection `json:"section,omitempty"`
}

// SectionMode represents how the entry list was prepared
type SectionMode string

const (
	SectionModeWorking  SectionMode = "working"
	SectionModeSnapshot SectionMode = "snapshot"
	SectionModeChanges  SectionMode = "changes"
)

// DocumentReference represents a reference to a document of any kind
type DocumentReference struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "DocumentReference"

	// Master Version Specific Identifier
	MasterIdentifier *common.Identifier `json:"masterIdentifier,omitempty"`

	// Other identifiers for the document
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// current | superseded | entered-in-error
	Status DocumentReferenceStatus `json:"status"`

	// preliminary | final | amended | entered-in-error
	DocStatus *DocumentReferenceDocStatus `json:"docStatus,omitempty"`

	// Kind of document (LOINC if possible)
	Type *common.CodeableConcept `json:"type,omitempty"`

	// Categorization of document
	Category []common.CodeableConcept `json:"category,omitempty"`

	// Who/what is the subject of the document
	Subject *common.Reference `json:"subject,omitempty"`

	// When this document reference was created
	Date *string `json:"date,omitempty"`

	// Who and/or what authored the document
	Author []common.Reference `json:"author,omitempty"`

	// Who/what authenticated the document
	Authenticator *common.Reference `json:"authenticator,omitempty"`

	// Organization which maintains the document
	Custodian *common.Reference `json:"custodian,omitempty"`

	// Relationships to other documents
	RelatesTo []DocumentReferenceRelatesTo `json:"relatesTo,omitempty"`

	// Human-readable description
	Description *string `json:"description,omitempty"`

	// Document security-tags
	SecurityLabel []common.CodeableConcept `json:"securityLabel,omitempty"`

	// Document referenced
	Content []DocumentReferenceContent `json:"content"`

	// Clinical context of document
	Context *DocumentReferenceContext `json:"context,omitempty"`
}

// DocumentReferenceStatus represents the status of the DocumentReference
type DocumentReferenceStatus string

const (
	DocumentReferenceStatusCurrent        DocumentReferenceStatus = "current"
	DocumentReferenceStatusSuperseded     DocumentReferenceStatus = "superseded"
	DocumentReferenceStatusEnteredInError DocumentReferenceStatus = "entered-in-error"
)

// DocumentReferenceDocStatus represents the status of the underlying document
type DocumentReferenceDocStatus string

const (
	DocumentReferenceDocStatusPreliminary    DocumentReferenceDocStatus = "preliminary"
	DocumentReferenceDocStatusFinal          DocumentReferenceDocStatus = "final"
	DocumentReferenceDocStatusAmended        DocumentReferenceDocStatus = "amended"
	DocumentReferenceDocStatusEnteredInError DocumentReferenceDocStatus = "entered-in-error"
)

// DocumentReferenceRelatesTo represents relationships to other documents
type DocumentReferenceRelatesTo struct {
	common.BackboneElement

	// replaces | transforms | signs | appends
	Code DocumentReferenceRelatesToCode `json:"code"`

	// Target of the relationship
	Target common.Reference `json:"target"`
}

// DocumentReferenceRelatesToCode represents the type of relationship between documents
type DocumentReferenceRelatesToCode string

const (
	DocumentReferenceRelatesToCodeReplaces   DocumentReferenceRelatesToCode = "replaces"
	DocumentReferenceRelatesToCodeTransforms DocumentReferenceRelatesToCode = "transforms"
	DocumentReferenceRelatesToCodeSigns      DocumentReferenceRelatesToCode = "signs"
	DocumentReferenceRelatesToCodeAppends    DocumentReferenceRelatesToCode = "appends"
)

// DocumentReferenceContent represents the document and format referenced
type DocumentReferenceContent struct {
	common.BackboneElement

	// Where to access the document
	Attachment Attachment `json:"attachment"`

	// Format/content rules for the document
	Format *common.Coding `json:"format,omitempty"`
}

// DocumentReferenceContext represents clinical context of document
type DocumentReferenceContext struct {
	common.BackboneElement

	// Context of the document content
	Encounter []common.Reference `json:"encounter,omitempty"`

	// Main clinical acts documented
	Event []common.CodeableConcept `json:"event,omitempty"`

	// Time of service that is being documented
	Period *common.Period `json:"period,omitempty"`

	// Kind of facility where patient was seen
	FacilityType *common.CodeableConcept `json:"facilityType,omitempty"`

	// Additional details about where the content was created (e.g. clinical specialty)
	PracticeSetting *common.CodeableConcept `json:"practiceSetting,omitempty"`

	// Patient demographics from source
	SourcePatientInfo *common.Reference `json:"sourcePatientInfo,omitempty"`

	// Related identifiers or resources
	Related []common.Reference `json:"related,omitempty"`
}

// Communication represents an occurrence of information being transmitted
type Communication struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Communication"

	// Unique identifier
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Instantiates FHIR protocol or definition
	InstantiatesCanonical []string `json:"instantiatesCanonical,omitempty"`

	// Instantiates external protocol or definition
	InstantiatesUri []string `json:"instantiatesUri,omitempty"`

	// Request fulfilled by this communication
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// Part of this action
	PartOf []common.Reference `json:"partOf,omitempty"`

	// Reply to
	InResponseTo []common.Reference `json:"inResponseTo,omitempty"`

	// preparation | in-progress | not-done | on-hold | stopped | completed | entered-in-error | unknown
	Status CommunicationStatus `json:"status"`

	// Reason for current status
	StatusReason *common.CodeableConcept `json:"statusReason,omitempty"`

	// Message category
	Category []common.CodeableConcept `json:"category,omitempty"`

	// routine | urgent | asap | stat
	Priority *CommunicationPriority `json:"priority,omitempty"`

	// A channel of communication
	Medium []common.CodeableConcept `json:"medium,omitempty"`

	// Focus of message
	Subject *common.Reference `json:"subject,omitempty"`

	// Description of the purpose/content
	Topic *common.CodeableConcept `json:"topic,omitempty"`

	// Resources that pertain to this communication
	About []common.Reference `json:"about,omitempty"`

	// Encounter created as part of
	Encounter *common.Reference `json:"encounter,omitempty"`

	// When sent
	Sent *string `json:"sent,omitempty"`

	// When received
	Received *string `json:"received,omitempty"`

	// Message recipient
	Recipient []common.Reference `json:"recipient,omitempty"`

	// Message sender
	Sender *common.Reference `json:"sender,omitempty"`

	// Indication for message
	ReasonCode []common.CodeableConcept `json:"reasonCode,omitempty"`

	// Why was communication done?
	ReasonReference []common.Reference `json:"reasonReference,omitempty"`

	// Message payload
	Payload []CommunicationPayload `json:"payload,omitempty"`

	// Comments made about the communication
	Note []Annotation `json:"note,omitempty"`
}

// CommunicationStatus represents the status of the communication
type CommunicationStatus string

const (
	CommunicationStatusPreparation     CommunicationStatus = "preparation"
	CommunicationStatusInProgress      CommunicationStatus = "in-progress"
	CommunicationStatusNotDone         CommunicationStatus = "not-done"
	CommunicationStatusOnHold          CommunicationStatus = "on-hold"
	CommunicationStatusStopped         CommunicationStatus = "stopped"
	CommunicationStatusCompleted       CommunicationStatus = "completed"
	CommunicationStatusEnteredInError  CommunicationStatus = "entered-in-error"
	CommunicationStatusUnknown         CommunicationStatus = "unknown"
)

// CommunicationPriority represents the priority of the communication
type CommunicationPriority string

const (
	CommunicationPriorityRoutine CommunicationPriority = "routine"
	CommunicationPriorityUrgent  CommunicationPriority = "urgent"
	CommunicationPriorityAsap    CommunicationPriority = "asap"
	CommunicationPriorityStat    CommunicationPriority = "stat"
)

// CommunicationPayload represents message payload
type CommunicationPayload struct {
	common.BackboneElement

	// Message part content
	ContentString     *string           `json:"contentString,omitempty"`
	ContentAttachment *Attachment       `json:"contentAttachment,omitempty"`
	ContentReference  *common.Reference `json:"contentReference,omitempty"`
}

// AllergyIntolerance represents risk of harmful or undesirable, physiological response to a substance
type AllergyIntolerance struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "AllergyIntolerance"

	// External ids for this item
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// active | inactive | resolved
	ClinicalStatus *common.CodeableConcept `json:"clinicalStatus,omitempty"`

	// unconfirmed | confirmed | refuted | entered-in-error
	VerificationStatus *common.CodeableConcept `json:"verificationStatus,omitempty"`

	// allergy | intolerance - Underlying mechanism (if known)
	Type *AllergyIntoleranceType `json:"type,omitempty"`

	// food | medication | environment | biologic
	Category []AllergyIntoleranceCategory `json:"category,omitempty"`

	// low | high | unable-to-assess
	Criticality *AllergyIntoleranceCriticality `json:"criticality,omitempty"`

	// Code that identifies the allergy or intolerance
	Code *common.CodeableConcept `json:"code,omitempty"`

	// Who the sensitivity is for
	Patient common.Reference `json:"patient"`

	// Encounter when the allergy or intolerance was asserted
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Date first version of the resource instance was recorded
	OnsetDateTime *string        `json:"onsetDateTime,omitempty"`
	OnsetAge      *Age           `json:"onsetAge,omitempty"`
	OnsetPeriod   *common.Period `json:"onsetPeriod,omitempty"`
	OnsetRange    *Range         `json:"onsetRange,omitempty"`
	OnsetString   *string        `json:"onsetString,omitempty"`

	// Date(/time) of last known occurrence of a reaction
	RecordedDate *string `json:"recordedDate,omitempty"`

	// Who recorded the sensitivity
	Recorder *common.Reference `json:"recorder,omitempty"`

	// Source of the information about the allergy
	Asserter *common.Reference `json:"asserter,omitempty"`

	// Date(/time) of last known occurrence of a reaction
	LastOccurrence *string `json:"lastOccurrence,omitempty"`

	// Additional text not captured in other fields
	Note []Annotation `json:"note,omitempty"`

	// Adverse Reaction Events linked to exposure to substance
	Reaction []AllergyIntoleranceReaction `json:"reaction,omitempty"`
}

// AllergyIntoleranceType represents the type of allergy
type AllergyIntoleranceType string

const (
	AllergyIntoleranceTypeAllergy     AllergyIntoleranceType = "allergy"
	AllergyIntoleranceTypeIntolerance AllergyIntoleranceType = "intolerance"
)

// AllergyIntoleranceCategory represents the category of the identified substance
type AllergyIntoleranceCategory string

const (
	AllergyIntoleranceCategoryFood        AllergyIntoleranceCategory = "food"
	AllergyIntoleranceCategoryMedication  AllergyIntoleranceCategory = "medication"
	AllergyIntoleranceCategoryEnvironment AllergyIntoleranceCategory = "environment"
	AllergyIntoleranceCategoryBiologic    AllergyIntoleranceCategory = "biologic"
)

// AllergyIntoleranceCriticality represents the potential clinical harm
type AllergyIntoleranceCriticality string

const (
	AllergyIntoleranceCriticalityLow            AllergyIntoleranceCriticality = "low"
	AllergyIntoleranceCriticalityHigh           AllergyIntoleranceCriticality = "high"
	AllergyIntoleranceCriticalityUnableToAssess AllergyIntoleranceCriticality = "unable-to-assess"
)

// AllergyIntoleranceReaction represents details about each adverse reaction event
type AllergyIntoleranceReaction struct {
	common.BackboneElement

	// Specific substance or pharmaceutical product considered to be responsible
	Substance *common.CodeableConcept `json:"substance,omitempty"`

	// Clinical symptoms/signs associated with the Event
	Manifestation []common.CodeableConcept `json:"manifestation"`

	// Description of the event as a whole
	Description *string `json:"description,omitempty"`

	// Date(/time) when manifestations showed
	Onset *string `json:"onset,omitempty"`

	// mild | moderate | severe (of event as a whole)
	Severity *AllergyIntoleranceReactionSeverity `json:"severity,omitempty"`

	// How the subject was exposed to the substance
	ExposureRoute *common.CodeableConcept `json:"exposureRoute,omitempty"`

	// Text about event not captured in other fields
	Note []Annotation `json:"note,omitempty"`
}

// AllergyIntoleranceReactionSeverity represents the clinical assessment of severity
type AllergyIntoleranceReactionSeverity string

const (
	AllergyIntoleranceReactionSeverityMild     AllergyIntoleranceReactionSeverity = "mild"
	AllergyIntoleranceReactionSeverityModerate AllergyIntoleranceReactionSeverity = "moderate"
	AllergyIntoleranceReactionSeveritySevere   AllergyIntoleranceReactionSeverity = "severe"
)

// Immunization represents the event of a patient being administered a vaccine
type Immunization struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Immunization"

	// Business identifier
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// completed | entered-in-error | not-done
	Status ImmunizationStatus `json:"status"`

	// Reason not done
	StatusReason *common.CodeableConcept `json:"statusReason,omitempty"`

	// Vaccine product administered
	VaccineCode common.CodeableConcept `json:"vaccineCode"`

	// Who was immunized
	Patient common.Reference `json:"patient"`

	// Encounter immunization was part of
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Vaccine administration date
	OccurrenceDateTime *string `json:"occurrenceDateTime,omitempty"`
	OccurrenceString   *string `json:"occurrenceString,omitempty"`

	// When the immunization was first captured in the subject's record
	Recorded *string `json:"recorded,omitempty"`

	// Indicates if record is from primary source
	PrimarySource *bool `json:"primarySource,omitempty"`

	// Indicates the source of the vaccine actually administered
	ReportOrigin *common.CodeableConcept `json:"reportOrigin,omitempty"`

	// Where immunization occurred
	Location *common.Reference `json:"location,omitempty"`

	// Vaccine manufacturer
	Manufacturer *common.Reference `json:"manufacturer,omitempty"`

	// Vaccine lot number
	LotNumber *string `json:"lotNumber,omitempty"`

	// Vaccine expiration date
	ExpirationDate *string `json:"expirationDate,omitempty"`

	// Body site vaccine was administered
	Site *common.CodeableConcept `json:"site,omitempty"`

	// How vaccine entered body
	Route *common.CodeableConcept `json:"route,omitempty"`

	// Amount of vaccine administered
	DoseQuantity *common.Quantity `json:"doseQuantity,omitempty"`

	// Who performed event
	Performer []ImmunizationPerformer `json:"performer,omitempty"`

	// Additional immunization notes
	Note []Annotation `json:"note,omitempty"`

	// Why immunization occurred
	ReasonCode []common.CodeableConcept `json:"reasonCode,omitempty"`

	// Why immunization occurred
	ReasonReference []common.Reference `json:"reasonReference,omitempty"`

	// Dose potency
	IsSubpotent *bool `json:"isSubpotent,omitempty"`

	// Reason for being subpotent
	SubpotentReason []common.CodeableConcept `json:"subpotentReason,omitempty"`

	// Educational material presented to patient
	Education []ImmunizationEducation `json:"education,omitempty"`

	// Patient eligibility for a vaccination program
	ProgramEligibility []common.CodeableConcept `json:"programEligibility,omitempty"`

	// Funding source for the vaccine
	FundingSource *common.CodeableConcept `json:"fundingSource,omitempty"`

	// Details of a reaction that follows immunization
	Reaction []ImmunizationReaction `json:"reaction,omitempty"`

	// Protocol followed by the provider
	ProtocolApplied []ImmunizationProtocolApplied `json:"protocolApplied,omitempty"`
}

// ImmunizationStatus represents the current status of the vaccination event
type ImmunizationStatus string

const (
	ImmunizationStatusCompleted       ImmunizationStatus = "completed"
	ImmunizationStatusEnteredInError  ImmunizationStatus = "entered-in-error"
	ImmunizationStatusNotDone         ImmunizationStatus = "not-done"
)

// ImmunizationPerformer represents who performed the immunization event
type ImmunizationPerformer struct {
	common.BackboneElement

	// What type of performance was done
	Function *common.CodeableConcept `json:"function,omitempty"`

	// Individual or organization who was performing
	Actor common.Reference `json:"actor"`
}

// ImmunizationEducation represents educational material presented to the patient
type ImmunizationEducation struct {
	common.BackboneElement

	// Educational material document identifier
	DocumentType *string `json:"documentType,omitempty"`

	// Educational material reference pointer
	Reference *string `json:"reference,omitempty"`

	// Educational material publication date
	PublicationDate *string `json:"publicationDate,omitempty"`

	// Educational material presentation date
	PresentationDate *string `json:"presentationDate,omitempty"`
}

// ImmunizationReaction represents details of a reaction following immunization
type ImmunizationReaction struct {
	common.BackboneElement

	// When reaction started
	Date *string `json:"date,omitempty"`

	// Additional information on reaction
	Detail *common.Reference `json:"detail,omitempty"`

	// Indicates self-reported reaction
	Reported *bool `json:"reported,omitempty"`
}

// ImmunizationProtocolApplied represents protocol followed by the provider
type ImmunizationProtocolApplied struct {
	common.BackboneElement

	// Name of vaccine series
	Series *string `json:"series,omitempty"`

	// Who is responsible for publishing the recommendations
	Authority *common.Reference `json:"authority,omitempty"`

	// Vaccine preventatable disease being targeted
	TargetDisease []common.CodeableConcept `json:"targetDisease,omitempty"`

	// Dose number within series
	DoseNumberPositiveInt *int    `json:"doseNumberPositiveInt,omitempty"`
	DoseNumberString      *string `json:"doseNumberString,omitempty"`

	// Recommended number of doses for immunity
	SeriesDosesPositiveInt *int    `json:"seriesDosesPositiveInt,omitempty"`
	SeriesDosesString      *string `json:"seriesDosesString,omitempty"`
}

// CarePlan represents the intention of how one or more practitioners intend to deliver care
type CarePlan struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "CarePlan"

	// External Ids for this plan
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Instantiates FHIR protocol or definition
	InstantiatesCanonical []string `json:"instantiatesCanonical,omitempty"`

	// Instantiates external protocol or definition
	InstantiatesUri []string `json:"instantiatesUri,omitempty"`

	// Fulfills CarePlan
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// CarePlans with some sort of formal relationship to the current plan
	Replaces []common.Reference `json:"replaces,omitempty"`

	// Part of referenced CarePlan
	PartOf []common.Reference `json:"partOf,omitempty"`

	// draft | active | on-hold | revoked | completed | entered-in-error | unknown
	Status CarePlanStatus `json:"status"`

	// proposal | plan | order | option
	Intent CarePlanIntent `json:"intent"`

	// Type of plan
	Category []common.CodeableConcept `json:"category,omitempty"`

	// Human-friendly name for the care plan
	Title *string `json:"title,omitempty"`

	// Summary of nature of plan
	Description *string `json:"description,omitempty"`

	// Who the care plan is for
	Subject common.Reference `json:"subject"`

	// Encounter created as part of
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Time period plan covers
	Period *common.Period `json:"period,omitempty"`

	// Date record was first recorded
	Created *string `json:"created,omitempty"`

	// Who is the designated responsible party
	Author *common.Reference `json:"author,omitempty"`

	// Who provided the content of the care plan
	Contributor []common.Reference `json:"contributor,omitempty"`

	// Who's involved in plan?
	CareTeam []common.Reference `json:"careTeam,omitempty"`

	// Health issues this plan addresses
	Addresses []common.Reference `json:"addresses,omitempty"`

	// Information considered as part of plan
	SupportingInfo []common.Reference `json:"supportingInfo,omitempty"`

	// Desired outcome of plan
	Goal []common.Reference `json:"goal,omitempty"`

	// Action to occur as part of plan
	Activity []CarePlanActivity `json:"activity,omitempty"`

	// Comments about the plan
	Note []Annotation `json:"note,omitempty"`
}

// CarePlanStatus represents the workflow state of the care plan
type CarePlanStatus string

const (
	CarePlanStatusDraft         CarePlanStatus = "draft"
	CarePlanStatusActive        CarePlanStatus = "active"
	CarePlanStatusOnHold        CarePlanStatus = "on-hold"
	CarePlanStatusRevoked       CarePlanStatus = "revoked"
	CarePlanStatusCompleted     CarePlanStatus = "completed"
	CarePlanStatusEnteredInError CarePlanStatus = "entered-in-error"
	CarePlanStatusUnknown       CarePlanStatus = "unknown"
)

// CarePlanIntent represents the level of authority/intentionality
type CarePlanIntent string

const (
	CarePlanIntentProposal CarePlanIntent = "proposal"
	CarePlanIntentPlan     CarePlanIntent = "plan"
	CarePlanIntentOrder    CarePlanIntent = "order"
	CarePlanIntentOption   CarePlanIntent = "option"
)

// CarePlanActivity represents action to occur as part of plan
type CarePlanActivity struct {
	common.BackboneElement

	// Results of the activity
	OutcomeCodeableConcept []common.CodeableConcept `json:"outcomeCodeableConcept,omitempty"`

	// Appointment, Encounter, Procedure, etc.
	OutcomeReference []common.Reference `json:"outcomeReference,omitempty"`

	// Comments about the activity status/progress
	Progress []Annotation `json:"progress,omitempty"`

	// Activity details defined in specific resource
	Reference *common.Reference `json:"reference,omitempty"`

	// In-line definition of activity
	Detail *CarePlanActivityDetail `json:"detail,omitempty"`
}

// CarePlanActivityDetail represents in-line definition of activity
type CarePlanActivityDetail struct {
	common.BackboneElement

	// Kind of resource
	Kind *CarePlanActivityKind `json:"kind,omitempty"`

	// Instantiates FHIR protocol or definition
	InstantiatesCanonical []string `json:"instantiatesCanonical,omitempty"`

	// Instantiates external protocol or definition
	InstantiatesUri []string `json:"instantiatesUri,omitempty"`

	// Detail type of activity
	Code *common.CodeableConcept `json:"code,omitempty"`

	// Why activity should be done or why activity was prohibited
	ReasonCode []common.CodeableConcept `json:"reasonCode,omitempty"`

	// Why activity is needed
	ReasonReference []common.Reference `json:"reasonReference,omitempty"`

	// Goals this activity relates to
	Goal []common.Reference `json:"goal,omitempty"`

	// not-started | scheduled | in-progress | on-hold | completed | cancelled | stopped | unknown | entered-in-error
	Status CarePlanActivityStatus `json:"status"`

	// Reason for current status
	StatusReason *common.CodeableConcept `json:"statusReason,omitempty"`

	// If true, activity is prohibiting action
	DoNotPerform *bool `json:"doNotPerform,omitempty"`

	// When activity is to occur
	ScheduledTiming  *Timing        `json:"scheduledTiming,omitempty"`
	ScheduledPeriod  *common.Period `json:"scheduledPeriod,omitempty"`
	ScheduledString  *string        `json:"scheduledString,omitempty"`

	// Where it should happen
	Location *common.Reference `json:"location,omitempty"`

	// Who will be responsible?
	Performer []common.Reference `json:"performer,omitempty"`

	// What is administered/supplied
	ProductCodeableConcept *common.CodeableConcept `json:"productCodeableConcept,omitempty"`
	ProductReference       *common.Reference       `json:"productReference,omitempty"`

	// How to consume/day?
	DailyAmount *common.Quantity `json:"dailyAmount,omitempty"`

	// How much to administer/supply/consume
	Quantity *common.Quantity `json:"quantity,omitempty"`

	// Extra info describing activity to perform
	Description *string `json:"description,omitempty"`
}

// CarePlanActivityKind represents the kind of resource the activity is
type CarePlanActivityKind string

const (
	CarePlanActivityKindAppointment       CarePlanActivityKind = "Appointment"
	CarePlanActivityKindCommunicationRequest CarePlanActivityKind = "CommunicationRequest"
	CarePlanActivityKindDeviceRequest     CarePlanActivityKind = "DeviceRequest"
	CarePlanActivityKindMedicationRequest CarePlanActivityKind = "MedicationRequest"
	CarePlanActivityKindNutritionOrder    CarePlanActivityKind = "NutritionOrder"
	CarePlanActivityKindTask              CarePlanActivityKind = "Task"
	CarePlanActivityKindServiceRequest    CarePlanActivityKind = "ServiceRequest"
	CarePlanActivityKindVisionPrescription CarePlanActivityKind = "VisionPrescription"
)

// CarePlanActivityStatus represents the current state of the activity
type CarePlanActivityStatus string

const (
	CarePlanActivityStatusNotStarted     CarePlanActivityStatus = "not-started"
	CarePlanActivityStatusScheduled      CarePlanActivityStatus = "scheduled"
	CarePlanActivityStatusInProgress     CarePlanActivityStatus = "in-progress"
	CarePlanActivityStatusOnHold         CarePlanActivityStatus = "on-hold"
	CarePlanActivityStatusCompleted      CarePlanActivityStatus = "completed"
	CarePlanActivityStatusCancelled      CarePlanActivityStatus = "cancelled"
	CarePlanActivityStatusStopped        CarePlanActivityStatus = "stopped"
	CarePlanActivityStatusUnknown        CarePlanActivityStatus = "unknown"
	CarePlanActivityStatusEnteredInError CarePlanActivityStatus = "entered-in-error"
)

// Specimen represents a sample to be used for analysis
type Specimen struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Specimen"

	// External Identifier
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Identifier assigned by the lab
	AccessionIdentifier *common.Identifier `json:"accessionIdentifier,omitempty"`

	// available | unavailable | unsatisfactory | entered-in-error
	Status *SpecimenStatus `json:"status,omitempty"`

	// Kind of material that forms the specimen
	Type *common.CodeableConcept `json:"type,omitempty"`

	// Where the specimen came from
	Subject *common.Reference `json:"subject,omitempty"`

	// Time when specimen was received for processing
	ReceivedTime *string `json:"receivedTime,omitempty"`

	// Specimen from which this specimen originated
	Parent []common.Reference `json:"parent,omitempty"`

	// Why the specimen was collected
	Request []common.Reference `json:"request,omitempty"`

	// Collection details
	Collection *SpecimenCollection `json:"collection,omitempty"`

	// Processing and processing step details
	Processing []SpecimenProcessing `json:"processing,omitempty"`

	// The container holding the specimen
	Container []SpecimenContainer `json:"container,omitempty"`

	// The specimen's condition, quality, and adequacy
	Condition []common.CodeableConcept `json:"condition,omitempty"`

	// Comments
	Note []Annotation `json:"note,omitempty"`
}

// SpecimenStatus represents the availability of the specimen
type SpecimenStatus string

const (
	SpecimenStatusAvailable      SpecimenStatus = "available"
	SpecimenStatusUnavailable    SpecimenStatus = "unavailable"
	SpecimenStatusUnsatisfactory SpecimenStatus = "unsatisfactory"
	SpecimenStatusEnteredInError SpecimenStatus = "entered-in-error"
)

// SpecimenCollection represents details concerning the specimen collection
type SpecimenCollection struct {
	common.BackboneElement

	// Who collected the specimen
	Collector *common.Reference `json:"collector,omitempty"`

	// Collection time
	CollectedDateTime *string        `json:"collectedDateTime,omitempty"`
	CollectedPeriod   *common.Period `json:"collectedPeriod,omitempty"`

	// The span of time over which the collection of a specimen occurred
	Duration *Duration `json:"duration,omitempty"`

	// How much was collected
	Quantity *common.Quantity `json:"quantity,omitempty"`

	// Technique used to perform collection
	Method *common.CodeableConcept `json:"method,omitempty"`

	// Anatomical collection site
	BodySite *common.CodeableConcept `json:"bodySite,omitempty"`

	// Whether or how long patient abstained from food and/or drink
	FastingStatusCodeableConcept *common.CodeableConcept `json:"fastingStatusCodeableConcept,omitempty"`
	FastingStatusDuration        *Duration               `json:"fastingStatusDuration,omitempty"`
}

// SpecimenProcessing represents processing and processing steps for the specimen
type SpecimenProcessing struct {
	common.BackboneElement

	// Textual description of procedure
	Description *string `json:"description,omitempty"`

	// Indicates the treatment step applied to the specimen
	Procedure *common.CodeableConcept `json:"procedure,omitempty"`

	// Material used in the processing step
	Additive []common.Reference `json:"additive,omitempty"`

	// Date and time of specimen processing
	TimeDateTime *string        `json:"timeDateTime,omitempty"`
	TimePeriod   *common.Period `json:"timePeriod,omitempty"`
}

// SpecimenContainer represents the container holding the specimen
type SpecimenContainer struct {
	common.BackboneElement

	// Id for the container
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Textual description of the container
	Description *string `json:"description,omitempty"`

	// Kind of container directly associated with specimen
	Type *common.CodeableConcept `json:"type,omitempty"`

	// Container volume or size
	Capacity *common.Quantity `json:"capacity,omitempty"`

	// Quantity of specimen within container
	SpecimenQuantity *common.Quantity `json:"specimenQuantity,omitempty"`

	// Additive associated with container
	AdditiveCodeableConcept *common.CodeableConcept `json:"additiveCodeableConcept,omitempty"`
	AdditiveReference       *common.Reference       `json:"additiveReference,omitempty"`
}
