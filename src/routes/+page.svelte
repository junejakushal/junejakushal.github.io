<script lang="ts">
	declare const __GIT_HASH__: string;
	declare const __BUILD_TIME__: string;

	const hash = __GIT_HASH__;
	const builtAt = new Date(__BUILD_TIME__).toLocaleDateString('en-US', {
		year: 'numeric', month: 'long', day: 'numeric'
	});
</script>

<svelte:head>
	<title>Kushal | About</title>
</svelte:head>

<section class="space-y-10">
	<div>
		<h1 class="text-2xl font-semibold tracking-tight mb-1">Kushal Juneja</h1>
		<p class="text-gray-500 text-sm">MBA @ XLRI Jamshedpur · Incoming Summer Intern @ Unilever · ex-Goldman Sachs · CS @ IIIT-Delhi</p>
	</div>

	<div class="space-y-4 text-gray-700 leading-relaxed text-[15px]">
		<p>
			I'm an MBA student at XLRI Jamshedpur with a background in Computer Science from IIIT-Delhi,
			where my thesis sat at the intersection of Climate Science and Deep Learning. Before business
			school, I was an Analyst at Goldman Sachs in their Global Compliance and Engineering division.
		</p>
		<p>
			This summer I'll be joining the Unilever team as part of my MBA internship, excited to work
			at the intersection of business and technology at scale.
		</p>
		<p>
			My interests span technology, finance, and sustainability. I'm currently exploring agentic AI,
			scalable distributed systems, and finance. Staying fit and reading widely.
		</p>
		<p class="text-gray-500 text-sm italic">
			Overprepare. Always. You can't wing it and assume someone else will do it for you.
		</p>
	</div>

	<div class="space-y-3">
		<h2 class="text-xs font-semibold tracking-widest uppercase text-gray-400">Links</h2>
		<ul class="space-y-1.5 text-sm">
			<li>
				<a href="https://github.com/junejakushal" class="text-gray-700 hover:text-gray-900 transition-colors" target="_blank" rel="noopener noreferrer">
					GitHub
				</a>
				<span class="text-gray-400 mx-2">/</span>
				<span class="text-gray-500">@junejakushal</span>
			</li>
			<li>
				<a href="https://in.linkedin.com/in/kushal-juneja-40133419a" class="text-gray-700 hover:text-gray-900 transition-colors" target="_blank" rel="noopener noreferrer">
					LinkedIn
				</a>
				<span class="text-gray-400 mx-2">/</span>
				<span class="text-gray-500">Kushal Juneja</span>
			</li>
			<li>
				<a href="https://x.com/junejakushal" class="text-gray-700 hover:text-gray-900 transition-colors" target="_blank" rel="noopener noreferrer">
					X / Twitter
				</a>
				<span class="text-gray-400 mx-2">/</span>
				<span class="text-gray-500">@junejakushal</span>
			</li>
			<li>
				<a href="mailto:b25079@astra.xlri.ac.in" class="text-gray-700 hover:text-gray-900 transition-colors">
					b25079@astra.xlri.ac.in
				</a>
			</li>
		</ul>
	</div>

</section>
