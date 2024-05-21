<script lang="ts">
	import * as Popover from '$lib/components/ui/popover';
	import { Button } from '$lib/components/ui/button';
	import { Switch } from '$lib/components/ui/switch';
	import Separator from '$lib/components/ui/separator/separator.svelte';
	import { Label } from '$lib/components/ui/label';
	export let icon: any;
	export let title: string;
	export let open = false;
	export let isPublic = true;
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
		<div class="flex items-center justify-between">
			<h4 class="font-medium leading-none">{title}</h4>
			<div class="flex items-center space-x-2">
				<Switch includeInput id="is-public" bind:checked={isPublic} />
				<Label for="is-public" class="text-xs">Show publicly</Label>
			</div>
		</div>
		<Separator class="mt-4" />
		<div class="py-4">
			<slot />
		</div>
	</Popover.Content>
</Popover.Root>
