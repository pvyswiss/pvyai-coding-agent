# PVYAI GitHub Action

Run PVYAI headlessly inside a GitHub workflow. The action is a thin wrapper around
`pvyai exec`: it installs a pinned PVYAI release, runs your prompt in the checked-out
repository, surfaces PVYAI's exit code as the step status, captures the output as a
file, and can optionally post a summary to the triggering pull request or to Slack.

PVYAI is model- and provider-agnostic, and so is this action: **you** choose the
provider and supply the API key. Nothing is hardcoded to any single provider.

## Quick start

```yaml
# .github/workflows/pvyai.yml
name: PVYAI
on:
  workflow_dispatch:
    inputs:
      task:
        description: What should PVYAI do?
        required: true

permissions:
  contents: write

jobs:
  run:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: pvyswiss/pvyai@v1
        with:
          prompt: ${{ inputs.task }}
          provider: openai
          api-key-env: OPENAI_API_KEY
          api-key: ${{ secrets.OPENAI_API_KEY }}
          model: gpt-4.1
```

> The `provider`, `api-key-env`, and `api-key` trio is how you stay
> provider-neutral: `api-key-env` is the environment variable name the chosen
> provider reads its key from (for example `ANTHROPIC_API_KEY`, `GEMINI_API_KEY`,
> `OPENROUTER_API_KEY`), and `api-key` is the secret value. The action exports
> that variable only for the PVYAI step and never prints it. If your repository
> already commits a `.pvyai/config.json` with an active provider, you can omit
> `provider`.

## Inputs

| Input | Required | Default | Description |
| --- | --- | --- | --- |
| `prompt` | one of `prompt`/`prompt-file` | `""` | The instruction for PVYAI to execute. |
| `prompt-file` | one of `prompt`/`prompt-file` | `""` | Path (relative to `working-directory`) to a file whose contents are the prompt. |
| `provider` | no | `""` | Provider id to activate (e.g. `openai`, `anthropic`, `gemini`, `ollama`, or any compatible endpoint). |
| `api-key` | no | `""` | The provider API key. Pass from a secret; exported only for the PVYAI step and never logged. |
| `api-key-env` | no | `""` | Env var name the provider reads its key from (e.g. `OPENAI_API_KEY`). Exported only for the PVYAI step when used with `api-key`. |
| `model` | no | `""` | Model id. Defaults to the resolved provider's default. |
| `mode` | no | `""` | Run mode (`pvyai exec --mode`), e.g. `smart`, `deep`, `fast`. |
| `auto` | no | `low` | Autonomy ceiling (`pvyai exec --auto`): `low`, `medium`, or `high`. Conservative by default. |
| `self-correct` | no | `false` | Allow mid-run model escalation (`pvyai exec --allow-escalation`). |
| `add-dir` | no | `""` | Newline-/comma-separated extra write roots (`pvyai exec --add-dir`). |
| `worktree` | no | `false` | Run in an isolated git worktree (`pvyai exec --worktree`). |
| `output-format` | no | `stream-json` | `text`, `json`, or `stream-json`. Captured to a file. |
| `post-to` | no | `none` | `pr-comment`, `slack`, or `none`. Where to post a summary after the run. |
| `slack-webhook-url` | no | `""` | Slack incoming-webhook (or generic webhook) URL for `post-to: slack`. Pass from a secret. |
| `github-token` | no | `${{ github.token }}` | Token used to post a PR comment. Requires `pull-requests: write`. |
| `working-directory` | no | `${{ github.workspace }}` | Directory to run PVYAI in. |
| `pvyai-version` | no | (action ref → `latest`) | PVYAI release version/tag to install, e.g. `v1.2.3` or `latest`. |
| `pvyai-repo` | no | `pvyswiss/pvyai` | Repository to install the PVYAI release from. |

## Outputs

| Output | Description |
| --- | --- |
| `exit-code` | PVYAI's exit code (`0` success, `2` usage, `3` provider, non-zero otherwise). |
| `output-file` | Path to the captured PVYAI stdout (the raw `output-format` stream). |
| `summary` | A short, single-line summary parsed from the run, when available. |

The step **fails when PVYAI returns a non-zero exit code**, so a failed run fails
the job by default. Use `continue-on-error: true` on the step (and read
`exit-code`) if you want to handle failures yourself.

