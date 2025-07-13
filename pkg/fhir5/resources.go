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

// EncounterHistoryStatus represents the status of an encounter history (reuses EncounterStatus)
type EncounterHistoryStatus = EncounterStatus

// EncounterHistoryLocation represents location information in encounter history
type EncounterHistoryLocation struct {
	common.BackboneElement

	// Physical form of the location
	Form *common.CodeableConcept `json:"form,omitempty"`

	// The location where the encounter takes place
	Location common.Reference `json:"location"`
}

// EncounterHistory represents significant events throughout the history of an Encounter
type EncounterHistory struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "EncounterHistory"

	// The start and end time associated with this set of values
	ActualPeriod *common.Period `json:"actualPeriod,omitempty"`

	// Classification of patient encounter
	Class common.CodeableConcept `json:"class"`

	// The encounter this history relates to
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Identifier(s) by which this encounter history is known
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Quantity of time the encounter lasted
	Length *Duration `json:"length,omitempty"`

	// List of locations where the patient has been during this encounter
	Location []EncounterHistoryLocation `json:"location,omitempty"`

	// The planned end date/time of the encounter
	PlannedEndDate *string `json:"plannedEndDate,omitempty"`

	// The planned start date/time of the encounter
	PlannedStartDate *string `json:"plannedStartDate,omitempty"`

	// Broad categorization of the service that is to be provided
	ServiceType []CodeableReference `json:"serviceType,omitempty"`

	// Current status of the encounter history
	Status EncounterHistoryStatus `json:"status"`

	// Who/what is the subject of the encounter
	Subject *common.Reference `json:"subject,omitempty"`

	// Patient's status within the encounter
	SubjectStatus *common.CodeableConcept `json:"subjectStatus,omitempty"`

	// Specific type of encounter
	Type []common.CodeableConcept `json:"type,omitempty"`
}

// EndpointStatus represents the status of an endpoint
type EndpointStatus string

const (
	EndpointStatusActive         EndpointStatus = "active"
	EndpointStatusSuspended      EndpointStatus = "suspended"
	EndpointStatusError          EndpointStatus = "error"
	EndpointStatusOff            EndpointStatus = "off"
	EndpointStatusEnteredInError EndpointStatus = "entered-in-error"
)

// EndpointPayload represents the payload type and format for the endpoint
type EndpointPayload struct {
	common.BackboneElement

	// MIME types supported by the endpoint
	MimeType []string `json:"mimeType,omitempty"`

	// The payload type
	Type []common.CodeableConcept `json:"type,omitempty"`
}

// Endpoint represents a technical endpoint providing access to services
type Endpoint struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Endpoint"

	// The technical base address for connecting to this endpoint
	Address string `json:"address"`

	// Connection type details
	ConnectionType []common.CodeableConcept `json:"connectionType"`

	// Contact details for endpoint
	Contact []ContactPoint `json:"contact,omitempty"`

	// Description of the endpoint
	Description *string `json:"description,omitempty"`

	// Environment type
	EnvironmentType []common.CodeableConcept `json:"environmentType,omitempty"`

	// Usage depends on the channel type
	Header []string `json:"header,omitempty"`

	// Identifier for the endpoint
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Organization that manages this endpoint
	ManagingOrganization *common.Reference `json:"managingOrganization,omitempty"`

	// A name that this endpoint can be referred to with
	Name *string `json:"name,omitempty"`

	// Set of payloads that are provided by this endpoint
	Payload []EndpointPayload `json:"payload,omitempty"`

	// Interval during which the endpoint is expected to be operational
	Period *common.Period `json:"period,omitempty"`

	// Active | suspended | error | off | entered-in-error
	Status EndpointStatus `json:"status"`
}

// EnrollmentRequestStatus represents the status of an enrollment request
type EnrollmentRequestStatus string

const (
	EnrollmentRequestStatusActive         EnrollmentRequestStatus = "active"
	EnrollmentRequestStatusCancelled      EnrollmentRequestStatus = "cancelled"
	EnrollmentRequestStatusDraft          EnrollmentRequestStatus = "draft"
	EnrollmentRequestStatusEnteredInError EnrollmentRequestStatus = "entered-in-error"
)

// EnrollmentRequest provides insurance enrollment details to an insurer
type EnrollmentRequest struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "EnrollmentRequest"

	// Patient Resource
	Candidate *common.Reference `json:"candidate,omitempty"`

	// Reference to the program or plan identification, underwriter or payor
	Coverage *common.Reference `json:"coverage,omitempty"`

	// The date when this resource was created
	Created *string `json:"created,omitempty"`

	// The Response business identifier
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The Insurer who is target of the request
	Insurer *common.Reference `json:"insurer,omitempty"`

	// The practitioner who is responsible for the services rendered to the patient
	Provider *common.Reference `json:"provider,omitempty"`

	// This element is labeled as a modifier because the status contains codes that mark the request as not currently valid
	Status *EnrollmentRequestStatus `json:"status,omitempty"`
}

// EnrollmentResponseStatus represents the status of an enrollment response
type EnrollmentResponseStatus string

const (
	EnrollmentResponseStatusActive         EnrollmentResponseStatus = "active"
	EnrollmentResponseStatusCancelled      EnrollmentResponseStatus = "cancelled"
	EnrollmentResponseStatusDraft          EnrollmentResponseStatus = "draft"
	EnrollmentResponseStatusEnteredInError EnrollmentResponseStatus = "entered-in-error"
)

// EnrollmentResponseOutcome represents the outcome of an enrollment response
type EnrollmentResponseOutcome string

const (
	EnrollmentResponseOutcomeQueued   EnrollmentResponseOutcome = "queued"
	EnrollmentResponseOutcomeComplete EnrollmentResponseOutcome = "complete"
	EnrollmentResponseOutcomeError    EnrollmentResponseOutcome = "error"
	EnrollmentResponseOutcomePartial  EnrollmentResponseOutcome = "partial"
)

// EnrollmentResponse provides enrollment and plan details from the processing of an EnrollmentRequest resource
type EnrollmentResponse struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "EnrollmentResponse"

	// The date when the enclosed suite of services were performed or completed
	Created *string `json:"created,omitempty"`

	// A description of the status of the adjudication
	Disposition *string `json:"disposition,omitempty"`

	// The Response business identifier
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The Insurer who produced this adjudicated response
	Organization *common.Reference `json:"organization,omitempty"`

	// Processing status: error, complete
	Outcome *EnrollmentResponseOutcome `json:"outcome,omitempty"`

	// Original request resource reference
	Request *common.Reference `json:"request,omitempty"`

	// The practitioner who is responsible for the services rendered to the patient
	RequestProvider *common.Reference `json:"requestProvider,omitempty"`

	// This element is labeled as a modifier because the status contains codes that mark the response as not currently valid
	Status *EnrollmentResponseStatus `json:"status,omitempty"`
}

// EpisodeOfCareStatus represents the status of an episode of care
type EpisodeOfCareStatus string

const (
	EpisodeOfCareStatusPlanned        EpisodeOfCareStatus = "planned"
	EpisodeOfCareStatusWaitlist       EpisodeOfCareStatus = "waitlist"
	EpisodeOfCareStatusActive         EpisodeOfCareStatus = "active"
	EpisodeOfCareStatusOnHold         EpisodeOfCareStatus = "onhold"
	EpisodeOfCareStatusFinished       EpisodeOfCareStatus = "finished"
	EpisodeOfCareStatusCancelled      EpisodeOfCareStatus = "cancelled"
	EpisodeOfCareStatusEnteredInError EpisodeOfCareStatus = "entered-in-error"
)

// EpisodeOfCareStatusHistory represents the history of statuses that the EpisodeOfCare has been through
type EpisodeOfCareStatusHistory struct {
	common.BackboneElement

	// The period during this EpisodeOfCare that the specific status applied
	Period common.Period `json:"period"`

	// planned | waitlist | active | onhold | finished | cancelled | entered-in-error
	Status EpisodeOfCareStatus `json:"status"`
}

// EpisodeOfCareReason represents the reason for the episode of care
type EpisodeOfCareReason struct {
	common.BackboneElement

	// What the reason value should be used as e.g. Chief Complaint, Health Concern, Health Maintenance (including screening)
	Use *common.CodeableConcept `json:"use,omitempty"`

	// The medical reason that is expected to be addressed during the episode of care
	Value []CodeableReference `json:"value,omitempty"`
}

// EpisodeOfCareDiagnosis represents the diagnosis that was addressed during the episode of care
type EpisodeOfCareDiagnosis struct {
	common.BackboneElement

	// The medical condition that was addressed during the episode of care
	Condition []CodeableReference `json:"condition,omitempty"`

	// Role that this diagnosis has within the episode of care (e.g. admission, billing, discharge)
	Use *common.CodeableConcept `json:"use,omitempty"`
}

// EpisodeOfCare represents an association between a patient and an organization/healthcare provider during which encounters may occur
type EpisodeOfCare struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "EpisodeOfCare"

	// The billing system may choose to allocate billable items associated with the EpisodeOfCare to different referenced Accounts
	Account []common.Reference `json:"account,omitempty"`

	// The practitioner that is the care manager/care coordinator for this patient
	CareManager *common.Reference `json:"careManager,omitempty"`

	// The list of practitioners that may be facilitating this episode of care for specific purposes
	CareTeam []common.Reference `json:"careTeam,omitempty"`

	// The diagnosis communicates what medical conditions were actually addressed during the episode of care
	Diagnosis []EpisodeOfCareDiagnosis `json:"diagnosis,omitempty"`

	// The EpisodeOfCare may be known by different identifiers for different contexts of use
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The organization that is primarily responsible for the care management
	ManagingOrganization *common.Reference `json:"managingOrganization,omitempty"`

	// The patient who is the focus of this episode of care
	Patient common.Reference `json:"patient"`

	// The interval during which the managing organization assumes the defined responsibility
	Period *common.Period `json:"period,omitempty"`

	// The reason communicates what medical problem the patient has that should be addressed during the episode of care
	Reason []EpisodeOfCareReason `json:"reason,omitempty"`

	// Referral Request(s) that are fulfilled by this EpisodeOfCare, incoming referrals
	ReferralRequest []common.Reference `json:"referralRequest,omitempty"`

	// This element is labeled as a modifier because the status contains codes that mark the episode as not currently valid
	Status EpisodeOfCareStatus `json:"status"`

	// The history of statuses that the EpisodeOfCare has been through
	StatusHistory []EpisodeOfCareStatusHistory `json:"statusHistory,omitempty"`

	// The type can be very important in processing as this could be used in determining if the EpisodeOfCare is relevant to specific government reporting
	Type []common.CodeableConcept `json:"type,omitempty"`
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

// TriggerDefinition represents a description of a triggering event
type TriggerDefinition struct {
	common.Element

	// A code that identifies the event
	Code *common.CodeableConcept `json:"code,omitempty"`

	// Additional semantics for the trigger (optional for data type triggers)
	Condition *Expression `json:"condition,omitempty"`

	// This element shall be present for any data type trigger (simplified for now)
	Data []interface{} `json:"data,omitempty"`

	// An event name can be provided for all event types, but is required for named events
	Name *string `json:"name,omitempty"`

	// A reference to a SubscriptionTopic resource that defines the event
	SubscriptionTopic *string `json:"subscriptionTopic,omitempty"`

	// The timing of the event (if this is a periodic trigger) - choice type
	TimingTiming    *Timing           `json:"timingTiming,omitempty"`
	TimingReference *common.Reference `json:"timingReference,omitempty"`
	TimingDate      *string           `json:"timingDate,omitempty"`
	TimingDateTime  *string           `json:"timingDateTime,omitempty"`

	// The type of triggering event
	Type TriggerDefinitionType `json:"type"`
}

// EventDefinitionStatus represents the status of an event definition
type EventDefinitionStatus string

const (
	EventDefinitionStatusDraft   EventDefinitionStatus = "draft"
	EventDefinitionStatusActive  EventDefinitionStatus = "active"
	EventDefinitionStatusRetired EventDefinitionStatus = "retired"
	EventDefinitionStatusUnknown EventDefinitionStatus = "unknown"
)

// EventDefinition provides a reusable description of when a particular event can occur
type EventDefinition struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "EventDefinition"

	// The 'date' element may be more recent than the approval date because of minor changes or editorial corrections
	ApprovalDate *string `json:"approvalDate,omitempty"`

	// An individual or organization primarily involved in the creation and maintenance of the content
	Author []ContactDetail `json:"author,omitempty"`

	// May be a web site, an email address, a telephone number, etc
	Contact []ContactDetail `json:"contact,omitempty"`

	// The short copyright declaration
	Copyright *string `json:"copyright,omitempty"`

	// The (c) symbol should NOT be included in this string
	CopyrightLabel *string `json:"copyrightLabel,omitempty"`

	// The date is often not tracked until the resource is published
	Date *string `json:"date,omitempty"`

	// This description can be used to capture details such as comments about misuse
	Description *string `json:"description,omitempty"`

	// An individual or organization primarily responsible for internal coherence of the content
	Editor []ContactDetail `json:"editor,omitempty"`

	// The effective period for an event definition determines when the content is applicable for usage
	EffectivePeriod *common.Period `json:"effectivePeriod,omitempty"`

	// See guidance around (not) making local changes to elements
	Endorser []ContactDetail `json:"endorser,omitempty"`

	// Allows filtering of event definitions that are appropriate for use versus not
	Experimental *bool `json:"experimental,omitempty"`

	// Typically, this is used for identifiers that can go in an HL7 V3 II (instance identifier) data type
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// It may be possible for the event definition to be used in jurisdictions other than those for which it was originally designed
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`

	// If specified, this date follows the original approval date
	LastReviewDate *string `json:"lastReviewDate,omitempty"`

	// The name is not expected to be globally unique
	Name *string `json:"name,omitempty"`

	// The publisher (or steward) of the event definition
	Publisher *string `json:"publisher,omitempty"`

	// This element does not describe the usage of the event definition
	Purpose *string `json:"purpose,omitempty"`

	// Each related resource is either an attachment, or a reference to another resource, but not both
	RelatedArtifact []RelatedArtifact `json:"relatedArtifact,omitempty"`

	// See guidance around (not) making local changes to elements
	Reviewer []ContactDetail `json:"reviewer,omitempty"`

	// Allows filtering of event definitions that are appropriate for use versus not
	Status EventDefinitionStatus `json:"status"`

	// A code or group definition that describes the intended subject of the event definition - choice type
	SubjectCodeableConcept *common.CodeableConcept `json:"subjectCodeableConcept,omitempty"`
	SubjectReference       *common.Reference       `json:"subjectReference,omitempty"`

	// An explanatory or alternate title for the event definition
	Subtitle *string `json:"subtitle,omitempty"`

	// This name does not need to be machine-processing friendly
	Title *string `json:"title,omitempty"`

	// DEPRECATION NOTE: For consistency, implementations are encouraged to migrate to using the new 'topic' code in the useContext element
	Topic []common.CodeableConcept `json:"topic,omitempty"`

	// The trigger element defines when the event occurs
	Trigger []TriggerDefinition `json:"trigger"`

	// Can be a urn:uuid: or a urn:oid: but real http: addresses are preferred
	Url *string `json:"url,omitempty"`

	// A detailed description of how the event definition is used from a clinical perspective
	Usage *string `json:"usage,omitempty"`

	// When multiple useContexts are specified, there is no expectation that all or any of the contexts apply
	UseContext []UsageContext `json:"useContext,omitempty"`

	// There may be different event definition instances that have the same identifier but different versions
	Version *string `json:"version,omitempty"`

	// If set as a string, this is a FHIRPath expression
	VersionAlgorithmString *string        `json:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *common.Coding `json:"versionAlgorithmCoding,omitempty"`
}

