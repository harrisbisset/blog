<script lang="ts">
	const {
		heading,
		heading_level,
		sub_title
	}: { heading: string; heading_level: number; sub_title: string | undefined } = $props();
</script>

<div class="flex flex-col gap-1.5 pb-3">
	<h1
		class={[
			heading_level == 1 && 'text-3xl',
			heading_level == 2 && 'text-2xl',
			heading_level == 3 && 'text-xl',
			'font-semibold'
		]}
	>
		{heading}
	</h1>

	{#if sub_title! + undefined}
		<blockquote class="flex bg-zinc-200">
			<div class="w-5 bg-highlight-2"></div>
			<p class="pl-1">
				{sub_title}
			</p>
		</blockquote>
	{/if}
</div>
