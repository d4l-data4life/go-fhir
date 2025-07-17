package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// Contract represents the Contract resource
type Contract struct {
	DomainResource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Contract"

	// Alternative representation of the title for this Contract definition
	Alias []string `json:"alias,omitempty"`

	// Relevant time or time-period when this Contract is applicable
	Applies *common.Period `json:"applies,omitempty"`

	// The individual or organization that authored the Contract definition
	Author *common.Reference `json:"author,omitempty"`

	// A formally or informally recognized grouping of people, principals, organizations
	Authority []common.Reference `json:"authority,omitempty"`

	// Precusory content developed with a focus and intent of supporting the formation a Contract instance
	ContentDefinition *ContractContentDefinition `json:"contentDefinition,omitempty"`

	// The minimal content derived from the basal information source at a specific stage in its lifecycle
	ContentDerivative *common.CodeableConcept `json:"contentDerivative,omitempty"`

	// Recognized governance framework or system operating with a circumscribed scope
	Domain []common.Reference `json:"domain,omitempty"`

	// Event resulting in discontinuation or termination of this Contract instance
	ExpirationType *common.CodeableConcept `json:"expirationType,omitempty"`

	// The "patient friendly language" version of the Contract in whole or in parts
	Friendly []ContractFriendly `json:"friendly,omitempty"`

	// Unique identifier for this Contract or a derivative that references a Source Contract
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The URL pointing to a FHIR-defined Contract Definition that is adhered to in whole or part by this Contract
	InstantiatesCanonical *common.Reference `json:"instantiatesCanonical,omitempty"`

	// The URL pointing to an externally maintained definition that is adhered to in whole or in part by this Contract
	InstantiatesUri *string `json:"instantiatesUri,omitempty"`

	// When this Contract was issued
	Issued *string `json:"issued,omitempty"`

	// List of Legal expressions or representations of this Contract
	Legal []ContractLegal `json:"legal,omitempty"`

	// Legally binding Contract: This is the signed and legally recognized representation of the Contract
	LegallyBindingAttachment *common.Attachment `json:"legallyBindingAttachment,omitempty"`

	// Legally binding Contract: This is the signed and legally recognized representation of the Contract
	LegallyBindingReference *common.Reference `json:"legallyBindingReference,omitempty"`

	// Legal states of the formation of a legal instrument
	LegalState *common.CodeableConcept `json:"legalState,omitempty"`

	// The name is not expected to be globally unique
	Name *string `json:"name,omitempty"`

	// Links to Provenance records for past versions of this Contract definition
	RelevantHistory []common.Reference `json:"relevantHistory,omitempty"`

	// List of Computable Policy Rule Language Representations of this Contract
	Rule []ContractRule `json:"rule,omitempty"`

	// A selector of legal concerns for this Contract definition, derivative, or instance in any legal state
	Scope *common.CodeableConcept `json:"scope,omitempty"`

	// Signers who are principal parties to the contract are bound by the Contract.activity
	Signer []ContractSigner `json:"signer,omitempty"`

	// Sites in which the contract is complied with, exercised, or in force
	Site []common.Reference `json:"site,omitempty"`

	// This element is labeled as a modifier because the status contains codes that mark the contract as not currently valid or active
	Status *ContractStatus `json:"status,omitempty"`

	// The Contract.subject is an entity that has some role with respect to the Contract.topic
	Subject []common.Reference `json:"subject,omitempty"`

	// An explanatory or alternate user-friendly title for this Contract definition
	Subtitle *string `json:"subtitle,omitempty"`

	// Sub-category for the Contract that distinguishes the kinds of systems that would be interested in the Contract
	SubType []common.CodeableConcept `json:"subType,omitempty"`

	// Information that may be needed by/relevant to the performer in their execution of this term action
	SupportingInfo []common.Reference `json:"supportingInfo,omitempty"`

	// One or more Contract Provisions, which may be related and conveyed as a group
	Term []ContractTerm `json:"term,omitempty"`

	// A short, descriptive, user-friendly title for this Contract definition
	Title *string `json:"title,omitempty"`

	// Narrows the range of legal concerns to focus on or achieve one or more Contract.topic to one or more Contract.topic.term
	TopicCodeableConcept *common.CodeableConcept `json:"topicCodeableConcept,omitempty"`

	// Narrows the range of legal concerns to focus on or achieve one or more Contract.topic to one or more Contract.topic.term
	TopicReference *common.Reference `json:"topicReference,omitempty"`

	// Precusory content structure and use, i.e., a boilerplate, template, application for a contract
	Type *common.CodeableConcept `json:"type,omitempty"`

	// Contract Valued Item List
	ValuedItem []ContractValuedItem `json:"valuedItem,omitempty"`
}

