// Package fhir5 contains FHIR R5 resource definitions
package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// TaskPerformer represents the entity who performed the requested task
type TaskPerformer struct {
	common.BackboneElement

	// The actor or entity who performed the task
	Actor common.Reference `json:"actor"`

	// A code or description of the performer of the task
	Function *common.CodeableConcept `json:"function,omitempty"`
}

// TaskRestriction represents restriction information for a task
type TaskRestriction struct {
	common.BackboneElement

	// This is distinct from Task.executionPeriod
	Period *common.Period `json:"period,omitempty"`

	// For requests that are targeted to more than one potential recipient/target, to identify who is fulfillment is sought for
	Recipient []common.Reference `json:"recipient,omitempty"`

	// Indicates the number of times the requested action should occur
	Repetitions *int `json:"repetitions,omitempty"`
}

// TaskInput represents input information for a task
type TaskInput struct {
	common.BackboneElement

	// A code or description of how the input is intended to be used as guidance when executing the task
	Type common.CodeableConcept `json:"type"`

	// The value of the input parameter as a basic type
	ValueBase64Binary        *string                     `json:"valueBase64Binary,omitempty"`
	ValueBoolean             *bool                       `json:"valueBoolean,omitempty"`
	ValueCanonical           *string                     `json:"valueCanonical,omitempty"`
	ValueCode                *string                     `json:"valueCode,omitempty"`
	ValueDate                *string                     `json:"valueDate,omitempty"`
	ValueDateTime            *string                     `json:"valueDateTime,omitempty"`
	ValueDecimal             *float64                    `json:"valueDecimal,omitempty"`
	ValueId                  *string                     `json:"valueId,omitempty"`
	ValueInstant             *string                     `json:"valueInstant,omitempty"`
	ValueInteger             *int                        `json:"valueInteger,omitempty"`
	ValueInteger64           *int64                      `json:"valueInteger64,omitempty"`
	ValueMarkdown            *string                     `json:"valueMarkdown,omitempty"`
	ValueOid                 *string                     `json:"valueOid,omitempty"`
	ValuePositiveInt         *int                        `json:"valuePositiveInt,omitempty"`
	ValueString              *string                     `json:"valueString,omitempty"`
	ValueTime                *string                     `json:"valueTime,omitempty"`
	ValueUnsignedInt         *int                        `json:"valueUnsignedInt,omitempty"`
	ValueUri                 *string                     `json:"valueUri,omitempty"`
	ValueUrl                 *string                     `json:"valueUrl,omitempty"`
	ValueUuid                *string                     `json:"valueUuid,omitempty"`
	ValueAddress             *Address                    `json:"valueAddress,omitempty"`
	ValueAge                 *Age                        `json:"valueAge,omitempty"`
	ValueAnnotation          *Annotation                 `json:"valueAnnotation,omitempty"`
	ValueAttachment          *Attachment                 `json:"valueAttachment,omitempty"`
	ValueCodeableConcept     *common.CodeableConcept     `json:"valueCodeableConcept,omitempty"`
	ValueCoding              *common.Coding              `json:"valueCoding,omitempty"`
	ValueContactPoint        *ContactPoint               `json:"valueContactPoint,omitempty"`
	ValueCount               *Count                      `json:"valueCount,omitempty"`
	ValueDistance            *Distance                   `json:"valueDistance,omitempty"`
	ValueDuration            *Duration                   `json:"valueDuration,omitempty"`
	ValueHumanName           *HumanName                  `json:"valueHumanName,omitempty"`
	ValueIdentifier          *common.Identifier          `json:"valueIdentifier,omitempty"`
	ValueMoney               *Money                      `json:"valueMoney,omitempty"`
	ValuePeriod              *common.Period              `json:"valuePeriod,omitempty"`
	ValueQuantity            *common.Quantity            `json:"valueQuantity,omitempty"`
	ValueRange               *Range                      `json:"valueRange,omitempty"`
	ValueRatio               *Ratio                      `json:"valueRatio,omitempty"`
	ValueReference           *common.Reference           `json:"valueReference,omitempty"`
	ValueSampledData         *common.SampledData         `json:"valueSampledData,omitempty"`
	ValueSignature           *Signature                  `json:"valueSignature,omitempty"`
	ValueTiming              *Timing                     `json:"valueTiming,omitempty"`
	ValueContactDetail       *ContactDetail              `json:"valueContactDetail,omitempty"`
	ValueContributor         *common.Contributor         `json:"valueContributor,omitempty"`
	ValueDataRequirement     *DataRequirement            `json:"valueDataRequirement,omitempty"`
	ValueExpression          *Expression                 `json:"valueExpression,omitempty"`
	ValueParameterDefinition *common.ParameterDefinition `json:"valueParameterDefinition,omitempty"`
	ValueRelatedArtifact     *RelatedArtifact            `json:"valueRelatedArtifact,omitempty"`
	ValueTriggerDefinition   *TriggerDefinition          `json:"valueTriggerDefinition,omitempty"`
	ValueUsageContext        *UsageContext               `json:"valueUsageContext,omitempty"`
	ValueDosage              *Dosage                     `json:"valueDosage,omitempty"`
	ValueMeta                *Meta                       `json:"valueMeta,omitempty"`
}

