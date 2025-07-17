package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// MeasureReportStatus represents the status of a measure report
type MeasureReportStatus string

const (
	MeasureReportStatusComplete MeasureReportStatus = "complete"
	MeasureReportStatusPending  MeasureReportStatus = "pending"
	MeasureReportStatusError    MeasureReportStatus = "error"
)

// MeasureReportType represents the type of measure report
type MeasureReportType string

const (
	MeasureReportTypeIndividual   MeasureReportType = "individual"
	MeasureReportTypeSubjectList  MeasureReportType = "subject-list"
	MeasureReportTypeSummary      MeasureReportType = "summary"
	MeasureReportTypeDataExchange MeasureReportType = "data-exchange"
)

// MeasureReportDataUpdateType represents the data update type
type MeasureReportDataUpdateType string

const (
	MeasureReportDataUpdateTypeIncremental MeasureReportDataUpdateType = "incremental"
	MeasureReportDataUpdateTypeSnapshot    MeasureReportDataUpdateType = "snapshot"
)

// MeasureReportGroupPopulation represents a population in a measure report group
type MeasureReportGroupPopulation struct {
	common.BackboneElement

	// The code for the population criteria
	Code *common.CodeableConcept `json:"code,omitempty"`

	// The number of members of the population
	Count *int `json:"count,omitempty"`

	// A unique identifier for the population criteria
	LinkId *string `json:"linkId,omitempty"`

	// Optional subject identifying the individual or individuals the report is for
	SubjectReference []common.Reference `json:"subjectReference,omitempty"`

	// Optional result references if the measure result is computed for the individual
	SubjectResults []common.Reference `json:"subjectResults,omitempty"`

	// Optional subject identifying the list of subjects the report is for
	Subjects *common.Reference `json:"subjects,omitempty"`
}

// MeasureReportGroupStratifierStratum represents a stratum in a stratifier
type MeasureReportGroupStratifierStratum struct {
	common.BackboneElement

	// A component of this stratifier stratum
	Component []MeasureReportGroupStratifierStratumComponent `json:"component,omitempty"`

	// The measure score for this stratum
	MeasureScoreQuantity        *common.Quantity        `json:"measureScoreQuantity,omitempty"`
	MeasureScoreDateTime        *string                 `json:"measureScoreDateTime,omitempty"`
	MeasureScoreCodeableConcept *common.CodeableConcept `json:"measureScoreCodeableConcept,omitempty"`
	MeasureScorePeriod          *common.Period          `json:"measureScorePeriod,omitempty"`
	MeasureScoreRange           *Range                  `json:"measureScoreRange,omitempty"`
	MeasureScoreDuration        *Duration               `json:"measureScoreDuration,omitempty"`

	// A population criteria for the measure report
	Population []MeasureReportGroupPopulation `json:"population,omitempty"`

	// The value for this stratum, expressed as a CodeableConcept
	ValueCodeableConcept *common.CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueBoolean         *bool                   `json:"valueBoolean,omitempty"`
	ValueQuantity        *common.Quantity        `json:"valueQuantity,omitempty"`
	ValueRange           *Range                  `json:"valueRange,omitempty"`
	ValueReference       *common.Reference       `json:"valueReference,omitempty"`
}

// MeasureReportGroupStratifierStratumComponent represents a component of a stratifier stratum
type MeasureReportGroupStratifierStratumComponent struct {
	common.BackboneElement

	// The code for the stratifier component
	Code *common.CodeableConcept `json:"code,omitempty"`

	// A unique identifier for the stratifier component
	LinkId *string `json:"linkId,omitempty"`

	// The value for this component, expressed as a CodeableConcept
	ValueCodeableConcept *common.CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueBoolean         *bool                   `json:"valueBoolean,omitempty"`
	ValueQuantity        *common.Quantity        `json:"valueQuantity,omitempty"`
	ValueRange           *Range                  `json:"valueRange,omitempty"`
	ValueReference       *common.Reference       `json:"valueReference,omitempty"`
}

// MeasureReportGroupStratifier represents a stratifier for a measure report group
type MeasureReportGroupStratifier struct {
	common.BackboneElement

	// The code for this stratifier
	Code *common.CodeableConcept `json:"code,omitempty"`

	// A unique identifier for the stratifier
	LinkId *string `json:"linkId,omitempty"`

	// Stratifier Results
	Stratum []MeasureReportGroupStratifierStratum `json:"stratum,omitempty"`
}

// MeasureReportGroup represents a group of populations for a measure report
type MeasureReportGroup struct {
	common.BackboneElement

	// Meaning of the group
	Code *common.CodeableConcept `json:"code,omitempty"`

	// A unique identifier for the group
	LinkId *string `json:"linkId,omitempty"`

	// The measure score for this population group
	MeasureScoreQuantity        *common.Quantity        `json:"measureScoreQuantity,omitempty"`
	MeasureScoreDateTime        *string                 `json:"measureScoreDateTime,omitempty"`
	MeasureScoreCodeableConcept *common.CodeableConcept `json:"measureScoreCodeableConcept,omitempty"`
	MeasureScorePeriod          *common.Period          `json:"measureScorePeriod,omitempty"`
	MeasureScoreRange           *Range                  `json:"measureScoreRange,omitempty"`
	MeasureScoreDuration        *Duration               `json:"measureScoreDuration,omitempty"`

	// A population criteria for the measure report
	Population []MeasureReportGroupPopulation `json:"population,omitempty"`

	// The stratifier criteria for the measure report
	Stratifier []MeasureReportGroupStratifier `json:"stratifier,omitempty"`
}

// MeasureReport represents the results of the calculation of a measure
type MeasureReport struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "MeasureReport"

	// This element only applies to Data-collection reports
	DataUpdateType *MeasureReportDataUpdateType `json:"dataUpdateType,omitempty"`

	// The date this measure was calculated
	Date *string `json:"date,omitempty"`

	// Evaluated resources are only reported for individual reports
	EvaluatedResource []common.Reference `json:"evaluatedResource,omitempty"`

	// The results of the calculation, one for each population group in the measure
	Group []MeasureReportGroup `json:"group,omitempty"`

	// Typically, this is used for identifiers that can go in an HL7 V3 II data type
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Information about whether an improvement in the measure is noted by an increase or decrease
	ImprovementNotation *common.CodeableConcept `json:"improvementNotation,omitempty"`

	// Input parameters from operation
	InputParameters *common.Reference `json:"inputParameters,omitempty"`

	// A reference to the location for which the data is being reported
	Location *common.Reference `json:"location,omitempty"`

	// A reference to the Measure that was calculated to produce this report
	Measure *string `json:"measure,omitempty"`

	// The reporting period for which the report was calculated
	Period common.Period `json:"period"`

	// The individual or organization that is reporting the data
	Reporter *common.Reference `json:"reporter,omitempty"`

	// A reference to the vendor who queried the data, calculated results and/or generated the report
	ReportingVendor *common.Reference `json:"reportingVendor,omitempty"`

	// Indicates how the calculation is performed for the measure
	Scoring *common.CodeableConcept `json:"scoring,omitempty"`

	// This element is labeled as a modifier because the status contains codes that mark the resource as not currently valid
	Status MeasureReportStatus `json:"status"`

	// Optional subject identifying the individual or individuals the report is for
	Subject *common.Reference `json:"subject,omitempty"`

	// For individual measure reports, the supplementalData elements represent the direct result of evaluating the supplementalData expression
	SupplementalData []common.Reference `json:"supplementalData,omitempty"`

	// Data-exchange reports are used only to communicate data-of-interest for a measure
	Type MeasureReportType `json:"type"`
}
