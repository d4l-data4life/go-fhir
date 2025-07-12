// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"time"

	"github.com/go-fhir/go-fhir/pkg/common"
)

// Meta represents metadata about a resource (R5 specific)
type Meta struct {
	common.Element

	// Version specific id
	VersionId *string `json:"versionId,omitempty"`

	// When the resource version last changed
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`

	// Profiles this resource claims to conform to
	Profile []string `json:"profile,omitempty"`

	// Security Labels applied to this resource
	Security []common.Coding `json:"security,omitempty"`

	// Tags applied to this resource
	Tag []common.Coding `json:"tag,omitempty"`

	// Identifies where the resource comes from
	Source *string `json:"source,omitempty"`
}

// Narrative represents human-readable content (R5 specific)
type Narrative struct {
	common.Element

	// generated | extensions | additional | empty
	Status NarrativeStatus `json:"status"`

	// Limited xhtml content
	Div string `json:"div"`
}

// NarrativeStatus represents the status of the narrative
type NarrativeStatus string

const (
	NarrativeStatusGenerated   NarrativeStatus = "generated"
	NarrativeStatusExtensions  NarrativeStatus = "extensions"
	NarrativeStatusAdditional  NarrativeStatus = "additional"
	NarrativeStatusEmpty       NarrativeStatus = "empty"
)

// Signature represents a digital signature (R5 specific)
type Signature struct {
	common.Element

	// Indication of the reason the entity signed the object(s)
	Type []common.Coding `json:"type"`

	// When the signature was created
	When time.Time `json:"when"`

	// Who signed
	Who common.Reference `json:"who"`

	// The party represented
	OnBehalfOf *common.Reference `json:"onBehalfOf,omitempty"`

	// The technical format of the signature
	TargetFormat *string `json:"targetFormat,omitempty"`

	// The technical format of the signed resources
	SigFormat *string `json:"sigFormat,omitempty"`

	// The actual signature content (XML DigSig. JWS, picture, etc.)
	Data *string `json:"data,omitempty"`
}

// CodeableReference represents a reference to a resource with a code (R5 specific)
type CodeableReference struct {
	common.Element

	// Reference to a concept
	Concept *common.CodeableConcept `json:"concept,omitempty"`

	// Reference to a resource
	Reference *common.Reference `json:"reference,omitempty"`
}

// ContactPoint represents details for all kinds of technology-mediated contact points (R5 enhanced)
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
	ContactPointSystemPhone  ContactPointSystem = "phone"
	ContactPointSystemFax    ContactPointSystem = "fax"
	ContactPointSystemEmail  ContactPointSystem = "email"
	ContactPointSystemPager  ContactPointSystem = "pager"
	ContactPointSystemUrl    ContactPointSystem = "url"
	ContactPointSystemSms    ContactPointSystem = "sms"
	ContactPointSystemOther  ContactPointSystem = "other"
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

// Resource is the base type for all FHIR resources in R5
type Resource struct {
	common.Element

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType,omitempty"`

	// Logical id of this artifact
	ID *string `json:"id,omitempty"`

	// A set of rules under which this content was created
	ImplicitRules *string `json:"implicitRules,omitempty"`

	// Language of the resource content
	Language *string `json:"language,omitempty"`

	// Metadata about the resource
	Meta *Meta `json:"meta,omitempty"`
}

// DomainResource is the base type for all domain resources in R5
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
	Signature []Signature `json:"signature,omitempty"`

	// Issues with the Bundle
	Issues *Resource `json:"issues,omitempty"`
}

// BundleType represents the purpose of this bundle
type BundleType string

const (
	BundleTypeDocument       BundleType = "document"
	BundleTypeMessage        BundleType = "message"
	BundleTypeTransaction    BundleType = "transaction"
	BundleTypeTransactionResponse BundleType = "transaction-response"
	BundleTypeBatch          BundleType = "batch"
	BundleTypeBatchResponse  BundleType = "batch-response"
	BundleTypeHistory        BundleType = "history"
	BundleTypeSearchset      BundleType = "searchset"
	BundleTypeCollection     BundleType = "collection"
	BundleTypeSubscriptionNotification BundleType = "subscription-notification"
)

// BundleLink represents a series of links that provide context to this bundle
type BundleLink struct {
	common.BackboneElement

	// Relation to the current entry
	Relation string `json:"relation"`

	// Reference details for the link
	Url string `json:"url"`
}

// BundleEntry represents an entry in a bundle resource
type BundleEntry struct {
	common.BackboneElement

	// Links related to this entry
	Link []BundleLink `json:"link,omitempty"`

	// URI for resource (Absolute URL)
	FullUrl *string `json:"fullUrl,omitempty"`

	// A resource in the bundle
	Resource *DomainResource `json:"resource,omitempty"`

	// Search related information
	Search *BundleEntrySearch `json:"search,omitempty"`

	// Additional execution information (transaction/batch/history)
	Request *BundleEntryRequest `json:"request,omitempty"`

	// Results of execution (transaction/batch/history)
	Response *BundleEntryResponse `json:"response,omitempty"`
}

// BundleEntrySearch represents search related information
type BundleEntrySearch struct {
	common.Element

	// Search ranking (between 0 and 1)
	Score *float64 `json:"score,omitempty"`

	// Why this entry is in the result set
	Mode *BundleEntrySearchMode `json:"mode,omitempty"`
}

// BundleEntrySearchMode represents why this entry is in the result set
type BundleEntrySearchMode string

const (
	BundleEntrySearchModeMatch   BundleEntrySearchMode = "match"
	BundleEntrySearchModeInclude BundleEntrySearchMode = "include"
	BundleEntrySearchModeOutcome BundleEntrySearchMode = "outcome"
)

