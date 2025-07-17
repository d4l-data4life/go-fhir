package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// Linkage represents identifies two or more records (resource instances) that refer to the same real-world "occurrence"
type Linkage struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Linkage"

	// If false, any asserted linkages should not be considered current/relevant/applicable
	Active *bool `json:"active,omitempty"`

	// Identifies the user or organization responsible for asserting the linkages
	Author *common.Reference `json:"author,omitempty"`

	// Identifies which record considered as the reference to the same real-world occurrence
	Item []LinkageItem `json:"item"`
}

// LinkageItem represents identifies which record considered as the reference to the same real-world occurrence
type LinkageItem struct {
	common.BackboneElement

	// The resource instance being linked as part of the group
	Resource common.Reference `json:"resource"`

	// Distinguishes which item is "source of truth" (if any) and which items are no longer considered to be current representations
	Type LinkageItemType `json:"type"`
}

// LinkageItemType represents the type of linkage item
type LinkageItemType string

const (
	LinkageItemTypeSource     LinkageItemType = "source"
	LinkageItemTypeAlternate  LinkageItemType = "alternate"
	LinkageItemTypeHistorical LinkageItemType = "historical"
)
