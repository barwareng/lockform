<script lang="ts">
	import {
		FacebookIcon,
		InstagramIcon,
		LinkedinIcon,
		LucideFacebook,
		MailCheckIcon,
		MailIcon,
		TrafficConeIcon,
		TwitterIcon
	} from 'lucide-svelte';
	import Button from '$lib/components/ui/button/button.svelte';
	import { fade, slide } from 'svelte/transition';

	export let title = '';
	export let subtitle = '';
</script>

<div
	transition:slide
	class="flex h-96 w-full flex-col items-center justify-center text-center {$$props.class}"
>
	<TrafficConeIcon class="h-16 w-16" />
	<h3 class="mt-2 text-sm font-semibold">WIP Feature</h3>
	<p class="mt-1 max-w-md text-sm text-muted-foreground">
		We are hard at work to get this feature in your hands. <br /> We'll let you know once it's ready.
	</p>
	<div class="mt-6 flex items-center">
		<Button size="icon" variant="ghost">
			<MailIcon class="h-5 w-5" />
		</Button>
		<Button size="icon" variant="ghost">
			<FacebookIcon class="h-5 w-5" />
		</Button>
		<Button size="icon" variant="ghost">
			<InstagramIcon class="h-5 w-5" />
		</Button>
		<Button size="icon" variant="ghost">
			<TwitterIcon class="h-5 w-5" />
		</Button>
		<Button size="icon" variant="ghost">
			<LinkedinIcon class="h-5 w-5" />
		</Button>
	</div>
</div>
