---
title: "Generate Project Status Report"
---

Create a project status report for {{What project name?|text}}.

Project info:
- Project manager: {{Who is leading this?|text}}
- Reporting period: {{What time period?|text}}
- Overall status: {{How is it going?|select<on_track,at_risk,delayed,completed>}}

Metrics:
- Budget status: {{What is the budget situation?|select<on_budget,over_budget,slightly_over,under_budget>}}
- Completion percentage: {{How complete?|select<0-25%,25-50%,50-75%,75-100%>}}
- Team size: {{How many people?|select<1-5,6-10,11-20,20+>}}

Include:
- Executive summary
- Key accomplishments this period
- Upcoming milestones
- Risks and issues
- Resource needs
- Next steps
- Visual: {{What visuals to include?|multiselect<gantt_chart,timeline,burndown_chart,pie_chart>}}