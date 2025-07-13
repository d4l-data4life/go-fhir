// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// EncounterStatus represents the status of an encounter
type EncounterStatus string

const (
	EncounterStatusPlanned        EncounterStatus = "planned"
	EncounterStatusInProgress     EncounterStatus = "in-progress"
	EncounterStatusOnHold         EncounterStatus = "on-hold"
	EncounterStatusDischarged     EncounterStatus = "discharged"
	EncounterStatusCompleted      EncounterStatus = "completed"
	EncounterStatusCancelled      EncounterStatus = "cancelled"
	EncounterStatusDiscontinued   EncounterStatus = "discontinued"
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

	// For planning purposes, Appointments may include a CareTeam participant
	Actor *common.Reference `json:"actor,omitempty"`

	// The period of time that the specified participant participated in the encounter
	Period *common.Period `json:"period,omitempty"`

	// The participant type indicates how an individual actor participates in an encounter
	Type []common.CodeableConcept `json:"type,omitempty"`
}

// EncounterReason represents the reason for the encounter
type EncounterReason struct {
	common.BackboneElement

	// What the reason value should be used as e.g. Chief Complaint, Health Concern, Health Maintenance
	Use []common.CodeableConcept `json:"use,omitempty"`

	// Reason the encounter takes place, expressed as a code or a reference to another resource
	Value []CodeableReference `json:"value,omitempty"`
}

// EncounterDiagnosis represents diagnoses relevant to this encounter
type EncounterDiagnosis struct {
	common.BackboneElement

	// The coded diagnosis or a reference to a Condition
	Condition []CodeableReference `json:"condition,omitempty"`

	// Role that this diagnosis has within the encounter (e.g. admission, billing, discharge)
	Use []common.CodeableConcept `json:"use,omitempty"`
}

// EncounterAdmission represents admission details (replaces Hospitalization in R5)
type EncounterAdmission struct {
	common.BackboneElement

	// From where patient was admitted (physician referral, transfer)
	AdmitSource *common.CodeableConcept `json:"admitSource,omitempty"`

	// Location/organization to which the patient is discharged
	Destination *common.Reference `json:"destination,omitempty"`

	// Category or kind of location after discharge
	DischargeDisposition *common.CodeableConcept `json:"dischargeDisposition,omitempty"`

	// The location/organization from which the patient came before admission
	Origin *common.Reference `json:"origin,omitempty"`

	// Pre-admission identifier
	PreAdmissionIdentifier *common.Identifier `json:"preAdmissionIdentifier,omitempty"`

	// Indicates that this encounter is directly related to a prior admission
	ReAdmission *common.CodeableConcept `json:"reAdmission,omitempty"`
}

// EncounterLocation represents locations where the patient has been during this encounter
type EncounterLocation struct {
	common.BackboneElement

	// This information is de-normalized from the Location resource
	Form *common.CodeableConcept `json:"form,omitempty"`

	// The location where the encounter takes place
	Location common.Reference `json:"location"`

	// Time period during which the patient was present at the location
	Period *common.Period `json:"period,omitempty"`

	// When the patient is no longer active at a location
	Status *EncounterLocationStatus `json:"status,omitempty"`
}

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

// Encounter represents an interaction between a patient and healthcare provider(s)
type Encounter struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Encounter"

	// The billing system may choose to allocate billable items associated with the Encounter
	Account []common.Reference `json:"account,omitempty"`

	// If not (yet) known, the end of the Period may be omitted
	ActualPeriod *common.Period `json:"actualPeriod,omitempty"`

	// An Encounter may cover more than just the inpatient stay
	Admission *EncounterAdmission `json:"admission,omitempty"`

	// The appointment that scheduled this encounter
	Appointment []common.Reference `json:"appointment,omitempty"`

	// The request this encounter satisfies (e.g. incoming referral or procedure request)
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// The group(s) of individuals, organizations that are allocated to participate in this encounter
	CareTeam []common.Reference `json:"careTeam,omitempty"`

	// Concepts representing classification of patient encounter
	Class []common.CodeableConcept `json:"class,omitempty"`

	// Also note that for the purpose of billing, the diagnoses are recorded in the account
	Diagnosis []EncounterDiagnosis `json:"diagnosis,omitempty"`

	// For example, a patient may request both a dairy-free and nut-free diet preference
	DietPreference []common.CodeableConcept `json:"dietPreference,omitempty"`

	// Where a specific encounter should be classified as a part of a specific episode(s) of care
	EpisodeOfCare []common.Reference `json:"episodeOfCare,omitempty"`

	// Identifier(s) by which this encounter is known
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// If the precision on these values is low then this may be considered was an all day encounter
	Length *Duration `json:"length,omitempty"`

	// Virtual encounters can be recorded in the Encounter by specifying a location reference
	Location []EncounterLocation `json:"location,omitempty"`

	// Any Patient or Group present in the participation.actor must also be the subject
	Participant []EncounterParticipant `json:"participant,omitempty"`

	// This is also used for associating a child's encounter back to the mother's encounter
	PartOf *common.Reference `json:"partOf,omitempty"`

	// The planned end date/time (or discharge date) of the encounter
	PlannedEndDate *string `json:"plannedEndDate,omitempty"`

	// The planned start date/time (or admission date) of the encounter
	PlannedStartDate *string `json:"plannedStartDate,omitempty"`

	// Indicates the urgency of the encounter
	Priority *common.CodeableConcept `json:"priority,omitempty"`

	// The reason communicates what medical problem the patient has
	Reason []EncounterReason `json:"reason,omitempty"`

	// The organization that is primarily responsible for this Encounter's services
	ServiceProvider *common.Reference `json:"serviceProvider,omitempty"`

	// Broad categorization of the service that is to be provided (e.g. cardiology)
	ServiceType []CodeableReference `json:"serviceType,omitempty"`

	// Any special requests that have been made for this encounter
	SpecialArrangement []common.CodeableConcept `json:"specialArrangement,omitempty"`

	// Any special requests that have been made for this encounter
	SpecialCourtesy []common.CodeableConcept `json:"specialCourtesy,omitempty"`

	// Note that internal business rules will determine the appropriate transitions
	Status EncounterStatus `json:"status"`

	// While the encounter is always about the patient, the patient might not actually be known
	Subject *common.Reference `json:"subject,omitempty"`

	// Different use-cases are likely to have different permitted transitions between states
	SubjectStatus *common.CodeableConcept `json:"subjectStatus,omitempty"`

	// Since there are many ways to further classify encounters, this element is 0..*
	Type []common.CodeableConcept `json:"type,omitempty"`

	// There are two types of virtual meetings that often exist
	VirtualService []VirtualServiceDetail `json:"virtualService,omitempty"`
}

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

// PractitionerCommunication represents communication details for a practitioner in R5
type PractitionerCommunication struct {
	common.BackboneElement

	// The language which can be used to communicate with the practitioner
	Language common.CodeableConcept `json:"language"`

	// Language preference indicator
	Preferred *bool `json:"preferred,omitempty"`
}

// PractitionerQualification represents qualifications obtained by training and certification
type PractitionerQualification struct {
	common.BackboneElement

	// Coded representation of the qualification
	Code common.CodeableConcept `json:"code"`

	// An identifier for this qualification for the practitioner
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Organization that regulates and issues the qualification
	Issuer *common.Reference `json:"issuer,omitempty"`

	// Period during which the qualification is valid
	Period *common.Period `json:"period,omitempty"`
}

// Practitioner represents a person who is directly or indirectly involved in the provisioning of healthcare or related services
type Practitioner struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Practitioner"

	// Whether this practitioner's record is in active use
	Active *bool `json:"active,omitempty"`

	// Addresses where the practitioner can be found or visited or to which mail can be delivered
	Address []Address `json:"address,omitempty"`

	// The date of birth for the practitioner
	BirthDate *string `json:"birthDate,omitempty"`

	// A language the practitioner can use in patient communication (R5 uses structured communication)
	Communication []PractitionerCommunication `json:"communication,omitempty"`

	// Indicates if the practitioner is deceased (R5 feature)
	DeceasedBoolean  *bool   `json:"deceasedBoolean,omitempty"`
	DeceasedDateTime *string `json:"deceasedDateTime,omitempty"`

	// Administrative Gender - the gender that the person is considered to have for administration and record keeping purposes
	Gender *PractitionerGender `json:"gender,omitempty"`

	// An identifier that applies to this person in this role
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

// PractitionerGender represents the gender of a practitioner
type PractitionerGender string

const (
	PractitionerGenderMale    PractitionerGender = "male"
	PractitionerGenderFemale  PractitionerGender = "female"
	PractitionerGenderOther   PractitionerGender = "other"
	PractitionerGenderUnknown PractitionerGender = "unknown"
)

// OrganizationQualification represents official certifications and accreditations
type OrganizationQualification struct {
	common.BackboneElement

	// Coded representation of the qualification
	Code common.CodeableConcept `json:"code"`

	// An identifier allocated to this qualification for this organization
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Organization that regulates and issues the qualification
	Issuer *common.Reference `json:"issuer,omitempty"`

	// Period during which the qualification is valid
	Period *common.Period `json:"period,omitempty"`
}

// Organization represents a formally or informally recognized grouping of people or organizations (R5)
type Organization struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Organization"

	// Whether the organization record is still in active use
	Active *bool `json:"active,omitempty"`

	// A list of alternate names that the organization is known as, or was known as in the past
	Alias []string `json:"alias,omitempty"`

	// Official contact details for the organization
	Contact []ExtendedContactDetail `json:"contact,omitempty"`

	// Additional details about the organization that could be displayed as further information
	Description *string `json:"description,omitempty"`

	// Technical endpoints providing access to services operated for the organization
	Endpoint []common.Reference `json:"endpoint,omitempty"`

	// Identifier for the organization that is used to identify the organization across multiple disparate systems
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Name used for the organization
	Name *string `json:"name,omitempty"`

	// The organization of which this organization forms a part
	PartOf *common.Reference `json:"partOf,omitempty"`

	// Official certifications, accreditations, training, designations and licenses
	Qualification []OrganizationQualification `json:"qualification,omitempty"`

	// Kind of organization
	Type []common.CodeableConcept `json:"type,omitempty"`
}

// ExtendedContactDetail represents extended contact information (R5)
type ExtendedContactDetail struct {
	common.Element

	// The type of contact
	Purpose *common.CodeableConcept `json:"purpose,omitempty"`

	// Name of an individual to contact
	Name []HumanName `json:"name,omitempty"`

	// Contact details (telephone, email, etc.) for a contact
	Telecom []ContactPoint `json:"telecom,omitempty"`

	// Address for the contact
	Address *Address `json:"address,omitempty"`

	// The organization that is associated with the contact
	Organization *common.Reference `json:"organization,omitempty"`

	// Period that this contact was valid for usage
	Period *common.Period `json:"period,omitempty"`
}

// ConditionParticipant represents people involved in the condition
type ConditionParticipant struct {
	common.BackboneElement

	// Indicates who or what participated in the activities related to the condition
	Actor common.Reference `json:"actor"`

	// Distinguishes the type of involvement of the actor in the activities related to the condition
	Function *common.CodeableConcept `json:"function,omitempty"`
}

// ConditionStage represents clinical stage or grade of a condition
type ConditionStage struct {
	common.BackboneElement

	// Reference to a formal record of the evidence on which the staging assessment is based
	Assessment []common.Reference `json:"assessment,omitempty"`

	// A simple summary of the stage such as "Stage 3" or "Early Onset"
	Summary *common.CodeableConcept `json:"summary,omitempty"`

	// The kind of staging, such as pathological or clinical staging
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// Condition represents a clinical condition, problem, diagnosis, or other event (R5)
type Condition struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Condition"

	// The date or estimated date that the condition resolved or went into remission
	AbatementDateTime *string        `json:"abatementDateTime,omitempty"`
	AbatementAge      *Age           `json:"abatementAge,omitempty"`
	AbatementPeriod   *common.Period `json:"abatementPeriod,omitempty"`
	AbatementRange    *Range         `json:"abatementRange,omitempty"`
	AbatementString   *string        `json:"abatementString,omitempty"`

	// The anatomical location where this condition manifests itself
	BodySite []common.CodeableConcept `json:"bodySite,omitempty"`

	// A category assigned to the condition
	Category []common.CodeableConcept `json:"category,omitempty"`

	// The clinical status of the condition (required in R5)
	ClinicalStatus common.CodeableConcept `json:"clinicalStatus"`

	// Identification of the condition, problem or diagnosis
	Code *common.CodeableConcept `json:"code,omitempty"`

	// Encounter during which the condition was first asserted
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Supporting evidence/manifestations that are the basis of this condition (R5 uses CodeableReference)
	Evidence []CodeableReference `json:"evidence,omitempty"`

	// Business identifier for this condition
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Additional information about the Condition
	Note []Annotation `json:"note,omitempty"`

	// Estimated or actual date or date-time the condition began
	OnsetDateTime *string        `json:"onsetDateTime,omitempty"`
	OnsetAge      *Age           `json:"onsetAge,omitempty"`
	OnsetPeriod   *common.Period `json:"onsetPeriod,omitempty"`
	OnsetRange    *Range         `json:"onsetRange,omitempty"`
	OnsetString   *string        `json:"onsetString,omitempty"`

	// Indicates who or what participated in the activities related to the condition (R5 feature)
	Participant []ConditionParticipant `json:"participant,omitempty"`

	// The recordedDate represents when this particular Condition record was created in the system
	RecordedDate *string `json:"recordedDate,omitempty"`

	// A subjective assessment of the severity of the condition
	Severity *common.CodeableConcept `json:"severity,omitempty"`

	// Clinical stage or grade of a condition
	Stage []ConditionStage `json:"stage,omitempty"`

	// Indicates the patient or group who the condition record is associated with
	Subject common.Reference `json:"subject"`

	// The verification status to support the clinical status of the condition
	VerificationStatus *common.CodeableConcept `json:"verificationStatus,omitempty"`
}

// MedicationIngredient represents medication ingredients
type MedicationIngredient struct {
	common.BackboneElement

	// Indication of whether this ingredient affects the therapeutic action of the drug
	IsActive *bool `json:"isActive,omitempty"`

	// The ingredient (substance or medication) that the ingredient.strength relates to
	Item CodeableReference `json:"item"`

	// Specifies how many (or how much) of the items there are in this Medication
	StrengthRatio           *Ratio                  `json:"strengthRatio,omitempty"`
	StrengthCodeableConcept *common.CodeableConcept `json:"strengthCodeableConcept,omitempty"`
	StrengthQuantity        *common.Quantity        `json:"strengthQuantity,omitempty"`
}

// MedicationBatch represents information that only applies to packages (not products)
type MedicationBatch struct {
	common.BackboneElement

	// When this specific batch of product will expire
	ExpirationDate *string `json:"expirationDate,omitempty"`

	// The assigned lot number of a batch of the specified product
	LotNumber *string `json:"lotNumber,omitempty"`
}

// Medication represents a medication definition (R5)
type Medication struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Medication"

	// Information that only applies to packages (not products)
	Batch *MedicationBatch `json:"batch,omitempty"`

	// A code that specifies this medication
	Code *common.CodeableConcept `json:"code,omitempty"`

	// A reference to a knowledge resource that provides more information about this medication
	Definition *common.Reference `json:"definition,omitempty"`

	// The dose form for a medication (R5 uses doseForm instead of form)
	DoseForm *common.CodeableConcept `json:"doseForm,omitempty"`

	// Business identifier for this medication
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Active or inactive ingredients
	Ingredient []MedicationIngredient `json:"ingredient,omitempty"`

	// Marketing authorization holder
	MarketingAuthorizationHolder *common.Reference `json:"marketingAuthorizationHolder,omitempty"`

	// Active | inactive | entered-in-error
	Status *MedicationStatus `json:"status,omitempty"`

	// When the specified product code does not infer a package size, this is the specific amount of drug in the product
	TotalVolume *common.Quantity `json:"totalVolume,omitempty"`
}

// MedicationStatus represents the status of a medication
type MedicationStatus string

const (
	MedicationStatusActive         MedicationStatus = "active"
	MedicationStatusInactive       MedicationStatus = "inactive"
	MedicationStatusEnteredInError MedicationStatus = "entered-in-error"
)

// MedicationRequestDispenseRequestInitialFill represents initial fill details
type MedicationRequestDispenseRequestInitialFill struct {
	common.BackboneElement

	// The length of time that the first dispense is expected to last
	Duration *Duration `json:"duration,omitempty"`

	// The amount or quantity to provide as part of the first dispense
	Quantity *common.Quantity `json:"quantity,omitempty"`
}

