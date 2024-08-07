<script lang="ts">
	import { Label } from '$lib/components/ui/label';
	import { Input } from '$lib/components/ui/input';
	import BaseDialog from './base-dialog.svelte';
	import { MapPinIcon } from 'lucide-svelte';
	import Textarea from '$lib/components/ui/textarea/textarea.svelte';
	import { CONTACT, type ISaveContactBody } from '$utils/interfaces/contacts.interface';

	export let saveContactBody: ISaveContactBody = {
		contact: {
			type: CONTACT.ADDRESS
		},
		teamContact: {}
	};
	export let isEditing = false;
	let open = false;
</script>

<BaseDialog title="Address" icon={MapPinIcon} bind:open bind:isEditing bind:saveContactBody>
	<div class="space-y-3">
		<div class="flex-1 space-y-1">
			<Label>Value</Label>
			<Textarea bind:value={saveContactBody.contact.value} class="resize-none" />
		</div>
		<div class="flex-1 space-y-1">
			<Label>Label <span class="text-muted-foreground text-[10px]">(optional)</span></Label>
			<Input
				type="text"
				bind:value={saveContactBody.contact.label}
				placeholder="e.g Acme Corp Depot"
			/>
		</div>
	</div>
</BaseDialog>
