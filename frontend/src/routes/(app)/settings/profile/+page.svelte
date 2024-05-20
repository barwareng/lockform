<script lang="ts">
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
	import * as Popover from '$lib/components/ui/popover';
	import ButtonLoadingSpinner from '$lib/components/reusable/loading-spinners/ButtonLoadingSpinner.svelte';
	import Session from 'supertokens-web-js/recipe/session';
	import * as Command from '$lib/components/ui/command';
	import { TelInput, normalizedCountries } from 'svelte-tel-input';
	import type {
		DetailedValue,
		E164Number,
		CountryCode,
		TelInputOptions
	} from 'svelte-tel-input/types';
	import { Check, ChevronsUpDownIcon } from 'lucide-svelte';
	import { closeAndRefocusTrigger } from '$utils';
	import { cn } from '$lib/utils';
	import PhoneNumber from '$lib/components/reusable/inputs/PhoneNumber.svelte';
	export let data: PageData;
	$: ({ profile } = data);
	$: loadingProfile = data?.loadingProfile ?? true;
	let updatingProfile = false;
	let phoneNumber: E164Number | null;
	$: phoneNumber = null;

	// Selected country
	let country: CountryCode | null = 'US';

	// Validity
	let valid: boolean;

	// Phone number details
	let detailedValue: DetailedValue | null = null;

	let options: TelInputOptions = {
		invalidateOnCountryChange: true,
		format: 'international'
	};
	let showCountryCodePopOver = false;
	const updateProfile = async () => {
		try {
			updatingProfile = true;
			await client.users.update(profile);
			await Session.attemptRefreshingSession();
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
					<!-- <Input bind:value={profile.phoneNumber} placeholder="(702)-234-5566" /> -->
					<PhoneNumber bind:phoneNumber={profile.phoneNumber} />
				</div>
			</div>
			<Button disabled={updatingProfile} on:click={updateProfile}>
				<ButtonLoadingSpinner bind:state={updatingProfile} />
				Update profile
			</Button>
		</div>
	{/if}
</div>
