<script lang="ts">
	import * as Popover from '$lib/components/ui/popover';
	import { Button } from '$lib/components/ui/button';

	import * as Card from '$lib/components/ui/card/index.js';
	import Separator from '$lib/components/ui/separator/separator.svelte';
	export let icon: any;
	export let title: string;
	export let open = false;
</script>

<Popover.Root bind:open>
	<Popover.Trigger asChild let:builder>
		<Button
			builders={[builder]}
			variant="outline"
			class="flex aspect-square min-h-20 flex-col gap-y-2 text-xs"
		>
			<svelte:component this={icon} class="!h-8 !w-8" />
			{title}
		</Button>
	</Popover.Trigger>
	<Popover.Content class="relative w-80">
		<div class="space-y-2">
			<h4 class="font-medium leading-none">{title}</h4>
		</div>
		<Separator class="mt-4" />
		<div class="py-4">
			<slot />
		</div>
	</Popover.Content>
</Popover.Root>
