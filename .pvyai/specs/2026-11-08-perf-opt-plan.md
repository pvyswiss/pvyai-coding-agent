---
 title: performance-opt — Team/Swarm, Loop Config & Execution Plan (CTO-APPROVED v0.6)
 status: CTO-APPROVED v0.6 — CTO approved 2026-11-08 (§6 resolved: Coder Flash = `pvy-coder-flash-2`, Bubble Tea upgrade = follow-up / watchdog-only in P0, gate score 68); skill reference (§3.6) + config-reconciliation note (§3.1) added; execution green-lit (E → P0 → (P1 ∥ P2) → P3 → P4); supersedes v0.5
branch: performance-opt (from pvy-develop)
author: PVY Architect M (pvy-architect-m)
memory_wing: PVY DevOps Shared Memory (pvy-devops)
date: 2026-11-08
repo: /Volumes/Data/Repos/ollama-plus-openweb-ui/coding-agent
---

# performance-opt — PVYai Agent Team/Swarm, Loop Configuration & Plans (CTO-APPROVED v0.6)

Programme: **TUI resilience (stale-after-idle bug) + performance optimisation** of the PVYai Coder agent Go codebase.

---

## 1. Situation Analysis (summary of code audit)

### 1.1 TUI staleness after ~30 min idle — root-cause chain

There is **no literal 30-minute timer** in the repo — the staleness is emergent. The ranked chain:

| # | Root cause | Evidence |
| --- | --- | --- |
| 1 | **Event-loop stall with no watchdog.** Every runtime event funnels through a *blocking* `program.Send` (`internal/tui/run.go:28-35`). The `textCoalescer` forwards **while holding its mutex** (`internal/tui/coalesce.go` drainAndForwardLocked), so one stalled Send blocks every agent callback and the PR watcher behind a lock. When the loop finally ingests one message, the backlog flushes in a burst — exactly the "stale until an event wakes it" symptom. Precedent: `internal/tui/model.go:1850` documents a prior "stuck 33 minutes" deadlock. This happens primarerly for Orchestrator or just Singe Agents without a Swarm/Loop Configuration in Place. | run.go, coalesce.go, model.go |
| 2 | **No keepalive / repaint watchdog.** The only unconditional idle repaint is the 530 ms composer blink chain (`model.go:849/1038`); the spinner tick is *deliberately stopped* when idle (`model.go:1773-1787`). If the blink TickMsg stops being delivered, the screen is permanently static — nothing detects or recovers. | model.go, streaming_fade.go |
| 3 | **PR watcher is the only idle event producer and leaks goroutines.** 60 s ticker (`pr_status.go:53,165-176`) spawns an unbounded goroutine per tick running `gh`/`glab`; wedged probes accumulate ~1/min (~30 in 30 min) and pile onto the blocked pipeline. It is also the likely "wake" event. | pr_status.go |
| Ruled out | Provider streams (5 min idle watchdog, `providerio.go:56`), reconnect logic (`reconnect.go` — establishment only), daemon/control sockets (no idle deadlines), HTTP IdleConnTimeout 30 s (irrelevant while idle). | — |

**Fix strategy (P0):** (a) non-blocking `forward`/`Send` with bounded queue + drop-with-flag semantics; (b) move `c.forward` outside `c.mu`; (c) add a 30 s self-healing repaint/keepalive watchdog (force a frame + detect a dead blink chain and re-arm it); (d) bound the PR watcher to one reusable worker with a hard context deadline.

### 1.2 Performance — ranked findings

