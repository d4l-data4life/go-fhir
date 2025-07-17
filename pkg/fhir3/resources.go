// Package fhir3 contains FHIR R3 (version 3.0.2) resource definitions
package fhir3

import (
	"time"

	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// Helper functions for creating pointers to basic types
func StringPtr(s string) *string     { return &s }
func IntPtr(i int) *int              { return &i }
func BoolPtr(b bool) *bool           { return &b }
func TimePtr(t time.Time) *time.Time { return &t }

// Base types for FHIR R3

// Resource is the base resource type for everything
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

// DomainResource represents a resource that includes narrative, extensions, and contained resources
type DomainResource struct {
	Resource

	// Text summary of the resource, for human interpretation
	Text *Narrative `json:"text,omitempty"`

	// Contained, inline Resources
	Contained []interface{} `json:"contained,omitempty"`

	// Additional content defined by implementations
	Extension []common.Extension `json:"extension,omitempty"`

	// Extensions that cannot be ignored
	ModifierExtension []common.Extension `json:"modifierExtension,omitempty"`
}

// Meta provides metadata about a resource
type Meta struct {
	common.Element

	// Version specific identifier
	VersionId *string `json:"versionId,omitempty"`

	// When the resource version last changed
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`

	// Profiles this resource claims to conform to
	Profile []string `json:"profile,omitempty"`

	// Security Labels applied to this resource
	Security []common.Coding `json:"security,omitempty"`

	// Tags applied to this resource
	Tag []common.Coding `json:"tag,omitempty"`
}

// HumanName represents a human name, with the ability to identify parts and usage
type HumanName struct {
	common.Element

	// usual | official | temp | nickname | anonymous | old | maiden
	Use *HumanNameUse `json:"use,omitempty"`

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

// HumanNameUse represents the use of a human name
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

// Patient represents information about an individual or animal receiving care or other health-related services
type Patient struct {
	DomainResource

	// Whether this patient's record is in active use
	Active *bool `json:"active,omitempty"`

	// Addresses for the individual
	Address []Address `json:"address,omitempty"`

	// The date of birth for the individual
	BirthDate *string `json:"birthDate,omitempty"`

	// A contact party (e.g. guardian, partner, friend) for the patient
	Contact []PatientContact `json:"contact,omitempty"`

	// Indicates if the individual is deceased or not
	DeceasedBoolean  *bool      `json:"deceasedBoolean,omitempty"`
	DeceasedDateTime *time.Time `json:"deceasedDateTime,omitempty"`

	// male | female | other | unknown
	Gender *PatientGender `json:"gender,omitempty"`

	// Identifier for this individual
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Link to another patient resource that concerns the same actual person
	Link []PatientLink `json:"link,omitempty"`

	// Organization that is the custodian of the patient record
	ManagingOrganization *common.Reference `json:"managingOrganization,omitempty"`

	// Marital (civil) status of a patient
	MaritalStatus *common.CodeableConcept `json:"maritalStatus,omitempty"`

	// Whether patient is part of a multiple birth
	MultipleBirthBoolean *bool `json:"multipleBirthBoolean,omitempty"`
	MultipleBirthInteger *int  `json:"multipleBirthInteger,omitempty"`

	// A name associated with the patient
	Name []HumanName `json:"name,omitempty"`

	// Image of the patient
	Photo []Attachment `json:"photo,omitempty"`

	// A contact detail for the individual
	Telecom []ContactPoint `json:"telecom,omitempty"`
}

// PatientContact represents a contact party for the patient
type PatientContact struct {
	common.BackboneElement

	// Address for the contact person
	Address *Address `json:"address,omitempty"`

	// male | female | other | unknown
	Gender *PatientGender `json:"gender,omitempty"`

	// A name associated with the contact person
	Name *HumanName `json:"name,omitempty"`

	// The kind of relationship
	Relationship []common.CodeableConcept `json:"relationship,omitempty"`

	// A contact detail for the person
	Telecom []ContactPoint `json:"telecom,omitempty"`
}

// PatientLink represents a link to another patient resource
type PatientLink struct {
	common.BackboneElement

	// The other patient resource that the link refers to
	Other *common.Reference `json:"other"`

	// replaced-by | replaces | refer | seealso
	Type PatientLinkType `json:"type"`
}

// PatientGender represents the gender of a patient
type PatientGender string

const (
	PatientGenderMale    PatientGender = "male"
	PatientGenderFemale  PatientGender = "female"
	PatientGenderOther   PatientGender = "other"
	PatientGenderUnknown PatientGender = "unknown"
)

// PatientLinkType represents the type of link between patient resources
type PatientLinkType string

const (
	PatientLinkTypeReplacedBy PatientLinkType = "replaced-by"
	PatientLinkTypeReplaces   PatientLinkType = "replaces"
	PatientLinkTypeRefer      PatientLinkType = "refer"
	PatientLinkTypeSeeAlso    PatientLinkType = "seealso"
)

// Bundle represents a container for a collection of resources (R3 version)
type Bundle struct {
	Resource

	// An entry in a bundle resource - will either contain a resource, or information about a resource
	Entry []BundleEntry `json:"entry,omitempty"`

	// Persistent identity generally only matters for batches of type Document, Message, and Collection
	Identifier *common.Identifier `json:"identifier,omitempty"`

	// A series of links that provide context to this bundle
	Link []BundleLink `json:"link,omitempty"`

	// Digital Signature - R3 version
	Signature *Signature `json:"signature,omitempty"`

	// Only used if the bundle is a search result set
	Total *int `json:"total,omitempty"`

	// document | message | transaction | transaction-response | batch | batch-response | history | searchset | collection
	Type BundleType `json:"type"`
}

// BundleEntry represents an entry in a bundle resource (R3 version)
type BundleEntry struct {
	common.BackboneElement

	// URI for resource (Absolute or relative)
	FullURL *string `json:"fullUrl,omitempty"`

	// Links related to this entry
	Link []BundleLink `json:"link,omitempty"`

	// Additional execution information (transaction/batch/history)
	Request *BundleEntryRequest `json:"request,omitempty"`

	// The Resource for the entry
	Resource interface{} `json:"resource,omitempty"`

	// Transaction Related Information
	Response *BundleEntryResponse `json:"response,omitempty"`

	// Search related information
	Search *BundleEntrySearch `json:"search,omitempty"`
}

// BundleLink represents a series of links that provide context to this bundle
type BundleLink struct {
	common.BackboneElement

	// See http://www.iana.org/assignments/link-relations/link-relations.xhtml#link-relations-1
	Relation string `json:"relation"`

	// Reference details for the link
	URL string `json:"url"`
}

// BundleEntryRequest represents additional execution information
type BundleEntryRequest struct {
	common.BackboneElement

	// For managing update contention
	IfMatch *string `json:"ifMatch,omitempty"`

	// For managing update contention
	IfModifiedSince *time.Time `json:"ifModifiedSince,omitempty"`

	// For conditional creates
	IfNoneExist *string `json:"ifNoneExist,omitempty"`

	// For conditional read
	IfNoneMatch *string `json:"ifNoneMatch,omitempty"`

	// GET | POST | PUT | DELETE
	Method BundleEntryRequestMethod `json:"method"`

	// URL for HTTP equivalent of this entry
	URL string `json:"url"`
}

// BundleEntryResponse represents transaction related information
type BundleEntryResponse struct {
	common.BackboneElement

	// The Etag for the resource
	Etag *string `json:"etag,omitempty"`

	// Server's date time modified
	LastModified *time.Time `json:"lastModified,omitempty"`

	// The location (if the operation returns a location)
	Location *string `json:"location,omitempty"`

	// OperationOutcome with hints and warnings
	Outcome interface{} `json:"outcome,omitempty"`

	// Status response code
	Status string `json:"status"`
}

// BundleEntrySearch represents search related information
type BundleEntrySearch struct {
	common.BackboneElement

	// match | include | outcome - why this is in the result set
	Mode *BundleEntrySearchMode `json:"mode,omitempty"`

	// Search ranking
	Score *float64 `json:"score,omitempty"`
}

// Bundle-related enums
type BundleType string

const (
	BundleTypeDocument            BundleType = "document"
	BundleTypeMessage             BundleType = "message"
	BundleTypeTransaction         BundleType = "transaction"
	BundleTypeTransactionResponse BundleType = "transaction-response"
	BundleTypeBatch               BundleType = "batch"
	BundleTypeBatchResponse       BundleType = "batch-response"
	BundleTypeHistory             BundleType = "history"
	BundleTypeSearchset           BundleType = "searchset"
	BundleTypeCollection          BundleType = "collection"
)

type BundleEntryRequestMethod string

const (
	BundleEntryRequestMethodGET    BundleEntryRequestMethod = "GET"
	BundleEntryRequestMethodPOST   BundleEntryRequestMethod = "POST"
	BundleEntryRequestMethodPUT    BundleEntryRequestMethod = "PUT"
	BundleEntryRequestMethodDELETE BundleEntryRequestMethod = "DELETE"
)

type BundleEntrySearchMode string

const (
	BundleEntrySearchModeMatch   BundleEntrySearchMode = "match"
	BundleEntrySearchModeInclude BundleEntrySearchMode = "include"
	BundleEntrySearchModeOutcome BundleEntrySearchMode = "outcome"
)

// Organization represents a formally or informally recognized grouping of people or organizations
type Organization struct {
	DomainResource

	// Whether the organization's record is still in active use
	Active *bool `json:"active,omitempty"`

	// Addresses for the organization
	Address []Address `json:"address,omitempty"`

	// A list of contact details for the organization
	Contact []OrganizationContact `json:"contact,omitempty"`

	// Technical endpoints providing access to services operated for the organization
	Endpoint []common.Reference `json:"endpoint,omitempty"`

	// Identifies this organization across multiple systems
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Name used for the organization
	Name *string `json:"name,omitempty"`

	// The organization of which this organization forms a part
	PartOf *common.Reference `json:"partOf,omitempty"`

	// A contact detail for the organization
	Telecom []ContactPoint `json:"telecom,omitempty"`

	// Kind of organization
	Type []common.CodeableConcept `json:"type,omitempty"`
}

// OrganizationContact represents contact information for the organization
type OrganizationContact struct {
	common.BackboneElement

	// Visiting or postal addresses for the contact
	Address *Address `json:"address,omitempty"`

	// A name associated with the contact
	Name *HumanName `json:"name,omitempty"`

	// The type of contact
	Purpose *common.CodeableConcept `json:"purpose,omitempty"`

	// Contact details (telephone, email, etc.)
	Telecom []ContactPoint `json:"telecom,omitempty"`
}

// Practitioner represents a person who is directly or indirectly involved in the provisioning of healthcare
type Practitioner struct {
	DomainResource

	// Whether this practitioner's record is in active use
	Active *bool `json:"active,omitempty"`

	// Addresses where the practitioner can be found or visited or to which mail can be delivered
	Address []Address `json:"address,omitempty"`

	// The date of birth for the practitioner
	BirthDate *string `json:"birthDate,omitempty"`

	// A language the practitioner can use in patient communication
	Communication []common.CodeableConcept `json:"communication,omitempty"`

	// male | female | other | unknown
	Gender *PractitionerGender `json:"gender,omitempty"`

	// An identifier for this practitioner
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The name of the practitioner
	Name []HumanName `json:"name,omitempty"`

	// Image of the person
	Photo []Attachment `json:"photo,omitempty"`

	// Qualifications obtained by training and certification
	Qualification []PractitionerQualification `json:"qualification,omitempty"`

	// A contact detail for the practitioner
	Telecom []ContactPoint `json:"telecom,omitempty"`
}

// PractitionerQualification represents qualifications obtained by training and certification
type PractitionerQualification struct {
	common.BackboneElement

	// Coded representation of the qualification
	Code *common.CodeableConcept `json:"code"`

	// An identifier for this qualification for the practitioner
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Organization that regulates and issues the qualification
	Issuer *common.Reference `json:"issuer,omitempty"`

	// Period during which the qualification is valid
	Period *common.Period `json:"period,omitempty"`
}

// PractitionerGender represents the gender of a practitioner
type PractitionerGender string

const (
	PractitionerGenderMale    PractitionerGender = "male"
	PractitionerGenderFemale  PractitionerGender = "female"
	PractitionerGenderOther   PractitionerGender = "other"
	PractitionerGenderUnknown PractitionerGender = "unknown"
)

// Encounter represents an interaction between a patient and healthcare provider(s)
type Encounter struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Encounter"

	// The billing system may choose to allocate billable items associated with the Encounter
	Account []common.Reference `json:"account,omitempty"`

	// The appointment that scheduled this encounter (single reference in R3)
	Appointment *common.Reference `json:"appointment,omitempty"`

	// inpatient | outpatient | ambulatory | emergency + (optional in R3)
	Class *common.Coding `json:"class,omitempty"`

	// The class history permits the tracking of the encounters transitions
	ClassHistory []EncounterClassHistory `json:"classHistory,omitempty"`

	// The list of diagnosis relevant to this encounter
	Diagnosis []EncounterDiagnosis `json:"diagnosis,omitempty"`

	// Where a specific encounter should be classified as a part of a specific episode(s) of care
	EpisodeOfCare []common.Reference `json:"episodeOfCare,omitempty"`

	// An Encounter may cover more than just the inpatient stay
	Hospitalization *EncounterHospitalization `json:"hospitalization,omitempty"`

	// Identifier(s) by which this encounter is known
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The referral request this encounter satisfies (incoming referral)
	IncomingReferral []common.Reference `json:"incomingReferral,omitempty"`

	// May differ from the time the Encounter.period lasted because of leave of absence
	Length *Duration `json:"length,omitempty"`

	// Virtual encounters can be recorded in the Encounter by specifying a location reference
	Location []EncounterLocation `json:"location,omitempty"`

	// The list of people responsible for providing the service
	Participant []EncounterParticipant `json:"participant,omitempty"`

	// This is also used for associating a child's encounter back to the mother's encounter
	PartOf *common.Reference `json:"partOf,omitempty"`

	// If not (yet) known, the end of the Period may be omitted
	Period *common.Period `json:"period,omitempty"`

	// Indicates the urgency of the encounter
	Priority *common.CodeableConcept `json:"priority,omitempty"`

	// For systems that need to know which was the primary diagnosis (simple array in R3)
	Reason []common.CodeableConcept `json:"reason,omitempty"`

	// The organization that is primarily responsible for this Encounter's services
	ServiceProvider *common.Reference `json:"serviceProvider,omitempty"`

	// Note that internal business rules will determine the appropriate transitions
	Status EncounterStatus `json:"status"`

	// The current status is always found in the current version of the resource
	StatusHistory []EncounterStatusHistory `json:"statusHistory,omitempty"`

	// While the encounter is always about the patient, the patient might not actually be known
	Subject *common.Reference `json:"subject,omitempty"`

	// Since there are many ways to further classify encounters, this element is 0..*
	Type []common.CodeableConcept `json:"type,omitempty"`
}

// EncounterClassHistory represents the class history permitting tracking of encounter transitions
type EncounterClassHistory struct {
	common.BackboneElement

	// inpatient | outpatient | ambulatory | emergency +
	Class common.Coding `json:"class"`

	// The time that the episode was in the specified class
	Period common.Period `json:"period"`
}

// EncounterDiagnosis represents diagnoses relevant to this encounter
type EncounterDiagnosis struct {
	common.BackboneElement

	// For systems that need to know which was the primary diagnosis
	Condition common.Reference `json:"condition"`

	// Ranking of the diagnosis (for each role type)
	Rank *int `json:"rank,omitempty"`

	// Role that this diagnosis has within the encounter (e.g. admission, billing, discharge) - Note: "role" in R3, "use" in R4+
	Role *common.CodeableConcept `json:"role,omitempty"`
}

// EncounterHospitalization represents hospitalization details
type EncounterHospitalization struct {
	common.BackboneElement

	// From where patient was admitted (physician referral, transfer)
	AdmitSource *common.CodeableConcept `json:"admitSource,omitempty"`

	// Location to which the patient is discharged
	Destination *common.Reference `json:"destination,omitempty"`

	// For example a patient may request both a dairy-free and nut-free diet preference
	DietPreference []common.CodeableConcept `json:"dietPreference,omitempty"`

	// Category or kind of location after discharge
	DischargeDisposition *common.CodeableConcept `json:"dischargeDisposition,omitempty"`

	// The location from which the patient came before admission
	Origin *common.Reference `json:"origin,omitempty"`

	// Pre-admission identifier
	PreAdmissionIdentifier *common.Identifier `json:"preAdmissionIdentifier,omitempty"`

	// Whether this hospitalization is a readmission and why if known
	ReAdmission *common.CodeableConcept `json:"reAdmission,omitempty"`

	// Any special requests that have been made for this hospitalization encounter
	SpecialArrangement []common.CodeableConcept `json:"specialArrangement,omitempty"`

	// Special courtesies (VIP, board member)
	SpecialCourtesy []common.CodeableConcept `json:"specialCourtesy,omitempty"`
}

// EncounterLocation represents locations where the patient has been during this encounter
type EncounterLocation struct {
	common.BackboneElement

	// The location where the encounter takes place
	Location common.Reference `json:"location"`

	// Time period during which the patient was present at the location
	Period *common.Period `json:"period,omitempty"`

	// When the patient is no longer active at a location
	Status *EncounterLocationStatus `json:"status,omitempty"`
}

// Duration represents a length of time (R3 version)

// EncounterStatus represents the status of an encounter
type EncounterStatus string

const (
	EncounterStatusPlanned        EncounterStatus = "planned"
	EncounterStatusArrived        EncounterStatus = "arrived"
	EncounterStatusTriaged        EncounterStatus = "triaged"
	EncounterStatusInProgress     EncounterStatus = "in-progress"
	EncounterStatusOnLeave        EncounterStatus = "onleave"
	EncounterStatusFinished       EncounterStatus = "finished"
	EncounterStatusCancelled      EncounterStatus = "cancelled"
	EncounterStatusEnteredInError EncounterStatus = "entered-in-error"
	EncounterStatusUnknown        EncounterStatus = "unknown"
)

// EncounterLocationStatus represents the status of a location during an encounter
type EncounterLocationStatus string

const (
	EncounterLocationStatusPlanned   EncounterLocationStatus = "planned"
	EncounterLocationStatusActive    EncounterLocationStatus = "active"
	EncounterLocationStatusReserved  EncounterLocationStatus = "reserved"
	EncounterLocationStatusCompleted EncounterLocationStatus = "completed"
)

// EncounterParticipant represents people involved in the encounter
type EncounterParticipant struct {
	common.BackboneElement

	// Persons involved in the encounter other than the patient
	Individual *common.Reference `json:"individual,omitempty"`

	// The period of time that the specified participant participated in the encounter
	Period *common.Period `json:"period,omitempty"`

	// The participant type indicates how an individual participates in an encounter
	Type []common.CodeableConcept `json:"type,omitempty"`
}

// EncounterStatusHistory represents the status history of an encounter
type EncounterStatusHistory struct {
	common.BackboneElement

	// The time that the episode was in the specified status
	Period common.Period `json:"period"`

	// planned | arrived | triaged | in-progress | onleave | finished | cancelled | entered-in-error | unknown
	Status EncounterStatus `json:"status"`
}

// Condition represents a clinical condition, problem, diagnosis, or other event, situation, issue, or clinical concept
type Condition struct {
	DomainResource

	// Estimated or actual date or date-time the condition was resolved
	AbatementDateTime *time.Time     `json:"abatementDateTime,omitempty"`
	AbatementAge      *Age           `json:"abatementAge,omitempty"`
	AbatementPeriod   *common.Period `json:"abatementPeriod,omitempty"`
	AbatementRange    *Range         `json:"abatementRange,omitempty"`
	AbatementString   *string        `json:"abatementString,omitempty"`

	// Person who asserts this condition
	Asserter *common.Reference `json:"asserter,omitempty"`

	// Anatomical location, if relevant
	BodySite []common.CodeableConcept `json:"bodySite,omitempty"`

	// The category of the condition
	Category []common.CodeableConcept `json:"category,omitempty"`

	// active | recurrence | relapse | inactive | remission | resolved
	ClinicalStatus *ConditionClinicalStatus `json:"clinicalStatus,omitempty"`

	// Identification of the condition, problem or diagnosis
	Code *common.CodeableConcept `json:"code,omitempty"`

	// Encounter or episode when condition first asserted
	Context *common.Reference `json:"context,omitempty"`

	// Supporting evidence
	Evidence []ConditionEvidence `json:"evidence,omitempty"`

	// External IDs for this condition
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Additional information about the Condition
	Note []Annotation `json:"note,omitempty"`

	// Date record was believed accurate
	OnsetDateTime *time.Time     `json:"onsetDateTime,omitempty"`
	OnsetAge      *Age           `json:"onsetAge,omitempty"`
	OnsetPeriod   *common.Period `json:"onsetPeriod,omitempty"`
	OnsetRange    *Range         `json:"onsetRange,omitempty"`
	OnsetString   *string        `json:"onsetString,omitempty"`

	// A subjective assessment of the severity of the condition
	Severity *common.CodeableConcept `json:"severity,omitempty"`

	// Stage/grade, usually assessed formally
	Stage *ConditionStage `json:"stage,omitempty"`

	// Who has the condition?
	Subject *common.Reference `json:"subject"`

	// provisional | differential | confirmed | refuted | entered-in-error | unknown
	VerificationStatus *ConditionVerificationStatus `json:"verificationStatus,omitempty"`
}

// ConditionEvidence represents supporting evidence
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
	Assessment []common.Reference `json:"assessment,omitempty"`

	// Kind of staging
	Summary *common.CodeableConcept `json:"summary,omitempty"`
}

// Condition-related enums
type ConditionClinicalStatus string

const (
	ConditionClinicalStatusActive     ConditionClinicalStatus = "active"
	ConditionClinicalStatusRecurrence ConditionClinicalStatus = "recurrence"
	ConditionClinicalStatusRelapse    ConditionClinicalStatus = "relapse"
	ConditionClinicalStatusInactive   ConditionClinicalStatus = "inactive"
	ConditionClinicalStatusRemission  ConditionClinicalStatus = "remission"
	ConditionClinicalStatusResolved   ConditionClinicalStatus = "resolved"
)

type ConditionVerificationStatus string

const (
	ConditionVerificationStatusProvisional    ConditionVerificationStatus = "provisional"
	ConditionVerificationStatusDifferential   ConditionVerificationStatus = "differential"
	ConditionVerificationStatusConfirmed      ConditionVerificationStatus = "confirmed"
	ConditionVerificationStatusRefuted        ConditionVerificationStatus = "refuted"
	ConditionVerificationStatusEnteredInError ConditionVerificationStatus = "entered-in-error"
	ConditionVerificationStatusUnknown        ConditionVerificationStatus = "unknown"
)

// Procedure represents an action that is or was performed on a patient
type Procedure struct {
	DomainResource

	// Target body sites
	BodySite []common.CodeableConcept `json:"bodySite,omitempty"`

	// A request for this procedure
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// Classification of the procedure
	Category *common.CodeableConcept `json:"category,omitempty"`

	// Identification of the procedure
	Code *common.CodeableConcept `json:"code,omitempty"`

	// Complication following the procedure
	Complication []common.CodeableConcept `json:"complication,omitempty"`

	// A condition that is a result of the procedure
	ComplicationDetail []common.Reference `json:"complicationDetail,omitempty"`

	// Encounter or episode associated with the procedure
	Context *common.Reference `json:"context,omitempty"`

	// Instantiates FHIR protocol or definition
	Definition []common.Reference `json:"definition,omitempty"`

	// Device changed in procedure
	FocalDevice []ProcedureFocalDevice `json:"focalDevice,omitempty"`

	// Instructions for follow up
	FollowUp []common.CodeableConcept `json:"followUp,omitempty"`

	// External Identifiers for this procedure
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Where the procedure happened
	Location *common.Reference `json:"location,omitempty"`

	// Additional information about the procedure
	Note []Annotation `json:"note,omitempty"`

	// The result of procedure
	Outcome *common.CodeableConcept `json:"outcome,omitempty"`

	// Part of referenced event
	PartOf []common.Reference `json:"partOf,omitempty"`

	// Date/Period the procedure was performed
	PerformedDateTime *time.Time     `json:"performedDateTime,omitempty"`
	PerformedPeriod   *common.Period `json:"performedPeriod,omitempty"`

	// The people who performed the procedure
	Performer []ProcedurePerformer `json:"performer,omitempty"`

	// Coded reason procedure performed
	ReasonCode []common.CodeableConcept `json:"reasonCode,omitempty"`

	// Condition that justifies procedure
	ReasonReference []common.Reference `json:"reasonReference,omitempty"`

	// Any report resulting from the procedure
	Report []common.Reference `json:"report,omitempty"`

	// preparation | in-progress | suspended | aborted | completed | entered-in-error | unknown
	Status ProcedureStatus `json:"status"`

	// Who the procedure was performed on
	Subject *common.Reference `json:"subject"`

	// Coded items used during the procedure
	UsedCode []common.CodeableConcept `json:"usedCode,omitempty"`

	// Items used during the procedure
	UsedReference []common.Reference `json:"usedReference,omitempty"`
}

// ProcedureFocalDevice represents device changed in procedure
type ProcedureFocalDevice struct {
	common.BackboneElement

	// Kind of change to device
	Action *common.CodeableConcept `json:"action,omitempty"`

	// Device that was changed
	Manipulated *common.Reference `json:"manipulated"`
}

// ProcedurePerformer represents the people who performed the procedure
type ProcedurePerformer struct {
	common.BackboneElement

	// The reference to the practitioner
	Actor *common.Reference `json:"actor"`

	// Organization the device or practitioner was acting for
	OnBehalfOf *common.Reference `json:"onBehalfOf,omitempty"`

	// The role the actor was in
	Role *common.CodeableConcept `json:"role,omitempty"`
}

// Procedure-related enums
type ProcedureStatus string

const (
	ProcedureStatusPreparation    ProcedureStatus = "preparation"
	ProcedureStatusInProgress     ProcedureStatus = "in-progress"
	ProcedureStatusSuspended      ProcedureStatus = "suspended"
	ProcedureStatusAborted        ProcedureStatus = "aborted"
	ProcedureStatusCompleted      ProcedureStatus = "completed"
	ProcedureStatusEnteredInError ProcedureStatus = "entered-in-error"
	ProcedureStatusUnknown        ProcedureStatus = "unknown"
)

// Location represents details and position information for a physical place
type Location struct {
	DomainResource

	// Physical location
	Address *Address `json:"address,omitempty"`

	// Description of the Location
	Description *string `json:"description,omitempty"`

	// Technical endpoints providing access to services operated for the location
	Endpoint []common.Reference `json:"endpoint,omitempty"`

	// Unique code or number identifying the location to its users
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Organization responsible for provisioning and upkeep
	ManagingOrganization *common.Reference `json:"managingOrganization,omitempty"`

	// instance | kind
	Mode *LocationMode `json:"mode,omitempty"`

	// Name of the location as used by humans
	Name *string `json:"name,omitempty"`

	// What days/times during a week is this location usually open
	OperationalStatus *common.Coding `json:"operationalStatus,omitempty"`

	// Another Location this one is physically part of
	PartOf *common.Reference `json:"partOf,omitempty"`

	// Physical form of the location
	PhysicalType *common.CodeableConcept `json:"physicalType,omitempty"`

	// The absolute geographic location
	Position *LocationPosition `json:"position,omitempty"`

	// active | suspended | inactive
	Status *LocationStatus `json:"status,omitempty"`

	// Contact details of the location
	Telecom []ContactPoint `json:"telecom,omitempty"`

	// Type of function performed
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// LocationPosition represents the absolute geographic location
type LocationPosition struct {
	common.BackboneElement

	// Altitude with WGS84 datum
	Altitude *float64 `json:"altitude,omitempty"`

	// Latitude with WGS84 datum
	Latitude float64 `json:"latitude"`

	// Longitude with WGS84 datum
	Longitude float64 `json:"longitude"`
}

// Location-related enums
type LocationMode string

const (
	LocationModeInstance LocationMode = "instance"
	LocationModeKind     LocationMode = "kind"
)

type LocationStatus string

const (
	LocationStatusActive    LocationStatus = "active"
	LocationStatusSuspended LocationStatus = "suspended"
	LocationStatusInactive  LocationStatus = "inactive"
)

// DiagnosticReport represents the findings and interpretation of diagnostic tests
type DiagnosticReport struct {
	DomainResource

	// Request details that this diagnostic report is based on
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// Service category
	Category *common.CodeableConcept `json:"category,omitempty"`

	// Name/Code for this diagnostic report
	Code *common.CodeableConcept `json:"code"`

	// Health care event when test ordered
	Context *common.Reference `json:"context,omitempty"`

	// Clinically relevant time/time-period for report
	EffectiveDateTime *time.Time     `json:"effectiveDateTime,omitempty"`
	EffectivePeriod   *common.Period `json:"effectivePeriod,omitempty"`

	// Business identifier for report
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Reference to full details of imaging associated with the diagnostic report
	Image []DiagnosticReportImage `json:"image,omitempty"`

	// Key images associated with this report
	ImagingStudy []common.Reference `json:"imagingStudy,omitempty"`

	// DateTime this version was released
	Issued *time.Time `json:"issued,omitempty"`

	// Responsible Diagnostic Service
	Performer []DiagnosticReportPerformer `json:"performer,omitempty"`

	// Entire report as issued
	PresentedForm []Attachment `json:"presentedForm,omitempty"`

	// Observations
	Result []common.Reference `json:"result,omitempty"`

	// Specimens this report is based on
	Specimen []common.Reference `json:"specimen,omitempty"`

	// registered | partial | preliminary | final | amended | corrected | appended | cancelled | entered-in-error | unknown
	Status DiagnosticReportStatus `json:"status"`

	// The subject of the report - usually, but not always, the patient
	Subject *common.Reference `json:"subject,omitempty"`
}

// DiagnosticReportImage represents key images associated with this report
type DiagnosticReportImage struct {
	common.BackboneElement

	// Comment about the image
	Comment *string `json:"comment,omitempty"`

	// Reference to the image source
	Link *common.Reference `json:"link"`
}

// DiagnosticReportPerformer represents responsible diagnostic service
type DiagnosticReportPerformer struct {
	common.BackboneElement

	// Practitioner or Organization participant
	Actor *common.Reference `json:"actor"`

	// Type of performer
	Role *common.CodeableConcept `json:"role,omitempty"`
}

// DiagnosticReport-related enums
type DiagnosticReportStatus string

const (
	DiagnosticReportStatusRegistered     DiagnosticReportStatus = "registered"
	DiagnosticReportStatusPartial        DiagnosticReportStatus = "partial"
	DiagnosticReportStatusPreliminary    DiagnosticReportStatus = "preliminary"
	DiagnosticReportStatusFinal          DiagnosticReportStatus = "final"
	DiagnosticReportStatusAmended        DiagnosticReportStatus = "amended"
	DiagnosticReportStatusCorrected      DiagnosticReportStatus = "corrected"
	DiagnosticReportStatusAppended       DiagnosticReportStatus = "appended"
	DiagnosticReportStatusCancelled      DiagnosticReportStatus = "cancelled"
	DiagnosticReportStatusEnteredInError DiagnosticReportStatus = "entered-in-error"
	DiagnosticReportStatusUnknown        DiagnosticReportStatus = "unknown"
)

// Observation represents measurements and simple assertions made about a patient, device or other subject
type Observation struct {
	DomainResource

	// Request that the observation is fulfilling
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// Observed body part
	BodySite *common.CodeableConcept `json:"bodySite,omitempty"`

	// Classification of type of observation
	Category []common.CodeableConcept `json:"category,omitempty"`

	// Type of observation (code / type)
	Code *common.CodeableConcept `json:"code"`

	// Why the result is missing
	DataAbsentReason *common.CodeableConcept `json:"dataAbsentReason,omitempty"`

	// (Measurement) Device
	Device *common.Reference `json:"device,omitempty"`

	// Clinically relevant time/time-period for observation
	EffectiveDateTime *time.Time     `json:"effectiveDateTime,omitempty"`
	EffectivePeriod   *common.Period `json:"effectivePeriod,omitempty"`

	// Encounter or episode related to observation
	Context *common.Reference `json:"context,omitempty"`

	// Business Identifier for observation
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// High, low, normal, etc.
	Interpretation *common.CodeableConcept `json:"interpretation,omitempty"`

	// Date/Time this was made available
	Issued *time.Time `json:"issued,omitempty"`

	// The observation method
	Method *common.CodeableConcept `json:"method,omitempty"`

	// Comments about result
	Note []Annotation `json:"note,omitempty"`

	// Who is responsible for the observation
	Performer []common.Reference `json:"performer,omitempty"`

	// Provides guide for interpretation
	ReferenceRange []ObservationReferenceRange `json:"referenceRange,omitempty"`

	// Resource related to this observation
	Related []ObservationRelated `json:"related,omitempty"`

	// Specimen used for this observation
	Specimen *common.Reference `json:"specimen,omitempty"`

	// registered | preliminary | final | amended | corrected | cancelled | entered-in-error | unknown
	Status ObservationStatus `json:"status"`

	// Who and/or what this is about
	Subject *common.Reference `json:"subject,omitempty"`

	// Actual result
	ValueQuantity        *common.Quantity        `json:"valueQuantity,omitempty"`
	ValueCodeableConcept *common.CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueString          *string                 `json:"valueString,omitempty"`
	ValueBoolean         *bool                   `json:"valueBoolean,omitempty"`
	ValueRange           *Range                  `json:"valueRange,omitempty"`
	ValueRatio           *Ratio                  `json:"valueRatio,omitempty"`
	ValueSampledData     *SampledData            `json:"valueSampledData,omitempty"`
	ValueAttachment      *Attachment             `json:"valueAttachment,omitempty"`
	ValueTime            *string                 `json:"valueTime,omitempty"`
	ValueDateTime        *time.Time              `json:"valueDateTime,omitempty"`
	ValuePeriod          *common.Period          `json:"valuePeriod,omitempty"`
}

// ObservationReferenceRange represents provides guide for interpretation
type ObservationReferenceRange struct {
	common.BackboneElement

	// Applicable age range, if relevant
	Age *Range `json:"age,omitempty"`

	// Reference range population
	AppliesTo []common.CodeableConcept `json:"appliesTo,omitempty"`

	// High Range, if relevant
	High *common.Quantity `json:"high,omitempty"`

	// Low Range, if relevant
	Low *common.Quantity `json:"low,omitempty"`

	// Text based reference range in an observation
	Text *string `json:"text,omitempty"`

	// Reference range qualifier
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// ObservationRelated represents resource related to this observation (R3 version)
type ObservationRelated struct {
	common.BackboneElement

	// Reference to the related resource
	Target *common.Reference `json:"target"`

	// has-member | derived-from | sequel-to | replaces | qualified-by | interfered-by
	Type *ObservationRelatedType `json:"type,omitempty"`
}

// Observation-related enums
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

type ObservationRelatedType string

const (
	ObservationRelatedTypeHasMember    ObservationRelatedType = "has-member"
	ObservationRelatedTypeDerivedFrom  ObservationRelatedType = "derived-from"
	ObservationRelatedTypeSequelTo     ObservationRelatedType = "sequel-to"
	ObservationRelatedTypeReplaces     ObservationRelatedType = "replaces"
	ObservationRelatedTypeQualifiedBy  ObservationRelatedType = "qualified-by"
	ObservationRelatedTypeInterferedBy ObservationRelatedType = "interfered-by"
)

// Medication represents a medication for an R3 context
type Medication struct {
	DomainResource

	// Codes that identify this medication
	Code *common.CodeableConcept `json:"code,omitempty"`

	// powder | tablets | capsule +
	Form *common.CodeableConcept `json:"form,omitempty"`

	// Business identifier for this medication
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Active or inactive ingredient
	Ingredient []MedicationIngredient `json:"ingredient,omitempty"`

	// True if a brand
	IsBrand *bool `json:"isBrand,omitempty"`

	// True if medication does not require a prescription
	IsOverTheCounter *bool `json:"isOverTheCounter,omitempty"`

	// Manufacturer of the item
	Manufacturer *common.Reference `json:"manufacturer,omitempty"`

	// Details about packaged medications
	Package *MedicationPackage `json:"package,omitempty"`

	// active | inactive | entered-in-error
	Status *MedicationStatus `json:"status,omitempty"`
}

// MedicationIngredient represents active or inactive ingredient
type MedicationIngredient struct {
	common.BackboneElement

	// The actual ingredient or content
	ItemCodeableConcept *common.CodeableConcept `json:"itemCodeableConcept,omitempty"`
	ItemReference       *common.Reference       `json:"itemReference,omitempty"`

	// Quantity of ingredient present
	Amount *Ratio `json:"amount,omitempty"`

	// Active ingredient indicator
	IsActive *bool `json:"isActive,omitempty"`
}

// MedicationPackage represents details about packaged medications
type MedicationPackage struct {
	common.BackboneElement

	// Batch information
	Batch []MedicationPackageBatch `json:"batch,omitempty"`

	// What is in the package
	Content []MedicationPackageContent `json:"content,omitempty"`

	// E.g. box, vial, blister-pack
	Container *common.CodeableConcept `json:"container,omitempty"`
}

// MedicationPackageBatch represents batch information
type MedicationPackageBatch struct {
	common.BackboneElement

	// When batch will expire
	ExpirationDate *time.Time `json:"expirationDate,omitempty"`

	// Identifier assigned to batch
	LotNumber *string `json:"lotNumber,omitempty"`
}

// MedicationPackageContent represents what is in the package
type MedicationPackageContent struct {
	common.BackboneElement

	// Quantity present in the package
	Amount *common.Quantity `json:"amount,omitempty"`

	// The item in the package
	ItemCodeableConcept *common.CodeableConcept `json:"itemCodeableConcept,omitempty"`
	ItemReference       *common.Reference       `json:"itemReference,omitempty"`
}

// Medication-related enums
type MedicationStatus string

const (
	MedicationStatusActive         MedicationStatus = "active"
	MedicationStatusInactive       MedicationStatus = "inactive"
	MedicationStatusEnteredInError MedicationStatus = "entered-in-error"
)

// MedicationRequest represents a prescription for R3
type MedicationRequest struct {
	DomainResource

	// When request was initially authored
	AuthoredOn *time.Time `json:"authoredOn,omitempty"`

	// Request fulfilled by this request
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// Type of medication usage
	Category *common.CodeableConcept `json:"category,omitempty"`

	// Encounter or Episode of Care created during
	Context *common.Reference `json:"context,omitempty"`

	// Protocol or definition
	Definition []common.Reference `json:"definition,omitempty"`

	// Medication supply authorization
	DispenseRequest *MedicationRequestDispenseRequest `json:"dispenseRequest,omitempty"`

	// How the medication should be taken
	DosageInstruction []Dosage `json:"dosageInstruction,omitempty"`

	// Order identifier
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// active | on-hold | cancelled | completed | entered-in-error | stopped | draft | unknown
	Intent MedicationRequestIntent `json:"intent"`

	// Medication to be taken
	MedicationCodeableConcept *common.CodeableConcept `json:"medicationCodeableConcept,omitempty"`
	MedicationReference       *common.Reference       `json:"medicationReference,omitempty"`

	// Information about the prescription
	Note []Annotation `json:"note,omitempty"`

	// What request fulfills
	PriorPrescription *common.Reference `json:"priorPrescription,omitempty"`

	// routine | urgent | stat | asap
	Priority *MedicationRequestPriority `json:"priority,omitempty"`

	// Reason or indication for writing the prescription
	ReasonCode []common.CodeableConcept `json:"reasonCode,omitempty"`

	// Condition or Observation that supports why the prescription is being written
	ReasonReference []common.Reference `json:"reasonReference,omitempty"`

	// Person who entered the request
	Recorder *common.Reference `json:"recorder,omitempty"`

	// Who/What requested the Request
	Requester *MedicationRequestRequester `json:"requester,omitempty"`

	// active | on-hold | cancelled | completed | entered-in-error | stopped | draft | unknown
	Status MedicationRequestStatus `json:"status"`

	// Who treatment is for
	Subject *common.Reference `json:"subject"`

	// Generic and brand substitution instructions
	Substitution *MedicationRequestSubstitution `json:"substitution,omitempty"`

	// Supporting information for the medication
	SupportingInformation []common.Reference `json:"supportingInformation,omitempty"`
}

// MedicationRequestDispenseRequest represents medication supply authorization
type MedicationRequestDispenseRequest struct {
	common.BackboneElement

	// Minimum period of time between dispenses
	ExpectedSupplyDuration *Duration `json:"expectedSupplyDuration,omitempty"`

	// Number of refills authorized
	NumberOfRepeatsAllowed *int `json:"numberOfRepeatsAllowed,omitempty"`

	// Intended dispenser
	Performer *common.Reference `json:"performer,omitempty"`

	// Amount of medication to supply per dispense
	Quantity *common.Quantity `json:"quantity,omitempty"`

	// Time period supply is authorized for
	ValidityPeriod *common.Period `json:"validityPeriod,omitempty"`
}

// MedicationRequestRequester represents who/what requested the request
type MedicationRequestRequester struct {
	common.BackboneElement

	// Who ordered the initial medication(s)
	Agent *common.Reference `json:"agent"`

	// Organization agent is acting for
	OnBehalfOf *common.Reference `json:"onBehalfOf,omitempty"`
}

// MedicationRequestSubstitution represents generic and brand substitution instructions
type MedicationRequestSubstitution struct {
	common.BackboneElement

	// Whether substitution is allowed or not
	Allowed bool `json:"allowed"`

	// Why should (not) substitution be made
	Reason *common.CodeableConcept `json:"reason,omitempty"`
}

// MedicationRequest-related enums
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

type MedicationRequestIntent string

const (
	MedicationRequestIntentProposal      MedicationRequestIntent = "proposal"
	MedicationRequestIntentPlan          MedicationRequestIntent = "plan"
	MedicationRequestIntentOrder         MedicationRequestIntent = "order"
	MedicationRequestIntentInstanceOrder MedicationRequestIntent = "instance-order"
)

type MedicationRequestPriority string

const (
	MedicationRequestPriorityRoutine MedicationRequestPriority = "routine"
	MedicationRequestPriorityUrgent  MedicationRequestPriority = "urgent"
	MedicationRequestPriorityStat    MedicationRequestPriority = "stat"
	MedicationRequestPriorityASAP    MedicationRequestPriority = "asap"
)

// QuestionnaireResponse represents a structured set of questions and their answers (R3 version)
type QuestionnaireResponse struct {
	DomainResource

	// Person who received and recorded the answers
	Author *common.Reference `json:"author,omitempty"`

	// Date the answers were gathered
	Authored *time.Time `json:"authored,omitempty"`

	// Request fulfilled by this QuestionnaireResponse
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// Encounter or episode during which questionnaire was completed
	Context *common.Reference `json:"context,omitempty"`

	// Unique id for this set of answers
	Identifier *common.Identifier `json:"identifier,omitempty"`

	// Groups and questions
	Item []QuestionnaireResponseItem `json:"item,omitempty"`

	// Composition of questionnaire responses
	Parent []common.Reference `json:"parent,omitempty"`

	// Form being answered (optional in R3)
	Questionnaire *common.Reference `json:"questionnaire,omitempty"`

	// Person who provided the information about the subject
	Source *common.Reference `json:"source,omitempty"`

	// in-progress | completed | amended | entered-in-error | stopped
	Status QuestionnaireResponseStatus `json:"status"`

	// The subject of the questions
	Subject *common.Reference `json:"subject,omitempty"`
}

// QuestionnaireResponseItem represents groups and questions
type QuestionnaireResponseItem struct {
	common.BackboneElement

	// The response(s) to the question
	Answer []QuestionnaireResponseItemAnswer `json:"answer,omitempty"`

	// ElementDefinition - details for the item
	Definition *string `json:"definition,omitempty"`

	// Nested questionnaire response items
	Item []QuestionnaireResponseItem `json:"item,omitempty"`

	// Pointer to specific item from Questionnaire
	LinkId string `json:"linkId"`

	// The subject this group's answers are about
	Subject *common.Reference `json:"subject,omitempty"`

	// Name for group or question text
	Text *string `json:"text,omitempty"`
}

// QuestionnaireResponseItemAnswer represents the response(s) to the question
type QuestionnaireResponseItemAnswer struct {
	common.BackboneElement

	// Nested groups and/or questions
	Item []QuestionnaireResponseItem `json:"item,omitempty"`

	// Single-valued answer to the question
	ValueBoolean    *bool             `json:"valueBoolean,omitempty"`
	ValueDecimal    *float64          `json:"valueDecimal,omitempty"`
	ValueInteger    *int              `json:"valueInteger,omitempty"`
	ValueDate       *string           `json:"valueDate,omitempty"`
	ValueDateTime   *time.Time        `json:"valueDateTime,omitempty"`
	ValueTime       *string           `json:"valueTime,omitempty"`
	ValueString     *string           `json:"valueString,omitempty"`
	ValueUri        *string           `json:"valueUri,omitempty"`
	ValueAttachment *Attachment       `json:"valueAttachment,omitempty"`
	ValueCoding     *common.Coding    `json:"valueCoding,omitempty"`
	ValueQuantity   *common.Quantity  `json:"valueQuantity,omitempty"`
	ValueReference  *common.Reference `json:"valueReference,omitempty"`
}

// QuestionnaireResponse-related enums
type QuestionnaireResponseStatus string

const (
	QuestionnaireResponseStatusInProgress     QuestionnaireResponseStatus = "in-progress"
	QuestionnaireResponseStatusCompleted      QuestionnaireResponseStatus = "completed"
	QuestionnaireResponseStatusAmended        QuestionnaireResponseStatus = "amended"
	QuestionnaireResponseStatusEnteredInError QuestionnaireResponseStatus = "entered-in-error"
	QuestionnaireResponseStatusStopped        QuestionnaireResponseStatus = "stopped"
)

// ResearchSubject represents information about an individual participating in a research study (R3 version)
type ResearchSubject struct {
	DomainResource

	// The actual group assignment for this subject
	ActualArm *string `json:"actualArm,omitempty"`

	// What path should be followed by the subject during the study
	AssignedArm *string `json:"assignedArm,omitempty"`

	// Agreement to participate in study
	Consent *common.Reference `json:"consent,omitempty"`

	// Business Identifier for research subject
	Identifier *common.Identifier `json:"identifier,omitempty"`

	// Who is part of study
	Individual *common.Reference `json:"individual"`

	// Start and end of participation
	Period *common.Period `json:"period,omitempty"`

	// candidate | eligible | follow-up | ineligible | not-registered | off-study | on-study | on-study-intervention | on-study-observation | pending-on-study | potential-candidate | screening | withdrawn
	Status ResearchSubjectStatus `json:"status"`

	// Study subject is part of
	Study *common.Reference `json:"study"`
}

// ResearchSubject-related enums (R3 version - simpler than R4B)
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

// Consent represents patient's or personal representative's formal agreement to allow something to happen (R3 version)
type Consent struct {
	DomainResource

	// Actions controlled by this consent
	Action []common.CodeableConcept `json:"action,omitempty"`

	// Who|what controlled by this consent (subject to this policy)
	Actor []ConsentActor `json:"actor,omitempty"`

	// Classification of the consent statement
	Category []common.CodeableConcept `json:"category,omitempty"`

	// Who is agreeing to the policy and exceptions
	ConsentingParty []common.Reference `json:"consentingParty,omitempty"`

	// Data controlled by this consent
	Data []ConsentData `json:"data,omitempty"`

	// Effective period for this Consent Resource
	DataPeriod *common.Period `json:"dataPeriod,omitempty"`

	// When this Consent was created or indexed
	DateTime *time.Time `json:"dateTime,omitempty"`

	// Additional rule - addition or removal of permissions
	Except []ConsentExcept `json:"except,omitempty"`

	// Identifier for this record (external references)
	Identifier *common.Identifier `json:"identifier,omitempty"`

	// Custodian of the consent
	Organization []common.Reference `json:"organization,omitempty"`

	// Who the consent applies to
	Patient *common.Reference `json:"patient"`

	// Timeframe for this rule
	Period *common.Period `json:"period,omitempty"`

	// Policy that this consents to
	Policy []ConsentPolicy `json:"policy,omitempty"`

	// Policy that this consents to
	PolicyRule *string `json:"policyRule,omitempty"`

	// Context of activities for which the agreement is made
	Purpose []common.Coding `json:"purpose,omitempty"`

	// Security labels that define affected resources
	SecurityLabel []common.Coding `json:"securityLabel,omitempty"`

	// Source from which this consent is taken
	SourceAttachment *Attachment        `json:"sourceAttachment,omitempty"`
	SourceIdentifier *common.Identifier `json:"sourceIdentifier,omitempty"`
	SourceReference  *common.Reference  `json:"sourceReference,omitempty"`

	// draft | proposed | active | rejected | inactive | entered-in-error
	Status ConsentStatus `json:"status"`
}

// ConsentActor represents who|what controlled by this consent
type ConsentActor struct {
	common.BackboneElement

	// Resource for the actor (or group, by role)
	Reference *common.Reference `json:"reference"`

	// How the individual is involved in the resources content that is described in the consent
	Role *common.CodeableConcept `json:"role"`
}

// ConsentPolicy represents policy that this consents to
type ConsentPolicy struct {
	common.BackboneElement

	// Enforcement source for policy
	Authority *string `json:"authority,omitempty"`

	// Specific policy covered by this consent
	URI *string `json:"uri,omitempty"`
}

// ConsentData represents data controlled by this consent
type ConsentData struct {
	common.BackboneElement

	// How the resource reference is interpreted when testing consent restrictions
	Meaning ConsentDataMeaning `json:"meaning"`

	// The actual data reference
	Reference *common.Reference `json:"reference"`
}

// ConsentExcept represents additional rule - addition or removal of permissions
type ConsentExcept struct {
	common.BackboneElement

	// Actions controlled by this exception
	Action []common.CodeableConcept `json:"action,omitempty"`

	// Who|what controlled by this exception
	Actor []ConsentExceptActor `json:"actor,omitempty"`

	// e.g. Resource Type, Profile, or CDA etc
	Class []common.Coding `json:"class,omitempty"`

	// e.g. LOINC or SNOMED CT code, etc in the content
	Code []common.Coding `json:"code,omitempty"`

	// Data controlled by this exception
	Data []ConsentExceptData `json:"data,omitempty"`

	// Timeframe for data controlled by this exception
	DataPeriod *common.Period `json:"dataPeriod,omitempty"`

	// Timeframe for this exception
	Period *common.Period `json:"period,omitempty"`

	// Context of activities covered by this exception
	Purpose []common.Coding `json:"purpose,omitempty"`

	// Security Labels that define affected resources
	SecurityLabel []common.Coding `json:"securityLabel,omitempty"`

	// deny | permit
	Type ConsentExceptType `json:"type"`
}

// ConsentExceptActor represents who|what controlled by this exception
type ConsentExceptActor struct {
	common.BackboneElement

	// Resource for the actor (or group, by role)
	Reference *common.Reference `json:"reference"`

	// How the individual is involved in the resources content that is described in the exception
	Role *common.CodeableConcept `json:"role"`
}

// ConsentExceptData represents data controlled by this exception
type ConsentExceptData struct {
	common.BackboneElement

	// How the resource reference is interpreted when testing consent restrictions
	Meaning ConsentDataMeaning `json:"meaning"`

	// The actual data reference
	Reference *common.Reference `json:"reference"`
}

// Consent-related enums
type ConsentStatus string

const (
	ConsentStatusDraft          ConsentStatus = "draft"
	ConsentStatusProposed       ConsentStatus = "proposed"
	ConsentStatusActive         ConsentStatus = "active"
	ConsentStatusRejected       ConsentStatus = "rejected"
	ConsentStatusInactive       ConsentStatus = "inactive"
	ConsentStatusEnteredInError ConsentStatus = "entered-in-error"
)

type ConsentDataMeaning string

const (
	ConsentDataMeaningInstance   ConsentDataMeaning = "instance"
	ConsentDataMeaningRelated    ConsentDataMeaning = "related"
	ConsentDataMeaningDependents ConsentDataMeaning = "dependents"
	ConsentDataMeaningAuthoredBy ConsentDataMeaning = "authoredby"
)

type ConsentExceptType string

const (
	ConsentExceptTypeDeny   ConsentExceptType = "deny"
	ConsentExceptTypePermit ConsentExceptType = "permit"
)

// Device represents an instance of a medical device used in the provision of patient care (R3 version)
type Device struct {
	DomainResource

	// Details for human/organization for support
	Contact []ContactPoint `json:"contact,omitempty"`

	// Date and time of expiry of this device (if applicable)
	ExpirationDate *time.Time `json:"expirationDate,omitempty"`

	// Instance identifier
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Where the resource is found
	Location *common.Reference `json:"location,omitempty"`

	// Lot number of manufacture
	LotNumber *string `json:"lotNumber,omitempty"`

	// Date when the device was made
	ManufactureDate *time.Time `json:"manufactureDate,omitempty"`

	// Name of device manufacturer
	Manufacturer *string `json:"manufacturer,omitempty"`

	// Model id assigned by the manufacturer
	Model *string `json:"model,omitempty"`

	// Device notes and comments
	Note []Annotation `json:"note,omitempty"`

	// Organization responsible for device
	Owner *common.Reference `json:"owner,omitempty"`

	// Patient to whom Device is affixed
	Patient *common.Reference `json:"patient,omitempty"`

	// The capabilities supported by the device
	Safety []common.CodeableConcept `json:"safety,omitempty"`

	// active | inactive | entered-in-error | unknown
	Status *DeviceStatus `json:"status,omitempty"`

	// The kind or type of device
	Type *common.CodeableConcept `json:"type,omitempty"`

	// Unique Device Identifier (UDI) Barcode string
	Udi *DeviceUdi `json:"udi,omitempty"`

	// Network address to contact device
	URL *string `json:"url,omitempty"`

	// Version number (i.e. software)
	Version *string `json:"version,omitempty"`
}

// DeviceUdi represents unique Device Identifier (UDI) Barcode string (R3 version)
type DeviceUdi struct {
	common.BackboneElement

	// UDI Machine Readable Barcode String
	CarrierAIDC *string `json:"carrierAIDC,omitempty"`

	// UDI Human Readable Barcode String
	CarrierHRF *string `json:"carrierHRF,omitempty"`

	// A fixed portion of the UDI that identifies the labeler and the specific version or model of a device
	DeviceIdentifier *string `json:"deviceIdentifier,omitempty"`

	// UDI Issuing Organization
	Issuer *string `json:"issuer,omitempty"`

	// Regional UDI authority
	Jurisdiction *string `json:"jurisdiction,omitempty"`

	// Device Name as appears on UDI label
	Name *string `json:"name,omitempty"`
}

// Device-related enums
type DeviceStatus string

const (
	DeviceStatusActive         DeviceStatus = "active"
	DeviceStatusInactive       DeviceStatus = "inactive"
	DeviceStatusEnteredInError DeviceStatus = "entered-in-error"
	DeviceStatusUnknown        DeviceStatus = "unknown"
)

// Composition represents a document authored/attested by humans, organizations and devices (R3 version)
type Composition struct {
	DomainResource

	// Attests to accuracy of composition
	Attester []CompositionAttester `json:"attester,omitempty"`

	// Who and/or what authored the composition
	Author []common.Reference `json:"author"`

	// Categorization of Composition
	Class *common.CodeableConcept `json:"class,omitempty"`

	// As defined by affinity domain
	Confidentiality *string `json:"confidentiality,omitempty"`

	// Organization which maintains the composition
	Custodian *common.Reference `json:"custodian,omitempty"`

	// Composition editing time
	Date time.Time `json:"date"`

	// Context of the Composition
	Encounter *common.Reference `json:"encounter,omitempty"`

	// For documents, the persistent identifier
	Identifier *common.Identifier `json:"identifier,omitempty"`

	// Composition is broken into sections
	Section []CompositionSection `json:"section,omitempty"`

	// preliminary | final | amended | entered-in-error
	Status CompositionStatus `json:"status"`

	// Who and/or what the composition is about
	Subject *common.Reference `json:"subject"`

	// Human Readable name/title
	Title string `json:"title"`

	// Kind of composition (LOINC if possible)
	Type *common.CodeableConcept `json:"type"`
}

// CompositionAttester represents attests to accuracy of composition
type CompositionAttester struct {
	common.BackboneElement

	// individual | organization
	Mode []CompositionAttesterMode `json:"mode"`

	// Who attested the composition
	Party *common.Reference `json:"party,omitempty"`

	// When composition attested
	Time *time.Time `json:"time,omitempty"`
}

// CompositionSection represents composition is broken into sections
type CompositionSection struct {
	common.BackboneElement

	// Classification of section (recommended)
	Code *common.CodeableConcept `json:"code,omitempty"`

	// How the entry list was prepared - whether it is a working list, snapshot, or changes
	EmptyReason *common.CodeableConcept `json:"emptyReason,omitempty"`

	// A reference to data that supports this section
	Entry []common.Reference `json:"entry,omitempty"`

	// working | snapshot | changes
	Mode *CompositionSectionMode `json:"mode,omitempty"`

	// Order of section entries
	OrderedBy *common.CodeableConcept `json:"orderedBy,omitempty"`

	// Nested Section
	Section []CompositionSection `json:"section,omitempty"`

	// Text summary of the section, for human interpretation
	Text *Narrative `json:"text,omitempty"`

	// Label for section (e.g. for ToC)
	Title *string `json:"title,omitempty"`
}

// Composition-related enums
type CompositionStatus string

const (
	CompositionStatusPreliminary    CompositionStatus = "preliminary"
	CompositionStatusFinal          CompositionStatus = "final"
	CompositionStatusAmended        CompositionStatus = "amended"
	CompositionStatusEnteredInError CompositionStatus = "entered-in-error"
)

type CompositionAttesterMode string

const (
	CompositionAttesterModePersonal     CompositionAttesterMode = "personal"
	CompositionAttesterModeProfessional CompositionAttesterMode = "professional"
	CompositionAttesterModeLegal        CompositionAttesterMode = "legal"
	CompositionAttesterModeOfficial     CompositionAttesterMode = "official"
)

type CompositionSectionMode string

const (
	CompositionSectionModeWorking  CompositionSectionMode = "working"
	CompositionSectionModeSnapshot CompositionSectionMode = "snapshot"
	CompositionSectionModeChanges  CompositionSectionMode = "changes"
)

// DocumentReference represents a reference to a document (R3 version)
type DocumentReference struct {
	DomainResource

	// Who/what authenticated the document
	Authenticator *common.Reference `json:"authenticator,omitempty"`

	// Who and/or what authored the document
	Author []common.Reference `json:"author,omitempty"`

	// Categorization of document
	Class *common.CodeableConcept `json:"class,omitempty"`

	// Document referenced
	Content []DocumentReferenceContent `json:"content"`

	// Clinical context of document
	Context *DocumentReferenceContext `json:"context,omitempty"`

	// When this document reference was created
	Created *time.Time `json:"created,omitempty"`

	// Organization which maintains the document
	Custodian *common.Reference `json:"custodian,omitempty"`

	// Human-readable description (title)
	Description *string `json:"description,omitempty"`

	// Unique identifier for the document
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// When this document reference was indexed
	Indexed time.Time `json:"indexed"`

	// Document security-tags
	SecurityLabel []common.CodeableConcept `json:"securityLabel,omitempty"`

	// current | superseded | entered-in-error
	Status DocumentReferenceStatus `json:"status"`

	// Who/what is the subject of the document
	Subject *common.Reference `json:"subject,omitempty"`

	// Kind of document (LOINC if possible)
	Type *common.CodeableConcept `json:"type"`
}

// DocumentReferenceContent represents document referenced
type DocumentReferenceContent struct {
	common.BackboneElement

	// Where to access the document
	Attachment *Attachment `json:"attachment"`

	// Format/content rules for the document
	Format *common.Coding `json:"format,omitempty"`
}

// DocumentReferenceContext represents clinical context of document
type DocumentReferenceContext struct {
	common.BackboneElement

	// Context of the document content
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Main clinical acts documented
	Event []common.CodeableConcept `json:"event,omitempty"`

	// Kind of facility where patient was seen
	FacilityType *common.CodeableConcept `json:"facilityType,omitempty"`

	// Time of service that is being documented
	Period *common.Period `json:"period,omitempty"`

	// Additional details about where the content was created (e.g. clinical specialty)
	PracticeSetting *common.CodeableConcept `json:"practiceSetting,omitempty"`

	// Related identifiers or resources
	Related []DocumentReferenceContextRelated `json:"related,omitempty"`

	// Patient demographics from source
	SourcePatientInfo *common.Reference `json:"sourcePatientInfo,omitempty"`
}

// DocumentReferenceContextRelated represents related identifiers or resources
type DocumentReferenceContextRelated struct {
	common.BackboneElement

	// Identifier of related objects or events
	Identifier *common.Identifier `json:"identifier,omitempty"`

	// Related Resource
	Ref *common.Reference `json:"ref,omitempty"`
}

// DocumentReference-related enums
type DocumentReferenceStatus string

const (
	DocumentReferenceStatusCurrent        DocumentReferenceStatus = "current"
	DocumentReferenceStatusSuperseded     DocumentReferenceStatus = "superseded"
	DocumentReferenceStatusEnteredInError DocumentReferenceStatus = "entered-in-error"
)

// Communication represents an occurrence of information being transmitted (R3 version)
type Communication struct {
	DomainResource

	// Request fulfilled by this communication
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// Message category
	Category []common.CodeableConcept `json:"category,omitempty"`

	// Encounter or episode leading to message
	Context *common.Reference `json:"context,omitempty"`

	// Instantiates protocol or definition
	Definition []common.Reference `json:"definition,omitempty"`

	// Unique identifier
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// A channel of communication
	Medium []common.CodeableConcept `json:"medium,omitempty"`

	// Comments made about the communication
	Note []Annotation `json:"note,omitempty"`

	// Part of this action
	PartOf []common.Reference `json:"partOf,omitempty"`

	// Message payload
	Payload []CommunicationPayload `json:"payload,omitempty"`

	// Indication for message
	ReasonCode []common.CodeableConcept `json:"reasonCode,omitempty"`

	// Why was communication done?
	ReasonReference []common.Reference `json:"reasonReference,omitempty"`

	// When received
	Received *time.Time `json:"received,omitempty"`

	// Message recipient
	Recipient []common.Reference `json:"recipient,omitempty"`

	// Message sender
	Sender *common.Reference `json:"sender,omitempty"`

	// When sent
	Sent *time.Time `json:"sent,omitempty"`

	// preparation | in-progress | suspended | aborted | completed | entered-in-error
	Status CommunicationStatus `json:"status"`

	// Focus of message
	Subject *common.Reference `json:"subject,omitempty"`

	// Communication did not occur
	NotDone *bool `json:"notDone,omitempty"`

	// Why communication did not occur
	NotDoneReason *common.CodeableConcept `json:"notDoneReason,omitempty"`
}

// CommunicationPayload represents message payload
type CommunicationPayload struct {
	common.BackboneElement

	// Message part content
	ContentString     *string           `json:"contentString,omitempty"`
	ContentAttachment *Attachment       `json:"contentAttachment,omitempty"`
	ContentReference  *common.Reference `json:"contentReference,omitempty"`
}

// Communication-related enums
type CommunicationStatus string

const (
	CommunicationStatusPreparation    CommunicationStatus = "preparation"
	CommunicationStatusInProgress     CommunicationStatus = "in-progress"
	CommunicationStatusSuspended      CommunicationStatus = "suspended"
	CommunicationStatusAborted        CommunicationStatus = "aborted"
	CommunicationStatusCompleted      CommunicationStatus = "completed"
	CommunicationStatusEnteredInError CommunicationStatus = "entered-in-error"
)

// Dosage represents how medication should be taken (R3 version)
type Dosage struct {
	common.BackboneElement

	// Additional instruction such as "Swallow with plenty of water"
	AdditionalInstruction []common.CodeableConcept `json:"additionalInstruction,omitempty"`

	// Take "as needed" (for x)
	AsNeededBoolean         *bool                   `json:"asNeededBoolean,omitempty"`
	AsNeededCodeableConcept *common.CodeableConcept `json:"asNeededCodeableConcept,omitempty"`

	// Amount of medication per dose
	DoseQuantity *common.Quantity `json:"doseQuantity,omitempty"`
	DoseRange    *Range           `json:"doseRange,omitempty"`

	// Upper limit on medication per unit of time
	MaxDosePerAdministration *common.Quantity `json:"maxDosePerAdministration,omitempty"`
	MaxDosePerLifetime       *common.Quantity `json:"maxDosePerLifetime,omitempty"`
	MaxDosePerPeriod         *Ratio           `json:"maxDosePerPeriod,omitempty"`

	// Technique for administering medication
	Method *common.CodeableConcept `json:"method,omitempty"`

	// Patient or consumer oriented instructions
	PatientInstruction *string `json:"patientInstruction,omitempty"`

	// Amount of medication per unit of time
	RateQuantity *common.Quantity `json:"rateQuantity,omitempty"`
	RateRange    *Range           `json:"rateRange,omitempty"`
	RateRatio    *Ratio           `json:"rateRatio,omitempty"`

	// How drug should enter body
	Route *common.CodeableConcept `json:"route,omitempty"`

	// The order of the dosage instructions
	Sequence *int `json:"sequence,omitempty"`

	// Body site to administer to
	Site *common.CodeableConcept `json:"site,omitempty"`

	// Free text dosage instructions e.g. SIG
	Text *string `json:"text,omitempty"`

	// When medication should be administered
	Timing *Timing `json:"timing,omitempty"`
}

// Timing represents a timing schedule that specifies an event that may occur multiple times (R3 version)
type Timing struct {
	common.BackboneElement

	// When the event occurs
	Code *common.CodeableConcept `json:"code,omitempty"`

	// When the event is to occur
	Event []time.Time `json:"event,omitempty"`

	// When the event is to occur
	Repeat *TimingRepeat `json:"repeat,omitempty"`
}

// TimingRepeat represents when the event is to occur
type TimingRepeat struct {
	common.BackboneElement

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
	DurationUnit *TimingRepeatDurationUnit `json:"durationUnit,omitempty"`

	// Event occurs frequency times per period
	Frequency *int `json:"frequency,omitempty"`

	// Event occurs up to frequencyMax times per period
	FrequencyMax *int `json:"frequencyMax,omitempty"`

	// Minutes from event (before or after)
	Offset *int `json:"offset,omitempty"`

	// Event occurs frequency times per period
	Period *float64 `json:"period,omitempty"`

	// Upper limit of period (3-4 hours)
	PeriodMax *float64 `json:"periodMax,omitempty"`

	// s | min | h | d | wk | mo | a - unit of time (UCUM)
	PeriodUnit *TimingRepeatPeriodUnit `json:"periodUnit,omitempty"`

	// Time of day for action
	TimeOfDay []string `json:"timeOfDay,omitempty"`

	// mon | tue | wed | thu | fri | sat | sun
	DayOfWeek []TimingRepeatDayOfWeek `json:"dayOfWeek,omitempty"`

	// Regular life events the event is tied to
	When []TimingRepeatWhen `json:"when,omitempty"`
}

// Timing-related enums
type TimingRepeatDurationUnit string

const (
	TimingRepeatDurationUnitS   TimingRepeatDurationUnit = "s"
	TimingRepeatDurationUnitMin TimingRepeatDurationUnit = "min"
	TimingRepeatDurationUnitH   TimingRepeatDurationUnit = "h"
	TimingRepeatDurationUnitD   TimingRepeatDurationUnit = "d"
	TimingRepeatDurationUnitWk  TimingRepeatDurationUnit = "wk"
	TimingRepeatDurationUnitMo  TimingRepeatDurationUnit = "mo"
	TimingRepeatDurationUnitA   TimingRepeatDurationUnit = "a"
)

type TimingRepeatPeriodUnit string

const (
	TimingRepeatPeriodUnitS   TimingRepeatPeriodUnit = "s"
	TimingRepeatPeriodUnitMin TimingRepeatPeriodUnit = "min"
	TimingRepeatPeriodUnitH   TimingRepeatPeriodUnit = "h"
	TimingRepeatPeriodUnitD   TimingRepeatPeriodUnit = "d"
	TimingRepeatPeriodUnitWk  TimingRepeatPeriodUnit = "wk"
	TimingRepeatPeriodUnitMo  TimingRepeatPeriodUnit = "mo"
	TimingRepeatPeriodUnitA   TimingRepeatPeriodUnit = "a"
)

type TimingRepeatDayOfWeek string

const (
	TimingRepeatDayOfWeekMon TimingRepeatDayOfWeek = "mon"
	TimingRepeatDayOfWeekTue TimingRepeatDayOfWeek = "tue"
	TimingRepeatDayOfWeekWed TimingRepeatDayOfWeek = "wed"
	TimingRepeatDayOfWeekThu TimingRepeatDayOfWeek = "thu"
	TimingRepeatDayOfWeekFri TimingRepeatDayOfWeek = "fri"
	TimingRepeatDayOfWeekSat TimingRepeatDayOfWeek = "sat"
	TimingRepeatDayOfWeekSun TimingRepeatDayOfWeek = "sun"
)

type TimingRepeatWhen string

const (
	TimingRepeatWhenHS   TimingRepeatWhen = "HS"
	TimingRepeatWhenWAKE TimingRepeatWhen = "WAKE"
	TimingRepeatWhenC    TimingRepeatWhen = "C"
	TimingRepeatWhenCM   TimingRepeatWhen = "CM"
	TimingRepeatWhenCD   TimingRepeatWhen = "CD"
	TimingRepeatWhenCV   TimingRepeatWhen = "CV"
	TimingRepeatWhenAC   TimingRepeatWhen = "AC"
	TimingRepeatWhenACM  TimingRepeatWhen = "ACM"
	TimingRepeatWhenACD  TimingRepeatWhen = "ACD"
	TimingRepeatWhenACV  TimingRepeatWhen = "ACV"
	TimingRepeatWhenPC   TimingRepeatWhen = "PC"
	TimingRepeatWhenPCM  TimingRepeatWhen = "PCM"
	TimingRepeatWhenPCD  TimingRepeatWhen = "PCD"
	TimingRepeatWhenPCV  TimingRepeatWhen = "PCV"
)

// SampledData represents a series of measurements taken by a device (R3 version)
type SampledData struct {
	common.Element

	// Decimal values with spaces, or "E" | "U" | "L"
	Data string `json:"data"`

	// Number of sample points at each time point
	Dimensions int `json:"dimensions"`

	// Multiply data by this before adding to origin
	Factor *float64 `json:"factor,omitempty"`

	// Lower limit of detection
	LowerLimit *float64 `json:"lowerLimit,omitempty"`

	// Zero value and units
	Origin *common.Quantity `json:"origin"`

	// Number of milliseconds between samples
	Period float64 `json:"period"`

	// Upper limit of detection
	UpperLimit *float64 `json:"upperLimit,omitempty"`
}

// Narrative represents a human-readable formatted text, including images (R3 version)
type Narrative struct {
	common.Element

	// generated | extensions | additional | empty
	Status NarrativeStatus `json:"status"`

	// Limited xhtml content
	Div string `json:"div"`
}

// Narrative-related enums
type NarrativeStatus string

const (
	NarrativeStatusGenerated  NarrativeStatus = "generated"
	NarrativeStatusExtensions NarrativeStatus = "extensions"
	NarrativeStatusAdditional NarrativeStatus = "additional"
	NarrativeStatusEmpty      NarrativeStatus = "empty"
)

// Signature represents a digital signature along with supporting context (R3 version)
type Signature struct {
	common.Element

	// The actual signature content (XML DigSig, JWT, picture, etc.)
	Blob *string `json:"blob,omitempty"`

	// The technical format of the signature
	ContentType *string `json:"contentType,omitempty"`

	// Indication of the reason the entity signed the object(s)
	Type []common.Coding `json:"type"`

	// Who signed
	WhoURI       *string           `json:"whoUri,omitempty"`
	WhoReference *common.Reference `json:"whoReference,omitempty"`

	// When the signature was created
	When time.Time `json:"when"`
}

// AllergyIntolerance represents risk of harmful or undesirable, physiological response to a substance (R3)
type AllergyIntolerance struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "AllergyIntolerance"

	// External ids for this item
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// active | inactive | resolved (simplified in R3)
	ClinicalStatus *AllergyIntoleranceClinicalStatus `json:"clinicalStatus,omitempty"`

	// unconfirmed | confirmed | refuted | entered-in-error (required in R3)
	VerificationStatus AllergyIntoleranceVerificationStatus `json:"verificationStatus"`

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

	// Date first version of the resource instance was recorded
	OnsetDateTime *time.Time     `json:"onsetDateTime,omitempty"`
	OnsetAge      *Age           `json:"onsetAge,omitempty"`
	OnsetPeriod   *common.Period `json:"onsetPeriod,omitempty"`
	OnsetRange    *Range         `json:"onsetRange,omitempty"`
	OnsetString   *string        `json:"onsetString,omitempty"`

	// When allergy or intolerance was identified
	AssertedDate *time.Time `json:"assertedDate,omitempty"`

	// Who recorded the sensitivity
	Recorder *common.Reference `json:"recorder,omitempty"`

	// Source of the information about the allergy
	Asserter *common.Reference `json:"asserter,omitempty"`

	// Date(/time) of last known occurrence of a reaction
	LastOccurrence *time.Time `json:"lastOccurrence,omitempty"`

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

// AllergyIntoleranceClinicalStatus represents the clinical status (R3 simplified)
type AllergyIntoleranceClinicalStatus string

const (
	AllergyIntoleranceClinicalStatusActive   AllergyIntoleranceClinicalStatus = "active"
	AllergyIntoleranceClinicalStatusInactive AllergyIntoleranceClinicalStatus = "inactive"
	AllergyIntoleranceClinicalStatusResolved AllergyIntoleranceClinicalStatus = "resolved"
)

// AllergyIntoleranceVerificationStatus represents the verification status (required in R3)
type AllergyIntoleranceVerificationStatus string

const (
	AllergyIntoleranceVerificationStatusUnconfirmed    AllergyIntoleranceVerificationStatus = "unconfirmed"
	AllergyIntoleranceVerificationStatusConfirmed      AllergyIntoleranceVerificationStatus = "confirmed"
	AllergyIntoleranceVerificationStatusRefuted        AllergyIntoleranceVerificationStatus = "refuted"
	AllergyIntoleranceVerificationStatusEnteredInError AllergyIntoleranceVerificationStatus = "entered-in-error"
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
	Onset *time.Time `json:"onset,omitempty"`

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

// Immunization represents the event of a patient being administered a vaccine (R3)
type Immunization struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Immunization"

	// Business identifier
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// completed | entered-in-error (R3 has different status values)
	Status ImmunizationStatus `json:"status"`

	// Indicates if the immunization was given (R3 uses notGiven)
	NotGiven bool `json:"notGiven"`

	// Vaccine product administered
	VaccineCode common.CodeableConcept `json:"vaccineCode"`

	// Who was immunized
	Patient common.Reference `json:"patient"`

	// Encounter immunization was part of
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Vaccine administration date
	Date *time.Time `json:"date,omitempty"`

	// Indicates context the data was recorded in (R3)
	PrimarySource bool `json:"primarySource"`

	// Indicates the source of the reported record
	ReportOrigin *common.CodeableConcept `json:"reportOrigin,omitempty"`

	// Where immunization occurred
	Location *common.Reference `json:"location,omitempty"`

	// Vaccine manufacturer
	Manufacturer *common.Reference `json:"manufacturer,omitempty"`

	// Vaccine lot number
	LotNumber *string `json:"lotNumber,omitempty"`

	// Vaccine expiration date
	ExpirationDate *time.Time `json:"expirationDate,omitempty"`

	// Body site vaccine was administered
	Site *common.CodeableConcept `json:"site,omitempty"`

	// How vaccine entered body
	Route *common.CodeableConcept `json:"route,omitempty"`

	// Amount of vaccine administered
	DoseQuantity *common.Quantity `json:"doseQuantity,omitempty"`

	// Who performed event
	Practitioner []ImmunizationPractitioner `json:"practitioner,omitempty"`

	// Additional immunization notes
	Note []Annotation `json:"note,omitempty"`

	// Why immunization occurred
	Explanation *ImmunizationExplanation `json:"explanation,omitempty"`

	// Details of a reaction that follows immunization
	Reaction []ImmunizationReaction `json:"reaction,omitempty"`

	// Protocol followed by the provider
	VaccinationProtocol []ImmunizationVaccinationProtocol `json:"vaccinationProtocol,omitempty"`
}

// ImmunizationStatus represents the current status of the vaccination event (R3)
type ImmunizationStatus string

const (
	ImmunizationStatusCompleted      ImmunizationStatus = "completed"
	ImmunizationStatusEnteredInError ImmunizationStatus = "entered-in-error"
)

// ImmunizationPractitioner represents who performed the immunization event (R3)
type ImmunizationPractitioner struct {
	common.BackboneElement

	// What type of performance was done
	Role *common.CodeableConcept `json:"role,omitempty"`

	// Individual who was performing
	Actor common.Reference `json:"actor"`
}

// ImmunizationExplanation represents why immunization occurred or did not occur (R3)
type ImmunizationExplanation struct {
	common.BackboneElement

	// Why immunization occurred
	Reason []common.CodeableConcept `json:"reason,omitempty"`

	// Why immunization did not occur
	ReasonNotGiven []common.CodeableConcept `json:"reasonNotGiven,omitempty"`
}

// ImmunizationReaction represents details of a reaction following immunization
type ImmunizationReaction struct {
	common.BackboneElement

	// When reaction started
	Date *time.Time `json:"date,omitempty"`

	// Additional information on reaction
	Detail *common.Reference `json:"detail,omitempty"`

	// Indicates self-reported reaction
	Reported *bool `json:"reported,omitempty"`
}

// ImmunizationVaccinationProtocol represents protocol followed by the provider (R3)
type ImmunizationVaccinationProtocol struct {
	common.BackboneElement

	// Dose number within series
	DoseSequence int `json:"doseSequence"`

	// Details of the series
	Description *string `json:"description,omitempty"`

	// Who is responsible for publishing the recommendations
	Authority *common.Reference `json:"authority,omitempty"`

	// Name of vaccine series
	Series *string `json:"series,omitempty"`

	// Recommended number of doses for immunity
	SeriesDoses *int `json:"seriesDoses,omitempty"`

	// Vaccine preventable disease being targeted
	TargetDisease []common.CodeableConcept `json:"targetDisease"`

	// Indicates if the immunization event should "count" against the protocol
	DoseStatus common.CodeableConcept `json:"doseStatus"`

	// Why dose does (not) count
	DoseStatusReason *common.CodeableConcept `json:"doseStatusReason,omitempty"`
}

// CarePlan represents the intention of how one or more practitioners intend to deliver care (R3)
type CarePlan struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "CarePlan"

	// External Ids for this plan
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Protocol or definition followed by this plan
	Definition []common.Reference `json:"definition,omitempty"`

	// Fulfills CarePlan
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// CarePlans with some sort of formal relationship to the current plan
	Replaces []common.Reference `json:"replaces,omitempty"`

	// Part of referenced CarePlan
	PartOf []common.Reference `json:"partOf,omitempty"`

	// draft | active | suspended | completed | entered-in-error | cancelled | unknown
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

	// Created in context of
	Context *common.Reference `json:"context,omitempty"`

	// Time period plan covers
	Period *common.Period `json:"period,omitempty"`

	// Who's involved in plan?
	Author []common.Reference `json:"author,omitempty"`

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

// CarePlanStatus represents the workflow state of the care plan (R3)
type CarePlanStatus string

const (
	CarePlanStatusDraft          CarePlanStatus = "draft"
	CarePlanStatusActive         CarePlanStatus = "active"
	CarePlanStatusSuspended      CarePlanStatus = "suspended"
	CarePlanStatusCompleted      CarePlanStatus = "completed"
	CarePlanStatusEnteredInError CarePlanStatus = "entered-in-error"
	CarePlanStatusCancelled      CarePlanStatus = "cancelled"
	CarePlanStatusUnknown        CarePlanStatus = "unknown"
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

	// High-level categorization of the type of activity
	Category *common.CodeableConcept `json:"category,omitempty"`

	// Protocol or definition followed by this activity
	Definition *common.Reference `json:"definition,omitempty"`

	// Detail type of activity
	Code *common.CodeableConcept `json:"code,omitempty"`

	// Why activity should be done or why activity was prohibited
	ReasonCode []common.CodeableConcept `json:"reasonCode,omitempty"`

	// Condition triggering need for activity
	ReasonReference []common.Reference `json:"reasonReference,omitempty"`

	// Goals this activity relates to
	Goal []common.Reference `json:"goal,omitempty"`

	// not-started | scheduled | in-progress | on-hold | completed | cancelled | stopped | unknown | entered-in-error
	Status CarePlanActivityStatus `json:"status"`

	// Do NOT do
	Prohibited bool `json:"prohibited"`

	// When activity is to occur
	ScheduledTiming *Timing        `json:"scheduledTiming,omitempty"`
	ScheduledPeriod *common.Period `json:"scheduledPeriod,omitempty"`
	ScheduledString *string        `json:"scheduledString,omitempty"`

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

// CarePlanActivityStatus represents the current state of the activity (R3)
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

// Specimen represents a sample to be used for analysis (R3)
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
	Subject common.Reference `json:"subject"`

	// Time when specimen was received for processing
	ReceivedTime *time.Time `json:"receivedTime,omitempty"`

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
	CollectedDateTime *time.Time     `json:"collectedDateTime,omitempty"`
	CollectedPeriod   *common.Period `json:"collectedPeriod,omitempty"`

	// How much was collected
	Quantity *common.Quantity `json:"quantity,omitempty"`

	// Technique used to perform collection
	Method *common.CodeableConcept `json:"method,omitempty"`

	// Anatomical collection site
	BodySite *common.CodeableConcept `json:"bodySite,omitempty"`
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
	TimeDateTime *time.Time     `json:"timeDateTime,omitempty"`
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
