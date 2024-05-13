<script>
	import Navbar from '$lib/components/reusable/navigation/Navbar.svelte';
	import '../app.css';
	import { ModeWatcher } from 'mode-watcher';
	import { Separator } from '$lib/components/ui/separator';
	import { MetaTags } from 'svelte-meta-tags';
	import { VITE_APP_NAME } from '$lib/env';

	import { Toaster } from '$lib/components/ui/sonner';
	import Footer from '$lib/components/reusable/navigation/Footer.svelte';
</script>

<MetaTags title={VITE_APP_NAME} />
<ModeWatcher />
<Toaster />
<main class="py-4">
	<div class=" sticky top-2 z-10 backdrop-blur-sm">
		<div class="container mx-auto max-w-screen-xl">
			<Navbar />
		</div>
		<Separator class="my-2" />
	</div>
	<div class="container mx-auto max-w-screen-xl">
		<slot />
	</div>
	<div class="">
		<Separator class="my-2" />
		<div class="container mx-auto max-w-screen-xl">
			<Footer />
		</div>
	</div>
</main>
