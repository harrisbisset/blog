<script lang="ts">
	import type { board_t } from '$lib/data/board';

	const { boards }: { boards: board_t[] } = $props();
</script>

<div class="flex flex-wrap justify-center gap-x-5 gap-y-3">
	{#each boards as board}
		<div
			class="flex flex-col gap-0.5 bg-zinc-50 px-1 pb-2"
		>
			<a class="text-lg font-semibold bg-blue-200 px-1 hover:bg-border transition-all" href={`/board/${board.title}`}>{board.title}</a>
			<div class="flex flex-col gap-1.5 px-1">
				{#each board.comments.slice(0, 3) as comment}
					<div>
						<a href={`/board/${board.title}/${comment.name}`} class="underline">
							{comment.name}
						</a>
						{#if comment.metadata.description != ''}
							<p class="text-sm italic">
								"{comment.metadata.description}"
							</p>
						{/if}
					</div>
				{/each}
			</div>
		</div>
	{/each}
</div>