// MedicationRequestDispenseRequest represents dispense details
type MedicationRequestDispenseRequest struct {
	common.BackboneElement

	// The minimum period of time that must occur between dispenses of the medication
	DispenseInterval *Duration `json:"dispenseInterval,omitempty"`

	// Indicates the intended performing Organization that will dispense the medication
	Dispenser *common.Reference `json:"dispenser,omitempty"`

	// Provides additional information to the dispenser
	DispenserInstruction []Annotation `json:"dispenserInstruction,omitempty"`

	// Provides information about the type of adherence packaging to be supplied
	DoseAdministrationAid *common.CodeableConcept `json:"doseAdministrationAid,omitempty"`

	// Expected supply duration
	ExpectedSupplyDuration *Duration `json:"expectedSupplyDuration,omitempty"`

	// Initial fill details
	InitialFill *MedicationRequestDispenseRequestInitialFill `json:"initialFill,omitempty"`

	// If displaying "number of authorized fills", add 1 to this number
	NumberOfRepeatsAllowed *int `json:"numberOfRepeatsAllowed,omitempty"`

	// The amount that is to be dispensed for one fill
	Quantity *common.Quantity `json:"quantity,omitempty"`

	// Prescription validity period
	ValidityPeriod *common.Period `json:"validityPeriod,omitempty"`
}

// MedicationRequestSubstitution represents substitution permissions
type MedicationRequestSubstitution struct {
	common.BackboneElement

	// Whether substitution is allowed
	AllowedBoolean         *bool                   `json:"allowedBoolean,omitempty"`
	AllowedCodeableConcept *common.CodeableConcept `json:"allowedCodeableConcept,omitempty"`

	// Indicates the reason for the substitution
	Reason *common.CodeableConcept `json:"reason,omitempty"`
}

// MedicationRequest represents an order or request for medication (R5)
type MedicationRequest struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "MedicationRequest"

	// The date when the prescription was initially written or authored on
	AuthoredOn *string `json:"authoredOn,omitempty"`

	// A plan or request that is fulfilled in whole or in part by this medication request
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// The category can be used to include where the medication is expected to be consumed
	Category []common.CodeableConcept `json:"category,omitempty"`

	// Course of therapy type
	CourseOfTherapyType *common.CodeableConcept `json:"courseOfTherapyType,omitempty"`

	// The intended type of device that is to be used for the administration
	Device []CodeableReference `json:"device,omitempty"`

	// Specific details for the dispense or medication supply part
	DispenseRequest *MedicationRequestDispenseRequest `json:"dispenseRequest,omitempty"`

	// If do not perform is not specified, the request is a positive request
	DoNotPerform *bool `json:"doNotPerform,omitempty"`

	// How to take the medication
	DosageInstruction []Dosage `json:"dosageInstruction,omitempty"`

	// The period over which the medication is to be taken (R5 feature)
	EffectiveDosePeriod *common.Period `json:"effectiveDosePeriod,omitempty"`

	// Encounter during which the request was created
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Event history references
	EventHistory []common.Reference `json:"eventHistory,omitempty"`

	// Shared identifier for multiple requests
	GroupIdentifier *common.Identifier `json:"groupIdentifier,omitempty"`

	// Business identifier
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Who provided the information about this request (R5 feature)
	InformationSource []common.Reference `json:"informationSource,omitempty"`

	// Associated insurance coverage
	Insurance []common.Reference `json:"insurance,omitempty"`

	// proposal | plan | order | original-order | reflex-order | filler-order | instance-order | option
	Intent MedicationRequestIntent `json:"intent"`

	// Medication to be taken (R5 uses CodeableReference)
	Medication CodeableReference `json:"medication"`

	// Extra information about the prescription
	Note []Annotation `json:"note,omitempty"`

	// The specified desired performer of the medication treatment (R5 allows multiple)
	Performer []common.Reference `json:"performer,omitempty"`

	// Performer type requirements
	PerformerType *common.CodeableConcept `json:"performerType,omitempty"`

	// routine | urgent | asap | stat
	Priority *MedicationRequestPriority `json:"priority,omitempty"`

	// Reference to an order/prescription that is being replaced
	PriorPrescription *common.Reference `json:"priorPrescription,omitempty"`

	// The reason for the order (R5 uses CodeableReference)
	Reason []CodeableReference `json:"reason,omitempty"`

	// The person who entered the order on behalf of another individual
	Recorder *common.Reference `json:"recorder,omitempty"`

	// Full representation of the dose of the medication (R5 feature)
	RenderedDosageInstruction *string `json:"renderedDosageInstruction,omitempty"`

	// If not populated, then assume that this is the original record
	Reported *bool `json:"reported,omitempty"`

	// The individual, organization, or device that initiated the request
	Requester *common.Reference `json:"requester,omitempty"`

	// active | on-hold | ended | stopped | completed | cancelled | entered-in-error | draft | unknown
	Status MedicationRequestStatus `json:"status"`

	// The date when the status was changed (R5 feature)
	StatusChanged *string `json:"statusChanged,omitempty"`

	// Reason for current status
	StatusReason *common.CodeableConcept `json:"statusReason,omitempty"`

	// The subject on a medication request is mandatory
	Subject common.Reference `json:"subject"`

	// Whether or not substitution can or should be part of the dispense
	Substitution *MedicationRequestSubstitution `json:"substitution,omitempty"`

	// Additional information that supports the ordering of the medication
	SupportingInformation []common.Reference `json:"supportingInformation,omitempty"`
}

// MedicationRequestIntent represents the intent of the medication request
type MedicationRequestIntent string

const (
	MedicationRequestIntentProposal      MedicationRequestIntent = "proposal"
	MedicationRequestIntentPlan          MedicationRequestIntent = "plan"
	MedicationRequestIntentOrder         MedicationRequestIntent = "order"
	MedicationRequestIntentOriginalOrder MedicationRequestIntent = "original-order"
	MedicationRequestIntentReflexOrder   MedicationRequestIntent = "reflex-order"
	MedicationRequestIntentFillerOrder   MedicationRequestIntent = "filler-order"
	MedicationRequestIntentInstanceOrder MedicationRequestIntent = "instance-order"
	MedicationRequestIntentOption        MedicationRequestIntent = "option"
)

// MedicationRequestStatus represents the status of the medication request
type MedicationRequestStatus string

const (
	MedicationRequestStatusActive         MedicationRequestStatus = "active"
	MedicationRequestStatusOnHold         MedicationRequestStatus = "on-hold"
	MedicationRequestStatusEnded          MedicationRequestStatus = "ended"
	MedicationRequestStatusStopped        MedicationRequestStatus = "stopped"
	MedicationRequestStatusCompleted      MedicationRequestStatus = "completed"
	MedicationRequestStatusCancelled      MedicationRequestStatus = "cancelled"
	MedicationRequestStatusEnteredInError MedicationRequestStatus = "entered-in-error"
	MedicationRequestStatusDraft          MedicationRequestStatus = "draft"
	MedicationRequestStatusUnknown        MedicationRequestStatus = "unknown"
)

// MedicationRequestPriority represents the priority of the medication request
type MedicationRequestPriority string

const (
	MedicationRequestPriorityRoutine MedicationRequestPriority = "routine"
	MedicationRequestPriorityUrgent  MedicationRequestPriority = "urgent"
	MedicationRequestPriorityASAP    MedicationRequestPriority = "asap"
	MedicationRequestPriorityStat    MedicationRequestPriority = "stat"
)

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

// AccountStatus represents the status of an account
type AccountStatus string

const (
	AccountStatusActive         AccountStatus = "active"
	AccountStatusInactive       AccountStatus = "inactive"
	AccountStatusEnteredInError AccountStatus = "entered-in-error"
	AccountStatusOnHold         AccountStatus = "on-hold"
	AccountStatusUnknown        AccountStatus = "unknown"
)

// AccountCoverage represents coverage for an account
type AccountCoverage struct {
	common.BackboneElement

	// The party(s) that contribute to payment (or part of) of the charges applied to this account
	Coverage common.Reference `json:"coverage"`

	// The priority of the coverage in the context of this account
	Priority *int `json:"priority,omitempty"`
}

// AccountGuarantor represents parties responsible for balancing the account
type AccountGuarantor struct {
	common.BackboneElement

	// A guarantor may be placed on credit hold or otherwise have their role temporarily suspended
	OnHold *bool `json:"onHold,omitempty"`

	// The entity who is responsible
	Party common.Reference `json:"party"`

	// The timeframe during which the guarantor accepts responsibility for the account
	Period *common.Period `json:"period,omitempty"`
}

// AccountDiagnosis represents diagnoses relevant to the account
type AccountDiagnosis struct {
	common.BackboneElement

	// The diagnosis relevant to the account
	Condition CodeableReference `json:"condition"`

	// Date of the diagnosis
	DateOfDiagnosis *string `json:"dateOfDiagnosis,omitempty"`

	// Was the diagnosis present on admission
	OnAdmission *bool `json:"onAdmission,omitempty"`

	// Package code
	PackageCode []common.CodeableConcept `json:"packageCode,omitempty"`

	// Ranking of the diagnosis
	Sequence *int `json:"sequence,omitempty"`

	// Type that this diagnosis has relevant to the account
	Type []common.CodeableConcept `json:"type,omitempty"`
}

// AccountProcedure represents procedures relevant to the account
type AccountProcedure struct {
	common.BackboneElement

	// The procedure relevant to the account
	Code CodeableReference `json:"code"`

	// Date of the procedure
	DateOfService *string `json:"dateOfService,omitempty"`

	// Devices associated with the procedure
	Device []common.Reference `json:"device,omitempty"`

	// Package code
	PackageCode []common.CodeableConcept `json:"packageCode,omitempty"`

	// Ranking of the procedure
	Sequence *int `json:"sequence,omitempty"`

	// How this procedure value should be used in charging the account
	Type []common.CodeableConcept `json:"type,omitempty"`
}

// AccountRelatedAccount represents related accounts
type AccountRelatedAccount struct {
	common.BackboneElement

	// Reference to an associated Account
	Account common.Reference `json:"account"`

	// Relationship of the associated Account
	Relationship *common.CodeableConcept `json:"relationship,omitempty"`
}

// Money represents an amount of economic utility in some recognized currency
type Money struct {
	common.Element

	// Numerical value (with implicit precision)
	Value *float64 `json:"value,omitempty"`

	// ISO 4217 Currency Code
	Currency *string `json:"currency,omitempty"`
}

// AccountBalance represents calculated account balances
type AccountBalance struct {
	common.BackboneElement

	// Who is expected to pay this part of the balance
	Aggregate *common.CodeableConcept `json:"aggregate,omitempty"`

	// The actual balance value calculated
	Amount Money `json:"amount"`

	// The amount is only an estimated value
	Estimate *bool `json:"estimate,omitempty"`

	// The term of the account balances
	Term *common.CodeableConcept `json:"term,omitempty"`
}

// Account represents a financial tool for tracking value accrued for a particular purpose
type Account struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Account"

	// The calculated account balances
	Balance []AccountBalance `json:"balance,omitempty"`

	// The BillingStatus tracks the lifecycle of the account
	BillingStatus *common.CodeableConcept `json:"billingStatus,omitempty"`

	// Time the balances were calculated
	CalculatedAt *string `json:"calculatedAt,omitempty"`

	// The coverage accounts to be used for billing
	Coverage []AccountCoverage `json:"coverage,omitempty"`

	// The default currency for the account
	Currency *common.CodeableConcept `json:"currency,omitempty"`

	// Explanation of purpose/use
	Description *string `json:"description,omitempty"`

	// The set of diagnoses that are relevant for billing
	Diagnosis []AccountDiagnosis `json:"diagnosis,omitempty"`

	// The parties responsible for balancing the account
	Guarantor []AccountGuarantor `json:"guarantor,omitempty"`

	// Unique identifier used to reference the account
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Human-readable label
	Name *string `json:"name,omitempty"`

	// Entity managing the Account
	Owner *common.Reference `json:"owner,omitempty"`

	// The set of procedures that are relevant for billing
	Procedure []AccountProcedure `json:"procedure,omitempty"`

	// Other associated accounts related to this account
	RelatedAccount []AccountRelatedAccount `json:"relatedAccount,omitempty"`

	// Transaction window
	ServicePeriod *common.Period `json:"servicePeriod,omitempty"`

	// active | inactive | entered-in-error | on-hold | unknown
	Status AccountStatus `json:"status"`

	// The entity that caused the expenses
	Subject []common.Reference `json:"subject,omitempty"`

	// E.g. patient, expense, depreciation
	Type *common.CodeableConcept `json:"type,omitempty"`
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

// ActivityDefinitionParticipant represents who should participate in performing the action described
type ActivityDefinitionParticipant struct {
	common.BackboneElement

	// Indicates how the actor will be involved in the action
	Function *common.CodeableConcept `json:"function,omitempty"`

	// The role the participant should play in performing the described action
	Role *common.CodeableConcept `json:"role,omitempty"`

	// The type of participant in the action
	Type *ParticipantType `json:"type,omitempty"`

	// The type of participant in the action (canonical)
	TypeCanonical *string `json:"typeCanonical,omitempty"`

	// When this element is a reference, it SHOULD be a reference to a definitional resource
	TypeReference *common.Reference `json:"typeReference,omitempty"`
}

// ActivityDefinitionDynamicValue represents dynamic values that are applied when the activity is performed
type ActivityDefinitionDynamicValue struct {
	common.BackboneElement

	// The expression may be inlined, or may be a reference to a named expression within a logic library
	Expression Expression `json:"expression"`

	// The path attribute contains a Simple FHIRPath Subset that allows path traversal
	Path string `json:"path"`
}

// ActivityDefinition represents a definition of an activity to be performed
type ActivityDefinition struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "ActivityDefinition"

	// The date on which the resource content was approved by the publisher
	ApprovalDate *string `json:"approvalDate,omitempty"`

	// If a CodeableConcept is present, it indicates the pre-condition for performing the service
	AsNeededBoolean         *bool                   `json:"asNeededBoolean,omitempty"`
	AsNeededCodeableConcept *common.CodeableConcept `json:"asNeededCodeableConcept,omitempty"`

	// An individual or organization primarily involved in the creation and maintenance of the content
	Author []ContactDetail `json:"author,omitempty"`

	// Only used if not implicit in the code found in ServiceRequest.type
	BodySite []common.CodeableConcept `json:"bodySite,omitempty"`

	// Tends to be less relevant for activities involving particular products
	Code *common.CodeableConcept `json:"code,omitempty"`

	// May be a web site, an email address, a telephone number, etc.
	Contact []ContactDetail `json:"contact,omitempty"`

	// The short copyright declaration
	Copyright *string `json:"copyright,omitempty"`

	// The (c) symbol should NOT be included in this string
	CopyrightLabel *string `json:"copyrightLabel,omitempty"`

	// The date (and optionally time) when the activity definition was published
	Date *string `json:"date,omitempty"`

	// Natural language description of the activity definition
	Description *string `json:"description,omitempty"`

	// This element is not intended to be used to communicate a decision support response to cancel an order in progress
	DoNotPerform *bool `json:"doNotPerform,omitempty"`

	// If a dosage instruction is used, the definition should not specify timing or quantity
	Dosage []Dosage `json:"dosage,omitempty"`

	// Dynamic values are applied in the order in which they are defined in the ActivityDefinition
	DynamicValue []ActivityDefinitionDynamicValue `json:"dynamicValue,omitempty"`

	// An individual or organization primarily responsible for internal coherence of the content
	Editor []ContactDetail `json:"editor,omitempty"`

	// The effective period for an activity definition
	EffectivePeriod *common.Period `json:"effectivePeriod,omitempty"`

	// See guidance around (not) making local changes to elements
	Endorser []ContactDetail `json:"endorser,omitempty"`

	// For testing purposes, not real usage
	Experimental *bool `json:"experimental,omitempty"`

	// Additional identifier for the activity definition
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Indicates the level of authority/intentionality associated with the activity
	Intent *RequestIntent `json:"intent,omitempty"`

	// Intended jurisdiction for activity definition (if applicable)
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`

	// The kind element may only specify Request resource types
	Kind *string `json:"kind,omitempty"`

	// When the activity definition was last reviewed
	LastReviewDate *string `json:"lastReviewDate,omitempty"`

	// A reference to a Library resource containing any formal logic used by the activity definition
	Library []string `json:"library,omitempty"`

	// May reference a specific clinical location or may just identify a type of location
	Location *CodeableReference `json:"location,omitempty"`

	// Name for this activity definition (computer friendly)
	Name *string `json:"name,omitempty"`

	// Defines observation requirements for the action to be performed
	ObservationRequirement []string `json:"observationRequirement,omitempty"`

	// Defines the observations that are expected to be produced by the action
	ObservationResultRequirement []string `json:"observationResultRequirement,omitempty"`

	// Indicates who should participate in performing the action described
	Participant []ActivityDefinitionParticipant `json:"participant,omitempty"`

	// Indicates how quickly the activity should be addressed with respect to other requests
	Priority *RequestPriority `json:"priority,omitempty"`

	// Identifies the food, drug or other product being consumed or supplied in the activity
	ProductReference       *common.Reference       `json:"productReference,omitempty"`
	ProductCodeableConcept *common.CodeableConcept `json:"productCodeableConcept,omitempty"`

	// A profile to which the target of the activity definition is expected to conform
	Profile *string `json:"profile,omitempty"`

	// Name of the publisher (organization or individual)
	Publisher *string `json:"publisher,omitempty"`

	// Why this activity definition is defined
	Purpose *string `json:"purpose,omitempty"`

	// Identifies the quantity expected to be consumed at once
	Quantity *common.Quantity `json:"quantity,omitempty"`

	// Each related artifact is either an attachment, or a reference to another resource
	RelatedArtifact []RelatedArtifact `json:"relatedArtifact,omitempty"`

	// See guidance around (not) making local changes to elements
	Reviewer []ContactDetail `json:"reviewer,omitempty"`

	// Defines specimen requirements for the action to be performed
	SpecimenRequirement []string `json:"specimenRequirement,omitempty"`

	// draft | active | retired | unknown
	Status RequestStatus `json:"status"`

	// Type of individual the activity definition is intended for
	SubjectCodeableConcept *common.CodeableConcept `json:"subjectCodeableConcept,omitempty"`
	SubjectReference       *common.Reference       `json:"subjectReference,omitempty"`
	SubjectCanonical       *string                 `json:"subjectCanonical,omitempty"`

	// An explanatory or alternate title for the activity definition
	Subtitle *string `json:"subtitle,omitempty"`

	// The intent of the timing element is to provide timing when the action should be performed
	TimingTiming   *Timing   `json:"timingTiming,omitempty"`
	TimingAge      *Age      `json:"timingAge,omitempty"`
	TimingRange    *Range    `json:"timingRange,omitempty"`
	TimingDuration *Duration `json:"timingDuration,omitempty"`

	// Name for this activity definition (human friendly)
	Title *string `json:"title,omitempty"`

	// E.g. Education, Treatment, Assessment, etc.
	Topic []common.CodeableConcept `json:"topic,omitempty"`

	// Note that if both a transform and dynamic values are specified, the dynamic values will be applied to the result of the transform
	Transform *string `json:"transform,omitempty"`

	// Canonical identifier for this activity definition
	URL *string `json:"url,omitempty"`

	// A detailed description of how the activity definition is used from a clinical perspective
	Usage *string `json:"usage,omitempty"`

	// The context that the content is intended to support
	UseContext []UsageContext `json:"useContext,omitempty"`

	// Business version of the activity definition
	Version *string `json:"version,omitempty"`

	// How to compare versions
	VersionAlgorithmString *string        `json:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *common.Coding `json:"versionAlgorithmCoding,omitempty"`
}

