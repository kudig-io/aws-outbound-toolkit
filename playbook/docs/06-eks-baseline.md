# 06 · EKS Baseline

A production EKS cluster is mostly a set of defaults. This chapter lists the defaults worth deciding
before the first workload lands. For runnable Terraform, see
[EKS production blueprint](../../infrastructure/eks-blueprint).

## Cluster mode

| Option | When it fits | Notes |
| --- | --- | --- |
| EKS Auto Mode | Teams wanting managed compute, networking, and storage defaults | Least operational surface |
| Managed node groups | Standard production baseline | You own addons and scaling policy |
| Karpenter | Dynamic, cost-sensitive, diverse instance needs | Requires subnet/SG discovery and disruption policy |
| Fargate | Small, bursty, isolated workloads | Per-pod cost and networking constraints |

Pick one primary compute model per cluster; mixing all of them increases operational cost.

## Networking

- VPC CNI with enough IP capacity for pod density; plan subnet sizing accordingly.
- Use **EKS Pod Identity** or IRSA for AWS API access; no node-level shared credentials.
- AWS Load Balancer Controller for ingress; avoid hand-managed load balancers.
- Network policies where the CNI supports them; default-deny for sensitive namespaces.
- Private API endpoint for production; restrict public access to approved CIDRs if required.

## Identity and permissions

- Map IAM roles to Kubernetes groups through EKS access entries.
- Separate cluster-admin, platform-operator, and developer roles.
- Use `eksctl`/Terraform access entries rather than editing the `aws-auth` ConfigMap by hand.
- Audit who can create cluster role bindings.

## Scaling and resilience

- Spread nodes across at least three AZs.
- Set PodDisruptionBudgets for critical services.
- Use topology spread constraints for multi-AZ balance.
- Configure Karpenter consolidation or cluster autoscaler with sane limits.
- Set resource requests/limits; avoid unbounded memory and CPU.
- Define priority classes for critical workloads.

## Addons and upgrades

- Pin addon versions; upgrade deliberately, not automatically.
- Track EKS version support windows and plan upgrades quarterly.
- Test upgrades in non-production with the same workload manifests.
- Keep node AMIs patched; use managed node group release versions or Karpenter AMI selection.

## Security

- Enable EKS audit logs to CloudWatch.
- Run GuardDuty EKS Runtime Monitoring where available.
- Scan images in CI and in the registry.
- Restrict hostPath, privileged containers, and host networking with admission policies.
- Use signed images and a trusted registry for production.
- Keep secrets in AWS Secrets Manager / SSM, mounted via CSI driver or external secrets.

## Observability

- Cluster-level: control plane logs, node metrics, kube-state-metrics.
- Workload-level: OTel metrics/traces/logs with resource attributes for service and environment.
- Alert on saturation, error rate, and latency, not only on pod restarts.
- Track addon health and certificate expiry.

## Cost

- Right-size requests; over-requested capacity is the most common EKS waste.
- Use Graviton node pools where images support arm64.
- Use Spot for stateless and batch workloads with interruption handling.
- Monitor per-namespace cost and cluster overhead.
- Review idle namespaces and orphaned load balancers monthly.

## Checklist

- [ ] Compute model chosen and documented
- [ ] Private API endpoint for production (or restricted public access)
- [ ] Pod Identity / IRSA configured; no static keys on nodes
- [ ] AWS Load Balancer Controller installed and tested
- [ ] Nodes spread across three AZs
- [ ] PodDisruptionBudgets and topology spread for critical services
- [ ] Resource requests/limits set for all workloads
- [ ] Priority classes defined
- [ ] Addon versions pinned; upgrade plan documented
- [ ] EKS audit logs enabled
- [ ] Image scanning and admission policies in place
- [ ] Secrets via Secrets Manager / SSM, not manifests
- [ ] Cluster and workload observability live
- [ ] Per-namespace cost visibility enabled

## Next

→ [07 · Observability](./07-observability.md)
