terraform {
  required_providers {
    gqldenring = {
      source = "kamsandhu93/gqldenring"
    }
  }
}

provider "gqldenring" {
  endpoint        = "http://localhost:8080/query"
  status_endpoint = "http://localhost:8080/health"
}

data "gqldenring_weapons" "weapons" {}

output "weapons" {
  value = data.gqldenring_weapons.weapons
}