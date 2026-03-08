<script lang="ts">
	let { data } = $props();
</script>

<svelte:head>
	<title>{data.title} — Kushal</title>
	{#if data.description}
		<meta name="description" content={data.description} />
	{/if}
</svelte:head>

<article>
	<header class="mb-12 space-y-2">
		<h1 class="text-2xl font-semibold tracking-tight">{data.title}</h1>
		<div class="flex items-center gap-3 text-sm text-gray-400">
			<time datetime={data.date}>
				{new Date(data.date).toLocaleDateString('en-US', { year: 'numeric', month: 'long', day: 'numeric' })}
			</time>
		</div>
	</header>

	<div class="prose prose-gray prose-sm max-w-none
		prose-headings:font-semibold prose-headings:tracking-tight
		prose-a:text-gray-900 prose-a:underline prose-a:underline-offset-2
		prose-code:text-gray-800 prose-code:bg-gray-100 prose-code:px-1 prose-code:py-0.5 prose-code:rounded prose-code:text-[13px] prose-code:font-normal
		prose-pre:bg-gray-50 prose-pre:border prose-pre:border-gray-200
		prose-blockquote:border-l-gray-300 prose-blockquote:text-gray-600
		prose-hr:border-gray-200">
		<svelte:component this={data.content} />
	</div>

	<div class="mt-16 pt-8 border-t border-gray-100">
		<a href="/blog" class="text-sm text-gray-500 hover:text-gray-900 transition-colors">
			&larr; All posts
		</a>
	</div>
</article>