// ActorType represents whether the actor is a person or system
type ActorType string

const (
	ActorTypePerson ActorType = "person"
	ActorTypeSystem ActorType = "system"
)

// ActorDefinition represents a description of an actor
type ActorDefinition struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "ActorDefinition"

	// The capability statement for the actor (if the concept is applicable)
	Capabilities *string `json:"capabilities,omitempty"`

	// May be a web site, an email address, a telephone number, etc.
	Contact []ContactDetail `json:"contact,omitempty"`

	// Use and/or publishing restrictions
	Copyright *string `json:"copyright,omitempty"`

	// Copyright holder and year(s)
	CopyrightLabel *string `json:"copyrightLabel,omitempty"`

	// Date this was last changed
	Date *string `json:"date,omitempty"`

	// A url that identifies the definition of this actor in another IG
	DerivedFrom []string `json:"derivedFrom,omitempty"`

	// Natural language description of the actor
	Description *string `json:"description,omitempty"`

	// What the actor does (or is expected to do)
	Documentation *string `json:"documentation,omitempty"`

	// For testing purposes, not real usage
	Experimental *bool `json:"experimental,omitempty"`

	// Additional identifier for the actor definition
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Intended jurisdiction for actor definition (if applicable)
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`

	// Name for this actor definition (computer friendly)
	Name *string `json:"name,omitempty"`

	// Name of the publisher (organization or individual)
	Publisher *string `json:"publisher,omitempty"`

	// Why this actor definition is defined
	Purpose *string `json:"purpose,omitempty"`

	// A reference to additional documentation about the actor
	Reference []string `json:"reference,omitempty"`

	// draft | active | retired | unknown
	Status RequestStatus `json:"status"`

	// Name for this actor definition (human friendly)
	Title *string `json:"title,omitempty"`

	// Whether the actor represents a human or an application
	Type ActorType `json:"type"`

	// Canonical identifier for this actor definition
	URL *string `json:"url,omitempty"`

	// The context that the content is intended to support
	UseContext []UsageContext `json:"useContext,omitempty"`

	// Business version of the actor definition
	Version *string `json:"version,omitempty"`

	// How to compare versions
	VersionAlgorithmString *string        `json:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *common.Coding `json:"versionAlgorithmCoding,omitempty"`
}

// AdministrableProductDefinitionProperty represents characteristics of a product
type AdministrableProductDefinitionProperty struct {
	common.BackboneElement

	// The status of characteristic e.g. assigned or pending
	Status *common.CodeableConcept `json:"status,omitempty"`

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

// AdministrableProductDefinitionRouteOfAdministrationTargetSpeciesWithdrawalPeriod represents a species specific time during which consumption of animal product is not appropriate
type AdministrableProductDefinitionRouteOfAdministrationTargetSpeciesWithdrawalPeriod struct {
	common.BackboneElement

	// Extra information about the withdrawal period
	SupportingInformation *string `json:"supportingInformation,omitempty"`

	// Coded expression for the type of tissue for which the withdrawal period applies
	Tissue common.CodeableConcept `json:"tissue"`

	// A value for the time
	Value common.Quantity `json:"value"`
}

// AdministrableProductDefinitionRouteOfAdministrationTargetSpecies represents a species for which this route applies
type AdministrableProductDefinitionRouteOfAdministrationTargetSpecies struct {
	common.BackboneElement

	// Coded expression for the species
	Code common.CodeableConcept `json:"code"`

	// A species specific time during which consumption of animal product is not appropriate
	WithdrawalPeriod []AdministrableProductDefinitionRouteOfAdministrationTargetSpeciesWithdrawalPeriod `json:"withdrawalPeriod,omitempty"`
}

// AdministrableProductDefinitionRouteOfAdministration represents the path by which the product is taken into or makes contact with the body
type AdministrableProductDefinitionRouteOfAdministration struct {
	common.BackboneElement

	// Coded expression for the route
	Code common.CodeableConcept `json:"code"`

	// The first dose (dose quantity) administered can be specified for the product
	FirstDose *common.Quantity `json:"firstDose,omitempty"`

	// The maximum dose per day that can be administered
	MaxDosePerDay *common.Quantity `json:"maxDosePerDay,omitempty"`

	// The maximum dose per treatment period that can be administered
	MaxDosePerTreatmentPeriod *Ratio `json:"maxDosePerTreatmentPeriod,omitempty"`

	// The maximum single dose that can be administered
	MaxSingleDose *common.Quantity `json:"maxSingleDose,omitempty"`

	// The maximum treatment period during which the product can be administered
	MaxTreatmentPeriod *Duration `json:"maxTreatmentPeriod,omitempty"`

	// A species for which this route applies
	TargetSpecies []AdministrableProductDefinitionRouteOfAdministrationTargetSpecies `json:"targetSpecies,omitempty"`
}

// AdministrableProductDefinition represents a medicinal product in the final form which is suitable for administering to a patient
type AdministrableProductDefinition struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "AdministrableProductDefinition"

	// The dose form of the final product after necessary reconstitution or processing
	AdministrableDoseForm *common.CodeableConcept `json:"administrableDoseForm,omitempty"`

	// A general description of the product, when in its final form, suitable for administration
	Description *string `json:"description,omitempty"`

	// A device that is integral to the medicinal product
	Device *common.Reference `json:"device,omitempty"`

	// References a product from which one or more of the constituent parts of that product can be prepared and used
	FormOf []common.Reference `json:"formOf,omitempty"`

	// An identifier for the administrable product
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The ingredients of this administrable medicinal product
	Ingredient []common.CodeableConcept `json:"ingredient,omitempty"`

	// Indicates the specific manufactured items that are part of the 'formOf' product that are used in the preparation
	ProducedFrom []common.Reference `json:"producedFrom,omitempty"`

	// Characteristics e.g. a product's onset of action
	Property []AdministrableProductDefinitionProperty `json:"property,omitempty"`

	// The path by which the product is taken into or makes contact with the body
	RouteOfAdministration []AdministrableProductDefinitionRouteOfAdministration `json:"routeOfAdministration"`

	// draft | active | retired | unknown
	Status RequestStatus `json:"status"`

	// The presentation type in which this item is given to a patient
	UnitOfPresentation *common.CodeableConcept `json:"unitOfPresentation,omitempty"`
}

// AdverseEventActuality represents whether the event actually happened or was potential
type AdverseEventActuality string

const (
	AdverseEventActualityActual    AdverseEventActuality = "actual"
	AdverseEventActualityPotential AdverseEventActuality = "potential"
)

// AdverseEventStatus represents the status of the adverse event
type AdverseEventStatus string

const (
	AdverseEventStatusInProgress     AdverseEventStatus = "in-progress"
	AdverseEventStatusCompleted      AdverseEventStatus = "completed"
	AdverseEventStatusEnteredInError AdverseEventStatus = "entered-in-error"
	AdverseEventStatusUnknown        AdverseEventStatus = "unknown"
)

// AdverseEventParticipant represents who participated in the adverse event and how they were involved
type AdverseEventParticipant struct {
	common.BackboneElement

	// The physician prescribing a drug, a nurse administering the drug, a device that administered the drug, a witness to the event, or an informant of clinical history
	Actor common.Reference `json:"actor"`

	// Distinguishes the type of involvement of the actor in the adverse event, such as contributor or informant
	Function *common.CodeableConcept `json:"function,omitempty"`
}

// AdverseEventSuspectEntityCausality represents information on the possible cause of the event
type AdverseEventSuspectEntityCausality struct {
	common.BackboneElement

	// The method of evaluating the relatedness of the suspected entity to the event
	AssessmentMethod *common.CodeableConcept `json:"assessmentMethod,omitempty"`

	// The author of the information on the possible cause of the event
	Author *common.Reference `json:"author,omitempty"`

	// The result of the assessment regarding the relatedness of the suspected entity to the event
	EntityRelatedness *common.CodeableConcept `json:"entityRelatedness,omitempty"`
}

// AdverseEventSuspectEntity represents the entity that is suspected to have caused the adverse event
type AdverseEventSuspectEntity struct {
	common.BackboneElement

	// Information on the possible cause of the event
	Causality *AdverseEventSuspectEntityCausality `json:"causality,omitempty"`

	// Identifies the actual instance of what caused the adverse event
	InstanceCodeableConcept *common.CodeableConcept `json:"instanceCodeableConcept,omitempty"`
	InstanceReference       *common.Reference       `json:"instanceReference,omitempty"`
}

// AdverseEventContributingFactor represents the contributing factors suspected to have increased the probability or severity of the adverse event
type AdverseEventContributingFactor struct {
	common.BackboneElement

	// The item that is suspected to have increased the probability or severity of the adverse event
	ItemReference       *common.Reference       `json:"itemReference,omitempty"`
	ItemCodeableConcept *common.CodeableConcept `json:"itemCodeableConcept,omitempty"`
}

// AdverseEventPreventiveAction represents preventive actions that contributed to avoiding the adverse event
type AdverseEventPreventiveAction struct {
	common.BackboneElement

	// The action that contributed to avoiding the adverse event
	ItemReference       *common.Reference       `json:"itemReference,omitempty"`
	ItemCodeableConcept *common.CodeableConcept `json:"itemCodeableConcept,omitempty"`
}

// AdverseEventMitigatingAction represents the ameliorating action taken after the adverse event occurred in order to reduce the extent of harm
type AdverseEventMitigatingAction struct {
	common.BackboneElement

	// The ameliorating action taken after the adverse event occurred in order to reduce the extent of harm
	ItemReference       *common.Reference       `json:"itemReference,omitempty"`
	ItemCodeableConcept *common.CodeableConcept `json:"itemCodeableConcept,omitempty"`
}

// AdverseEventSupportingInfo represents supporting information relevant to the event
type AdverseEventSupportingInfo struct {
	common.BackboneElement

	// Relevant past history for the subject
	ItemReference       *common.Reference       `json:"itemReference,omitempty"`
	ItemCodeableConcept *common.CodeableConcept `json:"itemCodeableConcept,omitempty"`
}

