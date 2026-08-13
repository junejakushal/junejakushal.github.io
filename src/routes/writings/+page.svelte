<script lang="ts">
	let { data } = $props();
</script>

<svelte:head>
	<title>Writings — Kushal Juneja</title>
</svelte:head>

<section class="space-y-10">
	<div>
		<h1 class="text-2xl font-semibold tracking-tight mb-1">Writings</h1>
		<p class="text-gray-500 text-sm">Hands-on work, shared online.</p>
	</div>

	{#if data.writings.length > 0}
		<ul class="space-y-8">
			{#each data.writings as writing}
				<li>
					<a href="/writings/{writing.slug}" class="group block space-y-1">
						<div class="flex items-baseline justify-between gap-4">
							<h2 class="text-[15px] font-medium text-gray-900 group-hover:text-gray-600 transition-colors">
								{writing.title}
							</h2>
							<time class="text-xs text-gray-400 shrink-0" datetime={writing.date}>
								{new Date(writing.date).toLocaleDateString('en-US', { year: 'numeric', month: 'long', day: 'numeric' })}
							</time>
						</div>
						{#if writing.description}
							<p class="text-sm text-gray-500">{writing.description}</p>
						{/if}
					</a>
				</li>
			{/each}
		</ul>
	{:else}
		<p class="text-sm text-gray-500 italic">Brewing in progress. Check back soon.</p>
	{/if}
</section>
