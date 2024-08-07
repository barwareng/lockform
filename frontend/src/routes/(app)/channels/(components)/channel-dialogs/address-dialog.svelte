<script lang="ts">
	import { Label } from '$lib/components/ui/label';
	import { Input } from '$lib/components/ui/input';
	import BaseDialog from './base-dialog.svelte';
	import { MapPinIcon } from 'lucide-svelte';
	import { CHANNEL, type IChannel } from '$utils/interfaces/channels.interface';
	import Textarea from '$lib/components/ui/textarea/textarea.svelte';

	export let channel: Partial<IChannel> = {
		type: CHANNEL.ADDRESS,
		category: 'General',
		isPublic: true
	};
	export let isEditing = false;
	let open = false;
</script>

<BaseDialog title="Address" icon={MapPinIcon} bind:open bind:isEditing bind:channel>
	<div class="space-y-3">
		<div class="flex-1 space-y-1">
			<Label>Value</Label>
			<Textarea bind:value={channel.value} class="resize-none" />
		</div>
		<div class="flex-1 space-y-1">
			<Label>Label <span class="text-[10px] text-muted-foreground">(optional)</span></Label>
			<Input type="text" bind:value={channel.label} placeholder="e.g Work, Mailing etc." />
		</div>
	</div>
</BaseDialog>
