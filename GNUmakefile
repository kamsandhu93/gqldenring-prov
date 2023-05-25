default: testacc

# Install binary at ~/go/bin
install:
	go build -o ~/go/bin/terraform-provider-gqldenring-tfprov

init-%:
	cd examples/$* && terraform init

plan-%:
	cd examples/$* && terraform plan

apply-%:
	cd examples/$* && terraform apply

# Run acceptance tests
.PHONY: testacc
testacc:
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m
