// Package model defines the inventory input consumed by aws-cost-doctor.
package model

// Inventory is a normalized snapshot of AWS resources. It is intentionally
// decoupled from the AWS SDK: any exporter (Cost Explorer, Resource Explorer,
// a custom script) can produce this shape.
type Inventory struct {
	AccountID     string         `json:"account_id"`
	Region        string         `json:"region"`
	Currency      string         `json:"currency,omitempty"`
	EC2Instances  []EC2Instance  `json:"ec2_instances,omitempty"`
	EBSVolumes    []EBSVolume    `json:"ebs_volumes,omitempty"`
	ElasticIPs    []ElasticIP    `json:"elastic_ips,omitempty"`
	LoadBalancers []LoadBalancer `json:"load_balancers,omitempty"`
	EKSNodeGroups []EKSNodeGroup `json:"eks_node_groups,omitempty"`
	SavingsPlans  []SavingsPlan  `json:"savings_plans,omitempty"`
}

// EC2Instance describes an EC2 instance and, optionally, observed utilization
// and its current monthly cost. Utilization is expressed as a percentage
// (0-100) over the observation window.
type EC2Instance struct {
	ID             string            `json:"id"`
	State          string            `json:"state"`
	InstanceType   string            `json:"instance_type,omitempty"`
	MonthlyCostUSD float64           `json:"monthly_cost_usd,omitempty"`
	CPUUtilization float64           `json:"cpu_utilization_pct,omitempty"`
	Tags           map[string]string `json:"tags,omitempty"`
}

// EBSVolume describes an EBS volume. State is typically "in-use" or
// "available" (the latter means unattached).
type EBSVolume struct {
	ID             string  `json:"id"`
	State          string  `json:"state"`
	SizeGB         int     `json:"size_gb,omitempty"`
	AttachedTo     string  `json:"attached_to,omitempty"`
	MonthlyCostUSD float64 `json:"monthly_cost_usd,omitempty"`
}

// ElasticIP describes an Elastic IP. Unassociated EIPs are commonly billable.
type ElasticIP struct {
	ID             string  `json:"id"`
	Associated     bool    `json:"associated"`
	MonthlyCostUSD float64 `json:"monthly_cost_usd,omitempty"`
}

// LoadBalancer describes an ELB/ALB/NLB with optional traffic observations.
type LoadBalancer struct {
	ID             string  `json:"id"`
	Type           string  `json:"type,omitempty"`
	HealthyTargets int     `json:"healthy_targets"`
	RequestsPerDay int64   `json:"requests_per_day,omitempty"`
	MonthlyCostUSD float64 `json:"monthly_cost_usd,omitempty"`
}

// EKSNodeGroup describes a managed node group and its observed utilization.
type EKSNodeGroup struct {
	Name           string  `json:"name"`
	DesiredSize    int     `json:"desired_size"`
	CPUUtilization float64 `json:"cpu_utilization_pct,omitempty"`
	MemUtilization float64 `json:"mem_utilization_pct,omitempty"`
	MonthlyCostUSD float64 `json:"monthly_cost_usd,omitempty"`
}

// SavingsPlan describes an active commitment. Utilization is a percentage.
type SavingsPlan struct {
	Name                 string  `json:"name,omitempty"`
	Type                 string  `json:"type,omitempty"`
	CommitmentUSDPerHour float64 `json:"commitment_usd_per_hour,omitempty"`
	Utilization          float64 `json:"utilization_pct,omitempty"`
	ExpiresAt            string  `json:"expires_at,omitempty"`
}
