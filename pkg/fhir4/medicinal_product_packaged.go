package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// MedicinalProductPackaged represents a medicinal product in a container or package
type MedicinalProductPackaged struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "MedicinalProductPackaged"

	// Batch numbering
	BatchIdentifier []MedicinalProductPackagedBatchIdentifier `json:"batchIdentifier,omitempty"`

	// Textual description
	Description *string `json:"description,omitempty"`

	// Unique identifier
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The legal status of supply of the medicinal product as classified by the regulator
	LegalStatusOfSupply *common.CodeableConcept `json:"legalStatusOfSupply,omitempty"`

	// Manufacturer of this Package Item
	Manufacturer []common.Reference `json:"manufacturer,omitempty"`

	// Marketing authorization
	MarketingAuthorization *common.Reference `json:"marketingAuthorization,omitempty"`

	// Marketing information
	MarketingStatus []MedicinalProductPackagedMarketingStatus `json:"marketingStatus,omitempty"`

	// A packaging item, as a contained for medicine, possibly with other packaging items within
	PackageItem []MedicinalProductPackagedPackageItem `json:"packageItem"`

	// The product with this is a pack for
	Subject []common.Reference `json:"subject,omitempty"`
}

// MedicinalProductPackagedBatchIdentifier represents batch numbering
type MedicinalProductPackagedBatchIdentifier struct {
	common.BackboneElement

	// A number appearing on the immediate packaging (and not the outer packaging)
	ImmediatePackaging *common.Identifier `json:"immediatePackaging,omitempty"`

	// A number appearing on the outer packaging of a specific batch
	OuterPackaging common.Identifier `json:"outerPackaging"`
}

// MedicinalProductPackagedPackageItem represents a packaging item, as a contained for medicine, possibly with other packaging items within
type MedicinalProductPackagedPackageItem struct {
	common.BackboneElement

	// A possible alternate material for the packaging
	AlternateMaterial []common.CodeableConcept `json:"alternateMaterial,omitempty"`

	// A device accompanying a medicinal product
	Device []common.Reference `json:"device,omitempty"`

	// Including possibly Data Carrier Identifier
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The manufactured item as contained in the packaged medicinal product
	ManufacturedItem []common.Reference `json:"manufacturedItem,omitempty"`

	// Manufacturer of this Package Item
	Manufacturer []common.Reference `json:"manufacturer,omitempty"`

	// Material type of the package item
	Material []common.CodeableConcept `json:"material,omitempty"`

	// Other codeable characteristics
	OtherCharacteristics []common.CodeableConcept `json:"otherCharacteristics,omitempty"`

	// Allows containers within containers
	PackageItem []MedicinalProductPackagedPackageItem `json:"packageItem,omitempty"`

	// Dimensions, color etc
	PhysicalCharacteristics *MedicinalProductPackagedPhysicalCharacteristics `json:"physicalCharacteristics,omitempty"`

	// The quantity of this package in the medicinal product, at the current level of packaging. The outermost is always 1
	Quantity common.Quantity `json:"quantity"`

	// Shelf Life and storage information
	ShelfLifeStorage []MedicinalProductPackagedShelfLife `json:"shelfLifeStorage,omitempty"`

	// The physical type of the container of the medicine
	Type common.CodeableConcept `json:"type"`
}

// MedicinalProductPackagedPhysicalCharacteristics represents dimensions, color etc
type MedicinalProductPackagedPhysicalCharacteristics struct {
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

// MedicinalProductPackagedShelfLife represents the shelf-life and storage information for a medicinal product item or container
type MedicinalProductPackagedShelfLife struct {
	common.BackboneElement

	// Unique identifier for the packaged Medicinal Product
	Identifier *common.Identifier `json:"identifier,omitempty"`

	// The shelf life time period can be specified using a numerical value for the period of time and its unit of time measurement
	Period common.Quantity `json:"period"`

	// Special precautions for storage, if any, can be specified using an appropriate controlled vocabulary
	SpecialPrecautionsForStorage []common.CodeableConcept `json:"specialPrecautionsForStorage,omitempty"`

	// This describes the shelf life, taking into account various scenarios such as shelf life of the packaged Medicinal Product itself
	Type common.CodeableConcept `json:"type"`
}

// MedicinalProductPackagedMarketingStatus represents the marketing status describes the date when a medicinal product is actually put on the market
type MedicinalProductPackagedMarketingStatus struct {
	common.BackboneElement

	// The country in which the marketing authorisation has been granted
	Country common.CodeableConcept `json:"country"`

	// The date when the Medicinal Product is placed on the market
	DateRange common.Period `json:"dateRange"`

	// Where a Medicines Regulatory Agency has granted a marketing authorisation
	Jurisdiction *common.CodeableConcept `json:"jurisdiction,omitempty"`

	// The date when the Medicinal Product is placed on the market
	RestoreDate *string `json:"restoreDate,omitempty"`

	// This attribute provides information on the status of the marketing of the medicinal product
	Status common.CodeableConcept `json:"status"`
}
