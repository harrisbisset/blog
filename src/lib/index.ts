export type { project_t } from '$lib/components/projects';

export { default as BoxCase } from '$lib/components/boxcase.svelte';

export function formatDate(date: Date): string {
	return `${appendZero(date.getDate())}/${appendZero(date.getMonth())}/${date.getFullYear()}`;
}

function appendZero(n: number): string {
	return n < 10 ? `0${n}` : `${n}`;
}
