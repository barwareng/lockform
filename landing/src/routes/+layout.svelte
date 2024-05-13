<script>
	import Navbar from '$lib/components/reusable/navigation/Navbar.svelte';
	import '../app.css';
	import { ModeWatcher } from 'mode-watcher';
	import { Separator } from '$lib/components/ui/separator';
	import { MetaTags } from 'svelte-meta-tags';
	import { VITE_APP_NAME } from '$lib/env';

	import { Toaster } from '$lib/components/ui/sonner';
</script>

<MetaTags title={VITE_APP_NAME} />
<ModeWatcher />
<Toaster />
<main class="py-4">
	<div class="container mx-auto max-w-screen-xl">
		<Navbar />
	</div>
	<Separator class="my-2" />
	<div class="container mx-auto max-w-screen-xl">
		<slot />
	</div>
</main>
