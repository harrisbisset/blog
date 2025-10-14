<script lang="ts">
	import { BoxCase, type project_t } from '$lib';

	const { projects }: { projects: project_t[] } = $props();
</script>

<BoxCase name={'latest articles'} width={undefined} colour={'blue'}>
	<div class="flex flex-col gap-2">
		{#each projects as project, i}
			{#if i > 0}
				<div class="h-0.5 w-full bg-highlight-2"></div>
			{/if}
			<div class="flex flex-col">
				<h3
					class="w-fit text-lg underline decoration-highlight-2 decoration-1 hover:bg-highlight-1"
					title={project.href}
				>
					<a target="_blank" rel="noopener noreferrer" href={project.href}>{project.name}</a>
				</h3>
				<div>{project.description}</div>
			</div>
		{/each}
	</div>
</BoxCase>
