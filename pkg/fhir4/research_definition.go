package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// ResearchDefinition represents the ResearchDefinition resource describes the conditional state and outcome that the knowledge is about
type ResearchDefinition struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "ResearchDefinition"

	// The 'date' element may be more recent than the approval date because of minor changes or editorial corrections
	ApprovalDate *string `json:"approvalDate,omitempty"`

	// An individiual or organization primarily involved in the creation and maintenance of the content
	Author []common.ContactDetail `json:"author,omitempty"`

	// A human-readable string to clarify or explain concepts about the resource
	Comment []string `json:"comment,omitempty"`

	// May be a web site, an email address, a telephone number, etc
	Contact []common.ContactDetail `json:"contact,omitempty"`

	// A copyright statement relating to the research definition and/or its contents
	Copyright *string `json:"copyright,omitempty"`

	// Note that this is not the same as the resource last-modified-date
	Date *string `json:"date,omitempty"`

	// This description can be used to capture details such as why the research definition was built
	Description *string `json:"description,omitempty"`

	// An individual or organization primarily responsible for internal coherence of the content
	Editor []common.ContactDetail `json:"editor,omitempty"`

	// The effective period for a research definition determines when the content is applicable for usage
	EffectivePeriod *common.Period `json:"effectivePeriod,omitempty"`

	// An individual or organization responsible for officially endorsing the content for use in some setting
	Endorser []common.ContactDetail `json:"endorser,omitempty"`

	// Allows filtering of research definitions that are appropriate for use versus not
	Experimental *bool `json:"experimental,omitempty"`

	// A reference to a ResearchElementDefinition resource that defines the exposure for the research
	Exposure *common.Reference `json:"exposure,omitempty"`

	// A reference to a ResearchElementDefinition resource that defines the exposureAlternative for the research
	ExposureAlternative *common.Reference `json:"exposureAlternative,omitempty"`

	// Typically, this is used for identifiers that can go in an HL7 V3 II (instance identifier) data type
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// It may be possible for the research definition to be used in jurisdictions other than those for which it was originally designed
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`

	// If specified, this date follows the original approval date
	LastReviewDate *string `json:"lastReviewDate,omitempty"`

	// A reference to a Library resource containing the formal logic used by the ResearchDefinition
	Library []string `json:"library,omitempty"`

	// The name is not expected to be globally unique
	Name *string `json:"name,omitempty"`

	// A reference to a ResearchElementDefinition resomece that defines the outcome for the research
	Outcome *common.Reference `json:"outcome,omitempty"`

	// A reference to a ResearchElementDefinition resource that defines the population for the research
	Population common.Reference `json:"population"`

	// Usually an organization but may be an individual
	Publisher *string `json:"publisher,omitempty"`

	// This element does not describe the usage of the research definition
	Purpose *string `json:"purpose,omitempty"`

	// Each related artifact is either an attachment, or a reference to another resource, but not both
	RelatedArtifact []common.RelatedArtifact `json:"relatedArtifact,omitempty"`

	// An individual or organization primarily responsible for review of some aspect of the content
	Reviewer []common.ContactDetail `json:"reviewer,omitempty"`

	// The short title provides an alternate title for use in informal descriptive contexts
	ShortTitle *string `json:"shortTitle,omitempty"`

	// Allows filtering of research definitions that are appropriate for use versus not
	Status ResearchDefinitionStatus `json:"status"`

	// The subject of the ResearchDefinition is critical in interpreting the criteria definitions
	SubjectCodeableConcept *common.CodeableConcept `json:"subjectCodeableConcept,omitempty"`

	// The subject of the ResearchDefinition is critical in interpreting the criteria definitions
	SubjectReference *common.Reference `json:"subjectReference,omitempty"`

	// An explanatory or alternate title for the ResearchDefinition giving additional information about its content
	Subtitle *string `json:"subtitle,omitempty"`

	// This name does not need to be machine-processing friendly and may contain punctuation, white-space, etc
	Title *string `json:"title,omitempty"`

	// Descriptive topics related to the content of the ResearchDefinition
	Topic []common.CodeableConcept `json:"topic,omitempty"`

	// Can be a urn:uuid: or a urn:oid: but real http: addresses are preferred
	Url *string `json:"url,omitempty"`

	// A detailed description, from a clinical perspective, of how the ResearchDefinition is used
	Usage *string `json:"usage,omitempty"`

	// When multiple useContexts are specified, there is no expectation that all or any of the contexts apply
	UseContext []common.UsageContext `json:"useContext,omitempty"`

	// There may be different research definition instances that have the same identifier but different versions
	Version *string `json:"version,omitempty"`
}

// ResearchDefinitionStatus represents allows filtering of research definitions that are appropriate for use versus not
type ResearchDefinitionStatus string

const (
	ResearchDefinitionStatusDraft   ResearchDefinitionStatus = "draft"
	ResearchDefinitionStatusActive  ResearchDefinitionStatus = "active"
	ResearchDefinitionStatusRetired ResearchDefinitionStatus = "retired"
	ResearchDefinitionStatusUnknown ResearchDefinitionStatus = "unknown"
)
