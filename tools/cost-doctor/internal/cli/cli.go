// Package cli implements the aws-cost-doctor command line.
package cli

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/kudig-io/aws-outbound-toolkit/tools/cost-doctor/internal/model"
	"github.com/kudig-io/aws-outbound-toolkit/tools/cost-doctor/internal/report"
	"github.com/kudig-io/aws-outbound-toolkit/tools/cost-doctor/internal/rules"
)

const usage = `aws-cost-doctor - deterministic AWS cost checks from an inventory snapshot

Usage:
  aws-cost-doctor scan --input inventory.json [--format table|markdown|json] [--fail-on high]

Commands:
  scan      Run cost rules against an inventory file
  version   Print version
  help      Show this help

Examples:
  aws-cost-doctor scan --input examples/inventory.json
  aws-cost-doctor scan --input examples/inventory.json --format markdown > report.md
  aws-cost-doctor scan --input examples/inventory.json --fail-on high
`

// Version is set at build time via -ldflags.
var Version = "dev"

// Run executes the CLI and returns a process exit code.
func Run(args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		fmt.Fprint(stderr, usage)
		return 2
	}

	switch args[0] {
	case "help", "-h", "--help":
		fmt.Fprint(stdout, usage)
		return 0
	case "version", "--version":
		fmt.Fprintln(stdout, Version)
		return 0
	case "scan":
		return runScan(args[1:], stdout, stderr)
	default:
		fmt.Fprintf(stderr, "unknown command %q\n\n%s", args[0], usage)
		return 2
	}
}

func runScan(args []string, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("scan", flag.ContinueOnError)
	fs.SetOutput(stderr)

	input := fs.String("input", "", "path to an inventory JSON file (required)")
	format := fs.String("format", "table", "output format: table, markdown, or json")
	failOn := fs.String("fail-on", "", "exit non-zero when a finding at or above this severity is present: low, medium, high, critical")

	if err := fs.Parse(args); err != nil {
		return 2
	}
	if *input == "" {
		fmt.Fprintln(stderr, "error: --input is required")
		return 2
	}

	data, err := os.ReadFile(*input)
	if err != nil {
		fmt.Fprintf(stderr, "error: read input: %v\n", err)
		return 2
	}

	var inv model.Inventory
	if err := json.Unmarshal(data, &inv); err != nil {
		fmt.Fprintf(stderr, "error: parse inventory: %v\n", err)
		return 2
	}

	findings := rules.Scan(inv)
	if err := report.Write(stdout, inv, findings, *format); err != nil {
		fmt.Fprintf(stderr, "error: render report: %v\n", err)
		return 2
	}

	if *failOn != "" {
		threshold, err := parseSeverity(*failOn)
		if err != nil {
			fmt.Fprintf(stderr, "error: %v\n", err)
			return 2
		}
		if hasFindingAtOrAbove(findings, threshold) {
			return 1
		}
	}
	return 0
}

func parseSeverity(s string) (string, error) {
	switch strings.ToLower(s) {
	case rules.SeverityLow:
		return rules.SeverityLow, nil
	case rules.SeverityMedium:
		return rules.SeverityMedium, nil
	case rules.SeverityHigh:
		return rules.SeverityHigh, nil
	case rules.SeverityCritical:
		return rules.SeverityCritical, nil
	default:
		return "", errors.New("--fail-on must be one of: low, medium, high, critical")
	}
}

func severityRank(s string) int {
	switch s {
	case rules.SeverityCritical:
		return 4
	case rules.SeverityHigh:
		return 3
	case rules.SeverityMedium:
		return 2
	case rules.SeverityLow:
		return 1
	default:
		return 0
	}
}

func hasFindingAtOrAbove(findings []rules.Finding, threshold string) bool {
	rank := severityRank(threshold)
	for _, f := range findings {
		if severityRank(f.Severity) >= rank {
			return true
		}
	}
	return false
}