| # | Finding | Impact | Effort | Where |
| --- | --- | --- | --- | --- |
| 1 | Per-frame O(transcript bytes) cache-**key** rebuilding: `renderRowCacheKey` embeds full row text per row per frame (~1 MB string building on a long session per keystroke/16 ms flush) | **HIGH** | med | `internal/tui/render_cache.go`, `transcript_selection.go:256,361,659` |
| 2 | Streaming markdown re-renders the **entire accumulated text** per ~16 ms flush: full copy + sha256 + markdown parse + syntax highlighting → O(n²) per turn | **HIGH** | med-high | `internal/tui/assistant_markdown.go:180-197`, `model.go:2986` |
| 3 | Body pipeline runs 2–3× per frame (syncChatScroll + View + two-column sidebar each rebuild body items & metrics) | **HIGH** | med | `model.go:1013,2374,2725,2888` |
| 4 | `buildRowContext` allocates 5 maps and scans all rows every frame; string map keys allocate per lookup | MED | low-med | `internal/tui/rendering.go:64-114` |
| 5 | `repomap.Scan` + `workspaceindex.Scan` run synchronously on the critical path **per user message and per specialist Run** (TTFT cost); no caching | MED | low | `internal/agent/loop.go:135`, `system_prompt.go:604-616,583-602` |
| 6 | `streamingTextString()` copies the whole stream buffer 3+× per frame | MED-LOW | trivial | `model.go:4939-4941` |
| 7 | `staticRenderCache` rune-counts values O(n) and renders under mutex | LOW | trivial | `render_cache.go` |
| 8 | Compaction does repeated O(context) passes per check | LOW | — | `compaction.go:393-484` |
| Good | Viewport virtualization exists; 16 ms coalescer with buffer reuse is well designed; token counting is heuristic; git sweeps are async. **Gap: no Go micro-benchmarks for the render pipeline** (`internal/perfbench` covers cold start / first-output only) | — | — | — |

---

## 2. PVYai Agent Team/Swarm Configuration (draft)

### 2.1 Role roster

