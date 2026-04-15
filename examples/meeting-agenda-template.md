---
title: "Generate Meeting Agenda Template"
---

Create a meeting agenda template for {{What type of meeting?|select<team_sync,one_on_one,project_review,board_meeting,client_meeting>}}.

Meeting details:
- Duration: {{How long is the meeting?|select<15min,30min,45min,1hour,2hours>}}
- Number of attendees: {{How many people?|select<2-5,6-10,11-20,20+>}}
- Meeting style: {{What style?|select<informal,formal,workshop,brainstorm>}}

Include sections for:
- Previous action items review
- Topics to discuss (prioritized)
- Time allocation per topic
- New action items
- Next meeting date

Also include {{What additional elements?|multiselect<ice_breaker,round_robin,timer,break,recording>}}