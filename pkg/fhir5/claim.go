// Package fhir5 contains FHIR R5 resource definitions
package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// Claim represents a claim for reimbursement for healthcare products and services
// (FHIR R5)
type Claim struct {
	DomainResource

	ResourceType         string                   `json:"resourceType"` // Always "Claim"
	Identifier           []common.Identifier      `json:"identifier,omitempty"`
	Status               ClaimStatus              `json:"status"`
	Type                 common.CodeableConcept   `json:"type"`
	SubType              []common.CodeableConcept `json:"subType,omitempty"`
	Use                  ClaimUse                 `json:"use"`
	Patient              common.Reference         `json:"patient"`
	BillablePeriod       *common.Period           `json:"billablePeriod,omitempty"`
	Created              string                   `json:"created"`
	Enterer              *common.Reference        `json:"enterer,omitempty"`
	Insurer              *common.Reference        `json:"insurer,omitempty"`
	Provider             common.Reference         `json:"provider"`
	Priority             common.CodeableConcept   `json:"priority"`
	FundsReserve         *common.CodeableConcept  `json:"fundsReserve,omitempty"`
	Related              []ClaimRelated           `json:"related,omitempty"`
	Prescription         *common.Reference        `json:"prescription,omitempty"`
	OriginalPrescription *common.Reference        `json:"originalPrescription,omitempty"`
	Payee                *ClaimPayee              `json:"payee,omitempty"`
	Referral             *common.Reference        `json:"referral,omitempty"`
	Facility             *common.Reference        `json:"facility,omitempty"`
	CareTeam             []ClaimCareTeam          `json:"careTeam,omitempty"`
	SupportingInfo       []ClaimSupportingInfo    `json:"supportingInfo,omitempty"`
	Diagnosis            []ClaimDiagnosis         `json:"diagnosis,omitempty"`
	Procedure            []ClaimProcedure         `json:"procedure,omitempty"`
	Insurance            []ClaimInsurance         `json:"insurance"`
	Accident             *ClaimAccident           `json:"accident,omitempty"`
	Item                 []ClaimItem              `json:"item,omitempty"`
	Total                *Money                   `json:"total,omitempty"`
}

type ClaimStatus string

const (
	ClaimStatusActive         ClaimStatus = "active"
	ClaimStatusCancelled      ClaimStatus = "cancelled"
	ClaimStatusDraft          ClaimStatus = "draft"
	ClaimStatusEnteredInError ClaimStatus = "entered-in-error"
)

type ClaimUse string

const (
	ClaimUseClaim   ClaimUse = "claim"
	ClaimUsePreauth ClaimUse = "preauthorization"
	ClaimUsePredet  ClaimUse = "predetermination"
)

type ClaimRelated struct {
	common.BackboneElement
	Claim        *common.Reference       `json:"claim,omitempty"`
	Relationship *common.CodeableConcept `json:"relationship,omitempty"`
	Reference    *common.Identifier      `json:"reference,omitempty"`
}

type ClaimPayee struct {
	common.BackboneElement
	Type  common.CodeableConcept `json:"type"`
	Party *common.Reference      `json:"party,omitempty"`
}

type ClaimCareTeam struct {
	common.BackboneElement
	Sequence      int                     `json:"sequence"`
	Provider      common.Reference        `json:"provider"`
	Responsible   *bool                   `json:"responsible,omitempty"`
	Role          *common.CodeableConcept `json:"role,omitempty"`
	Qualification *common.CodeableConcept `json:"qualification,omitempty"`
}

type ClaimSupportingInfo struct {
	common.BackboneElement
	Sequence        int                     `json:"sequence"`
	Category        common.CodeableConcept  `json:"category"`
	Code            *common.CodeableConcept `json:"code,omitempty"`
	TimingDate      *string                 `json:"timingDate,omitempty"`
	TimingPeriod    *common.Period          `json:"timingPeriod,omitempty"`
	ValueBoolean    *bool                   `json:"valueBoolean,omitempty"`
	ValueString     *string                 `json:"valueString,omitempty"`
	ValueQuantity   *common.Quantity        `json:"valueQuantity,omitempty"`
	ValueAttachment *Attachment             `json:"valueAttachment,omitempty"`
	ValueReference  *common.Reference       `json:"valueReference,omitempty"`
	Reason          *common.CodeableConcept `json:"reason,omitempty"`
}

type ClaimDiagnosis struct {
	common.BackboneElement
	Sequence                 int                      `json:"sequence"`
	DiagnosisCodeableConcept *common.CodeableConcept  `json:"diagnosisCodeableConcept,omitempty"`
	DiagnosisReference       *common.Reference        `json:"diagnosisReference,omitempty"`
	Type                     []common.CodeableConcept `json:"type,omitempty"`
	OnAdmission              *common.CodeableConcept  `json:"onAdmission,omitempty"`
	PackageCode              *common.CodeableConcept  `json:"packageCode,omitempty"`
}

