<script lang="ts">
	import Separator from '$lib/components/ui/separator/separator.svelte';
	import { MetaTags } from 'svelte-meta-tags';
	import type { PageData } from './$types';
	import { VITE_APP_NAME } from '$lib/env';
	import * as Tabs from '$lib/components/ui/tabs';
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import { slide } from 'svelte/transition';

	export let data: PageData;
	const settingsRoutes = [
		{
			name: 'Team members',
			value: 'team-members',
			path: `/settings`
		},
		{
			name: 'Billing',
			value: 'billing',
			path: `/settings/billing`
		},
		{
			name: 'Profile',
			value: 'profile',
			path: `/settings/profile`
		}
	];
	$: value = settingsRoutes.find((route) => route.path == $page.url.pathname)?.value;
</script>

<MetaTags
	title="{VITE_APP_NAME} | Settings"
	description="Manage your team and personal profile settings."
/>
<div class="relative pb-16">
	<div class="bg-background sticky top-0 z-10">
		<div class="flex flex-col items-start gap-6 md:flex-row md:justify-between">
			<div class="space-y-0.5">
				<h2 class="text-2xl font-bold tracking-tight">Settings</h2>
				<p class="text-muted-foreground text-sm">Manage your team and personal profile settings.</p>
			</div>
		</div>
		<Separator class="my-6" />
		<div>
			<Tabs.Root bind:value class="w-[400px]">
				<Tabs.List>
					{#each settingsRoutes as route}
						<Tabs.Trigger
							value={route.value}
							on:click={() => goto(route.path, { invalidateAll: true })}
						>
							{route.name}
						</Tabs.Trigger>
					{/each}
				</Tabs.List>
			</Tabs.Root>
		</div>
	</div>
	<div transition:slide class="mt-10">
		<slot />
	</div>
</div>
