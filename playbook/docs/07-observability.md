# 07 · Observability

Observability is part of the launch, not a follow-up project. The goal is to answer "is the user
experience healthy?" and "why not?" without SSH.

## Signals

| Signal | Minimum | Good |
| --- | --- | --- |
| Metrics | Request rate, errors, duration, saturation | Per-tenant and per-dependency breakdowns |
| Logs | Structured JSON with request/trace IDs | Correlated with traces and deployment version |
| Traces | Critical user journeys | End-to-end across services and queues |
| Events | Deployments, scaling, config changes | Annotated on dashboards |

## Collection architecture

- Use **OpenTelemetry** as the instrumentation standard; avoid vendor-specific SDKs in application code.
- Run ADOT Collector (or OTel Collector) as a DaemonSet/gateway to batch and route telemetry.
- Send metrics to Amazon Managed Prometheus or CloudWatch; logs to CloudWatch Logs or OpenSearch;
  traces to AWS X-Ray or a compatible backend.
- Set resource attributes consistently: `service.name`, `deployment.environment`, `service.version`,
  `cloud.region`, `k8s.namespace.name`.

## SLOs and alerting

- Define SLIs from the user's perspective: availability, latency, correctness, freshness.
- Set SLOs with error budgets; alert on **burn rate**, not raw thresholds.
- Route alerts by ownership; every alert needs a runbook and a clear action.
- Separate pages (user-impacting) from tickets (hygiene).
- Review alert noise monthly and delete alerts nobody acts on.

## Dashboards

Start with four dashboards:

1. **Service overview:** golden signals per service, error budget burn.
2. **Dependencies:** database, cache, queue, and third-party latency/error rates.
3. **Cluster:** node/pod capacity, scheduling failures, addon health.
4. **Business:** signups, orders, revenue-impacting flows (where applicable).

Annotate deployments and incidents on every dashboard.

## Logging cost control

- Sample high-volume debug logs; keep errors and audit logs complete.
- Set retention per log class; archive rather than retain everything hot.
- Avoid logging personal or secret data; redact at the collector.
- Use structured fields instead of free-text parsing.

## Runbooks

Every page should link to a runbook that states:

- what the alert means and its user impact
- first three diagnostic steps
- known mitigations and rollback
- escalation path and owner

## Checklist

- [ ] OTel instrumentation standard adopted
- [ ] Collector deployed with batching and retry
- [ ] Consistent resource attributes across services
- [ ] Metrics, logs, and traces correlated by trace ID
- [ ] SLOs and error budgets defined for user-facing services
- [ ] Burn-rate alerts with ownership and runbooks
- [ ] Four baseline dashboards live
- [ ] Deployment and incident annotations enabled
- [ ] Log retention and sampling policies set
- [ ] No personal/secret data in logs
- [ ] Alert noise reviewed monthly

## Next

→ [Launch Readiness Checklist](../checklists/launch-readiness.md)