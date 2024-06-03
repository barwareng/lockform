<script lang="ts">
	import WipPlaceHolder from '$lib/components/reusable/WIPPlaceHolder.svelte';
	import type { PageData } from './$types';
	import * as Card from '$lib/components/ui/card';
	import Button from '$lib/components/ui/button/button.svelte';
	import { page } from '$app/stores';
	import { parseSearchParams } from '$utils';
	import type { IOauthAuthorizationRequest } from '$utils/interfaces/oauth.interface';
	import { client } from '$lib/api/Client';
	import { toastError } from '$utils/toasts';
	export let data: PageData;

	$: oauthAuthorizationRequestParams = parseSearchParams<IOauthAuthorizationRequest>(
		$page.url.searchParams
	);
	const sendOauthAuthorizationRequest = async () => {
		try {
			await client.oauth.authorize(oauthAuthorizationRequestParams);
		} catch (error) {
			toastError(error);
		}
	};
</script>

<div class="flex h-screen w-full items-center justify-center">
	<Card.Root class="min-w-96 max-w-screen-sm">
		<Card.Header>
			<Card.Title>Account Access</Card.Title>
			<Card.Description>Lockform G-Mail Addon Like to access your account</Card.Description>
		</Card.Header>

		<Card.Footer class="justify-end gap-x-4">
			<Button variant="secondary">Deny</Button>
			<Button on:click={sendOauthAuthorizationRequest}>Allow</Button>
		</Card.Footer>
	</Card.Root>
</div>
