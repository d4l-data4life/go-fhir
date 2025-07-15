// Package fhir5 contains FHIR R5 resource definitions
package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// QuestionnaireItem represents a single item within a questionnaire
type QuestionnaireItem struct {
	common.BackboneElement

	// A constraint indicating that this item should only be enabled (displayed/allow answers) when the specified condition is true
	EnableBehavior *string `json:"enableBehavior,omitempty"`

	// A constraint indicating that this item should only be enabled when the specified condition is true
	EnableWhen []interface{} `json:"enableWhen,omitempty"`

	// The identifier that is unique to the questionnaire allowing referencing to this item from other questionnaires
	LinkId string `json:"linkId"`

	// The maximum number of characters that are permitted in the answer to be considered a "valid" QuestionnaireResponse
	MaxLength *int `json:"maxLength,omitempty"`

	// The name of a section, the text of a question or text displayed for the item
	Prefix *string `json:"prefix,omitempty"`

	// A particular question, question grouping or display text that is part of the questionnaire
	Question *string `json:"question,omitempty"`

	// An indication, when true, that the item may occur multiple times in the response, collecting multiple answers for questions or multiple sets of answers for groups
	Repeats *bool `json:"repeats,omitempty"`

	// An indication, when true, that the item must be present in a "completed" QuestionnaireResponse
	Required *bool `json:"required,omitempty"`

	// The type of questionnaire item this is - whether text for display, a grouping of other items or a particular type of data to be captured (string, integer, coded choice, etc.)
	Type string `json:"type"`

	// The value that should be defaulted when initially rendering the questionnaire for user input
	Initial []interface{} `json:"initial,omitempty"`

	// Text, questions and other groups to be nested beneath a question or group
	Item []QuestionnaireItem `json:"item,omitempty"`

	// A reference to a value set containing a list of codes representing permitted answers for a "choice" or "open-choice" question
	AnswerOption []interface{} `json:"answerOption,omitempty"`

	// One of the permitted answers for a "choice" or "open-choice" question
	AnswerValueSet *string `json:"answerValueSet,omitempty"`

	// A constraint indicating that this item should only be enabled when the specified condition is true
	Definition *string `json:"definition,omitempty"`

	// The text to be displayed in the questionnaire when this item is displayed
	Text *string `json:"text,omitempty"`
}

// Questionnaire represents a structured set of questions intended to guide the collection of answers from end-users
type Questionnaire struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Questionnaire"

	// The 'date' element may be more recent than the approval date because of minor changes or editorial corrections
	ApprovalDate *string `json:"approvalDate,omitempty"`

	// An individiual or organization primarily involved in the creation and maintenance of the content
	Author []ContactDetail `json:"author,omitempty"`

	// May be a web site, an email address, a telephone number, etc
	Contact []ContactDetail `json:"contact,omitempty"`

	// A copyright statement relating to the questionnaire and/or its contents
	Copyright *string `json:"copyright,omitempty"`

	// Note that this is not the same as the resource last-modified-date, since the resource may be a secondary representation of the questionnaire
	Date *string `json:"date,omitempty"`

	// This description can be used to capture details such as why the questionnaire was built
	Description *string `json:"description,omitempty"`

	// An individual or organization primarily responsible for internal coherence of the content
	Editor []ContactDetail `json:"editor,omitempty"`

	// The effective period for a questionnaire determines when the content is applicable for usage
	EffectivePeriod *common.Period `json:"effectivePeriod,omitempty"`

	// An individual or organization responsible for officially endorsing the content for use in some setting
	Endorser []ContactDetail `json:"endorser,omitempty"`

	// Allows filtering of questionnaires that are appropriate for use versus not
	Experimental *bool `json:"experimental,omitempty"`

	// Typically, this is used for identifiers that can go in an HL7 V3 II (instance identifier) data type
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// It may be possible for the questionnaire to be used in jurisdictions other than those for which it was originally designed or intended
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`

	// If specified, this date follows the original approval date
	LastReviewDate *string `json:"lastReviewDate,omitempty"`

	// The name is not expected to be globally unique
	Name *string `json:"name,omitempty"`

	// Usually an organization but may be an individual
	Publisher *string `json:"publisher,omitempty"`

	// This element does not describe the usage of the questionnaire
	Purpose *string `json:"purpose,omitempty"`

	// Each related artifact is either an attachment, or a reference to another resource, but not both
	RelatedArtifact []interface{} `json:"relatedArtifact,omitempty"`

	// An individual or organization primarily responsible for review of some aspect of the content
	Reviewer []ContactDetail `json:"reviewer,omitempty"`

	// Allows filtering of questionnaires that are appropriate for use versus not
	Status string `json:"status"`

	// An explanatory or alternate title for the questionnaire giving additional information about its content
	Subtitle *string `json:"subtitle,omitempty"`

	// This name does not need to be machine-processing friendly and may contain punctuation, white-space, etc
	Title *string `json:"title,omitempty"`

	// Descriptive topics related to the content of the questionnaire
	Topic []common.CodeableConcept `json:"topic,omitempty"`

	// A high-level category for the questionnaire that distinguishes the kinds of systems that would be interested in the questionnaire
	Type *common.CodeableConcept `json:"type,omitempty"`

	// Can be a urn:uuid: or a urn:oid: but real http: addresses are preferred
	URL *string `json:"url,omitempty"`

	// A detailed description of how the questionnaire is used from a clinical perspective
	Usage *string `json:"usage,omitempty"`

	// When multiple useContexts are specified, there is no expectation that all or any of the contexts apply
	UseContext []UsageContext `json:"useContext,omitempty"`

	// There may be different questionnaire instances that have the same identifier but different versions
	Version *string `json:"version,omitempty"`

	// A particular question, question grouping or display text that is part of the questionnaire
	Item []QuestionnaireItem `json:"item,omitempty"`

	// A reference to a value set containing a list of codes representing permitted answers for a "choice" or "open-choice" question
	Code []common.Coding `json:"code,omitempty"`

	// The subject of the questionnaire
	SubjectType []string `json:"subjectType,omitempty"`
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
