<script lang="ts">
	import { MetaTags } from 'svelte-meta-tags';
	import { VITE_APP_NAME } from '$lib/env';
	import WipPlaceHolder from '$lib/components/reusable/WIPPlaceHolder.svelte';
</script>

<MetaTags title="{VITE_APP_NAME} | Domains" />
<WipPlaceHolder />
