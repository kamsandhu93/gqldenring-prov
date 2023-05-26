TF_DIR ?= examples/resources/weapon

default: install

# Install binary at ~/go/bin
install:
	go build -o ~/go/bin/terraform-provider-gqldenring

gen: # broken on arm mac
	go generate ./...

lint:
	PWD=$(pwd) docker run -t --rm -v ${PWD}:/app -w /app golangci/golangci-lint:latest-alpine golangci-lint run -v

gen-doc: #workaround for arm mac
	terraform fmt -recursive ./examples/
	PWD=$(pwd) docker run --rm -v ${PWD}:/code -w /code golang:1.20.4-bullseye \
		go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs@v0.14.1 \
		generate --tf-version v1.4.0 --provider-name gqldenring --rendered-provider-name GQLdenring

init:
	terraform -chdir=${TF_DIR} init

plan:
	terraform -chdir=${TF_DIR} plan

apply:
	terraform -chdir=${TF_DIR} apply

clean:
	find . -name "*.tfstate" -type f -delete

# Run acceptance tests
testacc: install
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m


.PHONY: testacc install gen gen-doc init plan apply clean
