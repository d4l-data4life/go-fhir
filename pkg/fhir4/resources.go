// Package fhir4 contains FHIR R4 (version 4.0.1) resource definitions
package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
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
	ResearchSubjectStatusCandidate           ResearchSubjectStatus = "candidate"
	ResearchSubjectStatusEligible            ResearchSubjectStatus = "eligible"
	ResearchSubjectStatusFollowUp            ResearchSubjectStatus = "follow-up"
	ResearchSubjectStatusIneligible          ResearchSubjectStatus = "ineligible"
	ResearchSubjectStatusNotRegistered       ResearchSubjectStatus = "not-registered"
	ResearchSubjectStatusOffStudy            ResearchSubjectStatus = "off-study"
	ResearchSubjectStatusOnStudy             ResearchSubjectStatus = "on-study"
	ResearchSubjectStatusOnStudyIntervention ResearchSubjectStatus = "on-study-intervention"
	ResearchSubjectStatusOnStudyObservation  ResearchSubjectStatus = "on-study-observation"
	ResearchSubjectStatusPendingOnStudy      ResearchSubjectStatus = "pending-on-study"
	ResearchSubjectStatusPotentialCandidate  ResearchSubjectStatus = "potential-candidate"
	ResearchSubjectStatusScreening           ResearchSubjectStatus = "screening"
	ResearchSubjectStatusWithdrawn           ResearchSubjectStatus = "withdrawn"
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
	ConsentStatusDraft          ConsentStatus = "draft"
	ConsentStatusProposed       ConsentStatus = "proposed"
	ConsentStatusActive         ConsentStatus = "active"
	ConsentStatusRejected       ConsentStatus = "rejected"
	ConsentStatusInactive       ConsentStatus = "inactive"
	ConsentStatusEnteredInError ConsentStatus = "entered-in-error"
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
