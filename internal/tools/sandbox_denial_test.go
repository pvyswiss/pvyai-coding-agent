package tools

import (
	"testing"

	pvySandbox "github.com/pvyswiss/pvyai-coding-agent/internal/sandbox"
)

func TestLikelySandboxDeniedDetectsReferenceKeywords(t *testing.T) {
	plan := pvySandbox.CommandPlan{
		Wrapped:       true,
		TargetBackend: pvySandbox.BackendLinuxBwrap,
	}
	output := "touch: cannot touch '/home/user/.npm/cache': Read-only file system"
	if !likelySandboxDenied(plan, 1, output) {
		t.Fatalf("expected reference sandbox denial keyword to be classified as sandbox denied")
	}
}

func TestLikelySandboxDeniedDetectsNetworkDenialEvenWithZeroExit(t *testing.T) {
	plan := pvySandbox.CommandPlan{
		Wrapped:       true,
		TargetBackend: pvySandbox.BackendLinuxBwrap,
		Policy:        pvySandbox.Policy{Network: pvySandbox.NetworkDeny},
		PermissionProfile: pvySandbox.PermissionProfile{
			Network: pvySandbox.NetworkPolicy{Mode: pvySandbox.NetworkDeny},
		},
	}
	if !likelySandboxDenied(plan, 0, "Cannot open a network socket.") {
		t.Fatal("network-denied socket output with exit 0 must be classified as sandbox denied")
	}
	meta := map[string]string{}
	markLikelySandboxDenial(meta, plan, 0, "Cannot open a network socket.")
	if meta[SandboxLikelyDeniedMeta] != "true" || meta[SandboxDenialKindMeta] != SandboxDenialKindNetwork {
		t.Fatalf("network denial meta = %#v", meta)
	}
}

func TestLikelySandboxDeniedIgnoresUnsandboxedFailure(t *testing.T) {
	plan := pvySandbox.CommandPlan{Wrapped: false}
	if likelySandboxDenied(plan, 1, "permission denied") {
		t.Fatal("unsandboxed command output must not be classified as a sandbox denial")
	}
}
