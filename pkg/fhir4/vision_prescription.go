package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// VisionPrescription represents an authorization for the provision of glasses and/or contact lenses to a patient
type VisionPrescription struct {
	DomainResource

	ResourceType string `json:"resourceType"` // Always "VisionPrescription"

	Created           string                                `json:"created"`
	DateWritten       string                                `json:"dateWritten"`
	Encounter         *common.Reference                     `json:"encounter,omitempty"`
	Identifier        []common.Identifier                   `json:"identifier,omitempty"`
	LensSpecification []VisionPrescriptionLensSpecification `json:"lensSpecification"`
	Patient           common.Reference                      `json:"patient"`
	Prescriber        common.Reference                      `json:"prescriber"`
	Status            string                                `json:"status"` // "active" | "cancelled" | "draft" | "entered-in-error"
}

type VisionPrescriptionLensSpecification struct {
	common.BackboneElement

	Add       *float64                                   `json:"add,omitempty"`
	Axis      *int                                       `json:"axis,omitempty"`
	BackCurve *float64                                   `json:"backCurve,omitempty"`
	Brand     *string                                    `json:"brand,omitempty"`
	Color     *string                                    `json:"color,omitempty"`
	Cylinder  *float64                                   `json:"cylinder,omitempty"`
	Diameter  *float64                                   `json:"diameter,omitempty"`
	Duration  *common.Quantity                           `json:"duration,omitempty"`
	Eye       string                                     `json:"eye"` // "right" | "left"
	Note      []common.Annotation                        `json:"note,omitempty"`
	Power     *float64                                   `json:"power,omitempty"`
	Prism     []VisionPrescriptionLensSpecificationPrism `json:"prism,omitempty"`
	Product   common.CodeableConcept                     `json:"product"`
	Sphere    *float64                                   `json:"sphere,omitempty"`
}

type VisionPrescriptionLensSpecificationPrism struct {
	common.BackboneElement

	Amount float64 `json:"amount"`
	Base   string  `json:"base"` // "up" | "down" | "in" | "out"
}