// ContractStatus represents the status of the contract
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

// ContractContentDefinition represents the content definition
type ContractContentDefinition struct {
	common.BackboneElement

	// A copyright statement relating to Contract precursor content
	Copyright *string `json:"copyright,omitempty"`

	// The date (and optionally time) when the contract was published
	PublicationDate *string `json:"publicationDate,omitempty"`

	// amended | appended | cancelled | disputed | entered-in-error | executable | executed | negotiable | offered | policy | rejected | renewed | revoked | resolved | terminated
	PublicationStatus ContractContentDefinitionPublicationStatus `json:"publicationStatus"`

	// The individual or organization that published the Contract precursor content
	Publisher *common.Reference `json:"publisher,omitempty"`

	// Detailed Precusory content type
	SubType *common.CodeableConcept `json:"subType,omitempty"`

	// Precusory content structure and use, i.e., a boilerplate, template, application for a contract
	Type common.CodeableConcept `json:"type"`
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

// ContractSigner represents a signer of the contract
type ContractSigner struct {
	common.BackboneElement

	// Party which is a signator to this Contract
	Party common.Reference `json:"party"`

	// Legally binding Contract DSIG signature contents in Base64
	Signature []common.Signature `json:"signature"`

	// Role of this Contract signer, e.g. notary, grantee
	Type common.Coding `json:"type"`
}

// ContractFriendly represents the "patient friendly language" version of the Contract
type ContractFriendly struct {
	common.BackboneElement

	// Human readable rendering of this Contract in a format and representation intended to enhance comprehension
	ContentAttachment *common.Attachment `json:"contentAttachment,omitempty"`

	// Human readable rendering of this Contract in a format and representation intended to enhance comprehension
	ContentReference *common.Reference `json:"contentReference,omitempty"`
}

// ContractLegal represents legal expressions or representations of this Contract
type ContractLegal struct {
	common.BackboneElement

	// Contract legal text in human renderable form
	ContentAttachment *common.Attachment `json:"contentAttachment,omitempty"`

	// Contract legal text in human renderable form
	ContentReference *common.Reference `json:"contentReference,omitempty"`
}

// ContractRule represents computable policy rule language representations of this Contract
type ContractRule struct {
	common.BackboneElement

	// Computable Contract conveyed using a policy rule language (e.g. XACML, DKAL, SecPal)
	ContentAttachment *common.Attachment `json:"contentAttachment,omitempty"`

	// Computable Contract conveyed using a policy rule language (e.g. XACML, DKAL, SecPal)
	ContentReference *common.Reference `json:"contentReference,omitempty"`
}

// ContractValuedItem represents a valued item in the contract
type ContractValuedItem struct {
	common.BackboneElement

	// Indicates the time during which this Contract ValuedItem information is effective
	EffectiveTime *string `json:"effectiveTime,omitempty"`

	// Specific type of Contract Valued Item that may be priced
	EntityCodeableConcept *common.CodeableConcept `json:"entityCodeableConcept,omitempty"`

	// Specific type of Contract Valued Item that may be priced
	EntityReference *common.Reference `json:"entityReference,omitempty"`

	// A real number that represents a multiplier used in determining the overall value
	Factor *float64 `json:"factor,omitempty"`

	// Identifies a Contract Valued Item instance
	Identifier *common.Identifier `json:"identifier,omitempty"`

	// Id of the clause or question text related to the context of this valuedItem
	LinkId []string `json:"linkId,omitempty"`

	// Expresses the product of the Contract Valued Item unitQuantity and the unitPriceAmt
	Net *common.Money `json:"net,omitempty"`

	// Terms of valuation
	Payment *string `json:"payment,omitempty"`

	// When payment is due
	PaymentDate *string `json:"paymentDate,omitempty"`

	// An amount that expresses the weighting associated with the Contract Valued Item delivered
	Points *float64 `json:"points,omitempty"`

	// Specifies the units by which the Contract Valued Item is measured or counted
	Quantity *common.Quantity `json:"quantity,omitempty"`

	// Who will receive payment
	Recipient *common.Reference `json:"recipient,omitempty"`

	// Who will make payment
	Responsible *common.Reference `json:"responsible,omitempty"`

	// A set of security labels that define which terms are controlled by this condition
	SecurityLabelNumber []int `json:"securityLabelNumber,omitempty"`

	// A Contract Valued Item unit valuation measure
	UnitPrice *common.Money `json:"unitPrice,omitempty"`
}

// ContractTerm represents one or more Contract Provisions
type ContractTerm struct {
	common.BackboneElement

	// Several agents may be associated with an activity and vice-versa
	Action []ContractTermAction `json:"action,omitempty"`

	// Relevant time or time-period when this Contract Provision is applicable
	Applies *common.Period `json:"applies,omitempty"`

	// Contract Term Asset List
	Asset []ContractTermAsset `json:"asset,omitempty"`

	// Nested group of Contract Provisions
	Group []ContractTerm `json:"group,omitempty"`

	// Unique identifier for this particular Contract Provision
	Identifier *common.Identifier `json:"identifier,omitempty"`

	// When this Contract Provision was issued
	Issued *string `json:"issued,omitempty"`

	// The matter of concern in the context of this provision of the agreement
	Offer ContractTermOffer `json:"offer"`

	// Security labels that protect the handling of information about the term and its elements
	SecurityLabel []ContractTermSecurityLabel `json:"securityLabel,omitempty"`

	// A specialized legal clause or condition based on overarching contract type
	SubType *common.CodeableConcept `json:"subType,omitempty"`

	// Statement of a provision in a policy or a contract
	Text *string `json:"text,omitempty"`

	// The entity that the term applies to
	TopicCodeableConcept *common.CodeableConcept `json:"topicCodeableConcept,omitempty"`

	// The entity that the term applies to
	TopicReference *common.Reference `json:"topicReference,omitempty"`

	// A legal clause or condition contained within a contract
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// ContractTermAction represents several agents associated with an activity
type ContractTermAction struct {
	common.BackboneElement

	// Encounter or Episode with primary association to specified term activity
	Context *common.Reference `json:"context,omitempty"`

	// Id of the clause or question text related to the requester of this action
	ContextLinkId []string `json:"contextLinkId,omitempty"`

	// True if the term prohibits the action
	DoNotPerform *bool `json:"doNotPerform,omitempty"`

	// Reason or purpose for the action stipulated by this Contract Provision
	Intent common.CodeableConcept `json:"intent"`

	// Id of the clause or question text related to this action
	LinkId []string `json:"linkId,omitempty"`

	// Comments made about the term action made by the requester, performer, subject or other participants
	Note []common.Annotation `json:"note,omitempty"`

	// When action happens
	OccurrenceDateTime *string `json:"occurrenceDateTime,omitempty"`

	// When action happens
	OccurrencePeriod *common.Period `json:"occurrencePeriod,omitempty"`

	// When action happens
	OccurrenceTiming *common.Timing `json:"occurrenceTiming,omitempty"`

	// Indicates who or what is being asked to perform (or not perform) the action
	Performer *common.Reference `json:"performer,omitempty"`

	// Id of the clause or question text related to the reason type or reference of this action
	PerformerLinkId []string `json:"performerLinkId,omitempty"`

	// The type of role or competency of an individual desired or required to perform or not perform the action
	PerformerRole *common.CodeableConcept `json:"performerRole,omitempty"`

	// The type of individual that is desired or required to perform or not perform the action
	PerformerType []common.CodeableConcept `json:"performerType,omitempty"`

	// Describes why the action is to be performed or not performed in textual form
	Reason []string `json:"reason,omitempty"`

	// Rationale for the action to be performed or not performed
	ReasonCode []common.CodeableConcept `json:"reasonCode,omitempty"`

	// Id of the clause or question text related to the reason type or reference of this action
	ReasonLinkId []string `json:"reasonLinkId,omitempty"`

	// Indicates another resource whose existence justifies permitting or not permitting this action
	ReasonReference []common.Reference `json:"reasonReference,omitempty"`

	// Who or what initiated the action and has responsibility for its activation
	Requester []common.Reference `json:"requester,omitempty"`

	// Id of the clause or question text related to the requester of this action
	RequesterLinkId []string `json:"requesterLinkId,omitempty"`

	// Security labels that protects the action
	SecurityLabelNumber []int `json:"securityLabelNumber,omitempty"`

	// Current state of the term action
	Status common.CodeableConcept `json:"status"`

	// Entity of the action
	Subject []ContractTermActionSubject `json:"subject,omitempty"`

	// Activity or service obligation to be done or not done, performed or not performed
	Type common.CodeableConcept `json:"type"`
}

// ContractTermActionSubject represents entity of the action
type ContractTermActionSubject struct {
	common.BackboneElement

	// The entity the action is performed or not performed on or for
	Reference []common.Reference `json:"reference"`

	// Role type of agent assigned roles in this Contract
	Role *common.CodeableConcept `json:"role,omitempty"`
}

// ContractTermAsset represents Contract Term Asset List
type ContractTermAsset struct {
	common.BackboneElement

	// Response to assets
	Answer []ContractTermOfferAnswer `json:"answer,omitempty"`

	// Description of the quality and completeness of the asset that may be a factor in its valuation
	Condition *string `json:"condition,omitempty"`

	// Circumstance of the asset
	Context []ContractTermAssetContext `json:"context,omitempty"`

	// Id of the clause or question text about the asset in the referenced form or QuestionnaireResponse
	LinkId []string `json:"linkId,omitempty"`

	// Asset relevant contractual time period
	Period []common.Period `json:"period,omitempty"`

	// Type of Asset availability for use or ownership
	PeriodType []common.CodeableConcept `json:"periodType,omitempty"`

	// Specifies the applicability of the term to an asset resource instance
	Relationship *common.Coding `json:"relationship,omitempty"`

	// Differentiates the kind of the asset
	Scope *common.CodeableConcept `json:"scope,omitempty"`

	// Security labels that protects the asset
	SecurityLabelNumber []int `json:"securityLabelNumber,omitempty"`

	// May be a subtype or part of an offered asset
	Subtype []common.CodeableConcept `json:"subtype,omitempty"`

	// Clause or question text concerning the asset in a linked form
	Text *string `json:"text,omitempty"`

	// Target entity type about which the term may be concerned
	Type []common.CodeableConcept `json:"type,omitempty"`

	// Associated entities
	TypeReference []common.Reference `json:"typeReference,omitempty"`

	// Time period of asset use
	UsePeriod []common.Period `json:"usePeriod,omitempty"`

	// Contract Valued Item List
	ValuedItem []ContractTermAssetValuedItem `json:"valuedItem,omitempty"`
}

// ContractTermAssetContext represents circumstance of the asset
type ContractTermAssetContext struct {
	common.BackboneElement

	// Coded representation of the context generally or of the Referenced entity
	Code []common.CodeableConcept `json:"code,omitempty"`

	// Asset context reference may include the creator, custodian, or owning Person or Organization
	Reference *common.Reference `json:"reference,omitempty"`

	// Context description
	Text *string `json:"text,omitempty"`
}

// ContractTermAssetValuedItem represents Contract Valued Item List
type ContractTermAssetValuedItem struct {
	common.BackboneElement

	// Indicates the time during which this Contract ValuedItem information is effective
	EffectiveTime *string `json:"effectiveTime,omitempty"`

	// Specific type of Contract Valued Item that may be priced
	EntityCodeableConcept *common.CodeableConcept `json:"entityCodeableConcept,omitempty"`

	// Specific type of Contract Valued Item that may be priced
	EntityReference *common.Reference `json:"entityReference,omitempty"`

	// A real number that represents a multiplier used in determining the overall value
	Factor *float64 `json:"factor,omitempty"`

	// Identifies a Contract Valued Item instance
	Identifier *common.Identifier `json:"identifier,omitempty"`

	// Id of the clause or question text related to the context of this valuedItem
	LinkId []string `json:"linkId,omitempty"`

	// Expresses the product of the Contract Valued Item unitQuantity and the unitPriceAmt
	Net *common.Money `json:"net,omitempty"`

	// Terms of valuation
	Payment *string `json:"payment,omitempty"`

	// When payment is due
	PaymentDate *string `json:"paymentDate,omitempty"`

	// An amount that expresses the weighting associated with the Contract Valued Item delivered
	Points *float64 `json:"points,omitempty"`

	// Specifies the units by which the Contract Valued Item is measured or counted
	Quantity *common.Quantity `json:"quantity,omitempty"`

	// Who will receive payment
	Recipient *common.Reference `json:"recipient,omitempty"`

	// Who will make payment
	Responsible *common.Reference `json:"responsible,omitempty"`

	// A set of security labels that define which terms are controlled by this condition
	SecurityLabelNumber []int `json:"securityLabelNumber,omitempty"`

	// A Contract Valued Item unit valuation measure
	UnitPrice *common.Money `json:"unitPrice,omitempty"`
}

// ContractTermOffer represents the matter of concern in the context of this provision of the agreement
type ContractTermOffer struct {
	common.BackboneElement

	// Response to offer text
	Answer []ContractTermOfferAnswer `json:"answer,omitempty"`

	// Type of choice made by accepting party with respect to an offer made by an offeror/ grantee
	Decision *common.CodeableConcept `json:"decision,omitempty"`

	// How the decision about a Contract was conveyed
	DecisionMode []common.CodeableConcept `json:"decisionMode,omitempty"`

	// Unique identifier for this particular Contract Provision
	Identifier []common.Identifier `json:"identifier,omitempty"`

	// The id of the clause or question text of the offer in the referenced questionnaire/response
	LinkId []string `json:"linkId,omitempty"`

	// Offer Recipient
	Party []ContractTermOfferParty `json:"party,omitempty"`

	// Security labels that protects the offer
	SecurityLabelNumber []int `json:"securityLabelNumber,omitempty"`

	// Human readable form of this Contract Offer
	Text *string `json:"text,omitempty"`

	// The Contract.topic may be an application for or offer of a policy or service
	Topic *common.Reference `json:"topic,omitempty"`

	// Type of Contract Provision such as specific requirements, purposes for actions, obligations, prohibitions
	Type *common.CodeableConcept `json:"type,omitempty"`
}

// ContractTermOfferAnswer represents response to offer text
type ContractTermOfferAnswer struct {
	common.BackboneElement

	// Response to an offer clause or question text
	ValueBoolean *bool `json:"valueBoolean,omitempty"`

	// Response to an offer clause or question text
	ValueDecimal *float64 `json:"valueDecimal,omitempty"`

	// Response to an offer clause or question text
	ValueInteger *int `json:"valueInteger,omitempty"`

	// Response to an offer clause or question text
	ValueDate *string `json:"valueDate,omitempty"`

	// Response to an offer clause or question text
	ValueDateTime *string `json:"valueDateTime,omitempty"`

	// Response to an offer clause or question text
	ValueTime *string `json:"valueTime,omitempty"`

	// Response to an offer clause or question text
	ValueString *string `json:"valueString,omitempty"`

	// Response to an offer clause or question text
	ValueUri *string `json:"valueUri,omitempty"`

	// Response to an offer clause or question text
	ValueAttachment *common.Attachment `json:"valueAttachment,omitempty"`

	// Response to an offer clause or question text
	ValueCoding *common.Coding `json:"valueCoding,omitempty"`

	// Response to an offer clause or question text
	ValueQuantity *common.Quantity `json:"valueQuantity,omitempty"`

	// Response to an offer clause or question text
	ValueReference *common.Reference `json:"valueReference,omitempty"`
}

// ContractTermOfferParty represents Offer Recipient
type ContractTermOfferParty struct {
	common.BackboneElement

	// Participant in the offer
	Reference []common.Reference `json:"reference"`

	// How the party participates in the offer
	Role common.CodeableConcept `json:"role"`
}

// ContractTermSecurityLabel represents security labels that protect the handling of information
type ContractTermSecurityLabel struct {
	common.BackboneElement

	// Security label privacy tag that species the applicable privacy and security policies
	Category []common.Coding `json:"category,omitempty"`

	// Security label privacy tag that species the level of confidentiality protection required
	Classification common.Coding `json:"classification"`

	// Security label privacy tag that species the manner in which term and/or term elements are to be protected
	Control []common.Coding `json:"control,omitempty"`

	// Number used to link this term or term element to the applicable Security Label
	Number []int `json:"number,omitempty"`
}
