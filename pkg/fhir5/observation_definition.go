package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// ObservationDefinitionStatus represents the status of an observation definition
type ObservationDefinitionStatus string

const (
	ObservationDefinitionStatusDraft   ObservationDefinitionStatus = "draft"
	ObservationDefinitionStatusActive  ObservationDefinitionStatus = "active"
	ObservationDefinitionStatusRetired ObservationDefinitionStatus = "retired"
	ObservationDefinitionStatusUnknown ObservationDefinitionStatus = "unknown"
)

// ObservationDefinitionComponent represents some observations have multiple component observations
type ObservationDefinitionComponent struct {
	common.BackboneElement

	// Describes what will be observed
	Code common.CodeableConcept `json:"code"`

	// The data types allowed for the value element
	PermittedDataType []string `json:"permittedDataType,omitempty"`

	// Units allowed for the valueQuantity element
	PermittedUnit []common.Coding `json:"permittedUnit,omitempty"`

	// A set of qualified values associated with a context and conditions
	QualifiedValue []interface{} `json:"qualifiedValue,omitempty"`
}

// ObservationDefinition represents a definition of an observation
type ObservationDefinition struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "ObservationDefinition"

	// The date may be more recent than the approval date
	ApprovalDate *string `json:"approvalDate,omitempty"`

	// Only used if the defined observation is to be made directly on a body part
	BodySite *common.CodeableConcept `json:"bodySite,omitempty"`

	// This element allows various categorization schemes
	Category []common.CodeableConcept `json:"category,omitempty"`

	// Describes what will be observed
	Code common.CodeableConcept `json:"code"`

	// Some observations have multiple component observations
	Component []ObservationDefinitionComponent `json:"component,omitempty"`

	// May be a web site, an email address, a telephone number, etc.
	Contact []ContactDetail `json:"contact,omitempty"`

	// Copyright statement relating to the ObservationDefinition
	Copyright *string `json:"copyright,omitempty"`

	// The (c) symbol should NOT be included in this string
	CopyrightLabel *string `json:"copyrightLabel,omitempty"`

	// The date is often not tracked until the resource is published
	Date *string `json:"date,omitempty"`

	// The canonical URL pointing to another FHIR-defined ObservationDefinition
	DerivedFromCanonical []string `json:"derivedFromCanonical,omitempty"`

	// The URL pointing to an externally-defined observation definition
	DerivedFromUri []string `json:"derivedFromUri,omitempty"`

	// This description can be used to capture details
	Description *string `json:"description,omitempty"`

	// When multiple occurrences of device are present
	Device []common.Reference `json:"device,omitempty"`

	// The effective period for an ObservationDefinition
	EffectivePeriod *common.Period `json:"effectivePeriod,omitempty"`

	// Allows filtering of ObservationDefinition that are appropriate for use vs. not
	Experimental *bool `json:"experimental,omitempty"`

	// This ObservationDefinition defines a group observation
	HasMember []common.Reference `json:"hasMember,omitempty"`

	// This is a business identifier, not a resource identifier
	Identifier *common.Identifier `json:"identifier,omitempty"`

	// It may be possible for the ObservationDefinition to be used in jurisdictions
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`

	// If specified, this is usually after the approval date
	LastReviewDate *string `json:"lastReviewDate,omitempty"`

	// Only used if not implicit in observation code
	Method *common.CodeableConcept `json:"method,omitempty"`

	// An example of observation allowing multiple results
	MultipleResultsAllowed *bool `json:"multipleResultsAllowed,omitempty"`

	// The name is not expected to be globally unique
	Name *string `json:"name,omitempty"`

	// The type of individual/organization/device that is expected to act
	PerformerType *common.CodeableConcept `json:"performerType,omitempty"`

	// The data types allowed for the value element
	PermittedDataType []string `json:"permittedDataType,omitempty"`

	// Units allowed for the valueQuantity element
	PermittedUnit []common.Coding `json:"permittedUnit,omitempty"`

	// The preferred name to be used when reporting the results
	PreferredReportName *string `json:"preferredReportName,omitempty"`

	// Helps establish the "authority/credibility" of the ObservationDefinition
	Publisher *string `json:"publisher,omitempty"`

	// This element does not describe the usage of the ObservationDefinition
	Purpose *string `json:"purpose,omitempty"`

	// A set of qualified values associated with a context and conditions
	QualifiedValue []interface{} `json:"qualifiedValue,omitempty"`

	// Only used for in vitro observations
	Specimen []common.Reference `json:"specimen,omitempty"`

	// draft | active | retired | unknown
	Status ObservationDefinitionStatus `json:"status"`

	// Examples: person, animal, device, air, surface
	Subject []common.CodeableConcept `json:"subject,omitempty"`

	// A short, descriptive, user-friendly title
	Title *string `json:"title,omitempty"`

	// Can be a urn:uuid: or a urn:oid:, but real http: addresses are preferred
	Url *string `json:"url,omitempty"`

	// When multiple useContexts are specified
	UseContext []UsageContext `json:"useContext,omitempty"`

	// There may be different observation definitions that have the same url
	Version *string `json:"version,omitempty"`
}
