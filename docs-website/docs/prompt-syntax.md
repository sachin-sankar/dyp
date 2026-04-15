---
sidebar_position: 3
---

# Prompt Syntax

Prompts are Markdown files with YAML frontmatter. The frontmatter defines the prompt title, and the body contains your template with placeholders.

## File Format

```markdown
---
title: "Your Prompt Title"
---

Your template content with {{placeholder|type<options>}} placeholders.
```

## Placeholder Syntax

```
{{VariableName|type<options>}}
```

| Component | Description |
|-----------|-------------|
| `VariableName` | The label shown to the user |
| `type` | Input type: `text`, `input`, `select`, `multiselect` |
| `options` | Comma-separated options (for select/multiselect) |

## Input Types

### Text Input

```markdown
{{Your Name|text}}
```

Renders as a multi-line text field.

### Input Field

```markdown
{{Your Email|input}}
```

Renders as a single-line input field.

### Select Dropdown

```markdown
{{Your Role|select<Developer,Designer,Manager>}}
```

Renders as a dropdown with predefined options.

### Multi-Select

```markdown
{{Your Skills|multiselect<Go,Python,React,TypeScript>}}
```

Renders as a multi-select list where users can pick multiple options.

## Example

```markdown
---
title: "Meeting Notes Generator"
---

# Meeting Notes

**Date**: {{Date|text}}
**Attendees**: {{Attendees|multiselect<Alice,Bob,Charlie,David>}}

## Agenda
{{Agenda|text}}

## Action Items
{{Action Items|text}}

## Next Steps
{{Next Steps|text}}
```

This creates an interactive form where users can fill in the date, select attendees from a list, and provide text for agenda, action items, and next steps.

## Rendering Variables

After collecting answers, DYP replaces each `{{placeholder}}` with the user's input in the final output.

```markdown
Input: {{Name|text}} → User enters "John" → Output: "John"
```

## Tips

- Keep variable names short and descriptive
- Use consistent naming across prompts
- Test your prompts by running `dyp render <filename>`