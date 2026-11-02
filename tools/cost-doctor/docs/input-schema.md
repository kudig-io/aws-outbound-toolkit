# Input schema

`aws-cost-doctor` consumes a single JSON object. All fields are optional except `id` for resources
you want checked. Unknown fields are ignored, so exporters can add metadata.

## Top level

| Field | Type | Description |
| --- | --- | --- |
| `account_id` | string | AWS account ID, printed in the report |
| `region` | string | Region name, printed in the report |
| `currency` | string | Optional currency label; estimates are always USD |
| `ec2_instances` | array | See below |
| `ebs_volumes` | array | See below |
| `elastic_ips` | array | See below |
| `load_balancers` | array | See below |
| `eks_node_groups` | array | See below |
| `savings_plans` | array | See below |

## `ec2_instances[]`

| Field | Type | Description |
| --- | --- | --- |
| `id` | string | Instance ID |
| `state` | string | `running` or `stopped` |
| `instance_type` | string | Optional, for display |
| `monthly_cost_usd` | number | Current monthly cost, including attached storage |
| `cpu_utilization_pct` | number | Average CPU over the observation window, 0–100 |
| `tags` | object | Optional key/value tags |

## `ebs_volumes[]`

| Field | Type | Description |
| --- | --- | --- |
| `id` | string | Volume ID |
| `state` | string | `in-use` or `available` |
| `size_gb` | number | Volume size |
| `attached_to` | string | Optional instance ID |
| `monthly_cost_usd` | number | Current monthly cost |

## `elastic_ips[]`

| Field | Type | Description |
| --- | --- | --- |
| `id` | string | Allocation ID |
| `associated` | boolean | Whether the address is associated |
| `monthly_cost_usd` | number | Current monthly cost |

## `load_balancers[]`

| Field | Type | Description |
| --- | --- | --- |
| `id` | string | Load balancer ARN or name |
| `type` | string | `application`, `network`, or `gateway` |
| `healthy_targets` | number | Number of healthy targets |
| `requests_per_day` | number | Recent request count; use `0` when unknown or idle |
| `monthly_cost_usd` | number | Current monthly cost |

## `eks_node_groups[]`

| Field | Type | Description |
| --- | --- | --- |
| `name` | string | Node group name |
| `desired_size` | number | Desired node count |
| `cpu_utilization_pct` | number | Average CPU, 0–100 |
| `mem_utilization_pct` | number | Average memory, 0–100 |
| `monthly_cost_usd` | number | Current monthly cost |

## `savings_plans[]`

| Field | Type | Description |
| --- | --- | --- |
| `name` | string | Optional plan name |
| `type` | string | `ComputeSavingsPlan` or `EC2InstanceSavingsPlan` |
| `commitment_usd_per_hour` | number | Hourly commitment |
| `utilization_pct` | number | Utilization, 0–100 |
| `expires_at` | string | Optional ISO date |

## Minimal example

```json
{
  "account_id": "123456789012",
  "region": "ap-southeast-1",
  "ebs_volumes": [
    { "id": "vol-123", "state": "available", "size_gb": 100, "monthly_cost_usd": 8 }
  ]
}
```

## Notes

- Cost fields are optional. Rules still detect waste, but savings estimates will be `0`.
- Utilization fields are optional. Rules that depend on them are skipped when the value is missing.
- The schema is versioned by Git tags; breaking changes bump the minor version until v1.0.