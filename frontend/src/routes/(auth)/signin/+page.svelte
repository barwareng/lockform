<script lang="ts">
	// import { Icons } from '$lib/components/docs/icons/index.js';
	import { Button } from '$lib/components/ui/button';
	import * as Card from '$lib/components/ui/card';
	import { Label } from '$lib/components/ui/label';
	import { Input } from '$lib/components/ui/input';
	import { oauthLogin, signinWithEmailAndPassword } from '$utils/supertokens';
	import Google from '$lib/icons/Google.svelte';
	import { TriangleAlertIcon } from 'lucide-svelte';

	let email: string;
	let password: string;
	let emailErrors: string[] = [];
	let passwordErrors: string[] = [];
	const signIn = async () => {
		const res = await signinWithEmailAndPassword(email, password);
		emailErrors = res?.emailErrors!;
		passwordErrors = res?.passwordErrors!;
	};
</script>

<div class="mx-auto flex h-screen max-w-md items-center justify-center">
	<Card.Root>
		<Card.Header class="space-y-1">
			<Card.Title class="text-2xl">Signin to your account</Card.Title>
			<!-- <Card.Description>Enter your email below to create your account</Card.Description> -->
		</Card.Header>
		<Card.Content class="grid gap-4">
			<!-- <Button variant="outline" on:click={() => oauthLogin('google')}>
				<Google class="mr-2 h-4 w-4" />
				Google
			</Button>
			<div class="relative">
				<div class="absolute inset-0 flex items-center">
					<span class="w-full border-t" />
				</div>
				<div class="relative flex justify-center text-xs uppercase">
					<span class="bg-card px-2 text-muted-foreground"> Or continue with </span>
				</div>
			</div> -->
			<div class="grid gap-2">
				<Label for="email">Email</Label>
				<Input id="email" type="email" bind:value={email} placeholder="m@example.com" />
				{#if emailErrors.length}
					{#each emailErrors as error}
						<div class="text-destructive flex items-center gap-x-1 text-xs">
							<TriangleAlertIcon class="h-3 w-3" />
							<p class="">{error}</p>
						</div>
					{/each}
				{/if}
			</div>
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
			<div class="flex justify-end">
				<Button href="/reset-password" variant="link" class="-my-6 text-right text-xs"
					>Forgot password?</Button
				>
			</div>
		</Card.Content>
		<Card.Footer class="flex flex-col gap-2">
			<Button disabled={!password || !email} class="w-full" on:click={signIn}>Sign in</Button>

			<Button href="/signup" variant="link">Or Sign Up</Button>
		</Card.Footer>
	</Card.Root>
</div>
