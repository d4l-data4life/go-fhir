package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// Library represents the Library resource is a general-purpose container for knowledge asset definitions
type Library struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Library"

	// The 'date' element may be more recent than the approval date because of minor changes or editorial corrections
	ApprovalDate *string `json:"approvalDate,omitempty"`

	// An individual or organization primarily involved in the creation and maintenance of the content
	Author []common.ContactDetail `json:"author,omitempty"`

	// May be a web site, an email address, a telephone number, etc
	Contact []common.ContactDetail `json:"contact,omitempty"`

	// The content of the library as an Attachment
	Content []common.Attachment `json:"content,omitempty"`

	// A copyright statement relating to the library and/or its contents
	Copyright *string `json:"copyright,omitempty"`

	// Describes a set of data that must be provided in order to be able to successfully perform the computations
	DataRequirement []common.DataRequirement `json:"dataRequirement,omitempty"`

	// Note that this is not the same as the resource last-modified-date
	Date *string `json:"date,omitempty"`

	// This description can be used to capture details such as why the library was built
	Description *string `json:"description,omitempty"`

	// An individual or organization primarily responsible for internal coherence of the content
	Editor []common.ContactDetail `json:"editor,omitempty"`

	// The effective period for a library determines when the content is applicable for usage
	EffectivePeriod *common.Period `json:"effectivePeriod,omitempty"`

	// An individual or organization responsible for officially endorsing the content for use in some setting
	Endorser []common.ContactDetail `json:"endorser,omitempty"`

	// Allows filtering of libraries that are appropriate for use versus not
	Experimental *bool `json:"experimental,omitempty"`

	// Typically, this is used for identifiers that can go in an HL7 V3 II (instance identifier) data type
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// It may be possible for the library to be used in jurisdictions other than those for which it was originally designed
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`

	// If specified, this date follows the original approval date
	LastReviewDate *string `json:"lastReviewDate,omitempty"`

	// The name is not expected to be globally unique
	Name *string `json:"name,omitempty"`

	// The parameter element defines parameters used by the library
	Parameter []common.ParameterDefinition `json:"parameter,omitempty"`

	// Usually an organization but may be an individual
	Publisher *string `json:"publisher,omitempty"`

	// This element does not describe the usage of the library
	Purpose *string `json:"purpose,omitempty"`

	// Each related artifact is either an attachment, or a reference to another resource, but not both
	RelatedArtifact []common.RelatedArtifact `json:"relatedArtifact,omitempty"`

	// An individual or organization primarily responsible for review of some aspect of the content
	Reviewer []common.ContactDetail `json:"reviewer,omitempty"`

	// Allows filtering of libraries that are appropriate for use vs. not
	Status LibraryStatus `json:"status"`

	// A code or group definition that describes the intended subject of the contents of the library
	SubjectCodeableConcept *common.CodeableConcept `json:"subjectCodeableConcept,omitempty"`

	// A code or group definition that describes the intended subject of the contents of the library
	SubjectReference *common.Reference `json:"subjectReference,omitempty"`

	// An explanatory or alternate title for the library giving additional information about its content
	Subtitle *string `json:"subtitle,omitempty"`

	// This name does not need to be machine-processing friendly and may contain punctuation, white-space, etc
	Title *string `json:"title,omitempty"`

	// Descriptive topics related to the content of the library
	Topic []common.CodeableConcept `json:"topic,omitempty"`

	// Identifies the type of library such as a Logic Library, Model Definition, Asset Collection, or Module Definition
	Type common.CodeableConcept `json:"type"`

	// Can be a urn:uuid: or a urn:oid: but real http: addresses are preferred
	Url *string `json:"url,omitempty"`

	// A detailed description of how the library is used from a clinical perspective
	Usage *string `json:"usage,omitempty"`

	// When multiple useContexts are specified, there is no expectation that all or any of the contexts apply
	UseContext []common.UsageContext `json:"useContext,omitempty"`

	// There may be different library instances that have the same identifier but different versions
	Version *string `json:"version,omitempty"`
}

// LibraryStatus represents the status of the library
type LibraryStatus string

const (
	LibraryStatusDraft   LibraryStatus = "draft"
	LibraryStatusActive  LibraryStatus = "active"
	LibraryStatusRetired LibraryStatus = "retired"
	LibraryStatusUnknown LibraryStatus = "unknown"
)