// EvidenceReportSubjectCharacteristic represents a characteristic of the subject
type EvidenceReportSubjectCharacteristic struct {
	common.BackboneElement

	// Example 1 is a Citation. Example 2 is a type of outcome. Example 3 is a specific outcome
	Code common.CodeableConcept `json:"code"`

	// Is used to express not the characteristic
	Exclude *bool `json:"exclude,omitempty"`

	// Timeframe for the characteristic
	Period *common.Period `json:"period,omitempty"`

	// Example 1 is Citation #37. Example 2 is selecting clinical outcomes. Example 3 is 1-year mortality
	ValueReference       *common.Reference       `json:"valueReference,omitempty"`
	ValueCodeableConcept *common.CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueBoolean         *bool                   `json:"valueBoolean,omitempty"`
	ValueQuantity        *common.Quantity        `json:"valueQuantity,omitempty"`
	ValueRange           *Range                  `json:"valueRange,omitempty"`
}

// EvidenceReportSubject represents the subject of the evidence report
type EvidenceReportSubject struct {
	common.BackboneElement

	// Characteristic
	Characteristic []EvidenceReportSubjectCharacteristic `json:"characteristic,omitempty"`

	// Used for general notes and annotations not coded elsewhere
	Note []Annotation `json:"note,omitempty"`
}

// EvidenceReportRelatesToTarget represents the target of a relationship
type EvidenceReportRelatesToTarget struct {
	common.BackboneElement

	// Target of the relationship Display
	Display *string `json:"display,omitempty"`

	// Target of the relationship Identifier
	Identifier *common.Identifier `json:"identifier,omitempty"`

	// Target of the relationship Resource reference
	Resource *common.Reference `json:"resource,omitempty"`

	// Target of the relationship URL
	URL *string `json:"url,omitempty"`
}

// EvidenceReportRelatesTo represents relationships to other documents
type EvidenceReportRelatesTo struct {
	common.BackboneElement

	// replaces | amends | appends | transforms | replacedWith | amendedWith | appendedWith | transformedWith
	Code EvidenceReportRelatesToCode `json:"code"`

	// The target composition/document of this relationship
	Target EvidenceReportRelatesToTarget `json:"target"`
}

// EvidenceReportRelatesToCode represents the relationship type
type EvidenceReportRelatesToCode string

const (
	EvidenceReportRelatesToCodeReplaces        EvidenceReportRelatesToCode = "replaces"
	EvidenceReportRelatesToCodeAmends          EvidenceReportRelatesToCode = "amends"
	EvidenceReportRelatesToCodeAppends         EvidenceReportRelatesToCode = "appends"
	EvidenceReportRelatesToCodeTransforms      EvidenceReportRelatesToCode = "transforms"
	EvidenceReportRelatesToCodeReplacedWith    EvidenceReportRelatesToCode = "replacedWith"
	EvidenceReportRelatesToCodeAmendedWith     EvidenceReportRelatesToCode = "amendedWith"
	EvidenceReportRelatesToCodeAppendedWith    EvidenceReportRelatesToCode = "appendedWith"
	EvidenceReportRelatesToCodeTransformedWith EvidenceReportRelatesToCode = "transformedWith"
)

// EvidenceReportSection represents a section of the evidence report
type EvidenceReportSection struct {
	common.BackboneElement

	// Identifies who is responsible for the information in this section
	Author []common.Reference `json:"author,omitempty"`

	// The various reasons for an empty section
	EmptyReason *common.CodeableConcept `json:"emptyReason,omitempty"`

	// Specifies any type of classification of the evidence report
	EntryClassifier []common.CodeableConcept `json:"entryClassifier,omitempty"`

	// Quantity as content
	EntryQuantity []common.Quantity `json:"entryQuantity,omitempty"`

	// If there are no entries in the list, an emptyReason SHOULD be provided
	EntryReference []common.Reference `json:"entryReference,omitempty"`

	// The code identifies the section for an automated processor
	Focus *common.CodeableConcept `json:"focus,omitempty"`

	// A definitional Resource identifying the kind of content contained within the section
	FocusReference *common.Reference `json:"focusReference,omitempty"`

	// working | snapshot | changes
	Mode *EvidenceReportSectionMode `json:"mode,omitempty"`

	// Applications SHOULD render ordered lists in the order provided
	OrderedBy *common.CodeableConcept `json:"orderedBy,omitempty"`

	// Nested sections are primarily used to help human readers navigate
	Section []EvidenceReportSection `json:"section,omitempty"`

	// Document profiles may define what content should be represented in the narrative
	Text *Narrative `json:"text,omitempty"`

	// The title identifies the section for a human reader
	Title *string `json:"title,omitempty"`
}

// EvidenceReportSectionMode represents the mode of a section
type EvidenceReportSectionMode string

const (
	EvidenceReportSectionModeWorking  EvidenceReportSectionMode = "working"
	EvidenceReportSectionModeSnapshot EvidenceReportSectionMode = "snapshot"
	EvidenceReportSectionModeChanges  EvidenceReportSectionMode = "changes"
)

// EvidenceReport represents a specialized container for collections of evidence
type EvidenceReport struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "EvidenceReport"

	// Extensions to ContactDetail include: contactReference, contactAddress, and contributionTime
	Author []ContactDetail `json:"author,omitempty"`

	// Used for reports for which external citation is expected
	CiteAsReference *common.Reference `json:"citeAsReference,omitempty"`
	CiteAsMarkdown  *string           `json:"citeAsMarkdown,omitempty"`

	// Extensions to ContactDetail include: contactReference, contactAddress, and contributionTime
	Contact []ContactDetail `json:"contact,omitempty"`

	// Extensions to ContactDetail include: contactReference, contactAddress, and contributionTime
	Editor []ContactDetail `json:"editor,omitempty"`

	// Extensions to ContactDetail include: contactReference, contactAddress, and contributionTime
	Endorser []ContactDetail `json:"endorser,omitempty"`

	// This element will contain unique identifiers that support de-duplication of EvidenceReports
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Used for footnotes and annotations
	Note []Annotation `json:"note,omitempty"`

	// Usually an organization but may be an individual
	Publisher *string `json:"publisher,omitempty"`

	// Link, description or reference to artifact associated with the report
	RelatedArtifact []RelatedArtifact `json:"relatedArtifact,omitempty"`

	// May include trial registry identifiers
	RelatedIdentifier []common.Identifier `json:"relatedIdentifier,omitempty"`

	// A document is a version specific composition
	RelatesTo []EvidenceReportRelatesTo `json:"relatesTo,omitempty"`

	// Extensions to ContactDetail include: contactReference, contactAddress, and contributionTime
	Reviewer []ContactDetail `json:"reviewer,omitempty"`

	// The root of the sections that make up the composition
	Section []EvidenceReportSection `json:"section,omitempty"`

	// Allows filtering of summaries that are appropriate for use versus not
	Status EventDefinitionStatus `json:"status"`

	// May be used as an expression for search queries and search results
	Subject EvidenceReportSubject `json:"subject"`

	// Specifies the kind of report
	Type *common.CodeableConcept `json:"type,omitempty"`

	// Can be a urn:uuid: or a urn:oid: but real http: addresses are preferred
	URL *string `json:"url,omitempty"`

	// When multiple useContexts are specified, there is no expectation that all or any of the contexts apply
	UseContext []UsageContext `json:"useContext,omitempty"`
}

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

