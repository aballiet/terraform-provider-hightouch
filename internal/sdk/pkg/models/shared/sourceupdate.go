// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// SourceUpdate - The input for updating a Source
type SourceUpdate struct {
	// The source's configuration. This specifies general metadata about sources, like connection details
	// Hightouch will use this configuration to connect to underlying source.
	//
	// The schema depends on the source type.
	//
	// Consumers should NOT make assumptions on the contents of the
	// configuration. It may change as Hightouch updates its internal code.
	Configuration map[string]interface{} `json:"configuration,omitempty"`
	// The source's name
	Name *string `json:"name,omitempty"`
}
