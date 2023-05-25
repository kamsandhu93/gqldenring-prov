terraform {
  required_providers {
    gqldenring = {
      source = "github.com/kamsandhu93/gqldenring-tfprov"
    }
  }
}

provider "gqldenring" {
  endpoint = "http://localhost:8080/query"
}

data "gqldenring_weapons" "weapons" {}

output "weapons" {
  value = data.gqldenring_weapons.weapons
}