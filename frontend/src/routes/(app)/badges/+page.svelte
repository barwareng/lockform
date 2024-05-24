<script lang="ts">
	import Separator from '$lib/components/ui/separator/separator.svelte';
	import { MetaTags } from 'svelte-meta-tags';
	import type { PageData } from './$types';
	import { VITE_APP_NAME } from '$lib/env';
	import * as Tabs from '$lib/components/ui/tabs';
	import * as Card from '$lib/components/ui/card';
	import { Skeleton } from '$lib/components/ui/skeleton';
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
			<Tabs.List class="max-w-full overflow-x-scroll">
				{#each badgeVariants as variant}
					<Tabs.Trigger value={variant.name}>{variant.name}</Tabs.Trigger>
				{/each}
			</Tabs.List>
			{#each badgeVariants as variant}
				<Tabs.Content value={variant.name} class="w-full ">
					<Card.Root class="mx-auto mt-6 w-full lg:w-1/2">
						<Card.Header>
							<div class="flex w-full items-center space-x-4">
								<Skeleton class="h-12 w-12 rounded-full" />
								<div class="flex-1 space-y-2">
									<Skeleton class="h-4 w-4/5"></Skeleton>
									<Skeleton class="h-4 w-3/5" />
								</div>
							</div>
						</Card.Header>
						<Card.Content class="space-y-2">
							<Skeleton class="h-16 w-full" />
							<Skeleton class="h-12 w-full" />
							<Skeleton class="h-12 w-1/3" />
							<Skeleton class="mx-auto h-10 w-1/2" />
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
