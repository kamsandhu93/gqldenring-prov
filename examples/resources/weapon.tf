terraform {
  required_providers {
    gqldenring = {
      source = "github.com/kamsandhu93/gqldenring"
    }
  }
}

provider "gqldenring" {
  endpoint = "http://localhost:8080/query"
}

resource "gqldenring_weapon" "example" {
  name = "tf-wep"
}

resource "gqldenring_weapon" "import" {
  name = "hello"
}

output "weapon" {
  value = gqldenring_weapon.example
}

output "weapon2" {
  value = gqldenring_weapon.import
}