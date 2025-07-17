// Package fhir5 contains FHIR R5 resource definitions
package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// SpecimenDefinitionTypeTestedContainerAdditive represents an additive in a specimen container
type SpecimenDefinitionTypeTestedContainerAdditive struct {
	common.BackboneElement
	AdditiveCodeableConcept *common.CodeableConcept `json:"additiveCodeableConcept,omitempty"`
	AdditiveReference       *common.Reference       `json:"additiveReference,omitempty"`
}

// SpecimenDefinitionTypeTestedContainer represents the container for a specimen
type SpecimenDefinitionTypeTestedContainer struct {
	common.BackboneElement
	Additive              []SpecimenDefinitionTypeTestedContainerAdditive `json:"additive,omitempty"`
	Cap                   *common.CodeableConcept                         `json:"cap,omitempty"`
	Capacity              *common.Quantity                                `json:"capacity,omitempty"`
	Material              *common.CodeableConcept                         `json:"material,omitempty"`
	MinimumVolumeQuantity *common.Quantity                                `json:"minimumVolumeQuantity,omitempty"`
	MinimumVolumeString   *string                                         `json:"minimumVolumeString,omitempty"`
	Preparation           *string                                         `json:"preparation,omitempty"`
	Type                  *common.CodeableConcept                         `json:"type,omitempty"`
}

// SpecimenDefinitionTypeTestedHandling represents handling instructions for a specimen
type SpecimenDefinitionTypeTestedHandling struct {
	common.BackboneElement
	Instruction          *string                 `json:"instruction,omitempty"`
	MaxDuration          *Duration               `json:"maxDuration,omitempty"`
	TemperatureQualifier *common.CodeableConcept `json:"temperatureQualifier,omitempty"`
	TemperatureRange     *Range                  `json:"temperatureRange,omitempty"`
}

// SpecimenDefinitionTypeTested represents a kind of specimen conditioned for testing
type SpecimenDefinitionTypeTested struct {
	common.BackboneElement
	Container          *SpecimenDefinitionTypeTestedContainer `json:"container,omitempty"`
	Diagnostic         *common.Reference                      `json:"diagnostic,omitempty"`
	Handling           []SpecimenDefinitionTypeTestedHandling `json:"handling,omitempty"`
	IsDerived          *bool                                  `json:"isDerived,omitempty"`
	Type               *common.CodeableConcept                `json:"type,omitempty"`
	Preference         string                                 `json:"preference"`
	RejectionCriterion []common.CodeableConcept               `json:"rejectionCriterion,omitempty"`
	Requirement        *string                                `json:"requirement,omitempty"`
	RetentionTime      *Duration                              `json:"retentionTime,omitempty"`
}

// SpecimenDefinition represents a kind of specimen with associated set of requirements
type SpecimenDefinition struct {
	DomainResource
	ResourceType    string                         `json:"resourceType"` // Always "SpecimenDefinition"
	ApprovalDate    *string                        `json:"approvalDate,omitempty"`
	Author          []ContactDetail                `json:"author,omitempty"`
	Contact         []ContactDetail                `json:"contact,omitempty"`
	Copyright       *string                        `json:"copyright,omitempty"`
	Date            *string                        `json:"date,omitempty"`
	Description     *string                        `json:"description,omitempty"`
	Editor          []ContactDetail                `json:"editor,omitempty"`
	EffectivePeriod *common.Period                 `json:"effectivePeriod,omitempty"`
	Endorser        []ContactDetail                `json:"endorser,omitempty"`
	Experimental    *bool                          `json:"experimental,omitempty"`
	Identifier      []common.Identifier            `json:"identifier,omitempty"`
	Jurisdiction    []common.CodeableConcept       `json:"jurisdiction,omitempty"`
	LastReviewDate  *string                        `json:"lastReviewDate,omitempty"`
	Name            *string                        `json:"name,omitempty"`
	Publisher       *string                        `json:"publisher,omitempty"`
	Purpose         *string                        `json:"purpose,omitempty"`
	RelatedArtifact []interface{}                  `json:"relatedArtifact,omitempty"`
	Reviewer        []ContactDetail                `json:"reviewer,omitempty"`
	Status          string                         `json:"status"`
	Subtitle        *string                        `json:"subtitle,omitempty"`
	Title           *string                        `json:"title,omitempty"`
	Topic           []common.CodeableConcept       `json:"topic,omitempty"`
	Type            *common.CodeableConcept        `json:"type,omitempty"`
	URL             *string                        `json:"url,omitempty"`
	Usage           *string                        `json:"usage,omitempty"`
	UseContext      []interface{}                  `json:"useContext,omitempty"`
	Version         *string                        `json:"version,omitempty"`
	TypeTested      []SpecimenDefinitionTypeTested `json:"typeTested,omitempty"`
}
