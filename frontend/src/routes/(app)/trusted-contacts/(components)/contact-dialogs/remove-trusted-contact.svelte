<script lang="ts">
	import { client } from '$lib/api/Client';
	import ButtonLoadingSpinner from '$lib/components/reusable/loading-spinners/ButtonLoadingSpinner.svelte';
	import * as AlertDialog from '$lib/components/ui/alert-dialog/index.js';
	import { invalidateAll } from '$app/navigation';
	import { toastError, toastSuccess } from '$utils/toasts';
	import { Button } from '$lib/components/ui/button';
	export let id: number;
	let removingTrustedContact = false;
	const removeTrustedContact = async () => {
		try {
			removingTrustedContact = true;
			await client.trustedContacts.delete({ id });
			removingTrustedContact = false;
			invalidateAll();
			toastSuccess('Contact has been removed.');
		} catch (error) {
			removingTrustedContact = false;
			toastError(error);
		}
	};
</script>

<AlertDialog.Root>
	<AlertDialog.Trigger asChild let:builder>
		<Button
			builders={[builder]}
			variant="ghost"
			size="sm"
			class="flex w-full justify-start px-2 text-sm font-normal"
		>
			Delete
		</Button>
	</AlertDialog.Trigger>
	<AlertDialog.Content>
		<AlertDialog.Header>
			<AlertDialog.Title>Are you absolutely sure?</AlertDialog.Title>
			<AlertDialog.Description>
				This action cannot be undone. This will permanently remove this channel from your team.
			</AlertDialog.Description>
		</AlertDialog.Header>
		<AlertDialog.Footer>
			<AlertDialog.Cancel>Cancel</AlertDialog.Cancel>
			<AlertDialog.Action disabled={removingTrustedContact} on:click={removeTrustedContact}>
				<ButtonLoadingSpinner bind:state={removingTrustedContact} />
				Continue
			</AlertDialog.Action>
		</AlertDialog.Footer>
	</AlertDialog.Content>
</AlertDialog.Root>
