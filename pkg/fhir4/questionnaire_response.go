package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

type QuestionnaireResponse struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "QuestionnaireResponse"

	// Business identifier assigned to a particular completed questionnaire
	Identifier *common.Identifier `json:"identifier,omitempty"`

	// Plan/proposal/order fulfilled by this QuestionnaireResponse
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// Composition this QuestionnaireResponse is part of
	PartOf []common.Reference `json:"partOf,omitempty"`

	// Form being answered
	Questionnaire *string `json:"questionnaire,omitempty"`

	// in-progress | completed | amended | entered-in-error | stopped
	Status QuestionnaireResponseStatus `json:"status"`

	// The subject of the questionnaire response
	Subject *common.Reference `json:"subject,omitempty"`

	// Encounter created as part of
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Date the answers were gathered
	Authored *string `json:"authored,omitempty"`

	// Person who received and recorded the answers
	Author *common.Reference `json:"author,omitempty"`

	// The person who answered the questions
	Source *common.Reference `json:"source,omitempty"`

	// Groups and questions
	Item []QuestionnaireResponseItem `json:"item,omitempty"`
}

type QuestionnaireResponseStatus string

const (
	QuestionnaireResponseStatusInProgress     QuestionnaireResponseStatus = "in-progress"
	QuestionnaireResponseStatusCompleted      QuestionnaireResponseStatus = "completed"
	QuestionnaireResponseStatusAmended        QuestionnaireResponseStatus = "amended"
	QuestionnaireResponseStatusEnteredInError QuestionnaireResponseStatus = "entered-in-error"
	QuestionnaireResponseStatusStopped        QuestionnaireResponseStatus = "stopped"
)

type QuestionnaireResponseItem struct {
	common.BackboneElement

	// Pointer to specific item from Questionnaire
	LinkId string `json:"linkId"`

	// ElementDefinition - details for the item
	Definition *string `json:"definition,omitempty"`

	// Name for group or question text
	Text *string `json:"text,omitempty"`

	// The response(s) to the question
	Answer []QuestionnaireResponseItemAnswer `json:"answer,omitempty"`

	// Nested questionnaire response items
	Item []QuestionnaireResponseItem `json:"item,omitempty"`
}

type QuestionnaireResponseItemAnswer struct {
	common.BackboneElement

	// Single-valued answer to the question
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

	// Nested items inside this answer
	Item []QuestionnaireResponseItem `json:"item,omitempty"`
}
