package fhir4

import (
	"encoding/json"

	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// CodeSystem represents a FHIR R4 CodeSystem resource
type CodeSystem struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"`

	URL              *string                  `json:"url,omitempty"`
	Identifier       []common.Identifier      `json:"identifier,omitempty"`
	Version          *string                  `json:"version,omitempty"`
	Name             *string                  `json:"name,omitempty"`
	Title            *string                  `json:"title,omitempty"`
	Status           string                   `json:"status"`
	Experimental     *bool                    `json:"experimental,omitempty"`
	Date             *string                  `json:"date,omitempty"`
	Publisher        *string                  `json:"publisher,omitempty"`
	Contact          []common.ContactDetail   `json:"contact,omitempty"`
	Description      *string                  `json:"description,omitempty"`
	UseContext       []common.UsageContext    `json:"useContext,omitempty"`
	Jurisdiction     []common.CodeableConcept `json:"jurisdiction,omitempty"`
	Purpose          *string                  `json:"purpose,omitempty"`
	Copyright        *string                  `json:"copyright,omitempty"`
	CaseSensitive    *bool                    `json:"caseSensitive,omitempty"`
	ValueSet         *string                  `json:"valueSet,omitempty"`
	HierarchyMeaning *string                  `json:"hierarchyMeaning,omitempty"`
	Compositional    *bool                    `json:"compositional,omitempty"`
	VersionNeeded    *bool                    `json:"versionNeeded,omitempty"`
	Content          string                   `json:"content"`
	Supplements      *string                  `json:"supplements,omitempty"`
	Count            *int                     `json:"count,omitempty"`
	Filter           []CodeSystemFilter       `json:"filter,omitempty"`
	Property         []CodeSystemProperty     `json:"property,omitempty"`
	Concept          []CodeSystemConcept      `json:"concept,omitempty"`
}

// CodeSystemFilter represents filters
type CodeSystemFilter struct {
	common.BackboneElement

	Code        string   `json:"code"`
	Description *string  `json:"description,omitempty"`
	Operator    []string `json:"operator"`
	Value       string   `json:"value"`
}

// CodeSystemProperty represents properties
type CodeSystemProperty struct {
	common.BackboneElement

	Code        string  `json:"code"`
	URI         *string `json:"uri,omitempty"`
	Description *string `json:"description,omitempty"`
	Type        string  `json:"type"`
}

// CodeSystemConcept represents concepts
type CodeSystemConcept struct {
	common.BackboneElement

	Code        string                         `json:"code"`
	Display     *string                        `json:"display,omitempty"`
	Definition  *string                        `json:"definition,omitempty"`
	Designation []CodeSystemConceptDesignation `json:"designation,omitempty"`
	Property    []CodeSystemConceptProperty    `json:"property,omitempty"`
	Concept     []CodeSystemConcept            `json:"concept,omitempty"`
}

// CodeSystemConceptDesignation represents concept designations
type CodeSystemConceptDesignation struct {
	common.BackboneElement

	Language *string        `json:"language,omitempty"`
	Use      *common.Coding `json:"use,omitempty"`
	Value    string         `json:"value"`
}

// CodeSystemConceptProperty represents concept properties
type CodeSystemConceptProperty struct {
	common.BackboneElement

	Code          string         `json:"code"`
	ValueCode     *string        `json:"valueCode,omitempty"`
	ValueCoding   *common.Coding `json:"valueCoding,omitempty"`
	ValueString   *string        `json:"valueString,omitempty"`
	ValueInteger  *int           `json:"valueInteger,omitempty"`
	ValueBoolean  *bool          `json:"valueBoolean,omitempty"`
	ValueDateTime *string        `json:"valueDateTime,omitempty"`
	ValueDecimal  *float64       `json:"valueDecimal,omitempty"`
}

// MarshalJSON marshals the CodeSystem to JSON
func (c *CodeSystem) MarshalJSON() ([]byte, error) {
	type Alias CodeSystem
	return json.Marshal(&struct {
		*Alias
		ResourceType string `json:"resourceType"`
	}{
		Alias:        (*Alias)(c),
		ResourceType: "CodeSystem",
	})
}
