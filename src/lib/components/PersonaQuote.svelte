<script lang="ts">
	interface Props {
		name: string;
		role: string;
		quote: string;
		avatar?: 1 | 2 | 3 | 4;
	}

	let { name, role, quote, avatar = 1 }: Props = $props();
</script>

<figure class="not-prose my-10 flex flex-col gap-5 sm:flex-row sm:items-start">
	<div class="shrink-0 self-start sm:self-auto" aria-hidden="true">
		<svg
			viewBox="0 0 160 160"
			class="h-32 w-32"
			fill="none"
			stroke="#111"
			stroke-width="1.6"
			stroke-linecap="round"
			stroke-linejoin="round"
		>
			<defs>
				<filter id="ink-{avatar}" x="-25%" y="-25%" width="150%" height="150%">
					<feTurbulence type="fractalNoise" baseFrequency="0.035" numOctaves="4" result="noise" />
					<feDisplacementMap in="SourceGraphic" in2="noise" scale="1.8" xChannelSelector="R" yChannelSelector="G" />
				</filter>
			</defs>

			<g filter="url(#ink-{avatar})">
				{#if avatar === 1}
					<!-- Foram: inquisitive researcher with glasses and curly hair -->
					<path d="M56 68c-3 10-3 22 1 32 5 14 18 24 35 25 16 1 30-10 35-24 4-11 3-23-1-33" />
					<path d="M52 62c-2 12-2 26 2 38" />
					<path d="M112 62c2 12 2 26-2 38" />
					<path d="M48 72c-7 6-9 18-5 28" />
					<path d="M116 72c7 6 9 18 5 28" />
					<path d="M50 58c8-14 26-18 42-16 12 1 24 8 30 18" />
					<path d="M44 52c-6-4-8-14-4-22 6-12 22-18 36-18 16 0 32 6 40 18" />
					<path d="M40 46c8-16 26-22 42-22 16 0 34 6 42 22" />
					<path d="M36 40c10-14 30-20 48-20 18 0 38 6 48 20" />
					<path d="M34 50c-8 8-10 22-4 32" />
					<path d="M126 50c8 8 10 22 4 32" />
					<path d="M40 38c4-6 12-10 20-10" />
					<path d="M58 30c8-4 18-4 26 0" />
					<path d="M88 32c8 4 14 10 18 18" />
					<path d="M102 44c6 8 8 18 4 26" />
					<path d="M108 56c4 10 2 22-4 30" />
					<path d="M42 56c-4 10-2 22 4 30" />
					<path d="M36 68c-2 12 2 24 10 32" />
					<g stroke-width="2.2">
						<circle cx="66" cy="78" r="10" />
						<circle cx="94" cy="78" r="10" />
						<path d="M56 78h20M84 78h20" />
					</g>
					<circle cx="66" cy="78" r="5" />
					<circle cx="94" cy="78" r="5" />
					<path d="M58 70c5-2 12-2 16 0" />
					<path d="M86 70c5-2 12-2 16 0" />
					<path d="M74 106c6 4 14 4 20 0" />
				{:else if avatar === 2}
					<!-- Sandeepa: homemaker with short wavy hair -->
					<path d="M54 70c-3 10-3 22 2 32 6 14 20 24 38 24 16 0 30-10 36-24 5-12 4-24 0-34" />
					<path d="M50 64c-2 12-2 26 2 38" />
					<path d="M114 64c2 12 2 26-2 38" />
					<path d="M48 72c-8 6-10 18-6 28" />
					<path d="M116 72c8 6 10 18 6 28" />
					<path d="M50 60c8-14 26-18 42-16 12 1 24 8 30 18" />
					<path d="M48 54c-4-2-6-10-4-16 4-12 18-18 30-18 14 0 28 6 34 18" />
					<path d="M52 46c8-10 22-14 34-14 12 0 26 4 34 14" />
					<path d="M56 38c10-6 24-8 36-6" />
					<path d="M62 32c10-4 22-4 32 0" />
					<path d="M64 86c0 3 2 5 5 5s5-2 5-5" />
					<path d="M88 86c0 3 2 5 5 5s5-2 5-5" />
					<path d="M80 78v6" />
					<path d="M70 112c8 5 18 5 26 0" />
					<path d="M58 54c10-2 28-2 42 2" />
					<path d="M76 70c2-2 6-2 8 0" />
					<circle cx="80" cy="66" r="2" />
				{:else if avatar === 3}
					<!-- Niharika: practical evaluator with long straight hair -->
					<path d="M54 70c-4 10-4 22 0 32 5 14 20 24 38 24 16 0 30-10 36-24 5-11 4-23 0-33" />
					<path d="M50 64c-2 12-2 26 2 38" />
					<path d="M114 64c2 12 2 26-2 38" />
					<path d="M48 74c-8 4-12 14-8 24" />
					<path d="M116 74c8 4 12 14 8 24" />
					<path d="M52 56c8-16 30-20 48-12 10 4 18 14 22 24" />
					<path d="M46 54c8-14 28-18 44-14 14 3 26 14 32 26" />
					<path d="M54 46c10-8 26-10 38-6" />
					<path d="M62 40c8-4 20-4 28 0" />
					<path d="M72 36c6 0 12 2 16 6" />
					<path d="M44 52c-8 8-10 22-6 36-2 18 0 36 4 52" />
					<path d="M116 52c8 8 10 22 6 36 2 18 0 36-4 52" />
					<path d="M40 58c-6 14-8 36-4 56" />
					<path d="M120 58c6 14 8 36 4 56" />
					<path d="M48 66c-4 14-4 34 0 48" />
					<path d="M112 66c4 14 4 34 0 48" />
					<path d="M80 54v20" />
					<path d="M72 50c4-2 10-2 14 0" />
					<circle cx="68" cy="84" r="3.5" />
					<circle cx="96" cy="84" r="3.5" />
					<path d="M62 76c5-1 10-1 14 0" />
					<path d="M86 76c5-1 10-1 14 0" />
					<path d="M72 108c8 4 18 4 26 0" />
					<path d="M64 68c8-2 20-2 28 0" />
					<path d="M82 68c8-2 20-2 28 0" />
				{:else}
					<!-- Disha: cautious skeptic with curly hair -->
					<path d="M54 68c-3 10-3 22 2 32 6 14 20 24 38 24 16 0 30-10 36-24 5-12 4-24 0-34" />
					<path d="M50 62c-2 12-2 26 2 38" />
					<path d="M114 62c2 12 2 26-2 38" />
					<path d="M48 72c-8 6-10 18-6 28" />
					<path d="M116 72c8 6 10 18 6 28" />
					<path d="M50 58c8-14 26-18 42-16 12 1 24 8 30 18" />
					<path d="M44 52c-6-4-8-14-4-22 6-12 22-18 36-18 16 0 32 6 40 18" />
					<path d="M40 46c8-16 26-22 42-22 16 0 34 6 42 22" />
					<path d="M36 40c10-14 30-20 48-20 18 0 38 6 48 20" />
					<path d="M34 50c-8 8-10 22-4 32" />
					<path d="M126 50c8 8 10 22 4 32" />
					<path d="M40 38c4-6 12-10 20-10" />
					<path d="M58 30c8-4 18-4 26 0" />
					<path d="M88 32c8 4 14 10 18 18" />
					<path d="M102 44c6 8 8 18 4 26" />
					<path d="M108 56c4 10 2 22-4 30" />
					<path d="M42 56c-4 10-2 22 4 30" />
					<path d="M36 68c-2 12 2 24 10 32" />
					<path d="M64 86c0 3 2 5 5 5s5-2 5-5" />
					<path d="M88 86c0 3 2 5 5 5s5-2 5-5" />
					<path d="M72 112c8 5 18 5 26 0" />
					<path d="M58 58c10-2 28-2 42 2" />
				{/if}
			</g>
		</svg>
	</div>

	<figcaption class="min-w-0 border-l-2 border-gray-300 pl-5 pt-1">
			<blockquote class="mb-2 text-sm italic leading-relaxed text-gray-600">
				“{quote}”
			</blockquote>
		<div class="flex items-baseline gap-2">
			<cite class="not-italic text-sm font-medium text-gray-900">{name}</cite>
			<span class="text-sm text-gray-500">{role}</span>
		</div>
	</figcaption>
</figure>
