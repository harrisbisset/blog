<script lang="ts">
	import Nav from '$lib/components/nav/nav.svelte';
	import Search from '$lib/components/nav/search.svelte';
	import '../app.css';
	// import favicon from '$lib/assets/favicon.svg';

	let { children } = $props();
</script>

<svelte:head>
	<!-- <link rel="icon" href={favicon} /> -->
</svelte:head>

<Nav />
<div class="nav-buffer">
	{@render children?.()}
</div>
<Search />

<style>
	.nav-buffer {
		padding-top: var(--navbar-height);
	}
</style>
