package fhir4

import (
	"encoding/json"

	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// ClinicalImpression represents a FHIR R4 ClinicalImpression resource
type ClinicalImpression struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"`

	Identifier               []common.Identifier               `json:"identifier,omitempty"`
	Status                   string                            `json:"status"`
	StatusReason             *common.CodeableConcept           `json:"statusReason,omitempty"`
	Code                     *common.CodeableConcept           `json:"code,omitempty"`
	Description              *string                           `json:"description,omitempty"`
	Subject                  common.Reference                  `json:"subject"`
	Encounter                *common.Reference                 `json:"encounter,omitempty"`
	EffectiveDateTime        *string                           `json:"effectiveDateTime,omitempty"`
	EffectivePeriod          *common.Period                    `json:"effectivePeriod,omitempty"`
	Date                     *string                           `json:"date,omitempty"`
	Assessor                 *common.Reference                 `json:"assessor,omitempty"`
	Previous                 *common.Reference                 `json:"previous,omitempty"`
	Problem                  []common.Reference                `json:"problem,omitempty"`
	Investigation            []ClinicalImpressionInvestigation `json:"investigation,omitempty"`
	Protocol                 []string                          `json:"protocol,omitempty"`
	Finding                  []ClinicalImpressionFinding       `json:"finding,omitempty"`
	PrognosisCodeableConcept []common.CodeableConcept          `json:"prognosisCodeableConcept,omitempty"`
	PrognosisReference       []common.Reference                `json:"prognosisReference,omitempty"`
	SupportingInfo           []common.Reference                `json:"supportingInfo,omitempty"`
	Note                     []common.Annotation               `json:"note,omitempty"`
}

// ClinicalImpressionInvestigation represents investigations
type ClinicalImpressionInvestigation struct {
	common.BackboneElement

	Code common.CodeableConcept `json:"code"`
	Item []common.Reference     `json:"item,omitempty"`
}

// ClinicalImpressionFinding represents findings
type ClinicalImpressionFinding struct {
	common.BackboneElement

	ItemCodeableConcept *common.CodeableConcept `json:"itemCodeableConcept,omitempty"`
	ItemReference       *common.Reference       `json:"itemReference,omitempty"`
	Basis               *string                 `json:"basis,omitempty"`
}

// MarshalJSON marshals the ClinicalImpression to JSON
func (c *ClinicalImpression) MarshalJSON() ([]byte, error) {
	type Alias ClinicalImpression
	return json.Marshal(&struct {
		*Alias
		ResourceType string `json:"resourceType"`
	}{
		Alias:        (*Alias)(c),
		ResourceType: "ClinicalImpression",
	})
}
