// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// FormularyItemStatus represents the status of a formulary item
type FormularyItemStatus string

const (
	FormularyItemStatusActive         FormularyItemStatus = "active"
	FormularyItemStatusEnteredInError FormularyItemStatus = "entered-in-error"
	FormularyItemStatusInactive       FormularyItemStatus = "inactive"
)

// FormularyItem represents a product or service that is available through a program
type FormularyItem struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "FormularyItem"

	// A code (or set of codes) that specify the product or service that is identified by this formulary item
	Code *common.CodeableConcept `json:"code,omitempty"`

	// Business identifier for this formulary item
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// This status is intended to identify if the formulary item in a local system is in active use within the formulary
	Status *FormularyItemStatus `json:"status,omitempty"`
}
