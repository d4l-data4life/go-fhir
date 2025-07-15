package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// InventoryReportCountType represents whether the report is about the current inventory count or a differential change
type InventoryReportCountType string

const (
	InventoryReportCountTypeSnapshot   InventoryReportCountType = "snapshot"
	InventoryReportCountTypeDifference InventoryReportCountType = "difference"
)

// InventoryReportStatus represents the status of the inventory check or notification
type InventoryReportStatus string

const (
	InventoryReportStatusDraft          InventoryReportStatus = "draft"
	InventoryReportStatusRequested      InventoryReportStatus = "requested"
	InventoryReportStatusActive         InventoryReportStatus = "active"
	InventoryReportStatusEnteredInError InventoryReportStatus = "entered-in-error"
)

// InventoryReportInventoryListingItem represents the item or items in this listing
type InventoryReportInventoryListingItem struct {
	common.BackboneElement

	// The inventory category or classification of the items being reported
	Category *common.CodeableConcept `json:"category,omitempty"`

	// The code or reference to the item type
	Item CodeableReference `json:"item"`

	// The quantity of the item or items being reported
	Quantity common.Quantity `json:"quantity"`
}

// InventoryReportInventoryListing represents an inventory listing section (grouped by any of the attributes)
type InventoryReportInventoryListing struct {
	common.BackboneElement

	// The date and time when the items were counted
	CountingDateTime *string `json:"countingDateTime,omitempty"`

	// The item or items in this listing
	Item []InventoryReportInventoryListingItem `json:"item,omitempty"`

	// The status of the items
	ItemStatus *common.CodeableConcept `json:"itemStatus,omitempty"`

	// Location of the inventory items
	Location *common.Reference `json:"location,omitempty"`
}

// InventoryReport represents a report of inventory items
type InventoryReport struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "InventoryReport"

	// Whether the report is about the current inventory count (snapshot) or a differential change in inventory (change)
	CountType InventoryReportCountType `json:"countType"`

	// Business identifier for the InvoiceReport
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// An inventory listing section (grouped by any of the attributes)
	InventoryListing []InventoryReportInventoryListing `json:"inventoryListing,omitempty"`

	// A note associated with the InvoiceReport
	Note []Annotation `json:"note,omitempty"`

	// What type of operation is being performed - addition or subtraction
	OperationType *common.CodeableConcept `json:"operationType,omitempty"`

	// The reason for this count - regular count, ad-hoc count, new arrivals, etc.
	OperationTypeReason *common.CodeableConcept `json:"operationTypeReason,omitempty"`

	// When the report has been submitted
	ReportedDateTime string `json:"reportedDateTime"`

	// Who submits the report
	Reporter *common.Reference `json:"reporter,omitempty"`

	// The period the report refers to
	ReportingPeriod *common.Period `json:"reportingPeriod,omitempty"`

	// The status of the inventory check or notification
	Status InventoryReportStatus `json:"status"`
}
