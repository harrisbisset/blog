<script lang="ts">
	import Posts from '$lib/components/posts/posts.svelte';

	const { data } = $props();
</script>

<div class="flex h-full w-full justify-center">
	<div class="w-[var(--standard-width)] md:w-full border-l-4 border-dashed border-highlight-1 pl-4">
		<Posts posts={data.posts} />
	</div>
</div>
