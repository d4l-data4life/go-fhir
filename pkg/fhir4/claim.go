package fhir4

import (
	"encoding/json"

	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// Claim represents a FHIR R4 Claim resource
type Claim struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"`

	Identifier           []common.Identifier     `json:"identifier,omitempty"`
	Status               string                  `json:"status"`
	Type                 common.CodeableConcept  `json:"type"`
	SubType              *common.CodeableConcept `json:"subType,omitempty"`
	Use                  string                  `json:"use"`
	Patient              common.Reference        `json:"patient"`
	BillablePeriod       *common.Period          `json:"billablePeriod,omitempty"`
	Created              string                  `json:"created"`
	Enterer              *common.Reference       `json:"enterer,omitempty"`
	Insurer              *common.Reference       `json:"insurer,omitempty"`
	Provider             common.Reference        `json:"provider"`
	Priority             common.CodeableConcept  `json:"priority"`
	FundsReserve         *common.CodeableConcept `json:"fundsReserve,omitempty"`
	Related              []ClaimRelated          `json:"related,omitempty"`
	Prescription         *common.Reference       `json:"prescription,omitempty"`
	OriginalPrescription *common.Reference       `json:"originalPrescription,omitempty"`
	Payee                *ClaimPayee             `json:"payee,omitempty"`
	Referral             *common.Reference       `json:"referral,omitempty"`
	Facility             *common.Reference       `json:"facility,omitempty"`
	CareTeam             []ClaimCareTeam         `json:"careTeam,omitempty"`
	Insurance            []ClaimInsurance        `json:"insurance"`
	Accident             *ClaimAccident          `json:"accident,omitempty"`
	Item                 []ClaimItem             `json:"item,omitempty"`
	Total                *common.Money           `json:"total,omitempty"`
}

// ClaimRelated represents related claims
type ClaimRelated struct {
	common.BackboneElement

	Claim        *common.Reference       `json:"claim,omitempty"`
	Relationship *common.CodeableConcept `json:"relationship,omitempty"`
	Reference    *common.Identifier      `json:"reference,omitempty"`
}

// ClaimPayee represents payee information
type ClaimPayee struct {
	common.BackboneElement

	Type  common.CodeableConcept `json:"type"`
	Party *common.Reference      `json:"party,omitempty"`
}

// ClaimCareTeam represents care team members
type ClaimCareTeam struct {
	common.BackboneElement

	Sequence      int                     `json:"sequence"`
	Provider      common.Reference        `json:"provider"`
	Responsible   *bool                   `json:"responsible,omitempty"`
	Role          *common.CodeableConcept `json:"role,omitempty"`
	Qualification *common.CodeableConcept `json:"qualification,omitempty"`
}

// ClaimInsurance represents insurance information
type ClaimInsurance struct {
	common.BackboneElement

	Sequence            int                `json:"sequence"`
	Focal               bool               `json:"focal"`
	Identifier          *common.Identifier `json:"identifier,omitempty"`
	Coverage            common.Reference   `json:"coverage"`
	BusinessArrangement *string            `json:"businessArrangement,omitempty"`
	PreAuthRef          []string           `json:"preAuthRef,omitempty"`
	ClaimResponse       *common.Reference  `json:"claimResponse,omitempty"`
}

// ClaimAccident represents accident information
type ClaimAccident struct {
	common.BackboneElement

	Date              string                  `json:"date"`
	Type              *common.CodeableConcept `json:"type,omitempty"`
	LocationAddress   *common.Address         `json:"locationAddress,omitempty"`
	LocationReference *common.Reference       `json:"locationReference,omitempty"`
}

