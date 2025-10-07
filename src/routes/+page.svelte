<script lang="ts">
	import Articles from '$lib/components/articles/articles.svelte';
	import Projects from '$lib/components/projects/projects.svelte';
	import Posts from '../lib/components/posts/homeposts.svelte';
	import ProfileCard from './profilecard.svelte';

	const { data } = $props();
</script>

<div class="flex w-full flex-col gap-y-3 md:flex-row pb-40 md:pb-0">
	<div class="flex w-full flex-col justify-center gap-5 pt-10 md:flex-row md:gap-0">
		<Posts posts={data.posts} />
		<div class="flex h-fit flex-col gap-7 items-center md:items-baseline">
			<Articles articles={data.articles} />
			<Projects projects={data.projects} />
		</div>
	</div>

	<div class="flex w-full md:w-2/5 md:justify-end md:px-10 md:pt-5 justify-center">
		<ProfileCard />
	</div>
</div>
