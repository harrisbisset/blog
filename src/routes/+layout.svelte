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
<div class="pt-24 md:pt-[100px] md:px-10">
	{@render children?.()}
</div>