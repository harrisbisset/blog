<script lang="ts">
	import { onMount } from 'svelte';

	let ctrl = $state(false);
	let q = $state(false);
	let toggled = $state(false);

	let search: HTMLElement;
	let searchbar: HTMLElement;

	onMount(() => {
		document.addEventListener('keydown', function (e) {
			if (e.ctrlKey) ctrl = true;
			if (e.key == 'q') q = true;
			if (e.key == 'Escape' && search != null) {
				search.style.display = 'none';
				toggled = false;
			}

			console.log('hello');
			console.log(ctrl);
			console.log(q);

			if (ctrl && q) {
				if (toggled) {
					if (search != null) search.style.display = 'none';
					toggled = false;
					return;
				}

				if (search != null) search.style.display = 'flex';
				searchbar.focus();
				toggled = true;
			}
		});

		document.addEventListener('keyup', function (e) {
			if (e.ctrlKey) ctrl = false;
			if (e.key == 'q') q = false;
		});
	});
</script>

<div
	class="absolute top-0 hidden h-lvh w-lvw items-center justify-center backdrop-blur-xs"
	bind:this={search}
>
	<div class="border border-background-1 bg-white px-6 py-2">
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
	</div>
</div>
