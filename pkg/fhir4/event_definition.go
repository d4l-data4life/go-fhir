package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// EventDefinition represents the EventDefinition resource
type EventDefinition struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "EventDefinition"

	// The 'date' element may be more recent than the approval date because of minor changes or editorial corrections
	ApprovalDate *string `json:"approvalDate,omitempty"`

	// An individual or organization primarily involved in the creation and maintenance of the content
	Author []common.ContactDetail `json:"author,omitempty"`

	// May be a web site, an email address, a telephone number, etc
	Contact []common.ContactDetail `json:"contact,omitempty"`

	// A copyright statement relating to the event definition and/or its contents
	Copyright *string `json:"copyright,omitempty"`

	// Note that this is not the same as the resource last-modified-date
	Date *string `json:"date,omitempty"`

	// This description can be used to capture details such as why the event definition was built
	Description *string `json:"description,omitempty"`

	// An individual or organization primarily responsible for internal coherence of the content
	Editor []common.ContactDetail `json:"editor,omitempty"`

	// The effective period for a event definition determines when the content is applicable for usage
	EffectivePeriod *common.Period `json:"effectivePeriod,omitempty"`

	// An individual or organization responsible for officially endorsing the content for use in some setting
	Endorser []common.ContactDetail `json:"endorser,omitempty"`

	// Allows filtering of event definitions that are appropriate for use versus not
	Experimental *bool `json:"experimental,omitempty"`

	// Typically, this is used for identifiers that can go in an HL7 V3 II (instance identifier) data type
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// It may be possible for the event definition to be used in jurisdictions other than those for which it was originally designed or intended
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`

	// If specified, this date follows the original approval date
	LastReviewDate *string `json:"lastReviewDate,omitempty"`

	// The name is not expected to be globally unique
	Name *string `json:"name,omitempty"`

	// Usually an organization but may be an individual
	Publisher *string `json:"publisher,omitempty"`

	// This element does not describe the usage of the event definition
	Purpose *string `json:"purpose,omitempty"`

	// Each related resource is either an attachment, or a reference to another resource, but not both
	RelatedArtifact []common.RelatedArtifact `json:"relatedArtifact,omitempty"`

	// An individual or organization primarily responsible for review of some aspect of the content
	Reviewer []common.ContactDetail `json:"reviewer,omitempty"`

	// Allows filtering of event definitions that are appropriate for use versus not
	Status EventDefinitionStatus `json:"status"`

	// A code or group definition that describes the intended subject of the event definition
	SubjectCodeableConcept *common.CodeableConcept `json:"subjectCodeableConcept,omitempty"`

	// A code or group definition that describes the intended subject of the event definition
	SubjectReference *common.Reference `json:"subjectReference,omitempty"`

	// An explanatory or alternate title for the event definition giving additional information about its content
	Subtitle *string `json:"subtitle,omitempty"`

	// This name does not need to be machine-processing friendly and may contain punctuation, white-space, etc
	Title *string `json:"title,omitempty"`

	// Descriptive topics related to the module
	Topic []common.CodeableConcept `json:"topic,omitempty"`

	// The trigger element defines when the event occurs
	Trigger []common.TriggerDefinition `json:"trigger"`

	// Can be a urn:uuid: or a urn:oid: but real http: addresses are preferred
	Url *string `json:"url,omitempty"`

	// A detailed description of how the event definition is used from a clinical perspective
	Usage *string `json:"usage,omitempty"`

	// When multiple useContexts are specified, there is no expectation that all or any of the contexts apply
	UseContext []common.UsageContext `json:"useContext,omitempty"`

	// There may be different event definition instances that have the same identifier but different versions
	Version *string `json:"version,omitempty"`
}

// EventDefinitionStatus represents the status of the event definition
type EventDefinitionStatus string

const (
	EventDefinitionStatusDraft   EventDefinitionStatus = "draft"
	EventDefinitionStatusActive  EventDefinitionStatus = "active"
	EventDefinitionStatusRetired EventDefinitionStatus = "retired"
	EventDefinitionStatusUnknown EventDefinitionStatus = "unknown"
)
