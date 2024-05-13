<script lang="ts">
	import { buttonVariants } from '$lib/components/ui/button';
	import * as Drawer from '$lib/components/ui/drawer';
	import { PlusCircleIcon } from 'lucide-svelte';
	import { isMobile } from '$utils';
	import { cn } from '$lib/utils';
	import { channelDialogs } from './channel-dialogs/dialogs';
</script>

<Drawer.Root>
	<Drawer.Trigger class={buttonVariants({ variant: 'default' })}>
		<PlusCircleIcon class="mr-1 h-4 w-4" />
		Add channel
	</Drawer.Trigger>
	<Drawer.Content class={cn('max-h-[80%] pb-10', $isMobile ? '' : 'mx-auto h-auto max-w-xl')}>
		<Drawer.Header>
			<Drawer.Title>Select channel</Drawer.Title>
			<Drawer.Description>Select a channel to add.</Drawer.Description>
		</Drawer.Header>
		<div class="space-y-5 overflow-y-scroll px-4">
			{#each channelDialogs as dialogGroup}
				<div class="space-y-2">
					<p class="font-semibold">{dialogGroup.type}</p>
					<div class="flex flex-wrap justify-start gap-4">
						{#each dialogGroup.dialogs as dialog}
							<svelte:component this={dialog} />
						{/each}
					</div>
				</div>
			{/each}
		</div>
	</Drawer.Content>
</Drawer.Root>
