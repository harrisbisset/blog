<script lang="ts">
	import { formatDate } from '$lib';
	import type { post_metadata_t } from '$lib/blog';

	const { metadata, path }: { metadata: post_metadata_t, path: string } = $props();
</script>

<a href={`/blog/${path}`} class="border-b py-3">
	<div class="flex flex-col gap-2">
		<div class="gap flex flex-col">
			<h2 class="text-3xl font-semibold">{metadata.title}</h2>
			<div class="text-base">Published: {formatDate(metadata.date)}</div>
		</div>
		<div class="text-lg">{metadata.description}</div>
	</div>
</a>
