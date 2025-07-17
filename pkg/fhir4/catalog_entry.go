package fhir4

import (
	"encoding/json"

	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// CatalogEntry represents a FHIR R4 CatalogEntry resource
type CatalogEntry struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"`

	Identifier               []common.Identifier        `json:"identifier,omitempty"`
	Type                     *common.CodeableConcept    `json:"type,omitempty"`
	Orderable                bool                       `json:"orderable"`
	ReferencedItem           common.Reference           `json:"referencedItem"`
	AdditionalIdentifier     []common.Identifier        `json:"additionalIdentifier,omitempty"`
	Classification           []common.CodeableConcept   `json:"classification,omitempty"`
	Status                   *string                    `json:"status,omitempty"`
	ValidFrom                *string                    `json:"validFrom,omitempty"`
	ValidTo                  *string                    `json:"validTo,omitempty"`
	LastUpdated              *string                    `json:"lastUpdated,omitempty"`
	AdditionalCharacteristic []common.CodeableConcept   `json:"additionalCharacteristic,omitempty"`
	AdditionalClassification []common.CodeableConcept   `json:"additionalClassification,omitempty"`
	RelatedEntry             []CatalogEntryRelatedEntry `json:"relatedEntry,omitempty"`
}

// CatalogEntryRelatedEntry represents related catalog entries
type CatalogEntryRelatedEntry struct {
	common.BackboneElement

	RelationType string           `json:"relationtype"`
	Item         common.Reference `json:"item"`
}

// MarshalJSON marshals the CatalogEntry to JSON
func (c *CatalogEntry) MarshalJSON() ([]byte, error) {
	type Alias CatalogEntry
	return json.Marshal(&struct {
		*Alias
		ResourceType string `json:"resourceType"`
	}{
		Alias:        (*Alias)(c),
		ResourceType: "CatalogEntry",
	})
}
