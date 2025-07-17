// Package fhir5 contains FHIR R5 resource definitions
package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// RequestOrchestration represents a set of related requests (FHIR R5)
type RequestOrchestration struct {
	DomainResource

	ResourceType          string                        `json:"resourceType"` // Always "RequestOrchestration"
	Identifier            []common.Identifier           `json:"identifier,omitempty"`
	InstantiatesCanonical []string                      `json:"instantiatesCanonical,omitempty"`
	InstantiatesUri       []string                      `json:"instantiatesUri,omitempty"`
	BasedOn               []common.Reference            `json:"basedOn,omitempty"`
	Replaces              []common.Reference            `json:"replaces,omitempty"`
	GroupIdentifier       *common.Identifier            `json:"groupIdentifier,omitempty"`
	Status                RequestOrchestrationStatus    `json:"status"`
	Intent                RequestOrchestrationIntent    `json:"intent"`
	Priority              *RequestOrchestrationPriority `json:"priority,omitempty"`
	Subject               *common.Reference             `json:"subject,omitempty"`
	Encounter             *common.Reference             `json:"encounter,omitempty"`
	AuthoredOn            *string                       `json:"authoredOn,omitempty"`
	Author                *common.Reference             `json:"author,omitempty"`
	Reason                []CodeableReference           `json:"reason,omitempty"`
	Note                  []Annotation                  `json:"note,omitempty"`
	Action                []RequestOrchestrationAction  `json:"action,omitempty"`
}

type RequestOrchestrationStatus string

const (
	RequestOrchestrationStatusDraft          RequestOrchestrationStatus = "draft"
	RequestOrchestrationStatusActive         RequestOrchestrationStatus = "active"
	RequestOrchestrationStatusOnHold         RequestOrchestrationStatus = "on-hold"
	RequestOrchestrationStatusRevoked        RequestOrchestrationStatus = "revoked"
	RequestOrchestrationStatusCompleted      RequestOrchestrationStatus = "completed"
	RequestOrchestrationStatusEnteredInError RequestOrchestrationStatus = "entered-in-error"
	RequestOrchestrationStatusUnknown        RequestOrchestrationStatus = "unknown"
)

type RequestOrchestrationIntent string

const (
	RequestOrchestrationIntentProposal      RequestOrchestrationIntent = "proposal"
	RequestOrchestrationIntentPlan          RequestOrchestrationIntent = "plan"
	RequestOrchestrationIntentOrder         RequestOrchestrationIntent = "order"
	RequestOrchestrationIntentOriginalOrder RequestOrchestrationIntent = "original-order"
	RequestOrchestrationIntentReflexOrder   RequestOrchestrationIntent = "reflex-order"
	RequestOrchestrationIntentFillerOrder   RequestOrchestrationIntent = "filler-order"
	RequestOrchestrationIntentInstanceOrder RequestOrchestrationIntent = "instance-order"
	RequestOrchestrationIntentOption        RequestOrchestrationIntent = "option"
)

type RequestOrchestrationPriority string

const (
	RequestOrchestrationPriorityRoutine RequestOrchestrationPriority = "routine"
	RequestOrchestrationPriorityUrgent  RequestOrchestrationPriority = "urgent"
	RequestOrchestrationPriorityASAP    RequestOrchestrationPriority = "asap"
	RequestOrchestrationPriorityStat    RequestOrchestrationPriority = "stat"
)

type RequestOrchestrationAction struct {
	common.BackboneElement
	Prefix                 *string                                   `json:"prefix,omitempty"`
	Title                  *string                                   `json:"title,omitempty"`
	Description            *string                                   `json:"description,omitempty"`
	TextEquivalent         *string                                   `json:"textEquivalent,omitempty"`
	Priority               *RequestOrchestrationPriority             `json:"priority,omitempty"`
	Code                   []common.CodeableConcept                  `json:"code,omitempty"`
	Reason                 []common.CodeableConcept                  `json:"reason,omitempty"`
	Documentation          []RelatedArtifact                         `json:"documentation,omitempty"`
	Goal                   []common.Reference                        `json:"goal,omitempty"`
	SubjectCodeableConcept *common.CodeableConcept                   `json:"subjectCodeableConcept,omitempty"`
	SubjectReference       *common.Reference                         `json:"subjectReference,omitempty"`
	Trigger                []TriggerDefinition                       `json:"trigger,omitempty"`
	Condition              []RequestOrchestrationActionCondition     `json:"condition,omitempty"`
	Input                  []RequestOrchestrationActionInput         `json:"input,omitempty"`
	Output                 []RequestOrchestrationActionOutput        `json:"output,omitempty"`
	RelatedAction          []RequestOrchestrationActionRelatedAction `json:"relatedAction,omitempty"`
	TimingDateTime         *string                                   `json:"timingDateTime,omitempty"`
	TimingAge              *common.Age                               `json:"timingAge,omitempty"`
	TimingPeriod           *common.Period                            `json:"timingPeriod,omitempty"`
	TimingDuration         *Duration                                 `json:"timingDuration,omitempty"`
	TimingRange            *Range                                    `json:"timingRange,omitempty"`
	TimingTiming           *Timing                                   `json:"timingTiming,omitempty"`
	Participant            []RequestOrchestrationActionParticipant   `json:"participant,omitempty"`
	Type                   *common.CodeableConcept                   `json:"type,omitempty"`
	GroupingBehavior       *string                                   `json:"groupingBehavior,omitempty"`
	SelectionBehavior      *string                                   `json:"selectionBehavior,omitempty"`
	RequiredBehavior       *string                                   `json:"requiredBehavior,omitempty"`
	PrecheckBehavior       *string                                   `json:"precheckBehavior,omitempty"`
	CardinalityBehavior    *string                                   `json:"cardinalityBehavior,omitempty"`
	Resource               *common.Reference                         `json:"resource,omitempty"`
	DefinitionCanonical    *string                                   `json:"definitionCanonical,omitempty"`
	DefinitionUri          *string                                   `json:"definitionUri,omitempty"`
	Transform              *string                                   `json:"transform,omitempty"`
	DynamicValue           []RequestOrchestrationActionDynamicValue  `json:"dynamicValue,omitempty"`
	Action                 []RequestOrchestrationAction              `json:"action,omitempty"`
}

type RequestOrchestrationActionCondition struct {
	common.BackboneElement
	Kind       string      `json:"kind"`
	Expression *Expression `json:"expression,omitempty"`
}

type RequestOrchestrationActionInput struct {
	common.BackboneElement
	Title       *string          `json:"title,omitempty"`
	Requirement *DataRequirement `json:"requirement,omitempty"`
}

type RequestOrchestrationActionOutput struct {
	common.BackboneElement
	Title       *string          `json:"title,omitempty"`
	Requirement *DataRequirement `json:"requirement,omitempty"`
}

type RequestOrchestrationActionRelatedAction struct {
	common.BackboneElement
	TargetId       string    `json:"targetId"`
	Relationship   string    `json:"relationship"`
	OffsetDuration *Duration `json:"offsetDuration,omitempty"`
	OffsetRange    *Range    `json:"offsetRange,omitempty"`
}

type RequestOrchestrationActionParticipant struct {
	common.BackboneElement
	Type string                  `json:"type"`
	Role *common.CodeableConcept `json:"role,omitempty"`
}

type RequestOrchestrationActionDynamicValue struct {
	common.BackboneElement
	Path       string      `json:"path"`
	Expression *Expression `json:"expression,omitempty"`
}
