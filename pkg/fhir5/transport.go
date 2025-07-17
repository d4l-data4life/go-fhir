package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// TransportRestriction represents restrictions on transport fulfillment
type TransportRestriction struct {
	common.BackboneElement

	// Note that period.high is the due date representing the time by which the transport should be completed
	Period *common.Period `json:"period,omitempty"`

	// For requests that are targeted to more than one potential recipient/target, to identify who is fulfillment is sought for
	Recipient []common.Reference `json:"recipient,omitempty"`

	// Indicates the number of times the requested action should occur
	Repetitions *int `json:"repetitions,omitempty"`
}

// TransportInput represents additional information that may be needed in the execution of the transport
type TransportInput struct {
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
	ValueID *string `json:"valueId,omitempty"`

	// The value of the input parameter as a basic type
	ValueInstant *string `json:"valueInstant,omitempty"`

	// The value of the input parameter as a basic type
	ValueInteger *int `json:"valueInteger,omitempty"`

	// The value of the input parameter as a basic type
	ValueInteger64 *int64 `json:"valueInteger64,omitempty"`

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
	ValueCodeableReference *common.CodeableReference `json:"valueCodeableReference,omitempty"`

	// The value of the input parameter as a basic type
	ValueCoding *common.Coding `json:"valueCoding,omitempty"`

	// The value of the input parameter as a basic type
	ValueContactPoint *common.ContactPoint `json:"valueContactPoint,omitempty"`

	// The value of the input parameter as a basic type
	ValueCount *common.Count `json:"valueCount,omitempty"`

	// The value of the input parameter as a basic type
	ValueDistance *common.Distance `json:"valueDistance,omitempty"`

	// The value of the input parameter as a basic type
	ValueDuration *Duration `json:"valueDuration,omitempty"`

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
	ValueRange *Range `json:"valueRange,omitempty"`

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

// TransportOutput represents outputs produced by the Transport
type TransportOutput struct {
	common.BackboneElement

	// The type of the output parameter
	Type common.CodeableConcept `json:"type"`

	// The value of the output parameter as a basic type
	ValueBase64Binary *string `json:"valueBase64Binary,omitempty"`

	// The value of the output parameter as a basic type
	ValueBoolean *bool `json:"valueBoolean,omitempty"`

	// The value of the output parameter as a basic type
	ValueCanonical *string `json:"valueCanonical,omitempty"`

	// The value of the output parameter as a basic type
	ValueCode *string `json:"valueCode,omitempty"`

	// The value of the output parameter as a basic type
	ValueDate *string `json:"valueDate,omitempty"`

	// The value of the output parameter as a basic type
	ValueDateTime *string `json:"valueDateTime,omitempty"`

	// The value of the output parameter as a basic type
	ValueDecimal *float64 `json:"valueDecimal,omitempty"`

	// The value of the output parameter as a basic type
	ValueID *string `json:"valueId,omitempty"`

	// The value of the output parameter as a basic type
	ValueInstant *string `json:"valueInstant,omitempty"`

	// The value of the output parameter as a basic type
	ValueInteger *int `json:"valueInteger,omitempty"`

	// The value of the output parameter as a basic type
	ValueInteger64 *int64 `json:"valueInteger64,omitempty"`

	// The value of the output parameter as a basic type
	ValueMarkdown *string `json:"valueMarkdown,omitempty"`

	// The value of the output parameter as a basic type
	ValueOid *string `json:"valueOid,omitempty"`

	// The value of the output parameter as a basic type
	ValuePositiveInt *int `json:"valuePositiveInt,omitempty"`

	// The value of the output parameter as a basic type
	ValueString *string `json:"valueString,omitempty"`

	// The value of the output parameter as a basic type
	ValueTime *string `json:"valueTime,omitempty"`

	// The value of the output parameter as a basic type
	ValueUnsignedInt *int `json:"valueUnsignedInt,omitempty"`

	// The value of the output parameter as a basic type
	ValueUri *string `json:"valueUri,omitempty"`

	// The value of the output parameter as a basic type
	ValueUrl *string `json:"valueUrl,omitempty"`

	// The value of the output parameter as a basic type
	ValueUuid *string `json:"valueUuid,omitempty"`

	// The value of the output parameter as a basic type
	ValueAddress *common.Address `json:"valueAddress,omitempty"`

	// The value of the output parameter as a basic type
	ValueAge *common.Age `json:"valueAge,omitempty"`

	// The value of the output parameter as a basic type
	ValueAnnotation *common.Annotation `json:"valueAnnotation,omitempty"`

	// The value of the output parameter as a basic type
	ValueAttachment *common.Attachment `json:"valueAttachment,omitempty"`

	// The value of the output parameter as a basic type
	ValueCodeableConcept *common.CodeableConcept `json:"valueCodeableConcept,omitempty"`

	// The value of the output parameter as a basic type
	ValueCodeableReference *common.CodeableReference `json:"valueCodeableReference,omitempty"`

	// The value of the output parameter as a basic type
	ValueCoding *common.Coding `json:"valueCoding,omitempty"`

	// The value of the output parameter as a basic type
	ValueContactPoint *common.ContactPoint `json:"valueContactPoint,omitempty"`

	// The value of the output parameter as a basic type
	ValueCount *common.Count `json:"valueCount,omitempty"`

	// The value of the output parameter as a basic type
	ValueDistance *common.Distance `json:"valueDistance,omitempty"`

	// The value of the output parameter as a basic type
	ValueDuration *Duration `json:"valueDuration,omitempty"`

	// The value of the output parameter as a basic type
	ValueHumanName *common.HumanName `json:"valueHumanName,omitempty"`

	// The value of the output parameter as a basic type
	ValueIdentifier *common.Identifier `json:"valueIdentifier,omitempty"`

	// The value of the output parameter as a basic type
	ValueMoney *common.Money `json:"valueMoney,omitempty"`

	// The value of the output parameter as a basic type
	ValuePeriod *common.Period `json:"valuePeriod,omitempty"`

	// The value of the output parameter as a basic type
	ValueQuantity *common.Quantity `json:"valueQuantity,omitempty"`

	// The value of the output parameter as a basic type
	ValueRange *Range `json:"valueRange,omitempty"`

	// The value of the output parameter as a basic type
	ValueRatio *common.Ratio `json:"valueRatio,omitempty"`

	// The value of the output parameter as a basic type
	ValueReference *common.Reference `json:"valueReference,omitempty"`

	// The value of the output parameter as a basic type
	ValueSampledData *common.SampledData `json:"valueSampledData,omitempty"`

	// The value of the output parameter as a basic type
	ValueSignature *common.Signature `json:"valueSignature,omitempty"`

	// The value of the output parameter as a basic type
	ValueTiming *common.Timing `json:"valueTiming,omitempty"`

	// The value of the output parameter as a basic type
	ValueContactDetail *common.ContactDetail `json:"valueContactDetail,omitempty"`

	// The value of the output parameter as a basic type
	ValueContributor *common.Contributor `json:"valueContributor,omitempty"`

	// The value of the output parameter as a basic type
	ValueDataRequirement *common.DataRequirement `json:"valueDataRequirement,omitempty"`

	// The value of the output parameter as a basic type
	ValueExpression *common.Expression `json:"valueExpression,omitempty"`

	// The value of the output parameter as a basic type
	ValueParameterDefinition *common.ParameterDefinition `json:"valueParameterDefinition,omitempty"`

	// The value of the output parameter as a basic type
	ValueRelatedArtifact *common.RelatedArtifact `json:"valueRelatedArtifact,omitempty"`

	// The value of the output parameter as a basic type
	ValueTriggerDefinition *common.TriggerDefinition `json:"valueTriggerDefinition,omitempty"`

	// The value of the output parameter as a basic type
	ValueUsageContext *common.UsageContext `json:"valueUsageContext,omitempty"`

	// The value of the output parameter as a basic type
	ValueDosage *common.Dosage `json:"valueDosage,omitempty"`

	// The value of the output parameter as a basic type
	ValueMeta *common.Meta `json:"valueMeta,omitempty"`
}

// TransportIntent represents the intent of the transport
type TransportIntent string

const (
	TransportIntentUnknown       TransportIntent = "unknown"
	TransportIntentProposal      TransportIntent = "proposal"
	TransportIntentPlan          TransportIntent = "plan"
	TransportIntentOrder         TransportIntent = "order"
	TransportIntentOriginalOrder TransportIntent = "original-order"
	TransportIntentReflexOrder   TransportIntent = "reflex-order"
	TransportIntentFillerOrder   TransportIntent = "filler-order"
	TransportIntentInstanceOrder TransportIntent = "instance-order"
	TransportIntentOption        TransportIntent = "option"
)

// TransportPriority represents the priority of the transport
type TransportPriority string

const (
	TransportPriorityRoutine TransportPriority = "routine"
	TransportPriorityUrgent  TransportPriority = "urgent"
	TransportPriorityASAP    TransportPriority = "asap"
	TransportPriorityStat    TransportPriority = "stat"
)

// TransportStatus represents the status of the transport
type TransportStatus string

const (
	TransportStatusInProgress     TransportStatus = "in-progress"
	TransportStatusCompleted      TransportStatus = "completed"
	TransportStatusAbandoned      TransportStatus = "abandoned"
	TransportStatusCancelled      TransportStatus = "cancelled"
	TransportStatusPlanned        TransportStatus = "planned"
	TransportStatusEnteredInError TransportStatus = "entered-in-error"
)

// Transport represents record of transport of item
type Transport struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Transport"

	// The date and time this transport was created
	AuthoredOn *string `json:"authoredOn,omitempty"`

	// BasedOn refers to a higher-level authorization that triggered the creation of the transport
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// The title should go into the code
	Code *common.CodeableConcept `json:"code,omitempty"`

	// Identifies the completion time of the event
	CompletionTime *string `json:"completionTime,omitempty"`

	// The current location for the entity to be transported
	CurrentLocation common.Reference `json:"currentLocation"`

	// A free-text description of what is to be performed
	Description *string `json:"description,omitempty"`

	// The healthcare event during which this transport was created
	Encounter *common.Reference `json:"encounter,omitempty"`

	// If multiple resources need to be manipulated, use sub-transports
	Focus *common.Reference `json:"focus,omitempty"`

	// The entity who benefits from the performance of the service specified in the transport
	For *common.Reference `json:"for,omitempty"`

	// A shared identifier common to multiple independent Request instances
	GroupIdentifier *common.Identifier `json:"groupIdentifier,omitempty"`

	// The transport event prior to this one
	History *common.Reference `json:"history,omitempty"`

	// This identifier is typically assigned by the dispenser
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Additional information that may be needed in the execution of the transport
	Input []TransportInput `json:"input,omitempty"`

	// The URL pointing to a FHIR-defined protocol, guideline, orderset or other definition
	InstantiatesCanonical *string `json:"instantiatesCanonical,omitempty"`

	// The URL pointing to an externally maintained protocol, guideline, orderset or other definition
	InstantiatesUri *string `json:"instantiatesUri,omitempty"`

	// Insurance plans, coverage extensions, pre-authorizations and/or pre-determinations that may be relevant to the Transport
	Insurance []common.Reference `json:"insurance,omitempty"`

	// This element is immutable
	Intent TransportIntent `json:"intent"`

	// The date and time of last modification to this transport
	LastModified *string `json:"lastModified,omitempty"`

	// Principal physical location where this transport is performed
	Location *common.Reference `json:"location,omitempty"`

	// Free-text information captured about the transport as it progresses
	Note []common.Annotation `json:"note,omitempty"`

	// Outputs produced by the Transport
	Output []TransportOutput `json:"output,omitempty"`

	// Transports may be created with an owner not yet identified
	Owner *common.Reference `json:"owner,omitempty"`

	// Not to be used to link an event to an Encounter
	PartOf []common.Reference `json:"partOf,omitempty"`

	// The kind of participant that should perform the transport
	PerformerType []common.CodeableConcept `json:"performerType,omitempty"`

	// Indicates how quickly the Transport should be addressed with respect to other requests
	Priority *TransportPriority `json:"priority,omitempty"`

	// Transports might be justified based on an Observation, a Condition, a past or planned procedure, etc
	Reason *common.CodeableReference `json:"reason,omitempty"`

	// This element does not point to the Provenance associated with the current version of the resource
	RelevantHistory []common.Reference `json:"relevantHistory,omitempty"`

	// The desired or final location for the transport
	RequestedLocation common.Reference `json:"requestedLocation"`

	// The creator of the transport
	Requester *common.Reference `json:"requester,omitempty"`

	// If the Transport.focus is a request resource and the transport is seeking fulfillment
	Restriction *TransportRestriction `json:"restriction,omitempty"`

	// A code specifying the state of the transport event
	Status *TransportStatus `json:"status,omitempty"`

	// This applies to the current status
	StatusReason *common.CodeableConcept `json:"statusReason,omitempty"`
}
