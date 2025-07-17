package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// StructureMap represents a Map of relationships between 2 structures that can be used to transform data
type StructureMap struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "StructureMap"

	// May be a web site, an email address, a telephone number, etc
	Contact []common.ContactDetail `json:"contact,omitempty"`

	// A copyright statement relating to the structure map and/or its contents
	Copyright *string `json:"copyright,omitempty"`

	// Note that this is not the same as the resource last-modified-date
	Date *string `json:"date,omitempty"`

	// This description can be used to capture details such as why the structure map was built
	Description *string `json:"description,omitempty"`

	// Allows filtering of structure maps that are appropriate for use versus not
	Experimental *bool `json:"experimental,omitempty"`

	// Organizes the mapping into manageable chunks for human review/ease of maintenance
	Group []StructureMapGroup `json:"group"`

	// Typically, this is used for identifiers that can go in an HL7 V3 II (instance identifier) data type
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Other maps used by this map (canonical URLs)
	Import []string `json:"import,omitempty"`

	// It may be possible for the structure map to be used in jurisdictions other than those for which it was originally designed
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`

	// The name is not expected to be globally unique
	Name string `json:"name"`

	// Usually an organization but may be an individual
	Publisher *string `json:"publisher,omitempty"`

	// This element does not describe the usage of the structure map
	Purpose *string `json:"purpose,omitempty"`

	// Allows filtering of structure maps that are appropriate for use versus not
	Status StructureMapStatus `json:"status"`

	// It is not necessary for a structure map to identify any dependent structures
	Structure []StructureMapStructure `json:"structure,omitempty"`

	// This name does not need to be machine-processing friendly and may contain punctuation, white-space, etc
	Title *string `json:"title,omitempty"`

	// Can be a urn:uuid: or a urn:oid: but real http: addresses are preferred
	Url string `json:"url"`

	// When multiple useContexts are specified, there is no expectation that all or any of the contexts apply
	UseContext []common.UsageContext `json:"useContext,omitempty"`

	// There may be different structure map instances that have the same identifier but different versions
	Version *string `json:"version,omitempty"`
}

// StructureMapStatus represents allows filtering of structure maps that are appropriate for use versus not
type StructureMapStatus string

const (
	StructureMapStatusDraft   StructureMapStatus = "draft"
	StructureMapStatusActive  StructureMapStatus = "active"
	StructureMapStatusRetired StructureMapStatus = "retired"
	StructureMapStatusUnknown StructureMapStatus = "unknown"
)

// StructureMapStructure represents it is not necessary for a structure map to identify any dependent structures
type StructureMapStructure struct {
	common.BackboneElement

	// This is needed if both types have the same name (e.g. version conversion)
	Alias *string `json:"alias,omitempty"`

	// Documentation that describes how the structure is used in the mapping
	Documentation *string `json:"documentation,omitempty"`

	// How the referenced structure is used in this mapping
	Mode StructureMapStructureMode `json:"mode"`

	// The canonical reference to the structure
	Url string `json:"url"`
}

// StructureMapStructureMode represents how the referenced structure is used in this mapping
type StructureMapStructureMode string

const (
	StructureMapStructureModeSource   StructureMapStructureMode = "source"
	StructureMapStructureModeQueried  StructureMapStructureMode = "queried"
	StructureMapStructureModeTarget   StructureMapStructureMode = "target"
	StructureMapStructureModeProduced StructureMapStructureMode = "produced"
)

// StructureMapGroup represents organizes the mapping into manageable chunks for human review/ease of maintenance
type StructureMapGroup struct {
	common.BackboneElement

	// If no inputs are named, then the entry mappings are type based
	Input []StructureMapGroupInput `json:"input"`

	// A unique name for the group for the convenience of human readers
	Name string `json:"name"`

	// Another group that this group adds rules to
	Extends *string `json:"extends,omitempty"`

	// If this is the default rule set to apply for the source type or this combination of types
	TypeMode *StructureMapGroupTypeMode `json:"typeMode,omitempty"`

	// Additional supporting documentation that explains the purpose of the group and the types of mappings within it
	Documentation *string `json:"documentation,omitempty"`

	// Transform Rule from source to target
	Rule []StructureMapGroupRule `json:"rule"`
}

// StructureMapGroupTypeMode represents if this is the default rule set to apply for the source type or this combination of types
type StructureMapGroupTypeMode string

const (
	StructureMapGroupTypeModeNone      StructureMapGroupTypeMode = "none"
	StructureMapGroupTypeModeTypes     StructureMapGroupTypeMode = "types"
	StructureMapGroupTypeModeFirstType StructureMapGroupTypeMode = "first-type"
)

// StructureMapGroupInput represents if no inputs are named, then the entry mappings are type based
type StructureMapGroupInput struct {
	common.BackboneElement

	// Documentation for this instance of data
	Documentation *string `json:"documentation,omitempty"`

	// Mode for this instance of data
	Mode StructureMapGroupInputMode `json:"mode"`

	// Name for this instance of data
	Name string `json:"name"`

	// Type for this instance of data
	Type *string `json:"type,omitempty"`
}

// StructureMapGroupInputMode represents mode for this instance of data
type StructureMapGroupInputMode string

const (
	StructureMapGroupInputModeSource StructureMapGroupInputMode = "source"
	StructureMapGroupInputModeTarget StructureMapGroupInputMode = "target"
)

// StructureMapGroupRule represents transform Rule from source to target
type StructureMapGroupRule struct {
	common.BackboneElement

	// Name of the rule for internal references
	Name string `json:"name"`

	// Source inputs to the mapping
	Source []StructureMapGroupRuleSource `json:"source"`

	// Content to create because of this mapping rule
	Target []StructureMapGroupRuleTarget `json:"target,omitempty"`

	// Rules contained in this rule
	Rule []StructureMapGroupRule `json:"rule,omitempty"`

	// Which other rules to apply in the context of this rule
	Dependent []StructureMapGroupRuleDependent `json:"dependent,omitempty"`

	// Documentation for this instance of data
	Documentation *string `json:"documentation,omitempty"`
}

// StructureMapGroupRuleSource represents source inputs to the mapping
type StructureMapGroupRuleSource struct {
	common.BackboneElement

	// FHIRPath expression - must be true or the mapping engine throws an error instead of completing
	Check *string `json:"check,omitempty"`

	// FHIRPath expression - must be true or the rule does not apply
	Condition *string `json:"condition,omitempty"`

	// Type or variable this rule applies to
	Context string `json:"context"`

	// If there's a default value on an item that can repeat, it will only be used once
	DefaultValueBase64Binary    *string                 `json:"defaultValueBase64Binary,omitempty"`
	DefaultValueBoolean         *bool                   `json:"defaultValueBoolean,omitempty"`
	DefaultValueCanonical       *string                 `json:"defaultValueCanonical,omitempty"`
	DefaultValueCode            *string                 `json:"defaultValueCode,omitempty"`
	DefaultValueDate            *string                 `json:"defaultValueDate,omitempty"`
	DefaultValueDateTime        *string                 `json:"defaultValueDateTime,omitempty"`
	DefaultValueDecimal         *float64                `json:"defaultValueDecimal,omitempty"`
	DefaultValueId              *string                 `json:"defaultValueId,omitempty"`
	DefaultValueInstant         *string                 `json:"defaultValueInstant,omitempty"`
	DefaultValueInteger         *int                    `json:"defaultValueInteger,omitempty"`
	DefaultValueMarkdown        *string                 `json:"defaultValueMarkdown,omitempty"`
	DefaultValueOid             *string                 `json:"defaultValueOid,omitempty"`
	DefaultValuePositiveInt     *int                    `json:"defaultValuePositiveInt,omitempty"`
	DefaultValueString          *string                 `json:"defaultValueString,omitempty"`
	DefaultValueTime            *string                 `json:"defaultValueTime,omitempty"`
	DefaultValueUnsignedInt     *int                    `json:"defaultValueUnsignedInt,omitempty"`
	DefaultValueUri             *string                 `json:"defaultValueUri,omitempty"`
	DefaultValueUrl             *string                 `json:"defaultValueUrl,omitempty"`
	DefaultValueUuid            *string                 `json:"defaultValueUuid,omitempty"`
	DefaultValueAddress         *common.Address         `json:"defaultValueAddress,omitempty"`
	DefaultValueAge             *common.Age             `json:"defaultValueAge,omitempty"`
	DefaultValueAnnotation      *common.Annotation      `json:"defaultValueAnnotation,omitempty"`
	DefaultValueAttachment      *common.Attachment      `json:"defaultValueAttachment,omitempty"`
	DefaultValueCodeableConcept *common.CodeableConcept `json:"defaultValueCodeableConcept,omitempty"`
	DefaultValueCoding          *common.Coding          `json:"defaultValueCoding,omitempty"`
	DefaultValueContactPoint    *common.ContactPoint    `json:"defaultValueContactPoint,omitempty"`
	DefaultValueCount           *common.Count           `json:"defaultValueCount,omitempty"`
	DefaultValueDistance        *common.Distance        `json:"defaultValueDistance,omitempty"`
	DefaultValueDuration        *common.Duration        `json:"defaultValueDuration,omitempty"`
	DefaultValueHumanName       *common.HumanName       `json:"defaultValueHumanName,omitempty"`
	DefaultValueIdentifier      *common.Identifier      `json:"defaultValueIdentifier,omitempty"`
	DefaultValueMoney           *common.Money           `json:"defaultValueMoney,omitempty"`
	DefaultValuePeriod          *common.Period          `json:"defaultValuePeriod,omitempty"`
	DefaultValueQuantity        *common.Quantity        `json:"defaultValueQuantity,omitempty"`
	DefaultValueRange           *common.Range           `json:"defaultValueRange,omitempty"`
	DefaultValueRatio           *common.Ratio           `json:"defaultValueRatio,omitempty"`
	DefaultValueReference       *common.Reference       `json:"defaultValueReference,omitempty"`
	DefaultValueSampledData     *common.SampledData     `json:"defaultValueSampledData,omitempty"`
	DefaultValueSignature       *common.Signature       `json:"defaultValueSignature,omitempty"`
	DefaultValueTiming          *common.Timing          `json:"defaultValueTiming,omitempty"`
	DefaultValueMeta            *common.Meta            `json:"defaultValueMeta,omitempty"`

	// Specified type for the element
	Element *string `json:"element,omitempty"`

	// How to handle the list mode for this element
	ListMode *StructureMapGroupRuleSourceListMode `json:"listMode,omitempty"`

	// Named context for field, if a field is specified
	Variable *string `json:"variable,omitempty"`

	// FHIRPath expression to specify the first element
	First *string `json:"first,omitempty"`

	// FHIRPath expression to specify the last element
	Last *string `json:"last,omitempty"`

	// FHIRPath expression to specify the last element
	LogMessage *string `json:"logMessage,omitempty"`
}

// StructureMapGroupRuleSourceListMode represents how to handle the list mode for this element
type StructureMapGroupRuleSourceListMode string

const (
	StructureMapGroupRuleSourceListModeFirst    StructureMapGroupRuleSourceListMode = "first"
	StructureMapGroupRuleSourceListModeNotFirst StructureMapGroupRuleSourceListMode = "not_first"
	StructureMapGroupRuleSourceListModeLast     StructureMapGroupRuleSourceListMode = "last"
	StructureMapGroupRuleSourceListModeNotLast  StructureMapGroupRuleSourceListMode = "not_last"
	StructureMapGroupRuleSourceListModeOnlyOne  StructureMapGroupRuleSourceListMode = "only_one"
)

// StructureMapGroupRuleTarget represents content to create because of this mapping rule
type StructureMapGroupRuleTarget struct {
	common.BackboneElement

	// Type or variable this rule applies to
	Context *string `json:"context,omitempty"`

	// How to interpret the context
	ContextType *StructureMapGroupRuleTargetContextType `json:"contextType,omitempty"`

	// Field to create in the context
	Element *string `json:"element,omitempty"`

	// Named context for field, if desired, and a field is specified
	Variable *string `json:"variable,omitempty"`

	// If field is a list, how to manage the list
	ListMode []StructureMapGroupRuleTargetListMode `json:"listMode,omitempty"`

	// Internal rule reference for shared list items
	ListRuleId *string `json:"listRuleId,omitempty"`

	// Parameters to the transform
	Transform *StructureMapGroupRuleTargetTransform `json:"transform,omitempty"`

	// Parameters to the transform
	Parameter []StructureMapGroupRuleTargetParameter `json:"parameter,omitempty"`
}

// StructureMapGroupRuleTargetContextType represents how to interpret the context
type StructureMapGroupRuleTargetContextType string

const (
	StructureMapGroupRuleTargetContextTypeVariable StructureMapGroupRuleTargetContextType = "variable"
	StructureMapGroupRuleTargetContextTypeType     StructureMapGroupRuleTargetContextType = "type"
)

// StructureMapGroupRuleTargetListMode represents if field is a list, how to manage the list
type StructureMapGroupRuleTargetListMode string

const (
	StructureMapGroupRuleTargetListModeCollate StructureMapGroupRuleTargetListMode = "collate"
	StructureMapGroupRuleTargetListModeShare   StructureMapGroupRuleTargetListMode = "share"
)

// StructureMapGroupRuleTargetTransform represents parameters to the transform
type StructureMapGroupRuleTargetTransform string

const (
	StructureMapGroupRuleTargetTransformCreate    StructureMapGroupRuleTargetTransform = "create"
	StructureMapGroupRuleTargetTransformCopy      StructureMapGroupRuleTargetTransform = "copy"
	StructureMapGroupRuleTargetTransformTruncate  StructureMapGroupRuleTargetTransform = "truncate"
	StructureMapGroupRuleTargetTransformEscape    StructureMapGroupRuleTargetTransform = "escape"
	StructureMapGroupRuleTargetTransformCast      StructureMapGroupRuleTargetTransform = "cast"
	StructureMapGroupRuleTargetTransformAppend    StructureMapGroupRuleTargetTransform = "append"
	StructureMapGroupRuleTargetTransformTranslate StructureMapGroupRuleTargetTransform = "translate"
	StructureMapGroupRuleTargetTransformReference StructureMapGroupRuleTargetTransform = "reference"
	StructureMapGroupRuleTargetTransformDateOp    StructureMapGroupRuleTargetTransform = "dateOp"
	StructureMapGroupRuleTargetTransformUuid      StructureMapGroupRuleTargetTransform = "uuid"
	StructureMapGroupRuleTargetTransformPointer   StructureMapGroupRuleTargetTransform = "pointer"
	StructureMapGroupRuleTargetTransformEvaluate  StructureMapGroupRuleTargetTransform = "evaluate"
	StructureMapGroupRuleTargetTransformCc        StructureMapGroupRuleTargetTransform = "cc"
	StructureMapGroupRuleTargetTransformC         StructureMapGroupRuleTargetTransform = "c"
	StructureMapGroupRuleTargetTransformQty       StructureMapGroupRuleTargetTransform = "qty"
	StructureMapGroupRuleTargetTransformId        StructureMapGroupRuleTargetTransform = "id"
	StructureMapGroupRuleTargetTransformCp        StructureMapGroupRuleTargetTransform = "cp"
)

// StructureMapGroupRuleTargetParameter represents parameters to the transform
type StructureMapGroupRuleTargetParameter struct {
	common.BackboneElement

	// Parameter value - variable or literal
	ValueId      *string  `json:"valueId,omitempty"`
	ValueString  *string  `json:"valueString,omitempty"`
	ValueBoolean *bool    `json:"valueBoolean,omitempty"`
	ValueInteger *int     `json:"valueInteger,omitempty"`
	ValueDecimal *float64 `json:"valueDecimal,omitempty"`
}

// StructureMapGroupRuleDependent represents which other rules to apply in the context of this rule
type StructureMapGroupRuleDependent struct {
	common.BackboneElement

	// Name of a rule or group to apply
	Name string `json:"name"`

	// Variable to pass to the rule or group
	Variable []string `json:"variable"`
}
