// Package fhir5 contains FHIR R5 resource definitions
package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// Questionnaire represents a set of questions for gathering information (FHIR R5)
type Questionnaire struct {
	DomainResource

	ResourceType    string                   `json:"resourceType"` // Always "Questionnaire"
	Url             *string                  `json:"url,omitempty"`
	Identifier      []common.Identifier      `json:"identifier,omitempty"`
	Version         *string                  `json:"version,omitempty"`
	Name            *string                  `json:"name,omitempty"`
	Title           *string                  `json:"title,omitempty"`
	DerivedFrom     []string                 `json:"derivedFrom,omitempty"`
	Status          PublicationStatus        `json:"status"`
	Experimental    *bool                    `json:"experimental,omitempty"`
	SubjectType     []string                 `json:"subjectType,omitempty"`
	Date            *string                  `json:"date,omitempty"`
	Publisher       *string                  `json:"publisher,omitempty"`
	Contact         []ContactDetail          `json:"contact,omitempty"`
	Description     *string                  `json:"description,omitempty"`
	UseContext      []UsageContext           `json:"useContext,omitempty"`
	Jurisdiction    []common.CodeableConcept `json:"jurisdiction,omitempty"`
	Purpose         *string                  `json:"purpose,omitempty"`
	Copyright       *string                  `json:"copyright,omitempty"`
	ApprovalDate    *string                  `json:"approvalDate,omitempty"`
	LastReviewDate  *string                  `json:"lastReviewDate,omitempty"`
	EffectivePeriod *common.Period           `json:"effectivePeriod,omitempty"`
	Code            []common.Coding          `json:"code,omitempty"`
	Item            []QuestionnaireItem      `json:"item,omitempty"`
}

type QuestionnaireItem struct {
	common.BackboneElement
	LinkId          string                          `json:"linkId"`
	Definition      *string                         `json:"definition,omitempty"`
	Code            []common.Coding                 `json:"code,omitempty"`
	Prefix          *string                         `json:"prefix,omitempty"`
	Text            *string                         `json:"text,omitempty"`
	Type            string                          `json:"type"`
	EnableWhen      []QuestionnaireItemEnableWhen   `json:"enableWhen,omitempty"`
	EnableBehavior  *string                         `json:"enableBehavior,omitempty"`
	DisabledDisplay *string                         `json:"disabledDisplay,omitempty"`
	Required        *bool                           `json:"required,omitempty"`
	Repeats         *bool                           `json:"repeats,omitempty"`
	ReadOnly        *bool                           `json:"readOnly,omitempty"`
	MaxLength       *int                            `json:"maxLength,omitempty"`
	AnswerValueSet  *string                         `json:"answerValueSet,omitempty"`
	AnswerOption    []QuestionnaireItemAnswerOption `json:"answerOption,omitempty"`
	Initial         []QuestionnaireItemInitial      `json:"initial,omitempty"`
	Item            []QuestionnaireItem             `json:"item,omitempty"`
}

type QuestionnaireItemEnableWhen struct {
	common.BackboneElement
	Question        string            `json:"question"`
	Operator        string            `json:"operator"`
	AnswerBoolean   *bool             `json:"answerBoolean,omitempty"`
	AnswerDecimal   *float64          `json:"answerDecimal,omitempty"`
	AnswerInteger   *int              `json:"answerInteger,omitempty"`
	AnswerDate      *string           `json:"answerDate,omitempty"`
	AnswerDateTime  *string           `json:"answerDateTime,omitempty"`
	AnswerTime      *string           `json:"answerTime,omitempty"`
	AnswerString    *string           `json:"answerString,omitempty"`
	AnswerCoding    *common.Coding    `json:"answerCoding,omitempty"`
	AnswerQuantity  *common.Quantity  `json:"answerQuantity,omitempty"`
	AnswerReference *common.Reference `json:"answerReference,omitempty"`
}

type QuestionnaireItemAnswerOption struct {
	common.BackboneElement
	ValueInteger    *int              `json:"valueInteger,omitempty"`
	ValueDate       *string           `json:"valueDate,omitempty"`
	ValueTime       *string           `json:"valueTime,omitempty"`
	ValueString     *string           `json:"valueString,omitempty"`
	ValueCoding     *common.Coding    `json:"valueCoding,omitempty"`
	ValueReference  *common.Reference `json:"valueReference,omitempty"`
	InitialSelected *bool             `json:"initialSelected,omitempty"`
}

