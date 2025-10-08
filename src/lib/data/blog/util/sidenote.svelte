<script lang="ts">
	const { children } = $props();
</script>

<div class="float-none md:float-right m-1 md:w-1/4">
	<div class="w-full bg-background-2 text-lg px-1 pt-1 pb-0.5">Side note:</div>

	<div class="bg-zinc-50 p-1 text-lg">
		{@render children?.()}
	</div>
</div>
