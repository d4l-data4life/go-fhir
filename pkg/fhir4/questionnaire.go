package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// Questionnaire represents a structured set of questions intended to guide the collection of answers from end-users
type Questionnaire struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Questionnaire"

	// The 'date' element may be more recent than the approval date because of minor changes or editorial corrections
	ApprovalDate *string `json:"approvalDate,omitempty"`

	// An identifier for this question or group of questions in a particular terminology such as LOINC
	Code []common.Coding `json:"code,omitempty"`

	// May be a web site, an email address, a telephone number, etc
	Contact []common.ContactDetail `json:"contact,omitempty"`

	// A copyright statement relating to the questionnaire and/or its contents
	Copyright *string `json:"copyright,omitempty"`

	// Note that this is not the same as the resource last-modified-date
	Date *string `json:"date,omitempty"`

	// The URL of a Questionnaire that this Questionnaire is based on
	DerivedFrom []string `json:"derivedFrom,omitempty"`

	// This description can be used to capture details such as why the questionnaire was built
	Description *string `json:"description,omitempty"`

	// The effective period for a questionnaire determines when the content is applicable for usage
	EffectivePeriod *common.Period `json:"effectivePeriod,omitempty"`

	// Allows filtering of questionnaires that are appropriate for use versus not
	Experimental *bool `json:"experimental,omitempty"`

	// Typically, this is used for identifiers that can go in an HL7 V3 II (instance identifier) data type
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The content of the questionnaire is constructed from an ordered, hierarchical collection of items
	Item []QuestionnaireItem `json:"item,omitempty"`

	// It may be possible for the questionnaire to be used in jurisdictions other than those for which it was originally designed
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`

	// If specified, this date follows the original approval date
	LastReviewDate *string `json:"lastReviewDate,omitempty"`

	// The name is not expected to be globally unique
	Name *string `json:"name,omitempty"`

	// Usually an organization but may be an individual
	Publisher *string `json:"publisher,omitempty"`

	// This element does not describe the usage of the questionnaire
	Purpose *string `json:"purpose,omitempty"`

	// Allows filtering of questionnaires that are appropriate for use versus not
	Status QuestionnaireStatus `json:"status"`

	// If none are specified, then the subject is unlimited
	SubjectType []string `json:"subjectType,omitempty"`

	// This name does not need to be machine-processing friendly and may contain punctuation, white-space, etc
	Title *string `json:"title,omitempty"`

	// The name of the referenced questionnaire can be conveyed using the http://hl7.org/fhir/StructureDefinition/display extension
	Url *string `json:"url,omitempty"`

	// When multiple useContexts are specified, there is no expectation that all or any of the contexts apply
	UseContext []common.UsageContext `json:"useContext,omitempty"`

	// There may be different questionnaire instances that have the same identifier but different versions
	Version *string `json:"version,omitempty"`
}

// QuestionnaireStatus represents allows filtering of questionnaires that are appropriate for use versus not
type QuestionnaireStatus string

const (
	QuestionnaireStatusDraft   QuestionnaireStatus = "draft"
	QuestionnaireStatusActive  QuestionnaireStatus = "active"
	QuestionnaireStatusRetired QuestionnaireStatus = "retired"
	QuestionnaireStatusUnknown QuestionnaireStatus = "unknown"
)

// QuestionnaireItem represents the content of the questionnaire is constructed from an ordered, hierarchical collection of items
type QuestionnaireItem struct {
	common.BackboneElement

	// This element can be used when the value set machinery of answerValueSet is deemed too cumbersome
	AnswerOption []QuestionnaireItemAnswerOption `json:"answerOption,omitempty"`

	// LOINC defines many useful value sets for questionnaire responses
	AnswerValueSet *string `json:"answerValueSet,omitempty"`

	// The value may come from the ElementDefinition referred to by .definition
	Code []common.Coding `json:"code,omitempty"`

	// The uri refers to an ElementDefinition in a StructureDefinition
	Definition *string `json:"definition,omitempty"`

	// This element must be specified if more than one enableWhen value is provided
	EnableBehavior *QuestionnaireItemEnableBehavior `json:"enableBehavior,omitempty"`

	// If multiple repetitions of this extension are present, the item should be enabled when the condition for *any* of the repetitions is true
	EnableWhen []QuestionnaireItemEnableWhen `json:"enableWhen,omitempty"`

	// The user is allowed to change the value and override the default (unless marked as read-only)
	Initial []QuestionnaireItemInitial `json:"initial,omitempty"`

	// There is no specified limit to the depth of nesting
	Item []QuestionnaireItem `json:"item,omitempty"`

	// This ''can'' be a meaningful identifier (e.g. a LOINC code) but is not intended to have any meaning
	LinkId string `json:"linkId"`

	// For base64binary, reflects the number of characters representing the encoded data, not the number of bytes of the binary data
	MaxLength *int `json:"maxLength,omitempty"`

	// These are generally unique within a questionnaire, though this is not guaranteed
	Prefix *string `json:"prefix,omitempty"`

	// The value of readOnly elements can be established by asserting extensions for defaultValues
	ReadOnly *bool `json:"readOnly,omitempty"`

	// If a question is marked as repeats=true, then multiple answers can be provided for the question
	Repeats *bool `json:"repeats,omitempty"`

	// Questionnaire.item.required only has meaning for elements that are conditionally enabled with enableWhen
	Required *bool `json:"required,omitempty"`

	// When using this element to represent the name of a section, use group type item
	Text *string `json:"text,omitempty"`

	// Additional constraints on the type of answer can be conveyed by extensions
	Type QuestionnaireItemType `json:"type"`
}

// QuestionnaireItemType represents additional constraints on the type of answer can be conveyed by extensions
type QuestionnaireItemType string

const (
	QuestionnaireItemTypeGroup      QuestionnaireItemType = "group"
	QuestionnaireItemTypeDisplay    QuestionnaireItemType = "display"
	QuestionnaireItemTypeQuestion   QuestionnaireItemType = "question"
	QuestionnaireItemTypeBoolean    QuestionnaireItemType = "boolean"
	QuestionnaireItemTypeDecimal    QuestionnaireItemType = "decimal"
	QuestionnaireItemTypeInteger    QuestionnaireItemType = "integer"
	QuestionnaireItemTypeDate       QuestionnaireItemType = "date"
	QuestionnaireItemTypeDateTime   QuestionnaireItemType = "dateTime"
	QuestionnaireItemTypeTime       QuestionnaireItemType = "time"
	QuestionnaireItemTypeString     QuestionnaireItemType = "string"
	QuestionnaireItemTypeText       QuestionnaireItemType = "text"
	QuestionnaireItemTypeUrl        QuestionnaireItemType = "url"
	QuestionnaireItemTypeChoice     QuestionnaireItemType = "choice"
	QuestionnaireItemTypeOpenChoice QuestionnaireItemType = "open-choice"
	QuestionnaireItemTypeAttachment QuestionnaireItemType = "attachment"
	QuestionnaireItemTypeReference  QuestionnaireItemType = "reference"
	QuestionnaireItemTypeQuantity   QuestionnaireItemType = "quantity"
)

// QuestionnaireItemEnableBehavior represents this element must be specified if more than one enableWhen value is provided
type QuestionnaireItemEnableBehavior string

const (
	QuestionnaireItemEnableBehaviorAll QuestionnaireItemEnableBehavior = "all"
	QuestionnaireItemEnableBehaviorAny QuestionnaireItemEnableBehavior = "any"
)

// QuestionnaireItemEnableWhen represents if multiple repetitions of this extension are present, the item should be enabled when the condition for *any* of the repetitions is true
type QuestionnaireItemEnableWhen struct {
	common.BackboneElement

	// A value that the referenced question is tested using the specified operator in order for the item to be enabled
	AnswerBoolean *bool `json:"answerBoolean,omitempty"`

	// A value that the referenced question is tested using the specified operator in order for the item to be enabled
	AnswerDecimal *float64 `json:"answerDecimal,omitempty"`

	// A value that the referenced question is tested using the specified operator in order for the item to be enabled
	AnswerInteger *int `json:"answerInteger,omitempty"`

	// A value that the referenced question is tested using the specified operator in order for the item to be enabled
	AnswerDate *string `json:"answerDate,omitempty"`

	// A value that the referenced question is tested using the specified operator in order for the item to be enabled
	AnswerDateTime *string `json:"answerDateTime,omitempty"`

	// A value that the referenced question is tested using the specified operator in order for the item to be enabled
	AnswerTime *string `json:"answerTime,omitempty"`

	// A value that the referenced question is tested using the specified operator in order for the item to be enabled
	AnswerString *string `json:"answerString,omitempty"`

	// A value that the referenced question is tested using the specified operator in order for the item to be enabled
	AnswerCoding *common.Coding `json:"answerCoding,omitempty"`

	// A value that the referenced question is tested using the specified operator in order for the item to be enabled
	AnswerQuantity *common.Quantity `json:"answerQuantity,omitempty"`

	// A value that the referenced question is tested using the specified operator in order for the item to be enabled
	AnswerReference *common.Reference `json:"answerReference,omitempty"`

	// Specifies the criteria by which the question is enabled
	Operator QuestionnaireItemEnableWhenOperator `json:"operator"`

	// If multiple question occurrences are present for the same question (same linkId), then this refers to the nearest question occurrence
	Question string `json:"question"`
}

// QuestionnaireItemEnableWhenOperator represents specifies the criteria by which the question is enabled
type QuestionnaireItemEnableWhenOperator string

const (
	QuestionnaireItemEnableWhenOperatorExists         QuestionnaireItemEnableWhenOperator = "exists"
	QuestionnaireItemEnableWhenOperatorEquals         QuestionnaireItemEnableWhenOperator = "="
	QuestionnaireItemEnableWhenOperatorNotEquals      QuestionnaireItemEnableWhenOperator = "!="
	QuestionnaireItemEnableWhenOperatorGreater        QuestionnaireItemEnableWhenOperator = ">"
	QuestionnaireItemEnableWhenOperatorLess           QuestionnaireItemEnableWhenOperator = "<"
	QuestionnaireItemEnableWhenOperatorGreaterOrEqual QuestionnaireItemEnableWhenOperator = ">="
	QuestionnaireItemEnableWhenOperatorLessOrEqual    QuestionnaireItemEnableWhenOperator = "<="
)

// QuestionnaireItemAnswerOption represents this element can be used when the value set machinery of answerValueSet is deemed too cumbersome
type QuestionnaireItemAnswerOption struct {
	common.BackboneElement

	// Use this instead of initial[v] if answerValueSet is present
	InitialSelected *bool `json:"initialSelected,omitempty"`

	// The data type of the value must agree with the item.type
	ValueInteger *int `json:"valueInteger,omitempty"`

	// The data type of the value must agree with the item.type
	ValueDate *string `json:"valueDate,omitempty"`

	// The data type of the value must agree with the item.type
	ValueTime *string `json:"valueTime,omitempty"`

	// The data type of the value must agree with the item.type
	ValueString *string `json:"valueString,omitempty"`

	// The data type of the value must agree with the item.type
	ValueCoding *common.Coding `json:"valueCoding,omitempty"`

	// The data type of the value must agree with the item.type
	ValueReference *common.Reference `json:"valueReference,omitempty"`
}

// QuestionnaireItemInitial represents the user is allowed to change the value and override the default (unless marked as read-only)
type QuestionnaireItemInitial struct {
	common.BackboneElement

	// The type of the initial value must be consistent with the type of the item
	ValueBoolean *bool `json:"valueBoolean,omitempty"`

	// The type of the initial value must be consistent with the type of the item
	ValueDecimal *float64 `json:"valueDecimal,omitempty"`

	// The type of the initial value must be consistent with the type of the item
	ValueInteger *int `json:"valueInteger,omitempty"`

	// The type of the initial value must be consistent with the type of the item
	ValueDate *string `json:"valueDate,omitempty"`

	// The type of the initial value must be consistent with the type of the item
	ValueDateTime *string `json:"valueDateTime,omitempty"`

	// The type of the initial value must be consistent with the type of the item
	ValueTime *string `json:"valueTime,omitempty"`

	// The type of the initial value must be consistent with the type of the item
	ValueString *string `json:"valueString,omitempty"`

	// The type of the initial value must be consistent with the type of the item
	ValueUri *string `json:"valueUri,omitempty"`

	// The type of the initial value must be consistent with the type of the item
	ValueAttachment *common.Attachment `json:"valueAttachment,omitempty"`

	// The type of the initial value must be consistent with the type of the item
	ValueCoding *common.Coding `json:"valueCoding,omitempty"`

	// The type of the initial value must be consistent with the type of the item
	ValueQuantity *common.Quantity `json:"valueQuantity,omitempty"`

	// The type of the initial value must be consistent with the type of the item
	ValueReference *common.Reference `json:"valueReference,omitempty"`
}
