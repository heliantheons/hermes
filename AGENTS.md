# Hermes

This repository owns Helios identity persistence, provisioning, and relationship data.

## Boundaries

- Authentication flows and token issuance belong to Aegis.
- Hermes owns its language-neutral gRPC Schema under `proto/v1`.
- Reusable guards belong to `heliantheon/aegis-go`.
- Domain-independent infrastructure belongs to `heliantheon/common`.
- Kubernetes desired state belongs to the private `heliantheons/applications`
  repository. This public repository owns the image, not deployment manifests.

## Commands

```bash
make test
make lint
make build
make run
```

## Verification

- Run all Go tests after model, query, or gRPC changes.
- Keep SQL schema changes compatible with existing data or document the required migration.
- Keep generated seed data and deployment keys out of this repository.
- Run `make proto-lint`, regenerate the private server bindings, and publish a new `schema/*` tag when a contract changes.
