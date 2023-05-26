# GQLdenring Terraform Provider
_This repository is built on the [Terraform Plugin Framework](https://github.com/hashicorp/terraform-plugin-framework)._

## What is this?

[GQLdenring](https://github.com/kamsandhu93/gqldenring) is a personal project I have developed for learning purposes. It 
is a Go gqlgen Graphql service that serves data on the popular video game Eldenring. This repo hosts a custom terraform provider
that interacts with this service. 

This is an educational repo, and is not intended for production use. 

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) >= 1.0
- [Go](https://golang.org/doc/install) >= 1.20

## Building The Provider

1. Clone the repository
1. Enter the repository directory
1. Build the provider using the make `make install` command:

```shell
make install
```

## Adding Dependencies

This provider uses [Go modules](https://github.com/golang/go/wiki/Modules).
Please see the Go documentation for the most up to date information about using Go modules.

To add a new dependency `github.com/author/dependency` to your Terraform provider:

```shell
go get github.com/author/dependency
go mod tidy
```

Then commit the changes to `go.mod` and `go.sum`.

## Using the provider

See the [docs](docs/index.md)


## Developing the provider 
Update ~/.terraformrc to something like:

```text
plugin_cache_dir   = "$HOME/.terraform.d/plugin-cache"

provider_installation {
        dev_overrides {
                "kamsandhu93/gqldenring" = "/Users/YOU/go/bin/"
        }
        direct{}
}
```

Make code changes.

Install the provider binary to go bin  
```bash
make install
```

Run acceptance tests (currently broke due to https://github.com/hashicorp/terraform-plugin-sdk/issues/1171)
```bash
make testacc
```

Run example terraform (requires gqldenring server to be running)
```bash
make apply
```

