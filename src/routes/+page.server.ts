import { getArticles } from '$lib/data/articles/index.server';
import { getProjects } from '$lib/data/projects/index.server';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async () => {
	return {
		articles: await getArticles(4),
		projects: await getProjects(4),
	};
};
