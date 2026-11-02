# 01 · Account & Landing Zone

The landing zone is the set of accounts, identities, guardrails, and logging that every workload
inherits. Getting it right early is the highest-leverage work in the whole playbook.

## Account strategy

Start with a small, explicit set of accounts. A common baseline:

| Account | Purpose | Notes |
| --- | --- | --- |
| `management` | AWS Organizations, billing, SSO | No workloads |
| `security` | GuardDuty, Security Hub, audit tooling | Delegated admin |
| `log-archive` | CloudTrail, VPC Flow Logs, Config history | Immutable retention |
| `network` | Transit Gateway, shared VPCs, DNS | Owns connectivity |
| `shared-services` | CI/CD, artifact registry, observability | Shared platform |
| `prod` | Production workloads | Strict guardrails |
| `non-prod` | Staging, QA, sandboxes | Looser, cost-capped |

Principles:

- Separate **prod** from **non-prod** at the account boundary, not just by naming convention.
- Give each product team its own account when blast radius or chargeback matters.
- Keep the management account free of workloads and long-lived access keys.
- Use **account vending** (Control Tower or IaC) so new accounts are identical every time.

## Identity

- Use **IAM Identity Center** (SSO) as the human access path; no long-lived IAM users.
- Map groups to **permission sets**, not individual policies.
- Require MFA for every human identity, including administrators.
- Give workloads **IAM roles** (IRSA / EKS Pod Identity / instance profiles), never static keys.
- Review unused roles and permissions quarterly.

## Guardrails

- Enable **AWS Organizations SCPs** to deny risky actions organization-wide, for example:
  - disabling CloudTrail or GuardDuty
  - creating IAM users with console access in production accounts
  - leaving the organization
  - using unapproved regions
- Use **AWS Control Tower** controls or equivalent Config rules for detective guardrails.
- Document every exception with an owner, reason, and expiry date.

## Logging & detection

- Organization CloudTrail trail to the log-archive account, with log-file validation.
- AWS Config in all accounts, aggregated to a central account.
- VPC Flow Logs for every VPC that carries production traffic.
- GuardDuty and Security Hub enabled organization-wide with delegated administration.
- Central retention policy aligned to legal and incident-response requirements.

## Tagging

Define a mandatory tag schema before onboarding teams. A workable minimum:

| Tag | Example | Purpose |
| --- | --- | --- |
| `env` | `prod` | Environment separation |
| `owner` | `team-payments` | Accountability |
| `service` | `checkout-api` | Cost and impact mapping |
| `cost-center` | `cc-1042` | Chargeback |
| `data-class` | `confidential` | Compliance and access review |

Enforce tags with SCPs, Config rules, or IaC policy checks. Untagged resources are the main reason
cost attribution fails later.

## Infrastructure as code

- Pick **Terraform/OpenTofu or CDK** and use it for the landing zone itself, not only workloads.
- Store state in a dedicated account with locking and versioning.
- Require review and automated policy checks (for example `tflint`, `checkov`, `cfn-nag`).
- Keep account vending and baseline configuration in one repository.

## Checklist

- [ ] AWS Organizations enabled; management account has no workloads
- [ ] Account structure documented, with owner and purpose per account
- [ ] IAM Identity Center configured; MFA enforced; no long-lived admin keys
- [ ] Permission sets mapped to groups; least privilege reviewed
- [ ] SCP guardrails applied and tested
- [ ] Organization CloudTrail enabled with log-file validation
- [ ] AWS Config, GuardDuty, Security Hub enabled organization-wide
- [ ] VPC Flow Logs enabled for production VPCs
- [ ] Mandatory tag schema defined and enforced
- [ ] Landing zone managed as code; account vending repeatable
- [ ] Break-glass procedure documented and tested

## Next

→ [02 · Network & Latency](./02-network-and-latency.md)