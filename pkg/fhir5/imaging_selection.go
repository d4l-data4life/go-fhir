package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// ImagingSelectionStatus represents the status of an imaging selection
type ImagingSelectionStatus string

const (
	ImagingSelectionStatusAvailable      ImagingSelectionStatus = "available"
	ImagingSelectionStatusEnteredInError ImagingSelectionStatus = "entered-in-error"
	ImagingSelectionStatusUnknown        ImagingSelectionStatus = "unknown"
)

// ImagingSelectionInstanceImageRegion2DType represents the type of 2D image region
type ImagingSelectionInstanceImageRegion2DType string

const (
	ImagingSelectionInstanceImageRegion2DTypePoint        ImagingSelectionInstanceImageRegion2DType = "point"
	ImagingSelectionInstanceImageRegion2DTypePolyline     ImagingSelectionInstanceImageRegion2DType = "polyline"
	ImagingSelectionInstanceImageRegion2DTypeInterpolated ImagingSelectionInstanceImageRegion2DType = "interpolated"
	ImagingSelectionInstanceImageRegion2DTypeCircle       ImagingSelectionInstanceImageRegion2DType = "circle"
	ImagingSelectionInstanceImageRegion2DTypeEllipse      ImagingSelectionInstanceImageRegion2DType = "ellipse"
)

// ImagingSelectionInstanceImageRegion3DType represents the type of 3D image region
type ImagingSelectionInstanceImageRegion3DType string

const (
	ImagingSelectionInstanceImageRegion3DTypePoint      ImagingSelectionInstanceImageRegion3DType = "point"
	ImagingSelectionInstanceImageRegion3DTypeMultipoint ImagingSelectionInstanceImageRegion3DType = "multipoint"
	ImagingSelectionInstanceImageRegion3DTypePolyline   ImagingSelectionInstanceImageRegion3DType = "polyline"
	ImagingSelectionInstanceImageRegion3DTypePolygon    ImagingSelectionInstanceImageRegion3DType = "polygon"
	ImagingSelectionInstanceImageRegion3DTypeEllipse    ImagingSelectionInstanceImageRegion3DType = "ellipse"
	ImagingSelectionInstanceImageRegion3DTypeEllipsoid  ImagingSelectionInstanceImageRegion3DType = "ellipsoid"
)

// ImagingSelectionPerformer represents selector of the instances – human or machine
type ImagingSelectionPerformer struct {
	common.BackboneElement

	// Author – human or machine
	Actor *common.Reference `json:"actor,omitempty"`

	// Distinguishes the type of involvement of the performer
	Function *common.CodeableConcept `json:"function,omitempty"`
}

// ImagingSelectionInstanceImageRegion2D represents each imaging selection instance or frame list 2D image region
type ImagingSelectionInstanceImageRegion2D struct {
	common.BackboneElement

	// For a description of how 2D coordinates are encoded
	Coordinate []float64 `json:"coordinate"`

	// See DICOM PS3.3 C.10.5.1.2
	RegionType ImagingSelectionInstanceImageRegion2DType `json:"regionType"`
}

// ImagingSelectionInstanceImageRegion3D represents each imaging selection 3D image region
type ImagingSelectionInstanceImageRegion3D struct {
	common.BackboneElement

	// For a description of how 3D coordinates are encoded
	Coordinate []float64 `json:"coordinate"`

	// See DICOM PS3.3 C.18.9.1.2
	RegionType ImagingSelectionInstanceImageRegion3DType `json:"regionType"`
}

// ImagingSelectionInstance represents each imaging selection instance or frame list
type ImagingSelectionInstance struct {
	common.BackboneElement

	// Each imaging selection instance or frame list might includes 2D image regions
	ImageRegion2D []ImagingSelectionInstanceImageRegion2D `json:"imageRegion2D,omitempty"`

	// Each imaging selection might includes 3D image regions
	ImageRegion3D []ImagingSelectionInstanceImageRegion3D `json:"imageRegion3D,omitempty"`

	// Note: A multiframe instance has a single instance number
	Number *int `json:"number,omitempty"`

	// See DICOM PS3.3 C.12.1
	SopClass *common.Coding `json:"sopClass,omitempty"`

	// Selected subset of the SOP Instance
	Subset []string `json:"subset,omitempty"`

	// See DICOM PS3.3 C.12.1
	Uid string `json:"uid"`
}

// ImagingSelection represents a selection of DICOM SOP instances and/or frames within a single Study and Series
type ImagingSelection struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "ImagingSelection"

	// A list of the diagnostic requests that resulted in this imaging selection being performed
	BasedOn []common.Reference `json:"basedOn,omitempty"`

	// The anatomic structures examined
	BodySite *CodeableReference `json:"bodySite,omitempty"`

	// Classifies the imaging selection
	Category []common.CodeableConcept `json:"category,omitempty"`

	// All code-value and, if present, component.code-component.value pairs need to be taken into account
	Code common.CodeableConcept `json:"code"`

	// The imaging study from which the imaging selection is made
	DerivedFrom []common.Reference `json:"derivedFrom,omitempty"`

	// Typical endpoint types include DICOM WADO-RS
	Endpoint []common.Reference `json:"endpoint,omitempty"`

	// An imaging selection may reference a DICOM resource
	Focus []common.Reference `json:"focus,omitempty"`

	// See DICOM PS3.3 C.7.4.1
	FrameOfReferenceUid *string `json:"frameOfReferenceUid,omitempty"`

	// This is a business identifier, not a resource identifier
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// Each imaging selection includes one or more selected DICOM SOP instances
	Instance []ImagingSelectionInstance `json:"instance,omitempty"`

	// The date and time this imaging selection was created
	Issued *string `json:"issued,omitempty"`

	// Selector of the instances – human or machine
	Performer []ImagingSelectionPerformer `json:"performer,omitempty"`

	// See DICOM PS3.3 C.7.3
	SeriesNumber *int `json:"seriesNumber,omitempty"`

	// See DICOM PS3.3 C.7.3
	SeriesUid *string `json:"seriesUid,omitempty"`

	// Unknown does not represent "other" - one of the defined statuses must apply
	Status ImagingSelectionStatus `json:"status"`

	// See DICOM PS3.3 C.7.2
	StudyUid *string `json:"studyUid,omitempty"`

	// The patient, or group of patients, location, device, organization, procedure or practitioner
	Subject *common.Reference `json:"subject,omitempty"`
}
