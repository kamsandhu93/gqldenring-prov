package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hasura/go-graphql-client"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &WeaponResource{}
var _ resource.ResourceWithConfigure = &WeaponResource{}
var _ resource.ResourceWithImportState = &WeaponResource{}

func NewWeaponResource() resource.Resource {
	return &WeaponResource{}
}

// WeaponResource defines the resource implementation.
type WeaponResource struct {
	client *graphql.Client
}

// WeaponResourceModel describes the resource data model.
type WeaponResourceModel struct {
	Name   types.String `tfsdk:"name"`
	Custom types.Bool   `tfsdk:"custom"`
	Id     types.String `tfsdk:"id"`
}

func (r *WeaponResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_weapon"
}

func (r *WeaponResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		Description: "Manage a weapon.",

		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Description: "The name of the weapon.",
				Optional:    false,
				Required:    true,
				Computed:    false,
			},
			"id": schema.StringAttribute{
				Description: "The ID (uuid) of the weapon.",
				Optional:    false,
				Required:    false,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"custom": schema.BoolAttribute{
				Description: "Whether the weapon is custom.",
				Optional:    false,
				Required:    false,
				Computed:    true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *WeaponResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*graphql.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *graphql.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *WeaponResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *WeaponResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	type newWeapon struct {
		Name string `json:"name"`
	}

	var mut struct {
		CreateWeapon struct {
			Name   string
			Id     string
			Custom bool
		} `graphql:"createWeapon(input: $input)"`
	}

	vars := map[string]interface{}{
		"input": newWeapon{
			Name: data.Name.ValueString(),
		},
	}

	err := r.client.Mutate(context.Background(), &mut, vars)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to send create weapon mutation to Gqldenring",
			err.Error(),
		)
		return
	}

	data.Id = types.StringValue(mut.CreateWeapon.Id)
	data.Custom = types.BoolValue(mut.CreateWeapon.Custom)

	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "created a resource")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WeaponResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *WeaponResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var query struct {
		WeaponById struct {
			Name   string
			Id     string `graphql:"id"`
			Custom bool
		} `graphql:"WeaponById(id: $id)"`
	}

	vars := map[string]interface{}{
		"id": graphql.ID(data.Id.ValueString()),
	}

	err := r.client.Query(context.Background(), &query, vars)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to send WeaponById query to Gqldenring.",
			err.Error(),
		)
		return
	}

	data.Name = types.StringValue(query.WeaponById.Name)
	data.Custom = types.BoolValue(query.WeaponById.Custom)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WeaponResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *WeaponResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	type newWeapon struct {
		Name string `json:"name"`
	}

	var mut struct {
		UpdateWeapon struct {
			Name   string
			Id     string
			Custom bool
		} `graphql:"updateWeapon(id: $id, input: $input)"`
	}

	vars := map[string]interface{}{
		"input": newWeapon{
			Name: data.Name.ValueString(),
		},
		"id": graphql.ID(data.Id.ValueString()),
	}

	err := r.client.Mutate(context.Background(), &mut, vars)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to send UpdateWeapon mutation to Gqldenring.",
			err.Error(),
		)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WeaponResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *WeaponResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var mut struct {
		DeleteWeapon struct {
			Id string
		} `graphql:"deleteWeapon(id: $id)"`
	}

	vars := map[string]interface{}{
		"id": graphql.ID(data.Id.ValueString()),
	}

	err := r.client.Mutate(context.Background(), &mut, vars)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to send DeleteWeapon mutation to Gqldenring.",
			err.Error(),
		)
		return
	}
}

func (r *WeaponResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
