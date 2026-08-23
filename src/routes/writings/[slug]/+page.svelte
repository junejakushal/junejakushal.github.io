<script lang="ts">
  import type { PageData } from './$types';
  export let data: PageData;
  const { article } = data;
  const Component = article.component;
  const metadata = article.metadata || {};
</script>

<svelte:head>
	{#if metadata.title}
		<title>{metadata.title} — Kushal</title>
	{/if}
	{#if metadata.description}
		<meta name="description" content={metadata.description} />
	{/if}
</svelte:head>

<article>
	{#if metadata.title || metadata.date}
		<header class="mb-12 space-y-2">
			{#if metadata.title}
				<h1 class="text-2xl font-semibold tracking-tight">{metadata.title}</h1>
			{/if}
			{#if metadata.date}
				<div class="flex items-center gap-3 text-sm text-gray-400">
					<time datetime={metadata.date}>
						{new Date(metadata.date).toLocaleDateString('en-US', { year: 'numeric', month: 'long', day: 'numeric' })}
					</time>
				</div>
			{/if}
		</header>
	{/if}

	<div class="prose prose-gray prose-sm max-w-none
		prose-headings:font-semibold prose-headings:tracking-tight
		prose-a:text-gray-900 prose-a:underline prose-a:underline-offset-2
		prose-code:text-gray-800 prose-code:bg-gray-100 prose-code:px-1 prose-code:py-0.5 prose-code:rounded prose-code:text-[13px] prose-code:font-normal
		prose-pre:bg-gray-50 prose-pre:border prose-pre:border-gray-200
		prose-blockquote:border-l-gray-300 prose-blockquote:text-gray-600
		prose-hr:border-gray-200">
		<svelte:component this={Component} />
	</div>

</article>
