// Package fhir5 contains FHIR R5 resource definitions
package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// RelatedPerson represents a person related to a patient (FHIR R5)
type RelatedPerson struct {
	DomainResource

	ResourceType string                   `json:"resourceType"` // Always "RelatedPerson"
	Identifier   []common.Identifier      `json:"identifier,omitempty"`
	Active       *bool                    `json:"active,omitempty"`
	Patient      common.Reference         `json:"patient"`
	Relationship []common.CodeableConcept `json:"relationship,omitempty"`
	Name         []HumanName              `json:"name,omitempty"`
	Telecom      []ContactPoint           `json:"telecom,omitempty"`
	Gender       *string                  `json:"gender,omitempty"`
	BirthDate    *string                  `json:"birthDate,omitempty"`
	Address      []Address                `json:"address,omitempty"`
	Photo        []Attachment             `json:"photo,omitempty"`
	Period       *common.Period           `json:"period,omitempty"`
}
