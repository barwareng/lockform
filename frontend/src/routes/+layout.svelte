<script>
	import CreateTeam from '$lib/components/reusable/dialogs/CreateTeam.svelte';
	import '../app.pcss';
	import { ModeWatcher } from 'mode-watcher';
	import { Toaster } from '$lib/components/ui/sonner';
	import { MetaTags } from 'svelte-meta-tags';
</script>

<MetaTags title="Veriform" />
<ModeWatcher />
<CreateTeam />
<Toaster />
<slot />
