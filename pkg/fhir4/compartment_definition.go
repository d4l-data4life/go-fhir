package fhir4

import (
	"encoding/json"

	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// CompartmentDefinition represents a FHIR R4 CompartmentDefinition resource
type CompartmentDefinition struct {
	DomainResource

	ResourceType string                          `json:"resourceType"`
	URL          string                          `json:"url"`
	Version      *string                         `json:"version,omitempty"`
	Name         string                          `json:"name"`
	Status       string                          `json:"status"`
	Experimental *bool                           `json:"experimental,omitempty"`
	Date         *string                         `json:"date,omitempty"`
	Publisher    *string                         `json:"publisher,omitempty"`
	Contact      []common.ContactDetail          `json:"contact,omitempty"`
	Description  *string                         `json:"description,omitempty"`
	UseContext   []common.UsageContext           `json:"useContext,omitempty"`
	Purpose      *string                         `json:"purpose,omitempty"`
	Code         string                          `json:"code"`
	Search       bool                            `json:"search"`
	Resource     []CompartmentDefinitionResource `json:"resource,omitempty"`
}

type CompartmentDefinitionResource struct {
	common.BackboneElement
	Code          string   `json:"code"`
	Param         []string `json:"param,omitempty"`
	Documentation *string  `json:"documentation,omitempty"`
}

func (c *CompartmentDefinition) MarshalJSON() ([]byte, error) {
	type Alias CompartmentDefinition
	return json.Marshal(&struct {
		*Alias
		ResourceType string `json:"resourceType"`
	}{
		Alias:        (*Alias)(c),
		ResourceType: "CompartmentDefinition",
	})
}