// BundleEntryRequest represents additional execution information for transaction/batch/history
type BundleEntryRequest struct {
	common.Element

	// In a transaction or batch, this is the HTTP action to be executed for this entry
	Method BundleEntryRequestMethod `json:"method"`

	// URL for HTTP equivalent of this entry
	Url string `json:"url"`

	// For managing cache currency
	IfNoneMatch *string `json:"ifNoneMatch,omitempty"`

	// For managing cache currency
	IfModifiedSince *string `json:"ifModifiedSince,omitempty"`

	// For managing update contention
	IfMatch *string `json:"ifMatch,omitempty"`

	// For conditional creates
	IfNoneExist *string `json:"ifNoneExist,omitempty"`
}

// BundleEntryRequestMethod represents HTTP methods
type BundleEntryRequestMethod string

const (
	BundleEntryRequestMethodGET    BundleEntryRequestMethod = "GET"
	BundleEntryRequestMethodHEAD   BundleEntryRequestMethod = "HEAD"
	BundleEntryRequestMethodPOST   BundleEntryRequestMethod = "POST"
	BundleEntryRequestMethodPUT    BundleEntryRequestMethod = "PUT"
	BundleEntryRequestMethodDELETE BundleEntryRequestMethod = "DELETE"
	BundleEntryRequestMethodPATCH  BundleEntryRequestMethod = "PATCH"
)

// BundleEntryResponse represents results of execution for transaction/batch/history
type BundleEntryResponse struct {
	common.Element

	// Status response code (text optional)
	Status string `json:"status"`

	// The location (if the operation returns a location)
	Location *string `json:"location,omitempty"`

	// The Etag for the resource (if relevant)
	Etag *string `json:"etag,omitempty"`

	// Server's date time modified
	LastModified *time.Time `json:"lastModified,omitempty"`

	// OperationOutcome with hints and warnings (for batch/transaction)
	Outcome *Resource `json:"outcome,omitempty"`
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

	// Form being answered (required in R5)
	Questionnaire string `json:"questionnaire"`

	// in-progress | completed | amended | entered-in-error | stopped
	Status QuestionnaireResponseStatus `json:"status"`

	// The subject of the questionnaire response
	Subject *common.Reference `json:"subject,omitempty"`

	// Encounter created as part of
	Encounter *common.Reference `json:"encounter,omitempty"`

	// When the questionnaire response was authored
	Authored *time.Time `json:"authored,omitempty"`

	// Person who received and recorded the answers
	Author *common.Reference `json:"author,omitempty"`

	// The person who answered the questions
	Source *common.Reference `json:"source,omitempty"`

	// Groups and questions
	Item []QuestionnaireResponseItem `json:"item,omitempty"`
}

// QuestionnaireResponseStatus represents the lifecycle status
type QuestionnaireResponseStatus string

const (
	QuestionnaireResponseStatusInProgress     QuestionnaireResponseStatus = "in-progress"
	QuestionnaireResponseStatusCompleted      QuestionnaireResponseStatus = "completed"
	QuestionnaireResponseStatusAmended        QuestionnaireResponseStatus = "amended"
	QuestionnaireResponseStatusEnteredInError QuestionnaireResponseStatus = "entered-in-error"
	QuestionnaireResponseStatusStopped        QuestionnaireResponseStatus = "stopped"
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
	ValueBoolean    *bool                    `json:"valueBoolean,omitempty"`
	ValueDecimal    *float64                 `json:"valueDecimal,omitempty"`
	ValueInteger    *int                     `json:"valueInteger,omitempty"`
	ValueDate       *string                  `json:"valueDate,omitempty"`
	ValueDateTime   *time.Time               `json:"valueDateTime,omitempty"`
	ValueTime       *string                  `json:"valueTime,omitempty"`
	ValueString     *string                  `json:"valueString,omitempty"`
	ValueUri        *string                  `json:"valueUri,omitempty"`
	ValueAttachment *Attachment              `json:"valueAttachment,omitempty"`
	ValueCoding     *common.Coding           `json:"valueCoding,omitempty"`
	ValueQuantity   *common.Quantity         `json:"valueQuantity,omitempty"`
	ValueReference  *common.Reference        `json:"valueReference,omitempty"`

	// Nested groups and questions
	Item []QuestionnaireResponseItem `json:"item,omitempty"`
}

// ResearchSubject represents a ResearchSubject in R5 with enhanced progress tracking
type ResearchSubject struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "ResearchSubject"

	// Business identifier for research subject in a study
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// draft | active | retired | unknown (simplified in R5)
	Status ResearchSubjectStatus `json:"status"`

	// Dates the subject began and ended their participation in the study
	Period *common.Period `json:"period,omitempty"`

	// Study subject is participating in
	Study common.Reference `json:"study"`

	// Who or what is part of study
	Subject common.Reference `json:"subject"`

	// The name of the arm in the study the subject actually followed as part of this study
	ActualComparisonGroup *string `json:"actualComparisonGroup,omitempty"`

	// What path should be followed
	AssignedComparisonGroup *string `json:"assignedComparisonGroup,omitempty"`

	// Agreement to participate in study (R5 enhancement - now array)
	Consent []common.Reference `json:"consent,omitempty"`

	// Subject progress through study (NEW in R5 - major enhancement)
	Progress []ResearchSubjectProgress `json:"progress,omitempty"`
}

// ResearchSubjectStatus represents the status (simplified in R5)
type ResearchSubjectStatus string

const (
	ResearchSubjectStatusDraft   ResearchSubjectStatus = "draft"
	ResearchSubjectStatusActive  ResearchSubjectStatus = "active"
	ResearchSubjectStatusRetired ResearchSubjectStatus = "retired"
	ResearchSubjectStatusUnknown ResearchSubjectStatus = "unknown"
)

