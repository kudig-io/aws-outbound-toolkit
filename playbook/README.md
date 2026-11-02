# AWS Outbound Playbook

> A practical, opinionated playbook for teams taking their workloads global on AWS.

[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](../LICENSE)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](./CONTRIBUTING.md)
[![Docs](https://img.shields.io/badge/docs-playbook-0f1689.svg)](./docs/README.md)

**English** | [简体中文](./README.zh-CN.md)

This repository collects the decisions, checklists, and reference snippets we wish we had when
helping teams move production workloads from China/APAC to AWS global regions. It is intentionally
**opinionated**: every chapter ends with a checklist you can copy into your own launch plan.

## Who this is for

- Platform / SRE / DevOps teams running Kubernetes or serverless workloads on AWS outside China
- Startups and SaaS teams going global (出海) that need a repeatable landing-zone baseline
- Teams migrating from Alibaba Cloud, Tencent Cloud, or on-premises IDC to AWS

## What's inside

| Chapter | Focus | Checklist |
| --- | --- | --- |
| [00 · Overview](./docs/00-overview.md) | The five stages of going global | — |
| [01 · Account & Landing Zone](./docs/01-account-and-landing-zone.md) | AWS Organizations, Control Tower, guardrails | ✅ |
| [02 · Network & Latency](./docs/02-network-and-latency.md) | Region choice, VPC, CloudFront, Global Accelerator | ✅ |
| [03 · Security & Compliance](./docs/03-security-and-compliance.md) | Data residency, GDPR/PIPL, IAM, encryption | ✅ |
| [04 · Cost Governance](./docs/04-cost-governance.md) | Budgets, tagging, Savings Plans, data transfer | ✅ |
| [05 · Migration from China Clouds](./docs/05-migration-from-china-clouds.md) | Service mapping, waves, cutover, rollback | ✅ |
| [06 · EKS Baseline](./docs/06-eks-baseline.md) | Production EKS defaults and day-2 operations | ✅ |
| [07 · Observability](./docs/07-observability.md) | OTel, CloudWatch, Prometheus, SLOs | ✅ |
| [Launch Readiness](./checklists/launch-readiness.md) | One-page go-live checklist | ✅ |

## Principles

1. **Landing zone before workloads.** Account structure, identity, logging, and guardrails are
   cheaper to fix on day 0 than on day 300.
2. **Network is a product feature.** Region and edge choices determine latency, cost, and
   resilience more than most application code does.
3. **Compliance by design.** Data residency and cross-border transfer requirements shape
   architecture; they are not paperwork to finish later.
4. **Cost is a budget line, not a surprise.** Tag, budget, and review from the first workload.
5. **Migrate in waves.** Move stateless first, then stateful; keep a tested rollback path.
6. **Observable from day one.** No production launch without SLOs, dashboards, and alert routing.

## Reference snippets

Small, copy-friendly IaC and configuration examples live under [`templates/`](./templates/).
They are starting points, not certified modules — review them before production use.

## Related projects

- [kudig-io/eks-production-blueprint](../infrastructure/eks-blueprint) — production EKS baseline
- [kudig-io/aws-cost-doctor](../tools/cost-doctor) — AWS cost and idle-resource scanner
- [kudig-io/kudig-database](https://github.com/kudig-io/kudig-database) — Kubernetes + AI infrastructure knowledge base

## Contributing

Corrections, war stories, and new checklists are welcome. See [CONTRIBUTING.md](./CONTRIBUTING.md).

## Disclaimer

This project is **not affiliated with, endorsed by, or sponsored by Amazon Web Services (AWS)**.
AWS and related marks are trademarks of Amazon.com, Inc. or its affiliates. All recommendations are
community-maintained and must be validated against your own security, legal, and commercial requirements.

## License

[Apache-2.0](../LICENSE)