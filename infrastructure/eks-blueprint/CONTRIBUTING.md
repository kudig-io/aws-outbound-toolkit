# Contributing

Thanks for improving the blueprint.

## Before opening a PR

```bash
make fmt-check
make validate
```

`make validate` runs `terraform init -backend=false` and `terraform validate`. It downloads providers,
so it needs network access.

## Guidelines

- Keep the root module minimal; put reusable logic in `terraform/modules/`.
- Do not hard-code account IDs, regions, AMI IDs, or ARNs.
- Pin provider and addon versions where compatibility matters.
- Every new variable needs a description and a sensible default where possible.
- Update both `README.md` and `README.zh-CN.md` when behavior changes.
- Add a note to `docs/architecture.md` for anything that changes the topology.
- Avoid installing controllers implicitly; document them under `helm/` instead.

## Testing changes

- Run `terraform plan` in a sandbox account and include a redacted summary in the PR.
- For module changes, test with at least one non-default input set.
- Do not commit `terraform.tfvars`, state files, or plan files.

## Commit messages

Short imperative subjects, for example:

```text
feat: add Karpenter IRSA module
fix: correct subnet tagging for internal ELB
docs: clarify private endpoint access
```

## License

By contributing, you agree that your contributions are licensed under [Apache-2.0](../../LICENSE).
