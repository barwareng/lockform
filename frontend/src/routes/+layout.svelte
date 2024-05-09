<script>
	import CreateTeam from '$lib/components/reusable/dialogs/CreateTeam.svelte';
	import '../app.pcss';
	import { ModeWatcher } from 'mode-watcher';
	import { Toaster } from '$lib/components/ui/sonner';
	import { MetaTags } from 'svelte-meta-tags';
	import { VITE_APP_NAME } from '$lib/env';
</script>

<MetaTags title={VITE_APP_NAME} />
<ModeWatcher />
<CreateTeam />
<Toaster />
<slot />
