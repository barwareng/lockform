<script>
	import PageHeader from '$lib/components/reusable/PageHeader.svelte';
	import CreateTeam from '$lib/components/reusable/dialogs/CreateTeam.svelte';
	import '../app.pcss';
	import { ModeWatcher } from 'mode-watcher';
	import { Toaster } from '$lib/components/ui/sonner';
</script>

<PageHeader />
<ModeWatcher />
<CreateTeam />
<Toaster />
<slot />
