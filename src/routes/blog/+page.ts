export const prerender = true;

export async function load() {
	const modules = import.meta.glob(['/src/posts/*.md', '/src/posts/*.mdx'], { eager: true });

	const posts = Object.entries(modules)
		.map(([path, mod]: [string, any]) => {
			const slug = path.split('/').pop()!.replace(/\.mdx?$/, '');
			return {
				slug,
				title: mod.metadata?.title ?? slug,
				date: mod.metadata?.date ?? '',
				description: mod.metadata?.description ?? '',
				pinned: mod.metadata?.pinned ?? false,
				draft: mod.metadata?.draft ?? false,
			};
		})
		.filter((post) => import.meta.env.DEV || !post.draft);

	// Pinned posts always first, then by date descending
	posts.sort((a, b) => {
		if (a.pinned !== b.pinned) return a.pinned ? -1 : 1;
		return a.date < b.date ? 1 : -1;
	});

	return { posts };
}
