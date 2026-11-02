# Getting started

## Prerequisites

- Terraform >= 1.6 (or OpenTofu >= 1.6)
- AWS credentials with permission to create VPC, EKS, IAM, and CloudWatch resources
- `kubectl` and `helm` for post-install steps
- An S3 bucket and DynamoDB table if you want remote state (recommended)

## 1. Configure variables

```bash
cp examples/prod.tfvars terraform/terraform.tfvars
$EDITOR terraform/terraform.tfvars
```

At minimum set `region`, `cluster_name`, and `kubernetes_version`. Use a version currently in EKS
standard support. Decide whether the public API endpoint is acceptable; the default is private-only.

## 2. Configure remote state

Add a backend block to `terraform/versions.tf` (or a separate `backend.tf`):

```hcl
terraform {
  backend "s3" {
    bucket         = "my-terraform-state"
    key            = "eks-production-blueprint/terraform.tfstate"
    region         = "ap-southeast-1"
    dynamodb_table = "terraform-locks"
    encrypt        = true
  }
}
```

Commit the backend configuration, not the state file.

## 3. Plan and apply

```bash
make init
make plan
make apply
```

Review the plan for unexpected replacements, especially subnet CIDRs and node group settings.

## 4. Configure kubectl

```bash
aws eks update-kubeconfig \
  --name "$(terraform -chdir=terraform output -raw cluster_name)" \
  --region ap-southeast-1
kubectl get nodes
```

With a private-only endpoint, run this from a network path that can reach the VPC (VPN, bastion,
or a CI runner inside the VPC).

## 5. Grant human access

Access is managed through EKS access entries. Create an access entry for each IAM role, then
associate a policy:

```bash
aws eks create-access-entry \
  --cluster-name outbound-prod \
  --principal-arn arn:aws:iam::123456789012:role/platform-admin

aws eks associate-access-policy \
  --cluster-name outbound-prod \
  --principal-arn arn:aws:iam::123456789012:role/platform-admin \
  --policy-arn arn:aws:eks::aws:cluster-access-policy/AmazonEKSClusterAdminPolicy \
  --access-scope type=cluster
```

## 6. Install cluster addons

- AWS Load Balancer Controller: see [helm/README.md](../helm/README.md)
- Karpenter: see [helm/README.md](../helm/README.md) and create NodePool/EC2NodeClass resources
- Observability: deploy your OpenTelemetry collector and dashboards

## 7. Destroy

```bash
make destroy
```

Destroying the cluster deletes node groups, load balancers created by the controller, and the VPC.
Ensure you have exported any data you need first. Some controllers create resources outside
Terraform's graph; check for orphaned load balancers and security groups if destroy fails.
