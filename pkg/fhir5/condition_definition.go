// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// ConditionDefinitionObservation represents observations relevant to a condition
type ConditionDefinitionObservation struct {
	common.BackboneElement

	Category *common.CodeableConcept `json:"category,omitempty"`
	Code     *common.CodeableConcept `json:"code,omitempty"`
}

// ConditionDefinitionMedication represents medications relevant to a condition
type ConditionDefinitionMedication struct {
	common.BackboneElement

	Category *common.CodeableConcept `json:"category,omitempty"`
	Code     *common.CodeableConcept `json:"code,omitempty"`
}

// ConditionDefinitionPrecondition represents a precondition for a condition
type ConditionDefinitionPrecondition struct {
	common.BackboneElement

	Code                 common.CodeableConcept              `json:"code"`
	Type                 ConditionDefinitionPreconditionType `json:"type"`
	ValueCodeableConcept *common.CodeableConcept             `json:"valueCodeableConcept,omitempty"`
	ValueQuantity        *common.Quantity                    `json:"valueQuantity,omitempty"`
}

// ConditionDefinitionQuestionnaire represents a questionnaire for a condition
type ConditionDefinitionQuestionnaire struct {
	common.BackboneElement

	Purpose   ConditionDefinitionQuestionnairePurpose `json:"purpose"`
	Reference common.Reference                        `json:"reference"`
}

// ConditionDefinitionPlan represents a plan for a condition
type ConditionDefinitionPlan struct {
	common.BackboneElement

	Reference common.Reference        `json:"reference"`
	Role      *common.CodeableConcept `json:"role,omitempty"`
}

// ConditionDefinition represents a definition of a condition and information relevant to managing it
type ConditionDefinition struct {
	DomainResource

	ResourceType string `json:"resourceType"` // Always "ConditionDefinition"

	BodySite               *common.CodeableConcept            `json:"bodySite,omitempty"`
	Code                   common.CodeableConcept             `json:"code"`
	Contact                []ContactDetail                    `json:"contact,omitempty"`
	Date                   *string                            `json:"date,omitempty"`
	Definition             []string                           `json:"definition,omitempty"`
	Description            *string                            `json:"description,omitempty"`
	Experimental           *bool                              `json:"experimental,omitempty"`
	HasBodySite            *bool                              `json:"hasBodySite,omitempty"`
	HasSeverity            *bool                              `json:"hasSeverity,omitempty"`
	HasStage               *bool                              `json:"hasStage,omitempty"`
	Identifier             []common.Identifier                `json:"identifier,omitempty"`
	Jurisdiction           []common.CodeableConcept           `json:"jurisdiction,omitempty"`
	Medication             []ConditionDefinitionMedication    `json:"medication,omitempty"`
	Name                   *string                            `json:"name,omitempty"`
	Observation            []ConditionDefinitionObservation   `json:"observation,omitempty"`
	Plan                   []ConditionDefinitionPlan          `json:"plan,omitempty"`
	Precondition           []ConditionDefinitionPrecondition  `json:"precondition,omitempty"`
	Publisher              *string                            `json:"publisher,omitempty"`
	Questionnaire          []ConditionDefinitionQuestionnaire `json:"questionnaire,omitempty"`
	Severity               *common.CodeableConcept            `json:"severity,omitempty"`
	Stage                  *common.CodeableConcept            `json:"stage,omitempty"`
	Status                 PublicationStatus                  `json:"status"`
	Subtitle               *string                            `json:"subtitle,omitempty"`
	Team                   []common.Reference                 `json:"team,omitempty"`
	Title                  *string                            `json:"title,omitempty"`
	URL                    *string                            `json:"url,omitempty"`
	UseContext             []UsageContext                     `json:"useContext,omitempty"`
	Version                *string                            `json:"version,omitempty"`
	VersionAlgorithmString *string                            `json:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *common.Coding                     `json:"versionAlgorithmCoding,omitempty"`
}

// ConditionDefinitionPreconditionType represents the type of precondition
type ConditionDefinitionPreconditionType string

const (
	ConditionDefinitionPreconditionTypeSensitive ConditionDefinitionPreconditionType = "sensitive"
	ConditionDefinitionPreconditionTypeSpecific  ConditionDefinitionPreconditionType = "specific"
)

// ConditionDefinitionQuestionnairePurpose represents the purpose of a questionnaire
type ConditionDefinitionQuestionnairePurpose string

const (
	ConditionDefinitionQuestionnairePurposePreadmit      ConditionDefinitionQuestionnairePurpose = "preadmit"
	ConditionDefinitionQuestionnairePurposeDiffDiagnosis ConditionDefinitionQuestionnairePurpose = "diff-diagnosis"
	ConditionDefinitionQuestionnairePurposeOutcome       ConditionDefinitionQuestionnairePurpose = "outcome"
)
