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