// ResearchSubjectProgress represents subject progress through study (NEW in R5)
type ResearchSubjectProgress struct {
	common.BackboneElement

	// Identifies the aspect of the subject's journey that the state refers to
	Type *common.CodeableConcept `json:"type,omitempty"`

	// The current state of the subject
	SubjectState *common.CodeableConcept `json:"subjectState,omitempty"`

	// There can be multiple entries but it is also valid to just have the most recent
	Milestone *common.CodeableConcept `json:"milestone,omitempty"`

	// The reason for the state change
	Reason *common.CodeableConcept `json:"reason,omitempty"`

	// The date when the state started
	StartDate *string `json:"startDate,omitempty"`

	// The date when the state ended
	EndDate *string `json:"endDate,omitempty"`
}

// Consent represents a record of a healthcare consumer's choices (R5 enhanced)
type Consent struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Consent"

	// Identifier for this copy of the Consent Statement
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// draft | active | inactive | not-done | entered-in-error | unknown (updated in R5)
	Status ConsentStatus `json:"status"`

	// Classification of the consent statement - for indexing/retrieval
	Category []common.CodeableConcept `json:"category,omitempty"`

	// Who the consent applies to
	Subject *common.Reference `json:"subject,omitempty"`

	// When consent was agreed to
	Date *string `json:"date,omitempty"`

	// Effective period for this Consent Resource and all provisions unless specified in that provision
	Period *common.Period `json:"period,omitempty"`

	// Who is granting rights according to the policy and rules (NEW in R5)
	Grantor []common.Reference `json:"grantor,omitempty"`

	// Who is agreeing to the policy and rules (NEW in R5)
	Grantee []common.Reference `json:"grantee,omitempty"`

	// Who is managing the consent through its lifecycle (NEW in R5)
	Manager []common.Reference `json:"manager,omitempty"`

	// Who controls/enforces the access according to the consent (NEW in R5)
	Controller []common.Reference `json:"controller,omitempty"`

	// Source from which this consent is taken (enhanced in R5 - now arrays)
	SourceAttachment []Attachment        `json:"sourceAttachment,omitempty"`
	SourceReference  []common.Reference  `json:"sourceReference,omitempty"`

	// A set of codes that indicate the regulatory basis (NEW in R5)
	RegulatoryBasis []common.CodeableConcept `json:"regulatoryBasis,omitempty"`

	// Policies covered by this consent (enhanced in R5)
	PolicyBasis *ConsentPolicyBasis `json:"policyBasis,omitempty"`

	// A Reference to the human readable policy explaining the basis for the Consent (enhanced in R5)
	PolicyText []common.Reference `json:"policyText,omitempty"`

	// Consent Verified by patient or family (enhanced in R5)
	Verification []ConsentVerification `json:"verification,omitempty"`

	// Action to take - permit or deny - as default (NEW in R5)
	Decision *ConsentDecision `json:"decision,omitempty"`

	// Constraints to the base Consent.policyRule/Consent.policy
	Provision []ConsentProvision `json:"provision,omitempty"`
}

// ConsentStatus represents the consent status (updated in R5)
type ConsentStatus string

const (
	ConsentStatusDraft           ConsentStatus = "draft"
	ConsentStatusActive          ConsentStatus = "active"
	ConsentStatusInactive        ConsentStatus = "inactive"
	ConsentStatusNotDone         ConsentStatus = "not-done"
	ConsentStatusEnteredInError  ConsentStatus = "entered-in-error"
	ConsentStatusUnknown         ConsentStatus = "unknown"
)

// ConsentDecision represents action to take as default (NEW in R5)
type ConsentDecision string

const (
	ConsentDecisionDeny   ConsentDecision = "deny"
	ConsentDecisionPermit ConsentDecision = "permit"
)

// ConsentPolicyBasis represents policies covered by this consent (enhanced in R5)
type ConsentPolicyBasis struct {
	common.Element

	// Reference to a specific policy
	Reference *common.Reference `json:"reference,omitempty"`

	// URL to a policy
	Url *string `json:"url,omitempty"`
}

// ConsentVerification represents consent verified by patient or family (enhanced in R5)
type ConsentVerification struct {
	common.BackboneElement

	// Has the instruction been verified
	Verified bool `json:"verified"`

	// Type of verification
	VerificationType *common.CodeableConcept `json:"verificationType,omitempty"`

	// Person conducting verification
	VerifiedBy *common.Reference `json:"verifiedBy,omitempty"`

	// Person who verified the instruction
	VerifiedWith *common.Reference `json:"verifiedWith,omitempty"`

	// When consent verified
	VerificationDate []string `json:"verificationDate,omitempty"`
}

// ConsentProvision represents constraints to the base policy
type ConsentProvision struct {
	common.BackboneElement

	// Timeframe for this provision
	Period *common.Period `json:"period,omitempty"`

	// Who|what controlled by this provision (UsageContext)
	Actor []ConsentProvisionActor `json:"actor,omitempty"`

	// Actions controlled by this provision
	Action []common.CodeableConcept `json:"action,omitempty"`

	// Security Labels that define affected resources
	SecurityLabel []common.Coding `json:"securityLabel,omitempty"`

	// Context of activities covered by this provision
	Purpose []common.Coding `json:"purpose,omitempty"`

	// e.g. Resource Type, Profile, CDA, etc.
	DocumentType []common.Coding `json:"documentType,omitempty"`

	// e.g. Resource Type, Profile, etc.
	ResourceType []common.Coding `json:"resourceType,omitempty"`

	// Codes that identify fine-grained provisions
	Code []common.CodeableConcept `json:"code,omitempty"`

	// Timeframe for data controlled by this provision
	DataPeriod *common.Period `json:"dataPeriod,omitempty"`

	// Data controlled by this provision
	Data []ConsentProvisionData `json:"data,omitempty"`

	// A computable expression of the consent
	Expression *Expression `json:"expression,omitempty"`

	// Nested Exception provisions
	Provision []ConsentProvision `json:"provision,omitempty"`
}