// AdverseEvent represents an event that may be related to unintended effects on a patient or research participant
type AdverseEvent struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "AdverseEvent"

	// Whether the event actually happened or was potential
	Actuality AdverseEventActuality `json:"actuality"`

	// The overall type of event, intended for search and filtering purposes
	Category []common.CodeableConcept `json:"category,omitempty"`

	// Specific event that occurred or that was averted
	Code *common.CodeableConcept `json:"code,omitempty"`

	// The contributing factors suspected to have increased the probability or severity of the adverse event
	ContributingFactor []AdverseEventContributingFactor `json:"contributingFactor,omitempty"`

	// Estimated or actual date the AdverseEvent began, in the opinion of the reporter
	Detected *string `json:"detected,omitempty"`

	// The encounter the event occurred within
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Considered likely or probable or anticipated in the research study
	ExpectedInResearchStudy *bool `json:"expectedInResearchStudy,omitempty"`

	// Business identifier for the adverse event
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The information about where the adverse event occurred
	Location *common.Reference `json:"location,omitempty"`

	// The ameliorating action taken after the adverse event occurred in order to reduce the extent of harm
	MitigatingAction []AdverseEventMitigatingAction `json:"mitigatingAction,omitempty"`

	// Comments made about the adverse event by the performer, subject or other participants
	Note []Annotation `json:"note,omitempty"`

	// The date (and perhaps time) when the adverse event occurred
	OccurrenceDateTime *string        `json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod   *common.Period `json:"occurrencePeriod,omitempty"`
	OccurrenceTiming   *Timing        `json:"occurrenceTiming,omitempty"`

	// Describes the type of outcome from the adverse event
	Outcome []common.CodeableConcept `json:"outcome,omitempty"`

	// Indicates who or what participated in the adverse event and how they were involved
	Participant []AdverseEventParticipant `json:"participant,omitempty"`

	// Preventive actions that contributed to avoiding the adverse event
	PreventiveAction []AdverseEventPreventiveAction `json:"preventiveAction,omitempty"`

	// The recordedDate represents the date when this particular AdverseEvent record was created in the system
	RecordedDate *string `json:"recordedDate,omitempty"`

	// Information on who recorded the adverse event
	Recorder *common.Reference `json:"recorder,omitempty"`

	// Information about the condition that occurred as a result of the adverse event
	ResultingEffect []common.Reference `json:"resultingEffect,omitempty"`

	// Assessment of if the event was of clinical importance
	Seriousness *common.CodeableConcept `json:"seriousness,omitempty"`

	// in-progress | completed | entered-in-error | unknown
	Status AdverseEventStatus `json:"status"`

	// The research study that the subject is enrolled in
	Study []common.Reference `json:"study,omitempty"`

	// Subject impacted by event
	Subject common.Reference `json:"subject"`

	// Supporting information relevant to the event
	SupportingInfo []AdverseEventSupportingInfo `json:"supportingInfo,omitempty"`

	// Describes the entity that is suspected to have caused the adverse event
	SuspectEntity []AdverseEventSuspectEntity `json:"suspectEntity,omitempty"`
}

// AllergyIntoleranceCategory represents the category of an allergy/intolerance
type AllergyIntoleranceCategory string

const (
	AllergyIntoleranceCategoryFood        AllergyIntoleranceCategory = "food"
	AllergyIntoleranceCategoryMedication  AllergyIntoleranceCategory = "medication"
	AllergyIntoleranceCategoryEnvironment AllergyIntoleranceCategory = "environment"
	AllergyIntoleranceCategoryBiologic    AllergyIntoleranceCategory = "biologic"
)

// AllergyIntoleranceCriticality represents the criticality of an allergy/intolerance
type AllergyIntoleranceCriticality string

const (
	AllergyIntoleranceCriticalityLow            AllergyIntoleranceCriticality = "low"
	AllergyIntoleranceCriticalityHigh           AllergyIntoleranceCriticality = "high"
	AllergyIntoleranceCriticalityUnableToAssess AllergyIntoleranceCriticality = "unable-to-assess"
)

// AllergyIntoleranceType represents the type of an allergy/intolerance
type AllergyIntoleranceType string

const (
	AllergyIntoleranceTypeAllergy     AllergyIntoleranceType = "allergy"
	AllergyIntoleranceTypeIntolerance AllergyIntoleranceType = "intolerance"
)

// AllergyIntoleranceReactionSeverity represents the severity of an allergy/intolerance reaction
type AllergyIntoleranceReactionSeverity string

const (
	AllergyIntoleranceReactionSeverityMild     AllergyIntoleranceReactionSeverity = "mild"
	AllergyIntoleranceReactionSeverityModerate AllergyIntoleranceReactionSeverity = "moderate"
	AllergyIntoleranceReactionSeveritySevere   AllergyIntoleranceReactionSeverity = "severe"
)

// AllergyIntoleranceParticipant represents participants in the allergy/intolerance activities
type AllergyIntoleranceParticipant struct {
	common.BackboneElement

	// Indicates who or what participated in the activities related to the allergy or intolerance
	Actor common.Reference `json:"actor"`

	// Distinguishes the type of involvement of the actor in the activities related to the allergy or intolerance
	Function *common.CodeableConcept `json:"function,omitempty"`
}

// AllergyIntoleranceReaction represents adverse reaction event details
type AllergyIntoleranceReaction struct {
	common.BackboneElement

	// Use the description to provide any details of a particular event of the occurred reaction
	Description *string `json:"description,omitempty"`

	// Coding of the route of exposure with a terminology should be used wherever possible
	ExposureRoute *common.CodeableConcept `json:"exposureRoute,omitempty"`

	// Manifestation can be expressed as a single word, phrase or brief description
	Manifestation []CodeableReference `json:"manifestation"`

	// Use this field to record information indirectly related to a particular event
	Note []Annotation `json:"note,omitempty"`

	// Record of the date and/or time of the onset of the Reaction
	Onset *string `json:"onset,omitempty"`

	// It is acknowledged that this assessment is very subjective
	Severity *AllergyIntoleranceReactionSeverity `json:"severity,omitempty"`

	// Identification of the specific substance considered to be responsible for the Adverse Reaction event
	Substance *common.CodeableConcept `json:"substance,omitempty"`
}

// AllergyIntolerance represents a risk of harmful or undesirable physiological response
type AllergyIntolerance struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "AllergyIntolerance"

	// This data element has been included because it is currently being captured in some clinical systems
	Category []AllergyIntoleranceCategory `json:"category,omitempty"`

	// AllergyIntolerance.clinicalStatus should be present if verificationStatus is not entered-in-error
	ClinicalStatus *common.CodeableConcept `json:"clinicalStatus,omitempty"`

	// It is strongly recommended that this element be populated using a terminology
	Code *common.CodeableConcept `json:"code,omitempty"`

	// The default criticality value for any propensity to an adverse reaction should be 'Low Risk'
	Criticality *AllergyIntoleranceCriticality `json:"criticality,omitempty"`

	// The encounter when the allergy or intolerance was asserted
	Encounter *common.Reference `json:"encounter,omitempty"`

	// This is a business identifier, not a resource identifier
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// This date may be replicated by one of the Onset of Reaction dates
	LastOccurrence *string `json:"lastOccurrence,omitempty"`

	// For example: including reason for flagging a seriousness of 'High Risk'
	Note []Annotation `json:"note,omitempty"`

	// Age is generally used when the patient reports an age at which the AllergyIntolerance was noted
	OnsetDateTime *string        `json:"onsetDateTime,omitempty"`
	OnsetAge      *Age           `json:"onsetAge,omitempty"`
	OnsetPeriod   *common.Period `json:"onsetPeriod,omitempty"`
	OnsetRange    *Range         `json:"onsetRange,omitempty"`
	OnsetString   *string        `json:"onsetString,omitempty"`

	// Indicates who or what participated in the activities related to the allergy or intolerance and how they were involved
	Participant []AllergyIntoleranceParticipant `json:"participant,omitempty"`

	// The patient who has the allergy or intolerance
	Patient common.Reference `json:"patient"`

	// Details about each adverse reaction event linked to exposure to the identified substance
	Reaction []AllergyIntoleranceReaction `json:"reaction,omitempty"`

	// When onset date is unknown, recordedDate can be used to establish if the allergy or intolerance was present
	RecordedDate *string `json:"recordedDate,omitempty"`

	// Individual who recorded the record and takes responsibility for its content
	Recorder *common.Reference `json:"recorder,omitempty"`

	// Allergic (typically immune-mediated) reactions have been traditionally regarded as an indicator
	Type *AllergyIntoleranceType `json:"type,omitempty"`

	// The verification status to support the clinical status of the condition
	VerificationStatus *common.CodeableConcept `json:"verificationStatus,omitempty"`
}

// AppointmentStatus represents the status of an appointment
type AppointmentStatus string

const (
	AppointmentStatusProposed       AppointmentStatus = "proposed"
	AppointmentStatusPending        AppointmentStatus = "pending"
	AppointmentStatusBooked         AppointmentStatus = "booked"
	AppointmentStatusArrived        AppointmentStatus = "arrived"
	AppointmentStatusFulfilled      AppointmentStatus = "fulfilled"
	AppointmentStatusCancelled      AppointmentStatus = "cancelled"
	AppointmentStatusNoshow         AppointmentStatus = "noshow"
	AppointmentStatusEnteredInError AppointmentStatus = "entered-in-error"
	AppointmentStatusCheckedIn      AppointmentStatus = "checked-in"
	AppointmentStatusWaitlist       AppointmentStatus = "waitlist"
)

// AppointmentParticipantStatus represents the status of an appointment participant
type AppointmentParticipantStatus string

const (
	AppointmentParticipantStatusAccepted       AppointmentParticipantStatus = "accepted"
	AppointmentParticipantStatusDeclined       AppointmentParticipantStatus = "declined"
	AppointmentParticipantStatusTentative      AppointmentParticipantStatus = "tentative"
	AppointmentParticipantStatusNeedsAction    AppointmentParticipantStatus = "needs-action"
	AppointmentParticipantStatusEnteredInError AppointmentParticipantStatus = "entered-in-error"
)

// AppointmentRecurrenceTemplateWeeklyTemplate represents weekly template for appointment recurrence
type AppointmentRecurrenceTemplateWeeklyTemplate struct {
	common.BackboneElement

	// Days of the week when the appointment recurs
	Monday    *bool `json:"monday,omitempty"`
	Tuesday   *bool `json:"tuesday,omitempty"`
	Wednesday *bool `json:"wednesday,omitempty"`
	Thursday  *bool `json:"thursday,omitempty"`
	Friday    *bool `json:"friday,omitempty"`
	Saturday  *bool `json:"saturday,omitempty"`
	Sunday    *bool `json:"sunday,omitempty"`

	// The number of weeks between recurrences
	WeekInterval *int `json:"weekInterval,omitempty"`
}

// AppointmentRecurrenceTemplateMonthlyTemplate represents monthly template for appointment recurrence
type AppointmentRecurrenceTemplateMonthlyTemplate struct {
	common.BackboneElement

	// The day of the month when the appointment recurs
	DayOfMonth *int `json:"dayOfMonth,omitempty"`

	// Indicates which week within the month the appointment should recur
	NthWeekOfMonth *int `json:"nthWeekOfMonth,omitempty"`

	// Indicates which day of the week the appointment should recur
	DayOfWeek *common.Coding `json:"dayOfWeek,omitempty"`

	// The number of months between recurrences
	MonthInterval *int `json:"monthInterval,omitempty"`
}

// AppointmentRecurrenceTemplateYearlyTemplate represents yearly template for appointment recurrence
type AppointmentRecurrenceTemplateYearlyTemplate struct {
	common.BackboneElement

	// The number of years between recurrences
	YearInterval *int `json:"yearInterval,omitempty"`
}

// AppointmentRecurrenceTemplate represents the recurrence pattern for appointments
type AppointmentRecurrenceTemplate struct {
	common.BackboneElement

	// The timezone of the recurring appointment occurrences
	Timezone *common.CodeableConcept `json:"timezone,omitempty"`

	// How often the appointment series should recur
	RecurrenceType common.CodeableConcept `json:"recurrenceType"`

	// The date when the recurrence should end
	LastOccurrenceDate *string `json:"lastOccurrenceDate,omitempty"`

	// How many appointment occurrences are planned for the series
	OccurrenceCount *int `json:"occurrenceCount,omitempty"`

	// Recurring dates that should be excluded from the series
	ExcludingDate []string `json:"excludingDate,omitempty"`

	// Identifiers of occurrences that should be excluded from the series
	ExcludingRecurrenceId []int `json:"excludingRecurrenceId,omitempty"`

	// Template for weekly recurring appointments
	WeeklyTemplate *AppointmentRecurrenceTemplateWeeklyTemplate `json:"weeklyTemplate,omitempty"`

	// Template for monthly recurring appointments
	MonthlyTemplate *AppointmentRecurrenceTemplateMonthlyTemplate `json:"monthlyTemplate,omitempty"`

	// Template for yearly recurring appointments
	YearlyTemplate *AppointmentRecurrenceTemplateYearlyTemplate `json:"yearlyTemplate,omitempty"`
}

// AppointmentParticipant represents participants involved in appointment
type AppointmentParticipant struct {
	common.BackboneElement

	// Person, Location, HealthcareService, PractitionerRole, RelatedPerson or Device
	Actor *common.Reference `json:"actor,omitempty"`

	// Whether this participant is required to be present at the meeting
	Required *bool `json:"required,omitempty"`

	// Participation status of the participant
	Status AppointmentParticipantStatus `json:"status"`

	// Role of participant in the appointment
	Type []common.CodeableConcept `json:"type,omitempty"`

	// Participation period of the participant
	Period *common.Period `json:"period,omitempty"`
}

// Appointment represents a booking of a healthcare event among patient(s), practitioner(s), related person(s) and/or device(s)
type Appointment struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Appointment"

	// The style of appointment or patient that has been booked
	AppointmentType *common.CodeableConcept `json:"appointmentType,omitempty"`

	// The request this appointment is allocated to assess
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// The coded reason for the appointment being cancelled
	CancellationReason *common.CodeableConcept `json:"cancellationReason,omitempty"`

	// Concepts representing classification of patient encounter
	Class []common.CodeableConcept `json:"class,omitempty"`

	// The date that this appointment was initially created
	Created *string `json:"created,omitempty"`

	// Shown on a subject line in a meeting request, or appointment list
	Description *string `json:"description,omitempty"`

	// When appointment is to conclude
	End *string `json:"end,omitempty"`

	// External identifier for the appointment
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Can be less than start/end
	MinutesDuration *int `json:"minutesDuration,omitempty"`

	// Additional information to support the appointment
	Note []Annotation `json:"note,omitempty"`

	// The originating appointment in a recurring set of appointments
	OriginatingAppointment *common.Reference `json:"originatingAppointment,omitempty"`

	// List of participants involved in the appointment
	Participant []AppointmentParticipant `json:"participant"`

	// Detailed information and instructions for the patient
	PatientInstruction []CodeableReference `json:"patientInstruction,omitempty"`

	// Used to make informed decisions if needing to re-prioritize
	Priority *common.CodeableConcept `json:"priority,omitempty"`

	// Coded reason this appointment is scheduled
	Reason []CodeableReference `json:"reason,omitempty"`

	// Details of the recurrence pattern/template for recurring appointments
	RecurrenceTemplate []AppointmentRecurrenceTemplate `json:"recurrenceTemplate,omitempty"`

	// Appointment replaced by this Appointment
	Replaces []common.Reference `json:"replaces,omitempty"`

	// Potential date/time interval(s) requested to allocate the appointment within
	RequestedPeriod []common.Period `json:"requestedPeriod,omitempty"`

	// A broad categorization of the service that is to be performed during this appointment
	ServiceCategory []common.CodeableConcept `json:"serviceCategory,omitempty"`

	// The specific service that is to be performed during this appointment
	ServiceType []CodeableReference `json:"serviceType,omitempty"`

	// The slots that this appointment is filling
	Slot []common.Reference `json:"slot,omitempty"`

	// The specialty of a practitioner that would be required to perform the service requested
	Specialty []common.CodeableConcept `json:"specialty,omitempty"`

	// When appointment is to take place
	Start *string `json:"start,omitempty"`

	// The overall status of the Appointment
	Status AppointmentStatus `json:"status"`

	// The patient or group associated with the appointment
	Subject *common.Reference `json:"subject,omitempty"`

	// Additional information to support the appointment
	SupportingInformation []common.Reference `json:"supportingInformation,omitempty"`

	// Connection details of a virtual service
	VirtualService []VirtualServiceDetail `json:"virtualService,omitempty"`
}

// AppointmentResponse represents a reply to an appointment request for a patient and/or practitioner(s)
type AppointmentResponse struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "AppointmentResponse"

	// A Person, Location, HealthcareService, or Device that is participating in the appointment
	Actor *common.Reference `json:"actor,omitempty"`

	// Appointment that this response is replying to
	Appointment common.Reference `json:"appointment"`

	// This comment is particularly important when the responder is declining, tentatively accepting or requesting another time
	Comment *string `json:"comment,omitempty"`

	// This may be either the same as the appointment request to confirm the details of the appointment, or alternately a new time to request a re-negotiation of the end time
	End *string `json:"end,omitempty"`

	// Business identifier for this appointment response
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The original date within a recurring request
	OccurrenceDate *string `json:"occurrenceDate,omitempty"`

	// This element is labeled as a modifier because the status contains the code entered-in-error
	ParticipantStatus AppointmentParticipantStatus `json:"participantStatus"`

	// The role of the participant can be used to declare what the actor will be doing in the scope of the referenced appointment
	ParticipantType []common.CodeableConcept `json:"participantType,omitempty"`

	// Indicates that the response is proposing a different time that was initially requested
	ProposedNewTime *bool `json:"proposedNewTime,omitempty"`

	// The recurrence ID (sequence number) of the specific appointment when responding to a recurring request
	RecurrenceId *int `json:"recurrenceId,omitempty"`

	// When a recurring appointment is requested, the participant may choose to respond to each individual occurrence
	Recurring *bool `json:"recurring,omitempty"`

	// This may be either the same as the appointment request to confirm the details of the appointment, or alternately a new time to request a re-negotiation of the start time
	Start *string `json:"start,omitempty"`
}

// Basic represents a resource that is used to represent concepts not yet defined in FHIR
type Basic struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Basic"

	// Indicates who was responsible for creating the resource instance
	Author *common.Reference `json:"author,omitempty"`

	// The type of resource being represented - this defines the meaning of the resource and cannot be ignored
	Code common.CodeableConcept `json:"code"`

	// Identifies when the resource was first created
	Created *string `json:"created,omitempty"`

	// Identifier assigned to the resource for business purposes, outside the context of FHIR
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Optional as not all potential resources will have subjects
	Subject *common.Reference `json:"subject,omitempty"`
}

// BundleType represents the type of bundle
type BundleType string

const (
	BundleTypeDocument                 BundleType = "document"
	BundleTypeMessage                  BundleType = "message"
	BundleTypeTransaction              BundleType = "transaction"
	BundleTypeTransactionResponse      BundleType = "transaction-response"
	BundleTypeBatch                    BundleType = "batch"
	BundleTypeBatchResponse            BundleType = "batch-response"
	BundleTypeHistory                  BundleType = "history"
	BundleTypeSearchset                BundleType = "searchset"
	BundleTypeCollection               BundleType = "collection"
	BundleTypeSubscriptionNotification BundleType = "subscription-notification"
)

// BundleEntrySearchMode represents the mode of a search entry
type BundleEntrySearchMode string

const (
	BundleEntrySearchModeMatch   BundleEntrySearchMode = "match"
	BundleEntrySearchModeInclude BundleEntrySearchMode = "include"
	BundleEntrySearchModeOutcome BundleEntrySearchMode = "outcome"
)

// BundleEntryRequestMethod represents the HTTP method for a bundle entry request
type BundleEntryRequestMethod string

const (
	BundleEntryRequestMethodGET    BundleEntryRequestMethod = "GET"
	BundleEntryRequestMethodHEAD   BundleEntryRequestMethod = "HEAD"
	BundleEntryRequestMethodPOST   BundleEntryRequestMethod = "POST"
	BundleEntryRequestMethodPUT    BundleEntryRequestMethod = "PUT"
	BundleEntryRequestMethodDELETE BundleEntryRequestMethod = "DELETE"
	BundleEntryRequestMethodPATCH  BundleEntryRequestMethod = "PATCH"
)

// Signature represents a digital signature
type Signature struct {
	DataType

	// The technical format of the signature
	Type []common.Coding `json:"type"`

	// When the signature was created
	When *string `json:"when,omitempty"`

	// Who signed
	Who common.Reference `json:"who"`

	// The party represented
	OnBehalfOf *common.Reference `json:"onBehalfOf,omitempty"`

	// The technical format of the signed resources
	TargetFormat *string `json:"targetFormat,omitempty"`

	// The technical format of the signature
	SigFormat *string `json:"sigFormat,omitempty"`

	// The actual signature content (XML DigSig, JWS, etc.)
	Data *string `json:"data,omitempty"`
}

// BundleLink represents a link in a bundle
type BundleLink struct {
	common.BackboneElement

	// A name which details the functional use for this link
	Relation string `json:"relation"`

	// The reference details for the link
	URL string `json:"url"`
}

// BundleEntrySearch represents information about the search process that led to the creation of this entry
type BundleEntrySearch struct {
	common.BackboneElement

	// There is only one mode. In some corner cases, a resource may be included because it is both a match and an include
	Mode *BundleEntrySearchMode `json:"mode,omitempty"`

	// Servers are not required to return a ranking score. 1 is most relevant, and 0 is least relevant
	Score *float64 `json:"score,omitempty"`
}

// BundleEntryRequest represents additional information about how this entry should be processed as part of a transaction or batch
type BundleEntryRequest struct {
	common.BackboneElement

	// Only perform the operation if the Etag value matches
	IfMatch *string `json:"ifMatch,omitempty"`

	// Only perform the operation if the last updated date matches
	IfModifiedSince *string `json:"ifModifiedSince,omitempty"`

	// Instruct the server not to perform the create if a specified resource already exists
	IfNoneExist *string `json:"ifNoneExist,omitempty"`

	// If the ETag values match, return a 304 Not Modified status
	IfNoneMatch *string `json:"ifNoneMatch,omitempty"`

	// In a transaction or batch, this is the HTTP action to be executed for this entry
	Method BundleEntryRequestMethod `json:"method"`

	// The URL for the entry
	URL string `json:"url"`
}

// BundleEntryResponse represents the results of processing the corresponding 'request' entry in the batch or transaction
type BundleEntryResponse struct {
	common.BackboneElement

	// Etags match the Resource.meta.versionId
	ETag *string `json:"etag,omitempty"`

	// This has to match the same time in the meta header (meta.lastUpdated) if a resource is included
	LastModified *string `json:"lastModified,omitempty"`

	// The location header created by processing this operation
	Location *string `json:"location,omitempty"`

	// For a POST/PUT operation, this is the equivalent outcome that would be returned for prefer = operationoutcome
	Outcome interface{} `json:"outcome,omitempty"`

	// The status code returned by processing this entry
	Status string `json:"status"`
}

// BundleEntry represents an entry in a bundle resource
type BundleEntry struct {
	common.BackboneElement

	// The Absolute URL for the resource
	FullURL *string `json:"fullUrl,omitempty"`

	// A series of links that provide context to this entry
	Link []BundleLink `json:"link,omitempty"`

	// Additional information about how this entry should be processed as part of a transaction or batch
	Request *BundleEntryRequest `json:"request,omitempty"`

	// The Resource for the entry
	Resource interface{} `json:"resource,omitempty"`

	// Indicates the results of processing the corresponding 'request' entry in the batch or transaction
	Response *BundleEntryResponse `json:"response,omitempty"`

	// Information about the search process that lead to the creation of this entry
	Search *BundleEntrySearch `json:"search,omitempty"`
}

// Bundle represents a container for a collection of resources
type Bundle struct {
	Resource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Bundle"

	// An entry in a bundle resource - will either contain a resource or information about a resource
	Entry []BundleEntry `json:"entry,omitempty"`

	// Persistent identity generally only matters for batches of type Document, Message, and Collection
	Identifier *common.Identifier `json:"identifier,omitempty"`

	// Issues must apply to the Bundle as a whole, not to individual entries
	Issues interface{} `json:"issues,omitempty"`

	// A series of links that provide context to this bundle
	Link []BundleLink `json:"link,omitempty"`

	// The signature could be created by the "author" of the bundle or by the originating device
	Signature *Signature `json:"signature,omitempty"`

	// For many bundles, the timestamp is equal to .meta.lastUpdated
	Timestamp *string `json:"timestamp,omitempty"`

	// Only used if the bundle is a search result set
	Total *int `json:"total,omitempty"`

	// It's possible to use a bundle for other purposes (e.g. a document can be accepted as a transaction)
	Type BundleType `json:"type"`
}

// CapabilityStatementKind represents the way that this statement is intended to be used
type CapabilityStatementKind string

const (
	CapabilityStatementKindInstance     CapabilityStatementKind = "instance"
	CapabilityStatementKindCapability   CapabilityStatementKind = "capability"
	CapabilityStatementKindRequirements CapabilityStatementKind = "requirements"
)

// CapabilityStatementStatus represents the status of a capability statement
type CapabilityStatementStatus string

const (
	CapabilityStatementStatusDraft   CapabilityStatementStatus = "draft"
	CapabilityStatementStatusActive  CapabilityStatementStatus = "active"
	CapabilityStatementStatusRetired CapabilityStatementStatus = "retired"
	CapabilityStatementStatusUnknown CapabilityStatementStatus = "unknown"
)

// CapabilityStatementRestMode represents the mode of a REST interface
type CapabilityStatementRestMode string

const (
	CapabilityStatementRestModeClient CapabilityStatementRestMode = "client"
	CapabilityStatementRestModeServer CapabilityStatementRestMode = "server"
)

// CapabilityStatementRestResourceInteractionCode represents the code of a resource interaction
type CapabilityStatementRestResourceInteractionCode string

const (
	CapabilityStatementRestResourceInteractionCodeRead            CapabilityStatementRestResourceInteractionCode = "read"
	CapabilityStatementRestResourceInteractionCodeVread           CapabilityStatementRestResourceInteractionCode = "vread"
	CapabilityStatementRestResourceInteractionCodeUpdate          CapabilityStatementRestResourceInteractionCode = "update"
	CapabilityStatementRestResourceInteractionCodePatch           CapabilityStatementRestResourceInteractionCode = "patch"
	CapabilityStatementRestResourceInteractionCodeDelete          CapabilityStatementRestResourceInteractionCode = "delete"
	CapabilityStatementRestResourceInteractionCodeHistoryInstance CapabilityStatementRestResourceInteractionCode = "history-instance"
	CapabilityStatementRestResourceInteractionCodeHistoryType     CapabilityStatementRestResourceInteractionCode = "history-type"
	CapabilityStatementRestResourceInteractionCodeCreate          CapabilityStatementRestResourceInteractionCode = "create"
	CapabilityStatementRestResourceInteractionCodeSearchType      CapabilityStatementRestResourceInteractionCode = "search-type"
)

// CapabilityStatementRestInteractionCode represents the code of a system interaction
type CapabilityStatementRestInteractionCode string

const (
	CapabilityStatementRestInteractionCodeTransaction   CapabilityStatementRestInteractionCode = "transaction"
	CapabilityStatementRestInteractionCodeBatch         CapabilityStatementRestInteractionCode = "batch"
	CapabilityStatementRestInteractionCodeSearchSystem  CapabilityStatementRestInteractionCode = "search-system"
	CapabilityStatementRestInteractionCodeHistorySystem CapabilityStatementRestInteractionCode = "history-system"
)

// CapabilityStatementRestResourceConditionalDelete represents conditional delete support
type CapabilityStatementRestResourceConditionalDelete string

const (
	CapabilityStatementRestResourceConditionalDeleteNotSupported CapabilityStatementRestResourceConditionalDelete = "not-supported"
	CapabilityStatementRestResourceConditionalDeleteSingle       CapabilityStatementRestResourceConditionalDelete = "single"
	CapabilityStatementRestResourceConditionalDeleteMultiple     CapabilityStatementRestResourceConditionalDelete = "multiple"
)

// CapabilityStatementRestResourceConditionalRead represents conditional read support
type CapabilityStatementRestResourceConditionalRead string

const (
	CapabilityStatementRestResourceConditionalReadNotSupported  CapabilityStatementRestResourceConditionalRead = "not-supported"
	CapabilityStatementRestResourceConditionalReadModifiedSince CapabilityStatementRestResourceConditionalRead = "modified-since"
	CapabilityStatementRestResourceConditionalReadNotMatch      CapabilityStatementRestResourceConditionalRead = "not-match"
	CapabilityStatementRestResourceConditionalReadFullSupport   CapabilityStatementRestResourceConditionalRead = "full-support"
)

// CapabilityStatementRestResourceReferencePolicy represents reference policy
type CapabilityStatementRestResourceReferencePolicy string

const (
	CapabilityStatementRestResourceReferencePolicyLiteral  CapabilityStatementRestResourceReferencePolicy = "literal"
	CapabilityStatementRestResourceReferencePolicyLogical  CapabilityStatementRestResourceReferencePolicy = "logical"
	CapabilityStatementRestResourceReferencePolicyResolves CapabilityStatementRestResourceReferencePolicy = "resolves"
	CapabilityStatementRestResourceReferencePolicyEnforced CapabilityStatementRestResourceReferencePolicy = "enforced"
	CapabilityStatementRestResourceReferencePolicyLocal    CapabilityStatementRestResourceReferencePolicy = "local"
)

// CapabilityStatementRestResourceVersioning represents versioning support
type CapabilityStatementRestResourceVersioning string

const (
	CapabilityStatementRestResourceVersioningNoVersion       CapabilityStatementRestResourceVersioning = "no-version"
	CapabilityStatementRestResourceVersioningVersioned       CapabilityStatementRestResourceVersioning = "versioned"
	CapabilityStatementRestResourceVersioningVersionedUpdate CapabilityStatementRestResourceVersioning = "versioned-update"
)

// CapabilityStatementRestResourceSearchParamType represents search parameter type
type CapabilityStatementRestResourceSearchParamType string

const (
	CapabilityStatementRestResourceSearchParamTypeNumber    CapabilityStatementRestResourceSearchParamType = "number"
	CapabilityStatementRestResourceSearchParamTypeDate      CapabilityStatementRestResourceSearchParamType = "date"
	CapabilityStatementRestResourceSearchParamTypeString    CapabilityStatementRestResourceSearchParamType = "string"
	CapabilityStatementRestResourceSearchParamTypeToken     CapabilityStatementRestResourceSearchParamType = "token"
	CapabilityStatementRestResourceSearchParamTypeReference CapabilityStatementRestResourceSearchParamType = "reference"
	CapabilityStatementRestResourceSearchParamTypeComposite CapabilityStatementRestResourceSearchParamType = "composite"
	CapabilityStatementRestResourceSearchParamTypeQuantity  CapabilityStatementRestResourceSearchParamType = "quantity"
	CapabilityStatementRestResourceSearchParamTypeUri       CapabilityStatementRestResourceSearchParamType = "uri"
	CapabilityStatementRestResourceSearchParamTypeSpecial   CapabilityStatementRestResourceSearchParamType = "special"
)

// CapabilityStatementMessagingMode represents messaging mode
type CapabilityStatementMessagingMode string

const (
	CapabilityStatementMessagingModeSender   CapabilityStatementMessagingMode = "sender"
	CapabilityStatementMessagingModeReceiver CapabilityStatementMessagingMode = "receiver"
)

// CapabilityStatementDocumentMode represents document mode
type CapabilityStatementDocumentMode string

const (
	CapabilityStatementDocumentModeProducer CapabilityStatementDocumentMode = "producer"
	CapabilityStatementDocumentModeConsumer CapabilityStatementDocumentMode = "consumer"
)

// CapabilityStatementSoftware represents software covered by this capability statement
type CapabilityStatementSoftware struct {
	common.BackboneElement

	// Name the software is known by
	Name string `json:"name"`

	// Date this version of the software was released
	ReleaseDate *string `json:"releaseDate,omitempty"`

	// If possible, a version should be specified, as statements are likely to be different for different versions of software
	Version *string `json:"version,omitempty"`
}

// CapabilityStatementImplementation represents a specific implementation instance
type CapabilityStatementImplementation struct {
	common.BackboneElement

	// The organization responsible for the management of the instance and oversight of the data on the server
	Custodian *common.Reference `json:"custodian,omitempty"`

	// Information about the specific installation that this capability statement relates to
	Description string `json:"description"`

	// An absolute base URL for the implementation
	URL *string `json:"url,omitempty"`
}

// CapabilityStatementRestSecurity represents security information
type CapabilityStatementRestSecurity struct {
	common.BackboneElement

	// The easiest CORS headers to add are Access-Control-Allow-Origin: * & Access-Control-Request-Method: GET, POST, PUT, DELETE
	CORS *bool `json:"cors,omitempty"`

	// General description of how security works
	Description *string `json:"description,omitempty"`

	// Types of security services that are supported/required by the system
	Service []common.CodeableConcept `json:"service,omitempty"`
}

// CapabilityStatementRestResourceInteraction represents a restful operation supported by the solution
type CapabilityStatementRestResourceInteraction struct {
	common.BackboneElement

	// Coded identifier of the operation, supported by the system resource
	Code CapabilityStatementRestResourceInteractionCode `json:"code"`

	// Guidance specific to the implementation of this operation
	Documentation *string `json:"documentation,omitempty"`
}

// CapabilityStatementRestResourceSearchParam represents a search parameter supported by the system
type CapabilityStatementRestResourceSearchParam struct {
	common.BackboneElement

	// This SHOULD be present, and matches refers to a SearchParameter by its canonical URL
	Definition *string `json:"definition,omitempty"`

	// This allows documentation of any distinct behaviors about how the search parameter is used
	Documentation *string `json:"documentation,omitempty"`

	// Parameter names cannot overlap with standard parameter names
	Name string `json:"name"`

	// While this can be looked up from the definition, it is included here as a convenience
	Type CapabilityStatementRestResourceSearchParamType `json:"type"`
}

// CapabilityStatementRestResourceOperation represents an operation supported by the system
type CapabilityStatementRestResourceOperation struct {
	common.BackboneElement

	// This can be used to build an HTML form to invoke the operation, for instance
	Definition string `json:"definition"`

	// Documentation that describes anything special about the operation behavior
	Documentation *string `json:"documentation,omitempty"`

	// The name here SHOULD be the same as the OperationDefinition.code in the referenced OperationDefinition
	Name string `json:"name"`
}

// CapabilityStatementRestResource represents a resource type supported by the system
type CapabilityStatementRestResource struct {
	common.BackboneElement

	// Conditional Create is mainly appropriate for interface engine scripts converting from other formats
	ConditionalCreate *bool `json:"conditionalCreate,omitempty"`

	// Conditional Delete is mainly appropriate for interface engine scripts converting from other formats
	ConditionalDelete *CapabilityStatementRestResourceConditionalDelete `json:"conditionalDelete,omitempty"`

	// Conditional Patch is mainly appropriate for interface engine scripts converting from other formats
	ConditionalPatch *bool `json:"conditionalPatch,omitempty"`

	// Conditional Read is mainly appropriate for interface engine scripts converting from other formats
	ConditionalRead *CapabilityStatementRestResourceConditionalRead `json:"conditionalRead,omitempty"`

	// Conditional Update is mainly appropriate for interface engine scripts converting from other formats
	ConditionalUpdate *bool `json:"conditionalUpdate,omitempty"`

	// Additional information about the resource type used by the system
	Documentation *string `json:"documentation,omitempty"`

	// In general, a Resource will only appear in a CapabilityStatement if the server actually has some capabilities
	Interaction []CapabilityStatementRestResourceInteraction `json:"interaction,omitempty"`

	// Operations linked from CapabilityStatement.rest.resource.operation
	Operation []CapabilityStatementRestResourceOperation `json:"operation,omitempty"`

	// All other profiles for this type that are listed in `.rest.resource.supportedProfile` must conform to this profile
	Profile *string `json:"profile,omitempty"`

	// It is useful to support the vRead operation for current operations, even if past versions aren't available
	ReadHistory *bool `json:"readHistory,omitempty"`

	// A set of flags that defines how references are supported
	ReferencePolicy []CapabilityStatementRestResourceReferencePolicy `json:"referencePolicy,omitempty"`

	// Documenting `_include` support helps set conformance expectations for the desired system
	SearchInclude []string `json:"searchInclude,omitempty"`

	// The search parameters should include the control search parameters such as _sort, _count, etc.
	SearchParam []CapabilityStatementRestResourceSearchParam `json:"searchParam,omitempty"`

	// See `CapabilityStatement.rest.resource.searchInclude` comments
	SearchRevInclude []string `json:"searchRevInclude,omitempty"`

	// Supported profiles must conform to the resource profile in the `.rest.resource.profile` element
	SupportedProfile []string `json:"supportedProfile,omitempty"`

	// A type of resource exposed via the restful interface
	Type string `json:"type"`

	// Allowing the clients to create new identities on the server
	UpdateCreate *bool `json:"updateCreate,omitempty"`

	// If a server supports versionIds correctly, it SHOULD support vread too
	Versioning *CapabilityStatementRestResourceVersioning `json:"versioning,omitempty"`
}

// CapabilityStatementRestInteraction represents a specification of restful operations supported by the system
type CapabilityStatementRestInteraction struct {
	common.BackboneElement

	// A coded identifier of the operation, supported by the system
	Code CapabilityStatementRestInteractionCode `json:"code"`

	// Guidance specific to the implementation of this operation
	Documentation *string `json:"documentation,omitempty"`
}

// CapabilityStatementRest represents the RESTful interface for this system
type CapabilityStatementRest struct {
	common.BackboneElement

	// At present, the only defined compartments are at CompartmentDefinition
	Compartment []string `json:"compartment,omitempty"`

	// Information about the system's restful capabilities that apply across all applications
	Documentation *string `json:"documentation,omitempty"`

	// A specification of restful operations supported by the system
	Interaction []CapabilityStatementRestInteraction `json:"interaction,omitempty"`

	// Identifies whether this portion of the statement is describing the ability to initiate or receive restful operations
	Mode CapabilityStatementRestMode `json:"mode"`

	// CapabilityStatement.rest.operation is for operations invoked at the system level
	Operation []CapabilityStatementRestResourceOperation `json:"operation,omitempty"`

	// Max of one repetition per resource type
	Resource []CapabilityStatementRestResource `json:"resource,omitempty"`

	// Typically, the only search parameters supported for all searches are those that apply to all resources
	SearchParam []CapabilityStatementRestResourceSearchParam `json:"searchParam,omitempty"`

	// Information about security implementation from an interface perspective
	Security *CapabilityStatementRestSecurity `json:"security,omitempty"`
}

// CapabilityStatementMessagingEndpoint represents an endpoint to which messages and/or replies are to be sent
type CapabilityStatementMessagingEndpoint struct {
	common.BackboneElement

	// The network address of the endpoint
	Address string `json:"address"`

	// A list of the messaging transport protocol(s) identifiers, supported by this endpoint
	Protocol common.Coding `json:"protocol"`
}

// CapabilityStatementMessagingSupportedMessage represents a supported message
type CapabilityStatementMessagingSupportedMessage struct {
	common.BackboneElement

	// Points to a message definition that identifies the messaging event, message structure, allowed responses, etc.
	Definition string `json:"definition"`

	// The mode of this event declaration - whether application is sender or receiver
	Mode CapabilityStatementMessagingMode `json:"mode"`
}

// CapabilityStatementMessaging represents the messaging capabilities of the system
type CapabilityStatementMessaging struct {
	common.BackboneElement

	// Documentation about the system's messaging capabilities for this endpoint
	Documentation *string `json:"documentation,omitempty"`

	// An endpoint (network accessible address) to which messages and/or replies are to be sent
	Endpoint []CapabilityStatementMessagingEndpoint `json:"endpoint,omitempty"`

	// If this value is missing then the application does not implement (receiver) or depend on (sender) reliable messaging
	ReliableCache *int `json:"reliableCache,omitempty"`

	// This is a proposed alternative to the messaging.event structure
	SupportedMessage []CapabilityStatementMessagingSupportedMessage `json:"supportedMessage,omitempty"`
}

// CapabilityStatementDocument represents a document definition
type CapabilityStatementDocument struct {
	common.BackboneElement

	// A description of how the application supports or uses the specified document profile
	Documentation *string `json:"documentation,omitempty"`

	// Mode of this document declaration - whether an application is a producer or consumer
	Mode CapabilityStatementDocumentMode `json:"mode"`

	// The profile is actually on the Bundle
	Profile string `json:"profile"`
}

// CapabilityStatement represents a capability statement documenting the capabilities of a FHIR Server or Client
type CapabilityStatement struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "CapabilityStatement"

	// In general, if a server gets a request with an Accept-Language that it doesn't support, it should still return the resource
	AcceptLanguage []string `json:"acceptLanguage,omitempty"`

	// May be a web site, an email address, a telephone number, etc.
	Contact []ContactDetail `json:"contact,omitempty"`

	// Use and/or publishing restrictions
	Copyright *string `json:"copyright,omitempty"`

	// Copyright holder and year(s)
	CopyrightLabel *string `json:"copyrightLabel,omitempty"`

	// The date is often not tracked until the resource is published, but may be present on draft content
	Date string `json:"date"`

	// This description can be used to capture details such as comments about misuse, instructions for clinical use and interpretation
	Description *string `json:"description,omitempty"`

	// A document definition
	Document []CapabilityStatementDocument `json:"document,omitempty"`

	// Allows filtering of capability statements that are appropriate for use versus not
	Experimental *bool `json:"experimental,omitempty"`

	// Servers may implement multiple versions
	FhirVersion string `json:"fhirVersion"`

	// "xml", "json" and "ttl" are allowed, which describe the simple encodings described in the specification
	Format []string `json:"format"`

	// A formal identifier that is used to identify this CapabilityStatement when it is represented in other formats
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Identifies a specific implementation instance that is described by the capability statement
	Implementation *CapabilityStatementImplementation `json:"implementation,omitempty"`

	// Note: this is primarily only relevant in terms of ImplementationGuides that don't define specific CapabilityStatements
	ImplementationGuide []string `json:"implementationGuide,omitempty"`

	// the contents of any directly or indirectly imported CapabilityStatements SHALL NOT overlap
	Imports []string `json:"imports,omitempty"`

	// HL7 defines the following Services: Terminology Service
	Instantiates []string `json:"instantiates,omitempty"`

	// It may be possible for the capability statement to be used in jurisdictions other than those for which it was originally designed
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`

	// The way that this statement is intended to be used, to describe an actual running instance of software
	Kind CapabilityStatementKind `json:"kind"`

	// Multiple repetitions allow the documentation of multiple endpoints per solution
	Messaging []CapabilityStatementMessaging `json:"messaging,omitempty"`

	// The name is not expected to be globally unique
	Name *string `json:"name,omitempty"`

	// At present, the patch mime types application/json-patch+json and application/xml-patch+xml are legal
	PatchFormat []string `json:"patchFormat,omitempty"`

	// Usually an organization but may be an individual
	Publisher *string `json:"publisher,omitempty"`

	// This element does not describe the usage of the capability statement
	Purpose *string `json:"purpose,omitempty"`

	// Multiple repetitions allow definition of both client and/or server behaviors
	Rest []CapabilityStatementRest `json:"rest,omitempty"`

	// Software that is covered by this capability statement
	Software *CapabilityStatementSoftware `json:"software,omitempty"`

	// Allows filtering of capability statements that are appropriate for use versus not
	Status CapabilityStatementStatus `json:"status"`

	// This name does not need to be machine-processing friendly and may contain punctuation, white-space, etc.
	Title *string `json:"title,omitempty"`

	// Can be a urn:uuid: or a urn:oid: but real http: addresses are preferred
	URL *string `json:"url,omitempty"`

	// When multiple useContexts are specified, there is no expectation that all or any of the contexts apply
	UseContext []UsageContext `json:"useContext,omitempty"`

	// There may be different capability statement instances that have the same identifier but different versions
	Version *string `json:"version,omitempty"`

	// If set as a string, this is a FHIRPath expression that has two additional context variables passed in
	VersionAlgorithmString *string        `json:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *common.Coding `json:"versionAlgorithmCoding,omitempty"`
}

// CarePlanStatus represents the status of a care plan
type CarePlanStatus string

const (
	CarePlanStatusDraft          CarePlanStatus = "draft"
	CarePlanStatusActive         CarePlanStatus = "active"
	CarePlanStatusOnHold         CarePlanStatus = "on-hold"
	CarePlanStatusRevoked        CarePlanStatus = "revoked"
	CarePlanStatusCompleted      CarePlanStatus = "completed"
	CarePlanStatusEnteredInError CarePlanStatus = "entered-in-error"
	CarePlanStatusUnknown        CarePlanStatus = "unknown"
)

// CarePlanIntent represents the intention/purpose of a care plan
type CarePlanIntent string

const (
	CarePlanIntentProposal  CarePlanIntent = "proposal"
	CarePlanIntentPlan      CarePlanIntent = "plan"
	CarePlanIntentOrder     CarePlanIntent = "order"
	CarePlanIntentOption    CarePlanIntent = "option"
	CarePlanIntentDirective CarePlanIntent = "directive"
)

// CarePlanActivity represents an action that has occurred or is a planned action to occur as part of the plan
type CarePlanActivity struct {
	common.BackboneElement

	// Note that this should not duplicate the activity status (e.g. completed or in progress)
	PerformedActivity []CodeableReference `json:"performedActivity,omitempty"`

	// The goal should be visible when the resource referenced by CarePlan.activity.plannedActivityReference is viewed independently
	PlannedActivityReference *common.Reference `json:"plannedActivityReference,omitempty"`

	// This element should NOT be used to describe the activity to be performed
	Progress []Annotation `json:"progress,omitempty"`
}

// CarePlan represents the intention of how one or more practitioners intend to deliver care for a particular patient, group or community
type CarePlan struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "CarePlan"

	// Identifies an action that has occurred or is a planned action to occur as part of the plan
	Activity []CarePlanActivity `json:"activity,omitempty"`

	// Use CarePlan.addresses.concept when a code sufficiently describes the concern
	Addresses []CodeableReference `json:"addresses,omitempty"`

	// A higher-level request resource (i.e. a plan, proposal or order) that is fulfilled in whole or in part by this care plan
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// Identifies all people and organizations who are expected to be involved in the care envisioned by this plan
	CareTeam []common.Reference `json:"careTeam,omitempty"`

	// There may be multiple axes of categorization and one plan may serve multiple purposes
	Category []common.CodeableConcept `json:"category,omitempty"`

	// Collaborative care plans may have multiple contributors
	Contributor []common.Reference `json:"contributor,omitempty"`

	// Represents when this particular CarePlan record was created in the system
	Created *string `json:"created,omitempty"`

	// The custodian might or might not be a contributor
	Custodian *common.Reference `json:"custodian,omitempty"`

	// A description of the scope and nature of the plan
	Description *string `json:"description,omitempty"`

	// This will typically be the encounter the event occurred within
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Goal can be achieving a particular change or merely maintaining a current state or even slowing a decline
	Goal []common.Reference `json:"goal,omitempty"`

	// Business identifier for this care plan
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The URL pointing to a FHIR-defined protocol, guideline, questionnaire or other definition that is adhered to in whole or in part
	InstantiatesCanonical []string `json:"instantiatesCanonical,omitempty"`

	// This might be an HTML page, PDF, etc. or could just be a non-resolvable URI identifier
	InstantiatesUri []string `json:"instantiatesUri,omitempty"`

	// This element is labeled as a modifier because the intent alters when and how the resource is actually applicable
	Intent CarePlanIntent `json:"intent"`

	// General notes about the care plan not covered elsewhere
	Note []Annotation `json:"note,omitempty"`

	// Each care plan is an independent request, such that having a care plan be part of another care plan can cause issues
	PartOf []common.Reference `json:"partOf,omitempty"`

	// Any activities scheduled as part of the plan should be constrained to the specified period
	Period *common.Period `json:"period,omitempty"`

	// The replacement could be because the initial care plan was immediately rejected or because the previous care plan was completed
	Replaces []common.Reference `json:"replaces,omitempty"`

	// The unknown code is not to be used to convey other statuses
	Status CarePlanStatus `json:"status"`

	// Identifies the patient or group whose intended care is described by the plan
	Subject common.Reference `json:"subject"`

	// Use "concern" to identify specific conditions addressed by the care plan
	SupportingInfo []common.Reference `json:"supportingInfo,omitempty"`

	// Human-friendly name for the care plan
	Title *string `json:"title,omitempty"`
}

// CareTeamStatus represents the status of a care team
type CareTeamStatus string

const (
	CareTeamStatusProposed       CareTeamStatus = "proposed"
	CareTeamStatusActive         CareTeamStatus = "active"
	CareTeamStatusSuspended      CareTeamStatus = "suspended"
	CareTeamStatusInactive       CareTeamStatus = "inactive"
	CareTeamStatusEnteredInError CareTeamStatus = "entered-in-error"
)

// CareTeamParticipant represents all people and organizations who are expected to be involved in the care team
type CareTeamParticipant struct {
	common.BackboneElement

	// This is populated while creating / managing the CareTeam to ensure there is coverage when servicing CarePlan activities
	CoveragePeriod *common.Period `json:"coveragePeriod,omitempty"`

	// This is populated while creating / managing the CareTeam to ensure there is coverage when servicing CarePlan activities
	CoverageTiming *Timing `json:"coverageTiming,omitempty"`

	// Patient only needs to be listed if they have a role other than "subject of care"
	Member *common.Reference `json:"member,omitempty"`

	// The organization of the practitioner
	OnBehalfOf *common.Reference `json:"onBehalfOf,omitempty"`

	// Roles may sometimes be inferred by type of Practitioner
	Role *common.CodeableConcept `json:"role,omitempty"`
}

// CareTeam represents the people and organizations who plan to participate in the coordination and delivery of care
type CareTeam struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "CareTeam"

	// There may be multiple axis of categorization and one team may serve multiple purposes
	Category []common.CodeableConcept `json:"category,omitempty"`

	// Business identifier for this care team
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The organization responsible for the care team
	ManagingOrganization []common.Reference `json:"managingOrganization,omitempty"`

	// The meaning/purpose of the team is conveyed in CareTeam.category
	Name *string `json:"name,omitempty"`

	// Comments made about the CareTeam
	Note []Annotation `json:"note,omitempty"`

	// Identifies all people and organizations who are expected to be involved in the care team
	Participant []CareTeamParticipant `json:"participant,omitempty"`

	// Indicates when the team did (or is intended to) come into effect and end
	Period *common.Period `json:"period,omitempty"`

	// Describes why the care team exists
	Reason []CodeableReference `json:"reason,omitempty"`

	// This element is labeled as a modifier because the status contains the code entered-in-error
	Status *CareTeamStatus `json:"status,omitempty"`

	// Use Group for care provision to all members of the group (e.g. group therapy)
	Subject *common.Reference `json:"subject,omitempty"`

	// The ContactPoint.use code of home is not appropriate to use
	Telecom []ContactPoint `json:"telecom,omitempty"`
}

// ChargeItemStatus represents the status of a charge item
type ChargeItemStatus string

const (
	ChargeItemStatusPlanned        ChargeItemStatus = "planned"
	ChargeItemStatusBillable       ChargeItemStatus = "billable"
	ChargeItemStatusNotBillable    ChargeItemStatus = "not-billable"
	ChargeItemStatusAborted        ChargeItemStatus = "aborted"
	ChargeItemStatusBilled         ChargeItemStatus = "billed"
	ChargeItemStatusEnteredInError ChargeItemStatus = "entered-in-error"
	ChargeItemStatusUnknown        ChargeItemStatus = "unknown"
)

// MonetaryComponentType represents the type of a monetary component
type MonetaryComponentType string

const (
	MonetaryComponentTypeBase          MonetaryComponentType = "base"
	MonetaryComponentTypeSurcharge     MonetaryComponentType = "surcharge"
	MonetaryComponentTypeDeduction     MonetaryComponentType = "deduction"
	MonetaryComponentTypeDiscount      MonetaryComponentType = "discount"
	MonetaryComponentTypeTax           MonetaryComponentType = "tax"
	MonetaryComponentTypeInformational MonetaryComponentType = "informational"
)

// MonetaryComponent represents a monetary value component
type MonetaryComponent struct {
	DataType

	// Explicit value amount to be used
	Amount *Money `json:"amount,omitempty"`

	// Codes may be used to differentiate between kinds of taxes, surcharges, discounts etc.
	Code *common.CodeableConcept `json:"code,omitempty"`

	// Factor used for calculating this component
	Factor *float64 `json:"factor,omitempty"`

	// base | surcharge | deduction | discount | tax | informational
	Type MonetaryComponentType `json:"type"`
}

// ChargeItemPerformer represents who performed or participated in the charged service
type ChargeItemPerformer struct {
	common.BackboneElement

	// The device, practitioner, etc. who performed or participated in the service
	Actor common.Reference `json:"actor"`

	// Describes the type of performance or participation(e.g. primary surgeon, anesthesiologist, etc.)
	Function *common.CodeableConcept `json:"function,omitempty"`
}

// ChargeItem represents the provision of healthcare provider products for a certain patient
type ChargeItem struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "ChargeItem"

	// Systems posting the ChargeItems might not always be able to determine, which accounts the Items need to be placed into
	Account []common.Reference `json:"account,omitempty"`

	// Only used if not implicit in code found in Condition.code
	Bodysite []common.CodeableConcept `json:"bodysite,omitempty"`

	// A code that identifies the charge, like a billing code
	Code common.CodeableConcept `json:"code"`

	// The costCenter could either be given as a reference to an Organization(Role) resource
	CostCenter *common.Reference `json:"costCenter,omitempty"`

	// References the source of pricing information, rules of application for the code this ChargeItem uses
	DefinitionCanonical []string `json:"definitionCanonical,omitempty"`

	// References the (external) source of pricing information, rules of application for the code this ChargeItem uses
	DefinitionUri []string `json:"definitionUri,omitempty"`

	// This ChargeItem may be recorded during planning, execution or after the actual encounter takes place
	Encounter *common.Reference `json:"encounter,omitempty"`

	// The actual date when the service associated with the charge has been rendered is captured in occurrence[x]
	EnteredDate *string `json:"enteredDate,omitempty"`

	// The enterer is also the person considered responsible for factor/price overrides if applicable
	Enterer *common.Reference `json:"enterer,omitempty"`

	// Identifiers assigned to this event performer or other systems
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Comments made about the event by the performer, subject or other participants
	Note []Annotation `json:"note,omitempty"`

	// The list of types may be constrained as appropriate for the type of charge item
	OccurrenceDateTime *string        `json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod   *common.Period `json:"occurrencePeriod,omitempty"`
	OccurrenceTiming   *Timing        `json:"occurrenceTiming,omitempty"`

	// If the list price or the rule-based factor associated with the code is overridden
	OverrideReason *common.CodeableConcept `json:"overrideReason,omitempty"`

	// ChargeItems can be grouped to larger ChargeItems covering the whole set
	PartOf []common.Reference `json:"partOf,omitempty"`

	// Indicates who or what performed or participated in the charged service
	Performer []ChargeItemPerformer `json:"performer,omitempty"`

	// Practitioners and Devices can be associated with multiple organizations
	PerformingOrganization *common.Reference `json:"performingOrganization,omitempty"`

	// Identifies the device, food, drug or other product being charged either by type code or reference to an instance
	Product []CodeableReference `json:"product,omitempty"`

	// In many cases this may just be a value, if the underlying units are implicit in the definition of the charge item code
	Quantity *common.Quantity `json:"quantity,omitempty"`

	// If the application of the charge item requires a reason to be given, it can be captured here
	Reason []common.CodeableConcept `json:"reason,omitempty"`

	// The rendered Service might not be associated with a Request
	RequestingOrganization *common.Reference `json:"requestingOrganization,omitempty"`

	// Indicated the rendered service that caused this charge
	Service []CodeableReference `json:"service,omitempty"`

	// Unknown does not represent "other" - one of the defined statuses must apply
	Status ChargeItemStatus `json:"status"`

	// The individual or set of individuals the action is being or was performed on
	Subject common.Reference `json:"subject"`

	// Further information supporting this charge
	SupportingInformation []common.Reference `json:"supportingInformation,omitempty"`

	// Often, the total price may be calculated and recorded on the Invoice
	TotalPriceComponent *MonetaryComponent `json:"totalPriceComponent,omitempty"`

	// This could be communicated in ChargeItemDefinition
	UnitPriceComponent *MonetaryComponent `json:"unitPriceComponent,omitempty"`
}

