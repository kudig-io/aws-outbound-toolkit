variable "cluster_name" {
  description = "EKS cluster name."
  type        = string
}

variable "kubernetes_version" {
  description = "EKS Kubernetes version currently in standard support."
  type        = string
}

variable "private_subnet_ids" {
  description = "Private subnet IDs for the control plane ENIs and nodes."
  type        = list(string)

  validation {
    condition     = length(var.private_subnet_ids) >= 2
    error_message = "Provide at least two private subnets."
  }
}

variable "endpoint_public_access" {
  description = "Expose the public EKS API endpoint."
  type        = bool
  default     = false
}

variable "public_access_cidrs" {
  description = "CIDRs allowed to reach the public endpoint. Empty means 0.0.0.0/0 when enabled."
  type        = list(string)
  default     = []
}

variable "node_group_name" {
  description = "Managed node group name."
  type        = string
  default     = "general"
}

variable "node_instance_types" {
  description = "Instance types for the node group."
  type        = list(string)
  default     = ["m7g.large"]
}

variable "node_capacity_type" {
  description = "ON_DEMAND or SPOT."
  type        = string
  default     = "ON_DEMAND"
}

variable "node_desired_size" {
  description = "Desired node count."
  type        = number
  default     = 3
}

variable "node_min_size" {
  description = "Minimum node count."
  type        = number
  default     = 3
}

variable "node_max_size" {
  description = "Maximum node count."
  type        = number
  default     = 9
}

variable "node_labels" {
  description = "Labels applied to the node group."
  type        = map(string)
  default     = {}
}

variable "cluster_addons" {
  description = "Managed addons to install. Set version to null to use the default version."
  type = map(object({
    version = optional(string)
  }))
  default = {
    vpc-cni    = {}
    coredns    = {}
    kube-proxy = {}
  }
}

variable "tags" {
  description = "Tags applied to all resources."
  type        = map(string)
  default     = {}
}