// Evidence represents machine-interpretable expression of evidence
type Evidence struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Evidence"

	// The 'date' element may be more recent than the approval date
	ApprovalDate *string `json:"approvalDate,omitempty"`

	// Declarative description of the Evidence
	Assertion *string `json:"assertion,omitempty"`

	// An individual or organization primarily involved in the creation and maintenance of the content
	Author []ContactDetail `json:"author,omitempty"`

	// Assessment of certainty, confidence in the estimates, or quality of the evidence
	Certainty []EvidenceCertainty `json:"certainty,omitempty"`

	// Citation Resource or display of suggested citation for this evidence
	CiteAsReference *common.Reference `json:"citeAsReference,omitempty"`
	CiteAsMarkdown  *string           `json:"citeAsMarkdown,omitempty"`

	// May be a web site, an email address, a telephone number, etc
	Contact []ContactDetail `json:"contact,omitempty"`

	// Use and/or publishing restrictions
	Copyright *string `json:"copyright,omitempty"`

	// Copyright holder and year(s)
	CopyrightLabel *string `json:"copyrightLabel,omitempty"`

	// Date last changed
	Date *string `json:"date,omitempty"`

	// Natural language description of the evidence
	Description *string `json:"description,omitempty"`

	// An individual or organization primarily responsible for internal coherence of the content
	Editor []ContactDetail `json:"editor,omitempty"`

	// The effective period for an evidence determines when the content is applicable for usage
	EffectivePeriod *common.Period `json:"effectivePeriod,omitempty"`

	// See guidance around (not) making local changes to elements
	Endorser []ContactDetail `json:"endorser,omitempty"`

	// For testing purposes, not real usage
	Experimental *bool `json:"experimental,omitempty"`

	// Additional identifier for the evidence
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// If specified, this date follows the original approval date
	LastReviewDate *string `json:"lastReviewDate,omitempty"`

	// Name for this evidence (computer friendly)
	Name *string `json:"name,omitempty"`

	// Used for footnotes or explanatory notes
	Note []Annotation `json:"note,omitempty"`

	// Name of the publisher (organization or individual)
	Publisher *string `json:"publisher,omitempty"`

	// Why this evidence is defined
	Purpose *string `json:"purpose,omitempty"`

	// Link or citation to artifact associated with the summary
	RelatedArtifact []RelatedArtifact `json:"relatedArtifact,omitempty"`

	// See guidance around (not) making local changes to elements
	Reviewer []ContactDetail `json:"reviewer,omitempty"`

	// Values and parameters for a single statistic
	Statistic []EvidenceStatistic `json:"statistic,omitempty"`

	// The status of this evidence
	Status EventDefinitionStatus `json:"status"`

	// The design of the study that produced this evidence
	StudyDesign []common.CodeableConcept `json:"studyDesign,omitempty"`

	// The method to combine studies
	SynthesisType *common.CodeableConcept `json:"synthesisType,omitempty"`

	// Name for this evidence (human friendly)
	Title *string `json:"title,omitempty"`

	// Canonical identifier for this evidence
	URL *string `json:"url,omitempty"`

	// The context that the content is intended to support
	UseContext []UsageContext `json:"useContext,omitempty"`

	// Evidence variable such as population, exposure, or outcome
	VariableDefinition []EvidenceVariableDefinition `json:"variableDefinition"`

	// Business version of the evidence
	Version *string `json:"version,omitempty"`

	// If set as a string, this is a FHIRPath expression
	VersionAlgorithmString *string        `json:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *common.Coding `json:"versionAlgorithmCoding,omitempty"`
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

// EvidenceVariable represents the EvidenceVariable resource
type EvidenceVariable struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "EvidenceVariable"

	// True if the actual variable measured, false if a conceptual representation
	Actual *bool `json:"actual,omitempty"`

	// The date on which the resource content was approved by the publisher
	ApprovalDate *string `json:"approvalDate,omitempty"`

	// Contact details for the publisher
	Author []ContactDetail `json:"author,omitempty"`

	// A grouping for ordinal or polychotomous variables
	Category []EvidenceVariableCategory `json:"category,omitempty"`

	// Characteristics can be defined flexibly to accommodate different use cases
	Characteristic []EvidenceVariableCharacteristic `json:"characteristic,omitempty"`

	// Contact details for the publisher
	Contact []ContactDetail `json:"contact,omitempty"`

	// A copyright statement relating to the resource
	Copyright *string `json:"copyright,omitempty"`

	// Copyright label for the resource
	CopyrightLabel *string `json:"copyrightLabel,omitempty"`

	// The date on which the resource content was created
	Date *string `json:"date,omitempty"`

	// This description can be used to capture details such as comments about misuse
	Description *string `json:"description,omitempty"`

	// Contact details for the publisher
	Editor []ContactDetail `json:"editor,omitempty"`

	// The effective period for an EvidenceVariable
	EffectivePeriod *common.Period `json:"effectivePeriod,omitempty"`

	// Contact details for the publisher
	Endorser []ContactDetail `json:"endorser,omitempty"`

	// A Boolean value to indicate that this resource is authored for testing purposes
	Experimental *bool `json:"experimental,omitempty"`

	// The method of handling in statistical analysis
	Handling *EvidenceVariableHandling `json:"handling,omitempty"`

	// Typically, this is used for identifiers that can go in an HL7 V3 II data type
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The intended jurisdiction for the resource
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`

	// If specified, this date follows the original approval date
	LastReviewDate *string `json:"lastReviewDate,omitempty"`

	// The name is not expected to be globally unique
	Name *string `json:"name,omitempty"`

	// A human-readable string to clarify or explain concepts about the resource
	Note []Annotation `json:"note,omitempty"`

	// The publisher of the resource
	Publisher *string `json:"publisher,omitempty"`

	// This element does not describe the usage of the EvidenceVariable
	Purpose *string `json:"purpose,omitempty"`

	// Each related artifact is either an attachment, or a reference to another resource
	RelatedArtifact []RelatedArtifact `json:"relatedArtifact,omitempty"`

	// Contact details for the publisher
	Reviewer []ContactDetail `json:"reviewer,omitempty"`

	// The short title provides an alternate title for use in informal descriptive contexts
	ShortTitle *string `json:"shortTitle,omitempty"`

	// Allows filtering of evidence variables that are appropriate for use versus not
	Status EventDefinitionStatus `json:"status"`

	// An explanatory or alternate title for the EvidenceVariable
	Subtitle *string `json:"subtitle,omitempty"`

	// This name does not need to be machine-processing friendly
	Title *string `json:"title,omitempty"`

	// In some cases, the resource can no longer be found at the stated url
	URL *string `json:"url,omitempty"`

	// When multiple useContexts are specified, there is no expectation that all apply
	UseContext []UsageContext `json:"useContext,omitempty"`

	// The version of the EvidenceVariable
	Version *string `json:"version,omitempty"`

	// The version algorithm
	VersionAlgorithmString *string        `json:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *common.Coding `json:"versionAlgorithmCoding,omitempty"`
}

// ExampleScenarioActorType represents the type of actor
type ExampleScenarioActorType string

const (
	ExampleScenarioActorTypePerson ExampleScenarioActorType = "person"
	ExampleScenarioActorTypeSystem ExampleScenarioActorType = "system"
)

// ExampleScenarioActor represents a system or person who shares or receives an instance within the scenario
type ExampleScenarioActor struct {
	common.BackboneElement

	// An explanation of who/what the actor is and its role in the scenario
	Description *string `json:"description,omitempty"`

	// A unique string within the scenario that is used to reference the actor
	Key string `json:"key"`

	// The human-readable name for the actor used when rendering the scenario
	Title string `json:"title"`

	// The category of actor - person or system
	Type ExampleScenarioActorType `json:"type"`
}

// ExampleScenarioInstanceContainedInstance represents references to other instances
type ExampleScenarioInstanceContainedInstance struct {
	common.BackboneElement

	// A unique string within the instance that is used to reference the instance
	InstanceReference string `json:"instanceReference"`

	// A human-readable description of the instance
	VersionReference *string `json:"versionReference,omitempty"`
}

// ExampleScenarioInstanceVersion represents a specific version of an instance
type ExampleScenarioInstanceVersion struct {
	common.BackboneElement

	// If not conveying FHIR data or not using the same version of FHIR
	Content *common.Reference `json:"content,omitempty"`

	// An explanation of what this specific version of the instance contains
	Description *string `json:"description,omitempty"`

	// A unique string within the instance that is used to reference the version
	Key string `json:"key"`

	// A short descriptive label the version to be used in tables or diagrams
	Title string `json:"title"`
}

// ExampleScenarioInstance represents a single data collection that is shared as part of the scenario
type ExampleScenarioInstance struct {
	common.BackboneElement

	// References to other instances that can be found within this instance
	ContainedInstance []ExampleScenarioInstanceContainedInstance `json:"containedInstance,omitempty"`

	// If not conveying FHIR data or not using the same version of FHIR
	Content *common.Reference `json:"content,omitempty"`

	// An explanation of what the instance contains and what it's for
	Description *string `json:"description,omitempty"`

	// A unique string within the scenario that is used to reference the instance
	Key string `json:"key"`

	// Refers to a profile, template or other ruleset the instance adheres to
	StructureProfileCanonical *string `json:"structureProfileCanonical,omitempty"`
	StructureProfileUri       *string `json:"structureProfileUri,omitempty"`

	// A code indicating the kind of data structure
	StructureType common.Coding `json:"structureType"`

	// Conveys the version of the data structure instantiated
	StructureVersion *string `json:"structureVersion,omitempty"`

	// A short descriptive label the instance to be used in tables or diagrams
	Title string `json:"title"`

	// Not used if an instance doesn't change
	Version []ExampleScenarioInstanceVersion `json:"version,omitempty"`
}

// ExampleScenarioProcessStepOperation represents the step that represents a single operation invoked on receiver by sender
type ExampleScenarioProcessStepOperation struct {
	common.BackboneElement

	// This should contain information not already present in the process step
	Description *string `json:"description,omitempty"`

	// This must either be the 'key' of one of the actors defined in this scenario or the special keyword 'OTHER'
	Initiator *string `json:"initiator,omitempty"`

	// De-activation of an actor means they have no further role
	InitiatorActive *bool `json:"initiatorActive,omitempty"`

	// This must either be the 'key' of one of the actors defined in this scenario or the special keyword 'OTHER'
	Receiver *string `json:"receiver,omitempty"`

	// De-activation of an actor means they have no further role
	ReceiverActive *bool `json:"receiverActive,omitempty"`

	// A reference to the instance that is transmitted from requester to receiver
	Request *ExampleScenarioInstanceContainedInstance `json:"request,omitempty"`

	// A reference to the instance that is transmitted from receiver to requester
	Response *ExampleScenarioInstanceContainedInstance `json:"response,omitempty"`

	// The type of operation being performed
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// ExampleScenarioProcessStepAlternative represents an alternative step
type ExampleScenarioProcessStepAlternative struct {
	common.BackboneElement

	// A human-readable description of the alternative explaining when the alternative should occur
	Description *string `json:"description,omitempty"`

	// What happens in each alternative option
	Step []ExampleScenarioProcessStep `json:"step,omitempty"`

	// The label to display for the alternative
	Title string `json:"title"`
}

// ExampleScenarioProcessStep represents a significant action that occurs as part of the process
type ExampleScenarioProcessStep struct {
	common.BackboneElement

	// Indicates an alternative step that can be taken instead of the sub-process
	Alternative []ExampleScenarioProcessStepAlternative `json:"alternative,omitempty"`

	// If step numbers are simultaneous, they will be the same
	Number *string `json:"number,omitempty"`

	// The step represents a single operation invoked on receiver by sender
	Operation *ExampleScenarioProcessStepOperation `json:"operation,omitempty"`

	// If true, indicates that, following this step, there is a pause in the flow
	Pause *bool `json:"pause,omitempty"`

	// Indicates that the step is a complex sub-process with its own steps
	Process *ExampleScenarioProcess `json:"process,omitempty"`

	// Indicates that the step is defined by a separate scenario instance
	Workflow *string `json:"workflow,omitempty"`
}

// ExampleScenarioProcess represents some scenarios might describe only one process
type ExampleScenarioProcess struct {
	common.BackboneElement

	// An explanation of what the process represents and what it does
	Description *string `json:"description,omitempty"`

	// Alternate steps might not result in all post conditions holding
	PostConditions *string `json:"postConditions,omitempty"`

	// Description of the initial state of the actors, environment and data
	PreConditions *string `json:"preConditions,omitempty"`

	// A significant action that occurs as part of the process
	Step []ExampleScenarioProcessStep `json:"step,omitempty"`

	// A short descriptive label the process to be used in tables or diagrams
	Title string `json:"title"`
}

// ExampleScenario represents an example of workflow instance
type ExampleScenario struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "ExampleScenario"

	// A system or person who shares or receives an instance within the scenario
	Actor []ExampleScenarioActor `json:"actor,omitempty"`

	// May be a web site, an email address, a telephone number, etc
	Contact []ContactDetail `json:"contact,omitempty"`

	// Use and/or publishing restrictions
	Copyright *string `json:"copyright,omitempty"`

	// Copyright holder and year(s)
	CopyrightLabel *string `json:"copyrightLabel,omitempty"`

	// Date last changed
	Date *string `json:"date,omitempty"`

	// Natural language description of the example scenario
	Description *string `json:"description,omitempty"`

	// For testing purposes, not real usage
	Experimental *bool `json:"experimental,omitempty"`

	// Additional identifier for the example scenario
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// A single data collection that is shared as part of the scenario
	Instance []ExampleScenarioInstance `json:"instance,omitempty"`

	// Intended jurisdiction for example scenario
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`

	// Name for this example scenario (computer friendly)
	Name *string `json:"name,omitempty"`

	// Some scenarios might describe only one process
	Process []ExampleScenarioProcess `json:"process,omitempty"`

	// Name of the publisher (organization or individual)
	Publisher *string `json:"publisher,omitempty"`

	// Why this example scenario is defined
	Purpose *string `json:"purpose,omitempty"`

	// The status of this example scenario
	Status EventDefinitionStatus `json:"status"`

	// Name for this example scenario (human friendly)
	Title *string `json:"title,omitempty"`

	// Canonical identifier for this example scenario
	URL *string `json:"url,omitempty"`

	// The context that the content is intended to support
	UseContext []UsageContext `json:"useContext,omitempty"`

	// Business version of the example scenario
	Version *string `json:"version,omitempty"`

	// How to compare versions
	VersionAlgorithmString *string        `json:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *common.Coding `json:"versionAlgorithmCoding,omitempty"`
}

// ExplanationOfBenefitStatus represents the status of an explanation of benefit
type ExplanationOfBenefitStatus string

const (
	ExplanationOfBenefitStatusActive         ExplanationOfBenefitStatus = "active"
	ExplanationOfBenefitStatusCancelled      ExplanationOfBenefitStatus = "cancelled"
	ExplanationOfBenefitStatusDraft          ExplanationOfBenefitStatus = "draft"
	ExplanationOfBenefitStatusEnteredInError ExplanationOfBenefitStatus = "entered-in-error"
)

// ExplanationOfBenefitOutcome represents the outcome of the claim processing
type ExplanationOfBenefitOutcome string

const (
	ExplanationOfBenefitOutcomeQueued   ExplanationOfBenefitOutcome = "queued"
	ExplanationOfBenefitOutcomeComplete ExplanationOfBenefitOutcome = "complete"
	ExplanationOfBenefitOutcomeError    ExplanationOfBenefitOutcome = "error"
	ExplanationOfBenefitOutcomePartial  ExplanationOfBenefitOutcome = "partial"
)

// ExplanationOfBenefitUse represents the use of the explanation of benefit
type ExplanationOfBenefitUse string

const (
	ExplanationOfBenefitUseClaim            ExplanationOfBenefitUse = "claim"
	ExplanationOfBenefitUsePreauthorization ExplanationOfBenefitUse = "preauthorization"
	ExplanationOfBenefitUsePredetermination ExplanationOfBenefitUse = "predetermination"
)

// ExplanationOfBenefitEvent represents information code for an event
type ExplanationOfBenefitEvent struct {
	common.BackboneElement

	// A coded event such as when a service is expected or a card printed
	Type common.CodeableConcept `json:"type"`

	// A date or period in the past or future indicating when the event occurred
	WhenDateTime *string        `json:"whenDateTime,omitempty"`
	WhenPeriod   *common.Period `json:"whenPeriod,omitempty"`
}

// ExplanationOfBenefitSupportingInfo represents supporting information for the claim
type ExplanationOfBenefitSupportingInfo struct {
	common.BackboneElement

	// This may contain a category for the local bill type codes
	Category common.CodeableConcept `json:"category"`

	// This may contain the local bill type codes
	Code *common.CodeableConcept `json:"code,omitempty"`

	// For example: the reason for the additional stay, or why a tooth is missing
	Reason *common.Coding `json:"reason,omitempty"`

	// A number to uniquely identify supporting information entries
	Sequence int `json:"sequence"`

	// The date when or period to which this information refers
	TimingDate   *string        `json:"timingDate,omitempty"`
	TimingPeriod *common.Period `json:"timingPeriod,omitempty"`

	// Could be used to provide references to other resources, document
	ValueBoolean    *bool              `json:"valueBoolean,omitempty"`
	ValueString     *string            `json:"valueString,omitempty"`
	ValueQuantity   *common.Quantity   `json:"valueQuantity,omitempty"`
	ValueAttachment *Attachment        `json:"valueAttachment,omitempty"`
	ValueReference  *common.Reference  `json:"valueReference,omitempty"`
	ValueIdentifier *common.Identifier `json:"valueIdentifier,omitempty"`
}

// ExplanationOfBenefitDiagnosis represents information about diagnoses relevant to the claim
type ExplanationOfBenefitDiagnosis struct {
	common.BackboneElement

	// The nature of illness or problem in a coded form or as a reference to an external defined Condition
	DiagnosisCodeableConcept *common.CodeableConcept `json:"diagnosisCodeableConcept,omitempty"`
	DiagnosisReference       *common.Reference       `json:"diagnosisReference,omitempty"`

	// Indication of whether the diagnosis was present on admission to a facility
	OnAdmission *common.CodeableConcept `json:"onAdmission,omitempty"`

	// Diagnosis are presented in list order to their expected importance
	Sequence int `json:"sequence"`

	// For example: admitting, primary, secondary, discharge
	Type []common.CodeableConcept `json:"type,omitempty"`
}

// ExplanationOfBenefitProcedure represents procedures performed on the patient
type ExplanationOfBenefitProcedure struct {
	common.BackboneElement

	// Date and optionally time the procedure was performed
	Date *string `json:"date,omitempty"`

	// Specific clinical procedure
	ProcedureCodeableConcept *common.CodeableConcept `json:"procedureCodeableConcept,omitempty"`
	ProcedureReference       *common.Reference       `json:"procedureReference,omitempty"`

	// A number to uniquely identify procedure entries
	Sequence int `json:"sequence"`

	// Category of Procedure
	Type []common.CodeableConcept `json:"type,omitempty"`

	// Unique device identifier
	Udi []common.Reference `json:"udi,omitempty"`
}

// ExplanationOfBenefitInsurance represents patient insurance information
type ExplanationOfBenefitInsurance struct {
	common.BackboneElement

	// Insurance information
	Coverage common.Reference `json:"coverage"`

	// Coverage to be used for adjudication
	Focal bool `json:"focal"`

	// Prior authorization reference number
	PreAuthRef []string `json:"preAuthRef,omitempty"`
}

// ExplanationOfBenefitAccident represents details of an accident
type ExplanationOfBenefitAccident struct {
	common.BackboneElement

	// When the incident occurred
	Date *string `json:"date,omitempty"`

	// Where the event occurred
	LocationAddress   *Address          `json:"locationAddress,omitempty"`
	LocationReference *common.Reference `json:"locationReference,omitempty"`

	// The nature of the accident
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// ExplanationOfBenefitCareTeam represents the members of the team who provided the products and services
type ExplanationOfBenefitCareTeam struct {
	common.BackboneElement

	// Practitioner or organization
	Provider common.Reference `json:"provider"`

	// Practitioner credential or specialization
	Qualification *common.CodeableConcept `json:"qualification,omitempty"`

	// Indicator of the lead practitioner
	Responsible *bool `json:"responsible,omitempty"`

	// Function within the team
	Role *common.CodeableConcept `json:"role,omitempty"`

	// Order of care team
	Sequence int `json:"sequence"`
}

// ExplanationOfBenefitPayee represents the recipient of the payment
type ExplanationOfBenefitPayee struct {
	common.BackboneElement

	// Recipient reference
	Party *common.Reference `json:"party,omitempty"`

	// Category of recipient
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// ExplanationOfBenefitPayment represents payment details for the adjudication
type ExplanationOfBenefitPayment struct {
	common.BackboneElement

	// Partial or complete payment adjustment
	Adjustment *Money `json:"adjustment,omitempty"`

	// Explanation for the adjustment
	AdjustmentReason *common.CodeableConcept `json:"adjustmentReason,omitempty"`

	// Payable amount after adjustment
	Amount *Money `json:"amount,omitempty"`

	// Expected date of payment
	Date *string `json:"date,omitempty"`

	// Business identifier for the payment
	Identifier *common.Identifier `json:"identifier,omitempty"`

	// Payment instrument
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// ExplanationOfBenefitProcessNote represents a note that describes or explains adjudication results
type ExplanationOfBenefitProcessNote struct {
	common.BackboneElement

	// Language of the text
	Language *common.CodeableConcept `json:"language,omitempty"`

	// A number to uniquely identify a note entry
	Number *int `json:"number,omitempty"`

	// Note explanatory text
	Text *string `json:"text,omitempty"`

	// The business purpose of the note text
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// ExplanationOfBenefitRelated represents other claims which are related to this claim
type ExplanationOfBenefitRelated struct {
	common.BackboneElement

	// File or case reference
	Claim *common.Reference `json:"claim,omitempty"`

	// Identifier of the related claim
	Reference *common.Identifier `json:"reference,omitempty"`

	// How the reference claim is related
	Relationship *common.CodeableConcept `json:"relationship,omitempty"`
}

// ExplanationOfBenefitTotal represents totals for amounts submitted, co-pays, benefits payable etc
type ExplanationOfBenefitTotal struct {
	common.BackboneElement

	// Financial total for the category
	Amount Money `json:"amount"`

	// Type of adjudication information
	Category common.CodeableConcept `json:"category"`
}

// ExplanationOfBenefitBenefitBalance represents balance by benefit category
type ExplanationOfBenefitBenefitBalance struct {
	common.BackboneElement

	// Examples include Medical Care, Periodontics, Renal Dialysis, Vision Coverage
	Category common.CodeableConcept `json:"category"`

	// For example, 'DENT2 covers 100% of basic, 50% of major but excludes Ortho, Implants and Cosmetic services'
	Description *string `json:"description,omitempty"`

	// True if the indicated class of service is excluded from the plan
	Excluded *bool `json:"excluded,omitempty"`

	// For example: MED01, or DENT2
	Name *string `json:"name,omitempty"`

	// Is a flag to indicate whether the benefits refer to in-network providers
	Network *common.CodeableConcept `json:"network,omitempty"`

	// The term or period of the values such as 'maximum lifetime benefit'
	Term *common.CodeableConcept `json:"term,omitempty"`

	// Indicates if the benefits apply to an individual or to the family
	Unit *common.CodeableConcept `json:"unit,omitempty"`
}

// ExplanationOfBenefitItem represents a claim line
type ExplanationOfBenefitItem struct {
	common.BackboneElement

	// Care team members related to this service or product
	CareTeamSequence []int `json:"careTeamSequence,omitempty"`

	// Examples include Medical Care, Periodontics, Renal Dialysis, Vision Coverage
	Category *common.CodeableConcept `json:"category,omitempty"`

	// Diagnoses applicable for this service or product
	DiagnosisSequence []int `json:"diagnosisSequence,omitempty"`

	// Healthcare encounters related to this claim
	Encounter []common.Reference `json:"encounter,omitempty"`

	// To show a 10% senior's discount, the value entered is: 0.90 (1.00 - 0.10)
	Factor *float64 `json:"factor,omitempty"`

	// Exceptions, special conditions and supporting information
	InformationSequence []int `json:"informationSequence,omitempty"`

	// Where the product or service was provided
	LocationCodeableConcept *common.CodeableConcept `json:"locationCodeableConcept,omitempty"`
	LocationAddress         *Address                `json:"locationAddress,omitempty"`
	LocationReference       *common.Reference       `json:"locationReference,omitempty"`

	// For example, in Oral whether the treatment is cosmetic or associated with TMJ
	Modifier []common.CodeableConcept `json:"modifier,omitempty"`

	// When the value is a group code then this item collects a set of related claim details
	ProductOrService *common.CodeableConcept `json:"productOrService,omitempty"`

	// Item typification or modifiers codes
	ProgramCode []common.CodeableConcept `json:"programCode,omitempty"`

	// Count of products or services
	Quantity *common.Quantity `json:"quantity,omitempty"`

	// Applicable exception and supporting information
	Revenue *common.CodeableConcept `json:"revenue,omitempty"`

	// A number to uniquely identify item entries
	Sequence int `json:"sequence"`

	// Date or dates of service or product delivery
	ServicedDate   *string        `json:"servicedDate,omitempty"`
	ServicedPeriod *common.Period `json:"servicedPeriod,omitempty"`

	// Anatomical sub-location
	SubSite []common.CodeableConcept `json:"subSite,omitempty"`

	// Unique device identifier
	Udi []common.Reference `json:"udi,omitempty"`

	// Fee, charge or cost per item
	UnitPrice *Money `json:"unitPrice,omitempty"`
}

// ExplanationOfBenefit represents the explanation of benefit resource
type ExplanationOfBenefit struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "ExplanationOfBenefit"

	// Details of a accident which resulted in injuries
	Accident *ExplanationOfBenefitAccident `json:"accident,omitempty"`

	// Balance by Benefit Category
	BenefitBalance []ExplanationOfBenefitBenefitBalance `json:"benefitBalance,omitempty"`

	// Not applicable when use=claim
	BenefitPeriod *common.Period `json:"benefitPeriod,omitempty"`

	// Typically this would be today or in the past for a claim
	BillablePeriod *common.Period `json:"billablePeriod,omitempty"`

	// The members of the team who provided the products and services
	CareTeam []ExplanationOfBenefitCareTeam `json:"careTeam,omitempty"`

	// The business identifier for the instance of the adjudication request
	Claim *common.Reference `json:"claim,omitempty"`

	// The business identifier for the instance of the adjudication response
	ClaimResponse *common.Reference `json:"claimResponse,omitempty"`

	// The date this resource was created
	Created string `json:"created"`

	// The element is used to indicate the current state of the adjudication
	Decision *common.CodeableConcept `json:"decision,omitempty"`

	// Information about diagnoses relevant to the claim items
	Diagnosis []ExplanationOfBenefitDiagnosis `json:"diagnosis,omitempty"`

	// For example DRG (Diagnosis Related Group) or a bundled billing code
	DiagnosisRelatedGroup *common.CodeableConcept `json:"diagnosisRelatedGroup,omitempty"`

	// A human readable description of the status of the adjudication
	Disposition *string `json:"disposition,omitempty"`

	// Healthcare encounters related to this claim
	Encounter []common.Reference `json:"encounter,omitempty"`

	// Individual who created the claim, predetermination or preauthorization
	Enterer *common.Reference `json:"enterer,omitempty"`

	// Information code for an event with a corresponding date or period
	Event []ExplanationOfBenefitEvent `json:"event,omitempty"`

	// Facility where the services were provided
	Facility *common.Reference `json:"facility,omitempty"`

	// Needed to permit insurers to include the actual form
	Form *Attachment `json:"form,omitempty"`

	// May be needed to identify specific jurisdictional forms
	FormCode *common.CodeableConcept `json:"formCode,omitempty"`

	// Fund would be release by a future claim
	FundsReserve *common.CodeableConcept `json:"fundsReserve,omitempty"`

	// This field is only used for preauthorizations
	FundsReserveRequested *common.CodeableConcept `json:"fundsReserveRequested,omitempty"`

	// A unique identifier assigned to this explanation of benefit
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// All insurance coverages for the patient
	Insurance []ExplanationOfBenefitInsurance `json:"insurance,omitempty"`

	// The party responsible for authorization, adjudication and reimbursement
	Insurer *common.Reference `json:"insurer,omitempty"`

	// A claim line
	Item []ExplanationOfBenefitItem `json:"item,omitempty"`

	// For example, a physician may prescribe a medication
	OriginalPrescription *common.Reference `json:"originalPrescription,omitempty"`

	// The resource may be used to indicate that the Claim has been received
	Outcome ExplanationOfBenefitOutcome `json:"outcome"`

	// The party to whom the professional services have been supplied
	Patient common.Reference `json:"patient"`

	// The amount paid by the patient
	PatientPaid *Money `json:"patientPaid,omitempty"`

	// Often providers agree to receive the benefits payable
	Payee *ExplanationOfBenefitPayee `json:"payee,omitempty"`

	// Payment details for the adjudication of the claim
	Payment *ExplanationOfBenefitPayment `json:"payment,omitempty"`

	// This value is only present on preauthorization adjudications
	PreAuthRef []string `json:"preAuthRef,omitempty"`

	// This value is only present on preauthorization adjudications
	PreAuthRefPeriod []common.Period `json:"preAuthRefPeriod,omitempty"`

	// This indicates the relative order of a series of EOBs
	Precedence *int `json:"precedence,omitempty"`

	// Prescription is the document/authorization given to the claim author
	Prescription *common.Reference `json:"prescription,omitempty"`

	// If a claim processor is unable to complete the processing
	Priority *common.CodeableConcept `json:"priority,omitempty"`

	// Procedures performed on the patient relevant to the billing items
	Procedure []ExplanationOfBenefitProcedure `json:"procedure,omitempty"`

	// A note that describes or explains adjudication results
	ProcessNote []ExplanationOfBenefitProcessNote `json:"processNote,omitempty"`

	// Typically this field would be 1..1 where this party is accountable
	Provider *common.Reference `json:"provider,omitempty"`

	// The referral resource which lists the date, practitioner, reason
	Referral *common.Reference `json:"referral,omitempty"`

	// For example, for the original treatment and follow-up exams
	Related []ExplanationOfBenefitRelated `json:"related,omitempty"`

	// This element is labeled as a modifier because the status contains codes
	Status ExplanationOfBenefitStatus `json:"status"`

	// This may contain the local bill type codes such as the US UB-04 bill type code
	SubType *common.CodeableConcept `json:"subType,omitempty"`

	// Often there are multiple jurisdiction specific valuesets which are required
	SupportingInfo []ExplanationOfBenefitSupportingInfo `json:"supportingInfo,omitempty"`

	// Totals for amounts submitted, co-pays, benefits payable etc
	Total []ExplanationOfBenefitTotal `json:"total,omitempty"`

	// Trace number for tracking purposes
	TraceNumber []common.Identifier `json:"traceNumber,omitempty"`

	// The majority of jurisdictions use: oral, pharmacy, vision, professional
	Type common.CodeableConcept `json:"type"`

	// A code to indicate whether the nature of the request is
	Use ExplanationOfBenefitUse `json:"use"`
}

// FamilyMemberHistoryStatus represents the status of a family member history record
type FamilyMemberHistoryStatus string

const (
	FamilyMemberHistoryStatusPartial        FamilyMemberHistoryStatus = "partial"
	FamilyMemberHistoryStatusCompleted      FamilyMemberHistoryStatus = "completed"
	FamilyMemberHistoryStatusEnteredInError FamilyMemberHistoryStatus = "entered-in-error"
	FamilyMemberHistoryStatusHealthUnknown  FamilyMemberHistoryStatus = "health-unknown"
)

// FamilyMemberHistoryParticipant represents who or what participated in the activities related to the family member history
type FamilyMemberHistoryParticipant struct {
	common.BackboneElement

	// Indicates who or what participated in the activities related to the family member history
	Actor common.Reference `json:"actor"`

	// Distinguishes the type of involvement of the actor in the activities related to the family member history
	Function *common.CodeableConcept `json:"function,omitempty"`
}

// FamilyMemberHistoryCondition represents the significant conditions that the family member had
type FamilyMemberHistoryCondition struct {
	common.BackboneElement

	// The actual condition specified
	Code common.CodeableConcept `json:"code"`

	// This condition contributed to the cause of death of the related person
	ContributedToDeath *bool `json:"contributedToDeath,omitempty"`

	// An area where general notes can be placed about this specific condition
	Note []Annotation `json:"note,omitempty"`

	// Either the age of onset, range of approximate age or descriptive string can be recorded
	OnsetAge    *Age           `json:"onsetAge,omitempty"`
	OnsetRange  *Range         `json:"onsetRange,omitempty"`
	OnsetPeriod *common.Period `json:"onsetPeriod,omitempty"`
	OnsetString *string        `json:"onsetString,omitempty"`

	// Indicates what happened following the condition
	Outcome *common.CodeableConcept `json:"outcome,omitempty"`
}

// FamilyMemberHistoryProcedure represents the significant procedures that the family member had
type FamilyMemberHistoryProcedure struct {
	common.BackboneElement

	// The actual procedure specified
	Code common.CodeableConcept `json:"code"`

	// This procedure contributed to the cause of death of the related person
	ContributedToDeath *bool `json:"contributedToDeath,omitempty"`

	// An area where general notes can be placed about this specific procedure
	Note []Annotation `json:"note,omitempty"`

	// Indicates what happened following the procedure
	Outcome *common.CodeableConcept `json:"outcome,omitempty"`

	// Estimated or actual date, date-time, period, or age when the procedure was performed
	PerformedAge    *Age           `json:"performedAge,omitempty"`
	PerformedRange  *Range         `json:"performedRange,omitempty"`
	PerformedPeriod *common.Period `json:"performedPeriod,omitempty"`
	PerformedString *string        `json:"performedString,omitempty"`
}

// FamilyMemberHistory represents significant health conditions for a person related to the patient
type FamilyMemberHistory struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "FamilyMemberHistory"

	// Use estimatedAge to indicate whether the age is actual or not
	AgeAge    *Age    `json:"ageAge,omitempty"`
	AgeRange  *Range  `json:"ageRange,omitempty"`
	AgeString *string `json:"ageString,omitempty"`

	// The actual or approximate date of birth of the relative
	BornPeriod *common.Period `json:"bornPeriod,omitempty"`
	BornDate   *string        `json:"bornDate,omitempty"`
	BornString *string        `json:"bornString,omitempty"`

	// The significant conditions that the family member had
	Condition []FamilyMemberHistoryCondition `json:"condition,omitempty"`

	// Describes why the family member's history is not available
	DataAbsentReason *common.CodeableConcept `json:"dataAbsentReason,omitempty"`

	// This should be captured even if the same as the date on the List aggregating the full family history
	Date *string `json:"date,omitempty"`

	// Deceased flag or the actual or approximate age of the relative at the time of death
	DeceasedBoolean *bool   `json:"deceasedBoolean,omitempty"`
	DeceasedAge     *Age    `json:"deceasedAge,omitempty"`
	DeceasedRange   *Range  `json:"deceasedRange,omitempty"`
	DeceasedDate    *string `json:"deceasedDate,omitempty"`
	DeceasedString  *string `json:"deceasedString,omitempty"`

	// This element is labeled as a modifier because the fact that age is estimated can/should change the results
	EstimatedAge *bool `json:"estimatedAge,omitempty"`

	// This is a business identifier, not a resource identifier
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The URL pointing to a FHIR-defined protocol, guideline, orderset or other definition
	InstantiatesCanonical []string `json:"instantiatesCanonical,omitempty"`

	// This might be an HTML page, PDF, etc. or could just be a non-resolvable URI identifier
	InstantiatesUri []string `json:"instantiatesUri,omitempty"`

	// This will either be a name or a description; e.g. "Aunt Susan", "my cousin with the red hair"
	Name *string `json:"name,omitempty"`

	// This property allows a non condition-specific note to be made about the related person
	Note []Annotation `json:"note,omitempty"`

	// Indicates who or what participated in the activities related to the family member history
	Participant []FamilyMemberHistoryParticipant `json:"participant,omitempty"`

	// This is not the family member
	Patient common.Reference `json:"patient"`

	// The significant procedures that the family member had
	Procedure []FamilyMemberHistoryProcedure `json:"procedure,omitempty"`

	// Textual reasons can be captured using reasonCode.text
	Reason []CodeableReference `json:"reason,omitempty"`

	// The type of relationship this person has to the patient
	Relationship common.CodeableConcept `json:"relationship"`

	// This element should ideally reflect whether the individual is genetically male or female
	Sex *common.CodeableConcept `json:"sex,omitempty"`

	// This element is labeled as a modifier because the status contains codes that mark the resource as not currently valid
	Status FamilyMemberHistoryStatus `json:"status"`
}

// FlagStatus represents the status of a flag
type FlagStatus string

const (
	FlagStatusActive         FlagStatus = "active"
	FlagStatusInactive       FlagStatus = "inactive"
	FlagStatusEnteredInError FlagStatus = "entered-in-error"
)

// Flag represents prospective warnings of potential issues when providing care to the patient
type Flag struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Flag"

	// The person, organization or device that created the flag
	Author *common.Reference `json:"author,omitempty"`

	// The value set will often need to be adjusted based on local business rules and usage context
	Category []common.CodeableConcept `json:"category,omitempty"`

	// If non-coded, use CodeableConcept.text. This element should always be included in the narrative
	Code common.CodeableConcept `json:"code"`

	// If both Flag.encounter and Flag.period are valued, then Flag.period.start shall not be before Encounter.period.start
	Encounter *common.Reference `json:"encounter,omitempty"`

	// This is a business identifier, not a resource identifier
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The period of time from the activation of the flag to inactivation of the flag
	Period *common.Period `json:"period,omitempty"`

	// This element is labeled as a modifier because the status contains codes that mark the resource as not currently valid
	Status FlagStatus `json:"status"`

	// The patient, related person, location, group, organization, or practitioner this flag is associated with
	Subject common.Reference `json:"subject"`
}

// FormularyItemStatus represents the status of a formulary item
type FormularyItemStatus string

const (
	FormularyItemStatusActive         FormularyItemStatus = "active"
	FormularyItemStatusEnteredInError FormularyItemStatus = "entered-in-error"
	FormularyItemStatusInactive       FormularyItemStatus = "inactive"
)

// FormularyItem represents a product or service that is available through a program
type FormularyItem struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "FormularyItem"

	// A code (or set of codes) that specify the product or service that is identified by this formulary item
	Code *common.CodeableConcept `json:"code,omitempty"`

	// Business identifier for this formulary item
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// This status is intended to identify if the formulary item in a local system is in active use within the formulary
	Status *FormularyItemStatus `json:"status,omitempty"`
}

// GoalLifecycleStatus represents the lifecycle status of a goal
type GoalLifecycleStatus string

const (
	GoalLifecycleStatusProposed       GoalLifecycleStatus = "proposed"
	GoalLifecycleStatusPlanned        GoalLifecycleStatus = "planned"
	GoalLifecycleStatusAccepted       GoalLifecycleStatus = "accepted"
	GoalLifecycleStatusActive         GoalLifecycleStatus = "active"
	GoalLifecycleStatusOnHold         GoalLifecycleStatus = "on-hold"
	GoalLifecycleStatusCompleted      GoalLifecycleStatus = "completed"
	GoalLifecycleStatusCancelled      GoalLifecycleStatus = "cancelled"
	GoalLifecycleStatusEnteredInError GoalLifecycleStatus = "entered-in-error"
	GoalLifecycleStatusRejected       GoalLifecycleStatus = "rejected"
)

// GoalTarget represents when multiple targets are present for a single goal instance
type GoalTarget struct {
	common.BackboneElement

	// A CodeableConcept with just a text would be used instead of a string if the field was usually coded
	DetailQuantity        *common.Quantity        `json:"detailQuantity,omitempty"`
	DetailRange           *Range                  `json:"detailRange,omitempty"`
	DetailCodeableConcept *common.CodeableConcept `json:"detailCodeableConcept,omitempty"`
	DetailString          *string                 `json:"detailString,omitempty"`
	DetailBoolean         *bool                   `json:"detailBoolean,omitempty"`
	DetailInteger         *int                    `json:"detailInteger,omitempty"`
	DetailRatio           *Ratio                  `json:"detailRatio,omitempty"`

	// Indicates either the date or the duration after start by which the goal should be met
	DueDate     *string   `json:"dueDate,omitempty"`
	DueDuration *Duration `json:"dueDuration,omitempty"`

	// The parameter whose value is being tracked, e.g. body weight, blood pressure, or hemoglobin A1c level
	Measure *common.CodeableConcept `json:"measure,omitempty"`
}

// Goal represents the intended objective(s) for a patient, group or organization care
type Goal struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Goal"

	// Describes the progression, or lack thereof, towards the goal against the target
	AchievementStatus *common.CodeableConcept `json:"achievementStatus,omitempty"`

	// The identified conditions and other health record elements that are intended to be addressed by the goal
	Addresses []common.Reference `json:"addresses,omitempty"`

	// Indicates a category the goal falls within
	Category []common.CodeableConcept `json:"category,omitempty"`

	// For example, getting a yellow fever vaccination for a planned trip is a goal that is designed to be completed
	Continuous *bool `json:"continuous,omitempty"`

	// If no code is available, use CodeableConcept.text
	Description common.CodeableConcept `json:"description"`

	// This is a business identifier, not a resource identifier
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// This element is labeled as a modifier because the lifecycleStatus contains codes that mark the resource as not currently valid
	LifecycleStatus GoalLifecycleStatus `json:"lifecycleStatus"`

	// May be used for progress notes, concerns or other related information that doesn't actually describe the goal itself
	Note []Annotation `json:"note,omitempty"`

	// Note that this should not duplicate the goal status
	Outcome []CodeableReference `json:"outcome,omitempty"`

	// Extensions are available to track priorities as established by each participant
	Priority *common.CodeableConcept `json:"priority,omitempty"`

	// This is the individual or team responsible for establishing the goal
	Source *common.Reference `json:"source,omitempty"`

	// The date or event after which the goal should begin being pursued
	StartDate            *string                 `json:"startDate,omitempty"`
	StartCodeableConcept *common.CodeableConcept `json:"startCodeableConcept,omitempty"`

	// To see the date for past statuses, query history
	StatusDate *string `json:"statusDate,omitempty"`

	// This will typically be captured for statuses such as rejected, on-hold or cancelled
	StatusReason *string `json:"statusReason,omitempty"`

	// Identifies the patient, group or organization for whom the goal is being established
	Subject common.Reference `json:"subject"`

	// When multiple targets are present for a single goal instance, all targets must be met for the overall goal to be met
	Target []GoalTarget `json:"target,omitempty"`
}

// GroupType represents the type of group
type GroupType string

const (
	GroupTypePerson            GroupType = "person"
	GroupTypeAnimal            GroupType = "animal"
	GroupTypePractitioner      GroupType = "practitioner"
	GroupTypeDevice            GroupType = "device"
	GroupTypeCareTeam          GroupType = "careteam"
	GroupTypeHealthcareService GroupType = "healthcareservice"
	GroupTypeLocation          GroupType = "location"
	GroupTypeOrganization      GroupType = "organization"
	GroupTypeRelatedPerson     GroupType = "relatedperson"
	GroupTypeSpecimen          GroupType = "specimen"
)

// GroupMembership represents the basis for membership in the group
type GroupMembership string

const (
	GroupMembershipDefinitional GroupMembership = "definitional"
	GroupMembershipEnumerated   GroupMembership = "enumerated"
)

// GroupCharacteristic represents characteristics that must be true for an entity to be a member of the group
type GroupCharacteristic struct {
	common.BackboneElement

	// A code that identifies the kind of trait being asserted
	Code common.CodeableConcept `json:"code"`

	// This is labeled as "Is Modifier" because applications cannot wrongly include excluded members
	Exclude bool `json:"exclude"`

	// The period over which the characteristic is tested
	Period *common.Period `json:"period,omitempty"`

	// For Range, it means members of the group have a value that falls somewhere within the specified range
	ValueCodeableConcept *common.CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueBoolean         *bool                   `json:"valueBoolean,omitempty"`
	ValueQuantity        *common.Quantity        `json:"valueQuantity,omitempty"`
	ValueRange           *Range                  `json:"valueRange,omitempty"`
	ValueReference       *common.Reference       `json:"valueReference,omitempty"`
}

// GroupMember represents resource instances that are members of the group
type GroupMember struct {
	common.BackboneElement

	// A reference to the entity that is a member of the group
	Entity common.Reference `json:"entity"`

	// A flag to indicate that the member is no longer in the group
	Inactive *bool `json:"inactive,omitempty"`

	// The period that the member was in the group, if known
	Period *common.Period `json:"period,omitempty"`
}

// Group represents a defined collection of entities that may be discussed or acted upon collectively
type Group struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Group"

	// Indicates whether the record for the group is available for use or is merely being retained for historical purposes
	Active *bool `json:"active,omitempty"`

	// All the identified characteristics must be true for an entity to a member of the group
	Characteristic []GroupCharacteristic `json:"characteristic,omitempty"`

	// This would generally be omitted for Person resources
	Code *common.CodeableConcept `json:"code,omitempty"`

	// Explanation of what the group represents and how it is intended to be used
	Description *string `json:"description,omitempty"`

	// This is a business identifier, not a resource identifier
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// This does not strictly align with ownership of a herd or flock, but may suffice to represent that relationship
	ManagingEntity *common.Reference `json:"managingEntity,omitempty"`

	// Identifies the resource instances that are members of the group
	Member []GroupMember `json:"member,omitempty"`

	// Basis for membership in the Group
	Membership GroupMembership `json:"membership"`

	// A label assigned to the group for human identification and communication
	Name *string `json:"name,omitempty"`

	// Note that the quantity may be less than the number of members if some of the members are not active
	Quantity *int `json:"quantity,omitempty"`

	// Group members SHALL be of the appropriate resource type
	Type GroupType `json:"type"`
}

// GenomicStudyStatus represents the status of a genomic study
type GenomicStudyStatus string

const (
	GenomicStudyStatusRegistered     GenomicStudyStatus = "registered"
	GenomicStudyStatusAvailable      GenomicStudyStatus = "available"
	GenomicStudyStatusCancelled      GenomicStudyStatus = "cancelled"
	GenomicStudyStatusEnteredInError GenomicStudyStatus = "entered-in-error"
	GenomicStudyStatusUnknown        GenomicStudyStatus = "unknown"
)

// GenomicStudyAnalysisPerformer represents the performer of an analysis
type GenomicStudyAnalysisPerformer struct {
	common.BackboneElement

	// Performer for the analysis event
	Actor *common.Reference `json:"actor,omitempty"`

	// Indicates the role of the performer
	Role *common.CodeableConcept `json:"role,omitempty"`
}

// GenomicStudyAnalysisInput represents inputs for the analysis event
type GenomicStudyAnalysisInput struct {
	common.BackboneElement

	// File containing input data
	File *common.Reference `json:"file,omitempty"`

	// The analysis event or other GenomicStudy that generated this input file
	GeneratedByIdentifier *common.Identifier `json:"generatedByIdentifier,omitempty"`

	// The analysis event or other GenomicStudy that generated this input file
	GeneratedByReference *common.Reference `json:"generatedByReference,omitempty"`

	// Type of input data, e.g., BAM, CRAM, or FASTA
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// GenomicStudyAnalysisOutput represents outputs for the analysis event
type GenomicStudyAnalysisOutput struct {
	common.BackboneElement

	// File containing output data
	File *common.Reference `json:"file,omitempty"`

	// Type of output data, e.g., VCF, MAF, or BAM
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// GenomicStudyAnalysisDevice represents devices used for the analysis
type GenomicStudyAnalysisDevice struct {
	common.BackboneElement

	// Device used for the analysis
	Device *common.Reference `json:"device,omitempty"`

	// Specific function for the device used for the analysis
	Function *common.CodeableConcept `json:"function,omitempty"`
}

// GenomicStudyAnalysis represents the details about a specific analysis
type GenomicStudyAnalysis struct {
	common.BackboneElement

	// Type of the genomic changes studied in the analysis
	ChangeType []common.CodeableConcept `json:"changeType,omitempty"`

	// The date of the analysis event
	Date *string `json:"date,omitempty"`

	// Devices used for the analysis
	Device []GenomicStudyAnalysisDevice `json:"device,omitempty"`

	// The focus of the analysis
	Focus []common.Reference `json:"focus,omitempty"`

	// The reference genome build that is used in this analysis
	GenomeBuild *common.CodeableConcept `json:"genomeBuild,omitempty"`

	// Identifiers for the analysis event
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Inputs for the analysis event
	Input []GenomicStudyAnalysisInput `json:"input,omitempty"`

	// The defined protocol that describes the analysis
	InstantiatesCanonical *string `json:"instantiatesCanonical,omitempty"`

	// The URL pointing to an externally maintained protocol
	InstantiatesUri *string `json:"instantiatesUri,omitempty"`

	// Type of the methods used in the analysis
	MethodType []common.CodeableConcept `json:"methodType,omitempty"`

	// Any notes capture with the analysis event
	Note []Annotation `json:"note,omitempty"`

	// Outputs for the analysis event
	Output []GenomicStudyAnalysisOutput `json:"output,omitempty"`

	// Performer for the analysis event
	Performer []GenomicStudyAnalysisPerformer `json:"performer,omitempty"`

	// The protocol that was performed for the analysis event
	ProtocolPerformed *common.Reference `json:"protocolPerformed,omitempty"`

	// Genomic regions actually called in the analysis event
	RegionsCalled []common.Reference `json:"regionsCalled,omitempty"`

	// The genomic regions to be studied in the analysis
	RegionsStudied []common.Reference `json:"regionsStudied,omitempty"`

	// The specimen used in the analysis event
	Specimen []common.Reference `json:"specimen,omitempty"`

	// Name of the analysis event (human friendly)
	Title *string `json:"title,omitempty"`
}

// GenomicStudy represents a set of analyses performed to analyze and generate genomic data
type GenomicStudy struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "GenomicStudy"

	// The details about a specific analysis that was performed
	Analysis []GenomicStudyAnalysis `json:"analysis,omitempty"`

	// Event resources that the genomic study is based on
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// Description of the genomic study
	Description *string `json:"description,omitempty"`

	// The healthcare event with which this genomics study is associated
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Identifiers for this genomic study
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The defined protocol that describes the study
	InstantiatesCanonical *string `json:"instantiatesCanonical,omitempty"`

	// The URL pointing to an externally maintained protocol
	InstantiatesUri *string `json:"instantiatesUri,omitempty"`

	// Healthcare professionals who interpreted the genomic study
	Interpreter []common.Reference `json:"interpreter,omitempty"`

	// Comments related to the genomic study
	Note []Annotation `json:"note,omitempty"`

	// Why the genomic study was performed
	Reason []CodeableReference `json:"reason,omitempty"`

	// Healthcare professional who requested or referred the genomic study
	Referrer *common.Reference `json:"referrer,omitempty"`

	// When the genomic study was started
	StartDate *string `json:"startDate,omitempty"`

	// The status of the genomic study
	Status GenomicStudyStatus `json:"status"`

	// The primary subject of the genomic study
	Subject common.Reference `json:"subject"`

	// The type of the study
	Type []common.CodeableConcept `json:"type,omitempty"`
}

// GraphDefinitionNode represents a potential target for the link
type GraphDefinitionNode struct {
	common.BackboneElement

	// Information about why this node is of interest in this graph definition
	Description *string `json:"description,omitempty"`

	// Internal ID of node - target for link references
	NodeId string `json:"nodeId"`

	// Profile for the target resource
	Profile *string `json:"profile,omitempty"`

	// Type of resource this link refers to
	Type string `json:"type"`
}

// GraphDefinitionLinkCompartment represents compartment consistency rules
type GraphDefinitionLinkCompartment struct {
	common.BackboneElement

	// Identifies the compartment
	Code common.CodeableConcept `json:"code"`

	// Describes how the compartment rule is used
	Description *string `json:"description,omitempty"`

	// Identifies the compartment
	Expression *string `json:"expression,omitempty"`

	// identical | matching | different | no-rule
	Rule GraphDefinitionLinkCompartmentRule `json:"rule"`

	// Custom rule, as a FHIRPath expression
	Use common.CodeableConcept `json:"use"`
}

// GraphDefinitionLinkCompartmentRule represents the compartment rule
type GraphDefinitionLinkCompartmentRule string

const (
	GraphDefinitionLinkCompartmentRuleIdentical GraphDefinitionLinkCompartmentRule = "identical"
	GraphDefinitionLinkCompartmentRuleMatching  GraphDefinitionLinkCompartmentRule = "matching"
	GraphDefinitionLinkCompartmentRuleDifferent GraphDefinitionLinkCompartmentRule = "different"
	GraphDefinitionLinkCompartmentRuleNoRule    GraphDefinitionLinkCompartmentRule = "no-rule"
)

// GraphDefinitionLink represents links this graph makes rules about
type GraphDefinitionLink struct {
	common.BackboneElement

	// Compartment consistency rules
	Compartment []GraphDefinitionLinkCompartment `json:"compartment,omitempty"`

	// Information about why this link is of interest in this graph definition
	Description *string `json:"description,omitempty"`

	// Maximum occurrences for this link
	Max *string `json:"max,omitempty"`

	// Minimum occurrences for this link
	Min *int `json:"min,omitempty"`

	// At least one of the parameters must have the value {ref} which identifies the focus resource
	Params *string `json:"params,omitempty"`

	// The path expression cannot contain a resolve() function
	Path *string `json:"path,omitempty"`

	// Which slice (if profiled)
	SliceName *string `json:"sliceName,omitempty"`

	// The source node for this link
	SourceId string `json:"sourceId"`

	// The target node for this link
	TargetId string `json:"targetId"`
}

// GraphDefinition represents a formal computable definition of a graph of resources
type GraphDefinition struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "GraphDefinition"

	// May be a web site, an email address, a telephone number, etc.
	Contact []ContactDetail `json:"contact,omitempty"`

	// Copyright statement
	Copyright *string `json:"copyright,omitempty"`

	// Copyright label
	CopyrightLabel *string `json:"copyrightLabel,omitempty"`

	// The date is often not tracked until the resource is published
	Date *string `json:"date,omitempty"`

	// This description can be used to capture details such as comments about misuse
	Description *string `json:"description,omitempty"`

	// Allows filtering of graph definitions that are appropriate for use versus not
	Experimental *bool `json:"experimental,omitempty"`

	// A formal identifier that is used to identify this GraphDefinition
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// It may be possible for the graph definition to be used in jurisdictions
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`

	// Links this graph makes rules about
	Link []GraphDefinitionLink `json:"link,omitempty"`

	// The name is not expected to be globally unique
	Name string `json:"name"`

	// Potential target for the link
	Node []GraphDefinitionNode `json:"node,omitempty"`

	// Usually an organization but may be an individual
	Publisher *string `json:"publisher,omitempty"`

	// This element does not describe the usage of the graph definition
	Purpose *string `json:"purpose,omitempty"`

	// The Node at which instances of this graph start
	Start *string `json:"start,omitempty"`

	// Allows filtering of graph definitions that are appropriate for use versus not
	Status GraphDefinitionStatus `json:"status"`

	// This name does not need to be machine-processing friendly
	Title *string `json:"title,omitempty"`

	// Can be a urn:uuid: or a urn:oid: but real http: addresses are preferred
	Url *string `json:"url,omitempty"`

	// When multiple useContexts are specified, there is no expectation that all or any of the contexts apply
	UseContext []UsageContext `json:"useContext,omitempty"`

	// There may be different graph definition instances that have the same identifier
	Version *string `json:"version,omitempty"`

	// Version algorithm
	VersionAlgorithmString *string `json:"versionAlgorithmString,omitempty"`
}

