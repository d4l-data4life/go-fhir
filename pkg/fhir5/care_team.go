// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

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
