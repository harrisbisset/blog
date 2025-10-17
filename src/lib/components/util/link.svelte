<script lang="ts">
	const { href, text } = $props();
</script>

<h3 class="text-lg underline decoration-1" title={href}>
	<a target="_blank" rel="noopener noreferrer" {href}>{text}</a>
</h3>
