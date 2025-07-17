package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// MeasureReport represents the MeasureReport resource contains the results of the calculation of a measure
type MeasureReport struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "MeasureReport"

	// The date this measure report was generated
	Date *string `json:"date,omitempty"`

	// A reference to a Bundle containing the Resources that were used in the calculation of this measure
	EvaluatedResource []common.Reference `json:"evaluatedResource,omitempty"`

	// The results of the calculation, one for each population group in the measure
	Group []MeasureReportGroup `json:"group,omitempty"`

	// Typically, this is used for identifiers that can go in an HL7 V3 II (instance identifier) data type
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// A reference to the Measure that was calculated to produce this report
	Measure string `json:"measure"`

	// The reporting period for which the report was calculated
	Period common.Period `json:"period"`

	// The individual, location, or organization that is reporting the data
	Reporter *common.Reference `json:"reporter,omitempty"`

	// The status of the measure report
	Status MeasureReportStatus `json:"status"`

	// The type of measure report
	Type MeasureReportType `json:"type"`
}

// MeasureReportStatus represents the status of the measure report
type MeasureReportStatus string

const (
	MeasureReportStatusComplete MeasureReportStatus = "complete"
	MeasureReportStatusPending  MeasureReportStatus = "pending"
	MeasureReportStatusError    MeasureReportStatus = "error"
)

// MeasureReportType represents the type of measure report
type MeasureReportType string

const (
	MeasureReportTypeIndividual     MeasureReportType = "individual"
	MeasureReportTypeSubjectList    MeasureReportType = "subject-list"
	MeasureReportTypeSummary        MeasureReportType = "summary"
	MeasureReportTypeDataCollection MeasureReportType = "data-collection"
)

// MeasureReportGroup represents the results of the calculation, one for each population group in the measure
type MeasureReportGroup struct {
	common.BackboneElement

	// The results of the calculation, one for each population group in the measure
	Code *common.CodeableConcept `json:"code,omitempty"`

	// The results of the calculation, one for each population group in the measure
	Identifier *common.Identifier `json:"identifier,omitempty"`

	// The populations that make up the population group, one for each type of population appropriate for the measure
	Population []MeasureReportGroupPopulation `json:"population,omitempty"`

	// The measure score for this population group, calculated as appropriate for the measure type and scoring method
	Score *common.Quantity `json:"score,omitempty"`

	// When a measure includes multiple stratifiers, there will be a stratifier group for each stratifier defined by the measure
	Stratifier []MeasureReportGroupStratifier `json:"stratifier,omitempty"`
}

// MeasureReportGroupPopulation represents the populations that make up the population group
type MeasureReportGroupPopulation struct {
	common.BackboneElement

	// The type of the population
	Code *common.CodeableConcept `json:"code,omitempty"`

	// The number of members of the population
	Count *int `json:"count,omitempty"`

	// The identifier of the population being reported, as defined by the population element of the measure
	Identifier *common.Identifier `json:"identifier,omitempty"`

	// This element refers to a List of patient level MeasureReport resources, one for each patient in this population
	Patients *common.Reference `json:"patients,omitempty"`
}

// MeasureReportGroupStratifier represents when a measure includes multiple stratifiers
type MeasureReportGroupStratifier struct {
	common.BackboneElement

	// The identifier of the stratifier, as defined by the stratifier element of the measure
	Code []common.CodeableConcept `json:"code,omitempty"`

	// This element contains the results for a single stratum within the stratifier
	Stratum []MeasureReportGroupStratifierStratum `json:"stratum,omitempty"`
}

// MeasureReportGroupStratifierStratum represents this element contains the results for a single stratum within the stratifier
type MeasureReportGroupStratifierStratum struct {
	common.BackboneElement

	// The measure score for this stratum, calculated as appropriate for the measure type and scoring method
	Score *common.Quantity `json:"score,omitempty"`

	// The value for this stratum, expressed as a CodeableConcept
	Value *common.CodeableConcept `json:"value,omitempty"`

	// The populations that make up the stratum, one for each type of population appropriate to the measure
	Population []MeasureReportGroupStratifierStratumPopulation `json:"population,omitempty"`

	// A component of the stratifier value for this stratum
	Component []MeasureReportGroupStratifierStratumComponent `json:"component,omitempty"`
}

// MeasureReportGroupStratifierStratumPopulation represents the populations that make up the stratum
type MeasureReportGroupStratifierStratumPopulation struct {
	common.BackboneElement

	// The type of the population
	Code *common.CodeableConcept `json:"code,omitempty"`

	// The number of members of the population in this stratum
	Count *int `json:"count,omitempty"`

	// The identifier of the population being reported, as defined by the population element of the measure
	Identifier *common.Identifier `json:"identifier,omitempty"`

	// This element refers to a List of patient level MeasureReport resources, one for each patient in this population in this stratum
	Patients *common.Reference `json:"patients,omitempty"`
}

// MeasureReportGroupStratifierStratumComponent represents a component of the stratifier value for this stratum
type MeasureReportGroupStratifierStratumComponent struct {
	common.BackboneElement

	// The code for the stratum component value
	Code common.CodeableConcept `json:"code"`

	// The stratum component value
	Value common.CodeableConcept `json:"value"`
}
