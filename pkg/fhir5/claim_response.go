// Package fhir5 contains FHIR R5 resource definitions
package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// ClaimResponse represents the adjudication result for a Claim (FHIR R5)
type ClaimResponse struct {
	DomainResource

	ResourceType         string                      `json:"resourceType"` // Always "ClaimResponse"
	Identifier           []common.Identifier         `json:"identifier,omitempty"`
	Status               ClaimResponseStatus         `json:"status"`
	Type                 common.CodeableConcept      `json:"type"`
	SubType              *common.CodeableConcept     `json:"subType,omitempty"`
	Use                  ClaimResponseUse            `json:"use"`
	Patient              common.Reference            `json:"patient"`
	Created              string                      `json:"created"`
	Insurer              common.Reference            `json:"insurer"`
	Requestor            *common.Reference           `json:"requestor,omitempty"`
	Request              *common.Reference           `json:"request,omitempty"`
	Outcome              ClaimResponseOutcome        `json:"outcome"`
	Disposition          *string                     `json:"disposition,omitempty"`
	PreAuthRef           *string                     `json:"preAuthRef,omitempty"`
	PreAuthPeriod        *common.Period              `json:"preAuthPeriod,omitempty"`
	PayeeType            *common.CodeableConcept     `json:"payeeType,omitempty"`
	Item                 []ClaimResponseItem         `json:"item,omitempty"`
	AddItem              []ClaimResponseAddItem      `json:"addItem,omitempty"`
	Adjudication         []ClaimResponseAdjudication `json:"adjudication,omitempty"`
	Total                []ClaimResponseTotal        `json:"total,omitempty"`
	Payment              *ClaimResponsePayment       `json:"payment,omitempty"`
	FundsReserve         *common.CodeableConcept     `json:"fundsReserve,omitempty"`
	FormCode             *common.CodeableConcept     `json:"formCode,omitempty"`
	Form                 *Attachment                 `json:"form,omitempty"`
	ProcessNote          []ClaimResponseProcessNote  `json:"processNote,omitempty"`
	CommunicationRequest []common.Reference          `json:"communicationRequest,omitempty"`
	Insurance            []ClaimResponseInsurance    `json:"insurance"`
	Error                []ClaimResponseError        `json:"error,omitempty"`
}

type ClaimResponseStatus string

const (
	ClaimResponseStatusActive         ClaimResponseStatus = "active"
	ClaimResponseStatusCancelled      ClaimResponseStatus = "cancelled"
	ClaimResponseStatusDraft          ClaimResponseStatus = "draft"
	ClaimResponseStatusEnteredInError ClaimResponseStatus = "entered-in-error"
)

type ClaimResponseUse string

const (
	ClaimResponseUseClaim   ClaimResponseUse = "claim"
	ClaimResponseUsePreauth ClaimResponseUse = "preauthorization"
	ClaimResponseUsePredet  ClaimResponseUse = "predetermination"
)

type ClaimResponseOutcome string

const (
	ClaimResponseOutcomeQueued   ClaimResponseOutcome = "queued"
	ClaimResponseOutcomeComplete ClaimResponseOutcome = "complete"
	ClaimResponseOutcomeError    ClaimResponseOutcome = "error"
	ClaimResponseOutcomePartial  ClaimResponseOutcome = "partial"
)

type ClaimResponseItem struct {
	common.BackboneElement
	ItemSequence int                         `json:"itemSequence"`
	NoteNumber   []int                       `json:"noteNumber,omitempty"`
	Adjudication []ClaimResponseAdjudication `json:"adjudication,omitempty"`
	Detail       []ClaimResponseItemDetail   `json:"detail,omitempty"`
}

type ClaimResponseItemDetail struct {
	common.BackboneElement
	DetailSequence int                                `json:"detailSequence"`
	NoteNumber     []int                              `json:"noteNumber,omitempty"`
	Adjudication   []ClaimResponseAdjudication        `json:"adjudication,omitempty"`
	SubDetail      []ClaimResponseItemDetailSubDetail `json:"subDetail,omitempty"`
}

type ClaimResponseItemDetailSubDetail struct {
	common.BackboneElement
	ProductOrService    common.CodeableConcept      `json:"productOrService"`
	ProductOrServiceEnd *common.CodeableConcept     `json:"productOrServiceEnd,omitempty"`
	Modifier            []common.CodeableConcept    `json:"modifier,omitempty"`
	ProgramCode         []common.CodeableConcept    `json:"programCode,omitempty"`
	Quantity            *common.Quantity            `json:"quantity,omitempty"`
	UnitPrice           *Money                      `json:"unitPrice,omitempty"`
	Factor              *float64                    `json:"factor,omitempty"`
	Tax                 *Money                      `json:"tax,omitempty"`
	Net                 *Money                      `json:"net,omitempty"`
	NoteNumber          []int                       `json:"noteNumber,omitempty"`
	Adjudication        []ClaimResponseAdjudication `json:"adjudication,omitempty"`
}