// GraphDefinitionStatus represents the status of a graph definition
type GraphDefinitionStatus string

const (
	GraphDefinitionStatusDraft   GraphDefinitionStatus = "draft"
	GraphDefinitionStatusActive  GraphDefinitionStatus = "active"
	GraphDefinitionStatusRetired GraphDefinitionStatus = "retired"
	GraphDefinitionStatusUnknown GraphDefinitionStatus = "unknown"
)

// GuidanceResponseStatus represents the status of a guidance response
type GuidanceResponseStatus string

const (
	GuidanceResponseStatusSuccess        GuidanceResponseStatus = "success"
	GuidanceResponseStatusDataRequested  GuidanceResponseStatus = "data-requested"
	GuidanceResponseStatusDataRequired   GuidanceResponseStatus = "data-required"
	GuidanceResponseStatusInProgress     GuidanceResponseStatus = "in-progress"
	GuidanceResponseStatusFailure        GuidanceResponseStatus = "failure"
	GuidanceResponseStatusEnteredInError GuidanceResponseStatus = "entered-in-error"
)

// GuidanceResponse represents a formal response to a guidance request
type GuidanceResponse struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "GuidanceResponse"

	// If the evaluation could not be completed due to lack of information
	DataRequirement []interface{} `json:"dataRequirement,omitempty"`

	// This will typically be the encounter the event occurred within
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Messages resulting from the evaluation of the artifact or artifacts
	EvaluationMessage *common.Reference `json:"evaluationMessage,omitempty"`

	// Allows a service to provide unique, business identifiers for the response
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// An identifier, CodeableConcept or canonical reference to the guidance that was requested
	ModuleUri             *string                 `json:"moduleUri,omitempty"`
	ModuleCanonical       *string                 `json:"moduleCanonical,omitempty"`
	ModuleCodeableConcept *common.CodeableConcept `json:"moduleCodeableConcept,omitempty"`

	// Provides a mechanism to communicate additional information about the response
	Note []Annotation `json:"note,omitempty"`

	// Indicates when the guidance response was processed
	OccurrenceDateTime *string `json:"occurrenceDateTime,omitempty"`

	// The output parameters of the evaluation, if any
	OutputParameters *common.Reference `json:"outputParameters,omitempty"`

	// Provides a reference to the device that performed the guidance
	Performer *common.Reference `json:"performer,omitempty"`

	// Indicates the reason the request was initiated
	Reason []CodeableReference `json:"reason,omitempty"`

	// The identifier of the request associated with this response
	RequestIdentifier *common.Identifier `json:"requestIdentifier,omitempty"`

	// The actions, if any, produced by the evaluation of the artifact
	Result []common.Reference `json:"result,omitempty"`

	// This element is labeled as a modifier because the status contains codes that mark the resource as not currently valid
	Status GuidanceResponseStatus `json:"status"`

	// The patient for which the request was processed
	Subject *common.Reference `json:"subject,omitempty"`
}

