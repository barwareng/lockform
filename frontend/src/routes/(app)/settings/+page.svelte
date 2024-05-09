<script lang="ts">
	import * as Card from '$lib/components/ui/card';
	import InviteMember from './(components)/invite-member.svelte';
	import { requireRoles } from '$utils/guards';
	import { ROLE_VALUES } from '$utils/interfaces/roles.interface';
	import type { PageData } from '../$types';
	import LoadingSpinner from '$lib/components/reusable/loading-spinners/LoadingSpinner.svelte';
	import Separator from '$lib/components/ui/separator/separator.svelte';

	import MemberRow from './(components)/member-row.svelte';
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
		<Card.Content class="grid gap-2 p-0 sm:p-4">
			{#each members as member, i}
				<MemberRow bind:member />
				{#if i < members.length - 1}
					<Separator />
				{/if}
			{/each}
		</Card.Content>
	{/if}
</Card.Root>
<!-- </div> -->
