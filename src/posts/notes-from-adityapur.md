+++
title = "Notes from Adityapur"
date = "2026-03-09"
description = "On doing a live consulting project with an MSME in Jamshedpur, what a Goldman and CS background does and does not prepare you for, and what happens when you use Claude to prototype a sustainability dashboard on a shop floor."
draft = true
+++

Every MBA program has some version of a live project. You get assigned a real company, a real problem, and a semester to do something useful. XLRI calls it BAS. Our team was placed with Utkal Auto Coach, a Tier 2 vendor in Adityapur Industrial Area, Jamshedpur. They make tipper bodies for Tata Motors. 57 years old. About 400 workers. One major client accounting for 95 percent of revenue.

The brief: ESG Health Card, sustainability report, and a digital production dashboard.

I said yes before fully understanding what any of that would involve on a shop floor.

---

There is a version of this project that stays comfortably in the classroom. You pull production records, build an emissions baseline in Excel, run the numbers on solar economics, write up a report, present it. That version exists and we did parts of it. The part that actually taught me something was the floor.

Walking through the plant on the first visit, you realize quickly that the competencies you brought in are not quite the competencies that are useful here. My Goldman Sachs background meant I knew how to read a compliance framework and construct an audit trail. My CS background meant I could think about data pipelines and system design. Neither of these told me anything about why a worker without a welding mask is making a rational decision given the workflow incentives in front of him, or why the semi-paint booth has open sides, or how the owner had already thought through the solar economics before we arrived.

What prepared me for this was not the technical background. It was the disposition the technical background had trained: get comfortable being in rooms where you do not know the most important things yet, ask the question that reveals you do not know, and stay in the conversation long enough to actually learn something.

The GS compliance work taught me to look for the gap between the process as documented and the process as practiced. That instinct transferred. The gap on the shop floor was everywhere.

---

The dashboard is where the tech background actually showed up.

The client wanted to skip phased implementation and go directly to an app-based data entry system integrated with a live production and emissions dashboard. His dealership operations already run on full digital automation. He saw no reason to spend a year getting there via Excel. He was right.

What would normally take weeks of scoping and back-and-forth, I started prototyping the same day using Claude. Not to ship production code, but to get something visual and interactive in front of the client fast, so the conversation could move from what we planned to build to what he actually wanted. That compression matters enormously in a consulting context. Clients respond to something real in a way they cannot respond to a slide. When you can show a working data flow and a rough dashboard in an afternoon, the feedback loop tightens from weeks to hours.

This is the version of AI-assisted work that I find genuinely interesting: not replacing the thinking, but collapsing the distance between an idea and something you can put in front of a human and watch them react to. The shop floor context made it more vivid because the feedback was immediate and the stakes were concrete. Either the line manager can enter the data in thirty seconds at the end of a shift, or the system fails. No abstraction available.

---

In [The Plate is Wobbling](/blog/the-plate-is-wobbling), I wrote about a day doing market immersion with a sales team for my Unilever internship. Walking through neighbourhood shops, watching a distribution rep operate from memory and instinct, seeing the gap between how a tool is designed and how it is used on the ground. The MSME project has the same texture.

Both experiences are variations of the same practice: showing up somewhere you do not fully understand, staying long enough to get past the surface, and paying attention to the things that do not fit the framework. The sales rep who knew every shopkeeper's buying patterns. The plant owner who had already evaluated all four of our environmental recommendations. In both cases, the person on the ground had knowledge that no briefing document had captured and no model had encoded.

What the MBA provides, if you use it right, is repeated structured exposure to this. Finance, operations, marketing, strategy: each domain has its own version of the gap between the model and the room. Learning to navigate that gap is the actual skill.

---

Sustainability as a concept is easy to be cynical about. In the Tier 2 manufacturing context it lands differently. Tata Motors is building out a platform to track emissions across its vendor ecosystem. When it goes live, vendors who cannot provide structured emissions data in a specified format will be at risk. For a company like Utkal, with 95 percent revenue concentration in one client, ESG compliance is not a values question. It is a survival question.

That reframes the work. You are not trying to convince anyone that sustainability matters. You are helping a 57-year-old plant become legible to a reporting infrastructure that is being built around it whether it is ready or not. The emissions that cannot be reduced (welding fumes, paint VOCs, the diesel crane with no electric alternative) go into the report as measured, monitored, and acknowledged. The ones that can be reduced (electricity through solar and efficiency measures, diesel forklifts) get a cost-benefit case and a roadmap.

The interesting ethical question is not whether this is real sustainability. It is whether making it measurable and reported is a meaningful step toward it, or a substitute for it. I do not have a clean answer. I think it is both, simultaneously, and that the honest thing is to say so.

---

In [Prepared to Play](/blog/prepared-to-play), I wrote that deep preparation in the fundamentals is what makes wandering in the applications rich rather than shallow. The CS background, the Goldman compliance work, the coursework: none of it prepared me for a tipper fabrication plant specifically. All of it prepared me to be useful in a room I had never been in before.

The wobbling plate was interesting to Feynman because he had prepared obsessively in the underlying physics. He was not starting from nothing. He had the foundation that made the curiosity generative.

That is what I keep finding in the MBA: situations where the foundations are visible in retrospect, where the preparation shows up as surplus rather than script. The live project is the application. The plate is wobbling. The equations are not the same ones, but the motion looks familiar.

---

The field notes and visit reports from this project are documented in the [project repository](https://github.com/junejakushal/bas-utkal-project) if you want the detail underneath this.
