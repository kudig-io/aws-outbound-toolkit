# Templates

Small, copy-friendly starting points. These are **not** certified modules and are not maintained as
a Terraform registry product. Review, adapt, and test them in a non-production account first.

| Template | Purpose | Status |
| --- | --- | --- |
| [`terraform/account-baseline/`](./terraform/account-baseline/) | CloudTrail, Config, GuardDuty baseline sketch | draft |
| [`terraform/vpc/`](./terraform/vpc/) | Three-AZ VPC with private/public subnets | draft |
| [`kubernetes/`](./kubernetes/) | Namespace, PDB, and network-policy examples | draft |

## Conventions

- One template per directory with its own README and `variables.tf`.
- No real account IDs, ARNs, domains, or secrets.
- Pin provider and module versions.
- Include a short "what this creates / what it does not" section.
- Add a cost note for anything that creates billable resources.

For a fuller EKS implementation, see
[kudig-io/eks-production-blueprint](https://github.com/kudig-io/eks-production-blueprint).