// TaskOutput represents output information for a task
type TaskOutput struct {
	common.BackboneElement

	// The name of the Output parameter
	Name string `json:"name"`

	// The value of the Output parameter as a basic type
	ValueBase64Binary        *string                     `json:"valueBase64Binary,omitempty"`
	ValueBoolean             *bool                       `json:"valueBoolean,omitempty"`
	ValueCanonical           *string                     `json:"valueCanonical,omitempty"`
	ValueCode                *string                     `json:"valueCode,omitempty"`
	ValueDate                *string                     `json:"valueDate,omitempty"`
	ValueDateTime            *string                     `json:"valueDateTime,omitempty"`
	ValueDecimal             *float64                    `json:"valueDecimal,omitempty"`
	ValueId                  *string                     `json:"valueId,omitempty"`
	ValueInstant             *string                     `json:"valueInstant,omitempty"`
	ValueInteger             *int                        `json:"valueInteger,omitempty"`
	ValueInteger64           *int64                      `json:"valueInteger64,omitempty"`
	ValueMarkdown            *string                     `json:"valueMarkdown,omitempty"`
	ValueOid                 *string                     `json:"valueOid,omitempty"`
	ValuePositiveInt         *int                        `json:"valuePositiveInt,omitempty"`
	ValueString              *string                     `json:"valueString,omitempty"`
	ValueTime                *string                     `json:"valueTime,omitempty"`
	ValueUnsignedInt         *int                        `json:"valueUnsignedInt,omitempty"`
	ValueUri                 *string                     `json:"valueUri,omitempty"`
	ValueUrl                 *string                     `json:"valueUrl,omitempty"`
	ValueUuid                *string                     `json:"valueUuid,omitempty"`
	ValueAddress             *Address                    `json:"valueAddress,omitempty"`
	ValueAge                 *Age                        `json:"valueAge,omitempty"`
	ValueAnnotation          *Annotation                 `json:"valueAnnotation,omitempty"`
	ValueAttachment          *Attachment                 `json:"valueAttachment,omitempty"`
	ValueCodeableConcept     *common.CodeableConcept     `json:"valueCodeableConcept,omitempty"`
	ValueCoding              *common.Coding              `json:"valueCoding,omitempty"`
	ValueContactPoint        *ContactPoint               `json:"valueContactPoint,omitempty"`
	ValueCount               *Count                      `json:"valueCount,omitempty"`
	ValueDistance            *Distance                   `json:"valueDistance,omitempty"`
	ValueDuration            *Duration                   `json:"valueDuration,omitempty"`
	ValueHumanName           *HumanName                  `json:"valueHumanName,omitempty"`
	ValueIdentifier          *common.Identifier          `json:"valueIdentifier,omitempty"`
	ValueMoney               *Money                      `json:"valueMoney,omitempty"`
	ValuePeriod              *common.Period              `json:"valuePeriod,omitempty"`
	ValueQuantity            *common.Quantity            `json:"valueQuantity,omitempty"`
	ValueRange               *Range                      `json:"valueRange,omitempty"`
	ValueRatio               *Ratio                      `json:"valueRatio,omitempty"`
	ValueReference           *common.Reference           `json:"valueReference,omitempty"`
	ValueSampledData         *common.SampledData         `json:"valueSampledData,omitempty"`
	ValueSignature           *Signature                  `json:"valueSignature,omitempty"`
	ValueTiming              *Timing                     `json:"valueTiming,omitempty"`
	ValueContactDetail       *ContactDetail              `json:"valueContactDetail,omitempty"`
	ValueContributor         *common.Contributor         `json:"valueContributor,omitempty"`
	ValueDataRequirement     *DataRequirement            `json:"valueDataRequirement,omitempty"`
	ValueExpression          *Expression                 `json:"valueExpression,omitempty"`
	ValueParameterDefinition *common.ParameterDefinition `json:"valueParameterDefinition,omitempty"`
	ValueRelatedArtifact     *RelatedArtifact            `json:"valueRelatedArtifact,omitempty"`
	ValueTriggerDefinition   *TriggerDefinition          `json:"valueTriggerDefinition,omitempty"`
	ValueUsageContext        *UsageContext               `json:"valueUsageContext,omitempty"`
	ValueDosage              *Dosage                     `json:"valueDosage,omitempty"`
	ValueMeta                *Meta                       `json:"valueMeta,omitempty"`
}

