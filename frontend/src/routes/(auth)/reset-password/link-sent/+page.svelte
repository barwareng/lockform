<script lang="ts">
	import type { PageData } from './$types';

	import * as Alert from '$lib/components/ui/alert';
	export let data: PageData;
</script>

<div class="flex h-screen w-full items-center justify-center">
	<Alert.Root class="max-w-sm sm:max-w-md">
		<Alert.Title class="text-base">Password Reset</Alert.Title>
		<Alert.Description
			>We've sent you a password reset link. Please click it to continue resetting your password</Alert.Description
		>
	</Alert.Root>
</div>
