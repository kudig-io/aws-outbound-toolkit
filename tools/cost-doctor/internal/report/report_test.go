package report

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"github.com/kudig-io/aws-outbound-toolkit/tools/cost-doctor/internal/model"
	"github.com/kudig-io/aws-outbound-toolkit/tools/cost-doctor/internal/rules"
)

func fixture() (model.Inventory, []rules.Finding) {
	inv := model.Inventory{AccountID: "123456789012", Region: "ap-southeast-1"}
	findings := []rules.Finding{
		{
			Rule:                       "unattached-ebs",
			Severity:                   rules.SeverityHigh,
			ResourceID:                 "vol-1",
			ResourceType:               "ebs-volume",
			Message:                    "EBS volume is unattached.",
			Recommendation:             "Snapshot and delete.",
			EstimatedMonthlySavingsUSD: 8,
		},
	}
	return inv, findings
}

func TestWriteTable(t *testing.T) {
	inv, findings := fixture()
	var buf bytes.Buffer
	if err := Write(&buf, inv, findings, "table"); err != nil {
		t.Fatal(err)
	}
	out := buf.String()
	for _, want := range []string{"unattached-ebs", "vol-1", "8.00", "SEVERITY"} {
		if !strings.Contains(out, want) {
			t.Errorf("table output missing %q:\n%s", want, out)
		}
	}
}

func TestWriteMarkdown(t *testing.T) {
	inv, findings := fixture()
	var buf bytes.Buffer
	if err := Write(&buf, inv, findings, "markdown"); err != nil {
		t.Fatal(err)
	}
	out := buf.String()
	for _, want := range []string{"# AWS Cost Doctor report", "| Severity |", "Snapshot and delete."} {
		if !strings.Contains(out, want) {
			t.Errorf("markdown output missing %q:\n%s", want, out)
		}
	}
}

func TestWriteJSON(t *testing.T) {
	inv, findings := fixture()
	var buf bytes.Buffer
	if err := Write(&buf, inv, findings, "json"); err != nil {
		t.Fatal(err)
	}

	var doc Document
	if err := json.Unmarshal(buf.Bytes(), &doc); err != nil {
		t.Fatalf("invalid JSON: %v\n%s", err, buf.String())
	}
	if doc.Summary.TotalFindings != 1 {
		t.Fatalf("expected 1 finding, got %d", doc.Summary.TotalFindings)
	}
	if doc.Summary.EstimatedMonthlySavingsUSD != 8 {
		t.Fatalf("expected 8 savings, got %v", doc.Summary.EstimatedMonthlySavingsUSD)
	}
}

func TestWriteUnsupportedFormat(t *testing.T) {
	inv, findings := fixture()
	var buf bytes.Buffer
	if err := Write(&buf, inv, findings, "yaml"); err == nil {
		t.Fatal("expected error for unsupported format")
	}
}

func TestWriteEmptyFindings(t *testing.T) {
	inv := model.Inventory{}
	var buf bytes.Buffer
	if err := Write(&buf, inv, nil, "table"); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(buf.String(), "No cost issues found.") {
		t.Fatalf("expected empty message, got:\n%s", buf.String())
	}
}
