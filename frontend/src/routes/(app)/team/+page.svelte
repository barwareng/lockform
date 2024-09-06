<script lang="ts">
	import EmptyState from '$lib/components/reusable/EmptyState.svelte';
	import Button from '$lib/components/ui/button/button.svelte';
	import { showCreateTeamDialog } from '$stores';
	import { Boxes, PlusCircle } from 'lucide-svelte';
</script>

<div class="flex h-full w-full items-center justify-center">
	<EmptyState
		title="Start by creating a team"
		subtitle="A team helps you manage your organization's resources. Create a team to get started."
	>
		<div slot="leading-icon"><Boxes class="h-12 w-12" /></div>
		<Button class="mb-4 w-full" on:click={() => ($showCreateTeamDialog = true)}>
			<PlusCircle class="mr-2 h-5 w-5" />
			Create Team</Button
		>
	</EmptyState>
</div>
