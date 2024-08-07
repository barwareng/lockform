<script lang="ts">
	import { invalidateAll } from '$app/navigation';
	import { client } from '$lib/api/Client';
	import Avatar from '$lib/components/reusable/images/Avatar.svelte';
	import Badge from '$lib/components/ui/badge/badge.svelte';
	import { requireRoles } from '$utils/guards';
	import { ROLE_VALUES } from '$utils/interfaces/roles.interface';
	import type { IMember } from '$utils/interfaces/users.interfaces';
	import { toastError } from '$utils/toasts';
	import * as Command from '$lib/components/ui//command';
	import * as Popover from '$lib/components/ui//popover';
	import { Button } from '$lib/components/ui/button';
	import { ChevronDownIcon, LoaderCircle } from 'lucide-svelte';
	export let member: IMember;
	import { page } from '$app/stores';
	import { closeAndRefocusTrigger } from '$utils';
	import { ROLES } from '$utils/constants/roles.constants';
	import RemovalConfirmation from './removal-confirmation.svelte';
	import ButtonLoadingSpinner from '$lib/components/reusable/loading-spinners/ButtonLoadingSpinner.svelte';
	let updatingRole = false;
	const updateRole = async (userId: string, email: string, role: string) => {
		try {
			updatingRole = true;
			await client.members.update({ userId, email, role });
			updatingRole = false;
			await invalidateAll();
		} catch (error) {
			updatingRole = false;
			toastError(error);
		}
	};
	let popoverOpen = false;
	let showRemovalConfirmation = false;
</script>

<div class="flex flex-col justify-between gap-1 rounded p-2 md:flex-row md:items-center">
	<div class="flex items-center space-x-4">
		<Avatar seed={member.id} />
		<div class="flex flex-col">
			<div>
				{#if member?.firstName || member?.lastName}
					<p class="text-sm font-medium leading-none">
						{member.firstName ?? ''}
						{member.lastName ?? ''}
					</p>
				{/if}
				<p class="text-sm text-muted-foreground">{member.email ?? ''}</p>
			</div>
			{#if $page.data.userId == member.id}
				<Badge class="max-w-fit text-[10px] capitalize">YOU</Badge>
			{/if}
		</div>
	</div>
	{#if requireRoles([ROLE_VALUES.ADMIN, ROLE_VALUES.OWNER]) && member.role !== ROLE_VALUES.OWNER}
		<Popover.Root let:ids bind:open={popoverOpen}>
			<Popover.Trigger asChild let:builder>
				<Button
					builders={[builder]}
					variant="outline"
					disabled={updatingRole}
					class="ml-auto max-w-fit capitalize"
				>
					{member.role}
					{#if updatingRole}
						<ButtonLoadingSpinner bind:state={updatingRole} />
					{:else}
						<ChevronDownIcon class="ml-2 h-4 w-4 text-muted-foreground" />
					{/if}
				</Button>
			</Popover.Trigger>
			<Popover.Content class="p-0" align="end">
				<Command.Root>
					<Command.Input placeholder="Select new role..." />
					<Command.List>
						<Command.Empty>No roles found.</Command.Empty>
						<Command.Group>
							{#each ROLES.filter((r) => r.value !== 'owner') as role}
								<Command.Item
									onSelect={() => {
										updateRole(member.id, member.email, role.value);
										popoverOpen = closeAndRefocusTrigger(ids.trigger);
									}}
									class="flex flex-col items-start space-y-1 px-4 py-2"
								>
									<p>{role.name}</p>
									<p class="text-left text-xs text-muted-foreground">
										{role.description}
									</p>
								</Command.Item>
							{/each}

							<Command.Item
								onSelect={() => {
									showRemovalConfirmation = true;
									popoverOpen = closeAndRefocusTrigger(ids.trigger);
								}}
								class="flex  flex-col items-start space-y-1 px-4 py-2"
							>
								<p class="text-destructive">Remove</p>
								<p class="text-left text-xs text-destructive">
									Removes this member from your team.
								</p>
							</Command.Item>
						</Command.Group>
					</Command.List>
				</Command.Root>
			</Popover.Content>
		</Popover.Root>
	{:else}
		<Badge variant="outline" class="ml-auto max-w-fit px-4 py-2 capitalize">{member.role}</Badge>
	{/if}
</div>
<RemovalConfirmation bind:open={showRemovalConfirmation} id={member.id} />