// ConsentProvisionActor represents who|what controlled by this provision
type ConsentProvisionActor struct {
	common.BackboneElement

	// How the actor is involved
	Role common.CodeableConcept `json:"role"`

	// Resource for the actor (or group, by role)
	Reference *common.Reference `json:"reference,omitempty"`
}

// ConsentProvisionData represents data controlled by this provision
type ConsentProvisionData struct {
	common.BackboneElement

	// How the resource reference is interpreted when testing consent restrictions
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
	ConsentProvisionDataMeaningAuthoredBy ConsentProvisionDataMeaning = "authoredby"
)

// Device represents a medical or non-medical device used in healthcare (R5 enhanced)
type Device struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Device"

	// Instance identifier
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The name used to display by default when the device is referenced (NEW in R5)
	DisplayName *string `json:"displayName,omitempty"`

	// The reference to the definition for the device (enhanced in R5)
	Definition *CodeableReference `json:"definition,omitempty"`

	// Unique Device Identifier (UDI) Barcode string
	UdiCarrier []DeviceUdiCarrier `json:"udiCarrier,omitempty"`

	// active | inactive | entered-in-error (simplified in R5)
	Status *DeviceStatus `json:"status,omitempty"`

	// The availability of the device (NEW in R5)
	AvailabilityStatus *common.CodeableConcept `json:"availabilityStatus,omitempty"`

	// Necessary to support mandatory requirements for traceability (NEW in R5)
	BiologicalSourceEvent *common.Identifier `json:"biologicalSourceEvent,omitempty"`

	// A name of the manufacturer
	Manufacturer *string `json:"manufacturer,omitempty"`

	// Date when the device was made
	ManufactureDate *string `json:"manufactureDate,omitempty"`

	// Date and time of expiry of this device (if applicable)
	ExpirationDate *string `json:"expirationDate,omitempty"`

	// Lot number assigned by the manufacturer
	LotNumber *string `json:"lotNumber,omitempty"`

	// The manufacturer's model number for the device
	ModelNumber *string `json:"modelNumber,omitempty"`

	// Alphanumeric Maximum 20
	PartNumber *string `json:"partNumber,omitempty"`

	// Device category (enhanced in R5 - now array)
	Category []common.CodeableConcept `json:"category,omitempty"`

	// Multiple device types (enhanced in R5 - now array)
	Type []common.CodeableConcept `json:"type,omitempty"`

	// The actual design of the device or software version running on the device
	Version []DeviceVersion `json:"version,omitempty"`

	// Standards, specifications, or formal guidances (NEW in R5)
	ConformsTo []DeviceConformsTo `json:"conformsTo,omitempty"`

	// Device properties (enhanced in R5)
	Property []DeviceProperty `json:"property,omitempty"`

	// The designated condition for performing a task with the device (NEW in R5)
	Mode *common.CodeableConcept `json:"mode,omitempty"`

	// The series of occurrences that repeats during the operation of the device (NEW in R5)
	Cycle *Count `json:"cycle,omitempty"`

	// A measurement of time during the device's operation (NEW in R5)
	Duration *Duration `json:"duration,omitempty"`

	// An organization that is responsible for the provision and ongoing maintenance
	Owner *common.Reference `json:"owner,omitempty"`

	// Used for troubleshooting etc. (NEW in R5)
	Contact []ContactPoint `json:"contact,omitempty"`

	// Where the device is found
	Location *common.Reference `json:"location,omitempty"`

	// Network address to contact device
	Url *string `json:"url,omitempty"`

	// Technical endpoints providing access to services (NEW in R5)
	Endpoint []common.Reference `json:"endpoint,omitempty"`

	// The Device.gateway may be used when the Device being referenced has a gateway (NEW in R5)
	Gateway []CodeableReference `json:"gateway,omitempty"`

	// Device notes and comments
	Note []Annotation `json:"note,omitempty"`

	// Safety Characteristics of Device
	Safety []common.CodeableConcept `json:"safety,omitempty"`

	// The parent device
	Parent *common.Reference `json:"parent,omitempty"`
}

// DeviceStatus represents the status (simplified in R5)
type DeviceStatus string

const (
	DeviceStatusActive          DeviceStatus = "active"
	DeviceStatusInactive        DeviceStatus = "inactive"
	DeviceStatusEnteredInError  DeviceStatus = "entered-in-error"
)

// DeviceVersion represents the actual design of the device or software version
type DeviceVersion struct {
	common.BackboneElement

	// The type of the device version, e.g. manufacturer, approved, internal
	Type *common.CodeableConcept `json:"type,omitempty"`

	// The hardware or software module of the device to which the version applies
	Component *common.Identifier `json:"component,omitempty"`

	// The date the version was installed on the device
	InstallDate *string `json:"installDate,omitempty"`

	// The version text
	Value string `json:"value"`
}

// DeviceConformsTo represents standards, specifications, or formal guidances (NEW in R5)
type DeviceConformsTo struct {
	common.BackboneElement

	// Describes the type of the standard, specification, or formal guidance
	Category *common.CodeableConcept `json:"category,omitempty"`

	// Identifies the standard, specification, or formal guidance that the device adheres to
	Specification common.CodeableConcept `json:"specification"`

	// The specific form or variant of the standard
	Version []string `json:"version,omitempty"`
}

// DeviceProperty represents device properties (enhanced in R5)
type DeviceProperty struct {
	common.BackboneElement

	// Code that specifies the property being represented
	Type common.CodeableConcept `json:"type"`

	// Property value as a quantity
	ValueQuantity []common.Quantity `json:"valueQuantity,omitempty"`

	// Property value as a code
	ValueCodeableConcept []common.CodeableConcept `json:"valueCodeableConcept,omitempty"`

	// Property value as a string
	ValueString []string `json:"valueString,omitempty"`

	// Property value as a boolean
	ValueBoolean []bool `json:"valueBoolean,omitempty"`

	// Property value as an integer
	ValueInteger []int `json:"valueInteger,omitempty"`

	// Property value as a range
	ValueRange []Range `json:"valueRange,omitempty"`

	// Property value as an attachment
	ValueAttachment []Attachment `json:"valueAttachment,omitempty"`
}

