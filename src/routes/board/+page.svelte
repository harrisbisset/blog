<script lang="ts">
	import Board from '$lib/components/board/board.svelte';

	const { data } = $props();
</script>

<div class="flex h-full w-full justify-center">
	<div class="w-[var(--standard-width)] md:w-full">
		<Board boards={data.boards} />
	</div>
</div>
