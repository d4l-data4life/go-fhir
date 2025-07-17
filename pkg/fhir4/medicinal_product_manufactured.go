package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// MedicinalProductManufactured represents the manufactured item as contained in the packaged medicinal product
type MedicinalProductManufactured struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "MedicinalProductManufactured"

	// Ingredient
	Ingredient []common.Reference `json:"ingredient,omitempty"`

	// Dose form as manufactured and before any transformation into the pharmaceutical product
	ManufacturedDoseForm common.CodeableConcept `json:"manufacturedDoseForm"`

	// Manufacturer of the item (Note that this should be named "manufacturer" but it currently causes technical issues)
	Manufacturer []common.Reference `json:"manufacturer,omitempty"`

	// Other codeable characteristics
	OtherCharacteristics []common.CodeableConcept `json:"otherCharacteristics,omitempty"`

	// Dimensions, color etc
	PhysicalCharacteristics *MedicinalProductManufacturedPhysicalCharacteristics `json:"physicalCharacteristics,omitempty"`

	// The quantity or "count number" of the manufactured item
	Quantity common.Quantity `json:"quantity"`

	// The "real world" units in which the quantity of the manufactured item is described
	UnitOfPresentation *common.CodeableConcept `json:"unitOfPresentation,omitempty"`
}

// MedicinalProductManufacturedPhysicalCharacteristics represents dimensions, color etc
type MedicinalProductManufacturedPhysicalCharacteristics struct {
	common.BackboneElement

	// Where applicable, the color can be specified
	Color []string `json:"color,omitempty"`

	// Where applicable, the depth can be specified using a numerical value and its unit of measurement
	Depth *common.Quantity `json:"depth,omitempty"`

	// Where applicable, the external diameter can be specified using a numerical value and its unit of measurement
	ExternalDiameter *common.Quantity `json:"externalDiameter,omitempty"`

	// Where applicable, the height can be specified using a numerical value and its unit of measurement
	Height *common.Quantity `json:"height,omitempty"`

	// Where applicable, the image can be provided
	Image []common.Attachment `json:"image,omitempty"`

	// Where applicable, the imprint can be specified as text
	Imprint []string `json:"imprint,omitempty"`

	// Where applicable, the nominal volume can be specified using a numerical value and its unit of measurement
	NominalVolume *common.Quantity `json:"nominalVolume,omitempty"`

	// Where applicable, the scoring can be specified
	Scoring *common.CodeableConcept `json:"scoring,omitempty"`

	// Where applicable, the shape can be specified
	Shape *string `json:"shape,omitempty"`

	// Where applicable, the weight can be specified using a numerical value and its unit of measurement
	Weight *common.Quantity `json:"weight,omitempty"`

	// Where applicable, the width can be specified using a numerical value and its unit of measurement
	Width *common.Quantity `json:"width,omitempty"`
}
