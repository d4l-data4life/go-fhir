// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// ConsentPolicyBasis represents a policy basis for consent
type ConsentPolicyBasis struct {
	common.BackboneElement

	Reference *common.Reference `json:"reference,omitempty"`
	URL       *string           `json:"url,omitempty"`
}

// ConsentVerification represents verification of consent
type ConsentVerification struct {
	common.BackboneElement

	VerificationDate []string                `json:"verificationDate,omitempty"`
	VerificationType *common.CodeableConcept `json:"verificationType,omitempty"`
	Verified         bool                    `json:"verified"`
	VerifiedBy       *common.Reference       `json:"verifiedBy,omitempty"`
	VerifiedWith     *common.Reference       `json:"verifiedWith,omitempty"`
}

// ConsentProvisionActor represents an actor in a consent provision
type ConsentProvisionActor struct {
	common.BackboneElement

	Reference *common.Reference       `json:"reference,omitempty"`
	Role      *common.CodeableConcept `json:"role,omitempty"`
}

// ConsentProvisionData represents data controlled by a consent provision
type ConsentProvisionData struct {
	common.BackboneElement

	Meaning   ConsentProvisionDataMeaning `json:"meaning"`
	Reference common.Reference            `json:"reference"`
}

// ConsentProvision represents a provision in consent
type ConsentProvision struct {
	common.BackboneElement

	Action        []common.CodeableConcept `json:"action,omitempty"`
	Actor         []ConsentProvisionActor  `json:"actor,omitempty"`
	Code          []common.CodeableConcept `json:"code,omitempty"`
	Data          []ConsentProvisionData   `json:"data,omitempty"`
	DataPeriod    *common.Period           `json:"dataPeriod,omitempty"`
	DocumentType  []common.Coding          `json:"documentType,omitempty"`
	Expression    *Expression              `json:"expression,omitempty"`
	Period        *common.Period           `json:"period,omitempty"`
	Provision     []ConsentProvision       `json:"provision,omitempty"`
	Purpose       []common.Coding          `json:"purpose,omitempty"`
	ResourceType  []common.Coding          `json:"resourceType,omitempty"`
	SecurityLabel []common.Coding          `json:"securityLabel,omitempty"`
}

// Consent represents a record of healthcare consumer choices
type Consent struct {
	DomainResource

	ResourceType string `json:"resourceType"` // Always "Consent"

	Category         []common.CodeableConcept `json:"category,omitempty"`
	Controller       []common.Reference       `json:"controller,omitempty"`
	Date             *string                  `json:"date,omitempty"`
	Decision         *ConsentDecision         `json:"decision,omitempty"`
	Grantee          []common.Reference       `json:"grantee,omitempty"`
	Grantor          []common.Reference       `json:"grantor,omitempty"`
	Identifier       []common.Identifier      `json:"identifier,omitempty"`
	Manager          []common.Reference       `json:"manager,omitempty"`
	Period           *common.Period           `json:"period,omitempty"`
	PolicyBasis      *ConsentPolicyBasis      `json:"policyBasis,omitempty"`
	PolicyText       []common.Reference       `json:"policyText,omitempty"`
	Provision        []ConsentProvision       `json:"provision,omitempty"`
	RegulatoryBasis  []common.CodeableConcept `json:"regulatoryBasis,omitempty"`
	SourceAttachment []Attachment             `json:"sourceAttachment,omitempty"`
	SourceReference  []common.Reference       `json:"sourceReference,omitempty"`
	Status           ConsentStatus            `json:"status"`
	Subject          *common.Reference        `json:"subject,omitempty"`
	Verification     []ConsentVerification    `json:"verification,omitempty"`
}

// ConsentDecision represents the decision for consent
type ConsentDecision string

const (
	ConsentDecisionDeny   ConsentDecision = "deny"
	ConsentDecisionPermit ConsentDecision = "permit"
)

// ConsentStatus represents the status of consent
type ConsentStatus string

const (
	ConsentStatusDraft          ConsentStatus = "draft"
	ConsentStatusActive         ConsentStatus = "active"
	ConsentStatusInactive       ConsentStatus = "inactive"
	ConsentStatusNotDone        ConsentStatus = "not-done"
	ConsentStatusEnteredInError ConsentStatus = "entered-in-error"
	ConsentStatusUnknown        ConsentStatus = "unknown"
)

// ConsentProvisionDataMeaning represents how the resource reference is interpreted
type ConsentProvisionDataMeaning string

const (
	ConsentProvisionDataMeaningInstance   ConsentProvisionDataMeaning = "instance"
	ConsentProvisionDataMeaningRelated    ConsentProvisionDataMeaning = "related"
	ConsentProvisionDataMeaningDependents ConsentProvisionDataMeaning = "dependents"
	ConsentProvisionDataMeaningAuthoredBy ConsentProvisionDataMeaning = "authoredby"
)
