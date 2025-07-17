package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// DeviceMetricCalibration represents the calibrations that have been performed or that are required to be performed
type DeviceMetricCalibration struct {
	common.BackboneElement

	// Describes the state of the calibration
	State *DeviceMetricCalibrationState `json:"state,omitempty"`

	// Describes the time last calibration has been performed
	Time *string `json:"time,omitempty"`

	// Describes the type of the calibration method
	Type *DeviceMetricCalibrationType `json:"type,omitempty"`
}

// DeviceMetricCalibrationState represents the state of the calibration
type DeviceMetricCalibrationState string

const (
	DeviceMetricCalibrationStateNotCalibrated       DeviceMetricCalibrationState = "not-calibrated"
	DeviceMetricCalibrationStateCalibrationRequired DeviceMetricCalibrationState = "calibration-required"
	DeviceMetricCalibrationStateCalibrated          DeviceMetricCalibrationState = "calibrated"
	DeviceMetricCalibrationStateUnspecified         DeviceMetricCalibrationState = "unspecified"
)

// DeviceMetricCalibrationType represents the type of the calibration method
type DeviceMetricCalibrationType string

const (
	DeviceMetricCalibrationTypeUnspecified DeviceMetricCalibrationType = "unspecified"
	DeviceMetricCalibrationTypeOffset      DeviceMetricCalibrationType = "offset"
	DeviceMetricCalibrationTypeGain        DeviceMetricCalibrationType = "gain"
	DeviceMetricCalibrationTypeTwoPoint    DeviceMetricCalibrationType = "two-point"
)

// DeviceMetric represents a measurement, calculation or setting capability of a device
type DeviceMetric struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "DeviceMetric"

	// Describes the calibrations that have been performed or that are required to be performed
	Calibration []DeviceMetricCalibration `json:"calibration,omitempty"`

	// Indicates the category of the observation generation process
	Category DeviceMetricCategory `json:"category"`

	// The preferred color associated with the metric (e.g., display color)
	Color *string `json:"color,omitempty"`

	// Note if this is not the top level device - also known as the source device
	Device common.Reference `json:"device"`

	// For identifiers assigned to a device by the device or gateway software
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The frequency at which the metric is taken or recorded
	MeasurementFrequency *common.Quantity `json:"measurementFrequency,omitempty"`

	// Indicates current operational state of the device
	OperationalStatus *DeviceMetricOperationalStatus `json:"operationalStatus,omitempty"`

	// DeviceMetric.type can be referred to either IEEE 11073-10101 or LOINC
	Type common.CodeableConcept `json:"type"`

	// DeviceMetric.unit can refer to either UCUM or preferable a RTMMS coding system
	Unit *common.CodeableConcept `json:"unit,omitempty"`
}

// DeviceMetricCategory represents the category of the observation generation process
type DeviceMetricCategory string

const (
	DeviceMetricCategoryMeasurement DeviceMetricCategory = "measurement"
	DeviceMetricCategorySetting     DeviceMetricCategory = "setting"
	DeviceMetricCategoryCalculation DeviceMetricCategory = "calculation"
	DeviceMetricCategoryUnspecified DeviceMetricCategory = "unspecified"
)

// DeviceMetricOperationalStatus represents the current operational state of the device
type DeviceMetricOperationalStatus string

const (
	DeviceMetricOperationalStatusOn             DeviceMetricOperationalStatus = "on"
	DeviceMetricOperationalStatusOff            DeviceMetricOperationalStatus = "off"
	DeviceMetricOperationalStatusStandby        DeviceMetricOperationalStatus = "standby"
	DeviceMetricOperationalStatusEnteredInError DeviceMetricOperationalStatus = "entered-in-error"
)
