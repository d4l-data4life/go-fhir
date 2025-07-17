package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// VisionPrescriptionLensSpecificationPrism represents allows for adjustment on two axis
type VisionPrescriptionLensSpecificationPrism struct {
	common.BackboneElement

	// Amount of prism to compensate for eye alignment in fractional units
	Amount float64 `json:"amount"`

	// The relative base, or reference lens edge, for the prism
	Base VisionPrescriptionLensSpecificationPrismBase `json:"base"`
}

// VisionPrescriptionLensSpecificationPrismBase represents the relative base for the prism
type VisionPrescriptionLensSpecificationPrismBase string

const (
	VisionPrescriptionLensSpecificationPrismBaseUp   VisionPrescriptionLensSpecificationPrismBase = "up"
	VisionPrescriptionLensSpecificationPrismBaseDown VisionPrescriptionLensSpecificationPrismBase = "down"
	VisionPrescriptionLensSpecificationPrismBaseIn   VisionPrescriptionLensSpecificationPrismBase = "in"
	VisionPrescriptionLensSpecificationPrismBaseOut  VisionPrescriptionLensSpecificationPrismBase = "out"
)

// VisionPrescriptionLensSpecification represents contain the details of the individual lens specifications
type VisionPrescriptionLensSpecification struct {
	common.BackboneElement

	// Power adjustment for multifocal lenses measured in dioptres (0.25 units)
	Add *float64 `json:"add,omitempty"`

	// The limits are +180 and -180 degrees
	Axis *int `json:"axis,omitempty"`

	// Back curvature measured in millimetres
	BackCurve *float64 `json:"backCurve,omitempty"`

	// Brand recommendations or restrictions
	Brand *string `json:"brand,omitempty"`

	// Special color or pattern
	Color *string `json:"color,omitempty"`

	// Power adjustment for astigmatism measured in dioptres (0.25 units)
	Cylinder *float64 `json:"cylinder,omitempty"`

	// Contact lens diameter measured in millimetres
	Diameter *float64 `json:"diameter,omitempty"`

	// The recommended maximum wear period for the lens
	Duration *common.Quantity `json:"duration,omitempty"`

	// May also appear on the paper claim form or in the Medical Records as as OD (oculus dexter) for the right eye and OS (oculus sinister) for the left eye
	Eye VisionPrescriptionLensSpecificationEye `json:"eye"`

	// Notes for special requirements such as coatings and lens materials
	Note []common.Annotation `json:"note,omitempty"`

	// Contact lens power measured in dioptres (0.25 units)
	Power *float64 `json:"power,omitempty"`

	// Allows for adjustment on two axis
	Prism []VisionPrescriptionLensSpecificationPrism `json:"prism,omitempty"`

	// Identifies the type of vision correction product which is required for the patient
	Product common.CodeableConcept `json:"product"`

	// The value is negative for near-sighted and positive for far sighted
	Sphere *float64 `json:"sphere,omitempty"`
}

// VisionPrescriptionLensSpecificationEye represents the eye for the lens specification
type VisionPrescriptionLensSpecificationEye string

const (
	VisionPrescriptionLensSpecificationEyeRight VisionPrescriptionLensSpecificationEye = "right"
	VisionPrescriptionLensSpecificationEyeLeft  VisionPrescriptionLensSpecificationEye = "left"
)

// VisionPrescriptionStatus represents the status of a vision prescription
type VisionPrescriptionStatus string

const (
	VisionPrescriptionStatusActive         VisionPrescriptionStatus = "active"
	VisionPrescriptionStatusCancelled      VisionPrescriptionStatus = "cancelled"
	VisionPrescriptionStatusDraft          VisionPrescriptionStatus = "draft"
	VisionPrescriptionStatusEnteredInError VisionPrescriptionStatus = "entered-in-error"
)

// VisionPrescription represents an authorization for the provision of glasses and/or contact lenses to a patient
type VisionPrescription struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "VisionPrescription"

	// The date this resource was created
	Created string `json:"created"`

	// Jurisdictions determine the valid lifetime of a prescription
	DateWritten string `json:"dateWritten"`

	// A reference to a resource that identifies the particular occurrence of contact between patient and health care provider during which the prescription was issued
	Encounter *common.Reference `json:"encounter,omitempty"`

	// A unique identifier assigned to this vision prescription
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Contain the details of the individual lens specifications and serves as the authorization for the fullfillment by certified professionals
	LensSpecification []VisionPrescriptionLensSpecification `json:"lensSpecification"`

	// A resource reference to the person to whom the vision prescription applies
	Patient common.Reference `json:"patient"`

	// The healthcare professional responsible for authorizing the prescription
	Prescriber common.Reference `json:"prescriber"`

	// This element is labeled as a modifier because the status contains codes that mark the resource as not currently valid
	Status VisionPrescriptionStatus `json:"status"`
}
