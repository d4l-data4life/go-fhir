package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// MedicinalProductAuthorization represents the regulatory authorization of a medicinal product
type MedicinalProductAuthorization struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "MedicinalProductAuthorization"

	// The country in which the marketing authorization has been granted
	Country []common.CodeableConcept `json:"country,omitempty"`

	// A period of time after authorization before generic product applications can be submitted
	DataExclusivityPeriod *common.Period `json:"dataExclusivityPeriod,omitempty"`

	// The date when the first authorization was granted by a Medicines Regulatory Agency
	DateOfFirstAuthorization *string `json:"dateOfFirstAuthorization,omitempty"`

	// Marketing Authorization Holder
	Holder *common.Reference `json:"holder,omitempty"`

	// Business identifier for the marketing authorization, as assigned by a regulator
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Date of first marketing authorization for a company's new medicinal product in any country in the World
	InternationalBirthDate *string `json:"internationalBirthDate,omitempty"`

	// Jurisdiction within a country
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`

	// Authorization in areas within a country
	JurisdictionalAuthorization []MedicinalProductAuthorizationJurisdictionalAuthorization `json:"jurisdictionalAuthorization,omitempty"`

	// The legal framework against which this authorization is granted
	LegalBasis *common.CodeableConcept `json:"legalBasis,omitempty"`

	// The regulatory procedure for granting or amending a marketing authorization
	Procedure *MedicinalProductAuthorizationProcedure `json:"procedure,omitempty"`

	// Medicines Regulatory Agency
	Regulator *common.Reference `json:"regulator,omitempty"`

	// The date when a suspended the marketing or the marketing authorization of the product is anticipated to be restored
	RestoreDate *string `json:"restoreDate,omitempty"`

	// The status of the marketing authorization
	Status *common.CodeableConcept `json:"status,omitempty"`

	// The date at which the given status has become applicable
	StatusDate *string `json:"statusDate,omitempty"`

	// The medicinal product that is being authorized
	Subject *common.Reference `json:"subject,omitempty"`

	// The beginning of the time period in which the marketing authorization is in the specific status
	ValidityPeriod *common.Period `json:"validityPeriod,omitempty"`
}

// MedicinalProductAuthorizationJurisdictionalAuthorization represents authorization in areas within a country
type MedicinalProductAuthorizationJurisdictionalAuthorization struct {
	common.BackboneElement

	// Country of authorization
	Country *common.CodeableConcept `json:"country,omitempty"`

	// The assigned number for the marketing authorization
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Jurisdiction within a country
	Jurisdiction []common.CodeableConcept `json:"jurisdiction,omitempty"`

	// The legal status of supply in a jurisdiction or region
	LegalStatusOfSupply *common.CodeableConcept `json:"legalStatusOfSupply,omitempty"`

	// The start and expected end date of the authorization
	ValidityPeriod *common.Period `json:"validityPeriod,omitempty"`
}

// MedicinalProductAuthorizationProcedure represents the regulatory procedure for granting or amending a marketing authorization
type MedicinalProductAuthorizationProcedure struct {
	common.BackboneElement

	// Applications submitted to obtain a marketing authorization
	Application []MedicinalProductAuthorizationProcedure `json:"application,omitempty"`

	// Date of procedure
	DatePeriod *common.Period `json:"datePeriod,omitempty"`

	// Date of procedure
	DateDateTime *string `json:"dateDateTime,omitempty"`

	// Identifier for this procedure
	Identifier *common.Identifier `json:"identifier,omitempty"`

	// Type of procedure
	Type common.CodeableConcept `json:"type"`
}
