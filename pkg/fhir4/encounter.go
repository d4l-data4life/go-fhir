// Package fhir4 contains FHIR R4 (version 4.0.1) resource definitions
package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// Encounter represents an interaction during which services are provided to the patient
type Encounter struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Encounter"

	// The billing system may choose to allocate billable items associated with the Encounter
	Account []common.Reference `json:"account,omitempty"`

	// The appointment that scheduled this encounter
	Appointment []common.Reference `json:"appointment,omitempty"`

	// The request this encounter satisfies (e.g. incoming referral or procedure request)
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// Concepts representing classification of patient encounter (required in R4)
	Class common.Coding `json:"class"`

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

	// For systems that need to know which was the primary diagnosis
	ReasonCode []common.CodeableConcept `json:"reasonCode,omitempty"`

	// For systems that need to know which was the primary diagnosis
	ReasonReference []common.Reference `json:"reasonReference,omitempty"`

	// The organization that is primarily responsible for this Encounter's services
	ServiceProvider *common.Reference `json:"serviceProvider,omitempty"`

	// Broad categorization of the service that is to be provided (e.g. cardiology)
	ServiceType *common.CodeableConcept `json:"serviceType,omitempty"`

	// Note that internal business rules will determine the appropriate transitions
	Status EncounterStatus `json:"status"`

	// The current status is always found in the current version of the resource
	StatusHistory []EncounterStatusHistory `json:"statusHistory,omitempty"`

	// While the encounter is always about the patient, the patient might not actually be known
	Subject *common.Reference `json:"subject,omitempty"`

	// Since there are many ways to further classify encounters, this element is 0..*
	Type []common.CodeableConcept `json:"type,omitempty"`
}

// EncounterStatus represents the status of the encounter
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

// EncounterLocationStatus represents the status of the location
type EncounterLocationStatus string

const (
	EncounterLocationStatusPlanned   EncounterLocationStatus = "planned"
	EncounterLocationStatusActive    EncounterLocationStatus = "active"
	EncounterLocationStatusReserved  EncounterLocationStatus = "reserved"
	EncounterLocationStatusCompleted EncounterLocationStatus = "completed"
)

// EncounterStatusHistory represents the status history of the encounter
type EncounterStatusHistory struct {
	common.BackboneElement

	// The time that the episode was in the specified status
	Period common.Period `json:"period"`

	// planned | arrived | triaged | in-progress | onleave | finished | cancelled | entered-in-error | unknown
	Status EncounterStatus `json:"status"`
}

// EncounterClassHistory represents the class history of the encounter
type EncounterClassHistory struct {
	common.BackboneElement

	// inpatient | outpatient | ambulatory | emergency +
	Class common.Coding `json:"class"`

	// The time that the episode was in the specified class
	Period common.Period `json:"period"`
}

// EncounterParticipant represents a participant in the encounter
type EncounterParticipant struct {
	common.BackboneElement

	// Persons involved in the encounter other than the patient
	Individual *common.Reference `json:"individual,omitempty"`

	// The period of time that the specified participant participated in the encounter
	Period *common.Period `json:"period,omitempty"`

	// The participant type indicates how an individual participates in an encounter
	Type []common.CodeableConcept `json:"type,omitempty"`
}

// EncounterDiagnosis represents a diagnosis relevant to the encounter
type EncounterDiagnosis struct {
	common.BackboneElement

	// For systems that need to know which was the primary diagnosis
	Condition common.Reference `json:"condition"`

	// Ranking of the diagnosis (for each role type)
	Rank *int `json:"rank,omitempty"`

	// Role that this diagnosis has within the encounter (e.g. admission, billing, discharge)
	Use *common.CodeableConcept `json:"use,omitempty"`
}

// EncounterHospitalization represents hospitalization details for the encounter
type EncounterHospitalization struct {
	common.BackboneElement

	// From where patient was admitted (physician referral, transfer)
	AdmitSource *common.CodeableConcept `json:"admitSource,omitempty"`

	// Location/organization to which the patient is discharged
	Destination *common.Reference `json:"destination,omitempty"`

	// For example, a patient may request both a dairy-free and nut-free diet preference
	DietPreference []common.CodeableConcept `json:"dietPreference,omitempty"`

	// Category or kind of location after discharge
	DischargeDisposition *common.CodeableConcept `json:"dischargeDisposition,omitempty"`

	// The location/organization from which the patient came before admission
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

// EncounterLocation represents a location of the encounter
type EncounterLocation struct {
	common.BackboneElement

	// The location where the encounter takes place
	Location common.Reference `json:"location"`

	// Time period during which the patient was present at the location
	Period *common.Period `json:"period,omitempty"`

	// This information is de-normalized from the Location resource
	PhysicalType *common.CodeableConcept `json:"physicalType,omitempty"`

	// When the patient is no longer active at a location
	Status *EncounterLocationStatus `json:"status,omitempty"`
}
