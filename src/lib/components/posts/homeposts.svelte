<script lang="ts">
	import { type blog_post_t, BoxCase } from '$lib';
	import Posts from './posts.svelte';

	const { posts }: { posts: blog_post_t[] } = $props();
</script>

{#snippet title_component()}
	<div>
		<span>lastest</span>
		<a href="/blog" class="cursor-pointer transition-colors hover:bg-background-2">posts</a>
	</div>
{/snippet}

<BoxCase colour={'orange'} TitleComponent={title_component} width={'80%'}>
	<div class="w-full md:w-4/5">
		<Posts {posts} />
	</div>
</BoxCase>
