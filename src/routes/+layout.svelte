<script lang="ts">
	import '../app.css';
	import { page } from '$app/stores';

	declare const __GIT_HASH__: string;
	declare const __BUILD_TIME__: string;

	let { children } = $props();

	const hash = __GIT_HASH__;
	const builtAt = new Date(__BUILD_TIME__).toLocaleDateString('en-US', {
		year: 'numeric', month: 'long', day: 'numeric'
	});

	const links = [
		{ href: '/', label: 'about' },
		{ href: '/blog', label: 'blog' },
	];
</script>

<svelte:head>
	<link rel="preconnect" href="https://fonts.googleapis.com" />
	<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin="anonymous" />
	<link href="https://fonts.googleapis.com/css2?family=Inter:ital,opsz,wght@0,14..32,300;0,14..32,400;0,14..32,500;1,14..32,300;1,14..32,400&display=swap" rel="stylesheet" />
</svelte:head>

<div class="min-h-screen bg-white text-gray-900">
	<header class="border-b border-gray-100">
		<nav class="max-w-2xl mx-auto px-6 py-5 flex items-center justify-between">
			<a href="/" class="text-sm font-medium tracking-wide text-gray-900 hover:text-gray-600 transition-colors">
				Kushal Juneja
			</a>
			<div class="flex gap-6">
				{#each links as link}
					<a
						href={link.href}
						class="text-sm transition-colors {$page.url.pathname === link.href || ($page.url.pathname.startsWith('/blog') && link.href === '/blog')
							? 'text-gray-900 font-medium'
							: 'text-gray-500 hover:text-gray-900'}"
					>
						{link.label}
					</a>
				{/each}
			</div>
		</nav>
	</header>

	<main class="max-w-2xl mx-auto px-6 py-16">
		{@render children()}
	</main>

	<footer class="border-t border-gray-100 mt-16">
		<div class="max-w-2xl mx-auto px-6 py-8 text-center text-xs text-gray-400 space-y-1">
			<div>&copy; {new Date().getFullYear()} Kushal Juneja</div>
			<div>
				Updated {builtAt} &middot;
				<a
					href="https://github.com/junejakushal/junejakushal.github.io/commit/{hash}"
					target="_blank"
					rel="noopener noreferrer"
					class="font-mono hover:text-gray-600 transition-colors"
				>{hash}</a>
			</div>
		</div>
	</footer>
</div>
