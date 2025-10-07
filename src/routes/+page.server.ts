import type { project_t } from '$lib';
import type { article_t } from '$lib/components/articles';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = () => {
	return {
		articles_data: <Array<article_t>>[
			{
				name: 'ambigram.am',
				published: new Date(),
				description: 'interesting site',
				href: 'https://ambigr.am/hall-of-fame'
			},
			{
				name: 'ambigram.am',
				published: new Date(),
				description: 'interesting site',
				href: 'https://ambigr.am/hall-of-fame'
			}
		],
		projects_data: <Array<project_t>>[
			{
				href: "https://www.github.com",
				description: "example"
			},
		]
	};
};
