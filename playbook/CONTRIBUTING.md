# Contributing

Thanks for helping improve the playbook.

## What we want

- Corrections to technical claims, with a link to official documentation where possible
- New checklists or decision records from real outbound projects
- Small, tested IaC or configuration templates
- Translations that keep the same structure as the English source

## What we avoid

- Vendor marketing or affiliate links
- Recommendations without a stated trade-off
- Unverifiable cost or latency numbers presented as universal facts
- Secrets, account IDs, ARNs, or customer names

## Pull request checklist

- [ ] The change is scoped to one topic
- [ ] Links resolve and headings follow the existing structure
- [ ] Checklists use `- [ ]` and are actionable
- [ ] Any AWS-specific claim links to official documentation
- [ ] No credentials, account IDs, or personal data

## Style

- English documents in `docs/`; keep `README.zh-CN.md` in sync for user-facing changes.
- Prefer short sentences and tables over long prose.
- Use "we recommend" for opinions and "must" only for hard requirements.
- Mark draft content clearly.

## Commit messages

Use a short imperative subject, for example:

```text
docs: add multi-region failover checklist
fix: correct VPC endpoint cost note
```

## License

By contributing, you agree that your contributions are licensed under [Apache-2.0](../LICENSE).
