package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// Task represents a task to be performed
type Task struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Task"

	// The date and time this task was created
	AuthoredOn *string `json:"authoredOn,omitempty"`

	// BasedOn refers to a higher-level authorization that triggered the creation of the task
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// Contains business-specific nuances of the business state
	BusinessStatus *common.CodeableConcept `json:"businessStatus,omitempty"`

	// The title (eg "My Tasks", "Outstanding Tasks for Patient X") should go into the code
	Code *common.CodeableConcept `json:"code,omitempty"`

	// A free-text description of what is to be performed
	Description *string `json:"description,omitempty"`

	// The healthcare event (e.g. a patient and healthcare provider interaction) during which this task was created
	Encounter *common.Reference `json:"encounter,omitempty"`

	// Identifies the time action was first taken against the task (start) and/or the time final action was taken against the task prior to marking it as completed (end)
	ExecutionPeriod *common.Period `json:"executionPeriod,omitempty"`

	// If multiple resources need to be manipulated, use sub-tasks
	Focus *common.Reference `json:"focus,omitempty"`

	// The entity who benefits from the performance of the service specified in the task (e.g., the patient)
	For *common.Reference `json:"for,omitempty"`

	// An identifier that links together multiple tasks and other requests that were created in the same context
	GroupIdentifier *common.Identifier `json:"groupIdentifier,omitempty"`

	// The business identifier for this task
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Additional information that may be needed in the execution of the task
	Input []TaskInput `json:"input,omitempty"`

	// The URL pointing to a *FHIR*-defined protocol, guideline, orderset or other definition that is adhered to in whole or in part by this Task
	InstantiatesCanonical *string `json:"instantiatesCanonical,omitempty"`

	// The URL pointing to an *externally* maintained protocol, guideline, orderset or other definition that is adhered to in whole or in part by this Task
	InstantiatesUri *string `json:"instantiatesUri,omitempty"`

	// Insurance plans, coverage extensions, pre-authorizations and/or pre-determinations that may be relevant to the Task
	Insurance []common.Reference `json:"insurance,omitempty"`

	// This element is immutable. Proposed tasks, planned tasks, etc. must be distinct instances
	Intent string `json:"intent"` // "unknown" | "proposal" | "plan" | "order" | "original-order" | "reflex-order" | "filler-order" | "instance-order" | "option"

	// The date and time of last modification to this task
	LastModified *string `json:"lastModified,omitempty"`

	// Principal physical location where the this task is performed
	Location *common.Reference `json:"location,omitempty"`

	// Free-text information captured about the task as it progresses
	Note []common.Annotation `json:"note,omitempty"`

	// Outputs produced by the Task
	Output []TaskOutput `json:"output,omitempty"`

	// Tasks may be created with an owner not yet identified
	Owner *common.Reference `json:"owner,omitempty"`

	// This should usually be 0..1
	PartOf []common.Reference `json:"partOf,omitempty"`

	// The kind of participant that should perform the task
	PerformerType []common.CodeableConcept `json:"performerType,omitempty"`

	// Indicates how quickly the Task should be addressed with respect to other requests
	Priority *string `json:"priority,omitempty"` // "routine" | "urgent" | "asap" | "stat"

	// This should only be included if there is no focus or if it differs from the reason indicated on the focus
	ReasonCode *common.CodeableConcept `json:"reasonCode,omitempty"`

	// Tasks might be justified based on an Observation, a Condition, a past or planned procedure, etc
	ReasonReference *common.Reference `json:"reasonReference,omitempty"`

	// This element does not point to the Provenance associated with the *current* version of the resource
	RelevantHistory []common.Reference `json:"relevantHistory,omitempty"`

	// The creator of the task
	Requester *common.Reference `json:"requester,omitempty"`

	// If the Task.focus is a request resource and the task is seeking fulfillment, this element identifies any limitations on what parts of the referenced request should be actioned
	Restriction *TaskRestriction `json:"restriction,omitempty"`

	// The current status of the task
	Status string `json:"status"` // "draft" | "requested" | "received" | "accepted" | "rejected" | "ready" | "cancelled" | "in-progress" | "on-hold" | "failed" | "completed" | "entered-in-error"

	// This applies to the current status. Look at the history of the task to see reasons for past statuses
	StatusReason *common.CodeableConcept `json:"statusReason,omitempty"`
}

