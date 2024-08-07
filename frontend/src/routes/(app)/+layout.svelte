<script lang="ts">
	import Sidenav from '$lib/components/reusable/navigation/Sidenav.svelte';
	import { Button } from '$lib/components/ui/button';
	import { Separator } from '$lib/components/ui/separator';
	import { MenuIcon } from 'lucide-svelte';
	import { fade, slide } from 'svelte/transition';
	let openMobileNav = false;
</script>

<main class="relative mx-auto flex h-screen w-screen max-w-screen-2xl">
	{#if openMobileNav}
		<aside
			transition:slide
			class="absolute z-40 min-h-screen w-5/6 min-w-[300px] max-w-40 shrink-0 basis-1/5 bg-background lg:hidden"
		>
			<Sidenav bind:openMobileNav />
		</aside>
	{/if}
	<aside class="sticky top-0 hidden max-h-screen min-w-[300px] shrink-0 basis-1/5 lg:block">
		<Sidenav bind:openMobileNav />
	</aside>
	<Separator orientation="vertical" />
	<section class="min-h-screen flex-1 overflow-y-scroll px-6 pb-10 lg:px-10">
		<div class="h-10 bg-background">
			<Button
				size="icon"
				variant="ghost"
				class="lg:hidden"
				on:click={() => {
					openMobileNav = true;
				}}
			>
				<MenuIcon />
			</Button>
		</div>
		<slot />
	</section>
</main>
