package fhir4

// Binary represents a binary resource that contains content
type Binary struct {
	Resource

	// Resource Type Name (for serialization)
	ResourceType string `json:"resourceType"` // Always "Binary"

	// MimeType of the binary content represented as a standard MimeType (BCP 13)
	ContentType string `json:"contentType"`

	// The actual content, base64 encoded
	Data *string `json:"data,omitempty"`

	// A location where the data can be accessed
	URL *string `json:"url,omitempty"`

	// The number of bytes of data that make up this attachment
	Size *int64 `json:"size,omitempty"`

	// The calculated hash of the data using SHA-1
	Hash *string `json:"hash,omitempty"`

	// A label or set of text to display in place of the data
	Title *string `json:"title,omitempty"`

	// The date that the attachment was first created
	Creation *string `json:"creation,omitempty"`
}
