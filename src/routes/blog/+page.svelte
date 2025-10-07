<script lang="ts">
	import Posts from '$lib/components/posts/posts.svelte';

	const { data } = $props();
</script>

<div class="flex h-full w-full justify-center">
	<div class="flex w-4/5 flex-col">
		<Posts posts={data.posts} />
	</div>
</div>