// DiagnosticReportStatus represents the status of a diagnostic report
type DiagnosticReportStatus string

const (
	DiagnosticReportStatusRegistered     DiagnosticReportStatus = "registered"
	DiagnosticReportStatusPartial        DiagnosticReportStatus = "partial"
	DiagnosticReportStatusPreliminary    DiagnosticReportStatus = "preliminary"
	DiagnosticReportStatusModified       DiagnosticReportStatus = "modified"
	DiagnosticReportStatusFinal          DiagnosticReportStatus = "final"
	DiagnosticReportStatusAmended        DiagnosticReportStatus = "amended"
	DiagnosticReportStatusCorrected      DiagnosticReportStatus = "corrected"
	DiagnosticReportStatusAppended       DiagnosticReportStatus = "appended"
	DiagnosticReportStatusCancelled      DiagnosticReportStatus = "cancelled"
	DiagnosticReportStatusEnteredInError DiagnosticReportStatus = "entered-in-error"
	DiagnosticReportStatusUnknown        DiagnosticReportStatus = "unknown"
)

// DiagnosticReportSupportingInfo represents supporting information used in the creation of the report
type DiagnosticReportSupportingInfo struct {
	common.BackboneElement

	// The reference for the supporting information in the diagnostic report
	Reference common.Reference `json:"reference"`

	// The code value for the role of the supporting information in the diagnostic report
	Type common.CodeableConcept `json:"type"`
}

