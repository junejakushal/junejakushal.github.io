+++
title = "Leaving the Plate Still Wobbling"
date = "2026-05-30"
description = "Eight weeks at HUL, from charter to blockade to multi-agent orchestration, and what Feynman's summer job at Metaplast has to do with selling your own project back to yourself."
draft = false
+++

Richard Feynman spent a summer in college as the entire chemistry department of Metaplast Corporation. Their problem was getting metal plating to stick to plastic. His staff was himself.

Metaplast ran a full-page ad in *Modern Plastics* magazine. Chrome-bright photographs. A serious operation, to all appearances. A British lab saw the ad and shut down their own competing program. No point, they figured. Metaplast clearly had dozens of chemists, a whole department, a chief research chemist.

The whole department was a college student on summer break, fiddling until something stuck.

---

I arrived at HUL with a charter. A document that described what my eight weeks were for, what success looked like, what the business needed from the project. It was sensible. It was signed off. 

The charter assumed the problem was known. What I found in the first ten days was that the problem was still being negotiated. The business knew it wanted something around agentic AI. The business did not yet know what question that system was supposed to answer. So I wrote my own problem statement. Not in rebellion. In necessity. You cannot build a solution to a problem that has not been defined, and waiting for someone else to define it would have consumed the internship without touching the ground.

---

The ground, when I touched it, pushed back.

InfoSec blocked my laptop. Not temporarily. Not as a formality. Blocked. The machine I was supposed to build on was rendered unusable by the same security architecture that was protecting the data I needed to access. I spent days in the kind of bureaucratic loop that does not appear in any internship handbook: filling forms, chasing approvals, watching calendar days disappear while the project sat inert. At one point I genuinely considered building the prototype on my personal laptop and screen-sharing it in. I did not. But the fact that it felt like a reasonable option tells you something about how stuck I was.

This is the part no one puts in the final presentation. It is also the part that determines whether the project happens at all.

---

What got built in the gaps was not the chatbot everyone expected.

The first impulse in a corporate AI project is to build a conversational interface: ask a question, get an answer, call it insight. That is a parlour trick. The real problem was not that people lacked answers. It was that insight existed in fragments across systems, documents, dashboards, and heads, and no one had the time to synthesise it at the speed decisions actually get made.

[Aadit Palicha said at YC India 2026](https://www.youtube.com/watch?v=YKZCU0ynEbs) that you remove all laws of physics first, see what the art of the possible is, then add the constraints back. The chatbot was what you build when you leave the physics in place. I wanted to see what the system looked like when I removed them.

So I built a multi-agent orchestration system. Not one model pretending to know everything. Several specialised agents, each handling a slice of the pipeline — data retrieval, synthesis, validation, narrative formatting — coordinated to produce something a manager could use in a morning meeting without needing to know what a latent space is.

It was not elegant at first. It was held together by the digital equivalent of tape and optimism. The first territory summary it generated was wrong. I caught it because I happened to know the number was off. A manager relying on it would not have known until they stood in front of a shopkeeper who had never heard of the SKU. I rebuilt the retrieval agent, added a validation layer, and tried again. The second version worked. A rep in the field could get a synthesized view of their territory in seconds instead of hours. The gap between the data and the decision compressed.

---

Somewhere around week five, I realised the system was only half the product.

The other half was the story. No one funds a project they cannot narrate. No one adopts a tool they cannot describe to their own boss. I had spent weeks learning to build multi-agent pipelines. I now had to learn to build a hero's arc: the villain (the time cost of manual synthesis), the journey (the blockade, the pivot, the build), the impact (decisions made faster, coverage made wider, hours returned to people who would rather be in the field than in a spreadsheet).

I was Metaplast running the ad. The fiddling was real. The plating was starting to stick. But the full-page spread in *Modern Plastics* — the narrative that made the fiddling legible to people who were not in the room while it happened — that was a separate skill, and I had to learn it fast.

There is an honest question underneath this that I am still sitting with. The British lab quit because they were deceived. The ad projected a scale Metaplast did not have. I was not lying about what the system could do. But I was compressing. Thirty seconds does not leave room for the hallucinated territory summary, or the three days lost to InfoSec, or the version that only worked after it failed once. The hero's arc is true. It is not the whole truth. I do not yet know whether organisations can tolerate the whole truth, or whether they need the compressed version to act.

---

The test came in a room with senior management.

I pitched it Y Combinator style: the problem, the solution, the traction, the ask. No forty-slide buildup. Just the arc, delivered standing up, ending with what the team would need after my internship ended to keep the plate spinning. The sentence that moved the conversation was not about agents or orchestration. It was: *"How future product managers need to be ready to get their hands dirty and not just be coordinators."*

They responded to the clarity, not the complexity. But not everyone in the room did. One stakeholder wanted the forty slides. They wanted to see the architecture diagram, the three-year roadmap. The YC pitch worked because the most senior person in the room preferred the compressed version. I left not knowing whether I had persuaded the room, or just the person who already agreed with me.

This is what Metaplast understood without trying: the ad does not lie about the product. It makes the product legible. But legibility is a gamble. The British lab read exactly what Metaplast projected. Someone else might have looked closer and seen the fiction.

---

In [The Name of the Bird](/blog/the-name-of-the-bird), I wrote that I had been handed the plate and was watching it spin. In [The Messy Reactor](/blog/the-messy-reactor), I wrote that the messy reactor was where the energy was. Both are still true.

The multi-agent system is running. The account managers are using it. The pipeline that synthesises insight at scale is no longer an idea. But the plate is still wobbling. There are edge cases I did not resolve. There are handoffs I designed but did not live long enough to stress-test. There is a conversation about production infrastructure that belongs to the team staying, not to me.

I am leaving at the right time, which is to say: before the wobble stops. If the wobble were solved, the project would be dead. Maintenance is not the same as life. The wobble means the thing is still moving, still generating questions, still alive enough to be interesting.

In [The Eye Level](/blog/the-eye-level), I wrote about the kirana owner who knows what moves on which days by instinct, and about the brand manager who wants that knowledge systematised. I do not know whether the system I built closes that gap or widens it. The territory summary is faster now. But speed is not the same as knowledge. The rep still knows things the model does not.

The Metaplast ad was a fiction of scale. The fiction worked because the fiddling was real. I leave behind both: the system, and the story that makes the system comprehensible to the next person who has to carry it.

The plate is still wobbling. I am leaving it that way on purpose.

---

*This is the last post. This blog is now archived. I am going to focus on building and learning something new, and I will not write about it until something tangible comes out of it.*
