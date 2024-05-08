<script lang="ts">
	import * as Dialog from '$lib/components/ui/dialog';
	import { Input } from '$lib/components/ui/input';
	import { Button, buttonVariants } from '$lib/components/ui/button';
	import * as Command from '$lib/components/ui//command';
	import * as Popover from '$lib/components/ui//popover';
	import { client } from '$lib/api/Client';
	import { goto, invalidateAll } from '$app/navigation';
	import { toastError } from '$utils/toasts';
	import { ChevronDownIcon, PlusCircleIcon } from 'lucide-svelte';
	import { ROLES } from '$utils/constants/roles.constants';

	let email: string;
	let selectedRole = ROLES?.filter((r) => r.value != 'owner')?.[0];
	const createTeam = async () => {
		try {
			await client.members.create({ email, role: selectedRole.value });
			email = '';
			selectedRole = ROLES?.filter((r) => r.value != 'owner')?.[0];
			await invalidateAll();
			open = false;
		} catch (error) {
			toastError(error);
		}
	};

	let open = false;
</script>

<Dialog.Root bind:open>
	<Dialog.Trigger class={buttonVariants({ variant: 'default' })}>
		<PlusCircleIcon class="mr-1 h-4 w-4" />
		Invite member
	</Dialog.Trigger>
	<Dialog.Content>
		<Dialog.Header>
			<Dialog.Title>Add member</Dialog.Title>
			<Dialog.Description>Invite a new member to collaborate.</Dialog.Description>
		</Dialog.Header>
		<div>
			<div class="space-y-4 py-2 pb-4">
				<div class="flex gap-2">
					<Input id="name" placeholder="member@acme.com" bind:value={email} />
					<Popover.Root>
						<Popover.Trigger asChild let:builder>
							<Button builders={[builder]} variant="outline" class="ml-auto">
								{selectedRole.name}
								<ChevronDownIcon class="text-muted-foreground ml-2 h-4 w-4" />
							</Button>
						</Popover.Trigger>
						<Popover.Content class="p-0" align="end">
							<Command.Root>
								<Command.Input placeholder="Select new role..." />
								<Command.List>
									<Command.Empty>No roles found.</Command.Empty>
									<Command.Group>
										{#each ROLES.filter((r) => r.value !== 'owner') as role}
											<Command.Item>
												<button
													type="button"
													class="flex flex-col items-start space-y-1 px-4 py-2"
													on:click={() => (selectedRole = role)}
												>
													<p>{role.name}</p>
													<p class="text-muted-foreground text-left text-sm">
														{role.description}
													</p>
												</button>
											</Command.Item>
										{/each}
									</Command.Group>
								</Command.List>
							</Command.Root>
						</Popover.Content>
					</Popover.Root>
				</div>
			</div>
			<Dialog.Footer>
				<Button variant="outline" on:click={() => (open = false)}>Cancel</Button>
				<Button type="button" on:click={createTeam}>Invite</Button>
			</Dialog.Footer>
		</div></Dialog.Content
	>
</Dialog.Root>
