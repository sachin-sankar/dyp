---
sidebar_position: 5
---

# Command Reference

## Commands

### `dyp`

The main command launches an interactive TUI to select and render prompts.

```bash
dyp [flags]
```

#### Flags

| Flag | Description | Default |
|------|-------------|---------|
| `--prompts <path>` | Custom prompts directory | `~/.prompts` |
| `--verbose` | Enable debug logging | `false` |
| `--help` | Show help | - |

#### Examples

```bash
# Run interactive prompt selector
dyp

# Run with custom prompts directory
dyp --prompts /path/to/my/prompts

# Enable verbose logging
dyp --verbose
```

---

### `dyp list`

Lists all available prompts in the prompts directory.

```bash
dyp list [flags]
```

#### Examples

```bash
# List all prompts
dyp list

# List prompts from custom directory
dyp list --prompts /path/to/prompts
```

#### Output

```
┌─────────────────────────────────┬────────────────────────────────┐
│ Title                           │ File Path                      │
├─────────────────────────────────┼────────────────────────────────┤
│ Generate Business Proposal      │ /home/user/.prompts/...        │
│ Podcast Episode Script         │ /home/user/.prompts/...        │
│ Meeting Agenda Template        │ /home/user/.prompts/...        │
└─────────────────────────────────┴────────────────────────────────┘
```

---

### `dyp render`

Renders a specific prompt file directly by filename (without interactive selection).

```bash
dyp render <filename> [flags]
```

#### Arguments

| Argument | Description |
|----------|-------------|
| `<filename>` | The filename of the prompt to render |

#### Examples

```bash
# Render a specific prompt
dyp render business-proposal.md

# Render from custom directory
dyp render meeting-agenda.md --prompts /path/to/prompts
```

---

### `dyp sink`

Renders a prompt and sends it to an AI chat sink. Opens the rendered prompt in your default browser at the selected AI platform.

```bash
dyp sink <sink> [flags]
```

#### Arguments

| Argument | Description |
|----------|-------------|
| `<sink>` | The AI sink to send the prompt to (required). Supported sinks are listed below. |

#### Flags

| Flag | Short | Description | Default |
|------|-------|-------------|---------|
| `--prompt-file` | `-p` | Render a specific prompt file directly | Interactive selection |

#### Examples

```bash
# Interactive prompt selection, then send to ChatGPT
dyp sink chatgpt

# Send to Perplexity
dyp sink perplexity

# Send to Claude
dyp sink claude

# Render a specific prompt file and send to ChatGPT
dyp sink chatgpt --prompt-file business-proposal.md

# Use custom prompts directory
dyp sink perplexity --prompts /path/to/prompts

# Send to multiple sinks at once
dyp sink chatgpt perplexity claude
```

#### Supported Web based Sinks
| Sink | URL |
|------|-----|
| `chatgpt` | https://chatgpt.com |
| `perplexity` | https://www.perplexity.ai |
| `claude` | https://claude.ai |


#### Supported Terminal based Sinks
| Sink | Command |
|------|-----|
| `opencode` | opencode --prompt |
| `gemini` | gemini -i |
| `copilot` | copilot -i |
---

## Environment

- `HOME` - Used to determine default prompts directory (`~/.prompts`)
- `PATH` - Where the binary should be installed

## Exit Codes

| Code | Description |
|------|-------------|
| 0 | Success |
| 1 | Error (invalid arguments, missing prompts directory, etc.) |
