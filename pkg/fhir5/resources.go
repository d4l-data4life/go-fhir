// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// VirtualServiceDetail represents virtual service contact details
type VirtualServiceDetail struct {
	DataType

	// Contact details for virtual service
	ContactDetail *ContactDetail `json:"contactDetail,omitempty"`

	// Virtual service type
	VirtualService *common.CodeableConcept `json:"virtualService,omitempty"`
}

// CodeableReference represents a reference that may be typed
type CodeableReference struct {
	DataType

	// Type of reference
	Concept *common.CodeableConcept `json:"concept,omitempty"`

	// Reference to another resource
	Reference *common.Reference `json:"reference,omitempty"`
}

// ContactDetail represents contact information
type ContactDetail struct {
	DataType

	// Name of the contact
	Name *string `json:"name,omitempty"`

	// Contact details
	Telecom []ContactPoint `json:"telecom,omitempty"`
}

// ContactPoint represents details for contacting
type ContactPoint struct {
	DataType

	// phone | fax | email | pager | url | sms | other
	System *ContactPointSystem `json:"system,omitempty"`

	// The actual contact point details
	Value *string `json:"value,omitempty"`

	// home | work | temp | old | mobile - purpose of this contact point
	Use *ContactPointUse `json:"use,omitempty"`

	// Specify preferred order of use (1 = highest)
	Rank *int `json:"rank,omitempty"`

	// Time period when the contact point was/is in use
	Period *common.Period `json:"period,omitempty"`
}

// ContactPointSystem represents the contact point system
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

// ContactPointUse represents the use of a contact point
type ContactPointUse string

const (
	ContactPointUseHome   ContactPointUse = "home"
	ContactPointUseWork   ContactPointUse = "work"
	ContactPointUseTemp   ContactPointUse = "temp"
	ContactPointUseOld    ContactPointUse = "old"
	ContactPointUseMobile ContactPointUse = "mobile"
)

// Resource is the base definition for all resources (R5 version)
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

// DomainResource is a resource that includes narrative, extensions, and contained resources (R5 version)
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

// Meta represents metadata about a resource (R5 version)
type Meta struct {
	DataType

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

// Narrative represents human-readable summary of the resource (R5 version)
type Narrative struct {
	DataType

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

// Helper functions for pointer creation
func StringPtr(s string) *string    { return &s }
func BoolPtr(b bool) *bool          { return &b }
func IntPtr(i int) *int             { return &i }
func Float64Ptr(f float64) *float64 { return &f }

// Dosage represents how medication should be taken
type Dosage struct {
	common.Element

	// Free text dosage instructions
	Text *string `json:"text,omitempty"`

	// Patient or consumer oriented instructions
	PatientInstruction *string `json:"patientInstruction,omitempty"`

	// Timing for the dosage
	Timing *Timing `json:"timing,omitempty"`

	// Body site to administer to
	Site *common.CodeableConcept `json:"site,omitempty"`

	// How drug should enter body
	Route *common.CodeableConcept `json:"route,omitempty"`

	// Technique for administering medication
	Method *common.CodeableConcept `json:"method,omitempty"`

	// Amount of medication per dose
	DoseAndRate []DosageDoseAndRate `json:"doseAndRate,omitempty"`
}

// DosageDoseAndRate represents dosage dose and rate information
type DosageDoseAndRate struct {
	common.Element

	// The kind of dose or rate specified
	Type *common.CodeableConcept `json:"type,omitempty"`

	// Amount of medication per dose
	DoseRange    *Range           `json:"doseRange,omitempty"`
	DoseQuantity *common.Quantity `json:"doseQuantity,omitempty"`

	// Amount of medication per unit of time
	RateRatio    *Ratio           `json:"rateRatio,omitempty"`
	RateRange    *Range           `json:"rateRange,omitempty"`
	RateQuantity *common.Quantity `json:"rateQuantity,omitempty"`
}

// Timing represents when medication should be administered
type Timing struct {
	DataType

	// When the event occurs
	Event []*string `json:"event,omitempty"`

	// When the event is to occur
	Repeat *TimingRepeat `json:"repeat,omitempty"`

	// BID | TID | QID | AM | PM | QD | QOD | Q4H | Q6H +
	Code *common.CodeableConcept `json:"code,omitempty"`
}

// TimingRepeat represents timing repeat information
type TimingRepeat struct {
	common.Element

	// Length/Range of lengths, or (Start and/or end) limits
	BoundsDuration *Duration      `json:"boundsDuration,omitempty"`
	BoundsRange    *Range         `json:"boundsRange,omitempty"`
	BoundsPeriod   *common.Period `json:"boundsPeriod,omitempty"`

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
	TimeOfDay []*string `json:"timeOfDay,omitempty"`

	// Code for time period of occurrence
	When []string `json:"when,omitempty"`

	// Minutes from event (before or after)
	Offset *int `json:"offset,omitempty"`
}

// Expression represents a single FHIRPath expression
type Expression struct {
	common.Element

	// Natural language description of the condition
	Description *string `json:"description,omitempty"`

	// Short name assigned to expression for reuse
	Name *string `json:"name,omitempty"`

	// text/cql | text/fhirpath | application/x-fhir-query | text/cql-identifier | text/cql-expression | etc.
	Language string `json:"language"`

	// Expression in specified language
	Expression *string `json:"expression,omitempty"`

	// Where the expression is found
	Reference *string `json:"reference,omitempty"`
}

// RelatedArtifact represents related artifacts such as additional documentation, justification, or bibliographic references
type RelatedArtifact struct {
	common.Element

	// documentation | justification | citation | predecessor | successor | derived-from | depends-on | composed-of | part-of | amends | amended-with | appends | appended-with | cites | cited-by | comments-on | comment-in | contains | contained-in | corrects | correction-in | replaces | replaced-with | retracts | retracted-by | signs | similar-to | supports | supported-with | transforms | transformed-by | transformed-with | depends-on | endorses
	Type RelatedArtifactType `json:"type"`

	// Additional classifiers
	Classifier []common.CodeableConcept `json:"classifier,omitempty"`

	// Short label
	Label *string `json:"label,omitempty"`

	// Brief description of the related artifact
	Display *string `json:"display,omitempty"`

	// Bibliographic citation for the artifact
	Citation *string `json:"citation,omitempty"`

	// What document is being referenced
	Document *Attachment `json:"document,omitempty"`

	// What resource is being referenced
	Resource *string `json:"resource,omitempty"`

	// What resource type is being referenced
	ResourceReference *common.Reference `json:"resourceReference,omitempty"`

	// Publication year
	PublicationYear *int `json:"publicationYear,omitempty"`

	// Publication date
	PublicationDate *string `json:"publicationDate,omitempty"`
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
	RelatedArtifactTypePartOf        RelatedArtifactType = "part-of"
)

// UsageContext represents the context of use for a resource
type UsageContext struct {
	common.Element

	// Type of context being specified
	Code common.Coding `json:"code"`

	// Value that defines the context
	ValueCodeableConcept *common.CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueQuantity        *common.Quantity        `json:"valueQuantity,omitempty"`
	ValueRange           *Range                  `json:"valueRange,omitempty"`
	ValueReference       *common.Reference       `json:"valueReference,omitempty"`
}

// RequestIntent represents the intent of the request
type RequestIntent string

const (
	RequestIntentProposal      RequestIntent = "proposal"
	RequestIntentPlan          RequestIntent = "plan"
	RequestIntentDirective     RequestIntent = "directive"
	RequestIntentOrder         RequestIntent = "order"
	RequestIntentOriginalOrder RequestIntent = "original-order"
	RequestIntentReflexOrder   RequestIntent = "reflex-order"
	RequestIntentFillerOrder   RequestIntent = "filler-order"
	RequestIntentInstanceOrder RequestIntent = "instance-order"
	RequestIntentOption        RequestIntent = "option"
)

// RequestPriority represents the priority of the request
type RequestPriority string

const (
	RequestPriorityRoutine RequestPriority = "routine"
	RequestPriorityUrgent  RequestPriority = "urgent"
	RequestPriorityASAP    RequestPriority = "asap"
	RequestPriorityStat    RequestPriority = "stat"
)

// RequestStatus represents the status of the request
type RequestStatus string

const (
	RequestStatusDraft   RequestStatus = "draft"
	RequestStatusActive  RequestStatus = "active"
	RequestStatusRetired RequestStatus = "retired"
	RequestStatusUnknown RequestStatus = "unknown"
)

// ParticipantType represents the type of participant
type ParticipantType string

const (
	ParticipantTypeCareTeam          ParticipantType = "careteam"
	ParticipantTypeDevice            ParticipantType = "device"
	ParticipantTypeGroup             ParticipantType = "group"
	ParticipantTypeHealthcareService ParticipantType = "healthcareservice"
	ParticipantTypeLocation          ParticipantType = "location"
	ParticipantTypeOrganization      ParticipantType = "organization"
	ParticipantTypePatient           ParticipantType = "patient"
	ParticipantTypePractitioner      ParticipantType = "practitioner"
	ParticipantTypePractitionerRole  ParticipantType = "practitionerrole"
	ParticipantTypeRelatedPerson     ParticipantType = "relatedperson"
)

// CapabilityStatementKind represents the way that this statement is intended to be used

// EpisodeOfCareStatus represents the status of an episode of care

// EvidenceVariableDefinition represents evidence variable definition
type EvidenceVariableDefinition struct {
	common.BackboneElement

	// A text description or summary of the variable
	Description *string `json:"description,omitempty"`

	// Indication of quality of match between intended variable to actual variable
	DirectnessMatch *common.CodeableConcept `json:"directnessMatch,omitempty"`

	// Definition of the intended variable related to the Evidence
	Intended *common.Reference `json:"intended,omitempty"`

	// Footnotes and/or explanatory notes
	Note []Annotation `json:"note,omitempty"`

	// Definition of the actual variable related to the statistic(s)
	Observed *common.Reference `json:"observed,omitempty"`

	// population | subpopulation | exposure | referenceExposure | measuredVariable | confounder
	VariableRole common.CodeableConcept `json:"variableRole"`
}

// EvidenceStatisticSampleSize represents sample size information
type EvidenceStatisticSampleSize struct {
	common.BackboneElement

	// Human-readable summary of population sample size
	Description *string `json:"description,omitempty"`

	// Number of participants with known results for measured variables
	KnownDataCount *int `json:"knownDataCount,omitempty"`

	// Footnote or explanatory note about the sample size
	Note []Annotation `json:"note,omitempty"`

	// Number of participants in the population
	NumberOfParticipants *int `json:"numberOfParticipants,omitempty"`

	// Number of studies
	NumberOfStudies *int `json:"numberOfStudies,omitempty"`
}

// EvidenceStatisticAttributeEstimate represents attribute estimate information
type EvidenceStatisticAttributeEstimate struct {
	common.BackboneElement

	// A nested attribute estimate
	AttributeEstimate []EvidenceStatisticAttributeEstimate `json:"attributeEstimate,omitempty"`

	// Human-readable summary of the estimate
	Description *string `json:"description,omitempty"`

	// Use 95 for a 95% confidence interval
	Level *float64 `json:"level,omitempty"`

	// Footnote or explanatory note about the estimate
	Note []Annotation `json:"note,omitempty"`

	// Often the p value
	Quantity *common.Quantity `json:"quantity,omitempty"`

	// Lower bound of confidence interval
	Range *Range `json:"range,omitempty"`

	// The type of attribute estimate, e.g., confidence interval or p value
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// EvidenceStatisticModelCharacteristicVariable represents model characteristic variable
type EvidenceStatisticModelCharacteristicVariable struct {
	common.BackboneElement

	// How the variable is classified for use in adjusted analysis
	Handling *EvidenceVariableHandling `json:"handling,omitempty"`

	// Description for grouping of ordinal or polychotomous variables
	ValueCategory []common.CodeableConcept `json:"valueCategory,omitempty"`

	// Discrete value for grouping of ordinal or polychotomous variables
	ValueQuantity []common.Quantity `json:"valueQuantity,omitempty"`

	// Range of values for grouping of ordinal or polychotomous variables
	ValueRange []Range `json:"valueRange,omitempty"`

	// Description of the variable
	VariableDefinition common.Reference `json:"variableDefinition"`
}

// EvidenceVariableHandling represents how variables are handled in statistical analysis
type EvidenceVariableHandling string

const (
	EvidenceVariableHandlingContinuous    EvidenceVariableHandling = "continuous"
	EvidenceVariableHandlingDichotomous   EvidenceVariableHandling = "dichotomous"
	EvidenceVariableHandlingOrdinal       EvidenceVariableHandling = "ordinal"
	EvidenceVariableHandlingPolychotomous EvidenceVariableHandling = "polychotomous"
)

// EvidenceStatisticModelCharacteristic represents model characteristic
type EvidenceStatisticModelCharacteristic struct {
	common.BackboneElement

	// An attribute of the statistic used as a model characteristic
	AttributeEstimate []EvidenceStatisticAttributeEstimate `json:"attributeEstimate,omitempty"`

	// Description of a component of the method to generate the statistic
	Code common.CodeableConcept `json:"code"`

	// Further specification of the quantified value of the component
	Value *common.Quantity `json:"value,omitempty"`

	// A variable adjusted for in the adjusted analysis
	Variable []EvidenceStatisticModelCharacteristicVariable `json:"variable,omitempty"`
}

// EvidenceStatistic represents values and parameters for a single statistic
type EvidenceStatistic struct {
	common.BackboneElement

	// A statistical attribute of the statistic such as a measure of heterogeneity
	AttributeEstimate []EvidenceStatisticAttributeEstimate `json:"attributeEstimate,omitempty"`

	// Simple strings can be used for descriptive purposes
	Category *common.CodeableConcept `json:"category,omitempty"`

	// A description of the content value of the statistic
	Description *string `json:"description,omitempty"`

	// A component of the method to generate the statistic
	ModelCharacteristic []EvidenceStatisticModelCharacteristic `json:"modelCharacteristic,omitempty"`

	// Footnotes and/or explanatory notes
	Note []Annotation `json:"note,omitempty"`

	// Number of participants with events
	NumberAffected *int `json:"numberAffected,omitempty"`

	// Number of events
	NumberOfEvents *int `json:"numberOfEvents,omitempty"`

	// Statistic value
	Quantity *common.Quantity `json:"quantity,omitempty"`

	// Number of samples in the statistic
	SampleSize *EvidenceStatisticSampleSize `json:"sampleSize,omitempty"`

	// Type of statistic, e.g., relative risk
	StatisticType *common.CodeableConcept `json:"statisticType,omitempty"`
}

// EvidenceCertainty represents assessment of certainty
type EvidenceCertainty struct {
	common.BackboneElement

	// Textual description of certainty
	Description *string `json:"description,omitempty"`

	// Footnotes and/or explanatory notes
	Note []Annotation `json:"note,omitempty"`

	// Individual or group who did the rating
	Rater *string `json:"rater,omitempty"`

	// Assessment or judgement of the aspect
	Rating *common.CodeableConcept `json:"rating,omitempty"`

	// A domain or subdomain of certainty
	Subcomponent []EvidenceCertainty `json:"subcomponent,omitempty"`

	// Aspect of certainty being rated
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// EvidenceVariableCharacteristicDefinitionByCombination represents combination definition
type EvidenceVariableCharacteristicDefinitionByCombination struct {
	common.BackboneElement

	// Provide a code for the characteristic
	Code common.CodeableConcept `json:"code"`

	// Describes the characteristic
	Characteristic []EvidenceVariableCharacteristic `json:"characteristic"`
}

// EvidenceVariableCharacteristicDefinitionByTypeAndValue represents type and value definition
type EvidenceVariableCharacteristicDefinitionByTypeAndValue struct {
	common.BackboneElement

	// Defines the characteristic when paired with characteristic.definitionByTypeAndValue.value[x]
	Type common.CodeableConcept `json:"type"`

	// Defines the characteristic when paired with characteristic.definitionByTypeAndValue.type
	ValueCodeableConcept *common.CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueBoolean         *bool                   `json:"valueBoolean,omitempty"`
	ValueQuantity        *common.Quantity        `json:"valueQuantity,omitempty"`
	ValueRange           *Range                  `json:"valueRange,omitempty"`
	ValueReference       *common.Reference       `json:"valueReference,omitempty"`
	ValueId              *string                 `json:"valueId,omitempty"`
}

// EvidenceVariableCharacteristicTimeFromEvent represents time from event
type EvidenceVariableCharacteristicTimeFromEvent struct {
	common.BackboneElement

	// Human readable description
	Description *string `json:"description,omitempty"`

	// The event used as a base point (reference point) in time
	EventCodeableConcept *common.CodeableConcept `json:"eventCodeableConcept,omitempty"`
	EventReference       *common.Reference       `json:"eventReference,omitempty"`
	EventDateTime        *string                 `json:"eventDateTime,omitempty"`
	EventId              *string                 `json:"eventId,omitempty"`

	// Used to express the observation at a defined number of time units after the defined event
	Quantity *common.Quantity `json:"quantity,omitempty"`
	Range    *Range           `json:"range,omitempty"`
}

// EvidenceVariableCharacteristic represents characteristics of the evidence variable
type EvidenceVariableCharacteristic struct {
	common.BackboneElement

	// Defines the characteristic as a combination of two or more characteristics
	DefinitionByCombination *EvidenceVariableCharacteristicDefinitionByCombination `json:"definitionByCombination,omitempty"`

	// Defines the characteristic using both a type and value[x] elements
	DefinitionByTypeAndValue *EvidenceVariableCharacteristicDefinitionByTypeAndValue `json:"definitionByTypeAndValue,omitempty"`

	// Defines the characteristic using Canonical
	DefinitionCanonical *string `json:"definitionCanonical,omitempty"`

	// Defines the characteristic using CodeableConcept
	DefinitionCodeableConcept *common.CodeableConcept `json:"definitionCodeableConcept,omitempty"`

	// Defines the characteristic using Expression
	DefinitionExpression *Expression `json:"definitionExpression,omitempty"`

	// Defines the characteristic using id
	DefinitionId *string `json:"definitionId,omitempty"`

	// Defines the characteristic using Reference
	DefinitionReference *common.Reference `json:"definitionReference,omitempty"`

	// A short, natural language description of the characteristic
	Description *string `json:"description,omitempty"`

	// Device used for determining characteristic
	Device *common.Reference `json:"device,omitempty"`

	// When true, members with this characteristic are excluded from the element
	Exclude *bool `json:"exclude,omitempty"`

	// The grouping characteristic
	GroupMeasure *common.CodeableConcept `json:"groupMeasure,omitempty"`

	// Defines the characteristic using instances
	InstancesQuantity *common.Quantity `json:"instancesQuantity,omitempty"`
	InstancesRange    *Range           `json:"instancesRange,omitempty"`

	// The intended method of use for the characteristic
	Method []common.CodeableConcept `json:"method,omitempty"`

	// Length of time in which the characteristic is met
	DurationQuantity *common.Quantity `json:"durationQuantity,omitempty"`
	DurationRange    *Range           `json:"durationRange,omitempty"`

	// Timing for the characteristic
	TimeFromEvent []EvidenceVariableCharacteristicTimeFromEvent `json:"timeFromEvent,omitempty"`
}

// EvidenceVariableCategory represents a grouping for ordinal or polychotomous variables
type EvidenceVariableCategory struct {
	common.BackboneElement

	// Description of the grouping
	Name *string `json:"name,omitempty"`

	// Definition of the grouping
	ValueCodeableConcept *common.CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueQuantity        *common.Quantity        `json:"valueQuantity,omitempty"`
	ValueRange           *Range                  `json:"valueRange,omitempty"`
}

// ImagingStudyStatus represents the status of an imaging study
type ImagingStudyStatus string

const (
	ImagingStudyStatusRegistered     ImagingStudyStatus = "registered"
	ImagingStudyStatusAvailable      ImagingStudyStatus = "available"
	ImagingStudyStatusCancelled      ImagingStudyStatus = "cancelled"
	ImagingStudyStatusEnteredInError ImagingStudyStatus = "entered-in-error"
	ImagingStudyStatusUnknown        ImagingStudyStatus = "unknown"
)

// ImagingStudySeriesPerformer represents the performer of an imaging study series
type ImagingStudySeriesPerformer struct {
	common.BackboneElement

	// Indicates who or what performed the series
	Actor common.Reference `json:"actor"`

	// Distinguishes the type of involvement of the performer in the series
	Function *common.CodeableConcept `json:"function,omitempty"`
}

// ImagingStudySeriesInstance represents a single SOP instance within the series
type ImagingStudySeriesInstance struct {
	common.BackboneElement

	// The number of instance in the series
	Number *int `json:"number,omitempty"`

	// DICOM instance type
	SopClass common.Coding `json:"sopClass"`

	// Particularly for post-acquisition analytic objects
	Title *string `json:"title,omitempty"`

	// See DICOM PS3.3 C.12.1
	Uid string `json:"uid"`
}

// ImagingStudySeries represents each study series of images or other content
type ImagingStudySeries struct {
	common.BackboneElement

	// The anatomic structures examined
	BodySite *CodeableReference `json:"bodySite,omitempty"`

	// A description of the series
	Description *string `json:"description,omitempty"`

	// Typical endpoint types include DICOM WADO-RS
	Endpoint []common.Reference `json:"endpoint,omitempty"`

	// A single SOP instance within the series
	Instance []ImagingStudySeriesInstance `json:"instance,omitempty"`

	// The laterality of the anatomic structures examined
	Laterality *common.CodeableConcept `json:"laterality,omitempty"`

	// The distinct modality for this series
	Modality common.CodeableConcept `json:"modality"`

	// The numeric identifier of this series in the study
	Number *int `json:"number,omitempty"`

	// Number of SOP Instances in the Study
	NumberOfInstances *int `json:"numberOfInstances,omitempty"`

	// If the person who performed the series is not known, their Organization may be recorded
	Performer []ImagingStudySeriesPerformer `json:"performer,omitempty"`

	// The specimen imaged
	Specimen []common.Reference `json:"specimen,omitempty"`

	// The date and time the series was started
	Started *string `json:"started,omitempty"`

	// See DICOM PS3.3 C.7.3
	Uid string `json:"uid"`
}

// ImagingStudy represents the content produced in a DICOM imaging study
type ImagingStudy struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "ImagingStudy"

	// A list of the diagnostic requests that resulted in this imaging study
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// The Imaging Manager description of the study
	Description *string `json:"description,omitempty"`

	// This will typically be the encounter the event occurred within
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Typical endpoint types include DICOM WADO-RS
	Endpoint []common.Reference `json:"endpoint,omitempty"`

	// See discussion under Imaging Study Implementation Notes
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The principal physical location where the ImagingStudy was performed
	Location *common.Reference `json:"location,omitempty"`

	// A list of all the distinct values of series.modality
	Modality []common.CodeableConcept `json:"modality,omitempty"`

	// Per the recommended DICOM mapping
	Note []Annotation `json:"note,omitempty"`

	// Number of SOP Instances in Study
	NumberOfInstances *int `json:"numberOfInstances,omitempty"`

	// Number of Series in the Study
	NumberOfSeries *int `json:"numberOfSeries,omitempty"`

	// To link an ImagingStudy to an Encounter use encounter
	PartOf []common.Reference `json:"partOf,omitempty"`

	// This field corresponds to the DICOM Procedure Code Sequence
	Procedure []CodeableReference `json:"procedure,omitempty"`

	// Description of clinical condition indicating why the ImagingStudy was requested
	Reason []CodeableReference `json:"reason,omitempty"`

	// The requesting/referring physician
	Referrer *common.Reference `json:"referrer,omitempty"`

	// Each study has one or more series of images or other content
	Series []ImagingStudySeries `json:"series,omitempty"`

	// Date and time the study started
	Started *string `json:"started,omitempty"`

	// Unknown does not represent "other" - one of the defined statuses must apply
	Status ImagingStudyStatus `json:"status"`

	// QA phantoms can be recorded with a Device; multiple subjects can be recorded with a Group
	Subject common.Reference `json:"subject"`
}

// ImmunizationStatus represents the status of an immunization
type ImmunizationStatus string

const (
	ImmunizationStatusCompleted      ImmunizationStatus = "completed"
	ImmunizationStatusEnteredInError ImmunizationStatus = "entered-in-error"
	ImmunizationStatusNotDone        ImmunizationStatus = "not-done"
)

// ImmunizationPerformer represents who performed the immunization event
type ImmunizationPerformer struct {
	common.BackboneElement

	// When the individual practitioner who performed the action is known, it is best to send
	Actor common.Reference `json:"actor"`

	// Describes the type of performance (e.g. ordering provider, administering provider, etc.)
	Function *common.CodeableConcept `json:"function,omitempty"`
}

// ImmunizationProgramEligibility represents a patient's eligibility for a funding program
type ImmunizationProgramEligibility struct {
	common.BackboneElement

	// Indicates which program the patient had their eligility evaluated for
	Program common.CodeableConcept `json:"program"`

	// Indicates the patient's eligility status for for a specific payment program
	ProgramStatus common.CodeableConcept `json:"programStatus"`
}

// ImmunizationReaction represents a reaction to the immunization
type ImmunizationReaction struct {
	common.BackboneElement

	// Date of reaction to the immunization
	Date *string `json:"date,omitempty"`

	// Details of the reaction
	Manifestation *CodeableReference `json:"manifestation,omitempty"`

	// Self-reported indicator
	Reported *bool `json:"reported,omitempty"`
}

// ImmunizationProtocolApplied represents the protocol being followed by the provider
type ImmunizationProtocolApplied struct {
	common.BackboneElement

	// Indicates the authority who published the protocol (e.g. ACIP) that is being followed
	Authority *common.Reference `json:"authority,omitempty"`

	// The use of an integer is preferred if known
	DoseNumber string `json:"doseNumber"`

	// One possible path to achieve presumed immunity against a disease
	Series *string `json:"series,omitempty"`

	// The use of an integer is preferred if known
	SeriesDoses *string `json:"seriesDoses,omitempty"`

	// The vaccine preventable disease the dose is being administered against
	TargetDisease []common.CodeableConcept `json:"targetDisease,omitempty"`
}

// Immunization represents the event of a patient being administered a vaccine
type Immunization struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Immunization"

	// An indication of which product was administered to the patient
	AdministeredProduct *CodeableReference `json:"administeredProduct,omitempty"`

	// Allows tracing of an authorization for the Immunization
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// The quantity of vaccine product that was administered
	DoseQuantity *common.Quantity `json:"doseQuantity,omitempty"`

	// The visit or admission or other contact between patient and health care provider
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Date vaccine batch expires
	ExpirationDate *string `json:"expirationDate,omitempty"`

	// Indicates the source of the vaccine actually administered
	FundingSource *common.CodeableConcept `json:"fundingSource,omitempty"`

	// A unique identifier assigned to this immunization record
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Typically will not be populated if primarySource = True
	InformationSource *CodeableReference `json:"informationSource,omitempty"`

	// Typically, the recognition of the dose being sub-potent is retrospective
	IsSubpotent *bool `json:"isSubpotent,omitempty"`

	// The service delivery location where the vaccine administration occurred
	Location *common.Reference `json:"location,omitempty"`

	// Lot number of the vaccine product
	LotNumber *string `json:"lotNumber,omitempty"`

	// Name of vaccine manufacturer
	Manufacturer *CodeableReference `json:"manufacturer,omitempty"`

	// Extra information about the immunization
	Note []Annotation `json:"note,omitempty"`

	// When immunizations are given a specific date and time should always be known
	OccurrenceDateTime *string `json:"occurrenceDateTime,omitempty"`
	OccurrenceString   *string `json:"occurrenceString,omitempty"`

	// The patient who either received or did not receive the immunization
	Patient common.Reference `json:"patient"`

	// Indicates who performed the immunization event
	Performer []ImmunizationPerformer `json:"performer,omitempty"`

	// Reflects the "reliability" of the content
	PrimarySource *bool `json:"primarySource,omitempty"`

	// Indicates a patient's eligibility for a funding program
	ProgramEligibility []ImmunizationProgramEligibility `json:"programEligibility,omitempty"`

	// The protocol (set of recommendations) being followed by the provider
	ProtocolApplied []ImmunizationProtocolApplied `json:"protocolApplied,omitempty"`

	// A reaction may be an indication of an allergy or intolerance
	Reaction []ImmunizationReaction `json:"reaction,omitempty"`

	// Describes why the immunization occurred in coded or textual form
	Reason []CodeableReference `json:"reason,omitempty"`

	// The path by which the vaccine product is taken into the body
	Route *common.CodeableConcept `json:"route,omitempty"`

	// Body site where vaccine was administered
	Site *common.CodeableConcept `json:"site,omitempty"`

	// Will generally be set to show that the immunization has been completed or not done
	Status ImmunizationStatus `json:"status"`

	// This is generally only used for the status of "not-done"
	StatusReason *common.CodeableConcept `json:"statusReason,omitempty"`

	// Reason why a dose is considered to be subpotent
	SubpotentReason []common.CodeableConcept `json:"subpotentReason,omitempty"`

	// Additional information that is relevant to the immunization
	SupportingInformation []common.Reference `json:"supportingInformation,omitempty"`

	// The code for the administered vaccine
	VaccineCode common.CodeableConcept `json:"vaccineCode"`
}

// ImmunizationEvaluationStatus represents the status of an immunization evaluation
type ImmunizationEvaluationStatus string

const (
	ImmunizationEvaluationStatusCompleted      ImmunizationEvaluationStatus = "completed"
	ImmunizationEvaluationStatusEnteredInError ImmunizationEvaluationStatus = "entered-in-error"
)

// ImmunizationEvaluation represents a comparison of an immunization event against published recommendations
type ImmunizationEvaluation struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "ImmunizationEvaluation"

	// Indicates the authority who published the protocol (e.g. ACIP)
	Authority *common.Reference `json:"authority,omitempty"`

	// The date the evaluation of the vaccine administration event was performed
	Date *string `json:"date,omitempty"`

	// Additional information about the evaluation
	Description *string `json:"description,omitempty"`

	// The use of an integer is preferred if known
	DoseNumber *string `json:"doseNumber,omitempty"`

	// Indicates if the dose is valid or not valid with respect to the published recommendations
	DoseStatus common.CodeableConcept `json:"doseStatus"`

	// Provides an explanation as to why the vaccine administration event is valid or not
	DoseStatusReason []common.CodeableConcept `json:"doseStatusReason,omitempty"`

	// A unique identifier assigned to this immunization evaluation record
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The vaccine administration event being evaluated
	ImmunizationEvent common.Reference `json:"immunizationEvent"`

	// The individual for whom the evaluation is being done
	Patient common.Reference `json:"patient"`

	// One possible path to achieve presumed immunity against a disease
	Series *string `json:"series,omitempty"`

	// The use of an integer is preferred if known
	SeriesDoses *string `json:"seriesDoses,omitempty"`

	// Indicates the current status of the evaluation
	Status ImmunizationEvaluationStatus `json:"status"`

	// The vaccine preventable disease the dose is being evaluated against
	TargetDisease common.CodeableConcept `json:"targetDisease"`
}

// ImmunizationRecommendationRecommendationDateCriterion represents date criteria for recommendations
type ImmunizationRecommendationRecommendationDateCriterion struct {
	common.BackboneElement

	// Date classification of recommendation
	Code common.CodeableConcept `json:"code"`

	// The date whose meaning is specified by dateCriterion.code
	Value string `json:"value"`
}

// ImmunizationRecommendationRecommendation represents a single recommended administration
type ImmunizationRecommendationRecommendation struct {
	common.BackboneElement

	// Vaccine(s) which should not be used to fulfill the recommendation
	ContraindicatedVaccineCode []common.CodeableConcept `json:"contraindicatedVaccineCode,omitempty"`

	// Vaccine date recommendations
	DateCriterion []ImmunizationRecommendationRecommendationDateCriterion `json:"dateCriterion,omitempty"`

	// Contains the description about the protocol under which the vaccine was administered
	Description *string `json:"description,omitempty"`

	// The use of an integer is preferred if known
	DoseNumber *string `json:"doseNumber,omitempty"`

	// The reason for the assigned forecast status
	ForecastReason []common.CodeableConcept `json:"forecastReason,omitempty"`

	// Indicates the patient status with respect to the path to immunity for the target disease
	ForecastStatus common.CodeableConcept `json:"forecastStatus"`

	// One possible path to achieve presumed immunity against a disease
	Series *string `json:"series,omitempty"`

	// The use of an integer is preferred if known
	SeriesDoses *string `json:"seriesDoses,omitempty"`

	// Immunization event history and/or evaluation that supports the status and recommendation
	SupportingImmunization []common.Reference `json:"supportingImmunization,omitempty"`

	// Patient Information that supports the status and recommendation
	SupportingPatientInformation []common.Reference `json:"supportingPatientInformation,omitempty"`

	// A given instance of the .recommendation backbone element should correspond to a single recommended administration
	TargetDisease []common.CodeableConcept `json:"targetDisease,omitempty"`

	// Vaccine(s) or vaccine group that pertain to the recommendation
	VaccineCode []common.CodeableConcept `json:"vaccineCode,omitempty"`
}

// ImmunizationRecommendation represents a patient's point-in-time set of recommendations
type ImmunizationRecommendation struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "ImmunizationRecommendation"

	// Indicates the authority who published the protocol (e.g. ACIP)
	Authority *common.Reference `json:"authority,omitempty"`

	// The date the immunization recommendation(s) were created
	Date string `json:"date"`

	// A unique identifier assigned to this particular recommendation record
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The patient the recommendation(s) are for
	Patient common.Reference `json:"patient"`

	// A given instance of the .recommendation backbone element should correspond to a single recommended administration
	Recommendation []ImmunizationRecommendationRecommendation `json:"recommendation"`
}

