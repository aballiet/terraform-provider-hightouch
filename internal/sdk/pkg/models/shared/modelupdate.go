// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ModelUpdateCustom - Custom query for sources that doesn't support sql. For example, Airtable.
type ModelUpdateCustom struct {
	Query interface{} `json:"query"`
}

type ModelUpdateDbt struct {
	// Model id that refers to a dbt model
	ModelID string `json:"modelId"`
}

// ModelUpdateRaw - Standard raw SQL query
type ModelUpdateRaw struct {
	SQL string `json:"sql"`
}

// ModelUpdateTable - Table-based query that fetches on a table instead of SQL
type ModelUpdateTable struct {
	Name string `json:"name"`
}

// ModelUpdateVisual - Visual query, used by audience
type ModelUpdateVisual struct {
	Filter interface{} `json:"filter"`
	// Parent id of the schema that visual query is based on
	ParentID       string `json:"parentId"`
	PrimaryLabel   string `json:"primaryLabel"`
	SecondaryLabel string `json:"secondaryLabel"`
}

// ModelUpdate - The input for updating a Model
type ModelUpdate struct {
	// Custom query for sources that doesn't support sql. For example, Airtable.
	Custom *ModelUpdateCustom `json:"custom,omitempty"`
	Dbt    *ModelUpdateDbt    `json:"dbt,omitempty"`
	// If is_schema is true, the model is just used to build other models.
	// Either as part of visual querying, or as the root of a visual query.
	IsSchema *bool `json:"isSchema,omitempty"`
	// The name of the model
	Name *string `json:"name,omitempty"`
	// The primary key will be null if the query doesn't get directly synced (e.g. a relationship table for visual querying)
	PrimaryKey *string `json:"primaryKey,omitempty"`
	// Standard raw SQL query
	Raw *ModelUpdateRaw `json:"raw,omitempty"`
	// Table-based query that fetches on a table instead of SQL
	Table *ModelUpdateTable `json:"table,omitempty"`
	// Visual query, used by audience
	Visual *ModelUpdateVisual `json:"visual,omitempty"`
}
