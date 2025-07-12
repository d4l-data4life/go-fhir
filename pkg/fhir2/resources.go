// Package fhir2 contains FHIR R2 (version 1.0.2) resource definitions
package fhir2

import (
	"time"

	"github.com/go-fhir/go-fhir/pkg/common"
)

// Helper functions for creating pointers to basic types
func StringPtr(s string) *string     { return &s }
func IntPtr(i int) *int              { return &i }
func BoolPtr(b bool) *bool           { return &b }
func TimePtr(t time.Time) *time.Time { return &t }

// Base types for FHIR R2

// Resource is the base resource type for everything (R2 version)
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

// DomainResource represents a resource that includes narrative, extensions, and contained resources (R2 version)
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

// Meta provides metadata about a resource (R2 version)
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

// HumanName represents a human name, with the ability to identify parts and usage (R2 version)
type HumanName struct {
	common.Element

	// usual | official | temp | nickname | anonymous | old | maiden
	Use *HumanNameUse `json:"use,omitempty"`

	// Text representation of the full name
	Text *string `json:"text,omitempty"`

	// Family name (often called 'Surname')
	Family []string `json:"family,omitempty"` // R2 uses array, not single string

	// Given names (not always 'first'). Includes middle names
	Given []string `json:"given,omitempty"`

	// Parts that come before the name
	Prefix []string `json:"prefix,omitempty"`

	// Parts that come after the name
	Suffix []string `json:"suffix,omitempty"`

	// Time period when name was/is in use
	Period *common.Period `json:"period,omitempty"`
}

// HumanNameUse represents the use of a human name (R2 version)
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