// ImplementationGuideStatus represents the status of an implementation guide
type ImplementationGuideStatus string

const (
	ImplementationGuideStatusDraft   ImplementationGuideStatus = "draft"
	ImplementationGuideStatusActive  ImplementationGuideStatus = "active"
	ImplementationGuideStatusRetired ImplementationGuideStatus = "retired"
	ImplementationGuideStatusUnknown ImplementationGuideStatus = "unknown"
)

// ImplementationGuideDefinitionPageGeneration represents how the page is generated
type ImplementationGuideDefinitionPageGeneration string

const (
	ImplementationGuideDefinitionPageGenerationHTML      ImplementationGuideDefinitionPageGeneration = "html"
	ImplementationGuideDefinitionPageGenerationMarkdown  ImplementationGuideDefinitionPageGeneration = "markdown"
	ImplementationGuideDefinitionPageGenerationXML       ImplementationGuideDefinitionPageGeneration = "xml"
	ImplementationGuideDefinitionPageGenerationGenerated ImplementationGuideDefinitionPageGeneration = "generated"
)

// ImplementationGuideDependsOn represents another implementation guide that this implementation depends on
type ImplementationGuideDependsOn struct {
	common.BackboneElement

	// The NPM package name for the Implementation Guide that this IG depends on
	PackageId *string `json:"packageId,omitempty"`

	// This doesn't need to enumerate every resource used, but should give some sense of why the dependency exists
	Reason *string `json:"reason,omitempty"`

	// Usually, A canonical reference to the implementation guide is the same as the master location
	Uri string `json:"uri"`

	// This follows the syntax of the NPM packaging version field
	Version *string `json:"version,omitempty"`
}

// ImplementationGuideGlobal represents global profiles for resources
type ImplementationGuideGlobal struct {
	common.BackboneElement

	// A reference to the profile that all instances must conform to
	Profile string `json:"profile"`

	// The type must match that of the profile that is referred to
	Type string `json:"type"`
}

// ImplementationGuideDefinitionGrouping represents groupings of content
type ImplementationGuideDefinitionGrouping struct {
	common.BackboneElement

	// Human readable text describing the package
	Description *string `json:"description,omitempty"`

	// The human-readable title to display for the package of resources
	Name string `json:"name"`
}

// ImplementationGuideDefinitionResource represents a resource that is part of the implementation guide
type ImplementationGuideDefinitionResource struct {
	common.BackboneElement

	// This is mostly used with examples to explain why it is present
	Description *string `json:"description,omitempty"`

	// The resource SHALL be valid against all the versions it is specified to apply to
	FhirVersion []string `json:"fhirVersion,omitempty"`

	// This must correspond to a group.id element within this implementation guide
	GroupingId *string `json:"groupingId,omitempty"`

	// If true, indicates the resource is an example instance
	IsExample *bool `json:"isExample,omitempty"`

	// A human assigned name for the resource
	Name *string `json:"name,omitempty"`

	// Examples: StructureDefinition -> Any, ValueSet -> expansion, etc.
	Profile []string `json:"profile,omitempty"`

	// Usually this is a relative URL that locates the resource within the implementation guide
	Reference common.Reference `json:"reference"`
}

// ImplementationGuideDefinitionPage represents pages in the implementation guide
type ImplementationGuideDefinitionPage struct {
	common.BackboneElement

	// A code that indicates how the page is generated
	Generation ImplementationGuideDefinitionPageGeneration `json:"generation"`

	// This SHALL be a local reference, expressed with respect to the root of the IG output folder
	Name string `json:"name"`

	// Sub-pages (recursive structure)
	Page []ImplementationGuideDefinitionPage `json:"page,omitempty"`

	// If absent and the page isn't a generated page, this may be inferred from the page name
	SourceUrl      *string `json:"sourceUrl,omitempty"`
	SourceString   *string `json:"sourceString,omitempty"`
	SourceMarkdown *string `json:"sourceMarkdown,omitempty"`

	// A short title used to represent this page in navigational structures
	Title string `json:"title"`
}

// ImplementationGuideDefinitionParameter represents parameters for the implementation guide
type ImplementationGuideDefinitionParameter struct {
	common.BackboneElement

	// A tool-specific code that defines the parameter
	Code common.Coding `json:"code"`

	// Value for named type
	Value string `json:"value"`
}

// ImplementationGuideDefinitionTemplate represents a template for building resources
type ImplementationGuideDefinitionTemplate struct {
	common.BackboneElement

	// Type of template specified
	Code string `json:"code"`

	// The scope in which the template applies
	Scope *string `json:"scope,omitempty"`

	// The source location for the template
	Source string `json:"source"`
}

// ImplementationGuideDefinition represents the definition of the implementation guide
type ImplementationGuideDefinition struct {
	common.BackboneElement

	// Groupings are arbitrary sub-divisions of content
	Grouping []ImplementationGuideDefinitionGrouping `json:"grouping,omitempty"`

	// Pages automatically become sections if they have sub-pages
	Page *ImplementationGuideDefinitionPage `json:"page,omitempty"`

	// Parameters for the implementation guide
	Parameter []ImplementationGuideDefinitionParameter `json:"parameter,omitempty"`

	// A resource that is part of the implementation guide
	Resource []ImplementationGuideDefinitionResource `json:"resource,omitempty"`

	// A template for building resources
	Template []ImplementationGuideDefinitionTemplate `json:"template,omitempty"`
}

// ImplementationGuideManifestResource represents a resource that is part of the implementation guide
type ImplementationGuideManifestResource struct {
	common.BackboneElement

	// If true, indicates the resource is an example instance
	IsExample *bool `json:"isExample,omitempty"`

	// Examples: StructureDefinition -> Any, ValueSet -> expansion, etc.
	Profile []string `json:"profile,omitempty"`

	// Usually this is a relative URL that locates the resource within the implementation guide
	Reference common.Reference `json:"reference"`

	// Appending 'rendering' + "/" + this should resolve to the resource page
	RelativePath *string `json:"relativePath,omitempty"`
}

// ImplementationGuideManifestPage represents information about a page within the IG
type ImplementationGuideManifestPage struct {
	common.BackboneElement

	// Appending 'rendering' + "/" + page.name + "#" + page.anchor should resolve to the anchor
	Anchor []string `json:"anchor,omitempty"`

	// Appending 'rendering' + "/" + this should resolve to the page
	Name string `json:"name"`

	// Label for the page intended for human display
	Title *string `json:"title,omitempty"`
}

// ImplementationGuideManifest represents information about an assembled implementation guide
type ImplementationGuideManifest struct {
	common.BackboneElement

	// Indicates a relative path to an image that exists within the IG
	Image []string `json:"image,omitempty"`

	// Indicates the relative path of an additional non-page, non-image file that is part of the IG
	Other []string `json:"other,omitempty"`

	// Information about a page within the IG
	Page []ImplementationGuideManifestPage `json:"page,omitempty"`

	// A pointer to official web page, PDF or other rendering of the implementation guide
	Rendering *string `json:"rendering,omitempty"`

	// A resource that is part of the implementation guide
	Resource []ImplementationGuideManifestResource `json:"resource"`
}

// ImplementationGuide represents a set of rules of how a particular interoperability or standards problem is solved
type ImplementationGuide struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "ImplementationGuide"

	// May be a web site, an email address, a telephone number, etc.
	Contact []ContactDetail `json:"contact,omitempty"`

	// Use and/or publishing restrictions
	Copyright *string `json:"copyright,omitempty"`

	// Copyright holder and year(s)
	CopyrightLabel *string `json:"copyrightLabel,omitempty"`

	// Date last changed
	Date *string `json:"date,omitempty"`

	// Principally, this consists of information about source resource and file locations
	Definition *ImplementationGuideDefinition `json:"definition,omitempty"`

	// Another implementation guide that this implementation depends on
	DependsOn []ImplementationGuideDependsOn `json:"dependsOn,omitempty"`

	// Natural language description of the implementation guide
	Description *string `json:"description,omitempty"`

	// For testing purposes, not real usage
	Experimental *bool `json:"experimental,omitempty"`

	// FHIR Version(s) this Implementation Guide targets
	FhirVersion []string `json:"fhirVersion"`

	// See Default Profiles for a discussion of which resources are 'covered' by an implementation guide
	Global []ImplementationGuideGlobal `json:"global,omitempty"`

	// Additional identifier for the implementation guide
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Intended jurisdiction for implementation guide (if applicable)
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`

	// The license that applies to this Implementation Guide, using an SPDX license code
	License *string `json:"license,omitempty"`

	// Information about an assembled implementation guide, created by the publication tooling
	Manifest *ImplementationGuideManifest `json:"manifest,omitempty"`

	// Name for this implementation guide (computer friendly)
	Name string `json:"name"`

	// Many (if not all) IG publishing tools will require that this element be present
	PackageId string `json:"packageId"`

	// Name of the publisher (organization or individual)
	Publisher *string `json:"publisher,omitempty"`

	// Why this implementation guide is defined
	Purpose *string `json:"purpose,omitempty"`

	// draft | active | retired | unknown
	Status ImplementationGuideStatus `json:"status"`

	// Name for this implementation guide (human friendly)
	Title *string `json:"title,omitempty"`

	// Canonical identifier for this implementation guide, represented as a URI
	Url string `json:"url"`

	// The context that the content is intended to support
	UseContext []UsageContext `json:"useContext,omitempty"`

	// Business version of the implementation guide
	Version *string `json:"version,omitempty"`

	// How to compare versions
	VersionAlgorithmString *string        `json:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *common.Coding `json:"versionAlgorithmCoding,omitempty"`
}

// IngredientStatus represents the status of an ingredient
type IngredientStatus string

const (
	IngredientStatusDraft   IngredientStatus = "draft"
	IngredientStatusActive  IngredientStatus = "active"
	IngredientStatusRetired IngredientStatus = "retired"
	IngredientStatusUnknown IngredientStatus = "unknown"
)

// IngredientManufacturer represents the organization(s) that manufacture this ingredient
type IngredientManufacturer struct {
	common.BackboneElement

	// The organization that manufactures this ingredient
	Manufacturer string `json:"manufacturer"`

	// The way in which this manufacturer is associated with the ingredient
	Role *common.CodeableConcept `json:"role,omitempty"`
}

// IngredientSubstanceStrengthReferenceStrength represents strength expressed in terms of a reference substance
type IngredientSubstanceStrengthReferenceStrength struct {
	common.BackboneElement

	// The country or countries for which the strength range applies
	Country []common.CodeableConcept `json:"country,omitempty"`

	// For when strength is measured at a particular point or distance
	MeasurementPoint *string `json:"measurementPoint,omitempty"`

	// Strength expressed in terms of a reference substance
	StrengthRatio    *Ratio           `json:"strengthRatio,omitempty"`
	StrengthQuantity *common.Quantity `json:"strengthQuantity,omitempty"`

	// Relevant reference substance
	Substance CodeableReference `json:"substance"`
}

// IngredientSubstanceStrength represents the quantity of substance in the unit of presentation
type IngredientSubstanceStrength struct {
	common.BackboneElement

	// A code that indicates if the strength is, for example, based on the ingredient substance as stated or on the substance base
	Basis *common.CodeableConcept `json:"basis,omitempty"`

	// The strength per unitary volume (or mass)
	ConcentrationRatio           *Ratio                  `json:"concentrationRatio,omitempty"`
	ConcentrationCodeableConcept *common.CodeableConcept `json:"concentrationCodeableConcept,omitempty"`
	ConcentrationQuantity        *common.Quantity        `json:"concentrationQuantity,omitempty"`

	// The country or countries for which the strength range applies
	Country []common.CodeableConcept `json:"country,omitempty"`

	// For when strength is measured at a particular point or distance
	MeasurementPoint *string `json:"measurementPoint,omitempty"`

	// The quantity of substance in the unit of presentation
	PresentationRatio           *Ratio                  `json:"presentationRatio,omitempty"`
	PresentationCodeableConcept *common.CodeableConcept `json:"presentationCodeableConcept,omitempty"`
	PresentationQuantity        *common.Quantity        `json:"presentationQuantity,omitempty"`

	// Strength expressed in terms of a reference substance
	ReferenceStrength []IngredientSubstanceStrengthReferenceStrength `json:"referenceStrength,omitempty"`

	// A textual represention of either the whole of the concentration strength or a part of it
	TextConcentration *string `json:"textConcentration,omitempty"`

	// A textual represention of either the whole of the presentation strength or a part of it
	TextPresentation *string `json:"textPresentation,omitempty"`
}

// IngredientSubstance represents the substance that comprises this ingredient
type IngredientSubstance struct {
	common.BackboneElement

	// A code or full resource that represents the ingredient's substance
	Code CodeableReference `json:"code"`

	// The quantity of substance in the unit of presentation
	Strength []IngredientSubstanceStrength `json:"strength,omitempty"`
}

// Ingredient represents an ingredient of a manufactured item or pharmaceutical product
type Ingredient struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Ingredient"

	// If the ingredient is a known or suspected allergen
	AllergenicIndicator *bool `json:"allergenicIndicator,omitempty"`

	// A place for providing any notes that are relevant to the component
	Comment *string `json:"comment,omitempty"`

	// The product which this ingredient is a constituent part of
	For []common.Reference `json:"for,omitempty"`

	// A classification of the ingredient identifying its precise purpose(s) in the drug product
	Function []common.CodeableConcept `json:"function,omitempty"`

	// A classification of the ingredient according to where in the physical item it tends to be used
	Group *common.CodeableConcept `json:"group,omitempty"`

	// The identifier(s) of this Ingredient that are assigned by business processes
	Identifier *common.Identifier `json:"identifier,omitempty"`

	// The organization(s) that manufacture this ingredient
	Manufacturer []IngredientManufacturer `json:"manufacturer,omitempty"`

	// A classification of the ingredient identifying its purpose within the product
	Role common.CodeableConcept `json:"role"`

	// Allows filtering of ingredient that are appropriate for use versus not
	Status IngredientStatus `json:"status"`

	// The substance that comprises this ingredient
	Substance IngredientSubstance `json:"substance"`
}

// InsurancePlanStatus represents the status of an insurance plan
type InsurancePlanStatus string

const (
	InsurancePlanStatusDraft   InsurancePlanStatus = "draft"
	InsurancePlanStatusActive  InsurancePlanStatus = "active"
	InsurancePlanStatusRetired InsurancePlanStatus = "retired"
	InsurancePlanStatusUnknown InsurancePlanStatus = "unknown"
)

// InsurancePlanCoverageBenefitLimit represents the specific limits on the benefit
type InsurancePlanCoverageBenefitLimit struct {
	common.BackboneElement

	// The specific limit on the benefit
	Code *common.CodeableConcept `json:"code,omitempty"`

	// May also be called "eligible expense," "payment allowance," or "negotiated rate"
	Value *common.Quantity `json:"value,omitempty"`
}

// InsurancePlanCoverageBenefit represents specific benefits under this type of coverage
type InsurancePlanCoverageBenefit struct {
	common.BackboneElement

	// The specific limits on the benefit
	Limit []InsurancePlanCoverageBenefitLimit `json:"limit,omitempty"`

	// The referral requirements to have access/coverage for this benefit
	Requirement *string `json:"requirement,omitempty"`

	// Type of benefit (primary care; speciality care; inpatient; outpatient)
	Type common.CodeableConcept `json:"type"`
}

// InsurancePlanCoverage represents details about the coverage offered by the insurance product
type InsurancePlanCoverage struct {
	common.BackboneElement

	// Specific benefits under this type of coverage
	Benefit []InsurancePlanCoverageBenefit `json:"benefit,omitempty"`

	// Networks and tiers
	Network []common.Reference `json:"network,omitempty"`

	// Type of coverage (Medical; Dental; Mental Health; Substance Abuse; Vision; Drug; Short Term; Long Term Care; Hospice; Home Health)
	Type common.CodeableConcept `json:"type"`
}

