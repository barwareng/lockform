<script lang="ts">
	import * as Alert from '$lib/components/ui/alert';
	import { VITE_SUPERTOKENS_APP_NAME } from '$lib/env';
	import { sendEmailVerificationLink } from '$utils/supertokens';
	import { onMount } from 'svelte';
	onMount(async () => {
		await sendEmailVerificationLink();
	});
</script>

<div class="flex h-screen w-full items-center justify-center">
	<Alert.Root class="max-w-sm sm:max-w-md">
		<Alert.Title class="text-base">Email Verification</Alert.Title>
		<Alert.Description
			>We've sent you an email confirmation link. Please verify your email address to access the
			{VITE_SUPERTOKENS_APP_NAME} services.</Alert.Description
		>
	</Alert.Root>
</div>