// Patient represents information about an individual or animal receiving care (R2 version)
type Patient struct {
	DomainResource

	// Whether this patient's record is in active use
	Active *bool `json:"active,omitempty"`

	// Addresses for the individual
	Address []Address `json:"address,omitempty"`

	// Animal-specific information (R2 version)
	Animal *PatientAnimal `json:"animal,omitempty"`

	// The date of birth for the individual
	BirthDate *string `json:"birthDate,omitempty"`

	// Patient's nominated care provider
	CareProvider []common.Reference `json:"careProvider,omitempty"`

	// Languages spoken by the patient
	Communication []PatientCommunication `json:"communication,omitempty"`

	// A contact party (e.g. guardian, partner, friend) for the patient
	Contact []PatientContact `json:"contact,omitempty"`

	// Indicates if the individual is deceased or not
	DeceasedBoolean  *bool   `json:"deceasedBoolean,omitempty"`
	DeceasedDateTime *string `json:"deceasedDateTime,omitempty"` // R2 uses string, not time.Time

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

// PatientAnimal represents animal-specific information (R2 version)
type PatientAnimal struct {
	common.BackboneElement

	// Identifies the high level taxonomic categorization of the kind of animal
	Species *common.CodeableConcept `json:"species"`

	// Identifies the detailed categorization of the kind of animal
	Breed *common.CodeableConcept `json:"breed,omitempty"`

	// Indicates the current state of the animal's reproductive organs
	GenderStatus *common.CodeableConcept `json:"genderStatus,omitempty"`
}

// PatientCommunication represents languages spoken by the patient (R2 version)
type PatientCommunication struct {
	common.BackboneElement

	// The language which can be used to communicate with the patient about their health
	Language *common.CodeableConcept `json:"language"`

	// Language preference indicator
	Preferred *bool `json:"preferred,omitempty"`
}

// PatientContact represents a contact party for the patient (R2 version)
type PatientContact struct {
	common.BackboneElement

	// Address for the contact person
	Address *Address `json:"address,omitempty"`

	// male | female | other | unknown
	Gender *PatientGender `json:"gender,omitempty"`

	// A name associated with the contact person
	Name *HumanName `json:"name,omitempty"`

	// The period during which this contact person or organization is valid to be contacted
	Period *common.Period `json:"period,omitempty"`

	// The kind of relationship
	Relationship []common.CodeableConcept `json:"relationship,omitempty"`

	// A contact detail for the person
	Telecom []ContactPoint `json:"telecom,omitempty"`
}

// PatientLink represents a link to another patient resource (R2 version)
type PatientLink struct {
	common.BackboneElement

	// The other patient resource that the link refers to
	Other *common.Reference `json:"other"`

	// replace | replaces | refer | seealso
	Type PatientLinkType `json:"type"`
}

// PatientGender represents the gender of a patient (R2 version)
type PatientGender string

const (
	PatientGenderMale    PatientGender = "male"
	PatientGenderFemale  PatientGender = "female"
	PatientGenderOther   PatientGender = "other"
	PatientGenderUnknown PatientGender = "unknown"
)

// PatientLinkType represents the type of link between patient resources (R2 version)
type PatientLinkType string

const (
	PatientLinkTypeReplace  PatientLinkType = "replace"  // R2 uses "replace" instead of "replaced-by"
	PatientLinkTypeReplaces PatientLinkType = "replaces"
	PatientLinkTypeRefer    PatientLinkType = "refer"
	PatientLinkTypeSeeAlso  PatientLinkType = "seealso"
)

// Bundle represents a container for a collection of resources (R2 version)
type Bundle struct {
	Resource

	// An entry in a bundle resource - will either contain a resource, or information about a resource
	Entry []BundleEntry `json:"entry,omitempty"`

	// A series of links that provide context to this bundle
	Link []BundleLink `json:"link,omitempty"`

	// Digital Signature - R2 version
	Signature *Signature `json:"signature,omitempty"`

	// Only used if the bundle is a search result set
	Total *int `json:"total,omitempty"`

	// document | message | transaction | transaction-response | batch | batch-response | history | searchset | collection
	Type BundleType `json:"type"`
}

// BundleEntry represents an entry in a bundle resource (R2 version)
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

// BundleLink represents a series of links that provide context to this bundle (R2 version)
type BundleLink struct {
	common.BackboneElement

	// See http://www.iana.org/assignments/link-relations/link-relations.xhtml#link-relations-1
	Relation string `json:"relation"`

	// Reference details for the link
	URL string `json:"url"`
}

// BundleEntryRequest represents additional execution information (R2 version)
type BundleEntryRequest struct {
	common.BackboneElement

	// For managing update contention
	IfMatch *string `json:"ifMatch,omitempty"`

	// For managing update contention
	IfModifiedSince *string `json:"ifModifiedSince,omitempty"` // R2 uses string, not time

	// For conditional creates
	IfNoneExist *string `json:"ifNoneExist,omitempty"`

	// For conditional read
	IfNoneMatch *string `json:"ifNoneMatch,omitempty"`

	// GET | POST | PUT | DELETE
	Method BundleEntryRequestMethod `json:"method"`

	// URL for HTTP equivalent of this entry
	URL string `json:"url"`
}

// BundleEntryResponse represents transaction related information (R2 version)
type BundleEntryResponse struct {
	common.BackboneElement

	// The Etag for the resource
	Etag *string `json:"etag,omitempty"`

	// Server's date time modified
	LastModified *string `json:"lastModified,omitempty"` // R2 uses string, not time

	// The location (if the operation returns a location)
	Location *string `json:"location,omitempty"`

	// Status response code
	Status string `json:"status"`
}

// BundleEntrySearch represents search related information (R2 version)
type BundleEntrySearch struct {
	common.BackboneElement

	// match | include | outcome - why this is in the result set
	Mode *BundleEntrySearchMode `json:"mode,omitempty"`

	// Search ranking
	Score *float64 `json:"score,omitempty"`
}

// Bundle-related enums (R2 version)
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

// Organization represents a formally or informally recognized grouping of people or organizations (R2 version)
type Organization struct {
	DomainResource

	// Whether the organization's record is still in active use
	Active *bool `json:"active,omitempty"`

	// Addresses for the organization
	Address []Address `json:"address,omitempty"`

	// A list of contact details for the organization
	Contact []OrganizationContact `json:"contact,omitempty"`

	// Identifies this organization across multiple systems
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Name used for the organization
	Name *string `json:"name,omitempty"`

	// The organization of which this organization forms a part
	PartOf *common.Reference `json:"partOf,omitempty"`

	// A contact detail for the organization
	Telecom []ContactPoint `json:"telecom,omitempty"`

	// Kind of organization
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// OrganizationContact represents contact information for the organization (R2 version)
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

// Practitioner represents a person who is directly or indirectly involved in the provisioning of healthcare (R2 version)
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
	Name *HumanName `json:"name,omitempty"` // R2 uses single HumanName, not array

	// Image of the person
	Photo []Attachment `json:"photo,omitempty"`

	// Roles which this practitioner may perform
	PractitionerRole []PractitionerRole `json:"practitionerRole,omitempty"`

	// Qualifications obtained by training and certification
	Qualification []PractitionerQualification `json:"qualification,omitempty"`

	// A contact detail for the practitioner
	Telecom []ContactPoint `json:"telecom,omitempty"`
}

// PractitionerRole represents roles which this practitioner may perform (R2 version)
type PractitionerRole struct {
	common.BackboneElement

	// The organization where the Practitioner performs the roles associated
	ManagingOrganization *common.Reference `json:"managingOrganization,omitempty"`

	// The period during which the person is authorized to act as a practitioner in these role(s)
	Period *common.Period `json:"period,omitempty"`

	// Roles which this practitioner is authorized to perform for the organization
	Role *common.CodeableConcept `json:"role,omitempty"`

	// Specific specialty of the practitioner
	Specialty []common.CodeableConcept `json:"specialty,omitempty"`
}

// PractitionerQualification represents qualifications obtained by training and certification (R2 version)
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

// PractitionerGender represents the gender of a practitioner (R2 version)
type PractitionerGender string

const (
	PractitionerGenderMale    PractitionerGender = "male"
	PractitionerGenderFemale  PractitionerGender = "female"
	PractitionerGenderOther   PractitionerGender = "other"
	PractitionerGenderUnknown PractitionerGender = "unknown"
)

// Encounter represents an interaction between a patient and healthcare provider(s) (R2 version)
type Encounter struct {
	DomainResource

	// The appointment that scheduled this encounter
	Appointment *common.Reference `json:"appointment,omitempty"` // R2 uses single reference, not array

	// inpatient | outpatient | ambulatory | emergency +
	Class *EncounterClass `json:"class,omitempty"` // R2 uses enum, not Coding

	// Where the encounter took place
	Hospitalization *EncounterHospitalization `json:"hospitalization,omitempty"`

	// Identifier(s) by which this encounter is known
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Incoming Referral Request
	IncomingReferral []common.Reference `json:"incomingReferral,omitempty"`

	// Quantity of time the encounter lasted
	Length *Duration `json:"length,omitempty"`

	// List of locations where the patient has been
	Location []EncounterLocation `json:"location,omitempty"`

	// Another Encounter this encounter is part of
	PartOf *common.Reference `json:"partOf,omitempty"`

	// List of participants involved in the encounter
	Participant []EncounterParticipant `json:"participant,omitempty"`

	// The start and end time of the encounter
	Period *common.Period `json:"period,omitempty"`

	// Indicates the urgency of the encounter
	Priority *common.CodeableConcept `json:"priority,omitempty"`

	// Reason the encounter takes place
	Reason []common.CodeableConcept `json:"reason,omitempty"`

	// The custodian organization of this Encounter record
	ServiceProvider *common.Reference `json:"serviceProvider,omitempty"`

	// planned | arrived | in-progress | onleave | finished | cancelled
	Status EncounterStatus `json:"status"`

	// List of past encounter statuses
	StatusHistory []EncounterStatusHistory `json:"statusHistory,omitempty"`

	// The patient present at the encounter
	Patient *common.Reference `json:"patient,omitempty"` // R2 uses "patient" instead of "subject"

	// Specific type of encounter
	Type []common.CodeableConcept `json:"type,omitempty"`
}

// EncounterHospitalization represents details about the admission to a healthcare service (R2 version)
type EncounterHospitalization struct {
	common.BackboneElement

	// From where patient was admitted (physician referral, transfer)
	AdmitSource *common.CodeableConcept `json:"admitSource,omitempty"`

	// Diet preferences reported by the patient
	DietPreference []common.CodeableConcept `json:"dietPreference,omitempty"`

	// Location to which the patient is discharged
	Destination *common.Reference `json:"destination,omitempty"`

	// Category or kind of location after discharge
	DischargeDisposition *common.CodeableConcept `json:"dischargeDisposition,omitempty"`

	// The origin/source
	Origin *common.Reference `json:"origin,omitempty"`

	// Pre-admission identifier
	PreAdmissionIdentifier *common.Identifier `json:"preAdmissionIdentifier,omitempty"`

	// The type of hospital re-admission that has occurred
	ReAdmission *common.CodeableConcept `json:"reAdmission,omitempty"`

	// Special courtesies (VIP, board member)
	SpecialArrangement []common.CodeableConcept `json:"specialArrangement,omitempty"`

	// Wheelchair, translator, stretcher, etc.
	SpecialCourtesy []common.CodeableConcept `json:"specialCourtesy,omitempty"`
}

// EncounterLocation represents list of locations where the patient has been (R2 version)
type EncounterLocation struct {
	common.BackboneElement

	// The location where the encounter takes place
	Location *common.Reference `json:"location"`

	// Time period during which the patient was present at the location
	Period *common.Period `json:"period,omitempty"`

	// planned | active | reserved | completed
	Status *EncounterLocationStatus `json:"status,omitempty"`
}

// EncounterParticipant represents list of participants involved in the encounter (R2 version)
type EncounterParticipant struct {
	common.BackboneElement

	// Persons involved in the encounter other than the patient
	Individual *common.Reference `json:"individual,omitempty"`

	// Period of time during the encounter that the participant participated
	Period *common.Period `json:"period,omitempty"`

	// Role of participant in encounter
	Type []common.CodeableConcept `json:"type,omitempty"`
}

// EncounterStatusHistory represents list of past encounter statuses (R2 version)
type EncounterStatusHistory struct {
	common.BackboneElement

	// The time that the episode was in the specified status
	Period *common.Period `json:"period"`

	// planned | arrived | in-progress | onleave | finished | cancelled
	Status EncounterStatus `json:"status"`
}

// Encounter-related enums (R2 version)
type EncounterClass string

const (
	EncounterClassInpatient   EncounterClass = "inpatient"
	EncounterClassOutpatient  EncounterClass = "outpatient"
	EncounterClassAmbulatory  EncounterClass = "ambulatory"
	EncounterClassEmergency   EncounterClass = "emergency"
	EncounterClassHome        EncounterClass = "home"
	EncounterClassField       EncounterClass = "field"
	EncounterClassDaytime     EncounterClass = "daytime"
	EncounterClassVirtual     EncounterClass = "virtual"
)

type EncounterStatus string

const (
	EncounterStatusPlanned    EncounterStatus = "planned"
	EncounterStatusArrived    EncounterStatus = "arrived"
	EncounterStatusInProgress EncounterStatus = "in-progress"
	EncounterStatusOnLeave    EncounterStatus = "onleave"
	EncounterStatusFinished   EncounterStatus = "finished"
	EncounterStatusCancelled  EncounterStatus = "cancelled"
)

type EncounterLocationStatus string

const (
	EncounterLocationStatusPlanned   EncounterLocationStatus = "planned"
	EncounterLocationStatusActive    EncounterLocationStatus = "active"
	EncounterLocationStatusReserved  EncounterLocationStatus = "reserved"
	EncounterLocationStatusCompleted EncounterLocationStatus = "completed"
)

// Condition represents a clinical condition, problem, diagnosis, or other event (R2 version)
type Condition struct {
	DomainResource

	// Estimated or actual date or date-time the condition was resolved
	AbatementDateTime *string        `json:"abatementDateTime,omitempty"` // R2 uses string
	AbatementQuantity *common.Quantity `json:"abatementQuantity,omitempty"` // R2 uses Quantity, not Age
	AbatementBoolean  *bool          `json:"abatementBoolean,omitempty"`
	AbatementPeriod   *common.Period `json:"abatementPeriod,omitempty"`
	AbatementRange    *Range         `json:"abatementRange,omitempty"`
	AbatementString   *string        `json:"abatementString,omitempty"`

	// Person who asserts this condition
	Asserter *common.Reference `json:"asserter,omitempty"`

	// Anatomical location, if relevant
	BodySite []common.CodeableConcept `json:"bodySite,omitempty"`

	// The category of the condition
	Category *common.CodeableConcept `json:"category,omitempty"` // R2 uses single, not array

	// active | relapse | remission | resolved
	ClinicalStatus *ConditionClinicalStatus `json:"clinicalStatus,omitempty"` // R2 has different values

	// Identification of the condition, problem or diagnosis
	Code *common.CodeableConcept `json:"code"`

	// Encounter when condition first asserted
	Encounter *common.Reference `json:"encounter,omitempty"` // R2 uses "encounter" not "context"

	// Supporting evidence
	Evidence []ConditionEvidence `json:"evidence,omitempty"`

	// External IDs for this condition
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Additional information about the Condition
	Notes *string `json:"notes,omitempty"` // R2 uses single string, not array

	// Date record was believed accurate
	OnsetDateTime *string          `json:"onsetDateTime,omitempty"` // R2 uses string
	OnsetQuantity *common.Quantity `json:"onsetQuantity,omitempty"` // R2 uses Quantity, not Age
	OnsetPeriod   *common.Period   `json:"onsetPeriod,omitempty"`
	OnsetRange    *Range           `json:"onsetRange,omitempty"`
	OnsetString   *string          `json:"onsetString,omitempty"`

	// Who has the condition?
	Patient *common.Reference `json:"patient"` // R2 uses "patient" not "subject"

	// A subjective assessment of the severity of the condition
	Severity *common.CodeableConcept `json:"severity,omitempty"`

	// Stage/grade, usually assessed formally
	Stage *ConditionStage `json:"stage,omitempty"`

	// provisional | differential | confirmed | refuted | entered-in-error | unknown
	VerificationStatus ConditionVerificationStatus `json:"verificationStatus"` // R2 required field
}

// ConditionEvidence represents supporting evidence (R2 version)
type ConditionEvidence struct {
	common.BackboneElement

	// Manifestation/symptom
	Code []common.CodeableConcept `json:"code,omitempty"`

	// Supporting information found elsewhere
	Detail []common.Reference `json:"detail,omitempty"`
}

// ConditionStage represents stage/grade, usually assessed formally (R2 version)
type ConditionStage struct {
	common.BackboneElement

	// Simple summary (disease specific)
	Assessment []common.Reference `json:"assessment,omitempty"`

	// Kind of staging
	Summary *common.CodeableConcept `json:"summary,omitempty"`
}

// Condition-related enums (R2 version)
type ConditionClinicalStatus string

const (
	ConditionClinicalStatusActive     ConditionClinicalStatus = "active"
	ConditionClinicalStatusRelapse    ConditionClinicalStatus = "relapse"
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

// Procedure represents an action that is or was performed on a patient (R2 version)
type Procedure struct {
	DomainResource

	// Target body sites
	BodySite []common.CodeableConcept `json:"bodySite,omitempty"`

	// Classification of the procedure
	Category *common.CodeableConcept `json:"category,omitempty"`

	// Identification of the procedure
	Code *common.CodeableConcept `json:"code"`

	// Complication following the procedure
	Complication []common.CodeableConcept `json:"complication,omitempty"`

	// Device changed in procedure
	Device []ProcedureDevice `json:"device,omitempty"` // R2 different structure

	// Encounter associated with the procedure
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Instructions for follow up
	FollowUp []common.CodeableConcept `json:"followUp,omitempty"`

	// External Identifiers for this procedure
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Where the procedure happened
	Location *common.Reference `json:"location,omitempty"`

	// Additional information about the procedure
	Notes []Annotation `json:"notes,omitempty"` // R2 uses array of Annotation

	// The result of procedure
	Outcome *common.CodeableConcept `json:"outcome,omitempty"`

	// Date/Period the procedure was performed
	PerformedDateTime *string        `json:"performedDateTime,omitempty"` // R2 uses string
	PerformedPeriod   *common.Period `json:"performedPeriod,omitempty"`

	// The people who performed the procedure
	Performer []ProcedurePerformer `json:"performer,omitempty"`

	// Coded reason procedure performed
	ReasonCodeableConcept []common.CodeableConcept `json:"reasonCodeableConcept,omitempty"`

	// Condition that justifies procedure
	ReasonReference []common.Reference `json:"reasonReference,omitempty"`

	// Any report resulting from the procedure
	Report []common.Reference `json:"report,omitempty"`

	// A request for this procedure
	Request *common.Reference `json:"request,omitempty"` // R2 uses single reference

	// in-progress | aborted | completed | entered-in-error
	Status ProcedureStatus `json:"status"`

	// Who the procedure was performed on
	Subject *common.Reference `json:"subject"` // R2 uses "subject"

	// Items used during the procedure
	Used []common.Reference `json:"used,omitempty"`
}

// ProcedureDevice represents device changed in procedure (R2 version)
type ProcedureDevice struct {
	common.BackboneElement

	// Kind of change to device
	Action *common.CodeableConcept `json:"action,omitempty"`

	// Device that was changed
	Manipulated *common.Reference `json:"manipulated"`
}

// ProcedurePerformer represents the people who performed the procedure (R2 version)
type ProcedurePerformer struct {
	common.BackboneElement

	// The reference to the practitioner
	Actor *common.Reference `json:"actor,omitempty"` // R2 optional

	// The role the actor was in
	Role *common.CodeableConcept `json:"role,omitempty"`
}

// Procedure-related enums (R2 version)
type ProcedureStatus string

const (
	ProcedureStatusInProgress     ProcedureStatus = "in-progress"
	ProcedureStatusAborted        ProcedureStatus = "aborted"
	ProcedureStatusCompleted      ProcedureStatus = "completed"
	ProcedureStatusEnteredInError ProcedureStatus = "entered-in-error"
)

// Location represents details and position information for a physical place (R2 version)
type Location struct {
	DomainResource

	// Physical location
	Address *Address `json:"address,omitempty"`

	// Description of the Location
	Description *string `json:"description,omitempty"`

	// Unique code or number identifying the location to its users
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Organization responsible for provisioning and upkeep
	ManagingOrganization *common.Reference `json:"managingOrganization,omitempty"`

	// instance | kind
	Mode *LocationMode `json:"mode,omitempty"`

	// Name of the location as used by humans
	Name *string `json:"name,omitempty"`

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

// LocationPosition represents the absolute geographic location (R2 version)
type LocationPosition struct {
	common.BackboneElement

	// Altitude with WGS84 datum
	Altitude *float64 `json:"altitude,omitempty"`

	// Latitude with WGS84 datum
	Latitude float64 `json:"latitude"`

	// Longitude with WGS84 datum
	Longitude float64 `json:"longitude"`
}

// Location-related enums (R2 version)
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

// DiagnosticReport represents the findings and interpretation of diagnostic tests (R2 version)
type DiagnosticReport struct {
	DomainResource

	// Service category
	Category *common.CodeableConcept `json:"category,omitempty"`

	// Name/Code for this diagnostic report
	Code *common.CodeableConcept `json:"code"`

	// Codes for the conclusion
	CodedDiagnosis []common.CodeableConcept `json:"codedDiagnosis,omitempty"`

	// Clinical Interpretation of test results
	Conclusion *string `json:"conclusion,omitempty"`

	// Clinically relevant time/time-period for report
	EffectiveDateTime *string        `json:"effectiveDateTime,omitempty"` // R2 uses string
	EffectivePeriod   *common.Period `json:"effectivePeriod,omitempty"`

	// Health care event when test ordered
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Business identifier for report
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Key images associated with this report
	Image []DiagnosticReportImage `json:"image,omitempty"`

	// DateTime this version was released
	Issued *string `json:"issued,omitempty"` // R2 uses string

	// Responsible Diagnostic Service
	Performer *common.Reference `json:"performer"`

	// Entire report as issued
	PresentedForm []Attachment `json:"presentedForm,omitempty"`

	// Observations - simple observations
	Result []common.Reference `json:"result,omitempty"`

	// Specimens this report is based on
	Specimen []common.Reference `json:"specimen,omitempty"`

	// registered | partial | preliminary | final | amended | corrected | appended | cancelled | entered-in-error
	Status DiagnosticReportStatus `json:"status"`

	// The subject of the report - usually, but not always, the patient
	Subject *common.Reference `json:"subject"`
}

// DiagnosticReportImage represents key images associated with this report (R2 version)
type DiagnosticReportImage struct {
	common.BackboneElement

	// Comment about the image
	Comment *string `json:"comment,omitempty"`

	// Reference to the image source
	Link *common.Reference `json:"link"`
}

// DiagnosticReport-related enums (R2 version)
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
)

// Observation represents measurements and simple assertions made about a patient, device or other subject (R2 version)
type Observation struct {
	DomainResource

	// Classification of type of observation
	Category *common.CodeableConcept `json:"category,omitempty"` // R2 uses single, not array

	// Type of observation (code / type)
	Code *common.CodeableConcept `json:"code"`

	// Comments about result
	Comments *string `json:"comments,omitempty"` // R2 uses single string

	// Why the result is missing
	DataAbsentReason *common.CodeableConcept `json:"dataAbsentReason,omitempty"`

	// (Measurement) Device
	Device *common.Reference `json:"device,omitempty"`

	// Clinically relevant time/time-period for observation
	EffectiveDateTime *string        `json:"effectiveDateTime,omitempty"` // R2 uses string
	EffectivePeriod   *common.Period `json:"effectivePeriod,omitempty"`

	// Healthcare event during which this observation is made
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Business Identifier for observation
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// High, low, normal, etc.
	Interpretation *common.CodeableConcept `json:"interpretation,omitempty"`

	// Date/Time this was made available
	Issued *string `json:"issued,omitempty"` // R2 uses string

	// The observation method
	Method *common.CodeableConcept `json:"method,omitempty"`

	// Who is responsible for the observation
	Performer []common.Reference `json:"performer,omitempty"`

	// Provides guide for interpretation
	ReferenceRange []ObservationReferenceRange `json:"referenceRange,omitempty"`

	// Observations related to this observation
	Related []ObservationRelated `json:"related,omitempty"`

	// Specimen used for this observation
	Specimen *common.Reference `json:"specimen,omitempty"`

	// registered | preliminary | final | amended | cancelled | entered-in-error | unknown
	Status ObservationStatus `json:"status"`

	// Who and/or what this is about
	Subject *common.Reference `json:"subject,omitempty"`

	// Actual result
	ValueQuantity        *common.Quantity        `json:"valueQuantity,omitempty"`
	ValueCodeableConcept *common.CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueString          *string                 `json:"valueString,omitempty"`
	ValueRange           *Range                  `json:"valueRange,omitempty"`
	ValueRatio           *Ratio                  `json:"valueRatio,omitempty"`
	ValueSampledData     *SampledData            `json:"valueSampledData,omitempty"`
	ValueAttachment      *Attachment             `json:"valueAttachment,omitempty"`
	ValueTime            *string                 `json:"valueTime,omitempty"`
	ValueDateTime        *string                 `json:"valueDateTime,omitempty"` // R2 uses string
	ValuePeriod          *common.Period          `json:"valuePeriod,omitempty"`
}

// ObservationReferenceRange represents provides guide for interpretation (R2 version)
type ObservationReferenceRange struct {
	common.BackboneElement

	// Applicable age range, if relevant
	Age *Range `json:"age,omitempty"`

	// High Range, if relevant
	High *common.Quantity `json:"high,omitempty"`

	// Low Range, if relevant
	Low *common.Quantity `json:"low,omitempty"`

	// Indicates what aspects apply
	Meaning []common.CodeableConcept `json:"meaning,omitempty"`

	// Text based reference range in an observation
	Text *string `json:"text,omitempty"`
}

// ObservationRelated represents observations related to this observation (R2 version)
type ObservationRelated struct {
	common.BackboneElement

	// Reference to the related observation
	Target *common.Reference `json:"target"`

	// has-member | derived-from | sequel-to | replaces | qualified-by | interfered-by
	Type *ObservationRelatedType `json:"type,omitempty"`
}

// Observation-related enums (R2 version)
type ObservationStatus string

const (
	ObservationStatusRegistered     ObservationStatus = "registered"
	ObservationStatusPreliminary    ObservationStatus = "preliminary"
	ObservationStatusFinal          ObservationStatus = "final"
	ObservationStatusAmended        ObservationStatus = "amended"
	ObservationStatusCancelled      ObservationStatus = "cancelled"
	ObservationStatusEnteredInError ObservationStatus = "entered-in-error"
	ObservationStatusUnknown        ObservationStatus = "unknown"
)

type ObservationRelatedType string

const (
	ObservationRelatedTypeHasMember     ObservationRelatedType = "has-member"
	ObservationRelatedTypeDerivedFrom   ObservationRelatedType = "derived-from"
	ObservationRelatedTypeSequelTo      ObservationRelatedType = "sequel-to"
	ObservationRelatedTypeReplaces      ObservationRelatedType = "replaces"
	ObservationRelatedTypeQualifiedBy   ObservationRelatedType = "qualified-by"
	ObservationRelatedTypeInterferedBy  ObservationRelatedType = "interfered-by"
)

// Medication represents a medication for an R2 context
type Medication struct {
	DomainResource

	// Codes that identify this medication
	Code *common.CodeableConcept `json:"code,omitempty"`

	// True if a brand
	IsBrand *bool `json:"isBrand,omitempty"`

	// Manufacturer of the item
	Manufacturer *common.Reference `json:"manufacturer,omitempty"`

	// Administrable medication details
	Product *MedicationProduct `json:"product,omitempty"`

	// Details about packaged medications
	Package *MedicationPackage `json:"package,omitempty"`
}

// MedicationProduct represents administrable medication details (R2 version)
type MedicationProduct struct {
	common.BackboneElement

	// Dosage form
	Form *common.CodeableConcept `json:"form,omitempty"`

	// Active or inactive ingredient
	Ingredient []MedicationProductIngredient `json:"ingredient,omitempty"`

	// Identifies a particular constituent of interest in the product
	Batch []MedicationProductBatch `json:"batch,omitempty"`
}

// MedicationProductIngredient represents active or inactive ingredient (R2 version)
type MedicationProductIngredient struct {
	common.BackboneElement

	// The actual ingredient - either a substance (simple ingredient) or another medication
	Item *common.Reference `json:"item"`

	// Quantity of ingredient present
	Amount *Ratio `json:"amount,omitempty"`
}

// MedicationProductBatch represents batch identifier (R2 version)
type MedicationProductBatch struct {
	common.BackboneElement

	// When batch will expire
	ExpirationDate *string `json:"expirationDate,omitempty"` // R2 uses string

	// The assigned lot number of a batch of the specified product
	LotNumber *string `json:"lotNumber,omitempty"`
}

// MedicationPackage represents details about packaged medications (R2 version)
type MedicationPackage struct {
	common.BackboneElement

	// E.g. box, vial, blister-pack
	Container *common.CodeableConcept `json:"container,omitempty"`

	// What is in the package
	Content []MedicationPackageContent `json:"content,omitempty"`
}

// MedicationPackageContent represents what is in the package (R2 version)
type MedicationPackageContent struct {
	common.BackboneElement

	// A product in the package
	Item *common.Reference `json:"item"`

	// Quantity present in the package
	Amount *common.Quantity `json:"amount,omitempty"`
}

// MedicationOrder represents a prescription for R2 (Note: R2 uses "MedicationOrder" not "MedicationRequest")
type MedicationOrder struct {
	DomainResource

	// When prescription was initially authorized
	DateWritten *string `json:"dateWritten,omitempty"` // R2 uses string

	// Medication supply authorization
	DispenseRequest *MedicationOrderDispenseRequest `json:"dispenseRequest,omitempty"`

	// How the medication should be taken
	DosageInstruction []DosageInstruction `json:"dosageInstruction,omitempty"`

	// Created during encounter/admission/stay
	Encounter *common.Reference `json:"encounter,omitempty"`

	// External identifier
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Medication to be taken
	MedicationCodeableConcept *common.CodeableConcept `json:"medicationCodeableConcept,omitempty"`
	MedicationReference       *common.Reference       `json:"medicationReference,omitempty"`

	// Information about the prescription
	Note *string `json:"note,omitempty"` // R2 uses single string

	// Who ordered the medication(s)
	Prescriber *common.Reference `json:"prescriber,omitempty"`

	// An order/prescription that this supersedes
	PriorPrescription *common.Reference `json:"priorPrescription,omitempty"`

	// routine | urgent | stat | asap
	Priority *MedicationOrderPriority `json:"priority,omitempty"`

	// Reason or indication for writing the prescription
	ReasonCodeableConcept *common.CodeableConcept `json:"reasonCodeableConcept,omitempty"`
	ReasonReference       *common.Reference       `json:"reasonReference,omitempty"`

	// active | on-hold | completed | entered-in-error | stopped | draft
	Status *MedicationOrderStatus `json:"status,omitempty"`

	// Who prescription is for
	Patient *common.Reference `json:"patient"` // R2 uses "patient" not "subject"

	// Any restrictions on medication substitution
	Substitution *MedicationOrderSubstitution `json:"substitution,omitempty"`
}

// MedicationOrderDispenseRequest represents medication supply authorization (R2 version)
type MedicationOrderDispenseRequest struct {
	common.BackboneElement

	// Number of days supply per dispense
	ExpectedSupplyDuration *Duration `json:"expectedSupplyDuration,omitempty"`

	// Product to be supplied
	MedicationCodeableConcept *common.CodeableConcept `json:"medicationCodeableConcept,omitempty"`
	MedicationReference       *common.Reference       `json:"medicationReference,omitempty"`

	// Number of refills authorized
	NumberOfRepeatsAllowed *int `json:"numberOfRepeatsAllowed,omitempty"`

	// Amount of medication to supply per dispense
	Quantity *common.Quantity `json:"quantity,omitempty"`

	// Time period supply is authorized for
	ValidityPeriod *common.Period `json:"validityPeriod,omitempty"`
}

// MedicationOrderSubstitution represents any restrictions on medication substitution (R2 version)
type MedicationOrderSubstitution struct {
	common.BackboneElement

	// Whether substitution is allowed or not
	Type *common.CodeableConcept `json:"type"`

	// Why should substitution (not) be made
	Reason *common.CodeableConcept `json:"reason,omitempty"`
}

// MedicationOrder-related enums (R2 version)
type MedicationOrderStatus string

const (
	MedicationOrderStatusActive          MedicationOrderStatus = "active"
	MedicationOrderStatusOnHold          MedicationOrderStatus = "on-hold"
	MedicationOrderStatusCompleted       MedicationOrderStatus = "completed"
	MedicationOrderStatusEnteredInError  MedicationOrderStatus = "entered-in-error"
	MedicationOrderStatusStopped         MedicationOrderStatus = "stopped"
	MedicationOrderStatusDraft           MedicationOrderStatus = "draft"
)

type MedicationOrderPriority string

const (
	MedicationOrderPriorityRoutine MedicationOrderPriority = "routine"
	MedicationOrderPriorityUrgent  MedicationOrderPriority = "urgent"
	MedicationOrderPriorityStat    MedicationOrderPriority = "stat"
	MedicationOrderPriorityASAP    MedicationOrderPriority = "asap"
)

// QuestionnaireResponse represents a structured set of questions and their answers (R2 version)
type QuestionnaireResponse struct {
	DomainResource

	// Person who received and recorded the answers
	Author *common.Reference `json:"author,omitempty"`

	// Date the answers were gathered
	Authored *string `json:"authored,omitempty"` // R2 uses string

	// Primary encounter during which questionnaire was completed
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Grouped questions
	Group *QuestionnaireResponseGroup `json:"group,omitempty"` // R2 has different structure

	// Unique id for this set of answers
	Identifier *common.Identifier `json:"identifier,omitempty"`

	// Form being answered
	Questionnaire *common.Reference `json:"questionnaire,omitempty"`

	// Person who provided the information about the subject
	Source *common.Reference `json:"source,omitempty"`

	// in-progress | completed | amended
	Status QuestionnaireResponseStatus `json:"status"`

	// The subject of the questions
	Subject *common.Reference `json:"subject,omitempty"`
}

// QuestionnaireResponseGroup represents grouped questions (R2 version)
type QuestionnaireResponseGroup struct {
	common.BackboneElement

	// Nested questionnaire response groups
	Group []QuestionnaireResponseGroup `json:"group,omitempty"`

	// Groups can repeat in the answers, so a direct 1..1 correspondence may not exist
	LinkId *string `json:"linkId,omitempty"`

	// Questions in this group
	Question []QuestionnaireResponseGroupQuestion `json:"question,omitempty"`

	// The subject this group's answers are about
	Subject *common.Reference `json:"subject,omitempty"`

	// Additional text for the group
	Text *string `json:"text,omitempty"`

	// Name for the group
	Title *string `json:"title,omitempty"`
}

// QuestionnaireResponseGroupQuestion represents questions in a group (R2 version)
type QuestionnaireResponseGroupQuestion struct {
	common.BackboneElement

	// The response(s) to the question
	Answer []QuestionnaireResponseGroupQuestionAnswer `json:"answer,omitempty"`

	// Groups can repeat in the answers, so a direct 1..1 correspondence may not exist
	LinkId *string `json:"linkId,omitempty"`

	// Text of the question as it is shown to the user
	Text *string `json:"text,omitempty"`
}

// QuestionnaireResponseGroupQuestionAnswer represents the response(s) to the question (R2 version)
type QuestionnaireResponseGroupQuestionAnswer struct {
	common.BackboneElement

	// Nested groups and/or questions
	Group []QuestionnaireResponseGroup `json:"group,omitempty"`

	// Single-valued answer to the question
	ValueBoolean     *bool                       `json:"valueBoolean,omitempty"`
	ValueDecimal     *float64                    `json:"valueDecimal,omitempty"`
	ValueInteger     *int                        `json:"valueInteger,omitempty"`
	ValueDate        *string                     `json:"valueDate,omitempty"`
	ValueDateTime    *string                     `json:"valueDateTime,omitempty"` // R2 uses string
	ValueInstant     *string                     `json:"valueInstant,omitempty"`  // R2 has instant type
	ValueTime        *string                     `json:"valueTime,omitempty"`
	ValueString      *string                     `json:"valueString,omitempty"`
	ValueUri         *string                     `json:"valueUri,omitempty"`
	ValueAttachment  *Attachment                 `json:"valueAttachment,omitempty"`
	ValueCoding      *common.Coding              `json:"valueCoding,omitempty"`
	ValueQuantity    *common.Quantity            `json:"valueQuantity,omitempty"`
	ValueReference   *common.Reference           `json:"valueReference,omitempty"`
}

// QuestionnaireResponse-related enums (R2 version)
type QuestionnaireResponseStatus string

const (
	QuestionnaireResponseStatusInProgress QuestionnaireResponseStatus = "in-progress"
	QuestionnaireResponseStatusCompleted  QuestionnaireResponseStatus = "completed"
	QuestionnaireResponseStatusAmended    QuestionnaireResponseStatus = "amended"
)

// Communication represents an occurrence of information being transmitted (R2 version)
type Communication struct {
	DomainResource

	// Message category
	Category *common.CodeableConcept `json:"category,omitempty"` // R2 uses single, not array

	// Encounter leading to message
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Unique identifier
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// A channel of communication
	Medium []common.CodeableConcept `json:"medium,omitempty"`

	// Message payload
	Payload []CommunicationPayload `json:"payload,omitempty"`

	// Indication for message
	Reason []common.CodeableConcept `json:"reason,omitempty"`

	// When received
	Received *string `json:"received,omitempty"` // R2 uses string

	// Message recipient
	Recipient []common.Reference `json:"recipient,omitempty"`

	// An order/prescription that this supersedes
	RequestDetail *common.Reference `json:"requestDetail,omitempty"` // R2 uses this field

	// Message sender
	Sender *common.Reference `json:"sender,omitempty"`

	// When sent
	Sent *string `json:"sent,omitempty"` // R2 uses string

	// in-progress | completed | suspended | rejected | failed
	Status CommunicationStatus `json:"status"`

	// Focus of message
	Subject *common.Reference `json:"subject,omitempty"`
}

// CommunicationPayload represents message payload (R2 version)
type CommunicationPayload struct {
	common.BackboneElement

	// Message part content
	ContentString     *string           `json:"contentString,omitempty"`
	ContentAttachment *Attachment       `json:"contentAttachment,omitempty"`
	ContentReference  *common.Reference `json:"contentReference,omitempty"`
}

// Communication-related enums (R2 version)
type CommunicationStatus string

const (
	CommunicationStatusInProgress CommunicationStatus = "in-progress"
	CommunicationStatusCompleted  CommunicationStatus = "completed"
	CommunicationStatusSuspended  CommunicationStatus = "suspended"
	CommunicationStatusRejected   CommunicationStatus = "rejected"
	CommunicationStatusFailed     CommunicationStatus = "failed"
)

// Composition represents a document authored/attested by humans, organizations and devices (R2 version)
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
	Date string `json:"date"` // R2 uses string

	// Context of the Composition
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Version-independent identifier for the Composition
	Identifier *common.Identifier `json:"identifier,omitempty"`

	// The clinical service(s) being documented
	Event []CompositionEvent `json:"event,omitempty"`

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

// CompositionAttester represents attests to accuracy of composition (R2 version)
type CompositionAttester struct {
	common.BackboneElement

	// personal | professional | legal | official
	Mode []CompositionAttesterMode `json:"mode"`

	// Who attested the composition
	Party *common.Reference `json:"party,omitempty"`

	// When composition attested
	Time *string `json:"time,omitempty"` // R2 uses string
}

// CompositionEvent represents the clinical service being documented (R2 version)
type CompositionEvent struct {
	common.BackboneElement

	// Code(s) that apply to the event being documented
	Code []common.CodeableConcept `json:"code,omitempty"`

	// Full details for the event(s) the composition consents
	Detail []common.Reference `json:"detail,omitempty"`

	// The period covered by the documentation
	Period *common.Period `json:"period,omitempty"`
}

// CompositionSection represents composition is broken into sections (R2 version)
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

// Composition-related enums (R2 version)
type CompositionStatus string

const (
	CompositionStatusPreliminary     CompositionStatus = "preliminary"
	CompositionStatusFinal           CompositionStatus = "final"
	CompositionStatusAmended         CompositionStatus = "amended"
	CompositionStatusEnteredInError  CompositionStatus = "entered-in-error"
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

// DocumentReference represents a reference to a document (R2 version)
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
	Created *string `json:"created,omitempty"` // R2 uses string

	// Organization which maintains the document
	Custodian *common.Reference `json:"custodian,omitempty"`

	// Human-readable description (title)
	Description *string `json:"description,omitempty"`

	// Document identifier and version
	DocStatus *common.CodeableConcept `json:"docStatus,omitempty"` // R2 has this field

	// Master Version Specific Identifier
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// When this document reference was indexed
	Indexed string `json:"indexed"` // R2 uses string

	// Document security-tags
	SecurityLabel []common.CodeableConcept `json:"securityLabel,omitempty"`

	// current | superseded | entered-in-error
	Status DocumentReferenceStatus `json:"status"`

	// Who/what is the subject of the document
	Subject *common.Reference `json:"subject,omitempty"`

	// Kind of document (LOINC if possible)
	Type *common.CodeableConcept `json:"type"`
}

// DocumentReferenceContent represents document referenced (R2 version)
type DocumentReferenceContent struct {
	common.BackboneElement

	// Where to access the document
	Attachment *Attachment `json:"attachment"`

	// Format/content rules for the document
	Format []common.Coding `json:"format,omitempty"` // R2 uses array
}

// DocumentReferenceContext represents clinical context of document (R2 version)
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

	// Related things
	Related []DocumentReferenceContextRelated `json:"related,omitempty"`

	// Patient demographics from source
	SourcePatientInfo *common.Reference `json:"sourcePatientInfo,omitempty"`
}

// DocumentReferenceContextRelated represents related things (R2 version)
type DocumentReferenceContextRelated struct {
	common.BackboneElement

	// Identifier of related objects or events
	Identifier *common.Identifier `json:"identifier,omitempty"`

	// Related Resource
	Ref *common.Reference `json:"ref,omitempty"`
}

// DocumentReference-related enums (R2 version)
type DocumentReferenceStatus string

const (
	DocumentReferenceStatusCurrent       DocumentReferenceStatus = "current"
	DocumentReferenceStatusSuperseded    DocumentReferenceStatus = "superseded"
	DocumentReferenceStatusEnteredInError DocumentReferenceStatus = "entered-in-error"
)

// Device represents an instance of a medical device used in the provision of patient care (R2 version)
type Device struct {
	DomainResource

	// Details for human/organization for support
	Contact []ContactPoint `json:"contact,omitempty"`

	// Date and time of expiry of this device (if applicable)
	Expiry *string `json:"expiry,omitempty"` // R2 uses string

	// Instance identifier
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Where the resource is found
	Location *common.Reference `json:"location,omitempty"`

	// Lot number of manufacture
	LotNumber *string `json:"lotNumber,omitempty"`

	// Date when the device was made
	ManufactureDate *string `json:"manufactureDate,omitempty"` // R2 uses string

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

	// available | not-available | entered-in-error
	Status *DeviceStatus `json:"status,omitempty"`

	// The kind or type of device
	Type *common.CodeableConcept `json:"type"`

	// FDA mandated Unique Device Identifier
	Udi *string `json:"udi,omitempty"` // R2 uses simple string

	// Network address to contact device
	URL *string `json:"url,omitempty"`

	// Version number (i.e. software)
	Version *string `json:"version,omitempty"`
}

// Device-related enums (R2 version)
type DeviceStatus string

const (
	DeviceStatusAvailable      DeviceStatus = "available"
	DeviceStatusNotAvailable   DeviceStatus = "not-available"
	DeviceStatusEnteredInError DeviceStatus = "entered-in-error"
)

// DosageInstruction represents how medication should be taken (R2 version)
type DosageInstruction struct {
	common.BackboneElement

	// Supplemental instructions
	AdditionalInstructions *common.CodeableConcept `json:"additionalInstructions,omitempty"`

	// Take "as needed" (for x)
	AsNeededBoolean        *bool                    `json:"asNeededBoolean,omitempty"`
	AsNeededCodeableConcept *common.CodeableConcept `json:"asNeededCodeableConcept,omitempty"`

	// Amount of medication per dose
	DoseQuantity *common.Quantity `json:"doseQuantity,omitempty"`
	DoseRange    *Range           `json:"doseRange,omitempty"`

	// Upper limit on medication per administration
	MaxDosePerPeriod *Ratio `json:"maxDosePerPeriod,omitempty"`

	// Technique for administering medication
	Method *common.CodeableConcept `json:"method,omitempty"`

	// Amount of medication per unit of time
	Rate *Ratio `json:"rate,omitempty"` // R2 uses simple Ratio

	// How drug should enter body
	Route *common.CodeableConcept `json:"route,omitempty"`

	// Body site to administer to
	SiteCodeableConcept *common.CodeableConcept `json:"siteCodeableConcept,omitempty"`
	SiteReference       *common.Reference       `json:"siteReference,omitempty"`

	// Free text dosage instructions e.g. SIG
	Text *string `json:"text,omitempty"`

	// When medication should be administered
	Timing *Timing `json:"timing,omitempty"`
}

// Timing represents a timing schedule that specifies an event that may occur multiple times (R2 version)
type Timing struct {
	common.Element

	// When the event occurs
	Code *common.CodeableConcept `json:"code,omitempty"`

	// When the event is to occur
	Event []string `json:"event,omitempty"` // R2 uses string array

	// When the event is to occur
	Repeat *TimingRepeat `json:"repeat,omitempty"`
}

// TimingRepeat represents when the event is to occur (R2 version)
type TimingRepeat struct {
	common.Element

	// Length/Range of lengths, or (Start and/or end) limits
	BoundsPeriod   *common.Period   `json:"boundsPeriod,omitempty"`
	BoundsQuantity *common.Quantity `json:"boundsQuantity,omitempty"` // R2 uses Quantity for duration

	// Number of times to repeat
	Count *int `json:"count,omitempty"`

	// How long when it happens
	Duration *float64 `json:"duration,omitempty"`

	// s | min | h | d | wk | mo | a - unit of time (UCUM)
	DurationUnits *TimingRepeatDurationUnits `json:"durationUnits,omitempty"`

	// Event occurs frequency times per period
	Frequency *int `json:"frequency,omitempty"`

	// Event occurs up to frequencyMax times per period
	FrequencyMax *int `json:"frequencyMax,omitempty"`

	// Event occurs frequency times per period
	Period *float64 `json:"period,omitempty"`

	// s | min | h | d | wk | mo | a - unit of time (UCUM)
	PeriodUnits *TimingRepeatPeriodUnits `json:"periodUnits,omitempty"`

	// Regular life events the event is tied to
	When *TimingRepeatWhen `json:"when,omitempty"` // R2 uses single value
}

// Timing-related enums (R2 version)
type TimingRepeatDurationUnits string

const (
	TimingRepeatDurationUnitsS   TimingRepeatDurationUnits = "s"
	TimingRepeatDurationUnitsMin TimingRepeatDurationUnits = "min"
	TimingRepeatDurationUnitsH   TimingRepeatDurationUnits = "h"
	TimingRepeatDurationUnitsD   TimingRepeatDurationUnits = "d"
	TimingRepeatDurationUnitsWk  TimingRepeatDurationUnits = "wk"
	TimingRepeatDurationUnitsMo  TimingRepeatDurationUnits = "mo"
	TimingRepeatDurationUnitsA   TimingRepeatDurationUnits = "a"
)

type TimingRepeatPeriodUnits string

const (
	TimingRepeatPeriodUnitsS   TimingRepeatPeriodUnits = "s"
	TimingRepeatPeriodUnitsMin TimingRepeatPeriodUnits = "min"
	TimingRepeatPeriodUnitsH   TimingRepeatPeriodUnits = "h"
	TimingRepeatPeriodUnitsD   TimingRepeatPeriodUnits = "d"
	TimingRepeatPeriodUnitsWk  TimingRepeatPeriodUnits = "wk"
	TimingRepeatPeriodUnitsMo  TimingRepeatPeriodUnits = "mo"
	TimingRepeatPeriodUnitsA   TimingRepeatPeriodUnits = "a"
)

type TimingRepeatWhen string

const (
	TimingRepeatWhenHS  TimingRepeatWhen = "HS"
	TimingRepeatWhenWAKE TimingRepeatWhen = "WAKE"
	TimingRepeatWhenAC  TimingRepeatWhen = "AC"
	TimingRepeatWhenACD TimingRepeatWhen = "ACD"
	TimingRepeatWhenACM TimingRepeatWhen = "ACM"
	TimingRepeatWhenACV TimingRepeatWhen = "ACV"
	TimingRepeatWhenC   TimingRepeatWhen = "C"
	TimingRepeatWhenCD  TimingRepeatWhen = "CD"
	TimingRepeatWhenCM  TimingRepeatWhen = "CM"
	TimingRepeatWhenCV  TimingRepeatWhen = "CV"
	TimingRepeatWhenPC  TimingRepeatWhen = "PC"
	TimingRepeatWhenPCD TimingRepeatWhen = "PCD"
	TimingRepeatWhenPCM TimingRepeatWhen = "PCM"
	TimingRepeatWhenPCV TimingRepeatWhen = "PCV"
)

// SampledData represents a series of measurements taken by a device (R2 version)
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

// Narrative represents a human-readable formatted text, including images (R2 version)
type Narrative struct {
	common.Element

	// generated | extensions | additional | empty
	Status NarrativeStatus `json:"status"`

	// Limited xhtml content
	Div string `json:"div"`
}

// Narrative-related enums (R2 version)
type NarrativeStatus string

const (
	NarrativeStatusGenerated  NarrativeStatus = "generated"
	NarrativeStatusExtensions NarrativeStatus = "extensions"
	NarrativeStatusAdditional NarrativeStatus = "additional"
	NarrativeStatusEmpty      NarrativeStatus = "empty"
)

// Signature represents a digital signature along with supporting context (R2 version)
type Signature struct {
	common.Element

	// The actual signature content (XML DigSig, JWT, picture, etc.)
	Blob *string `json:"blob,omitempty"`

	// The technical format of the signature
	ContentType string `json:"contentType"`

	// Indication of the reason the entity signed the object(s)
	Type []common.Coding `json:"type"`

	// Who signed
	WhoURI       *string           `json:"whoUri,omitempty"`
	WhoReference *common.Reference `json:"whoReference,omitempty"`

	// When the signature was created
	When string `json:"when"` // R2 uses string
}

// AllergyIntolerance represents risk of harmful or undesirable, physiological response to a substance (R2)
type AllergyIntolerance struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "AllergyIntolerance"

	// External ids for this item
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Date record was believed accurate
	RecordedDate *string `json:"recordedDate,omitempty"` // R2 uses string

	// Who recorded the sensitivity
	Recorder *common.Reference `json:"recorder,omitempty"`

	// Who the sensitivity is for
	Patient common.Reference `json:"patient"`

	// Substance, agent, or class considered to be responsible for risk
	Substance *common.CodeableConcept `json:"substance"`

	// active | unconfirmed | confirmed | inactive | resolved | refuted | entered-in-error
	Status AllergyIntoleranceStatus `json:"status"`

	// CRITL | CRITH | CRITU (R2 has simpler criticality)
	Criticality *AllergyIntoleranceCriticality `json:"criticality,omitempty"`

	// allergy | intolerance | adverse reaction (R2 categories)
	Type AllergyIntoleranceType `json:"type"`

	// food | medication | environment | other (R2 simpler categories)
	Category AllergyIntoleranceCategory `json:"category"`

	// Date(/time) of last known occurrence of a reaction
	LastOccurenceDate *string `json:"lastOccurenceDate,omitempty"` // R2 uses string

	// Additional information about the Allergy or Intolerance
	Note *Annotation `json:"note,omitempty"` // R2 uses single Annotation

	// Adverse Reaction Events linked to exposure to substance
	Reaction []AllergyIntoleranceReaction `json:"reaction,omitempty"`
}

