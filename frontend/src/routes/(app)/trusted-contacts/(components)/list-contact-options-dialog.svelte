<script lang="ts">
	import { buttonVariants } from '$lib/components/ui/button';
	import * as Drawer from '$lib/components/ui/drawer';
	import { PlusCircleIcon } from 'lucide-svelte';
	import { isMobile } from '$utils';
	import { cn } from '$lib/utils';
	import { contactDialogs } from './contact-dialogs/dialogs';
</script>

<Drawer.Root>
	<Drawer.Trigger class={buttonVariants({ variant: 'default' })}>
		<PlusCircleIcon class="mr-1 h-4 w-4" />
		Add contact
	</Drawer.Trigger>
	<Drawer.Content class={cn('max-h-[80%] pb-10', $isMobile ? '' : 'mx-auto h-auto max-w-xl')}>
		<Drawer.Header>
			<Drawer.Title>Select contact</Drawer.Title>
			<Drawer.Description>Select a contact to add.</Drawer.Description>
		</Drawer.Header>
		<div class="space-y-5 overflow-y-scroll px-4">
			<div class="flex w-full flex-row flex-wrap justify-start gap-4">
				{#each contactDialogs as dialog}
					<svelte:component this={dialog.component} />
				{/each}
			</div>
		</div>
	</Drawer.Content>
</Drawer.Root>
