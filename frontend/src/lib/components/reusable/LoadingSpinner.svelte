<script lang="ts">
	import { LoaderCircleIcon } from 'lucide-svelte';

	export let title = '';
	export let subtitle = '';
</script>

<div class="flex h-96 w-full flex-col items-center justify-center text-center {$$props.class}">
	<LoaderCircleIcon class="h-12 w-12 animate-spin" />
</div>
