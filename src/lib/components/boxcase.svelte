<script lang="ts">
	const { children } = $props();
</script>

<section
	class="box-case w-[300px] rounded-b-md border-x border-t border-r-2 border-b-2 border-background-1 p-1"
>
	{@render children?.()}
</section>

<style>
	.box-case:hover {
		box-shadow: 2px 2px var(--color-highlight-2);
	}
</style>
