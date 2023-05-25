package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hasura/go-graphql-client"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &WeaponsDataSource{}
var _ datasource.DataSourceWithConfigure = &WeaponsDataSource{}

func NewWeaponsDataSource() datasource.DataSource {
	return &WeaponsDataSource{}
}

// WeaponsDataSource defines the data source implementation.
type WeaponsDataSource struct {
	client *graphql.Client
}

// WeaponsDataSourceModel describes the data source data model.
type WeaponsDataSourceModel struct {
	Weapons []WeaponModel `tfsdk:"weapons"`
	Id      types.String  `tfsdk:"id"`
}

type WeaponModel struct {
	Name types.String `tfsdk:"name"`
	Id   types.String `tfsdk:"id"`
}

func (d *WeaponsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_weapons"
}

func (d *WeaponsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		Description: "Fetches a list of Weapons",
		// Schema defines the schema for the data source.

		Attributes: map[string]schema.Attribute{
			"weapons": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed:    true,
							Description: "The ID of a weapon.",
						},
						"name": schema.StringAttribute{
							Computed:    true,
							Description: "The name of a weapon.",
						},
					},
				},
			},
			"id": schema.StringAttribute{
				Computed: true,
			}},
	}
}

func (d *WeaponsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*graphql.Client)

}

func (d *WeaponsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state WeaponsDataSourceModel

	var query struct {
		Weapons []struct {
			Name string
			Id   string
		}
	}

	err := d.client.Query(context.Background(), &query, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to send gql query",
			err.Error(),
		)
		return
	}

	// Map response body to model
	for _, weapon := range query.Weapons {
		_weapon := WeaponModel{
			Id:   types.StringValue(weapon.Id),
			Name: types.StringValue(weapon.Name),
		}

		state.Weapons = append(state.Weapons, _weapon)

	}

	state.Id = types.StringValue("placeholder")

	// Set state
	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
