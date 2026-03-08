import { error } from '@sveltejs/kit';

export const prerender = true;

export async function load({ params }) {
	const modules = import.meta.glob('/src/posts/*.md');

	const postModule = modules[`/src/posts/${params.slug}.md`];

	if (!postModule) {
		throw error(404, 'Post not found');
	}

	const mod = await postModule() as any;

	return {
		content: mod.default,
		title: mod.metadata?.title ?? params.slug,
		date: mod.metadata?.date ?? '',
		description: mod.metadata?.description ?? '',
	};
}

export async function entries() {
	const modules = import.meta.glob('/src/posts/*.md');
	return Object.keys(modules).map((path) => ({
		slug: path.split('/').pop()!.replace('.md', ''),
	}));
}