// InsurancePlanPlanGeneralCost represents general costs associated with the plan
type InsurancePlanPlanGeneralCost struct {
	common.BackboneElement

	// Additional information about the general costs
	Comment *string `json:"comment,omitempty"`

	// Value of the cost
	Cost *Money `json:"cost,omitempty"`

	// Number of enrollees
	GroupSize *int `json:"groupSize,omitempty"`

	// Type of cost (copay; individual-deductible; family-deductible; out-of-pocket-maximum)
	Type common.CodeableConcept `json:"type"`
}

// InsurancePlanPlanSpecificCostBenefitCost represents costs associated with the benefit
type InsurancePlanPlanSpecificCostBenefitCost struct {
	common.BackboneElement

	// Additional information about the cost
	Applicability *common.CodeableConcept `json:"applicability,omitempty"`

	// Qualifiers that apply to the cost
	Qualifiers []common.CodeableConcept `json:"qualifiers,omitempty"`

	// Type of cost (copay; individual-deductible; family-deductible; out-of-pocket-maximum)
	Type common.CodeableConcept `json:"type"`

	// The actual cost value
	Value *common.Quantity `json:"value,omitempty"`
}

// InsurancePlanPlanSpecificCostBenefit represents benefits and their costs for a specific coverage area
type InsurancePlanPlanSpecificCostBenefit struct {
	common.BackboneElement

	// List of the costs associated with this benefit
	Cost []InsurancePlanPlanSpecificCostBenefitCost `json:"cost,omitempty"`

	// Type of the benefit
	Type common.CodeableConcept `json:"type"`
}

// InsurancePlanPlanSpecificCost represents costs for a specific coverage area
type InsurancePlanPlanSpecificCost struct {
	common.BackboneElement

	// Benefits and their costs for a specific coverage area
	Benefit []InsurancePlanPlanSpecificCostBenefit `json:"benefit,omitempty"`

	// General category of benefit
	Category common.CodeableConcept `json:"category"`
}

// InsurancePlanPlan represents details about an insurance plan
type InsurancePlanPlan struct {
	common.BackboneElement

	// The geographic region in which a health insurance plan's benefits apply
	CoverageArea []common.Reference `json:"coverageArea,omitempty"`

	// Overall costs associated with the plan
	GeneralCost []InsurancePlanPlanGeneralCost `json:"generalCost,omitempty"`

	// Business identifier for the plan
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Networks and tiers
	Network []common.Reference `json:"network,omitempty"`

	// Costs for specific coverage areas
	SpecificCost []InsurancePlanPlanSpecificCost `json:"specificCost,omitempty"`

	// Type of plan (Bronze; Silver; Gold; Platinum)
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// InventoryItemStatus represents the status of an inventory item
type InventoryItemStatus string

const (
	InventoryItemStatusActive         InventoryItemStatus = "active"
	InventoryItemStatusInactive       InventoryItemStatus = "inactive"
	InventoryItemStatusEnteredInError InventoryItemStatus = "entered-in-error"
	InventoryItemStatusUnknown        InventoryItemStatus = "unknown"
)

// InventoryItemName represents the item name(s) - the brand name, or common name, functional name, generic name
type InventoryItemName struct {
	common.BackboneElement

	// The language that the item name is expressed in
	Language string `json:"language"`

	// The name or designation that the item is given
	Name string `json:"name"`

	// The type of name e.g. 'brand-name', 'functional-name', 'common-name'
	NameType common.Coding `json:"nameType"`
}

// InventoryItemResponsibleOrganization represents organization(s) responsible for the product
type InventoryItemResponsibleOrganization struct {
	common.BackboneElement

	// An organization that has an association with the item, e.g. manufacturer, distributor, responsible, etc.
	Organization common.Reference `json:"organization"`

	// The role of the organization e.g. manufacturer, distributor, etc.
	Role common.CodeableConcept `json:"role"`
}

// InventoryItemDescription represents the descriptive characteristics of the inventory item
type InventoryItemDescription struct {
	common.BackboneElement

	// Textual description of the item
	Description *string `json:"description,omitempty"`

	// The language for the item description
	Language *string `json:"language,omitempty"`
}

// InventoryItemAssociation represents association with other items or products
type InventoryItemAssociation struct {
	common.BackboneElement

	// This attribute defined the type of association when establishing associations or relations between items
	AssociationType common.CodeableConcept `json:"associationType"`

	// The quantity of the related product in this product
	Quantity Ratio `json:"quantity"`

	// The related item or product
	RelatedItem common.Reference `json:"relatedItem"`
}

// InventoryItemCharacteristic represents the descriptive or identifying characteristics of the item
type InventoryItemCharacteristic struct {
	common.BackboneElement

	// The type of characteristic that is being defined
	CharacteristicType common.CodeableConcept `json:"characteristicType"`

	// The string value is used for characteristics that are descriptive and not codeable information
	ValueString          *string                 `json:"valueString,omitempty"`
	ValueInteger         *int                    `json:"valueInteger,omitempty"`
	ValueDecimal         *float64                `json:"valueDecimal,omitempty"`
	ValueBoolean         *bool                   `json:"valueBoolean,omitempty"`
	ValueUrl             *string                 `json:"valueUrl,omitempty"`
	ValueDateTime        *string                 `json:"valueDateTime,omitempty"`
	ValueQuantity        *common.Quantity        `json:"valueQuantity,omitempty"`
	ValueRange           *Range                  `json:"valueRange,omitempty"`
	ValueRatio           *Ratio                  `json:"valueRatio,omitempty"`
	ValueAnnotation      *Annotation             `json:"valueAnnotation,omitempty"`
	ValueAddress         *Address                `json:"valueAddress,omitempty"`
	ValueDuration        *Duration               `json:"valueDuration,omitempty"`
	ValueCodeableConcept *common.CodeableConcept `json:"valueCodeableConcept,omitempty"`
}

// InventoryItemInstance represents instances or occurrences of the product
type InventoryItemInstance struct {
	common.BackboneElement

	// The expiry date or date and time for the product
	Expiry *string `json:"expiry,omitempty"`

	// The identifier for the physical instance, typically a serial number
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The location that the item is associated with
	Location *common.Reference `json:"location,omitempty"`

	// The lot or batch number of the item
	LotNumber *string `json:"lotNumber,omitempty"`

	// The subject that the item is associated with
	Subject *common.Reference `json:"subject,omitempty"`
}

// InventoryItem represents a functional description of an inventory item used in inventory and supply-related workflows
type InventoryItem struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "InventoryItem"

	// Association with other items or products
	Association []InventoryItemAssociation `json:"association,omitempty"`

	// The base unit of measure - the unit in which the product is used or counted
	BaseUnit *common.CodeableConcept `json:"baseUnit,omitempty"`

	// Category or class of the item
	Category []common.CodeableConcept `json:"category,omitempty"`

	// The descriptive or identifying characteristics of the item
	Characteristic []InventoryItemCharacteristic `json:"characteristic,omitempty"`

	// Code designating the specific type of item
	Code []common.CodeableConcept `json:"code,omitempty"`

	// The descriptive characteristics of the inventory item
	Description *InventoryItemDescription `json:"description,omitempty"`

	// Business identifier for the inventory item
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Instances or occurrences of the product
	Instance *InventoryItemInstance `json:"instance,omitempty"`

	// The usage status e.g. recalled, in use, discarded... This can be used to indicate that the items have been taken out of inventory
	InventoryStatus []common.CodeableConcept `json:"inventoryStatus,omitempty"`

	// The item name(s) - the brand name, or common name, functional name, generic name
	Name []InventoryItemName `json:"name,omitempty"`

	// Net content or amount present in the item
	NetContent *common.Quantity `json:"netContent,omitempty"`

	// Link to a product resource used in clinical workflows
	ProductReference *common.Reference `json:"productReference,omitempty"`

	// Organization(s) responsible for the product
	ResponsibleOrganization []InventoryItemResponsibleOrganization `json:"responsibleOrganization,omitempty"`

	// Status of the item entry
	Status InventoryItemStatus `json:"status"`
}

// InventoryReportCountType represents whether the report is about the current inventory count or a differential change
type InventoryReportCountType string

const (
	InventoryReportCountTypeSnapshot   InventoryReportCountType = "snapshot"
	InventoryReportCountTypeDifference InventoryReportCountType = "difference"
)

// InventoryReportStatus represents the status of the inventory check or notification
type InventoryReportStatus string

const (
	InventoryReportStatusDraft          InventoryReportStatus = "draft"
	InventoryReportStatusRequested      InventoryReportStatus = "requested"
	InventoryReportStatusActive         InventoryReportStatus = "active"
	InventoryReportStatusEnteredInError InventoryReportStatus = "entered-in-error"
)

// InventoryReportInventoryListingItem represents the item or items in this listing
type InventoryReportInventoryListingItem struct {
	common.BackboneElement

	// The inventory category or classification of the items being reported
	Category *common.CodeableConcept `json:"category,omitempty"`

	// The code or reference to the item type
	Item CodeableReference `json:"item"`

	// The quantity of the item or items being reported
	Quantity common.Quantity `json:"quantity"`
}

// InventoryReportInventoryListing represents an inventory listing section (grouped by any of the attributes)
type InventoryReportInventoryListing struct {
	common.BackboneElement

	// The date and time when the items were counted
	CountingDateTime *string `json:"countingDateTime,omitempty"`

	// The item or items in this listing
	Item []InventoryReportInventoryListingItem `json:"item,omitempty"`

	// The status of the items
	ItemStatus *common.CodeableConcept `json:"itemStatus,omitempty"`

	// Location of the inventory items
	Location *common.Reference `json:"location,omitempty"`
}

// InventoryReport represents a report of inventory or stock items
type InventoryReport struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "InventoryReport"

	// Whether the report is about the current inventory count (snapshot) or a differential change in inventory (change)
	CountType InventoryReportCountType `json:"countType"`

	// Business identifier for the InvoiceReport
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// An inventory listing section (grouped by any of the attributes)
	InventoryListing []InventoryReportInventoryListing `json:"inventoryListing,omitempty"`

	// A note associated with the InvoiceReport
	Note []Annotation `json:"note,omitempty"`

	// What type of operation is being performed - addition or subtraction
	OperationType *common.CodeableConcept `json:"operationType,omitempty"`

	// The reason for this count - regular count, ad-hoc count, new arrivals, etc.
	OperationTypeReason *common.CodeableConcept `json:"operationTypeReason,omitempty"`

	// When the report has been submitted
	ReportedDateTime string `json:"reportedDateTime"`

	// Who submits the report
	Reporter *common.Reference `json:"reporter,omitempty"`

	// The period the report refers to
	ReportingPeriod *common.Period `json:"reportingPeriod,omitempty"`

	// The status of the inventory check or notification
	Status InventoryReportStatus `json:"status"`
}

// InvoiceStatus represents the current state of the Invoice
type InvoiceStatus string

const (
	InvoiceStatusDraft          InvoiceStatus = "draft"
	InvoiceStatusIssued         InvoiceStatus = "issued"
	InvoiceStatusBalanced       InvoiceStatus = "balanced"
	InvoiceStatusCancelled      InvoiceStatus = "cancelled"
	InvoiceStatusEnteredInError InvoiceStatus = "entered-in-error"
)

// InvoiceParticipant represents participants who performed or participated in the charged service
type InvoiceParticipant struct {
	common.BackboneElement

	// The device, practitioner, etc. who performed or participated in the service
	Actor common.Reference `json:"actor"`

	// Describes the type of involvement (e.g. transcriptionist, creator etc.)
	Role *common.CodeableConcept `json:"role,omitempty"`
}

// InvoiceLineItem represents each line item representing one charge for goods and services rendered
type InvoiceLineItem struct {
	common.BackboneElement

	// The ChargeItem contains information such as the billing code, date, amount etc.
	ChargeItemReference       *common.Reference       `json:"chargeItemReference,omitempty"`
	ChargeItemCodeableConcept *common.CodeableConcept `json:"chargeItemCodeableConcept,omitempty"`

	// The price for a ChargeItem may be calculated as a base price with surcharges/deductions
	PriceComponent []MonetaryComponent `json:"priceComponent,omitempty"`

	// Sequence in which the items appear on the invoice
	Sequence *int `json:"sequence,omitempty"`

	// Date/time(s) range when this service was delivered or completed
	ServicedDate   *string        `json:"servicedDate,omitempty"`
	ServicedPeriod *common.Period `json:"servicedPeriod,omitempty"`
}

// Invoice represents invoice containing collected ChargeItems from an Account with calculated individual and total price
type Invoice struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Invoice"

	// Account that is supposed to be balanced with this Invoice
	Account *common.Reference `json:"account,omitempty"`

	// Reason for cancellation of this Invoice
	CancelledReason *string `json:"cancelledReason,omitempty"`

	// When this Invoice was posted
	Creation *string `json:"creation,omitempty"`

	// Deprecated by the element below
	Date *string `json:"date,omitempty"`

	// Identifier of this Invoice
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The issuer of the Invoice
	Issuer *common.Reference `json:"issuer,omitempty"`

	// Each line item represents one charge for goods and services rendered
	LineItem []InvoiceLineItem `json:"lineItem,omitempty"`

	// Comments made about the invoice by the issuer, subject, or other participants
	Note []Annotation `json:"note,omitempty"`

	// Indicates who or what performed or participated in the charged service
	Participant []InvoiceParticipant `json:"participant,omitempty"`

	// Payment details
	PaymentTerms *string `json:"paymentTerms,omitempty"`

	// Date/time(s) range of services included in this invoice
	PeriodDate   *string        `json:"periodDate,omitempty"`
	PeriodPeriod *common.Period `json:"periodPeriod,omitempty"`

	// The individual or Organization responsible for balancing of this invoice
	Recipient *common.Reference `json:"recipient,omitempty"`

	// The current state of the Invoice
	Status InvoiceStatus `json:"status"`

	// The individual or set of individuals receiving the goods and services billed in this invoice
	Subject *common.Reference `json:"subject,omitempty"`

	// Invoice Totals
	TotalGross          *Money              `json:"totalGross,omitempty"`
	TotalNet            *Money              `json:"totalNet,omitempty"`
	TotalPriceComponent []MonetaryComponent `json:"totalPriceComponent,omitempty"`

	// Type of Invoice depending on domain, realm an usage
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// LibraryStatus represents the status of a library
type LibraryStatus string

const (
	LibraryStatusDraft   LibraryStatus = "draft"
	LibraryStatusActive  LibraryStatus = "active"
	LibraryStatusRetired LibraryStatus = "retired"
	LibraryStatusUnknown LibraryStatus = "unknown"
)

// Library represents a general-purpose container for knowledge asset definitions
type Library struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Library"

	// The 'date' element may be more recent than the approval date because of minor changes or editorial corrections
	ApprovalDate *string `json:"approvalDate,omitempty"`

	// An individual or organization primarily involved in the creation and maintenance of the content
	Author []ContactDetail `json:"author,omitempty"`

	// May be a web site, an email address, a telephone number, etc.
	Contact []ContactDetail `json:"contact,omitempty"`

	// The content of the library as an Attachment
	Content []Attachment `json:"content,omitempty"`

	// Use and/or publishing restrictions
	Copyright *string `json:"copyright,omitempty"`

	// Copyright holder and year(s)
	CopyrightLabel *string `json:"copyrightLabel,omitempty"`

	// Describes a set of data that must be provided in order to be able to successfully perform the computations
	DataRequirement []interface{} `json:"dataRequirement,omitempty"`

	// Date last changed
	Date *string `json:"date,omitempty"`

	// Natural language description of the library
	Description *string `json:"description,omitempty"`

	// An individual or organization primarily responsible for internal coherence of the content
	Editor []ContactDetail `json:"editor,omitempty"`

	// When the library is expected to be used
	EffectivePeriod *common.Period `json:"effectivePeriod,omitempty"`

	// Who endorsed the content
	Endorser []ContactDetail `json:"endorser,omitempty"`

	// For testing purposes, not real usage
	Experimental *bool `json:"experimental,omitempty"`

	// Additional identifier for the library
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Intended jurisdiction for library (if applicable)
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`

	// When the library was last reviewed
	LastReviewDate *string `json:"lastReviewDate,omitempty"`

	// Name for this library (computer friendly)
	Name *string `json:"name,omitempty"`

	// The parameter element defines parameters used by the library
	Parameter []interface{} `json:"parameter,omitempty"`

	// Name of the publisher (organization or individual)
	Publisher *string `json:"publisher,omitempty"`

	// Why this library is defined
	Purpose *string `json:"purpose,omitempty"`

	// Additional documentation, citations, etc.
	RelatedArtifact []interface{} `json:"relatedArtifact,omitempty"`

	// Who reviewed the content
	Reviewer []ContactDetail `json:"reviewer,omitempty"`

	// draft | active | retired | unknown
	Status LibraryStatus `json:"status"`

	// Type of individual the library content is focused on
	SubjectCodeableConcept *common.CodeableConcept `json:"subjectCodeableConcept,omitempty"`
	SubjectReference       *common.Reference       `json:"subjectReference,omitempty"`

	// Subordinate title of the library
	Subtitle *string `json:"subtitle,omitempty"`

	// Name for this library (human friendly)
	Title *string `json:"title,omitempty"`

	// The category of the library, such as education, treatment, assessment, etc.
	Topic []common.CodeableConcept `json:"topic,omitempty"`

	// Identifies the type of library such as a Logic Library, Model Definition, Asset Collection, or Module Definition
	Type common.CodeableConcept `json:"type"`

	// Canonical identifier for this library, represented as a URI
	Url *string `json:"url,omitempty"`

	// Describes the clinical usage of the library
	Usage *string `json:"usage,omitempty"`

	// The context that the content is intended to support
	UseContext []interface{} `json:"useContext,omitempty"`

	// Business version of the library
	Version *string `json:"version,omitempty"`

	// How to compare versions
	VersionAlgorithmString *string        `json:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *common.Coding `json:"versionAlgorithmCoding,omitempty"`
}

// LinkageItemType represents the type of linkage item
type LinkageItemType string

const (
	LinkageItemTypeSource     LinkageItemType = "source"
	LinkageItemTypeAlternate  LinkageItemType = "alternate"
	LinkageItemTypeHistorical LinkageItemType = "historical"
)

// LinkageItem represents an item in a linkage
type LinkageItem struct {
	common.BackboneElement

	// The resource instance being linked as part of the group
	Resource common.Reference `json:"resource"`

	// Distinguishes which item is "source of truth" (if any) and which items are no longer considered to be current representations
	Type LinkageItemType `json:"type"`
}

// Linkage represents a linkage between resources
type Linkage struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Linkage"

	// If false, any asserted linkages should not be considered current/relevant/applicable
	Active *bool `json:"active,omitempty"`

	// Identifies the user or organization responsible for asserting the linkages
	Author *common.Reference `json:"author,omitempty"`

	// Identifies which record considered as the reference to the same real-world occurrence
	Item []LinkageItem `json:"item"`
}

// ListMode represents the mode of a list
type ListMode string

const (
	ListModeWorking  ListMode = "working"
	ListModeSnapshot ListMode = "snapshot"
	ListModeChanges  ListMode = "changes"
)

// ListStatus represents the status of a list
type ListStatus string

const (
	ListStatusCurrent        ListStatus = "current"
	ListStatusRetired        ListStatus = "retired"
	ListStatusEnteredInError ListStatus = "entered-in-error"
)

// ListEntry represents an entry in a list
type ListEntry struct {
	common.BackboneElement

	// When this item was added to the list
	Date *string `json:"date,omitempty"`

	// If the flag means that the entry has actually been deleted from the list, the deleted element SHALL be true
	Deleted *bool `json:"deleted,omitempty"`

	// The flag can only be understood in the context of the List.code
	Flag *common.CodeableConcept `json:"flag,omitempty"`

	// A reference to the actual resource from which data was derived
	Item common.Reference `json:"item"`
}

// List represents a curated collection of resources
type List struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "List"

	// If there is no code, the purpose of the list is implied where it is used
	Code *common.CodeableConcept `json:"code,omitempty"`

	// The actual important date is the date of currency of the resources that were summarized
	Date *string `json:"date,omitempty"`

	// The various reasons for an empty list make a significant interpretation to its interpretation
	EmptyReason *common.CodeableConcept `json:"emptyReason,omitempty"`

	// The encounter that is the context in which this list was created
	Encounter *common.Reference `json:"encounter,omitempty"`

	// If there are no entries in the list, an emptyReason SHOULD be provided
	Entry []ListEntry `json:"entry,omitempty"`

	// Identifier for the List assigned for business purposes outside the context of FHIR
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// This element is labeled as a modifier because a change list must not be misunderstood as a complete list
	Mode ListMode `json:"mode"`

	// Comments that apply to the overall list
	Note []Annotation `json:"note,omitempty"`

	// Applications SHOULD render ordered lists in the order provided
	OrderedBy *common.CodeableConcept `json:"orderedBy,omitempty"`

	// The primary source is the entity that made the decisions what items are in the list
	Source *common.Reference `json:"source,omitempty"`

	// This element is labeled as a modifier because the status contains codes that mark the resource as not currently valid
	Status ListStatus `json:"status"`

	// Some purely arbitrary lists do not have a common subject, so this is optional
	Subject []common.Reference `json:"subject,omitempty"`

	// A label for the list assigned by the author
	Title *string `json:"title,omitempty"`
}

// ManufacturedItemDefinitionStatus represents the status of a manufactured item definition
type ManufacturedItemDefinitionStatus string

const (
	ManufacturedItemDefinitionStatusDraft   ManufacturedItemDefinitionStatus = "draft"
	ManufacturedItemDefinitionStatusActive  ManufacturedItemDefinitionStatus = "active"
	ManufacturedItemDefinitionStatusRetired ManufacturedItemDefinitionStatus = "retired"
	ManufacturedItemDefinitionStatusUnknown ManufacturedItemDefinitionStatus = "unknown"
)

// ManufacturedItemDefinitionProperty represents general characteristics of an item
type ManufacturedItemDefinitionProperty struct {
	common.BackboneElement

	// A code expressing the type of characteristic
	Type common.CodeableConcept `json:"type"`

	// A value for the characteristic
	ValueCodeableConcept *common.CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueQuantity        *common.Quantity        `json:"valueQuantity,omitempty"`
	ValueDate            *string                 `json:"valueDate,omitempty"`
	ValueBoolean         *bool                   `json:"valueBoolean,omitempty"`
	ValueMarkdown        *string                 `json:"valueMarkdown,omitempty"`
	ValueAttachment      *Attachment             `json:"valueAttachment,omitempty"`
	ValueReference       *common.Reference       `json:"valueReference,omitempty"`
}

// ManufacturedItemDefinitionComponentConstituent represents a constituent of a component
type ManufacturedItemDefinitionComponentConstituent struct {
	common.BackboneElement

	// The measurable amount of the constituent within the component
	Amount []common.Quantity `json:"amount,omitempty"`

	// A reference to a constituent of the manufactured item
	Constituent CodeableReference `json:"constituent"`

	// A reference to a component of a constituent within this component
	Component []ManufacturedItemDefinitionComponent `json:"component,omitempty"`

	// The function of this constituent within the component
	Function []common.CodeableConcept `json:"function,omitempty"`

	// The physical location of the constituent within the component
	Location []common.CodeableConcept `json:"location,omitempty"`
}

// ManufacturedItemDefinitionComponent represents physical parts of the manufactured item
type ManufacturedItemDefinitionComponent struct {
	common.BackboneElement

	// The measurable amount of total quantity of all substances in the component
	Amount []common.Quantity `json:"amount,omitempty"`

	// A component that this component contains or is made from
	Component []ManufacturedItemDefinitionComponent `json:"component,omitempty"`

	// A reference to a constituent of the manufactured item
	Constituent []ManufacturedItemDefinitionComponentConstituent `json:"constituent,omitempty"`

	// The function of this component within the item
	Function []common.CodeableConcept `json:"function,omitempty"`

	// General characteristics of this component
	Property []ManufacturedItemDefinitionProperty `json:"property,omitempty"`

	// Defining type of the component e.g. shell, layer, ink
	Type common.CodeableConcept `json:"type"`
}

