# aws-cost-doctor

> 基于资源快照的确定性 AWS 成本检查——不需要凭证、不联网、没有意外。

[![CI](https://github.com/kudig-io/aws-outbound-toolkit/actions/workflows/cost-doctor.yml/badge.svg)](https://github.com/kudig-io/aws-outbound-toolkit/actions/workflows/cost-doctor.yml)
[![Go Version](https://img.shields.io/badge/Go-1.24+-00ADD8.svg)](https://go.dev)
[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](../../LICENSE)

[English](./README.md) | **简体中文**

大多数成本工具要求较宽的只读权限，或者需要接入 SaaS。`aws-cost-doctor` 读取一份
**归一化的资源 JSON 快照**，套用透明规则，输出表格、Markdown 或 JSON。它刻意保持简单和确定性，
因此很适合放进 CI、贴到 Pull Request，或作为引入更重工具前的第一轮体检。

## 状态

Alpha。规则集和输入 schema 在 v0.1.0 前可能变化。

## 快速开始

```bash
git clone https://github.com/kudig-io/aws-outbound-toolkit.git
cd aws-outbound-toolkit/tools/cost-doctor
go run ./cmd/aws-cost-doctor scan --input examples/inventory.json
```

输出示例：

```text
Account: 123456789012   Region: ap-southeast-1   Findings: 7   Estimated savings: $379.16/month

SEVERITY  RULE                          RESOURCE                          EST. SAVINGS  MESSAGE
high      unattached-ebs                vol-0exampleunattached1           $40.00        EBS volume is unattached (500 GiB).
high      idle-load-balancer            app/example-api/...               $22.00        Load balancer has no healthy targets and no recent requests.
medium    underutilized-eks-nodegroup   example-general                   $155.00       Node group utilization is low (CPU 9.5%, memory 21.0%).
```

## 安装

```bash
go install github.com/kudig-io/aws-outbound-toolkit/tools/cost-doctor/cmd/aws-cost-doctor@latest
```

或本地构建：

```bash
make build          # 生成 bin/aws-cost-doctor
```

## 用法

```bash
aws-cost-doctor scan --input inventory.json
aws-cost-doctor scan --input inventory.json --format markdown > cost-report.md
aws-cost-doctor scan --input inventory.json --format json | jq '.summary'
aws-cost-doctor scan --input inventory.json --fail-on high   # CI 门禁
```

| 参数 | 默认值 | 说明 |
| --- | --- | --- |
| `--input` | 必填 | 资源快照 JSON 路径 |
| `--format` | `table` | `table`、`markdown` 或 `json` |
| `--fail-on` | 关闭 | 存在 `low`/`medium`/`high`/`critical` 及以上问题时退出码为 `1` |

### 退出码

| 退出码 | 含义 |
| --- | --- |
| `0` | 扫描完成，未触发 `--fail-on` 阈值 |
| `1` | 存在达到 `--fail-on` 阈值的问题 |
| `2` | 用法错误、文件不可读或 JSON 非法 |

## 规则

| 规则 | 级别 | 检测内容 |
| --- | --- | --- |
| `unattached-ebs` | high | 处于 `available` 状态的 EBS 卷 |
| `idle-load-balancer` | high | 没有健康后端且近期无请求的负载均衡器 |
| `low-savings-plan-utilization` | high | Savings Plan 利用率低于 80% |
| `underutilized-eks-nodegroup` | medium | CPU < 15% 且内存 < 30% 的节点组 |
| `stopped-ec2-cost` | medium | 已停止但仍产生月成本的实例 |
| `unattached-eip` | medium | 未关联资源的弹性 IP |
| `underutilized-ec2` | low | 平均 CPU < 5% 的运行中实例 |

节省估算使用快照中的 `monthly_cost_usd` 字段。对于只能判断“部分可回收”的规则（例如降配），
会使用明确写出的比例，而不是猜测价格。**所有估算仅用于判断方向，不是账单。**

## 输入 schema

见 [docs/input-schema.md](./docs/input-schema.md) 与
[examples/inventory.json](./examples/inventory.json)。

本工具不调用 AWS API。快照可以来自 Cost Explorer、Resource Explorer、自定义导出脚本，
或规划中的原生采集器（见 Roadmap）。

## Roadmap

- [ ] 基于 AWS SDK v2 的原生采集器（Cost Explorer + Resource Explorer + EKS）
- [ ] 对 Terraform/OpenTofu plan 输出做成本检查
- [ ] 更多规则：孤立快照、空闲 NAT 网关、过期 AMI、未打标签资源
- [ ] HTML 报告与 GitHub Actions Summary
- [ ] 阈值与规则白名单可配置
- [ ] AWS Lambda / SAR 部署

## 参与贡献

欢迎 Issue 和 PR。新增规则请附带测试，并保持既有输出格式稳定。详见
[CONTRIBUTING.md](./CONTRIBUTING.md)。

## 相关项目

- [kudig-io/aws-outbound-playbook](../../playbook) — AWS 出海实战手册
- [kudig-io/eks-production-blueprint](../../infrastructure/eks-blueprint) — 生产级 EKS 基线

## 声明

本项目**与 Amazon Web Services (AWS) 无任何隶属、背书或赞助关系**。AWS 及相关标识是
Amazon.com, Inc. 或其关联公司的商标。

## 许可证

[Apache-2.0](../../LICENSE)