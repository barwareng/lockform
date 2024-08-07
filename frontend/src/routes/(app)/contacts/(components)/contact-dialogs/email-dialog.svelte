<script lang="ts">
	import { Label } from '$lib/components/ui/label';
	import { Input } from '$lib/components/ui/input';
	import BaseDialog from './base-dialog.svelte';
	import { MailIcon } from 'lucide-svelte';
	import { CONTACT, type ISaveContactBody } from '$utils/interfaces/contacts.interface';

	export let saveContactBody: ISaveContactBody = {
		contact: {
			type: CONTACT.EMAIL
		},
		teamContact: {}
	};
	let open = false;
	export let isEditing = false;
</script>

<BaseDialog title="Email" icon={MailIcon} bind:open bind:isEditing bind:saveContactBody>
	<div class="space-y-3">
		<div class="flex-1 space-y-1">
			<Label>Value</Label>
			<Input
				type="email"
				bind:value={saveContactBody.contact.value}
				placeholder="support@acme.com"
			/>
		</div>
		<div class="flex-1 space-y-1">
			<Label>Label <span class="text-muted-foreground text-[10px]">(optional)</span></Label>
			<Input
				type="text"
				bind:value={saveContactBody.contact.label}
				placeholder="Acme Corp Support"
			/>
		</div>
	</div>
</BaseDialog>
