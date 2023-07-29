// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// SyncCreateSchedule - The scheduling configuration. It can be triggerd based on several ways:
//
// Interval: the sync will be trigged based on certain interval(minutes/hours/days/weeks)
//
// Cron: the sync will be trigged based on cron expression https://en.wikipedia.org/wiki/Cron.
//
// Visual: the sync will be trigged based a visual cron configuration on UI
//
// DBT-cloud: the sync will be trigged based on a dbt cloud job
type SyncCreateSchedule struct {
	Schedule interface{} `json:"schedule"`
	Type     string      `json:"type"`
}

// SyncCreate - The input for creating a Sync
type SyncCreate struct {
	// The sync's configuration. This specifies how data is mapped, among other
	// configuration.
	//
	// The schema depends on the destination type.
	//
	// Consumers should NOT make assumptions on the contents of the
	// configuration. It may change as Hightouch updates its internal code.
	Configuration map[string]interface{} `json:"configuration"`
	// The id of the Destination that sync is connected to
	DestinationID string `json:"destinationId"`
	// Whether the sync has been disabled by the user.
	Disabled bool `json:"disabled"`
	// The id of the Model that sync is connected to
	ModelID string `json:"modelId"`
	// The scheduling configuration. It can be triggerd based on several ways:
	//
	// Interval: the sync will be trigged based on certain interval(minutes/hours/days/weeks)
	//
	// Cron: the sync will be trigged based on cron expression https://en.wikipedia.org/wiki/Cron.
	//
	// Visual: the sync will be trigged based a visual cron configuration on UI
	//
	// DBT-cloud: the sync will be trigged based on a dbt cloud job
	Schedule SyncCreateSchedule `json:"schedule"`
	// The sync's slug
	Slug string `json:"slug"`
}