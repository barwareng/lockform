<script>
	import { LoaderCircleIcon } from 'lucide-svelte';
	export let state = false;
</script>

{#if state}
	<LoaderCircleIcon class="h-4 w-4 animate-spin" />
{/if}
