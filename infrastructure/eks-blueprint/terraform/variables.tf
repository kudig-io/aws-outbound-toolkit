variable "region" {
  description = "AWS region for the cluster."
  type        = string
}

variable "cluster_name" {
  description = "EKS cluster name. Also used as the resource name prefix."
  type        = string
}

variable "kubernetes_version" {
  description = "EKS Kubernetes version. Must be a version currently in standard support."
  type        = string
}

variable "vpc_cidr" {
  description = "CIDR block for the VPC. Plan for pod density and future growth."
  type        = string
  default     = "10.40.0.0/16"
}

variable "availability_zones" {
  description = "Availability zones. Empty means the first three AZs in the region."
  type        = list(string)
  default     = []
}

variable "endpoint_public_access" {
  description = "Expose the EKS public API endpoint. Prefer false with a private CI/VPN path."
  type        = bool
  default     = false
}

variable "public_access_cidrs" {
  description = "CIDRs allowed to reach the public endpoint when enabled."
  type        = list(string)
  default     = []
}

variable "node_group_name" {
  description = "Managed node group name."
  type        = string
  default     = "general"
}

variable "node_instance_types" {
  description = "Instance types for the general node group. Graviton is the default."
  type        = list(string)
  default     = ["m7g.large"]
}

variable "node_capacity_type" {
  description = "ON_DEMAND or SPOT for the general node group."
  type        = string
  default     = "ON_DEMAND"

  validation {
    condition     = contains(["ON_DEMAND", "SPOT"], var.node_capacity_type)
    error_message = "node_capacity_type must be ON_DEMAND or SPOT."
  }
}

variable "node_desired_size" {
  description = "Desired node count."
  type        = number
  default     = 3
}

variable "node_min_size" {
  description = "Minimum node count. Keep at least one per AZ for production."
  type        = number
  default     = 3
}

variable "node_max_size" {
  description = "Maximum node count."
  type        = number
  default     = 9
}

variable "node_labels" {
  description = "Labels applied to the managed node group."
  type        = map(string)
  default     = {}
}

variable "tags" {
  description = "Tags applied to all resources."
  type        = map(string)
  default = {
    project = "eks-production-blueprint"
    env     = "prod"
  }
}