<script lang="ts">
	import Nav from '$lib/components/nav/nav.svelte';
	import '../app.css';
	// import favicon from '$lib/assets/favicon.svg';

	let { data, children } = $props();
</script>

<svelte:head>
	<!-- <link rel="icon" href={favicon} /> -->
</svelte:head>

<Nav {data} />
<div class="nav-buffer">
	{@render children?.()}
</div>

<style>
	.nav-buffer {
		padding-top: var(--navbar-height);
	}
</style>
