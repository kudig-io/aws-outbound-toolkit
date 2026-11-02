# Kubernetes templates

Small manifests to copy into application repositories.

| File | Purpose |
| --- | --- |
| [`namespace-baseline.yaml`](./namespace-baseline.yaml) | Namespace labels, resource quota, and default limits |
| [`pdb.yaml`](./pdb.yaml) | PodDisruptionBudget example for a critical service |
| [`network-policy-default-deny.yaml`](./network-policy-default-deny.yaml) | Default-deny ingress/egress baseline |

## Conventions

- Replace `example` names and labels before use.
- Keep quotas and limits aligned with measured usage, not guesses.
- Test network policies in a non-production namespace first.