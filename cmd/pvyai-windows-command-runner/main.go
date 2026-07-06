package main

import (
	"os"

	"github.com/pvyswiss/pvyai-coding-agent/internal/sandbox"
)

func main() {
	os.Exit(sandbox.RunWindowsSandboxCommandRunner(os.Args[1:], os.Stderr))
}
