import type { PageServerLoad } from './$types';
import { createHash } from 'crypto';
import { error } from '@sveltejs/kit';

export const load: PageServerLoad = async ({ params }) => {
  const { slug } = params;
  
  const modules = import.meta.glob('../../../writings/*.mdx', { query: '?url', import: 'default' });
  
  let matchedFilename: string | null = null;
  
  for (const path of Object.keys(modules)) {
      const filename = path.split('/').pop();
      if (!filename) continue;
      
      const hash = createHash('md5').update(`src/writings/${filename}`).digest('hex').slice(0, 16);
      if (hash === slug) {
          matchedFilename = filename;
          break;
      }
  }
  
  if (!matchedFilename) {
      throw error(404, 'Article not found');
  }
  
  return { matchedFilename };
};
