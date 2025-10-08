<script>
	const { children } = $props();
</script>

<div class="my-2 border border-highlight-1 w-fit py-1 mx-2 px-1.5 text-lg font-semibold">
	{@render children?.()}
</div>
