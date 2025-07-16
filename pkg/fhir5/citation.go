// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// CitationSummary represents a human-readable display of key concepts to represent the citation
type CitationSummary struct {
	common.BackboneElement

	// Format for display of the citation summary
	Style *common.CodeableConcept `json:"style,omitempty"`

	// The human-readable display of the citation summary
	Text string `json:"text"`
}

// CitationClassification represents classification of the citation record
type CitationClassification struct {
	common.BackboneElement

	// The specific classification value
	Classifier []common.CodeableConcept `json:"classifier,omitempty"`

	// The kind of classifier (e.g. publication type, keyword)
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// CitationStatusDate represents the state or status of the citation record
type CitationStatusDate struct {
	common.BackboneElement

	// The state or status of the citation record
	Activity common.CodeableConcept `json:"activity"`

	// Whether the status date is actual (has occurred) or expected (estimated or anticipated)
	Actual *bool `json:"actual,omitempty"`

	// For an instance, place the same value in both start and end elements
	Period common.Period `json:"period"`
}

// CitationCitedArtifactVersion represents the defined version of the cited artifact
type CitationCitedArtifactVersion struct {
	common.BackboneElement

	// When referencing a baseCitation, one may inherit any data from the referenced Citation Resource
	BaseCitation *common.Reference `json:"baseCitation,omitempty"`

	// The version number or other version identifier
	Value string `json:"value"`
}

// CitationCitedArtifact represents the article or artifact being described
type CitationCitedArtifact struct {
	common.BackboneElement

	// The defined version of the cited artifact
	Version *CitationCitedArtifactVersion `json:"version,omitempty"`

	// The current state of the cited artifact
	CurrentState []common.CodeableConcept `json:"currentState,omitempty"`

	// An effective date or period, historical or future, actual or expected
	StatusDate []CitationStatusDate `json:"statusDate,omitempty"`

	// The title details of the cited artifact
	Title []CitationCitedArtifactTitle `json:"title,omitempty"`

	// Summary of the article or artifact
	Abstract []CitationCitedArtifactAbstract `json:"abstract,omitempty"`

	// The component of the cited artifact
	Part []CitationCitedArtifactPart `json:"part,omitempty"`

	// The publication form of the cited artifact
	PublicationForm []CitationCitedArtifactPublicationForm `json:"publicationForm,omitempty"`

	// The web location of the cited artifact
	WebLocation []CitationCitedArtifactWebLocation `json:"webLocation,omitempty"`

	// The classification of the cited artifact
	Classification []CitationCitedArtifactClassification `json:"classification,omitempty"`

	// The contributorship of the cited artifact
	Contributorship *CitationCitedArtifactContributorship `json:"contributorship,omitempty"`

	// The related artifacts
	RelatesTo []CitationCitedArtifactRelatesTo `json:"relatesTo,omitempty"`
}

// CitationCitedArtifactTitle represents the title details of the cited artifact
type CitationCitedArtifactTitle struct {
	common.BackboneElement

	// The kind of title
	Type *common.CodeableConcept `json:"type,omitempty"`

	// Used to express the specific language
	Language *common.CodeableConcept `json:"language,omitempty"`

	// The title of the article or artifact
	Text string `json:"text"`
}

// CitationCitedArtifactAbstract represents summary of the article or artifact
type CitationCitedArtifactAbstract struct {
	common.BackboneElement

	// Used to express the specific language
	Language *common.CodeableConcept `json:"language,omitempty"`

	// Copyright holder for this abstract
	Copyright *string `json:"copyright,omitempty"`

	// Abstract content
	Text string `json:"text"`

	// Used for the display of the abstract
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// CitationCitedArtifactPart represents the component of the cited artifact
type CitationCitedArtifactPart struct {
	common.BackboneElement

	// The kind of component
	Type *common.CodeableConcept `json:"type,omitempty"`

	// The specification of the component
	Value *string `json:"value,omitempty"`

	// The citation for the full article or artifact
	BaseCitation *common.Reference `json:"baseCitation,omitempty"`
}

// CitationCitedArtifactPublicationForm represents the publication form of the cited artifact
type CitationCitedArtifactPublicationForm struct {
	common.BackboneElement

	// The collection the cited article or artifact is published in
	PublishedIn *CitationCitedArtifactPublicationFormPublishedIn `json:"publishedIn,omitempty"`

	// The specific issue in which the cited article resides
	PeriodicRelease *CitationCitedArtifactPublicationFormPeriodicRelease `json:"periodicRelease,omitempty"`

	// The date the article was added to the database
	ArticleDate *string `json:"articleDate,omitempty"`

	// The date the article was last revised or updated in the database
	LastRevisionDate *string `json:"lastRevisionDate,omitempty"`

	// Language in which the form is written
	Language []common.CodeableConcept `json:"language,omitempty"`

	// The accession number for the publication
	AccessionNumber *string `json:"accessionNumber,omitempty"`

	// The page numbers of the publication
	PageString *string `json:"pageString,omitempty"`

	// The page numbers of the publication
	FirstPage *string `json:"firstPage,omitempty"`

	// The page numbers of the publication
	LastPage *string `json:"lastPage,omitempty"`

	// The page numbers of the publication
	PageCount *string `json:"pageCount,omitempty"`

	// Copyright notice for the full article or artifact
	Copyright *string `json:"copyright,omitempty"`
}

// CitationCitedArtifactPublicationFormPublishedIn represents the collection the cited article is published in
type CitationCitedArtifactPublicationFormPublishedIn struct {
	common.BackboneElement

	// Kind of container (e.g. Periodical, database, or book)
	Type *common.CodeableConcept `json:"type,omitempty"`

	// Name of the database or title of the book or journal
	Name *string `json:"name,omitempty"`

	// Name of the publisher
	Publisher *common.Reference `json:"publisher,omitempty"`

	// Geographic location of the publisher
	PublisherLocation *string `json:"publisherLocation,omitempty"`

	// When the article was published
	PublicationDate *string `json:"publicationDate,omitempty"`

	// When the article was published
	PublicationDateString *string `json:"publicationDateString,omitempty"`
}

// CitationCitedArtifactPublicationFormPeriodicRelease represents the specific issue in which the cited article resides
type CitationCitedArtifactPublicationFormPeriodicRelease struct {
	common.BackboneElement

	// Designation of the issue
	CitedMedium *common.CodeableConcept `json:"citedMedium,omitempty"`

	// Volume number of journal in which the article is published
	Volume *string `json:"volume,omitempty"`

	// Issue, part or supplement of journal in which the article is published
	Issue *string `json:"issue,omitempty"`

	// Date of publication of the issue
	DateOfPublication *CitationCitedArtifactPublicationFormPeriodicReleaseDateOfPublication `json:"dateOfPublication,omitempty"`
}

// CitationCitedArtifactPublicationFormPeriodicReleaseDateOfPublication represents the date of publication of the issue
type CitationCitedArtifactPublicationFormPeriodicReleaseDateOfPublication struct {
	common.BackboneElement

	// Date on which the issue of the journal was published
	Date *string `json:"date,omitempty"`

	// Year on which the issue of the journal was published
	Year *string `json:"year,omitempty"`

	// Month on which the issue of the journal was published
	Month *string `json:"month,omitempty"`

	// Day on which the issue of the journal was published
	Day *string `json:"day,omitempty"`

	// Spring, Summer, Fall/Autumn, Winter
	Season *string `json:"season,omitempty"`

	// Text representation of the date of which the issue of the journal was published
	Text *string `json:"text,omitempty"`
}

// CitationCitedArtifactWebLocation represents the web location of the cited artifact
type CitationCitedArtifactWebLocation struct {
	common.BackboneElement

	// Code the reason for different URLs, e.g. abstract and full-text
	Classifier []common.CodeableConcept `json:"classifier,omitempty"`

	// The specific URL
	Url *string `json:"url,omitempty"`
}

// CitationCitedArtifactClassification represents the classification of the cited artifact
type CitationCitedArtifactClassification struct {
	common.BackboneElement

	// The kind of classifier (e.g. publication type, keyword)
	Type *common.CodeableConcept `json:"type,omitempty"`

	// The specific classification value
	Classifier []common.CodeableConcept `json:"classifier,omitempty"`
}

// CitationCitedArtifactContributorship represents the contributorship of the cited artifact
type CitationCitedArtifactContributorship struct {
	common.BackboneElement

	// Indicates if the list includes all authors and/or contributors
	Complete *bool `json:"complete,omitempty"`

	// An individual entity named as a contributor
	Entry []CitationCitedArtifactContributorshipEntry `json:"entry,omitempty"`

	// Used to record a display of the author/contributor list without separate coding for each list member
	Summary []CitationCitedArtifactContributorshipSummary `json:"summary,omitempty"`
}

// CitationCitedArtifactContributorshipEntry represents an individual entity named as a contributor
type CitationCitedArtifactContributorshipEntry struct {
	common.BackboneElement

	// The identity of the individual entity
	Contributor *common.Reference `json:"contributor,omitempty"`

	// For citation of contributorship list where the order is important
	ForenameInitials *string `json:"forenameInitials,omitempty"`

	// For citation of contributorship list where the order is important
	SurnameInitials *string `json:"surnameInitials,omitempty"`

	// The initials for forename if used to disambiguate the contributor
	Initials *string `json:"initials,omitempty"`

	// Used to code contribution based on appropriate role taxonomy
	ContributionType []common.CodeableConcept `json:"contributionType,omitempty"`

	// The role of the contributor (e.g. author, editor, reviewer)
	Role *common.CodeableConcept `json:"role,omitempty"`

	// Contributions with accounting for time or number
	ContributionInstance []CitationCitedArtifactContributorshipEntryContributionInstance `json:"contributionInstance,omitempty"`

	// Indication of which contributor is the corresponding contributor for the role
	CorrespondingContact *bool `json:"correspondingContact,omitempty"`

	// Ranked order of contribution e.g. 1, 2, 3, etc
	ListOrder *int `json:"listOrder,omitempty"`
}

// CitationCitedArtifactContributorshipEntryContributionInstance represents contributions with accounting for time or number
type CitationCitedArtifactContributorshipEntryContributionInstance struct {
	common.BackboneElement

	// The specific time
	Time *string `json:"time,omitempty"`

	// The amount of time
	TimeRange *common.Period `json:"timeRange,omitempty"`

	// The number of occurrences
	Count *int `json:"count,omitempty"`

	// The number of occurrences
	CountRange *Range `json:"countRange,omitempty"`

	// The type of contribution
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// CitationCitedArtifactContributorshipSummary represents a display of the author/contributor list
type CitationCitedArtifactContributorshipSummary struct {
	common.BackboneElement

	// Used most commonly to express an author list or a contributorship statement
	Type *common.CodeableConcept `json:"type,omitempty"`

	// The format for the display of the citation
	Style *common.CodeableConcept `json:"style,omitempty"`

	// Used to code the producer or rule for creating the display string
	Source *common.CodeableConcept `json:"source,omitempty"`

	// The display string for the author list, contributor list, or contributorship statement
	Value string `json:"value"`
}

// CitationCitedArtifactRelatesTo represents the related artifacts
type CitationCitedArtifactRelatesTo struct {
	common.BackboneElement

	// How the cited artifact relates to the target artifact
	Type *common.CodeableConcept `json:"type,omitempty"`

	// The article or artifact that the cited artifact is related to
	TargetClassifier []common.CodeableConcept `json:"targetClassifier,omitempty"`

	// The article or artifact that the cited artifact is related to
	TargetUri *string `json:"targetUri,omitempty"`

	// The article or artifact that the cited artifact is related to
	TargetIdentifier *common.Identifier `json:"targetIdentifier,omitempty"`

	// The article or artifact that the cited artifact is related to
	TargetReference *common.Reference `json:"targetReference,omitempty"`

	// The article or artifact that the cited artifact is related to
	TargetAttachment *Attachment `json:"targetAttachment,omitempty"`
}

// Citation represents a citation
type Citation struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Citation"

	// The 'date' element may be more recent than the approval date
	ApprovalDate *string `json:"approvalDate,omitempty"`

	// The Citation.author element is structured to support a name and contact point information
	Author []ContactDetail `json:"author,omitempty"`

	// The article or artifact being described
	CitedArtifact *CitationCitedArtifact `json:"citedArtifact,omitempty"`

	// Use this element if you need to classify the citation record independently
	Classification []CitationClassification `json:"classification,omitempty"`

	// May be a web site, an email address, a telephone number, etc.
	Contact []ContactDetail `json:"contact,omitempty"`

	// Use and/or publishing restrictions for the citation record
	Copyright *string `json:"copyright,omitempty"`

	// The (c) symbol should NOT be included in this string
	CopyrightLabel *string `json:"copyrightLabel,omitempty"`

	// Use this if needed for reporting the state or status of the citation record
	CurrentState []common.CodeableConcept `json:"currentState,omitempty"`

	// The date is often not tracked until the resource is published
	Date *string `json:"date,omitempty"`

	// This description can be used to capture details such as comments about misuse
	Description *string `json:"description,omitempty"`

	// The Citation.editor element is structured to support a name and contact point information
	Editor []ContactDetail `json:"editor,omitempty"`

	// The effective period for a citation record determines when the content is applicable for usage
	EffectivePeriod *common.Period `json:"effectivePeriod,omitempty"`

	// The Citation.endorser element is structured to support a name and contact point information
	Endorser []ContactDetail `json:"endorser,omitempty"`

	// Allows filtering of citation records that are appropriate for use versus not
	Experimental *bool `json:"experimental,omitempty"`

	// Use this element if you need to identify the citation record independently
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// It may be possible for the citation record to be used in jurisdictions other than those for which it was originally designed
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`

	// If specified, this date follows the original approval date
	LastReviewDate *string `json:"lastReviewDate,omitempty"`

	// The name is not expected to be globally unique
	Name *string `json:"name,omitempty"`

	// Used for general notes and annotations not coded elsewhere
	Note []Annotation `json:"note,omitempty"`

	// Usually an organization but may be an individual
	Publisher *string `json:"publisher,omitempty"`

	// This element does not describe the usage of the citation
	Purpose *string `json:"purpose,omitempty"`

	// Use this if needed for reporting artifacts related to the citation record
	RelatedArtifact []RelatedArtifact `json:"relatedArtifact,omitempty"`

	// The Citation.reviewer element is structured to support a name and contact point information
	Reviewer []ContactDetail `json:"reviewer,omitempty"`

	// Allows filtering of summaries that are appropriate for use versus not
	Status CitationStatus `json:"status"`

	// Use this if needed for reporting the state or status of the citation record
	StatusDate []CitationStatusDate `json:"statusDate,omitempty"`

	// A human-readable display of key concepts to represent the citation
	Summary []CitationSummary `json:"summary,omitempty"`

	// This name does not need to be machine-processing friendly
	Title *string `json:"title,omitempty"`

	// In some cases, the resource can no longer be found at the stated url
	Url *string `json:"url,omitempty"`

	// When multiple useContexts are specified, there is no expectation that all or any of the contexts apply
	UseContext []UsageContext `json:"useContext,omitempty"`

	// There may be different citation record instances that have the same identifier but different versions
	Version *string `json:"version,omitempty"`

	// If set as a string, this is a FHIRPath expression
	VersionAlgorithmString *string `json:"versionAlgorithmString,omitempty"`

	// If set as a string, this is a FHIRPath expression
	VersionAlgorithmCoding *common.Coding `json:"versionAlgorithmCoding,omitempty"`
}

// CitationStatus represents the status of a citation
type CitationStatus string

const (
	CitationStatusDraft   CitationStatus = "draft"
	CitationStatusActive  CitationStatus = "active"
	CitationStatusRetired CitationStatus = "retired"
	CitationStatusUnknown CitationStatus = "unknown"
)
