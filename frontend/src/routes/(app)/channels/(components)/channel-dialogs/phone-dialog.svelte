<script lang="ts">
	import { Label } from '$lib/components/ui/label';
	import { Button } from '$lib/components/ui/button';
	import { Input } from '$lib/components/ui/input';
	import BaseDialog from './base-dialog.svelte';
	import { CHANNEL, type IChannel } from '$utils/interfaces/channels.interface';
	import { toastError, toastSuccess } from '$utils/toasts';
	import ButtonLoadingSpinner from '$lib/components/reusable/loading-spinners/ButtonLoadingSpinner.svelte';

	import { client } from '$lib/api/Client';
	import { invalidateAll } from '$app/navigation';
	import { PhoneIcon } from 'lucide-svelte';
	import PhoneNumber from '$lib/components/reusable/inputs/PhoneNumber.svelte';
	let channel: Partial<IChannel> = {};
	let addingChannel = false;
	let open = false;

	let phoneNumber: string;

	const addChannel = async () => {
		try {
			addingChannel = true;
			channel.value = phoneNumber?.replaceAll(' ', '');
			channel.type = CHANNEL.PHONE;
			channel.category = 'General';
			await client.channels.create(channel);
			await invalidateAll();
			// TODO API call
			addingChannel = false;
			open = false;
			toastSuccess('Channel successfully added.');
		} catch (error) {
			addingChannel = false;
			toastError(error);
		}
	};
</script>

<BaseDialog title="Phone" icon={PhoneIcon} bind:open bind:isPublic={channel.isPublic}>
	<div class="space-y-3">
		<div class="flex-1 space-y-1">
			<Label>Number</Label>
			<PhoneNumber bind:phoneNumber />
			<!-- <Input type="email" bind:value={channel.value} placeholder="mail@acme.com" /> -->
		</div>
		<div class="flex-1 space-y-1">
			<Label>Label <span class="text-muted-foreground text-[10px]">(optional)</span></Label>
			<Input type="text" bind:value={channel.label} placeholder="e.g Work" />
		</div>
		<div class="flex items-center justify-end gap-x-4">
			<Button variant="secondary" on:click={() => (open = false)}>Cancel</Button>
			<Button disabled={addingChannel} on:click={addChannel}>
				<ButtonLoadingSpinner bind:state={addingChannel} />
				Save
			</Button>
		</div>
	</div>
</BaseDialog>
