// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// EpisodeOfCareStatus represents the status of an episode of care
type EpisodeOfCareStatus string

const (
	EpisodeOfCareStatusPlanned        EpisodeOfCareStatus = "planned"
	EpisodeOfCareStatusWaitlist       EpisodeOfCareStatus = "waitlist"
	EpisodeOfCareStatusActive         EpisodeOfCareStatus = "active"
	EpisodeOfCareStatusOnhold         EpisodeOfCareStatus = "onhold"
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

// EpisodeOfCareDiagnosis represents diagnoses relevant to this episode of care
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
