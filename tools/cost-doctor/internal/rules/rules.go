// Package rules contains deterministic checks over an inventory snapshot.
package rules

import (
	"fmt"
	"sort"

	"github.com/kudig-io/aws-outbound-toolkit/tools/cost-doctor/internal/model"
)

// Severity levels, ordered from least to most urgent.
const (
	SeverityLow      = "low"
	SeverityMedium   = "medium"
	SeverityHigh     = "high"
	SeverityCritical = "critical"
)

// Finding is a single actionable cost issue.
type Finding struct {
	Rule                       string  `json:"rule"`
	Severity                   string  `json:"severity"`
	ResourceID                 string  `json:"resource_id"`
	ResourceType               string  `json:"resource_type"`
	Message                    string  `json:"message"`
	Recommendation             string  `json:"recommendation"`
	EstimatedMonthlySavingsUSD float64 `json:"estimated_monthly_savings_usd,omitempty"`
}

// Scan runs every rule against the inventory and returns findings sorted by
// severity (most urgent first), then by estimated savings (largest first).
func Scan(inv model.Inventory) []Finding {
	var findings []Finding

	for _, v := range inv.EBSVolumes {
		if v.State == "available" {
			findings = append(findings, Finding{
				Rule:                       "unattached-ebs",
				Severity:                   SeverityHigh,
				ResourceID:                 v.ID,
				ResourceType:               "ebs-volume",
				Message:                    fmt.Sprintf("EBS volume is unattached (%d GiB).", v.SizeGB),
				Recommendation:             "Snapshot if retention is required, then delete the volume.",
				EstimatedMonthlySavingsUSD: v.MonthlyCostUSD,
			})
		}
	}

	for _, e := range inv.ElasticIPs {
		if !e.Associated {
			findings = append(findings, Finding{
				Rule:                       "unattached-eip",
				Severity:                   SeverityMedium,
				ResourceID:                 e.ID,
				ResourceType:               "elastic-ip",
				Message:                    "Elastic IP is not associated with a resource.",
				Recommendation:             "Release the address or attach it to a load balancer/NAT gateway.",
				EstimatedMonthlySavingsUSD: e.MonthlyCostUSD,
			})
		}
	}

	for _, lb := range inv.LoadBalancers {
		if lb.HealthyTargets == 0 && lb.RequestsPerDay == 0 {
			findings = append(findings, Finding{
				Rule:                       "idle-load-balancer",
				Severity:                   SeverityHigh,
				ResourceID:                 lb.ID,
				ResourceType:               "load-balancer",
				Message:                    "Load balancer has no healthy targets and no recent requests.",
				Recommendation:             "Delete the load balancer or restore its target group.",
				EstimatedMonthlySavingsUSD: lb.MonthlyCostUSD,
			})
		}
	}

	for _, i := range inv.EC2Instances {
		if i.State == "stopped" && i.MonthlyCostUSD > 0 {
			findings = append(findings, Finding{
				Rule:                       "stopped-ec2-cost",
				Severity:                   SeverityMedium,
				ResourceID:                 i.ID,
				ResourceType:               "ec2-instance",
				Message:                    "Stopped instance still reports monthly cost (likely attached storage).",
				Recommendation:             "Terminate the instance or right-size and detach unused volumes.",
				EstimatedMonthlySavingsUSD: i.MonthlyCostUSD,
			})
		}
		if i.State == "running" && i.CPUUtilization > 0 && i.CPUUtilization < 5 && i.MonthlyCostUSD > 0 {
			findings = append(findings, Finding{
				Rule:                       "underutilized-ec2",
				Severity:                   SeverityLow,
				ResourceID:                 i.ID,
				ResourceType:               "ec2-instance",
				Message:                    fmt.Sprintf("Average CPU utilization is %.1f%%.", i.CPUUtilization),
				Recommendation:             "Downsize, consolidate, or move the workload to containers/serverless.",
				EstimatedMonthlySavingsUSD: i.MonthlyCostUSD * 0.5,
			})
		}
	}

	for _, ng := range inv.EKSNodeGroups {
		if ng.DesiredSize > 0 && ng.CPUUtilization > 0 && ng.CPUUtilization < 15 && ng.MemUtilization < 30 {
			findings = append(findings, Finding{
				Rule:                       "underutilized-eks-nodegroup",
				Severity:                   SeverityMedium,
				ResourceID:                 ng.Name,
				ResourceType:               "eks-nodegroup",
				Message:                    fmt.Sprintf("Node group utilization is low (CPU %.1f%%, memory %.1f%%).", ng.CPUUtilization, ng.MemUtilization),
				Recommendation:             "Enable consolidation (Karpenter/cluster autoscaler), reduce desired size, or use smaller instances.",
				EstimatedMonthlySavingsUSD: ng.MonthlyCostUSD * 0.25,
			})
		}
	}

	for _, sp := range inv.SavingsPlans {
		if sp.CommitmentUSDPerHour > 0 && sp.Utilization > 0 && sp.Utilization < 80 {
			findings = append(findings, Finding{
				Rule:           "low-savings-plan-utilization",
				Severity:       SeverityHigh,
				ResourceID:     savingsPlanID(sp),
				ResourceType:   "savings-plan",
				Message:        fmt.Sprintf("Savings Plan utilization is %.1f%%; unused commitment is still billed.", sp.Utilization),
				Recommendation: "Shift eligible workloads onto the commitment or adjust future commitments at renewal.",
			})
		}
	}

	sortFindings(findings)
	return findings
}

func savingsPlanID(sp model.SavingsPlan) string {
	if sp.Name != "" {
		return sp.Name
	}
	if sp.Type != "" {
		return sp.Type
	}
	return "savings-plan"
}

func severityRank(s string) int {
	switch s {
	case SeverityCritical:
		return 4
	case SeverityHigh:
		return 3
	case SeverityMedium:
		return 2
	case SeverityLow:
		return 1
	default:
		return 0
	}
}

func sortFindings(findings []Finding) {
	sort.SliceStable(findings, func(i, j int) bool {
		ri, rj := severityRank(findings[i].Severity), severityRank(findings[j].Severity)
		if ri != rj {
			return ri > rj
		}
		return findings[i].EstimatedMonthlySavingsUSD > findings[j].EstimatedMonthlySavingsUSD
	})
}

// TotalEstimatedSavings sums the estimated monthly savings across findings.
func TotalEstimatedSavings(findings []Finding) float64 {
	var total float64
	for _, f := range findings {
		total += f.EstimatedMonthlySavingsUSD
	}
	return total
}

// CountBySeverity returns the number of findings per severity.
func CountBySeverity(findings []Finding) map[string]int {
	counts := map[string]int{}
	for _, f := range findings {
		counts[f.Severity]++
	}
	return counts
}
