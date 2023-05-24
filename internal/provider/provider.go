package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/path"
	graphql "github.com/hasura/go-graphql-client"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure GqldenringProvider satisfies various provider interfaces.
var _ provider.Provider = &GqldenringProvider{}

// GqldenringProvider defines the provider implementation.
type GqldenringProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// GqldenringProviderModel describes the provider data model.
type GqldenringProviderModel struct {
	Endpoint types.String `tfsdk:"endpoint"`
}

func (p *GqldenringProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "gqldenring"
	resp.Version = p.version
}

func (p *GqldenringProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"endpoint": schema.StringAttribute{
				MarkdownDescription: "Gqldenring GQL endpoint",
				Optional:            true,
			},
		},
	}
}

func (p *GqldenringProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data GqldenringProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if data.Endpoint.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("endpoint"),
			"Unknown gqldenring endpoint",
			"The provider cannot create the gqldenring API client as there is an unknown configuration value for the endpoint. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the GQLDENRING_ENDPOINT environment variable.",
		)
	}

	endpoint := os.Getenv("GQLDENRING_ENDPOINT")

	if !data.Endpoint.IsNull() {
		endpoint = data.Endpoint.ValueString()
	}

	if endpoint == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Missing Gqldenring API Endpoint",
			"The provider cannot create the Gqldenring API client as there is a missing or empty value for the Gqldenring API endpoint. "+
				"Set the host value in the configuration or use the GQLDENRING_ENDPOINT environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}
	// Init GQL client
	client := graphql.NewClient(endpoint, nil)

	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *GqldenringProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewExampleResource,
	}
}

func (p *GqldenringProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewWeaponsDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &GqldenringProvider{
			version: version,
		}
	}
}
