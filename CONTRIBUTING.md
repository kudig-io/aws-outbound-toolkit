# Contributing

Thanks for improving the toolkit. This repository has three modules with different toolchains;
please keep changes scoped to one module per pull request where possible.

## Repository conventions

- `playbook/` — Markdown chapters, checklists, and small templates
- `tools/cost-doctor/` — Go CLI; every rule needs a test
- `infrastructure/eks-blueprint/` — Terraform; run `terraform fmt` and `terraform validate`

## Before opening a PR

```bash
make test        # Go tests
make fmt-check   # Go formatting + Terraform formatting
make validate    # terraform init -backend=false + terraform validate
```

Markdown is checked in CI with markdownlint and lychee. Keep relative links valid.

## Guidelines

- Do not commit credentials, account IDs, ARNs, state files, or customer data.
- Keep examples generic; use `123456789012` and placeholder domains.
- Update both the English and Chinese README when user-facing behavior changes.
- Record architectural changes in `docs/design-decisions.md`.
- For Playbook changes, make recommendations actionable: a decision, a default, or a check.
- For cost-doctor changes, keep JSON output stable; add fields instead of renaming them.

## Commit messages

Short imperative subjects, for example:

```
docs: add multi-region failover checklist
feat(cost-doctor): add idle NAT gateway rule
fix(eks-blueprint): correct subnet tagging for internal ELB
```

## License

By contributing, you agree that your contributions are licensed under [Apache-2.0](./LICENSE).