// DeviceUdiCarrier represents Unique device identifier (UDI) assigned to device label or package
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

	// barcode | rfid | manual + (how the UDI was captured)
	EntryType *DeviceUdiCarrierEntryType `json:"entryType,omitempty"`
}

// DeviceUdiCarrierEntryType represents how the UDI was captured
type DeviceUdiCarrierEntryType string

const (
	DeviceUdiCarrierEntryTypeBarcode    DeviceUdiCarrierEntryType = "barcode"
	DeviceUdiCarrierEntryTypeRfid       DeviceUdiCarrierEntryType = "rfid"
	DeviceUdiCarrierEntryTypeManual     DeviceUdiCarrierEntryType = "manual"
	DeviceUdiCarrierEntryTypeCard       DeviceUdiCarrierEntryType = "card"
	DeviceUdiCarrierEntryTypeSelfReported DeviceUdiCarrierEntryType = "self-reported"
	DeviceUdiCarrierEntryTypeElectronicTransmission DeviceUdiCarrierEntryType = "electronic-transmission"
	DeviceUdiCarrierEntryTypeUnknown    DeviceUdiCarrierEntryType = "unknown"
)

// Composition represents a document structure for clinical and administrative content (R5 updated)
type Composition struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Composition"

	// Business identifier assigned to a composition
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The URL that this Composition is about (NEW in R5)
	Url *string `json:"url,omitempty"`

	// preliminary | final | amended | entered-in-error + (updated status codes in R5)
	Status CompositionStatus `json:"status"`

	// Kind of composition (LOINC if possible)
	Type common.CodeableConcept `json:"type"`

	// Categorization of Composition
	Category []common.CodeableConcept `json:"category,omitempty"`

	// Who and/or what the composition is about
	Subject []common.Reference `json:"subject,omitempty"`

	// Context of the Composition
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Composition editing time
	Date time.Time `json:"date"`

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

// CompositionStatus represents the composition status (updated in R5)
type CompositionStatus string

const (
	CompositionStatusRegistered      CompositionStatus = "registered"
	CompositionStatusPartial         CompositionStatus = "partial"
	CompositionStatusPreliminary     CompositionStatus = "preliminary"
	CompositionStatusFinal           CompositionStatus = "final"
	CompositionStatusAmended         CompositionStatus = "amended"
	CompositionStatusCorrected       CompositionStatus = "corrected"
	CompositionStatusAppended        CompositionStatus = "appended"
	CompositionStatusCancelled       CompositionStatus = "cancelled"
	CompositionStatusEnteredInError  CompositionStatus = "entered-in-error"
	CompositionStatusDeprecated      CompositionStatus = "deprecated"
	CompositionStatusUnknown         CompositionStatus = "unknown"
)

// CompositionAttester represents who attested the composition
type CompositionAttester struct {
	common.BackboneElement

	// personal | professional | legal | official
	Mode []CompositionAttesterMode `json:"mode"`

	// When the composition was attested
	Time *time.Time `json:"time,omitempty"`

	// Who attested the composition
	Party *common.Reference `json:"party,omitempty"`
}

// CompositionAttesterMode represents the type of attestation
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

	// The type of relationship that this composition has with anther composition or document
	Type CompositionRelatesToType `json:"type"`

	// Target of the relationship
	TargetIdentifier *common.Identifier `json:"targetIdentifier,omitempty"`
	TargetReference  *common.Reference  `json:"targetReference,omitempty"`
}

// CompositionRelatesToType represents the type of relationship
type CompositionRelatesToType string

const (
	CompositionRelatesToTypeReplaces  CompositionRelatesToType = "replaces"
	CompositionRelatesToTypeTransforms CompositionRelatesToType = "transforms"
	CompositionRelatesToTypeSigns     CompositionRelatesToType = "signs"
	CompositionRelatesToTypeAppends   CompositionRelatesToType = "appends"
)

// CompositionEvent represents the clinical service(s) being documented
type CompositionEvent struct {
	common.BackboneElement

	// Code(s) that apply to the event being documented
	Detail []common.Reference `json:"detail,omitempty"`

	// The period covered by the documentation
	Period *common.Period `json:"period,omitempty"`
}

// CompositionSection represents composition sections
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

	// How the entry list was prepared - whether it is a working list that is suitable for being maintained on an ongoing basis, or if it represents a snapshot of a list of items from another source, or whether it is a prepared list where items may be marked as added, modified or deleted
	Mode *CompositionSectionMode `json:"mode,omitempty"`

	// Order of section entries
	OrderedBy *common.CodeableConcept `json:"orderedBy,omitempty"`

	// A reference to the actual resource from which the narrative in the section is derived
	Entry []common.Reference `json:"entry,omitempty"`

	// Why the section is empty
	EmptyReason *common.CodeableConcept `json:"emptyReason,omitempty"`

	// Nested Sections
	Section []CompositionSection `json:"section,omitempty"`
}

// CompositionSectionMode represents how the entry list was prepared
type CompositionSectionMode string

const (
	CompositionSectionModeWorking  CompositionSectionMode = "working"
	CompositionSectionModeSnapshot CompositionSectionMode = "snapshot"
	CompositionSectionModeChanges  CompositionSectionMode = "changes"
)

