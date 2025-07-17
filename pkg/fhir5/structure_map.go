// Package fhir5 contains FHIR R5 resource definitions
package fhir5

import "github.com/d4l-data4life/go-fhir/pkg/common"

// StructureMapStructure represents a referenced structure in a StructureMap
// See: https://hl7.org/fhir/structuremap.html
type StructureMapStructure struct {
	common.BackboneElement
	Alias         *string `json:"alias,omitempty"`
	Documentation *string `json:"documentation,omitempty"`
	Mode          string  `json:"mode"`
	URL           string  `json:"url"`
}

type StructureMapConst struct {
	common.BackboneElement
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

type StructureMapGroupInput struct {
	common.BackboneElement
	Documentation *string `json:"documentation,omitempty"`
	Mode          string  `json:"mode"`
	Name          string  `json:"name"`
	Type          *string `json:"type,omitempty"`
}

type StructureMapGroupRuleSource struct {
	common.BackboneElement
	Check        *string `json:"check,omitempty"`
	Condition    *string `json:"condition,omitempty"`
	Context      string  `json:"context"`
	DefaultValue *string `json:"defaultValue,omitempty"`
	Element      *string `json:"element,omitempty"`
	ListMode     *string `json:"listMode,omitempty"`
	LogMessage   *string `json:"logMessage,omitempty"`
	Max          *string `json:"max,omitempty"`
	Min          *int    `json:"min,omitempty"`
	Type         *string `json:"type,omitempty"`
	Variable     *string `json:"variable,omitempty"`
}

type StructureMapGroupRuleTargetParameter struct {
	common.BackboneElement
	ValueId       *string  `json:"valueId,omitempty"`
	ValueString   *string  `json:"valueString,omitempty"`
	ValueBoolean  *bool    `json:"valueBoolean,omitempty"`
	ValueInteger  *int     `json:"valueInteger,omitempty"`
	ValueDecimal  *float64 `json:"valueDecimal,omitempty"`
	ValueDate     *string  `json:"valueDate,omitempty"`
	ValueTime     *string  `json:"valueTime,omitempty"`
	ValueDateTime *string  `json:"valueDateTime,omitempty"`
}

type StructureMapGroupRuleTarget struct {
	common.BackboneElement
	Context    *string                                `json:"context,omitempty"`
	Element    *string                                `json:"element,omitempty"`
	ListMode   []string                               `json:"listMode,omitempty"`
	ListRuleId *string                                `json:"listRuleId,omitempty"`
	Parameter  []StructureMapGroupRuleTargetParameter `json:"parameter,omitempty"`
	Transform  *string                                `json:"transform,omitempty"`
	Variable   *string                                `json:"variable,omitempty"`
}

type StructureMapGroupRuleDependent struct {
	common.BackboneElement
	Name      string                                 `json:"name"`
	Parameter []StructureMapGroupRuleTargetParameter `json:"parameter"`
}

type StructureMapGroupRule struct {
	common.BackboneElement
	Dependent     []StructureMapGroupRuleDependent `json:"dependent,omitempty"`
	Documentation *string                          `json:"documentation,omitempty"`
	Name          *string                          `json:"name,omitempty"`
	Rule          []StructureMapGroupRule          `json:"rule,omitempty"`
	Source        []StructureMapGroupRuleSource    `json:"source"`
	Target        []StructureMapGroupRuleTarget    `json:"target,omitempty"`
}

type StructureMapGroup struct {
	common.BackboneElement
	Documentation *string                  `json:"documentation,omitempty"`
	Extends       *string                  `json:"extends,omitempty"`
	Input         []StructureMapGroupInput `json:"input"`
	Name          string                   `json:"name"`
	Rule          []StructureMapGroupRule  `json:"rule,omitempty"`
	TypeMode      *string                  `json:"typeMode,omitempty"`
}

// StructureMap is a FHIR R5 resource for mapping between structures
type StructureMap struct {
	DomainResource
	ResourceType           string                   `json:"resourceType"` // Always "StructureMap"
	Const                  []StructureMapConst      `json:"const,omitempty"`
	Contact                []ContactDetail          `json:"contact,omitempty"`
	Copyright              *string                  `json:"copyright,omitempty"`
	CopyrightLabel         *string                  `json:"copyrightLabel,omitempty"`
	Date                   *string                  `json:"date,omitempty"`
	Description            *string                  `json:"description,omitempty"`
	Experimental           *bool                    `json:"experimental,omitempty"`
	Group                  []StructureMapGroup      `json:"group"`
	Identifier             []common.Identifier      `json:"identifier,omitempty"`
	Import                 []string                 `json:"import,omitempty"`
	Jurisdiction           []common.CodeableConcept `json:"jurisdiction,omitempty"`
	Name                   string                   `json:"name"`
	Publisher              *string                  `json:"publisher,omitempty"`
	Purpose                *string                  `json:"purpose,omitempty"`
	Status                 string                   `json:"status"`
	Structure              []StructureMapStructure  `json:"structure,omitempty"`
	Title                  *string                  `json:"title,omitempty"`
	URL                    string                   `json:"url"`
	UseContext             []UsageContext           `json:"useContext,omitempty"`
	Version                *string                  `json:"version,omitempty"`
	VersionAlgorithmString *string                  `json:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *common.Coding           `json:"versionAlgorithmCoding,omitempty"`
}