// AllergyIntoleranceStatus represents the status (R2 specific values)
type AllergyIntoleranceStatus string

const (
	AllergyIntoleranceStatusActive       AllergyIntoleranceStatus = "active"
	AllergyIntoleranceStatusUnconfirmed  AllergyIntoleranceStatus = "unconfirmed"
	AllergyIntoleranceStatusConfirmed    AllergyIntoleranceStatus = "confirmed"
	AllergyIntoleranceStatusInactive     AllergyIntoleranceStatus = "inactive"
	AllergyIntoleranceStatusResolved     AllergyIntoleranceStatus = "resolved"
	AllergyIntoleranceStatusRefuted      AllergyIntoleranceStatus = "refuted"
	AllergyIntoleranceStatusEnteredInError AllergyIntoleranceStatus = "entered-in-error"
)

// AllergyIntoleranceCriticality represents the potential clinical harm (R2 simpler)
type AllergyIntoleranceCriticality string

const (
	AllergyIntoleranceCriticalityLow      AllergyIntoleranceCriticality = "CRITL"
	AllergyIntoleranceCriticalityHigh     AllergyIntoleranceCriticality = "CRITH"
	AllergyIntoleranceCriticalityUnknown  AllergyIntoleranceCriticality = "CRITU"
)

// AllergyIntoleranceType represents the type (R2 includes adverse reaction)
type AllergyIntoleranceType string

