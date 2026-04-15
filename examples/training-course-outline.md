---
title: "Generate Training Course Outline"
---

Create a training course outline for teaching {{What topic?|text}}.

Course details:
- Duration: {{How long is the course?|select<1_day,1_week,2_weeks,1_month>}}
- Delivery format: {{How is it delivered?|select<in_person,virtual,hybrid,self_paced>}}
- Difficulty level: {{What skill level?|select<beginner,intermediate,advanced,all_levels>}}

Audience:
- Who is this for?: {{Target learners?|text}}
- Prior knowledge: {{What prerequisites?|text}}

Include:
- Course objectives
- Module breakdown (how many?: {{How many modules?|select<3,5,7,10>}})
- Learning outcomes per module
- Assessment methods: {{How to assess?|multiselect<quiz,assignment,project,exam,presentation>}}
- Required materials
- Certification details