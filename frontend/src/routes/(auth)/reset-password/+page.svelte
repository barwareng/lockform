<script lang="ts">
	import { Button } from '$lib/components/ui/button';
	import * as Card from '$lib/components/ui/card';
	import { Input } from '$lib/components/ui/input';
	import { sendResetPasswordLink } from '$utils/supertokens';
	import { TriangleAlertIcon } from 'lucide-svelte';

	let email: string;
	let emailErrors: string[] = [];
	const sendLink = async () => {
		emailErrors = (await sendResetPasswordLink(email))!;
	};
</script>

<div class="mx-auto flex h-screen max-w-md items-center justify-center">
	<Card.Root>
		<Card.Header class="space-y-1">
			<Card.Title class="text-2xl">Reset your password</Card.Title>
			<Card.Description
				>Please enter your email address to get a password reset link.</Card.Description
			>
		</Card.Header>
		<Card.Content class="grid gap-4">
			<div class="grid gap-2">
				<Input id="email" type="email" bind:value={email} placeholder="m@example.com" />
				{#if emailErrors?.length}
					{#each emailErrors as error}
						<div class="text-destructive flex items-center gap-x-1 text-xs">
							<TriangleAlertIcon class="h-3 w-3" />
							<p class="">{error}</p>
						</div>
					{/each}
				{/if}
			</div>
		</Card.Content>
		<Card.Footer class="flex flex-col gap-2">
			<Button disabled={!email} class="w-full" on:click={sendLink}>Send Reset Link</Button>
		</Card.Footer>
	</Card.Root>
</div>
