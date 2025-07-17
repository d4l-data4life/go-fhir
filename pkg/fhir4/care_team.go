package fhir4

import (
	"encoding/json"

	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// CareTeam represents a FHIR R4 CareTeam resource
type CareTeam struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"`

	Identifier           []common.Identifier      `json:"identifier,omitempty"`
	Status               *string                  `json:"status,omitempty"`
	Category             []common.CodeableConcept `json:"category,omitempty"`
	Name                 *string                  `json:"name,omitempty"`
	Subject              *common.Reference        `json:"subject,omitempty"`
	Encounter            *common.Reference        `json:"encounter,omitempty"`
	Period               *common.Period           `json:"period,omitempty"`
	Participant          []CareTeamParticipant    `json:"participant,omitempty"`
	ReasonCode           []common.CodeableConcept `json:"reasonCode,omitempty"`
	ReasonReference      []common.Reference       `json:"reasonReference,omitempty"`
	ManagingOrganization []common.Reference       `json:"managingOrganization,omitempty"`
	Telecom              []common.ContactPoint    `json:"telecom,omitempty"`
	Note                 []common.Annotation      `json:"note,omitempty"`
}

// CareTeamParticipant represents a participant in a care team
type CareTeamParticipant struct {
	common.BackboneElement

	Role       []common.CodeableConcept `json:"role,omitempty"`
	Member     *common.Reference        `json:"member,omitempty"`
	OnBehalfOf *common.Reference        `json:"onBehalfOf,omitempty"`
	Period     *common.Period           `json:"period,omitempty"`
}

// MarshalJSON marshals the CareTeam to JSON
func (c *CareTeam) MarshalJSON() ([]byte, error) {
	type Alias CareTeam
	return json.Marshal(&struct {
		*Alias
		ResourceType string `json:"resourceType"`
	}{
		Alias:        (*Alias)(c),
		ResourceType: "CareTeam",
	})
}
