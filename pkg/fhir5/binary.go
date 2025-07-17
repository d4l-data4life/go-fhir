// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/d4l-data4life/go-fhir/pkg/common"
)

// Binary represents a binary resource
type Binary struct {
	Resource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Binary"

	// MimeType of the binary content represented as a standard MimeType (BCP 13)
	ContentType string `json:"contentType"`

	// If the content type is itself base64 encoding, then this will be base64 encoded twice
	Data *string `json:"data,omitempty"`

	// Very often, a server will also know of a resource that references the binary
	SecurityContext *common.Reference `json:"securityContext,omitempty"`
}
