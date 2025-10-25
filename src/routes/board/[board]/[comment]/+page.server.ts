import type { PageServerLoad } from './$types';

export const load: PageServerLoad = ({params}) => {
    return {
        comment_name: params.comment
    };
};
