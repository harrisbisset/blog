import projectsFile from '$lib/assets/projects.txt';
import { read } from '$app/server';
import type { project_t } from '.';

export async function getProjects(max: number): Promise<project_t[]> {
	let projects: project_t[] = [];

	const resp = await read(projectsFile).text();
	const projects_text = resp.split('\n');

	for (let i = 0; i < projects_text.length; i++) {
		if (max === i) break;

		const element = projects_text[i];

		const items = element.split(',');
		if (items.length != 3) {
			console.log('warning: incorrect length of project - ' + element);
			continue;
		}

		const project: project_t = {
			name: items[0],
			href: items[1],
			description: items[2]
		};

		projects.push(project);
	}

	return projects;
}
