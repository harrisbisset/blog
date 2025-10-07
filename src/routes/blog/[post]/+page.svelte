<script lang="ts">
	import type { blog_post_t } from '$lib/blog/index.js';

	const { data } = $props();

	let component: blog_post_t | undefined = $state(undefined);

	for (let i = 0; i < data.posts.length; i++) {
		const element = data.posts[i];

		if (element.name == data.post_name) {
			component = element;
			break;
		}
	}
</script>

<div class="flex w-full flex-col items-center">
	<div class="w-4/5">
		{#if component != undefined}
			<component.component />
		{/if}
	</div>
</div>

<style>
	:global(p) {
		padding-bottom: 3px;
		padding-top: 3px;
        font-size: var(--text-lg);
	}
</style>
