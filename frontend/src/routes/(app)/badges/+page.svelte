<script lang="ts">
	import Separator from '$lib/components/ui/separator/separator.svelte';
	import { MetaTags } from 'svelte-meta-tags';
	import type { PageData } from './$types';
	import { VITE_APP_NAME } from '$lib/env';
	import * as Tabs from '$lib/components/ui/tabs';
	import * as Card from '$lib/components/ui/card';
	import { Button } from '$lib/components/ui/button';
	import { Input } from '$lib/components/ui/input';
	import { Label } from '$lib/components/ui/label';
	import BadgeBottomLeft from './(components)/badge-bottom-left.svelte';
	import BadgeBottomCenter from './(components)/badge-bottom-center.svelte';
	import BadgeBottomRight from './(components)/badge-bottom-right.svelte';
	export let data: PageData;
	const badgeVariants = [
		{
			name: 'Bottom Left',
			component: BadgeBottomLeft,
			html: ``
		},
		{
			name: 'Bottom Center',
			component: BadgeBottomCenter,
			html: ``
		},
		{
			name: 'Bottom Right',
			component: BadgeBottomRight,
			html: ``
		}
	];
</script>

<MetaTags
	title="{VITE_APP_NAME} | Badges"
	description="Manage how your verification badges show emails to your customers."
/>
<div class="relative pb-16">
	<div class="bg-background sticky top-0 z-10">
		<div class="flex flex-col items-start gap-6 md:flex-row md:justify-between">
			<div class="space-y-0.5">
				<h2 class="text-2xl font-bold tracking-tight">Badges</h2>
				<p class="text-muted-foreground text-sm">
					Manage how your verification badges show emails to your customers.
				</p>
			</div>
		</div>
		<Separator class="my-6" />

		<Tabs.Root value="Bottom Left">
			<Tabs.List class="">
				{#each badgeVariants as variant}
					<Tabs.Trigger value={variant.name}>{variant.name}</Tabs.Trigger>
				{/each}
			</Tabs.List>
			{#each badgeVariants as variant}
				<Tabs.Content value={variant.name} class="w-full ">
					<Card.Root class="mx-auto mt-6 w-1/2">
						<Card.Header>
							<!-- <Card.Title>Account</Card.Title>
							<Card.Description>
								Make changes to your account here. Click save when you're done.
							</Card.Description> -->
						</Card.Header>
						<Card.Content class="space-y-2">
							<!-- <div class="space-y-1">
								<Label for="name">Name</Label>
								<Input id="name" value="Pedro Duarte" />
							</div>
							<div class="space-y-1">
								<Label for="username">Username</Label>
								<Input id="username" value="@peduarte" />
							</div> -->
						</Card.Content>
						<Card.Footer>
							<svelte:component this={variant.component}></svelte:component>
						</Card.Footer>
					</Card.Root>
				</Tabs.Content>
			{/each}
		</Tabs.Root>
	</div>
</div>