// ClaimItem represents claim items
type ClaimItem struct {
	common.BackboneElement

	Sequence                int                      `json:"sequence"`
	CareTeamLinkId          []int                    `json:"careTeamLinkId,omitempty"`
	DiagnosisLinkId         []int                    `json:"diagnosisLinkId,omitempty"`
	ProcedureLinkId         []int                    `json:"procedureLinkId,omitempty"`
	InformationLinkId       []int                    `json:"informationLinkId,omitempty"`
	Revenue                 *common.CodeableConcept  `json:"revenue,omitempty"`
	Category                *common.CodeableConcept  `json:"category,omitempty"`
	Service                 *common.CodeableConcept  `json:"service,omitempty"`
	Modifier                []common.CodeableConcept `json:"modifier,omitempty"`
	ProgramCode             []common.CodeableConcept `json:"programCode,omitempty"`
	ServicedDate            *string                  `json:"servicedDate,omitempty"`
	ServicedPeriod          *common.Period           `json:"servicedPeriod,omitempty"`
	LocationCodeableConcept *common.CodeableConcept  `json:"locationCodeableConcept,omitempty"`
	LocationAddress         *common.Address          `json:"locationAddress,omitempty"`
	LocationReference       *common.Reference        `json:"locationReference,omitempty"`
	Quantity                *common.Quantity         `json:"quantity,omitempty"`
	UnitPrice               *common.Money            `json:"unitPrice,omitempty"`
	Factor                  *float64                 `json:"factor,omitempty"`
	Net                     *common.Money            `json:"net,omitempty"`
	Udi                     []common.Reference       `json:"udi,omitempty"`
	BodySite                *common.CodeableConcept  `json:"bodySite,omitempty"`
	SubSite                 []common.CodeableConcept `json:"subSite,omitempty"`
	Encounter               []common.Reference       `json:"encounter,omitempty"`
	Detail                  []ClaimItemDetail        `json:"detail,omitempty"`
}

// ClaimItemDetail represents claim item details
type ClaimItemDetail struct {
	common.BackboneElement

	Sequence    int                        `json:"sequence"`
	Revenue     *common.CodeableConcept    `json:"revenue,omitempty"`
	Category    *common.CodeableConcept    `json:"category,omitempty"`
	Service     *common.CodeableConcept    `json:"service,omitempty"`
	Modifier    []common.CodeableConcept   `json:"modifier,omitempty"`
	ProgramCode []common.CodeableConcept   `json:"programCode,omitempty"`
	Quantity    *common.Quantity           `json:"quantity,omitempty"`
	UnitPrice   *common.Money              `json:"unitPrice,omitempty"`
	Factor      *float64                   `json:"factor,omitempty"`
	Net         *common.Money              `json:"net,omitempty"`
	Udi         []common.Reference         `json:"udi,omitempty"`
	SubDetail   []ClaimItemDetailSubDetail `json:"subDetail,omitempty"`
}

// ClaimItemDetailSubDetail represents claim item sub-details
type ClaimItemDetailSubDetail struct {
	common.BackboneElement

	Sequence    int                      `json:"sequence"`
	Revenue     *common.CodeableConcept  `json:"revenue,omitempty"`
	Category    *common.CodeableConcept  `json:"category,omitempty"`
	Service     *common.CodeableConcept  `json:"service,omitempty"`
	Modifier    []common.CodeableConcept `json:"modifier,omitempty"`
	ProgramCode []common.CodeableConcept `json:"programCode,omitempty"`
	Quantity    *common.Quantity         `json:"quantity,omitempty"`
	UnitPrice   *common.Money            `json:"unitPrice,omitempty"`
	Factor      *float64                 `json:"factor,omitempty"`
	Net         *common.Money            `json:"net,omitempty"`
	Udi         []common.Reference       `json:"udi,omitempty"`
}

// MarshalJSON marshals the Claim to JSON
func (c *Claim) MarshalJSON() ([]byte, error) {
	type Alias Claim
	return json.Marshal(&struct {
		*Alias
		ResourceType string `json:"resourceType"`
	}{
		Alias:        (*Alias)(c),
		ResourceType: "Claim",
	})
}
