# 03 · Security & Compliance

Compliance requirements shape architecture. Treat this chapter as a set of questions to answer with
your legal and security teams, not as legal advice.

## Data residency and cross-border transfer

Before choosing a region, answer:

- Which data classes are in scope (personal data, payment data, health data, logs, backups)?
- Must any data stay in a specific jurisdiction?
- Is cross-border transfer allowed, and under which mechanism (contracts, consent, adequacy)?
- Do backups, snapshots, and logs inherit the same restrictions as primary data?
- Can support staff in another country access the data? Under what controls?

Document the answers as architecture constraints. Revisit when adding regions or vendors.

## Common regulatory frameworks

| Framework | Typical implication |
| --- | --- |
| GDPR / UK GDPR | Lawful basis, data subject rights, transfer mechanisms, breach notification |
| CCPA / CPRA | Consumer rights, opt-out, disclosure obligations |
| China PIPL | Separate consent and security assessment for certain cross-border transfers |
| SOC 2 | Controls for security, availability, confidentiality; customer audit expectations |
| ISO 27001 | Information security management system and evidence trail |
| PCI DSS | Cardholder data scope reduction, segmentation, logging |

AWS provides shared-responsibility documentation and compliance reports; your controls and evidence
remain your responsibility.

## Identity and access

- Human access through IAM Identity Center with MFA; no shared accounts.
- Workload access through IAM roles; no static access keys in code or images.
- Separate roles by function and environment; review permissions quarterly.
- Use permission boundaries or SCPs to contain delegated administration.
- Rotate and audit break-glass credentials.

## Encryption

- Encrypt data at rest with KMS customer-managed keys where control or auditability matters.
- Enforce TLS in transit; prefer modern TLS versions and strong cipher suites.
- Separate key policies by environment and data class.
- Enable automatic rotation for supported secrets and keys.
- Store application secrets in Secrets Manager or Parameter Store, never in images or repos.

## Network security

- Default-deny security groups; allow only required paths.
- Use WAF on public HTTP endpoints; consider Shield Advanced for high-risk workloads.
- Restrict management access through SSM Session Manager or a bastion, not open SSH.
- Segment production from non-production at the account and network level.
- Log and alert on unusual egress, DNS, and authentication patterns.

## Detection and response

- Enable GuardDuty, Security Hub, and CloudTrail organization-wide.
- Centralize findings and route them to an owned on-call channel.
- Define severity levels, response times, and escalation paths.
- Run at least one incident-response tabletop before launch.
- Keep forensic evidence retention aligned with legal requirements.

## Vendor and subprocessor review

- Maintain a list of subprocessors and their data locations.
- Review data processing agreements before onboarding new services.
- Check whether the service stores data in the chosen region or replicates it elsewhere.
- Include support-access terms in the review.

## Checklist

- [ ] Data classes and residency requirements documented
- [ ] Cross-border transfer mechanism approved by legal/compliance
- [ ] Region and service choices validated against residency requirements
- [ ] Backups, snapshots, and logs covered by the same controls
- [ ] IAM Identity Center + MFA enforced; no static workload keys
- [ ] KMS keys and rotation strategy defined
- [ ] Encryption in transit enforced end to end
- [ ] WAF / Shield decision recorded for public endpoints
- [ ] GuardDuty, Security Hub, CloudTrail centralized and alerting
- [ ] Incident-response plan and tabletop completed
- [ ] Subprocessor list and data locations reviewed
- [ ] Evidence retention and audit process defined

## Next

→ [04 · Cost Governance](./04-cost-governance.md)