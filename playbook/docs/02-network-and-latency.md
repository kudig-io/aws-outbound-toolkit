# 02 · Network & Latency

For outbound teams, network design decides two things users feel immediately: **latency** and
**reliability**. It also drives a large share of the bill.

## Region selection

Score candidate regions against four axes:

| Axis | Questions |
| --- | --- |
| User latency | Where are the primary users? What is the p95 budget? |
| Service availability | Are the required services and instance families available? |
| Compliance | Does data residency or cross-border transfer restrict the choice? |
| Cost | Are data transfer, NAT, and compute prices acceptable? |

Practical approach:

1. Shortlist 2–3 regions.
2. Measure real latency from representative user locations, not from your office.
3. Validate service quotas before committing.
4. Choose one **primary** region and one **disaster-recovery** region, not three half-built ones.

## VPC design

- Use non-overlapping CIDRs across accounts and regions; plan for growth and future peering.
- Prefer a **shared network account** with Transit Gateway when you have more than a handful of VPCs.
- Keep public subnets small and only for load balancers / NAT; put workloads in private subnets.
- Spread production subnets across at least three AZs.
- Use VPC endpoints (Gateway for S3/DynamoDB, Interface for other services) to avoid NAT charges
  and keep traffic private.
- Enable VPC Flow Logs and DNS logging for production VPCs.

## Edge and acceleration

| Need | Option | Notes |
| --- | --- | --- |
| Static/content delivery | CloudFront | Cache at edge, reduce origin load and egress |
| Global TCP/UDP acceleration | Global Accelerator | Anycast IPs, useful for APIs and gaming |
| Regional HTTP API | ALB | Standard choice for EKS/ECS ingress |
| Private connectivity to China | Direct Connect / Site-to-Site VPN | Validate legal and compliance constraints first |

Rules of thumb:

- Put CloudFront in front of cacheable content and static assets.
- Use Global Accelerator when users are far from the region and TCP handshake latency dominates.
- Avoid routing all traffic through a single region "for simplicity" if users are spread globally.

## DNS and service discovery

- Host public zones in Route 53; use health checks and failover routing where needed.
- Keep internal zones private; use Route 53 Resolver endpoints for hybrid DNS.
- Standardize naming so environments and regions are obvious from the hostname.

## Cross-border connectivity

Cross-border links are a legal, compliance, and operational topic, not only a networking one:

- Confirm whether data can leave the origin jurisdiction and under which conditions.
- Prefer encrypted tunnels (IPsec VPN or Direct Connect with MACsec where available).
- Monitor latency, packet loss, and jitter continuously; cross-border paths are less predictable.
- Design for degraded connectivity: queue, retry, and batch rather than synchronous dependency.

## Latency validation

- Measure from real user networks and mobile carriers, not only from cloud regions.
- Track p50/p95/p99, not averages.
- Test at peak hours across at least a week before launch.
- Include DNS resolution and TLS handshake time in the budget.

## Checklist

- [ ] Primary and DR regions selected with latency and compliance evidence
- [ ] VPC CIDR plan documented and non-overlapping
- [ ] Production subnets span at least three AZs
- [ ] Workloads in private subnets; NAT only where required
- [ ] VPC endpoints configured for S3, DynamoDB, and high-traffic services
- [ ] Transit Gateway or peering design documented
- [ ] CloudFront / Global Accelerator decision recorded
- [ ] Route 53 zones, health checks, and failover tested
- [ ] Cross-border connectivity reviewed by legal/compliance
- [ ] Latency measured from representative user locations
- [ ] VPC Flow Logs and DNS logging enabled

## Next

→ [03 · Security & Compliance](./03-security-and-compliance.md)