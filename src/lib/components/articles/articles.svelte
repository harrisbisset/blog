<script lang="ts">
	import { BoxCase, type article_t } from '$lib';
	import Link from '../util/link.svelte';

	const { articles }: { articles: article_t[] } = $props();
</script>

<BoxCase TitleComponent={'interesting articles'} width={undefined} colour={'orange'}>
	<div class="flex flex-col gap-2">
		{#each articles as article, i}
			{#if i > 0}
				<div class="h-0.5 w-full bg-highlight-2"></div>
			{/if}
			<div class="flex flex-col">
				<div class="flex flex-row items-baseline gap-2">
					<Link href={article.href} text={article.name} />
					<div class="text-sm italic">{article.published}</div>
				</div>
				<div>{article.description}</div>
			</div>
		{/each}
	</div>
</BoxCase>
