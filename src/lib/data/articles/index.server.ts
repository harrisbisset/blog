import articlesFile from '$lib/assets/articles.txt';
import { read } from '$app/server';
import type { article_t } from '.';

export async function getArticles(max: number): Promise<article_t[]> {
	let articles: article_t[] = [];

	const resp = await read(articlesFile).text();
	const articles_text = resp.split('\n');

	for (let i = 0; i < articles_text.length; i++) {

		if (max === i) break;

		const element = articles_text[i];

		const items = element.split(',');
		if (items.length != 4) {
			console.log('warning: incorrect length of article - ' + element);
			continue;
		}

		const article: article_t = {
			name: items[0],
			published: items[1],
			description: items[2],
			href: items[3]
		};

		articles.push(article);
	}

	return articles;
}