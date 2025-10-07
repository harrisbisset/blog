<script lang="ts">
	import type { blog_post_t } from '$lib';
	import { onMount } from 'svelte';

	const { data }: { data: { posts: blog_post_t[] } } = $props();

	let navbar: HTMLElement;
	onMount(() => {
		document.addEventListener('scroll', function (e) {
			if (navbar == null) return;

			if (window.pageYOffset > 15) {
				navbar.style.boxShadow = '10px 10px #dba45c';
			} else {
				navbar.style.boxShadow = '';
			}
		});
	});

	let ctrl = $state(false);
	let q = $state(false);
	let toggled = $state(false);

	let search: HTMLElement;
	let searchbar: HTMLElement;

	function showSearch() {
		if (search != null) search.style.display = 'flex';
		if (searchbar != null) searchbar.focus();
		toggled = true;
	}

	onMount(() => {
		document.addEventListener('keydown', function (e) {
			if (e.ctrlKey) ctrl = true;
			if (e.key == 'q') q = true;
			if (e.key == 'Escape' && search != null) {
				search.style.display = 'none';
				toggled = false;
			}

			if (ctrl && q) {
				if (toggled) {
					if (search != null) search.style.display = 'none';
					toggled = false;
					return;
				}

				showSearch();
			}
		});

		document.addEventListener('keyup', function (e) {
			if (e.ctrlKey) ctrl = false;
			if (e.key == 'q') q = false;
		});
	});

	// const routes = ['', 'blog', 'articles', 'projects'];
	const routes = ['', 'blog'];
</script>

<div class="fixed flex w-full items-center justify-center pt-1">
	<nav
		class="flex h-16 w-fit flex-row items-center justify-center gap-8 rounded-b-sm border border-background-1 bg-white px-10 py-0.5"
		bind:this={navbar}
	>
		<div class="flex flex-row gap-2">
			{#each routes as route}
				<a class="px-1 italic" href={`/${route}`}>
					/<span class="underline-highlight-2 underline hover:bg-highlight-1">
						{#if route != ''}
							{route}
						{:else}
							bisset.dev
						{/if}
					</span>
				</a>
			{/each}
		</div>

		<div class="flex w-fit flex-row items-center justify-end">
			<form class="flex w-48 flex-row gap-1 border border-border px-3 py-1" action="/search">
				<input
					type="search"
					name="search"
					placeholder="search"
					aria-label="search"
					spellcheck="false"
					class="text-text-500 w-full outline-0"
					onclick={showSearch}
				/>
				<kbd>ctrl</kbd>
				<kbd>q</kbd>
			</form>
		</div>
	</nav>
</div>

<div
	class="absolute top-0 hidden h-lvh w-lvw items-center justify-center backdrop-blur-xs"
	bind:this={search}
>
	<div class="border border-background-1 bg-white px-6 py-2">
		<div class="flex flex-col gap-2.5">
			<div class="flex flex-row gap-8">
				<form>
					<input
						bind:this={searchbar}
						placeholder="search"
						class="border-b border-background-1 outline-0"
						type="text"
					/>
				</form>
				<button
					class="cursor-pointer"
					onclick={() => {
						if (search != null) search.style.display = 'none';
					}}>Close</button
				>
			</div>
			<div>
				🚧 Under Maintenence 🚧
			</div>
		</div>
	</div>
</div>
