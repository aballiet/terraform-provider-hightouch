// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"hightouch/internal/sdk"
	"hightouch/internal/sdk/pkg/models/operations"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"hightouch/internal/validators"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &ModelDataSource{}
var _ datasource.DataSourceWithConfigure = &ModelDataSource{}

func NewModelDataSource() datasource.DataSource {
	return &ModelDataSource{}
}

// ModelDataSource is the data source implementation.
type ModelDataSource struct {
	client *sdk.Hightouch
}

// ModelDataSourceModel describes the data model.
type ModelDataSourceModel struct {
	CreatedAt   types.String            `tfsdk:"created_at"`
	Custom      *ModelCreateCustom      `tfsdk:"custom"`
	Dbt         *ModelCreateDbt         `tfsdk:"dbt"`
	ID          types.String            `tfsdk:"id"`
	IsSchema    types.Bool              `tfsdk:"is_schema"`
	ModelID     types.String            `tfsdk:"model_id"`
	Name        types.String            `tfsdk:"name"`
	PrimaryKey  types.String            `tfsdk:"primary_key"`
	QueryType   types.String            `tfsdk:"query_type"`
	Raw         *ModelCreateRaw         `tfsdk:"raw"`
	Slug        types.String            `tfsdk:"slug"`
	SourceID    types.String            `tfsdk:"source_id"`
	Syncs       []types.String          `tfsdk:"syncs"`
	Table       *ModelCreateTable       `tfsdk:"table"`
	Tags        map[string]types.String `tfsdk:"tags"`
	UpdatedAt   types.String            `tfsdk:"updated_at"`
	Visual      *ModelCreateVisual      `tfsdk:"visual"`
	WorkspaceID types.String            `tfsdk:"workspace_id"`
}

// Metadata returns the data source type name.
func (r *ModelDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_model"
}

// Schema defines the schema for the data source.
func (r *ModelDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Model DataSource",

		Attributes: map[string]schema.Attribute{
			"created_at": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
				Description: `The timestamp when model was created`,
			},
			"custom": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"query": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							validators.IsValidJSON(),
						},
						Description: `Parsed as JSON.`,
					},
				},
				Description: `Custom query for sources that doesn't support sql. For example, Airtable.`,
			},
			"dbt": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"compiled_sql": schema.StringAttribute{
						Computed:    true,
						Description: `Compiled SQL in the dbt model`,
					},
					"database": schema.StringAttribute{
						Computed:    true,
						Description: `Name of the database containing the generated table`,
					},
					"dbt_unique_id": schema.StringAttribute{
						Computed:    true,
						Description: `Unique ID of the model assigned by dbt (usually some combination of the schema and table name)`,
					},
					"model_id": schema.StringAttribute{
						Computed:    true,
						Description: `Model id that refer to a dbt model`,
					},
					"name": schema.StringAttribute{
						Computed:    true,
						Description: `Name of the table generated by the dbt model`,
					},
					"raw_sql": schema.StringAttribute{
						Computed:    true,
						Description: `Raw SQL in the dbt model`,
					},
					"schema": schema.StringAttribute{
						Computed:    true,
						Description: `Name of the schema containing the generated table`,
					},
				},
				Description: `Query that is based on a dbt model`,
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Description: `The id of the model`,
			},
			"is_schema": schema.BoolAttribute{
				Computed: true,
				MarkdownDescription: `If is_schema is true, the model is just used to build other models.` + "\n" +
					`Either as part of visual querying, or as the root of a visual query.`,
			},
			"model_id": schema.StringAttribute{
				Required:    true,
				Description: `The id of the model`,
			},
			"name": schema.StringAttribute{
				Computed:    true,
				Description: `The name of the model`,
			},
			"primary_key": schema.StringAttribute{
				Computed:    true,
				Description: `The primary key will be null if the query doesn't get directly synced (e.g. a relationship table for visual querying)`,
			},
			"query_type": schema.StringAttribute{
				Computed:    true,
				Description: `The type of the query. Available options: custom, raw_sql, tabel, dbt and visual.`,
			},
			"raw": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"sql": schema.StringAttribute{
						Computed: true,
					},
				},
				Description: `Standard raw SQL query`,
			},
			"slug": schema.StringAttribute{
				Computed:    true,
				Description: `The slug of the model`,
			},
			"source_id": schema.StringAttribute{
				Computed:    true,
				Description: `The id of the source that model is connected to`,
			},
			"syncs": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: `The list of id of syncs that uses this model`,
			},
			"table": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						Computed: true,
					},
				},
				Description: `Table-based query that fetches on a table instead of SQL`,
			},
			"tags": schema.MapAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: `The tags of the model`,
			},
			"updated_at": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
				Description: `The timestamp when model was lastly updated`,
			},
			"visual": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"filter": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							validators.IsValidJSON(),
						},
						Description: `Parsed as JSON.`,
					},
					"parent_id": schema.StringAttribute{
						Computed:    true,
						Description: `Parent id of the schema that visual query is based on`,
					},
					"primary_label": schema.StringAttribute{
						Computed: true,
					},
					"secondary_label": schema.StringAttribute{
						Computed: true,
					},
				},
				Description: `Visual query, used by audience`,
			},
			"workspace_id": schema.StringAttribute{
				Computed:    true,
				Description: `The id of the workspace where the model belongs to`,
			},
		},
	}
}

func (r *ModelDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.Hightouch)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected DataSource Configure Type",
			fmt.Sprintf("Expected *sdk.Hightouch, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *ModelDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *ModelDataSourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Config.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	modelID := data.ModelID.ValueString()
	request := operations.GetModelRequest{
		ModelID: modelID,
	}
	res, err := r.client.GetModel(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.Model == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromGetResponse(res.Model)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}