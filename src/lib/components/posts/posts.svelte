<script lang="ts">
	import type { blog_post_t } from '$lib';
	import Post from './post.svelte';

	const { posts }: { posts: blog_post_t[] } = $props();
</script>


<div class="flex flex-col gap-4">
	{#each posts as post, i}
		<Post metadata={post.metadata} path={post.name} />
	{/each}
</div>