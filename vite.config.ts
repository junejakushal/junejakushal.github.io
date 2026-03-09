import { sveltekit } from '@sveltejs/kit/vite';
import tailwindcss from '@tailwindcss/vite';
import { defineConfig } from 'vite';
import { execSync } from 'child_process';

const gitHash = execSync('git rev-parse --short HEAD').toString().trim();
const buildTime = new Date().toISOString();

export default defineConfig({
	plugins: [tailwindcss(), sveltekit()],
	define: {
		__GIT_HASH__: JSON.stringify(gitHash),
		__BUILD_TIME__: JSON.stringify(buildTime),
	}
});
