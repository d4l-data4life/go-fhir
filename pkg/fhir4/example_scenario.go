package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// ExampleScenario represents an example of workflow instance
type ExampleScenario struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "ExampleScenario"

	// Actor participating in the resource
	Actor []ExampleScenarioActor `json:"actor,omitempty"`

	// May be a web site, an email address, a telephone number, etc
	Contact []common.ContactDetail `json:"contact,omitempty"`

	// Frequently, the copyright differs between the value set and the codes that are included
	Copyright *string `json:"copyright,omitempty"`

	// Note that this is not the same as the resource last-modified-date
	Date *string `json:"date,omitempty"`

	// Allows filtering of example scenarios that are appropriate for use versus not
	Experimental *bool `json:"experimental,omitempty"`

	// Typically, this is used for identifiers that can go in an HL7 V3 II (instance identifier) data type
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Each resource and each version that is present in the workflow
	Instance []ExampleScenarioInstance `json:"instance,omitempty"`

	// It may be possible for the example scenario to be used in jurisdictions other than those for which it was originally designed or intended
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`

	// The name is not expected to be globally unique
	Name *string `json:"name,omitempty"`

	// Each major process - a group of operations
	Process []ExampleScenarioProcess `json:"process,omitempty"`

	// Usually an organization but may be an individual
	Publisher *string `json:"publisher,omitempty"`

	// This element does not describe the usage of the example scenario
	Purpose *string `json:"purpose,omitempty"`

	// Allows filtering of example scenarios that are appropriate for use versus not
	Status ExampleScenarioStatus `json:"status"`

	// Can be a urn:uuid: or a urn:oid: but real http: addresses are preferred
	Url *string `json:"url,omitempty"`

	// When multiple useContexts are specified, there is no expectation that all or any of the contexts apply
	UseContext []common.UsageContext `json:"useContext,omitempty"`

	// There may be different example scenario instances that have the same identifier but different versions
	Version *string `json:"version,omitempty"`

	// Another nested workflow
	Workflow []string `json:"workflow,omitempty"`
}

// ExampleScenarioStatus represents the status of the example scenario
type ExampleScenarioStatus string

const (
	ExampleScenarioStatusDraft   ExampleScenarioStatus = "draft"
	ExampleScenarioStatusActive  ExampleScenarioStatus = "active"
	ExampleScenarioStatusRetired ExampleScenarioStatus = "retired"
	ExampleScenarioStatusUnknown ExampleScenarioStatus = "unknown"
)

// ExampleScenarioActor represents an actor participating in the resource
type ExampleScenarioActor struct {
	common.BackboneElement

	// Should this be called ID or acronym?
	ActorId string `json:"actorId"`

	// Cardinality: is name and description 1..1?
	Description *string `json:"description,omitempty"`

	// Cardinality: is name and description 1..1?
	Name *string `json:"name,omitempty"`

	// The type of actor - person or system
	Type ExampleScenarioActorType `json:"type"`
}

// ExampleScenarioActorType represents the type of actor
type ExampleScenarioActorType string

const (
	ExampleScenarioActorTypePerson ExampleScenarioActorType = "person"
	ExampleScenarioActorTypeEntity ExampleScenarioActorType = "entity"
)

// ExampleScenarioInstance represents each resource and each version that is present in the workflow
type ExampleScenarioInstance struct {
	common.BackboneElement

	// Resources contained in the instance (e.g. the observations contained in a bundle)
	ContainedInstance []ExampleScenarioInstanceContainedInstance `json:"containedInstance,omitempty"`

	// Human-friendly description of the resource instance
	Description *string `json:"description,omitempty"`

	// A short name for the resource instance
	Name *string `json:"name,omitempty"`

	// The id of the resource for referencing
	ResourceId string `json:"resourceId"`

	// The type of the resource
	ResourceType string `json:"resourceType"`

	// A specific version of the resource
	Version []ExampleScenarioInstanceVersion `json:"version,omitempty"`
}

// ExampleScenarioInstanceVersion represents a specific version of the resource
type ExampleScenarioInstanceVersion struct {
	common.BackboneElement

	// The description of the resource version
	Description string `json:"description"`

	// The identifier of a specific version of a resource
	VersionId string `json:"versionId"`
}

// ExampleScenarioInstanceContainedInstance represents resources contained in the instance
type ExampleScenarioInstanceContainedInstance struct {
	common.BackboneElement

	// Each resource contained in the instance
	ResourceId string `json:"resourceId"`

	// A specific version of a resource contained in the instance
	VersionId *string `json:"versionId,omitempty"`
}

// ExampleScenarioProcess represents each major process - a group of operations
type ExampleScenarioProcess struct {
	common.BackboneElement

	// A longer description of the group of operations
	Description *string `json:"description,omitempty"`

	// Description of final status after the process ends
	PostConditions *string `json:"postConditions,omitempty"`

	// Description of initial status before the process starts
	PreConditions *string `json:"preConditions,omitempty"`

	// Each step of the process
	Step []ExampleScenarioProcessStep `json:"step,omitempty"`

	// The diagram title of the group of operations
	Title string `json:"title"`
}

// ExampleScenarioProcessStep represents each step of the process
type ExampleScenarioProcessStep struct {
	common.BackboneElement

	// Indicates an alternative step that can be taken instead of the operations on the base step
	Alternative []ExampleScenarioProcessStepAlternative `json:"alternative,omitempty"`

	// Each interaction or action
	Operation *ExampleScenarioProcessStepOperation `json:"operation,omitempty"`

	// If there is a pause in the flow
	Pause *bool `json:"pause,omitempty"`

	// Nested process
	Process []ExampleScenarioProcess `json:"process,omitempty"`
}

// ExampleScenarioProcessStepAlternative represents an alternative step that can be taken
type ExampleScenarioProcessStepAlternative struct {
	common.BackboneElement

	// A human-readable description of the alternative explaining when the alternative should occur
	Description *string `json:"description,omitempty"`

	// What happens in each alternative option
	Step []ExampleScenarioProcessStep `json:"step,omitempty"`

	// The label to display for the alternative that gives a sense of the circumstance
	Title string `json:"title"`
}

// ExampleScenarioProcessStepOperation represents each interaction or action
type ExampleScenarioProcessStepOperation struct {
	common.BackboneElement

	// A comment to be inserted in the diagram
	Description *string `json:"description,omitempty"`

	// Who starts the transaction
	Initiator *string `json:"initiator,omitempty"`

	// Whether the initiator is deactivated right after the transaction
	InitiatorActive *bool `json:"initiatorActive,omitempty"`

	// The human-friendly name of the interaction
	Name *string `json:"name,omitempty"`

	// The sequential number of the interaction, e.g. 1.2.5
	Number string `json:"number"`

	// Who receives the transaction
	Receiver *string `json:"receiver,omitempty"`

	// Whether the receiver is deactivated right after the transaction
	ReceiverActive *bool `json:"receiverActive,omitempty"`

	// Each resource instance used by the initiator
	Request *ExampleScenarioInstanceContainedInstance `json:"request,omitempty"`

	// Each resource instance used by the responder
	Response *ExampleScenarioInstanceContainedInstance `json:"response,omitempty"`

	// The type of operation - CRUD
	Type *string `json:"type,omitempty"`
}
