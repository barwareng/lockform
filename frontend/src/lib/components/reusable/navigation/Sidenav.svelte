<script lang="ts">
	import { page } from '$app/stores';
	import { Button } from '$lib/components/ui/button';
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu';
	import { setMode } from 'mode-watcher';
	import {
		AntennaIcon,
		BadgeIcon,
		Building,
		CircleX,
		Cog,
		ConstructionIcon,
		ContactIcon,
		ExternalLinkIcon,
		HomeIcon,
		LogOutIcon,
		PieChartIcon,
		SunMoon,
		SunMoonIcon,
		User
	} from 'lucide-svelte';
	import { logout } from '$utils/supertokens';
	import TeamSwitcher from './TeamSwitcher.svelte';
	export let openMobileNav = false;
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
			name: 'Trusted Contacts',
			path: `/trusted-contacts`,
			position: 'top'
		},
		{
			icon: ExternalLinkIcon,
			name: 'Trusted Domains',
			path: `/trusted-domains`,
			position: 'top'
		},
		// {
		// 	icon: BadgeIcon,
		// 	name: 'Badges',
		// 	path: `/badges`,
		// 	position: 'top'
		// },

		{
			icon: Cog,
			name: 'Settings',
			path: `/settings`,
			position: 'bottom'
		}
	];
</script>

<nav class="flex min-h-screen flex-col justify-between px-3 py-6">
	<Button
		size="icon"
		variant="ghost"
		class="absolute -right-6 top-0 !h-12 !w-12 rounded-full lg:hidden"
		on:click={() => (openMobileNav = false)}
	>
		<CircleX />
	</Button>
	<div class="space-y-0.5">
		<TeamSwitcher />

		{#each routes.filter((route) => route.position == 'top') as route}
			<Button
				variant={route.path == $page.url.pathname ? 'secondary' : 'ghost'}
				class="w-full items-center justify-start rounded-full text-xs"
				href={route.path}
				on:click={() => (openMobileNav = false)}
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
				on:click={() => (openMobileNav = false)}
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
				on:click={() => (openMobileNav = false)}
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
						<DropdownMenu.Item
							on:click={() => {
								setMode('dark');
								openMobileNav = false;
							}}>Dark</DropdownMenu.Item
						>
						<DropdownMenu.Item
							on:click={() => {
								setMode('light');
								openMobileNav = false;
							}}>Light</DropdownMenu.Item
						>
						<DropdownMenu.Item
							on:click={() => {
								setMode('system');
								openMobileNav = false;
							}}>System</DropdownMenu.Item
						>
					</DropdownMenu.Group>
				</DropdownMenu.Content>
			</DropdownMenu.Root>
		</div>
	</div>
</nav>
