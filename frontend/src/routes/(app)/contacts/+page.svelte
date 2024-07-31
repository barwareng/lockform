<script lang="ts">
	import { MetaTags } from 'svelte-meta-tags';
	import type { PageData } from './$types';
	import { VITE_APP_NAME } from '$lib/env';
	import Separator from '$lib/components/ui/separator/separator.svelte';
	import * as Table from '$lib/components/ui/table';
	import { requireRoles } from '$utils/guards';
	import { ROLE_VALUES } from '$utils/interfaces/roles.interface';
	import LoadingSpinner from '$lib/components/reusable/loading-spinners/LoadingSpinner.svelte';
	import EmptyState from '$lib/components/reusable/EmptyState.svelte';
	import { ContactIcon, EllipsisIcon, MailIcon, MapPinIcon, PhoneIcon } from 'lucide-svelte';
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu';
	import { Button } from '$lib/components/ui/button';
	import parsePhoneNumber from 'libphonenumber-js';
	import ListContactOptionsDialog from './(components)/list-contact-options-dialog.svelte';
	import { CONTACT } from '$utils/interfaces/contacts.interface';
	import RemoveContact from './(components)/contact-dialogs/remove-contact.svelte';
	import { contactDialogs } from './(components)/contact-dialogs/dialogs';

	export let data: PageData;
	$: ({ contacts } = data);
	$: loadingContacts = data.loadingContacts ?? true;
	const getIcon = (type: CONTACT) => {
		let icon: any;
		switch (type) {
			case CONTACT.EMAIL:
				icon = MailIcon;
				break;
			case CONTACT.PHONE:
				icon = PhoneIcon;
				break;
			default:
				icon = MapPinIcon;
				break;
		}
		return icon;
	};
</script>

<MetaTags
	title="{VITE_APP_NAME} | Contacts"
	description="Manage all your organization's contacts. This includes those add from add-ons from different workspaces like GMail, Outlook, etc."
/>
<div class="relative pb-16">
	<div class="bg-background sticky top-0 z-10">
		<div class="flex flex-col items-start gap-6 md:flex-row md:justify-between">
			<div class="space-y-0.5">
				<h2 class="text-2xl font-bold tracking-tight">Contacts</h2>
				<p class="text-muted-foreground max-w-lg text-sm">
					Manage all your organization's contacts. This includes those add from add-ons from
					different workspaces like GMail, Outlook etc.
				</p>
			</div>
			{#if requireRoles([ROLE_VALUES.ADMIN, ROLE_VALUES.OWNER])}
				<ListContactOptionsDialog />
			{/if}
		</div>
		<Separator class="my-6" />
	</div>
	{#if loadingContacts}
		<LoadingSpinner />
	{:else if contacts?.length}
		<Table.Root>
			<Table.Header>
				<Table.Row>
					<Table.Head class="hidden w-[100px] sm:table-cell">Contact</Table.Head>
					<Table.Head>Value</Table.Head>
					<Table.Head class="hidden md:table-cell">Label</Table.Head>
					<Table.Head class="hidden md:table-cell">Domain</Table.Head>
					<Table.Head class="hidden md:table-cell">Added By</Table.Head>

					{#if requireRoles([ROLE_VALUES.OWNER, ROLE_VALUES.ADMIN])}
						<Table.Head class="sr-only">Action</Table.Head>
					{/if}
				</Table.Row>
			</Table.Header>
			<Table.Body>
				{#each contacts as contact}
					{@const icon = getIcon(contact.type)}
					{@const dialog = contactDialogs.find((d) => d.type == contact.type)}
					<Table.Row>
						<Table.Cell class="hidden sm:table-cell">
							<svelte:component this={icon} class="h-6 w-6" />
						</Table.Cell>
						{#if contact.type == CONTACT.PHONE}
							<Table.Cell class="font-medium"
								>{parsePhoneNumber(contact.value)?.formatInternational()}</Table.Cell
							>
						{:else}
							<Table.Cell class="font-medium">{contact.value}</Table.Cell>
						{/if}
						<Table.Cell class="hidden md:table-cell">{contact.label || '--'}</Table.Cell>
						<Table.Cell class="hidden md:table-cell">{contact.domain || '--'}</Table.Cell>
						<Table.Cell class="hidden md:table-cell">{contact.addedBy?.name || ''}</Table.Cell>
						{#if requireRoles([ROLE_VALUES.OWNER, ROLE_VALUES.ADMIN])}
							<Table.Cell class="text-right">
								<DropdownMenu.Root>
									<DropdownMenu.Trigger asChild let:builder>
										<Button aria-haspopup="true" size="icon" variant="ghost" builders={[builder]}>
											<EllipsisIcon class="h-4 w-4" />
											<span class="sr-only">Toggle menu</span>
										</Button>
									</DropdownMenu.Trigger>
									<DropdownMenu.Content align="end">
										<DropdownMenu.Label>Actions</DropdownMenu.Label>
										{#if dialog?.component && contact}
											<svelte:component this={dialog.component} {contact} isEditing />
										{/if}
										<RemoveContact id={contact.id} />
									</DropdownMenu.Content>
								</DropdownMenu.Root>
							</Table.Cell>
						{/if}
					</Table.Row>
				{/each}
			</Table.Body>
		</Table.Root>
	{:else}
		<EmptyState
			title="No contacts yet"
			subtitle="When you add your team's contacts, they will show up here."
		>
			<div slot="leading-icon"><ContactIcon class="h-12 w-12" /></div>
			{#if requireRoles([ROLE_VALUES.ADMIN, ROLE_VALUES.OWNER])}
				<ListContactOptionsDialog />
			{/if}
		</EmptyState>
	{/if}
</div>
