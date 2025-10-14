<script lang="ts">
	import { type blog_post_t, BoxCase } from '$lib';
	import Posts from './posts.svelte';

	const { posts }: { posts: blog_post_t[] } = $props();
</script>

<BoxCase colour={'orange'} name={'latest posts'} width={'80%'}>
	<div class="w-full md:w-4/5">
		<Posts {posts} />
		<a href="/blog" class="flex items-center gap-1 pt-3">
			<span class="text-lg italic decoration-highlight-2 hover:underline"> See more posts </span>
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
	</div>
</BoxCase>
