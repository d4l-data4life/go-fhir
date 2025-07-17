package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// ImmunizationRecommendation represents a patient's point-in-time set of recommendations
type ImmunizationRecommendation struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "ImmunizationRecommendation"

	// Indicates the authority who published the protocol (e.g. ACIP)
	Authority *common.Reference `json:"authority,omitempty"`

	// The date the immunization recommendation(s) were created
	Date string `json:"date"`

	// A unique identifier assigned to this particular recommendation record
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The patient the recommendation(s) are for
	Patient common.Reference `json:"patient"`

	// Vaccine administration recommendations
	Recommendation []ImmunizationRecommendationRecommendation `json:"recommendation"`
}

// ImmunizationRecommendationRecommendation represents vaccine administration recommendations
type ImmunizationRecommendationRecommendation struct {
	common.BackboneElement

	// Vaccine(s) which should not be used to fulfill the recommendation
	ContraindicatedVaccineCode []common.CodeableConcept `json:"contraindicatedVaccineCode,omitempty"`

	// Vaccine date recommendations. For example, earliest date to administer, latest date to administer, etc
	DateCriterion []ImmunizationRecommendationRecommendationDateCriterion `json:"dateCriterion,omitempty"`

	// Contains the description about the protocol under which the vaccine was administered
	Description *string `json:"description,omitempty"`

	// The use of an integer is preferred if known
	DoseNumberPositiveInt *int `json:"doseNumberPositiveInt,omitempty"`

	// The use of an integer is preferred if known
	DoseNumberString *string `json:"doseNumberString,omitempty"`

	// The reason for the assigned forecast status
	ForecastReason []common.CodeableConcept `json:"forecastReason,omitempty"`

	// Indicates the patient status with respect to the path to immunity for the target disease
	ForecastStatus common.CodeableConcept `json:"forecastStatus"`

	// One possible path to achieve presumed immunity against a disease - within the context of an authority
	Series *string `json:"series,omitempty"`

	// The use of an integer is preferred if known
	SeriesDosesPositiveInt *int `json:"seriesDosesPositiveInt,omitempty"`

	// The use of an integer is preferred if known
	SeriesDosesString *string `json:"seriesDosesString,omitempty"`

	// Immunization event history and/or evaluation that supports the status and recommendation
	SupportingImmunization []common.Reference `json:"supportingImmunization,omitempty"`

	// Patient Information that supports the status and recommendation
	SupportingPatientInformation []common.Reference `json:"supportingPatientInformation,omitempty"`

	// The targeted disease for the recommendation
	TargetDisease *common.CodeableConcept `json:"targetDisease,omitempty"`

	// Vaccine(s) or vaccine group that pertain to the recommendation
	VaccineCode []common.CodeableConcept `json:"vaccineCode,omitempty"`
}

// ImmunizationRecommendationRecommendationDateCriterion represents vaccine date recommendations
type ImmunizationRecommendationRecommendationDateCriterion struct {
	common.BackboneElement

	// Date classification of recommendation. For example, earliest date to give, latest date to give, etc
	Code common.CodeableConcept `json:"code"`

	// The date whose meaning is specified by dateCriterion.code
	Value string `json:"value"`
}