type ClaimProcedure struct {
	common.BackboneElement
	Sequence                 int                      `json:"sequence"`
	Type                     []common.CodeableConcept `json:"type,omitempty"`
	Date                     *string                  `json:"date,omitempty"`
	ProcedureCodeableConcept *common.CodeableConcept  `json:"procedureCodeableConcept,omitempty"`
	ProcedureReference       *common.Reference        `json:"procedureReference,omitempty"`
}

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

type ClaimAccident struct {
	common.BackboneElement
	Date              string                  `json:"date"`
	Type              *common.CodeableConcept `json:"type,omitempty"`
	LocationAddress   *Address                `json:"locationAddress,omitempty"`
	LocationReference *common.Reference       `json:"locationReference,omitempty"`
}

type ClaimItem struct {
	common.BackboneElement
	Sequence                int                      `json:"sequence"`
	CareTeamSequence        []int                    `json:"careTeamSequence,omitempty"`
	DiagnosisSequence       []int                    `json:"diagnosisSequence,omitempty"`
	ProcedureSequence       []int                    `json:"procedureSequence,omitempty"`
	InformationSequence     []int                    `json:"informationSequence,omitempty"`
	Revenue                 *common.CodeableConcept  `json:"revenue,omitempty"`
	Category                *common.CodeableConcept  `json:"category,omitempty"`
	ProductOrService        common.CodeableConcept   `json:"productOrService"`
	ProductOrServiceEnd     *common.CodeableConcept  `json:"productOrServiceEnd,omitempty"`
	Modifier                []common.CodeableConcept `json:"modifier,omitempty"`
	ProgramCode             []common.CodeableConcept `json:"programCode,omitempty"`
	ServicedDate            *string                  `json:"servicedDate,omitempty"`
	ServicedPeriod          *common.Period           `json:"servicedPeriod,omitempty"`
	LocationCodeableConcept *common.CodeableConcept  `json:"locationCodeableConcept,omitempty"`
	LocationAddress         *Address                 `json:"locationAddress,omitempty"`
	LocationReference       *common.Reference        `json:"locationReference,omitempty"`
	Quantity                *common.Quantity         `json:"quantity,omitempty"`
	UnitPrice               *Money                   `json:"unitPrice,omitempty"`
	Factor                  *float64                 `json:"factor,omitempty"`
	Tax                     *Money                   `json:"tax,omitempty"`
	Net                     *Money                   `json:"net,omitempty"`
	Udi                     []common.Reference       `json:"udi,omitempty"`
	BodySite                *common.CodeableConcept  `json:"bodySite,omitempty"`
	SubSite                 []common.CodeableConcept `json:"subSite,omitempty"`
	Encounter               []common.Reference       `json:"encounter,omitempty"`
	Detail                  []ClaimItemDetail        `json:"detail,omitempty"`
}

type ClaimItemDetail struct {
	common.BackboneElement
	Sequence            int                        `json:"sequence"`
	Revenue             *common.CodeableConcept    `json:"revenue,omitempty"`
	Category            *common.CodeableConcept    `json:"category,omitempty"`
	ProductOrService    common.CodeableConcept     `json:"productOrService"`
	ProductOrServiceEnd *common.CodeableConcept    `json:"productOrServiceEnd,omitempty"`
	Modifier            []common.CodeableConcept   `json:"modifier,omitempty"`
	ProgramCode         []common.CodeableConcept   `json:"programCode,omitempty"`
	Quantity            *common.Quantity           `json:"quantity,omitempty"`
	UnitPrice           *Money                     `json:"unitPrice,omitempty"`
	Factor              *float64                   `json:"factor,omitempty"`
	Tax                 *Money                     `json:"tax,omitempty"`
	Net                 *Money                     `json:"net,omitempty"`
	Udi                 []common.Reference         `json:"udi,omitempty"`
	NoteNumber          []int                      `json:"noteNumber,omitempty"`
	SubDetail           []ClaimItemDetailSubDetail `json:"subDetail,omitempty"`
}

type ClaimItemDetailSubDetail struct {
	common.BackboneElement
	Sequence            int                      `json:"sequence"`
	Revenue             *common.CodeableConcept  `json:"revenue,omitempty"`
	Category            *common.CodeableConcept  `json:"category,omitempty"`
	ProductOrService    common.CodeableConcept   `json:"productOrService"`
	ProductOrServiceEnd *common.CodeableConcept  `json:"productOrServiceEnd,omitempty"`
	Modifier            []common.CodeableConcept `json:"modifier,omitempty"`
	ProgramCode         []common.CodeableConcept `json:"programCode,omitempty"`
	Quantity            *common.Quantity         `json:"quantity,omitempty"`
	UnitPrice           *Money                   `json:"unitPrice,omitempty"`
	Factor              *float64                 `json:"factor,omitempty"`
	Tax                 *Money                   `json:"tax,omitempty"`
	Net                 *Money                   `json:"net,omitempty"`
	Udi                 []common.Reference       `json:"udi,omitempty"`
	NoteNumber          []int                    `json:"noteNumber,omitempty"`
}
