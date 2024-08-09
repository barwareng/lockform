<script lang="ts">
	import { client } from '$lib/api/Client';
	import ButtonLoadingSpinner from '$lib/components/reusable/loading-spinners/ButtonLoadingSpinner.svelte';
	import * as AlertDialog from '$lib/components/ui/alert-dialog/index.js';
	import { invalidateAll } from '$app/navigation';
	import { toastError, toastSuccess } from '$utils/toasts';
	import { Button } from '$lib/components/ui/button';
	import { Textarea } from '$lib/components/ui/textarea';
	import { Label } from '$lib/components/ui/label';
	export let id: number;
	export let isTrusted: boolean;
	let updatingTrustworthiness = false;
	const updateTrustworthiness = async () => {
		try {
			updatingTrustworthiness = true;
			// await client.contacts.delete({ id });
			updatingTrustworthiness = false;
			invalidateAll();
			toastSuccess(`Contact has been marked ${isTrusted ? 'untrusted' : 'trusted'}.`);
		} catch (error) {
			updatingTrustworthiness = false;
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
			{isTrusted ? 'Untrust' : 'Trust'}
		</Button>
	</AlertDialog.Trigger>
	<AlertDialog.Content>
		<AlertDialog.Header>
			<AlertDialog.Title>Are you absolutely sure?</AlertDialog.Title>
			<AlertDialog.Description>
				<!-- TODO better wording -->
				This action will change how this contact is displayed to team.
			</AlertDialog.Description>
		</AlertDialog.Header>
		{#if isTrusted}
			<div class="flex-1 space-y-1">
				<Label class="text-xs">Why?</Label>
				<Textarea
					class="resize-none placeholder:text-xs"
					placeholder="Reason not to trust this contact"
				/>
			</div>
		{/if}
		<AlertDialog.Footer>
			<AlertDialog.Cancel>Cancel</AlertDialog.Cancel>
			<AlertDialog.Action disabled={updatingTrustworthiness} on:click={updateTrustworthiness}>
				<ButtonLoadingSpinner bind:state={updatingTrustworthiness} />
				Continue
			</AlertDialog.Action>
		</AlertDialog.Footer>
	</AlertDialog.Content>
</AlertDialog.Root>
