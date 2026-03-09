+++
title = "Notes from Adityapur"
date = "2026-03-09"
description = "On advising a 57-year-old MSME on ESG, what the factory floor teaches you that the classroom cannot, and why sustainability in Tier 2 manufacturing is a harder problem than it looks."
draft = true
+++

Utkal Auto Coach has been making tipper bodies in Adityapur Industrial Area, Jamshedpur, since 1967. The plant is 57 years old. About 400 people work there, roughly 300 of them contractual. The owner, Shibu, runs it with the kind of operational clarity that only comes from decades of surviving thin margins. His son Arhaan, a new entrant, is curious about solar panels, rainwater harvesting, and automation.

Our team from XLRI has been working with them since December 2025 on an ESG Health Card, a sustainability report for FY 25-26, and a digital production dashboard. This post is a working account of what we found, what we got wrong, and what the factory floor teaches you that two semesters of strategy courses cannot.

I am documenting the work as I go. Field notes and visit reports are in the [project repository](https://github.com/junejakushal/bas-utkal-project).

---

## The 95 percent problem

The first thing to understand about Utkal is that 95 percent of its revenue comes from Tata Motors. One client. One relationship. The language for this in strategy class is "bargaining power," and the textbook verdict is clear: you do not want to be here.

The practical reality is more specific. Tata Motors runs a "Should Cost" department. Utkal submits a manufacturing cost of, say, 130 rupees per unit. Tata's team calculates their own figure, arrives at 110, and the gap is absorbed by the vendor. There is no negotiation. The owner described it plainly: "It's my way or the highway."

ESG compliance, in this context, is not a values question. It is a vendor viability question. Tata Motors is building out a platform called Prakriti, developed with TCS, to track emissions across its vendor ecosystem. When (not if) it goes live, vendors who cannot provide electricity consumption data, GHG emissions per process, and audit trails in digital format will be at risk. Our job is to make sure Utkal is ready.

This changes how you think about the work. You are not advising a company that wants to be sustainable. You are advising a company that needs to survive, and sustainability is becoming a precondition for that.

---

## What two visits looked like

We visited the factory twice. The [first visit on January 23](https://github.com/junejakushal/bas-utkal-project/blob/main/20250123_visit_report.md) was orientation: understanding the production process, mapping the data gaps, and beginning to see where the ESG exposures actually were. The [second visit on February 18](https://github.com/junejakushal/bas-utkal-project/blob/main/20260218_visit_report.md) was presentation: we came back with findings, recommendations, and a proposed digital roadmap.

The gap between what I expected the second visit to look like and what it actually was is probably the most useful thing to write about.

I expected pushback on the recommendations. I expected the owner to be skeptical about solar economics, resistant to the electric forklift analysis, cautious about the APFC panel investment. Instead, he told us all four were already on his radar. The 150 kW solar PO had been placed. The APFC work was pending internally. LED lighting was 80 percent complete. The electric forklift evaluation had started but stalled.

The lesson: when you advise a company that has been running for 57 years, the problem is rarely that they have not thought of the solution. The problem is implementation, sequencing, and organizational bandwidth.

Where we added value was not in the recommendations themselves. It was in the structure: a 3D prioritization matrix sorting initiatives by cost, time, and effort; a per-tipper emissions baseline built from 20 months of production records; a calculated payback period for each intervention. Making tacit knowledge explicit, quantified, and presentable to a Tata Motors audit team. That is the actual work.

---

## The number that was wrong

We built a waterfall diagram showing per-tipper cost reduction from our four environmental recommendations. The owner interrupted us to point out that our baseline electricity cost was off.

We had estimated 9.5 rupees per unit of electricity. He said it was 6 to 6.5. We had calculated the energy cost per tipper at 35,000 rupees. He said his monthly electricity bill divided by monthly production gave a number closer to 2,264.

The four recommendations remain valid. The savings are real. But the denominator was wrong, which meant the per-tipper impact figures were wrong, which undermined the financial case at exactly the moment we needed it to land.

This is a mundane lesson that is easy to forget in the comfort of working from shared Excel files. The number you are using to calculate everything else needs to be verified with the person who signs the electricity bills. We will correct it. But it should not have gone into the presentation unchecked.

---

## What you cannot fix

The semi-paint booth has open sides. Paint fumes and welding exhaust leave through those openings directly into the air. There are 115 welding machines. Installing centralized ducting for all of them is operationally impossible at current margins.

The owner described this as an occupational hazard of the tipper industry. He is not wrong. It is also a Scope 1 emissions source that we cannot reduce, only estimate and report.

This is where the classroom version of ESG and the factory floor version diverge most sharply. In the classroom, emissions are a problem to be solved. On the floor, some of them are structural features of the production process. The honest answer is: monitor it, build a proxy measurement methodology (emissions per electrode, per paint can, per store issue), and put it in the sustainability report as a known gap with a stated measurement approach.

The strategy professor might call this greenwashing. The plant operator would call it the difference between a report that reflects reality and one that does not get filed at all.

---

## The digital decision

We proposed a three-phase digital roadmap: Excel to dashboard, dashboard to ERP, ERP integration. Standard change management logic. Manage the transition, do not overwhelm the organization.

The owner rejected all three phases and asked us to go straight to app-based data entry by line managers at each production stage, feeding directly into a live dashboard.

His reasoning: the dealership operations he runs on the side (Mahindra, Kia) are fully automated with mobile apps. All the line managers at Utkal have smartphones. He sees no reason to take two years getting there via Excel.

I spent some time thinking about whether he was right. The standard argument for phased adoption is that organizations need time to build data literacy and process discipline before the technology is meaningful. The argument against it, in this case, is that the owner already has the data literacy from adjacent operations, the workforce has the devices, and each additional phase is a year of delayed Prakriti readiness.

He was right. We are building the app.

---

## What I am still working through

The social side of ESG is harder to think about clearly than the environmental side.

Utkal has roughly 280 to 300 contractual workers. Many are classified as unskilled or semi-skilled to reduce compliance costs. The safety observations from the first visit were significant: absent PPE, no fire hydrant line, workers walking under active EOT cranes, chemical storage without proper zoning. The owner's response was that Tata Motors conducts weekly audits and they have had zero incidents this year.

I do not know how to think about that. The audit passes. The incidents are zero. The conditions are also genuinely dangerous. Whether the audit is a ceiling (meet the standard, stop there) or a floor (meet the standard, then go further) is a governance question, not a measurement question, and I do not think I have a clean answer yet.

The fire hydrant line is being implemented. The AI-based CCTV system for PPE compliance is being scoped. These are real improvements. Whether they add up to a safe workplace or a compliant one is a distinction I am still sitting with.

---

The work continues. I will keep updating the [project repo](https://github.com/junejakushal/bas-utkal-project) as we build out the dashboard and finalize the sustainability report. What I know so far is that this is more interesting, and more genuinely difficult, than any case study I have read.

The factory floor is not a strategy problem. It is a real place, with real constraints, where real people work. The job is to be useful to it.
