<script lang="ts">
	import Articles from '$lib/components/articles/articles.svelte';
	import Projects from '$lib/components/projects/projects.svelte';
	import Posts from '../lib/components/posts/homeposts.svelte';
	import ProfileCard from './profilecard.svelte';

	const { data } = $props();
</script>

<div class="flex w-full justify-center md:block">
	<div
		class="flex w-[300px] flex-col gap-y-10 pb-40 md:w-full md:flex-row md:justify-center xl:grid xl:grid-cols-6"
	>
		<div
			class="flex flex-col gap-y-10 md:px-3 xl:col-span-4 xl:grid xl:grid-cols-4
		"
		>
			<div class="col-span-3 flex w-full justify-center">
				<div class="xl:w-7/8">
					<Posts posts={data.posts} />
				</div>
			</div>

			<div class="flex flex-col gap-y-10 md:gap-7 xl:items-center">
				<Projects projects={data.projects} />
				<Articles articles={data.articles} />
			</div>
		</div>
		<div class="col-span-2 flex justify-center md:justify-end">
			<ProfileCard />
		</div>
	</div>
</div>
