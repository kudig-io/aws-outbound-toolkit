package cli

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const sample = `{
  "account_id": "123456789012",
  "region": "ap-southeast-1",
  "ebs_volumes": [
    {"id": "vol-1", "state": "available", "size_gb": 100, "monthly_cost_usd": 8}
  ]
}`

func writeSample(t *testing.T) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "inventory.json")
	if err := os.WriteFile(path, []byte(sample), 0o600); err != nil {
		t.Fatal(err)
	}
	return path
}

func TestScanSucceeds(t *testing.T) {
	var out, errOut bytes.Buffer
	code := Run([]string{"scan", "--input", writeSample(t)}, &out, &errOut)
	if code != 0 {
		t.Fatalf("expected exit 0, got %d: %s", code, errOut.String())
	}
	if !strings.Contains(out.String(), "unattached-ebs") {
		t.Fatalf("expected finding in output, got:\n%s", out.String())
	}
}

func TestScanFailOnHighReturnsOne(t *testing.T) {
	var out, errOut bytes.Buffer
	code := Run([]string{"scan", "--input", writeSample(t), "--fail-on", "high"}, &out, &errOut)
	if code != 1 {
		t.Fatalf("expected exit 1, got %d", code)
	}
}

func TestScanFailOnCriticalReturnsZero(t *testing.T) {
	var out, errOut bytes.Buffer
	code := Run([]string{"scan", "--input", writeSample(t), "--fail-on", "critical"}, &out, &errOut)
	if code != 0 {
		t.Fatalf("expected exit 0, got %d", code)
	}
}

func TestScanRequiresInput(t *testing.T) {
	var out, errOut bytes.Buffer
	if code := Run([]string{"scan"}, &out, &errOut); code != 2 {
		t.Fatalf("expected exit 2, got %d", code)
	}
	if !strings.Contains(errOut.String(), "--input is required") {
		t.Fatalf("unexpected stderr: %s", errOut.String())
	}
}

func TestScanRejectsBadJSON(t *testing.T) {
	path := filepath.Join(t.TempDir(), "bad.json")
	if err := os.WriteFile(path, []byte("{"), 0o600); err != nil {
		t.Fatal(err)
	}
	var out, errOut bytes.Buffer
	if code := Run([]string{"scan", "--input", path}, &out, &errOut); code != 2 {
		t.Fatalf("expected exit 2, got %d", code)
	}
}

func TestUnknownCommand(t *testing.T) {
	var out, errOut bytes.Buffer
	if code := Run([]string{"frobnicate"}, &out, &errOut); code != 2 {
		t.Fatalf("expected exit 2, got %d", code)
	}
}

func TestVersion(t *testing.T) {
	var out, errOut bytes.Buffer
	if code := Run([]string{"version"}, &out, &errOut); code != 0 {
		t.Fatalf("expected exit 0, got %d", code)
	}
	if strings.TrimSpace(out.String()) == "" {
		t.Fatal("expected version output")
	}
}
