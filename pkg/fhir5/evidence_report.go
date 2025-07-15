// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// EvidenceReportSubjectCharacteristic represents a characteristic of the subject
type EvidenceReportSubjectCharacteristic struct {
	common.BackboneElement

	// Example 1 is a Citation. Example 2 is a type of outcome. Example 3 is a specific outcome
	Code common.CodeableConcept `json:"code"`

	// Is used to express not the characteristic
	Exclude *bool `json:"exclude,omitempty"`

	// Timeframe for the characteristic
	Period *common.Period `json:"period,omitempty"`

	// Example 1 is Citation #37. Example 2 is selecting clinical outcomes. Example 3 is 1-year mortality
	ValueReference       *common.Reference       `json:"valueReference,omitempty"`
	ValueCodeableConcept *common.CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueBoolean         *bool                   `json:"valueBoolean,omitempty"`
	ValueQuantity        *common.Quantity        `json:"valueQuantity,omitempty"`
	ValueRange           *Range                  `json:"valueRange,omitempty"`
}

// EvidenceReportSubject represents the subject of the evidence report
type EvidenceReportSubject struct {
	common.BackboneElement

	// Characteristic
	Characteristic []EvidenceReportSubjectCharacteristic `json:"characteristic,omitempty"`

	// Used for general notes and annotations not coded elsewhere
	Note []Annotation `json:"note,omitempty"`
}

// EvidenceReportRelatesToTarget represents the target of a relationship
type EvidenceReportRelatesToTarget struct {
	common.BackboneElement

	// Target of the relationship Display
	Display *string `json:"display,omitempty"`

	// Target of the relationship Identifier
	Identifier *common.Identifier `json:"identifier,omitempty"`

	// Target of the relationship Resource reference
	Resource *common.Reference `json:"resource,omitempty"`

	// Target of the relationship URL
	URL *string `json:"url,omitempty"`
}

// EvidenceReportRelatesTo represents relationships to other documents
type EvidenceReportRelatesTo struct {
	common.BackboneElement

	// replaces | amends | appends | transforms | replacedWith | amendedWith | appendedWith | transformedWith
	Code EvidenceReportRelatesToCode `json:"code"`

	// The target composition/document of this relationship
	Target EvidenceReportRelatesToTarget `json:"target"`
}

// EvidenceReportRelatesToCode represents the relationship type
type EvidenceReportRelatesToCode string

const (
	EvidenceReportRelatesToCodeReplaces        EvidenceReportRelatesToCode = "replaces"
	EvidenceReportRelatesToCodeAmends          EvidenceReportRelatesToCode = "amends"
	EvidenceReportRelatesToCodeAppends         EvidenceReportRelatesToCode = "appends"
	EvidenceReportRelatesToCodeTransforms      EvidenceReportRelatesToCode = "transforms"
	EvidenceReportRelatesToCodeReplacedWith    EvidenceReportRelatesToCode = "replacedWith"
	EvidenceReportRelatesToCodeAmendedWith     EvidenceReportRelatesToCode = "amendedWith"
	EvidenceReportRelatesToCodeAppendedWith    EvidenceReportRelatesToCode = "appendedWith"
	EvidenceReportRelatesToCodeTransformedWith EvidenceReportRelatesToCode = "transformedWith"
)

// EvidenceReportSection represents a section of the evidence report
type EvidenceReportSection struct {
	common.BackboneElement

	// Identifies who is responsible for the information in this section
	Author []common.Reference `json:"author,omitempty"`

	// The various reasons for an empty section
	EmptyReason *common.CodeableConcept `json:"emptyReason,omitempty"`

	// Specifies any type of classification of the evidence report
	EntryClassifier []common.CodeableConcept `json:"entryClassifier,omitempty"`

	// Quantity as content
	EntryQuantity []common.Quantity `json:"entryQuantity,omitempty"`

	// If there are no entries in the list, an emptyReason SHOULD be provided
	EntryReference []common.Reference `json:"entryReference,omitempty"`

	// The code identifies the section for an automated processor
	Focus *common.CodeableConcept `json:"focus,omitempty"`

	// A definitional Resource identifying the kind of content contained within the section
	FocusReference *common.Reference `json:"focusReference,omitempty"`

	// working | snapshot | changes
	Mode *EvidenceReportSectionMode `json:"mode,omitempty"`

	// Applications SHOULD render ordered lists in the order provided
	OrderedBy *common.CodeableConcept `json:"orderedBy,omitempty"`

	// Nested sections are primarily used to help human readers navigate
	Section []EvidenceReportSection `json:"section,omitempty"`

	// Document profiles may define what content should be represented in the narrative
	Text *Narrative `json:"text,omitempty"`

	// The title identifies the section for a human reader
	Title *string `json:"title,omitempty"`
}

// EvidenceReportSectionMode represents the mode of a section
type EvidenceReportSectionMode string

const (
	EvidenceReportSectionModeWorking  EvidenceReportSectionMode = "working"
	EvidenceReportSectionModeSnapshot EvidenceReportSectionMode = "snapshot"
	EvidenceReportSectionModeChanges  EvidenceReportSectionMode = "changes"
)

// EvidenceReport represents a specialized container for collections of evidence
type EvidenceReport struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "EvidenceReport"

	// Extensions to ContactDetail include: contactReference, contactAddress, and contributionTime
	Author []ContactDetail `json:"author,omitempty"`

	// Used for reports for which external citation is expected
	CiteAsReference *common.Reference `json:"citeAsReference,omitempty"`
	CiteAsMarkdown  *string           `json:"citeAsMarkdown,omitempty"`

	// Extensions to ContactDetail include: contactReference, contactAddress, and contributionTime
	Contact []ContactDetail `json:"contact,omitempty"`

	// Extensions to ContactDetail include: contactReference, contactAddress, and contributionTime
	Editor []ContactDetail `json:"editor,omitempty"`

	// Extensions to ContactDetail include: contactReference, contactAddress, and contributionTime
	Endorser []ContactDetail `json:"endorser,omitempty"`

	// This element will contain unique identifiers that support de-duplication of EvidenceReports
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Used for footnotes and annotations
	Note []Annotation `json:"note,omitempty"`

	// Usually an organization but may be an individual
	Publisher *string `json:"publisher,omitempty"`

	// Link, description or reference to artifact associated with the report
	RelatedArtifact []RelatedArtifact `json:"relatedArtifact,omitempty"`

	// May include trial registry identifiers
	RelatedIdentifier []common.Identifier `json:"relatedIdentifier,omitempty"`

	// A document is a version specific composition
	RelatesTo []EvidenceReportRelatesTo `json:"relatesTo,omitempty"`

	// Extensions to ContactDetail include: contactReference, contactAddress, and contributionTime
	Reviewer []ContactDetail `json:"reviewer,omitempty"`

	// The root of the sections that make up the composition
	Section []EvidenceReportSection `json:"section,omitempty"`

	// Allows filtering of summaries that are appropriate for use versus not
	Status EventDefinitionStatus `json:"status"`

	// May be used as an expression for search queries and search results
	Subject EvidenceReportSubject `json:"subject"`

	// Specifies the kind of report
	Type *common.CodeableConcept `json:"type,omitempty"`

	// Can be a urn:uuid: or a urn:oid: but real http: addresses are preferred
	URL *string `json:"url,omitempty"`

	// When multiple useContexts are specified, there is no expectation that all or any of the contexts apply
	UseContext []UsageContext `json:"useContext,omitempty"`
}