// HealthcareServiceEligibility represents specific eligibility requirements for using the service
type HealthcareServiceEligibility struct {
	common.BackboneElement

	// Coded value for the eligibility
	Code *common.CodeableConcept `json:"code,omitempty"`

	// The description of service eligibility
	Comment *string `json:"comment,omitempty"`
}

// HealthcareService represents the details of a healthcare service available at a location
type HealthcareService struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "HealthcareService"

	// This element is labeled as a modifier because it may be used to mark that the resource was created in error
	Active *bool `json:"active,omitempty"`

	// Indicates whether or not a prospective consumer will require an appointment for a particular service
	AppointmentRequired *bool `json:"appointmentRequired,omitempty"`

	// More detailed availability information may be provided in associated Schedule/Slot resources
	Availability []Availability `json:"availability,omitempty"`

	// Selecting a Service Category then determines the list of relevant service types
	Category []common.CodeableConcept `json:"category,omitempty"`

	// These could be such things as is wheelchair accessible
	Characteristic []common.CodeableConcept `json:"characteristic,omitempty"`

	// Would expect that a user would not see this information on a search results
	Comment *string `json:"comment,omitempty"`

	// When using this property it indicates that the service is available with this language
	Communication []common.CodeableConcept `json:"communication,omitempty"`

	// Official contacts for the HealthcareService itself for specific purposes
	Contact []ExtendedContactDetail `json:"contact,omitempty"`

	// The locations referenced by the coverage area can include both specific locations and areas
	CoverageArea []common.Reference `json:"coverageArea,omitempty"`

	// Does this service have specific eligibility requirements that need to be met in order to use the service?
	Eligibility []HealthcareServiceEligibility `json:"eligibility,omitempty"`

	// Technical endpoints providing access to services operated for the specific healthcare services
	Endpoint []common.Reference `json:"endpoint,omitempty"`

	// Extra details about the service that can't be placed in the other fields
	ExtraDetails *string `json:"extraDetails,omitempty"`

	// External identifiers for this item
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The location(s) where this healthcare service may be provided
	Location []common.Reference `json:"location,omitempty"`

	// Further description of the service as it would be presented to a consumer while searching
	Name *string `json:"name,omitempty"`

	// For example, if there is a generic Radiology service that offers CT Scans, MRIs, etc
	OfferedIn []common.Reference `json:"offeredIn,omitempty"`

	// If there is a photo/symbol associated with this HealthcareService
	Photo *Attachment `json:"photo,omitempty"`

	// Programs are often defined externally to an Organization
	Program []common.CodeableConcept `json:"program,omitempty"`

	// This property is recommended to be the same as the Location's managingOrganization
	ProvidedBy *common.Reference `json:"providedBy,omitempty"`

	// Ways that the service accepts referrals
	ReferralMethod []common.CodeableConcept `json:"referralMethod,omitempty"`

	// The provision means being commissioned by, contractually obliged or financially sourced
	ServiceProvisionCode []common.CodeableConcept `json:"serviceProvisionCode,omitempty"`

	// Collection of specialties handled by the Healthcare service
	Specialty []common.CodeableConcept `json:"specialty,omitempty"`

	// The specific type of service that may be delivered or performed
	Type []common.CodeableConcept `json:"type,omitempty"`
}

