<script lang="ts">
	import type { blog_post_t } from '$lib';

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

<div class="flex w-full flex-col items-center pb-36">
	<div class="blog-post w-4/5">
		{#if component != undefined}
			<component.component />
		{/if}
	</div>
</div>

<style>
	@media only screen and (min-width: 448px) {
		:global(.blog-post > * p) {
			width: 75%;
		}
	}

	:global(.blog-post > * p) {
		padding-bottom: 3px;
		padding-top: 3px;
		font-size: var(--text-lg);
	}

	:global(.blog-post > * a) {
		text-decoration: underline;
		text-decoration-color: var(--color-highlight-2);
		text-decoration-thickness: 1px;
	}

	:global(.blog-post > * a:hover) {
		background-color: var(--color-highlight-1);
	}
</style>
