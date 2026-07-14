<p align="center">
  <img src="https://raw.githubusercontent.com/pvyswiss/pvyai-coding-agent/main/docs/assets/pvai-logo.png" alt="PVYai" width="385">
</p>

<p align="center"><strong>A terminal coding agent you actually own.</strong></p>

<p align="center">
  <a href="LICENSE"><img alt="license" src="https://img.shields.io/badge/license-MIT-blue"></a>
  <img alt="Go 1.25+" src="https://img.shields.io/badge/Go-1.25+-00ADD8?logo=go&logoColor=white">
  <img alt="25+ providers" src="https://img.shields.io/badge/providers-25+-34E2EA">
  <br>
  <strong>English</strong> | <a href="README_ZH.md">中文</a>
</p>

PVYai is an AI coding agent for your local terminal. It can inspect a repository,
edit files, run commands, use browser/terminal helpers, and keep durable local
sessions while you choose the model and the permission level.

```bash
pvyai
pvyai exec "fix the failing test in ./pkg"
pvyai exec --output-format stream-json < turns.jsonl
```

## Why PVYai Agent

- **Use the model you want.** Bring PVYai, OpenAI, Anthropic, Gemini, Groq, OpenRouter,
  DeepSeek, Mistral, xAI, Qwen, Kimi, GitHub Models, Ollama, LM Studio, or any
  OpenAI-/Anthropic-compatible endpoint.
- **Stay in control.** File writes, shell commands, network access, and
  out-of-workspace writes go through PVYai Agent permission and sandbox policy.
- **Works in the terminal.** The TUI has model/provider pickers, image input,
  slash commands, live plan/tool rendering, scrollback, themes, and resume/fork
  support.
- **Works without the TUI.** `pvyai exec` is scriptable, supports text/JSON/
  stream-JSON I/O, isolated worktrees, spec-first runs, and meaningful exit
  codes for CI.
- **Keeps context local.** Sessions are stored on disk, searchable, resumable,
  and never uploaded as telemetry by PVYai.
- **Extensible when you need it.** Use local MCP servers, skills, plugins, hooks, and
  specialist subagents from the same CLI.
- **PVYai Sentinel integration.** PVYai Sentinel watch and simulates the outcome 
  of your LLM interations in real time, and prevents harmful actions before LLM can execute.
- **PVYai Memory integration.** PVYai Memory offers for PVYai LLM Models automatic and manual
  Memory management for findings, session handover, iniated and monitored by PVYai Sentinel Intelligence.
  It offers users Private Memory Wing or allocated Primary Shared Memory, Secondary Shared Memory and Third,
  for interdisciplinary cross-domain preservation.
- **PVYai Sentinel Orchestration.** As an Orchestrator, he manage your Team Agents and evict and re-onboard 
  them automatic based on real-time metrics instead of classical turn/tool call count.
- **PVYai Unified SSE API.** Chosing PVYai as Provider, drives our sophisticated, OpenAI API compatible,
  unified SSE Stream API, which offers native function call for our MCP Services. Your PVYai Agent can execute
  all tools offered in parallel instead of serial curl commands, also the Specialists/Sub-agents.

