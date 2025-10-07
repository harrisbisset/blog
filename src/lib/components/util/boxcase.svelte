<script lang="ts">
	import type { direction_t } from '.';


	let { children, direction, width="300" }: { children: any; direction: direction_t; width: string | undefined } = $props();
</script>

<section
	class={[
		direction === 'NE' && 'hover:shadow-ne rounded-t-md border-t-2 border-r-2',
		direction === 'SE' && 'hover:shadow-se rounded-b-md border-r-2 border-b-2',
		direction === 'SW' && 'hover:shadow-sw rounded-b-md border-r-2 border-b-2',
		direction === 'NW' && 'hover:shadow-nw rounded-t-md border-t-2 border-l-2',
		'border border-background-1 p-2 h-fit'
	]}
	style:width={width}
>
	{@render children?.()}
</section>
