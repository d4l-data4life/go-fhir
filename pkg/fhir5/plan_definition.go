package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

type PlanDefinitionStatus string

type PlanDefinitionGoal struct {
	common.BackboneElement

	Addresses     []common.CodeableConcept `json:"addresses,omitempty"`
	Category      *common.CodeableConcept  `json:"category,omitempty"`
	Description   common.CodeableConcept   `json:"description"`
	Documentation []interface{}            `json:"documentation,omitempty"`
	Priority      *common.CodeableConcept  `json:"priority,omitempty"`
	Start         *common.CodeableConcept  `json:"start,omitempty"`
	Target        []interface{}            `json:"target,omitempty"`
}

type PlanDefinitionAction struct {
	common.BackboneElement

	Action              []PlanDefinitionAction   `json:"action,omitempty"`
	CardinalityBehavior *string                  `json:"cardinalityBehavior,omitempty"`
	Code                *common.CodeableConcept  `json:"code,omitempty"`
	Condition           []interface{}            `json:"condition,omitempty"`
	DefinitionCanonical *string                  `json:"definitionCanonical,omitempty"`
	DefinitionUri       *string                  `json:"definitionUri,omitempty"`
	Description         *string                  `json:"description,omitempty"`
	Documentation       []interface{}            `json:"documentation,omitempty"`
	DynamicValue        []interface{}            `json:"dynamicValue,omitempty"`
	Expression          *interface{}             `json:"expression,omitempty"`
	GoalId              []string                 `json:"goalId,omitempty"`
	GroupingBehavior    *string                  `json:"groupingBehavior,omitempty"`
	Input               []interface{}            `json:"input,omitempty"`
	Output              []interface{}            `json:"output,omitempty"`
	Participant         []interface{}            `json:"participant,omitempty"`
	PrecheckBehavior    *string                  `json:"precheckBehavior,omitempty"`
	Priority            *string                  `json:"priority,omitempty"`
	Reason              []common.CodeableConcept `json:"reason,omitempty"`
	RelatedAction       []interface{}            `json:"relatedAction,omitempty"`
	RequiredBehavior    *string                  `json:"requiredBehavior,omitempty"`
	SelectionBehavior   *string                  `json:"selectionBehavior,omitempty"`
	TextEquivalent      *string                  `json:"textEquivalent,omitempty"`
	TimingAge           *interface{}             `json:"timingAge,omitempty"`
	TimingDateTime      *string                  `json:"timingDateTime,omitempty"`
	TimingDuration      *interface{}             `json:"timingDuration,omitempty"`
	TimingPeriod        *common.Period           `json:"timingPeriod,omitempty"`
	TimingRange         *interface{}             `json:"timingRange,omitempty"`
	TimingTiming        *interface{}             `json:"timingTiming,omitempty"`
	Title               *string                  `json:"title,omitempty"`
	Transform           *string                  `json:"transform,omitempty"`
	Type                *common.CodeableConcept  `json:"type,omitempty"`
}

type PlanDefinition struct {
	DomainResource

	ResourceType            string                   `json:"resourceType"` // Always "PlanDefinition"
	Action                  []PlanDefinitionAction   `json:"action,omitempty"`
	Actor                   []interface{}            `json:"actor,omitempty"`
	ApprovalDate            *string                  `json:"approvalDate,omitempty"`
	AsNeededBoolean         *bool                    `json:"asNeededBoolean,omitempty"`
	AsNeededCodeableConcept *common.CodeableConcept  `json:"asNeededCodeableConcept,omitempty"`
	Author                  []ContactDetail          `json:"author,omitempty"`
	Contact                 []ContactDetail          `json:"contact,omitempty"`
	Copyright               *string                  `json:"copyright,omitempty"`
	Date                    *string                  `json:"date,omitempty"`
	Description             *string                  `json:"description,omitempty"`
	Editor                  []ContactDetail          `json:"editor,omitempty"`
	EffectivePeriod         *common.Period           `json:"effectivePeriod,omitempty"`
	Endorser                []ContactDetail          `json:"endorser,omitempty"`
	Experimental            *bool                    `json:"experimental,omitempty"`
	Goal                    []PlanDefinitionGoal     `json:"goal,omitempty"`
	Identifier              []common.Identifier      `json:"identifier,omitempty"`
	Jurisdiction            []common.CodeableConcept `json:"jurisdiction,omitempty"`
	LastReviewDate          *string                  `json:"lastReviewDate,omitempty"`
	Library                 []string                 `json:"library,omitempty"`
	Name                    *string                  `json:"name,omitempty"`
	Publisher               *string                  `json:"publisher,omitempty"`
	Purpose                 *string                  `json:"purpose,omitempty"`
	RelatedArtifact         []interface{}            `json:"relatedArtifact,omitempty"`
	Reviewer                []ContactDetail          `json:"reviewer,omitempty"`
	Status                  PlanDefinitionStatus     `json:"status"`
	SubjectCodeableConcept  *common.CodeableConcept  `json:"subjectCodeableConcept,omitempty"`
	SubjectReference        *common.Reference        `json:"subjectReference,omitempty"`
	SubjectCanonical        *string                  `json:"subjectCanonical,omitempty"`
	Subtitle                *string                  `json:"subtitle,omitempty"`
	Title                   *string                  `json:"title,omitempty"`
	Topic                   []common.CodeableConcept `json:"topic,omitempty"`
	Type                    *common.CodeableConcept  `json:"type,omitempty"`
	URL                     *string                  `json:"url,omitempty"`
	Usage                   *string                  `json:"usage,omitempty"`
	UseContext              []UsageContext           `json:"useContext,omitempty"`
	Version                 *string                  `json:"version,omitempty"`
}
