import { error } from '@sveltejs/kit';

export const prerender = true;

export async function load({ params }) {
	const modules = import.meta.glob(['/src/writings/*.md', '/src/writings/*.mdx']);

	const writingModule = modules[`/src/writings/${params.slug}.md`] ?? modules[`/src/writings/${params.slug}.mdx`];

	if (!writingModule) {
		throw error(404, 'Writing not found');
	}

	const mod = await writingModule() as any;

	if (!import.meta.env.DEV && mod.metadata?.draft) {
		throw error(404, 'Writing not found');
	}

	return {
		content: mod.default,
		title: mod.metadata?.title ?? params.slug,
		date: mod.metadata?.date ?? '',
		description: mod.metadata?.description ?? '',
	};
}

export async function entries() {
	const modules = import.meta.glob(['/src/writings/*.md', '/src/writings/*.mdx'], { eager: true });
	return Object.entries(modules)
		.filter(([, mod]: [string, any]) => import.meta.env.DEV || !mod.metadata?.draft)
		.map(([path]) => ({
			slug: path.split('/').pop()!.replace(/\.mdx?$/, ''),
		}));
}