// ManufacturedItemDefinition represents the definition and characteristics of a medicinal manufactured item
type ManufacturedItemDefinition struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "ManufacturedItemDefinition"

	// Physical parts of the manufactured item
	Component []ManufacturedItemDefinitionComponent `json:"component,omitempty"`

	// Unique identifier
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The ingredients of this manufactured item
	Ingredient []common.CodeableConcept `json:"ingredient,omitempty"`

	// Dose form as manufactured and before any transformation into the pharmaceutical product
	ManufacturedDoseForm common.CodeableConcept `json:"manufacturedDoseForm"`

	// Manufacturer of the item, one of several possible
	Manufacturer []common.Reference `json:"manufacturer,omitempty"`

	// Allows specifying that an item is on the market for sale
	MarketingStatus []interface{} `json:"marketingStatus,omitempty"`

	// A descriptive name applied to this item
	Name *string `json:"name,omitempty"`

	// General characteristics of this item
	Property []ManufacturedItemDefinitionProperty `json:"property,omitempty"`

	// Allows filtering of manufactured items that are appropriate for use versus not
	Status ManufacturedItemDefinitionStatus `json:"status"`

	// The "real-world" units in which the quantity of the manufactured item is described
	UnitOfPresentation *common.CodeableConcept `json:"unitOfPresentation,omitempty"`
}

// MeasureStatus represents the status of a measure
type MeasureStatus string

const (
	MeasureStatusDraft   MeasureStatus = "draft"
	MeasureStatusActive  MeasureStatus = "active"
	MeasureStatusRetired MeasureStatus = "retired"
	MeasureStatusUnknown MeasureStatus = "unknown"
)

// MeasureTerm represents a term used within the measure
type MeasureTerm struct {
	common.BackboneElement

	// A codeable representation of the defined term
	Code *common.CodeableConcept `json:"code,omitempty"`

	// Provides a definition for the term as used within the measure
	Definition *string `json:"definition,omitempty"`
}

// MeasureGroupPopulation represents a population in a measure group
type MeasureGroupPopulation struct {
	common.BackboneElement

	// The code for the population criteria
	Code *common.CodeableConcept `json:"code,omitempty"`

	// The human readable description of this population criteria
	Description *string `json:"description,omitempty"`

	// An expression that specifies the criteria for this population
	Criteria interface{} `json:"criteria"`

	// A unique identifier for the population criteria
	LinkId *string `json:"linkId,omitempty"`
}

// MeasureGroupStratifier represents a stratifier for a measure group
type MeasureGroupStratifier struct {
	common.BackboneElement

	// The code for this stratifier
	Code *common.CodeableConcept `json:"code,omitempty"`

	// Indicates a meaning for the stratifier component
	Component []MeasureGroupStratifierComponent `json:"component,omitempty"`

	// The expression that is used to define the stratifier criteria
	Criteria *interface{} `json:"criteria,omitempty"`

	// The human readable description of this stratifier
	Description *string `json:"description,omitempty"`

	// A unique identifier for the stratifier
	LinkId *string `json:"linkId,omitempty"`
}

// MeasureGroupStratifierComponent represents a component of a stratifier
type MeasureGroupStratifierComponent struct {
	common.BackboneElement

	// The code for the stratifier component
	Code *common.CodeableConcept `json:"code,omitempty"`

	// The expression that is used to define the stratifier criteria
	Criteria interface{} `json:"criteria"`

	// The human readable description of this stratifier component
	Description *string `json:"description,omitempty"`

	// A unique identifier for the stratifier component
	LinkId *string `json:"linkId,omitempty"`
}

// MeasureGroup represents a group of populations for a measure
type MeasureGroup struct {
	common.BackboneElement

	// Meaning of the group
	Code *common.CodeableConcept `json:"code,omitempty"`

	// The human readable description of this population group
	Description *string `json:"description,omitempty"`

	// A unique identifier for the group
	LinkId *string `json:"linkId,omitempty"`

	// A population criteria for the measure
	Population []MeasureGroupPopulation `json:"population,omitempty"`

	// The stratifier criteria for the measure report
	Stratifier []MeasureGroupStratifier `json:"stratifier,omitempty"`
}

// MeasureSupplementalData represents supplemental data for a measure
type MeasureSupplementalData struct {
	common.BackboneElement

	// The code for the supplemental data element
	Code *common.CodeableConcept `json:"code,omitempty"`

	// The expression that is used to define the supplemental data
	Criteria interface{} `json:"criteria"`

	// The human readable description of this supplemental data
	Description *string `json:"description,omitempty"`

	// A unique identifier for the supplemental data element
	LinkId *string `json:"linkId,omitempty"`

	// Indicates how the calculation is performed for the measure
	Usage []common.CodeableConcept `json:"usage,omitempty"`
}

// Measure represents the definition of a quality measure
type Measure struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Measure"

	// The 'date' element may be more recent than the approval date
	ApprovalDate *string `json:"approvalDate,omitempty"`

	// An individual or organization primarily involved in the creation and maintenance of the content
	Author []ContactDetail `json:"author,omitempty"`

	// For a subject-based measure, the population basis is simply boolean
	Basis *string `json:"basis,omitempty"`

	// Provides a summary of relevant clinical guidelines or other clinical recommendations
	ClinicalRecommendationStatement *string `json:"clinicalRecommendationStatement,omitempty"`

	// If this is a composite measure, the scoring method used to combine the component measures
	CompositeScoring *common.CodeableConcept `json:"compositeScoring,omitempty"`

	// May be a web site, an email address, a telephone number, etc.
	Contact []ContactDetail `json:"contact,omitempty"`

	// A copyright statement relating to the measure and/or its contents
	Copyright *string `json:"copyright,omitempty"`

	// Copyright holder and year(s)
	CopyrightLabel *string `json:"copyrightLabel,omitempty"`

	// Note that this is not the same as the resource last-modified-date
	Date *string `json:"date,omitempty"`

	// Natural language description of the measure
	Description *string `json:"description,omitempty"`

	// Provides a summary of key clinical recommendation
	Disclaimer *string `json:"disclaimer,omitempty"`

	// An individual or organization primarily responsible for internal coherence of the content
	Editor []ContactDetail `json:"editor,omitempty"`

	// When the measure is expected to be used
	EffectivePeriod *common.Period `json:"effectivePeriod,omitempty"`

	// Who endorsed the content
	Endorser []ContactDetail `json:"endorser,omitempty"`

	// For testing purposes, not real usage
	Experimental *bool `json:"experimental,omitempty"`

	// Additional guidance for the measure including how it can be used in a clinical context
	Guidance *string `json:"guidance,omitempty"`

	// Population criteria group
	Group []MeasureGroup `json:"group,omitempty"`

	// Additional identifier for the measure
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Information about whether an improvement in the measure is noted by an increase or decrease
	ImprovementNotation *common.CodeableConcept `json:"improvementNotation,omitempty"`

	// Intended jurisdiction for measure (if applicable)
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`

	// When the measure was last reviewed
	LastReviewDate *string `json:"lastReviewDate,omitempty"`

	// Logic used by the measure
	Library []string `json:"library,omitempty"`

	// Name for this measure (computer friendly)
	Name *string `json:"name,omitempty"`

	// Name of the publisher (organization or individual)
	Publisher *string `json:"publisher,omitempty"`

	// Why this measure is defined
	Purpose *string `json:"purpose,omitempty"`

	// Summary of how the measure addresses the clinical quality construct
	RateAggregation *string `json:"rateAggregation,omitempty"`

	// Detailed description of why the measure exists
	Rationale *string `json:"rationale,omitempty"`

	// Additional documentation, citations, etc.
	RelatedArtifact []interface{} `json:"relatedArtifact,omitempty"`

	// Who reviewed the content
	Reviewer []ContactDetail `json:"reviewer,omitempty"`

	// Risk adjustment formula used to calculate the measure
	RiskAdjustment *string `json:"riskAdjustment,omitempty"`

	// Indicates how the calculation is performed for the measure
	Scoring *common.CodeableConcept `json:"scoring,omitempty"`

	// draft | active | retired | unknown
	Status MeasureStatus `json:"status"`

	// Subordinate title of the measure
	Subtitle *string `json:"subtitle,omitempty"`

	// Supplemental data for the measure
	SupplementalData []MeasureSupplementalData `json:"supplementalData,omitempty"`

	// Provides a description of an individual term used within the measure
	Term []MeasureTerm `json:"term,omitempty"`

	// Name for this measure (human friendly)
	Title *string `json:"title,omitempty"`

	// The category of the measure, such as education, treatment, assessment, etc.
	Topic []common.CodeableConcept `json:"topic,omitempty"`

	// Indicates whether the measure is used to examine a process, an outcome over time, a patient-reported outcome, or a structure measure
	Type []common.CodeableConcept `json:"type,omitempty"`

	// Canonical identifier for this measure, represented as a URI
	Url *string `json:"url,omitempty"`

	// Describes the clinical usage of the measure
	Usage *string `json:"usage,omitempty"`

	// The context that the content is intended to support
	UseContext []interface{} `json:"useContext,omitempty"`

	// Business version of the measure
	Version *string `json:"version,omitempty"`

	// How to compare versions
	VersionAlgorithmString *string        `json:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *common.Coding `json:"versionAlgorithmCoding,omitempty"`
}

// MeasureReportStatus represents the status of a measure report
type MeasureReportStatus string

const (
	MeasureReportStatusComplete MeasureReportStatus = "complete"
	MeasureReportStatusPending  MeasureReportStatus = "pending"
	MeasureReportStatusError    MeasureReportStatus = "error"
)

// MeasureReportType represents the type of measure report
type MeasureReportType string

const (
	MeasureReportTypeIndividual   MeasureReportType = "individual"
	MeasureReportTypeSubjectList  MeasureReportType = "subject-list"
	MeasureReportTypeSummary      MeasureReportType = "summary"
	MeasureReportTypeDataExchange MeasureReportType = "data-exchange"
)

// MeasureReportDataUpdateType represents the data update type
type MeasureReportDataUpdateType string

const (
	MeasureReportDataUpdateTypeIncremental MeasureReportDataUpdateType = "incremental"
	MeasureReportDataUpdateTypeSnapshot    MeasureReportDataUpdateType = "snapshot"
)

// MeasureReportGroupPopulation represents a population in a measure report group
type MeasureReportGroupPopulation struct {
	common.BackboneElement

	// The code for the population criteria
	Code *common.CodeableConcept `json:"code,omitempty"`

	// The number of members of the population
	Count *int `json:"count,omitempty"`

	// A unique identifier for the population criteria
	LinkId *string `json:"linkId,omitempty"`

	// Optional subject identifying the individual or individuals the report is for
	SubjectReference []common.Reference `json:"subjectReference,omitempty"`

	// Optional result references if the measure result is computed for the individual
	SubjectResults []common.Reference `json:"subjectResults,omitempty"`

	// Optional subject identifying the list of subjects the report is for
	Subjects *common.Reference `json:"subjects,omitempty"`
}

// MeasureReportGroupStratifierStratum represents a stratum in a stratifier
type MeasureReportGroupStratifierStratum struct {
	common.BackboneElement

	// A component of this stratifier stratum
	Component []MeasureReportGroupStratifierStratumComponent `json:"component,omitempty"`

	// The measure score for this stratum
	MeasureScoreQuantity        *common.Quantity        `json:"measureScoreQuantity,omitempty"`
	MeasureScoreDateTime        *string                 `json:"measureScoreDateTime,omitempty"`
	MeasureScoreCodeableConcept *common.CodeableConcept `json:"measureScoreCodeableConcept,omitempty"`
	MeasureScorePeriod          *common.Period          `json:"measureScorePeriod,omitempty"`
	MeasureScoreRange           *Range                  `json:"measureScoreRange,omitempty"`
	MeasureScoreDuration        *Duration               `json:"measureScoreDuration,omitempty"`

	// A population criteria for the measure report
	Population []MeasureReportGroupPopulation `json:"population,omitempty"`

	// The value for this stratum, expressed as a CodeableConcept
	ValueCodeableConcept *common.CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueBoolean         *bool                   `json:"valueBoolean,omitempty"`
	ValueQuantity        *common.Quantity        `json:"valueQuantity,omitempty"`
	ValueRange           *Range                  `json:"valueRange,omitempty"`
	ValueReference       *common.Reference       `json:"valueReference,omitempty"`
}

// MeasureReportGroupStratifierStratumComponent represents a component of a stratifier stratum
type MeasureReportGroupStratifierStratumComponent struct {
	common.BackboneElement

	// The code for the stratifier component
	Code *common.CodeableConcept `json:"code,omitempty"`

	// A unique identifier for the stratifier component
	LinkId *string `json:"linkId,omitempty"`

	// The value for this component, expressed as a CodeableConcept
	ValueCodeableConcept *common.CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueBoolean         *bool                   `json:"valueBoolean,omitempty"`
	ValueQuantity        *common.Quantity        `json:"valueQuantity,omitempty"`
	ValueRange           *Range                  `json:"valueRange,omitempty"`
	ValueReference       *common.Reference       `json:"valueReference,omitempty"`
}

// MeasureReportGroupStratifier represents a stratifier for a measure report group
type MeasureReportGroupStratifier struct {
	common.BackboneElement

	// The code for this stratifier
	Code *common.CodeableConcept `json:"code,omitempty"`

	// A unique identifier for the stratifier
	LinkId *string `json:"linkId,omitempty"`

	// Stratifier Results
	Stratum []MeasureReportGroupStratifierStratum `json:"stratum,omitempty"`
}

// MeasureReportGroup represents a group of populations for a measure report
type MeasureReportGroup struct {
	common.BackboneElement

	// Meaning of the group
	Code *common.CodeableConcept `json:"code,omitempty"`

	// A unique identifier for the group
	LinkId *string `json:"linkId,omitempty"`

	// The measure score for this population group
	MeasureScoreQuantity        *common.Quantity        `json:"measureScoreQuantity,omitempty"`
	MeasureScoreDateTime        *string                 `json:"measureScoreDateTime,omitempty"`
	MeasureScoreCodeableConcept *common.CodeableConcept `json:"measureScoreCodeableConcept,omitempty"`
	MeasureScorePeriod          *common.Period          `json:"measureScorePeriod,omitempty"`
	MeasureScoreRange           *Range                  `json:"measureScoreRange,omitempty"`
	MeasureScoreDuration        *Duration               `json:"measureScoreDuration,omitempty"`

	// A population criteria for the measure report
	Population []MeasureReportGroupPopulation `json:"population,omitempty"`

	// The stratifier criteria for the measure report
	Stratifier []MeasureReportGroupStratifier `json:"stratifier,omitempty"`
}

// MeasureReport represents the results of the calculation of a measure
type MeasureReport struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "MeasureReport"

	// This element only applies to Data-collection reports
	DataUpdateType *MeasureReportDataUpdateType `json:"dataUpdateType,omitempty"`

	// The date this measure was calculated
	Date *string `json:"date,omitempty"`

	// Evaluated resources are only reported for individual reports
	EvaluatedResource []common.Reference `json:"evaluatedResource,omitempty"`

	// The results of the calculation, one for each population group in the measure
	Group []MeasureReportGroup `json:"group,omitempty"`

	// Typically, this is used for identifiers that can go in an HL7 V3 II data type
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Information about whether an improvement in the measure is noted by an increase or decrease
	ImprovementNotation *common.CodeableConcept `json:"improvementNotation,omitempty"`

	// Input parameters from operation
	InputParameters *common.Reference `json:"inputParameters,omitempty"`

	// A reference to the location for which the data is being reported
	Location *common.Reference `json:"location,omitempty"`

	// A reference to the Measure that was calculated to produce this report
	Measure *string `json:"measure,omitempty"`

	// The reporting period for which the report was calculated
	Period common.Period `json:"period"`

	// The individual or organization that is reporting the data
	Reporter *common.Reference `json:"reporter,omitempty"`

	// A reference to the vendor who queried the data, calculated results and/or generated the report
	ReportingVendor *common.Reference `json:"reportingVendor,omitempty"`

	// Indicates how the calculation is performed for the measure
	Scoring *common.CodeableConcept `json:"scoring,omitempty"`

	// This element is labeled as a modifier because the status contains codes that mark the resource as not currently valid
	Status MeasureReportStatus `json:"status"`

	// Optional subject identifying the individual or individuals the report is for
	Subject *common.Reference `json:"subject,omitempty"`

	// For individual measure reports, the supplementalData elements represent the direct result of evaluating the supplementalData expression
	SupplementalData []common.Reference `json:"supplementalData,omitempty"`

	// Data-exchange reports are used only to communicate data-of-interest for a measure
	Type MeasureReportType `json:"type"`
}

// MedicationAdministrationStatus represents the status of a medication administration
type MedicationAdministrationStatus string

const (
	MedicationAdministrationStatusInProgress     MedicationAdministrationStatus = "in-progress"
	MedicationAdministrationStatusNotDone        MedicationAdministrationStatus = "not-done"
	MedicationAdministrationStatusOnHold         MedicationAdministrationStatus = "on-hold"
	MedicationAdministrationStatusCompleted      MedicationAdministrationStatus = "completed"
	MedicationAdministrationStatusEnteredInError MedicationAdministrationStatus = "entered-in-error"
	MedicationAdministrationStatusStopped        MedicationAdministrationStatus = "stopped"
	MedicationAdministrationStatusUnknown        MedicationAdministrationStatus = "unknown"
)

// MedicationAdministrationPerformer represents the performer of the medication administration
type MedicationAdministrationPerformer struct {
	common.BackboneElement

	// Indicates who or what performed the medication administration
	Actor common.Reference `json:"actor"`

	// Distinguishes the type of involvement of the actor in the medication administration
	Function *common.CodeableConcept `json:"function,omitempty"`
}

// MedicationAdministrationDosage represents dosage information for medication administration
type MedicationAdministrationDosage struct {
	common.BackboneElement

	// Amount of medication per dose
	Dose *common.Quantity `json:"dose,omitempty"`

	// How drug should enter body
	Method *common.CodeableConcept `json:"method,omitempty"`

	// Amount of medication per unit of time
	RateRatio    *Ratio           `json:"rateRatio,omitempty"`
	RateQuantity *common.Quantity `json:"rateQuantity,omitempty"`

	// Body site to administer to
	Route *common.CodeableConcept `json:"route,omitempty"`

	// Body site to administer to
	Site *common.CodeableConcept `json:"site,omitempty"`

	// Free text dosage instructions
	Text *string `json:"text,omitempty"`
}

