# Three-AZ VPC template

A minimal starting point: one VPC, public and private subnets across three AZs, one NAT gateway
per AZ, and flow logs.

## What this creates

- VPC with DNS hostnames enabled
- Three public subnets and three private subnets
- Internet gateway and NAT gateways
- Route tables and associations
- VPC Flow Logs to CloudWatch Logs

## What this does not create

- Transit Gateway, peering, or shared-VPC attachments
- VPC endpoints
- Any workload resources

## Cost note

NAT gateways and flow logs are billable. For non-production use, consider a single NAT gateway or
none at all.

## Usage

```hcl
module "vpc" {
  source = "./templates/terraform/vpc"

  name               = "outbound-prod"
  cidr_block         = "10.40.0.0/16"
  availability_zones = ["ap-southeast-1a", "ap-southeast-1b", "ap-southeast-1c"]
  tags = {
    env     = "prod"
    owner   = "platform"
    service = "network"
  }
}
```

Review CIDR planning and subnet sizing against your pod density before applying.
