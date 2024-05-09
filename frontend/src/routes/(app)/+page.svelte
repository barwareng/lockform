<script lang="ts">
	import { MetaTags } from 'svelte-meta-tags';
	import type { PageData } from './$types';
	import { VITE_APP_NAME } from '$lib/env';
	import WipPlaceHolder from '$lib/components/reusable/WIPPlaceHolder.svelte';

	export let data: PageData;
</script>

<MetaTags title="{VITE_APP_NAME} | Dashboard" />

<WipPlaceHolder />
