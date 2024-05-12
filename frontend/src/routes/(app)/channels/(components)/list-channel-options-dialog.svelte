<script lang="ts">
	import { ScrollArea } from '$lib/components/ui/scroll-area/index.js';
	import { Button, buttonVariants } from '$lib/components/ui/button/index.js';
	import * as Drawer from '$lib/components/ui/drawer';
	import { PlusCircleIcon } from 'lucide-svelte';
	import EmailDialog from './channel-dialogs/email-dialog.svelte';
	import PhoneDialog from './channel-dialogs/phone-dialog.svelte';
	import { isMobile } from '$utils';
	import { cn } from '$lib/utils';
</script>

<Drawer.Root open>
	<Drawer.Trigger class={buttonVariants({ variant: 'default' })}>
		<PlusCircleIcon class="mr-1 h-4 w-4" />
		Add channel
	</Drawer.Trigger>
	<Drawer.Content class={cn('max-h-[80%] pb-10', $isMobile ? '' : 'mx-auto h-auto max-w-xl')}>
		<Drawer.Header>
			<Drawer.Title>Select channel</Drawer.Title>
			<Drawer.Description>Select a channel to add.</Drawer.Description>
		</Drawer.Header>
		<div class="flex flex-wrap justify-center gap-4 overflow-y-scroll">
			<!-- {#each Array(20) as _} -->
			<EmailDialog />
			<PhoneDialog />
			<!-- {/each} -->
		</div>
	</Drawer.Content>
</Drawer.Root>