// Task represents a task to be performed
type Task struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Task"

	// The date and time this task was created
	AuthoredOn *string `json:"authoredOn,omitempty"`

	// BasedOn refers to a higher-level authorization that triggered the creation of the task
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// A business identifier for this task
	BusinessStatus *common.CodeableConcept `json:"businessStatus,omitempty"`

	// A name or code (or both) briefly describing what the task involves
	Code *common.CodeableConcept `json:"code,omitempty"`

	// The healthcare event (e.g. a patient and healthcare provider interaction) during which this task was created
	Encounter *common.Reference `json:"encounter,omitempty"`

	// The date and time of when the task should be performed
	ExecutionPeriod *common.Period `json:"executionPeriod,omitempty"`

	// The request being fulfilled by this task
	Focus *common.Reference `json:"focus,omitempty"`

	// An identifier for this task
	For *common.Reference `json:"for,omitempty"`

	// An identifier that links together multiple tasks and other requests that were created in the same context
	GroupIdentifier *common.Identifier `json:"groupIdentifier,omitempty"`

	// Additional information about the task
	Input []TaskInput `json:"input,omitempty"`

	// An identifier for this task
	InstantiatesCanonical *string `json:"instantiatesCanonical,omitempty"`

	// An identifier for this task
	InstantiatesUri *string `json:"instantiatesUri,omitempty"`

	// Insurance plans, coverage extensions, pre-authorizations and/or pre-determinations that may be relevant to the task
	Insurance []common.Reference `json:"insurance,omitempty"`

	// The date and time of when the task was last modified
	LastModified *string `json:"lastModified,omitempty"`

	// Additional information about the task
	Location *common.Reference `json:"location,omitempty"`

	// Additional information about the task
	Note []Annotation `json:"note,omitempty"`

	// Additional information about the task
	Output []TaskOutput `json:"output,omitempty"`

	// The entity who performed the requested task
	Performer []TaskPerformer `json:"performer,omitempty"`

	// The type of participant that can execute the task
	PerformerType []common.CodeableConcept `json:"performerType,omitempty"`

	// The priority of the task among other tasks of the same type
	Priority *TaskPriority `json:"priority,omitempty"`

	// A description or code indicating why this task needs to be performed
	Reason *common.CodeableConcept `json:"reason,omitempty"`

	// Links to Provenance records for past versions of this Task that identify key state transitions or updates that are likely to be relevant to a user looking at the current version of the task
	RelevantHistory []common.Reference `json:"relevantHistory,omitempty"`

	// The entity who requested the task
	Requester *common.Reference `json:"requester,omitempty"`

	// Task.restriction can only be present if the Task is seeking fulfillment of another Request resource
	Restriction *TaskRestriction `json:"restriction,omitempty"`

	// The current status of the task
	Status TaskStatus `json:"status"`

	// An explanation as to why this task is held, failed, was refused, etc
	StatusReason *common.CodeableConcept `json:"statusReason,omitempty"`

	// The subject of the task
	Subject *common.Reference `json:"subject,omitempty"`

	// A short description of the task used to provide a summary to the user
	Title *string `json:"title,omitempty"`

	// A URI that identifies the task
	Url *string `json:"url,omitempty"`
}

// TaskStatus represents the status of a task
type TaskStatus string

const (
	TaskStatusDraft          TaskStatus = "draft"
	TaskStatusRequested      TaskStatus = "requested"
	TaskStatusReceived       TaskStatus = "received"
	TaskStatusAccepted       TaskStatus = "accepted"
	TaskStatusRejected       TaskStatus = "rejected"
	TaskStatusReady          TaskStatus = "ready"
	TaskStatusCancelled      TaskStatus = "cancelled"
	TaskStatusInProgress     TaskStatus = "in-progress"
	TaskStatusOnHold         TaskStatus = "on-hold"
	TaskStatusFailed         TaskStatus = "failed"
	TaskStatusCompleted      TaskStatus = "completed"
	TaskStatusEnteredInError TaskStatus = "entered-in-error"
)

// TaskPriority represents the priority of a task
type TaskPriority string

const (
	TaskPriorityRoutine TaskPriority = "routine"
	TaskPriorityUrgent  TaskPriority = "urgent"
	TaskPriorityASAP    TaskPriority = "asap"
	TaskPriorityStat    TaskPriority = "stat"
)
