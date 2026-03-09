import adapter from '@sveltejs/adapter-static';
import { mdsvex } from 'mdsvex';
import remarkToc from 'remark-toc';
import rehypeSlug from 'rehype-slug';
import rehypeAutolinkHeadings from 'rehype-autolink-headings';
import { parse as parseToml } from 'smol-toml';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	extensions: ['.svelte', '.md'],
	preprocess: [
		mdsvex({
			extensions: ['.md'],
			frontmatter: {
				type: 'toml',
				marker: '+',
				parse: parseToml,
			},
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