| Role | Agent / Specialist ID | Model (platform ID → base) | Key parameters | Tools | Responsibility |
| --- | --- | --- | --- | --- | --- |
| Architect (this role) | `pvy-architect-x` | `pvy-architect-x` → **GLM 5.3** | ctx 1 048 576, max 128 000, T 1.0, top_p 0.95, think: high, native FC, vision, Sentinel quality ≥ 68 | read-only, plan | Root-cause analysis, ADRs, phase gates, code review of P0 |
| Lead Developer | `pvy-lead-dev-x` | `pvy-lead-dev-x` → **GLM 5.3** | ctx 1 048 576, max 128 000, T 1.0, top_p 0.95, think: high, native FC, vision | read-only, edit, execute, plan | Owns P0 (TUI staleness) implementation & integration, merges worker output |
| Senior Lead Dev | `pvy-senior-lead-dev-5` | `pvy-senior-lead-dev-5` → **GLM-Flash 5.3** (fast supervision) | ctx 1 048 576, max 32 000, no explicit sampling/reasoning preset, native FC | read-only, edit, execute, plan | Owns P1 render pipeline; reviews Coder Flash output; keeps diff scope tight |
| Coder | `pvy-coder-flash-2` | `pvy-coder-flash-2` → **DeepSeek Flash** | ctx 1 000 000, max 32 768, reasoning_effort: low, native FC | read-only, edit, execute | Mechanical, well-specified refactors (P1 items #1, #4, #6, #7; P2) |
| Bench/QA | `pvy-bench-qa` *(proposed)* | inherit session | — | read-only, execute, plan | P3 benchmarks, regression gates, profiling runs |

### 2.2 Topology

Star topology from the orchestrator (the platform constraint — specialists cannot delegate, `Task`/`swarm_*` are forbidden in member tool universes):

```
                    ┌─────────────────────────┐
                    │  pvy-architect-x (GLM)  │  plan gates, ADR, review
                    └───────────┬─────────────┘
                                │
              ┌─────────────────┼─────────────────┐
              ▼                                   ▼
   ┌─────────────────────┐             ┌─────────────────────┐
   │ pvy-lead-dev-x│  P0 owner   │ pvy-senior-lead-dev-5│ P1 owner
   │      (GLM 5.3)      │             │   (GLM-Flash 5.3)    │
   └──────────┬──────────┘             └──────────┬───────────┘
              ▼                                   ▼
      ┌───────────────┐                  ┌────────────────┐
      │ pvy-coder-flash-2│ P0 helpers      │ pvy-coder-flash-2 │ P1 items
      │ (DeepSeek V4p) │                 │ (DeepSeek V4p)  │
      └───────────────┘                  └────────────────┘
```

Parallelism rule: **one worktree per parallel stream** (`internal/worktrees/`), files disjoint per coder, lead integrates. Swarm messaging (`swarm_send`/`swarm_inbox`) for cross-stream coordination; `swarm_collect` for result harvesting (default 600 s, cap 1800 s).

### 2.3 Model binding — PVYai Custom Catalog resolution (approach corrected 2026-11-08)

**Run & deployment model (out of scope — do not change):** the binary runs **in the git repo**. The sandbox binds to the folder the binary is started in by design (a checked-out repo cannot widen its own sandbox — `AGENTS.md`); external directories are granted per invocation with `--add-dir` / `/add-dir`. The `/tmp/pvyai-perfopt` binary is a **test build only**. Sandbox behavior stays untouched.

**Resolution path:** the PVYai Custom Catalog resolves the `pvy-*` platform model IDs when the binary is started properly (repo root, `pvyai` provider): `pvyai-opengateway` → `FetchOpenGateway` (`internal/providermodelcatalog/remote.go:53-59`) → `IsCodingModel` reduction (`internal/providermodelcatalog/filter.go`) → model pickers (`internal/tui/picker.go:293,470`, `internal/tui/command_center.go:627`). The catalog lists **only the coder / architect / lead-dev role families — application design** (§2.5). No registry code change is needed; **E1 (registry passthrough) is dropped**.

**Binding path for specialists:** platform IDs pass through the openai-compatible provider as opaque strings (`internal/providercatalog/catalog.go:93`). Frontmatter `model:` stays **omitted by design** — manifest validation resolves against the static `modelregistry` (`internal/specialist/manifest.go:255`), which does not know platform IDs. Binding happens at launch:

| Option | Description | Recommendation |
| --- | --- | --- |
| **E2** | Bind models via swarm `Definition.Model` (bypasses registry validation today — `internal/swarm/definition.go:30-44`), IDs taken from the Custom Catalog | **Recommended** — no registry/sandbox code change |
| E3 | Leave `model:` empty and inherit the session model; launch each stream with `--model` explicitly | Workable fallback, weakest provenance |

> Note: today the swarm roster is hardcoded (`teammate`, `subagent`); named agent types above require registering `swarm.Definition`s (`internal/swarm/definition.go`, wired in `internal/cli/app.go:979-988`) — or project specialists launched via `Task`. This is included in the plan as **Phase E** (self-dogfooding: the team configures its own tooling).

---

### 2.4 Model binding contract (platform model-config export, 2026-11-08)

| Platform ID | Base | num_ctx | max_tokens | Temperature | Top-p | Reasoning | FC | Vision |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `pvy-architect-x` | glm-5.3 | 1 048 576 | 128 000 | 1.0 | 0.95 | think: high | native | yes |
| `pvy-lead-dev-x` | glm-5.3 | 1 048 576 | 128 000 | 1.0 | 0.95 | think: high | native | yes |
| `pvy-senior-lead-dev-5` | glm-5.3-flash | 1 048 576 | 32 000 | — | — | none (platform default) | native | no |
| `pvy-coder-flash-2` | deepseek-flash | 1 000 000 | 32 768 | — | — | reasoning_effort: low | native | no |

Contract rules:

1. Exact IDs: the roster binds `pvy-lead-dev-x` and `pvy-coder-flash-2` — catalog bindings must use these spellings. (The CSV's `pvy-coder-flash` is glm-5.3-flash with think: high, NOT DeepSeek; binding uses `pvy-coder-flash-2` (deepseek-flash, reasoning low) — CTO to confirm.)
2. `reasoningEffort` in a manifest is passed as `--reasoning-effort` (internal/specialist/exec.go:702-707). Launch binding must preserve the platform params_json spelling per model family — `think` for GLM 5.3 / GLM-Flash, `reasoning_effort` for deepseek-flash — and manifest values must never override the platform's sampling params (temperature, top_p).
3. All four models use native function calling — bindings can assume native tool-calling; no prompt-based FC fallback needed.
4. Vision split: only pvy-architect-x and pvy-lead-dev-x have vision → screenshot-based TUI soak evidence (P0/P3) goes only to those roles; coder and senior-lead tasks stay text-only.
5. Context budget: ~1M tokens on all four models; compaction thresholds must derive from per-model num_ctx (1 048 576 vs 1 000 000), not fixed constants — verify internal/agent/context_measurement.go reads num_ctx from model config.
6. Output budget: GLM 5.3 roles may emit up to 128 000 tokens/response; pvy-senior-lead-dev-5 (32 000) and pvy-coder-flash-2 (32 768) are capped → P1/P2 work items must be chunked to fit one response.
7. Sampling is T 1.0 / top_p 0.95 on GLM 5.3 (non-deterministic) → P3 regression gates must pin sampling params; the architect's `pvysentinel_qlty_threeshold: 68` becomes the code-review gate score for P0/P1 acceptance.

> Catalog hygiene (2026-11-08): `pvy-coder-qw36-deep` (PVY Coder Deep) is **not active** on the platform and has been removed from the model-config export; it must not be bound by any role.

### 2.5 Tag-based model reduction (Coder / Lead / Architect only)

The platform model-config export is a **reduction by tags**: only models tagged for the
coding roles (coder / lead / architect) are exposed, so **non-coding agents** (image gen,
TTS, embeddings, deep-research, moderation, rerank, …) can never be bound by a coding
specialist. Enforcement lives in `internal/providermodelcatalog/filter.go`:

| Guard | Mechanism | Effect |
| --- | --- | --- |
| `hasCodingTag(model.Tags)` | tag allowlist: `agentic, chat, code, coder, coding, instruct, reasoning, tools` | keeps coding-capable models |
| `IsKnownNonCodingModelID(id)` | ID blocklist: `audio, dall-e, deep-research, embedding, image, moderation, realtime, rerank, sora, speech, transcribe, translate, tts, whisper` | drops non-coding agents even if tagged |
| `LooksLikeCodingModelID(id/desc)` | family-name heuristic (claude/glm/qwen/deepseek/coder/pvy/…) | fallback for untagged models |

Contract rules (binding):

8. **Role tags are mandatory on every roster model.** Each bound model config must carry at least one coding tag from the allowlist above plus its role tag (`coder`, `lead`, or `architect`). The three roster bindings map as: `pvy-coder-flash-2` → `coder, coding, tools`; `pvy-lead-dev-x` / `pvy-senior-lead-dev-5` → `lead, coding, agentic`; architect model(s) → `architect, coding, reasoning`. The reduction keeps all of them in the export.
9. **Non-coding exclusion is hard.** A specialist manifest or swarm `Definition.Model` must never resolve to a model whose ID matches the non-coding blocklist — the filter rejects it before the catalog is built, so a mis-tagged binding would silently fall back to "inherit" or fail validation. Bindings (E2/E3) must stay inside the catalog surface — only IDs the Custom Catalog lists (the coder / architect / lead-dev families, application design) may be bound.
10. **Reduction is by tags, not by naming.** Do not rely on `pvy-*` naming to imply coding capability — the ID heuristic is only a fallback. When adding models to the export for this team, tag them explicitly; untagged models may be dropped or misclassified. This is as Taks for PVYai WebUI, to add tag. „orchestrator, researcher and agentic shall be added to allowed „tags“ for PVYai Coder.&#32;

---

## 3. Agent & Loop Configuration (draft)

### 3.1 `config.json` (delta)

```json
{
  "activeProvider": "pvyai",
  "maxTurns": 24,
  "swarm": { "maxTeamSize": 5 },
  "tools": { "deferThreshold": 8 }
}
```

> Tag reduction note: `FileConfig` (`internal/config/types.go:236`) has **no** model/tag
> allowlist field — the Coder/Lead/Architect reduction is enforced catalog-side
> (`internal/providermodelcatalog/filter.go`), so `config.json` needs no model filter.
> Specialist manifests carry no `model:` line (inherit session model), and everything the
> picker offers has already passed `IsCodingModel`.
> **Config-reconciliation note (approved 2026-11-08 — Phase E entry criterion):**
> `.pvyai/config.json` (delivered) carries `maxTurns: 200` and `swarm.maxTeamSize: 4`
> (autonomous-run values) while the §3.1 delta above specifies `24` and `5`. At Phase E
> entry, reconcile field-by-field — disputed keys: `maxTurns` (200 vs 24) and
> `swarm.maxTeamSize` (4 vs 5); `tools.deferThreshold` is uncontested — and commit the
> merged project config with the winning value per key recorded in the commit message.

### 3.2 Specialist manifests (`.pvyai/specialists/*.md`, project scope)

> `model:` frontmatter is omitted **by design** — manifest validation resolves against the static registry (`internal/specialist/manifest.go:255`) and does not know platform IDs. Per-role models bind at launch through the Custom Catalog (swarm `Definition.Model` / `--model`): `pvy-lead-dev-x`, `pvy-senior-lead-dev-5`, `pvy-coder-flash-2` (IDs and parameters in §2.4).

`pvy-lead-dev-x.md`:

```markdown
---
description: Lead developer for performance-opt. Owns TUI resilience (P0) and integration.
reasoningEffort: high
tools: read-only, edit, execute, plan
---
You own P0 (TUI staleness fix) on branch performance-opt.
Acceptance: 30-min idle soak test passes; no blocking program. Send on the
event path; coalescer never forwards under lock; PR watcher bounded.
Run `make lint && make test` before reporting done. Report file:line per change.
```

`pvy-senior-lead-dev-5.md`:

```markdown
---
description: Senior lead dev (fast model). Owns P1 render pipeline and reviews coder output.
tools: read-only, edit, execute, plan
---
You own P1 (render performance items #1-#4). Keep diffs minimal and scoped.
Verify each coder change compiles (`go build ./...`) and keeps `go test ./internal/tui/...` green.
```

`pvy-coder-flash-2.md`:

```markdown
---
description: Fast coder for mechanical, well-specified refactors.
reasoningEffort: low
tools: read-only, edit, execute
---
Implement only the assigned item, no drive-by refactors, match surrounding style.
Every change must compile and keep tests green. No new dependencies.
```

### 3.3 Loop configuration — built-in cron (by application design, verified 2026-11-08)

Loops are the **built-in cron system** (`internal/cron/`, `internal/cli/cron*.go`) — foreground,
file-backed, no daemon. The perf-opt autonomous loop is a cron job whose prompt lives in the repo.

**Cron job registration** (state lives in the user data dir, not the repo):

```bash
pvyai cron add "0 */2 * * *" \
  --cwd /Volumes/Data/Repos/ollama-plus-openweb-ui/coding-agent \
  --model pvy-lead-dev-x
```

> Delivered 2026-11-08: job `20261108-161202` registered (active, `0 */2 * * *`, prompt resolved
> from `.pvyai/loop.md`). It fires only while a runner is up: `pvyai cron run` (tmux/launchd) or
> `pvyai cron run --once` from system cron.

| Aspect | Design (code-verified) |
| --- | --- |
| Job storage | `$XDG_DATA_HOME/pvyai/cron` or `~/.local/share/pvyai/cron`; one dir per job with `metadata.json` + `runs.jsonl` (`internal/cron/store.go:DefaultRoot`) — **runtime state, per machine, not committed** |
| Loop prompt | `loop.md` resolution order: `<cwd>/.pvyai/loop.md` → `<cwd>/.agents/loop.md` → `~/.pvyai/loop.md` → `~/.agents/loop.md` → built-in fallback (`internal/cron/loopprompt.go:ResolveLoopPrompt`). ≤ 25 000 bytes, symlinks skipped. **The repo therefore carries `.pvyai/loop.md`** |
| Schedule | 5-field Vixie cron; DST gap/collapse guards; bounded 9-year search |
| Runner | `pvyai cron run` (foreground, 30 s tick) under tmux/launchd/systemd, or `pvyai cron run --once` driven by system cron/launchd/K8s CronJob (recommended for reliability) |
| Fire semantics | Atomic per-job claim (lock + stale-reclaim): `NextRunAt` advances before the run; two runners can't double-fire. Default: overdue jobs skip the backlog; `--catch-up` fires them |
| Job fields | `expr`, `prompt` (empty → `loop.md`), `--cwd`, `--model`, `--recipe`, `--run-now`; `pause`/`resume`/`rm`/`list` for management |
| Execution | Each fire runs `pvyai exec --output-format stream-json --session-title cron:<jobID>` (+ `--cwd`/`--model` when set) |
| Observability | `runs.jsonl` per job (`RunRecord`: jobId, at, exitCode, sessionTitle, error) — loop monitoring reads this |
| Best practice | Single-step loop prompts ("one work item per cycle"), `--cwd` always explicit, prompt in repo (version-controlled), no secrets in prompts |

> Note: the *swarm* scheduler (`internal/swarm/scheduler.go`) is separate — cron fires the
> orchestrator session; team members/streams are orchestrated inside that session (§2.2).

#### Runner supervision — decision: launchd LaunchAgent over tmux (agreed 2026-11-08)

The runner (`pvyai cron run`, foreground 30 s ticker) needs a supervisor. Decision: a
**launchd LaunchAgent** wrapping `pvyai cron run`; tmux stays a debugging convenience only.

| Dimension | tmux (foreground pane) | launchd LaunchAgent (**chosen**) |
| --- | --- | --- |
| Supervision | none — crash = scheduler gone silently; pane shows a shell | `KeepAlive` restarts on crash/exit; `ThrottleInterval` caps restart storms |
| Reboot / logout | gone until manually restarted | `RunAtLoad` restores it automatically |
| Logs | scrollback, lost on buffer overflow | `StandardOutPath`/`StandardErrorPath` files, rotated externally |
| Resource use | full tmux server + detached session | bare process, no terminal |
| Visibility | `tmux attach` (interactive, easy) | `launchctl list` / TUI loops panel (Phase P4) |
| Failure mode to watch | silent death (the exact class of bug P0 fixes) | restart throttle loop → alert on `runs.jsonl` gap > 2× interval |

LaunchAgent sketch (`~/Library/LaunchAgents/ch.pvyai.cron-runner.plist`):

```xml
<key>Label</key><string>ch.pvyai.cron-runner</string>
<key>ProgramArguments</key>
<array><string>/usr/local/bin/pvyai</string><string>cron</string><string>run</string></array>
<key>WorkingDirectory</key><string>/Volumes/Data/Repos/ollama-plus-openweb-ui/coding-agent</string>
<key>KeepAlive</key><true/>
<key>ThrottleInterval</key><integer>30</integer>
<key>RunAtLoad</key><true/>
<key>StandardOutPath</key><string>/Users/qtt/.local/share/pvyai/cron/runner.out.log</string>
<key>StandardErrorPath</key><string>/Users/qtt/.local/share/pvyai/cron/runner.err.log</string>
```

Implications of the choice (accepted):

- **Ticker semantics unchanged**: launchd supervises the *process*; the built-in cron
    store still owns scheduling, locking and reclaim (`internal/cron/lock.go`). Do **not**
    switch to `--once` + `StartInterval` — the 30 s ticker + atomic claim already prevents
    double-fires and keeps fire latency ≤ 30 s. `--once` remains the fallback for
    system-cron/K8s environments without launchd.
- **`ThrottleInterval: 30`** matches one tick, so a crash-loop can never fire jobs faster
    than the ticker would.
- **Out of scope / untouched**: sandbox behaviour, system-wide daemon (LaunchAgent is
    per-user), and Windows/Linux supervision (systemd unit is the equivalent there).

### 3.5 TUI loop wiring (agreed 2026-11-08) — Phase P4

Make the loop system visible and manageable from the TUI instead of raw `launchctl`/`cron`
commands. Design:

- **Loops panel** (command menu entry, e.g. `/loops`): lists jobs from the cron store
    (`metadata.json` per job) with expr, next fire, paused/active, and the last N runs from
    `runs.jsonl` (exit code, duration, session title).
- **Runner health indicator**: derived from `runs.jsonl` recency + the runner log —
    `running` / `stale (no fire > 2× interval)` / `stopped`.
- **Actions**: `pause` / `resume` / `run-now` per job — thin wrappers over the existing
    `internal/cron` store APIs (`pause`/`resume`/`run-now` already exist CLI-side); no new
    scheduling logic in the TUI.
- **Non-blocking**: the panel reads the file-backed store on demand (and on a slow ticker
    while open); it must never block the event loop — reuses the P0 non-blocking patterns.

### 3.4 Swarm runtime knobs

| Knob | Value | Rationale |
| --- | --- | --- |
| Orchestrator `maxTurns` | 24 | Long refactor phases with gate reviews |
| Parallel read batches | 8 (`parallel_tools.go:21`) — unchanged | Already optimal |
| Member autonomy | `member` (coders), `high` (leads) | Coder Flash stays gated on writes outside assigned files |
| Compaction | default proactive | Threshold = 75% of per-model num_ctx (~786k of 1 048 576 for GLM roles, 750k for coder) — verify context_measurement.go honors per-model num_ctx instead of a fixed constant |
| Output chunking | Work items sized to fit 32 000-token responses | Caps from platform max_tokens (pvy-senior-lead-dev-5: 32 000; pvy-coder-flash-2: 32 768; GLM 5.3 roles: 128 000) |
| `swarm_collect` wait | 600 s default | Enough for compile+test cycles |
| Worktree isolation | on for all parallel streams | Disjoint `internal/tui/` files per stream |

### 3.6 Skill layer — `spec-plan-fulfillment` (delivered 2026-11-08)

The project skill `.pvyai/skills/spec-plan-fulfillment/SKILL.md` (project skill root,
loader-validated 2026-11-08) encodes the per-cycle fulfilment workflow for team agents:
orient → pick **one** work item → delegate per §2.2 topology → gate-verify → report +
commit. It is the enforcement checklist; **this spec and the specialist manifests (§3.2)
remain the single source of truth** for scope, owners and acceptance gates.

---

## 4. Execution Plan

### Phase E — Enablement (E2 swarm definitions, catalog-bound models) — *first, small (Done, verify)*

- Owner: `pvy-lead-dev-x` + Architect review
- Deliverable: named swarm definitions registered (`internal/swarm/definition.go`, wired in `internal/cli/app.go:979-988`) with `Model` bound to Custom Catalog IDs; launch convention documented (`--model` / `Definition.Model`). **No registry and no sandbox code changes** — E1 dropped, the Custom Catalog resolves.
- Acceptance: a member launched with `Definition.Model: pvy-coder-flash-2` runs on that catalog model (session model report shows `pvy-coder-flash-2`); specialist manifests load with frontmatter `model:` omitted

### Phase P0 — TUI staleness fix (the user-visible bug) — *highest priority*

- Owner: `pvy-lead-dev-x`, helpers `pvy-coder-flash-2`
- Work items:
    1. Non-blocking `forward`/`program.Send` with bounded queue + drop-with-visible-flag (`run.go:28-35`)
    2. Move `c.forward` outside `c.mu` in `coalesce.go`
    3. 30 s keepalive watchdog: force frame + re-arm dead blink chain (`model.go`)
    4. Bound PR watcher: single reusable worker, hard context deadline, no goroutine-per-tick (`pr_status.go:101-176`)
- Acceptance: **30-min idle soak** (screen updates within 1 frame after idle; keystroke latency < 50 ms after 30 min idle); `go test ./internal/tui/... ./internal/agent/...` green
- Risk: Bubble Tea v2.0.7 upstream event-loop stall (residual uncertainty) — mitigation: watchdog (item 3) self-heals regardless of upstream cause; upgrade evaluation as follow-up

### Phase P1 — Render pipeline performance (items #1–#4)

- Owner: `pvy-senior-lead-dev-5`, items #1/#4/#6/#7 → `pvy-coder-flash-2`
- Work items: identity-keyed render cache (replace content-key building); incremental streaming-markdown block rendering (render closed blocks once); single body-items/metrics pass per frame; incremental `rowContext`
- Acceptance: benchmarks `BenchmarkTranscriptBodyItems` / `BenchmarkStreamingMarkdown` (100/1000-row fixtures) show **≥5×** reduction in per-frame allocs on long sessions; keystroke→render < 16 ms at 500 rows

### Phase P2 — Startup & I/O (items #5–#8)

- Owner: `pvy-coder-flash-2`, review `pvy-senior-lead-dev-5`
- Work items: cache repomap/workspaceindex snapshot per cwd (mtime/TTL invalidation, async rebuild); single `streamingTextString` conversion per frame; render-cache micro-fixes
- Acceptance: TTFT on repeat runs improves measurably (perfbench `MeasureFirstOutput`), no behavior change

### Phase P3 — Benchmarks & regression gates

- Owner: `pvy-bench-qa`
- Work items: Go micro-benchmarks for the render pipeline; wire into CI as regression thresholds; pprof before/after report for P1
- Acceptance: CI gate blocks >10% regressions; before/after numbers documented in `Sessions/`

### Phase P4 — TUI loops panel (§3.5) — *after P3, nice-to-have*

- Owner: `pvy-lead-dev-x`, helpers `pvy-coder-flash-2`
- Work items: `/loops` panel (job list + `runs.jsonl` history), runner health indicator, pause/resume/run-now actions over `internal/cron` store APIs
- Acceptance: panel renders and acts on jobs without blocking the event loop (P0 patterns); unit tests over a fixture store
- Depends on: P0 (non-blocking event path)

### Sequencing

```
E → P0 → (P1 ∥ P2) → P3 → P4
     └─ Architect gate review after P0 and after P1
```

---

## 5. Risks

| Risk | Mitigation |
| --- | --- |
| Bubble Tea v2.0.7 upstream stall not fully confirmed | Watchdog self-heals regardless; upstream issue filed if reproduced |
| P0 fixes change message delivery ordering | Bounded queue preserves per-sender order; soak + unit tests for coalescer |
| Frontmatter `model:` rejects platform IDs (static registry) | By design — bind via Custom Catalog (E2 `Definition.Model` / `--model`); frontmatter stays empty |
| Parallel TUI edits collide | Disjoint file ownership per stream; worktrees |
| Perf changes regress rendering correctness | Snapshot/visual tests in `go test ./internal/tui/...` run per item |

---

## 6. CTO decisions — all resolved (CTO approved 2026-11-08)

1. ~~Model IDs~~ RESOLVED from platform model-config export (2026-11-08): `pvy-architect-x`, `pvy-lead-dev-x`, `pvy-senior-lead-dev-5`, `pvy-coder-flash-2`. Coder Flash binding CONFIRMED (CTO, 2026-11-08): `pvy-coder-flash-2` (deepseek-flash, reasoning_effort: low).
2. ~~Reasoning defaults~~ RESOLVED: GLM 5.3 roles run think: high (T 1.0, top_p 0.95); `pvy-senior-lead-dev-5` has no explicit preset (inherit platform default).
3. ~~Bubble Tea upgrade~~ RESOLVED (CTO, 2026-11-08): strictly watchdog-only in P0; the v2.0.7 → newer upgrade evaluation remains a follow-up (per §5 risk mitigation).
4. ~~Registry passthrough~~ SUPERSEDED (2026-11-08): dropped per CTO — the PVYai Custom Catalog resolves platform IDs; Phase E is E2-only (swarm definitions).
5. ~~Gate score~~ RESOLVED (CTO, 2026-11-08): `pvysentinel_qlty_threeshold: 68` (on `pvy-architect-x` / `pvy-architect-m`) is the intended code-review gate score for P0/P1 acceptance (§2.4 note 7).

---

*CTO-approved 2026-11-08 (v0.6) — execution green-lit: `E → P0 → (P1 ∥ P2) → P3 → P4`. Branch `performance-opt` (from `pvy-develop`).*
