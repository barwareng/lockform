<script lang="ts">
	import * as Form from '$lib/components/ui/form';
	import LoadingSpinner from '$lib/components/reusable/loading-spinners/LoadingSpinner.svelte';
	import { MetaTags } from 'svelte-meta-tags';
	import type { PageData } from './$types';
	import { VITE_APP_NAME } from '$lib/env';
	import Separator from '$lib/components/ui/separator/separator.svelte';
	import Label from '$lib/components/ui/label/label.svelte';
	import Input from '$lib/components/ui/input/input.svelte';
	import Button from '$lib/components/ui/button/button.svelte';
	import { toastError, toastSuccess } from '$utils/toasts';
	import { client } from '$lib/api/Client';
	import { invalidateAll } from '$app/navigation';
	import ButtonLoadingSpinner from '$lib/components/reusable/loading-spinners/ButtonLoadingSpinner.svelte';

	export let data: PageData;
	$: ({ profile } = data);
	$: loadingProfile = data?.loadingProfile ?? true;
	let updatingProfile = false;
	const updateProfile = async () => {
		try {
			updatingProfile = true;
			await client.users.update(profile);
			await invalidateAll();
			toastSuccess('Your profile has been updated');
			updatingProfile = false;
		} catch (error) {
			updatingProfile = false;
			toastError(error);
		}
	};
</script>

<MetaTags title="{VITE_APP_NAME} | Settings | Profile" />
<div class="space-y-6">
	<div class="flex flex-col justify-between md:flex-row md:items-center">
		<div>
			<h3 class="text-lg font-medium">Profile</h3>
			<p class="text-muted-foreground text-sm">This is how others will see you on the site.</p>
		</div>
		<!-- <EditProfileSheet bind:profile /> -->
	</div>
	<Separator />
	<!-- <ProfileForm data={data.form} /> -->
	{#if loadingProfile}
		<LoadingSpinner />
	{:else if profile}
		<!-- {JSON.stringify(profile)} -->
		<div class="space-y-6">
			<div class="flex w-full flex-col justify-between gap-x-4 gap-y-3 md:flex-row md:items-start">
				<div class="flex-1 space-y-1">
					<Label>First name</Label>
					<Input bind:value={profile.firstName} placeholder="Luke" />
				</div>
				<div class="flex-1 space-y-1">
					<Label>Last name</Label>
					<Input bind:value={profile.lastName} placeholder="Skywalker" />
				</div>
			</div>
			<div class="flex w-full flex-col justify-between gap-x-4 gap-y-3 md:flex-row md:items-start">
				<div class="flex-1 space-y-1">
					<Label>Email</Label>
					<Input bind:value={profile.email} placeholder="luke@starwars.com" />
				</div>
				<div class="flex-1 space-y-1">
					<Label>Phone number</Label>
					<Input bind:value={profile.phoneNumber} placeholder="(702)-234-5566" />
				</div>
			</div>
			<Button disabled={updatingProfile} on:click={updateProfile}>
				<ButtonLoadingSpinner bind:state={updatingProfile} />
				Update profile
			</Button>
		</div>
	{/if}
</div>
