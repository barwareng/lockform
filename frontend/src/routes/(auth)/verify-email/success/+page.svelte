<script lang="ts">
	import * as Alert from '$lib/components/ui/alert';
	import { Button } from '$lib/components/ui/button';
</script>

<div class="flex h-screen w-full items-center justify-center">
	<Alert.Root class="max-w-sm sm:max-w-md">
		<Alert.Title class="text-base">Email Verified</Alert.Title>
		<Alert.Description
			>Thank you for verifying your email address. Please proceed to the dashboard to finalize your
			profile setup.</Alert.Description
		>
		<div class="mt-4 flex w-full justify-end gap-x-2">
			<Button href="/">Go to dashboard</Button>
		</div>
	</Alert.Root>
</div>
