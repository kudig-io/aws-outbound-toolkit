# Launch readiness checklist

Use this as the final gate before sending production traffic to AWS. Each item should have an owner
and evidence (link to dashboard, config, or test result).

## Foundation

- [ ] Account structure, owners, and guardrails documented
- [ ] IAM Identity Center + MFA enforced; no long-lived admin keys
- [ ] Organization CloudTrail, Config, GuardDuty, Security Hub enabled
- [ ] Mandatory tags enforced; cost allocation working
- [ ] Break-glass procedure tested

## Network

- [ ] VPC CIDRs non-overlapping; subnets span three AZs
- [ ] Private subnets for workloads; NAT/endpoints sized and monitored
- [ ] DNS zones, health checks, and failover tested
- [ ] Edge strategy (CloudFront / Global Accelerator) validated
- [ ] Cross-border connectivity reviewed by legal/compliance
- [ ] Latency measured from representative user locations

## Security & compliance

- [ ] Data residency and transfer constraints confirmed
- [ ] Encryption at rest and in transit verified
- [ ] Secrets in Secrets Manager / SSM; none in images or repos
- [ ] WAF and TLS configuration reviewed
- [ ] Least-privilege roles reviewed; unused permissions removed
- [ ] Incident-response plan and on-call rotation live

## Application platform

- [ ] IaC covers all production resources; drift detection enabled
- [ ] CI/CD pipeline builds, tests, scans, and deploys reproducibly
- [ ] Rollback tested for the current release
- [ ] Resource requests/limits and autoscaling configured
- [ ] PodDisruptionBudgets and multi-AZ spread for critical services
- [ ] Backup and restore tested, including restore time

## Observability

- [ ] SLOs and error budgets defined
- [ ] Dashboards for service, dependencies, cluster, and business flows
- [ ] Alerts route to an owned channel with runbooks
- [ ] Logs are structured, retained, and free of sensitive data
- [ ] Traces cover critical user journeys

## Cost

- [ ] Budgets and alerts per account/team
- [ ] Commitment strategy decided (or explicitly deferred)
- [ ] Top cost drivers reviewed with owners
- [ ] Idle resources scanned before launch

## Cutover

- [ ] Cutover window, decision owner, and communication plan agreed
- [ ] Smoke tests defined and passing in the target environment
- [ ] Data reconciliation completed
- [ ] Rollback trigger, steps, and time limit documented
- [ ] Source environment kept available until the rollback window closes

## Sign-off

| Role | Name | Date |
| --- | --- | --- |
| Platform / SRE lead | | |
| Security lead | | |
| Product owner | | |
| Compliance / legal | | |