// MedicationAdministration represents the event of a patient consuming or otherwise being administered a medication
type MedicationAdministration struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "MedicationAdministration"

	// A plan that is fulfilled in whole or in part by this MedicationAdministration
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// The type of medication administration
	Category []common.CodeableConcept `json:"category,omitempty"`

	// The device that is to be used for the administration of the medication
	Device []CodeableReference `json:"device,omitempty"`

	// Describes the medication dosage information details
	Dosage *MedicationAdministrationDosage `json:"dosage,omitempty"`

	// The visit, admission, or other contact between patient and health care provider
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Event history references
	EventHistory []common.Reference `json:"eventHistory,omitempty"`

	// Business identifier for this medication administration
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// An indication that the full dose was not administered
	IsSubPotent *bool `json:"isSubPotent,omitempty"`

	// If only a code is specified, then it needs to be a code for a specific product
	Medication CodeableReference `json:"medication"`

	// Extra information about the medication administration
	Note []Annotation `json:"note,omitempty"`

	// A specific date/time or interval of time during which the administration took place
	OccurrenceDateTime *string        `json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod   *common.Period `json:"occurrencePeriod,omitempty"`
	OccurrenceTiming   *Timing        `json:"occurrenceTiming,omitempty"`

	// A larger event of which this particular event is a component or step
	PartOf []common.Reference `json:"partOf,omitempty"`

	// The performer of the medication treatment
	Performer []MedicationAdministrationPerformer `json:"performer,omitempty"`

	// A code, Condition or observation that supports why the medication was administered
	Reason []CodeableReference `json:"reason,omitempty"`

	// The date the occurrence of the MedicationAdministration was first captured in the record
	Recorded *string `json:"recorded,omitempty"`

	// The original request that this medication administration is based on
	Request *common.Reference `json:"request,omitempty"`

	// Will generally be set to show that the administration has been completed
	Status MedicationAdministrationStatus `json:"status"`

	// A code indicating why the administration was not performed
	StatusReason []common.CodeableConcept `json:"statusReason,omitempty"`

	// The person or animal or group receiving the medication
	Subject common.Reference `json:"subject"`

	// The reason or reasons why the full dose was not administered
	SubPotentReason []common.CodeableConcept `json:"subPotentReason,omitempty"`

	// Additional information that supports the administration of the medication
	SupportingInformation []common.Reference `json:"supportingInformation,omitempty"`
}

// MedicationDispenseStatus represents the status of a medication dispense
type MedicationDispenseStatus string

const (
	MedicationDispenseStatusPreparation    MedicationDispenseStatus = "preparation"
	MedicationDispenseStatusInProgress     MedicationDispenseStatus = "in-progress"
	MedicationDispenseStatusCancelled      MedicationDispenseStatus = "cancelled"
	MedicationDispenseStatusOnHold         MedicationDispenseStatus = "on-hold"
	MedicationDispenseStatusCompleted      MedicationDispenseStatus = "completed"
	MedicationDispenseStatusEnteredInError MedicationDispenseStatus = "entered-in-error"
	MedicationDispenseStatusStopped        MedicationDispenseStatus = "stopped"
	MedicationDispenseStatusDeclined       MedicationDispenseStatus = "declined"
	MedicationDispenseStatusUnknown        MedicationDispenseStatus = "unknown"
)

// MedicationDispensePerformer represents the performer of the medication dispense
type MedicationDispensePerformer struct {
	common.BackboneElement

	// Indicates who or what performed the event
	Actor common.Reference `json:"actor"`

	// Distinguishes the type of performer in the dispense
	Function *common.CodeableConcept `json:"function,omitempty"`
}

// MedicationDispenseSubstitution represents whether a substitution was made as part of the dispense
type MedicationDispenseSubstitution struct {
	common.BackboneElement

	// A code signifying whether a different drug was dispensed from what was prescribed
	Type *common.CodeableConcept `json:"type,omitempty"`

	// Indicates the reason for the substitution
	Reason []common.CodeableConcept `json:"reason,omitempty"`

	// The person or organization that has primary responsibility for the substitution
	ResponsibleParty *common.Reference `json:"responsibleParty,omitempty"`

	// True if the dispenser dispensed a different drug or product from what was prescribed
	WasSubstituted *bool `json:"wasSubstituted,omitempty"`
}

// MedicationDispense represents the provision of a medication to a patient
type MedicationDispense struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "MedicationDispense"

	// Indicates the medication order that is being dispensed against
	AuthorizingPrescription []common.Reference `json:"authorizingPrescription,omitempty"`

	// A plan that is fulfilled in whole or in part by this MedicationDispense
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// The category can be used to include where the medication is expected to be consumed
	Category []common.CodeableConcept `json:"category,omitempty"`

	// The amount of medication expressed as a timing amount
	DaysSupply *common.Quantity `json:"daysSupply,omitempty"`

	// Identification of the facility/location where the medication was/will be shipped to
	Destination *common.Reference `json:"destination,omitempty"`

	// Indicates an actual or potential clinical issue with or between one or more active or proposed clinical actions
	DetectedIssue []common.Reference `json:"detectedIssue,omitempty"`

	// How the medication is to be used by the patient
	DosageInstruction []Dosage `json:"dosageInstruction,omitempty"`

	// The encounter that establishes the context for this event
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Event history references
	EventHistory []common.Reference `json:"eventHistory,omitempty"`

	// Business identifier for this medication dispense
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Where the medication was/will be dispensed to
	Location *common.Reference `json:"location,omitempty"`

	// What medication was supplied
	Medication CodeableReference `json:"medication"`

	// Information about the dispense
	Note []Annotation `json:"note,omitempty"`

	// A larger event of which this particular event is a component or step
	PartOf []common.Reference `json:"partOf,omitempty"`

	// Who performed the dispense and what they did
	Performer []MedicationDispensePerformer `json:"performer,omitempty"`

	// The amount of medication that has been dispensed
	Quantity *common.Quantity `json:"quantity,omitempty"`

	// Who collected the medication or where the medication was delivered
	Receiver []common.Reference `json:"receiver,omitempty"`

	// The date the dispensed product was recorded
	Recorded *string `json:"recorded,omitempty"`

	// The status of the dispense
	Status MedicationDispenseStatus `json:"status"`

	// Indicates the reason why a dispense was not performed
	StatusReason *CodeableReference `json:"statusReason,omitempty"`

	// A link to a resource representing the person or the group to whom the medication will be given
	Subject common.Reference `json:"subject"`

	// Whether a substitution was made as part of the dispense
	Substitution *MedicationDispenseSubstitution `json:"substitution,omitempty"`

	// Additional information that supports the medication being dispensed
	SupportingInformation []common.Reference `json:"supportingInformation,omitempty"`

	// The type of medication dispense
	Type *common.CodeableConcept `json:"type,omitempty"`

	// The time when the dispensed product was packaged and reviewed
	WhenHandedOver *string `json:"whenHandedOver,omitempty"`

	// The time the dispensed product was prepared
	WhenPrepared *string `json:"whenPrepared,omitempty"`
}

// MedicationKnowledgeStatus represents the status of medication knowledge
type MedicationKnowledgeStatus string

const (
	MedicationKnowledgeStatusActive         MedicationKnowledgeStatus = "active"
	MedicationKnowledgeStatusEnteredInError MedicationKnowledgeStatus = "entered-in-error"
	MedicationKnowledgeStatusInactive       MedicationKnowledgeStatus = "inactive"
)

// MedicationKnowledgeDefinitionalDrugCharacteristic represents descriptive properties of the medicine
type MedicationKnowledgeDefinitionalDrugCharacteristic struct {
	common.BackboneElement

	// A code specifying which characteristic of the medicine is being described
	Type *common.CodeableConcept `json:"type,omitempty"`

	// Description of the characteristic
	ValueCodeableConcept *common.CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueString          *string                 `json:"valueString,omitempty"`
	ValueQuantity        *common.Quantity        `json:"valueQuantity,omitempty"`
	ValueBase64Binary    *string                 `json:"valueBase64Binary,omitempty"`
	ValueAttachment      *Attachment             `json:"valueAttachment,omitempty"`
}

// MedicationKnowledgeDefinitionalIngredient represents ingredients of the medication
type MedicationKnowledgeDefinitionalIngredient struct {
	common.BackboneElement

	// The ingredient substance
	Item CodeableReference `json:"item"`

	// Whether the ingredient is active or inactive
	IsActive *bool `json:"isActive,omitempty"`

	// Specifies how many (or how much) of the items there are in this Medication
	StrengthRatio           *Ratio                  `json:"strengthRatio,omitempty"`
	StrengthCodeableConcept *common.CodeableConcept `json:"strengthCodeableConcept,omitempty"`
	StrengthQuantity        *common.Quantity        `json:"strengthQuantity,omitempty"`
}

// MedicationKnowledgeDefinitional represents definitional information about the medication
type MedicationKnowledgeDefinitional struct {
	common.BackboneElement

	// Associated definitions for this medication
	Definition []common.Reference `json:"definition,omitempty"`

	// The dose form for a medication
	DoseForm *common.CodeableConcept `json:"doseForm,omitempty"`

	// Specifies descriptive properties of the medicine
	DrugCharacteristic []MedicationKnowledgeDefinitionalDrugCharacteristic `json:"drugCharacteristic,omitempty"`

	// Identifies a particular constituent of interest in the product
	Ingredient []MedicationKnowledgeDefinitionalIngredient `json:"ingredient,omitempty"`

	// The intended or approved route of administration
	IntendedRoute []common.CodeableConcept `json:"intendedRoute,omitempty"`
}

// MedicationKnowledgeCost represents the price of the medication
type MedicationKnowledgeCost struct {
	common.BackboneElement

	// The price or representation of the cost
	CostMoney           *Money                  `json:"costMoney,omitempty"`
	CostCodeableConcept *common.CodeableConcept `json:"costCodeableConcept,omitempty"`

	// The date range for which the cost is effective
	EffectiveDate []common.Period `json:"effectiveDate,omitempty"`

	// The category of the cost information
	Source *common.CodeableConcept `json:"source,omitempty"`

	// The category of the cost information
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// MedicationKnowledge represents information about a medication that is used to support knowledge
type MedicationKnowledge struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "MedicationKnowledge"

	// Links to associated medications that could be prescribed, dispensed or administered
	AssociatedMedication []common.Reference `json:"associatedMedication,omitempty"`

	// The creator or owner of the knowledge or information about the medication
	Author *common.Reference `json:"author,omitempty"`

	// Potential clinical issue with or between medication(s)
	ClinicalUseIssue []common.Reference `json:"clinicalUseIssue,omitempty"`

	// A code that specifies this medication
	Code *common.CodeableConcept `json:"code,omitempty"`

	// The price of the medication
	Cost []MedicationKnowledgeCost `json:"cost,omitempty"`

	// Definitional information about the medication
	Definitional *MedicationKnowledgeDefinitional `json:"definitional,omitempty"`

	// Business identifier for this medication knowledge
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The programs that the medication is applicable for
	IndicationGuideline []interface{} `json:"indicationGuideline,omitempty"`

	// The maximum number of units of the medication that can be dispensed
	MaxDispenseQuantity *common.Quantity `json:"maxDispenseQuantity,omitempty"`

	// Categorization of the medication within a formulary or classification system
	MedicineClassification []interface{} `json:"medicineClassification,omitempty"`

	// Associated documentation about the medication
	MonitoringProgram []interface{} `json:"monitoringProgram,omitempty"`

	// Associated documentation about the medication
	Monograph []interface{} `json:"monograph,omitempty"`

	// The packaging details for the medication
	Packaging []interface{} `json:"packaging,omitempty"`

	// The time course of drug absorption, distribution, metabolism and excretion
	PharmacokineticProperty []interface{} `json:"pharmacokineticProperty,omitempty"`

	// Program under which the medication is reviewed
	RegulatoryInformation []interface{} `json:"regulatoryInformation,omitempty"`

	// Associated documentation about the medication
	RelatedMedicationKnowledge []interface{} `json:"relatedMedicationKnowledge,omitempty"`

	// The status of this medication knowledge
	Status *MedicationKnowledgeStatus `json:"status,omitempty"`

	// A code to indicate if the medication is in active use
	StorageGuideline []interface{} `json:"storageGuideline,omitempty"`
}

// MedicationStatementStatus represents the status of a medication statement
type MedicationStatementStatus string

const (
	MedicationStatementStatusRecorded       MedicationStatementStatus = "recorded"
	MedicationStatementStatusEnteredInError MedicationStatementStatus = "entered-in-error"
	MedicationStatementStatusDraft          MedicationStatementStatus = "draft"
)

// MedicationStatementAdherence represents adherence information for a medication statement
type MedicationStatementAdherence struct {
	common.BackboneElement

	// Type of the adherence for the medication
	Code common.CodeableConcept `json:"code"`

	// This is generally only used for "exception" statuses such as "entered-in-error"
	Reason *common.CodeableConcept `json:"reason,omitempty"`
}

// MedicationStatement represents a record of a medication that is being consumed by a patient
type MedicationStatement struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "MedicationStatement"

	// This element can be used to indicate whether a patient is following a course of treatment as instructed
	Adherence *MedicationStatementAdherence `json:"adherence,omitempty"`

	// Type of medication statement
	Category []common.CodeableConcept `json:"category,omitempty"`

	// The date when the Medication Statement was asserted by the information source
	DateAsserted *string `json:"dateAsserted,omitempty"`

	// Likely references would be to MedicationRequest, MedicationDispense, Claim, Observation or QuestionnaireAnswers
	DerivedFrom []common.Reference `json:"derivedFrom,omitempty"`

	// The dates included in the dosage on a Medication Statement reflect the dates for a given dose
	Dosage []Dosage `json:"dosage,omitempty"`

	// This attribute reflects the period over which the patient consumed the medication
	EffectiveDateTime *string        `json:"effectiveDateTime,omitempty"`
	EffectivePeriod   *common.Period `json:"effectivePeriod,omitempty"`
	EffectiveTiming   *Timing        `json:"effectiveTiming,omitempty"`

	// The encounter that establishes the context for this MedicationStatement
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Business identifier for this medication statement
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The person or organization that provided the information about the taking of this medication
	InformationSource []common.Reference `json:"informationSource,omitempty"`

	// If only a code is specified, then it needs to be a code for a specific product
	Medication CodeableReference `json:"medication"`

	// Provides extra information about the Medication Statement
	Note []Annotation `json:"note,omitempty"`

	// A larger event of which this particular MedicationStatement is a component or step
	PartOf []common.Reference `json:"partOf,omitempty"`

	// This could be a diagnosis code
	Reason []CodeableReference `json:"reason,omitempty"`

	// Link to information that is relevant to a medication statement
	RelatedClinicalInformation []common.Reference `json:"relatedClinicalInformation,omitempty"`

	// The full representation of the dose of the medication included in all dosage instructions
	RenderedDosageInstruction *string `json:"renderedDosageInstruction,omitempty"`

	// This status concerns just the recording of the medication statement
	Status MedicationStatementStatus `json:"status"`

	// The person, animal or group who is/was taking the medication
	Subject common.Reference `json:"subject"`
}

// MedicinalProductDefinitionContact represents a product specific contact
type MedicinalProductDefinitionContact struct {
	common.BackboneElement

	// A product specific contact, person (in a role), or an organization
	Contact common.Reference `json:"contact"`

	// Allows the contact to be classified
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// MedicinalProductDefinitionNamePart represents coding words or phrases of the name
type MedicinalProductDefinitionNamePart struct {
	common.BackboneElement

	// A fragment of a product name
	Part string `json:"part"`

	// Identifying type for this part of the name
	Type common.CodeableConcept `json:"type"`
}

// MedicinalProductDefinitionNameUsage represents country and jurisdiction where the name applies
type MedicinalProductDefinitionNameUsage struct {
	common.BackboneElement

	// Country code for where this name applies
	Country common.CodeableConcept `json:"country"`

	// Jurisdiction code for where this name applies
	Jurisdiction *common.CodeableConcept `json:"jurisdiction,omitempty"`

	// Language code for this name
	Language common.CodeableConcept `json:"language"`
}

// MedicinalProductDefinitionName represents the product's name
type MedicinalProductDefinitionName struct {
	common.BackboneElement

	// Coding words or phrases of the name
	Part []MedicinalProductDefinitionNamePart `json:"part,omitempty"`

	// The full product name
	ProductName string `json:"productName"`

	// Type of product name, such as rINN, BAN, Proprietary, Non-Proprietary
	Type *common.CodeableConcept `json:"type,omitempty"`

	// Country and jurisdiction where the name applies, and associated language
	Usage []MedicinalProductDefinitionNameUsage `json:"usage,omitempty"`
}

// MedicinalProductDefinitionCrossReference represents reference to another product
type MedicinalProductDefinitionCrossReference struct {
	common.BackboneElement

	// Reference to another product
	Product CodeableReference `json:"product"`

	// The type of relationship
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// MedicinalProductDefinitionOperation represents a manufacturing or administrative process
type MedicinalProductDefinitionOperation struct {
	common.BackboneElement

	// Specifies whether this particular business or manufacturing process is considered proprietary or confidential
	ConfidentialityIndicator *common.CodeableConcept `json:"confidentialityIndicator,omitempty"`

	// Date range of applicability
	EffectiveDate *common.Period `json:"effectiveDate,omitempty"`

	// The organization or establishment responsible for the particular process or step
	Organization []common.Reference `json:"organization,omitempty"`

	// The type of manufacturing operation
	Type *CodeableReference `json:"type,omitempty"`
}

// MedicinalProductDefinitionCharacteristic represents key product features
type MedicinalProductDefinitionCharacteristic struct {
	common.BackboneElement

	// A code expressing the type of characteristic
	Type common.CodeableConcept `json:"type"`

	// Description of the characteristic
	ValueCodeableConcept *common.CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueMarkdown        *string                 `json:"valueMarkdown,omitempty"`
	ValueQuantity        *common.Quantity        `json:"valueQuantity,omitempty"`
	ValueInteger         *int                    `json:"valueInteger,omitempty"`
	ValueDate            *string                 `json:"valueDate,omitempty"`
	ValueBoolean         *bool                   `json:"valueBoolean,omitempty"`
	ValueAttachment      *Attachment             `json:"valueAttachment,omitempty"`
}

// MedicinalProductDefinition represents a medicinal product
type MedicinalProductDefinition struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "MedicinalProductDefinition"

	// Whether the Medicinal Product is subject to additional monitoring for regulatory reasons
	AdditionalMonitoringIndicator *common.CodeableConcept `json:"additionalMonitoringIndicator,omitempty"`

	// Additional information or supporting documentation about the medicinal product
	AttachedDocument []common.Reference `json:"attachedDocument,omitempty"`

	// Allows the key product features to be recorded
	Characteristic []MedicinalProductDefinitionCharacteristic `json:"characteristic,omitempty"`

	// Allows the product to be classified by various systems, commonly WHO ATC
	Classification []common.CodeableConcept `json:"classification,omitempty"`

	// Clinical trials or studies that this product is involved in
	ClinicalTrial []common.Reference `json:"clinicalTrial,omitempty"`

	// A code that this product is known by
	Code []common.Coding `json:"code,omitempty"`

	// The dose form for a single part product, or combined form of a multiple part product
	CombinedPharmaceuticalDoseForm *common.CodeableConcept `json:"combinedPharmaceuticalDoseForm,omitempty"`

	// Types of medicinal manufactured items and/or devices that this product consists of
	ComprisedOf []common.Reference `json:"comprisedOf,omitempty"`

	// A product specific contact, person (in a role), or an organization
	Contact []MedicinalProductDefinitionContact `json:"contact,omitempty"`

	// Reference to another product
	CrossReference []MedicinalProductDefinitionCrossReference `json:"crossReference,omitempty"`

	// General description of this product
	Description *string `json:"description,omitempty"`

	// If this medicine applies to human or veterinary uses
	Domain *common.CodeableConcept `json:"domain,omitempty"`

	// Business identifier for this product
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Any component of the drug product which is not the chemical entity defined as the drug substance
	Impurity []CodeableReference `json:"impurity,omitempty"`

	// Description of indication(s) for this product
	Indication *string `json:"indication,omitempty"`

	// The ingredients of this medicinal product
	Ingredient []common.CodeableConcept `json:"ingredient,omitempty"`

	// The legal status of supply of the medicinal product as classified by the regulator
	LegalStatusOfSupply *common.CodeableConcept `json:"legalStatusOfSupply,omitempty"`

	// Marketing status of the medicinal product
	MarketingStatus []interface{} `json:"marketingStatus,omitempty"`

	// A master file for the medicinal product
	MasterFile []common.Reference `json:"masterFile,omitempty"`

	// The product's name, including full name and possibly coded parts
	Name []MedicinalProductDefinitionName `json:"name"`

	// A manufacturing or administrative process or step associated with the medicinal product
	Operation []MedicinalProductDefinitionOperation `json:"operation,omitempty"`

	// Package type for the product
	PackagedMedicinalProduct []common.CodeableConcept `json:"packagedMedicinalProduct,omitempty"`

	// If authorised for use in children, or infants, neonates etc.
	PediatricUseIndicator *common.CodeableConcept `json:"pediatricUseIndicator,omitempty"`

	// The path by which the product is taken into or makes contact with the body
	Route []common.CodeableConcept `json:"route,omitempty"`

	// Whether the Medicinal Product is subject to special measures for regulatory reasons
	SpecialMeasures []common.CodeableConcept `json:"specialMeasures,omitempty"`

	// The status within the lifecycle of this product record
	Status *common.CodeableConcept `json:"status,omitempty"`

	// The date at which the given status became applicable
	StatusDate *string `json:"statusDate,omitempty"`

	// Regulatory type, e.g. Investigational or Authorized
	Type *common.CodeableConcept `json:"type,omitempty"`

	// A business identifier relating to a specific version of the product
	Version *string `json:"version,omitempty"`
}

// MessageDefinitionStatus represents the status of a message definition
type MessageDefinitionStatus string

const (
	MessageDefinitionStatusDraft   MessageDefinitionStatus = "draft"
	MessageDefinitionStatusActive  MessageDefinitionStatus = "active"
	MessageDefinitionStatusRetired MessageDefinitionStatus = "retired"
	MessageDefinitionStatusUnknown MessageDefinitionStatus = "unknown"
)

// MessageDefinitionCategory represents the impact of the content of the message
type MessageDefinitionCategory string

const (
	MessageDefinitionCategoryConsequence  MessageDefinitionCategory = "consequence"
	MessageDefinitionCategoryCurrency     MessageDefinitionCategory = "currency"
	MessageDefinitionCategoryNotification MessageDefinitionCategory = "notification"
)

// MessageDefinitionResponseRequired represents when a response is required
type MessageDefinitionResponseRequired string

const (
	MessageDefinitionResponseRequiredAlways    MessageDefinitionResponseRequired = "always"
	MessageDefinitionResponseRequiredOnError   MessageDefinitionResponseRequired = "on-error"
	MessageDefinitionResponseRequiredNever     MessageDefinitionResponseRequired = "never"
	MessageDefinitionResponseRequiredOnSuccess MessageDefinitionResponseRequired = "on-success"
)

// MessageDefinitionFocus represents identifies the resource that are being addressed by the event
type MessageDefinitionFocus struct {
	common.BackboneElement

	// The kind of resource that must be the focus for this message
	Code string `json:"code"`

	// Maximum number of resources of this type that must be pointed to by a message
	Max *string `json:"max,omitempty"`

	// Minimum number of resources of this type that must be pointed to by a message
	Min int `json:"min"`

	// A profile that reflects constraints for the focal resource
	Profile *string `json:"profile,omitempty"`
}

// MessageDefinitionAllowedResponse represents allowed response to the message
type MessageDefinitionAllowedResponse struct {
	common.BackboneElement

	// A reference to the message definition that must be adhered to by this supported response
	Message string `json:"message"`

	// Provides a description of the circumstances in which this response should be used
	Situation *string `json:"situation,omitempty"`
}

// MessageDefinition represents the characteristics of a message that can be shared between systems
type MessageDefinition struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "MessageDefinition"

	// This indicates an application level response to "close" a transaction
	AllowedResponse []MessageDefinitionAllowedResponse `json:"allowedResponse,omitempty"`

	// The MessageDefinition that is the basis for the contents of this resource
	Base *string `json:"base,omitempty"`

	// The impact of the content of the message
	Category *MessageDefinitionCategory `json:"category,omitempty"`

	// May be a web site, an email address, a telephone number, etc.
	Contact []ContactDetail `json:"contact,omitempty"`

	// A copyright statement relating to the message definition and/or its contents
	Copyright *string `json:"copyright,omitempty"`

	// Copyright holder and year(s)
	CopyrightLabel *string `json:"copyrightLabel,omitempty"`

	// The date is often not tracked until the resource is published
	Date string `json:"date"`

	// This description can be used to capture details such as comments about misuse
	Description *string `json:"description,omitempty"`

	// Event code or link to the EventDefinition
	EventCoding *common.Coding `json:"eventCoding,omitempty"`
	EventUri    *string        `json:"eventUri,omitempty"`

	// Allows filtering of message definitions that are appropriate for use versus not
	Experimental *bool `json:"experimental,omitempty"`

	// Identifies the resource (or resources) that are being addressed by the event
	Focus []MessageDefinitionFocus `json:"focus,omitempty"`

	// Canonical reference to a GraphDefinition
	Graph *string `json:"graph,omitempty"`

	// Typically, this is used for identifiers that can go in an HL7 V3 II data type
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// It may be possible for the message definition to be used in jurisdictions other than those for which it was originally designed
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`

	// The name is not expected to be globally unique
	Name *string `json:"name,omitempty"`

	// It should be possible to use MessageDefinition to describe a message to be used by certain steps in a particular protocol
	Parent []string `json:"parent,omitempty"`

	// Usually an organization but may be an individual
	Publisher *string `json:"publisher,omitempty"`

	// This element does not describe the usage of the message definition
	Purpose *string `json:"purpose,omitempty"`

	// A MessageDefinition that is superseded by this definition
	Replaces []string `json:"replaces,omitempty"`

	// This enables the capability currently available through MSH-16 in HL7 Version 2
	ResponseRequired *MessageDefinitionResponseRequired `json:"responseRequired,omitempty"`

	// Allows filtering of message definitions that are appropriate for use versus not
	Status MessageDefinitionStatus `json:"status"`

	// This name does not need to be machine-processing friendly and may contain punctuation, white-space, etc.
	Title *string `json:"title,omitempty"`

	// Can be a urn:uuid: or a urn:oid: but real http: addresses are preferred
	Url *string `json:"url,omitempty"`

	// When multiple useContexts are specified, there is no expectation that all or any of the contexts apply
	UseContext []UsageContext `json:"useContext,omitempty"`

	// There may be different message definition instances that have the same identifier but different versions
	Version *string `json:"version,omitempty"`
}

// MessageHeaderDestination represents destination for message header
type MessageHeaderDestination struct {
	common.BackboneElement

	// URL for the destination
	EndpointUrl *string `json:"endpointUrl,omitempty"`

	// Reference to the destination endpoint
	EndpointReference *common.Reference `json:"endpointReference,omitempty"`

	// Name of the destination
	Name *string `json:"name,omitempty"`

	// Specific person or department to receive the message
	Receiver *common.Reference `json:"receiver,omitempty"`

	// Target system for the message
	Target *common.Reference `json:"target,omitempty"`
}

// MessageHeaderSource represents source of message header
type MessageHeaderSource struct {
	common.BackboneElement

	// Contact information for the source
	Contact *ContactPoint `json:"contact,omitempty"`

	// URL for the source
	EndpointUrl *string `json:"endpointUrl,omitempty"`

	// Reference to the source endpoint
	EndpointReference *common.Reference `json:"endpointReference,omitempty"`

	// Name of the source
	Name *string `json:"name,omitempty"`

	// Software running at the source
	Software *string `json:"software,omitempty"`

	// Version of the software
	Version *string `json:"version,omitempty"`
}

// MessageHeaderResponseCode represents response code for message header
type MessageHeaderResponseCode string

const (
	MessageHeaderResponseCodeOK             MessageHeaderResponseCode = "ok"
	MessageHeaderResponseCodeTransientError MessageHeaderResponseCode = "transient-error"
	MessageHeaderResponseCodeFatalError     MessageHeaderResponseCode = "fatal-error"
)

// MessageHeaderResponse represents response information for message header
type MessageHeaderResponse struct {
	common.BackboneElement

	// Response code
	Code MessageHeaderResponseCode `json:"code"`

	// Details about the response
	Details *common.Reference `json:"details,omitempty"`

	// Identifier of the message being responded to
	Identifier common.Identifier `json:"identifier"`
}

// MessageHeader represents the header for a message exchange
type MessageHeader struct {
	DomainResource

	// Resource type
	ResourceType string `json:"resourceType"`

	// Author of the message
	Author *common.Reference `json:"author,omitempty"`

	// Link to the message definition
	Definition *string `json:"definition,omitempty"`

	// Message destinations
	Destination []MessageHeaderDestination `json:"destination,omitempty"`

	// Event the message represents (coding)
	EventCoding *common.Coding `json:"eventCoding,omitempty"`

	// Event the message represents (canonical)
	EventCanonical *string `json:"eventCanonical,omitempty"`

	// The actual data content
	Focus []common.Reference `json:"focus,omitempty"`

	// Cause of the event
	Reason *common.CodeableConcept `json:"reason,omitempty"`

	// Response information
	Response *MessageHeaderResponse `json:"response,omitempty"`

	// Responsible party
	Responsible *common.Reference `json:"responsible,omitempty"`

	// Message sender
	Sender *common.Reference `json:"sender,omitempty"`

	// Message source
	Source MessageHeaderSource `json:"source"`
}

// MolecularSequenceType represents sequence type
type MolecularSequenceType string

const (
	MolecularSequenceTypeAA  MolecularSequenceType = "aa"
	MolecularSequenceTypeDNA MolecularSequenceType = "dna"
	MolecularSequenceTypeRNA MolecularSequenceType = "rna"
)

// MolecularSequenceRelativeStartingSequenceOrientation represents orientation
type MolecularSequenceRelativeStartingSequenceOrientation string

const (
	MolecularSequenceRelativeStartingSequenceOrientationSense     MolecularSequenceRelativeStartingSequenceOrientation = "sense"
	MolecularSequenceRelativeStartingSequenceOrientationAntisense MolecularSequenceRelativeStartingSequenceOrientation = "antisense"
)

// MolecularSequenceRelativeStartingSequenceStrand represents strand
type MolecularSequenceRelativeStartingSequenceStrand string

const (
	MolecularSequenceRelativeStartingSequenceStrandWatson MolecularSequenceRelativeStartingSequenceStrand = "watson"
	MolecularSequenceRelativeStartingSequenceStrandCrick  MolecularSequenceRelativeStartingSequenceStrand = "crick"
)

// MolecularSequenceRelativeStartingSequence represents starting sequence
type MolecularSequenceRelativeStartingSequence struct {
	common.BackboneElement

	// Structural unit composed of a nucleic acid molecule
	Chromosome *common.CodeableConcept `json:"chromosome,omitempty"`

	// The genome assembly used for starting sequence
	GenomeAssembly *common.CodeableConcept `json:"genomeAssembly,omitempty"`

	// A relative reference to a DNA strand based on gene orientation
	Orientation *MolecularSequenceRelativeStartingSequenceOrientation `json:"orientation,omitempty"`

	// A starting sequence may be represented in one of three ways
	SequenceCodeableConcept *common.CodeableConcept `json:"sequenceCodeableConcept,omitempty"`
	SequenceString          *string                 `json:"sequenceString,omitempty"`
	SequenceReference       *common.Reference       `json:"sequenceReference,omitempty"`

	// An absolute reference to a strand
	Strand *MolecularSequenceRelativeStartingSequenceStrand `json:"strand,omitempty"`

	// End position of the window on the starting sequence
	WindowEnd *int `json:"windowEnd,omitempty"`

	// Start position of the window on the starting sequence
	WindowStart *int `json:"windowStart,omitempty"`
}

