<script lang="ts">
	import { page } from '$app/stores';
	import { Button } from '$lib/components/ui/button';
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu';
	import { setMode } from 'mode-watcher';
	import {
		AntennaIcon,
		BadgeIcon,
		Building,
		Cog,
		ConstructionIcon,
		ContactIcon,
		HomeIcon,
		LogOutIcon,
		PieChartIcon,
		SunMoon,
		SunMoonIcon,
		User
	} from 'lucide-svelte';
	import { logout } from '$utils/supertokens';
	import TeamSwitcher from './TeamSwitcher.svelte';

	$: routes = [
		{
			icon: PieChartIcon,
			name: 'Dashboard',
			path: `/`,
			position: 'top'
		},
		{
			icon: AntennaIcon,
			name: 'Channels',
			path: `/channels`,
			position: 'top'
		},
		{
			icon: ContactIcon,
			name: 'Contacts',
			path: `/contacts`,
			position: 'top'
		},
		{
			icon: BadgeIcon,
			name: 'Badges',
			path: `/badges`,
			position: 'top'
		},

		{
			icon: Cog,
			name: 'Settings',
			path: `/settings`,
			position: 'bottom'
		}
	];
</script>

<nav class="flex h-full flex-col justify-between px-3 py-6">
	<div class="space-y-0.5">
		<TeamSwitcher />
		{#each routes.filter((route) => route.position == 'top') as route}
			<Button
				variant={route.path == $page.url.pathname ? 'secondary' : 'ghost'}
				class="w-full items-center justify-start rounded-full text-xs"
				href={route.path}
			>
				<svelte:component this={route.icon} class="mr-1 h-4 w-4" />
				<span>
					{route.name}
				</span>
			</Button>
		{/each}
	</div>
	<div class="space-y-0.5">
		{#each routes.filter((route) => route.position == 'bottom') as route}
			<Button
				variant={route.path == $page.url.pathname ? 'secondary' : 'ghost'}
				class="w-full items-center justify-start rounded-full text-xs"
				href={route.path}
			>
				<svelte:component this={route.icon} class="mr-1 h-4 w-4" />
				<span>
					{route.name}
				</span>
			</Button>
		{/each}
		<Button
			variant="ghost"
			class="w-full items-center justify-start rounded-full text-xs"
			on:click={logout}
		>
			<LogOutIcon class="mr-1 h-4 w-4" />
			Logout
		</Button>
		<div class="flex items-center justify-between pt-2">
			<Button
				variant="link"
				href="https://lockform.com"
				target="_blank"
				class="flex flex-col items-start justify-start gap-y-0"
			>
				<h1 class="text-lg font-black">lockform</h1>
				<p class="-mt-1 text-[10px]">&copy; {new Date().getFullYear()} All rights reserved</p>
			</Button>
			<DropdownMenu.Root>
				<DropdownMenu.Trigger>
					<Button size="icon" variant="ghost"><SunMoonIcon class="h-5 w-5" /></Button>
				</DropdownMenu.Trigger>
				<DropdownMenu.Content>
					<DropdownMenu.Group>
						<DropdownMenu.Item on:click={() => setMode('dark')}>Dark</DropdownMenu.Item>
						<DropdownMenu.Item on:click={() => setMode('light')}>Light</DropdownMenu.Item>
						<DropdownMenu.Item on:click={() => setMode('system')}>System</DropdownMenu.Item>
					</DropdownMenu.Group>
				</DropdownMenu.Content>
			</DropdownMenu.Root>
		</div>
	</div>
</nav>
