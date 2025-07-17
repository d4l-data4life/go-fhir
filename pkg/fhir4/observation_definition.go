package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// ObservationDefinition represents a formal computable definition of an observation
type ObservationDefinition struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "ObservationDefinition"

	// The set of abnormal coded results for the observation conforming to this ObservationDefinition
	AbnormalCodedValueSet *common.Reference `json:"abnormalCodedValueSet,omitempty"`

	// This element allows various categorization schemes based on the owner's definition of the category
	Category []common.CodeableConcept `json:"category,omitempty"`

	// Describes what will be observed. Sometimes this is called the observation "name"
	Code common.CodeableConcept `json:"code"`

	// The set of critical coded results for the observation conforming to this ObservationDefinition
	CriticalCodedValueSet *common.Reference `json:"criticalCodedValueSet,omitempty"`

	// A unique identifier assigned to this ObservationDefinition artifact
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Only used if not implicit in observation code
	Method *common.CodeableConcept `json:"method,omitempty"`

	// An example of observation allowing multiple results is "bacteria identified by culture"
	MultipleResultsAllowed *bool `json:"multipleResultsAllowed,omitempty"`

	// The set of normal coded results for the observations conforming to this ObservationDefinition
	NormalCodedValueSet *common.Reference `json:"normalCodedValueSet,omitempty"`

	// The data types allowed for the value element of the instance observations conforming to this ObservationDefinition
	PermittedDataType []ObservationDefinitionPermittedDataType `json:"permittedDataType,omitempty"`

	// The preferred name to be used when reporting the results of observations conforming to this ObservationDefinition
	PreferredReportName *string `json:"preferredReportName,omitempty"`

	// Multiple ranges of results qualified by different contexts for ordinal or continuous observations
	QualifiedInterval []ObservationDefinitionQualifiedInterval `json:"qualifiedInterval,omitempty"`

	// Characteristics for quantitative results of this observation
	QuantitativeDetails *ObservationDefinitionQuantitativeDetails `json:"quantitativeDetails,omitempty"`

	// The set of valid coded results for the observations conforming to this ObservationDefinition
	ValidCodedValueSet *common.Reference `json:"validCodedValueSet,omitempty"`
}

// ObservationDefinitionPermittedDataType represents the data types allowed for the value element
type ObservationDefinitionPermittedDataType string

const (
	ObservationDefinitionPermittedDataTypeQuantity        ObservationDefinitionPermittedDataType = "Quantity"
	ObservationDefinitionPermittedDataTypeCodeableConcept ObservationDefinitionPermittedDataType = "CodeableConcept"
	ObservationDefinitionPermittedDataTypeString          ObservationDefinitionPermittedDataType = "string"
	ObservationDefinitionPermittedDataTypeBoolean         ObservationDefinitionPermittedDataType = "boolean"
	ObservationDefinitionPermittedDataTypeInteger         ObservationDefinitionPermittedDataType = "integer"
	ObservationDefinitionPermittedDataTypeRange           ObservationDefinitionPermittedDataType = "Range"
	ObservationDefinitionPermittedDataTypeRatio           ObservationDefinitionPermittedDataType = "Ratio"
	ObservationDefinitionPermittedDataTypeSampledData     ObservationDefinitionPermittedDataType = "SampledData"
	ObservationDefinitionPermittedDataTypeTime            ObservationDefinitionPermittedDataType = "time"
	ObservationDefinitionPermittedDataTypeDateTime        ObservationDefinitionPermittedDataType = "dateTime"
	ObservationDefinitionPermittedDataTypePeriod          ObservationDefinitionPermittedDataType = "Period"
)

// ObservationDefinitionQualifiedInterval represents multiple ranges of results qualified by different contexts
type ObservationDefinitionQualifiedInterval struct {
	common.BackboneElement

	// The age at which this reference range is applicable
	Age *common.Range `json:"age,omitempty"`

	// Some analytes vary greatly over age
	AppliesTo []common.CodeableConcept `json:"appliesTo,omitempty"`

	// The category of interval of values for continuous or ordinal observations
	Category *ObservationDefinitionQualifiedIntervalCategory `json:"category,omitempty"`

	// The condition that is associated with the reference range
	Condition *string `json:"condition,omitempty"`

	// Sex of the population the range applies to
	Gender *common.CodeableConcept `json:"gender,omitempty"`

	// The gestational age to which this reference range is applicable
	GestationalAge *common.Range `json:"gestationalAge,omitempty"`

	// The low and high values determining the interval
	Range *common.Range `json:"range,omitempty"`

	// Text based reference range in an observation
	Text *string `json:"text,omitempty"`

	// The type of qualified interval
	Type *ObservationDefinitionQualifiedIntervalType `json:"type,omitempty"`
}

// ObservationDefinitionQualifiedIntervalCategory represents the category of interval of values
type ObservationDefinitionQualifiedIntervalCategory string

const (
	ObservationDefinitionQualifiedIntervalCategoryReference ObservationDefinitionQualifiedIntervalCategory = "reference"
	ObservationDefinitionQualifiedIntervalCategoryCritical  ObservationDefinitionQualifiedIntervalCategory = "critical"
	ObservationDefinitionQualifiedIntervalCategoryAbsolute  ObservationDefinitionQualifiedIntervalCategory = "absolute"
)

// ObservationDefinitionQualifiedIntervalType represents the type of qualified interval
type ObservationDefinitionQualifiedIntervalType string

const (
	ObservationDefinitionQualifiedIntervalTypeValue ObservationDefinitionQualifiedIntervalType = "value"
	ObservationDefinitionQualifiedIntervalTypeRatio ObservationDefinitionQualifiedIntervalType = "ratio"
)

// ObservationDefinitionQuantitativeDetails represents characteristics for quantitative results
type ObservationDefinitionQuantitativeDetails struct {
	common.BackboneElement

	// Customary unit used to report quantitative results of observations conforming to this ObservationDefinition
	CustomaryUnit *common.CodeableConcept `json:"customaryUnit,omitempty"`

	// Factor for converting value expressed with SI unit to value expressed with customary unit
	ConversionFactor *float64 `json:"conversionFactor,omitempty"`

	// Number of digits after decimal separator when the results of such observations are of type Quantity
	DecimalPrecision *int `json:"decimalPrecision,omitempty"`

	// SI unit used to report quantitative results of observations conforming to this ObservationDefinition
	Unit *common.CodeableConcept `json:"unit,omitempty"`
}
