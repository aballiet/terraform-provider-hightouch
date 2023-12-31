// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"hightouch/internal/sdk"
	"hightouch/internal/sdk/pkg/models/operations"

	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"hightouch/internal/validators"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &DestinationDataSource{}
var _ datasource.DataSourceWithConfigure = &DestinationDataSource{}

func NewDestinationDataSource() datasource.DataSource {
	return &DestinationDataSource{}
}

// DestinationDataSource is the data source implementation.
type DestinationDataSource struct {
	client *sdk.Hightouch
}

// DestinationDataSourceModel describes the data model.
type DestinationDataSourceModel struct {
	Configuration map[string]types.String `tfsdk:"configuration"`
	CreatedAt     types.String            `tfsdk:"created_at"`
	ID            types.String            `tfsdk:"id"`
	Name          types.String            `tfsdk:"name"`
	Slug          types.String            `tfsdk:"slug"`
	Syncs         []types.String          `tfsdk:"syncs"`
	Type          types.String            `tfsdk:"type"`
	UpdatedAt     types.String            `tfsdk:"updated_at"`
	WorkspaceID   types.String            `tfsdk:"workspace_id"`
}

// Metadata returns the data source type name.
func (r *DestinationDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_destination"
}

// Schema defines the schema for the data source.
func (r *DestinationDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Destination DataSource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.MapAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Validators: []validator.Map{
					mapvalidator.ValueStringsAre(validators.IsValidJSON()),
				},
				MarkdownDescription: `The destination's configuration. This specifies general metadata about destination, like hostname and username.` + "\n" +
					`Hightouch will be using this configuration to connect to destination.` + "\n" +
					`` + "\n" +
					`The schema depends on the destination type.` + "\n" +
					`` + "\n" +
					`Consumers should NOT make assumptions on the contents of the` + "\n" +
					`configuration. It may change as Hightouch updates its internal code.`,
			},
			"created_at": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
				Description: `The timestamp when the destination was created`,
			},
			"id": schema.StringAttribute{
				Required:    true,
				Description: `The destination's ID`,
			},
			"name": schema.StringAttribute{
				Computed:    true,
				Description: `The destination's name`,
			},
			"slug": schema.StringAttribute{
				Computed:    true,
				Description: `The destination's slug`,
			},
			"syncs": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: `A list of syncs that sync to this destination.`,
			},
			"type": schema.StringAttribute{
				Computed:    true,
				Description: `The destination's type (e.g. salesforce or hubspot).`,
			},
			"updated_at": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
				Description: `The timestamp when the destination was last updated`,
			},
			"workspace_id": schema.StringAttribute{
				Computed:    true,
				Description: `The id of the workspace that the destination belongs to`,
			},
		},
	}
}

func (r *DestinationDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *DestinationDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *DestinationDataSourceModel
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

	id := data.ID.ValueString()
	request := operations.GetDestinationRequest{
		ID: id,
	}
	res, err := r.client.GetDestination(ctx, request)
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
	if res.Destination == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromGetResponse(res.Destination)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
