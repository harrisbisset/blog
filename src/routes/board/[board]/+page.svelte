<script lang="ts">
	import type { board_t } from '$lib/data/board/index.js';

	const { data } = $props();

	let board: board_t | undefined = $state(undefined);

	data.boards.forEach((element) => {
		if (element.title == data.board_name) {
			board = element;
			return;
		}
	});
</script>

<div class="flex justify-center w-full">
	{#each board?.comments as comment}
		<div class="w-fit">
			<a href={`/board/${board?.title}/${comment.name}`} class="underline">
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
