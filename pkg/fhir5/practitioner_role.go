package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

type PractitionerRole struct {
	DomainResource

	ResourceType      string                   `json:"resourceType"` // Always "PractitionerRole"
	Active            *bool                    `json:"active,omitempty"`
	Availability      []interface{}            `json:"availability,omitempty"`
	Characteristic    []common.CodeableConcept `json:"characteristic,omitempty"`
	Code              []common.CodeableConcept `json:"code,omitempty"`
	Communication     []common.CodeableConcept `json:"communication,omitempty"`
	Contact           []interface{}            `json:"contact,omitempty"`
	Endpoint          []common.Reference       `json:"endpoint,omitempty"`
	HealthcareService []common.Reference       `json:"healthcareService,omitempty"`
	Identifier        []common.Identifier      `json:"identifier,omitempty"`
	Location          []common.Reference       `json:"location,omitempty"`
	Organization      *common.Reference        `json:"organization,omitempty"`
	Period            *common.Period           `json:"period,omitempty"`
	Practitioner      *common.Reference        `json:"practitioner,omitempty"`
	Specialty         []common.CodeableConcept `json:"specialty,omitempty"`
}
