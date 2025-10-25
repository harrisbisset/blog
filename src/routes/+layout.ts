import { getBlogPosts } from '$lib/data/blog';
import { getBoards } from '$lib/data/board';
import type { LayoutLoad } from './$types';

export const load: LayoutLoad = async () => {
    return {
        posts: await getBlogPosts(),
        boards: await getBoards(),
    }
}