// ImagingSelectionStatus represents the status of an imaging selection
type ImagingSelectionStatus string

const (
	ImagingSelectionStatusAvailable      ImagingSelectionStatus = "available"
	ImagingSelectionStatusEnteredInError ImagingSelectionStatus = "entered-in-error"
	ImagingSelectionStatusUnknown        ImagingSelectionStatus = "unknown"
)

// ImagingSelectionInstanceImageRegion2DType represents the type of 2D image region
type ImagingSelectionInstanceImageRegion2DType string

const (
	ImagingSelectionInstanceImageRegion2DTypePoint        ImagingSelectionInstanceImageRegion2DType = "point"
	ImagingSelectionInstanceImageRegion2DTypePolyline     ImagingSelectionInstanceImageRegion2DType = "polyline"
	ImagingSelectionInstanceImageRegion2DTypeInterpolated ImagingSelectionInstanceImageRegion2DType = "interpolated"
	ImagingSelectionInstanceImageRegion2DTypeCircle       ImagingSelectionInstanceImageRegion2DType = "circle"
	ImagingSelectionInstanceImageRegion2DTypeEllipse      ImagingSelectionInstanceImageRegion2DType = "ellipse"
)

// ImagingSelectionInstanceImageRegion3DType represents the type of 3D image region
type ImagingSelectionInstanceImageRegion3DType string

const (
	ImagingSelectionInstanceImageRegion3DTypePoint      ImagingSelectionInstanceImageRegion3DType = "point"
	ImagingSelectionInstanceImageRegion3DTypeMultipoint ImagingSelectionInstanceImageRegion3DType = "multipoint"
	ImagingSelectionInstanceImageRegion3DTypePolyline   ImagingSelectionInstanceImageRegion3DType = "polyline"
	ImagingSelectionInstanceImageRegion3DTypePolygon    ImagingSelectionInstanceImageRegion3DType = "polygon"
	ImagingSelectionInstanceImageRegion3DTypeEllipse    ImagingSelectionInstanceImageRegion3DType = "ellipse"
	ImagingSelectionInstanceImageRegion3DTypeEllipsoid  ImagingSelectionInstanceImageRegion3DType = "ellipsoid"
)

// ImagingSelectionPerformer represents selector of the instances – human or machine
type ImagingSelectionPerformer struct {
	common.BackboneElement

	// Author – human or machine
	Actor *common.Reference `json:"actor,omitempty"`

	// Distinguishes the type of involvement of the performer
	Function *common.CodeableConcept `json:"function,omitempty"`
}

// ImagingSelectionInstanceImageRegion2D represents each imaging selection instance or frame list 2D image region
type ImagingSelectionInstanceImageRegion2D struct {
	common.BackboneElement

	// For a description of how 2D coordinates are encoded
	Coordinate []float64 `json:"coordinate"`

	// See DICOM PS3.3 C.10.5.1.2
	RegionType ImagingSelectionInstanceImageRegion2DType `json:"regionType"`
}

// ImagingSelectionInstanceImageRegion3D represents each imaging selection 3D image region
type ImagingSelectionInstanceImageRegion3D struct {
	common.BackboneElement

	// For a description of how 3D coordinates are encoded
	Coordinate []float64 `json:"coordinate"`

	// See DICOM PS3.3 C.18.9.1.2
	RegionType ImagingSelectionInstanceImageRegion3DType `json:"regionType"`
}

// ImagingSelectionInstance represents each imaging selection includes one or more selected DICOM SOP instances
type ImagingSelectionInstance struct {
	common.BackboneElement

	// Each imaging selection instance or frame list might includes 2D image regions
	ImageRegion2D []ImagingSelectionInstanceImageRegion2D `json:"imageRegion2D,omitempty"`

	// Each imaging selection might includes 3D image regions
	ImageRegion3D []ImagingSelectionInstanceImageRegion3D `json:"imageRegion3D,omitempty"`

	// Note: A multiframe instance has a single instance number
	Number *int `json:"number,omitempty"`

	// See DICOM PS3.3 C.12.1
	SopClass *common.Coding `json:"sopClass,omitempty"`

	// Selected subset of the SOP Instance
	Subset []string `json:"subset,omitempty"`

	// See DICOM PS3.3 C.12.1
	Uid string `json:"uid"`
}

// ImagingSelection represents a selection of DICOM SOP instances and/or frames within a single Study and Series
type ImagingSelection struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "ImagingSelection"

	// A list of the diagnostic requests that resulted in this imaging selection being performed
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// The anatomic structures examined
	BodySite *CodeableReference `json:"bodySite,omitempty"`

	// Classifies the imaging selection
	Category []common.CodeableConcept `json:"category,omitempty"`

	// All code-value and, if present, component.code-component.value pairs need to be taken into account
	Code common.CodeableConcept `json:"code"`

	// The imaging study from which the imaging selection is made
	DerivedFrom []common.Reference `json:"derivedFrom,omitempty"`

	// Typical endpoint types include DICOM WADO-RS
	Endpoint []common.Reference `json:"endpoint,omitempty"`

	// An imaging selection may reference a DICOM resource
	Focus []common.Reference `json:"focus,omitempty"`

	// See DICOM PS3.3 C.7.4.1
	FrameOfReferenceUid *string `json:"frameOfReferenceUid,omitempty"`

	// This is a business identifier, not a resource identifier
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Each imaging selection includes one or more selected DICOM SOP instances
	Instance []ImagingSelectionInstance `json:"instance,omitempty"`

	// The date and time this imaging selection was created
	Issued *string `json:"issued,omitempty"`

	// Selector of the instances – human or machine
	Performer []ImagingSelectionPerformer `json:"performer,omitempty"`

	// See DICOM PS3.3 C.7.3
	SeriesNumber *int `json:"seriesNumber,omitempty"`

	// See DICOM PS3.3 C.7.3
	SeriesUid *string `json:"seriesUid,omitempty"`

	// Unknown does not represent "other" - one of the defined statuses must apply
	Status ImagingSelectionStatus `json:"status"`

	// See DICOM PS3.3 C.7.2
	StudyUid *string `json:"studyUid,omitempty"`

	// The patient, or group of patients, location, device, organization, procedure or practitioner
	Subject *common.Reference `json:"subject,omitempty"`
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

