+++
title = "Everything You Can Put in a Post"
date = "2026-03-05"
description = "A reference post showing images, code, callouts, table of contents, anchor links, YouTube embeds, and Mermaid diagrams."
draft = true
+++

<script>
  import Callout from '$lib/components/Callout.svelte';
  import YouTube from '$lib/components/YouTube.svelte';
  import Mermaid from '$lib/components/Mermaid.svelte';

  const flowchart = `
flowchart TD
    A([Write markdown]) --> B[git commit]
    B --> C{Build passes?}
    C -- Yes --> D([Deploy])
    C -- No --> E[Fix errors]
    E --> B
  `;

  const sequence = `
sequenceDiagram
    participant Browser
    participant SvelteKit
    participant Posts

    Browser->>SvelteKit: GET /blog/my-post
    SvelteKit->>Posts: import /src/posts/my-post.md
    Posts-->>SvelteKit: metadata + component
    SvelteKit-->>Browser: Rendered HTML
  `;

  const graph = `
graph LR
    md[Markdown file] --> mdsvex[mdsvex]
    mdsvex --> svelte[Svelte component]
    svelte --> html[Static HTML]
  `;
</script>

This post is a reference for every element you can use when writing. Jump to any section via the [contents](#contents) below.

## Contents

## Image

A standard Markdown image — hosted anywhere or placed in `static/`:

![A quiet mountain path through dense forest](https://images.unsplash.com/photo-1448375240586-882707db888b?w=800&q=80)

Caption goes here as regular text below the image if you want one.

## Code

Inline code: use `npm run dev` to start the dev server.

A fenced block with syntax highlighting:

```typescript
interface Post {
  slug: string;
  title: string;
  date: string;
}

async function loadPosts(): Promise<Post[]> {
  const modules = import.meta.glob('/src/posts/*.md', { eager: true });
  return Object.entries(modules).map(([path, mod]: [string, any]) => ({
    slug: path.split('/').pop()!.replace('.md', ''),
    title: mod.metadata.title,
    date: mod.metadata.date,
  }));
}
```

A shell snippet:

```bash
# Add a new post — that's all it takes
touch src/posts/my-idea.md
```

## Callouts

Four flavours of callout:

<Callout type="info">
This is a general note. Good for supplementary context that doesn't interrupt the main flow.
</Callout>

<Callout type="tip">
A quick shortcut: press `g` then `b` in the GitHub UI to jump straight to blame view on any file.
</Callout>

<Callout type="warning">
This approach works but has caveats. Make sure you understand the trade-offs before using it in production.
</Callout>

<Callout type="danger" title="Destructive operation">
Running this command will permanently delete the directory. There is no undo.
</Callout>

## Text Linking to Headings

Every heading on this page has an anchor you can link to. Hover a heading to see the `#` appear — click it to get a shareable URL.

You can also link directly in prose: see the [Image section](#image), the [Mermaid section](#mermaid-diagram), or jump back to the [table of contents](#contents).

The anchors are generated automatically from the heading text — spaces become hyphens, everything lowercased.

## YouTube Embed

Privacy-respecting embed via `youtube-nocookie.com`:

<YouTube id="dQw4w9WgXcQ" title="Rick Astley — Never Gonna Give You Up" />

Replace the `id` prop with any YouTube video ID.

## Mermaid Diagram

Diagrams are rendered client-side from plain text definitions — no image files needed.

A **flowchart**:

<Mermaid diagram={flowchart} />

A **sequence diagram**:

<Mermaid diagram={sequence} />

A **simple graph**:

<Mermaid diagram={graph} />
