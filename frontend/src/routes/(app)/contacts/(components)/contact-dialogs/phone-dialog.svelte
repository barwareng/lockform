<script lang="ts">
	import { Label } from '$lib/components/ui/label';
	import { Input } from '$lib/components/ui/input';
	import BaseDialog from './base-dialog.svelte';
	import { PhoneIcon } from 'lucide-svelte';
	import PhoneNumber from '$lib/components/reusable/inputs/PhoneNumber.svelte';
	import { parsePhoneNumber } from 'libphonenumber-js';
	import { CONTACT, type ISaveContactBody } from '$utils/interfaces/contacts.interface';
	export let saveContactBody: ISaveContactBody = {
		contact: {
			type: CONTACT.PHONE
		},
		teamContact: {}
	};
	let open = false;
	export let isEditing = false;
	$: if (saveContactBody.contact.value) {
		saveContactBody.contact.value = parsePhoneNumber(
			saveContactBody.contact.value
		)?.formatInternational();
	}
</script>

<BaseDialog title="Phone" icon={PhoneIcon} bind:open bind:isEditing bind:saveContactBody>
	<div class="space-y-3">
		<div class="flex-1 space-y-1">
			<Label>Number</Label>
			<PhoneNumber bind:phoneNumber={saveContactBody.contact.value} />
		</div>
		<div class="flex-1 space-y-1">
			<Label>Label <span class="text-muted-foreground text-[10px]">(optional)</span></Label>
			<Input
				type="text"
				bind:value={saveContactBody.contact.label}
				placeholder="e.g Luther Corp Front Desk"
			/>
		</div>
	</div>
</BaseDialog>
