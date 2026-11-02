package main

import (
	"os"

	"github.com/kudig-io/aws-outbound-toolkit/tools/cost-doctor/internal/cli"
)

func main() {
	os.Exit(cli.Run(os.Args[1:], os.Stdout, os.Stderr))
}
