<script lang="ts">
	import { BoxCase, type project_t } from '$lib';
	import Link from '../util/link.svelte';

	const { projects }: { projects: project_t[] } = $props();
</script>

<BoxCase TitleComponent={'featured projects'} width={undefined} colour={'blue'}>
	<div class="flex flex-col gap-2">
		{#each projects as project, i}
			{#if i > 0}
				<div class="h-0.5 w-full bg-highlight-2"></div>
			{/if}
			<div class="flex flex-col">
				<Link href={project.href} text={project.name} />
				<div>{project.description}</div>
			</div>
		{/each}
	</div>
</BoxCase>
