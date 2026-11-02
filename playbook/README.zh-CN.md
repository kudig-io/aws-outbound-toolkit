# AWS 出海实战手册

> 面向出海团队的 AWS 落地实操手册：账号地基、多区域网络、合规、成本、迁移与 Day-2 运维。

[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](../LICENSE)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](./CONTRIBUTING.md)

[English](./README.md) | **简体中文**

本仓库整理了我们在帮助团队把生产负载从中国/亚太迁往 AWS 全球区域时，最希望一开始就有的
决策清单、检查表和参考片段。内容刻意保持**有观点**：每一章都以一份可直接抄进发布计划的
Checklist 结尾。

## 适合谁

- 在中国境外使用 AWS 运行 Kubernetes / Serverless 的平台、SRE、DevOps 团队
- 需要可复制的 Landing Zone 基线来支撑业务出海的创业公司与 SaaS 团队
- 正从阿里云、腾讯云或自建 IDC 迁移到 AWS 的团队

## 内容导航

| 章节 | 关注点 | Checklist |
| --- | --- | --- |
| [00 · 总览](./docs/00-overview.md) | 出海的五个阶段 | — |
| [01 · 账号与 Landing Zone](./docs/01-account-and-landing-zone.md) | Organizations、Control Tower、护栏 | ✅ |
| [02 · 网络与延迟](./docs/02-network-and-latency.md) | 区域选择、VPC、CloudFront、Global Accelerator | ✅ |
| [03 · 安全与合规](./docs/03-security-and-compliance.md) | 数据驻留、GDPR/PIPL、IAM、加密 | ✅ |
| [04 · 成本治理](./docs/04-cost-governance.md) | 预算、标签、Savings Plans、流量成本 | ✅ |
| [05 · 从中国云迁移](./docs/05-migration-from-china-clouds.md) | 服务映射、迁移批次、割接、回滚 | ✅ |
| [06 · EKS 生产基线](./docs/06-eks-baseline.md) | 生产级 EKS 默认配置与 Day-2 运维 | ✅ |
| [07 · 可观测性](./docs/07-observability.md) | OTel、CloudWatch、Prometheus、SLO | ✅ |
| [上线就绪检查](./checklists/launch-readiness.md) | 一页式 Go-Live 清单 | ✅ |

## 六条原则

1. **先地基，后负载。** 账号结构、身份、日志和护栏在 Day 0 修的代价远低于 Day 300。
2. **网络是产品能力。** 区域和边缘选择对延迟、成本、韧性的影响，通常大于大多数应用代码。
3. **合规要前置设计。** 数据驻留与跨境传输要求决定架构，而不是最后补的文书。
4. **成本是预算项，不是惊喜。** 从第一个负载开始打标签、设预算、做复盘。
5. **按批次迁移。** 先无状态、后有状态；始终保留经过验证的回滚路径。
6. **Day 1 就要可观测。** 没有 SLO、看板和告警路由，不上生产。

## 参考片段

可复制的 IaC 与配置示例位于 [`templates/`](./templates/)。它们是起点，不是经过认证的模块，
上生产前请自行评审。

## 相关项目

- [EKS 生产基线](../infrastructure/eks-blueprint) — 生产级 EKS 基线
- [Cost Doctor](../tools/cost-doctor) — AWS 成本与闲置资源扫描器
- [kudig-io/kudig-database](https://github.com/kudig-io/kudig-database) — Kubernetes + AI 基础设施知识库

## 参与贡献

欢迎勘误、实战经验和新的 Checklist，详见 [CONTRIBUTING.md](./CONTRIBUTING.md)。

## 声明

本项目**与 Amazon Web Services (AWS) 无任何隶属、背书或赞助关系**。AWS 及相关标识是
Amazon.com, Inc. 或其关联公司的商标。所有建议均由社区维护，落地前请结合自身安全、法律与
商务要求进行评估。

## 许可证

[Apache-2.0](../LICENSE)