type QuestionnaireItemInitial struct {
	common.BackboneElement
	ValueBoolean    *bool             `json:"valueBoolean,omitempty"`
	ValueDecimal    *float64          `json:"valueDecimal,omitempty"`
	ValueInteger    *int              `json:"valueInteger,omitempty"`
	ValueDate       *string           `json:"valueDate,omitempty"`
	ValueDateTime   *string           `json:"valueDateTime,omitempty"`
	ValueTime       *string           `json:"valueTime,omitempty"`
	ValueString     *string           `json:"valueString,omitempty"`
	ValueUri        *string           `json:"valueUri,omitempty"`
	ValueAttachment *Attachment       `json:"valueAttachment,omitempty"`
	ValueCoding     *common.Coding    `json:"valueCoding,omitempty"`
	ValueQuantity   *common.Quantity  `json:"valueQuantity,omitempty"`
	ValueReference  *common.Reference `json:"valueReference,omitempty"`
}

// QuestionnaireResponseItemAnswer represents an answer to a question
type QuestionnaireResponseItemAnswer struct {
	common.BackboneElement

	// The answer (or one of the answers) provided by the respondent to the question
	ValueBoolean    *bool             `json:"valueBoolean,omitempty"`
	ValueDecimal    *float64          `json:"valueDecimal,omitempty"`
	ValueInteger    *int              `json:"valueInteger,omitempty"`
	ValueDate       *string           `json:"valueDate,omitempty"`
	ValueDateTime   *string           `json:"valueDateTime,omitempty"`
	ValueTime       *string           `json:"valueTime,omitempty"`
	ValueString     *string           `json:"valueString,omitempty"`
	ValueUri        *string           `json:"valueUri,omitempty"`
	ValueAttachment *Attachment       `json:"valueAttachment,omitempty"`
	ValueCoding     *common.Coding    `json:"valueCoding,omitempty"`
	ValueQuantity   *common.Quantity  `json:"valueQuantity,omitempty"`
	ValueReference  *common.Reference `json:"valueReference,omitempty"`

	// Nested groups and questions within this answer
	Item []QuestionnaireResponseItem `json:"item,omitempty"`
}

// QuestionnaireResponseItem represents a group or question item from the original questionnaire
type QuestionnaireResponseItem struct {
	common.BackboneElement

	// The respondent's answer(s) to the question
	Answer []QuestionnaireResponseItemAnswer `json:"answer,omitempty"`

	// The item from the Questionnaire that corresponds to this item in the QuestionnaireResponse resource
	Definition *string `json:"definition,omitempty"`

	// The identifier that is unique to the questionnaire allowing referencing to this item from other questionnaires
	LinkId string `json:"linkId"`

	// Text, questions and other groups to be nested beneath a question or group
	Item []QuestionnaireResponseItem `json:"item,omitempty"`

	// The name of a section, the text of a question or text displayed for the item
	Prefix *string `json:"prefix,omitempty"`

	// A question or group displayed in the questionnaire response
	Question *string `json:"question,omitempty"`

	// The type of questionnaire item this is - whether text for display, a grouping of other items or a particular type of data to be captured
	Type string `json:"type"`

	// The text to be displayed in the questionnaire response when this item is displayed
	Text *string `json:"text,omitempty"`
}

// QuestionnaireResponse represents a structured set of questions and their answers
type QuestionnaireResponse struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "QuestionnaireResponse"

	// The individual or organization that received and recorded the answers to the questionnaire
	Author *common.Reference `json:"author,omitempty"`

	// The date and/or time that this set of answers was last changed
	Authored *string `json:"authored,omitempty"`

	// The encounter or episode of care with primary association to the questionnaire response
	Encounter *common.Reference `json:"encounter,omitempty"`

	// A group or question item from the original questionnaire for which answers are provided
	Item []QuestionnaireResponseItem `json:"item,omitempty"`

	// A procedure or observation that this questionnaire was performed as part of the execution of
	PartOf []common.Reference `json:"partOf,omitempty"`

	// The Questionnaire that defines and organizes this questionnaire response
	Questionnaire *string `json:"questionnaire,omitempty"`

	// The individual or organization that received the responses to the questionnaire
	Source *common.Reference `json:"source,omitempty"`

	// The position of the questionnaire response within its overall lifecycle
	Status string `json:"status"`

	// The subject of the questionnaire response
	Subject *common.Reference `json:"subject,omitempty"`
}
