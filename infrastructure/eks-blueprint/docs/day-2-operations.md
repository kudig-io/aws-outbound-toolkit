# Day-2 operations

## Upgrades

| Component | Cadence | Notes |
| --- | --- | --- |
| Kubernetes control plane | Quarterly or when a version approaches end of support | Test in non-prod first |
| Managed node group | After control plane | Use a surge or rolling update |
| EKS addons | With Kubernetes upgrades | Check compatibility before changing versions |
| Node AMIs | Monthly | Managed node groups handle release versions |
| Application images | Continuous | Scan and sign in CI |

Upgrade order: control plane → addons → nodes → workloads. Keep a rollback plan; EKS does not
support downgrading the control plane.

## Scaling

- Increase `node_max_size` before relying on the cluster autoscaler.
- Watch for pods stuck in `Pending` due to subnet IP exhaustion.
- Use topology spread constraints so replicas land across AZs.
- For Karpenter, define NodePools with instance family and capacity-type constraints, plus
  disruption budgets.

## Cost

- Review node utilization weekly; under 20% CPU is a downsizing signal.
- Use Graviton where images support arm64.
- Use Spot for stateless and batch workloads with disruption handling.
- Use VPC endpoints so pod traffic to S3/ECR does not traverse NAT.
- Track per-namespace cost with OpenCost or a similar tool.

## Observability

- Control plane logs are enabled for all five log types; route them to a central account.
- Collect node and pod metrics with a DaemonSet collector.
- Define SLOs per service and alert on burn rate.
- Track addon health, certificate expiry, and node pressure conditions.

## Security

- Enable GuardDuty EKS Runtime Monitoring where available.
- Use EKS access entries; audit who can create cluster role bindings.
- Apply admission policies to block privileged pods, hostPath, and host networking.
- Keep secrets in Secrets Manager/SSM and mount them with the CSI driver.
- Rotate credentials and review IAM permissions quarterly.

## Backup and recovery

- Back up persistent volumes with AWS Backup or Velero.
- Export cluster manifests and Helm values to Git.
- Test restore into a separate cluster quarterly.
- Document the RTO/RPO for each workload and verify them against the backup policy.

## Incident response

- Keep runbooks for: node NotReady, IP exhaustion, control plane throttling, certificate expiry,
  and failed deployments.
- Route alerts to an owned on-call rotation.
- Preserve control plane and audit logs for post-incident review.

## Decommissioning

- Delete workloads before node groups to avoid orphaned load balancers.
- Check for controller-created resources (ALBs, NLBs, security groups) after `terraform destroy`.
- Remove access entries and IRSA roles that are no longer used.