## Informations: 
- Website: [https://pvy.swiss/ai](https://pvy.swiss/ai)
- Documentation: [https://docs.pvy.swiss/en/pvyapps/pvyai/pvyai-agent](https://docs.pvy.swiss/en/pvyapps/pvyai/pvyai-agent)

## Install

### npm

```bash
npm install -g @pvyswiss/pvyai-agent
pvyai
```

The npm package installs a small wrapper plus the matching PVYai binary for your
platform from GitHub Releases. It supports Linux, macOS, and Windows on x64 and
arm64.

### Bun

Bun does not run dependency lifecycle scripts by default, so the `postinstall`
that fetches the PVYai binary is skipped and the first run fails with
`No native binary found next to the npm wrapper`.

The simplest fix is to trust the package after installing, which runs the
blocked postinstall. This works for project and global installs:

```bash
# project install
bun add @pvyswiss/pvyai-agent
bun pm trust @pvyswiss/pvyai-agent

# global install
bun add -g @pvyswiss/pvyai-agent
bun pm -g trust @pvyswiss/pvyai-agent
```

Alternatives: allow the postinstall up front by adding
`"trustedDependencies": ["@pvyswiss/pvyai-agent"]` to your project's package.json
before `bun add`, or run the installer manually
(`node node_modules/@pvyswiss/pvyai/scripts/postinstall.mjs`) on Bun versions
that do not have `bun pm trust`.

### Install scripts

Linux/macOS:

```bash
curl -fsSL https://raw.githubusercontent.com/pvyswiss/pvyai-coding-agent/main/scripts/install.sh | bash
```

Windows PowerShell:

```powershell
irm https://raw.githubusercontent.com/pvyswiss/pvyai-coding-agent/main/scripts/install.ps1 | iex
```

### From source

Source builds require Go 1.25+.

```bash
git clone https://github.com/pvyswiss/pvyai-coding-agent.git
cd pvyai-coding-agent
go run ./cmd/pvyai
```

Release installers and the npm wrapper require published GitHub Release assets.
If you are testing before the first public release, build from source:

```bash
go build -o pvyai ./cmd/pvyai
```

On Linux, build the sandbox helper too if you want native sandboxing:

```bash
go build -o pvyai-sandbox ./cmd/pvyai-linux-sandbox
go build -o pvyai-seccomp ./cmd/pvyai-seccomp   # optional compatibility wrapper
```

Put `zero` and `pvyai-sandbox` in the same directory on `PATH`
(`~/.local/bin` is a good default). macOS does not need an extra helper binary.
Windows source builds can use the main `pvyai.exe` as their sandbox helper; release
archives still ship standalone Windows helper executables.

More install details: [docs/INSTALL.md](docs/INSTALL.md).

## First Run

Start the TUI:

```bash
pvyai
```

The setup wizard helps you pick a provider and model. You can also configure
providers from the command line:

```bash
pvyai setup
pvyai providers list
pvyai models list
pvyai doctor
```

For API providers, set the matching environment variable before setup or enter
the key in the wizard:

```bash
export OPENAI_API_KEY=sk-...
export ANTHROPIC_API_KEY=...
export GEMINI_API_KEY=...
export LONGCAT_API_KEY=...
```

To configure Meituan LongCat (LongCat-2.0) directly, run:

```bash
pvyai providers setup longcat --set-active
```

For local models, run Ollama or LM Studio and then use `pvyai setup` or
`pvyai providers detect`.

## Daily Use

### Interactive TUI

```bash
pvai
```

Useful controls:

| Control     | Action                         |
| ----------- | ------------------------------ |
| `Enter`     | send the prompt                |
| `/`         | open slash-command suggestions |
| `Shift+Tab` | cycle permission mode          |
| `Ctrl+X`    | show/hide the sidebar          |
| `Ctrl+C`    | cancel or exit                 |

Common slash commands:

| Command                        | Purpose                                         |
| ------------------------------ | ----------------------------------------------- |
| `/model`, `/provider`          | switch the active model/provider                |
| `/spec`, `/plan`               | draft and review a plan before building         |
| `/image`                       | attach an image for vision-capable models       |
| `/resume`, `/rewind`           | continue or roll back local sessions            |
| `/compact`, `/context`         | manage context usage                            |
| `/permissions`, `/tools`       | inspect available tools and policy              |
| `/add-dir`                     | allow an extra write directory for this session |
| `/theme`, `/doctor`, `/config` | adjust appearance and inspect setup             |

### Headless `exec`

```bash
pvyai exec "explain internal/agent/loop.go"
pvyai exec --model claude-sonnet-4.5 "refactor the config loader"
pvyai exec --use-spec "add rate limiting to the API client"
pvyai exec --worktree "try the migration in an isolated worktree"
pvyai exec --resume
pvyai exec --fork <session-id> "try the other approach"
```

Programmatic use:

```bash
pvyai exec --input-format stream-json --output-format stream-json < turns.jsonl
```

The stream-JSON contract is documented in
[docs/STREAM_JSON_PROTOCOL.md](docs/STREAM_JSON_PROTOCOL.md).

## Safety Model

PVYai is designed to make side effects visible.

- Workspace reads are allowed by default.
- File writes are limited to the workspace unless you grant another directory.
- Shell commands, network access, destructive commands, and elevated actions are
  permission-gated.
- `--add-dir <path>` and `/add-dir <path>` grant additional write roots without
  giving the agent the whole filesystem.
- Unsafe/autonomous modes are explicit opt-ins.
- Secrets are redacted from tool output and logs where PVYai controls the surface.
- PVYai Sentinel Real World Model watches, predicts and simulate real action outcome, blocks harfmul action before LLM can execute it (Only for PVYai as Provider, but Anthropic Models included)

Example:

```bash
pvyai --add-dir ../docs-site
pvyai exec --add-dir ../shared "update both repos"
```

Sandbox behavior can be inspected with:

```bash
pvyai sandbox policy
pvyai sandbox grants list
```

## Web And Local Control

PVYai includes local file/search/edit/shell tools, `web_fetch` for public URLs,
and MCP support for additional tools.

For local dev servers, use shell commands such as `curl` through `exec_command`
so the normal sandbox and permission policy applies. Long-running commands stay
attached to a background terminal session and can be listed or stopped from the
TUI.

The npm package also includes browser and terminal helper packages used by local
browser/terminal tools. Source builds can use the same helpers when they are on
`PATH` or configured in PVYai's local-control settings.

## Common Commands

```text
pvyai                  interactive TUI
pvyai exec             one-shot or scripted agent run
pvyai setup            first-run provider setup
pvyai auth             OAuth/login helpers for supported providers
pvyai models           model registry and capabilities
pvyai providers        provider profiles and detection
pvyai doctor           setup, key, and connectivity checks
pvyai context          context-budget report
pvyai repo-map         deterministic repository map
pvyai repo-info        local repository summary
pvyai search | find    search local session history
pvyai sessions         inspect, resume, fork, and rewind sessions
pvyai spec             manage spec-mode drafts
pvyai specialist       manage specialist subagents
pvyai skills           manage markdown instruction skills
pvyai plugins          manage plugins
pvyai hooks            manage lifecycle hooks
pvyai mcp              manage MCP servers and tools
pvyai serve --mcp      expose PVYai tools over MCP stdio
pvyai sandbox          inspect sandbox policy and grants
pvyai worktrees        prepare isolated git worktrees
pvyai verify           detect and run local verification checks
pvyai changes          inspect and commit local git changes
pvyai usage            token usage and estimated cost
pvyai cron             scheduled agent jobs
pvyai update           check for newer releases
```

## Extending PVYai

### Project and personal instructions

PVYai appends project-specific guidance to the system prompt from the first
`AGENTS.md`, `PVYAI.md`, or `.pvyai/AGENTS.md` file found in each directory from
the git root down to your current working directory (checked in that order
per directory). Files are injected general-to-specific, capped at 8 KiB per
file and 32 KiB total.

A personal `PVYAI.md` under `config.UserConfigDir()/pvyai/PVYAI.md`
(`$XDG_CONFIG_HOME/pvyai/PVYAI.md` or `~/.config/pvyai/PVYAI.md` on Linux/macOS,
`%AppData%\Roaming\pvyai\PVYAI.md` on Windows) applies across every workspace, ahead of any project guidelines.

### Plugins

Plugins are discovered from `~/.config/pvyai/plugins/<name>/plugin.json` (user
scope — `$XDG_CONFIG_HOME` or `~/.config` on every OS, independent of the
`config.UserConfigDir()` path used above) and `<cwd>/.pvyai/plugins/<name>/plugin.json`
(project scope — resolved from the current working directory, not the repo
root), and managed with `pvyai plugins`. A manifest can declare:

- `tools` — custom tools (`command`, `args`, `inputSchema`, and a
  `permission` of `prompt` or `deny`; `allow` is honored only when manifest tool
  auto-approval is enabled)
- `hooks` — commands run on `beforeTool`, `afterTool`, `sessionStart`, or
  `sessionEnd`
- `prompts` and `skills` — additional prompt/skill files

MCP servers (`pvyai mcp`) and standalone markdown skills (`pvyai skills`) use
the same extension points and can also be wired up outside of a plugin
manifest.

## Appearance And Accessibility

| Control               | Effect                                                                                                                                                                                                           |
| --------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `NO_COLOR=<anything>` | disables color output                                                                                                                                                                                            |
| `PVYAI_THEME=<name>`  | selects the startup theme (`auto`, `dark`, `light`, or a color theme like `dracula`, `nord`, `gruvbox`, `tokyo-night`, `catppuccin`, `one-dark`, `solarized-dark`, `rose-pine`, `everforest`, `solarized-light`) |
| `--theme <name>`      | selects the TUI theme from the CLI (same names)                                                                                                                                                                  |
| `/theme`              | opens the theme picker inside the TUI (live preview; `/theme <name>` switches directly)                                                                                                                          |
| `PVYAI_NO_FADE=1`     | disables streaming fade animation                                                                                                                                                                                |

Meaning does not rely on color alone; diffs, permissions, and statuses also use
text or glyph markers.

## Development

```bash
go test ./...
go run ./cmd/pvyai-release build
go run ./cmd/pvyai-release smoke
go run ./cmd/pvyai-perf-bench
```

Cross-compile examples:

```bash
go run ./cmd/pvyai-release build --goos linux --goarch amd64
go run ./cmd/pvyai-release build --goos windows --goarch amd64 --output dist/pvyai.exe
```

## Documentation

- [Install](docs/INSTALL.md)
- [Update flow](docs/UPDATE.md)
- [Stream-JSON protocol](docs/STREAM_JSON_PROTOCOL.md)
- [Specialists](docs/SPECIALISTS.md)
- [GitHub Action](docs/GITHUB_ACTION.md)
- [Benchmarks](docs/BENCHMARK.md)
- [Performance](docs/PERFORMANCE.md)
- [Agent evals](docs/AGENT_EVALS.md)

## Contributing

Contributions are welcome. Read [CONTRIBUTING.md](CONTRIBUTING.md), run the
relevant tests, and open a focused pull request.

Security reports should follow [SECURITY.md](SECURITY.md).

## License

PVYai is released under the [MIT License](LICENSE).
