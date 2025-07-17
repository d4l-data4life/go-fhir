package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// Parameters represents a parameter passed to or received from the operation
type Parameters struct {
	Resource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Parameters"

	// A parameter passed to or received from the operation
	Parameter []ParametersParameter `json:"parameter,omitempty"`
}

// ParametersParameter represents a parameter passed to or received from the operation
type ParametersParameter struct {
	common.BackboneElement

	// The name of the parameter (reference to the operation definition)
	Name string `json:"name"`

	// Only one level of nested parameters is allowed
	Part []ParametersParameter `json:"part,omitempty"`

	// When resolving references in resources, the operation definition may specify how references may be resolved between parameters
	Resource interface{} `json:"resource,omitempty"`

	// If the parameter is a data type
	ValueBase64Binary *string `json:"valueBase64Binary,omitempty"`

	// If the parameter is a data type
	ValueBoolean *bool `json:"valueBoolean,omitempty"`

	// If the parameter is a data type
	ValueCanonical *string `json:"valueCanonical,omitempty"`

	// If the parameter is a data type
	ValueCode *string `json:"valueCode,omitempty"`

	// If the parameter is a data type
	ValueDate *string `json:"valueDate,omitempty"`

	// If the parameter is a data type
	ValueDateTime *string `json:"valueDateTime,omitempty"`

	// If the parameter is a data type
	ValueDecimal *float64 `json:"valueDecimal,omitempty"`

	// If the parameter is a data type
	ValueId *string `json:"valueId,omitempty"`

	// If the parameter is a data type
	ValueInstant *string `json:"valueInstant,omitempty"`

	// If the parameter is a data type
	ValueInteger *int `json:"valueInteger,omitempty"`

	// If the parameter is a data type
	ValueMarkdown *string `json:"valueMarkdown,omitempty"`

	// If the parameter is a data type
	ValueOid *string `json:"valueOid,omitempty"`

	// If the parameter is a data type
	ValuePositiveInt *int `json:"valuePositiveInt,omitempty"`

	// If the parameter is a data type
	ValueString *string `json:"valueString,omitempty"`

	// If the parameter is a data type
	ValueTime *string `json:"valueTime,omitempty"`

	// If the parameter is a data type
	ValueUnsignedInt *int `json:"valueUnsignedInt,omitempty"`

	// If the parameter is a data type
	ValueUri *string `json:"valueUri,omitempty"`

	// If the parameter is a data type
	ValueUrl *string `json:"valueUrl,omitempty"`

	// If the parameter is a data type
	ValueUuid *string `json:"valueUuid,omitempty"`

	// If the parameter is a data type
	ValueAddress *common.Address `json:"valueAddress,omitempty"`

	// If the parameter is a data type
	ValueAge *common.Age `json:"valueAge,omitempty"`

	// If the parameter is a data type
	ValueAnnotation *common.Annotation `json:"valueAnnotation,omitempty"`

	// If the parameter is a data type
	ValueAttachment *common.Attachment `json:"valueAttachment,omitempty"`

	// If the parameter is a data type
	ValueCodeableConcept *common.CodeableConcept `json:"valueCodeableConcept,omitempty"`

	// If the parameter is a data type
	ValueCoding *common.Coding `json:"valueCoding,omitempty"`

	// If the parameter is a data type
	ValueContactPoint *common.ContactPoint `json:"valueContactPoint,omitempty"`

	// If the parameter is a data type
	ValueCount *common.Count `json:"valueCount,omitempty"`

	// If the parameter is a data type
	ValueDistance *common.Distance `json:"valueDistance,omitempty"`

	// If the parameter is a data type
	ValueDuration *common.Duration `json:"valueDuration,omitempty"`

	// If the parameter is a data type
	ValueHumanName *common.HumanName `json:"valueHumanName,omitempty"`

	// If the parameter is a data type
	ValueIdentifier *common.Identifier `json:"valueIdentifier,omitempty"`

	// If the parameter is a data type
	ValueMoney *common.Money `json:"valueMoney,omitempty"`

	// If the parameter is a data type
	ValuePeriod *common.Period `json:"valuePeriod,omitempty"`

	// If the parameter is a data type
	ValueQuantity *common.Quantity `json:"valueQuantity,omitempty"`

	// If the parameter is a data type
	ValueRange *common.Range `json:"valueRange,omitempty"`

	// If the parameter is a data type
	ValueRatio *common.Ratio `json:"valueRatio,omitempty"`

	// If the parameter is a data type
	ValueReference *common.Reference `json:"valueReference,omitempty"`

	// If the parameter is a data type
	ValueSampledData *common.SampledData `json:"valueSampledData,omitempty"`

	// If the parameter is a data type
	ValueSignature *common.Signature `json:"valueSignature,omitempty"`

	// If the parameter is a data type
	ValueTiming *common.Timing `json:"valueTiming,omitempty"`

	// If the parameter is a data type
	ValueDosage *common.Dosage `json:"valueDosage,omitempty"`

	// If the parameter is a data type
	ValueMeta *common.Meta `json:"valueMeta,omitempty"`
}
