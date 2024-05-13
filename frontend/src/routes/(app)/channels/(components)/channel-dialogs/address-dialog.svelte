<script lang="ts">
	import * as Popover from '$lib/components/ui/popover';
	import { Label } from '$lib/components/ui/label';
	import { Button } from '$lib/components/ui/button';
	import { Input } from '$lib/components/ui/input';
	import BaseDialog from './base-dialog.svelte';
	import { MapPinIcon } from 'lucide-svelte';
	import { CHANNEL, type IChannel } from '$utils/interfaces/channels.interface';
	import { toastError, toastSuccess } from '$utils/toasts';
	import ButtonLoadingSpinner from '$lib/components/reusable/loading-spinners/ButtonLoadingSpinner.svelte';
	import Textarea from '$lib/components/ui/textarea/textarea.svelte';
	import { client } from '$lib/api/Client';
	import { invalidateAll } from '$app/navigation';
	let channel: Partial<IChannel> = {};
	let addingChannel = false;
	let open = false;
	const addChannel = async () => {
		try {
			addingChannel = true;
			channel.type = CHANNEL.PHONE;
			await client.channels.create(channel);
			await invalidateAll();
			addingChannel = false;
			open = false;
			toastSuccess('Channel successfully added.');
		} catch (error) {
			addingChannel = false;
			toastError(error);
		}
	};
</script>

<BaseDialog title="Address" icon={MapPinIcon} bind:open>
	<div class="space-y-3">
		<div class="flex-1 space-y-1">
			<Label>Value</Label>
			<Textarea bind:value={channel.value} class="resize-none" />
		</div>
		<div class="flex-1 space-y-1">
			<Label>Label <span class="text-muted-foreground text-[10px]">(optional)</span></Label>
			<Input type="text" bind:value={channel.label} placeholder="e.g Work, Mailing etc." />
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