const (
	AllergyIntoleranceTypeAllergy        AllergyIntoleranceType = "allergy"
	AllergyIntoleranceTypeIntolerance    AllergyIntoleranceType = "intolerance"
	AllergyIntoleranceTypeAdverseReaction AllergyIntoleranceType = "adverse-reaction"
)

// AllergyIntoleranceCategory represents the category (R2 simpler)
type AllergyIntoleranceCategory string

const (
	AllergyIntoleranceCategoryFood        AllergyIntoleranceCategory = "food"
	AllergyIntoleranceCategoryMedication  AllergyIntoleranceCategory = "medication"
	AllergyIntoleranceCategoryEnvironment AllergyIntoleranceCategory = "environment"
	AllergyIntoleranceCategoryOther       AllergyIntoleranceCategory = "other"
)

// AllergyIntoleranceReaction represents details about each adverse reaction event (R2)
type AllergyIntoleranceReaction struct {
	common.BackboneElement

	// Specific substance considered to be responsible for event
	Substance *common.CodeableConcept `json:"substance,omitempty"`

	// Clinical symptoms/signs associated with the Event
	Manifestation []common.CodeableConcept `json:"manifestation"`

	// Description of the event as a whole
	Description *string `json:"description,omitempty"`

	// Date(/time) when manifestations showed
	Onset *string `json:"onset,omitempty"` // R2 uses string

	// mild | moderate | severe (of event as a whole)
	Severity *AllergyIntoleranceReactionSeverity `json:"severity,omitempty"`

	// How the subject was exposed to the substance
	ExposureRoute *common.CodeableConcept `json:"exposureRoute,omitempty"`

	// Text about event not captured in other fields
	Note *Annotation `json:"note,omitempty"` // R2 uses single Annotation
}

