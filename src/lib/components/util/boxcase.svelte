<script lang="ts">
	import type { Snippet } from 'svelte';
	import type { colour_t } from '.';

	let {
		children,
		TitleComponent,
		width = 'var(--standard-width)',
		colour
	}: {
		children: any;
		TitleComponent: Snippet | string;
		colour: colour_t;
		width: string | undefined;
	} = $props();
</script>

<section style={`width: ${width};`} class="h-fitsa min-w-(--standard-width)">
	<div class="flex h-fit flex-col gap-1">
		<div>
			<div class="text-2xl font-bold lowercase">
				{#if typeof TitleComponent === 'string'}
					{TitleComponent}
				{:else}
					{@render TitleComponent()}
				{/if}
			</div>
			<div
				class={[
					colour === 'orange' && 'border-highlight-1 shadow-highlight-1',
					colour === 'blue' && 'border-background-2 shadow-background-2',
					'relative -top-2.5 -z-10 w-full border-t-10  shadow-[0_4px_10px_-2px] '
				]}
			></div>
		</div>
		{@render children?.()}
	</div>
</section>
