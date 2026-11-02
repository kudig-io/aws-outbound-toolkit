# 05 · Migration from China Clouds

Most outbound teams are not starting from zero; they are moving from Alibaba Cloud, Tencent Cloud,
Huawei Cloud, or self-managed IDC. The hard parts are usually data transfer, identity cutover, and
DNS, not compute.

## Service mapping

| Source (example) | AWS target | Notes |
| --- | --- | --- |
| ECS | EC2 / EKS / Fargate | Prefer containers for portability |
| OSS | S3 | Check consistency, permissions, and lifecycle |
| SLB / CLB | ALB / NLB | Map listeners, certificates, and health checks |
| RDS | RDS / Aurora | Engine versions and extensions may differ |
| Redis | ElastiCache | Cluster mode and persistence differ |
| Kafka | MSK | Topic/ACL migration and client compatibility |
| NAS | EFS / FSx | Throughput and protocol differences |
| VPC | VPC | CIDR overlap is the most common blocker |
| RAM | IAM + Organizations | Rebuild roles and trust relationships |
| DNS | Route 53 | TTL planning is critical for cutover |
| CloudMonitor | CloudWatch | Dashboards and alarms must be rebuilt |
| WAF | AWS WAF | Rule semantics differ; retest |

## Migration waves

1. **Wave 0 — foundations:** accounts, network, identity, logging, CI/CD.
2. **Wave 1 — stateless services:** APIs and workers with no local state.
3. **Wave 2 — caches and queues:** rebuild rather than lift-and-shift where possible.
4. **Wave 3 — databases:** replicate, validate, then cut over with a rollback window.
5. **Wave 4 — storage and analytics:** bulk data transfer and lifecycle policies.

Rules:

- Never migrate stateful data before the network and backup paths are proven.
- Prefer **rebuild over lift-and-shift** for managed services with different semantics.
- Keep source and target running in parallel long enough to validate behavior under real traffic.

## Data transfer

- Estimate volume, duration, and cost before starting.
- For large datasets, use AWS DataSync, Snowball, or a dedicated line; test throughput first.
- Checksum and reconcile row/object counts after transfer.
- Plan for incremental sync and a final delta window.
- Verify encryption in transit and at rest throughout.

## Cutover patterns

| Pattern | Use when | Trade-off |
| --- | --- | --- |
| DNS weighted cutover | Stateless or replicated services | TTL determines rollback speed |
| Dual-write | Databases with strict continuity needs | Complexity and consistency risk |
| Blue/green | Full environment can run in parallel | Cost of duplicate capacity |
| Big-bang | Small, low-risk workloads | Highest risk; avoid for critical systems |

Always define:

- the cutover window and decision owner
- success criteria and smoke tests
- rollback trigger, steps, and time limit
- communication plan for users and support

## Identity and access cutover

- Recreate human identities in IAM Identity Center; do not copy static keys.
- Rebuild service roles and trust policies; audit for over-permission.
- Rotate all credentials that touched the source environment.
- Remove source-cloud access only after the rollback window closes.

## Validation

- Functional smoke tests for every critical path.
- Performance tests from representative user locations.
- Data reconciliation reports for databases and object storage.
- Security scan of the new environment before opening traffic.
- Cost baseline captured for the first 30 days.

## Checklist

- [ ] Workload inventory and dependency map complete
- [ ] Service mapping reviewed, especially managed-service semantics
- [ ] VPC CIDR plan checked for overlap
- [ ] Data transfer method, duration, and cost estimated
- [ ] Migration waves and owners defined
- [ ] Cutover pattern and rollback plan documented
- [ ] DNS TTL lowered ahead of cutover
- [ ] Identity and service roles rebuilt, old credentials rotated
- [ ] Smoke, performance, and data-reconciliation tests passed
- [ ] Observability and alarms live in the target
- [ ] Rollback window and decision owner agreed
- [ ] Source-cloud access decommissioned after validation

## Next

→ [06 · EKS Baseline](./06-eks-baseline.md)
