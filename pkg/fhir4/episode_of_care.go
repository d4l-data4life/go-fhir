package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// EpisodeOfCare represents an association between a patient and an organization / healthcare provider(s)
type EpisodeOfCare struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "EpisodeOfCare"

	// The billing system may choose to allocate billable items associated with the EpisodeOfCare to different referenced Accounts
	Account []common.Reference `json:"account,omitempty"`

	// The practitioner that is the care manager/care coordinator for this patient
	CareManager *common.Reference `json:"careManager,omitempty"`

	// The list of diagnosis relevant to this episode of care
	Diagnosis []EpisodeOfCareDiagnosis `json:"diagnosis,omitempty"`

	// The EpisodeOfCare may be known by different identifiers for different contexts of use
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The organization that has assumed the specific responsibilities for the specified duration
	ManagingOrganization *common.Reference `json:"managingOrganization,omitempty"`

	// The patient who is the focus of this episode of care
	Patient common.Reference `json:"patient"`

	// The interval during which the managing organization assumes the defined responsibility
	Period *common.Period `json:"period,omitempty"`

	// The practitioner that is the care manager/care coordinator for this patient
	ReferralRequest []common.Reference `json:"referralRequest,omitempty"`

	// The history of statuses that the EpisodeOfCare has been through
	StatusHistory []EpisodeOfCareStatusHistory `json:"statusHistory,omitempty"`

	// This element is labeled as a modifier because the status contains the code entered-in-error that marks the episode as not currently valid
	Status EpisodeOfCareStatus `json:"status"`

	// A classification of the type of episode of care; e.g. specialist referral, disease management, type of funded care
	Type []common.CodeableConcept `json:"type,omitempty"`
}

// EpisodeOfCareStatus represents the status of the episode of care
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

	// planned | waitlist | active | onhold | finished | cancelled
	Status EpisodeOfCareStatus `json:"status"`
}

// EpisodeOfCareDiagnosis represents the list of diagnosis relevant to this episode of care
type EpisodeOfCareDiagnosis struct {
	common.BackboneElement

	// A list of conditions/problems/diagnoses that this episode of care is intended to be providing care for
	Condition common.Reference `json:"condition"`

	// Ranking of the diagnosis (for each role type)
	Rank *int `json:"rank,omitempty"`

	// Role that this diagnosis has within the episode of care (e.g. admission, billing, discharge …)
	Role *common.CodeableConcept `json:"role,omitempty"`
}
