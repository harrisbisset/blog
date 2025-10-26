<script lang="ts">
	import type { commentary_t } from '$lib/data/board/index.js';

	const { data } = $props();

	let component: commentary_t | undefined = $state(undefined);
	data.boards.forEach((element) => {
		if (element.title == data.board_name) {
			element.comments.forEach((element) => {
				component = element;
				return;
			});
		}
	});
</script>

<div class="flex w-full flex-col items-center pb-36">
	<div class="w-4/5">
		{#if component != undefined}
			<component.component />
		{/if}
	</div>
</div>
