//go:build linux

// Command pvyai-seccomp is retained as a compatibility wrapper for existing Linux
// installs and scripts. The main sandbox path now applies the same optional
// Unix-socket filter inside pvyai-linux-sandbox when sandbox.blockUnixSockets is
// enabled.
package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github.com/pvyswiss/pvyai-coding-agent/internal/sandbox"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: pvyai-seccomp <command> [args...]")
		os.Exit(2)
	}
	if err := sandbox.ApplyUnixSocketBlock(); err != nil {
		fmt.Fprintln(os.Stderr, "pvyai-seccomp: warning: "+err.Error()+"; running without the Unix-socket filter")
	}
	binary, err := exec.LookPath(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "pvyai-seccomp: "+err.Error())
		os.Exit(127)
	}
	if err := syscall.Exec(binary, os.Args[1:], os.Environ()); err != nil {
		fmt.Fprintln(os.Stderr, "pvyai-seccomp: exec failed: "+err.Error())
		os.Exit(126)
	}
}
