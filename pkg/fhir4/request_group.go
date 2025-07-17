package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// RequestGroup represents a group of related requests that can be used to capture intended activities that have inter-dependencies
type RequestGroup struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "RequestGroup"

	// The actions, if any, produced by the evaluation of the artifact
	Action []RequestGroupAction `json:"action,omitempty"`

	// Provides a reference to the author of the request group
	Author *common.Reference `json:"author,omitempty"`

	// Indicates when the request group was created
	AuthoredOn *string `json:"authoredOn,omitempty"`

	// A plan, proposal or order that is fulfilled in whole or in part by this request
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// This element can be used to provide a code that captures the meaning of the request group as a whole
	Code *common.CodeableConcept `json:"code,omitempty"`

	// Describes the context of the request group, if any
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Requests are linked either by a "basedOn" relationship or by having a common requisition
	GroupIdentifier *common.Identifier `json:"groupIdentifier,omitempty"`

	// Allows a service to provide a unique, business identifier for the request
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// A canonical URL referencing a FHIR-defined protocol, guideline, orderset or other definition
	InstantiatesCanonical []string `json:"instantiatesCanonical,omitempty"`

	// A URL referencing an externally defined protocol, guideline, orderset or other definition
	InstantiatesUri []string `json:"instantiatesUri,omitempty"`

	// Indicates the level of authority/intentionality associated with the request and where the request fits into the workflow chain
	Intent RequestGroupIntent `json:"intent"`

	// Provides a mechanism to communicate additional information about the response
	Note []common.Annotation `json:"note,omitempty"`

	// Indicates how quickly the request should be addressed with respect to other requests
	Priority *RequestGroupPriority `json:"priority,omitempty"`

	// Describes the reason for the request group in coded or textual form
	ReasonCode []common.CodeableConcept `json:"reasonCode,omitempty"`

	// Indicates another resource whose existence justifies this request group
	ReasonReference []common.Reference `json:"reasonReference,omitempty"`

	// The replacement could be because the initial request was immediately rejected
	Replaces []common.Reference `json:"replaces,omitempty"`

	// The current state of the request. For request groups, the status reflects the status of all the requests in the group
	Status RequestGroupStatus `json:"status"`

	// The subject for which the request group was created
	Subject *common.Reference `json:"subject,omitempty"`
}

// RequestGroupIntent represents indicates the level of authority/intentionality associated with the request
type RequestGroupIntent string

const (
	RequestGroupIntentProposal      RequestGroupIntent = "proposal"
	RequestGroupIntentPlan          RequestGroupIntent = "plan"
	RequestGroupIntentDirective     RequestGroupIntent = "directive"
	RequestGroupIntentOrder         RequestGroupIntent = "order"
	RequestGroupIntentOriginalOrder RequestGroupIntent = "original-order"
	RequestGroupIntentReflexOrder   RequestGroupIntent = "reflex-order"
	RequestGroupIntentFillerOrder   RequestGroupIntent = "filler-order"
	RequestGroupIntentInstanceOrder RequestGroupIntent = "instance-order"
	RequestGroupIntentOption        RequestGroupIntent = "option"
)

// RequestGroupStatus represents the current state of the request
type RequestGroupStatus string

const (
	RequestGroupStatusDraft          RequestGroupStatus = "draft"
	RequestGroupStatusActive         RequestGroupStatus = "active"
	RequestGroupStatusOnHold         RequestGroupStatus = "on-hold"
	RequestGroupStatusRevoked        RequestGroupStatus = "revoked"
	RequestGroupStatusCompleted      RequestGroupStatus = "completed"
	RequestGroupStatusEnteredInError RequestGroupStatus = "entered-in-error"
	RequestGroupStatusUnknown        RequestGroupStatus = "unknown"
)

// RequestGroupPriority represents indicates how quickly the request should be addressed with respect to other requests
type RequestGroupPriority string

const (
	RequestGroupPriorityRoutine RequestGroupPriority = "routine"
	RequestGroupPriorityUrgent  RequestGroupPriority = "urgent"
	RequestGroupPriorityAsap    RequestGroupPriority = "asap"
	RequestGroupPriorityStat    RequestGroupPriority = "stat"
)

