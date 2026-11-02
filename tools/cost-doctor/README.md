# aws-cost-doctor

> Deterministic AWS cost checks from an inventory snapshot — no credentials, no network, no surprise.

[![CI](https://github.com/kudig-io/aws-outbound-toolkit/actions/workflows/cost-doctor.yml/badge.svg)](https://github.com/kudig-io/aws-outbound-toolkit/actions/workflows/cost-doctor.yml)
[![Go Version](https://img.shields.io/badge/Go-1.24+-00ADD8.svg)](https://go.dev)
[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](../../LICENSE)

**English** | [简体中文](./README.zh-CN.md)

Most cost tools want broad read permissions or a SaaS account. `aws-cost-doctor` takes a
**normalized JSON snapshot** of your AWS resources, applies transparent rules, and prints findings
as a table, Markdown, or JSON. It is deliberately boring and deterministic, which makes it easy to
run in CI, attach to a pull request, or use as a first-pass audit before buying a heavier tool.

## Status

Alpha. The rule set and input schema are intentionally small and may change until v0.1.0.

## Quick start

```bash
git clone https://github.com/kudig-io/aws-outbound-toolkit.git
cd aws-outbound-toolkit/tools/cost-doctor
go run ./cmd/aws-cost-doctor scan --input examples/inventory.json
```

Example output:

```text
Account: 123456789012   Region: ap-southeast-1   Findings: 7   Estimated savings: $379.16/month

SEVERITY  RULE                          RESOURCE                          EST. SAVINGS  MESSAGE
high      unattached-ebs                vol-0exampleunattached1           $40.00        EBS volume is unattached (500 GiB).
high      idle-load-balancer            app/example-api/...               $22.00        Load balancer has no healthy targets and no recent requests.
medium    underutilized-eks-nodegroup   example-general                   $155.00       Node group utilization is low (CPU 9.5%, memory 21.0%).
```

## Install

```bash
go install github.com/kudig-io/aws-outbound-toolkit/tools/cost-doctor/cmd/aws-cost-doctor@latest
```

Or build a binary:

```bash
make build          # writes bin/aws-cost-doctor
```

## Usage

```bash
aws-cost-doctor scan --input inventory.json
aws-cost-doctor scan --input inventory.json --format markdown > cost-report.md
aws-cost-doctor scan --input inventory.json --format json | jq '.summary'
aws-cost-doctor scan --input inventory.json --fail-on high   # CI gate
```

| Flag | Default | Description |
| --- | --- | --- |
| `--input` | required | Path to the inventory JSON file |
| `--format` | `table` | `table`, `markdown`, or `json` |
| `--fail-on` | off | Exit `1` when a finding at or above `low`, `medium`, `high`, or `critical` exists |

### Exit codes

| Code | Meaning |
| --- | --- |
| `0` | Scan completed; no `--fail-on` threshold hit |
| `1` | A finding at or above the `--fail-on` threshold was found |
| `2` | Usage error, unreadable input, or invalid JSON |

## Rules

| Rule | Severity | What it detects |
| --- | --- | --- |
| `unattached-ebs` | high | EBS volumes in `available` state |
| `idle-load-balancer` | high | Load balancer with no healthy targets and no recent requests |
| `low-savings-plan-utilization` | high | Savings Plan utilization below 80% |
| `underutilized-eks-nodegroup` | medium | Node group with CPU < 15% and memory < 30% |
| `stopped-ec2-cost` | medium | Stopped instance still reporting monthly cost |
| `unattached-eip` | medium | Elastic IP not associated with a resource |
| `underutilized-ec2` | low | Running instance with average CPU < 5% |

Savings estimates use the `monthly_cost_usd` field from the inventory. Where a rule only knows a
portion is recoverable (for example downsizing), it applies a clearly documented fraction rather
than guessing at prices. **Treat all estimates as directional, not invoices.**

## Input schema

See [docs/input-schema.md](./docs/input-schema.md) and
[examples/inventory.json](./examples/inventory.json).

The tool does not call AWS APIs. Produce the snapshot with Cost Explorer, Resource Explorer, a
custom exporter, or the planned native collector (see roadmap).

## Roadmap

- [ ] Native collector using AWS SDK v2 (Cost Explorer + Resource Explorer + EKS)
- [ ] Terraform/OpenTofu cost checks on plan output
- [ ] Additional rules: orphaned snapshots, idle NAT gateways, old AMIs, untagged resources
- [ ] HTML report and GitHub Actions summary
- [ ] Configurable thresholds and rule allowlists
- [ ] AWS Lambda / SAR deployment

## Contributing

Issues and pull requests are welcome. Please add a test for every new rule and keep output stable
for existing formats. See [CONTRIBUTING.md](./CONTRIBUTING.md).

## Related projects

- [kudig-io/aws-outbound-playbook](../../playbook) — AWS 出海实战手册
- [kudig-io/eks-production-blueprint](../../infrastructure/eks-blueprint) — production EKS baseline

## Disclaimer

This project is **not affiliated with, endorsed by, or sponsored by Amazon Web Services (AWS)**.
AWS and related marks are trademarks of Amazon.com, Inc. or its affiliates.

## License

[Apache-2.0](../../LICENSE)