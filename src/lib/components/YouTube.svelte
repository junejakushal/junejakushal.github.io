<script lang="ts">
	let { id, title = 'YouTube video' }: { id: string; title?: string } = $props();
</script>

<div class="not-prose my-6 overflow-hidden rounded-lg border border-gray-200 aspect-video">
	<iframe
		src="https://www.youtube-nocookie.com/embed/{id}"
		{title}
		allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
		allowfullscreen
		class="w-full h-full"
		loading="lazy"
	></iframe>
</div>