// DocumentReference represents a reference to a document of any kind (R5 enhanced)
type DocumentReference struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "DocumentReference"

	// Business identifiers for the document
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// An explicitly assigned identifer of a variation of the content in the DocumentReference
	Version *string `json:"version,omitempty"`

	// Procedure that caused this media to be created (enhanced in R5)
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// current | superseded | entered-in-error (updated in R5)
	Status DocumentReferenceStatus `json:"status"`

	// preliminary | final | amended | entered-in-error (enhanced in R5)
	DocStatus *DocumentReferenceDocStatus `json:"docStatus,omitempty"`

	// The referenced time(s) when the document was created (enhanced in R5)
	Date *time.Time `json:"date,omitempty"`

	// Who and/or what authored the document
	Author []common.Reference `json:"author,omitempty"`

	// Who/what is the subject of the document
	Subject *common.Reference `json:"subject,omitempty"`

	// Context of the document content
	Context []DocumentReferenceContext `json:"context,omitempty"`

	// Kind of document (LOINC if possible)
	Type *common.CodeableConcept `json:"type,omitempty"`

	// Categorization of document
	Category []common.CodeableConcept `json:"category,omitempty"`

	// Additional details about where the content was created (e.g. clinical specialty)
	Specialty []common.CodeableConcept `json:"specialty,omitempty"`

	// Organization which maintains the document
	Custodian *common.Reference `json:"custodian,omitempty"`

	// Relationships to other documents
	RelatesTo []DocumentReferenceRelatesTo `json:"relatesTo,omitempty"`

	// Human-readable description
	Description *string `json:"description,omitempty"`

	// Document security-tags
	SecurityLabel []common.CodeableConcept `json:"securityLabel,omitempty"`

	// Document referenced (enhanced content profiles in R5)
	Content []DocumentReferenceContent `json:"content"`
}

// DocumentReferenceStatus represents the status
type DocumentReferenceStatus string

const (
	DocumentReferenceStatusCurrent        DocumentReferenceStatus = "current"
	DocumentReferenceStatusSuperseded     DocumentReferenceStatus = "superseded"
	DocumentReferenceStatusEnteredInError DocumentReferenceStatus = "entered-in-error"
)

// DocumentReferenceDocStatus represents the document status
type DocumentReferenceDocStatus string

const (
	DocumentReferenceDocStatusPreliminary    DocumentReferenceDocStatus = "preliminary"
	DocumentReferenceDocStatusFinal          DocumentReferenceDocStatus = "final"
	DocumentReferenceDocStatusAmended        DocumentReferenceDocStatus = "amended"
	DocumentReferenceDocStatusEnteredInError DocumentReferenceDocStatus = "entered-in-error"
)

// DocumentReferenceContext represents context of the document content
type DocumentReferenceContext struct {
	common.BackboneElement

	// Context of the document content
	Encounter []common.Reference `json:"encounter,omitempty"`

	// Kind of event (code)
	Event []common.CodeableConcept `json:"event,omitempty"`

	// Time of service that is being documented
	Period *common.Period `json:"period,omitempty"`

	// Additional details about where the content was created (e.g. clinical specialty)
	FacilityType *common.CodeableConcept `json:"facilityType,omitempty"`

	// Additional details about where the content was created (e.g. clinical specialty)
	PracticeSetting *common.CodeableConcept `json:"practiceSetting,omitempty"`

	// Patient demographics from source
	SourcePatientInfo *common.Reference `json:"sourcePatientInfo,omitempty"`

	// Related identifiers or resources
	Related []common.Reference `json:"related,omitempty"`
}

// DocumentReferenceRelatesTo represents relationships to other documents
type DocumentReferenceRelatesTo struct {
	common.BackboneElement

	// The type of relationship that this document has with anther document
	Type DocumentReferenceRelatesToType `json:"type"`

	// Target of the relationship
	Target common.Reference `json:"target"`
}

// DocumentReferenceRelatesToType represents the type of relationship
type DocumentReferenceRelatesToType string

const (
	DocumentReferenceRelatesToTypeReplaces   DocumentReferenceRelatesToType = "replaces"
	DocumentReferenceRelatesToTypeTransforms DocumentReferenceRelatesToType = "transforms"
	DocumentReferenceRelatesToTypeSigns      DocumentReferenceRelatesToType = "signs"
	DocumentReferenceRelatesToTypeAppends    DocumentReferenceRelatesToType = "appends"
)

// DocumentReferenceContent represents document referenced (enhanced in R5)
type DocumentReferenceContent struct {
	common.BackboneElement

	// Where to access the document
	Attachment Attachment `json:"attachment"`

	// Content profile rules for the document (NEW in R5)
	Profile []DocumentReferenceContentProfile `json:"profile,omitempty"`
}

// DocumentReferenceContentProfile represents content profile rules (NEW in R5)
type DocumentReferenceContentProfile struct {
	common.BackboneElement

	// Code of the profile
	ValueCodeableConcept *common.CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueCoding          *common.Coding          `json:"valueCoding,omitempty"`
	ValueUri             *string                 `json:"valueUri,omitempty"`
	ValueCanonical       *string                 `json:"valueCanonical,omitempty"`
}

// Communication represents a record of communication between entities (R5 enhanced)
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

	// Characterizes how quickly the planned or in progress communication must be addressed (enhanced in R5)
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
	Sent *time.Time `json:"sent,omitempty"`

	// When received
	Received *time.Time `json:"received,omitempty"`

	// Who the information is sent to
	Recipient []common.Reference `json:"recipient,omitempty"`

	// Who initiated the communication
	Sender *common.Reference `json:"sender,omitempty"`

	// Indication for message (enhanced in R5)
	Reason []CodeableReference `json:"reason,omitempty"`

	// Message payload (enhanced types in R5)
	Payload []CommunicationPayload `json:"payload,omitempty"`

	// Comments made about the communication
	Note []Annotation `json:"note,omitempty"`
}

// CommunicationStatus represents the status
type CommunicationStatus string

