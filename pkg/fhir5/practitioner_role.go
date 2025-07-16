// Package fhir5 contains FHIR R5 resource definitions
package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// PractitionerRole represents roles/organizations the practitioner is associated with (FHIR R5)
type PractitionerRole struct {
	DomainResource

	ResourceType           string                          `json:"resourceType"` // Always "PractitionerRole"
	Identifier             []common.Identifier             `json:"identifier,omitempty"`
	Active                 *bool                           `json:"active,omitempty"`
	Period                 *common.Period                  `json:"period,omitempty"`
	Practitioner           *common.Reference               `json:"practitioner,omitempty"`
	Organization           *common.Reference               `json:"organization,omitempty"`
	Code                   []common.CodeableConcept        `json:"code,omitempty"`
	Specialty              []common.CodeableConcept        `json:"specialty,omitempty"`
	Location               []common.Reference              `json:"location,omitempty"`
	HealthcareService      []common.Reference              `json:"healthcareService,omitempty"`
	Telecom                []ContactPoint                  `json:"telecom,omitempty"`
	AvailableTime          []PractitionerRoleAvailableTime `json:"availableTime,omitempty"`
	NotAvailable           []PractitionerRoleNotAvailable  `json:"notAvailable,omitempty"`
	AvailabilityExceptions *string                         `json:"availabilityExceptions,omitempty"`
	Endpoint               []common.Reference              `json:"endpoint,omitempty"`
}

type PractitionerRoleAvailableTime struct {
	common.BackboneElement
	DaysOfWeek         []string `json:"daysOfWeek,omitempty"`
	AllDay             *bool    `json:"allDay,omitempty"`
	AvailableStartTime *string  `json:"availableStartTime,omitempty"`
	AvailableEndTime   *string  `json:"availableEndTime,omitempty"`
}

type PractitionerRoleNotAvailable struct {
	common.BackboneElement
	Description string         `json:"description"`
	During      *common.Period `json:"during,omitempty"`
}
