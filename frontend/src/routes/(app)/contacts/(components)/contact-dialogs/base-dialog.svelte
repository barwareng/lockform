<script lang="ts">
	import * as Popover from '$lib/components/ui/popover';
	import { Button } from '$lib/components/ui/button';
	import Separator from '$lib/components/ui/separator/separator.svelte';
	import { Switch } from '$lib/components/ui/switch';
	import { Label } from '$lib/components/ui/label';
	import ButtonLoadingSpinner from '$lib/components/reusable/loading-spinners/ButtonLoadingSpinner.svelte';
	import { client } from '$lib/api/Client';
	import { invalidateAll } from '$app/navigation';
	import { toastError, toastSuccess } from '$utils/toasts';
	import { CONTACT, type ISaveContactBody } from '$utils/interfaces/contacts.interface';
	import { Textarea } from '$lib/components/ui/textarea';
	export let icon: any;
	export let title: string;
	export let open = false;
	export let isEditing = false;
	let isSaving = false;
	export let saveContactBody: ISaveContactBody = { contact: {}, teamContact: { isTrusted: false } };
	const saveContact = async () => {
		try {
			isSaving = true;
			if (saveContactBody.contact.type == CONTACT.PHONE)
				saveContactBody.contact.value = saveContactBody.contact.value?.replaceAll(' ', '');
			await client.contacts.create(saveContactBody);
			await invalidateAll();
			isSaving = false;
			open = false;
			toastSuccess('Contact successfully added.');
		} catch (error) {
			isSaving = false;
			toastError(error);
		}
	};
</script>

<Popover.Root bind:open>
	<Popover.Trigger asChild let:builder>
		{#if isEditing}
			<Button
				builders={[builder]}
				variant="ghost"
				size="sm"
				class="flex w-full justify-start px-2 text-sm font-normal"
			>
				Edit
			</Button>
		{:else}
			<Button
				builders={[builder]}
				variant="outline"
				class="flex aspect-square min-h-20 flex-col gap-y-2 text-xs"
			>
				<svelte:component this={icon} class="!h-8 !w-8" />
				{title}
			</Button>
		{/if}
	</Popover.Trigger>
	<Popover.Content class="relative w-80">
		<div class="flex items-center justify-between">
			<h4 class="font-medium leading-none">{title}</h4>
		</div>
		<Separator class="mt-4" />
		<div class="space-y-4 py-4">
			<slot />
			<Separator class="mt-4" />
			<div class="space-y-2">
				<div class="flex items-center space-x-2">
					<Label for="is-public" class="text-xs">To be trusted?</Label>
					<Switch
						includeInput
						id="is-public"
						bind:checked={saveContactBody.teamContact.isTrusted}
					/>
				</div>
				{#if !saveContactBody.teamContact.isTrusted}
					<div class="flex-1 space-y-1">
						<Label class="text-xs">Why?</Label>
						<Textarea
							class="resize-none placeholder:text-xs"
							bind:value={saveContactBody.teamContact.reasonForUntrusting}
							placeholder="Reason not to trust this contact"
						/>
					</div>
				{/if}
			</div>
			<Separator class="mt-4" />
			<div class="flex items-center justify-end gap-x-4">
				<Button variant="secondary" on:click={() => (open = false)}>Cancel</Button>
				<Button disabled={isSaving} on:click={saveContact}>
					<ButtonLoadingSpinner bind:state={isSaving} />
					Save
				</Button>
			</div>
		</div>
	</Popover.Content>
</Popover.Root>
