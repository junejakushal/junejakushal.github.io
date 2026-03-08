<script lang="ts">
	import { onMount } from 'svelte';
	import { browser } from '$app/environment';

	let { diagram }: { diagram: string } = $props();

	let container: HTMLDivElement;
	let svg = $state('');

	onMount(async () => {
		if (!browser) return;
		const mermaid = (await import('mermaid')).default;
		mermaid.initialize({ startOnLoad: false, theme: 'neutral', fontFamily: 'Inter, sans-serif' });
		const id = `mermaid-${Math.random().toString(36).slice(2)}`;
		const { svg: rendered } = await mermaid.render(id, diagram.trim());
		svg = rendered;
	});
</script>

<div class="not-prose my-6 flex justify-center rounded-lg border border-gray-200 bg-gray-50 p-6 overflow-x-auto">
	{#if svg}
		<div bind:this={container}>{@html svg}</div>
	{:else}
		<div class="text-xs text-gray-400 font-mono whitespace-pre">{diagram.trim()}</div>
	{/if}
</div>
