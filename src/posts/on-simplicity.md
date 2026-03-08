---
title: On Simplicity
date: 2026-03-01
description: The hardest thing to build is something simple.
---

The hardest thing to build is something simple.

Simple software is rare. Not because engineers don't want it, but because simplicity requires restraint — saying no to features, resisting complexity, and doing the unglamorous work of deletion.

## The temptation of complexity

When you know a lot, you see all the edge cases. You see everything that *could* go wrong. And so you build for it. An extra layer of abstraction here, a configuration option there, a fallback for a scenario that will never happen in production.

Each addition seems reasonable in isolation. Collectively, they make the system incomprehensible.

> Perfection is achieved not when there is nothing more to add, but when there is nothing left to take away.
> — Antoine de Saint-Exupéry

## What simplicity isn't

Simplicity isn't the same as minimal features. A simple system can do a lot. What matters is whether the pieces fit together in a way that's easy to reason about.

Simple code reads like prose. You can hold it in your head. When something goes wrong, you know where to look.

## A practice

Before adding something, ask: what is the simplest possible thing that works? Build that. Then live with it. Resist the urge to generalize prematurely.

Delete more than you think you need to.