// MolecularSequenceRelativeEdit represents changes in sequence
type MolecularSequenceRelativeEdit struct {
	common.BackboneElement

	// End position of the edit on the starting sequence
	End *int `json:"end,omitempty"`

	// Allele in the starting sequence
	ReplacedSequence *string `json:"replacedSequence,omitempty"`

	// Allele that was observed
	ReplacementSequence *string `json:"replacementSequence,omitempty"`

	// Start position of the edit on the starting sequence
	Start *int `json:"start,omitempty"`
}

// MolecularSequenceRelative represents a sequence defined relative to another sequence
type MolecularSequenceRelative struct {
	common.BackboneElement

	// These are different ways of identifying nucleotides or amino acids within a sequence
	CoordinateSystem common.CodeableConcept `json:"coordinateSystem"`

	// Changes in sequence from the starting sequence
	Edit []MolecularSequenceRelativeEdit `json:"edit,omitempty"`

	// Indicates the order in which the sequence should be considered
	OrdinalPosition *int `json:"ordinalPosition,omitempty"`

	// Indicates the nucleotide range in the composed sequence
	SequenceRange *Range `json:"sequenceRange,omitempty"`

	// A sequence that is used as a starting sequence
	StartingSequence *MolecularSequenceRelativeStartingSequence `json:"startingSequence,omitempty"`
}

// MolecularSequence represents a molecular sequence
type MolecularSequence struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "MolecularSequence"

	// The method for sequencing
	Device *common.Reference `json:"device,omitempty"`

	// The actual focus of a molecular sequence
	Focus []common.Reference `json:"focus,omitempty"`

	// Sequence that was observed as file content
	Formatted []Attachment `json:"formatted,omitempty"`

	// A unique identifier for this particular sequence instance
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Sequence that was observed
	Literal *string `json:"literal,omitempty"`

	// The organization or lab that should be responsible for this result
	Performer *common.Reference `json:"performer,omitempty"`

	// A sequence defined relative to another sequence
	Relative []MolecularSequenceRelative `json:"relative,omitempty"`

	// Specimen used for sequencing
	Specimen *common.Reference `json:"specimen,omitempty"`

	// Indicates the subject this sequence is associated too
	Subject *common.Reference `json:"subject,omitempty"`

	// Amino Acid Sequence/ DNA Sequence / RNA Sequence
	Type *MolecularSequenceType `json:"type,omitempty"`
}

// NamingSystemStatus represents status of naming system
type NamingSystemStatus string

const (
	NamingSystemStatusDraft   NamingSystemStatus = "draft"
	NamingSystemStatusActive  NamingSystemStatus = "active"
	NamingSystemStatusRetired NamingSystemStatus = "retired"
	NamingSystemStatusUnknown NamingSystemStatus = "unknown"
)

// NamingSystemKind represents kind of naming system
type NamingSystemKind string

const (
	NamingSystemKindCodesystem NamingSystemKind = "codesystem"
	NamingSystemKindIdentifier NamingSystemKind = "identifier"
	NamingSystemKindRoot       NamingSystemKind = "root"
)

// NamingSystemUniqueIdType represents type of unique id
type NamingSystemUniqueIdType string

const (
	NamingSystemUniqueIdTypeOID          NamingSystemUniqueIdType = "oid"
	NamingSystemUniqueIdTypeUUID         NamingSystemUniqueIdType = "uuid"
	NamingSystemUniqueIdTypeURI          NamingSystemUniqueIdType = "uri"
	NamingSystemUniqueIdTypeIRIStem      NamingSystemUniqueIdType = "iri-stem"
	NamingSystemUniqueIdTypeV2CSMnemonic NamingSystemUniqueIdType = "v2csmnemonic"
	NamingSystemUniqueIdTypeOther        NamingSystemUniqueIdType = "other"
)

// NamingSystemUniqueId represents unique identifier for naming system
type NamingSystemUniqueId struct {
	common.BackboneElement

	// Indicates whether this identifier is endorsed by the official owner
	Authoritative *bool `json:"authoritative,omitempty"`

	// Comment about the identifier
	Comment *string `json:"comment,omitempty"`

	// The period of time this identifier is valid
	Period *common.Period `json:"period,omitempty"`

	// Indicates whether this identifier is the "preferred" identifier
	Preferred *bool `json:"preferred,omitempty"`

	// The type of identifier
	Type NamingSystemUniqueIdType `json:"type"`

	// The string that should be sent over the wire
	Value string `json:"value"`
}

// NamingSystem represents a namespace that issues unique symbols
type NamingSystem struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "NamingSystem"

	// The approval date
	ApprovalDate *string `json:"approvalDate,omitempty"`

	// An individual or organization primarily involved in the creation and maintenance
	Author []ContactDetail `json:"author,omitempty"`

	// Contact details for the publisher
	Contact []ContactDetail `json:"contact,omitempty"`

	// Copyright statement
	Copyright *string `json:"copyright,omitempty"`

	// Copyright holder and year(s)
	CopyrightLabel *string `json:"copyrightLabel,omitempty"`

	// The date when the naming system was published
	Date string `json:"date"`

	// Natural language description of the naming system
	Description *string `json:"description,omitempty"`

	// An individual or organization primarily responsible for internal coherence
	Editor []ContactDetail `json:"editor,omitempty"`

	// When the naming system is expected to be used
	EffectivePeriod *common.Period `json:"effectivePeriod,omitempty"`

	// Who endorsed the naming system
	Endorser []ContactDetail `json:"endorser,omitempty"`

	// For testing purposes, not real usage
	Experimental *bool `json:"experimental,omitempty"`

	// Typically, this is used for identifiers that can go in an HL7 V3 II data type
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Intended jurisdiction for naming system
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`

	// Indicates the purpose for the naming system
	Kind NamingSystemKind `json:"kind"`

	// Last review date
	LastReviewDate *string `json:"lastReviewDate,omitempty"`

	// Name for this naming system
	Name string `json:"name"`

	// Name of the publisher
	Publisher *string `json:"publisher,omitempty"`

	// Why this naming system is defined
	Purpose *string `json:"purpose,omitempty"`

	// Additional documentation, citations, etc.
	RelatedArtifact []interface{} `json:"relatedArtifact,omitempty"`

	// Organization that is responsible for issuing identifiers or codes
	Responsible *string `json:"responsible,omitempty"`

	// Who reviewed the naming system
	Reviewer []ContactDetail `json:"reviewer,omitempty"`

	// draft | active | retired | unknown
	Status NamingSystemStatus `json:"status"`

	// Human-readable title
	Title *string `json:"title,omitempty"`

	// E.g. Education, Treatment, Assessment, etc.
	Topic []common.CodeableConcept `json:"topic,omitempty"`

	// Categorizes a naming system for easier search
	Type *common.CodeableConcept `json:"type,omitempty"`

	// Indicates how the system may be identified when referenced
	UniqueId []NamingSystemUniqueId `json:"uniqueId"`

	// Canonical identifier for this naming system
	Url *string `json:"url,omitempty"`

	// Provides guidance on the use of the namespace
	Usage *string `json:"usage,omitempty"`

	// The context that the content is intended to support
	UseContext []interface{} `json:"useContext,omitempty"`
}

// PlanDefinitionStatus represents status of plan definition
type PlanDefinitionStatus string

const (
	PlanDefinitionStatusDraft   PlanDefinitionStatus = "draft"
	PlanDefinitionStatusActive  PlanDefinitionStatus = "active"
	PlanDefinitionStatusRetired PlanDefinitionStatus = "retired"
	PlanDefinitionStatusUnknown PlanDefinitionStatus = "unknown"
)

// PlanDefinitionGoal represents a goal in a plan definition
type PlanDefinitionGoal struct {
	common.BackboneElement

	// Identifies problems, conditions, issues, or concerns the goal is intended to address
	Addresses []common.CodeableConcept `json:"addresses,omitempty"`

	// Indicates a category the goal falls within
	Category *common.CodeableConcept `json:"category,omitempty"`

	// Human-readable and/or coded description of a specific desired objective of care
	Description common.CodeableConcept `json:"description"`

	// Didactic or other informational resources associated with the goal
	Documentation []interface{} `json:"documentation,omitempty"`

	// Identifies the expected level of importance associated with reaching/sustaining the defined goal
	Priority *common.CodeableConcept `json:"priority,omitempty"`

	// The event after which the goal should begin being pursued
	Start *common.CodeableConcept `json:"start,omitempty"`

	// Indicates what should be done and within what timeframe
	Target []interface{} `json:"target,omitempty"`
}

// PlanDefinitionAction represents an action in a plan definition
type PlanDefinitionAction struct {
	common.BackboneElement

	// Sub actions that are contained within the action
	Action []PlanDefinitionAction `json:"action,omitempty"`

	// Defines whether the action can be selected multiple times
	CardinalityBehavior *string `json:"cardinalityBehavior,omitempty"`

	// A code that provides a meaning, grouping, or classification for the action or action group
	Code *common.CodeableConcept `json:"code,omitempty"`

	// When multiple conditions of the same kind are present, the effects are combined using AND semantics
	Condition []interface{} `json:"condition,omitempty"`

	// Note that the definition is optional, and if no definition is specified, a dynamicValue with a root ($this) path can be used to define the entire resource dynamically
	DefinitionCanonical *string `json:"definitionCanonical,omitempty"`

	// Note that the definition is optional, and if no definition is specified, a dynamicValue with a root ($this) path can be used to define the entire resource dynamically
	DefinitionUri *string `json:"definitionUri,omitempty"`

	// A short description of the action used to provide a summary to display to the user
	Description *string `json:"description,omitempty"`

	// Didactic or other informational resources associated with the action
	Documentation []interface{} `json:"documentation,omitempty"`

	// Dynamic values that will be evaluated to produce values for elements of the resulting resource
	DynamicValue []interface{} `json:"dynamicValue,omitempty"`

	// An expression that describes applicability criteria, or start/stop conditions for the action
	Expression *interface{} `json:"expression,omitempty"`

	// A goal that this action helps to achieve within the plan
	GoalId []string `json:"goalId,omitempty"`

	// Defines the grouping behavior for the action and its children
	GroupingBehavior *string `json:"groupingBehavior,omitempty"`

	// Defines the input data requirements for the action
	Input []interface{} `json:"input,omitempty"`

	// Defines the output data requirements for the action
	Output []interface{} `json:"output,omitempty"`

	// The participant that should perform or be responsible for this action
	Participant []interface{} `json:"participant,omitempty"`

	// Defines whether the action should usually be preselected
	PrecheckBehavior *string `json:"precheckBehavior,omitempty"`

	// Defines the required behavior for the action
	RequiredBehavior *string `json:"requiredBehavior,omitempty"`

	// Defines the selection behavior for the action and its children
	SelectionBehavior *string `json:"selectionBehavior,omitempty"`

	// A text equivalent of the action to be performed
	TextEquivalent *string `json:"textEquivalent,omitempty"`

	// A time period when the action should occur
	TimingDateTime *string        `json:"timingDateTime,omitempty"`
	TimingAge      *Age           `json:"timingAge,omitempty"`
	TimingPeriod   *common.Period `json:"timingPeriod,omitempty"`
	TimingDuration *Duration      `json:"timingDuration,omitempty"`
	TimingRange    *Range         `json:"timingRange,omitempty"`
	TimingTiming   *Timing        `json:"timingTiming,omitempty"`

	// The title of the action displayed to a user
	Title *string `json:"title,omitempty"`

	// A transformation to apply to the input data before the data is used by the action
	Transform *string `json:"transform,omitempty"`

	// A description of when this action should be triggered
	Trigger []interface{} `json:"trigger,omitempty"`

	// The type of action to perform (create, update, remove)
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// PlanDefinition represents a plan definition
type PlanDefinition struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "PlanDefinition"

	// Note that there is overlap between many of the elements defined here and the ActivityDefinition resource
	Action []PlanDefinitionAction `json:"action,omitempty"`

	// Actors represent the individuals or groups involved in the execution of the defined set of activities
	Actor []interface{} `json:"actor,omitempty"`

	// The 'date' element may be more recent than the approval date because of minor changes or editorial corrections
	ApprovalDate *string `json:"approvalDate,omitempty"`

	// If a CodeableConcept is present, it indicates the pre-condition for performing the service
	AsNeededBoolean         *bool                   `json:"asNeededBoolean,omitempty"`
	AsNeededCodeableConcept *common.CodeableConcept `json:"asNeededCodeableConcept,omitempty"`

	// An individiual or organization primarily involved in the creation and maintenance of the content
	Author []ContactDetail `json:"author,omitempty"`

	// May be a web site, an email address, a telephone number, etc
	Contact []ContactDetail `json:"contact,omitempty"`

	// A copyright statement relating to the plan definition and/or its contents
	Copyright *string `json:"copyright,omitempty"`

	// Note that this is not the same as the resource last-modified-date, since the resource may be a secondary representation of the plan definition
	Date *string `json:"date,omitempty"`

	// This description can be used to capture details such as why the plan definition was built
	Description *string `json:"description,omitempty"`

	// An individual or organization primarily responsible for internal coherence of the content
	Editor []ContactDetail `json:"editor,omitempty"`

	// The effective period for a plan definition determines when the content is applicable for usage
	EffectivePeriod *common.Period `json:"effectivePeriod,omitempty"`

	// An individual or organization responsible for officially endorsing the content for use in some setting
	Endorser []ContactDetail `json:"endorser,omitempty"`

	// Allows filtering of plan definitions that are appropriate for use versus not
	Experimental *bool `json:"experimental,omitempty"`

	// A goal describes an expected outcome that activities within the plan are intended to achieve
	Goal []PlanDefinitionGoal `json:"goal,omitempty"`

	// Typically, this is used for identifiers that can go in an HL7 V3 II (instance identifier) data type
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// It may be possible for the plan definition to be used in jurisdictions other than those for which it was originally designed or intended
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`

	// If specified, this date follows the original approval date
	LastReviewDate *string `json:"lastReviewDate,omitempty"`

	// A reference to a Library resource containing any formal logic used by the plan definition
	Library []string `json:"library,omitempty"`

	// The name is not expected to be globally unique
	Name *string `json:"name,omitempty"`

	// Usually an organization but may be an individual
	Publisher *string `json:"publisher,omitempty"`

	// This element does not describe the usage of the plan definition
	Purpose *string `json:"purpose,omitempty"`

	// Each related artifact is either an attachment, or a reference to another resource, but not both
	RelatedArtifact []interface{} `json:"relatedArtifact,omitempty"`

	// An individual or organization primarily responsible for review of some aspect of the content
	Reviewer []ContactDetail `json:"reviewer,omitempty"`

	// Allows filtering of plan definitions that are appropriate for use versus not
	Status PlanDefinitionStatus `json:"status"`

	// Note that the choice of canonical for the subject element was introduced in R4B to support pharmaceutical quality use cases
	SubjectCodeableConcept *common.CodeableConcept `json:"subjectCodeableConcept,omitempty"`
	SubjectReference       *common.Reference       `json:"subjectReference,omitempty"`
	SubjectCanonical       *string                 `json:"subjectCanonical,omitempty"`

	// An explanatory or alternate title for the plan definition giving additional information about its content
	Subtitle *string `json:"subtitle,omitempty"`

	// This name does not need to be machine-processing friendly and may contain punctuation, white-space, etc
	Title *string `json:"title,omitempty"`

	// Descriptive topics related to the content of the plan definition
	Topic []common.CodeableConcept `json:"topic,omitempty"`

	// A high-level category for the plan definition that distinguishes the kinds of systems that would be interested in the plan definition
	Type *common.CodeableConcept `json:"type,omitempty"`

	// Can be a urn:uuid: or a urn:oid: but real http: addresses are preferred
	URL *string `json:"url,omitempty"`

	// A detailed description of how the plan definition is used from a clinical perspective
	Usage *string `json:"usage,omitempty"`

	// When multiple useContexts are specified, there is no expectation that all or any of the contexts apply
	UseContext []UsageContext `json:"useContext,omitempty"`

	// There may be different plan definition instances that have the same identifier but different versions
	Version *string `json:"version,omitempty"`
}

// PractitionerRole represents a specific set of roles/locations/specialties/services that a practitioner may perform at an organization for a period of time
type PractitionerRole struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "PractitionerRole"

	// If this value is false, you may refer to the period to see when the role was in active use
	Active *bool `json:"active,omitempty"`

	// More detailed availability information may be provided in associated Schedule/Slot resources
	Availability []interface{} `json:"availability,omitempty"`

	// These could be such things as is the service mode used by this role
	Characteristic []common.CodeableConcept `json:"characteristic,omitempty"`

	// A person may have more than one role
	Code []common.CodeableConcept `json:"code,omitempty"`

	// The structure aa-BB with this exact casing is one the most widely used notations for locale
	Communication []common.CodeableConcept `json:"communication,omitempty"`

	// The contact details of communication devices available relevant to the specific PractitionerRole
	Contact []interface{} `json:"contact,omitempty"`

	// Technical endpoints providing access to services operated for the practitioner with this role
	Endpoint []common.Reference `json:"endpoint,omitempty"`

	// The list of healthcare services that this worker provides for this role's Organization/Location(s)
	HealthcareService []common.Reference `json:"healthcareService,omitempty"`

	// A specific identifier value may appear on multiple PractitionerRole instances
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The location(s) at which this practitioner provides care
	Location []common.Reference `json:"location,omitempty"`

	// The organization where the Practitioner performs the roles associated
	Organization *common.Reference `json:"organization,omitempty"`

	// If a practitioner is performing a role within an organization over multiple, non-adjacent periods
	Period *common.Period `json:"period,omitempty"`

	// Practitioner that is able to provide the defined services for the organization
	Practitioner *common.Reference `json:"practitioner,omitempty"`

	// The specialty represents the functional role a practitioner is playing within an organization/location
	Specialty []common.CodeableConcept `json:"specialty,omitempty"`
}

// ProvenanceEntity represents an entity used in this activity
type ProvenanceEntity struct {
	common.BackboneElement

	// A usecase where one Provenance.entity.agent is used where the Entity that was used in the creation/updating of the Target
	Agent []interface{} `json:"agent,omitempty"`

	// How the entity was used during the activity
	Role string `json:"role"`

	// whatIdentity should be used for entities that are not a Resource type
	What common.Reference `json:"what"`
}

// Provenance represents provenance of a resource
type Provenance struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Provenance"

	// An activity is something that occurs over a period of time and acts upon or with entities
	Activity *common.CodeableConcept `json:"activity,omitempty"`

	// Several agents may be associated with an activity and vice-versa
	Agent []interface{} `json:"agent"`

	// The authorization that was used during the event being recorded
	Authorization []CodeableReference `json:"authorization,omitempty"`

	// Allows tracing of authorization for the events and tracking whether proposals/recommendations were acted upon
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// This will typically be the encounter the event occurred
	Encounter *common.Reference `json:"encounter,omitempty"`

	// An entity used in this activity
	Entity []ProvenanceEntity `json:"entity,omitempty"`

	// Where the activity occurred, if relevant
	Location *common.Reference `json:"location,omitempty"`

	// The period can be a little arbitrary; where possible, the time should correspond to human assessment of the activity time
	OccurredPeriod   *common.Period `json:"occurredPeriod,omitempty"`
	OccurredDateTime *string        `json:"occurredDateTime,omitempty"`

	// The patient element is available to enable deterministic tracking of activities that involve the patient as the subject
	Patient *common.Reference `json:"patient,omitempty"`

	// For example: Where an OAuth token authorizes, the unique identifier from the OAuth token is placed into the policy element
	Policy []string `json:"policy,omitempty"`

	// This can be a little different from the time stamp on the resource if there is a delay between recording the event
	Recorded *string `json:"recorded,omitempty"`

	// A digital signature on the target Reference(s)
	Signature []Signature `json:"signature,omitempty"`

	// Target references are usually version specific, but might not be
	Target []common.Reference `json:"target"`
}

// QuestionnaireItem represents a single item within a questionnaire
type QuestionnaireItem struct {
	common.BackboneElement

	// A constraint indicating that this item should only be enabled (displayed/allow answers) when the specified condition is true
	EnableBehavior *string `json:"enableBehavior,omitempty"`

	// A constraint indicating that this item should only be enabled when the specified condition is true
	EnableWhen []interface{} `json:"enableWhen,omitempty"`

	// The identifier that is unique to the questionnaire allowing referencing to this item from other questionnaires
	LinkId string `json:"linkId"`

	// The maximum number of characters that are permitted in the answer to be considered a "valid" QuestionnaireResponse
	MaxLength *int `json:"maxLength,omitempty"`

	// The name of a section, the text of a question or text displayed for the item
	Prefix *string `json:"prefix,omitempty"`

	// A particular question, question grouping or display text that is part of the questionnaire
	Question *string `json:"question,omitempty"`

	// An indication, when true, that the item may occur multiple times in the response, collecting multiple answers for questions or multiple sets of answers for groups
	Repeats *bool `json:"repeats,omitempty"`

	// An indication, when true, that the item must be present in a "completed" QuestionnaireResponse
	Required *bool `json:"required,omitempty"`

	// The type of questionnaire item this is - whether text for display, a grouping of other items or a particular type of data to be captured (string, integer, coded choice, etc.)
	Type string `json:"type"`

	// The value that should be defaulted when initially rendering the questionnaire for user input
	Initial []interface{} `json:"initial,omitempty"`

	// Text, questions and other groups to be nested beneath a question or group
	Item []QuestionnaireItem `json:"item,omitempty"`

	// A reference to a value set containing a list of codes representing permitted answers for a "choice" or "open-choice" question
	AnswerOption []interface{} `json:"answerOption,omitempty"`

	// One of the permitted answers for a "choice" or "open-choice" question
	AnswerValueSet *string `json:"answerValueSet,omitempty"`

	// A constraint indicating that this item should only be enabled when the specified condition is true
	Definition *string `json:"definition,omitempty"`

	// The text to be displayed in the questionnaire when this item is displayed
	Text *string `json:"text,omitempty"`
}

