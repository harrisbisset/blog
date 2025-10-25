<script lang="ts">
	import { BoxCase } from '$lib';
	import Board from './board.svelte';
	import type { board_t } from '$lib/data/board';

	const { boards }: { boards: board_t[] } = $props();
</script>

{#snippet title_component()}
	<div class="group">
		<a href="/board" class="board-heading text-transparent">boards</a>
	</div>
{/snippet}

<BoxCase colour={'orange'} TitleComponent={title_component} width={'80%'}>
	<div class="flex justify-center pt-1">
		<Board {boards} />
	</div>
</BoxCase>

<style>
	@property --myColor1 {
		syntax: '<color>';
		initial-value: #000;
		inherits: false;
	}

	@property --myColor2 {
		syntax: '<color>';
		initial-value: #000;
		inherits: false;
	}

	.board-heading {
		background: linear-gradient(var(--myColor1), var(--myColor2));
		background-clip: text;
		transition: --myColor2 0.3s;
	}

	.board-heading:hover {
		--myColor2: var(--color-highlight-1);
	}
</style>
