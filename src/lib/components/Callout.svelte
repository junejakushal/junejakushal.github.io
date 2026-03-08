<script lang="ts">
	type CalloutType = 'info' | 'warning' | 'danger' | 'tip';

	let { type = 'info', title, children }: { type?: CalloutType; title?: string; children: any } = $props();

	const styles: Record<CalloutType, { border: string; bg: string; icon: string; label: string }> = {
		info:    { border: 'border-blue-300',  bg: 'bg-blue-50',   icon: 'ℹ', label: 'Note'    },
		warning: { border: 'border-amber-300', bg: 'bg-amber-50',  icon: '⚠', label: 'Warning' },
		danger:  { border: 'border-red-300',   bg: 'bg-red-50',    icon: '✕', label: 'Danger'  },
		tip:     { border: 'border-green-300', bg: 'bg-green-50',  icon: '✦', label: 'Tip'     },
	};

	const s = styles[type];
</script>

<aside class="not-prose my-6 rounded-md border-l-4 {s.border} {s.bg} px-4 py-3">
	<div class="flex items-start gap-2.5">
		<span class="mt-0.5 text-sm select-none opacity-70">{s.icon}</span>
		<div class="text-sm leading-relaxed text-gray-800">
			<span class="font-semibold">{title ?? s.label}: </span>
			{@render children()}
		</div>
	</div>
</aside>
