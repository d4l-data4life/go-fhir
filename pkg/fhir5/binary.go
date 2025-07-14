// Package fhir5 contains FHIR R5 (version 5.0.0) resource definitions
package fhir5

import (
	"github.com/go-fhir/go-fhir/pkg/common"
)

// Binary represents a resource that contains a single raw artifact as digital content
// accessible in its native format. A Binary resource can contain any content, whether text, image, pdf, zip archive, etc.
type Binary struct {
	Resource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Binary"

	// MimeType of the binary content represented as a standard MimeType (BCP 13)
	ContentType string `json:"contentType"`

	// The actual content, base64 encoded
	Data *string `json:"data,omitempty"`

	// Very often, a server will also know of a resource that references the binary, and can automatically apply the appropriate access rules based on that reference
	SecurityContext *common.Reference `json:"securityContext,omitempty"`
}
