<script lang="ts">
	import * as Card from '$lib/components/ui/card';
	import * as Command from '$lib/components/ui//command';
	import * as Popover from '$lib/components/ui//popover';
	import { Button } from '$lib/components/ui/button';
	import { ChevronDownIcon } from 'lucide-svelte';
	import Avatar from '$lib/components/reusable/images/Avatar.svelte';
</script>

<!-- <div class="lg:w/2/3 mx-auto w-full max-w-xl md:w-3/4"> -->
<Card.Root>
	<Card.Header>
		<Card.Title>Team Members</Card.Title>
		<Card.Description>Invite your team members to collaborate</Card.Description>
	</Card.Header>
	<Card.Content class="grid gap-6">
		{#each Array(5) as _}
			<div class="hover:bg-muted flex items-center justify-between space-x-4 rounded p-1">
				<div class="flex items-center space-x-4">
					<Avatar />
					<div>
						<p class="text-sm font-medium leading-none">Sofia Davis</p>
						<p class="text-muted-foreground text-sm">m@example.com</p>
					</div>
				</div>
				<Popover.Root>
					<Popover.Trigger asChild let:builder>
						<Button builders={[builder]} variant="outline" class="ml-auto">
							Owner
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
			</div>
		{/each}
	</Card.Content>
</Card.Root>
<!-- </div> -->
