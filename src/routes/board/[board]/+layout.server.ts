import type { LayoutServerLoad } from './$types';

export const load: LayoutServerLoad = ({ params }) => {
	return {
		board_name: params.board
	};
};
