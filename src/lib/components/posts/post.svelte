<script lang="ts">
	import type { post_metadata_t } from "$lib";


	const { metadata, path }: { metadata: post_metadata_t, path: string } = $props();
</script>

<a href={`/blog/${path}`} class="group flex gap-1 h-fit">
	<div class="flex flex-col gap">
		<div class="flex flex-col">
			<h2 class="text-2xl font-semibold">{metadata.title}</h2>
			<div class="text-sm italic">{metadata.date}</div>
		</div>
		<div class="text-base font-sem">{metadata.description}</div>
	</div>
</a>
