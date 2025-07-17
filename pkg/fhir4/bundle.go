// Package fhir4 contains FHIR R4 (version 4.0.1) resource definitions
package fhir4

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// Bundle represents a container for a collection of resources
type Bundle struct {
	Resource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Bundle"

	// Identifies the purpose of this bundle
	Type BundleType `json:"type"`

	// Optional timestamp when this bundle was assembled
	Timestamp *string `json:"timestamp,omitempty"`

	// If a set of search matches, this is the total number of matches for the search
	Total *int `json:"total,omitempty"`

	// A series of links that provide context to this bundle
	Link []BundleLink `json:"link,omitempty"`

	// An entry in a bundle resource
	Entry []BundleEntry `json:"entry,omitempty"`

	// Digital Signature
	Signature *Signature `json:"signature,omitempty"`

	// Persistent identity generally only matters for batches of type Document, Message, and Collection
	Identifier *common.Identifier `json:"identifier,omitempty"`
}

// BundleType represents the type of bundle
type BundleType string

const (
	BundleTypeDocument            BundleType = "document"
	BundleTypeMessage             BundleType = "message"
	BundleTypeTransaction         BundleType = "transaction"
	BundleTypeTransactionResponse BundleType = "transaction-response"
	BundleTypeBatch               BundleType = "batch"
	BundleTypeBatchResponse       BundleType = "batch-response"
	BundleTypeHistory             BundleType = "history"
	BundleTypeSearchset           BundleType = "searchset"
	BundleTypeCollection          BundleType = "collection"
)

// BundleLink represents a link in a bundle
type BundleLink struct {
	common.BackboneElement

	// A name which details the functional use for this link
	Relation string `json:"relation"`

	// The reference details for the link
	URL string `json:"url"`
}

// BundleEntry represents an entry in a bundle
type BundleEntry struct {
	common.BackboneElement

	// URI for resource (Absolute URL or relative URL)
	FullURL *string `json:"fullUrl,omitempty"`

	// A series of links that provide context to this entry
	Link []BundleLink `json:"link,omitempty"`

	// The Resource for the entry
	Resource interface{} `json:"resource,omitempty"`

	// Information about the search process that lead to the creation of this entry
	Search *BundleEntrySearch `json:"search,omitempty"`

	// Additional information about how this entry should be processed as part of a transaction or batch
	Request *BundleEntryRequest `json:"request,omitempty"`

	// Indicates the results of processing the corresponding 'request' entry
	Response *BundleEntryResponse `json:"response,omitempty"`
}

// BundleEntrySearch represents search information for a bundle entry
type BundleEntrySearch struct {
	common.BackboneElement

	// search | match | include | outcome - why this entry is in the result set
	Mode *BundleEntrySearchMode `json:"mode,omitempty"`

	// Search ranking (between 0..1)
	Score *float64 `json:"score,omitempty"`
}

// BundleEntrySearchMode represents the mode of a bundle entry search
type BundleEntrySearchMode string

const (
	BundleEntrySearchModeSearch  BundleEntrySearchMode = "search"
	BundleEntrySearchModeMatch   BundleEntrySearchMode = "match"
	BundleEntrySearchModeInclude BundleEntrySearchMode = "include"
	BundleEntrySearchModeOutcome BundleEntrySearchMode = "outcome"
)

// BundleEntryRequest represents request information for a bundle entry
type BundleEntryRequest struct {
	common.BackboneElement

	// In a transaction or batch, this is the HTTP action to be executed for this entry
	Method BundleEntryRequestMethod `json:"method"`

	// URL for HTTP equivalent of this entry
	URL string `json:"url"`

	// For managing cache currency
	IfNoneMatch *string `json:"ifNoneMatch,omitempty"`

	// For managing cache currency
	IfModifiedSince *string `json:"ifModifiedSince,omitempty"`

	// For managing update contention
	IfMatch *string `json:"ifMatch,omitempty"`

	// For conditional creates
	IfNoneExist *string `json:"ifNoneExist,omitempty"`
}

// BundleEntryRequestMethod represents the HTTP method for a bundle entry request
type BundleEntryRequestMethod string

const (
	BundleEntryRequestMethodGET    BundleEntryRequestMethod = "GET"
	BundleEntryRequestMethodPOST   BundleEntryRequestMethod = "POST"
	BundleEntryRequestMethodPUT    BundleEntryRequestMethod = "PUT"
	BundleEntryRequestMethodDELETE BundleEntryRequestMethod = "DELETE"
	BundleEntryRequestMethodPATCH  BundleEntryRequestMethod = "PATCH"
)

// BundleEntryResponse represents response information for a bundle entry
type BundleEntryResponse struct {
	common.BackboneElement

	// Status response code (text optional)
	Status string `json:"status"`

	// The location header created by processing this operation
	Location *string `json:"location,omitempty"`

	// The etag for the resource (if relevant)
	Etag *string `json:"etag,omitempty"`

	// Server's date time modified
	LastModified *string `json:"lastModified,omitempty"`

	// OperationOutcome with hints and warnings (for batch/transaction)
	Outcome interface{} `json:"outcome,omitempty"`
}
