export const prerender = true;

export async function load() {
	const modules = import.meta.glob(['/src/writings/*.md', '/src/writings/*.mdx'], { eager: true });

	const writings = Object.entries(modules)
		.map(([path, mod]: [string, any]) => {
			const slug = path.split('/').pop()!.replace(/\.mdx?$/, '');
			return {
				slug,
				title: mod.metadata?.title ?? slug,
				date: mod.metadata?.date ?? '',
				description: mod.metadata?.description ?? '',
				draft: mod.metadata?.draft ?? false,
			};
		})
		.filter((writing) => import.meta.env.DEV || !writing.draft);

	writings.sort((a, b) => (a.date < b.date ? 1 : -1));

	return { writings };
}
