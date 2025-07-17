package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// PlanDefinition represents the definition of various types of plans as a sharable, consumable, and executable artifact
type PlanDefinition struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "PlanDefinition"

	// Note that there is overlap between many of the elements defined here and the ActivityDefinition resource
	Action []PlanDefinitionAction `json:"action,omitempty"`

	// The 'date' element may be more recent than the approval date because of minor changes or editorial corrections
	ApprovalDate *string `json:"approvalDate,omitempty"`

	// An individiual or organization primarily involved in the creation and maintenance of the content
	Author []common.ContactDetail `json:"author,omitempty"`

	// May be a web site, an email address, a telephone number, etc
	Contact []common.ContactDetail `json:"contact,omitempty"`

	// A copyright statement relating to the plan definition and/or its contents
	Copyright *string `json:"copyright,omitempty"`

	// Note that this is not the same as the resource last-modified-date
	Date *string `json:"date,omitempty"`

	// This description can be used to capture details such as why the plan definition was built
	Description *string `json:"description,omitempty"`

	// An individual or organization primarily responsible for internal coherence of the content
	Editor []common.ContactDetail `json:"editor,omitempty"`

	// The effective period for a plan definition determines when the content is applicable for usage
	EffectivePeriod *common.Period `json:"effectivePeriod,omitempty"`

	// An individual or organization responsible for officially endorsing the content for use in some setting
	Endorser []common.ContactDetail `json:"endorser,omitempty"`

	// Allows filtering of plan definitions that are appropriate for use versus not
	Experimental *bool `json:"experimental,omitempty"`

	// Goals that describe what the activities within the plan are intended to achieve
	Goal []PlanDefinitionGoal `json:"goal,omitempty"`

	// Typically, this is used for identifiers that can go in an HL7 V3 II (instance identifier) data type
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// It may be possible for the plan definition to be used in jurisdictions other than those for which it was originally designed
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`

	// If specified, this date follows the original approval date
	LastReviewDate *string `json:"lastReviewDate,omitempty"`

	// A natural language name identifying the plan definition
	Name *string `json:"name,omitempty"`

	// Usually an organization but may be an individual
	Publisher *string `json:"publisher,omitempty"`

	// This element does not describe the usage of the plan definition
	Purpose *string `json:"purpose,omitempty"`

	// Related artifacts such as additional documentation, justification, or bibliographic references
	RelatedArtifact []common.RelatedArtifact `json:"relatedArtifact,omitempty"`

	// Allows filtering of plan definitions that are appropriate for use versus not
	Status PlanDefinitionStatus `json:"status"`

	// A short, descriptive, user-friendly title for the plan definition
	Title *string `json:"title,omitempty"`

	// Descriptive topics related to the content of the plan definition
	Topic []common.CodeableConcept `json:"topic,omitempty"`

	// Can be a urn:uuid: or a urn:oid: but real http: addresses are preferred
	Url *string `json:"url,omitempty"`

	// When multiple useContexts are specified, there is no expectation that all or any of the contexts apply
	UseContext []common.UsageContext `json:"useContext,omitempty"`

	// There may be different plan definition instances that have the same identifier but different versions
	Version *string `json:"version,omitempty"`
}

// PlanDefinitionStatus represents allows filtering of plan definitions that are appropriate for use versus not
type PlanDefinitionStatus string

const (
	PlanDefinitionStatusDraft   PlanDefinitionStatus = "draft"
	PlanDefinitionStatusActive  PlanDefinitionStatus = "active"
	PlanDefinitionStatusRetired PlanDefinitionStatus = "retired"
	PlanDefinitionStatusUnknown PlanDefinitionStatus = "unknown"
)

// PlanDefinitionGoal represents goals that describe what the activities within the plan are intended to achieve
type PlanDefinitionGoal struct {
	common.BackboneElement

	// Identifies problems, conditions, issues, or concerns the goal is intended to address
	Addresses []common.CodeableConcept `json:"addresses,omitempty"`

	// Indicates a category the goal falls within
	Category *common.CodeableConcept `json:"category,omitempty"`

	// If no code is available, use CodeableConcept.text
	Description common.CodeableConcept `json:"description"`

	// Didactic or other informational resources associated with the goal that provide further supporting information about the goal
	Documentation []common.RelatedArtifact `json:"documentation,omitempty"`

	// Identifies the expected level of importance associated with reaching/sustaining the defined goal
	Priority *common.CodeableConcept `json:"priority,omitempty"`

	// The event after which the goal should begin being pursued
	Start *common.CodeableConcept `json:"start,omitempty"`

	// Indicates what should be done and within what timeframe
	Target []PlanDefinitionGoalTarget `json:"target,omitempty"`
}

// PlanDefinitionGoalTarget represents indicates what should be done and within what timeframe
type PlanDefinitionGoalTarget struct {
	common.BackboneElement

	// The target value of the measure to be achieved to signify fulfillment of the goal
	DetailQuantity *common.Quantity `json:"detailQuantity,omitempty"`

	// The target value of the measure to be achieved to signify fulfillment of the goal
	DetailRange *common.Range `json:"detailRange,omitempty"`

	// The target value of the measure to be achieved to signify fulfillment of the goal
	DetailCodeableConcept *common.CodeableConcept `json:"detailCodeableConcept,omitempty"`

	// Indicates the timeframe after the start of the goal in which the goal should be met
	Due *common.Duration `json:"due,omitempty"`

	// The parameter whose value is to be tracked, e.g. body weight, blood pressure, or hemoglobin A1c level
	Measure *common.CodeableConcept `json:"measure,omitempty"`
}

// PlanDefinitionAction represents note that there is overlap between many of the elements defined here and the ActivityDefinition resource
type PlanDefinitionAction struct {
	common.BackboneElement

	// Sub actions that are contained within the action
	Action []PlanDefinitionAction `json:"action,omitempty"`

	// Defines whether the action can be selected multiple times
	CardinalityBehavior *PlanDefinitionActionCardinalityBehavior `json:"cardinalityBehavior,omitempty"`

	// A code that provides meaning for the action or action group
	Code []common.CodeableConcept `json:"code,omitempty"`

	// When multiple conditions of the same kind are present, the effects are combined using AND semantics
	Condition []PlanDefinitionActionCondition `json:"condition,omitempty"`

	// Note that the definition is optional, and if no definition is specified, a dynamicValue with a root ($this) path can be used to define the entire resource dynamically
	DefinitionCanonical *string `json:"definitionCanonical,omitempty"`

	// Note that the definition is optional, and if no definition is specified, a dynamicValue with a root ($this) path can be used to define the entire resource dynamically
	DefinitionUri *string `json:"definitionUri,omitempty"`

	// A brief description of the action used to provide a summary to display to the user
	Description *string `json:"description,omitempty"`

	// Didactic or other informational resources associated with the action that can be provided to the CDS recipient
	Documentation []common.RelatedArtifact `json:"documentation,omitempty"`

	// Dynamic values are applied in the order in which they are defined in the PlanDefinition resource
	DynamicValue []PlanDefinitionActionDynamicValue `json:"dynamicValue,omitempty"`

	// Identifies goals that this action supports
	GoalId []string `json:"goalId,omitempty"`

	// Defines the grouping behavior for the action and its children
	GroupingBehavior *PlanDefinitionActionGroupingBehavior `json:"groupingBehavior,omitempty"`

	// Defines input data requirements for the action
	Input []common.DataRequirement `json:"input,omitempty"`

	// Defines the outputs of the action, if any
	Output []common.DataRequirement `json:"output,omitempty"`

	// Indicates who should participate in performing the action described
	Participant []PlanDefinitionActionParticipant `json:"participant,omitempty"`

	// Defines whether the action should usually be preselected
	PrecheckBehavior *PlanDefinitionActionPrecheckBehavior `json:"precheckBehavior,omitempty"`

	// A user-visible prefix for the action
	Prefix *string `json:"prefix,omitempty"`

	// Indicates how quickly the action should be addressed with respect to other actions
	Priority *PlanDefinitionActionPriority `json:"priority,omitempty"`

	// This is different than the clinical evidence documentation, it's an actual business description of the reason for performing the action
	Reason []common.CodeableConcept `json:"reason,omitempty"`

	// When an action depends on multiple actions, the meaning is that all actions are dependencies
	RelatedAction []PlanDefinitionActionRelatedAction `json:"relatedAction,omitempty"`

	// Defines the required behavior for the action
	RequiredBehavior *PlanDefinitionActionRequiredBehavior `json:"requiredBehavior,omitempty"`

	// Defines the selection behavior for the action and its children
	SelectionBehavior *PlanDefinitionActionSelectionBehavior `json:"selectionBehavior,omitempty"`

	// The subject of an action overrides the subject at a parent action or on the root of the PlanDefinition if specified
	SubjectCodeableConcept *common.CodeableConcept `json:"subjectCodeableConcept,omitempty"`

	// The subject of an action overrides the subject at a parent action or on the root of the PlanDefinition if specified
	SubjectReference *common.Reference `json:"subjectReference,omitempty"`

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

	// Note that when a referenced ActivityDefinition also defines a transform, the transform specified here generally takes precedence
	Transform *string `json:"transform,omitempty"`

	// A description of when the action should be triggered
	Trigger []common.TriggerDefinition `json:"trigger,omitempty"`

	// The type of action to perform (create, update, remove)
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// PlanDefinitionActionCardinalityBehavior represents defines whether the action can be selected multiple times
type PlanDefinitionActionCardinalityBehavior string

const (
	PlanDefinitionActionCardinalityBehaviorSingle   PlanDefinitionActionCardinalityBehavior = "single"
	PlanDefinitionActionCardinalityBehaviorMultiple PlanDefinitionActionCardinalityBehavior = "multiple"
)

// PlanDefinitionActionGroupingBehavior represents defines the grouping behavior for the action and its children
type PlanDefinitionActionGroupingBehavior string

const (
	PlanDefinitionActionGroupingBehaviorVisualGroup   PlanDefinitionActionGroupingBehavior = "visual-group"
	PlanDefinitionActionGroupingBehaviorLogicalGroup  PlanDefinitionActionGroupingBehavior = "logical-group"
	PlanDefinitionActionGroupingBehaviorSentenceGroup PlanDefinitionActionGroupingBehavior = "sentence-group"
)

// PlanDefinitionActionPrecheckBehavior represents defines whether the action should usually be preselected
type PlanDefinitionActionPrecheckBehavior string

const (
	PlanDefinitionActionPrecheckBehaviorYes PlanDefinitionActionPrecheckBehavior = "yes"
	PlanDefinitionActionPrecheckBehaviorNo  PlanDefinitionActionPrecheckBehavior = "no"
)

// PlanDefinitionActionPriority represents indicates how quickly the action should be addressed with respect to other actions
type PlanDefinitionActionPriority string

const (
	PlanDefinitionActionPriorityRoutine PlanDefinitionActionPriority = "routine"
	PlanDefinitionActionPriorityUrgent  PlanDefinitionActionPriority = "urgent"
	PlanDefinitionActionPriorityAsap    PlanDefinitionActionPriority = "asap"
	PlanDefinitionActionPriorityStat    PlanDefinitionActionPriority = "stat"
)

// PlanDefinitionActionRequiredBehavior represents defines the required behavior for the action
type PlanDefinitionActionRequiredBehavior string

const (
	PlanDefinitionActionRequiredBehaviorMust                 PlanDefinitionActionRequiredBehavior = "must"
	PlanDefinitionActionRequiredBehaviorCould                PlanDefinitionActionRequiredBehavior = "could"
	PlanDefinitionActionRequiredBehaviorMustUnlessDocumented PlanDefinitionActionRequiredBehavior = "must-unless-documented"
)

// PlanDefinitionActionSelectionBehavior represents defines the selection behavior for the action and its children
type PlanDefinitionActionSelectionBehavior string

const (
	PlanDefinitionActionSelectionBehaviorAny        PlanDefinitionActionSelectionBehavior = "any"
	PlanDefinitionActionSelectionBehaviorAll        PlanDefinitionActionSelectionBehavior = "all"
	PlanDefinitionActionSelectionBehaviorAllOrNone  PlanDefinitionActionSelectionBehavior = "all-or-none"
	PlanDefinitionActionSelectionBehaviorExactlyOne PlanDefinitionActionSelectionBehavior = "exactly-one"
	PlanDefinitionActionSelectionBehaviorAtMostOne  PlanDefinitionActionSelectionBehavior = "at-most-one"
	PlanDefinitionActionSelectionBehaviorOneOrMore  PlanDefinitionActionSelectionBehavior = "one-or-more"
)

// PlanDefinitionActionCondition represents when multiple conditions of the same kind are present, the effects are combined using AND semantics
type PlanDefinitionActionCondition struct {
	common.BackboneElement

	// The expression may be inlined or may be a reference to a named expression within a logic library referenced by the library element
	Expression *common.Expression `json:"expression,omitempty"`

	// Applicability criteria are used to determine immediate applicability when a plan definition is applied to a given context
	Kind PlanDefinitionActionConditionKind `json:"kind"`
}

// PlanDefinitionActionConditionKind represents applicability criteria are used to determine immediate applicability when a plan definition is applied to a given context
type PlanDefinitionActionConditionKind string

const (
	PlanDefinitionActionConditionKindApplicability PlanDefinitionActionConditionKind = "applicability"
	PlanDefinitionActionConditionKindStart         PlanDefinitionActionConditionKind = "start"
	PlanDefinitionActionConditionKindStop          PlanDefinitionActionConditionKind = "stop"
)

// PlanDefinitionActionRelatedAction represents when an action depends on multiple actions, the meaning is that all actions are dependencies
type PlanDefinitionActionRelatedAction struct {
	common.BackboneElement

	// The element id of the related action
	ActionId string `json:"actionId"`

	// A duration or range of durations to apply to the relationship
	OffsetDuration *common.Duration `json:"offsetDuration,omitempty"`

	// A duration or range of durations to apply to the relationship
	OffsetRange *common.Range `json:"offsetRange,omitempty"`

	// The relationship of this action to the related action
	Relationship PlanDefinitionActionRelatedActionRelationship `json:"relationship"`
}

// PlanDefinitionActionRelatedActionRelationship represents the relationship of this action to the related action
type PlanDefinitionActionRelatedActionRelationship string

const (
	PlanDefinitionActionRelatedActionRelationshipBeforeStart         PlanDefinitionActionRelatedActionRelationship = "before-start"
	PlanDefinitionActionRelatedActionRelationshipBefore              PlanDefinitionActionRelatedActionRelationship = "before"
	PlanDefinitionActionRelatedActionRelationshipBeforeEnd           PlanDefinitionActionRelatedActionRelationship = "before-end"
	PlanDefinitionActionRelatedActionRelationshipConcurrentWithStart PlanDefinitionActionRelatedActionRelationship = "concurrent-with-start"
	PlanDefinitionActionRelatedActionRelationshipConcurrent          PlanDefinitionActionRelatedActionRelationship = "concurrent"
	PlanDefinitionActionRelatedActionRelationshipConcurrentWithEnd   PlanDefinitionActionRelatedActionRelationship = "concurrent-with-end"
	PlanDefinitionActionRelatedActionRelationshipAfterStart          PlanDefinitionActionRelatedActionRelationship = "after-start"
	PlanDefinitionActionRelatedActionRelationshipAfter               PlanDefinitionActionRelatedActionRelationship = "after"
	PlanDefinitionActionRelatedActionRelationshipAfterEnd            PlanDefinitionActionRelatedActionRelationship = "after-end"
)

// PlanDefinitionActionParticipant represents indicates who should participate in performing the action described
type PlanDefinitionActionParticipant struct {
	common.BackboneElement

	// The role the participant should play in performing the described action
	Role *common.CodeableConcept `json:"role,omitempty"`

	// The type of participant in the action
	Type PlanDefinitionActionParticipantType `json:"type"`
}

// PlanDefinitionActionParticipantType represents the type of participant in the action
type PlanDefinitionActionParticipantType string

const (
	PlanDefinitionActionParticipantTypePatient       PlanDefinitionActionParticipantType = "patient"
	PlanDefinitionActionParticipantTypePractitioner  PlanDefinitionActionParticipantType = "practitioner"
	PlanDefinitionActionParticipantTypeRelatedPerson PlanDefinitionActionParticipantType = "related-person"
	PlanDefinitionActionParticipantTypeDevice        PlanDefinitionActionParticipantType = "device"
)

// PlanDefinitionActionDynamicValue represents dynamic values are applied in the order in which they are defined in the PlanDefinition resource
type PlanDefinitionActionDynamicValue struct {
	common.BackboneElement

	// The expression may be inlined or may be a reference to a named expression within a logic library referenced by the library element
	Expression *common.Expression `json:"expression,omitempty"`

	// To specify the path to the current action being realized, the %action environment variable is available in this path
	Path *string `json:"path,omitempty"`
}
