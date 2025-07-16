// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// ContractContentDefinition represents precursor content for a contract
type ContractContentDefinition struct {
	common.BackboneElement

	Copyright         *string                                    `json:"copyright,omitempty"`
	PublicationDate   *string                                    `json:"publicationDate,omitempty"`
	PublicationStatus ContractContentDefinitionPublicationStatus `json:"publicationStatus"`
	Publisher         *common.Reference                          `json:"publisher,omitempty"`
	SubType           *common.CodeableConcept                    `json:"subType,omitempty"`
	Type              common.CodeableConcept                     `json:"type"`
}

// ContractSigner represents a signer of a contract
type ContractSigner struct {
	common.BackboneElement

	Party     common.Reference `json:"party"`
	Signature []Signature      `json:"signature"`
	Type      common.Coding    `json:"type"`
}

// ContractFriendly represents a patient-friendly version of a contract
type ContractFriendly struct {
	common.BackboneElement

	ContentAttachment *Attachment       `json:"contentAttachment,omitempty"`
	ContentReference  *common.Reference `json:"contentReference,omitempty"`
}

// ContractLegal represents legal expressions of a contract
type ContractLegal struct {
	common.BackboneElement

	ContentAttachment *Attachment       `json:"contentAttachment,omitempty"`
	ContentReference  *common.Reference `json:"contentReference,omitempty"`
}

// ContractRule represents computable policy rule language representations
type ContractRule struct {
	common.BackboneElement

	ContentAttachment *Attachment       `json:"contentAttachment,omitempty"`
	ContentReference  *common.Reference `json:"contentReference,omitempty"`
}

// Contract represents a legally enforceable contract
type Contract struct {
	DomainResource

	ResourceType string `json:"resourceType"` // Always "Contract"

	Alias                    []string                   `json:"alias,omitempty"`
	Applies                  *common.Period             `json:"applies,omitempty"`
	Author                   *common.Reference          `json:"author,omitempty"`
	Authority                []common.Reference         `json:"authority,omitempty"`
	ContentDefinition        *ContractContentDefinition `json:"contentDefinition,omitempty"`
	ContentDerivative        *common.CodeableConcept    `json:"contentDerivative,omitempty"`
	Domain                   []common.Reference         `json:"domain,omitempty"`
	ExpirationType           *common.CodeableConcept    `json:"expirationType,omitempty"`
	Friendly                 []ContractFriendly         `json:"friendly,omitempty"`
	Identifier               []common.Identifier        `json:"identifier,omitempty"`
	InstantiatesCanonical    *common.Reference          `json:"instantiatesCanonical,omitempty"`
	InstantiatesUri          *string                    `json:"instantiatesUri,omitempty"`
	Issued                   *string                    `json:"issued,omitempty"`
	Legal                    []ContractLegal            `json:"legal,omitempty"`
	LegallyBindingAttachment *Attachment                `json:"legallyBindingAttachment,omitempty"`
	LegallyBindingReference  *common.Reference          `json:"legallyBindingReference,omitempty"`
	LegalState               *common.CodeableConcept    `json:"legalState,omitempty"`
	Name                     *string                    `json:"name,omitempty"`
	RelevantHistory          []common.Reference         `json:"relevantHistory,omitempty"`
	Rule                     []ContractRule             `json:"rule,omitempty"`
	Scope                    *common.CodeableConcept    `json:"scope,omitempty"`
	Signer                   []ContractSigner           `json:"signer,omitempty"`
	Site                     []common.Reference         `json:"site,omitempty"`
	Status                   *ContractStatus            `json:"status,omitempty"`
	Subject                  []common.Reference         `json:"subject,omitempty"`
	Subtitle                 *string                    `json:"subtitle,omitempty"`
	SubType                  []common.CodeableConcept   `json:"subType,omitempty"`
	SupportingInfo           []common.Reference         `json:"supportingInfo,omitempty"`
	Term                     []ContractTerm             `json:"term,omitempty"`
	Title                    *string                    `json:"title,omitempty"`
	TopicCodeableConcept     *common.CodeableConcept    `json:"topicCodeableConcept,omitempty"`
	TopicReference           *common.Reference          `json:"topicReference,omitempty"`
	Type                     *common.CodeableConcept    `json:"type,omitempty"`
	URL                      *string                    `json:"url,omitempty"`
	Version                  *string                    `json:"version,omitempty"`
}

// ContractContentDefinitionPublicationStatus represents the publication status
type ContractContentDefinitionPublicationStatus string

const (
	ContractContentDefinitionPublicationStatusAmended        ContractContentDefinitionPublicationStatus = "amended"
	ContractContentDefinitionPublicationStatusAppended       ContractContentDefinitionPublicationStatus = "appended"
	ContractContentDefinitionPublicationStatusCancelled      ContractContentDefinitionPublicationStatus = "cancelled"
	ContractContentDefinitionPublicationStatusDisputed       ContractContentDefinitionPublicationStatus = "disputed"
	ContractContentDefinitionPublicationStatusEnteredInError ContractContentDefinitionPublicationStatus = "entered-in-error"
	ContractContentDefinitionPublicationStatusExecutable     ContractContentDefinitionPublicationStatus = "executable"
	ContractContentDefinitionPublicationStatusExecuted       ContractContentDefinitionPublicationStatus = "executed"
	ContractContentDefinitionPublicationStatusNegotiable     ContractContentDefinitionPublicationStatus = "negotiable"
	ContractContentDefinitionPublicationStatusOffered        ContractContentDefinitionPublicationStatus = "offered"
	ContractContentDefinitionPublicationStatusPolicy         ContractContentDefinitionPublicationStatus = "policy"
	ContractContentDefinitionPublicationStatusRejected       ContractContentDefinitionPublicationStatus = "rejected"
	ContractContentDefinitionPublicationStatusRenewed        ContractContentDefinitionPublicationStatus = "renewed"
	ContractContentDefinitionPublicationStatusRevoked        ContractContentDefinitionPublicationStatus = "revoked"
	ContractContentDefinitionPublicationStatusResolved       ContractContentDefinitionPublicationStatus = "resolved"
	ContractContentDefinitionPublicationStatusTerminated     ContractContentDefinitionPublicationStatus = "terminated"
)

// ContractStatus represents the status of a contract
type ContractStatus string

const (
	ContractStatusAmended        ContractStatus = "amended"
	ContractStatusAppended       ContractStatus = "appended"
	ContractStatusCancelled      ContractStatus = "cancelled"
	ContractStatusDisputed       ContractStatus = "disputed"
	ContractStatusEnteredInError ContractStatus = "entered-in-error"
	ContractStatusExecutable     ContractStatus = "executable"
	ContractStatusExecuted       ContractStatus = "executed"
	ContractStatusNegotiable     ContractStatus = "negotiable"
	ContractStatusOffered        ContractStatus = "offered"
	ContractStatusPolicy         ContractStatus = "policy"
	ContractStatusRejected       ContractStatus = "rejected"
	ContractStatusRenewed        ContractStatus = "renewed"
	ContractStatusRevoked        ContractStatus = "revoked"
	ContractStatusResolved       ContractStatus = "resolved"
	ContractStatusTerminated     ContractStatus = "terminated"
)

// ContractTerm represents a term in a contract (placeholder for now)
type ContractTerm struct {
	common.BackboneElement
	// TODO: Add ContractTerm fields when found in the TypeScript definition
}