// InsurancePlan represents a health insurance product
type InsurancePlan struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "InsurancePlan"

	// The product name for marketing
	Alias []string `json:"alias,omitempty"`

	// Contact details for the plan
	Contact []ExtendedContactDetail `json:"contact,omitempty"`

	// Coverage details
	Coverage []InsurancePlanCoverage `json:"coverage,omitempty"`

	// The geographic region in which a health insurance plan's benefits apply
	CoverageArea []common.Reference `json:"coverageArea,omitempty"`

	// Technical endpoint
	Endpoint []common.Reference `json:"endpoint,omitempty"`

	// Business identifier for the plan
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Official name of the plan
	Name *string `json:"name,omitempty"`

	// Networks and tiers
	Network []common.Reference `json:"network,omitempty"`

	// Plan details
	Plan []InsurancePlanPlan `json:"plan,omitempty"`

	// Plan issuer
	OwnedBy *common.Reference `json:"ownedBy,omitempty"`

	// The period of time that the plan is active
	Period *common.Period `json:"period,omitempty"`

	// The current state of the plan
	Status *InsurancePlanStatus `json:"status,omitempty"`

	// Kind of product
	Type []common.CodeableConcept `json:"type,omitempty"`
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

	// Business identifier for the InventoryReport
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// An inventory listing section (grouped by any of the attributes)
	InventoryListing []InventoryReportInventoryListing `json:"inventoryListing,omitempty"`

	// A note associated with the InventoryReport
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
	RelatedArtifact []RelatedArtifact `json:"relatedArtifact,omitempty"`

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
	UseContext []UsageContext `json:"useContext,omitempty"`

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
	RelatedArtifact []RelatedArtifact `json:"relatedArtifact,omitempty"`

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
	UseContext []UsageContext `json:"useContext,omitempty"`

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
	RelatedArtifact []RelatedArtifact `json:"relatedArtifact,omitempty"`

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
	UseContext []UsageContext `json:"useContext,omitempty"`
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
	UseContext []UsageContext `json:"useContext,omitempty"`

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

	// This description can be used to capture details such as why the search parameter was built
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
	URL string `json:"url"`

	// A detailed description of how the search parameter is used from a clinical perspective
	Usage *string `json:"usage,omitempty"`

	// When multiple useContexts are specified, there is no expectation that all or any of the contexts apply
	UseContext []UsageContext `json:"useContext,omitempty"`

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
	UseContext []UsageContext `json:"useContext,omitempty"`

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
	URL string `json:"url"`

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

// NutritionProductStatus represents status of nutrition product
type NutritionProductStatus string

const (
	NutritionProductStatusActive         NutritionProductStatus = "active"
	NutritionProductStatusInactive       NutritionProductStatus = "inactive"
	NutritionProductStatusEnteredInError NutritionProductStatus = "entered-in-error"
)

// NutritionProductNutrient represents nutrient information
type NutritionProductNutrient struct {
	common.BackboneElement

	// The amount of nutrient expressed in one or more units
	Amount []Ratio `json:"amount,omitempty"`

	// The (relevant) nutrients in the product
	Item *CodeableReference `json:"item,omitempty"`
}

// NutritionProductIngredient represents ingredient information
type NutritionProductIngredient struct {
	common.BackboneElement

	// The amount of ingredient that is in the product
	Amount []Ratio `json:"amount,omitempty"`

	// The ingredient contained in the product
	Item CodeableReference `json:"item"`
}

// NutritionProductCharacteristic represents descriptive properties
type NutritionProductCharacteristic struct {
	common.BackboneElement

	// A code specifying which characteristic of the product is being described
	Type common.CodeableConcept `json:"type"`

	// The description of the characteristic
	ValueCodeableConcept *common.CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueString          *string                 `json:"valueString,omitempty"`
	ValueQuantity        *common.Quantity        `json:"valueQuantity,omitempty"`
	ValueBase64Binary    *string                 `json:"valueBase64Binary,omitempty"`
	ValueAttachment      *Attachment             `json:"valueAttachment,omitempty"`
	ValueBoolean         *bool                   `json:"valueBoolean,omitempty"`
}

// NutritionProductInstance represents instance-level information
type NutritionProductInstance struct {
	common.BackboneElement

	// Biological source event identifier
	BiologicalSourceEvent *common.Identifier `json:"biologicalSourceEvent,omitempty"`

	// The time after which the product is no longer expected to be in proper condition
	Expiry *string `json:"expiry,omitempty"`

	// The identifier for the physical instance
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The identification of the batch or lot of the product
	LotNumber *string `json:"lotNumber,omitempty"`

	// The name for the specific product
	Name *string `json:"name,omitempty"`

	// The amount of items or instances that the resource considers
	Quantity *common.Quantity `json:"quantity,omitempty"`

	// The time after which the product is no longer expected to be in proper condition
	UseBy *string `json:"useBy,omitempty"`
}

// NutritionProduct represents a food or supplement consumed by patients
type NutritionProduct struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "NutritionProduct"

	// Nutrition products can have different classifications
	Category []common.CodeableConcept `json:"category,omitempty"`

	// Specifies descriptive properties of the nutrition product
	Characteristic []NutritionProductCharacteristic `json:"characteristic,omitempty"`

	// The code assigned to the product
	Code *common.CodeableConcept `json:"code,omitempty"`

	// Ingredients contained in this product
	Ingredient []NutritionProductIngredient `json:"ingredient,omitempty"`

	// Conveys instance-level information about this product item
	Instance []NutritionProductInstance `json:"instance,omitempty"`

	// Allergens that are known or suspected to be a part of this nutrition product
	KnownAllergen []CodeableReference `json:"knownAllergen,omitempty"`

	// The organisation responsible for the product
	Manufacturer []common.Reference `json:"manufacturer,omitempty"`

	// Comments made about the product
	Note []Annotation `json:"note,omitempty"`

	// Nutrient information
	Nutrient []NutritionProductNutrient `json:"nutrient,omitempty"`

	// active | inactive | entered-in-error
	Status NutritionProductStatus `json:"status"`
}

// ObservationDefinitionStatus represents status of observation definition
type ObservationDefinitionStatus string

const (
	ObservationDefinitionStatusDraft   ObservationDefinitionStatus = "draft"
	ObservationDefinitionStatusActive  ObservationDefinitionStatus = "active"
	ObservationDefinitionStatusRetired ObservationDefinitionStatus = "retired"
	ObservationDefinitionStatusUnknown ObservationDefinitionStatus = "unknown"
)

// ObservationDefinitionComponent represents component observations
type ObservationDefinitionComponent struct {
	common.BackboneElement

	// Describes what will be observed
	Code common.CodeableConcept `json:"code"`

	// The data types allowed for the value element
	PermittedDataType []string `json:"permittedDataType,omitempty"`

	// Units allowed for the valueQuantity element
	PermittedUnit []common.Coding `json:"permittedUnit,omitempty"`

	// A set of qualified values associated with a context and conditions
	QualifiedValue []interface{} `json:"qualifiedValue,omitempty"`
}

// ObservationDefinition represents definitional characteristics for observations
type ObservationDefinition struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "ObservationDefinition"

	// The date may be more recent than the approval date
	ApprovalDate *string `json:"approvalDate,omitempty"`

	// Only used if the defined observation is to be made directly on a body part
	BodySite *common.CodeableConcept `json:"bodySite,omitempty"`

	// This element allows various categorization schemes
	Category []common.CodeableConcept `json:"category,omitempty"`

	// Describes what will be observed
	Code common.CodeableConcept `json:"code"`

	// Some observations have multiple component observations
	Component []ObservationDefinitionComponent `json:"component,omitempty"`

	// May be a web site, an email address, a telephone number, etc.
	Contact []ContactDetail `json:"contact,omitempty"`

	// Copyright statement relating to the ObservationDefinition
	Copyright *string `json:"copyright,omitempty"`

	// The (c) symbol should NOT be included in this string
	CopyrightLabel *string `json:"copyrightLabel,omitempty"`

	// The date is often not tracked until the resource is published
	Date *string `json:"date,omitempty"`

	// The canonical URL pointing to another FHIR-defined ObservationDefinition
	DerivedFromCanonical []string `json:"derivedFromCanonical,omitempty"`

	// The URL pointing to an externally-defined observation definition
	DerivedFromUri []string `json:"derivedFromUri,omitempty"`

	// This description can be used to capture details
	Description *string `json:"description,omitempty"`

	// When multiple occurrences of device are present
	Device []common.Reference `json:"device,omitempty"`

	// The effective period for an ObservationDefinition
	EffectivePeriod *common.Period `json:"effectivePeriod,omitempty"`

	// Allows filtering of ObservationDefinition that are appropriate for use vs. not
	Experimental *bool `json:"experimental,omitempty"`

	// This ObservationDefinition defines a group observation
	HasMember []common.Reference `json:"hasMember,omitempty"`

	// This is a business identifier, not a resource identifier
	Identifier *common.Identifier `json:"identifier,omitempty"`

	// It may be possible for the ObservationDefinition to be used in jurisdictions
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`

	// If specified, this is usually after the approval date
	LastReviewDate *string `json:"lastReviewDate,omitempty"`

	// Only used if not implicit in observation code
	Method *common.CodeableConcept `json:"method,omitempty"`

	// An example of observation allowing multiple results
	MultipleResultsAllowed *bool `json:"multipleResultsAllowed,omitempty"`

	// The name is not expected to be globally unique
	Name *string `json:"name,omitempty"`

	// The type of individual/organization/device that is expected to act
	PerformerType *common.CodeableConcept `json:"performerType,omitempty"`

	// The data types allowed for the value element
	PermittedDataType []string `json:"permittedDataType,omitempty"`

	// Units allowed for the valueQuantity element
	PermittedUnit []common.Coding `json:"permittedUnit,omitempty"`

	// The preferred name to be used when reporting the results
	PreferredReportName *string `json:"preferredReportName,omitempty"`

	// Helps establish the "authority/credibility" of the ObservationDefinition
	Publisher *string `json:"publisher,omitempty"`

	// This element does not describe the usage of the ObservationDefinition
	Purpose *string `json:"purpose,omitempty"`

	// A set of qualified values associated with a context and conditions
	QualifiedValue []interface{} `json:"qualifiedValue,omitempty"`

	// Only used for in vitro observations
	Specimen []common.Reference `json:"specimen,omitempty"`

	// draft | active | retired | unknown
	Status ObservationDefinitionStatus `json:"status"`

	// Examples: person, animal, device, air, surface
	Subject []common.CodeableConcept `json:"subject,omitempty"`

	// A short, descriptive, user-friendly title
	Title *string `json:"title,omitempty"`

	// Can be a urn:uuid: or a urn:oid:, but real http: addresses are preferred
	Url *string `json:"url,omitempty"`

	// When multiple usageContexts are specified
	UseContext []UsageContext `json:"useContext,omitempty"`

	// There may be different observation definitions that have the same url
	Version *string `json:"version,omitempty"`
}

// OperationDefinitionKind represents kind of operation definition
type OperationDefinitionKind string

const (
	OperationDefinitionKindOperation OperationDefinitionKind = "operation"
	OperationDefinitionKindQuery     OperationDefinitionKind = "query"
)

// OperationDefinitionStatus represents status of operation definition
type OperationDefinitionStatus string

const (
	OperationDefinitionStatusDraft   OperationDefinitionStatus = "draft"
	OperationDefinitionStatusActive  OperationDefinitionStatus = "active"
	OperationDefinitionStatusRetired OperationDefinitionStatus = "retired"
	OperationDefinitionStatusUnknown OperationDefinitionStatus = "unknown"
)

// OperationDefinitionParameter represents operation parameter
type OperationDefinitionParameter struct {
	common.BackboneElement

	// In previous versions of FHIR, there was an extension for this
	AllowedType []string `json:"allowedType,omitempty"`

	// Binds to a value set if this parameter is coded
	Binding *interface{} `json:"binding,omitempty"`

	// Describes the meaning or use of this parameter
	Documentation *string `json:"documentation,omitempty"`

	// The maximum number of times this element is permitted to appear
	Max string `json:"max"`

	// The minimum number of times this parameter SHALL appear
	Min int `json:"min"`

	// This name must be a token
	Name string `json:"name"`

	// Query Definitions only have one output parameter, named "result"
	Part []OperationDefinitionParameter `json:"part,omitempty"`

	// Resolution applies if the referenced parameter exists
	ReferencedFrom []interface{} `json:"referencedFrom,omitempty"`

	// If present, indicates that the parameter applies when the operation is being invoked
	Scope []string `json:"scope,omitempty"`

	// How the parameter is understood as a search parameter
	SearchType *string `json:"searchType,omitempty"`

	// The type for this parameter
	Type *string `json:"type,omitempty"`

	// Whether this is an input or an output parameter
	Use string `json:"use"`
}

// OperationDefinition represents a formal computable definition of an operation
type OperationDefinition struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "OperationDefinition"

	// What http methods can be used for the operation
	AffectsState *bool `json:"affectsState,omitempty"`

	// A constrained profile can make optional parameters required or not used
	Base *string `json:"base,omitempty"`

	// For maximum compatibility, use only lowercase ASCII characters
	Code string `json:"code"`

	// Additional information about how to use this operation or named query
	Comment *string `json:"comment,omitempty"`

	// May be a web site, an email address, a telephone number, etc.
	Contact []ContactDetail `json:"contact,omitempty"`

	// Note that this is not the same as the resource last-modified-date
	Date *string `json:"date,omitempty"`

	// This description can be used to capture details
	Description *string `json:"description,omitempty"`

	// Allows filtering of operation definition that are appropriate for use vs. not
	Experimental *bool `json:"experimental,omitempty"`

	// Operations that are idempotent may be invoked by performing an HTTP GET
	Idempotent *bool `json:"idempotent,omitempty"`

	// Indicates whether this operation can be invoked on a particular instance
	Instance bool `json:"instance"`

	// It may be possible for the operation definition to be used in jurisdictions
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`

	// Named queries are invoked differently, and have different capabilities
	Kind OperationDefinitionKind `json:"kind"`

	// The name is not expected to be globally unique
	Name string `json:"name"`

	// The combinations are suggestions as to which sets of parameters to use together
	Overload []interface{} `json:"overload,omitempty"`

	// Query Definitions only have one output parameter, named "result"
	Parameter []OperationDefinitionParameter `json:"parameter,omitempty"`

	// Usually an organization but may be an individual
	Publisher *string `json:"publisher,omitempty"`

	// This element does not describe the usage of the operation definition
	Purpose *string `json:"purpose,omitempty"`

	// If the type is an abstract resource then the operation can be invoked
	Resource []string `json:"resource,omitempty"`

	// draft | active | retired | unknown
	Status OperationDefinitionStatus `json:"status"`

	// Indicates whether this operation or named query can be invoked at the system level
	System bool `json:"system"`

	// This name does not need to be machine-processing friendly
	Title *string `json:"title,omitempty"`

	// Indicates whether this operation can be invoked at the resource type level
	Type bool `json:"type"`

	// Can be a urn:uuid: or a urn:oid: but real http: addresses are preferred
	Url *string `json:"url,omitempty"`

	// When multiple useContexts are specified
	UseContext []UsageContext `json:"useContext,omitempty"`

	// There may be different operation definitions that have the same url
	Version *string `json:"version,omitempty"`
}

// OperationOutcomeIssueSeverity represents severity of operation outcome issue
type OperationOutcomeIssueSeverity string

const (
	OperationOutcomeIssueSeverityFatal       OperationOutcomeIssueSeverity = "fatal"
	OperationOutcomeIssueSeverityError       OperationOutcomeIssueSeverity = "error"
	OperationOutcomeIssueSeverityWarning     OperationOutcomeIssueSeverity = "warning"
	OperationOutcomeIssueSeverityInformation OperationOutcomeIssueSeverity = "information"
	OperationOutcomeIssueSeveritySuccess     OperationOutcomeIssueSeverity = "success"
)

