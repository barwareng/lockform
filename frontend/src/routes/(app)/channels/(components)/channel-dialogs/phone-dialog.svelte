<script lang="ts">
	import { Label } from '$lib/components/ui/label';
	import { Input } from '$lib/components/ui/input';
	import BaseDialog from './base-dialog.svelte';
	import { CHANNEL, type IChannel } from '$utils/interfaces/channels.interface';
	import { PhoneIcon } from 'lucide-svelte';
	import PhoneNumber from '$lib/components/reusable/inputs/PhoneNumber.svelte';
	export let channel: Partial<IChannel> = {
		type: CHANNEL.PHONE,
		category: 'General',
		isPublic: true
	};
	let open = false;
	export let isEditing = false;
</script>

<BaseDialog title="Phone" icon={PhoneIcon} bind:open bind:isEditing bind:channel>
	<div class="space-y-3">
		<div class="flex-1 space-y-1">
			<Label>Number</Label>
			<PhoneNumber bind:phoneNumber={channel.value} />
		</div>
		<div class="flex-1 space-y-1">
			<Label>Label <span class="text-muted-foreground text-[10px]">(optional)</span></Label>
			<Input type="text" bind:value={channel.label} placeholder="e.g Work" />
		</div>
	</div>
</BaseDialog>
