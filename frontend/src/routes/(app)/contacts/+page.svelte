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
	import {
		ContactIcon,
		EllipsisIcon,
		InfoIcon,
		MailIcon,
		MapPinIcon,
		PhoneIcon
	} from 'lucide-svelte';
	import { Badge } from '$lib/components/ui/badge';
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu';
	import { Button } from '$lib/components/ui/button';
	import parsePhoneNumber from 'libphonenumber-js';
	import ListContactOptionsDialog from './(components)/list-contact-options-dialog.svelte';
	import { CONTACT } from '$utils/interfaces/contacts.interface';
	import * as Tooltip from '$lib/components/ui/tooltip';
	import UpdateTrustworthiness from './(components)/contact-dialogs/update-trustworthiness.svelte';
	import Pagination from '$lib/components/reusable/navigation/Pagination.svelte';

	export let data: PageData;
	$: ({ contacts } = data);
	$: loadingContacts = data.loading ?? true;
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
		<div class="space-y-4">
			<Table.Root>
				<Table.Header>
					<Table.Row>
						<Table.Head>Contact</Table.Head>
						<Table.Head>Trusted</Table.Head>
						<Table.Head class="hidden sm:table-cell">Added By</Table.Head>
						<Table.Head class="hidden md:table-cell">Domain</Table.Head>
						<Table.Head class="hidden md:table-cell">URL</Table.Head>
						{#if requireRoles([ROLE_VALUES.OWNER, ROLE_VALUES.ADMIN])}
							<Table.Head class="sr-only">Action</Table.Head>
						{/if}
					</Table.Row>
				</Table.Header>
				<Table.Body>
					{#each contacts as contact}
						{@const icon = getIcon(contact.type)}
						<Table.Row>
							<Table.Cell class="flex items-center gap-x-2">
								<div>
									<svelte:component this={icon} class="h-6 !w-6" />
								</div>
								<div>
									<p class="whitespace-nowrap">
										{contact.label || ''}
									</p>
									<p class="text-xs opacity-75">
										{#if contact.type == CONTACT.PHONE}
											{parsePhoneNumber(contact.value)?.formatInternational()}
										{:else}
											{contact.value}
										{/if}
									</p>
								</div>
							</Table.Cell>

							<Table.Cell>
								{#if contact.isTrusted}
									<Badge>Yes</Badge>
								{:else if !contact.isTrusted && contact.reasonForUntrusting}
									<div class="flex items-center gap-x-1">
										<Tooltip.Root>
											<Tooltip.Trigger><Badge variant="destructive">No</Badge></Tooltip.Trigger>
											<Tooltip.Content>
												<p class="text-xs">{contact.reasonForUntrusting}</p>
											</Tooltip.Content>
										</Tooltip.Root>
									</div>
								{:else}
									<Badge variant="destructive">No</Badge>
								{/if}
							</Table.Cell>
							<Table.Cell class="hidden sm:table-cell">
								<div>
									<p class="whitespace-nowrap">
										{contact.addedBy?.name || ''}
									</p>
									<p class="text-xs opacity-75">
										{contact.addedBy?.email}
									</p>
								</div>
							</Table.Cell>
							<Table.Cell class="hidden md:table-cell">{contact.domain || '--'}</Table.Cell>
							<Table.Cell class="hidden md:table-cell">{contact.url || '--'}</Table.Cell>
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
											<UpdateTrustworthiness
												id={contact.id}
												bind:trustworthiness={contact.isTrusted}
											/>
											<!-- {#if dialog?.component && contact}
											<svelte:component this={dialog.component} {contact} isEditing />
										{/if} -->
											<!-- <RemoveContact id={contact.id} /> -->
										</DropdownMenu.Content>
									</DropdownMenu.Root>
								</Table.Cell>
							{/if}
						</Table.Row>
					{/each}
				</Table.Body>
			</Table.Root>
			<Pagination />
		</div>
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
