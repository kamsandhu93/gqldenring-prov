default: testacc

# Install binary at ~/go/bin
install:
	go build -o ~/go/bin/terraform-provider-gqldenring

init-%:
	cd examples/$* && terraform init

plan-%:
	cd examples/$* && terraform plan

apply-%:
	cd examples/$* && terraform apply

# Run acceptance tests
.PHONY: testacc
testacc: install
	TF_ACC_PROVIDER_NAMESPACE=github.com/kamsandhu93/ TF_CLI_CONFIG_FILE=~/.terraformrc TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m
