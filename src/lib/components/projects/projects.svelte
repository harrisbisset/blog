<script lang="ts">
	import type { project_t } from '$lib';

	const { projects }: { projects: project_t[] } = $props();
</script>

<section
	class="box-case-3 w-[300px] rounded-l-md border border-b-2 border-l-2 border-background-1 p-1"
>
	<div class="pb-3 text-xl font-semibold">Projects</div>

	{#each projects as project, i}
		{#if i > 0}
			<div class="h-0.5 w-full bg-highlight-2"></div>
		{/if}
		<div>
			<a class="text-lg underline decoration-highlight-2 hover:bg-highlight-1" href={project.href}
				>{project.href}</a
			>
			<div>{project.description}</div>
		</div>
	{/each}
</section>

<style>
	.box-case-3:hover {
		box-shadow: -2px 2px var(--color-highlight-2);
	}
</style>
