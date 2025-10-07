<script lang="ts">
	import type { blog_post_t } from '$lib';
	import Post from './post.svelte';

	const { posts }: { posts: blog_post_t[] } = $props();
</script>


<div class="text-xl font-semibold">Latest Posts</div>
<div class="flex flex-col gap-4">
	{#each posts as post, i}
		<Post metadata={post.metadata} path={post.name} />
	{/each}
</div>

<a href="/blog" class="flex items-center gap-1 pt-3">
	<span class="text-lg font-semibold decoration-highlight-2 hover:underline"> See more posts </span>
	<svg
		class="stroke-current"
		xmlns="http://www.w3.org/2000/svg"
		viewBox="0 0 24 24"
		stroke-width="1.5"
		fill="none"
		width="24"
		height="24"
		aria-hidden="true"
		><path
			stroke-linecap="round"
			stroke-linejoin="round"
			d="M17.25 8.25L21 12m0 0l-3.75 3.75M21 12H3"
		></path>
	</svg>
</a>
