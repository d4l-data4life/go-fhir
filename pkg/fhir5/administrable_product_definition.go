// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// AdministrableProductDefinitionProperty represents characteristics of a product
type AdministrableProductDefinitionProperty struct {
	common.BackboneElement

	// The status of characteristic e.g. assigned or pending
	Status *common.CodeableConcept `json:"status,omitempty"`

	// A code expressing the type of characteristic
	Type common.CodeableConcept `json:"type"`

	// A value for the characteristic
	ValueCodeableConcept *common.CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueQuantity        *common.Quantity        `json:"valueQuantity,omitempty"`
	ValueDate            *string                 `json:"valueDate,omitempty"`
	ValueBoolean         *bool                   `json:"valueBoolean,omitempty"`
	ValueMarkdown        *string                 `json:"valueMarkdown,omitempty"`
	ValueAttachment      *Attachment             `json:"valueAttachment,omitempty"`
	ValueReference       *common.Reference       `json:"valueReference,omitempty"`
}

// AdministrableProductDefinitionRouteOfAdministrationTargetSpeciesWithdrawalPeriod represents a species specific time during which consumption of animal product is not appropriate
type AdministrableProductDefinitionRouteOfAdministrationTargetSpeciesWithdrawalPeriod struct {
	common.BackboneElement

	// Extra information about the withdrawal period
	SupportingInformation *string `json:"supportingInformation,omitempty"`

	// Coded expression for the type of tissue for which the withdrawal period applies
	Tissue common.CodeableConcept `json:"tissue"`

	// A value for the time
	Value common.Quantity `json:"value"`
}

// AdministrableProductDefinitionRouteOfAdministrationTargetSpecies represents a species for which this route applies
type AdministrableProductDefinitionRouteOfAdministrationTargetSpecies struct {
	common.BackboneElement

	// Coded expression for the species
	Code common.CodeableConcept `json:"code"`

	// A species specific time during which consumption of animal product is not appropriate
	WithdrawalPeriod []AdministrableProductDefinitionRouteOfAdministrationTargetSpeciesWithdrawalPeriod `json:"withdrawalPeriod,omitempty"`
}

// AdministrableProductDefinitionRouteOfAdministration represents the path by which the product is taken into or makes contact with the body
type AdministrableProductDefinitionRouteOfAdministration struct {
	common.BackboneElement

	// Coded expression for the route
	Code common.CodeableConcept `json:"code"`

	// The first dose (dose quantity) administered can be specified for the product
	FirstDose *common.Quantity `json:"firstDose,omitempty"`

	// The maximum dose per day that can be administered
	MaxDosePerDay *common.Quantity `json:"maxDosePerDay,omitempty"`

	// The maximum dose per treatment period that can be administered
	MaxDosePerTreatmentPeriod *Ratio `json:"maxDosePerTreatmentPeriod,omitempty"`

	// The maximum single dose that can be administered
	MaxSingleDose *common.Quantity `json:"maxSingleDose,omitempty"`

	// The maximum treatment period during which the product can be administered
	MaxTreatmentPeriod *Duration `json:"maxTreatmentPeriod,omitempty"`

	// A species for which this route applies
	TargetSpecies []AdministrableProductDefinitionRouteOfAdministrationTargetSpecies `json:"targetSpecies,omitempty"`
}

// AdministrableProductDefinition represents a medicinal product in the final form which is suitable for administering to a patient
type AdministrableProductDefinition struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "AdministrableProductDefinition"

	// The dose form of the final product after necessary reconstitution or processing
	AdministrableDoseForm *common.CodeableConcept `json:"administrableDoseForm,omitempty"`

	// A general description of the product, when in its final form, suitable for administration
	Description *string `json:"description,omitempty"`

	// A device that is integral to the medicinal product
	Device *common.Reference `json:"device,omitempty"`

	// References a product from which one or more of the constituent parts of that product can be prepared and used
	FormOf []common.Reference `json:"formOf,omitempty"`

	// An identifier for the administrable product
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The ingredients of this administrable medicinal product
	Ingredient []common.CodeableConcept `json:"ingredient,omitempty"`

	// Indicates the specific manufactured items that are part of the 'formOf' product that are used in the preparation
	ProducedFrom []common.Reference `json:"producedFrom,omitempty"`

	// Characteristics e.g. a product's onset of action
	Property []AdministrableProductDefinitionProperty `json:"property,omitempty"`

	// The path by which the product is taken into or makes contact with the body
	RouteOfAdministration []AdministrableProductDefinitionRouteOfAdministration `json:"routeOfAdministration"`

	// draft | active | retired | unknown
	Status RequestStatus `json:"status"`

	// The presentation type in which this item is given to a patient
	UnitOfPresentation *common.CodeableConcept `json:"unitOfPresentation,omitempty"`
}