const (
	CommunicationStatusPreparation    CommunicationStatus = "preparation"
	CommunicationStatusInProgress     CommunicationStatus = "in-progress"
	CommunicationStatusNotDone        CommunicationStatus = "not-done"
	CommunicationStatusOnHold         CommunicationStatus = "on-hold"
	CommunicationStatusStopped        CommunicationStatus = "stopped"
	CommunicationStatusCompleted      CommunicationStatus = "completed"
	CommunicationStatusEnteredInError CommunicationStatus = "entered-in-error"
	CommunicationStatusUnknown        CommunicationStatus = "unknown"
)

// CommunicationPriority represents the priority (enhanced in R5)
type CommunicationPriority string

const (
	CommunicationPriorityRoutine CommunicationPriority = "routine"
	CommunicationPriorityUrgent  CommunicationPriority = "urgent"
	CommunicationPriorityAsap    CommunicationPriority = "asap"
	CommunicationPriorityStat    CommunicationPriority = "stat"
)

// CommunicationPayload represents message payload (enhanced types in R5)
type CommunicationPayload struct {
	common.BackboneElement

	// Message part content
	ContentAttachment  *Attachment           `json:"contentAttachment,omitempty"`
	ContentReference   *common.Reference     `json:"contentReference,omitempty"`
	ContentCodeableConcept *common.CodeableConcept `json:"contentCodeableConcept,omitempty"`
}

// Expression represents a computable expression
type Expression struct {
	common.Element

	// Natural language description of the condition
	Description *string `json:"description,omitempty"`

	// Short name assigned to expression for reuse
	Name *string `json:"name,omitempty"`

	// text/cql | text/fhirpath | application/x-fhir-query | etc.
	Language string `json:"language"`

	// Expression in specified language
	Expression *string `json:"expression,omitempty"`

	// Where the expression is found
	Reference *string `json:"reference,omitempty"`
}

// Helper functions for creating pointers to basic types
func StringPtr(s string) *string { return &s }
func IntPtr(i int) *int { return &i }
func BoolPtr(b bool) *bool { return &b }
func Float64Ptr(f float64) *float64 { return &f }
func TimePtr(t time.Time) *time.Time { return &t } 

// AllergyIntolerance represents risk of harmful or undesirable, physiological response to a substance (R5 enhanced)
type AllergyIntolerance struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "AllergyIntolerance"

	// External ids for this item
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// active | inactive | resolved (enhanced in R5)
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

	// Who the allergy or intolerance is for
	Patient common.Reference `json:"patient"`

	// Encounter when the allergy or intolerance was asserted
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Date first version of the resource instance was recorded
	OnsetDateTime *time.Time     `json:"onsetDateTime,omitempty"`
	OnsetAge      *Age           `json:"onsetAge,omitempty"`
	OnsetPeriod   *common.Period `json:"onsetPeriod,omitempty"`
	OnsetRange    *Range         `json:"onsetRange,omitempty"`
	OnsetString   *string        `json:"onsetString,omitempty"`

	// Date(/time) of last known occurrence of a reaction
	RecordedDate *time.Time `json:"recordedDate,omitempty"`

	// Indicates who or what participated in the activities related to the allergy or intolerance (NEW in R5)
	Participant []AllergyIntoleranceParticipant `json:"participant,omitempty"`

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

// AllergyIntoleranceParticipant represents who or what participated in the activities (NEW in R5)
type AllergyIntoleranceParticipant struct {
	common.BackboneElement

	// Type of involvement
	Function *common.CodeableConcept `json:"function,omitempty"`

	// Who participated
	Actor common.Reference `json:"actor"`
}