// AllergyIntoleranceReactionSeverity represents the clinical assessment of severity
type AllergyIntoleranceReactionSeverity string

const (
	AllergyIntoleranceReactionSeverityMild     AllergyIntoleranceReactionSeverity = "mild"
	AllergyIntoleranceReactionSeverityModerate AllergyIntoleranceReactionSeverity = "moderate"
	AllergyIntoleranceReactionSeveritySevere   AllergyIntoleranceReactionSeverity = "severe"
)

// Immunization represents the event of a patient being administered a vaccine (R2)
type Immunization struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Immunization"

	// Business identifier
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Vaccine administration date
	Date string `json:"date"` // R2 uses string and is required

	// Vaccine that was administered or was to be administered
	VaccineCode common.CodeableConcept `json:"vaccineCode"`

	// Who was immunized
	Patient common.Reference `json:"patient"`

	// Flag for whether immunization was given
	WasNotGiven bool `json:"wasNotGiven"`

	// True if this administration was reported rather than directly administered
	Reported bool `json:"reported"`

	// Who performed the vaccination
	Performer *common.Reference `json:"performer,omitempty"`

	// Who ordered the vaccination
	Requester *common.Reference `json:"requester,omitempty"`

	// Encounter administered as part of
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Vaccine manufacturer
	Manufacturer *common.Reference `json:"manufacturer,omitempty"`

	// Where vaccination occurred
	Location *common.Reference `json:"location,omitempty"`

	// Vaccine lot number
	LotNumber *string `json:"lotNumber,omitempty"`

	// Vaccine expiration date
	ExpirationDate *string `json:"expirationDate,omitempty"` // R2 uses string

	// Body site vaccine was administered
	Site *common.CodeableConcept `json:"site,omitempty"`

	// How vaccine entered body
	Route *common.CodeableConcept `json:"route,omitempty"`

	// Amount of vaccine administered
	DoseQuantity *common.Quantity `json:"doseQuantity,omitempty"`

	// Vaccination notes
	Note []Annotation `json:"note,omitempty"`

	// Why immunization occurred
	Explanation *ImmunizationExplanation `json:"explanation,omitempty"`

	// Details of a reaction that follows immunization
	Reaction []ImmunizationReaction `json:"reaction,omitempty"`

	// Protocol followed by the provider
	VaccinationProtocol []ImmunizationVaccinationProtocol `json:"vaccinationProtocol,omitempty"`
}

