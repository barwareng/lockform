<script lang="ts">
	import { MetaTags } from 'svelte-meta-tags';
	import type { PageData } from './$types';
	import * as Popover from '$lib/components/ui/popover';
	import { VITE_APP_NAME } from '$lib/env';
	import Separator from '$lib/components/ui/separator/separator.svelte';
	import * as Table from '$lib/components/ui/table';
	import { requireRoles } from '$utils/guards';
	import ListChannelOptionsDialog from './(components)/list-channel-options-dialog.svelte';
	import { ROLE_VALUES } from '$utils/interfaces/roles.interface';
	import LoadingSpinner from '$lib/components/reusable/loading-spinners/LoadingSpinner.svelte';
	import EmptyState from '$lib/components/reusable/EmptyState.svelte';
	import { AntennaIcon, EllipsisIcon, MailIcon, MapPinIcon, PhoneIcon } from 'lucide-svelte';
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu';
	import { Button } from '$lib/components/ui/button';
	import { CHANNEL } from '$utils/interfaces/channels.interface';
	import parsePhoneNumber from 'libphonenumber-js';
	import { channelDialogs } from './(components)/channel-dialogs/dialogs';
	import EmailDialog from './(components)/channel-dialogs/email-dialog.svelte';
	import RemoveMember from './(components)/channel-dialogs/remove-member.svelte';

	export let data: PageData;
	$: ({ channels } = data);
	$: loadingChannels = data.loadingChannels ?? true;
	const getIcon = (type: CHANNEL) => {
		let icon: any;
		switch (type) {
			case CHANNEL.EMAIL:
				icon = MailIcon;
				break;
			case CHANNEL.PHONE:
				icon = PhoneIcon;
				break;
			default:
				icon = MapPinIcon;
				break;
		}
		return icon;
	};
	$: props = { channel: channels?.[0] };
</script>

<MetaTags
	title="{VITE_APP_NAME} | Channels"
	description="Manage all the channels that your team uses to communicate"
/>
<div class="relative pb-16">
	<div class="bg-background sticky top-0 z-10">
		<div class="flex flex-col items-start gap-6 md:flex-row md:justify-between">
			<div class="space-y-0.5">
				<h2 class="text-2xl font-bold tracking-tight">Channels</h2>
				<p class="text-muted-foreground text-sm">
					Manage all the channels that your team uses to communicate
				</p>
			</div>
			{#if requireRoles([ROLE_VALUES.ADMIN, ROLE_VALUES.OWNER])}
				<ListChannelOptionsDialog />
			{/if}
		</div>
		<Separator class="my-6" />
	</div>
	{#if loadingChannels}
		<LoadingSpinner />
	{:else if channels?.length}
		<Table.Root>
			<Table.Header>
				<Table.Row>
					<Table.Head class="hidden w-[100px] sm:table-cell">Channel</Table.Head>
					<Table.Head>Value</Table.Head>
					<Table.Head class="hidden md:table-cell">Label</Table.Head>
					<Table.Head class="hidden md:table-cell">Category</Table.Head>
					<Table.Head class="hidden md:table-cell">Public</Table.Head>
					<Table.Head class="hidden md:table-cell">URL</Table.Head>
					{#if requireRoles([ROLE_VALUES.OWNER, ROLE_VALUES.ADMIN])}
						<Table.Head class="sr-only">Action</Table.Head>
					{/if}
				</Table.Row>
			</Table.Header>
			<Table.Body>
				{#each channels as channel}
					{@const icon = getIcon(channel.type)}
					{@const dialog = channelDialogs
						.find((c) => c.category == channel.category)
						?.dialogs.find((d) => d.type == channel.type)}
					<Table.Row>
						<Table.Cell class="hidden sm:table-cell">
							<svelte:component this={icon} class="h-6 w-6" />
						</Table.Cell>
						{#if channel.type == CHANNEL.PHONE}
							<Table.Cell class="font-medium"
								>{parsePhoneNumber(channel.value)?.formatInternational()}</Table.Cell
							>
						{:else}
							<Table.Cell class="font-medium">{channel.value}</Table.Cell>
						{/if}
						<Table.Cell class="hidden md:table-cell">{channel.label || '--'}</Table.Cell>
						<Table.Cell class="hidden md:table-cell">{channel.category || '--'}</Table.Cell>
						<Table.Cell class="hidden md:table-cell">{channel.isPublic ? 'Yes' : 'No'}</Table.Cell>
						<Table.Cell class="hidden md:table-cell">{channel.url || '--'}</Table.Cell>
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
										{#if dialog?.component && channel}
											<svelte:component this={dialog.component} {channel} isEditing />
										{/if}
										<!-- <DropdownMenu.Item>Delete</DropdownMenu.Item> -->
										<RemoveMember id={channel.id} />
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
			title="No channels yet"
			subtitle="When you add your team's channels, they will show up here."
		>
			<div slot="leading-icon"><AntennaIcon class="h-12 w-12" /></div>
			{#if requireRoles([ROLE_VALUES.ADMIN, ROLE_VALUES.OWNER])}
				<ListChannelOptionsDialog />
			{/if}
		</EmptyState>
	{/if}
</div>
