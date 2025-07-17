package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// ListMode represents the mode of a list
type ListMode string

const (
	ListModeWorking  ListMode = "working"
	ListModeSnapshot ListMode = "snapshot"
	ListModeChanges  ListMode = "changes"
)

// ListStatus represents the status of a list
type ListStatus string

const (
	ListStatusCurrent        ListStatus = "current"
	ListStatusRetired        ListStatus = "retired"
	ListStatusEnteredInError ListStatus = "entered-in-error"
)

// ListEntry represents an entry in a list
type ListEntry struct {
	common.BackboneElement

	// When this item was added to the list
	Date *string `json:"date,omitempty"`

	// If the flag means that the entry has actually been deleted from the list, the deleted element SHALL be true
	Deleted *bool `json:"deleted,omitempty"`

	// The flag can only be understood in the context of the List.code
	Flag *common.CodeableConcept `json:"flag,omitempty"`

	// A reference to the actual resource from which data was derived
	Item common.Reference `json:"item"`
}

// List represents a curated collection of resources
type List struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "List"

	// If there is no code, the purpose of the list is implied where it is used
	Code *common.CodeableConcept `json:"code,omitempty"`

	// The actual important date is the date of currency of the resources that were summarized
	Date *string `json:"date,omitempty"`

	// The various reasons for an empty list make a significant interpretation to its interpretation
	EmptyReason *common.CodeableConcept `json:"emptyReason,omitempty"`

	// The encounter that is the context in which this list was created
	Encounter *common.Reference `json:"encounter,omitempty"`

	// If there are no entries in the list, an emptyReason SHOULD be provided
	Entry []ListEntry `json:"entry,omitempty"`

	// Identifier for the List assigned for business purposes outside the context of FHIR
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// This element is labeled as a modifier because a change list must not be misunderstood as a complete list
	Mode ListMode `json:"mode"`

	// Comments that apply to the overall list
	Note []Annotation `json:"note,omitempty"`

	// Applications SHOULD render ordered lists in the order provided
	OrderedBy *common.CodeableConcept `json:"orderedBy,omitempty"`

	// The primary source is the entity that made the decisions what items are in the list
	Source *common.Reference `json:"source,omitempty"`

	// This element is labeled as a modifier because the status contains codes that mark the resource as not currently valid
	Status ListStatus `json:"status"`

	// Some purely arbitrary lists do not have a common subject, so this is optional
	Subject []common.Reference `json:"subject,omitempty"`

	// A label for the list assigned by the author
	Title *string `json:"title,omitempty"`
}
