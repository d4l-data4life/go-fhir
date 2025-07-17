// Package fhir5 contains FHIR R5 resource definitions
package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// RegulatedAuthorization represents regulatory approval for a product (FHIR R5)
type RegulatedAuthorization struct {
	DomainResource

	ResourceType   string                      `json:"resourceType"` // Always "RegulatedAuthorization"
	Identifier     []common.Identifier         `json:"identifier,omitempty"`
	Subject        *common.Reference           `json:"subject,omitempty"`
	Type           *common.CodeableConcept     `json:"type,omitempty"`
	Description    *string                     `json:"description,omitempty"`
	Region         []common.CodeableConcept    `json:"region,omitempty"`
	Status         *common.CodeableConcept     `json:"status,omitempty"`
	StatusDate     *string                     `json:"statusDate,omitempty"`
	ValidityPeriod *common.Period              `json:"validityPeriod,omitempty"`
	Indication     *CodeableReference          `json:"indication,omitempty"`
	IntendedUse    *common.CodeableConcept     `json:"intendedUse,omitempty"`
	Basis          []common.CodeableConcept    `json:"basis,omitempty"`
	Holder         *common.Reference           `json:"holder,omitempty"`
	Regulator      *common.Reference           `json:"regulator,omitempty"`
	Case           *RegulatedAuthorizationCase `json:"case,omitempty"`
}

type RegulatedAuthorizationCase struct {
	common.BackboneElement
	Identifier   *common.Identifier      `json:"identifier,omitempty"`
	Type         *common.CodeableConcept `json:"type,omitempty"`
	Status       *common.CodeableConcept `json:"status,omitempty"`
	DatePeriod   *common.Period          `json:"datePeriod,omitempty"`
	DateDateTime *string                 `json:"dateDateTime,omitempty"`
}
