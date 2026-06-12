+++
title = "Leaving the Plate Still Wobbling"
date = "2026-05-30"
description = "Eight weeks at HUL, from charter to friction to multi-agent orchestration, and what Feynman's summer job at Metaplast has to do with selling your own project back to yourself."
draft = false
+++

Richard Feynman spent a summer in college as the entire chemistry department of Metaplast Corporation. Their problem was getting metal plating to stick to plastic. His staff was himself.

Metaplast ran a full-page ad in *Modern Plastics* magazine. Chrome-bright photographs. A serious operation, to all appearances. A British lab saw the ad and shut down their own competing program. No point, they figured. Metaplast clearly had dozens of chemists, a whole department, a chief research chemist.

The whole department was a college student on summer break, fiddling until something stuck.

---

I arrived at HUL with a charter. A document that described what my eight weeks were for, what success looked like, what the business needed from the project. It was sensible. It was signed off. 

The charter assumed the problem was known. What I found in the first ten days was that the problem was still being negotiated. The business knew it wanted something around agentic AI. The business did not yet know what question that system was supposed to answer. So I wrote my own problem statement. You cannot build a solution to a problem that has not been defined, and waiting for someone else to define it would have consumed the internship without touching the ground.

---

The ground, when I touched it, pushed back.

Most of the early weeks went not into building but into getting permission to build: forms, approvals, the small administrative frictions that eat calendar days without showing up on any project plan. The project sat inert while I chased signatures. It was tedious rather than dramatic, but it was the part that determined whether anything happened at all.

---

What got built in the gaps was not the chatbot everyone expected.

The first impulse in a corporate AI project is to build a conversational interface: ask a question, get an answer, call it insight. The real problem was not that people lacked answers. It was that insight existed in fragments across systems, documents, dashboards, and heads, and no one had the time to synthesise it at the speed decisions actually get made.

[Aadit Palicha said at YC India 2026](https://www.youtube.com/watch?v=YKZCU0ynEbs) that you remove all laws of physics first, see what the art of the possible is, then add the constraints back. The chatbot was what you build when you leave the physics in place. I wanted to see what the system looked like when I removed them.

So I built a multi-agent orchestration system. Not one model pretending to know everything. Several specialised agents, each handling a slice of the pipeline such as data retrieval, synthesis, validation, narrative formatting. All coordinated to produce something a manager could use in a morning meeting without needing to know what a latent space is.

It was not elegant at first. It was held together by the digital equivalent of duct tape and optimism. The first territory summary it generated was wrong. I caught it because I happened to know the number was off. A manager relying on it would not have known until they stood in front of a shopkeeper who had never heard of the SKU. I rebuilt the retrieval agent, added a validation layer, and tried again. The second version worked. A rep in the field could get a synthesized view of their territory in seconds instead of hours. The gap between the data and the decision compressed.

---

Somewhere around week five, I realised the system was only half the product.

The other half was the story. No one funds a project they cannot narrate. No one adopts a tool they cannot describe to their own boss. I had spent weeks learning to build multi-agent pipelines. I now had to learn to build a hero's arc: the villain (the time cost of manual synthesis), the journey (the friction, the pivot, the build), the impact (decisions made faster, coverage made wider, hours returned to people who would rather be in the field than in a spreadsheet).

I was Metaplast running the ad. The fiddling was real. The plating was starting to stick. But the full-page spread in *Modern Plastics*; the narrative that made the fiddling legible to people who were not in the room while it happened; that was a separate skill, and I had to learn it fast.


---

The test came in a room with senior management.

I pitched it Y Combinator style: the problem, the solution, the traction, the ask. No forty-slide buildup. Just the arc, delivered standing up, ending with what the team would need after my internship ended to keep the plate spinning. The sentence that moved the conversation was not about agents or orchestration. It was: *"How future product managers need to be ready to get their hands dirty and not just be coordinators."*

They responded to the clarity, not the complexity.


---

In [The Name of the Bird](/blog/the-name-of-the-bird), I wrote that I had been handed the plate and was watching it spin. In [The Messy Reactor](/blog/the-messy-reactor), I wrote that the messy reactor was where the energy was. Both are still true.

The multi-agent system is running. The account managers are using it. The pipeline that synthesises insight at scale is no longer an idea. But the plate is still wobbling. There are edge cases I did not resolve. There are handoffs I designed but did not live long enough to stress-test. There is a conversation about production infrastructure that belongs to the team staying, not to me.

In [The Eye Level](/blog/the-eye-level), I wrote about the kirana owner who knows what moves on which days by instinct, and about the brand manager who wants that knowledge systematised. I do not know whether the system I built closes that gap or widens it. The territory summary is faster now. But speed is not the same as knowledge. The rep still knows things the model does not.

The Metaplast ad worked because the fiddling was real. I leave behind both: the system, and the story that helps the next person carry it.

The plate is still wobbling. I am leaving it that way on purpose.

---

*With this the blog is archived. I am back in XLRI doing marketing & finance specialisation, the live consulting project in Jamshedpur & agentic AI self-study on my own time, case competitions when they come up. The plan is to rotate through sales, marketing, finance, legal, tech, operations and warehousing at HUL next year, one lens at a time, and keep enough open problems loaded that the connections start finding me.*
