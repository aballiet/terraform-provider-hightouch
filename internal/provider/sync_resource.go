// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"hightouch/internal/sdk"

	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"hightouch/internal/sdk/pkg/models/operations"
	"hightouch/internal/validators"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SyncResource{}
var _ resource.ResourceWithImportState = &SyncResource{}

func NewSyncResource() resource.Resource {
	return &SyncResource{}
}

// SyncResource defines the resource implementation.
type SyncResource struct {
	client *sdk.Hightouch
}

// SyncResourceModel describes the resource data model.
type SyncResourceModel struct {
	Configuration     map[string]types.String `tfsdk:"configuration"`
	CreatedAt         types.String            `tfsdk:"created_at"`
	DestinationID     types.String            `tfsdk:"destination_id"`
	Details           map[string]types.String `tfsdk:"details"`
	Disabled          types.Bool              `tfsdk:"disabled"`
	ID                types.String            `tfsdk:"id"`
	LastRunAt         types.String            `tfsdk:"last_run_at"`
	Message           types.String            `tfsdk:"message"`
	ModelID           types.String            `tfsdk:"model_id"`
	PrimaryKey        types.String            `tfsdk:"primary_key"`
	ReferencedColumns []types.String          `tfsdk:"referenced_columns"`
	Schedule          SyncCreateSchedule      `tfsdk:"schedule"`
	Slug              types.String            `tfsdk:"slug"`
	Status            types.String            `tfsdk:"status"`
	UpdatedAt         types.String            `tfsdk:"updated_at"`
	WorkspaceID       types.String            `tfsdk:"workspace_id"`
}

func (r *SyncResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sync"
}

func (r *SyncResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Sync Resource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.MapAttribute{
				Required:    true,
				ElementType: types.StringType,
				Validators: []validator.Map{
					mapvalidator.ValueStringsAre(validators.IsValidJSON()),
				},
				MarkdownDescription: `The sync's configuration. This specifies how data is mapped, among other` + "\n" +
					`configuration.` + "\n" +
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
				Description: `The timestamp when the sync was created`,
			},
			"destination_id": schema.StringAttribute{
				Required:    true,
				Description: `The id of the Destination that sync is connected to`,
			},
			"details": schema.MapAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Validators: []validator.Map{
					mapvalidator.ValueStringsAre(validators.IsValidJSON()),
				},
			},
			"disabled": schema.BoolAttribute{
				Required:    true,
				Description: `Whether the sync has been disabled by the user.`,
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Description: `The sync's id`,
			},
			"last_run_at": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
				Description: `The timestamp of the last sync run`,
			},
			"message": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"Validation failed",
					),
				},
				Description: `must be one of ["Validation failed"]`,
			},
			"model_id": schema.StringAttribute{
				Required:    true,
				Description: `The id of the Model that sync is connected to`,
			},
			"primary_key": schema.StringAttribute{
				Computed:    true,
				Description: `The primary key that sync uses to identify data from source`,
			},
			"referenced_columns": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: `The reference column that sync depends on to sync data from source`,
			},
			"schedule": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"schedule": schema.SingleNestedAttribute{
						Required: true,
						Attributes: map[string]schema.Attribute{
							"cron_schedule": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"expression": schema.StringAttribute{
										Required: true,
									},
								},
							},
							"dbt_schedule": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"account": schema.SingleNestedAttribute{
										Required: true,
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Required: true,
											},
										},
									},
									"dbt_credential_id": schema.StringAttribute{
										Required: true,
									},
									"job": schema.SingleNestedAttribute{
										Required: true,
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Required: true,
											},
										},
									},
								},
							},
							"interval_schedule": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"interval": schema.SingleNestedAttribute{
										Required: true,
										Attributes: map[string]schema.Attribute{
											"quantity": schema.NumberAttribute{
												Required: true,
											},
											"unit": schema.StringAttribute{
												Required: true,
												Validators: []validator.String{
													stringvalidator.OneOf(
														"minute",
														"hour",
														"day",
														"week",
													),
												},
												Description: `must be one of ["minute", "hour", "day", "week"]`,
											},
										},
									},
								},
							},
							"visual_cron_schedule": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"expressions": schema.ListNestedAttribute{
										Required: true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"days": schema.SingleNestedAttribute{
													Required: true,
													Attributes: map[string]schema.Attribute{
														"friday": schema.BoolAttribute{
															Optional: true,
														},
														"monday": schema.BoolAttribute{
															Optional: true,
														},
														"saturday": schema.BoolAttribute{
															Optional: true,
														},
														"sunday": schema.BoolAttribute{
															Optional: true,
														},
														"thursday": schema.BoolAttribute{
															Optional: true,
														},
														"tuesday": schema.BoolAttribute{
															Optional: true,
														},
														"wednesday": schema.BoolAttribute{
															Optional: true,
														},
													},
													Description: `Construct a type with a set of properties K of type T`,
												},
												"time": schema.StringAttribute{
													Required: true,
												},
											},
										},
									},
								},
							},
						},
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
					},
					"type": schema.StringAttribute{
						Required: true,
					},
				},
				MarkdownDescription: `The scheduling configuration. It can be triggerd based on several ways:` + "\n" +
					`` + "\n" +
					`Interval: the sync will be trigged based on certain interval(minutes/hours/days/weeks)` + "\n" +
					`` + "\n" +
					`Cron: the sync will be trigged based on cron expression https://en.wikipedia.org/wiki/Cron.` + "\n" +
					`` + "\n" +
					`Visual: the sync will be trigged based a visual cron configuration on UI` + "\n" +
					`` + "\n" +
					`DBT-cloud: the sync will be trigged based on a dbt cloud job`,
			},
			"slug": schema.StringAttribute{
				Required:    true,
				Description: `The sync's slug`,
			},
			"status": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"disabled",
						"pending",
						"cancelled",
						"failed",
						"queued",
						"success",
						"warning",
						"querying",
						"processing",
						"reporting",
						"interrupted",
					),
				},
				MarkdownDescription: `must be one of ["disabled", "pending", "cancelled", "failed", "queued", "success", "warning", "querying", "processing", "reporting", "interrupted"]` + "\n" +
					`SyncStatus`,
			},
			"updated_at": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
				Description: `The timestamp when the sync was last updated`,
			},
			"workspace_id": schema.StringAttribute{
				Computed:    true,
				Description: `The id of the workspace that the sync belongs to`,
			},
		},
	}
}

func (r *SyncResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.Hightouch)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.Hightouch, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *SyncResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *SyncResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &item)...)
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

	request := *data.ToCreateSDKType()
	res, err := r.client.CreateSync(ctx, request)
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
	if res.ValidateErrorJSON == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromCreateResponse(res.ValidateErrorJSON)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SyncResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *SyncResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
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
	request := operations.GetSyncRequest{
		ID: id,
	}
	res, err := r.client.GetSync(ctx, request)
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
	if res.Sync == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromGetResponse(res.Sync)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SyncResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *SyncResourceModel
	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	syncUpdate := *data.ToUpdateSDKType()
	id := data.ID.ValueString()
	request := operations.UpdateSyncRequest{
		SyncUpdate: syncUpdate,
		ID:         id,
	}
	res, err := r.client.UpdateSync(ctx, request)
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
	if res.ValidateErrorJSON == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromUpdateResponse(res.ValidateErrorJSON)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SyncResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *SyncResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
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

	// Not Implemented; entity does not have a configured DELETE operation
}

func (r *SyncResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