// ImmunizationExplanation represents why immunization occurred (R2)
type ImmunizationExplanation struct {
	common.BackboneElement

	// Why immunization occurred
	Reason []common.CodeableConcept `json:"reason,omitempty"`

	// Why immunization did not occur
	ReasonNotGiven []common.CodeableConcept `json:"reasonNotGiven,omitempty"`
}

// ImmunizationReaction represents details of a reaction following immunization (R2)
type ImmunizationReaction struct {
	common.BackboneElement

	// When reaction started
	Date *string `json:"date,omitempty"` // R2 uses string

	// Additional information on reaction
	Detail *common.Reference `json:"detail,omitempty"`

	// Indicates self-reported reaction
	Reported *bool `json:"reported,omitempty"`
}

// ImmunizationVaccinationProtocol represents protocol followed by the provider (R2)
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

	// Disease immunized against
	DoseTarget []common.CodeableConcept `json:"doseTarget"` // R2 uses doseTarget instead

	// Does dose count towards immunity?
	DoseStatus common.CodeableConcept `json:"doseStatus"`

	// Why dose does (not) count
	DoseStatusReason *common.CodeableConcept `json:"doseStatusReason,omitempty"`
}

// CarePlan represents describes the intention of how one or more practitioners intend to deliver care (R2)
type CarePlan struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "CarePlan"

	// External Ids for this plan
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Who care plan is for
	Patient *common.Reference `json:"patient,omitempty"` // R2 uses patient

	// planned | active | completed | cancelled
	Status CarePlanStatus `json:"status"`

	// Time period plan covers
	Period *common.Period `json:"period,omitempty"`

	// When last updated
	Modified *string `json:"modified,omitempty"` // R2 uses string

	// Who is responsible for contents of the care plan
	Author []common.Reference `json:"author,omitempty"`

	// Who's involved in plan
	Participant []CarePlanParticipant `json:"participant,omitempty"`

	// Goals that this plan address
	Goal []CarePlanGoal `json:"goal,omitempty"`

	// Action to occur as part of plan
	Activity []CarePlanActivity `json:"activity,omitempty"`

	// Comments about the plan
	Note []Annotation `json:"note,omitempty"`
}

