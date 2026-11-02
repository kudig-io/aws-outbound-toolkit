# 04 · Cost Governance

Cloud cost is driven by a few large levers: compute commitment, storage lifecycle, and data
transfer. Governance is about making those levers visible and owned.

## Budget and alerting

- Create a budget per account, environment, and product team.
- Alert at 50%, 80%, and 100% of forecast, not only at the monthly total.
- Route alerts to the team that owns the spend, not only to finance.
- Review anomalies weekly during the first quarter after launch.

## Tagging and chargeback

- Enforce the tag schema from [01 · Account & Landing Zone](./01-account-and-landing-zone.md).
- Use cost allocation tags and a cost-allocation report per team.
- Treat untagged spend as a defect with an owner.
- Publish a monthly cost review: actual vs budget, top movers, and actions.

## Compute commitments

| Lever | When to use | Caveat |
| --- | --- | --- |
| Savings Plans (Compute) | Stable baseline usage across instance families | Commitment term, hourly spend |
| Savings Plans (EC2) | Predictable EC2-only usage | Less flexible than Compute SP |
| Reserved Instances | Very stable, specific instance usage | Least flexible |
| Spot | Batch, CI, fault-tolerant workers | Interruption handling required |
| Graviton | Most modern workloads | Validate image and dependency support |

Start with **savings plans for the stable baseline**, keep burst capacity on-demand, and revisit
after 30–60 days of real usage. Do not commit before the workload shape is stable.

## Storage and data lifecycle

- Move infrequently accessed objects to cheaper storage classes with lifecycle rules.
- Delete incomplete multipart uploads.
- Expire logs and snapshots according to retention policy, not forever.
- Review EBS volumes: unattached volumes and old snapshots are common waste.
- Use EBS snapshots lifecycle policies with a clear retention owner.

## Data transfer and networking

Data transfer is often the largest hidden cost for outbound workloads:

- Keep chatty service-to-service traffic inside the same AZ or region where possible.
- Use VPC endpoints for S3 and DynamoDB to avoid NAT data processing charges.
- Consolidate NAT gateways where resilience requirements allow.
- Put a CDN in front of cacheable egress.
- Monitor inter-region and internet egress separately.

## Kubernetes cost visibility

- Allocate cost by namespace, workload, and team using labels and a cost tool
  (OpenCost, Kubecost, or cloud-native cost allocation).
- Right-size requests and limits; over-requested capacity is paid for but unused.
- Use Karpenter or cluster autoscaler consolidation for bin-packing.
- Keep system and addon overhead visible in the unit economics.

## FinOps cadence

| Cadence | Activity |
| --- | --- |
| Weekly (first quarter) | Budget variance, anomaly review, idle resources |
| Monthly | Cost per service, commitment coverage, top movers |
| Quarterly | Architecture review, commitment renewal, unit economics |

## Checklist

- [ ] Budgets and alerts per account/environment/team
- [ ] Cost allocation tags enforced and reported
- [ ] Monthly cost review owned by engineering
- [ ] Commitment strategy documented and reviewed after usage stabilizes
- [ ] Storage lifecycle rules for objects, logs, snapshots, and backups
- [ ] Unattached EBS volumes and old snapshots reviewed
- [ ] VPC endpoints used for S3/DynamoDB
- [ ] NAT and inter-region egress monitored
- [ ] Kubernetes cost allocated by namespace/workload
- [ ] Requests and limits right-sized
- [ ] Unit economics (cost per request/user/tenant) tracked

## Next

→ [05 · Migration from China Clouds](./05-migration-from-china-clouds.md)
