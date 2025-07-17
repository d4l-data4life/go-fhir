// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// ExampleScenarioActor represents a system or person who shares or receives an instance within the scenario
type ExampleScenarioActor struct {
	common.BackboneElement

	// ID of the actor
	ActorID string `json:"actorId"`

	// The description of the actor
	Description *string `json:"description,omitempty"`

	// The name of the actor as shown in the page
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

// ExampleScenarioInstanceContainedInstance represents contained instances
type ExampleScenarioInstanceContainedInstance struct {
	common.BackboneElement

	// Each resource contained in the instance
	ResourceID string `json:"resourceId"`

	// A specific version of a resource contained in the instance
	VersionID *string `json:"versionId,omitempty"`
}

// ExampleScenarioInstanceVersion represents version information for an instance
type ExampleScenarioInstanceVersion struct {
	common.BackboneElement

	// The identifier of a specific version of a resource described in the scenario
	VersionID string `json:"versionId"`

	// The description of the resource version
	Description string `json:"description"`
}

// ExampleScenarioInstance represents a single data collection that is shared as part of the scenario
type ExampleScenarioInstance struct {
	common.BackboneElement

	// The id of the resource for referencing
	ResourceID string `json:"resourceId"`

	// The type of the resource
	ResourceType string `json:"resourceType"`

	// A short name for the resource instance
	Name *string `json:"name,omitempty"`

	// Human-friendly description of the resource instance
	Description *string `json:"description,omitempty"`

	// A specific version of the resource
	Version []ExampleScenarioInstanceVersion `json:"version,omitempty"`

	// Resources contained in the instance
	ContainedInstance []ExampleScenarioInstanceContainedInstance `json:"containedInstance,omitempty"`
}

// ExampleScenarioProcessStepOperation represents operations performed in a process step
type ExampleScenarioProcessStepOperation struct {
	common.BackboneElement

	// The sequential number of the interaction
	Number string `json:"number"`

	// The type of operation - CRUD
	Type string `json:"type"`

	// The human-friendly name of the interaction
	Name *string `json:"name,omitempty"`

	// Who starts the transaction
	Initiator *string `json:"initiator,omitempty"`

	// Who receives the transaction
	Receiver *string `json:"receiver,omitempty"`

	// A comment to be inserted in the diagram
	Description *string `json:"description,omitempty"`

	// Whether the initiator is deactivated right after the transaction
	InitiatorActive *bool `json:"initiatorActive,omitempty"`

	// Whether the receiver is deactivated right after the transaction
	ReceiverActive *bool `json:"receiverActive,omitempty"`

	// Each resource instance used by the initiator
	Request *ExampleScenarioInstanceContainedInstance `json:"request,omitempty"`

	// Each resource instance used by the responder
	Response *ExampleScenarioInstanceContainedInstance `json:"response,omitempty"`
}

// ExampleScenarioProcessStepAlternative represents alternative paths in a process step
type ExampleScenarioProcessStepAlternative struct {
	common.BackboneElement

	// The label to display for the alternative that gives a sense of the circumstance in which the alternative should be invoked
	Title string `json:"title"`

	// A human-readable description of each option
	Description *string `json:"description,omitempty"`

	// What happens in each alternative option
	Step []ExampleScenarioProcessStep `json:"step,omitempty"`
}

// ExampleScenarioProcessStep represents a step in a process
type ExampleScenarioProcessStep struct {
	common.BackboneElement

	// Nested process
	Process []ExampleScenarioProcess `json:"process,omitempty"`

	// If there is a pause in the flow
	Pause *bool `json:"pause,omitempty"`

	// Each interaction or action
	Operation *ExampleScenarioProcessStepOperation `json:"operation,omitempty"`

	// Indicates an alternative step that can be taken instead of the operations on the base step in exceptional/atypical circumstances
	Alternative []ExampleScenarioProcessStepAlternative `json:"alternative,omitempty"`
}

// ExampleScenarioProcess represents a process in the scenario
type ExampleScenarioProcess struct {
	common.BackboneElement

	// The diagram title of the group of operations
	Title string `json:"title"`

	// A longer description of the group of operations
	Description *string `json:"description,omitempty"`

	// Description of final status after the process ends
	PreConditions *string `json:"preConditions,omitempty"`

	// Description of initial status before the process starts
	PostConditions *string `json:"postConditions,omitempty"`

	// Each step of the process
	Step []ExampleScenarioProcessStep `json:"step,omitempty"`
}

// ExampleScenario represents an example of workflow instance
type ExampleScenario struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "ExampleScenario"

	// A system or person who shares or receives an instance within the scenario
	Actor []ExampleScenarioActor `json:"actor,omitempty"`

	// May be a web site, an email address, a telephone number, etc
	Contact []ContactDetail `json:"contact,omitempty"`

	// Use and/or publishing restrictions
	Copyright *string `json:"copyright,omitempty"`

	// Copyright holder and year(s)
	CopyrightLabel *string `json:"copyrightLabel,omitempty"`

	// Date last changed
	Date *string `json:"date,omitempty"`

	// Natural language description of the example scenario
	Description *string `json:"description,omitempty"`

	// For testing purposes, not real usage
	Experimental *bool `json:"experimental,omitempty"`

	// Additional identifier for the example scenario
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// A single data collection that is shared as part of the scenario
	Instance []ExampleScenarioInstance `json:"instance,omitempty"`

	// Intended jurisdiction for example scenario
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`

	// Name for this example scenario (computer friendly)
	Name *string `json:"name,omitempty"`

	// Some scenarios might describe only one process
	Process []ExampleScenarioProcess `json:"process,omitempty"`

	// Name of the publisher (organization or individual)
	Publisher *string `json:"publisher,omitempty"`

	// Why this example scenario is defined
	Purpose *string `json:"purpose,omitempty"`

	// The status of this example scenario
	Status EventDefinitionStatus `json:"status"`

	// Name for this example scenario (human friendly)
	Title *string `json:"title,omitempty"`

	// Canonical identifier for this example scenario
	URL *string `json:"url,omitempty"`

	// The context that the content is intended to support
	UseContext []UsageContext `json:"useContext,omitempty"`

	// Business version of the example scenario
	Version *string `json:"version,omitempty"`

	// How to compare versions
	VersionAlgorithmString *string        `json:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *common.Coding `json:"versionAlgorithmCoding,omitempty"`
}
