import type { PageLoad } from './$types';
import { error } from '@sveltejs/kit';

export const load: PageLoad = async ({ data }) => {
  const { matchedFilename } = data;
  
  // Import all MDX files in the writings folder.
  const modules = import.meta.glob('../../../writings/*.mdx');
  
  const targetPath = `../../../writings/${matchedFilename}`;
  const resolver = modules[targetPath];
  
  if (!resolver) {
      throw error(404, 'Article component not found');
  }
  
  const mod = await resolver();
  
  return {
    article: {
      component: (mod as any).default,
      metadata: (mod as any).metadata ?? {}
    }
  };
};
