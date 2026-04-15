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

## Environment

- `HOME` - Used to determine default prompts directory (`~/.prompts`)
- `PATH` - Where the binary should be installed

## Exit Codes

| Code | Description |
|------|-------------|
| 0 | Success |
| 1 | Error (invalid arguments, missing prompts directory, etc.) |