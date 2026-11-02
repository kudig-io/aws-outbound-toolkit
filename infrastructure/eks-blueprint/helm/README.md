# Helm values

Example values files for controllers that sit on top of the cluster. They are **not installed by
Terraform** in this blueprint yet; install them after the cluster is up.

| File | Chart | Notes |
| --- | --- | --- |
| [`values/aws-load-balancer-controller.yaml`](./values/aws-load-balancer-controller.yaml) | `aws-load-balancer-controller` | Requires an IRSA role and a policy matching the chart version |
| [`values/karpenter.yaml`](./values/karpenter.yaml) | `karpenter` | Requires controller role, node role, SQS interruption queue, and NodePool resources |

## Install order

1. Create the IRSA role and attach the required policy.
2. Install the chart with the matching values file.
3. Verify the controller logs and a test workload.

```bash
helm repo add eks https://aws.github.io/eks-charts
helm repo update

helm upgrade --install aws-load-balancer-controller eks/aws-load-balancer-controller \
  --namespace kube-system \
  --values helm/values/aws-load-balancer-controller.yaml
```

Always check the upstream chart documentation for the policy version that matches the release you
deploy. IAM policies for these controllers are version-sensitive.
