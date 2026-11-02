# 00 · Overview: the five stages of going global

Going global on AWS is less a single migration and more a sequence of decisions. Teams that skip a
stage usually pay for it later in security findings, surprise bills, or a rebuild of their account
structure.

| Stage | Question to answer | Typical duration | Primary output |
| --- | --- | --- | --- |
| 1. Assess | What are we moving, for whom, and under which regulations? | 1–3 weeks | Workload inventory, latency targets, compliance constraints |
| 2. Foundation | Where does it live, and who can touch it? | 2–6 weeks | Landing zone, account structure, identity, logging |
| 3. Land | How does the first workload reach production? | 2–8 weeks | Network, CI/CD, EKS/serverless baseline, observability |
| 4. Operate | How do we run it safely and affordably? | continuous | SLOs, on-call, cost cadence, patch process |
| 5. Optimize | What should we change next? | quarterly | Savings plan, architecture review, region expansion |

## Stage 1 · Assess

- Inventory workloads: stateless vs stateful, data classification, dependencies, traffic sources.
- Define latency budgets per user cohort (for example p95 < 200 ms for the primary market).
- List regulatory constraints: data residency, cross-border transfer, sector rules, contractual terms.
- Decide the **first region** and the **first workload**; resist the urge to migrate everything at once.

## Stage 2 · Foundation

- Choose an account strategy (single account, environment accounts, or full multi-account).
- Stand up AWS Organizations + IAM Identity Center + organization-wide CloudTrail.
- Define mandatory tags and guardrails (SCPs) before the first team onboards.
- Write down the landing-zone ownership model: who approves new accounts, network changes, and exceptions.

See [01 · Account & Landing Zone](./01-account-and-landing-zone.md).

## Stage 3 · Land

- Design VPC CIDRs so they can be peered or attached to a Transit Gateway later.
- Choose the edge strategy: CloudFront, Global Accelerator, or neither.
- Build the deployment path before the workload: IaC, CI/CD, secrets, rollback.
- Ship observability with the first workload, not after the first incident.

See [02 · Network & Latency](./02-network-and-latency.md), [06 · EKS Baseline](./06-eks-baseline.md),
and [07 · Observability](./07-observability.md).

## Stage 4 · Operate

- Define SLOs and error budgets per user-facing service.
- Establish an on-call rotation with clear escalation and runbooks.
- Review cost weekly for the first quarter, then monthly.
- Track patch cadence for nodes, addons, and base images.

## Stage 5 · Optimize

- Re-evaluate commitments (Savings Plans / Reserved Instances) after usage stabilizes.
- Review Graviton and Spot suitability for stateless and batch workloads.
- Consider additional regions only when latency, resilience, or regulation justifies the cost.

## Anti-patterns

- **One giant account** with production and experiments mixed together.
- **Migrating stateful data first** before the network and backup paths are proven.
- **Treating compliance as a post-launch task.**
- **No tagging from day one**, making cost attribution impossible later.
- **Launching without a rollback plan** because "the cutover window is short".

## Next

→ [01 · Account & Landing Zone](./01-account-and-landing-zone.md)