// CarePlanStatus represents the workflow state of the care plan (R2)
type CarePlanStatus string

const (
	CarePlanStatusPlanned   CarePlanStatus = "planned"
	CarePlanStatusActive    CarePlanStatus = "active"
	CarePlanStatusCompleted CarePlanStatus = "completed"
	CarePlanStatusCancelled CarePlanStatus = "cancelled"
)

// CarePlanParticipant represents identifies all people and organizations who are expected to be involved in the care (R2)
type CarePlanParticipant struct {
	common.BackboneElement

	// Type of involvement
	Role *common.CodeableConcept `json:"role,omitempty"`

	// Who is involved
	Member *common.Reference `json:"member,omitempty"`
}

// CarePlanGoal represents describes the intended objective(s) for the patient (R2)
type CarePlanGoal struct {
	common.BackboneElement

	// Unique identifier
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Human-readable description of a specific desired objective
	Description string `json:"description"`

	// in-progress | achieved | sustaining | cancelled | accepted | rejected
	Status *CarePlanGoalStatus `json:"status,omitempty"`

	// Comments about the goal
	Note []Annotation `json:"note,omitempty"`

	// Health issues this goal addresses
	Concern []common.Reference `json:"concern,omitempty"`
}

// CarePlanGoalStatus represents the status (R2)
type CarePlanGoalStatus string