// RequestGroupAction represents the actions, if any, produced by the evaluation of the artifact
type RequestGroupAction struct {
	common.BackboneElement

	// Sub actions
	Action []RequestGroupAction `json:"action,omitempty"`

	// Defines whether the action can be selected multiple times
	CardinalityBehavior *RequestGroupActionCardinalityBehavior `json:"cardinalityBehavior,omitempty"`

	// A code that provides meaning for the action or action group
	Code []common.CodeableConcept `json:"code,omitempty"`

	// When multiple conditions of the same kind are present, the effects are combined using AND semantics
	Condition []RequestGroupActionCondition `json:"condition,omitempty"`

	// A short description of the action used to provide a summary to display to the user
	Description *string `json:"description,omitempty"`

	// Didactic or other informational resources associated with the action that can be provided to the CDS recipient
	Documentation []common.RelatedArtifact `json:"documentation,omitempty"`

	// Defines the grouping behavior for the action and its children
	GroupingBehavior *RequestGroupActionGroupingBehavior `json:"groupingBehavior,omitempty"`

	// The participant that should perform or be responsible for this action
	Participant []common.Reference `json:"participant,omitempty"`

	// Defines whether the action should usually be preselected
	PrecheckBehavior *RequestGroupActionPrecheckBehavior `json:"precheckBehavior,omitempty"`

	// A user-visible prefix for the action
	Prefix *string `json:"prefix,omitempty"`

	// Indicates how quickly the action should be addressed with respect to other actions
	Priority *RequestGroupActionPriority `json:"priority,omitempty"`

	// A relationship to another action such as "before" or "30-60 minutes after start of"
	RelatedAction []RequestGroupActionRelatedAction `json:"relatedAction,omitempty"`

	// Defines expectations around whether an action is required
	RequiredBehavior *RequestGroupActionRequiredBehavior `json:"requiredBehavior,omitempty"`

	// The target resource SHALL be a Request resource with a Request.intent set to "option"
	Resource *common.Reference `json:"resource,omitempty"`

	// Defines the selection behavior for the action and its children
	SelectionBehavior *RequestGroupActionSelectionBehavior `json:"selectionBehavior,omitempty"`

	// A text equivalent of the action to be performed
	TextEquivalent *string `json:"textEquivalent,omitempty"`

	// An optional value describing when the action should be performed
	TimingDateTime *string `json:"timingDateTime,omitempty"`

	// An optional value describing when the action should be performed
	TimingAge *common.Age `json:"timingAge,omitempty"`

	// An optional value describing when the action should be performed
	TimingPeriod *common.Period `json:"timingPeriod,omitempty"`

	// An optional value describing when the action should be performed
	TimingDuration *common.Duration `json:"timingDuration,omitempty"`

	// An optional value describing when the action should be performed
	TimingRange *common.Range `json:"timingRange,omitempty"`

	// An optional value describing when the action should be performed
	TimingTiming *common.Timing `json:"timingTiming,omitempty"`

	// The title of the action displayed to a user
	Title *string `json:"title,omitempty"`

	// The type of action to perform (create, update, remove)
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// RequestGroupActionCardinalityBehavior represents defines whether the action can be selected multiple times
type RequestGroupActionCardinalityBehavior string

const (
	RequestGroupActionCardinalityBehaviorSingle   RequestGroupActionCardinalityBehavior = "single"
	RequestGroupActionCardinalityBehaviorMultiple RequestGroupActionCardinalityBehavior = "multiple"
)

// RequestGroupActionGroupingBehavior represents defines the grouping behavior for the action and its children
type RequestGroupActionGroupingBehavior string

const (
	RequestGroupActionGroupingBehaviorVisualGroup   RequestGroupActionGroupingBehavior = "visual-group"
	RequestGroupActionGroupingBehaviorLogicalGroup  RequestGroupActionGroupingBehavior = "logical-group"
	RequestGroupActionGroupingBehaviorSentenceGroup RequestGroupActionGroupingBehavior = "sentence-group"
)

// RequestGroupActionPrecheckBehavior represents defines whether the action should usually be preselected
type RequestGroupActionPrecheckBehavior string

const (
	RequestGroupActionPrecheckBehaviorYes RequestGroupActionPrecheckBehavior = "yes"
	RequestGroupActionPrecheckBehaviorNo  RequestGroupActionPrecheckBehavior = "no"
)

// RequestGroupActionPriority represents indicates how quickly the action should be addressed with respect to other actions
type RequestGroupActionPriority string

const (
	RequestGroupActionPriorityRoutine RequestGroupActionPriority = "routine"
	RequestGroupActionPriorityUrgent  RequestGroupActionPriority = "urgent"
	RequestGroupActionPriorityAsap    RequestGroupActionPriority = "asap"
	RequestGroupActionPriorityStat    RequestGroupActionPriority = "stat"
)

// RequestGroupActionRequiredBehavior represents defines expectations around whether an action is required
type RequestGroupActionRequiredBehavior string

const (
	RequestGroupActionRequiredBehaviorMust                 RequestGroupActionRequiredBehavior = "must"
	RequestGroupActionRequiredBehaviorCould                RequestGroupActionRequiredBehavior = "could"
	RequestGroupActionRequiredBehaviorMustUnlessDocumented RequestGroupActionRequiredBehavior = "must-unless-documented"
)

// RequestGroupActionSelectionBehavior represents defines the selection behavior for the action and its children
type RequestGroupActionSelectionBehavior string

const (
	RequestGroupActionSelectionBehaviorAny        RequestGroupActionSelectionBehavior = "any"
	RequestGroupActionSelectionBehaviorAll        RequestGroupActionSelectionBehavior = "all"
	RequestGroupActionSelectionBehaviorAllOrNone  RequestGroupActionSelectionBehavior = "all-or-none"
	RequestGroupActionSelectionBehaviorExactlyOne RequestGroupActionSelectionBehavior = "exactly-one"
	RequestGroupActionSelectionBehaviorAtMostOne  RequestGroupActionSelectionBehavior = "at-most-one"
	RequestGroupActionSelectionBehaviorOneOrMore  RequestGroupActionSelectionBehavior = "one-or-more"
)

// RequestGroupActionCondition represents when multiple conditions of the same kind are present, the effects are combined using AND semantics
type RequestGroupActionCondition struct {
	common.BackboneElement

	// The expression may be inlined, or may be a reference to a named expression within a logic library referenced by the library element
	Expression *common.Expression `json:"expression,omitempty"`

	// Applicability criteria are used to determine immediate applicability when a plan definition is applied to a given context
	Kind RequestGroupActionConditionKind `json:"kind"`
}

// RequestGroupActionConditionKind represents applicability criteria are used to determine immediate applicability when a plan definition is applied to a given context
type RequestGroupActionConditionKind string

const (
	RequestGroupActionConditionKindApplicability RequestGroupActionConditionKind = "applicability"
	RequestGroupActionConditionKindStart         RequestGroupActionConditionKind = "start"
	RequestGroupActionConditionKindStop          RequestGroupActionConditionKind = "stop"
)

// RequestGroupActionRelatedAction represents a relationship to another action such as "before" or "30-60 minutes after start of"
type RequestGroupActionRelatedAction struct {
	common.BackboneElement

	// The element id of the action this is related to
	ActionId string `json:"actionId"`

	// A duration or range of durations to apply to the relationship
	OffsetDuration *common.Duration `json:"offsetDuration,omitempty"`

	// A duration or range of durations to apply to the relationship
	OffsetRange *common.Range `json:"offsetRange,omitempty"`

	// The relationship of this action to the related action
	Relationship RequestGroupActionRelatedActionRelationship `json:"relationship"`
}

// RequestGroupActionRelatedActionRelationship represents the relationship of this action to the related action
type RequestGroupActionRelatedActionRelationship string

const (
	RequestGroupActionRelatedActionRelationshipBeforeStart         RequestGroupActionRelatedActionRelationship = "before-start"
	RequestGroupActionRelatedActionRelationshipBefore              RequestGroupActionRelatedActionRelationship = "before"
	RequestGroupActionRelatedActionRelationshipBeforeEnd           RequestGroupActionRelatedActionRelationship = "before-end"
	RequestGroupActionRelatedActionRelationshipConcurrentWithStart RequestGroupActionRelatedActionRelationship = "concurrent-with-start"
	RequestGroupActionRelatedActionRelationshipConcurrent          RequestGroupActionRelatedActionRelationship = "concurrent"
	RequestGroupActionRelatedActionRelationshipConcurrentWithEnd   RequestGroupActionRelatedActionRelationship = "concurrent-with-end"
	RequestGroupActionRelatedActionRelationshipAfterStart          RequestGroupActionRelatedActionRelationship = "after-start"
	RequestGroupActionRelatedActionRelationshipAfter               RequestGroupActionRelatedActionRelationship = "after"
	RequestGroupActionRelatedActionRelationshipAfterEnd            RequestGroupActionRelatedActionRelationship = "after-end"
)
