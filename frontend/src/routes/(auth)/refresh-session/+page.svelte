<script lang="ts">
	import ButtonLoadingSpinner from '$lib/components/reusable/loading-spinners/ButtonLoadingSpinner.svelte';
	import * as Alert from '$lib/components/ui/alert/index.js';
</script>

<div class="min-w-screen flex min-h-screen items-center justify-center">
	<Alert.Root variant="destructive" class="max-w-xs">
		<ButtonLoadingSpinner state />
		<Alert.Title>Session</Alert.Title>
		<Alert.Description class="mt-2 text-sm"
			>Your session is invalid or has expired. We're refreshing it.</Alert.Description
		>
	</Alert.Root>
</div>
