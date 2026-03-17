+++
title = "The Pipe and the Pit"
date = "2026-03-17"
description = "On scope creep, cement pits, founders who make decisions without opening a deck, and whether measuring something is the same as improving it."
draft = true
+++

We went in to build something simple. Track the number of units produced at each stage of the factory floor. From that single number, you can calculate emissions per tipper, build a Scope 1 and 2 baseline, and give the client the data infrastructure they need for the sustainability report. Clean. Bounded. Achievable in a semester.

Then the operations team walked us through their day.

They spend hours every morning pulling together production numbers from across the floor, consolidating shift reports, chasing down line managers for counts that should have been entered the night before. All of it lands in an Excel file that someone emails to someone else. The ask was quiet and reasonable: since you are already capturing this data for the dashboard, could you also make it so we can generate the daily report from the same system?

You cannot say no to that. Not when you are there as a guest, not when the ask is genuinely sensible, not when the person making it has just spent two hours showing you exactly why their current process is broken. The scope doubled before we left the room.

This is how it always goes, I suspect. You arrive with a diagram. Reality has different ideas.

---

The meter reading was the moment that crystallised it.

One of the sustainability data points we needed was water consumption. To track water, you need meters. To install meters, you need to know where water enters and exits each part of the process. What started as a question about a single measurement turned into a twenty-minute conversation involving the founder, three line managers, a maintenance supervisor, and us, standing around a section of pipe trying to understand which way the water flows, where it branches, what feeds what.

No deck. No pre-read. No action items circulated forty-eight hours in advance. The founder called the relevant people over, they talked through the pipe layout in real time, a decision was made about where the meters should go, and that was that.

Business school has a lot to say about decision-making. Frameworks for stakeholder alignment. Structured approaches to cross-functional collaboration. The importance of bringing the right people into the room. What it does not quite prepare you for is a founder who simply is the alignment, who carries enough context and authority that a conversation around a pipe becomes a binding organisational decision in the time it takes to drink a cup of chai.

The Get Stuff Done instinct was something I recognised from engineering. It has largely gone quiet here. Most of the MBA operates on a different rhythm, and there are good reasons for that when organisations are large and decisions have wide consequences. But standing in that factory, watching a decision made without a single slide, I felt something I had not felt in a while. The pleasure of a problem actually closing.

---

The rainwater harvesting discussion was stranger.

As part of the ESG assessment, we were looking at what certifications the client might pursue and what would be required to achieve them. The CII sustainability frameworks have a points-based system. Rainwater harvesting is one of the items. To earn the certification points, the rainwater must be collected in a cement pit with a valve system, part of the water reusable on-site and part fed back into the groundwater separately through controlled drainage.

The founder's question was immediate: what is the point of rainwater harvesting?

The point, as everyone in the room understood it, is to replenish groundwater. Rain falls on the factory roof and on the surrounding land. If you reroute it to an open unpaved area and let it soak into the ground, you have returned that water to the aquifer. That is rainwater harvesting in any meaningful sense of the phrase.

The CII definition disagrees. Open ground does not count. You need the pit, the concrete, the valve, the infrastructure. The groundwater replenishment outcome is technically identical. The certified version costs significantly more and requires ongoing maintenance.

The founder was not wrong to ask the question. The certification captures the apparatus of the practice, not its purpose. You can build an expensive pit that feeds water back into the ground, or you can route it to open ground and achieve the same hydrological result for a fraction of the cost. Only one of these earns the certification point.

This generalises. Zero Liquid Discharge sounds absolute until you ask what it actually means and who is defining it and in what context. The standards are not wrong exactly. They are just built for a different kind of operation, by bodies that are closer to large corporates than to 400-person fabrication units in industrial estates, and the gap between the certification logic and the ground logic is wide enough to drive a tipper through.

The question the founder was really asking, underneath the rainwater harvesting discussion, was the same one I have been sitting with since [Notes from Adityapur](/blog/notes-from-adityapur): is making something measurable and certified the same thing as improving it? Does the pit actually do more for the aquifer than the open ground? Or does it just do more for the scorecard?

I still do not have a clean answer.

---

The AI safety monitoring idea came up in an earlier visit and has stayed with me.

The proposal was to use the existing CCTV infrastructure, run a computer vision model over the footage, detect PPE violations in real time, alert the line manager, generate a daily report. On paper it is elegant. The technology is not exotic. The hardware is already there.

The problem is not the technology.

If a worker is not wearing a welding mask, there is a reason. Maybe it is uncomfortable in the heat. Maybe it slows down a motion they do hundreds of times a day. Maybe the culture on that part of the floor has always been loose about it and nobody has pushed back. These are human compliance problems, and a model that flags the violation does not touch any of them. It just adds a notification that someone has to act on. If the line manager was not enforcing PPE before the AI system existed, the daily violation report does not change that calculus. It might actually make it worse, by creating a paper trail of unaddressed violations that becomes a liability without becoming a solution.

You can optimise the detection. You cannot optimise the will to act on it.

I do not think this means the system is useless. I think it means the system is not where the work is. The work is in why the norm is not being followed, which is a management question, a culture question, a question about whether line managers have the authority and incentive to enforce it. The AI sits downstream of all of that.

---

I wrote in [Prepared to Play](/blog/prepared-to-play) that deep preparation is what makes wandering in new domains rich rather than shallow. In [Notes from Adityapur](/blog/notes-from-adityapur) I wrote that the team is the preparation, and that the honest move was to show up, stay close, and pay attention.

Both of those things are still true. What I did not write, because I did not know it yet, is that paying attention on a factory floor eventually leads you to a set of questions that the frameworks do not resolve. Not because the frameworks are wrong but because the questions are more interesting than the frameworks anticipated.

The app wants to be small. The ops team needs it to be large. The certification wants a cement pit. The groundwater does not care. The AI can see the violation. It cannot see why.

The plate is still wobbling in Adityapur. The equation is still open in Jamshedpur. I am starting to think that might be the point.
