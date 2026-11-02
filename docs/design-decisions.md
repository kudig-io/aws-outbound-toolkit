# Design decisions

Each decision below states the context, the choice, the trade-off, and the condition that should
make you revisit it. The goal is not to claim these are the only correct answers, but to make the
reasoning inspectable.

---

## 1. Private EKS API endpoint by default

**Context.** A public Kubernetes API endpoint is the most common way clusters are exposed to the
internet, and the most common source of "why is this cluster in a scan result?" findings.

**Decision.** `endpoint_public_access = false` by default. Access goes through a VPN, bastion, or a
CI runner inside the VPC.

**Trade-off.** Onboarding is slower; developers need a network path before `kubectl` works. Public
access remains a one-variable opt-in with a CIDR allowlist.

**Revisit when.** You have a managed VPN/Zero Trust path and a strong reason to allow public access,
or you are running a non-production sandbox where convenience dominates.

---

## 2. EKS access entries instead of the `aws-auth` ConfigMap

**Context.** The legacy `aws-auth` ConfigMap is easy to drift, hard to audit, and can lock you out
if edited incorrectly.

**Decision.** Use `authentication_mode = "API"` and manage access with EKS access entries. The
cluster creator is bootstrapped as admin, then real roles are added explicitly.

**Trade-off.** Tooling and documentation must use the access-entry API; older runbooks that edit
`aws-auth` no longer apply.

**Revisit when.** You must support a tool that only understands ConfigMap-based access. In that case,
use `API_AND_CONFIG_MAP` deliberately, not as a default.

---

## 3. Managed node group first, Karpenter later

**Context.** Karpenter is excellent for cost and elasticity, but it adds controller IAM, SQS
interruption handling, NodePool/EC2NodeClass design, and disruption tuning.

**Decision.** Ship a managed node group as the baseline. Document Karpenter as the next step rather
than installing it implicitly.

**Trade-off.** Less dynamic bin-packing and fewer instance-family options at the start. The cluster
is still multi-AZ and replaceable.

**Revisit when.** Workloads have heterogeneous capacity needs, batch/Spot requirements, or node
utilization shows persistent fragmentation.

---

## 4. Graviton (`m7g`) as the default instance family

**Context.** Most modern container images support arm64, and Graviton typically offers better
price/performance than comparable x86 instances.

**Decision.** Default to `m7g.large`; make `node_instance_types` a variable.

**Trade-off.** Some third-party or legacy images are amd64-only. Teams must verify image architecture
before adopting the default.

**Revisit when.** A dependency cannot run on arm64, or a workload is benchmarked and performs
materially better on x86.

---

## 5. One NAT gateway per AZ by default

**Context.** A single NAT gateway is cheaper but creates an AZ dependency for all private egress.

**Decision.** Default to one NAT gateway per AZ. Expose `single_nat_gateway` for cost-sensitive
non-production environments.

**Trade-off.** Higher fixed cost and more Elastic IPs. The default favors availability because
egress failure is user-visible and hard to diagnose.

**Revisit when.** The environment is ephemeral, or egress volume is low and the availability
requirement is explicitly lower.

---

## 6. Cost tool reads a snapshot, not the AWS API

**Context.** Cost tools that call AWS APIs need broad read permissions, credentials management, and
network access. That makes them hard to run in CI and hard to reason about.

**Decision.** `cost-doctor` consumes a normalized JSON snapshot and applies deterministic rules.
The tool has no AWS SDK dependency, no credentials, and no network calls.

**Trade-off.** Someone must produce the snapshot. The planned native collector will address that,
but the core remains snapshot-based so rules stay testable and auditable.

**Revisit when.** A hosted or credential-based mode becomes necessary, but keep the offline path as
the reference implementation for tests.

---

## 7. Conservative, transparent savings estimates

**Context.** Cost tools often multiply resources by list prices and present the result as guaranteed
savings. That erodes trust when finance checks the bill.

**Decision.** Only use `monthly_cost_usd` values present in the snapshot. When a rule can only
recover part of a resource (for example downsizing), apply a documented fraction and label the result
as directional.

**Trade-off.** Estimates understate some opportunities and require the exporter to provide costs.

**Revisit when.** The tool can join against a versioned price list with an explicit effective date
and region; even then, keep the assumptions visible in the output.

---

## 8. Monorepo with three entry points

**Context.** The playbook, the cost tool, and the EKS baseline are useful alone but tell a stronger
story together: design → cost baseline → deploy → operate.

**Decision.** Keep them in one repository with a shared README, shared license, and independent CI
workflows filtered by path.

**Trade-off.** A larger repository and three toolchains. Tags, releases, and ownership need to be
scoped per directory.

**Revisit when.** One module gains external contributors and release cadence that justify its own
repository. Splitting later is straightforward; keeping three half-empty repositories early is not.

---

## 9. Own Terraform modules instead of wrapping registry modules

**Context.** Registry modules such as `terraform-aws-modules/eks` are mature and well-maintained, but
they hide many defaults behind a large variable surface.

**Decision.** Keep the blueprint small and explicit: a network module and an EKS module with the
variables that matter for this baseline.

**Trade-off.** More code to maintain and fewer features than a registry module. In exchange, every
resource and default is visible and reviewable.

**Revisit when.** The baseline needs features that would take significant effort to replicate, or
the team standardizes on registry modules for governance reasons.

---

## 10. Intentional gaps

The blueprint does not install the AWS Load Balancer Controller, Karpenter, external-dns,
cert-manager, or an observability stack. Those choices depend on account structure, domain,
certificate management, and compliance posture. Shipping them half-configured would create a false
sense of completeness.

The Helm values under `infrastructure/eks-blueprint/helm/values/` are starting points, with the
IAM prerequisites documented rather than guessed.

---

## Using this in a design review

A useful review exercise is to walk each decision and ask:

1. What failure mode does this default prevent?
2. What does it make slower or more expensive?
3. What evidence would justify changing it?
4. Who owns the exception, and when does it expire?

If a team cannot answer those four questions, the default should stay in place.
