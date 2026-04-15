---
sidebar_position: 2
---

# Installation

## Prerequisites

- **Go 1.21+** - Required to build from source

## Install from Source

```bash
# Clone the repository
git clone https://github.com/sachin-sankar/dyp.git
cd dyp

# Build the binary
go build -o dyp .

# Move to your PATH
sudo mv dyp /usr/local/bin/
```

Or use `go install`:

```bash
go install github.com/sachin-sankar/dyp@latest
```

## Verify Installation

```bash
dyp --help
```

You should see:

```
Dynamically render prompts on the fly.

Usage:
  dyp [flags]
  dyp [command]

Available Commands:
  list      List available prompts.
  render    Render a prompt from filename directly.

Flags:
  --prompts string   Specify a custom directory path to look for prompts in. (default "$HOME/.prompts")
  --verbose         Run dyp in verbose mode to observe debug logs.
  --help            Show help

Use "dyp [command] --help" for more information about a command.
```

## Setup Prompts Directory

DYP looks for prompts in `~/.prompts` by default. Create it:

```bash
mkdir -p ~/.prompts
```

Add some example prompts or copy from the examples directory:

```bash
cp examples/*.md ~/.prompts/
```