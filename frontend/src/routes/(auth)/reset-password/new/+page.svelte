<script lang="ts">
	// import { Icons } from '$lib/components/docs/icons/index.js';
	import { Button } from '$lib/components/ui/button';
	import * as Card from '$lib/components/ui/card';
	import { Label } from '$lib/components/ui/label';
	import { Input } from '$lib/components/ui/input';
	import { newPasswordEntered } from '$utils/supertokens';
	import { TriangleAlertIcon } from 'lucide-svelte';

	let password: string;
	let confirmPassword: string;
	let passwordErrors: string[] = [];
	const changePassword = async () => {
		passwordErrors = (await newPasswordEntered(password))!;
	};
	$: if (confirmPassword?.length && password?.length && confirmPassword !== password) {
		passwordErrors[0] = `Passwords don't match.`;
	} else {
		passwordErrors = [];
	}
</script>

<div class="mx-auto flex h-screen max-w-md items-center justify-center">
	<Card.Root>
		<Card.Header class="space-y-1">
			<Card.Title class="text-2xl">Enter new password</Card.Title>
		</Card.Header>
		<Card.Content class="grid gap-4">
			<div class="grid gap-2">
				<Label for="password">Password</Label>
				<Input id="password" type="password" bind:value={password} />
				{#if passwordErrors.length}
					{#each passwordErrors as error}
						<div class="text-destructive flex items-center gap-x-1 text-xs">
							<TriangleAlertIcon class="h-3 w-3" />
							<p class="">{error}</p>
						</div>
					{/each}
				{/if}
			</div>
			<div class="grid gap-2">
				<Label for="password">Confirm password</Label>
				<Input id="password" type="password" bind:value={confirmPassword} />
			</div>
		</Card.Content>
		<Card.Footer class="flex flex-col gap-2">
			<Button
				disabled={!password || !confirmPassword || password !== confirmPassword}
				class="w-full"
				on:click={changePassword}>Reset Password</Button
			>
		</Card.Footer>
	</Card.Root>
</div>
