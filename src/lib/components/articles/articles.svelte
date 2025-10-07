<script lang="ts">
	import { BoxCase, formatDate } from '$lib';
	import type { article_t } from '.';

	const { articles }: { articles: article_t[] } = $props();
</script>

<BoxCase>
	<div class="flex flex-col gap-2">
		<div class="text-xl font-semibold">Latest Articles</div>
		<div class="flex flex-col gap-2">
			{#each articles as article, i}
				{#if i > 0}
					<div class="h-0.5 w-full bg-highlight-2"></div>
				{/if}
				<div class="flex flex-col">
					<div class="flex flex-row items-center gap-1.5">
						<h3
							class="text-lg underline decoration-highlight-2 decoration-1 hover:bg-highlight-1"
							title={article.href}
						>
							<a href={article.href}>{article.name}</a>
						</h3>
						<div class="text-sm italic">{formatDate(article.published)}</div>
					</div>
					<div>{article.description}</div>
				</div>
			{/each}
		</div>
	</div>
</BoxCase>
