TF_DIR := infrastructure/eks-blueprint/terraform
GO_DIR := tools/cost-doctor

.PHONY: test test-go fmt fmt-go fmt-check fmt-check-go fmt-tf fmt-check-tf validate validate-tf run-cost

test: test-go

test-go:
	$(MAKE) -C $(GO_DIR) test

fmt: fmt-go fmt-tf

fmt-go:
	$(MAKE) -C $(GO_DIR) fmt

fmt-check: fmt-check-go fmt-check-tf

fmt-check-go:
	$(MAKE) -C $(GO_DIR) fmt-check

fmt-tf:
	terraform -chdir=$(TF_DIR) fmt -recursive

fmt-check-tf:
	terraform -chdir=$(TF_DIR) fmt -check -recursive

validate: validate-tf

validate-tf:
	terraform -chdir=$(TF_DIR) init -backend=false -input=false
	terraform -chdir=$(TF_DIR) validate

run-cost:
	$(MAKE) -C $(GO_DIR) run