package fhir4

import (
	"encoding/json"

	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// ConceptMap represents a FHIR R4 ConceptMap resource
type ConceptMap struct {
	DomainResource

	ResourceType    string                   `json:"resourceType"`
	URL             *string                  `json:"url,omitempty"`
	Identifier      *common.Identifier       `json:"identifier,omitempty"`
	Version         *string                  `json:"version,omitempty"`
	Name            *string                  `json:"name,omitempty"`
	Title           *string                  `json:"title,omitempty"`
	Status          string                   `json:"status"`
	Experimental    *bool                    `json:"experimental,omitempty"`
	Date            *string                  `json:"date,omitempty"`
	Publisher       *string                  `json:"publisher,omitempty"`
	Contact         []common.ContactDetail   `json:"contact,omitempty"`
	Description     *string                  `json:"description,omitempty"`
	UseContext      []common.UsageContext    `json:"useContext,omitempty"`
	Jurisdiction    []common.CodeableConcept `json:"jurisdiction,omitempty"`
	Purpose         *string                  `json:"purpose,omitempty"`
	Copyright       *string                  `json:"copyright,omitempty"`
	SourceUri       *string                  `json:"sourceUri,omitempty"`
	SourceCanonical *string                  `json:"sourceCanonical,omitempty"`
	TargetUri       *string                  `json:"targetUri,omitempty"`
	TargetCanonical *string                  `json:"targetCanonical,omitempty"`
	Group           []ConceptMapGroup        `json:"group,omitempty"`
}

type ConceptMapGroup struct {
	common.BackboneElement
	Source        *string                  `json:"source,omitempty"`
	SourceVersion *string                  `json:"sourceVersion,omitempty"`
	Target        *string                  `json:"target,omitempty"`
	TargetVersion *string                  `json:"targetVersion,omitempty"`
	Element       []ConceptMapGroupElement `json:"element,omitempty"`
	Unmapped      *ConceptMapGroupUnmapped `json:"unmapped,omitempty"`
}

type ConceptMapGroupElement struct {
	common.BackboneElement
	Code    *string                        `json:"code,omitempty"`
	Display *string                        `json:"display,omitempty"`
	Target  []ConceptMapGroupElementTarget `json:"target,omitempty"`
}

type ConceptMapGroupElementTarget struct {
	common.BackboneElement
	Code        *string                                 `json:"code,omitempty"`
	Display     *string                                 `json:"display,omitempty"`
	Equivalence string                                  `json:"equivalence"`
	Comment     *string                                 `json:"comment,omitempty"`
	DependsOn   []ConceptMapGroupElementTargetDependsOn `json:"dependsOn,omitempty"`
	Product     []ConceptMapGroupElementTargetDependsOn `json:"product,omitempty"`
}

type ConceptMapGroupElementTargetDependsOn struct {
	common.BackboneElement
	Property string  `json:"property"`
	System   *string `json:"system,omitempty"`
	Value    string  `json:"value"`
	Display  *string `json:"display,omitempty"`
}

type ConceptMapGroupUnmapped struct {
	common.BackboneElement
	Mode    string  `json:"mode"`
	Code    *string `json:"code,omitempty"`
	Display *string `json:"display,omitempty"`
	URL     *string `json:"url,omitempty"`
}

func (c *ConceptMap) MarshalJSON() ([]byte, error) {
	type Alias ConceptMap
	return json.Marshal(&struct {
		*Alias
		ResourceType string `json:"resourceType"`
	}{
		Alias:        (*Alias)(c),
		ResourceType: "ConceptMap",
	})
}
