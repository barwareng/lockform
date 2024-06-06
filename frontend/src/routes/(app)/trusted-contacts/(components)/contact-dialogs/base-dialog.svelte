<script lang="ts">
	import * as Popover from '$lib/components/ui/popover';
	import { Button } from '$lib/components/ui/button';
	import Separator from '$lib/components/ui/separator/separator.svelte';
	import ButtonLoadingSpinner from '$lib/components/reusable/loading-spinners/ButtonLoadingSpinner.svelte';
	import { client } from '$lib/api/Client';
	import { invalidateAll } from '$app/navigation';
	import { toastError, toastSuccess } from '$utils/toasts';
	import { CONTACT, type ITrustedContact } from '$utils/interfaces/trusted-contacts.interface';
	export let icon: any;
	export let title: string;
	export let open = false;
	export let isEditing = false;
	let isSaving = false;
	export let trustedContact: Partial<ITrustedContact> = {};
	const saveContact = async () => {
		try {
			isSaving = true;
			if (trustedContact.type == CONTACT.PHONE)
				trustedContact.value = trustedContact.value?.replaceAll(' ', '');
			await client.trustedContacts.create(trustedContact);
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
