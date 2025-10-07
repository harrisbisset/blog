<script lang="ts">
	import Articles from '$lib/components/articles/articles.svelte';
	import Projects from '$lib/components/projects/projects.svelte';
	import Posts from '../lib/components/posts/homeposts.svelte';
	import ProfileCard from './profilecard.svelte';

	const { data } = $props();
</script>

<div class="flex w-full flex-row gap-y-3">
	<div class="flex w-full justify-center pt-10">
		<Posts posts={data.posts} />
		<div class="h-fit flex flex-col gap-7">
			<Articles articles={data.articles} />
			<Projects projects={data.projects}/>
		</div>
	</div>

	<div class="flex w-2/5 justify-end px-10 pt-5">
		<ProfileCard />
	</div>
</div>
