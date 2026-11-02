# Account baseline sketch

A sketch of the organization-level services every workload account should inherit.

## Recommended baseline

| Service | Scope | Purpose |
| --- | --- | --- |
| CloudTrail | Organization trail | API audit log with validation |
| AWS Config | All accounts | Resource inventory and compliance rules |
| GuardDuty | Organization, delegated admin | Threat detection |
| Security Hub | Organization, delegated admin | Findings aggregation and standards |
| IAM Identity Center | Organization | Human SSO and permission sets |
| S3 log archive | Dedicated account | Immutable retention |

## Implementation notes

- Prefer AWS Control Tower when starting fresh; use IaC when you need custom account structure.
- Delegate GuardDuty and Security Hub administration to the security account.
- Protect the log-archive bucket with bucket policies, Object Lock where required, and restricted
  access.
- Test SCP guardrails in a sandbox account before applying organization-wide.

This directory is a documentation sketch, not a runnable module yet. Contributions with tested
Terraform/OpenTofu code are welcome.