// OperationOutcomeIssue represents an error, warning, or information message
type OperationOutcomeIssue struct {
	common.BackboneElement

	// Code values should align with the severity
	Code string `json:"code"`

	// A human readable description of the error issue
	Details *common.CodeableConcept `json:"details,omitempty"`

	// This may be a description of how a value is erroneous
	Diagnostics *string `json:"diagnostics,omitempty"`

	// The root of the FHIRPath is the resource or bundle that generated OperationOutcome
	Expression []string `json:"expression,omitempty"`

	// The root of the XPath is the resource or bundle that generated OperationOutcome
	Location []string `json:"location,omitempty"`

	// Indicates whether the issue indicates a variation from successful processing
	Severity OperationOutcomeIssueSeverity `json:"severity"`
}

// OperationOutcome represents a collection of error, warning, or information messages
type OperationOutcome struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "OperationOutcome"

	// An error, warning, or information message that results from a system action
	Issue []OperationOutcomeIssue `json:"issue"`
}

// OrganizationAffiliation represents an affiliation between organizations
type OrganizationAffiliation struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "OrganizationAffiliation"

	// If this value is false, you may refer to the period to see when the role was in active use
	Active *bool `json:"active,omitempty"`

	// Definition of the role the participatingOrganization plays in the association
	Code []common.CodeableConcept `json:"code,omitempty"`

	// The contact details of communication devices available at the participatingOrganization
	Contact []ExtendedContactDetail `json:"contact,omitempty"`

	// Technical endpoints providing access to services operated for this role
	Endpoint []common.Reference `json:"endpoint,omitempty"`

	// Healthcare services provided through the role
	HealthcareService []common.Reference `json:"healthcareService,omitempty"`

	// Business identifiers that are specific to this role
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The location(s) at which the role occurs
	Location []common.Reference `json:"location,omitempty"`

	// e.g. Commonly used for Health Insurance provider networks
	Network []common.Reference `json:"network,omitempty"`

	// For example, a Spotless Cleaning Services is a supplier to General Hospital
	Organization *common.Reference `json:"organization,omitempty"`

	// See comments for OrganizationAffiliation.organization above
	ParticipatingOrganization *common.Reference `json:"participatingOrganization,omitempty"`

	// The period during which the participatingOrganization is affiliated
	Period *common.Period `json:"period,omitempty"`

	// Specific specialty of the participatingOrganization in the context of the role
	Specialty []common.CodeableConcept `json:"specialty,omitempty"`
}

// PackagedProductDefinitionPackagingProperty represents packaging property
type PackagedProductDefinitionPackagingProperty struct {
	common.BackboneElement

	// A code expressing the type of characteristic
	Type common.CodeableConcept `json:"type"`

	// A value for the characteristic
	ValueCodeableConcept *common.CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueQuantity        *common.Quantity        `json:"valueQuantity,omitempty"`
	ValueDate            *string                 `json:"valueDate,omitempty"`
	ValueBoolean         *bool                   `json:"valueBoolean,omitempty"`
	ValueAttachment      *Attachment             `json:"valueAttachment,omitempty"`
}

// PackagedProductDefinitionPackagingContainedItem represents contained item
type PackagedProductDefinitionPackagingContainedItem struct {
	common.BackboneElement

	// The number of this type of item within this packaging
	Amount *common.Quantity `json:"amount,omitempty"`

	// The actual item(s) of medication, as manufactured
	Item CodeableReference `json:"item"`
}

// PackagedProductDefinitionPackaging represents packaging
type PackagedProductDefinitionPackaging struct {
	common.BackboneElement

	// A possible alternate material for this part of the packaging
	AlternateMaterial []common.CodeableConcept `json:"alternateMaterial,omitempty"`

	// Is this a part of the packaging rather than the packaging itself
	ComponentPart *bool `json:"componentPart,omitempty"`

	// The item(s) within the packaging
	ContainedItem []PackagedProductDefinitionPackagingContainedItem `json:"containedItem,omitempty"`

	// A business identifier that is specific to this particular part of the packaging
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Manufacturer of this packaging item
	Manufacturer []common.Reference `json:"manufacturer,omitempty"`

	// Material type of the package item
	Material []common.CodeableConcept `json:"material,omitempty"`

	// Allows containers within containers
	Packaging []PackagedProductDefinitionPackaging `json:"packaging,omitempty"`

	// General characteristics of this item
	Property []PackagedProductDefinitionPackagingProperty `json:"property,omitempty"`
}

// PackagedProductDefinition represents a medically related item in a container
type PackagedProductDefinition struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "PackagedProductDefinition"

	// Additional information or supporting documentation
	AttachedDocument []common.Reference `json:"attachedDocument,omitempty"`

	// Allows the key features to be recorded
	Characteristic []PackagedProductDefinitionPackagingProperty `json:"characteristic,omitempty"`

	// A total of the complete count of contained items
	ContainedItemQuantity []common.Quantity `json:"containedItemQuantity,omitempty"`

	// Identifies if the package contains different items
	CopackagedIndicator *bool `json:"copackagedIndicator,omitempty"`

	// Textual description
	Description *string `json:"description,omitempty"`

	// A unique identifier for this package as whole
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The legal status of supply of the packaged item
	LegalStatusOfSupply []interface{} `json:"legalStatusOfSupply,omitempty"`

	// Manufacturer of this package type
	Manufacturer []common.Reference `json:"manufacturer,omitempty"`

	// Allows specifying that an item is on the market for sale
	MarketingStatus []interface{} `json:"marketingStatus,omitempty"`

	// A name for this package
	Name *string `json:"name,omitempty"`

	// The product this package model relates to
	PackageFor []common.Reference `json:"packageFor,omitempty"`

	// A packaging item, as a container for medically related items
	Packaging *PackagedProductDefinitionPackaging `json:"packaging,omitempty"`

	// The status within the lifecycle of this item
	Status *common.CodeableConcept `json:"status,omitempty"`

	// The date at which the given status became applicable
	StatusDate *string `json:"statusDate,omitempty"`

	// A high level category
	Type *common.CodeableConcept `json:"type,omitempty"`
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

// PaymentNotice represents a payment notice
type PaymentNotice struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "PaymentNotice"

	// The amount sent to the payee
	Amount Money `json:"amount"`

	// The date when this resource was created
	Created *string `json:"created,omitempty"`

	// The party who will receive or has received payment that is the subject of this notification
	Payee *common.Reference `json:"payee,omitempty"`

	// The party who notifies the payer about the payment being made
	Payer *common.Reference `json:"payer,omitempty"`

	// A note that describes or explains the processing in a human readable form
	PaymentDate *string `json:"paymentDate,omitempty"`

	// The date when the above payment action occurred
	PaymentStatus *common.CodeableConcept `json:"paymentStatus,omitempty"`

	// The practitioner who is responsible for the services rendered to the patient
	Provider *common.Reference `json:"provider,omitempty"`

	// A reference to the payment which is the subject of this notice
	Request *common.Reference `json:"request,omitempty"`

	// A reference to the payment which is the subject of this notice
	Response *common.Reference `json:"response,omitempty"`

	// The status of the resource instance
	Status PaymentNoticeStatus `json:"status"`

	// The target of the payment notice
	Target *common.Reference `json:"target,omitempty"`
}

// PaymentNoticeStatus represents status of payment notice
type PaymentNoticeStatus string

const (
	PaymentNoticeStatusActive       PaymentNoticeStatus = "active"
	PaymentNoticeStatusCancelled    PaymentNoticeStatus = "cancelled"
	PaymentNoticeStatusDraft        PaymentNoticeStatus = "draft"
	PaymentNoticeStatusEnteredError PaymentNoticeStatus = "entered-in-error"
)

// PaymentReconciliation represents payment reconciliation
type PaymentReconciliation struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "PaymentReconciliation"

	// The date when the resource was created
	Created *string `json:"created,omitempty"`

	// A human readable description of the status of the request for the reconciliation
	Disposition *string `json:"disposition,omitempty"`

	// The date of payment as indicated on the financial instrument
	FormCode *common.CodeableConcept `json:"formCode,omitempty"`

	// A unique identifier assigned to this payment reconciliation
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The Insurer who produced this adjudicated response
	Issuer *common.Reference `json:"issuer,omitempty"`

	// The outcome of a request for a reconciliation
	Outcome *PaymentReconciliationOutcome `json:"outcome,omitempty"`

	// The period of time for which payments have been gathered into this bulk payment for settlement
	Period *common.Period `json:"period,omitempty"`

	// The practitioner who is responsible for the services rendered to the patient
	ProcessNote []PaymentReconciliationProcessNote `json:"processNote,omitempty"`

	// Original request resource reference
	Request *common.Reference `json:"request,omitempty"`

	// The organization which is responsible for the services rendered to the patient
	Requestor *common.Reference `json:"requestor,omitempty"`

	// The status of the resource instance
	Status PaymentReconciliationStatus `json:"status"`

	// Total payment amount as indicated on the financial instrument
	TotalAmount Money `json:"totalAmount"`

	// The date when the above payment action occurred
	TransactionDate *string `json:"transactionDate,omitempty"`
}

// PaymentReconciliationOutcome represents outcome of payment reconciliation
type PaymentReconciliationOutcome string

const (
	PaymentReconciliationOutcomeQueued   PaymentReconciliationOutcome = "queued"
	PaymentReconciliationOutcomeComplete PaymentReconciliationOutcome = "complete"
	PaymentReconciliationOutcomeError    PaymentReconciliationOutcome = "error"
	PaymentReconciliationOutcomePartial  PaymentReconciliationOutcome = "partial"
)

// PaymentReconciliationStatus represents status of payment reconciliation
type PaymentReconciliationStatus string

const (
	PaymentReconciliationStatusActive       PaymentReconciliationStatus = "active"
	PaymentReconciliationStatusCancelled    PaymentReconciliationStatus = "cancelled"
	PaymentReconciliationStatusDraft        PaymentReconciliationStatus = "draft"
	PaymentReconciliationStatusEnteredError PaymentReconciliationStatus = "entered-in-error"
)

// PaymentReconciliationProcessNote represents process note
type PaymentReconciliationProcessNote struct {
	common.BackboneElement

	// The explanation or description associated with the processing
	Text *string `json:"text,omitempty"`

	// The business purpose of the note text
	Type *PaymentReconciliationProcessNoteType `json:"type,omitempty"`
}

// PaymentReconciliationProcessNoteType represents type of process note
type PaymentReconciliationProcessNoteType string

const (
	PaymentReconciliationProcessNoteTypeDisplay   PaymentReconciliationProcessNoteType = "display"
	PaymentReconciliationProcessNoteTypePrint     PaymentReconciliationProcessNoteType = "print"
	PaymentReconciliationProcessNoteTypePrintoper PaymentReconciliationProcessNoteType = "printoper"
)

// Permission represents permission
type Permission struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Permission"

	// The person or entity that asserts the permission
	Asserter *common.Reference `json:"asserter,omitempty"`

	// The date that permission was asserted
	Date []*string `json:"date,omitempty"`

	// The period in which the permission is active
	Justification *PermissionJustification `json:"justification,omitempty"`

	// The person or entity that is being allowed or denied access
	ManagingEntity *common.Reference `json:"managingEntity,omitempty"`

	// The person or entity that is being allowed or denied access
	Note []Annotation `json:"note,omitempty"`

	// The person or entity that is being allowed or denied access
	Status PermissionStatus `json:"status"`

	// The person or entity that is being allowed or denied access
	Subject []PermissionSubject `json:"subject,omitempty"`

	// The person or entity that is being allowed or denied access
	ValidityPeriod *common.Period `json:"validityPeriod,omitempty"`
}

// PermissionStatus represents status of permission
type PermissionStatus string

const (
	PermissionStatusActive       PermissionStatus = "active"
	PermissionStatusInactive     PermissionStatus = "inactive"
	PermissionStatusEnteredError PermissionStatus = "entered-in-error"
)

// PermissionJustification represents justification
type PermissionJustification struct {
	common.BackboneElement

	// The regulatory grounds for the permission
	Basis []common.CodeableConcept `json:"basis,omitempty"`

	// Justification for the permission
	Evidence []common.Reference `json:"evidence,omitempty"`
}

// PermissionSubject represents subject
type PermissionSubject struct {
	common.BackboneElement

	// The person or entity that is being allowed or denied access
	Characteristic []PermissionSubjectCharacteristic `json:"characteristic,omitempty"`

	// The person or entity that is being allowed or denied access
	Reference []common.Reference `json:"reference,omitempty"`
}

// PermissionSubjectCharacteristic represents characteristic
type PermissionSubjectCharacteristic struct {
	common.BackboneElement

	// The person or entity that is being allowed or denied access
	Exclude *bool `json:"exclude,omitempty"`

	// The person or entity that is being allowed or denied access
	Period *common.Period `json:"period,omitempty"`

	// The person or entity that is being allowed or denied access
	Type common.CodeableConcept `json:"type"`

	// The person or entity that is being allowed or denied access
	ValueCodeableConcept *common.CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueReference       *common.Reference       `json:"valueReference,omitempty"`
	ValueBoolean         *bool                   `json:"valueBoolean,omitempty"`
	ValueExpression      *Expression             `json:"valueExpression,omitempty"`
}

// Person represents a person
type Person struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Person"

	// Whether this person's record is in active use
	Active *bool `json:"active,omitempty"`

	// One or more addresses for the person
	Address []Address `json:"address,omitempty"`

	// The date on which the person was born
	BirthDate *string `json:"birthDate,omitempty"`

	// Administrative Gender - the gender that the person is considered to have for administration and record keeping purposes
	Gender *PersonGender `json:"gender,omitempty"`

	// Identifier for a person within a particular scope
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// A link to a resource that concerns the same actual person
	Link []PersonLink `json:"link,omitempty"`

	// The organization that is the custodian of the person record
	ManagingOrganization *common.Reference `json:"managingOrganization,omitempty"`

	// A name associated with the person
	Name []HumanName `json:"name,omitempty"`

	// An image of the person
	Photo *Attachment `json:"photo,omitempty"`

	// A contact detail for the person
	Telecom []ContactPoint `json:"telecom,omitempty"`
}

// PersonGender represents gender of person
type PersonGender string

const (
	PersonGenderMale    PersonGender = "male"
	PersonGenderFemale  PersonGender = "female"
	PersonGenderOther   PersonGender = "other"
	PersonGenderUnknown PersonGender = "unknown"
)

// PersonLink represents link to another person
type PersonLink struct {
	common.BackboneElement

	// Level of assurance that this link is associated with the target resource
	Assurance *PersonLinkAssurance `json:"assurance,omitempty"`

	// The resource to which this actual person is associated
	Target common.Reference `json:"target"`
}

// PersonLinkAssurance represents assurance level
type PersonLinkAssurance string

const (
	PersonLinkAssuranceLevel1 PersonLinkAssurance = "level1"
	PersonLinkAssuranceLevel2 PersonLinkAssurance = "level2"
	PersonLinkAssuranceLevel3 PersonLinkAssurance = "level3"
	PersonLinkAssuranceLevel4 PersonLinkAssurance = "level4"
)