// DiagnosticReportMedia represents key images or data associated with the report
type DiagnosticReportMedia struct {
	common.BackboneElement

	// The comment should be displayed with the image or data
	Comment *string `json:"comment,omitempty"`

	// Reference to the image or data source
	Link common.Reference `json:"link"`
}

// DiagnosticReport represents diagnostic test results and interpretation
type DiagnosticReport struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "DiagnosticReport"

	// Note: Usually there is one test request for each result, however in some circumstances multiple test requests may be represented using a single test result resource
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// Multiple categories are allowed using various categorization schemes
	Category []common.CodeableConcept `json:"category,omitempty"`

	// A code or name that describes this diagnostic report
	Code common.CodeableConcept `json:"code"`

	// The Composition provides structure to the content of the DiagnosticReport
	Composition *common.Reference `json:"composition,omitempty"`

	// Concise and clinically contextualized summary conclusion (interpretation/impression) of the diagnostic report
	Conclusion *string `json:"conclusion,omitempty"`

	// One or more codes that represent the summary conclusion (interpretation/impression) of the diagnostic report
	ConclusionCode []common.CodeableConcept `json:"conclusionCode,omitempty"`

	// If the diagnostic procedure was performed on the patient, this is the time it was performed
	EffectiveDateTime *string        `json:"effectiveDateTime,omitempty"`
	EffectivePeriod   *common.Period `json:"effectivePeriod,omitempty"`

	// This will typically be the encounter the event occurred within
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Usually assigned by the Information System of the diagnostic service provider (filler id)
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// May be different from the update time of the resource itself, because that is the status of the record
	Issued *string `json:"issued,omitempty"`

	// A list of key images or data associated with this report
	Media []DiagnosticReportMedia `json:"media,omitempty"`

	// May include general statements about the diagnostic report, or statements about significant, unexpected or unreliable results values
	Note []Annotation `json:"note,omitempty"`

	// This is not necessarily the source of the atomic data items or the entity that interpreted the results
	Performer []common.Reference `json:"performer,omitempty"`

	// "application/pdf" is recommended as the most reliable and interoperable in this context
	PresentedForm []Attachment `json:"presentedForm,omitempty"`

	// Observations can contain observations
	Result []common.Reference `json:"result,omitempty"`

	// Might not be the same entity that takes responsibility for the clinical report
	ResultsInterpreter []common.Reference `json:"resultsInterpreter,omitempty"`

	// If the specimen is sufficiently specified with a code in the test result name, then this additional data may be redundant
	Specimen []common.Reference `json:"specimen,omitempty"`

	// The status of the diagnostic report
	Status DiagnosticReportStatus `json:"status"`

	// For laboratory-type studies like GenomeStudy, type resources will be used for tracking additional metadata
	Study []common.Reference `json:"study,omitempty"`

	// The subject of the report. Usually, but not always, this is a patient
	Subject *common.Reference `json:"subject,omitempty"`

	// This backbone element contains supporting information that was used in the creation of the report
	SupportingInfo []DiagnosticReportSupportingInfo `json:"supportingInfo,omitempty"`
}

