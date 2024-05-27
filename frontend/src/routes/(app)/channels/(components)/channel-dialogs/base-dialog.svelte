<script lang="ts">
	import * as Popover from '$lib/components/ui/popover';
	import { Button } from '$lib/components/ui/button';
	import { Switch } from '$lib/components/ui/switch';
	import Separator from '$lib/components/ui/separator/separator.svelte';
	import { Label } from '$lib/components/ui/label';
	import ButtonLoadingSpinner from '$lib/components/reusable/loading-spinners/ButtonLoadingSpinner.svelte';
	import { CHANNEL, type IChannel } from '$utils/interfaces/channels.interface';
	import { client } from '$lib/api/Client';
	import { invalidateAll } from '$app/navigation';
	import { toastError, toastSuccess } from '$utils/toasts';
	export let icon: any;
	export let title: string;
	export let open = false;
	export let isEditing = false;
	let isSaving = false;
	export let channel: Partial<IChannel> = { isPublic: true };
	const saveChannel = async () => {
		try {
			isSaving = true;
			if (channel.type == CHANNEL.PHONE) channel.value = channel.value?.replaceAll(' ', '');
			await client.channels.create(channel);
			// console.log(channel);
			await invalidateAll();
			isSaving = false;
			open = false;
			toastSuccess('Channel successfully added.');
		} catch (error) {
			isSaving = false;
			toastError(error);
		}
	};
</script>

<Popover.Root bind:open>
	<Popover.Trigger asChild let:builder>
		{#if isEditing}
			<Button
				builders={[builder]}
				variant="ghost"
				size="sm"
				class="flex w-full justify-start px-2 text-sm font-normal"
			>
				Edit
			</Button>
		{:else}
			<Button
				builders={[builder]}
				variant="outline"
				class="flex aspect-square min-h-20 flex-col gap-y-2 text-xs"
			>
				<svelte:component this={icon} class="!h-8 !w-8" />
				{title}
			</Button>
		{/if}
	</Popover.Trigger>
	<Popover.Content class="relative w-80">
		<div class="flex items-center justify-between">
			<h4 class="font-medium leading-none">{title}</h4>
			<div class="flex items-center space-x-2">
				<Switch includeInput id="is-public" bind:checked={channel.isPublic} />
				<Label for="is-public" class="text-xs">Show publicly</Label>
			</div>
		</div>
		<Separator class="mt-4" />
		<div class="space-y-4 py-4">
			<slot />
			<div class="flex items-center justify-end gap-x-4">
				<Button variant="secondary" on:click={() => (open = false)}>Cancel</Button>
				<Button disabled={isSaving} on:click={saveChannel}>
					<ButtonLoadingSpinner bind:state={isSaving} />
					Save
				</Button>
			</div>
		</div>
	</Popover.Content>
</Popover.Root>
