<script lang="ts">
	import { client } from '$lib/api/Client';
	import { Button } from '$lib/components/ui/button';
	import { Input } from '$lib/components/ui/input';
	import { toastError } from '../../utils/toasts';
	let searchPhrase: string;
	let verifying = false;
	const verify = async () => {
		try {
			verifying = true;
			await client.verifications.verify({ searchPhrase });
			verifying = false;
		} catch (error) {
			verifying = false;
			toastError(error);
		}
	};
</script>

<form class="mx-auto mt-8 flex w-2/3 items-center space-x-2">
	<Input
		type="text"
		placeholder="Email address, phone number, social media handle"
		bind:value={searchPhrase}
	/>
	<Button type="button" on:click={verify}>Verify</Button>
</form>
