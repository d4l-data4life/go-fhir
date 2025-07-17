package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// VerificationResult represents validation requirements, source(s), status and dates for one or more elements
type VerificationResult struct {
	DomainResource

	ResourceType string `json:"resourceType"` // Always "VerificationResult"

	Attestation       *VerificationResultAttestation    `json:"attestation,omitempty"`
	FailureAction     *common.CodeableConcept           `json:"failureAction,omitempty"`
	Frequency         *common.Timing                    `json:"frequency,omitempty"`
	LastPerformed     *string                           `json:"lastPerformed,omitempty"`
	Need              *common.CodeableConcept           `json:"need,omitempty"`
	NextScheduled     *string                           `json:"nextScheduled,omitempty"`
	PrimarySource     []VerificationResultPrimarySource `json:"primarySource,omitempty"`
	Status            *common.CodeableConcept           `json:"status,omitempty"`
	StatusDate        *string                           `json:"statusDate,omitempty"`
	Target            []common.Reference                `json:"target,omitempty"`
	TargetLocation    []string                          `json:"targetLocation,omitempty"`
	Validator         []VerificationResultValidator     `json:"validator,omitempty"`
	ValidationProcess []common.CodeableConcept          `json:"validationProcess,omitempty"`
	ValidationType    *common.CodeableConcept           `json:"validationType,omitempty"`
}

type VerificationResultPrimarySource struct {
	common.BackboneElement

	CanPushUpdates      *common.CodeableConcept  `json:"canPushUpdates,omitempty"`
	CommunicationMethod []common.CodeableConcept `json:"communicationMethod,omitempty"`
	PushTypeAvailable   []common.CodeableConcept `json:"pushTypeAvailable,omitempty"`
	Type                []common.CodeableConcept `json:"type,omitempty"`
	ValidationDate      *string                  `json:"validationDate,omitempty"`
	ValidationStatus    *common.CodeableConcept  `json:"validationStatus,omitempty"`
	Who                 *common.Reference        `json:"who,omitempty"`
}

type VerificationResultAttestation struct {
	common.BackboneElement

	CommunicationMethod       *common.CodeableConcept `json:"communicationMethod,omitempty"`
	Date                      *string                 `json:"date,omitempty"`
	OnBehalfOf                *common.Reference       `json:"onBehalfOf,omitempty"`
	ProxyIdentityCertificate  *string                 `json:"proxyIdentityCertificate,omitempty"`
	ProxySignature            *common.Signature       `json:"proxySignature,omitempty"`
	SourceIdentityCertificate *string                 `json:"sourceIdentityCertificate,omitempty"`
	SourceSignature           *common.Signature       `json:"sourceSignature,omitempty"`
	Who                       *common.Reference       `json:"who,omitempty"`
}

type VerificationResultValidator struct {
	common.BackboneElement

	AttestationSignature *common.Signature `json:"attestationSignature,omitempty"`
	IdentityCertificate  *string           `json:"identityCertificate,omitempty"`
	Organization         common.Reference  `json:"organization"`
}
