// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
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
