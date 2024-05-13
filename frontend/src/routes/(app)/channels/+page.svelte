<script lang="ts">
	import { MetaTags } from 'svelte-meta-tags';
	import type { PageData } from './$types';
	import { VITE_APP_NAME } from '$lib/env';
	import Separator from '$lib/components/ui/separator/separator.svelte';
	import WipPlaceHolder from '$lib/components/reusable/WIPPlaceHolder.svelte';
	import { requireRoles } from '$utils/guards';
	import ListChannelOptionsDialog from './(components)/list-channel-options-dialog.svelte';
	import { ROLE_VALUES } from '$utils/interfaces/roles.interface';
	import LoadingSpinner from '$lib/components/reusable/loading-spinners/LoadingSpinner.svelte';
	import EmptyState from '$lib/components/reusable/EmptyState.svelte';
	import { AntennaIcon } from 'lucide-svelte';

	export let data: PageData;
	$: ({ channels } = data);
	$: loadingChannels = data.loadingChannels ?? true;
</script>

<MetaTags
	title="{VITE_APP_NAME} | Channels"
	description="Manage all the channels that your team uses to communicate"
/>
<div class="relative pb-16">
	<div class="bg-background sticky top-0 z-10">
		<div class="flex flex-col items-start gap-6 md:flex-row md:justify-between">
			<div class="space-y-0.5">
				<h2 class="text-2xl font-bold tracking-tight">Channels</h2>
				<p class="text-muted-foreground text-sm">
					Manage all the channels that your team uses to communicate
				</p>
			</div>
			{#if requireRoles([ROLE_VALUES.ADMIN, ROLE_VALUES.OWNER])}
				<ListChannelOptionsDialog />
			{/if}
		</div>
		<Separator class="my-6" />
	</div>
	{#if loadingChannels}
		<LoadingSpinner />
	{:else if channels?.length}
		Channels
	{:else}
		<EmptyState
			title="No channels yet"
			subtitle="When you add your teams channels, they will show up here."
		>
			<div slot="leading-icon"><AntennaIcon class="h-12 w-12" /></div>
			{#if requireRoles([ROLE_VALUES.ADMIN, ROLE_VALUES.OWNER])}
				<ListChannelOptionsDialog />
			{/if}
		</EmptyState>
	{/if}
</div>
