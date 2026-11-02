variable "name" {
  description = "Name prefix for all resources."
  type        = string
}

variable "cidr_block" {
  description = "VPC CIDR block. Plan for growth and avoid overlap with other VPCs."
  type        = string
}

variable "availability_zones" {
  description = "Availability zones to use. Use at least three for production."
  type        = list(string)

  validation {
    condition     = length(var.availability_zones) >= 2 && length(var.availability_zones) <= 6
    error_message = "Use between 2 and 6 availability zones."
  }
}

variable "single_nat_gateway" {
  description = "Use one NAT gateway instead of one per AZ. Cheaper, but reduces AZ resilience."
  type        = bool
  default     = false
}

variable "enable_flow_logs" {
  description = "Enable VPC Flow Logs to CloudWatch Logs."
  type        = bool
  default     = true
}

variable "flow_log_retention_days" {
  description = "Retention for VPC Flow Logs."
  type        = number
  default     = 30
}

variable "tags" {
  description = "Tags applied to all resources."
  type        = map(string)
  default     = {}
}