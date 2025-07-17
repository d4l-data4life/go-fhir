package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// SpecimenDefinition represents a kind of specimen with associated set of requirements
type SpecimenDefinition struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "SpecimenDefinition"

	// The action to be performed for collecting the specimen
	Collection []common.CodeableConcept `json:"collection,omitempty"`

	// A business identifier associated with the kind of specimen
	Identifier *common.Identifier `json:"identifier,omitempty"`

	// Preparation of the patient for specimen collection
	PatientPreparation []common.CodeableConcept `json:"patientPreparation,omitempty"`

	// Time aspect of specimen collection (duration or offset)
	TimeAspect *string `json:"timeAspect,omitempty"`

	// The kind of material to be collected
	TypeCollected *common.CodeableConcept `json:"typeCollected,omitempty"`

	// Specimen conditioned in a container as expected by the testing laboratory
	TypeTested []SpecimenDefinitionTypeTested `json:"typeTested,omitempty"`
}

// SpecimenDefinitionTypeTested represents specimen conditioned in a container as expected by the testing laboratory
type SpecimenDefinitionTypeTested struct {
	common.BackboneElement

	// The specimen's container
	Container *SpecimenDefinitionTypeTestedContainer `json:"container,omitempty"`

	// Set of instructions for preservation/transport of the specimen at a defined temperature interval
	Handling []SpecimenDefinitionTypeTestedHandling `json:"handling,omitempty"`

	// Primary of secondary specimen
	IsDerived *bool `json:"isDerived,omitempty"`

	// The preferred patient preparation for this type of specimen
	PatientPreparation []common.CodeableConcept `json:"patientPreparation,omitempty"`

	// The time aspect of specimen collection (duration or offset)
	TimeAspect *string `json:"timeAspect,omitempty"`

	// The kind of specimen conditioned for testing expected by lab
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// SpecimenDefinitionTypeTestedContainer represents the specimen's container
type SpecimenDefinitionTypeTestedContainer struct {
	common.BackboneElement

	// Substance introduced in the kind of container to preserve, maintain or enhance the specimen
	Additive []SpecimenDefinitionTypeTestedContainerAdditive `json:"additive,omitempty"`

	// Color of container cap
	Cap *common.CodeableConcept `json:"cap,omitempty"`

	// The capacity (volume or other measure) of this kind of container
	Capacity *common.Quantity `json:"capacity,omitempty"`

	// The textual description of the kind of container
	Description *string `json:"description,omitempty"`

	// The type of material of the container
	Material *common.CodeableConcept `json:"material,omitempty"`

	// The minimum volume to be conditioned in the container
	MinimumVolumeQuantity *common.Quantity `json:"minimumVolumeQuantity,omitempty"`

	// The minimum volume to be conditioned in the container
	MinimumVolumeString *string `json:"minimumVolumeString,omitempty"`

	// Special processing that should be applied to the container for this kind of specimen
	Preparation *string `json:"preparation,omitempty"`

	// The type of container used to contain this kind of specimen
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// SpecimenDefinitionTypeTestedContainerAdditive represents substance introduced in the kind of container to preserve, maintain or enhance the specimen
type SpecimenDefinitionTypeTestedContainerAdditive struct {
	common.BackboneElement

	// Substance introduced in the kind of container to preserve, maintain or enhance the specimen
	AdditiveCodeableConcept *common.CodeableConcept `json:"additiveCodeableConcept,omitempty"`

	// Substance introduced in the kind of container to preserve, maintain or enhance the specimen
	AdditiveReference *common.Reference `json:"additiveReference,omitempty"`
}

// SpecimenDefinitionTypeTestedHandling represents set of instructions for preservation/transport of the specimen at a defined temperature interval
type SpecimenDefinitionTypeTestedHandling struct {
	common.BackboneElement

	// Additional textual instructions for the preservation or transport of the specimen
	Instruction *string `json:"instruction,omitempty"`

	// The maximum time interval of preservation of the specimen with these conditions
	MaxDuration *common.Duration `json:"maxDuration,omitempty"`

	// It qualifies the interval of temperature, which characterizes an occurrence of handling
	TemperatureQualifier *common.CodeableConcept `json:"temperatureQualifier,omitempty"`

	// The temperature interval for this set of handling instructions
	TemperatureRange *common.Range `json:"temperatureRange,omitempty"`
}
