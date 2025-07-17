package fhir4

import (
	"encoding/json"

	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// CapabilityStatement represents a FHIR R4 CapabilityStatement resource
type CapabilityStatement struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"`

	URL                 *string                            `json:"url,omitempty"`
	Version             *string                            `json:"version,omitempty"`
	Name                *string                            `json:"name,omitempty"`
	Title               *string                            `json:"title,omitempty"`
	Status              string                             `json:"status"`
	Experimental        *bool                              `json:"experimental,omitempty"`
	Date                string                             `json:"date"`
	Publisher           *string                            `json:"publisher,omitempty"`
	Contact             []common.ContactDetail             `json:"contact,omitempty"`
	Description         *string                            `json:"description,omitempty"`
	UseContext          []common.UsageContext              `json:"useContext,omitempty"`
	Jurisdiction        []common.CodeableConcept           `json:"jurisdiction,omitempty"`
	Purpose             *string                            `json:"purpose,omitempty"`
	Copyright           *string                            `json:"copyright,omitempty"`
	Kind                string                             `json:"kind"`
	Instantiates        []string                           `json:"instantiates,omitempty"`
	Imports             []string                           `json:"imports,omitempty"`
	Software            *CapabilityStatementSoftware       `json:"software,omitempty"`
	Implementation      *CapabilityStatementImplementation `json:"implementation,omitempty"`
	FHIRVersion         string                             `json:"fhirVersion"`
	Format              []string                           `json:"format"`
	PatchFormat         []string                           `json:"patchFormat,omitempty"`
	ImplementationGuide []string                           `json:"implementationGuide,omitempty"`
	Rest                []CapabilityStatementRest          `json:"rest,omitempty"`
	Messaging           []CapabilityStatementMessaging     `json:"messaging,omitempty"`
	Document            []CapabilityStatementDocument      `json:"document,omitempty"`
}

// CapabilityStatementSoftware represents software information
type CapabilityStatementSoftware struct {
	common.BackboneElement

	Name        string  `json:"name"`
	Version     *string `json:"version,omitempty"`
	ReleaseDate *string `json:"releaseDate,omitempty"`
}

// CapabilityStatementImplementation represents implementation details
type CapabilityStatementImplementation struct {
	common.BackboneElement

	Description string            `json:"description"`
	URL         *string           `json:"url,omitempty"`
	Custodian   *common.Reference `json:"custodian,omitempty"`
}

// CapabilityStatementRest represents REST capabilities
type CapabilityStatementRest struct {
	common.BackboneElement

	Mode          string                                       `json:"mode"`
	Documentation *string                                      `json:"documentation,omitempty"`
	Security      *CapabilityStatementRestSecurity             `json:"security,omitempty"`
	Resource      []CapabilityStatementRestResource            `json:"resource,omitempty"`
	Interaction   []CapabilityStatementRestInteraction         `json:"interaction,omitempty"`
	SearchParam   []CapabilityStatementRestResourceSearchParam `json:"searchParam,omitempty"`
	Operation     []CapabilityStatementRestResourceOperation   `json:"operation,omitempty"`
	Compartment   []string                                     `json:"compartment,omitempty"`
}

// CapabilityStatementRestSecurity represents security information
type CapabilityStatementRestSecurity struct {
	common.BackboneElement

	Cors        *bool                    `json:"cors,omitempty"`
	Service     []common.CodeableConcept `json:"service,omitempty"`
	Description *string                  `json:"description,omitempty"`
}

// CapabilityStatementRestResource represents resource capabilities
type CapabilityStatementRestResource struct {
	common.BackboneElement

	Type              string                                       `json:"type"`
	Profile           *string                                      `json:"profile,omitempty"`
	SupportedProfile  []string                                     `json:"supportedProfile,omitempty"`
	Documentation     *string                                      `json:"documentation,omitempty"`
	Interaction       []CapabilityStatementRestResourceInteraction `json:"interaction,omitempty"`
	Versioning        *string                                      `json:"versioning,omitempty"`
	ReadHistory       *bool                                        `json:"readHistory,omitempty"`
	UpdateCreate      *bool                                        `json:"updateCreate,omitempty"`
	ConditionalCreate *bool                                        `json:"conditionalCreate,omitempty"`
	ConditionalRead   *string                                      `json:"conditionalRead,omitempty"`
	ConditionalUpdate *bool                                        `json:"conditionalUpdate,omitempty"`
	ConditionalDelete *string                                      `json:"conditionalDelete,omitempty"`
	ReferencePolicy   []string                                     `json:"referencePolicy,omitempty"`
	SearchInclude     []string                                     `json:"searchInclude,omitempty"`
	SearchRevInclude  []string                                     `json:"searchRevInclude,omitempty"`
	SearchParam       []CapabilityStatementRestResourceSearchParam `json:"searchParam,omitempty"`
	Operation         []CapabilityStatementRestResourceOperation   `json:"operation,omitempty"`
}

// CapabilityStatementRestResourceInteraction represents resource interactions
type CapabilityStatementRestResourceInteraction struct {
	common.BackboneElement

	Code          string  `json:"code"`
	Documentation *string `json:"documentation,omitempty"`
}

// CapabilityStatementRestResourceSearchParam represents search parameters
type CapabilityStatementRestResourceSearchParam struct {
	common.BackboneElement

	Name          string  `json:"name"`
	Definition    *string `json:"definition,omitempty"`
	Type          string  `json:"type"`
	Documentation *string `json:"documentation,omitempty"`
}

// CapabilityStatementRestResourceOperation represents operations
type CapabilityStatementRestResourceOperation struct {
	common.BackboneElement

	Name          string  `json:"name"`
	Definition    string  `json:"definition"`
	Documentation *string `json:"documentation,omitempty"`
}

// CapabilityStatementRestInteraction represents REST interactions
type CapabilityStatementRestInteraction struct {
	common.BackboneElement

	Code          string  `json:"code"`
	Documentation *string `json:"documentation,omitempty"`
}

// CapabilityStatementMessaging represents messaging capabilities
type CapabilityStatementMessaging struct {
	common.BackboneElement

	Endpoint         []CapabilityStatementMessagingEndpoint         `json:"endpoint,omitempty"`
	ReliableCache    *int                                           `json:"reliableCache,omitempty"`
	Documentation    *string                                        `json:"documentation,omitempty"`
	SupportedMessage []CapabilityStatementMessagingSupportedMessage `json:"supportedMessage,omitempty"`
}

// CapabilityStatementMessagingEndpoint represents messaging endpoints
type CapabilityStatementMessagingEndpoint struct {
	common.BackboneElement

	Protocol common.Coding `json:"protocol"`
	Address  string        `json:"address"`
}

// CapabilityStatementMessagingSupportedMessage represents supported messages
type CapabilityStatementMessagingSupportedMessage struct {
	common.BackboneElement

	Mode       string `json:"mode"`
	Definition string `json:"definition"`
}

// CapabilityStatementDocument represents document capabilities
type CapabilityStatementDocument struct {
	common.BackboneElement

	Mode          string  `json:"mode"`
	Documentation *string `json:"documentation,omitempty"`
	Profile       string  `json:"profile"`
}

// MarshalJSON marshals the CapabilityStatement to JSON
func (c *CapabilityStatement) MarshalJSON() ([]byte, error) {
	type Alias CapabilityStatement
	return json.Marshal(&struct {
		*Alias
		ResourceType string `json:"resourceType"`
	}{
		Alias:        (*Alias)(c),
		ResourceType: "CapabilityStatement",
	})
}