// TaskRestriction represents if the Task.focus is a request resource and the task is seeking fulfillment, this element identifies any limitations on what parts of the referenced request should be actioned
type TaskRestriction struct {
	common.BackboneElement

	// Note that period.high is the due date representing the time by which the task should be completed
	Period *common.Period `json:"period,omitempty"`

	// For requests that are targeted to more than on potential recipient/target, for whom is fulfillment sought?
	Recipient []common.Reference `json:"recipient,omitempty"`

	// Indicates the number of times the requested action should occur
	Repetitions *int `json:"repetitions,omitempty"`
}

// TaskInput represents additional information that may be needed in the execution of the task
type TaskInput struct {
	common.BackboneElement

	// If referencing a BPMN workflow or Protocol, the "system" is the URL for the workflow definition and the code is the "name" of the required input
	Type common.CodeableConcept `json:"type"`

	// The value of the input parameter as a basic type
	ValueBase64Binary *string `json:"valueBase64Binary,omitempty"`

	// The value of the input parameter as a basic type
	ValueBoolean *bool `json:"valueBoolean,omitempty"`

	// The value of the input parameter as a basic type
	ValueCanonical *string `json:"valueCanonical,omitempty"`

	// The value of the input parameter as a basic type
	ValueCode *string `json:"valueCode,omitempty"`

	// The value of the input parameter as a basic type
	ValueDate *string `json:"valueDate,omitempty"`

	// The value of the input parameter as a basic type
	ValueDateTime *string `json:"valueDateTime,omitempty"`

	// The value of the input parameter as a basic type
	ValueDecimal *float64 `json:"valueDecimal,omitempty"`

	// The value of the input parameter as a basic type
	ValueId *string `json:"valueId,omitempty"`

	// The value of the input parameter as a basic type
	ValueInstant *string `json:"valueInstant,omitempty"`

	// The value of the input parameter as a basic type
	ValueInteger *int `json:"valueInteger,omitempty"`

	// The value of the input parameter as a basic type
	ValueMarkdown *string `json:"valueMarkdown,omitempty"`

	// The value of the input parameter as a basic type
	ValueOid *string `json:"valueOid,omitempty"`

	// The value of the input parameter as a basic type
	ValuePositiveInt *int `json:"valuePositiveInt,omitempty"`

	// The value of the input parameter as a basic type
	ValueString *string `json:"valueString,omitempty"`

	// The value of the input parameter as a basic type
	ValueTime *string `json:"valueTime,omitempty"`

	// The value of the input parameter as a basic type
	ValueUnsignedInt *int `json:"valueUnsignedInt,omitempty"`

	// The value of the input parameter as a basic type
	ValueUri *string `json:"valueUri,omitempty"`

	// The value of the input parameter as a basic type
	ValueUrl *string `json:"valueUrl,omitempty"`

	// The value of the input parameter as a basic type
	ValueUuid *string `json:"valueUuid,omitempty"`

	// The value of the input parameter as a basic type
	ValueAddress *common.Address `json:"valueAddress,omitempty"`

	// The value of the input parameter as a basic type
	ValueAge *common.Age `json:"valueAge,omitempty"`

	// The value of the input parameter as a basic type
	ValueAnnotation *common.Annotation `json:"valueAnnotation,omitempty"`

	// The value of the input parameter as a basic type
	ValueAttachment *common.Attachment `json:"valueAttachment,omitempty"`

	// The value of the input parameter as a basic type
	ValueCodeableConcept *common.CodeableConcept `json:"valueCodeableConcept,omitempty"`

	// The value of the input parameter as a basic type
	ValueCoding *common.Coding `json:"valueCoding,omitempty"`

	// The value of the input parameter as a basic type
	ValueContactPoint *common.ContactPoint `json:"valueContactPoint,omitempty"`

	// The value of the input parameter as a basic type
	ValueCount *common.Count `json:"valueCount,omitempty"`

	// The value of the input parameter as a basic type
	ValueDistance *common.Distance `json:"valueDistance,omitempty"`

	// The value of the input parameter as a basic type
	ValueDuration *common.Duration `json:"valueDuration,omitempty"`

	// The value of the input parameter as a basic type
	ValueHumanName *common.HumanName `json:"valueHumanName,omitempty"`

	// The value of the input parameter as a basic type
	ValueIdentifier *common.Identifier `json:"valueIdentifier,omitempty"`

	// The value of the input parameter as a basic type
	ValueMoney *common.Money `json:"valueMoney,omitempty"`

	// The value of the input parameter as a basic type
	ValuePeriod *common.Period `json:"valuePeriod,omitempty"`

	// The value of the input parameter as a basic type
	ValueQuantity *common.Quantity `json:"valueQuantity,omitempty"`

	// The value of the input parameter as a basic type
	ValueRange *common.Range `json:"valueRange,omitempty"`

	// The value of the input parameter as a basic type
	ValueRatio *common.Ratio `json:"valueRatio,omitempty"`

	// The value of the input parameter as a basic type
	ValueReference *common.Reference `json:"valueReference,omitempty"`

	// The value of the input parameter as a basic type
	ValueSampledData *common.SampledData `json:"valueSampledData,omitempty"`

	// The value of the input parameter as a basic type
	ValueSignature *common.Signature `json:"valueSignature,omitempty"`

	// The value of the input parameter as a basic type
	ValueTiming *common.Timing `json:"valueTiming,omitempty"`

	// The value of the input parameter as a basic type
	ValueContactDetail *common.ContactDetail `json:"valueContactDetail,omitempty"`

	// The value of the input parameter as a basic type
	ValueContributor *common.Contributor `json:"valueContributor,omitempty"`

	// The value of the input parameter as a basic type
	ValueDataRequirement *common.DataRequirement `json:"valueDataRequirement,omitempty"`

	// The value of the input parameter as a basic type
	ValueExpression *common.Expression `json:"valueExpression,omitempty"`

	// The value of the input parameter as a basic type
	ValueParameterDefinition *common.ParameterDefinition `json:"valueParameterDefinition,omitempty"`

	// The value of the input parameter as a basic type
	ValueRelatedArtifact *common.RelatedArtifact `json:"valueRelatedArtifact,omitempty"`

	// The value of the input parameter as a basic type
	ValueTriggerDefinition *common.TriggerDefinition `json:"valueTriggerDefinition,omitempty"`

	// The value of the input parameter as a basic type
	ValueUsageContext *common.UsageContext `json:"valueUsageContext,omitempty"`

	// The value of the input parameter as a basic type
	ValueDosage *common.Dosage `json:"valueDosage,omitempty"`

	// The value of the input parameter as a basic type
	ValueMeta *common.Meta `json:"valueMeta,omitempty"`
}

