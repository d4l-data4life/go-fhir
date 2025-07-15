// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// BundleType represents the type of bundle
type BundleType string

const (
	BundleTypeDocument                 BundleType = "document"
	BundleTypeMessage                  BundleType = "message"
	BundleTypeTransaction              BundleType = "transaction"
	BundleTypeTransactionResponse      BundleType = "transaction-response"
	BundleTypeBatch                    BundleType = "batch"
	BundleTypeBatchResponse            BundleType = "batch-response"
	BundleTypeHistory                  BundleType = "history"
	BundleTypeSearchset                BundleType = "searchset"
	BundleTypeCollection               BundleType = "collection"
	BundleTypeSubscriptionNotification BundleType = "subscription-notification"
)

// BundleEntrySearchMode represents the mode of a search entry
type BundleEntrySearchMode string

const (
	BundleEntrySearchModeMatch   BundleEntrySearchMode = "match"
	BundleEntrySearchModeInclude BundleEntrySearchMode = "include"
	BundleEntrySearchModeOutcome BundleEntrySearchMode = "outcome"
)

// BundleEntryRequestMethod represents the HTTP method for a bundle entry request
type BundleEntryRequestMethod string

const (
	BundleEntryRequestMethodGET    BundleEntryRequestMethod = "GET"
	BundleEntryRequestMethodHEAD   BundleEntryRequestMethod = "HEAD"
	BundleEntryRequestMethodPOST   BundleEntryRequestMethod = "POST"
	BundleEntryRequestMethodPUT    BundleEntryRequestMethod = "PUT"
	BundleEntryRequestMethodDELETE BundleEntryRequestMethod = "DELETE"
	BundleEntryRequestMethodPATCH  BundleEntryRequestMethod = "PATCH"
)

// Signature represents a digital signature
type Signature struct {
	DataType

	// The technical format of the signature
	Type []common.Coding `json:"type"`

	// When the signature was created
	When *string `json:"when,omitempty"`

	// Who signed
	Who common.Reference `json:"who"`

	// The party represented
	OnBehalfOf *common.Reference `json:"onBehalfOf,omitempty"`

	// The technical format of the signed resources
	TargetFormat *string `json:"targetFormat,omitempty"`

	// The technical format of the signature
	SigFormat *string `json:"sigFormat,omitempty"`

	// The actual signature content (XML DigSig, JWS, etc.)
	Data *string `json:"data,omitempty"`
}

// BundleLink represents a link in a bundle
type BundleLink struct {
	common.BackboneElement

	// A name which details the functional use for this link
	Relation string `json:"relation"`

	// The reference details for the link
	URL string `json:"url"`
}

// BundleEntrySearch represents information about the search process that led to the creation of this entry
type BundleEntrySearch struct {
	common.BackboneElement

	// There is only one mode. In some corner cases, a resource may be included because it is both a match and an include
	Mode *BundleEntrySearchMode `json:"mode,omitempty"`

	// Servers are not required to return a ranking score. 1 is most relevant, and 0 is least relevant
	Score *float64 `json:"score,omitempty"`
}

// BundleEntryRequest represents additional information about how this entry should be processed as part of a transaction or batch
type BundleEntryRequest struct {
	common.BackboneElement

	// Only perform the operation if the Etag value matches
	IfMatch *string `json:"ifMatch,omitempty"`

	// Only perform the operation if the last updated date matches
	IfModifiedSince *string `json:"ifModifiedSince,omitempty"`

	// Instruct the server not to perform the create if a specified resource already exists
	IfNoneExist *string `json:"ifNoneExist,omitempty"`

	// If the ETag values match, return a 304 Not Modified status
	IfNoneMatch *string `json:"ifNoneMatch,omitempty"`

	// In a transaction or batch, this is the HTTP action to be executed for this entry
	Method BundleEntryRequestMethod `json:"method"`

	// The URL for the entry
	URL string `json:"url"`
}

// BundleEntryResponse represents the results of processing the corresponding 'request' entry in the batch or transaction
type BundleEntryResponse struct {
	common.BackboneElement

	// Etags match the Resource.meta.versionId
	ETag *string `json:"etag,omitempty"`

	// This has to match the same time in the meta header (meta.lastUpdated) if a resource is included
	LastModified *string `json:"lastModified,omitempty"`

	// The location header created by processing this operation
	Location *string `json:"location,omitempty"`

	// For a POST/PUT operation, this is the equivalent outcome that would be returned for prefer = operationoutcome
	Outcome interface{} `json:"outcome,omitempty"`

	// The status code returned by processing this entry
	Status string `json:"status"`
}

// BundleEntry represents an entry in a bundle resource
type BundleEntry struct {
	common.BackboneElement

	// The Absolute URL for the resource
	FullURL *string `json:"fullUrl,omitempty"`

	// A series of links that provide context to this entry
	Link []BundleLink `json:"link,omitempty"`

	// Additional information about how this entry should be processed as part of a transaction or batch
	Request *BundleEntryRequest `json:"request,omitempty"`

	// The Resource for the entry
	Resource interface{} `json:"resource,omitempty"`

	// Indicates the results of processing the corresponding 'request' entry in the batch or transaction
	Response *BundleEntryResponse `json:"response,omitempty"`

	// Information about the search process that lead to the creation of this entry
	Search *BundleEntrySearch `json:"search,omitempty"`
}

// Bundle represents a container for a collection of resources
type Bundle struct {
	Resource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Bundle"

	// An entry in a bundle resource - will either contain a resource or information about a resource
	Entry []BundleEntry `json:"entry,omitempty"`

	// Persistent identifier generally only matters for batches of type Document, Message, and Collection
	Identifier *common.Identifier `json:"identifier,omitempty"`

	// Issues must apply to the Bundle as a whole, not to individual entries
	Issues interface{} `json:"issues,omitempty"`

	// A series of links that provide context to this bundle
	Link []BundleLink `json:"link,omitempty"`

	// The signature could be created by the "author" of the bundle or by the originating device
	Signature *Signature `json:"signature,omitempty"`

	// For many bundles, the timestamp is equal to .meta.lastUpdated
	Timestamp *string `json:"timestamp,omitempty"`

	// Only used if the bundle is a search result set
	Total *int `json:"total,omitempty"`

	// It's possible to use a bundle for other purposes (e.g. a document can be accepted as a transaction)
	Type BundleType `json:"type"`
}
