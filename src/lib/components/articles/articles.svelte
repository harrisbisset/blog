<script lang="ts">
	import { BoxCase, type article_t } from '$lib';

	const { articles }: { articles: article_t[] } = $props();
</script>

<BoxCase TitleComponent={'latest articles'} width={undefined} colour={"orange"}>
	<div class="flex flex-col gap-2">
		{#each articles as article, i}
			{#if i > 0}
				<div class="h-0.5 w-full bg-highlight-2"></div>
			{/if}
			<div class="flex flex-col">
				<div class="flex flex-row items-baseline gap-2">
					<h3
						class="text-lg underline decoration-highlight-2 decoration-1 hover:bg-background-2"
						title={article.href}
					>
						<a target="_blank" rel="noopener noreferrer" href={article.href}>{article.name}</a>
					</h3>
					<div class="text-sm italic">{article.published}</div>
				</div>
				<div>{article.description}</div>
			</div>
		{/each}
	</div>
</BoxCase>
