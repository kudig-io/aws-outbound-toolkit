# Contributing

Thanks for improving `aws-cost-doctor`.

## Development

```bash
make test     # go test ./...
make vet      # go vet ./...
make build    # build bin/aws-cost-doctor
```

CI runs `gofmt -l`, `go vet`, `go test`, and `go build`.

## Adding a rule

1. Add the rule to `internal/rules/rules.go` with a stable, kebab-case ID.
2. Choose a severity and document the threshold in the rule comment.
3. Compute savings only from fields present in the inventory; never hard-code AWS prices.
4. Add a positive test and a negative test in `internal/rules/rules_test.go`.
5. Update the rule table in both `README.md` and `README.zh-CN.md`.

## Changing the input schema

- Update `internal/model/model.go`, `docs/input-schema.md`, and `examples/inventory.json` together.
- Keep backward compatibility until v1.0; new fields must be optional.

## Output stability

Table and Markdown output are human-facing; JSON is the stable machine contract. Avoid renaming
existing JSON keys. Add fields instead.

## Commit messages

Short imperative subjects, for example:

```text
feat: add idle NAT gateway rule
fix: skip utilization rules when metrics are missing
docs: clarify savings estimate assumptions
```

## License

By contributing, you agree that your contributions are licensed under [Apache-2.0](../../LICENSE).
