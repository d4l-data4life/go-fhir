package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// LibraryStatus represents the status of a library
type LibraryStatus string

const (
	LibraryStatusDraft   LibraryStatus = "draft"
	LibraryStatusActive  LibraryStatus = "active"
	LibraryStatusRetired LibraryStatus = "retired"
	LibraryStatusUnknown LibraryStatus = "unknown"
)

// Library represents a general-purpose container for knowledge asset definitions
type Library struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Library"

	// The 'date' element may be more recent than the approval date because of minor changes or editorial corrections
	ApprovalDate *string `json:"approvalDate,omitempty"`

	// An individual or organization primarily involved in the creation and maintenance of the content
	Author []ContactDetail `json:"author,omitempty"`

	// May be a web site, an email address, a telephone number, etc.
	Contact []ContactDetail `json:"contact,omitempty"`

	// The content of the library as an Attachment
	Content []Attachment `json:"content,omitempty"`

	// Use and/or publishing restrictions
	Copyright *string `json:"copyright,omitempty"`

	// Copyright holder and year(s)
	CopyrightLabel *string `json:"copyrightLabel,omitempty"`

	// Describes a set of data that must be provided in order to be able to successfully perform the computations
	DataRequirement []interface{} `json:"dataRequirement,omitempty"`

	// Date last changed
	Date *string `json:"date,omitempty"`

	// Natural language description of the library
	Description *string `json:"description,omitempty"`

	// An individual or organization primarily responsible for internal coherence of the content
	Editor []ContactDetail `json:"editor,omitempty"`

	// When the library is expected to be used
	EffectivePeriod *common.Period `json:"effectivePeriod,omitempty"`

	// Who endorsed the content
	Endorser []ContactDetail `json:"endorser,omitempty"`

	// For testing purposes, not real usage
	Experimental *bool `json:"experimental,omitempty"`

	// Additional identifier for the library
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Intended jurisdiction for library (if applicable)
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`

	// When the library was last reviewed
	LastReviewDate *string `json:"lastReviewDate,omitempty"`

	// Name for this library (computer friendly)
	Name *string `json:"name,omitempty"`

	// The parameter element defines parameters used by the library
	Parameter []interface{} `json:"parameter,omitempty"`

	// Name of the publisher (organization or individual)
	Publisher *string `json:"publisher,omitempty"`

	// Why this library is defined
	Purpose *string `json:"purpose,omitempty"`

	// Additional documentation, citations, etc.
	RelatedArtifact []interface{} `json:"relatedArtifact,omitempty"`

	// Who reviewed the content
	Reviewer []ContactDetail `json:"reviewer,omitempty"`

	// draft | active | retired | unknown
	Status LibraryStatus `json:"status"`

	// Type of individual the library content is focused on
	SubjectCodeableConcept *common.CodeableConcept `json:"subjectCodeableConcept,omitempty"`
	SubjectReference       *common.Reference       `json:"subjectReference,omitempty"`

	// Subordinate title of the library
	Subtitle *string `json:"subtitle,omitempty"`

	// Name for this library (human friendly)
	Title *string `json:"title,omitempty"`

	// The category of the library, such as education, treatment, assessment, etc.
	Topic []common.CodeableConcept `json:"topic,omitempty"`

	// Identifies the type of library such as a Logic Library, Model Definition, Asset Collection, or Module Definition
	Type common.CodeableConcept `json:"type"`

	// Canonical identifier for this library, represented as a URI
	Url *string `json:"url,omitempty"`

	// Describes the clinical usage of the library
	Usage *string `json:"usage,omitempty"`

	// The context that the content is intended to support
	UseContext []interface{} `json:"useContext,omitempty"`

	// Business version of the library
	Version *string `json:"version,omitempty"`

	// How to compare versions
	VersionAlgorithmString *string        `json:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *common.Coding `json:"versionAlgorithmCoding,omitempty"`
}
