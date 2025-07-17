package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// ActivityDefinitionParticipantType represents the type of participant in the action
type ActivityDefinitionParticipantType string

const (
	ActivityDefinitionParticipantTypePatient       ActivityDefinitionParticipantType = "patient"
	ActivityDefinitionParticipantTypePractitioner  ActivityDefinitionParticipantType = "practitioner"
	ActivityDefinitionParticipantTypeRelatedPerson ActivityDefinitionParticipantType = "related-person"
	ActivityDefinitionParticipantTypeDevice        ActivityDefinitionParticipantType = "device"
)

// ActivityDefinitionIntent represents the level of authority/intentionality associated with the activity
type ActivityDefinitionIntent string

const (
	ActivityDefinitionIntentProposal      ActivityDefinitionIntent = "proposal"
	ActivityDefinitionIntentPlan          ActivityDefinitionIntent = "plan"
	ActivityDefinitionIntentDirective     ActivityDefinitionIntent = "directive"
	ActivityDefinitionIntentOrder         ActivityDefinitionIntent = "order"
	ActivityDefinitionIntentOriginalOrder ActivityDefinitionIntent = "original-order"
	ActivityDefinitionIntentReflexOrder   ActivityDefinitionIntent = "reflex-order"
	ActivityDefinitionIntentFillerOrder   ActivityDefinitionIntent = "filler-order"
	ActivityDefinitionIntentInstanceOrder ActivityDefinitionIntent = "instance-order"
	ActivityDefinitionIntentOption        ActivityDefinitionIntent = "option"
)

// ActivityDefinitionPriority represents the priority of the activity
type ActivityDefinitionPriority string

const (
	ActivityDefinitionPriorityRoutine ActivityDefinitionPriority = "routine"
	ActivityDefinitionPriorityUrgent  ActivityDefinitionPriority = "urgent"
	ActivityDefinitionPriorityASAP    ActivityDefinitionPriority = "asap"
	ActivityDefinitionPriorityStat    ActivityDefinitionPriority = "stat"
)

// ActivityDefinitionStatus represents the status of the activity definition
type ActivityDefinitionStatus string

const (
	ActivityDefinitionStatusDraft   ActivityDefinitionStatus = "draft"
	ActivityDefinitionStatusActive  ActivityDefinitionStatus = "active"
	ActivityDefinitionStatusRetired ActivityDefinitionStatus = "retired"
	ActivityDefinitionStatusUnknown ActivityDefinitionStatus = "unknown"
)

// ActivityDefinitionParticipant represents who should participate in performing the action described
type ActivityDefinitionParticipant struct {
	common.BackboneElement

	// The role the participant should play in performing the described action
	Role *common.CodeableConcept `json:"role,omitempty"`

	// The type of participant in the action
	Type ActivityDefinitionParticipantType `json:"type"`
}

// ActivityDefinitionDynamicValue represents dynamic values applied to the activity definition
type ActivityDefinitionDynamicValue struct {
	common.BackboneElement

	// The expression may be inlined, or may be a reference to a named expression within a logic library referenced by the library element
	Expression common.Expression `json:"expression"`

	// The path attribute contains a Simple FHIRPath Subset that allows path traversal, but not calculation
	Path string `json:"path"`
}

