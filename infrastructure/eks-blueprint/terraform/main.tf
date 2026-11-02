data "aws_availability_zones" "available" {
  state = "available"
}

locals {
  azs = length(var.availability_zones) > 0 ? var.availability_zones : slice(data.aws_availability_zones.available.names, 0, 3)
}

module "network" {
  source = "./modules/network"

  name               = var.cluster_name
  cidr_block         = var.vpc_cidr
  availability_zones = local.azs
  tags               = var.tags
}

module "eks" {
  source = "./modules/eks"

  cluster_name           = var.cluster_name
  kubernetes_version     = var.kubernetes_version
  private_subnet_ids     = module.network.private_subnet_ids
  endpoint_public_access = var.endpoint_public_access
  public_access_cidrs    = var.public_access_cidrs

  node_group_name     = var.node_group_name
  node_instance_types = var.node_instance_types
  node_capacity_type  = var.node_capacity_type
  node_desired_size   = var.node_desired_size
  node_min_size       = var.node_min_size
  node_max_size       = var.node_max_size
  node_labels         = var.node_labels

  tags = var.tags
}