<script lang="ts">
	import Separator from '$lib/components/ui/separator/separator.svelte';
	import { MetaTags } from 'svelte-meta-tags';
	import type { PageData } from './$types';
	import { VITE_APP_NAME } from '$lib/env';
	import WipPlaceHolder from '$lib/components/reusable/WIPPlaceHolder.svelte';

	export let data: PageData;
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
		<WipPlaceHolder />
	</div>
</div>