type ClaimResponseAddItem struct {
	common.BackboneElement
	ItemSequence            []int                        `json:"itemSequence,omitempty"`
	DetailSequence          []int                        `json:"detailSequence,omitempty"`
	SubDetailSequence       []int                        `json:"subDetailSequence,omitempty"`
	Provider                []common.Reference           `json:"provider,omitempty"`
	ProductOrService        common.CodeableConcept       `json:"productOrService"`
	ProductOrServiceEnd     *common.CodeableConcept      `json:"productOrServiceEnd,omitempty"`
	Modifier                []common.CodeableConcept     `json:"modifier,omitempty"`
	ProgramCode             []common.CodeableConcept     `json:"programCode,omitempty"`
	ServicedDate            *string                      `json:"servicedDate,omitempty"`
	ServicedPeriod          *common.Period               `json:"servicedPeriod,omitempty"`
	LocationCodeableConcept *common.CodeableConcept      `json:"locationCodeableConcept,omitempty"`
	LocationAddress         *Address                     `json:"locationAddress,omitempty"`
	LocationReference       *common.Reference            `json:"locationReference,omitempty"`
	Quantity                *common.Quantity             `json:"quantity,omitempty"`
	UnitPrice               *Money                       `json:"unitPrice,omitempty"`
	Factor                  *float64                     `json:"factor,omitempty"`
	Tax                     *Money                       `json:"tax,omitempty"`
	Net                     *Money                       `json:"net,omitempty"`
	BodySite                *common.CodeableConcept      `json:"bodySite,omitempty"`
	SubSite                 []common.CodeableConcept     `json:"subSite,omitempty"`
	NoteNumber              []int                        `json:"noteNumber,omitempty"`
	Adjudication            []ClaimResponseAdjudication  `json:"adjudication,omitempty"`
	Detail                  []ClaimResponseAddItemDetail `json:"detail,omitempty"`
}

type ClaimResponseAddItemDetail struct {
	common.BackboneElement
	ProductOrService    common.CodeableConcept                `json:"productOrService"`
	ProductOrServiceEnd *common.CodeableConcept               `json:"productOrServiceEnd,omitempty"`
	Modifier            []common.CodeableConcept              `json:"modifier,omitempty"`
	ProgramCode         []common.CodeableConcept              `json:"programCode,omitempty"`
	Quantity            *common.Quantity                      `json:"quantity,omitempty"`
	UnitPrice           *Money                                `json:"unitPrice,omitempty"`
	Factor              *float64                              `json:"factor,omitempty"`
	Tax                 *Money                                `json:"tax,omitempty"`
	Net                 *Money                                `json:"net,omitempty"`
	NoteNumber          []int                                 `json:"noteNumber,omitempty"`
	Adjudication        []ClaimResponseAdjudication           `json:"adjudication,omitempty"`
	SubDetail           []ClaimResponseAddItemDetailSubDetail `json:"subDetail,omitempty"`
}

type ClaimResponseAddItemDetailSubDetail struct {
	common.BackboneElement
	ProductOrService    common.CodeableConcept      `json:"productOrService"`
	ProductOrServiceEnd *common.CodeableConcept     `json:"productOrServiceEnd,omitempty"`
	Modifier            []common.CodeableConcept    `json:"modifier,omitempty"`
	ProgramCode         []common.CodeableConcept    `json:"programCode,omitempty"`
	Quantity            *common.Quantity            `json:"quantity,omitempty"`
	UnitPrice           *Money                      `json:"unitPrice,omitempty"`
	Factor              *float64                    `json:"factor,omitempty"`
	Tax                 *Money                      `json:"tax,omitempty"`
	Net                 *Money                      `json:"net,omitempty"`
	NoteNumber          []int                       `json:"noteNumber,omitempty"`
	Adjudication        []ClaimResponseAdjudication `json:"adjudication,omitempty"`
}

type ClaimResponseAdjudication struct {
	common.BackboneElement
	Category common.CodeableConcept  `json:"category"`
	Reason   *common.CodeableConcept `json:"reason,omitempty"`
	Amount   *Money                  `json:"amount,omitempty"`
	Value    *float64                `json:"value,omitempty"`
}

type ClaimResponseTotal struct {
	common.BackboneElement
	Category common.CodeableConcept `json:"category"`
	Amount   Money                  `json:"amount"`
}

type ClaimResponsePayment struct {
	common.BackboneElement
	Type             common.CodeableConcept  `json:"type"`
	Adjustment       *Money                  `json:"adjustment,omitempty"`
	AdjustmentReason *common.CodeableConcept `json:"adjustmentReason,omitempty"`
	Date             *string                 `json:"date,omitempty"`
	Amount           Money                   `json:"amount"`
	Identifier       *common.Identifier      `json:"identifier,omitempty"`
}

type ClaimResponseProcessNote struct {
	common.BackboneElement
	Number   *int                          `json:"number,omitempty"`
	Type     *ClaimResponseProcessNoteType `json:"type,omitempty"`
	Text     string                        `json:"text"`
	Language *common.CodeableConcept       `json:"language,omitempty"`
}

type ClaimResponseProcessNoteType string

const (
	ClaimResponseProcessNoteTypeDisplay   ClaimResponseProcessNoteType = "display"
	ClaimResponseProcessNoteTypePrint     ClaimResponseProcessNoteType = "print"
	ClaimResponseProcessNoteTypePrintoper ClaimResponseProcessNoteType = "printoper"
)

type ClaimResponseInsurance struct {
	common.BackboneElement
	Sequence            int                `json:"sequence"`
	Focal               bool               `json:"focal"`
	Identifier          *common.Identifier `json:"identifier,omitempty"`
	Coverage            common.Reference   `json:"coverage"`
	BusinessArrangement *string            `json:"businessArrangement,omitempty"`
	PreAuthRef          []string           `json:"preAuthRef,omitempty"`
	ClaimResponse       *common.Reference  `json:"claimResponse,omitempty"`
}

type ClaimResponseError struct {
	common.BackboneElement
	ItemSequence      *int                   `json:"itemSequence,omitempty"`
	DetailSequence    *int                   `json:"detailSequence,omitempty"`
	SubDetailSequence *int                   `json:"subDetailSequence,omitempty"`
	Code              common.CodeableConcept `json:"code"`
}
