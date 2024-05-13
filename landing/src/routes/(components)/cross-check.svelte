<script lang="ts">
	import { goto } from '$app/navigation';
	import { Button } from '$lib/components/ui/button';
	import { Input } from '$lib/components/ui/input';
	let searchPhrase: string;
</script>

<form class="mx-auto mt-8 flex w-5/6 flex-col items-center gap-2 md:flex-row">
	<Input
		type="text"
		placeholder="Email address, phone number, social media handle"
		bind:value={searchPhrase}
	/>
	<Button type="button" disabled={!searchPhrase} on:click={() => goto(`/result/${searchPhrase}`)}
		>Verify</Button
	>
</form>
