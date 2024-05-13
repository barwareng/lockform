<script lang="ts">
	import Logo from '$lib/icons/Logo.svelte';
	import { Button } from '$lib/components/ui/button';
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu';
	import { ShoppingBagIcon, SunMoonIcon, UserRoundIcon } from 'lucide-svelte';
	import { setMode } from 'mode-watcher';
</script>

<nav class="flex items-center justify-between">
	<div>
		<a href="/" class="flex items-center gap-x-1 text-xl font-black">
			<Logo class="h-6 w-6" />
			lockform
		</a>
	</div>
	<div>
		<Button size="sm">Verify Business</Button>
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
</nav>
