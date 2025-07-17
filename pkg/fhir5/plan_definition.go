// Package fhir5 contains FHIR R5 resource definitions
package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// PlanDefinition represents a definition of a plan for a set of actions (FHIR R5)
type PlanDefinition struct {
	DomainResource

	ResourceType           string                   `json:"resourceType"` // Always "PlanDefinition"
	Url                    *string                  `json:"url,omitempty"`
	Identifier             []common.Identifier      `json:"identifier,omitempty"`
	Version                *string                  `json:"version,omitempty"`
	VersionAlgorithmString *string                  `json:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *common.Coding           `json:"versionAlgorithmCoding,omitempty"`
	Name                   *string                  `json:"name,omitempty"`
	Title                  *string                  `json:"title,omitempty"`
	Subtitle               *string                  `json:"subtitle,omitempty"`
	Type                   *common.CodeableConcept  `json:"type,omitempty"`
	Status                 PublicationStatus        `json:"status"`
	Experimental           *bool                    `json:"experimental,omitempty"`
	SubjectCodeableConcept *common.CodeableConcept  `json:"subjectCodeableConcept,omitempty"`
	SubjectReference       *common.Reference        `json:"subjectReference,omitempty"`
	Date                   *string                  `json:"date,omitempty"`
	Publisher              *string                  `json:"publisher,omitempty"`
	Contact                []ContactDetail          `json:"contact,omitempty"`
	Description            *string                  `json:"description,omitempty"`
	UseContext             []UsageContext           `json:"useContext,omitempty"`
	Jurisdiction           []common.CodeableConcept `json:"jurisdiction,omitempty"`
	Purpose                *string                  `json:"purpose,omitempty"`
	Usage                  *string                  `json:"usage,omitempty"`
	Copyright              *string                  `json:"copyright,omitempty"`
	ApprovalDate           *string                  `json:"approvalDate,omitempty"`
	LastReviewDate         *string                  `json:"lastReviewDate,omitempty"`
	EffectivePeriod        *common.Period           `json:"effectivePeriod,omitempty"`
	Topic                  []common.CodeableConcept `json:"topic,omitempty"`
	Author                 []ContactDetail          `json:"author,omitempty"`
	Editor                 []ContactDetail          `json:"editor,omitempty"`
	Reviewer               []ContactDetail          `json:"reviewer,omitempty"`
	Endorser               []ContactDetail          `json:"endorser,omitempty"`
	RelatedArtifact        []RelatedArtifact        `json:"relatedArtifact,omitempty"`
	Library                []string                 `json:"library,omitempty"`
	Goal                   []PlanDefinitionGoal     `json:"goal,omitempty"`
	Action                 []PlanDefinitionAction   `json:"action,omitempty"`
}

type PlanDefinitionGoal struct {
	common.BackboneElement
	Category      *common.CodeableConcept    `json:"category,omitempty"`
	Description   common.CodeableConcept     `json:"description"`
	Priority      *common.CodeableConcept    `json:"priority,omitempty"`
	Start         *common.CodeableConcept    `json:"start,omitempty"`
	Address       []common.CodeableConcept   `json:"address,omitempty"`
	Documentation []RelatedArtifact          `json:"documentation,omitempty"`
	Target        []PlanDefinitionGoalTarget `json:"target,omitempty"`
}

type PlanDefinitionGoalTarget struct {
	common.BackboneElement
	Measure               *common.CodeableConcept `json:"measure,omitempty"`
	DetailQuantity        *common.Quantity        `json:"detailQuantity,omitempty"`
	DetailRange           *Range                  `json:"detailRange,omitempty"`
	DetailCodeableConcept *common.CodeableConcept `json:"detailCodeableConcept,omitempty"`
	Due                   *Duration               `json:"due,omitempty"`
}

type PlanDefinitionAction struct {
	common.BackboneElement
	Prefix                 *string                             `json:"prefix,omitempty"`
	Title                  *string                             `json:"title,omitempty"`
	Description            *string                             `json:"description,omitempty"`
	TextEquivalent         *string                             `json:"textEquivalent,omitempty"`
	Priority               *common.CodeableConcept             `json:"priority,omitempty"`
	Code                   []common.CodeableConcept            `json:"code,omitempty"`
	Reason                 []common.CodeableConcept            `json:"reason,omitempty"`
	Documentation          []RelatedArtifact                   `json:"documentation,omitempty"`
	GoalId                 []string                            `json:"goalId,omitempty"`
	SubjectCodeableConcept *common.CodeableConcept             `json:"subjectCodeableConcept,omitempty"`
	SubjectReference       *common.Reference                   `json:"subjectReference,omitempty"`
	Trigger                []TriggerDefinition                 `json:"trigger,omitempty"`
	Condition              []PlanDefinitionActionCondition     `json:"condition,omitempty"`
	Input                  []PlanDefinitionActionInput         `json:"input,omitempty"`
	Output                 []PlanDefinitionActionOutput        `json:"output,omitempty"`
	RelatedAction          []PlanDefinitionActionRelatedAction `json:"relatedAction,omitempty"`
	TimingDateTime         *string                             `json:"timingDateTime,omitempty"`
	TimingAge              *Age                                `json:"timingAge,omitempty"`
	TimingPeriod           *common.Period                      `json:"timingPeriod,omitempty"`
	TimingDuration         *Duration                           `json:"timingDuration,omitempty"`
	TimingRange            *Range                              `json:"timingRange,omitempty"`
	TimingTiming           *Timing                             `json:"timingTiming,omitempty"`
	Participant            []PlanDefinitionActionParticipant   `json:"participant,omitempty"`
	Type                   *common.CodeableConcept             `json:"type,omitempty"`
	GroupingBehavior       *string                             `json:"groupingBehavior,omitempty"`
	SelectionBehavior      *string                             `json:"selectionBehavior,omitempty"`
	RequiredBehavior       *string                             `json:"requiredBehavior,omitempty"`
	PrecheckBehavior       *string                             `json:"precheckBehavior,omitempty"`
	CardinalityBehavior    *string                             `json:"cardinalityBehavior,omitempty"`
	DefinitionCanonical    *string                             `json:"definitionCanonical,omitempty"`
	DefinitionUri          *string                             `json:"definitionUri,omitempty"`
	Transform              *string                             `json:"transform,omitempty"`
	DynamicValue           []PlanDefinitionActionDynamicValue  `json:"dynamicValue,omitempty"`
	Action                 []PlanDefinitionAction              `json:"action,omitempty"`
}

type PlanDefinitionActionCondition struct {
	common.BackboneElement
	Kind       string      `json:"kind"`
	Expression *Expression `json:"expression,omitempty"`
}

type PlanDefinitionActionInput struct {
	common.BackboneElement
	Title       *string          `json:"title,omitempty"`
	Requirement *DataRequirement `json:"requirement,omitempty"`
}

type PlanDefinitionActionOutput struct {
	common.BackboneElement
	Title       *string          `json:"title,omitempty"`
	Requirement *DataRequirement `json:"requirement,omitempty"`
}

type PlanDefinitionActionRelatedAction struct {
	common.BackboneElement
	TargetId       string    `json:"targetId"`
	Relationship   string    `json:"relationship"`
	OffsetDuration *Duration `json:"offsetDuration,omitempty"`
	OffsetRange    *Range    `json:"offsetRange,omitempty"`
}

type PlanDefinitionActionParticipant struct {
	common.BackboneElement
	Type string                  `json:"type"`
	Role *common.CodeableConcept `json:"role,omitempty"`
}

type PlanDefinitionActionDynamicValue struct {
	common.BackboneElement
	Path       string      `json:"path"`
	Expression *Expression `json:"expression,omitempty"`
}
