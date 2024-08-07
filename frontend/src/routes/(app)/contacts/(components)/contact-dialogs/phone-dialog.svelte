<script lang="ts">
	import { Label } from '$lib/components/ui/label';
	import { Input } from '$lib/components/ui/input';
	import BaseDialog from './base-dialog.svelte';
	import { PhoneIcon } from 'lucide-svelte';
	import PhoneNumber from '$lib/components/reusable/inputs/PhoneNumber.svelte';
	import { parsePhoneNumber } from 'libphonenumber-js';
	import { CONTACT, type IContact } from '$utils/interfaces/contacts.interface';
	export let contact: Partial<IContact> = {
		type: CONTACT.PHONE
	};
	let open = false;
	export let isEditing = false;
	$: if (contact.value) {
		contact.value = parsePhoneNumber(contact.value)?.formatInternational();
	}
</script>

<BaseDialog title="Phone" icon={PhoneIcon} bind:open bind:isEditing bind:contact>
	<div class="space-y-3">
		<div class="flex-1 space-y-1">
			<Label>Number</Label>
			<PhoneNumber bind:phoneNumber={contact.value} />
		</div>
		<div class="flex-1 space-y-1">
			<Label>Label <span class="text-[10px] text-muted-foreground">(optional)</span></Label>
			<Input type="text" bind:value={contact.label} placeholder="e.g Work" />
		</div>
	</div>
</BaseDialog>
