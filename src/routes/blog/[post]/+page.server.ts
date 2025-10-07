import type { PageServerLoad } from './$types';

export const load: PageServerLoad = ({params}) => {
	return {
		post_name: params.post
	};
};
