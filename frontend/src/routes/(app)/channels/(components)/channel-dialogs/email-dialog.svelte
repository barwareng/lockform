<script lang="ts">
	import { Label } from '$lib/components/ui/label';
	import { Input } from '$lib/components/ui/input';
	import BaseDialog from './base-dialog.svelte';
	import { MailIcon } from 'lucide-svelte';
	import { CHANNEL, type IChannel } from '$utils/interfaces/channels.interface';
	export let channel: Partial<IChannel> = {
		type: CHANNEL.EMAIL,
		category: 'General',
		isPublic: true
	};
	let open = false;
	export let isEditing = false;
</script>

<BaseDialog title="Email" icon={MailIcon} bind:open bind:isEditing bind:channel>
	<div class="space-y-3">
		<div class="flex-1 space-y-1">
			<Label>Value</Label>
			<Input type="email" bind:value={channel.value} placeholder="mail@acme.com" />
		</div>
		<div class="flex-1 space-y-1">
			<Label>Label <span class="text-[10px] text-muted-foreground">(optional)</span></Label>
			<Input type="text" bind:value={channel.label} placeholder="e.g Work" />
		</div>
	</div>
</BaseDialog>