// ActivityDefinition represents the definition of some activity to be performed
type ActivityDefinition struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "ActivityDefinition"

	// The 'date' element may be more recent than the approval date because of minor changes or editorial corrections
	ApprovalDate *string `json:"approvalDate,omitempty"`

	// An individual or organization primarily involved in the creation and maintenance of the content
	Author []common.ContactDetail `json:"author,omitempty"`

	// Only used if not implicit in the code found in ServiceRequest.type
	BodySite []common.CodeableConcept `json:"bodySite,omitempty"`

	// Tends to be less relevant for activities involving particular products
	Code *common.CodeableConcept `json:"code,omitempty"`

	// May be a web site, an email address, a telephone number, etc
	Contact []common.ContactDetail `json:"contact,omitempty"`

	// A copyright statement relating to the activity definition and/or its contents
	Copyright *string `json:"copyright,omitempty"`

	// Note that this is not the same as the resource last-modified-date, since the resource may be a secondary representation of the activity definition
	Date *string `json:"date,omitempty"`

	// This description can be used to capture details such as why the activity definition was built, comments about misuse, instructions for clinical use and interpretation, literature references, examples from the paper world, etc
	Description *string `json:"description,omitempty"`

	// This element is not intended to be used to communicate a decision support response to cancel an order in progress
	DoNotPerform *bool `json:"doNotPerform,omitempty"`

	// If a dosage instruction is used, the definition should not specify timing or quantity
	Dosage []common.Dosage `json:"dosage,omitempty"`

	// Dynamic values are applied in the order in which they are defined in the ActivityDefinition
	DynamicValue []ActivityDefinitionDynamicValue `json:"dynamicValue,omitempty"`

	// An individual or organization primarily responsible for internal coherence of the content
	Editor []common.ContactDetail `json:"editor,omitempty"`

	// The effective period for a activity definition determines when the content is applicable for usage and is independent of publication and review dates
	EffectivePeriod *common.Period `json:"effectivePeriod,omitempty"`

	// An individual or organization responsible for officially endorsing the content for use in some setting
	Endorser []common.ContactDetail `json:"endorser,omitempty"`

	// Allows filtering of activity definitions that are appropriate for use versus not
	Experimental *bool `json:"experimental,omitempty"`

	// Typically, this is used for identifiers that can go in an HL7 V3 II (instance identifier) data type
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Indicates the level of authority/intentionality associated with the activity and where the request should fit into the workflow chain
	Intent *ActivityDefinitionIntent `json:"intent,omitempty"`

	// It may be possible for the activity definition to be used in jurisdictions other than those for which it was originally designed or intended
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`

	// May determine what types of extensions are permitted
	Kind *string `json:"kind,omitempty"`

	// If specified, this date follows the original approval date
	LastReviewDate *string `json:"lastReviewDate,omitempty"`

	// A reference to a Library resource containing any formal logic used by the activity definition
	Library []string `json:"library,omitempty"`

	// May reference a specific clinical location or may just identify a type of location
	Location *common.Reference `json:"location,omitempty"`

	// The name is not expected to be globally unique. The name should be a simple alphanumeric type name to ensure that it is machine-processing friendly
	Name *string `json:"name,omitempty"`

	// Defines observation requirements for the action to be performed, such as body weight or surface area
	ObservationRequirement []common.Reference `json:"observationRequirement,omitempty"`

	// Defines the observations that are expected to be produced by the action
	ObservationResultRequirement []common.Reference `json:"observationResultRequirement,omitempty"`

	// Indicates who should participate in performing the action described
	Participant []ActivityDefinitionParticipant `json:"participant,omitempty"`

	// Indicates how quickly the activity should be addressed with respect to other requests
	Priority *ActivityDefinitionPriority `json:"priority,omitempty"`

	// Identifies the food, drug or other product being consumed or supplied in the activity
	ProductReference *common.Reference `json:"productReference,omitempty"`

	// Identifies the food, drug or other product being consumed or supplied in the activity
	ProductCodeableConcept *common.CodeableConcept `json:"productCodeableConcept,omitempty"`

	// A profile to which the target of the activity definition is expected to conform
	Profile *string `json:"profile,omitempty"`

	// Usually an organization but may be an individual. The publisher (or steward) of the activity definition is the organization or individual primarily responsible for the maintenance and upkeep of the activity definition
	Publisher *string `json:"publisher,omitempty"`

	// This element does not describe the usage of the activity definition. Instead, it provides traceability of 'why' the resource is either needed or 'why' it is defined as it is
	Purpose *string `json:"purpose,omitempty"`

	// Identifies the quantity expected to be consumed at once (per dose, per meal, etc.)
	Quantity *common.Quantity `json:"quantity,omitempty"`

	// Each related artifact is either an attachment, or a reference to another resource, but not both
	RelatedArtifact []common.RelatedArtifact `json:"relatedArtifact,omitempty"`

	// An individual or organization primarily responsible for review of some aspect of the content
	Reviewer []common.ContactDetail `json:"reviewer,omitempty"`

	// Defines specimen requirements for the action to be performed, such as required specimens for a lab test
	SpecimenRequirement []common.Reference `json:"specimenRequirement,omitempty"`

	// Allows filtering of activity definitions that are appropriate for use versus not
	Status ActivityDefinitionStatus `json:"status"`

	// A code or group definition that describes the intended subject of the activity being defined
	SubjectCodeableConcept *common.CodeableConcept `json:"subjectCodeableConcept,omitempty"`

	// A code or group definition that describes the intended subject of the activity being defined
	SubjectReference *common.Reference `json:"subjectReference,omitempty"`

	// An explanatory or alternate title for the activity definition giving additional information about its content
	Subtitle *string `json:"subtitle,omitempty"`

	// The period, timing or frequency upon which the described activity is to occur
	TimingTiming *common.Timing `json:"timingTiming,omitempty"`

	// The period, timing or frequency upon which the described activity is to occur
	TimingDateTime *string `json:"timingDateTime,omitempty"`

	// The period, timing or frequency upon which the described activity is to occur
	TimingAge *common.Age `json:"timingAge,omitempty"`

	// The period, timing or frequency upon which the described activity is to occur
	TimingPeriod *common.Period `json:"timingPeriod,omitempty"`

	// The period, timing or frequency upon which the described activity is to occur
	TimingRange *common.Range `json:"timingRange,omitempty"`

	// The period, timing or frequency upon which the described activity is to occur
	TimingDuration *common.Duration `json:"timingDuration,omitempty"`

	// This name does not need to be machine-processing friendly and may contain punctuation, white-space, etc
	Title *string `json:"title,omitempty"`

	// Descriptive topics related to the content of the activity. Topics provide a high-level categorization of the activity that can be useful for filtering and searching
	Topic []common.CodeableConcept `json:"topic,omitempty"`

	// Note that if both a transform and dynamic values are specified, the dynamic values will be applied to the result of the transform
	Transform *string `json:"transform,omitempty"`

	// Can be a urn:uuid: or a urn:oid: but real http: addresses are preferred
	URL *string `json:"url,omitempty"`

	// A detailed description of how the activity definition is used from a clinical perspective
	Usage *string `json:"usage,omitempty"`

	// When multiple useContexts are specified, there is no expectation that all or any of the contexts apply
	UseContext []common.UsageContext `json:"useContext,omitempty"`

	// There may be different activity definition instances that have the same identifier but different versions
	Version *string `json:"version,omitempty"`
}