// Questionnaire represents a structured set of questions intended to guide the collection of answers from end-users
type Questionnaire struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Questionnaire"

	// The 'date' element may be more recent than the approval date because of minor changes or editorial corrections
	ApprovalDate *string `json:"approvalDate,omitempty"`

	// An individiual or organization primarily involved in the creation and maintenance of the content
	Author []ContactDetail `json:"author,omitempty"`

	// May be a web site, an email address, a telephone number, etc
	Contact []ContactDetail `json:"contact,omitempty"`

	// A copyright statement relating to the questionnaire and/or its contents
	Copyright *string `json:"copyright,omitempty"`

	// Note that this is not the same as the resource last-modified-date, since the resource may be a secondary representation of the questionnaire
	Date *string `json:"date,omitempty"`

	// This description can be used to capture details such as why the questionnaire was built
	Description *string `json:"description,omitempty"`

	// An individual or organization primarily responsible for internal coherence of the content
	Editor []ContactDetail `json:"editor,omitempty"`

	// The effective period for a questionnaire determines when the content is applicable for usage
	EffectivePeriod *common.Period `json:"effectivePeriod,omitempty"`

	// An individual or organization responsible for officially endorsing the content for use in some setting
	Endorser []ContactDetail `json:"endorser,omitempty"`

	// Allows filtering of questionnaires that are appropriate for use versus not
	Experimental *bool `json:"experimental,omitempty"`

	// Typically, this is used for identifiers that can go in an HL7 V3 II (instance identifier) data type
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// It may be possible for the questionnaire to be used in jurisdictions other than those for which it was originally designed or intended
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`

	// If specified, this date follows the original approval date
	LastReviewDate *string `json:"lastReviewDate,omitempty"`

	// The name is not expected to be globally unique
	Name *string `json:"name,omitempty"`

	// Usually an organization but may be an individual
	Publisher *string `json:"publisher,omitempty"`

	// This element does not describe the usage of the questionnaire
	Purpose *string `json:"purpose,omitempty"`

	// Each related artifact is either an attachment, or a reference to another resource, but not both
	RelatedArtifact []interface{} `json:"relatedArtifact,omitempty"`

	// An individual or organization primarily responsible for review of some aspect of the content
	Reviewer []ContactDetail `json:"reviewer,omitempty"`

	// Allows filtering of questionnaires that are appropriate for use versus not
	Status string `json:"status"`

	// An explanatory or alternate title for the questionnaire giving additional information about its content
	Subtitle *string `json:"subtitle,omitempty"`

	// This name does not need to be machine-processing friendly and may contain punctuation, white-space, etc
	Title *string `json:"title,omitempty"`

	// Descriptive topics related to the content of the questionnaire
	Topic []common.CodeableConcept `json:"topic,omitempty"`

	// A high-level category for the questionnaire that distinguishes the kinds of systems that would be interested in the questionnaire
	Type *common.CodeableConcept `json:"type,omitempty"`

	// Can be a urn:uuid: or a urn:oid: but real http: addresses are preferred
	URL *string `json:"url,omitempty"`

	// A detailed description of how the questionnaire is used from a clinical perspective
	Usage *string `json:"usage,omitempty"`

	// When multiple useContexts are specified, there is no expectation that all or any of the contexts apply
	UseContext []UsageContext `json:"useContext,omitempty"`

	// There may be different questionnaire instances that have the same identifier but different versions
	Version *string `json:"version,omitempty"`

	// A particular question, question grouping or display text that is part of the questionnaire
	Item []QuestionnaireItem `json:"item,omitempty"`

	// A reference to a value set containing a list of codes representing permitted answers for a "choice" or "open-choice" question
	Code []common.Coding `json:"code,omitempty"`

	// The subject of the questionnaire
	SubjectType []string `json:"subjectType,omitempty"`
}

// QuestionnaireResponseItemAnswer represents an answer to a question
type QuestionnaireResponseItemAnswer struct {
	common.BackboneElement

	// The answer (or one of the answers) provided by the respondent to the question
	ValueBoolean    *bool             `json:"valueBoolean,omitempty"`
	ValueDecimal    *float64          `json:"valueDecimal,omitempty"`
	ValueInteger    *int              `json:"valueInteger,omitempty"`
	ValueDate       *string           `json:"valueDate,omitempty"`
	ValueDateTime   *string           `json:"valueDateTime,omitempty"`
	ValueTime       *string           `json:"valueTime,omitempty"`
	ValueString     *string           `json:"valueString,omitempty"`
	ValueUri        *string           `json:"valueUri,omitempty"`
	ValueAttachment *Attachment       `json:"valueAttachment,omitempty"`
	ValueCoding     *common.Coding    `json:"valueCoding,omitempty"`
	ValueQuantity   *common.Quantity  `json:"valueQuantity,omitempty"`
	ValueReference  *common.Reference `json:"valueReference,omitempty"`

	// Nested groups and questions within this answer
	Item []QuestionnaireResponseItem `json:"item,omitempty"`
}

// QuestionnaireResponseItem represents a group or question item from the original questionnaire
type QuestionnaireResponseItem struct {
	common.BackboneElement

	// The respondent's answer(s) to the question
	Answer []QuestionnaireResponseItemAnswer `json:"answer,omitempty"`

	// The item from the Questionnaire that corresponds to this item in the QuestionnaireResponse resource
	Definition *string `json:"definition,omitempty"`

	// The identifier that is unique to the questionnaire allowing referencing to this item from other questionnaires
	LinkId string `json:"linkId"`

	// Text, questions and other groups to be nested beneath a question or group
	Item []QuestionnaireResponseItem `json:"item,omitempty"`

	// The name of a section, the text of a question or text displayed for the item
	Prefix *string `json:"prefix,omitempty"`

	// A question or group displayed in the questionnaire response
	Question *string `json:"question,omitempty"`

	// The type of questionnaire item this is - whether text for display, a grouping of other items or a particular type of data to be captured
	Type string `json:"type"`

	// The text to be displayed in the questionnaire response when this item is displayed
	Text *string `json:"text,omitempty"`
}

// QuestionnaireResponse represents a structured set of questions and their answers
type QuestionnaireResponse struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "QuestionnaireResponse"

	// The individual or organization that received and recorded the answers to the questionnaire
	Author *common.Reference `json:"author,omitempty"`

	// The date and/or time that this set of answers was last changed
	Authored *string `json:"authored,omitempty"`

	// The encounter or episode of care with primary association to the questionnaire response
	Encounter *common.Reference `json:"encounter,omitempty"`

	// A group or question item from the original questionnaire for which answers are provided
	Item []QuestionnaireResponseItem `json:"item,omitempty"`

	// A procedure or observation that this questionnaire was performed as part of the execution of
	PartOf []common.Reference `json:"partOf,omitempty"`

	// The Questionnaire that defines and organizes this questionnaire response
	Questionnaire *string `json:"questionnaire,omitempty"`

	// The individual or organization that received the responses to the questionnaire
	Source *common.Reference `json:"source,omitempty"`

	// The position of the questionnaire response within its overall lifecycle
	Status string `json:"status"`

	// The subject of the questionnaire response
	Subject *common.Reference `json:"subject,omitempty"`
}

// RegulatedAuthorization represents a regulatory authorization for a medicinal product
type RegulatedAuthorization struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "RegulatedAuthorization"

	// The case or regulatory procedure for granting or amending a regulated authorization
	Case *interface{} `json:"case,omitempty"`

	// The country in which the authorization has been granted
	Country []common.CodeableConcept `json:"country,omitempty"`

	// The date at which the current status was assigned
	DateOfStatus *string `json:"dateOfStatus,omitempty"`

	// The date when a suspended the marketing or the marketing authorization of the product is anticipated to be restored
	RestoreDate *string `json:"restoreDate,omitempty"`

	// The holder of the regulated authorization
	Holder *common.Reference `json:"holder,omitempty"`

	// Business identifier for the authorization, as assigned by a regulator
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The legal or regulatory framework against which this authorization is granted
	LegalBasis *common.CodeableConcept `json:"legalBasis,omitempty"`

	// The organization that has been granted this authorization, by some authoritative body
	Regulator *common.Reference `json:"regulator,omitempty"`

	// The status that is authorised e.g. approved
	Status *common.CodeableConcept `json:"status,omitempty"`

	// The date at which the given status has become applicable
	StatusDate *string `json:"statusDate,omitempty"`

	// The type of regulated authorization, e.g. marketing or pediatric
	Type *common.CodeableConcept `json:"type,omitempty"`

	// The date when the authorization was granted
	ValidFrom *string `json:"validFrom,omitempty"`

	// The date when the authorization expires
	ValidUntil *string `json:"validUntil,omitempty"`

	// The medicinal product that is being granted the authorization
	Subject []common.Reference `json:"subject,omitempty"`
}

// RelatedPerson represents information about a person that is involved in the care for a patient
type RelatedPerson struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "RelatedPerson"

	// Whether this related person record is in active use
	Active *bool `json:"active,omitempty"`

	// Address where the related person can be contacted or visited
	Address []Address `json:"address,omitempty"`

	// The date on which the related person was born
	BirthDate *string `json:"birthDate,omitempty"`

	// Administrative Gender - the gender that the person is considered to have for administration and record keeping purposes
	Gender *string `json:"gender,omitempty"`

	// Identifier for a person within a particular scope
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The name of the related person
	Name []HumanName `json:"name,omitempty"`

	// The patient this person is related to
	Patient common.Reference `json:"patient"`

	// The nature of the relationship between a patient and the related person
	Relationship []common.CodeableConcept `json:"relationship,omitempty"`

	// A contact detail for the person
	Telecom []ContactPoint `json:"telecom,omitempty"`

	// Image of the person
	Photo []Attachment `json:"photo,omitempty"`

	// Period of time that this relationship is considered valid
	Period *common.Period `json:"period,omitempty"`

	// A language which may be used to communicate with the related person about the patient's health
	Communication []interface{} `json:"communication,omitempty"`
}

// RequestOrchestration represents a set of related requests that can be used to capture and track the state of a request
type RequestOrchestration struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "RequestOrchestration"

	// The actions, if any, produced by the evaluation of the artifact
	Action []interface{} `json:"action,omitempty"`

	// The author of the request orchestration
	Author *common.Reference `json:"author,omitempty"`

	// A plan or request that is fulfilled in whole or in part by this request orchestration
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// The date when this request orchestration was created
	Created *string `json:"created,omitempty"`

	// A group of related requests that can be used to capture and track the state of a request
	GroupIdentifier *common.Identifier `json:"groupIdentifier,omitempty"`

	// Business identifier for the request orchestration
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// A reference to a formal or informal definition of the request orchestration
	InstantiatesCanonical []string `json:"instantiatesCanonical,omitempty"`

	// A reference to a formal or informal definition of the request orchestration
	InstantiatesUri []string `json:"instantiatesUri,omitempty"`

	// The intended performer of the request orchestration
	Intent string `json:"intent"`

	// A description of the request orchestration
	Note []Annotation `json:"note,omitempty"`

	// The priority of the request orchestration
	Priority *string `json:"priority,omitempty"`

	// A reference to a formal or informal definition of the request orchestration
	Replaces []common.Reference `json:"replaces,omitempty"`

	// The current state of the request orchestration
	Status string `json:"status"`

	// The subject of the request orchestration
	Subject *common.Reference `json:"subject,omitempty"`

	// A human-readable narrative that contains a summary of the request orchestration
	Text *Narrative `json:"text,omitempty"`
}

// Requirements represents a set of requirements for a FHIR resource
type Requirements struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Requirements"

	// The 'date' element may be more recent than the approval date because of minor changes or editorial corrections
	ApprovalDate *string `json:"approvalDate,omitempty"`

	// An individiual or organization primarily involved in the creation and maintenance of the content
	Author []ContactDetail `json:"author,omitempty"`

	// May be a web site, an email address, a telephone number, etc
	Contact []ContactDetail `json:"contact,omitempty"`

	// A copyright statement relating to the requirements and/or its contents
	Copyright *string `json:"copyright,omitempty"`

	// Note that this is not the same as the resource last-modified-date, since the resource may be a secondary representation of the requirements
	Date *string `json:"date,omitempty"`

	// This description can be used to capture details such as why the requirements were built
	Description *string `json:"description,omitempty"`

	// An individual or organization primarily responsible for internal coherence of the content
	Editor []ContactDetail `json:"editor,omitempty"`

	// The effective period for a requirements determines when the content is applicable for usage
	EffectivePeriod *common.Period `json:"effectivePeriod,omitempty"`

	// An individual or organization responsible for officially endorsing the content for use in some setting
	Endorser []ContactDetail `json:"endorser,omitempty"`

	// Allows filtering of requirements that are appropriate for use versus not
	Experimental *bool `json:"experimental,omitempty"`

	// Typically, this is used for identifiers that can go in an HL7 V3 II (instance identifier) data type
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// It may be possible for the requirements to be used in jurisdictions other than those for which it was originally designed or intended
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`

	// If specified, this date follows the original approval date
	LastReviewDate *string `json:"lastReviewDate,omitempty"`

	// The name is not expected to be globally unique
	Name *string `json:"name,omitempty"`

	// Usually an organization but may be an individual
	Publisher *string `json:"publisher,omitempty"`

	// This element does not describe the usage of the requirements
	Purpose *string `json:"purpose,omitempty"`

	// Each related artifact is either an attachment, or a reference to another resource, but not both
	RelatedArtifact []interface{} `json:"relatedArtifact,omitempty"`

	// An individual or organization primarily responsible for review of some aspect of the content
	Reviewer []ContactDetail `json:"reviewer,omitempty"`

	// Allows filtering of requirements that are appropriate for use versus not
	Status string `json:"status"`

	// An explanatory or alternate title for the requirements giving additional information about its content
	Subtitle *string `json:"subtitle,omitempty"`

	// This name does not need to be machine-processing friendly and may contain punctuation, white-space, etc
	Title *string `json:"title,omitempty"`

	// Descriptive topics related to the content of the requirements
	Topic []common.CodeableConcept `json:"topic,omitempty"`

	// A high-level category for the requirements that distinguishes the kinds of systems that would be interested in the requirements
	Type *common.CodeableConcept `json:"type,omitempty"`

	// Can be a urn:uuid: or a urn:oid: but real http: addresses are preferred
	URL *string `json:"url,omitempty"`

	// A detailed description of how the requirements are used from a clinical perspective
	Usage *string `json:"usage,omitempty"`

	// When multiple useContexts are specified, there is no expectation that all or any of the contexts apply
	UseContext []interface{} `json:"useContext,omitempty"`

	// There may be different requirements instances that have the same identifier but different versions
	Version *string `json:"version,omitempty"`

	// A statement of business rules that describes the scope, concerns, and objectives of the requirements
	Statement []interface{} `json:"statement,omitempty"`
}

// ResearchStudy represents a process where a researcher or organization plans and then executes a series of steps intended to increase the field of healthcare-related knowledge
type ResearchStudy struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "ResearchStudy"

	// The date on which the research study was approved by the research ethics board
	ApprovalDate *string `json:"approvalDate,omitempty"`

	// Classification for the study
	Category []common.CodeableConcept `json:"category,omitempty"`

	// Contact details to assist a user in learning more about or engaging with the study
	Contact []ContactDetail `json:"contact,omitempty"`

	// A full description of how the study is being conducted
	Description *string `json:"description,omitempty"`

	// The date on which the research study was terminated
	EndDate *string `json:"endDate,omitempty"`

	// The facility where the study is being conducted
	Facility []common.Reference `json:"facility,omitempty"`

	// The condition that is the focus of the research study
	Focus []common.CodeableConcept `json:"focus,omitempty"`

	// Identifiers assigned to this research study by the sponsor or other systems
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// A key word to identify the study
	Keyword []common.CodeableConcept `json:"keyword,omitempty"`

	// The location where the study is being conducted
	Location []common.CodeableConcept `json:"location,omitempty"`

	// A description and/or code explaining the premature termination of the study
	Note []Annotation `json:"note,omitempty"`

	// The party responsible for the execution of the study
	Protocol []common.Reference `json:"protocol,omitempty"`

	// A description and/or code explaining the premature termination of the study
	ReasonStopped *common.CodeableConcept `json:"reasonStopped,omitempty"`

	// Related artifacts such as publications, references, and other resources
	RelatedArtifact []interface{} `json:"relatedArtifact,omitempty"`

	// The party responsible for the execution of the study
	Sponsor *common.Reference `json:"sponsor,omitempty"`

	// The date on which the research study began
	StartDate *string `json:"startDate,omitempty"`

	// The current state of the study
	Status string `json:"status"`

	// The title of the study
	Title *string `json:"title,omitempty"`

	// The URL where the study can be found
	URL *string `json:"url,omitempty"`

	// The version of the study
	Version *string `json:"version,omitempty"`

	// The subjects that are part of the study
	Enrollment []common.Reference `json:"enrollment,omitempty"`

	// The period during which the study is conducted
	Period *common.Period `json:"period,omitempty"`

	// The phase of the study
	Phase *common.CodeableConcept `json:"phase,omitempty"`

	// The primary purpose of the study
	PrimaryPurposeType *common.CodeableConcept `json:"primaryPurposeType,omitempty"`

	// The sites where the study is being conducted
	Site []common.Reference `json:"site,omitempty"`

	// The design of the study
	StudyDesign []interface{} `json:"studyDesign,omitempty"`

	// The target population for the study
	TargetPopulation []common.Reference `json:"targetPopulation,omitempty"`
}

// RiskAssessmentPrediction represents a prediction in a risk assessment
type RiskAssessmentPrediction struct {
	common.BackboneElement

	// One of the potential outcomes for the patient (e.g. remission, death, a particular condition)
	Outcome *common.CodeableConcept `json:"outcome,omitempty"`

	// If range is used, it represents the lower and upper bounds of certainty; e.g. 40-60%
	ProbabilityDecimal *float64 `json:"probabilityDecimal,omitempty"`
	ProbabilityRange   *Range   `json:"probabilityRange,omitempty"`

	// Indicates how likely the outcome is (in the specified timeframe), expressed as a qualitative value
	QualitativeRisk *common.CodeableConcept `json:"qualitativeRisk,omitempty"`

	// Additional information explaining the basis for the prediction
	Rationale *string `json:"rationale,omitempty"`

	// Indicates the risk for this particular subject divided by the risk of the population in general
	RelativeRisk *float64 `json:"relativeRisk,omitempty"`

	// If not specified, the risk applies "over the subject's lifespan"
	WhenPeriod *common.Period `json:"whenPeriod,omitempty"`
	WhenRange  *Range         `json:"whenRange,omitempty"`
}

// RiskAssessment represents an assessment of the likely outcome(s) for a patient or other subject
type RiskAssessment struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "RiskAssessment"

	// A reference to the request that is fulfilled by this risk assessment
	BasedOn *common.Reference `json:"basedOn,omitempty"`

	// Indicates the source data considered as part of the assessment
	Basis []common.Reference `json:"basis,omitempty"`

	// The type of the risk assessment performed
	Code *common.CodeableConcept `json:"code,omitempty"`

	// For assessments or prognosis specific to a particular condition, indicates the condition being assessed
	Condition *common.Reference `json:"condition,omitempty"`

	// The encounter where the assessment was performed
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Business identifier assigned to the risk assessment
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The algorithm, process or mechanism used to evaluate the risk
	Method *common.CodeableConcept `json:"method,omitempty"`

	// A description of the steps that might be taken to reduce the identified risk(s)
	Mitigation *string `json:"mitigation,omitempty"`

	// Additional comments about the risk assessment
	Note []Annotation `json:"note,omitempty"`

	// The date (and possibly time) the risk assessment was performed
	OccurrenceDateTime *string        `json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod   *common.Period `json:"occurrencePeriod,omitempty"`

	// A reference to a resource that this risk assessment is part of, such as a Procedure
	Parent *common.Reference `json:"parent,omitempty"`

	// The provider, patient, related person, or software application that performed the assessment
	Performer *common.Reference `json:"performer,omitempty"`

	// Multiple repetitions can be used to identify the same type of outcome in different timeframes
	Prediction []RiskAssessmentPrediction `json:"prediction,omitempty"`

	// The reason the risk assessment was performed
	Reason []CodeableReference `json:"reason,omitempty"`

	// The status of the RiskAssessment, using the same statuses as an Observation
	Status string `json:"status"`

	// The patient or group the risk assessment applies to
	Subject common.Reference `json:"subject"`
}

// Schedule represents a container for slots of time that may be available for booking appointments
type Schedule struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Schedule"

	// This element is labeled as a modifier because it may be used to mark that the resource was created in error
	Active *bool `json:"active,omitempty"`

	// The capacity to support multiple referenced resource types should be used in cases where the specific resources themselves cannot be scheduled without the other
	Actor []common.Reference `json:"actor"`

	// Comments on the availability to describe any extended information
	Comment *string `json:"comment,omitempty"`

	// External Ids for this item
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// This MAY be used to describe what the schedule is for where it is clearer than just the name of the single actor
	Name *string `json:"name,omitempty"`

	// The period of time that the slots that reference this Schedule resource cover
	PlanningHorizon *common.Period `json:"planningHorizon,omitempty"`

	// A broad categorization of the service that is to be performed during this appointment
	ServiceCategory []common.CodeableConcept `json:"serviceCategory,omitempty"`

	// The specific service that is to be performed during this appointment
	ServiceType []CodeableReference `json:"serviceType,omitempty"`

	// The specialty of a practitioner that would be required to perform the service requested in this appointment
	Specialty []common.CodeableConcept `json:"specialty,omitempty"`
}

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

// Slot represents a slot of time on a schedule that may be available for booking appointments
type Slot struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Slot"

	// A slot may be allow multiple appointment types, but when booked, would only be used for one of the given appointment types
	AppointmentType []common.CodeableConcept `json:"appointmentType,omitempty"`

	// Comments on the slot to describe any extended information
	Comment *string `json:"comment,omitempty"`

	// Date/Time that the slot is to conclude
	End string `json:"end"`

	// External Ids for this item
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// This slot has already been overbooked, appointments are unlikely to be accepted for this time
	Overbooked *bool `json:"overbooked,omitempty"`

	// The schedule resource that this slot defines an interval of status information
	Schedule common.Reference `json:"schedule"`

	// A broad categorization of the service that is to be performed during this appointment
	ServiceCategory []common.CodeableConcept `json:"serviceCategory,omitempty"`

	// The type of appointments that can be booked into this slot
	ServiceType []CodeableReference `json:"serviceType,omitempty"`

	// The specialty of a practitioner that would be required to perform the service requested in this appointment
	Specialty []common.CodeableConcept `json:"specialty,omitempty"`

	// Date/Time that the slot is to begin
	Start string `json:"start"`

	// busy | free | busy-unavailable | busy-tentative | entered-in-error
	Status string `json:"status"`
}

// SpecimenFeature represents a physical feature or landmark on a specimen
type SpecimenFeature struct {
	common.BackboneElement

	// Description of the feature of the specimen
	Description string `json:"description"`

	// The landmark or feature being highlighted
	Type common.CodeableConcept `json:"type"`
}

// SpecimenCollection represents details concerning the specimen collection
type SpecimenCollection struct {
	common.BackboneElement

	// Reasons for using BodyStructure reference include: 1.) Need to identify a specific site instance
	BodySite *CodeableReference `json:"bodySite,omitempty"`

	// Time when specimen was collected from subject - the physiologically relevant time
	CollectedDateTime *string        `json:"collectedDateTime,omitempty"`
	CollectedPeriod   *common.Period `json:"collectedPeriod,omitempty"`

	// Person who collected the specimen
	Collector *common.Reference `json:"collector,omitempty"`

	// A coded value specifying the technique that is used to perform the procedure
	Device *CodeableReference `json:"device,omitempty"`

	// The span of time over which the collection of a specimen occurred
	Duration *Duration `json:"duration,omitempty"`

	// Representing fasting status using this element is preferred to representing it with an observation
	FastingStatusCodeableConcept *common.CodeableConcept `json:"fastingStatusCodeableConcept,omitempty"`
	FastingStatusDuration        *Duration               `json:"fastingStatusDuration,omitempty"`

	// A coded value specifying the technique that is used to perform the procedure
	Method *common.CodeableConcept `json:"method,omitempty"`

	// The procedure event during which the specimen was collected
	Procedure *common.Reference `json:"procedure,omitempty"`

	// The quantity of specimen collected
	Quantity *common.Quantity `json:"quantity,omitempty"`
}

// SpecimenProcessing represents details concerning processing and processing steps for the specimen
type SpecimenProcessing struct {
	common.BackboneElement

	// Material used in the processing step
	Additive []common.Reference `json:"additive,omitempty"`

	// Textual description of procedure
	Description *string `json:"description,omitempty"`

	// A coded value specifying the method used to process the specimen
	Method *common.CodeableConcept `json:"method,omitempty"`

	// A record of the time or period when the specimen processing occurred
	TimeDateTime *string        `json:"timeDateTime,omitempty"`
	TimePeriod   *common.Period `json:"timePeriod,omitempty"`
}

// SpecimenContainer represents the container holding the specimen
type SpecimenContainer struct {
	common.BackboneElement

	// The device resource for the the container holding the specimen
	Device common.Reference `json:"device"`

	// The location of the container holding the specimen
	Location *common.Reference `json:"location,omitempty"`

	// The quantity of specimen in the container
	SpecimenQuantity *common.Quantity `json:"specimenQuantity,omitempty"`
}

// Specimen represents a sample to be used for analysis
type Specimen struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Specimen"

	// The identifier assigned by the lab when accessioning specimen(s)
	AccessionIdentifier *common.Identifier `json:"accessionIdentifier,omitempty"`

	// Details concerning the specimen collection
	Collection *SpecimenCollection `json:"collection,omitempty"`

	// This element signifies if the specimen is part of a group or pooled
	Combined *string `json:"combined,omitempty"`

	// Specimen condition is an observation made about the specimen
	Condition []common.CodeableConcept `json:"condition,omitempty"`

	// The container holding the specimen
	Container []SpecimenContainer `json:"container,omitempty"`

	// A physical feature or landmark on a specimen, highlighted for context by the collector
	Feature []SpecimenFeature `json:"feature,omitempty"`

	// Id for specimen
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// To communicate any details or issues about the specimen or during the specimen collection
	Note []Annotation `json:"note,omitempty"`

	// Details concerning processing and processing steps for the specimen
	Processing []SpecimenProcessing `json:"processing,omitempty"`

	// The parent specimen
	Parent []common.Reference `json:"parent,omitempty"`

	// Time when specimen was received and processed
	ReceivedTime *string `json:"receivedTime,omitempty"`

	// The request to be performed
	Request []common.Reference `json:"request,omitempty"`

	// The physical location where the specimen is located
	Status *string `json:"status,omitempty"`

	// Where the specimen came from
	Subject *common.Reference `json:"subject,omitempty"`

	// The kind of material that forms the specimen
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// SpecimenDefinitionTypeTestedContainerAdditive represents an additive in a specimen container
type SpecimenDefinitionTypeTestedContainerAdditive struct {
	common.BackboneElement

	// Substance introduced in the kind of container to preserve, maintain or enhance the specimen
	AdditiveCodeableConcept *common.CodeableConcept `json:"additiveCodeableConcept,omitempty"`
	AdditiveReference       *common.Reference       `json:"additiveReference,omitempty"`
}

// SpecimenDefinitionTypeTestedContainer represents the container for a specimen
type SpecimenDefinitionTypeTestedContainer struct {
	common.BackboneElement

	// Substance introduced in the kind of container to preserve, maintain or enhance the specimen
	Additive []SpecimenDefinitionTypeTestedContainerAdditive `json:"additive,omitempty"`

	// Color of container cap
	Cap *common.CodeableConcept `json:"cap,omitempty"`

	// The capacity (volume or other measure) the container may contain
	Capacity *common.Quantity `json:"capacity,omitempty"`

	// The type of material of the container
	Material *common.CodeableConcept `json:"material,omitempty"`

	// The minimum volume to be conditioned in the container
	MinimumVolumeQuantity *common.Quantity `json:"minimumVolumeQuantity,omitempty"`
	MinimumVolumeString   *string          `json:"minimumVolumeString,omitempty"`

	// Special processing that should be applied to the container for this type of specimen
	Preparation *string `json:"preparation,omitempty"`

	// The type of container used to contain this kind of specimen
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// SpecimenDefinitionTypeTestedHandling represents handling instructions for a specimen
type SpecimenDefinitionTypeTestedHandling struct {
	common.BackboneElement

	// Additional textual instructions for the preservation or transport of the specimen
	Instruction *string `json:"instruction,omitempty"`

	// The maximum time interval of preservation of the specimen with these conditions
	MaxDuration *Duration `json:"maxDuration,omitempty"`

	// It qualifies the interval of temperature, which characterizes an occurrence of handling
	TemperatureQualifier *common.CodeableConcept `json:"temperatureQualifier,omitempty"`

	// The temperature interval for this set of handling instructions
	TemperatureRange *Range `json:"temperatureRange,omitempty"`
}

// SpecimenDefinitionTypeTested represents a type of specimen that can be collected
type SpecimenDefinitionTypeTested struct {
	common.BackboneElement

	// The primary preferred container for this type of specimen
	Container *SpecimenDefinitionTypeTestedContainer `json:"container,omitempty"`

	// The preferred diagnostic service for this type of specimen
	Diagnostic *common.Reference `json:"diagnostic,omitempty"`

	// The preferred handling instructions for this type of specimen
	Handling []SpecimenDefinitionTypeTestedHandling `json:"handling,omitempty"`

	// Indicates if the specimen is preferred for this type of specimen
	IsDerived *bool `json:"isDerived,omitempty"`

	// The kind of specimen conditioned for testing expected by lab
	Type *common.CodeableConcept `json:"type,omitempty"`

	// The preference for this type of conditioned specimen
	Preference string `json:"preference"`

	// The usual or preferred rejection criteria for this type of specimen
	RejectionCriterion []common.CodeableConcept `json:"rejectionCriterion,omitempty"`

	// Requirements for delivery and special handling of this kind of conditioned specimen
	Requirement *string `json:"requirement,omitempty"`

	// The time period during which specimen is accepted
	RetentionTime *Duration `json:"retentionTime,omitempty"`
}

// SpecimenDefinition represents a kind of specimen with associated set of requirements
type SpecimenDefinition struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "SpecimenDefinition"

	// The 'date' element may be more recent than the approval date because of minor changes or editorial corrections
	ApprovalDate *string `json:"approvalDate,omitempty"`

	// An individiual or organization primarily involved in the creation and maintenance of the content
	Author []ContactDetail `json:"author,omitempty"`

	// May be a web site, an email address, a telephone number, etc
	Contact []ContactDetail `json:"contact,omitempty"`

	// A copyright statement relating to the specimen definition and/or its contents
	Copyright *string `json:"copyright,omitempty"`

	// Note that this is not the same as the resource last-modified-date, since the resource may be a secondary representation of the specimen definition
	Date *string `json:"date,omitempty"`

	// This description can be used to capture details such as why the specimen definition was built
	Description *string `json:"description,omitempty"`

	// An individual or organization primarily responsible for internal coherence of the content
	Editor []ContactDetail `json:"editor,omitempty"`

	// The effective period for a specimen definition determines when the content is applicable for usage
	EffectivePeriod *common.Period `json:"effectivePeriod,omitempty"`

	// An individual or organization responsible for officially endorsing the content for use in some setting
	Endorser []ContactDetail `json:"endorser,omitempty"`

	// Allows filtering of specimen definitions that are appropriate for use versus not
	Experimental *bool `json:"experimental,omitempty"`

	// Typically, this is used for identifiers that can go in an HL7 V3 II (instance identifier) data type
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// It may be possible for the specimen definition to be used in jurisdictions other than those for which it was originally designed or intended
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`

	// If specified, this date follows the original approval date
	LastReviewDate *string `json:"lastReviewDate,omitempty"`

	// The name is not expected to be globally unique
	Name *string `json:"name,omitempty"`

	// Usually an organization but may be an individual
	Publisher *string `json:"publisher,omitempty"`

	// This element does not describe the usage of the specimen definition
	Purpose *string `json:"purpose,omitempty"`

	// Each related artifact is either an attachment, or a reference to another resource, but not both
	RelatedArtifact []interface{} `json:"relatedArtifact,omitempty"`

	// An individual or organization primarily responsible for review of some aspect of the content
	Reviewer []ContactDetail `json:"reviewer,omitempty"`

	// Allows filtering of specimen definitions that are appropriate for use versus not
	Status string `json:"status"`

	// An explanatory or alternate title for the specimen definition giving additional information about its content
	Subtitle *string `json:"subtitle,omitempty"`

	// This name does not need to be machine-processing friendly and may contain punctuation, white-space, etc
	Title *string `json:"title,omitempty"`

	// Descriptive topics related to the content of the specimen definition
	Topic []common.CodeableConcept `json:"topic,omitempty"`

	// A high-level category for the specimen definition that distinguishes the kinds of systems that would be interested in the specimen definition
	Type *common.CodeableConcept `json:"type,omitempty"`

	// Can be a urn:uuid: or a urn:oid: but real http: addresses are preferred
	URL *string `json:"url,omitempty"`

	// A detailed description of how the specimen definition is used from a clinical perspective
	Usage *string `json:"usage,omitempty"`

	// When multiple useContexts are specified, there is no expectation that all or any of the contexts apply
	UseContext []interface{} `json:"useContext,omitempty"`

	// There may be different specimen definition instances that have the same identifier but different versions
	Version *string `json:"version,omitempty"`

	// A kind of specimen with associated set of requirements
	TypeTested []SpecimenDefinitionTypeTested `json:"typeTested,omitempty"`
}

