package rules

import (
	"testing"

	"github.com/kudig-io/aws-outbound-toolkit/tools/cost-doctor/internal/model"
)

func testInventory() model.Inventory {
	return model.Inventory{
		AccountID: "123456789012",
		Region:    "ap-southeast-1",
		EC2Instances: []model.EC2Instance{
			{ID: "i-idle", State: "running", MonthlyCostUSD: 100, CPUUtilization: 2},
			{ID: "i-stopped", State: "stopped", MonthlyCostUSD: 10},
			{ID: "i-busy", State: "running", MonthlyCostUSD: 100, CPUUtilization: 70},
		},
		EBSVolumes: []model.EBSVolume{
			{ID: "vol-free", State: "available", SizeGB: 100, MonthlyCostUSD: 8},
			{ID: "vol-inuse", State: "in-use", SizeGB: 100, MonthlyCostUSD: 8},
		},
		ElasticIPs: []model.ElasticIP{
			{ID: "eip-free", Associated: false, MonthlyCostUSD: 3.6},
		},
		LoadBalancers: []model.LoadBalancer{
			{ID: "lb-idle", HealthyTargets: 0, RequestsPerDay: 0, MonthlyCostUSD: 20},
			{ID: "lb-live", HealthyTargets: 2, RequestsPerDay: 1000, MonthlyCostUSD: 20},
		},
		EKSNodeGroups: []model.EKSNodeGroup{
			{Name: "ng-idle", DesiredSize: 3, CPUUtilization: 5, MemUtilization: 10, MonthlyCostUSD: 200},
		},
		SavingsPlans: []model.SavingsPlan{
			{Name: "sp-underused", CommitmentUSDPerHour: 1, Utilization: 50},
		},
	}
}

func TestScanFindsExpectedRules(t *testing.T) {
	findings := Scan(testInventory())

	want := []string{
		"unattached-ebs",
		"unattached-eip",
		"idle-load-balancer",
		"stopped-ec2-cost",
		"underutilized-ec2",
		"underutilized-eks-nodegroup",
		"low-savings-plan-utilization",
	}

	got := map[string]bool{}
	for _, f := range findings {
		got[f.Rule] = true
	}
	for _, rule := range want {
		if !got[rule] {
			t.Errorf("expected rule %q in findings, got %v", rule, got)
		}
	}

	if len(findings) != len(want) {
		t.Fatalf("expected %d findings, got %d", len(want), len(findings))
	}
}

func TestScanIgnoresHealthyResources(t *testing.T) {
	inv := model.Inventory{
		EC2Instances: []model.EC2Instance{
			{ID: "i-busy", State: "running", MonthlyCostUSD: 100, CPUUtilization: 70},
		},
		EBSVolumes: []model.EBSVolume{
			{ID: "vol-inuse", State: "in-use", MonthlyCostUSD: 8},
		},
		ElasticIPs: []model.ElasticIP{
			{ID: "eip-used", Associated: true, MonthlyCostUSD: 0},
		},
		LoadBalancers: []model.LoadBalancer{
			{ID: "lb-live", HealthyTargets: 2, RequestsPerDay: 1000},
		},
		SavingsPlans: []model.SavingsPlan{
			{Name: "sp-healthy", CommitmentUSDPerHour: 1, Utilization: 95},
		},
	}

	if findings := Scan(inv); len(findings) != 0 {
		t.Fatalf("expected no findings, got %+v", findings)
	}
}

func TestScanSortsBySeverityThenSavings(t *testing.T) {
	inv := model.Inventory{
		EBSVolumes: []model.EBSVolume{
			{ID: "vol-small", State: "available", MonthlyCostUSD: 1},
			{ID: "vol-large", State: "available", MonthlyCostUSD: 50},
		},
		ElasticIPs: []model.ElasticIP{
			{ID: "eip", Associated: false, MonthlyCostUSD: 100},
		},
	}

	findings := Scan(inv)
	if len(findings) != 3 {
		t.Fatalf("expected 3 findings, got %d", len(findings))
	}
	if findings[0].Rule != "unattached-ebs" || findings[0].ResourceID != "vol-large" {
		t.Fatalf("expected high-severity volume first, got %+v", findings[0])
	}
	if findings[2].Rule != "unattached-eip" {
		t.Fatalf("expected medium-severity EIP last, got %+v", findings[2])
	}
}

func TestTotalEstimatedSavings(t *testing.T) {
	findings := []Finding{
		{EstimatedMonthlySavingsUSD: 10.5},
		{EstimatedMonthlySavingsUSD: 2.25},
	}
	if got := TotalEstimatedSavings(findings); got != 12.75 {
		t.Fatalf("expected 12.75, got %v", got)
	}
}

func TestCountBySeverity(t *testing.T) {
	findings := Scan(testInventory())
	counts := CountBySeverity(findings)
	if counts[SeverityHigh] == 0 {
		t.Fatalf("expected at least one high severity finding, got %v", counts)
	}
}
