package cli

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/pvyswiss/pvyai-coding-agent/internal/acp"
	"github.com/pvyswiss/pvyai-coding-agent/internal/agent"
	"github.com/pvyswiss/pvyai-coding-agent/internal/config"
	"github.com/pvyswiss/pvyai-coding-agent/internal/sandbox"
	"github.com/pvyswiss/pvyai-coding-agent/internal/tools"
)

const acpUsage = `pvyai acp — serve the Agent Client Protocol (ACP) over stdio

Editors that speak ACP (Zed, JetBrains, Neovim, ...) spawn this command and drive
PVYai as a backend over JSON-RPC 2.0 on stdin/stdout. PVYai keeps your provider,
model, and API keys (BYOK); the editor only hosts the conversation thread.

Usage:
  pvyai acp

Not meant to be run interactively — point your editor's ACP / external-agent
setting at "pvyai acp".`

// runACP serves ACP over stdio so an editor can drive PVYai's agent core. It
// speaks JSON-RPC 2.0 (newline-delimited JSON) on stdin/stdout; stderr stays free
// for human-readable diagnostics. The session lifecycle maps onto PVYai's own
// session store, and provider/model/keys remain owned by PVYai.
func runACP(args []string, stdout io.Writer, stderr io.Writer, deps appDeps) int {
	for _, arg := range args {
		switch arg {
		case "-h", "--help", "help":
			fmt.Fprintln(stdout, acpUsage)
			return exitSuccess
		default:
			return writeExecUsageError(stderr, fmt.Sprintf("unknown acp flag %q", arg))
		}
	}

	conn := acp.NewConn(deps.stdin, stdout)
	acp.NewAgent(conn, acp.Deps{
		ResolveConfig: deps.resolveConfig,
		// deps.newProvider is wrapped in fillAppDeps to apply the stored API key,
		// so ACP is authenticated for apiKeyStored profiles like every other
		// surface — no ACP-specific credential handling needed.
		NewProvider: deps.newProvider,
		RunAgent:    agent.Run,
		// Build the SCOPED registry + sandbox engine per workspace, exactly like the
		// exec surface, so ACP shell/file tools are confined — never run unconfined.
		BuildWorkspace: func(workspaceRoot string, resolved config.ResolvedConfig) (*tools.Registry, *sandbox.Engine, error) {
			scope, err := sandbox.NewScope(workspaceRoot, resolved.Sandbox.AdditionalWriteRoots)
			if err != nil {
				return nil, nil, err
			}
			engine, err := buildExecSandboxEngine(workspaceRoot, resolved, deps, scope)
			if err != nil {
				return nil, nil, err
			}
			registry := newCoreRegistryScoped(workspaceRoot, scope)
			registerLocalControlTools(registry, workspaceRoot, resolved.LocalControl)
			return registry, engine, nil
		},
		ResolveWorkspaceRoot: acpWorkspaceRootResolver(deps),
		Store:                deps.newSessionStore(),
		AgentInfo:            acp.Implementation{Name: "pvyai", Version: version},
	})

	ctx, stop := signalContext()
	defer stop()
	if err := conn.Serve(ctx); err != nil && ctx.Err() == nil {
		return writeAppError(stderr, "acp: "+err.Error(), exitCrash)
	}
	return exitSuccess
}

// acpWorkspaceRootResolver validates a client-supplied cwd into a confinement
// root. It reuses exec's resolveWorkspaceRoot (abs+clean, must be an existing
// dir) and additionally rejects the filesystem root and the home directory — an
// editor must not be able to point PVYai's file/shell tools at the whole disk.
func acpWorkspaceRootResolver(deps appDeps) func(string) (string, error) {
	return func(cwd string) (string, error) {
		root, err := resolveWorkspaceRoot(cwd, deps)
		if err != nil {
			return "", err
		}
		if root == filepath.Dir(root) {
			return "", fmt.Errorf("cwd must not be the filesystem root: %s", root)
		}
		if home, herr := os.UserHomeDir(); herr == nil && home != "" && filepath.Clean(home) == root {
			return "", fmt.Errorf("cwd must not be the home directory: %s", root)
		}
		return root, nil
	}
}
