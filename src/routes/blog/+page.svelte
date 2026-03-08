<script lang="ts">
	let { data } = $props();
</script>

<svelte:head>
	<title>Blog — Kushal Juneja</title>
</svelte:head>

<section class="space-y-10">
	<div>
		<h1 class="text-2xl font-semibold tracking-tight mb-1">Writing</h1>
		<p class="text-gray-500 text-sm">Occasional thoughts on software and other things.</p>
	</div>

	<ul class="space-y-8">
		{#each data.posts as post}
			<li>
				<a href="/blog/{post.slug}" class="group block space-y-1">
					<div class="flex items-baseline justify-between gap-4">
						<div class="flex items-baseline gap-2">
							<h2 class="text-[15px] font-medium text-gray-900 group-hover:text-gray-600 transition-colors">
								{post.title}
							</h2>
							{#if post.pinned}
								<span class="text-[10px] font-medium tracking-widest uppercase text-gray-400 border border-gray-200 rounded px-1.5 py-0.5 leading-none">pinned</span>
							{/if}
						</div>
						<time class="text-xs text-gray-400 shrink-0" datetime={post.date}>
							{new Date(post.date).toLocaleDateString('en-US', { year: 'numeric', month: 'long', day: 'numeric' })}
						</time>
					</div>
					{#if post.description}
						<p class="text-sm text-gray-500">{post.description}</p>
					{/if}
				</a>
			</li>
		{/each}
	</ul>
</section>
