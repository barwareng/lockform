<script>
	import { Button } from '$lib/components/ui/button';
	import { Input } from '$lib/components/ui/input';
	import CrossCheck from './(components)/cross-check.svelte';
</script>

<section class="mx-auto max-w-2xl py-56">
	<h1 class="text-center text-6xl font-bold">Confirm the business is what it says it is.</h1>
	<h2 class="text-muted-foreground mt-4 text-center text-lg">
		Verify that email you received is not a phishing attempt. Or that text, or that call.
	</h2>
	<CrossCheck />
</section>
