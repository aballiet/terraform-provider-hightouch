// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// DestinationCreate - The input for creating a Destination
type DestinationCreate struct {
	// The destination's configuration. This specifies general metadata about destination, like hostname and username.
	// Hightouch will be using this configuration to connect to destination.
	//
	// The schema depends on the destination type.
	//
	// Consumers should NOT make assumptions on the contents of the
	// configuration. It may change as Hightouch updates its internal code.
	Configuration map[string]interface{} `json:"configuration"`
	// The destination's name
	Name string `json:"name"`
	// The destination's slug
	Slug string `json:"slug"`
	// The destination's type (e.g. salesforce or hubspot).
	Type string `json:"type"`
}