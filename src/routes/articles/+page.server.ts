import { getArticles } from '$lib/data/articles/index.server';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async () => {
	return {
		articles: await getArticles(-1),
	};
};