// StructureDefinitionMapping represents a mapping from one set of concepts to one or more other concepts
type StructureDefinitionMapping struct {
	common.BackboneElement

	// Comments about this mapping, including version notes, issues, scope limitations, and other important notes for usage
	Comment *string `json:"comment,omitempty"`

	// An Internal id that is used to identify this mapping within the mapping set
	Identity string `json:"identity"`

	// A name for the mapping set
	Name *string `json:"name,omitempty"`

	// An absolute URI that identifies the specification that this mapping is expressed to
	URI *string `json:"uri,omitempty"`
}

// StructureDefinitionContext represents an expression that defines where an extension can be used in resources
type StructureDefinitionContext struct {
	common.BackboneElement

	// An expression that defines where an extension can be used in resources
	Expression string `json:"expression"`

	// Defines how to interpret the context that is listed
	Type string `json:"type"`
}

// StructureDefinitionSnapshot represents a snapshot view is expressed in a standalone form that can be used and interpreted without considering the base StructureDefinition
type StructureDefinitionSnapshot struct {
	common.BackboneElement

	// Captures constraints on each element within the resource
	Element []interface{} `json:"element"`
}

// StructureDefinitionDifferential represents a differential view is expressed relative to the base StructureDefinition - a statement of differences that it applies
type StructureDefinitionDifferential struct {
	common.BackboneElement

	// Captures constraints on each element within the resource
	Element []interface{} `json:"element"`
}

// StructureDefinition represents a definition of a FHIR structure
type StructureDefinition struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "StructureDefinition"

	// Abstract Resources cannot be instantiated - a concrete sub-type must be used
	Abstract bool `json:"abstract"`

	// If differential constraints are specified in this structure, they are applied to the base in a "differential" fashion
	BaseDefinition *string `json:"baseDefinition,omitempty"`

	// May be a web site, an email address, a telephone number, etc
	Contact []ContactDetail `json:"contact,omitempty"`

	// Identifies the types of resource or data type elements to which the extension can be applied
	Context []StructureDefinitionContext `json:"context,omitempty"`

	// A differential view is expressed relative to the base StructureDefinition
	Differential *StructureDefinitionDifferential `json:"differential,omitempty"`

	// The effective period for a structure definition determines when the content is applicable for usage
	EffectivePeriod *common.Period `json:"effectivePeriod,omitempty"`

	// Allows filtering of structure definitions that are appropriate for use versus not
	Experimental *bool `json:"experimental,omitempty"`

	// Typically, this is used for identifiers that can go in an HL7 V3 II (instance identifier) data type
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// It may be possible for the structure definition to be used in jurisdictions other than those for which it was originally designed or intended
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`

	// If specified, this date follows the original approval date
	LastReviewDate *string `json:"lastReviewDate,omitempty"`

	// A mapping from one set of concepts to one or more other concepts
	Mapping []StructureDefinitionMapping `json:"mapping,omitempty"`

	// The name is not expected to be globally unique
	Name string `json:"name"`

	// Usually an organization but may be an individual
	Publisher *string `json:"publisher,omitempty"`

	// This element does not describe the usage of the structure definition
	Purpose *string `json:"purpose,omitempty"`

	// An individual or organization primarily responsible for review of some aspect of the content
	Reviewer []ContactDetail `json:"reviewer,omitempty"`

	// A snapshot view is expressed in a standalone form that can be used and interpreted without considering the base StructureDefinition
	Snapshot *StructureDefinitionSnapshot `json:"snapshot,omitempty"`

	// Allows filtering of structure definitions that are appropriate for use versus not
	Status string `json:"status"`

	// The type this structure describes
	Type string `json:"type"`

	// Can be a urn:uuid: or a urn:oid: but real http: addresses are preferred
	Url string `json:"url"`

	// A detailed description of how the structure definition is used from a clinical perspective
	Usage *string `json:"usage,omitempty"`

	// When multiple useContexts are specified, there is no expectation that all or any of the contexts apply
	UseContext []UsageContext `json:"useContext,omitempty"`

	// There may be different structure definition instances that have the same identifier but different versions
	Version *string `json:"version,omitempty"`

	// The version of the FHIR specification on which this StructureDefinition is based
	FhirVersion string `json:"fhirVersion"`

	// The kind of structure that this definition is describing
	Kind string `json:"kind"`

	// Whether this a 'primitive' type - has a value but no children
	Primitive *bool `json:"primitive,omitempty"`

	// Whether this is an extension or not
	Extension []string `json:"extension,omitempty"`

	// The name used to identify the constraint
	Key *string `json:"key,omitempty"`

	// Identifies the resource (or resources) that are being addressed by the event
	Target []string `json:"target,omitempty"`
}

// NutritionIntakeStatus represents status of nutrition intake
type NutritionIntakeStatus string

const (
	NutritionIntakeStatusPreparation  NutritionIntakeStatus = "preparation"
	NutritionIntakeStatusInProgress   NutritionIntakeStatus = "in-progress"
	NutritionIntakeStatusNotDone      NutritionIntakeStatus = "not-done"
	NutritionIntakeStatusOnHold       NutritionIntakeStatus = "on-hold"
	NutritionIntakeStatusStopped      NutritionIntakeStatus = "stopped"
	NutritionIntakeStatusCompleted    NutritionIntakeStatus = "completed"
	NutritionIntakeStatusEnteredError NutritionIntakeStatus = "entered-in-error"
	NutritionIntakeStatusUnknown      NutritionIntakeStatus = "unknown"
)

// NutritionIntakeConsumedItem represents consumed food or fluid item
type NutritionIntakeConsumedItem struct {
	common.BackboneElement

	// Quantity of the specified food
	Amount *common.Quantity `json:"amount,omitempty"`

	// Indicator when food was not consumed
	NotConsumed *bool `json:"notConsumed,omitempty"`

	// Reason the food or fluid was not consumed
	NotConsumedReason *common.CodeableConcept `json:"notConsumedReason,omitempty"`

	// The nutrition product that was consumed
	NutritionProduct CodeableReference `json:"nutritionProduct"`

	// Rate at which enteral feeding was administered
	Rate *common.Quantity `json:"rate,omitempty"`

	// Scheduled frequency of consumption
	Schedule *Timing `json:"schedule,omitempty"`

	// Category of item that was consumed
	Type common.CodeableConcept `json:"type"`
}

// NutritionIntakeIngredientLabel represents nutrient information
type NutritionIntakeIngredientLabel struct {
	common.BackboneElement

	// Total amount of nutrient consumed
	Amount common.Quantity `json:"amount"`

	// Total nutrient consumed
	Nutrient CodeableReference `json:"nutrient"`
}

// NutritionIntakePerformer represents who performed the intake
type NutritionIntakePerformer struct {
	common.BackboneElement

	// Who performed the intake
	Actor common.Reference `json:"actor"`

	// Type of performer
	Function *common.CodeableConcept `json:"function,omitempty"`
}

// NutritionIntake represents a record of food or fluid consumption
type NutritionIntake struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "NutritionIntake"

	// A plan, proposal or order that is fulfilled in whole or in part by this event
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// Overall type of nutrition intake
	Code *common.CodeableConcept `json:"code,omitempty"`

	// What food or fluid product or item was consumed
	ConsumedItem []NutritionIntakeConsumedItem `json:"consumedItem"`

	// Likely references would be to AllergyIntolerance, Observation or QuestionnaireAnswers
	DerivedFrom []common.Reference `json:"derivedFrom,omitempty"`

	// The encounter that establishes the context for this NutritionIntake
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Business identifier for this nutrition intake
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Individual item nutrients information
	IngredientLabel []NutritionIntakeIngredientLabel `json:"ingredientLabel,omitempty"`

	// Instantiates FHIR protocol or definition
	InstantiatesCanonical []string `json:"instantiatesCanonical,omitempty"`

	// Instantiates external protocol or definition
	InstantiatesUri []string `json:"instantiatesUri,omitempty"`

	// Where the intake occurred
	Location *common.Reference `json:"location,omitempty"`

	// Extra information about the Nutrition Intake
	Note []Annotation `json:"note,omitempty"`

	// When the consumption occurred
	OccurrenceDateTime *string        `json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod   *common.Period `json:"occurrencePeriod,omitempty"`

	// A larger event of which this particular event is a component or step
	PartOf []common.Reference `json:"partOf,omitempty"`

	// Who performed the intake and how they were involved
	Performer []NutritionIntakePerformer `json:"performer,omitempty"`

	// Reason for why the food or fluid is/was consumed
	Reason []CodeableReference `json:"reason,omitempty"`

	// The date when the Nutrition Intake was asserted
	Recorded *string `json:"recorded,omitempty"`

	// The person or organization that provided the information
	ReportedBoolean   *bool             `json:"reportedBoolean,omitempty"`
	ReportedReference *common.Reference `json:"reportedReference,omitempty"`

	// Status of the nutrition intake
	Status NutritionIntakeStatus `json:"status"`

	// Reason for current status
	StatusReason []common.CodeableConcept `json:"statusReason,omitempty"`

	// The person, animal or group who is/was consuming the food or fluid
	Subject common.Reference `json:"subject"`
}

// NutritionOrderIntent represents intent of nutrition order
type NutritionOrderIntent string

const (
	NutritionOrderIntentProposal      NutritionOrderIntent = "proposal"
	NutritionOrderIntentPlan          NutritionOrderIntent = "plan"
	NutritionOrderIntentDirective     NutritionOrderIntent = "directive"
	NutritionOrderIntentOrder         NutritionOrderIntent = "order"
	NutritionOrderIntentOriginalOrder NutritionOrderIntent = "original-order"
	NutritionOrderIntentReflexOrder   NutritionOrderIntent = "reflex-order"
	NutritionOrderIntentFillerOrder   NutritionOrderIntent = "filler-order"
	NutritionOrderIntentInstanceOrder NutritionOrderIntent = "instance-order"
	NutritionOrderIntentOption        NutritionOrderIntent = "option"
)

// NutritionOrderStatus represents status of nutrition order
type NutritionOrderStatus string

const (
	NutritionOrderStatusDraft        NutritionOrderStatus = "draft"
	NutritionOrderStatusActive       NutritionOrderStatus = "active"
	NutritionOrderStatusOnHold       NutritionOrderStatus = "on-hold"
	NutritionOrderStatusRevoked      NutritionOrderStatus = "revoked"
	NutritionOrderStatusCompleted    NutritionOrderStatus = "completed"
	NutritionOrderStatusEnteredError NutritionOrderStatus = "entered-in-error"
	NutritionOrderStatusUnknown      NutritionOrderStatus = "unknown"
)

// NutritionOrderPriority represents priority of nutrition order
type NutritionOrderPriority string

const (
	NutritionOrderPriorityRoutine NutritionOrderPriority = "routine"
	NutritionOrderPriorityUrgent  NutritionOrderPriority = "urgent"
	NutritionOrderPriorityASAP    NutritionOrderPriority = "asap"
	NutritionOrderPriorityStat    NutritionOrderPriority = "stat"
)

// NutritionOrder represents a request to supply nutrition
type NutritionOrder struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "NutritionOrder"

	// Information on a patient's food allergies and intolerances
	AllergyIntolerance []common.Reference `json:"allergyIntolerance,omitempty"`

	// A plan or request that is fulfilled in whole or in part by this nutrition order
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// The date and time that this nutrition order was requested
	DateTime string `json:"dateTime"`

	// An encounter that provides additional information about the healthcare context
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Feeding provided through the gastrointestinal tract via a tube, catheter, or stoma
	EnteralFormula *interface{} `json:"enteralFormula,omitempty"`

	// Information on food allergies, intolerances and preferences
	ExcludeFoodModifier []common.CodeableConcept `json:"excludeFoodModifier,omitempty"`

	// Information on food preferences
	FoodPreferenceModifier []common.CodeableConcept `json:"foodPreferenceModifier,omitempty"`

	// A shared identifier common to all nutrition orders
	GroupIdentifier *common.Identifier `json:"groupIdentifier,omitempty"`

	// Business identifier
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The URL pointing to a protocol, guideline, orderset or other definition
	Instantiates []string `json:"instantiates,omitempty"`

	// Instantiates canonical protocol or definition
	InstantiatesCanonical []string `json:"instantiatesCanonical,omitempty"`

	// Instantiates external protocol or definition
	InstantiatesUri []string `json:"instantiatesUri,omitempty"`

	// proposal | plan | directive | order | original-order | reflex-order | filler-order | instance-order | option
	Intent NutritionOrderIntent `json:"intent"`

	// Extra information about the nutrition order
	Note []Annotation `json:"note,omitempty"`

	// Diet given orally in contrast to enteral (tube) feeding
	OralDiet *interface{} `json:"oralDiet,omitempty"`

	// The practitioner that holds legal responsibility for ordering
	Orderer *common.Reference `json:"orderer,omitempty"`

	// Whether a food item is allowed to be brought in by the patient and/or family
	OutsideFoodAllowed *bool `json:"outsideFoodAllowed,omitempty"`

	// The specified desired performer of the nutrition order
	Performer []CodeableReference `json:"performer,omitempty"`

	// Indicates how quickly the Nutrition Order should be addressed
	Priority *NutritionOrderPriority `json:"priority,omitempty"`

	// draft | active | on-hold | revoked | completed | entered-in-error | unknown
	Status NutritionOrderStatus `json:"status"`

	// The person or set of individuals who needs the nutrition order
	Subject common.Reference `json:"subject"`

	// Oral nutritional products given to add further nutritional value
	Supplement []interface{} `json:"supplement,omitempty"`

	// Information to support fulfilling of the nutrition
	SupportingInformation []common.Reference `json:"supportingInformation,omitempty"`
}

// Parameters represents a set of parameters
type Parameters struct {
	Resource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Parameters"

	// A parameter passed to or received from the operation
	Parameter []ParametersParameter `json:"parameter,omitempty"`
}

// ParametersParameter represents a parameter
type ParametersParameter struct {
	common.Element

	// The name of the parameter
	Name string `json:"name"`

	// If parameter is a data type
	ValueBase64Binary        *string                 `json:"valueBase64Binary,omitempty"`
	ValueBoolean             *bool                   `json:"valueBoolean,omitempty"`
	ValueCanonical           *string                 `json:"valueCanonical,omitempty"`
	ValueCode                *string                 `json:"valueCode,omitempty"`
	ValueDate                *string                 `json:"valueDate,omitempty"`
	ValueDateTime            *string                 `json:"valueDateTime,omitempty"`
	ValueDecimal             *float64                `json:"valueDecimal,omitempty"`
	ValueId                  *string                 `json:"valueId,omitempty"`
	ValueInstant             *string                 `json:"valueInstant,omitempty"`
	ValueInteger             *int                    `json:"valueInteger,omitempty"`
	ValueInteger64           *int64                  `json:"valueInteger64,omitempty"`
	ValueMarkdown            *string                 `json:"valueMarkdown,omitempty"`
	ValueOid                 *string                 `json:"valueOid,omitempty"`
	ValuePositiveInt         *int                    `json:"valuePositiveInt,omitempty"`
	ValueString              *string                 `json:"valueString,omitempty"`
	ValueTime                *string                 `json:"valueTime,omitempty"`
	ValueUnsignedInt         *int                    `json:"valueUnsignedInt,omitempty"`
	ValueUri                 *string                 `json:"valueUri,omitempty"`
	ValueUrl                 *string                 `json:"valueUrl,omitempty"`
	ValueUuid                *string                 `json:"valueUuid,omitempty"`
	ValueAddress             *Address                `json:"valueAddress,omitempty"`
	ValueAge                 *Age                    `json:"valueAge,omitempty"`
	ValueAnnotation          *Annotation             `json:"valueAnnotation,omitempty"`
	ValueAttachment          *Attachment             `json:"valueAttachment,omitempty"`
	ValueCodeableConcept     *common.CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueCoding              *common.Coding          `json:"valueCoding,omitempty"`
	ValueContactPoint        *ContactPoint           `json:"valueContactPoint,omitempty"`
	ValueCount               *Count                  `json:"valueCount,omitempty"`
	ValueDistance            *Distance               `json:"valueDistance,omitempty"`
	ValueDuration            *Duration               `json:"valueDuration,omitempty"`
	ValueHumanName           *HumanName              `json:"valueHumanName,omitempty"`
	ValueIdentifier          *common.Identifier      `json:"valueIdentifier,omitempty"`
	ValueMoney               *Money                  `json:"valueMoney,omitempty"`
	ValuePeriod              *common.Period          `json:"valuePeriod,omitempty"`
	ValueQuantity            *common.Quantity        `json:"valueQuantity,omitempty"`
	ValueRange               *Range                  `json:"valueRange,omitempty"`
	ValueRatio               *Ratio                  `json:"valueRatio,omitempty"`
	ValueReference           *common.Reference       `json:"valueReference,omitempty"`
	ValueSampledData         *interface{}            `json:"valueSampledData,omitempty"`
	ValueSignature           *Signature              `json:"valueSignature,omitempty"`
	ValueTiming              *Timing                 `json:"valueTiming,omitempty"`
	ValueContactDetail       *ContactDetail          `json:"valueContactDetail,omitempty"`
	ValueContributor         *interface{}            `json:"valueContributor,omitempty"`
	ValueDataRequirement     *interface{}            `json:"valueDataRequirement,omitempty"`
	ValueExpression          *Expression             `json:"valueExpression,omitempty"`
	ValueParameterDefinition *interface{}            `json:"valueParameterDefinition,omitempty"`
	ValueRelatedArtifact     *RelatedArtifact        `json:"valueRelatedArtifact,omitempty"`
	ValueTriggerDefinition   *TriggerDefinition      `json:"valueTriggerDefinition,omitempty"`
	ValueUsageContext        *UsageContext           `json:"valueUsageContext,omitempty"`
	ValueDosage              *Dosage                 `json:"valueDosage,omitempty"`
	ValueMeta                *Meta                   `json:"valueMeta,omitempty"`

	// If parameter is a whole resource
	Resource interface{} `json:"resource,omitempty"`

	// Named part of a multi-part parameter
	Part []ParametersParameter `json:"part,omitempty"`
}
