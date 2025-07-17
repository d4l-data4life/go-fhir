package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// VerificationResultPrimarySource represents information about the primary source(s) involved in validation
type VerificationResultPrimarySource struct {
	common.BackboneElement

	// Ability of the primary source to push updates/alerts (yes; no; undetermined)
	CanPushUpdates *common.CodeableConcept `json:"canPushUpdates,omitempty"`

	// Method for communicating with the primary source (manual; API; Push)
	CommunicationMethod []common.CodeableConcept `json:"communicationMethod,omitempty"`

	// Type of alerts/updates the primary source can send (specific requested changes; any changes; as defined by source)
	PushTypeAvailable []common.CodeableConcept `json:"pushTypeAvailable,omitempty"`

	// Type of primary source (License Board; Primary Education; Continuing Education; Postal Service; Relationship owner; Registration Authority; legal source; issuing source; authoritative source)
	Type []common.CodeableConcept `json:"type,omitempty"`

	// When the target was validated against the primary source
	ValidationDate *string `json:"validationDate,omitempty"`

	// Status of the validation of the target against the primary source (successful; failed; unknown)
	ValidationStatus *common.CodeableConcept `json:"validationStatus,omitempty"`

	// Reference to the primary source
	Who *common.Reference `json:"who,omitempty"`
}

// VerificationResultAttestation represents information about the entity attesting to information
type VerificationResultAttestation struct {
	common.BackboneElement

	// The method by which attested information was submitted/retrieved (manual; API; Push)
	CommunicationMethod *common.CodeableConcept `json:"communicationMethod,omitempty"`

	// The date the information was attested to
	Date *string `json:"date,omitempty"`

	// A digital identity certificate associated with the attestation source
	OnBehalfOf *common.Reference `json:"onBehalfOf,omitempty"`

	// The procedure or code system that was used to perform the attestation
	ProxyIdentityCertificate *string `json:"proxyIdentityCertificate,omitempty"`

	// Signed assertion by the proxy entity indicating that they have the right to submit attested information on behalf of the attestation source
	ProxySignature *common.Signature `json:"proxySignature,omitempty"`

	// A digital identity certificate associated with the proxy entity submitting attested information on behalf of the attestation source
	SourceIdentityCertificate *string `json:"sourceIdentityCertificate,omitempty"`

	// Signed assertion by the attestation source that they have the right to submit attested information
	SourceSignature *common.Signature `json:"sourceSignature,omitempty"`

	// The individual or organization attesting to information
	Who *common.Reference `json:"who,omitempty"`
}

// VerificationResultValidator represents information about the entity validating information
type VerificationResultValidator struct {
	common.BackboneElement

	// A digital identity certificate associated with the validator
	AttestationSignature *common.Signature `json:"attestationSignature,omitempty"`

	// Validates a single target
	Organization common.Reference `json:"organization"`
}

// VerificationResultStatus represents the validation status of the target
type VerificationResultStatus string

const (
	VerificationResultStatusAttested       VerificationResultStatus = "attested"
	VerificationResultStatusValidated      VerificationResultStatus = "validated"
	VerificationResultStatusInProcess      VerificationResultStatus = "in-process"
	VerificationResultStatusReqRevalid     VerificationResultStatus = "req-revalid"
	VerificationResultStatusValFail        VerificationResultStatus = "val-fail"
	VerificationResultStatusRevalFail      VerificationResultStatus = "reval-fail"
	VerificationResultStatusEnteredInError VerificationResultStatus = "entered-in-error"
)

// VerificationResult represents describes validation requirements, source(s), status and dates for one or more elements
type VerificationResult struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "VerificationResult"

	// Information about the entity attesting to information
	Attestation *VerificationResultAttestation `json:"attestation,omitempty"`

	// The result if validation fails (fatal; warning; record only; none)
	FailureAction *common.CodeableConcept `json:"failureAction,omitempty"`

	// Frequency of revalidation
	Frequency *common.Timing `json:"frequency,omitempty"`

	// The date/time validation was last completed (including failed validations)
	LastPerformed *string `json:"lastPerformed,omitempty"`

	// The frequency with which the target must be validated (none; initial; periodic)
	Need *common.CodeableConcept `json:"need,omitempty"`

	// The date when target is next validated, if appropriate
	NextScheduled *string `json:"nextScheduled,omitempty"`

	// Information about the primary source(s) involved in validation
	PrimarySource []VerificationResultPrimarySource `json:"primarySource,omitempty"`

	// The validation status of the target (attested; validated; in process; requires revalidation; validation failed; revalidation failed)
	Status VerificationResultStatus `json:"status"`

	// When the validation status was updated
	StatusDate *string `json:"statusDate,omitempty"`

	// A resource that was validated
	Target []common.Reference `json:"target,omitempty"`

	// The fhirpath location(s) within the resource that was validated
	TargetLocation []string `json:"targetLocation,omitempty"`

	// The primary process by which the target is validated (edit check; value set; primary source; multiple sources; standalone; in context)
	ValidationProcess []common.CodeableConcept `json:"validationProcess,omitempty"`

	// What the target is validated against (nothing; primary source; multiple sources)
	ValidationType *common.CodeableConcept `json:"validationType,omitempty"`

	// Information about the entity validating information
	Validator []VerificationResultValidator `json:"validator,omitempty"`
}
