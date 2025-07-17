package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// MedicinalProductPharmaceutical represents a pharmaceutical product described in terms of its composition and dose form
type MedicinalProductPharmaceutical struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "MedicinalProductPharmaceutical"

	// The administrable dose form, after necessary reconstitution
	AdministrableDoseForm common.CodeableConcept `json:"administrableDoseForm"`

	// Characteristics e.g. a products onset of action
	Characteristics []MedicinalProductPharmaceuticalCharacteristics `json:"characteristics,omitempty"`

	// Accompanying device
	Device []common.Reference `json:"device,omitempty"`

	// An identifier for the pharmaceutical medicinal product
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Ingredient
	Ingredient []common.Reference `json:"ingredient,omitempty"`

	// The path by which the pharmaceutical product is taken into or makes contact with the body
	RouteOfAdministration []MedicinalProductPharmaceuticalRouteOfAdministration `json:"routeOfAdministration"`

	// Todo
	UnitOfPresentation *common.CodeableConcept `json:"unitOfPresentation,omitempty"`
}

// MedicinalProductPharmaceuticalCharacteristics represents characteristics e.g. a products onset of action
type MedicinalProductPharmaceuticalCharacteristics struct {
	common.BackboneElement

	// A coded characteristic
	Code common.CodeableConcept `json:"code"`

	// The status of characteristic e.g. assigned or pending
	Status *common.CodeableConcept `json:"status,omitempty"`
}

// MedicinalProductPharmaceuticalRouteOfAdministration represents the path by which the pharmaceutical product is taken into or makes contact with the body
type MedicinalProductPharmaceuticalRouteOfAdministration struct {
	common.BackboneElement

	// Coded expression for the route
	Code common.CodeableConcept `json:"code"`

	// The first dose (dose quantity) administered in humans can be specified, for a product under investigation
	FirstDose *common.Quantity `json:"firstDose,omitempty"`

	// The maximum dose per day (maximum dose quantity to be administered in any one 24-h period)
	MaxDosePerDay *common.Quantity `json:"maxDosePerDay,omitempty"`

	// The maximum dose per treatment period that can be administered as per the protocol
	MaxDosePerTreatmentPeriod *common.Ratio `json:"maxDosePerTreatmentPeriod,omitempty"`

	// The maximum single dose that can be administered as per the protocol of a clinical trial
	MaxSingleDose *common.Quantity `json:"maxSingleDose,omitempty"`

	// The maximum treatment period during which an Investigational Medicinal Product can be administered
	MaxTreatmentPeriod *common.Duration `json:"maxTreatmentPeriod,omitempty"`

	// A species for which this route applies
	TargetSpecies []MedicinalProductPharmaceuticalRouteOfAdministrationTargetSpecies `json:"targetSpecies,omitempty"`
}

// MedicinalProductPharmaceuticalRouteOfAdministrationTargetSpecies represents a species for which this route applies
type MedicinalProductPharmaceuticalRouteOfAdministrationTargetSpecies struct {
	common.BackboneElement

	// Coded expression for the species
	Code common.CodeableConcept `json:"code"`

	// A species specific time during which consumption of animal product is not appropriate
	WithdrawalPeriod []MedicinalProductPharmaceuticalRouteOfAdministrationTargetSpeciesWithdrawalPeriod `json:"withdrawalPeriod,omitempty"`
}

// MedicinalProductPharmaceuticalRouteOfAdministrationTargetSpeciesWithdrawalPeriod represents a species specific time during which consumption of animal product is not appropriate
type MedicinalProductPharmaceuticalRouteOfAdministrationTargetSpeciesWithdrawalPeriod struct {
	common.BackboneElement

	// Extra information about the withdrawal period
	SupportingInformation *string `json:"supportingInformation,omitempty"`

	// Coded expression for the type of tissue for which the withdrawal period applies, e.g. meat, milk
	Tissue common.CodeableConcept `json:"tissue"`

	// A value for the time
	Value common.Quantity `json:"value"`
}
