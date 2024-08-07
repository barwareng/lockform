<script lang="ts">
	import * as Alert from '$lib/components/ui/alert';
	import { Button } from '$lib/components/ui/button';
	import { consumeVerificationCode, sendEmailVerificationLink } from '$utils/supertokens';

	import { onMount } from 'svelte';
	onMount(async () => {
		await consumeVerificationCode();
	});
</script>

<div class="flex h-screen w-full items-center justify-center">
	<Alert.Root variant="destructive" class="max-w-sm sm:max-w-md">
		<Alert.Title class="text-base">Email Verification Error</Alert.Title>
		<Alert.Description>Verification link is expired or token is invalid</Alert.Description>
		<div class="mt-4 flex w-full justify-end">
			<Button on:click={sendEmailVerificationLink}>Resend Verification Link</Button>
		</div>
	</Alert.Root>
</div>
