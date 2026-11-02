// Package report renders findings in table, markdown, or JSON form.
package report

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"text/tabwriter"

	"github.com/kudig-io/aws-outbound-toolkit/tools/cost-doctor/internal/model"
	"github.com/kudig-io/aws-outbound-toolkit/tools/cost-doctor/internal/rules"
)

// Document is the JSON representation of a scan result.
type Document struct {
	AccountID string          `json:"account_id"`
	Region    string          `json:"region"`
	Currency  string          `json:"currency,omitempty"`
	Summary   Summary         `json:"summary"`
	Findings  []rules.Finding `json:"findings"`
}

// Summary aggregates scan results.
type Summary struct {
	TotalFindings              int            `json:"total_findings"`
	BySeverity                 map[string]int `json:"by_severity"`
	EstimatedMonthlySavingsUSD float64        `json:"estimated_monthly_savings_usd"`
}

// Build returns the normalized document for a scan.
func Build(inv model.Inventory, findings []rules.Finding) Document {
	return Document{
		AccountID: inv.AccountID,
		Region:    inv.Region,
		Currency:  inv.Currency,
		Summary: Summary{
			TotalFindings:              len(findings),
			BySeverity:                 rules.CountBySeverity(findings),
			EstimatedMonthlySavingsUSD: rules.TotalEstimatedSavings(findings),
		},
		Findings: findings,
	}
}

// Write renders a scan result in the requested format: table, markdown, or json.
func Write(w io.Writer, inv model.Inventory, findings []rules.Finding, format string) error {
	switch strings.ToLower(format) {
	case "table", "":
		return writeTable(w, inv, findings)
	case "markdown", "md":
		return writeMarkdown(w, inv, findings)
	case "json":
		return writeJSON(w, inv, findings)
	default:
		return fmt.Errorf("unsupported format %q (want table, markdown, or json)", format)
	}
}

func writeTable(w io.Writer, inv model.Inventory, findings []rules.Finding) error {
	fmt.Fprintf(w, "Account: %s   Region: %s   Findings: %d   Estimated savings: $%.2f/month\n\n",
		display(inv.AccountID), display(inv.Region), len(findings), rules.TotalEstimatedSavings(findings))

	if len(findings) == 0 {
		_, err := fmt.Fprintln(w, "No cost issues found.")
		return err
	}

	tw := tabwriter.NewWriter(w, 0, 4, 2, ' ', 0)
	fmt.Fprintln(tw, "SEVERITY\tRULE\tRESOURCE\tEST. SAVINGS\tMESSAGE")
	for _, f := range findings {
		fmt.Fprintf(tw, "%s\t%s\t%s\t$%.2f\t%s\n",
			f.Severity, f.Rule, f.ResourceID, f.EstimatedMonthlySavingsUSD, f.Message)
	}
	return tw.Flush()
}

func writeMarkdown(w io.Writer, inv model.Inventory, findings []rules.Finding) error {
	fmt.Fprintf(w, "# AWS Cost Doctor report\n\n")
	fmt.Fprintf(w, "- Account: `%s`\n- Region: `%s`\n- Findings: **%d**\n- Estimated monthly savings: **$%.2f**\n\n",
		display(inv.AccountID), display(inv.Region), len(findings), rules.TotalEstimatedSavings(findings))

	if len(findings) == 0 {
		_, err := fmt.Fprintln(w, "No cost issues found.")
		return err
	}

	fmt.Fprintln(w, "| Severity | Rule | Resource | Est. savings (USD/month) | Message |")
	fmt.Fprintln(w, "| --- | --- | --- | ---: | --- |")
	for _, f := range findings {
		fmt.Fprintf(w, "| %s | `%s` | `%s` | %.2f | %s |\n",
			f.Severity, f.Rule, f.ResourceID, f.EstimatedMonthlySavingsUSD, escapePipe(f.Message))
	}
	fmt.Fprintln(w)
	fmt.Fprintln(w, "## Recommendations")
	fmt.Fprintln(w)
	for _, f := range findings {
		fmt.Fprintf(w, "- **%s** (`%s`): %s\n", f.ResourceID, f.Rule, f.Recommendation)
	}
	return nil
}

func writeJSON(w io.Writer, inv model.Inventory, findings []rules.Finding) error {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(Build(inv, findings))
}

func display(s string) string {
	if s == "" {
		return "-"
	}
	return s
}

func escapePipe(s string) string {
	return strings.ReplaceAll(s, "|", "\\|")
}
