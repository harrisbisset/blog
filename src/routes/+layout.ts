import { getBlogPosts } from '$lib/blog';
import type { LayoutLoad } from './$types';

export const load: LayoutLoad = async () => {
    return {
        posts: await getBlogPosts()
    }
}
