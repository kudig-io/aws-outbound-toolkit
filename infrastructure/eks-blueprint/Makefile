TF_DIR := terraform

.PHONY: fmt fmt-check init validate plan apply destroy clean

fmt:
	terraform -chdir=$(TF_DIR) fmt -recursive

fmt-check:
	terraform -chdir=$(TF_DIR) fmt -check -recursive

init:
	terraform -chdir=$(TF_DIR) init

validate: init
	terraform -chdir=$(TF_DIR) validate

plan: init
	terraform -chdir=$(TF_DIR) plan

apply: init
	terraform -chdir=$(TF_DIR) apply

destroy: init
	terraform -chdir=$(TF_DIR) destroy

clean:
	rm -rf $(TF_DIR)/.terraform $(TF_DIR)/.terraform.lock.hcl