region             = "ap-southeast-1"
cluster_name       = "outbound-prod"
kubernetes_version = "1.33" # replace with a version currently in EKS standard support
vpc_cidr           = "10.40.0.0/16"

# Leave empty to use the first three AZs in the region.
availability_zones = []

# Prefer private endpoint access with a VPN/CI runner inside the VPC.
endpoint_public_access = false
public_access_cidrs    = []

node_group_name     = "general"
node_instance_types = ["m7g.large"]
node_capacity_type  = "ON_DEMAND"
node_desired_size   = 3
node_min_size       = 3
node_max_size       = 9

tags = {
  project = "eks-production-blueprint"
  env     = "prod"
  owner   = "platform"
}