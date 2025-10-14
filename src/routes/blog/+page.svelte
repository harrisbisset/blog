<script lang="ts">
	import Posts from '$lib/components/posts/posts.svelte';

	const { data } = $props();
</script>

<div class="flex h-full w-full justify-center pt-3">
	<Posts posts={data.posts} />
</div>
