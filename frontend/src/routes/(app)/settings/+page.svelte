<script lang="ts">
	import * as Card from '$lib/components/ui/card';
	import * as Command from '$lib/components/ui//command';
	import * as Popover from '$lib/components/ui//popover';
	import { Button } from '$lib/components/ui/button';
	import { ChevronDownIcon } from 'lucide-svelte';
	import Avatar from '$lib/components/reusable/images/Avatar.svelte';
	import InviteMember from './(components)/invite-member.svelte';
	import { requireRoles } from '$utils/guards';
	import { ROLE_VALUES } from '$utils/interfaces/roles.interface';
	import { Badge } from '$lib/components/ui/badge';
	import type { PageData } from '../$types';
	import LoadingSpinner from '$lib/components/reusable/LoadingSpinner.svelte';
	import Separator from '$lib/components/ui/separator/separator.svelte';
	export let data: PageData;
	$: ({ members } = data);
	$: loadingMembers = data?.loadingMembers ?? true;
</script>

<Card.Root>
	<Card.Header class="flex justify-between gap-3 md:flex-row md:items-center">
		<div>
			<Card.Title>Team Members</Card.Title>
			<Card.Description>Invite your team members to collaborate</Card.Description>
		</div>
		{#if requireRoles([ROLE_VALUES.ADMIN, ROLE_VALUES.OWNER])}
			<InviteMember />
		{/if}
	</Card.Header>
	<Separator class="mb-3" />
	{#if loadingMembers}
		<LoadingSpinner />
	{:else if members?.length}
		<Card.Content class="grid gap-2">
			{#each members as member}
				<div
					class="hover:bg-muted flex flex-col justify-between gap-1 space-x-4 rounded p-2 md:flex-row md:items-center"
				>
					<div class="flex items-center space-x-4">
						<Avatar seed={member.id} />
						<div>
							{#if member?.firstName || member?.lastName}
								<p class="text-sm font-medium leading-none">
									{member.firstName ?? ''}
									{member.lastName ?? ''}
								</p>
							{/if}
							<p class="text-muted-foreground text-sm">{member.email ?? ''}</p>
						</div>
					</div>
					{#if requireRoles([ROLE_VALUES.ADMIN, ROLE_VALUES.OWNER])}
						<Popover.Root>
							<Popover.Trigger asChild let:builder>
								<Button builders={[builder]} variant="outline" class="ml-auto capitalize">
									{member.role}
									<ChevronDownIcon class="text-muted-foreground ml-2 h-4 w-4" />
								</Button>
							</Popover.Trigger>
							<Popover.Content class="p-0" align="end">
								<Command.Root>
									<Command.Input placeholder="Select new role..." />
									<Command.List>
										<Command.Empty>No roles found.</Command.Empty>
										<Command.Group>
											<Command.Item class="flex flex-col items-start space-y-1 px-4 py-2">
												<p>Viewer</p>
												<p class="text-muted-foreground text-sm">Can view and comment.</p>
											</Command.Item>
											<Command.Item class="flex flex-col items-start space-y-1 px-4 py-2">
												<p>Developer</p>
												<p class="text-muted-foreground text-sm">Can view, comment, and edit.</p>
											</Command.Item>
											<Command.Item class="flex flex-col items-start space-y-1 px-4 py-2">
												<p>Billing</p>
												<p class="text-muted-foreground text-sm">
													Can view, comment and manage billing.
												</p>
											</Command.Item>
											<Command.Item class="flex flex-col items-start space-y-1 px-4 py-2">
												<p>Owner</p>
												<p class="text-muted-foreground text-sm">
													Admin-level access to all resources.
												</p>
											</Command.Item>
										</Command.Group>
									</Command.List>
								</Command.Root>
							</Popover.Content>
						</Popover.Root>
					{:else}
						<Badge variant="outline" class="px-4 py-2 capitalize">{member.role}</Badge>
					{/if}
				</div>
			{/each}
		</Card.Content>
	{/if}
</Card.Root>
<!-- </div> -->
