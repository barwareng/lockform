<script lang="ts">
	import { client } from '$lib/api/Client';
	import ButtonLoadingSpinner from '$lib/components/reusable/loading-spinners/ButtonLoadingSpinner.svelte';
	import * as AlertDialog from '$lib/components/ui/alert-dialog/index.js';
	import { invalidateAll } from '$app/navigation';
	import { toastError, toastSuccess } from '$utils/toasts';
	import { Button } from '$lib/components/ui/button';
	export let id: number;
	let removingChannel = false;
	const removeChannel = async () => {
		try {
			removingChannel = true;
			await client.channels.delete({ id });
			removingChannel = false;
			invalidateAll();
			toastSuccess('Channel has been removed.');
		} catch (error) {
			removingChannel = false;
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
			<AlertDialog.Action disabled={removingChannel} on:click={removeChannel}>
				<ButtonLoadingSpinner bind:state={removingChannel} />
				Continue
			</AlertDialog.Action>
		</AlertDialog.Footer>
	</AlertDialog.Content>
</AlertDialog.Root>