// AllergyIntoleranceReaction represents details about each adverse reaction event (enhanced in R5)
type AllergyIntoleranceReaction struct {
	common.BackboneElement

	// Specific substance or pharmaceutical product considered to be responsible
	Substance *common.CodeableConcept `json:"substance,omitempty"`

	// Clinical symptoms/signs associated with the Event (enhanced in R5)
	Manifestation []CodeableReference `json:"manifestation"`

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

// Immunization represents the event of a patient being administered a vaccine (R5 enhanced)
type Immunization struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Immunization"

	// Business identifier
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Instantiates protocol or definition (NEW in R5)
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// completed | entered-in-error | not-done
	Status ImmunizationStatus `json:"status"`

	// Reason for the vaccination status (enhanced in R5)
	StatusReason *common.CodeableConcept `json:"statusReason,omitempty"`

	// Vaccine administered (enhanced in R5)
	VaccineCode common.CodeableConcept `json:"vaccineCode"`

	// Product being administered (NEW in R5)
	AdministeredProduct *CodeableReference `json:"administeredProduct,omitempty"`

	// Vaccine manufacturer (enhanced in R5)
	Manufacturer *CodeableReference `json:"manufacturer,omitempty"`

	// Vaccine lot number
	LotNumber *string `json:"lotNumber,omitempty"`

	// Vaccine expiration date
	ExpirationDate *time.Time `json:"expirationDate,omitempty"`

	// Who was immunized
	Patient common.Reference `json:"patient"`

	// Encounter immunization was part of
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Additional information that is relevant to the immunization (NEW in R5)
	SupportingInformation []common.Reference `json:"supportingInformation,omitempty"`

	// Vaccine administration date
	OccurrenceDateTime *time.Time `json:"occurrenceDateTime,omitempty"`
	OccurrenceString   *string    `json:"occurrenceString,omitempty"`

	// Indicates if the dose is valid or not valid (enhanced in R5)
	PrimarySource *bool `json:"primarySource,omitempty"`

	// Indicates the source of the reported record (enhanced in R5)
	InformationSource *CodeableReference `json:"informationSource,omitempty"`

	// Where immunization occurred
	Location *common.Reference `json:"location,omitempty"`

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

	// Why immunization occurred (enhanced in R5)
	Reason []CodeableReference `json:"reason,omitempty"`

	// Dose potency
	IsSubpotent *bool `json:"isSubpotent,omitempty"`

	// Reason for being subpotent
	SubpotentReason []common.CodeableConcept `json:"subpotentReason,omitempty"`

	// Patient eligibility for a vaccination program (enhanced in R5)
	ProgramEligibility []ImmunizationProgramEligibility `json:"programEligibility,omitempty"`

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

// ImmunizationProgramEligibility represents patient eligibility for a vaccination program (NEW in R5)
type ImmunizationProgramEligibility struct {
	common.BackboneElement

	// The program
	Program common.CodeableConcept `json:"program"`

	// The patient's eligibility status for the program
	ProgramStatus common.CodeableConcept `json:"programStatus"`
}

// ImmunizationReaction represents details of a reaction following immunization (enhanced in R5)
type ImmunizationReaction struct {
	common.BackboneElement

	// When reaction started
	Date *time.Time `json:"date,omitempty"`

	// Additional information on reaction (enhanced in R5)
	Manifestation *CodeableReference `json:"manifestation,omitempty"`

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
	DoseNumber string `json:"doseNumber"`

	// Recommended number of doses for immunity
	SeriesDoses *string `json:"seriesDoses,omitempty"`
}

// CarePlan represents the intention of how one or more practitioners intend to deliver care (R5 enhanced)
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

	// A higher-level request resource (enhanced in R5)
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
	Created *time.Time `json:"created,omitempty"`

	// Who is the custodian (NEW in R5)
	Custodian *common.Reference `json:"custodian,omitempty"`

	// Who provided the content of the care plan
	Contributor []common.Reference `json:"contributor,omitempty"`

	// Who's involved in plan?
	CareTeam []common.Reference `json:"careTeam,omitempty"`

	// Health issues this plan addresses (enhanced in R5)
	Addresses []CodeableReference `json:"addresses,omitempty"`

	// Information considered as part of plan
	SupportingInfo []common.Reference `json:"supportingInfo,omitempty"`

	// Desired outcome of plan
	Goal []common.Reference `json:"goal,omitempty"`

	// Action to occur or has occurred as part of plan (enhanced in R5)
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

// CarePlanActivity represents action to occur or has occurred as part of plan (enhanced in R5)
type CarePlanActivity struct {
	common.BackboneElement

	// Results of the activity (enhanced in R5)
	PerformedActivity []CodeableReference `json:"performedActivity,omitempty"`

	// Comments about the activity status/progress
	Progress []Annotation `json:"progress,omitempty"`

	// Activity details defined in specific resource (enhanced in R5)
	PlannedActivityReference *common.Reference `json:"plannedActivityReference,omitempty"`
}

// Specimen represents a sample to be used for analysis (R5 enhanced)
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
	ReceivedTime *time.Time `json:"receivedTime,omitempty"`

	// Specimen from which this specimen originated
	Parent []common.Reference `json:"parent,omitempty"`

	// Why the specimen was collected
	Request []common.Reference `json:"request,omitempty"`

	// grouped | pooled (NEW in R5)
	Combined *SpecimenCombined `json:"combined,omitempty"`

	// The role or reason for the specimen in the testing workflow (NEW in R5)
	Role []common.CodeableConcept `json:"role,omitempty"`

	// A physical feature or landmark on a specimen (NEW in R5)
	Feature []SpecimenFeature `json:"feature,omitempty"`

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

// SpecimenCombined represents whether this is a group or pooled specimen (NEW in R5)
type SpecimenCombined string

const (
	SpecimenCombinedGrouped SpecimenCombined = "grouped"
	SpecimenCombinedPooled  SpecimenCombined = "pooled"
)

// SpecimenFeature represents a physical feature or landmark on a specimen (NEW in R5)
type SpecimenFeature struct {
	common.BackboneElement

	// Highlighted feature
	Type common.CodeableConcept `json:"type"`

	// Information about the feature
	Description string `json:"description"`
}

// SpecimenCollection represents details concerning the specimen collection (enhanced in R5)
type SpecimenCollection struct {
	common.BackboneElement

	// Who collected the specimen
	Collector *common.Reference `json:"collector,omitempty"`

	// Collection time
	CollectedDateTime *time.Time     `json:"collectedDateTime,omitempty"`
	CollectedPeriod   *common.Period `json:"collectedPeriod,omitempty"`

	// The span of time over which the collection of a specimen occurred
	Duration *Duration `json:"duration,omitempty"`

	// How much was collected
	Quantity *common.Quantity `json:"quantity,omitempty"`

	// Technique used to perform collection
	Method *common.CodeableConcept `json:"method,omitempty"`

	// A coded value specifying the technique that is used to perform the procedure (NEW in R5)
	Device *CodeableReference `json:"device,omitempty"`

	// The procedure event during which the specimen was collected (NEW in R5)
	Procedure *common.Reference `json:"procedure,omitempty"`

	// Anatomical collection site (enhanced in R5)
	BodySite *CodeableReference `json:"bodySite,omitempty"`

	// Whether or how long patient abstained from food and/or drink
	FastingStatusCodeableConcept *common.CodeableConcept `json:"fastingStatusCodeableConcept,omitempty"`
	FastingStatusDuration        *Duration               `json:"fastingStatusDuration,omitempty"`
}

// SpecimenProcessing represents processing and processing steps for the specimen (enhanced in R5)
type SpecimenProcessing struct {
	common.BackboneElement

	// Textual description of procedure
	Description *string `json:"description,omitempty"`

	// Indicates the treatment or processing step applied to the specimen (enhanced in R5)
	Method *common.CodeableConcept `json:"method,omitempty"`

	// Material used in the processing step
	Additive []common.Reference `json:"additive,omitempty"`

	// Date and time of specimen processing
	TimeDateTime *time.Time     `json:"timeDateTime,omitempty"`
	TimePeriod   *common.Period `json:"timePeriod,omitempty"`
}

// SpecimenContainer represents the container holding the specimen
type SpecimenContainer struct {
	common.BackboneElement

	// A list of different types of identifier
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
