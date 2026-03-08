import adapter from '@sveltejs/adapter-static';
import { mdsvex } from 'mdsvex';
import remarkToc from 'remark-toc';
import rehypeSlug from 'rehype-slug';
import rehypeAutolinkHeadings from 'rehype-autolink-headings';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	extensions: ['.svelte', '.md'],
	preprocess: [
		mdsvex({
			extensions: ['.md'],
			remarkPlugins: [
				[remarkToc, { heading: 'contents', tight: true }],
			],
			rehypePlugins: [
				rehypeSlug,
				[rehypeAutolinkHeadings, { behavior: 'append', content: { type: 'text', value: ' #' } }],
			],
		}),
	],
	kit: {
		adapter: adapter({
			fallback: '404.html'
		})
	}
};

export default config;
