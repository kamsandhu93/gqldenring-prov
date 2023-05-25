default: testacc

# Install binary at ~/go/bin
install:
	go build -o ~/go/bin/terraform-provider-gqldenring

gen: # broken bc m1
	go generate ./...

gen-doc:
	terraform fmt -recursive ./examples/
	PWD=$(pwd) docker run -v ${PWD}:/code -w /code golang:1.20.4-bullseye \
		go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs@v0.14.1 \
		generate --tf-version v1.4.0 --provider-name gqldenring --rendered-provider-name GQLdenring

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
