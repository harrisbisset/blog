import type { Component } from 'svelte';

export interface commentary_metadata_t {
	title: string;
	part: number;
	date: string;
	description: string;
}

export interface commentary_t {
	component: Component;
	path: string;
	name: string;
	metadata: commentary_metadata_t;
}

export interface board_t {
	title: string;
	comments: commentary_t[];
}

export async function getBoards(): Promise<Array<board_t>> {
	let boards: Array<board_t> = [];

	const modules: Record<string, any> = import.meta.glob('./boards/*/*.svx', {
		eager: true
	});

	console.log(modules);

	for (const [path, o] of Object.entries(modules)) {
		let str = path.replace('./boards/', '');
		str = str.replace('.svx', '');
		const board_name = str.split('/');

		let comment: commentary_t = {
			metadata: o.metadata,
			component: o.default,
			path: path,
			name: board_name[1]
		};

		let found = false;
		boards.forEach((element) => {
			if (element.title == board_name[0]) {
				found = true;
				element.comments.push(comment);
			}
		});

		if (!found) boards.push({ title: board_name[0], comments: [comment] });
	}

	return boards;
}
