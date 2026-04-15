# DYP - Dynamically Render Prompts on the Fly

DYP is a CLI tool that renders interactive markdown prompts. It reads prompt template files, asks you questions interactively in the terminal, and outputs the filled-in content.

## The Problem

If you frequently use LLMs like ChatGPT or Claude, you probably have a collection of prompt templates for different tasks—job descriptions, meeting agendas, business proposals, etc.

The frustration: every time you use a template, you need to manually edit it to fill in your specific details. You open the markdown file, find each placeholder, replace it with your actual value, and only then can you copy-paste it to your LLM.

This manual editing is:
- **Repetitive** - Same edits every time
- **Error-prone** - Easy to miss a placeholder
- **Slow** - Context switching between editor and LLM

## The Solution

DYP automates this workflow. You write your prompt templates once in markdown, define what questions need answering, and DYP handles the rest:

1. Shows you a list of your prompts
2. Asks each question interactively in the terminal
3. Substitutes your answers into the template
4. Outputs the rendered result—ready to copy and use

## Installation

### Quick Install

```bash
go install github.com/sachin-sankar/dyp@latest
```

### From Source

```bash
git clone https://github.com/sachin-sankar/dyp.git
cd dyp
go build -o dist/dyp main.go
```

Move the binary to your PATH:

```bash
sudo mv dist/dyp /usr/local/bin/
```

> **Note**: Package manager-specific installation (brew, apt, etc.) coming soon.

## Quick Start

### 1. Create a Prompt Template

Create a markdown file in `~/.prompts/` (or any directory):

```markdown
---
title: "Generate Job Description"
---

Create a job description for a {{What job title?|text}} position.

Company details:
- Company name: {{What is the company name?|text}}
- Industry: {{What industry?|select<technology,healthcare,finance>}}
- Work style: {{What work arrangement?|select<remote,hybrid,onsite>}}

Include:
- Job summary
- Key responsibilities
- Required qualifications
```

### 2. Run DYP

```bash
dyp
```

You'll see an interactive prompt to select which template to use. DYP will then ask each question defined in your template, and finally output the rendered result.

### 3. Use the Output

Copy the output and paste it into your LLM.

Since DYP outputs plain text to stdout, you can easily pipe it to other tools:

```bash
# Pipe directly to another CLI tool
dyp | opencode

# Save to a file
dyp > output.md

# Copy to clipboard (macOS)
dyp | pbcopy

# Copy to clipboard (Linux)
dyp | xclip -selection clipboard
```

## Template Syntax

Variables use the format: `{{Question|Type<options>}}`

| Type | Syntax | Description |
|------|--------|-------------|
| Text | `{{What is your name?\|text}}` | Free-form text input |
| Input | `{{Enter your email?\|input}}` | Single-line input |
| Select | `{{Pick a category?\|select<a,b,c>}}` | Single choice from options |
| MultiSelect | `{{Choose tags?\|multiselect<a,b,c>}}` | Multiple choice from options |

### Example Variables

```markdown
{{What is the job title?|text}}
{{What industry?|select<technology,healthcare,finance,education>}}
{{How large is the company?|select<startup,small,medium,large,enterprise>}}
{{What work arrangement?|select<remote,hybrid,onsite>}}
{{Required skills?|multiselect<Python,JavaScript,Go,Rust,SQL>}}
```

## Commands

### `dyp` (default)

Launches the interactive prompt selector, then renders the chosen prompt.

### `dyp list`

Lists all available prompts in your prompts directory.

```bash
dyp list
```

### `dyp render <filename>`

Renders a specific prompt file directly (bypasses the selection menu).

```bash
dyp render job-description.md
```

### Options

- `--prompts <path>`: Specify a custom prompts directory (default: `~/.prompts`)
- `--verbose`: Enable debug logging

```bash
dyp --prompts /path/to/my/prompts --verbose
```

## Prompts Directory

By default, DYP looks for prompts in `~/.prompts/`. You can change this with the `--prompts` flag.

Organize your prompts any way you like:

```
~/.prompts/
├── job-description.md
├── meeting-agenda.md
├── business-proposal.md
└── writing/
    ├── blog-post.md
    └── social-post.md
```

## Included Examples

DYP comes with example templates in the `examples/` directory. Copy them to your prompts folder to try them out:

```bash
cp examples/*.md ~/.prompts/
```

Available examples:
- `job-description.md` - Generate job descriptions
- `meeting-agenda.md` - Create meeting agendas
- `business-proposal.md` - Write business proposals
- `project-status-report.md` - Project status updates
- `marketing-email-campaign.md` - Email campaign drafts
- `podcast-episode-script.md` - Podcast planning
- `social-media-content-calendar.md` - Social media planning
- `training-course-outline.md` - Course structure
- `product-launch-plan.md` - Product launch docs
- `academic-research-paper.md` - Research paper outlines
