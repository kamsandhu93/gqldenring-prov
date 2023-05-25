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

resource "gqldenring_weapon" "example" {
  name = "tf-wep-up"
}

output "weapon" {
  value = gqldenring_weapon.example
}