const (
	CarePlanGoalStatusInProgress  CarePlanGoalStatus = "in-progress"
	CarePlanGoalStatusAchieved    CarePlanGoalStatus = "achieved"
	CarePlanGoalStatusSustaining  CarePlanGoalStatus = "sustaining"
	CarePlanGoalStatusCancelled   CarePlanGoalStatus = "cancelled"
	CarePlanGoalStatusAccepted    CarePlanGoalStatus = "accepted"
	CarePlanGoalStatusRejected    CarePlanGoalStatus = "rejected"
)

// CarePlanActivity represents identifies a planned action to occur as part of the care plan (R2)
type CarePlanActivity struct {
	common.BackboneElement

	// Unique identifier
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Goals that this activity relates to
	Goal []string `json:"goal,omitempty"` // R2 uses string array, not references

	// in-progress | completed | aborted | cancelled
	Status *CarePlanActivityStatus `json:"status,omitempty"`

	// Do NOT do
	Prohibited bool `json:"prohibited"`

	// Appointments, orders, etc.
	ActionResulting []common.Reference `json:"actionResulting,omitempty"`

	// Comments about the activity
	Note []Annotation `json:"note,omitempty"`

	// Activity details defined in specific resource
	Reference *common.Reference `json:"reference,omitempty"`

	// In-line definition of activity
	Detail *CarePlanActivityDetail `json:"detail,omitempty"`
}

// CarePlanActivityStatus represents the status (R2)
type CarePlanActivityStatus string

const (
	CarePlanActivityStatusNotStarted   CarePlanActivityStatus = "not-started"
	CarePlanActivityStatusScheduled    CarePlanActivityStatus = "scheduled"
	CarePlanActivityStatusInProgress   CarePlanActivityStatus = "in-progress"
	CarePlanActivityStatusOnHold       CarePlanActivityStatus = "on-hold"
	CarePlanActivityStatusCompleted    CarePlanActivityStatus = "completed"
	CarePlanActivityStatusCancelled    CarePlanActivityStatus = "cancelled"
)

// CarePlanActivityDetail represents a simple summary of a planned activity (R2)
type CarePlanActivityDetail struct {
	common.BackboneElement

	// diet | drug | encounter | observation | procedure | supply | other
	Category CarePlanActivityCategory `json:"category"`

	// Detail type of activity
	Code *common.CodeableConcept `json:"code,omitempty"`

	// Why activity should be done
	ReasonCode []common.CodeableConcept `json:"reasonCode,omitempty"`

	// Condition triggering need for activity
	ReasonReference []common.Reference `json:"reasonReference,omitempty"`

	// Goals this activity relates to
	Goal []common.Reference `json:"goal,omitempty"`

	// not-started | scheduled | in-progress | on-hold | completed | cancelled | stopped
	Status *CarePlanActivityStatus `json:"status,omitempty"`

	// Reason for current status
	StatusReason *common.CodeableConcept `json:"statusReason,omitempty"`

	// Do NOT do
	Prohibited bool `json:"prohibited"`

	// When activity is to occur
	ScheduledTiming  *Timing        `json:"scheduledTiming,omitempty"`
	ScheduledPeriod  *common.Period `json:"scheduledPeriod,omitempty"`
	ScheduledString  *string        `json:"scheduledString,omitempty"`

	// Where it should happen
	Location *common.Reference `json:"location,omitempty"`

	// Who will be responsible?
	Performer []common.Reference `json:"performer,omitempty"`

	// What is administered/supplied
	Product *common.Reference `json:"product,omitempty"` // R2 uses single reference

	// How to consume/day?
	DailyAmount *common.Quantity `json:"dailyAmount,omitempty"`

	// How much to administer/supply/consume
	Quantity *common.Quantity `json:"quantity,omitempty"`

	// Extra info describing activity to perform
	Description *string `json:"description,omitempty"`
}

// CarePlanActivityCategory represents the category (R2)
type CarePlanActivityCategory string

const (
	CarePlanActivityCategoryDiet        CarePlanActivityCategory = "diet"
	CarePlanActivityCategoryDrug        CarePlanActivityCategory = "drug"
	CarePlanActivityCategoryEncounter   CarePlanActivityCategory = "encounter"
	CarePlanActivityCategoryObservation CarePlanActivityCategory = "observation"
	CarePlanActivityCategoryProcedure   CarePlanActivityCategory = "procedure"
	CarePlanActivityCategorySupply      CarePlanActivityCategory = "supply"
	CarePlanActivityCategoryOther       CarePlanActivityCategory = "other"
)

// Specimen represents a sample to be used for analysis (R2)
type Specimen struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Specimen"

	// External Identifier
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Kind of material that forms the specimen
	Type *common.CodeableConcept `json:"type,omitempty"`

	// Specimen from which this specimen originated
	Parent []common.Reference `json:"parent,omitempty"`

	// Where the specimen came from
	Subject common.Reference `json:"subject"`

	// Identifier assigned by the lab
	AccessionIdentifier []common.Identifier `json:"accessionIdentifier,omitempty"` // R2 uses array

	// Time when specimen was received for processing
	ReceivedTime *string `json:"receivedTime,omitempty"` // R2 uses string

	// Collection details
	Collection *SpecimenCollection `json:"collection,omitempty"`

	// Treatment and processing step details
	Treatment []SpecimenTreatment `json:"treatment,omitempty"` // R2 uses treatment instead of processing

	// The container holding the specimen
	Container []SpecimenContainer `json:"container,omitempty"`
}

// SpecimenCollection represents details concerning the specimen collection (R2)
type SpecimenCollection struct {
	common.BackboneElement

	// Who collected the specimen
	Collector *common.Reference `json:"collector,omitempty"`

	// Collector comments
	Comment []string `json:"comment,omitempty"` // R2 uses string array

	// Collection time
	CollectedDateTime *string        `json:"collectedDateTime,omitempty"` // R2 uses string
	CollectedPeriod   *common.Period `json:"collectedPeriod,omitempty"`

	// How much was collected
	Quantity *common.Quantity `json:"quantity,omitempty"`

	// Technique used to perform collection
	Method *common.CodeableConcept `json:"method,omitempty"`

	// Anatomical collection site
	BodySiteCodeableConcept *common.CodeableConcept `json:"bodySiteCodeableConcept,omitempty"`
	BodySiteReference       *common.Reference       `json:"bodySiteReference,omitempty"`
}

// SpecimenTreatment represents treatment and processing step for the specimen (R2)
type SpecimenTreatment struct {
	common.BackboneElement

	// Textual description of procedure
	Description *string `json:"description,omitempty"`

	// Indicates the treatment or processing step applied to the specimen
	Procedure *common.CodeableConcept `json:"procedure,omitempty"`

	// Material used in the processing step
	Additive []common.Reference `json:"additive,omitempty"`
}

// SpecimenContainer represents the container holding the specimen (R2)
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
	Additive *common.Reference `json:"additive,omitempty"` // R2 uses single reference
} 
