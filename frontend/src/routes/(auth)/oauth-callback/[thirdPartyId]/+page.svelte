<script lang="ts">
	import { handleOauthCallback } from '$utils/supertokens';

	import * as Alert from '$lib/components/ui/alert';
	import { onMount } from 'svelte';
	import { LoaderCircleIcon } from 'lucide-svelte';
	onMount(() => {
		handleOauthCallback();
	});
</script>

<div class="centered flex h-screen w-full items-center justify-center">
	<Alert.Root class="max-w-sm">
		<LoaderCircleIcon class="h-4 w-4 animate-spin" />
		<Alert.Title>Verifying login!</Alert.Title>
		<Alert.Description>You're moments away from your account.</Alert.Description>
	</Alert.Root>
</div>