// ObservationStatus represents the status of an observation
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

// ObservationTriggeredByType represents the type of trigger for an observation
type ObservationTriggeredByType string

const (
	ObservationTriggeredByTypeReflex ObservationTriggeredByType = "reflex"
	ObservationTriggeredByTypeRepeat ObservationTriggeredByType = "repeat"
	ObservationTriggeredByTypeReRun  ObservationTriggeredByType = "re-run"
)

// ObservationTriggeredBy represents observations that triggered this observation
type ObservationTriggeredBy struct {
	common.BackboneElement

	// Reference to the triggering observation
	Observation common.Reference `json:"observation"`

	// Provides the reason why this observation was performed as a result of the observation(s) referenced
	Reason *string `json:"reason,omitempty"`

	// The type of trigger
	Type ObservationTriggeredByType `json:"type"`
}

// ObservationReferenceRange represents reference range for observation values
type ObservationReferenceRange struct {
	common.BackboneElement

	// The age at which this reference range is applicable
	Age *Range `json:"age,omitempty"`

	// This SHOULD be populated if there is more than one range
	AppliesTo []common.CodeableConcept `json:"appliesTo,omitempty"`

	// The value of the high bound of the reference range
	High *common.Quantity `json:"high,omitempty"`

	// The value of the low bound of the reference range
	Low *common.Quantity `json:"low,omitempty"`

	// The value of the normal value of the reference range
	NormalValue *common.CodeableConcept `json:"normalValue,omitempty"`

	// Text based reference range in an observation
	Text *string `json:"text,omitempty"`

	// This SHOULD be populated if there is more than one range
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// ObservationComponent represents component observations
type ObservationComponent struct {
	common.BackboneElement

	// Describes what was observed
	Code common.CodeableConcept `json:"code"`

	// For exceptional values in measurements
	DataAbsentReason *common.CodeableConcept `json:"dataAbsentReason,omitempty"`

	// Historically used for laboratory results (known as 'abnormal flag')
	Interpretation []common.CodeableConcept `json:"interpretation,omitempty"`

	// Reference ranges for the component
	ReferenceRange []ObservationReferenceRange `json:"referenceRange,omitempty"`

	// Component observation values - choice of types
	ValueQuantity        *common.Quantity        `json:"valueQuantity,omitempty"`
	ValueCodeableConcept *common.CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueString          *string                 `json:"valueString,omitempty"`
	ValueBoolean         *bool                   `json:"valueBoolean,omitempty"`
	ValueInteger         *int                    `json:"valueInteger,omitempty"`
	ValueRange           *Range                  `json:"valueRange,omitempty"`
	ValueRatio           *Ratio                  `json:"valueRatio,omitempty"`
	ValueSampledData     *common.SampledData     `json:"valueSampledData,omitempty"`
	ValueTime            *string                 `json:"valueTime,omitempty"`
	ValueDateTime        *string                 `json:"valueDateTime,omitempty"`
	ValuePeriod          *common.Period          `json:"valuePeriod,omitempty"`
	ValueAttachment      *Attachment             `json:"valueAttachment,omitempty"`
	ValueReference       *common.Reference       `json:"valueReference,omitempty"`
}

// Observation represents measurements and simple assertions made about a patient, device or other subject
type Observation struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Observation"

	// A plan, proposal or order that is fulfilled in whole or in part by this event
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// Only used if not implicit in code found in Observation.code
	BodySite *common.CodeableConcept `json:"bodySite,omitempty"`

	// Only used if not implicit in code found in Observation.code or bodySite is used
	BodyStructure *common.Reference `json:"bodyStructure,omitempty"`

	// In addition to the required category valueset, this element allows various categorization schemes
	Category []common.CodeableConcept `json:"category,omitempty"`

	// Describes what was observed
	Code common.CodeableConcept `json:"code"`

	// For a discussion on the ways Observations can be assembled in groups together
	Component []ObservationComponent `json:"component,omitempty"`

	// Null or exceptional values can be represented two ways in FHIR Observations
	DataAbsentReason *common.CodeableConcept `json:"dataAbsentReason,omitempty"`

	// All the reference choices that are listed in this element can represent clinical observations
	DerivedFrom []common.Reference `json:"derivedFrom,omitempty"`

	// Note that this is not meant to represent a device involved in the transmission of the result
	Device *common.Reference `json:"device,omitempty"`

	// At least a date should be present unless this observation is a historical report
	EffectiveDateTime *string        `json:"effectiveDateTime,omitempty"`
	EffectivePeriod   *common.Period `json:"effectivePeriod,omitempty"`
	EffectiveTiming   *Timing        `json:"effectiveTiming,omitempty"`
	EffectiveInstant  *string        `json:"effectiveInstant,omitempty"`

	// This will typically be the encounter the event occurred within
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Typically, an observation is made about the subject
	Focus []common.Reference `json:"focus,omitempty"`

	// When using this element, an observation will typically have either a value or a set of related resources
	HasMember []common.Reference `json:"hasMember,omitempty"`

	// A unique identifier assigned to this observation
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// ObservationDefinition can be referenced by its canonical url using instantiatesCanonical
	InstantiatesCanonical *string `json:"instantiatesCanonical,omitempty"`

	// ObservationDefinition can be referenced by its canonical url using instantiatesCanonical
	InstantiatesReference *common.Reference `json:"instantiatesReference,omitempty"`

	// Historically used for laboratory results (known as 'abnormal flag')
	Interpretation []common.CodeableConcept `json:"interpretation,omitempty"`

	// For Observations that don't require review and verification
	Issued *string `json:"issued,omitempty"`

	// Only used if not implicit in code for Observation.code
	Method *common.CodeableConcept `json:"method,omitempty"`

	// May include general statements about the observation
	Note []Annotation `json:"note,omitempty"`

	// To link an Observation to an Encounter use `encounter`
	PartOf []common.Reference `json:"partOf,omitempty"`

	// Who was responsible for asserting the observed value as "true"
	Performer []common.Reference `json:"performer,omitempty"`

	// Most observations only have one generic reference range
	ReferenceRange []ObservationReferenceRange `json:"referenceRange,omitempty"`

	// Should only be used if not implicit in code found in `Observation.code`
	Specimen *common.Reference `json:"specimen,omitempty"`

	// This element is labeled as a modifier because the status contains codes that mark the resource as not currently valid
	Status ObservationStatus `json:"status"`

	// One would expect this element to be a cardinality of 1..1
	Subject *common.Reference `json:"subject,omitempty"`

	// Identifies the observation(s) that triggered the performance of this observation
	TriggeredBy []ObservationTriggeredBy `json:"triggeredBy,omitempty"`

	// An observation may have a single value here, both a value and a set of related or component values, or only a set of related or component values
	ValueQuantity        *common.Quantity        `json:"valueQuantity,omitempty"`
	ValueCodeableConcept *common.CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueString          *string                 `json:"valueString,omitempty"`
	ValueBoolean         *bool                   `json:"valueBoolean,omitempty"`
	ValueInteger         *int                    `json:"valueInteger,omitempty"`
	ValueRange           *Range                  `json:"valueRange,omitempty"`
	ValueRatio           *Ratio                  `json:"valueRatio,omitempty"`
	ValueSampledData     *common.SampledData     `json:"valueSampledData,omitempty"`
	ValueTime            *string                 `json:"valueTime,omitempty"`
	ValueDateTime        *string                 `json:"valueDateTime,omitempty"`
	ValuePeriod          *common.Period          `json:"valuePeriod,omitempty"`
	ValueAttachment      *Attachment             `json:"valueAttachment,omitempty"`
	ValueReference       *common.Reference       `json:"valueReference,omitempty"`
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

// PatientContact represents contact information for a patient
type PatientContact struct {
	common.BackboneElement

	// The nature of the relationship between the patient and the contact person
	Relationship []common.CodeableConcept `json:"relationship,omitempty"`

	// A name associated with the contact person
	Name *HumanName `json:"name,omitempty"`

	// A contact detail for the person
	Telecom []ContactPoint `json:"telecom,omitempty"`

	// Address for the contact person
	Address *Address `json:"address,omitempty"`

	// Administrative Gender
	Gender *PatientGender `json:"gender,omitempty"`

	// Organization on behalf of which the contact is acting or for which the contact is working
	Organization *common.Reference `json:"organization,omitempty"`

	// The period during which this contact person or organization is valid to be contacted relating to this patient
	Period *common.Period `json:"period,omitempty"`
}

// PatientCommunication represents patient communication preferences
type PatientCommunication struct {
	common.BackboneElement

	// The language which can be used to communicate with the patient about his or her health
	Language common.CodeableConcept `json:"language"`

	// Indicates whether or not the patient prefers this language
	Preferred *bool `json:"preferred,omitempty"`
}

// PatientLink represents a link to another patient resource
type PatientLink struct {
	common.BackboneElement

	// Link to a Patient or RelatedPerson resource that concerns the same actual person
	Other common.Reference `json:"other"`

	// The type of link between this patient resource and another Patient resource
	Type PatientLinkType `json:"type"`
}

// Patient represents demographics and other administrative information about an individual or animal receiving care
type Patient struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Patient"

	// Whether this patient record is in active use
	Active *bool `json:"active,omitempty"`

	// An address for the individual
	Address []Address `json:"address,omitempty"`

	// The date of birth for the individual
	BirthDate *string `json:"birthDate,omitempty"`

	// A language which may be used to communicate with the patient about his or her health
	Communication []PatientCommunication `json:"communication,omitempty"`

	// A contact party for the patient
	Contact []PatientContact `json:"contact,omitempty"`

	// Indicates if the individual is deceased or not
	DeceasedBoolean  *bool   `json:"deceasedBoolean,omitempty"`
	DeceasedDateTime *string `json:"deceasedDateTime,omitempty"`

	// Administrative Gender
	Gender *PatientGender `json:"gender,omitempty"`

	// This may be the primary care provider
	GeneralPractitioner []common.Reference `json:"generalPractitioner,omitempty"`

	// An identifier for this patient
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Link to a Patient or RelatedPerson resource that concerns the same actual individual
	Link []PatientLink `json:"link,omitempty"`

	// Organization that is the custodian of the patient record
	ManagingOrganization *common.Reference `json:"managingOrganization,omitempty"`

	// Marital (civil) status of a patient
	MaritalStatus *common.CodeableConcept `json:"maritalStatus,omitempty"`

	// Whether patient is part of a multiple birth
	MultipleBirthBoolean *bool `json:"multipleBirthBoolean,omitempty"`
	MultipleBirthInteger *int  `json:"multipleBirthInteger,omitempty"`

	// A name associated with the individual
	Name []HumanName `json:"name,omitempty"`

	// Image of the patient
	Photo []Attachment `json:"photo,omitempty"`

	// A contact detail for the individual
	Telecom []ContactPoint `json:"telecom,omitempty"`
}

// ProcedureStatus represents the status of a procedure
type ProcedureStatus string

