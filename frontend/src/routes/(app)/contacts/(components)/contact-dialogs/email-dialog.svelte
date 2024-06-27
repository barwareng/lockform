<script lang="ts">
	import { Label } from '$lib/components/ui/label';
	import { Input } from '$lib/components/ui/input';
	import BaseDialog from './base-dialog.svelte';
	import { MailIcon } from 'lucide-svelte';
	import { CONTACT, type IContact } from '$utils/interfaces/contacts.interface';
	export let contact: Partial<IContact> = {
		type: CONTACT.EMAIL
	};
	let open = false;
	export let isEditing = false;
</script>

<BaseDialog title="Email" icon={MailIcon} bind:open bind:isEditing bind:contact>
	<div class="space-y-3">
		<div class="flex-1 space-y-1">
			<Label>Value</Label>
			<Input type="email" bind:value={contact.value} placeholder="mail@acme.com" />
		</div>
		<div class="flex-1 space-y-1">
			<Label>Label <span class="text-muted-foreground text-[10px]">(optional)</span></Label>
			<Input type="text" bind:value={contact.label} placeholder="e.g Work" />
		</div>
	</div>
</BaseDialog>
