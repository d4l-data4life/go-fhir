// Package fhir5 contains FHIR R5 resource definitions
package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// ParametersParameter represents a parameter passed to or received from the operation
type ParametersParameter struct {
	common.Element
	Name                     string                  `json:"name"`
	ValueBase64Binary        *string                 `json:"valueBase64Binary,omitempty"`
	ValueBoolean             *bool                   `json:"valueBoolean,omitempty"`
	ValueCanonical           *string                 `json:"valueCanonical,omitempty"`
	ValueCode                *string                 `json:"valueCode,omitempty"`
	ValueDate                *string                 `json:"valueDate,omitempty"`
	ValueDateTime            *string                 `json:"valueDateTime,omitempty"`
	ValueDecimal             *float64                `json:"valueDecimal,omitempty"`
	ValueId                  *string                 `json:"valueId,omitempty"`
	ValueInstant             *string                 `json:"valueInstant,omitempty"`
	ValueInteger             *int                    `json:"valueInteger,omitempty"`
	ValueInteger64           *int64                  `json:"valueInteger64,omitempty"`
	ValueMarkdown            *string                 `json:"valueMarkdown,omitempty"`
	ValueOid                 *string                 `json:"valueOid,omitempty"`
	ValuePositiveInt         *int                    `json:"valuePositiveInt,omitempty"`
	ValueString              *string                 `json:"valueString,omitempty"`
	ValueTime                *string                 `json:"valueTime,omitempty"`
	ValueUnsignedInt         *int                    `json:"valueUnsignedInt,omitempty"`
	ValueUri                 *string                 `json:"valueUri,omitempty"`
	ValueUrl                 *string                 `json:"valueUrl,omitempty"`
	ValueUuid                *string                 `json:"valueUuid,omitempty"`
	ValueAddress             *Address                `json:"valueAddress,omitempty"`
	ValueAge                 *Age                    `json:"valueAge,omitempty"`
	ValueAnnotation          *Annotation             `json:"valueAnnotation,omitempty"`
	ValueAttachment          *Attachment             `json:"valueAttachment,omitempty"`
	ValueCodeableConcept     *common.CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueCoding              *common.Coding          `json:"valueCoding,omitempty"`
	ValueContactPoint        *ContactPoint           `json:"valueContactPoint,omitempty"`
	ValueCount               *Count                  `json:"valueCount,omitempty"`
	ValueDistance            *Distance               `json:"valueDistance,omitempty"`
	ValueDuration            *Duration               `json:"valueDuration,omitempty"`
	ValueHumanName           *HumanName              `json:"valueHumanName,omitempty"`
	ValueIdentifier          *common.Identifier      `json:"valueIdentifier,omitempty"`
	ValueMoney               *Money                  `json:"valueMoney,omitempty"`
	ValuePeriod              *common.Period          `json:"valuePeriod,omitempty"`
	ValueQuantity            *common.Quantity        `json:"valueQuantity,omitempty"`
	ValueRange               *Range                  `json:"valueRange,omitempty"`
	ValueRatio               *Ratio                  `json:"valueRatio,omitempty"`
	ValueReference           *common.Reference       `json:"valueReference,omitempty"`
	ValueSampledData         *interface{}            `json:"valueSampledData,omitempty"`
	ValueSignature           *Signature              `json:"valueSignature,omitempty"`
	ValueTiming              *Timing                 `json:"valueTiming,omitempty"`
	ValueContactDetail       *ContactDetail          `json:"valueContactDetail,omitempty"`
	ValueContributor         *interface{}            `json:"valueContributor,omitempty"`
	ValueDataRequirement     *interface{}            `json:"valueDataRequirement,omitempty"`
	ValueExpression          *Expression             `json:"valueExpression,omitempty"`
	ValueParameterDefinition *interface{}            `json:"valueParameterDefinition,omitempty"`
	ValueRelatedArtifact     *RelatedArtifact        `json:"valueRelatedArtifact,omitempty"`
	ValueTriggerDefinition   *TriggerDefinition      `json:"valueTriggerDefinition,omitempty"`
	ValueUsageContext        *UsageContext           `json:"valueUsageContext,omitempty"`
	ValueDosage              *Dosage                 `json:"valueDosage,omitempty"`
	ValueMeta                *Meta                   `json:"valueMeta,omitempty"`
	Resource                 interface{}             `json:"resource,omitempty"`
	Part                     []ParametersParameter   `json:"part,omitempty"`
}

// Parameters represents a set of parameters
type Parameters struct {
	Resource
	ResourceType string                `json:"resourceType"` // Always "Parameters"
	Parameter    []ParametersParameter `json:"parameter,omitempty"`
}
