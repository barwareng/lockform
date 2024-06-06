<script lang="ts">
	import { Label } from '$lib/components/ui/label';
	import { Input } from '$lib/components/ui/input';
	import BaseDialog from './base-dialog.svelte';
	import { MapPinIcon } from 'lucide-svelte';
	import Textarea from '$lib/components/ui/textarea/textarea.svelte';
	import { CONTACT, type ITrustedContact } from '$utils/interfaces/trusted-contacts.interface';

	export let trustedContact: Partial<ITrustedContact> = {
		type: CONTACT.ADDRESS
	};
	export let isEditing = false;
	let open = false;
</script>

<BaseDialog title="Address" icon={MapPinIcon} bind:open bind:isEditing bind:trustedContact>
	<div class="space-y-3">
		<div class="flex-1 space-y-1">
			<Label>Value</Label>
			<Textarea bind:value={trustedContact.value} class="resize-none" />
		</div>
		<div class="flex-1 space-y-1">
			<Label>Label <span class="text-muted-foreground text-[10px]">(optional)</span></Label>
			<Input type="text" bind:value={trustedContact.label} placeholder="e.g Work, Mailing etc." />
		</div>
	</div>
</BaseDialog>