## Examples

### Run PVYAI on every issue labeled `pvyai`

```yaml
name: PVYAI issue triage
on:
  issues:
    types: [labeled]

permissions:
  contents: write
  issues: write
  pull-requests: write

jobs:
  triage:
    if: github.event.label.name == 'pvyai'
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: pvyswiss/pvyai@v1
        with:
          prompt: |
            Investigate this issue and propose a fix.

            Title: ${{ github.event.issue.title }}

            ${{ github.event.issue.body }}
          provider: anthropic
          api-key-env: ANTHROPIC_API_KEY
          api-key: ${{ secrets.ANTHROPIC_API_KEY }}
          auto: low
          post-to: slack
          slack-webhook-url: ${{ secrets.PVYAI_SLACK_WEBHOOK_URL }}
```

### Nightly dependency-upgrade PR

```yaml
name: PVYAI nightly deps
on:
  schedule:
    - cron: "0 6 * * 1" # Mondays 06:00 UTC

permissions:
  contents: write
  pull-requests: write

jobs:
  deps:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: pvyswiss/pvyai@v1
        id: pvyai
        with:
          prompt-file: .github/pvyai/upgrade-deps.md
          provider: openai
          api-key-env: OPENAI_API_KEY
          api-key: ${{ secrets.OPENAI_API_KEY }}
          worktree: true
          auto: medium
      - name: Open pull request
        run: |
          git checkout -b pvyai/deps-$(date +%Y%m%d)
          git commit -am "chore(deps): nightly upgrade via PVYAI" || exit 0
          git push -u origin HEAD
          gh pr create --fill --label dependencies
        env:
          GH_TOKEN: ${{ github.token }}
      - name: Upload PVYAI output
        if: always()
        uses: actions/upload-artifact@v4
        with:
          name: pvyai-output
          path: ${{ steps.pvyai.outputs.output-file }}
          if-no-files-found: warn
```

## Security notes

- **Secrets are passed as Action secrets and never logged.** The action exports
  `api-key` (under `api-key-env`) only for the PVYAI step. `slack-webhook-url`
  is passed only to the Slack post step. They are never echoed or persisted to
  later workflow steps, and PVYAI redacts secret-shaped strings from its own
  output.
- **The sandbox is always active.** This action never passes
  `--skip-permissions-unsafe`, so writes stay inside the checked-out repository
  (plus any roots you grant with `add-dir`). Unsafe mode is never enabled
  implicitly.
- **Autonomy defaults to `low`.** Raise it deliberately (`auto: medium`/`high`)
  only for tasks you trust to run unattended.
- **Least-privilege tokens.** Grant only the permissions the workflow needs
  (`contents: write` to edit files, `pull-requests: write` to comment). The
  default `GITHUB_TOKEN` is scoped to the repository.
- **Pin the action.** Reference a tag (`pvyswiss/pvyai@v1`) or a commit SHA so a
  workflow run uses a known PVYAI version. `pvyai-version` lets you pin the
  installed binary independently of the action ref.
- **Linux and macOS runners** are supported; Windows runners are rejected with a
  clear error.

## Slack / webhook notifier

The action's `post-to: slack` step sends a one-line summary to a Slack incoming
webhook after the run. PVYAI also has a built-in webhook notifier sink
(`internal/notify`) that an unattended run can use to report
"finished / needs input / verify failed after N retries" to Slack or any generic
webhook:

- Configure the destination with the `PVYAI_SLACK_WEBHOOK_URL` environment variable
  (or settings). A blank URL disables the sink.
- The sink POSTs a JSON body `{ "text", "type", "message", "summary?", "links?" }`.
  The `text` field is what Slack renders; the structured fields carry the
  machine-readable detail.
- **Fail-soft:** a non-2xx response or a transport error is logged (redacted) and
  swallowed — a webhook problem never crashes the run.
- **Redaction:** the message, summary, links, and the webhook URL itself are run
  through PVYAI's redaction before being sent or logged, so tokens never leak.
- **Egress/proxy:** the notifier uses the default HTTP transport, which honors
  `HTTP_PROXY`/`HTTPS_PROXY` when a proxy is configured.