// TaskOutput represents outputs produced by the Task
type TaskOutput struct {
	common.BackboneElement

	// The name of the Output parameter
	Type common.CodeableConcept `json:"type"`

	// The value of the Output parameter as a basic type
	ValueBase64Binary *string `json:"valueBase64Binary,omitempty"`

	// The value of the Output parameter as a basic type
	ValueBoolean *bool `json:"valueBoolean,omitempty"`

	// The value of the Output parameter as a basic type
	ValueCanonical *string `json:"valueCanonical,omitempty"`

	// The value of the Output parameter as a basic type
	ValueCode *string `json:"valueCode,omitempty"`

	// The value of the Output parameter as a basic type
	ValueDate *string `json:"valueDate,omitempty"`

	// The value of the Output parameter as a basic type
	ValueDateTime *string `json:"valueDateTime,omitempty"`

	// The value of the Output parameter as a basic type
	ValueDecimal *float64 `json:"valueDecimal,omitempty"`

	// The value of the Output parameter as a basic type
	ValueId *string `json:"valueId,omitempty"`

	// The value of the Output parameter as a basic type
	ValueInstant *string `json:"valueInstant,omitempty"`

	// The value of the Output parameter as a basic type
	ValueInteger *int `json:"valueInteger,omitempty"`

	// The value of the Output parameter as a basic type
	ValueMarkdown *string `json:"valueMarkdown,omitempty"`

	// The value of the Output parameter as a basic type
	ValueOid *string `json:"valueOid,omitempty"`

	// The value of the Output parameter as a basic type
	ValuePositiveInt *int `json:"valuePositiveInt,omitempty"`

	// The value of the Output parameter as a basic type
	ValueString *string `json:"valueString,omitempty"`

	// The value of the Output parameter as a basic type
	ValueTime *string `json:"valueTime,omitempty"`

	// The value of the Output parameter as a basic type
	ValueUnsignedInt *int `json:"valueUnsignedInt,omitempty"`

	// The value of the Output parameter as a basic type
	ValueUri *string `json:"valueUri,omitempty"`

	// The value of the Output parameter as a basic type
	ValueUrl *string `json:"valueUrl,omitempty"`

	// The value of the Output parameter as a basic type
	ValueUuid *string `json:"valueUuid,omitempty"`

	// The value of the Output parameter as a basic type
	ValueAddress *common.Address `json:"valueAddress,omitempty"`

	// The value of the Output parameter as a basic type
	ValueAge *common.Age `json:"valueAge,omitempty"`

	// The value of the Output parameter as a basic type
	ValueAnnotation *common.Annotation `json:"valueAnnotation,omitempty"`

	// The value of the Output parameter as a basic type
	ValueAttachment *common.Attachment `json:"valueAttachment,omitempty"`

	// The value of the Output parameter as a basic type
	ValueCodeableConcept *common.CodeableConcept `json:"valueCodeableConcept,omitempty"`

	// The value of the Output parameter as a basic type
	ValueCoding *common.Coding `json:"valueCoding,omitempty"`

	// The value of the Output parameter as a basic type
	ValueContactPoint *common.ContactPoint `json:"valueContactPoint,omitempty"`

	// The value of the Output parameter as a basic type
	ValueCount *common.Count `json:"valueCount,omitempty"`

	// The value of the Output parameter as a basic type
	ValueDistance *common.Distance `json:"valueDistance,omitempty"`

	// The value of the Output parameter as a basic type
	ValueDuration *common.Duration `json:"valueDuration,omitempty"`

	// The value of the Output parameter as a basic type
	ValueHumanName *common.HumanName `json:"valueHumanName,omitempty"`

	// The value of the Output parameter as a basic type
	ValueIdentifier *common.Identifier `json:"valueIdentifier,omitempty"`

	// The value of the Output parameter as a basic type
	ValueMoney *common.Money `json:"valueMoney,omitempty"`

	// The value of the Output parameter as a basic type
	ValuePeriod *common.Period `json:"valuePeriod,omitempty"`

	// The value of the Output parameter as a basic type
	ValueQuantity *common.Quantity `json:"valueQuantity,omitempty"`

	// The value of the Output parameter as a basic type
	ValueRange *common.Range `json:"valueRange,omitempty"`

	// The value of the Output parameter as a basic type
	ValueRatio *common.Ratio `json:"valueRatio,omitempty"`

	// The value of the Output parameter as a basic type
	ValueReference *common.Reference `json:"valueReference,omitempty"`

	// The value of the Output parameter as a basic type
	ValueSampledData *common.SampledData `json:"valueSampledData,omitempty"`

	// The value of the Output parameter as a basic type
	ValueSignature *common.Signature `json:"valueSignature,omitempty"`

	// The value of the Output parameter as a basic type
	ValueTiming *common.Timing `json:"valueTiming,omitempty"`

	// The value of the Output parameter as a basic type
	ValueContactDetail *common.ContactDetail `json:"valueContactDetail,omitempty"`

	// The value of the Output parameter as a basic type
	ValueContributor *common.Contributor `json:"valueContributor,omitempty"`

	// The value of the Output parameter as a basic type
	ValueDataRequirement *common.DataRequirement `json:"valueDataRequirement,omitempty"`

	// The value of the Output parameter as a basic type
	ValueExpression *common.Expression `json:"valueExpression,omitempty"`

	// The value of the Output parameter as a basic type
	ValueParameterDefinition *common.ParameterDefinition `json:"valueParameterDefinition,omitempty"`

	// The value of the Output parameter as a basic type
	ValueRelatedArtifact *common.RelatedArtifact `json:"valueRelatedArtifact,omitempty"`

	// The value of the Output parameter as a basic type
	ValueTriggerDefinition *common.TriggerDefinition `json:"valueTriggerDefinition,omitempty"`

	// The value of the Output parameter as a basic type
	ValueUsageContext *common.UsageContext `json:"valueUsageContext,omitempty"`

	// The value of the Output parameter as a basic type
	ValueDosage *common.Dosage `json:"valueDosage,omitempty"`

	// The value of the Output parameter as a basic type
	ValueMeta *common.Meta `json:"valueMeta,omitempty"`
}
