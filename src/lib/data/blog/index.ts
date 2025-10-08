import type { Component } from 'svelte';

export interface blog_post_t {
	metadata: post_metadata_t;
	component: Component;
	path: string;
	name: string;
}

export interface post_metadata_t {
	title: string;
	date: string;
	description: string;
}

export async function getBlogPosts(): Promise<Array<blog_post_t>> {
	let posts: Array<blog_post_t> = [];

	const modules: Record<string, any> = import.meta.glob('./posts/*.svx', {
		eager: true
	});

	for (const [path, o] of Object.entries(modules)) {
		let metadata = <post_metadata_t>o.metadata;
		metadata.date = metadata.date;

		let name = path.replace('./posts/', '');
		name = name.replace('.svx', '');

		posts.push({ metadata: o.metadata, component: o.default, path: path, name: name });
	}

	posts.sort(function (a: blog_post_t, b: blog_post_t): number {
		const a_date = new Date(a.metadata.date);
		const b_date = new Date(b.metadata.date);

		if (a_date > b_date) return -1;
		else if (a_date < b_date) return 1;
		else return 0;
	});

	return posts;
}
