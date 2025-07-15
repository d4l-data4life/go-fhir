// Package fhir5 contains FHIR R5 resource definitions
package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// RegulatedAuthorization represents a regulatory authorization for a medicinal product
type RegulatedAuthorization struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "RegulatedAuthorization"

	// The case or regulatory procedure for granting or amending a regulated authorization
	Case *interface{} `json:"case,omitempty"`

	// The country in which the authorization has been granted
	Country []common.CodeableConcept `json:"country,omitempty"`

	// The date at which the current status was assigned
	DateOfStatus *string `json:"dateOfStatus,omitempty"`

	// The date when a suspended the marketing or the marketing authorization of the product is anticipated to be restored
	RestoreDate *string `json:"restoreDate,omitempty"`

	// The holder of the regulated authorization
	Holder *common.Reference `json:"holder,omitempty"`

	// Business identifier for the authorization, as assigned by a regulator
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The legal or regulatory framework against which this authorization is granted
	LegalBasis *common.CodeableConcept `json:"legalBasis,omitempty"`

	// The organization that has been granted this authorization, by some authoritative body
	Regulator *common.Reference `json:"regulator,omitempty"`

	// The status that is authorised e.g. approved
	Status *common.CodeableConcept `json:"status,omitempty"`

	// The date at which the given status has become applicable
	StatusDate *string `json:"statusDate,omitempty"`

	// The type of regulated authorization, e.g. marketing or pediatric
	Type *common.CodeableConcept `json:"type,omitempty"`

	// The date when the authorization was granted
	ValidFrom *string `json:"validFrom,omitempty"`

	// The date when the authorization expires
	ValidUntil *string `json:"validUntil,omitempty"`

	// The medicinal product that is being granted the authorization
	Subject []common.Reference `json:"subject,omitempty"`
}
