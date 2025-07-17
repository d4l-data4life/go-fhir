package fhir4

import (
	"encoding/json"

	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// ChargeItem represents a FHIR R4 ChargeItem resource
type ChargeItem struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"`

	Identifier             []common.Identifier      `json:"identifier,omitempty"`
	DefinitionUri          []string                 `json:"definitionUri,omitempty"`
	DefinitionCanonical    []string                 `json:"definitionCanonical,omitempty"`
	Status                 string                   `json:"status"`
	PartOf                 []common.Reference       `json:"partOf,omitempty"`
	Code                   common.CodeableConcept   `json:"code"`
	Subject                common.Reference         `json:"subject"`
	Context                *common.Reference        `json:"context,omitempty"`
	OccurrenceDateTime     *string                  `json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod       *common.Period           `json:"occurrencePeriod,omitempty"`
	OccurrenceTiming       *Timing                  `json:"occurrenceTiming,omitempty"`
	Performer              []ChargeItemPerformer    `json:"performer,omitempty"`
	PerformingOrganization *common.Reference        `json:"performingOrganization,omitempty"`
	RequestingOrganization *common.Reference        `json:"requestingOrganization,omitempty"`
	CostCenter             *common.Reference        `json:"costCenter,omitempty"`
	Quantity               *common.Quantity         `json:"quantity,omitempty"`
	Bodysite               []common.CodeableConcept `json:"bodysite,omitempty"`
	FactorOverride         *float64                 `json:"factorOverride,omitempty"`
	PriceOverride          *common.Money            `json:"priceOverride,omitempty"`
	OverrideReason         *string                  `json:"overrideReason,omitempty"`
	Enterer                *common.Reference        `json:"enterer,omitempty"`
	EnteredDate            *string                  `json:"enteredDate,omitempty"`
	Reason                 []common.CodeableConcept `json:"reason,omitempty"`
	Service                []common.Reference       `json:"service,omitempty"`
	ProductReference       *common.Reference        `json:"productReference,omitempty"`
	ProductCodeableConcept *common.CodeableConcept  `json:"productCodeableConcept,omitempty"`
	Account                []common.Reference       `json:"account,omitempty"`
	Note                   []common.Annotation      `json:"note,omitempty"`
	SupportingInformation  []common.Reference       `json:"supportingInformation,omitempty"`
}

// ChargeItemPerformer represents a performer of a charge item
type ChargeItemPerformer struct {
	common.BackboneElement

	Function *common.CodeableConcept `json:"function,omitempty"`
	Actor    common.Reference        `json:"actor"`
}

// MarshalJSON marshals the ChargeItem to JSON
func (c *ChargeItem) MarshalJSON() ([]byte, error) {
	type Alias ChargeItem
	return json.Marshal(&struct {
		*Alias
		ResourceType string `json:"resourceType"`
	}{
		Alias:        (*Alias)(c),
		ResourceType: "ChargeItem",
	})
}