const (
	ProcedureStatusPreparation    ProcedureStatus = "preparation"
	ProcedureStatusInProgress     ProcedureStatus = "in-progress"
	ProcedureStatusNotDone        ProcedureStatus = "not-done"
	ProcedureStatusOnHold         ProcedureStatus = "on-hold"
	ProcedureStatusStopped        ProcedureStatus = "stopped"
	ProcedureStatusCompleted      ProcedureStatus = "completed"
	ProcedureStatusEnteredInError ProcedureStatus = "entered-in-error"
	ProcedureStatusUnknown        ProcedureStatus = "unknown"
)

// ProcedurePerformer represents who performed the procedure
type ProcedurePerformer struct {
	common.BackboneElement

	// Indicates who or what performed the procedure
	Actor common.Reference `json:"actor"`

	// Distinguishes the type of involvement of the performer in the procedure
	Function *common.CodeableConcept `json:"function,omitempty"`

	// Organization they were acting on behalf of when performing the action
	OnBehalfOf *common.Reference `json:"onBehalfOf,omitempty"`

	// Time period during which the performer performed the procedure
	Period *common.Period `json:"period,omitempty"`
}

// ProcedureFocalDevice represents a device that was manipulated during the procedure
type ProcedureFocalDevice struct {
	common.BackboneElement

	// The kind of change that happened to the device during the procedure
	Action *common.CodeableConcept `json:"action,omitempty"`

	// The device that was manipulated (changed) during the procedure
	Manipulated common.Reference `json:"manipulated"`
}

// Procedure represents an action that is or was performed on or for a patient
type Procedure struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Procedure"

	// A reference to a resource that contains details of the request for this procedure
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// Specific anatomical location where the procedure was performed
	BodySite []common.CodeableConcept `json:"bodySite,omitempty"`

	// A code that classifies the procedure for searching, sorting and display purposes
	Category []common.CodeableConcept `json:"category,omitempty"`

	// The specific procedure that is performed
	Code *common.CodeableConcept `json:"code,omitempty"`

	// Complications that occurred during the procedure
	Complication []CodeableReference `json:"complication,omitempty"`

	// The encounter during which the procedure was performed
	Encounter *common.Reference `json:"encounter,omitempty"`

	// A device that is implanted, removed or otherwise manipulated during the procedure
	FocalDevice []ProcedureFocalDevice `json:"focalDevice,omitempty"`

	// Who is the target of the procedure when it is not the subject of record only
	Focus *common.Reference `json:"focus,omitempty"`

	// If the procedure required specific follow up
	FollowUp []common.CodeableConcept `json:"followUp,omitempty"`

	// Business identifier for this procedure
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The URL pointing to a FHIR-defined protocol, guideline, order set or other definition
	InstantiatesCanonical []string `json:"instantiatesCanonical,omitempty"`

	// The URL pointing to an externally maintained protocol, guideline, order set
	InstantiatesUri []string `json:"instantiatesUri,omitempty"`

	// Where the procedure happened
	Location *common.Reference `json:"location,omitempty"`

	// Any other notes and comments about the procedure
	Note []Annotation `json:"note,omitempty"`

	// When the procedure occurred
	OccurrenceDateTime *string        `json:"occurrenceDateTime,omitempty"`
	OccurrenceAge      *Age           `json:"occurrenceAge,omitempty"`
	OccurrencePeriod   *common.Period `json:"occurrencePeriod,omitempty"`
	OccurrenceRange    *Range         `json:"occurrenceRange,omitempty"`
	OccurrenceString   *string        `json:"occurrenceString,omitempty"`
	OccurrenceTiming   *Timing        `json:"occurrenceTiming,omitempty"`

	// The outcome of the procedure
	Outcome *common.CodeableConcept `json:"outcome,omitempty"`

	// Part of referenced resource
	PartOf []common.Reference `json:"partOf,omitempty"`

	// Who performed the procedure and how they were involved
	Performer []ProcedurePerformer `json:"performer,omitempty"`

	// The reason why the procedure was performed
	Reason []CodeableReference `json:"reason,omitempty"`

	// The date the occurrence of the procedure was first captured in the record
	Recorded *string `json:"recorded,omitempty"`

	// Individual who recorded the record and takes responsibility for its content
	Recorder *common.Reference `json:"recorder,omitempty"`

	// There could potentially be multiple reports
	Report []common.Reference `json:"report,omitempty"`

	// Indicates if this record was captured as a secondary 'reported' record
	ReportedBoolean   *bool             `json:"reportedBoolean,omitempty"`
	ReportedReference *common.Reference `json:"reportedReference,omitempty"`

	// The status of the procedure
	Status ProcedureStatus `json:"status"`

	// Captures the reason for the current state of the procedure
	StatusReason *common.CodeableConcept `json:"statusReason,omitempty"`

	// On whom or on what the procedure was performed
	Subject common.Reference `json:"subject"`

	// Other resources from the patient record that may be relevant to the procedure
	SupportingInfo []common.Reference `json:"supportingInfo,omitempty"`

	// For devices actually implanted or removed, use Procedure.focalDevice.manipulated
	Used []CodeableReference `json:"used,omitempty"`
}

// ServiceRequestStatus represents the status of a service request
type ServiceRequestStatus string

const (
	ServiceRequestStatusDraft          ServiceRequestStatus = "draft"
	ServiceRequestStatusActive         ServiceRequestStatus = "active"
	ServiceRequestStatusOnHold         ServiceRequestStatus = "on-hold"
	ServiceRequestStatusRevoked        ServiceRequestStatus = "revoked"
	ServiceRequestStatusCompleted      ServiceRequestStatus = "completed"
	ServiceRequestStatusEnteredInError ServiceRequestStatus = "entered-in-error"
	ServiceRequestStatusUnknown        ServiceRequestStatus = "unknown"
)

// ServiceRequestPriority represents the priority of a service request
type ServiceRequestPriority string

const (
	ServiceRequestPriorityRoutine ServiceRequestPriority = "routine"
	ServiceRequestPriorityUrgent  ServiceRequestPriority = "urgent"
	ServiceRequestPriorityASAP    ServiceRequestPriority = "asap"
	ServiceRequestPriorityStat    ServiceRequestPriority = "stat"
)

// ServiceRequestIntent represents the intent of a service request
type ServiceRequestIntent string

const (
	ServiceRequestIntentProposal      ServiceRequestIntent = "proposal"
	ServiceRequestIntentPlan          ServiceRequestIntent = "plan"
	ServiceRequestIntentDirective     ServiceRequestIntent = "directive"
	ServiceRequestIntentOrder         ServiceRequestIntent = "order"
	ServiceRequestIntentOriginalOrder ServiceRequestIntent = "original-order"
	ServiceRequestIntentReflexOrder   ServiceRequestIntent = "reflex-order"
	ServiceRequestIntentFillerOrder   ServiceRequestIntent = "filler-order"
	ServiceRequestIntentInstanceOrder ServiceRequestIntent = "instance-order"
	ServiceRequestIntentOption        ServiceRequestIntent = "option"
)

// ServiceRequest represents a record of a request for service such as diagnostic investigations, treatments, or operations
type ServiceRequest struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "ServiceRequest"

	// If a CodeableConcept is present, it indicates the pre-condition for performing the service
	AsNeededBoolean         *bool                   `json:"asNeededBoolean,omitempty"`
	AsNeededCodeableConcept *common.CodeableConcept `json:"asNeededCodeableConcept,omitempty"`

	// When the request transitioned to being actionable
	AuthoredOn *string `json:"authoredOn,omitempty"`

	// Plan/proposal/order fulfilled by this request
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// Anatomical location where the procedure should be performed
	BodySite []common.CodeableConcept `json:"bodySite,omitempty"`

	// Anatomic location where the procedure should be performed (R5 feature)
	BodyStructure *common.Reference `json:"bodyStructure,omitempty"`

	// A code that classifies the procedure for searching, sorting and display purposes
	Category []common.CodeableConcept `json:"category,omitempty"`

	// What is being requested/ordered
	Code *common.CodeableConcept `json:"code,omitempty"`

	// If do not perform is not specified, the request is a positive request
	DoNotPerform *bool `json:"doNotPerform,omitempty"`

	// An encounter that provides additional context in which this request is made
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Business identifier for this service request
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The URL pointing to a FHIR-defined protocol, guideline, order set or other definition
	InstantiatesCanonical []string `json:"instantiatesCanonical,omitempty"`

	// The URL pointing to an externally maintained protocol, guideline, order set
	InstantiatesUri []string `json:"instantiatesUri,omitempty"`

	// Additional clinical information
	Instructions []Instruction `json:"instructions,omitempty"`

	// Insurance plans, coverage extensions, pre-authorizations and/or pre-determinations
	Insurance []common.Reference `json:"insurance,omitempty"`

	// proposal | plan | directive | order | original-order | reflex-order | filler-order | instance-order | option
	Intent ServiceRequestIntent `json:"intent"`

	// Where procedure is going to be done
	Location []common.Reference `json:"location,omitempty"`

	// Any other notes and comments about the service request
	Note []Annotation `json:"note,omitempty"`

	// When service should occur
	OccurrenceDateTime *string        `json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod   *common.Period `json:"occurrencePeriod,omitempty"`
	OccurrenceTiming   *Timing        `json:"occurrenceTiming,omitempty"`

	// What order this request fulfills
	OrderDetail []CodeableReference `json:"orderDetail,omitempty"`

	// The desired performer for doing the requested service
	Performer []common.Reference `json:"performer,omitempty"`

	// Performer role
	PerformerType *common.CodeableConcept `json:"performerType,omitempty"`

	// Indicates how quickly the ServiceRequest should be addressed
	Priority *ServiceRequestPriority `json:"priority,omitempty"`

	// An amount of service being requested which can be a quantity, ratio, or range
	QuantityQuantity *common.Quantity `json:"quantityQuantity,omitempty"`
	QuantityRatio    *Ratio           `json:"quantityRatio,omitempty"`
	QuantityRange    *Range           `json:"quantityRange,omitempty"`

	// Explanation/Justification for procedure or service
	Reason []CodeableReference `json:"reason,omitempty"`

	// Key events in the history of the request
	RelevantHistory []common.Reference `json:"relevantHistory,omitempty"`

	// What request replaces
	Replaces []common.Reference `json:"replaces,omitempty"`

	// Who/what is requesting service
	Requester *common.Reference `json:"requester,omitempty"`

	// Composite Request ID
	Requisition *common.Identifier `json:"requisition,omitempty"`

	// Procedure Samples
	Specimen []common.Reference `json:"specimen,omitempty"`

	// The status of the order
	Status ServiceRequestStatus `json:"status"`

	// Explanation for current status
	StatusReason *common.CodeableConcept `json:"statusReason,omitempty"`

	// Individual or Entity the service is ordered for
	Subject common.Reference `json:"subject"`

	// Additional clinical information
	SupportingInfo []CodeableReference `json:"supportingInfo,omitempty"`
}

// Instruction represents instruction details for the service request
type Instruction struct {
	DataType

	// Describes the instruction
	InstructionText *string `json:"instructionText,omitempty"`
}

// LocationStatus represents the status of a location
type LocationStatus string

const (
	LocationStatusActive    LocationStatus = "active"
	LocationStatusSuspended LocationStatus = "suspended"
	LocationStatusInactive  LocationStatus = "inactive"
)

// LocationMode represents the mode of a location
type LocationMode string

const (
	LocationModeInstance LocationMode = "instance"
	LocationModeKind     LocationMode = "kind"
)

// LocationPosition represents the position of a location
type LocationPosition struct {
	common.BackboneElement

	// Altitude with the same coordinate system as KML
	Altitude *float64 `json:"altitude,omitempty"`

	// Latitude with the same coordinate system as KML
	Latitude float64 `json:"latitude"`

	// Longitude with the same coordinate system as KML
	Longitude float64 `json:"longitude"`
}

// Location represents a place where services are provided and resources may be stored
type Location struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Location"

	// Address of the location
	Address *Address `json:"address,omitempty"`

	// Alternate names that the location is known as
	Alias []string `json:"alias,omitempty"`

	// Collection of characteristics that describe the location
	Characteristic []common.CodeableConcept `json:"characteristic,omitempty"`

	// Official contact details for the location
	Contact []ExtendedContactDetail `json:"contact,omitempty"`

	// Additional details about the location
	Description *string `json:"description,omitempty"`

	// Technical endpoints providing access to services operated for the location
	Endpoint []common.Reference `json:"endpoint,omitempty"`

	// Physical form of the location
	Form *common.CodeableConcept `json:"form,omitempty"`

	// What days/times during a week is this location usually open
	HoursOfOperation []Availability `json:"hoursOfOperation,omitempty"`

	// Unique code or number identifying the location to its users
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Organization responsible for provisioning and upkeep
	ManagingOrganization *common.Reference `json:"managingOrganization,omitempty"`

	// Whether this location is a class of locations or a specific location
	Mode *LocationMode `json:"mode,omitempty"`

	// Name of the location as used by humans
	Name *string `json:"name,omitempty"`

	// The operational status of the location
	OperationalStatus *common.Coding `json:"operationalStatus,omitempty"`

	// Another location this location is physically part of
	PartOf *common.Reference `json:"partOf,omitempty"`

	// The absolute geographic location
	Position *LocationPosition `json:"position,omitempty"`

	// Whether the location is still in active use
	Status *LocationStatus `json:"status,omitempty"`

	// Type of function performed at the location
	Type []common.CodeableConcept `json:"type,omitempty"`

	// Connection details for virtual services
	VirtualService []VirtualServiceDetail `json:"virtualService,omitempty"`
}

// DocumentReferenceStatus represents the status of a document reference
type DocumentReferenceStatus string

const (
	DocumentReferenceStatusCurrent        DocumentReferenceStatus = "current"
	DocumentReferenceStatusSuperseded     DocumentReferenceStatus = "superseded"
	DocumentReferenceStatusEnteredInError DocumentReferenceStatus = "entered-in-error"
)

// DocumentReferenceAttester represents who attested the document
type DocumentReferenceAttester struct {
	common.BackboneElement

	// The type of attestation the authenticator offers
	Mode common.CodeableConcept `json:"mode"`

	// Who attested the document in the specified way
	Party *common.Reference `json:"party,omitempty"`

	// When the document was attested by the party
	Time *string `json:"time,omitempty"`
}

// DocumentReferenceRelatesTo represents relationships this document has with other documents
type DocumentReferenceRelatesTo struct {
	common.BackboneElement

	// The type of relationship between documents
	Code common.CodeableConcept `json:"code"`

	// The target document of this relationship
	Target common.Reference `json:"target"`
}

// DocumentReferenceContentProfile represents the profile information for document content
type DocumentReferenceContentProfile struct {
	common.BackboneElement

	// Profile information as a coding, URI, or canonical reference
	ValueCoding    *common.Coding `json:"valueCoding,omitempty"`
	ValueUri       *string        `json:"valueUri,omitempty"`
	ValueCanonical *string        `json:"valueCanonical,omitempty"`
}

// DocumentReferenceContent represents the content of the document
type DocumentReferenceContent struct {
	common.BackboneElement

	// The document or URL of the document with critical metadata
	Attachment Attachment `json:"attachment"`

	// Format/content rules for the document
	Profile []DocumentReferenceContentProfile `json:"profile,omitempty"`
}

// DocumentReference represents a reference to a document of any kind
type DocumentReference struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "DocumentReference"

	// Who attested the document
	Attester []DocumentReferenceAttester `json:"attester,omitempty"`

	// Who and/or what authored the document
	Author []common.Reference `json:"author,omitempty"`

	// Procedure that is fulfilled in whole or part by the creation of this document
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// Anatomic structures included in the document
	BodySite []CodeableReference `json:"bodySite,omitempty"`

	// Categorization of document
	Category []common.CodeableConcept `json:"category,omitempty"`

	// Document referenced
	Content []DocumentReferenceContent `json:"content"`

	// Context of the document encounter
	Context []common.Reference `json:"context,omitempty"`

	// Organization which maintains the document
	Custodian *common.Reference `json:"custodian,omitempty"`

	// When this document reference was created
	Date *string `json:"date,omitempty"`

	// Human-readable description
	Description *string `json:"description,omitempty"`

	// Document security-tags
	Facility *common.CodeableConcept `json:"facility,omitempty"`

	// Business identifier
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Time period over which the service described by the document was provided
	Period *common.Period `json:"period,omitempty"`

	// Additional details about where the content was created
	PracticeSetting *common.CodeableConcept `json:"practiceSetting,omitempty"`

	// Relationships to other documents
	RelatesTo []DocumentReferenceRelatesTo `json:"relatesTo,omitempty"`

	// Document security-tags
	SecurityLabel []common.CodeableConcept `json:"securityLabel,omitempty"`

	// Current | superseded | entered-in-error
	Status DocumentReferenceStatus `json:"status"`

	// Who/what is the subject of the document
	Subject *common.Reference `json:"subject,omitempty"`

	// Kind of document
	Type *common.CodeableConcept `json:"type,omitempty"`

	// Document version
	Version *string `json:"version,omitempty"`
}
