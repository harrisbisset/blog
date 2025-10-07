<script lang="ts">
	import { type blog_post_t, BoxCase } from '$lib';
	import Posts from './posts.svelte';

	const { posts }: { posts: blog_post_t[] } = $props();
</script>

<div class="w-full flex justify-center">
	<BoxCase direction={'SE'} width={'80%'}>
		<div class="w-4/5">
			<Posts {posts} />
		</div>
	</BoxCase>
</div>
