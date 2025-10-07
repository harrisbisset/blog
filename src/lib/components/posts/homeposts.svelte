<script lang="ts">
	import type { blog_post_t } from '$lib/blog';
	import Posts from './posts.svelte';

	const { posts }: { posts: blog_post_t[] } = $props();
</script>

<section class="w-full px-14">
	<div class="box-case-2 w-4/5 rounded-l-md border border-t-2 border-l-2 border-background-1 p-4">
		<Posts {posts} />
	</div>
</section>

<style>
	.box-case-2:hover {
		box-shadow: -2px -2px var(--color-highlight-2);
	}
</style>
