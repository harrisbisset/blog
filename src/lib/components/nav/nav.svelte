<script lang="ts">
	import { onMount } from 'svelte';

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

	let search;

	const routes = ['', 'blog', 'articles', 'projects'];
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
				/>
				<kbd>ctrl</kbd>
				<kbd>q</kbd>
			</form>
		</div>
	</nav>
</div>
